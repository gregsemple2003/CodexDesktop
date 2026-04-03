from __future__ import annotations

import tempfile
import unittest
from pathlib import Path
from unittest import mock

from app.codex_dashboard.hotkey import MOD_ALT, MOD_CONTROL, parse_hotkey
from app.codex_dashboard.startup import startup_command


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


if __name__ == "__main__":
    unittest.main()
