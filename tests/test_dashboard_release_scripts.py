from __future__ import annotations

import json
import os
import shutil
import subprocess
import tempfile
import unittest
from pathlib import Path


REPO_ROOT = Path(__file__).resolve().parents[1]
PUBLISH_SCRIPT = REPO_ROOT / "scripts" / "Publish-DashboardRelease.ps1"
TEST_SCRIPT = REPO_ROOT / "scripts" / "Test-DashboardRelease.ps1"


class DashboardReleaseScriptsTests(unittest.TestCase):
    def run_powershell(
        self,
        command: str,
        *,
        env: dict[str, str] | None = None,
        check: bool = True,
        timeout: int = 60,
    ) -> subprocess.CompletedProcess[str]:
        executable = shutil.which("powershell") or shutil.which("powershell.exe")
        if executable is None:
            self.skipTest("Windows PowerShell is not available")

        completed = subprocess.run(
            [
                executable,
                "-NoProfile",
                "-ExecutionPolicy",
                "Bypass",
                "-Command",
                command,
            ],
            check=False,
            capture_output=True,
            text=True,
            timeout=timeout,
            env=env,
        )
        if check:
            self.assertEqual(
                completed.returncode,
                0,
                f"PowerShell failed:\nSTDOUT:\n{completed.stdout}\nSTDERR:\n{completed.stderr}",
            )
        return completed

    def release_test_env(self, root: str) -> dict[str, str]:
        env = os.environ.copy()
        env["LOCALAPPDATA"] = str(Path(root) / "Local")
        env["APPDATA"] = str(Path(root) / "Roaming")
        return env

    def test_publish_plan_uses_local_appdata_release_roots(self) -> None:
        with tempfile.TemporaryDirectory() as tmp:
            env = self.release_test_env(tmp)

            completed = self.run_powershell(f"& '{PUBLISH_SCRIPT}' -PlanOnly", env=env)
            data = json.loads(completed.stdout)

        self.assertEqual(Path(data["source_repo_root"]), REPO_ROOT)
        self.assertEqual(Path(data["source_app_root"]), REPO_ROOT / "app")
        self.assertEqual(
            Path(data["releases_root"]),
            Path(tmp) / "Local" / "CodexDashboard" / "dashboard-releases",
        )
        self.assertEqual(
            Path(data["current_release_manifest_path"]),
            Path(tmp) / "Local" / "CodexDashboard" / "dashboard-current-release.json",
        )
        self.assertEqual(
            Path(data["launcher_script_path"]),
            Path(tmp) / "Local" / "CodexDashboard" / "dashboard-launcher" / "Start-CodexDashboard.ps1",
        )
        self.assertEqual(
            Path(data["startup_path"]),
            Path(tmp)
            / "Roaming"
            / "Microsoft"
            / "Windows"
            / "Start Menu"
            / "Programs"
            / "Startup"
            / "CodexDashboard.cmd",
        )
        self.assertTrue(data["would_pin_current"])
        self.assertTrue(data["would_install_startup"])
        self.assertEqual(data["source_mode"], "git_commit")

    def test_publish_pins_dashboard_release_and_startup_launcher(self) -> None:
        with tempfile.TemporaryDirectory() as tmp:
            env = self.release_test_env(tmp)

            publish = self.run_powershell(f"& '{PUBLISH_SCRIPT}' -AllowDirty", env=env)
            release = json.loads(publish.stdout)
            status = self.run_powershell(f"& '{TEST_SCRIPT}'", env=env)
            data = json.loads(status.stdout)

        self.assertEqual(release["component"], "dashboard_frontend")
        self.assertEqual(release["source_mode"], "git_commit")
        self.assertFalse(release["source_dirty"])
        self.assertGreater(len(release["files"]), 0)
        self.assertIsNone(data["current_release_error"])
        self.assertEqual(data["current_release"]["release_id"], release["release_id"])
        self.assertEqual(data["current_release"]["source_mode"], "git_commit")
        self.assertTrue(data["launcher_exists"])
        self.assertTrue(data["startup_uses_pinned_launcher"])
        self.assertGreater(data["current_release"]["file_count"], 0)

    def test_release_status_reports_hash_mismatch_for_tampered_file(self) -> None:
        with tempfile.TemporaryDirectory() as tmp:
            env = self.release_test_env(tmp)

            publish = self.run_powershell(f"& '{PUBLISH_SCRIPT}' -AllowDirty -NoStartup", env=env)
            release = json.loads(publish.stdout)
            first_file = release["files"][0]
            release_file_path = Path(release["release_root"]) / Path(first_file["path"])
            with release_file_path.open("a", encoding="utf-8") as handle:
                handle.write("\n# tampered\n")

            status = self.run_powershell(f"& '{TEST_SCRIPT}'", env=env)
            data = json.loads(status.stdout)

        self.assertIn("hash mismatch", data["current_release_error"])
        self.assertEqual(data["running_process_count"], 0)


if __name__ == "__main__":
    unittest.main()
