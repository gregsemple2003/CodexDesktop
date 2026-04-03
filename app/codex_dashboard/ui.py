from __future__ import annotations

import ctypes
import os
import queue
import threading
import tkinter as tk
from datetime import UTC, datetime, timedelta, tzinfo
from pathlib import Path
from tkinter import ttk

from .aggregation import INTERVAL_SECONDS, build_buckets, is_over_redline, project_weekly_burn
from .config import DashboardConfig, load_config, save_config
from .hotkey import GlobalHotkey
from .paths import default_config_path
from .scanner import ingest_once
from .startup import is_startup_enabled, set_startup_enabled
from .storage import connect, initialize_db, load_events_since


def format_tick_label(start_at: datetime, interval_key: str) -> str:
    if interval_key == "1d":
        return start_at.strftime("%m-%d")

    hour = start_at.strftime("%I").lstrip("0") or "12"
    meridiem = start_at.strftime("%p")
    if start_at.minute == 0:
        return f"{hour}{meridiem}"
    return f"{hour}:{start_at.minute:02d}{meridiem}"


class DashboardApp:
    def __init__(
        self,
        config_path: Path | None = None,
        smoke_artifact_dir: Path | None = None,
    ) -> None:
        self.config_path = config_path or default_config_path()
        self.config = load_config(self.config_path)
        self.config.startup_enabled = is_startup_enabled()
        self.selected_interval = "1h"
        self.ingest_queue: queue.Queue[tuple[str, object]] = queue.Queue()
        self.ingest_in_flight = False
        self.smoke_artifact_dir = smoke_artifact_dir
        self._quitting = False
        self.smoke_hotkey_triggered = False
        self.smoke_overlay_fallback = False
        self.display_timezone = self._resolve_display_timezone()

        self.root = tk.Tk()
        self.root.withdraw()
        self.root.title("CodexDashboard")
        self.root.protocol("WM_DELETE_WINDOW", self.quit)

        self.overlay = tk.Toplevel(self.root)
        self.overlay.withdraw()
        self.overlay.overrideredirect(True)
        self.overlay.attributes("-topmost", True)
        self.overlay.geometry("860x560+980+110")
        self.overlay.configure(bg="#0b1524")
        self.overlay.bind("<Escape>", lambda _event: self.hide_overlay())

        self._configure_style()
        self._build_overlay()

        self.hotkey = GlobalHotkey(self.config.hotkey, self.toggle_overlay)
        self.hotkey.register()

        self.root.after(50, self._poll_hotkey)
        self.root.after(100, self._poll_ingest_results)
        self.root.after(100, self.refresh_data)
        self.root.after(250, self.schedule_ingest)
        if self.smoke_artifact_dir is not None:
            self.root.after(350, self._trigger_smoke_hotkey)
            self.root.after(1200, self._run_smoke_capture)

    def _configure_style(self) -> None:
        style = ttk.Style()
        style.theme_use("clam")
        style.configure("Overlay.TFrame", background="#0b1524")
        style.configure(
            "Title.TLabel",
            background="#0b1524",
            foreground="#f5f7fb",
            font=("Segoe UI Semibold", 20),
        )
        style.configure(
            "Body.TLabel",
            background="#0b1524",
            foreground="#bfd3ea",
            font=("Segoe UI", 10),
        )
        style.configure(
            "Value.TLabel",
            background="#0b1524",
            foreground="#f5f7fb",
            font=("Segoe UI Semibold", 11),
        )
        style.configure(
            "Accent.TButton",
            background="#0d8aa8",
            foreground="#ffffff",
            font=("Segoe UI Semibold", 10),
            borderwidth=0,
            focusthickness=0,
        )
        style.map(
            "Accent.TButton",
            background=[("active", "#12a4c7")],
        )
        style.configure(
            "Quiet.TButton",
            background="#18283a",
            foreground="#dce7f3",
            font=("Segoe UI", 10),
            borderwidth=0,
            focusthickness=0,
        )
        style.map(
            "Quiet.TButton",
            background=[("active", "#24374e")],
        )
        style.configure(
            "Overlay.TCheckbutton",
            background="#0b1524",
            foreground="#dce7f3",
            font=("Segoe UI", 10),
        )
        style.map(
            "Overlay.TCheckbutton",
            background=[("active", "#0b1524")],
            foreground=[("active", "#f5f7fb")],
        )

    def _build_overlay(self) -> None:
        self.container = ttk.Frame(self.overlay, style="Overlay.TFrame", padding=18)
        self.container.pack(fill="both", expand=True)

        header = ttk.Frame(self.container, style="Overlay.TFrame")
        header.pack(fill="x")
        ttk.Label(header, text="CodexDashboard", style="Title.TLabel").pack(side="left")
        ttk.Button(
            header,
            text="Hide",
            style="Quiet.TButton",
            command=self.hide_overlay,
        ).pack(side="right", padx=(8, 0))
        ttk.Button(
            header,
            text="Quit",
            style="Quiet.TButton",
            command=self.quit,
        ).pack(side="right")

        subtitle = ttk.Label(
            self.container,
            text="Live token velocity from local Codex session telemetry.",
            style="Body.TLabel",
        )
        subtitle.pack(anchor="w", pady=(6, 12))

        top_row = ttk.Frame(self.container, style="Overlay.TFrame")
        top_row.pack(fill="x", pady=(0, 10))

        self.status_label = ttk.Label(
            top_row,
            text="Waiting for first ingest...",
            style="Body.TLabel",
        )
        self.status_label.pack(side="left")
        ttk.Button(
            top_row,
            text="Refresh now",
            style="Accent.TButton",
            command=self.schedule_ingest,
        ).pack(side="right")

        interval_row = ttk.Frame(self.container, style="Overlay.TFrame")
        interval_row.pack(fill="x", pady=(0, 12))
        self.interval_buttons: dict[str, ttk.Button] = {}
        for interval_key in ("1m", "5m", "15m", "1h", "1d"):
            button = ttk.Button(
                interval_row,
                text=interval_key,
                style="Quiet.TButton",
                command=lambda key=interval_key: self.select_interval(key),
            )
            button.pack(side="left", padx=(0, 8))
            self.interval_buttons[interval_key] = button

        self.canvas = tk.Canvas(
            self.container,
            width=804,
            height=280,
            bg="#132338",
            highlightthickness=0,
        )
        self.canvas.pack(fill="x")

        metrics_row = ttk.Frame(self.container, style="Overlay.TFrame")
        metrics_row.pack(fill="x", pady=(14, 10))

        self.local_total_value = ttk.Label(metrics_row, text="0", style="Value.TLabel")
        self.local_total_value.pack(side="left")
        ttk.Label(
            metrics_row,
            text=" local tokens in last 7d",
            style="Body.TLabel",
        ).pack(side="left", padx=(4, 18))

        self.redline_value = ttk.Label(metrics_row, text="Projected burn: 0", style="Value.TLabel")
        self.redline_value.pack(side="left")

        budget_row = ttk.Frame(self.container, style="Overlay.TFrame")
        budget_row.pack(fill="x", pady=(0, 10))

        ttk.Label(budget_row, text="Weekly budget", style="Body.TLabel").pack(side="left")
        self.weekly_budget_var = tk.StringVar(
            value=str(self.config.weekly_budget_tokens)
        )
        self.weekly_budget_entry = tk.Entry(
            budget_row,
            textvariable=self.weekly_budget_var,
            width=14,
            bg="#132338",
            fg="#f5f7fb",
            insertbackground="#f5f7fb",
            relief="flat",
            font=("Segoe UI", 10),
        )
        self.weekly_budget_entry.pack(side="left", padx=(10, 8))
        ttk.Button(
            budget_row,
            text="Save",
            style="Accent.TButton",
            command=self.save_budget,
        ).pack(side="left")

        self.startup_var = tk.BooleanVar(value=self.config.startup_enabled)
        startup_toggle = ttk.Checkbutton(
            budget_row,
            text="Start with Windows",
            variable=self.startup_var,
            style="Overlay.TCheckbutton",
            command=self.toggle_startup,
        )
        startup_toggle.pack(side="right")

        self.advisory_label = ttk.Label(
            self.container,
            text="No weekly advisory yet.",
            style="Body.TLabel",
        )
        self.advisory_label.pack(anchor="w")

        self._refresh_interval_buttons()

    def _poll_hotkey(self) -> None:
        self.hotkey.poll()
        self.root.after(50, self._poll_hotkey)

    def _poll_ingest_results(self) -> None:
        try:
            while True:
                event_type, payload = self.ingest_queue.get_nowait()
                if event_type == "summary":
                    self.ingest_in_flight = False
                    files_scanned, files_updated, events_ingested = payload
                    self.status_label.configure(
                        text=(
                            f"Last ingest {datetime.now().strftime('%H:%M:%S')} | "
                            f"files {files_updated}/{files_scanned} | "
                            f"events +{events_ingested}"
                        )
                    )
                    self.refresh_data()
                elif event_type == "error":
                    self.ingest_in_flight = False
                    self.status_label.configure(text=f"Ingest error: {payload}")
        except queue.Empty:
            pass
        self.root.after(100, self._poll_ingest_results)

    def schedule_ingest(self) -> None:
        if self.ingest_in_flight:
            return
        self.ingest_in_flight = True

        def worker() -> None:
            try:
                connection = connect(Path(self.config.db_path))
                initialize_db(connection)
                summary = ingest_once(connection, self.config)
                connection.close()
                self.ingest_queue.put(
                    (
                        "summary",
                        (
                            summary.files_scanned,
                            summary.files_updated,
                            summary.events_ingested,
                        ),
                    )
                )
            except Exception as exc:  # pragma: no cover - GUI error surfacing
                self.ingest_queue.put(("error", str(exc)))

        thread = threading.Thread(target=worker, daemon=True)
        thread.start()
        self.root.after(self.config.polling_seconds * 1000, self.schedule_ingest)

    def refresh_data(self) -> None:
        connection = connect(Path(self.config.db_path))
        initialize_db(connection)
        now = datetime.now(self.display_timezone)
        events = load_events_since(connection, now.astimezone(UTC) - timedelta(days=7))
        connection.close()

        buckets = build_buckets(
            events,
            self.selected_interval,
            bucket_count=18,
            now=now,
            display_tz=self.display_timezone,
        )
        interval_seconds = INTERVAL_SECONDS[self.selected_interval]
        current_bucket_tokens = buckets[-1].total_tokens if buckets else 0
        projected = project_weekly_burn(current_bucket_tokens, interval_seconds)
        redline = is_over_redline(
            current_bucket_tokens,
            interval_seconds,
            self.config.weekly_budget_tokens,
        )
        total_7d = sum(event.total_tokens for event in events)

        self.local_total_value.configure(text=f"{total_7d:,}")
        self.redline_value.configure(
            text=f"Projected burn: {projected:,} {'REDLINE' if redline else 'within budget'}"
        )
        self.redline_value.configure(
            foreground="#ff7a59" if redline else "#6de0b0"
        )
        latest_advisory = next(
            (event for event in reversed(events) if event.weekly_used_percent is not None),
            None,
        )
        if latest_advisory is None:
            self.advisory_label.configure(text="No weekly advisory yet.")
        else:
            advisory = latest_advisory.weekly_used_percent
            self.advisory_label.configure(
                text=(
                    f"Codex advisory window: {advisory:.1f}% used | "
                    f"reset epoch {latest_advisory.weekly_resets_at}"
                )
            )
        self.draw_chart(buckets)

    def draw_chart(self, buckets) -> None:
        self.canvas.delete("all")
        width = int(self.canvas["width"])
        height = int(self.canvas["height"])
        left = 28
        right = width - 20
        top = 20
        bottom = height - 34
        chart_height = bottom - top
        chart_width = right - left

        self.canvas.create_rectangle(
            left,
            top,
            right,
            bottom,
            outline="#203651",
            fill="#132338",
        )
        for row in range(1, 5):
            y = top + row * chart_height / 5
            self.canvas.create_line(left, y, right, y, fill="#1d3148")

        if not buckets:
            self.canvas.create_text(
                width / 2,
                height / 2,
                text="No token data yet.",
                fill="#dce7f3",
                font=("Segoe UI", 14),
            )
            return

        max_tokens = max(bucket.total_tokens for bucket in buckets) or 1
        gap = 6
        bar_width = max(10, int((chart_width - gap * (len(buckets) - 1)) / len(buckets)))
        for index, bucket in enumerate(buckets):
            x0 = left + index * (bar_width + gap)
            x1 = x0 + bar_width
            bar_height = (bucket.total_tokens / max_tokens) * (chart_height - 8)
            y0 = bottom - bar_height
            fill = "#ff7a59" if index == len(buckets) - 1 and bar_height > 0 else "#22b8cf"
            self.canvas.create_rectangle(x0, y0, x1, bottom, fill=fill, outline="")
            if index % 3 == 0 or index == len(buckets) - 1:
                label = format_tick_label(bucket.start_at, self.selected_interval)
                self.canvas.create_text(
                    (x0 + x1) / 2,
                    bottom + 16,
                    text=label,
                    fill="#8ba6c2",
                    font=("Segoe UI", 8),
                )
        self.canvas.create_text(
            left,
            top - 8,
            anchor="w",
            text=f"Interval {self.selected_interval} | {self._timezone_label()}",
            fill="#dce7f3",
            font=("Segoe UI Semibold", 10),
        )

    def select_interval(self, interval_key: str) -> None:
        self.selected_interval = interval_key
        self._refresh_interval_buttons()
        self.refresh_data()

    def _refresh_interval_buttons(self) -> None:
        for key, button in self.interval_buttons.items():
            button.configure(style="Accent.TButton" if key == self.selected_interval else "Quiet.TButton")

    def save_budget(self) -> None:
        raw_value = self.weekly_budget_var.get().replace(",", "").strip()
        try:
            budget = int(raw_value)
        except ValueError:
            self.status_label.configure(text="Weekly budget must be an integer token count.")
            return
        self.config.weekly_budget_tokens = max(1, budget)
        save_config(self.config, self.config_path)
        self.status_label.configure(text=f"Saved weekly budget: {self.config.weekly_budget_tokens:,}")
        self.refresh_data()

    def toggle_startup(self) -> None:
        enabled = bool(self.startup_var.get())
        set_startup_enabled(enabled)
        self.config.startup_enabled = enabled
        save_config(self.config, self.config_path)
        state = "enabled" if enabled else "disabled"
        self.status_label.configure(text=f"Startup {state}.")

    def toggle_overlay(self) -> None:
        if self.smoke_artifact_dir is not None:
            self.smoke_hotkey_triggered = True
        if self.overlay.state() == "withdrawn":
            self.show_overlay()
        else:
            self.hide_overlay()

    def show_overlay(self) -> None:
        self.refresh_data()
        self.overlay.deiconify()
        self.overlay.lift()
        self.overlay.focus_force()

    def hide_overlay(self) -> None:
        self.overlay.withdraw()

    def quit(self) -> None:
        if self._quitting:
            return
        self._quitting = True
        self.hotkey.unregister()
        self.root.quit()
        self.overlay.destroy()
        self.root.destroy()

    def _run_smoke_capture(self) -> None:
        artifact_dir = self.smoke_artifact_dir
        if artifact_dir is None:
            return
        artifact_dir.mkdir(parents=True, exist_ok=True)
        if self.overlay.state() == "withdrawn":
            self.smoke_overlay_fallback = True
            self.show_overlay()
        self.overlay.update_idletasks()
        self.canvas.postscript(
            file=str(artifact_dir / "overlay-chart.ps"),
            colormode="color",
        )
        summary = "\n".join(
            [
                self.status_label.cget("text"),
                self.local_total_value.cget("text"),
                self.redline_value.cget("text"),
                self.advisory_label.cget("text"),
                f"hotkey_triggered={self.smoke_hotkey_triggered}",
                f"overlay_fallback={self.smoke_overlay_fallback}",
            ]
        )
        (artifact_dir / "overlay-summary.txt").write_text(summary, encoding="utf-8")
        os._exit(0)

    def _trigger_smoke_hotkey(self) -> None:
        user32 = ctypes.windll.user32
        keybd_event = user32.keybd_event
        KEYEVENTF_KEYUP = 0x0002
        VK_CONTROL = 0x11
        VK_MENU = 0x12
        VK_SPACE = 0x20
        keybd_event(VK_CONTROL, 0, 0, 0)
        keybd_event(VK_MENU, 0, 0, 0)
        keybd_event(VK_SPACE, 0, 0, 0)
        keybd_event(VK_SPACE, 0, KEYEVENTF_KEYUP, 0)
        keybd_event(VK_MENU, 0, KEYEVENTF_KEYUP, 0)
        keybd_event(VK_CONTROL, 0, KEYEVENTF_KEYUP, 0)

    def run(self) -> None:
        self.root.mainloop()

    def _resolve_display_timezone(self) -> tzinfo:
        return datetime.now().astimezone().tzinfo or UTC

    def _timezone_label(self) -> str:
        return datetime.now(self.display_timezone).strftime("%Z") or "local"
