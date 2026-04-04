from __future__ import annotations

import ctypes
import json
import os
import queue
import shutil
import subprocess
import threading
import tkinter as tk
from datetime import UTC, datetime, timedelta, tzinfo
from pathlib import Path
from tkinter import ttk

from .aggregation import (
    INTERVAL_SECONDS,
    METRIC_MODES,
    build_buckets,
    build_project_stacks,
    is_over_redline,
    project_weekly_burn,
)
from .config import DashboardConfig, load_config, maybe_upgrade_weekly_budget, save_config
from .hotkey import GlobalHotkey
from .investigation import (
    build_bucket_investigation,
    build_codex_launch_command,
    report_path_for_brief,
    write_bucket_investigation,
)
from .jobs import (
    DESIRED_STATE_ENABLED,
    JOB_STATUS_BLOCKED,
    JOB_STATUS_DISABLED,
    JOB_STATUS_DRIFTED,
    JOB_STATUS_MISSING,
    apply_registry,
    default_jobs_registry_path,
    ensure_jobs_registry,
    reconcile_registry,
)
from .paths import default_config_path, default_investigations_path
from .scanner import ingest_once
from .storage import connect, initialize_db, load_events_since, load_session_context_markers


INTERVAL_TITLES = {
    "1m": "1 Minute",
    "5m": "5 Minutes",
    "15m": "15 Minutes",
    "1h": "1 Hour",
    "1d": "1 Day",
}
CHART_MODES = {
    "velocity": "Velocity",
    "repo": "Repo",
}
FONT_ASSET_DIR = Path(__file__).resolve().parent / "assets" / "fonts"
ROLLING_PROJECTION_BUCKETS = 4
REPO_STACK_COLORS = (
    "#5eb8ff",
    "#ff9b52",
    "#3fd49f",
    "#d785ff",
    "#ffd65c",
    "#6e7b8c",
)
REPO_STACK_OUTLINE = "#0d131b"
JOBS_STATUS_COLORS = {
    "in_sync": "#16d9f5",
    "drifted": "#ff8a52",
    "disabled": "#ff8a52",
    "missing": "#ff8a52",
    "blocked": "#ff5a52",
    "unknown": "#8fa8bb",
}
TAB_ACTIVE_FOREGROUND = "#c3f5ff"
TAB_INACTIVE_FOREGROUND = "#9fbdcc"
TAB_ACTIVE_UNDERLINE = "#00e5ff"
HEADER_BACKGROUND = "#181c22"


def load_private_font_assets() -> list[Path]:
    font_paths = [
        FONT_ASSET_DIR / "Inter[opsz,wght].ttf",
        FONT_ASSET_DIR / "SpaceGrotesk[wght].ttf",
    ]
    loaded_fonts: list[Path] = []
    add_font_resource = ctypes.windll.gdi32.AddFontResourceExW
    FR_PRIVATE = 0x10
    for font_path in font_paths:
        if font_path.exists() and add_font_resource(str(font_path), FR_PRIVATE, 0):
            loaded_fonts.append(font_path)
    return loaded_fonts


def unload_private_font_assets(font_paths: list[Path]) -> None:
    remove_font_resource = ctypes.windll.gdi32.RemoveFontResourceExW
    FR_PRIVATE = 0x10
    for font_path in font_paths:
        remove_font_resource(str(font_path), FR_PRIVATE, 0)


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


def format_signed_token_value(value: int) -> str:
    if value == 0:
        return "0"
    prefix = "+" if value > 0 else "-"
    return f"{prefix}{format_token_value(abs(value))}"


def format_reset_remaining(reset_epoch: int, now: datetime | None = None) -> str:
    current_time = now or datetime.now(UTC)
    remaining_seconds = max(0, reset_epoch - int(current_time.timestamp()))
    if remaining_seconds >= 24 * 60 * 60:
        return f"{remaining_seconds / (24 * 60 * 60):.1f}d"
    if remaining_seconds >= 60 * 60:
        return f"{remaining_seconds / (60 * 60):.1f}h"
    if remaining_seconds >= 60:
        return f"{remaining_seconds / 60:.0f}m"
    return f"{remaining_seconds}s"


def format_budget_billions(weekly_budget_tokens: int) -> str:
    return f"{weekly_budget_tokens / 1_000_000_000:.1f}".rstrip("0").rstrip(".")


def parse_budget_billions(raw_value: str) -> int:
    normalized = raw_value.lower().replace(",", "").strip()
    if not normalized:
        raise ValueError("budget is required")
    if normalized.endswith("b"):
        normalized = normalized[:-1].strip()
    budget_value = float(normalized)
    if budget_value <= 0:
        raise ValueError("budget must be positive")
    if "." in normalized or budget_value < 10_000:
        return int(round(budget_value * 1_000_000_000))
    return int(round(budget_value))


def rolling_average_tokens(buckets, sample_size: int) -> int:
    if not buckets or sample_size <= 0:
        return 0
    recent_buckets = buckets[-sample_size:]
    return int(round(sum(bucket.total_tokens for bucket in recent_buckets) / len(recent_buckets)))


def format_chart_title(
    interval_key: str,
    chart_mode: str = "velocity",
    metric_mode: str = "total",
) -> str:
    prefix = "Normalized " if metric_mode == "norm" else ""
    if chart_mode == "repo":
        return f"{prefix}Repo Share per {INTERVAL_TITLES.get(interval_key, interval_key)}"
    return f"{prefix}Token Velocity per {INTERVAL_TITLES.get(interval_key, interval_key)}"


def format_velocity_tooltip(total_tokens: int) -> str:
    return format_token_value(total_tokens)


def format_repo_tooltip(
    bucket_totals: dict[str, int],
    repo_legend: list[tuple[str, str]],
) -> str:
    nonzero_segments = [
        (label, bucket_totals.get(project_key, 0))
        for project_key, label in repo_legend
        if bucket_totals.get(project_key, 0) > 0
    ]
    if not nonzero_segments:
        return "0"
    nonzero_segments.sort(key=lambda item: (-item[1], item[0].lower()))
    return "\n".join(f"{label}: {format_token_value(total_tokens)}" for label, total_tokens in nonzero_segments)


def interval_redline_tokens(weekly_budget_tokens: int, interval_seconds: int) -> int:
    return max(1, int(weekly_budget_tokens * interval_seconds / (7 * 24 * 60 * 60)))


def jobs_needs_attention_count(summary: dict[str, int]) -> int:
    return sum(count for status, count in summary.items() if status != "in_sync")


def format_jobs_timestamp(raw_value: str | None) -> str:
    if not raw_value:
        return "Not reconciled"
    parsed = datetime.fromisoformat(raw_value.replace("Z", "+00:00"))
    return parsed.astimezone().strftime("%I:%M %p").lstrip("0")


def jobs_mechanism_label(kind: str) -> str:
    if kind == "scheduled_task":
        return "Scheduled Task"
    if kind == "startup_launcher":
        return "Startup launcher"
    return kind.replace("_", " ").title()


