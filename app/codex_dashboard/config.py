from __future__ import annotations

import json
from dataclasses import asdict, dataclass
from pathlib import Path

from .paths import default_codex_root, default_config_path, default_db_path


@dataclass(slots=True)
class DashboardConfig:
    codex_root: str
    db_path: str
    polling_seconds: int = 5
    weekly_budget_tokens: int = 8_000_000
    startup_enabled: bool = False
    hotkey: str = "Ctrl+Alt+Space"

    @classmethod
    def defaults(cls) -> "DashboardConfig":
        return cls(
            codex_root=str(default_codex_root()),
            db_path=str(default_db_path()),
        )


def load_config(path: Path | None = None) -> DashboardConfig:
    config_path = path or default_config_path()
    if not config_path.exists():
        config = DashboardConfig.defaults()
        save_config(config, config_path)
        return config
    payload = json.loads(config_path.read_text(encoding="utf-8"))
    defaults = DashboardConfig.defaults()
    return DashboardConfig(
        codex_root=str(payload.get("codex_root", defaults.codex_root)),
        db_path=str(payload.get("db_path", defaults.db_path)),
        polling_seconds=int(payload.get("polling_seconds", defaults.polling_seconds)),
        weekly_budget_tokens=int(
            payload.get("weekly_budget_tokens", defaults.weekly_budget_tokens)
        ),
        startup_enabled=bool(payload.get("startup_enabled", defaults.startup_enabled)),
        hotkey=str(payload.get("hotkey", defaults.hotkey)),
    )


def save_config(config: DashboardConfig, path: Path | None = None) -> None:
    config_path = path or default_config_path()
    config_path.parent.mkdir(parents=True, exist_ok=True)
    config_path.write_text(
        json.dumps(asdict(config), indent=2, sort_keys=True),
        encoding="utf-8",
    )
