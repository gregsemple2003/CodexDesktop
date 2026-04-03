from __future__ import annotations

import sys
from pathlib import Path


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


def repo_root() -> Path:
    return Path(__file__).resolve().parents[2]


def startup_command() -> str:
    python_executable = startup_python_executable()
    root = repo_root()
    return (
        "@echo off\r\n"
        f'cd /d "{root}"\r\n'
        f'"{python_executable}" -m app.codex_dashboard\r\n'
    )


def is_startup_enabled() -> bool:
    return startup_script_path().exists()


def set_startup_enabled(enabled: bool) -> None:
    script_path = startup_script_path()
    if enabled:
        script_path.write_text(startup_command(), encoding="utf-8")
    elif script_path.exists():
        script_path.unlink()
