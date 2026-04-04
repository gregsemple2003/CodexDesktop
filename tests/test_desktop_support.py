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


if __name__ == "__main__":
    unittest.main()
