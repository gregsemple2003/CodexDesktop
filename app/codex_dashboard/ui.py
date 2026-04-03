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


INTERVAL_TITLES = {
    "1m": "1 Minute",
    "5m": "5 Minutes",
    "15m": "15 Minutes",
    "1h": "1 Hour",
    "1d": "1 Day",
}


def format_tick_label(start_at: datetime, interval_key: str) -> str:
    if interval_key == "1d":
        return start_at.strftime("%m-%d")

    hour = start_at.strftime("%I").lstrip("0") or "12"
    meridiem = start_at.strftime("%p")
    if start_at.minute == 0:
        return f"{hour}{meridiem}"
    return f"{hour}:{start_at.minute:02d}{meridiem}"


def format_token_value(value: int) -> str:
    absolute_value = abs(value)
    for divisor, suffix in (
        (1_000_000_000, "B"),
        (1_000_000, "M"),
        (1_000, "K"),
    ):
        if absolute_value >= divisor:
            scaled = value / divisor
            text = f"{scaled:.1f}".rstrip("0").rstrip(".")
            return f"{text}{suffix}"
    return str(value)


def format_chart_title(interval_key: str) -> str:
    return f"Token Velocity per {INTERVAL_TITLES.get(interval_key, interval_key)}"


def interval_redline_tokens(weekly_budget_tokens: int, interval_seconds: int) -> int:
    return max(1, int(weekly_budget_tokens * interval_seconds / (7 * 24 * 60 * 60)))


