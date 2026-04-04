from __future__ import annotations

from pathlib import Path


def repo_root() -> Path:
    return Path(__file__).resolve().parents[2]


def default_codex_root() -> Path:
    return Path.home() / ".codex"


def orchestration_root(codex_root: Path | None = None) -> Path:
    return (codex_root or default_codex_root()) / "Orchestration"


def default_jobs_registry_path(codex_root: Path | None = None) -> Path:
    return orchestration_root(codex_root) / "codex-jobs-registry.json"


def app_data_root() -> Path:
    return Path.home() / "AppData" / "Local" / "CodexDashboard"


def default_db_path() -> Path:
    return app_data_root() / "dashboard.db"


def default_config_path() -> Path:
    return app_data_root() / "config.json"


def default_investigations_path() -> Path:
    return repo_root() / "Tracking" / "Investigations"
