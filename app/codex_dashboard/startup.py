from __future__ import annotations

import os
import sys
from pathlib import Path

from .paths import app_data_root


def startup_script_path() -> Path:
    return (
        Path.home()
        / "AppData"
        / "Roaming"
        / "Microsoft"
        / "Windows"
        / "Start Menu"
        / "Programs"
        / "Startup"
        / "CodexDashboard.cmd"
    )


def startup_python_executable() -> Path:
    executable = Path(sys.executable)
    if executable.name.lower() == "python.exe":
        pythonw = executable.with_name("pythonw.exe")
        if pythonw.exists():
            return pythonw
    return executable


def dashboard_launcher_path() -> Path:
    return app_data_root() / "dashboard-launcher" / "Start-CodexDashboard.ps1"


def startup_powershell_executable() -> Path:
    system_root = Path(os.environ.get("SystemRoot", r"C:\Windows"))
    stable_path = system_root / "System32" / "WindowsPowerShell" / "v1.0" / "powershell.exe"
    if stable_path.exists():
        return stable_path
    return Path("powershell.exe")


def startup_command() -> str:
    powershell_executable = startup_powershell_executable()
    launcher_path = dashboard_launcher_path()
    return (
        "@echo off\r\n"
        f'"{powershell_executable}" -NoProfile -ExecutionPolicy Bypass -WindowStyle Hidden -File "{launcher_path}"\r\n'
    )


def is_startup_enabled() -> bool:
    return startup_script_path().exists()


def set_startup_enabled(enabled: bool) -> None:
    script_path = startup_script_path()
    if enabled:
        with script_path.open("w", encoding="utf-8", newline="") as handle:
            handle.write(startup_command())
    elif script_path.exists():
        script_path.unlink()