class DashboardApp:
    def __init__(
        self,
        config_path: Path | None = None,
        smoke_artifact_dir: Path | None = None,
        smoke_tab: str | None = None,
    ) -> None:
        self.config_path = config_path or default_config_path()
        self.config = load_config(self.config_path)
        self.active_tab = "usage"
        self.selected_interval = "15m"
        self.selected_chart_mode = "velocity"
        self.selected_metric_mode = "total"
        self.ingest_queue: queue.Queue[tuple[str, object]] = queue.Queue()
        self.ingest_in_flight = False
        self.last_ingest_error: str | None = None
        self.hotkey_registered = False
        self.smoke_artifact_dir = smoke_artifact_dir
        self.smoke_tab = smoke_tab
        self._quitting = False
        self.smoke_hotkey_triggered = False
        self.smoke_overlay_fallback = False
        self.display_timezone = self._resolve_display_timezone()
        self.loaded_font_assets = load_private_font_assets()
        self.chart_hover_regions: list[dict[str, object]] = []
        self.chart_context_region: dict[str, object] | None = None
        self.latest_events = []
        self.latest_session_context_markers: dict[str, list[object]] = {}
        self.latest_repo_legend: list[tuple[str, str]] = []
        self.latest_repo_totals: list[dict[str, int]] = []
        self.jobs_registry_path = default_jobs_registry_path(Path(self.config.codex_root))
        self.jobs_registry = {"jobs": []}
        self.jobs_snapshot: dict[str, object] = {
            "last_reconciled_at": None,
            "summary": {},
            "jobs": [],
        }
        self.jobs_detail_job_id: str | None = None
        self.jobs_status_message = "Press Refresh to inspect local Windows state."
        self.debug_log_path = self.config_path.parent / "dashboard-debug.log"
        self._append_debug_log("dashboard_started")

        self.root = tk.Tk()
        self.root.withdraw()
        self.root.title("CODEX DASHBOARD")
        self.root.protocol("WM_DELETE_WINDOW", self.quit)

        self.overlay = tk.Toplevel(self.root)
        self.overlay.withdraw()
        self.overlay_visible = False
        self.overlay.overrideredirect(True)
        self.overlay.attributes("-topmost", True)
        self.overlay.geometry(self._overlay_geometry())
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

    def _overlay_geometry(self) -> str:
        desired_width = 980
        desired_height = 660
        margin_x = 40
        margin_y = 40
        screen_width = self.root.winfo_screenwidth()
        screen_height = self.root.winfo_screenheight()
        width = min(desired_width, max(860, screen_width - (margin_x * 2)))
        height = min(desired_height, max(620, screen_height - (margin_y * 2)))
        x = min(940, max(20, screen_width - width - margin_x))
        y = min(100, max(20, screen_height - height - margin_y))
        return f"{width}x{height}+{x}+{y}"

    def _configure_style(self) -> None:
        style = ttk.Style()
        style.theme_use("clam")
        style.configure("Overlay.TFrame", background="#0a0e14")
        style.configure("Shell.TFrame", background="#1c2026")
        style.configure("Header.TFrame", background="#181c22")
        style.configure("BodyPanel.TFrame", background="#1c2026")
        style.configure("Card.TFrame", background="#181c22")
        style.configure(
            "Brand.TLabel",
            background="#181c22",
            foreground="#bff4ff",
            font=("Space Grotesk", 16, "bold"),
        )
        style.configure(
            "Badge.TLabel",
            background="#31353c",
            foreground="#8fa8bb",
            font=("Inter", 8, "bold"),
        )
        style.configure(
            "Status.TLabel",
            background="#1c2026",
            foreground="#8fa8bb",
            font=("Inter", 9),
        )
        style.configure(
            "MetricTitle.TLabel",
            background="#181c22",
            foreground="#8fa8bb",
            font=("Inter", 8, "bold"),
        )
        style.configure(
            "MetricValue.TLabel",
            background="#181c22",
            foreground="#dfe2eb",
            font=("Space Grotesk", 20, "bold"),
        )
        style.configure(
            "MetricUnit.TLabel",
            background="#181c22",
            foreground="#8fa8bb",
            font=("Inter", 9),
        )
        style.configure(
            "MetricDetail.TLabel",
            background="#181c22",
            foreground="#8fa8bb",
            font=("Inter", 9),
        )
        style.configure(
            "ChartTitle.TLabel",
            background="#1c2026",
            foreground="#dfe2eb",
            font=("Space Grotesk", 10, "bold"),
        )
        style.configure(
            "Tiny.TLabel",
            background="#1c2026",
            foreground="#6e8598",
            font=("Inter", 8),
        )
        style.configure(
            "Accent.TButton",
            background="#16d9f5",
            foreground="#10141a",
            font=("Inter", 9, "bold"),
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
            font=("Inter", 9, "bold"),
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
            font=("Inter", 8, "bold"),
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
            font=("Inter", 8, "bold"),
            borderwidth=0,
            focusthickness=0,
        )
        style.map(
            "HeaderAccent.TButton",
            background=[("active", "#2ee8ff")],
        )
        style.configure(
            "ToolbarQuiet.TButton",
            background="#303743",
            foreground="#dfe2eb",
            font=("Inter", 8, "bold"),
            borderwidth=0,
            focusthickness=0,
        )
        style.map(
            "ToolbarQuiet.TButton",
            background=[("active", "#3b4450")],
        )
        style.configure(
            "ToolbarAccent.TButton",
            background="#16d9f5",
            foreground="#10141a",
            font=("Inter", 8, "bold"),
            borderwidth=0,
            focusthickness=0,
        )
        style.map(
            "ToolbarAccent.TButton",
            background=[("active", "#2ee8ff")],
        )
        style.configure(
            "StatusValue.TLabel",
            background="#181c22",
            foreground="#bff4ff",
            font=("Inter", 12, "bold"),
        )
        style.configure(
            "StatusDetail.TLabel",
            background="#181c22",
            foreground="#8fa8bb",
            font=("Inter", 9),
        )
    def _build_overlay(self) -> None:
        self.container = ttk.Frame(self.overlay, style="Overlay.TFrame", padding=28)
        self.container.pack(fill="both", expand=True)

        self.shell = ttk.Frame(self.container, style="Shell.TFrame")
        self.shell.pack(fill="both", expand=True)

        header = ttk.Frame(self.shell, style="Header.TFrame", padding=(16, 12))
        header.pack(fill="x")
        header.columnconfigure(0, weight=1)
        brand_row = ttk.Frame(header, style="Header.TFrame")
        brand_row.grid(row=0, column=0, sticky="w")
        ttk.Label(brand_row, text="CODEX_DASHBOARD", style="Brand.TLabel").pack(side="left")
        self.tab_buttons: dict[str, tk.Label] = {}
        self.tab_underlines: dict[str, tk.Frame] = {}
        nav_row = ttk.Frame(brand_row, style="Header.TFrame")
        nav_row.pack(side="left", padx=(24, 0))
        for tab_id, label in (("usage", "Usage"), ("jobs", "Jobs")):
            tab_shell = tk.Frame(nav_row, bg=HEADER_BACKGROUND)
            tab_shell.pack(side="left", padx=(0, 20))
            tab_label = tk.Label(
                tab_shell,
                text=label.upper(),
                bg=HEADER_BACKGROUND,
                fg=TAB_INACTIVE_FOREGROUND,
                font=("Space Grotesk", 10, "bold"),
                cursor="hand2",
            )
            tab_label.pack(anchor="w")
            underline = tk.Frame(
                tab_shell,
                bg=HEADER_BACKGROUND,
                height=2,
                width=30,
            )
            underline.pack(anchor="w", pady=(5, 0))
            for widget in (tab_shell, tab_label, underline):
                widget.bind("<Button-1>", lambda _event, key=tab_id: self.select_tab(key))
            self.tab_buttons[tab_id] = tab_label
            self.tab_underlines[tab_id] = underline

        ttk.Button(
            header,
            text="X",
            style="HeaderQuiet.TButton",
            command=self.hide_overlay,
            width=3,
        ).grid(row=0, column=1, sticky="e")

        tk.Frame(self.shell, bg="#39424d", height=1).pack(fill="x")

        self.content_stack = ttk.Frame(self.shell, style="BodyPanel.TFrame")
        self.content_stack.pack(fill="both", expand=True)

        body = ttk.Frame(self.content_stack, style="BodyPanel.TFrame", padding=(16, 14))
        body.pack(fill="both", expand=True)
        self.usage_body = body

        usage_toolbar = ttk.Frame(body, style="BodyPanel.TFrame")
        usage_toolbar.pack(fill="x", pady=(0, 12))
        self.usage_header_controls = ttk.Frame(usage_toolbar, style="BodyPanel.TFrame")
        self.usage_header_controls.pack(side="left")
        self.usage_budget_controls = ttk.Frame(usage_toolbar, style="BodyPanel.TFrame")
        self.usage_budget_controls.pack(side="right")

        interval_shell = ttk.Frame(self.usage_header_controls, style="Shell.TFrame", padding=(8, 6))
        interval_shell.pack(side="left", padx=(0, 8))
        self.interval_buttons: dict[str, ttk.Button] = {}
        for interval_key in ("1m", "5m", "15m", "1h", "1d"):
            button = ttk.Button(
                interval_shell,
                text=interval_key,
                style="ToolbarQuiet.TButton",
                command=lambda key=interval_key: self.select_interval(key),
                width=4,
            )
            button.pack(side="left", padx=(0, 6))
            self.interval_buttons[interval_key] = button

        chart_mode_shell = ttk.Frame(self.usage_header_controls, style="Shell.TFrame", padding=(8, 6))
        chart_mode_shell.pack(side="left", padx=(0, 8))
        self.chart_mode_buttons: dict[str, ttk.Button] = {}
        for chart_mode, label in CHART_MODES.items():
            button = ttk.Button(
                chart_mode_shell,
                text=label,
                style="ToolbarQuiet.TButton",
                command=lambda mode=chart_mode: self.select_chart_mode(mode),
                width=6,
            )
            button.pack(side="left", padx=(0, 6))
            self.chart_mode_buttons[chart_mode] = button

        metric_mode_shell = ttk.Frame(self.usage_header_controls, style="Shell.TFrame", padding=(8, 6))
        metric_mode_shell.pack(side="left")
        self.metric_mode_buttons: dict[str, ttk.Button] = {}
        for metric_mode, label in METRIC_MODES.items():
            button = ttk.Button(
                metric_mode_shell,
                text=label,
                style="ToolbarQuiet.TButton",
                command=lambda mode=metric_mode: self.select_metric_mode(mode),
                width=6,
            )
            button.pack(side="left", padx=(0, 6))
            self.metric_mode_buttons[metric_mode] = button

        self.status_label = ttk.Label(
            body,
            text="Waiting for first ingest...",
            style="Status.TLabel",
        )
        ttk.Label(self.usage_budget_controls, text="Budget (B)", style="Status.TLabel").pack(side="left")
        self.weekly_budget_var = tk.StringVar(
            value=format_budget_billions(self.config.weekly_budget_tokens)
        )
        self.weekly_budget_entry = tk.Entry(
            self.usage_budget_controls,
            textvariable=self.weekly_budget_var,
            width=5,
            bg="#121820",
            fg="#dfe2eb",
            insertbackground="#dfe2eb",
            relief="flat",
            highlightthickness=1,
            highlightbackground="#2b323b",
            highlightcolor="#16d9f5",
            font=("Inter", 10),
        )
        self.weekly_budget_entry.pack(side="left", padx=(10, 8), ipady=4)
        ttk.Button(
            self.usage_budget_controls,
            text="Save",
            style="Accent.TButton",
            command=self.save_budget,
        ).pack(side="left", padx=(0, 10))

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

        card_headroom = ttk.Frame(metrics_row, style="Card.TFrame", padding=(12, 10))
        card_headroom.grid(row=0, column=2, sticky="nsew", padx=(0, 10))
        ttk.Label(card_headroom, text="HEADROOM", style="MetricTitle.TLabel").pack(anchor="w")
        headroom_row = ttk.Frame(card_headroom, style="Card.TFrame")
        headroom_row.pack(anchor="w", pady=(10, 4))
        self.headroom_value = ttk.Label(headroom_row, text="0", style="MetricValue.TLabel")
        self.headroom_value.pack(side="left")
        ttk.Label(headroom_row, text="TKN", style="MetricUnit.TLabel").pack(side="left", padx=(6, 0), pady=(9, 0))
        self.headroom_detail = ttk.Label(card_headroom, text="", style="MetricDetail.TLabel")
        self.headroom_detail.pack(anchor="w")

        self.status_card = ttk.Frame(metrics_row, style="Card.TFrame", padding=(0, 10))
        self.status_card.grid(row=0, column=3, sticky="nsew")
        self.status_card.columnconfigure(1, weight=1)
        self.status_accent = tk.Frame(self.status_card, bg="#16d9f5", width=2)
        self.status_accent.grid(row=0, column=0, rowspan=3, sticky="ns", padx=(0, 10))
        ttk.Label(self.status_card, text="STATUS", style="MetricTitle.TLabel").grid(row=0, column=1, sticky="w", padx=(2, 0))
        status_value_row = ttk.Frame(self.status_card, style="Card.TFrame")
        status_value_row.grid(row=1, column=1, sticky="w", padx=(2, 0), pady=(10, 0))
        self.status_dot = tk.Frame(status_value_row, bg="#16d9f5", width=7, height=7)
        self.status_dot.pack(side="left", padx=(0, 8), pady=(3, 0))
        self.status_metric_value = ttk.Label(status_value_row, text="Awaiting data", style="StatusValue.TLabel")
        self.status_metric_value.pack(side="left")
        self.status_metric_detail = ttk.Label(self.status_card, text="No ingest cycle completed yet.", style="StatusDetail.TLabel")
        self.status_metric_detail.grid(row=2, column=1, sticky="w", padx=(19, 0), pady=(4, 0))

        chart_header = ttk.Frame(body, style="BodyPanel.TFrame")
        chart_header.pack(fill="x", pady=(0, 8))
        self.chart_header_title = ttk.Label(
            chart_header,
            text=format_chart_title(
                self.selected_interval,
                self.selected_chart_mode,
                self.selected_metric_mode,
            ),
            style="ChartTitle.TLabel",
        )
        self.chart_header_title.pack(side="left")
        self.chart_header_context = ttk.Label(chart_header, text=self._timezone_label(), style="Tiny.TLabel")
        self.chart_header_context.pack(side="right")

        chart_shell = ttk.Frame(body, style="Shell.TFrame", padding=(10, 10))
        chart_shell.pack(fill="both", expand=True)
        self.canvas = tk.Canvas(
            chart_shell,
            width=880,
            height=270,
            bg="#10141a",
            highlightthickness=0,
        )
        self.canvas.pack(fill="both", expand=True)
        self.canvas.bind("<Motion>", self._on_chart_motion)
        self.canvas.bind("<Leave>", self._on_chart_leave)
        self.canvas.bind("<Button-3>", self._on_chart_right_click)
        self.chart_context_menu = tk.Menu(self.overlay, tearoff=0)
        self.chart_context_menu.add_command(
            label="Investigate with Codex",
            command=self._investigate_selected_bucket,
        )

        info_row = ttk.Frame(body, style="BodyPanel.TFrame")
        info_row.pack(fill="x", pady=(12, 0))
        info_row.columnconfigure(0, weight=1)
        self.advisory_label = ttk.Label(
            info_row,
            text="No weekly advisory yet.",
            style="Status.TLabel",
            wraplength=720,
            justify="left",
        )
        self.advisory_label.grid(row=0, column=0, sticky="w")
        self.hotkey_label = ttk.Label(
            info_row,
            text=f"Toggle: {self.config.hotkey}",
            style="Tiny.TLabel",
        )
        self.hotkey_label.grid(row=0, column=1, sticky="e", padx=(12, 0))

        self._refresh_interval_buttons()
        self._refresh_chart_mode_buttons()
        self._refresh_metric_mode_buttons()
        self._build_jobs_lane()
        self._refresh_tab_buttons()
        self._render_active_tab()

    def _build_jobs_lane(self) -> None:
        self.jobs_body = ttk.Frame(self.content_stack, style="BodyPanel.TFrame", padding=(16, 14))

        summary_row = ttk.Frame(self.jobs_body, style="BodyPanel.TFrame")
        summary_row.pack(fill="x", pady=(0, 12))
        for column in range(4):
            summary_row.columnconfigure(column, weight=1)

        self.jobs_declared_value = self._build_jobs_summary_card(summary_row, 0, "DECLARED JOBS", "0")
        self.jobs_synced_value = self._build_jobs_summary_card(summary_row, 1, "IN SYNC", "0")
        self.jobs_attention_value = self._build_jobs_summary_card(summary_row, 2, "NEEDS ATTENTION", "0")
        self.jobs_last_reconciled_value = self._build_jobs_summary_card(
            summary_row,
            3,
            "LAST RECONCILED",
            "Not reconciled",
        )

        action_row = ttk.Frame(self.jobs_body, style="Shell.TFrame", padding=(10, 10))
        action_row.pack(fill="x", pady=(0, 12))
        ttk.Label(
            action_row,
            text="FILTER: ALL_TYPES",
            style="Tiny.TLabel",
        ).pack(side="left")
        ttk.Button(
            action_row,
            text="REFRESH",
            style="Quiet.TButton",
            command=self.refresh_jobs_data,
        ).pack(side="right", padx=(8, 0))
        ttk.Button(
            action_row,
            text="FORCE RECONCILE",
            style="Accent.TButton",
            command=lambda: self.refresh_jobs_data(apply_changes=True),
        ).pack(side="right")

        self.jobs_detail_shell = ttk.Frame(self.jobs_body, style="Shell.TFrame", padding=(10, 10))
        self.jobs_detail_title = ttk.Label(self.jobs_detail_shell, text="", style="ChartTitle.TLabel")
        self.jobs_detail_title.pack(anchor="w", pady=(0, 6))
        self.jobs_detail_text = tk.Text(
            self.jobs_detail_shell,
            height=10,
            bg="#10141a",
            fg="#dfe2eb",
            relief="flat",
            wrap="word",
            font=("Inter", 9),
            insertbackground="#dfe2eb",
        )
        self.jobs_detail_text.pack(fill="x")
        self.jobs_detail_text.configure(state="disabled")

        self.jobs_rows_shell = ttk.Frame(self.jobs_body, style="Shell.TFrame", padding=(10, 10))
        self.jobs_rows_shell.pack(fill="both", expand=True)

        header_row = ttk.Frame(self.jobs_rows_shell, style="Shell.TFrame")
        header_row.pack(fill="x", pady=(0, 8))
        for text, width in (
            ("Job", 34),
            ("Mechanism", 16),
            ("Desired / observed", 20),
            ("Drift status", 18),
            ("Actions", 10),
        ):
            ttk.Label(header_row, text=text, style="Tiny.TLabel", width=width).pack(side="left")

        self.jobs_rows_container = ttk.Frame(self.jobs_rows_shell, style="Shell.TFrame")
        self.jobs_rows_container.pack(fill="both", expand=True)

    def _build_jobs_summary_card(
        self,
        parent: ttk.Frame,
        column: int,
        title: str,
        value: str,
    ) -> ttk.Label:
        card = ttk.Frame(parent, style="Card.TFrame", padding=(12, 10))
        card.grid(row=0, column=column, sticky="nsew", padx=(0, 10) if column < 3 else (0, 0))
        ttk.Label(card, text=title, style="MetricTitle.TLabel").pack(anchor="w")
        value_label = ttk.Label(card, text=value, style="MetricValue.TLabel")
        value_label.pack(anchor="w", pady=(8, 0))
        return value_label

    def select_tab(self, tab_id: str) -> None:
        self.active_tab = tab_id
        if tab_id == "jobs":
            self._prime_jobs_snapshot()
        self._render_active_tab()

    def _render_active_tab(self) -> None:
        self._refresh_tab_buttons()
        if self.active_tab == "jobs":
            self.usage_body.pack_forget()
            self.jobs_body.pack(fill="both", expand=True)
            return
        self.jobs_body.pack_forget()
        self.usage_body.pack(fill="both", expand=True)

    def _refresh_tab_buttons(self) -> None:
        for tab_id, label in self.tab_buttons.items():
            is_active = tab_id == self.active_tab
            label.configure(
                fg=TAB_ACTIVE_FOREGROUND if is_active else TAB_INACTIVE_FOREGROUND,
            )
            self.tab_underlines[tab_id].configure(
                bg=TAB_ACTIVE_UNDERLINE if is_active else HEADER_BACKGROUND,
            )

    def refresh_jobs_data(self, apply_changes: bool = False) -> None:
        try:
            self.jobs_registry = ensure_jobs_registry(codex_root=Path(self.config.codex_root))
            if apply_changes:
                self.jobs_snapshot = apply_registry(self.jobs_registry)
                self.jobs_status_message = "Jobs reconcile completed for supported drift."
            else:
                self.jobs_snapshot = reconcile_registry(self.jobs_registry)
                self.jobs_status_message = "Jobs state refreshed from local Windows state."
            if self.active_tab == "jobs":
                self.status_label.configure(text=self.jobs_status_message)
        except Exception as exc:
            self.jobs_snapshot = {
                "last_reconciled_at": None,
                "summary": {"blocked": 1},
                "jobs": [
                    {
                        "job_id": "jobs-backend-blocked",
                        "label": "Jobs backend",
                        "mechanism_label": "Registry",
                        "desired_label": "Enabled",
                        "observed_label": "Blocked",
                        "status": "blocked",
                        "reason": str(exc),
                        "details": {"error": str(exc)},
                    }
                ],
            }
            self.jobs_status_message = f"Jobs error: {exc}"
            self.status_label.configure(text=self.jobs_status_message)
        self._render_jobs_snapshot()

    def _prime_jobs_snapshot(self) -> None:
        existing_jobs = list(self.jobs_snapshot.get("jobs", []))
        if existing_jobs:
            return
        try:
            self.jobs_registry = ensure_jobs_registry(codex_root=Path(self.config.codex_root))
            self.jobs_snapshot = self._declared_jobs_snapshot()
        except Exception as exc:
            self.jobs_snapshot = {
                "last_reconciled_at": None,
                "summary": {"blocked": 1},
                "jobs": [
                    {
                        "job_id": "jobs-backend-blocked",
                        "label": "Jobs backend",
                        "mechanism_label": "Registry",
                        "desired_label": "Enabled",
                        "observed_label": "Blocked",
                        "status": "blocked",
                        "reason": str(exc),
                        "definition": {},
                        "details": {"error": str(exc)},
                    }
                ],
            }
            self.jobs_status_message = f"Jobs error: {exc}"
        self._render_jobs_snapshot()

    def _declared_jobs_snapshot(self) -> dict[str, object]:
        jobs: list[dict[str, object]] = []
        for job in self.jobs_registry.get("jobs", []):
            desired_state = str(job.get("desired_state", DESIRED_STATE_ENABLED))
            desired_label = "Enabled" if desired_state == DESIRED_STATE_ENABLED else "Disabled"
            jobs.append(
                {
                    "job_id": str(job.get("job_id", "unknown-job")),
                    "label": str(job.get("label", "Unnamed job")),
                    "kind": str(job.get("kind", "")),
                    "mechanism_label": jobs_mechanism_label(str(job.get("kind", ""))),
                    "desired_state": desired_state,
                    "desired_label": desired_label,
                    "observed_label": "Unknown",
                    "status": "unknown",
                    "reason": "State has not been checked yet. Press Refresh to inspect Windows state.",
                    "definition": dict(job.get("definition", {})),
                    "details": {},
                }
            )
        return {
            "last_reconciled_at": None,
            "summary": {},
            "jobs": jobs,
        }

    def _render_jobs_snapshot(self) -> None:
        snapshot = self.jobs_snapshot
        summary = dict(snapshot.get("summary", {}))
        jobs = list(snapshot.get("jobs", []))
        self.jobs_declared_value.configure(text=f"{len(jobs):02d}")
        self.jobs_synced_value.configure(text=f"{summary.get('in_sync', 0):02d}")
        self.jobs_attention_value.configure(text=f"{jobs_needs_attention_count(summary):02d}")
        self.jobs_last_reconciled_value.configure(
            text=format_jobs_timestamp(snapshot.get("last_reconciled_at"))
        )

        for child in self.jobs_rows_container.winfo_children():
            child.destroy()

        if not jobs:
            ttk.Label(
                self.jobs_rows_container,
                text=self.jobs_status_message,
                style="Status.TLabel",
                wraplength=760,
                justify="left",
            ).pack(anchor="w", padx=12, pady=(12, 0))
            self.jobs_detail_job_id = None
            self.jobs_detail_shell.pack_forget()
            return

        selected_job = None
        for job in jobs:
            if job["job_id"] == self.jobs_detail_job_id:
                selected_job = job
            self._build_jobs_row(job)

        if selected_job is None:
            self.jobs_detail_job_id = None
            self.jobs_detail_shell.pack_forget()
        else:
            self._show_job_details(selected_job)

    def _build_jobs_row(self, job: dict[str, object]) -> None:
        row = ttk.Frame(self.jobs_rows_container, style="Card.TFrame", padding=(12, 10))
        row.pack(fill="x", pady=(0, 8))

        title_column = ttk.Frame(row, style="Card.TFrame")
        title_column.pack(side="left", fill="x", expand=True)
        ttk.Label(
            title_column,
            text=str(job["label"]),
            style="ChartTitle.TLabel",
        ).pack(anchor="w")
        ttk.Label(
            title_column,
            text=str(job["reason"]),
            style="Status.TLabel",
            wraplength=320,
            justify="left",
        ).pack(anchor="w", pady=(4, 0))

        ttk.Label(
            row,
            text=str(job.get("mechanism_label", job.get("kind", ""))),
            style="Status.TLabel",
            width=16,
        ).pack(side="left", padx=(10, 0))
        ttk.Label(
            row,
            text=f"{job.get('desired_label', 'Unknown')} / {job.get('observed_label', 'Unknown')}",
            style="Status.TLabel",
            width=20,
        ).pack(side="left", padx=(10, 0))

        status = str(job["status"])
        status_chip = tk.Label(
            row,
            text=status.replace("_", " ").upper(),
            bg="#10141a",
            fg=JOBS_STATUS_COLORS.get(status, "#dfe2eb"),
            padx=8,
            pady=4,
            font=("Space Grotesk", 8, "bold"),
        )
        status_chip.pack(side="left", padx=(10, 0))

        ttk.Button(
            row,
            text="Details",
            style="Quiet.TButton",
            command=lambda payload=job: self.toggle_job_details(payload),
            width=8,
        ).pack(side="right")

    def toggle_job_details(self, job: dict[str, object]) -> None:
        if self.jobs_detail_job_id == job["job_id"]:
            self.jobs_detail_job_id = None
            self.jobs_detail_shell.pack_forget()
            return
        self.jobs_detail_job_id = str(job["job_id"])
        self._show_job_details(job)

    def _show_job_details(self, job: dict[str, object]) -> None:
        declared_job = {
            "job_id": job.get("job_id", ""),
            "label": job.get("label", ""),
            "kind": job.get("kind", ""),
            "desired_state": job.get("desired_state", ""),
            "definition": dict(job.get("definition", {})),
        }
        self.jobs_detail_title.configure(text=f"{job['label']} declared job")
        self.jobs_detail_text.configure(state="normal")
        self.jobs_detail_text.delete("1.0", "end")
        self.jobs_detail_text.insert("1.0", json.dumps(declared_job, indent=2, sort_keys=True))
        self.jobs_detail_text.configure(state="disabled")
        if not self.jobs_detail_shell.winfo_manager():
            self.jobs_detail_shell.pack(fill="x", pady=(0, 12), before=self.jobs_rows_shell)

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
        session_context_markers = (
            load_session_context_markers(
                connection,
                sorted({event.session_path for event in events}),
            )
            if self.selected_chart_mode == "repo"
            else {}
        )
        connection.close()
        self.latest_events = events
        self.latest_session_context_markers = session_context_markers
        self.latest_repo_legend = []
        self.latest_repo_totals = []

        raw_buckets = build_buckets(
            events,
            self.selected_interval,
            bucket_count=20,
            now=now,
            display_tz=self.display_timezone,
            metric_mode="total",
        )
        display_buckets = raw_buckets
        if self.selected_metric_mode != "total":
            display_buckets = build_buckets(
                events,
                self.selected_interval,
                bucket_count=20,
                now=now,
                display_tz=self.display_timezone,
                metric_mode=self.selected_metric_mode,
            )
        interval_seconds = INTERVAL_SECONDS[self.selected_interval]
        total_7d = sum(event.total_tokens for event in events)
        latest_advisory = next(
            (event for event in reversed(events) if event.weekly_used_percent is not None),
            None,
        )
        if maybe_upgrade_weekly_budget(
            self.config,
            total_7d,
            latest_advisory.weekly_used_percent if latest_advisory else None,
        ):
            save_config(self.config, self.config_path)
            self.weekly_budget_var.set(format_budget_billions(self.config.weekly_budget_tokens))

        pace_sample_size = min(ROLLING_PROJECTION_BUCKETS, len(raw_buckets))
        pace_tokens = rolling_average_tokens(raw_buckets, pace_sample_size)
        projected = project_weekly_burn(pace_tokens, interval_seconds)
        redline = projected > self.config.weekly_budget_tokens
        budget_line_tokens = interval_redline_tokens(
            self.config.weekly_budget_tokens,
            interval_seconds,
        )
        headroom_tokens = budget_line_tokens - pace_tokens

        self.local_total_value.configure(text=format_token_value(total_7d))
        self.local_total_detail.configure(text="in the last 7d")
        self.projected_value.configure(text=format_token_value(projected))
        self.projected_detail.configure(text=f"based on the last {pace_sample_size} bars")
        self.headroom_value.configure(text=format_signed_token_value(headroom_tokens))
        self.headroom_detail.configure(text="until exceeding budget")
        if latest_advisory is None:
            self.advisory_label.configure(text="No weekly advisory yet.")
        else:
            advisory = latest_advisory.weekly_used_percent
            reset_text = (
                f"reset in {format_reset_remaining(latest_advisory.weekly_resets_at)}"
                if latest_advisory.weekly_resets_at is not None
                else "reset time unavailable"
            )
            self.advisory_label.configure(
                text=(
                    f"Codex advisory window: {advisory:.1f}% used | "
                    f"{reset_text}"
                )
            )
        self.chart_header_title.configure(
            text=format_chart_title(
                self.selected_interval,
                self.selected_chart_mode,
                self.selected_metric_mode,
            )
        )
        self._refresh_status_surfaces(redline)
        chart_context_bits: list[str] = []
        if self.selected_chart_mode == "repo":
            repo_buckets, repo_legend, repo_totals = build_project_stacks(
                events,
                session_context_markers,
                self.selected_interval,
                bucket_count=20,
                now=now,
                display_tz=self.display_timezone,
                top_n=5,
                metric_mode=self.selected_metric_mode,
            )
            self.latest_repo_legend = repo_legend
            self.latest_repo_totals = repo_totals
            chart_context_bits.append("Top 5 repos")
            if self.selected_metric_mode == "norm":
                chart_context_bits.append("Norm")
            chart_context_bits.append(self._timezone_label())
            self.chart_header_context.configure(
                text=" | ".join(chart_context_bits)
            )
            self.draw_chart(
                repo_buckets,
                repo_legend=repo_legend,
                repo_totals=repo_totals,
                raw_buckets=raw_buckets,
                show_budget_line=self.selected_metric_mode == "total",
            )
            return
        if self.selected_metric_mode == "norm":
            chart_context_bits.append("Norm")
        chart_context_bits.append(self._timezone_label())
        self.chart_header_context.configure(text=" | ".join(chart_context_bits))
        self.draw_chart(
            display_buckets,
            raw_buckets=raw_buckets,
            show_budget_line=self.selected_metric_mode == "total",
        )

    def _refresh_status_surfaces(self, redline: bool) -> None:
        if self.last_ingest_error is not None:
            accent = "#ff5a52"
            value = "Attention"
            detail = "Ingest error detected."
        elif redline:
            accent = "#ff5a52"
            value = "Redline"
            detail = "Projected weekly burn exceeds budget."
        else:
            accent = "#bff4ff"
            value = "Operational"
            detail = "Within weekly budget."

        self.status_accent.configure(bg=accent)
        self.status_dot.configure(bg=accent)
        self.status_metric_value.configure(text=value, foreground=accent)
        self.status_metric_detail.configure(text=detail)

    def draw_chart(
        self,
        buckets,
        repo_legend: list[tuple[str, str]] | None = None,
        repo_totals: list[dict[str, int]] | None = None,
        raw_buckets=None,
        show_budget_line: bool = True,
    ) -> None:
        self.canvas.delete("all")
        self.chart_hover_regions = []
        self.chart_context_region = None
        self._hide_chart_tooltip()
        width = max(int(self.canvas.winfo_width()), int(self.canvas["width"]))
        height = max(int(self.canvas.winfo_height()), int(self.canvas["height"]))
        left = 56
        right = width - 24
        top = 18
        if repo_legend:
            top += 24
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

        if repo_legend:
            legend_x = left
            legend_y = top - 15
            for index, (_project_key, label) in enumerate(repo_legend):
                color = REPO_STACK_COLORS[index % len(REPO_STACK_COLORS)]
                self.canvas.create_rectangle(
                    legend_x,
                    legend_y - 5,
                    legend_x + 8,
                    legend_y + 3,
                    fill=color,
                    outline=REPO_STACK_OUTLINE,
                    width=1,
                )
                self.canvas.create_text(
                    legend_x + 14,
                    legend_y,
                    anchor="w",
                    text=label,
                    fill="#8fa8bb",
                    font=("Inter", 8),
                )
                legend_x += 14 + min(120, len(label) * 6 + 18)
                if legend_x > right - 120:
                    legend_x = left
                    legend_y += 14

        threshold_tokens = 0
        if show_budget_line:
            threshold_tokens = interval_redline_tokens(
                self.config.weekly_budget_tokens,
                INTERVAL_SECONDS[self.selected_interval],
            )
        max_tokens = max(
            max(bucket.total_tokens for bucket in buckets),
            threshold_tokens if show_budget_line else 0,
            1,
        )
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
                font=("Inter", 8),
            )

        if show_budget_line:
            threshold_y = bottom - ((threshold_tokens / max_tokens) * chart_height)
            threshold_color = "#8ec5ff"
            self.canvas.create_line(
                left,
                threshold_y,
                right,
                threshold_y,
                fill=threshold_color,
                width=2,
                dash=(2, 4),
            )
            label_left = left + 12
            label_right = label_left + 54
            self.canvas.create_rectangle(
                label_left,
                threshold_y - 9,
                label_right,
                threshold_y + 5,
                fill="#10141a",
                outline="",
            )
            self.canvas.create_text(
                (label_left + label_right) / 2,
                threshold_y - 2,
                text="BUDGET",
                fill=threshold_color,
                font=("Inter", 7, "bold"),
            )

        gap = 8
        bar_width = max(12, int((chart_width - gap * (len(buckets) - 1)) / len(buckets)))
        for index, bucket in enumerate(buckets):
            x0 = left + index * (bar_width + gap)
            x1 = x0 + bar_width
            hover_text = format_velocity_tooltip(bucket.total_tokens)
            raw_bucket = raw_buckets[index] if raw_buckets is not None and index < len(raw_buckets) else bucket
            if repo_legend and repo_totals is not None:
                segment_bottom = bottom
                bucket_segments = repo_totals[index] if index < len(repo_totals) else {}
                hover_text = format_repo_tooltip(bucket_segments, repo_legend)
                if bucket.total_tokens == 0:
                    self.canvas.create_rectangle(
                        x0,
                        bottom - 1,
                        x1,
                        bottom,
                        fill="#17314c",
                        outline=REPO_STACK_OUTLINE,
                        width=1,
                    )
                else:
                    for color_index, (project_key, _label) in enumerate(repo_legend):
                        segment_tokens = bucket_segments.get(project_key, 0)
                        if segment_tokens <= 0:
                            continue
                        segment_height = (segment_tokens / max_tokens) * (chart_height - 8)
                        y0 = segment_bottom - segment_height
                        self.canvas.create_rectangle(
                            x0,
                            y0,
                            x1,
                            segment_bottom,
                            fill=REPO_STACK_COLORS[color_index % len(REPO_STACK_COLORS)],
                            outline=REPO_STACK_OUTLINE,
                            width=1,
                        )
                        segment_bottom = y0
            else:
                bar_height = (bucket.total_tokens / max_tokens) * (chart_height - 8)
                y0 = bottom - bar_height
                if show_budget_line and bucket.total_tokens >= threshold_tokens and bucket.total_tokens > 0:
                    fill = "#ff7a6e"
                elif index == len(buckets) - 1:
                    fill = "#58a8ff"
                elif bucket.total_tokens == 0:
                    fill = "#17314c"
                elif index % 2 == 0:
                    fill = "#2f6fa3"
                else:
                    fill = "#265d8a"
                self.canvas.create_rectangle(x0, y0, x1, bottom, fill=fill, outline="")
            if index % 3 == 0 or index == len(buckets) - 1:
                label = format_tick_label(bucket.start_at, self.selected_interval)
                self.canvas.create_text(
                    (x0 + x1) / 2,
                    bottom + 14,
                    text=label,
                    fill="#6e8598",
                    font=("Inter", 8),
                )
                self.chart_hover_regions.append(
                {
                    "x0": x0,
                    "x1": x1,
                    "y0": top,
                    "y1": bottom,
                    "text": hover_text,
                    "bucket": raw_bucket,
                    "display_total": bucket.total_tokens,
                    "repo_totals": dict(repo_totals[index]) if repo_totals is not None and index < len(repo_totals) else {},
                }
            )

    def select_interval(self, interval_key: str) -> None:
        self.selected_interval = interval_key
        self._refresh_interval_buttons()
        self.refresh_data()

    def select_chart_mode(self, chart_mode: str) -> None:
        self.selected_chart_mode = chart_mode
        self._refresh_chart_mode_buttons()
        self.refresh_data()

    def select_metric_mode(self, metric_mode: str) -> None:
        self.selected_metric_mode = metric_mode
        self._refresh_metric_mode_buttons()
        self.refresh_data()

    def _refresh_interval_buttons(self) -> None:
        for key, button in self.interval_buttons.items():
            button.configure(
                style="ToolbarAccent.TButton" if key == self.selected_interval else "ToolbarQuiet.TButton"
            )

    def _refresh_chart_mode_buttons(self) -> None:
        for key, button in self.chart_mode_buttons.items():
            button.configure(
                style="ToolbarAccent.TButton" if key == self.selected_chart_mode else "ToolbarQuiet.TButton"
            )

    def _refresh_metric_mode_buttons(self) -> None:
        for key, button in self.metric_mode_buttons.items():
            button.configure(
                style="ToolbarAccent.TButton" if key == self.selected_metric_mode else "ToolbarQuiet.TButton"
            )

    def _chart_region_at(self, x: int, y: int) -> dict[str, object] | None:
        for region in self.chart_hover_regions:
            if (
                region["x0"] <= x <= region["x1"]
                and region["y0"] <= y <= region["y1"]
            ):
                return region
        return None

    def _on_chart_motion(self, event) -> None:
        region = self._chart_region_at(event.x, event.y)
        if region is not None:
            self._show_chart_tooltip(event.x, event.y, str(region["text"]))
            return
        self._hide_chart_tooltip()

    def _on_chart_leave(self, _event) -> None:
        self._hide_chart_tooltip()

    def _on_chart_right_click(self, event) -> None:
        region = self._chart_region_at(event.x, event.y)
        if region is None:
            self._append_debug_log(
                f"right_click_miss x={event.x} y={event.y} mode={self.selected_chart_mode} metric={self.selected_metric_mode} interval={self.selected_interval}"
            )
            return
        self.chart_context_region = region
        bucket = region.get("bucket")
        if bucket is not None:
            display_total = int(region.get("display_total") or 0)
            self._append_debug_log(
                "right_click_bucket "
                f"bucket_start={bucket.start_at.isoformat()} "
                f"bucket_end={bucket.end_at.isoformat()} "
                f"bucket_total={bucket.total_tokens} "
                f"display_total={display_total} "
                f"mode={self.selected_chart_mode} "
                f"metric={self.selected_metric_mode} "
                f"interval={self.selected_interval} "
                f"x={event.x} y={event.y}"
            )
        self._show_chart_tooltip(event.x, event.y, str(region["text"]))
        try:
            self.chart_context_menu.tk_popup(event.x_root, event.y_root)
        finally:
            self.chart_context_menu.grab_release()

    def _investigate_selected_bucket(self) -> None:
        region = self.chart_context_region
        if region is None:
            self.status_label.configure(text="Select a bucket before launching Codex investigation.")
            return
        codex_executable = shutil.which("codex")
        if codex_executable is None:
            self.status_label.configure(text="Codex CLI was not found in PATH.")
            return
        bucket = region.get("bucket")
        if bucket is None:
            self.status_label.configure(text="No bucket context is available for investigation.")
            return
        investigation = build_bucket_investigation(
            bucket,
            self.latest_events,
            self.latest_session_context_markers,
            self.selected_interval,
            self.selected_chart_mode,
            Path(self.config.codex_root),
        )
        brief_path = write_bucket_investigation(
            investigation,
            default_investigations_path(),
            datetime.now(),
        )
        report_path = report_path_for_brief(brief_path)
        launch_command = build_codex_launch_command(
            codex_executable,
            brief_path,
            report_path,
            investigation.workspace_root,
            investigation.add_dirs,
        )
        self._append_debug_log(f"investigation_brief path={brief_path}")
        self._append_debug_log(f"investigation_report path={report_path}")
        self._append_debug_log(
            f"investigation_command {subprocess.list2cmdline(launch_command)}"
        )
        try:
            subprocess.Popen(
                launch_command,
                cwd=str(investigation.workspace_root),
                creationflags=getattr(subprocess, "CREATE_NEW_CONSOLE", 0),
            )
        except OSError as exc:
            self._append_debug_log(f"investigation_launch_failed error={exc}")
            self.status_label.configure(text=f"Failed to launch Codex investigation: {exc}")
            return
        self._append_debug_log("investigation_launch_started")
        self.status_label.configure(
            text=(
                "Codex investigation launched for "
                f"{bucket.start_at.strftime('%I:%M %p').lstrip('0')}. "
                f"Report: {report_path.name}"
            )
        )

    def _show_chart_tooltip(self, x: int, y: int, text: str) -> None:
        self.canvas.delete("chart-tooltip")
        tooltip_text = self.canvas.create_text(
            x + 14,
            y + 14,
            anchor="nw",
            text=text,
            justify="left",
            fill="#dfe2eb",
            font=("Inter", 8),
            tags="chart-tooltip",
        )
        bbox = self.canvas.bbox(tooltip_text)
        if bbox is None:
            return
        left, top, right, bottom = bbox
        shift_x = 0
        shift_y = 0
        canvas_width = max(int(self.canvas.winfo_width()), int(self.canvas["width"]))
        canvas_height = max(int(self.canvas.winfo_height()), int(self.canvas["height"]))
        if right + 8 > canvas_width:
            shift_x = canvas_width - right - 8
        if bottom + 8 > canvas_height:
            shift_y = canvas_height - bottom - 8
        if shift_x or shift_y:
            self.canvas.move(tooltip_text, shift_x, shift_y)
            bbox = self.canvas.bbox(tooltip_text)
            if bbox is None:
                return
            left, top, right, bottom = bbox
        background = self.canvas.create_rectangle(
            left - 6,
            top - 5,
            right + 6,
            bottom + 5,
            fill="#0d131b",
            outline="#39424d",
            width=1,
            tags="chart-tooltip",
        )
        self.canvas.tag_lower(background, tooltip_text)

    def _hide_chart_tooltip(self) -> None:
        self.canvas.delete("chart-tooltip")

    def _append_debug_log(self, message: str) -> None:
        self.debug_log_path.parent.mkdir(parents=True, exist_ok=True)
        timestamp = datetime.now().isoformat(timespec="seconds")
        with self.debug_log_path.open("a", encoding="utf-8") as handle:
            handle.write(f"{timestamp} {message}\n")

    def save_budget(self) -> None:
        raw_value = self.weekly_budget_var.get().strip()
        try:
            budget = parse_budget_billions(raw_value)
        except ValueError:
            self.status_label.configure(text="Weekly budget must be a number of billions, for example 3.5.")
            return
        self.config.weekly_budget_tokens = max(1, budget)
        save_config(self.config, self.config_path)
        self.weekly_budget_var.set(format_budget_billions(self.config.weekly_budget_tokens))
        self.status_label.configure(text=f"Saved weekly budget: {format_budget_billions(self.config.weekly_budget_tokens)}B")
        self.refresh_data()

    def toggle_overlay(self) -> None:
        if self.smoke_artifact_dir is not None:
            self.smoke_hotkey_triggered = True
        if self.overlay_visible:
            self.hide_overlay()
        else:
            self.show_overlay()

    def show_overlay(self) -> None:
        self.refresh_data()
        self.overlay.deiconify()
        self.overlay_visible = True
        self.overlay.lift()
        self.overlay.focus_force()

    def hide_overlay(self) -> None:
        self.chart_context_region = None
        self._hide_chart_tooltip()
        self.overlay.withdraw()
        self.overlay_visible = False

    def quit(self) -> None:
        if self._quitting:
            return
        self._quitting = True
        if self.hotkey_registered:
            self.hotkey.unregister()
        unload_private_font_assets(self.loaded_font_assets)
        self.root.quit()
        self.overlay.destroy()
        self.root.destroy()

    def _run_smoke_capture(self) -> None:
        artifact_dir = self.smoke_artifact_dir
        if artifact_dir is None:
            return
        artifact_dir.mkdir(parents=True, exist_ok=True)
        if self.smoke_tab in {"usage", "jobs"}:
            self.select_tab(self.smoke_tab)
        if self.smoke_tab == "jobs":
            self.refresh_jobs_data()
        if not self.overlay_visible:
            self.smoke_overlay_fallback = True
            self.show_overlay()
        self.overlay.update_idletasks()
        if self.active_tab == "usage":
            self.canvas.postscript(
                file=str(artifact_dir / "overlay-chart.ps"),
                colormode="color",
            )
        summary_lines = [
            f"active_tab={self.active_tab}",
            self.status_label.cget("text"),
            f"hotkey_triggered={self.smoke_hotkey_triggered}",
            f"overlay_fallback={self.smoke_overlay_fallback}",
        ]
        if self.active_tab == "usage":
            summary_lines.extend(
                [
                    f"interval={self.selected_interval}",
                    f"metric_mode={self.selected_metric_mode}",
                    f"weekly_budget={self.config.weekly_budget_tokens}",
                    f"7d_total={self.local_total_value.cget('text')}",
                    f"projected={self.projected_value.cget('text')}",
                    f"headroom={self.headroom_value.cget('text')}",
                    (
                        f"budget_line={format_token_value(interval_redline_tokens(self.config.weekly_budget_tokens, INTERVAL_SECONDS[self.selected_interval]))}"
                        if self.selected_metric_mode == "total"
                        else "budget_line=hidden_in_norm_mode"
                    ),
                    f"status={self.status_metric_value.cget('text')}",
                    self.advisory_label.cget("text"),
                ]
            )
        else:
            summary_lines.extend(
                [
                    f"jobs_registry={self.jobs_registry_path}",
                    f"jobs_declared={self.jobs_declared_value.cget('text')}",
                    f"jobs_in_sync={self.jobs_synced_value.cget('text')}",
                    f"jobs_needs_attention={self.jobs_attention_value.cget('text')}",
                    f"jobs_last_reconciled={self.jobs_last_reconciled_value.cget('text')}",
                ]
            )
        summary = "\n".join(summary_lines)
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