class DashboardApp:
    def __init__(
        self,
        config_path: Path | None = None,
        smoke_artifact_dir: Path | None = None,
    ) -> None:
        self.config_path = config_path or default_config_path()
        self.config = load_config(self.config_path)
        self.config.startup_enabled = is_startup_enabled()
        self.selected_interval = "15m"
        self.ingest_queue: queue.Queue[tuple[str, object]] = queue.Queue()
        self.ingest_in_flight = False
        self.last_ingest_error: str | None = None
        self.hotkey_registered = False
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
        self.overlay.geometry("980x610+940+100")
        self.overlay.configure(bg="#0a0e14")
        self.overlay.bind("<Escape>", lambda _event: self.hide_overlay())

        self._configure_style()
        self._build_overlay()

        self.hotkey = GlobalHotkey(self.config.hotkey, self.toggle_overlay)
        try:
            self.hotkey.register()
            self.hotkey_registered = True
        except OSError:
            if self.smoke_artifact_dir is None:
                raise

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
        style.configure("Overlay.TFrame", background="#0a0e14")
        style.configure("Shell.TFrame", background="#1c2026")
        style.configure("Header.TFrame", background="#181c22")
        style.configure("BodyPanel.TFrame", background="#1c2026")
        style.configure("Card.TFrame", background="#181c22")
        style.configure("Footer.TFrame", background="#262a31")
        style.configure(
            "Brand.TLabel",
            background="#181c22",
            foreground="#bff4ff",
            font=("Bahnschrift SemiBold", 16),
        )
        style.configure(
            "Badge.TLabel",
            background="#31353c",
            foreground="#8fa8bb",
            font=("Segoe UI Semibold", 8),
        )
        style.configure(
            "Status.TLabel",
            background="#1c2026",
            foreground="#8fa8bb",
            font=("Segoe UI", 9),
        )
        style.configure(
            "MetricTitle.TLabel",
            background="#181c22",
            foreground="#8fa8bb",
            font=("Segoe UI Semibold", 8),
        )
        style.configure(
            "MetricValue.TLabel",
            background="#181c22",
            foreground="#dfe2eb",
            font=("Bahnschrift SemiBold", 20),
        )
        style.configure(
            "MetricUnit.TLabel",
            background="#181c22",
            foreground="#8fa8bb",
            font=("Segoe UI", 9),
        )
        style.configure(
            "MetricDetail.TLabel",
            background="#181c22",
            foreground="#8fa8bb",
            font=("Segoe UI", 9),
        )
        style.configure(
            "ChartTitle.TLabel",
            background="#1c2026",
            foreground="#dfe2eb",
            font=("Segoe UI Semibold", 10),
        )
        style.configure(
            "Tiny.TLabel",
            background="#1c2026",
            foreground="#6e8598",
            font=("Consolas", 8),
        )
        style.configure(
            "FooterLabel.TLabel",
            background="#262a31",
            foreground="#8fa8bb",
            font=("Segoe UI", 9),
        )
        style.configure(
            "FooterStatus.TLabel",
            background="#262a31",
            foreground="#8fa8bb",
            font=("Segoe UI Semibold", 8),
        )
        style.configure(
            "Accent.TButton",
            background="#16d9f5",
            foreground="#10141a",
            font=("Segoe UI Semibold", 9),
            borderwidth=0,
            focusthickness=0,
        )
        style.map(
            "Accent.TButton",
            background=[("active", "#2ee8ff")],
        )
        style.configure(
            "Quiet.TButton",
            background="#303743",
            foreground="#dfe2eb",
            font=("Segoe UI Semibold", 9),
            borderwidth=0,
            focusthickness=0,
        )
        style.map(
            "Quiet.TButton",
            background=[("active", "#3b4450")],
        )
        style.configure(
            "HeaderQuiet.TButton",
            background="#303743",
            foreground="#dfe2eb",
            font=("Segoe UI Semibold", 8),
            borderwidth=0,
            focusthickness=0,
        )
        style.map(
            "HeaderQuiet.TButton",
            background=[("active", "#3b4450")],
        )
        style.configure(
            "HeaderAccent.TButton",
            background="#16d9f5",
            foreground="#10141a",
            font=("Segoe UI Semibold", 8),
            borderwidth=0,
            focusthickness=0,
        )
        style.map(
            "HeaderAccent.TButton",
            background=[("active", "#2ee8ff")],
        )
        style.configure(
            "Overlay.TCheckbutton",
            background="#262a31",
            foreground="#dfe2eb",
            font=("Segoe UI", 9),
        )
        style.map(
            "Overlay.TCheckbutton",
            background=[("active", "#262a31")],
            foreground=[("active", "#dfe2eb")],
        )

    def _build_overlay(self) -> None:
        self.container = ttk.Frame(self.overlay, style="Overlay.TFrame", padding=28)
        self.container.pack(fill="both", expand=True)

        self.shell = ttk.Frame(self.container, style="Shell.TFrame")
        self.shell.pack(fill="both", expand=True)

        header = ttk.Frame(self.shell, style="Header.TFrame", padding=(16, 12))
        header.pack(fill="x")
        brand_row = ttk.Frame(header, style="Header.TFrame")
        brand_row.pack(side="left")
        ttk.Label(brand_row, text="CODEXDASHBOARD", style="Brand.TLabel").pack(side="left")
        ttk.Label(brand_row, text="VELOCITY.V2", style="Badge.TLabel").pack(side="left", padx=(8, 0))

        header_controls = ttk.Frame(header, style="Header.TFrame")
        header_controls.pack(side="right")
        self.interval_buttons: dict[str, ttk.Button] = {}
        for interval_key in ("1m", "5m", "15m", "1h", "1d"):
            button = ttk.Button(
                header_controls,
                text=interval_key,
                style="HeaderQuiet.TButton",
                command=lambda key=interval_key: self.select_interval(key),
                width=4,
            )
            button.pack(side="left", padx=(0, 6))
            self.interval_buttons[interval_key] = button
        ttk.Button(
            header_controls,
            text="X",
            style="HeaderQuiet.TButton",
            command=self.hide_overlay,
            width=3,
        ).pack(side="left", padx=(10, 0))

        tk.Frame(self.shell, bg="#39424d", height=1).pack(fill="x")

        body = ttk.Frame(self.shell, style="BodyPanel.TFrame", padding=(16, 14))
        body.pack(fill="both", expand=True)

        status_row = ttk.Frame(body, style="BodyPanel.TFrame")
        status_row.pack(fill="x", pady=(0, 12))
        self.status_label = ttk.Label(
            status_row,
            text="Waiting for first ingest...",
            style="Status.TLabel",
        )
        self.status_label.pack(side="left")
        self.hotkey_label = ttk.Label(
            status_row,
            text=f"Toggle: {self.config.hotkey}",
            style="Tiny.TLabel",
        )
        self.hotkey_label.pack(side="right")

        metrics_row = ttk.Frame(body, style="BodyPanel.TFrame")
        metrics_row.pack(fill="x", pady=(0, 14))
        for column in range(4):
            metrics_row.columnconfigure(column, weight=1)

        card_7d = ttk.Frame(metrics_row, style="Card.TFrame", padding=(12, 10))
        card_7d.grid(row=0, column=0, sticky="nsew", padx=(0, 10))
        ttk.Label(card_7d, text="7D TOTAL TOKENS", style="MetricTitle.TLabel").pack(anchor="w")
        value_row = ttk.Frame(card_7d, style="Card.TFrame")
        value_row.pack(anchor="w", pady=(10, 4))
        self.local_total_value = ttk.Label(value_row, text="0", style="MetricValue.TLabel")
        self.local_total_value.pack(side="left")
        ttk.Label(value_row, text="TKN", style="MetricUnit.TLabel").pack(side="left", padx=(6, 0), pady=(9, 0))
        self.local_total_detail = ttk.Label(card_7d, text="", style="MetricDetail.TLabel")
        self.local_total_detail.pack(anchor="w")

        card_projected = ttk.Frame(metrics_row, style="Card.TFrame", padding=(12, 10))
        card_projected.grid(row=0, column=1, sticky="nsew", padx=(0, 10))
        ttk.Label(card_projected, text="PROJECTED WEEKLY BURN", style="MetricTitle.TLabel").pack(anchor="w")
        projected_row = ttk.Frame(card_projected, style="Card.TFrame")
        projected_row.pack(anchor="w", pady=(10, 4))
        self.projected_value = ttk.Label(projected_row, text="0", style="MetricValue.TLabel")
        self.projected_value.pack(side="left")
        ttk.Label(projected_row, text="TKN", style="MetricUnit.TLabel").pack(side="left", padx=(6, 0), pady=(9, 0))
        self.projected_detail = ttk.Label(card_projected, text="", style="MetricDetail.TLabel")
        self.projected_detail.pack(anchor="w")

        card_current = ttk.Frame(metrics_row, style="Card.TFrame", padding=(12, 10))
        card_current.grid(row=0, column=2, sticky="nsew", padx=(0, 10))
        self.current_card_title = ttk.Label(card_current, text="CURRENT BUCKET", style="MetricTitle.TLabel")
        self.current_card_title.pack(anchor="w")
        current_row = ttk.Frame(card_current, style="Card.TFrame")
        current_row.pack(anchor="w", pady=(10, 4))
        self.current_bucket_value = ttk.Label(current_row, text="0", style="MetricValue.TLabel")
        self.current_bucket_value.pack(side="left")
        ttk.Label(current_row, text="TKN", style="MetricUnit.TLabel").pack(side="left", padx=(6, 0), pady=(9, 0))
        self.current_bucket_detail = ttk.Label(card_current, text="", style="MetricDetail.TLabel")
        self.current_bucket_detail.pack(anchor="w")

        self.status_card = ttk.Frame(metrics_row, style="Card.TFrame", padding=(0, 10))
        self.status_card.grid(row=0, column=3, sticky="nsew")
        self.status_card.columnconfigure(1, weight=1)
        self.status_accent = tk.Frame(self.status_card, bg="#16d9f5", width=3)
        self.status_accent.grid(row=0, column=0, rowspan=3, sticky="ns", padx=(0, 10))
        ttk.Label(self.status_card, text="STATUS", style="MetricTitle.TLabel").grid(row=0, column=1, sticky="w", padx=(2, 0))
        self.status_metric_value = ttk.Label(self.status_card, text="Awaiting data", style="MetricValue.TLabel")
        self.status_metric_value.grid(row=1, column=1, sticky="w", padx=(2, 0), pady=(8, 0))
        self.status_metric_detail = ttk.Label(self.status_card, text="No ingest cycle completed yet.", style="MetricDetail.TLabel")
        self.status_metric_detail.grid(row=2, column=1, sticky="w", padx=(2, 0), pady=(2, 0))

        chart_header = ttk.Frame(body, style="BodyPanel.TFrame")
        chart_header.pack(fill="x", pady=(0, 8))
        self.chart_header_title = ttk.Label(chart_header, text=format_chart_title(self.selected_interval), style="ChartTitle.TLabel")
        self.chart_header_title.pack(side="left")
        self.chart_header_context = ttk.Label(chart_header, text=self._timezone_label(), style="Tiny.TLabel")
        self.chart_header_context.pack(side="right")

        chart_shell = ttk.Frame(body, style="Shell.TFrame", padding=(10, 10))
        chart_shell.pack(fill="x")
        self.canvas = tk.Canvas(
            chart_shell,
            width=880,
            height=230,
            bg="#10141a",
            highlightthickness=0,
        )
        self.canvas.pack(fill="x")

        info_row = ttk.Frame(body, style="BodyPanel.TFrame")
        info_row.pack(fill="x", pady=(12, 0))
        self.advisory_label = ttk.Label(
            info_row,
            text="No weekly advisory yet.",
            style="Status.TLabel",
            wraplength=520,
            justify="left",
        )
        self.advisory_label.pack(side="left", fill="x", expand=True)

        footer = ttk.Frame(self.shell, style="Footer.TFrame", padding=(16, 12))
        footer.pack(fill="x", side="bottom")
        tk.Frame(footer, bg="#39424d", height=1).pack(fill="x", side="top")

        footer_inner = ttk.Frame(footer, style="Footer.TFrame")
        footer_inner.pack(fill="x", pady=(8, 0))
        self.live_status_label = ttk.Label(footer_inner, text="LIVE SYNC: ACTIVE", style="FooterStatus.TLabel")
        self.live_status_label.pack(side="left")

        controls_row = ttk.Frame(footer_inner, style="Footer.TFrame")
        controls_row.pack(side="left", padx=(28, 0))
        ttk.Label(controls_row, text="Weekly Budget", style="FooterLabel.TLabel").pack(side="left")
        self.weekly_budget_var = tk.StringVar(
            value=str(self.config.weekly_budget_tokens)
        )
        self.weekly_budget_entry = tk.Entry(
            controls_row,
            textvariable=self.weekly_budget_var,
            width=12,
            bg="#121820",
            fg="#dfe2eb",
            insertbackground="#dfe2eb",
            relief="flat",
            highlightthickness=1,
            highlightbackground="#2b323b",
            highlightcolor="#16d9f5",
            font=("Segoe UI", 10),
        )
        self.weekly_budget_entry.pack(side="left", padx=(10, 8), ipady=4)
        ttk.Button(
            controls_row,
            text="Save Budget",
            style="Accent.TButton",
            command=self.save_budget,
        ).pack(side="left", padx=(0, 10))

        self.startup_var = tk.BooleanVar(value=self.config.startup_enabled)
        startup_toggle = ttk.Checkbutton(
            controls_row,
            text="Start with Windows",
            variable=self.startup_var,
            style="Overlay.TCheckbutton",
            command=self.toggle_startup,
        )
        startup_toggle.pack(side="left")

        action_row = ttk.Frame(footer_inner, style="Footer.TFrame")
        action_row.pack(side="right")
        ttk.Button(action_row, text="Refresh Now", style="Quiet.TButton", command=self.schedule_ingest).pack(side="left", padx=(0, 8))
        ttk.Button(action_row, text="Dismiss", style="Quiet.TButton", command=self.hide_overlay).pack(side="left", padx=(0, 8))
        ttk.Button(action_row, text="Quit App", style="Accent.TButton", command=self.quit).pack(side="left")

        self._refresh_interval_buttons()

    def _poll_hotkey(self) -> None:
        if self.hotkey_registered:
            self.hotkey.poll()
        self.root.after(50, self._poll_hotkey)

    def _poll_ingest_results(self) -> None:
        try:
            while True:
                event_type, payload = self.ingest_queue.get_nowait()
                if event_type == "summary":
                    self.ingest_in_flight = False
                    self.last_ingest_error = None
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
                    self.last_ingest_error = str(payload)
                    self.status_label.configure(text=f"Ingest error: {payload}")
                    self._refresh_status_surfaces(False)
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
            bucket_count=20,
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

        self.local_total_value.configure(text=format_token_value(total_7d))
        self.local_total_detail.configure(text=f"{total_7d:,} tokens in the last 7 days")
        self.projected_value.configure(text=format_token_value(projected))
        self.projected_detail.configure(
            text=(
                f"Over {format_token_value(self.config.weekly_budget_tokens)} budget"
                if redline
                else f"Redline {format_token_value(self.config.weekly_budget_tokens)}"
            )
        )
        self.current_card_title.configure(
            text=f"CURRENT {INTERVAL_TITLES[self.selected_interval].upper()}"
        )
        self.current_bucket_value.configure(text=format_token_value(current_bucket_tokens))
        self.current_bucket_detail.configure(
            text=(
                "Bucket threshold "
                f"{format_token_value(interval_redline_tokens(self.config.weekly_budget_tokens, interval_seconds))}"
            )
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
        self.chart_header_title.configure(text=format_chart_title(self.selected_interval))
        self.chart_header_context.configure(text=self._timezone_label())
        self._refresh_status_surfaces(redline)
        self.draw_chart(buckets)

    def _refresh_status_surfaces(self, redline: bool) -> None:
        if self.last_ingest_error is not None:
            accent = "#ff5a52"
            value = "Attention"
            detail = "Ingest error detected."
            live = "SYNC BLOCKED"
        elif redline:
            accent = "#ff5a52"
            value = "Redline"
            detail = "Projected weekly burn exceeds budget."
            live = "LIVE SYNC: ACTIVE"
        else:
            accent = "#bff4ff"
            value = "Operational"
            detail = "Within weekly budget."
            live = "LIVE SYNC: ACTIVE"

        self.status_accent.configure(bg=accent)
        self.status_metric_value.configure(text=value, foreground=accent)
        self.status_metric_detail.configure(text=detail)
        self.live_status_label.configure(text=live)

    def draw_chart(self, buckets) -> None:
        self.canvas.delete("all")
        width = int(self.canvas["width"])
        height = int(self.canvas["height"])
        left = 56
        right = width - 24
        top = 18
        bottom = height - 28
        chart_height = bottom - top
        chart_width = right - left

        self.canvas.create_rectangle(
            left,
            top,
            right,
            bottom,
            outline="#39424d",
            fill="#10141a",
        )

        if not buckets:
            self.canvas.create_text(
                width / 2,
                height / 2,
                text="No token data yet.",
                fill="#dfe2eb",
                font=("Segoe UI Semibold", 14),
            )
            return

        threshold_tokens = interval_redline_tokens(
            self.config.weekly_budget_tokens,
            INTERVAL_SECONDS[self.selected_interval],
        )
        max_tokens = max(max(bucket.total_tokens for bucket in buckets), threshold_tokens, 1)
        grid_steps = 4
        for row in range(grid_steps + 1):
            y = top + row * chart_height / grid_steps
            self.canvas.create_line(left, y, right, y, fill="#31353c")
            label_value = int(round(max_tokens * (grid_steps - row) / grid_steps))
            self.canvas.create_text(
                left - 8,
                y,
                anchor="e",
                text=format_token_value(label_value),
                fill="#6e8598",
                font=("Consolas", 8),
            )

        threshold_y = bottom - ((threshold_tokens / max_tokens) * chart_height)
        self.canvas.create_line(left, threshold_y, right, threshold_y, fill="#ff5a52")
        self.canvas.create_rectangle(left + 2, threshold_y - 9, left + 68, threshold_y + 7, fill="#ff5a52", outline="")
        self.canvas.create_text(
            left + 35,
            threshold_y - 1,
            text="REDLINE",
            fill="#dfe2eb",
            font=("Segoe UI Semibold", 7),
        )

        gap = 8
        bar_width = max(12, int((chart_width - gap * (len(buckets) - 1)) / len(buckets)))
        for index, bucket in enumerate(buckets):
            x0 = left + index * (bar_width + gap)
            x1 = x0 + bar_width
            bar_height = (bucket.total_tokens / max_tokens) * (chart_height - 8)
            y0 = bottom - bar_height
            if bucket.total_tokens >= threshold_tokens and bucket.total_tokens > 0:
                fill = "#ff5a52"
            elif index == len(buckets) - 1:
                fill = "#16d9f5"
            elif bucket.total_tokens == 0:
                fill = "#0f5c69"
            elif index % 2 == 0:
                fill = "#20b5cc"
            else:
                fill = "#138ea2"
            self.canvas.create_rectangle(x0, y0, x1, bottom, fill=fill, outline="")
            if index % 3 == 0 or index == len(buckets) - 1:
                label = format_tick_label(bucket.start_at, self.selected_interval)
                self.canvas.create_text(
                    (x0 + x1) / 2,
                    bottom + 14,
                    text=label,
                    fill="#6e8598",
                    font=("Consolas", 8),
                )

    def select_interval(self, interval_key: str) -> None:
        self.selected_interval = interval_key
        self._refresh_interval_buttons()
        self.refresh_data()

    def _refresh_interval_buttons(self) -> None:
        for key, button in self.interval_buttons.items():
            button.configure(
                style="HeaderAccent.TButton" if key == self.selected_interval else "HeaderQuiet.TButton"
            )

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
        if self.hotkey_registered:
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
                f"interval={self.selected_interval}",
                f"weekly_budget={self.config.weekly_budget_tokens}",
                f"startup_enabled={self.startup_var.get()}",
                f"7d_total={self.local_total_value.cget('text')}",
                f"projected={self.projected_value.cget('text')}",
                f"current_bucket={self.current_bucket_value.cget('text')}",
                f"status={self.status_metric_value.cget('text')}",
                self.advisory_label.cget("text"),
                f"hotkey_triggered={self.smoke_hotkey_triggered}",
                f"overlay_fallback={self.smoke_overlay_fallback}",
            ]
        )
        (artifact_dir / "overlay-summary.txt").write_text(summary, encoding="utf-8")
        os._exit(0)

    def _trigger_smoke_hotkey(self) -> None:
        if not self.hotkey_registered:
            self.toggle_overlay()
            return
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
