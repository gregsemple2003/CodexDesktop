from __future__ import annotations

import tempfile
import unittest
from datetime import datetime
from pathlib import Path
from types import SimpleNamespace
from unittest import mock

from app.codex_dashboard.hotkey import MOD_ALT, MOD_CONTROL, GlobalHotkey, parse_hotkey
from app.codex_dashboard.config import (
    DashboardConfig,
    advisory_implied_weekly_budget_tokens,
    maybe_upgrade_weekly_budget,
)
from app.codex_dashboard.startup import startup_command
from app.codex_dashboard.ui import (
    DashboardApp,
    format_budget_billions,
    format_chart_title,
    format_jobs_timestamp,
    format_repo_tooltip,
    format_reset_remaining,
    format_tick_label,
    format_signed_token_value,
    format_token_value,
    format_velocity_tooltip,
    interval_redline_tokens,
    jobs_needs_attention_count,
    parse_budget_billions,
    rolling_average_tokens,
)


class DesktopSupportTests(unittest.TestCase):
    def test_parse_hotkey_supports_ctrl_alt_space(self) -> None:
        modifiers, virtual_key = parse_hotkey("Ctrl+Alt+Space")
        self.assertEqual(modifiers, MOD_CONTROL | MOD_ALT)
        self.assertEqual(virtual_key, 0x20)

    def test_startup_command_uses_python_module_entrypoint(self) -> None:
        fake_repo = Path("C:/Agent/CodexDashboard")
        fake_pythonw = Path("C:/Python313/pythonw.exe")
        with mock.patch("app.codex_dashboard.startup.repo_root", return_value=fake_repo):
            with mock.patch(
                "app.codex_dashboard.startup.startup_python_executable",
                return_value=fake_pythonw,
            ):
                command = startup_command()
        self.assertIn('cd /d "C:\\Agent\\CodexDashboard"', command)
        self.assertIn('"C:\\Python313\\pythonw.exe" -m app.codex_dashboard', command)

    def test_global_hotkey_poll_drains_pending_callbacks(self) -> None:
        callback = mock.Mock()
        hotkey = GlobalHotkey("Ctrl+Alt+Space", callback)
        hotkey._pending_callbacks.put(None)
        hotkey._pending_callbacks.put(None)

        hotkey.poll()

        self.assertEqual(callback.call_count, 2)

    def test_format_tick_label_uses_compact_ampm_hours(self) -> None:
        self.assertEqual(format_tick_label(datetime(2026, 4, 2, 19, 0), "1h"), "7PM")
        self.assertEqual(format_tick_label(datetime(2026, 4, 2, 5, 0), "1h"), "5AM")
        self.assertEqual(format_tick_label(datetime(2026, 4, 2, 17, 15), "15m"), "5:15PM")

    def test_format_token_value_uses_human_suffixes(self) -> None:
        self.assertEqual(format_token_value(999), "999")
        self.assertEqual(format_token_value(100_000), "100K")
        self.assertEqual(format_token_value(2_500_000), "2.5M")

    def test_format_signed_token_value_shows_direction(self) -> None:
        self.assertEqual(format_signed_token_value(250_000), "+250K")
        self.assertEqual(format_signed_token_value(-2_500_000), "-2.5M")
        self.assertEqual(format_signed_token_value(0), "0")

    def test_format_chart_title_uses_interval_name(self) -> None:
        self.assertEqual(format_chart_title("1h"), "Token Velocity per 1 Hour")
        self.assertEqual(format_chart_title("5m"), "Token Velocity per 5 Minutes")
        self.assertEqual(format_chart_title("15m", "repo"), "Repo Share per 15 Minutes")
        self.assertEqual(
            format_chart_title("15m", "velocity", "norm"),
            "Normalized Token Velocity per 15 Minutes",
        )
        self.assertEqual(
            format_chart_title("15m", "repo", "norm"),
            "Normalized Repo Share per 15 Minutes",
        )

    def test_velocity_tooltip_uses_scalar_human_format(self) -> None:
        self.assertEqual(format_velocity_tooltip(2_500_000), "2.5M")

    def test_repo_tooltip_lists_nonzero_repo_breakdown(self) -> None:
        tooltip = format_repo_tooltip(
            {
                "repo-a": 2_500_000,
                "repo-b": 0,
                "__other__": 1_200_000,
            },
            [
                ("repo-a", "RepoA"),
                ("repo-b", "RepoB"),
                ("__other__", "Other"),
            ],
        )
        self.assertEqual(tooltip, "RepoA: 2.5M\nOther: 1.2M")

    def test_budget_billions_helpers_format_and_parse(self) -> None:
        self.assertEqual(format_budget_billions(3_550_000_000), "3.5")
        self.assertEqual(parse_budget_billions("3.5"), 3_500_000_000)
        self.assertEqual(parse_budget_billions("3.5B"), 3_500_000_000)
        self.assertEqual(parse_budget_billions("3550000000"), 3_550_000_000)

    def test_format_reset_remaining_prefers_days_then_hours(self) -> None:
        now = datetime(2026, 4, 2, 22, 0)
        self.assertEqual(
            format_reset_remaining(int(datetime(2026, 4, 8, 0, 0).timestamp()), now=now),
            "5.1d",
        )
        self.assertEqual(
            format_reset_remaining(int(datetime(2026, 4, 3, 1, 0).timestamp()), now=now),
            "3.0h",
        )

    def test_interval_redline_tokens_scales_budget_to_bucket_size(self) -> None:
        self.assertEqual(interval_redline_tokens(8_000_000, 3600), 47_619)

    def test_rolling_average_tokens_uses_last_n_buckets(self) -> None:
        buckets = [
            SimpleNamespace(total_tokens=5),
            SimpleNamespace(total_tokens=10),
            SimpleNamespace(total_tokens=15),
            SimpleNamespace(total_tokens=20),
            SimpleNamespace(total_tokens=30),
        ]
        self.assertEqual(rolling_average_tokens(buckets, 4), 19)

    def test_advisory_implied_weekly_budget_rounds_to_nearest_50m(self) -> None:
        self.assertEqual(
            advisory_implied_weekly_budget_tokens(1_838_350_974, 52.0),
            3_550_000_000,
        )

    def test_maybe_upgrade_weekly_budget_replaces_legacy_default(self) -> None:
        config = DashboardConfig(
            codex_root="C:/Users/gregs/.codex",
            db_path="C:/Users/gregs/AppData/Local/CodexDashboard/dashboard.db",
            weekly_budget_tokens=8_000_000,
        )

        updated = maybe_upgrade_weekly_budget(config, 1_838_350_974, 52.0)

        self.assertTrue(updated)
        self.assertEqual(config.weekly_budget_tokens, 3_550_000_000)

    def test_jobs_needs_attention_count_excludes_in_sync(self) -> None:
        self.assertEqual(
            jobs_needs_attention_count(
                {
                    "in_sync": 2,
                    "drifted": 1,
                    "missing": 1,
                    "blocked": 1,
                }
            ),
            3,
        )

    def test_format_jobs_timestamp_uses_local_clock_label(self) -> None:
        self.assertEqual(
            format_jobs_timestamp("2026-04-04T12:01:55-04:00"),
            "12:01 PM",
        )

    def test_select_tab_does_not_trigger_jobs_refresh(self) -> None:
        app = SimpleNamespace(
            active_tab="usage",
            _prime_jobs_snapshot=mock.Mock(),
            _render_active_tab=mock.Mock(),
            refresh_jobs_data=mock.Mock(),
        )

        DashboardApp.select_tab(app, "jobs")

        self.assertEqual(app.active_tab, "jobs")
        app._prime_jobs_snapshot.assert_called_once_with()
        app._render_active_tab.assert_called_once_with()
        app.refresh_jobs_data.assert_not_called()

    def test_declared_jobs_snapshot_uses_unknown_observed_state(self) -> None:
        app = SimpleNamespace(
            jobs_registry={
                "jobs": [
                    {
                        "job_id": "codex-dashboard-startup",
                        "label": "CodexDashboard overlay at sign-in",
                        "kind": "startup_launcher",
                        "desired_state": "enabled",
                        "definition": {"script_path": "C:/Startup/CodexDashboard.cmd"},
                    }
                ]
            }
        )

        snapshot = DashboardApp._declared_jobs_snapshot(app)

        self.assertEqual(snapshot["last_reconciled_at"], None)
        self.assertEqual(len(snapshot["jobs"]), 1)
        self.assertEqual(snapshot["jobs"][0]["desired_label"], "Enabled")
        self.assertEqual(snapshot["jobs"][0]["observed_label"], "Unknown")
        self.assertEqual(snapshot["jobs"][0]["status"], "unknown")

    def test_show_job_details_includes_definition_and_observed_sections(self) -> None:
        jobs_detail_title = mock.Mock()
        jobs_detail_text = mock.Mock()
        jobs_detail_shell = mock.Mock()
        jobs_detail_shell.winfo_manager.return_value = False
        app = SimpleNamespace(
            jobs_detail_title=jobs_detail_title,
            jobs_detail_text=jobs_detail_text,
            jobs_detail_shell=jobs_detail_shell,
            jobs_rows_shell=object(),
        )

        DashboardApp._show_job_details(
            app,
            {
                "job_id": "codex-dashboard-startup",
                "label": "CodexDashboard overlay at sign-in",
                "kind": "startup_launcher",
                "desired_state": "enabled",
                "desired_label": "Enabled",
                "observed_label": "Enabled",
                "status": "in_sync",
                "reason": "Job matches the managed definition.",
                "definition": {"script_path": "C:/Startup/CodexDashboard.cmd"},
                "details": {"command_text": "@echo off"},
            },
        )

        inserted_text = jobs_detail_text.insert.call_args.args[1]
        self.assertIn('"definition"', inserted_text)
        self.assertNotIn("command_text", inserted_text)
        jobs_detail_shell.pack.assert_called_once()

    def test_toggle_overlay_hides_visible_overlay_from_explicit_flag(self) -> None:
        app = SimpleNamespace(
            smoke_artifact_dir=None,
            smoke_hotkey_triggered=False,
            overlay_visible=True,
            show_overlay=mock.Mock(),
            hide_overlay=mock.Mock(),
        )

        DashboardApp.toggle_overlay(app)

        app.hide_overlay.assert_called_once_with()
        app.show_overlay.assert_not_called()

    def test_show_and_hide_overlay_update_visible_flag(self) -> None:
        overlay = mock.Mock()
        app = SimpleNamespace(
            overlay=overlay,
            overlay_visible=False,
            refresh_data=mock.Mock(),
            chart_context_region="selected",
            _hide_chart_tooltip=mock.Mock(),
        )

        DashboardApp.show_overlay(app)

        self.assertTrue(app.overlay_visible)
        app.refresh_data.assert_called_once_with()
        overlay.deiconify.assert_called_once_with()
        overlay.lift.assert_called_once_with()
        overlay.focus_force.assert_called_once_with()

        DashboardApp.hide_overlay(app)

        self.assertFalse(app.overlay_visible)
        self.assertIsNone(app.chart_context_region)
        app._hide_chart_tooltip.assert_called_once_with()
        overlay.withdraw.assert_called_once_with()


if __name__ == "__main__":
    unittest.main()
