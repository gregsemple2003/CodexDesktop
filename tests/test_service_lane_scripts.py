from __future__ import annotations

import hashlib
import json
import os
import shutil
import subprocess
import tempfile
import unittest
from pathlib import Path


REPO_ROOT = Path(__file__).resolve().parents[1]
LANE_HELPERS = REPO_ROOT / "backend" / "orchestration" / "scripts" / "LaneHelpers.ps1"


class ServiceLaneScriptsTests(unittest.TestCase):
    def run_powershell(
        self,
        command: str,
        *,
        env: dict[str, str] | None = None,
        check: bool = True,
        timeout: int = 20,
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

    def resolve_powershell_path(self, executable: str) -> str:
        command = f". '{LANE_HELPERS}'; Get-PowerShellExecutablePath"
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
            timeout=20,
        )
        self.assertEqual(
            completed.returncode,
            0,
            f"{executable} failed:\nSTDOUT:\n{completed.stdout}\nSTDERR:\n{completed.stderr}",
        )
        return completed.stdout.strip()

    def assert_stable_powershell_path(self, resolved: str) -> None:
        self.assertTrue(resolved, "resolver returned an empty path")
        self.assertTrue(Path(resolved).exists(), f"resolved path does not exist: {resolved}")
        self.assertNotIn(
            "\\windowsapps\\microsoft.powershell_",
            resolved.lower(),
            "scheduled task runner must not capture Store PowerShell's versioned package path",
        )

    def test_windows_powershell_resolver_returns_stable_existing_path(self) -> None:
        executable = shutil.which("powershell") or shutil.which("powershell.exe")
        if executable is None:
            self.skipTest("Windows PowerShell is not available")

        self.assert_stable_powershell_path(self.resolve_powershell_path(executable))

    def test_pwsh_resolver_does_not_capture_windowsapps_package_path(self) -> None:
        executable = shutil.which("pwsh") or shutil.which("pwsh.exe")
        if executable is None:
            self.skipTest("PowerShell 7 is not available")

        self.assert_stable_powershell_path(self.resolve_powershell_path(executable))

    def test_service_lane_runner_is_outside_repo_checkout(self) -> None:
        with tempfile.TemporaryDirectory() as tmp:
            env = os.environ.copy()
            env["LOCALAPPDATA"] = tmp
            command = (
                f". '{LANE_HELPERS}'; "
                "$c = Get-OrchestrationLaneConfig -Lane service; "
                "[pscustomobject]@{ runner = $c.RunnerScriptPath; manifest = $c.CurrentReleaseManifestPath; releases = $c.ReleasesRoot } | ConvertTo-Json"
            )

            completed = self.run_powershell(command, env=env)
            data = json.loads(completed.stdout)

        self.assertIn("orchestration-service-lane", data["runner"])
        self.assertIn("launcher", data["runner"])
        self.assertNotIn(str(REPO_ROOT / "backend" / "orchestration" / "scripts"), data["runner"])
        self.assertTrue(data["manifest"].endswith("current-release.json"))
        self.assertTrue(data["releases"].endswith("releases"))

    def test_service_lane_requires_pinned_release_manifest(self) -> None:
        with tempfile.TemporaryDirectory() as tmp:
            env = os.environ.copy()
            env["LOCALAPPDATA"] = tmp
            command = f". '{LANE_HELPERS}'; $c = Get-OrchestrationLaneConfig -Lane service; Get-ServiceLaneCurrentRelease -Config $c"

            completed = self.run_powershell(command, env=env, check=False)

        self.assertNotEqual(completed.returncode, 0)
        self.assertIn("No pinned service-lane release manifest", completed.stderr + completed.stdout)

    def test_service_lane_current_release_validates_binary_and_compose_hashes(self) -> None:
        with tempfile.TemporaryDirectory() as tmp:
            env = os.environ.copy()
            env["LOCALAPPDATA"] = tmp
            runtime_root = Path(tmp) / "CodexDashboard" / "orchestration-service-lane"
            release_root = runtime_root / "releases" / "release-test"
            binary_path = release_root / "bin" / "controlplane-service-lane.exe"
            compose_path = release_root / "docker-compose.temporal-postgres.yml"
            binary_path.parent.mkdir(parents=True)
            binary_path.write_bytes(b"test binary")
            compose_path.write_text("services: {}\n", encoding="utf-8")
            manifest = {
                "schema_version": 1,
                "lane": "service",
                "release_id": "release-test",
                "release_root": str(release_root),
                "git_commit": "a" * 40,
                "source_dirty": False,
                "binary_path": str(binary_path),
                "binary_sha256": hashlib.sha256(binary_path.read_bytes()).hexdigest(),
                "compose_file_path": str(compose_path),
                "compose_file_sha256": hashlib.sha256(compose_path.read_bytes()).hexdigest(),
            }
            runtime_root.mkdir(parents=True, exist_ok=True)
            (runtime_root / "current-release.json").write_text(json.dumps(manifest), encoding="utf-8")
            command = (
                f". '{LANE_HELPERS}'; "
                "$c = Get-OrchestrationLaneConfig -Lane service; "
                "$r = Get-ServiceLaneCurrentRelease -Config $c; "
                "$r.release_id"
            )

            completed = self.run_powershell(command, env=env)

        self.assertEqual(completed.stdout.strip(), "release-test")

    def test_service_lane_current_release_rejects_hash_mismatch(self) -> None:
        with tempfile.TemporaryDirectory() as tmp:
            env = os.environ.copy()
            env["LOCALAPPDATA"] = tmp
            runtime_root = Path(tmp) / "CodexDashboard" / "orchestration-service-lane"
            release_root = runtime_root / "releases" / "release-test"
            binary_path = release_root / "bin" / "controlplane-service-lane.exe"
            compose_path = release_root / "docker-compose.temporal-postgres.yml"
            binary_path.parent.mkdir(parents=True)
            binary_path.write_bytes(b"test binary")
            compose_path.write_text("services: {}\n", encoding="utf-8")
            manifest = {
                "schema_version": 1,
                "lane": "service",
                "release_id": "release-test",
                "release_root": str(release_root),
                "git_commit": "a" * 40,
                "source_dirty": False,
                "binary_path": str(binary_path),
                "binary_sha256": "0" * 64,
                "compose_file_path": str(compose_path),
                "compose_file_sha256": hashlib.sha256(compose_path.read_bytes()).hexdigest(),
            }
            runtime_root.mkdir(parents=True, exist_ok=True)
            (runtime_root / "current-release.json").write_text(json.dumps(manifest), encoding="utf-8")
            command = f". '{LANE_HELPERS}'; $c = Get-OrchestrationLaneConfig -Lane service; Get-ServiceLaneCurrentRelease -Config $c"

            completed = self.run_powershell(command, env=env, check=False)

        self.assertNotEqual(completed.returncode, 0)
        self.assertIn("binary hash mismatch", completed.stderr + completed.stdout)
