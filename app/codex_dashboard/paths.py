from __future__ import annotations

from pathlib import Path


def default_codex_root() -> Path:
    return Path.home() / ".codex"


def app_data_root() -> Path:
    return Path.home() / "AppData" / "Local" / "CodexDashboard"


def default_db_path() -> Path:
    return app_data_root() / "dashboard.db"


def default_config_path() -> Path:
    return app_data_root() / "config.json"
