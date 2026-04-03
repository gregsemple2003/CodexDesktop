from __future__ import annotations

import tempfile
import unittest
from datetime import datetime
from pathlib import Path
from unittest import mock

from app.codex_dashboard.hotkey import MOD_ALT, MOD_CONTROL, GlobalHotkey, parse_hotkey
from app.codex_dashboard.startup import startup_command
from app.codex_dashboard.ui import (
    format_chart_title,
    format_tick_label,
    format_token_value,
    interval_redline_tokens,
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

    def test_format_chart_title_uses_interval_name(self) -> None:
        self.assertEqual(format_chart_title("1h"), "Token Velocity per 1 Hour")
        self.assertEqual(format_chart_title("5m"), "Token Velocity per 5 Minutes")

    def test_interval_redline_tokens_scales_budget_to_bucket_size(self) -> None:
        self.assertEqual(interval_redline_tokens(8_000_000, 3600), 47_619)


if __name__ == "__main__":
    unittest.main()
