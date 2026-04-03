from __future__ import annotations

from pathlib import Path


UNKNOWN_PROJECT_KEY = "__unknown__"
UNKNOWN_PROJECT_LABEL = "Unknown"


def _normalize_cwd(raw_cwd: str | None) -> str | None:
    if raw_cwd is None:
        return None
    normalized = raw_cwd.strip()
    if not normalized:
        return None
    try:
        return str(Path(normalized).resolve(strict=False))
    except OSError:
        return normalized


def _find_git_root(path: Path) -> Path | None:
    current = path
    if not current.exists():
        current = current.parent
    while True:
        if (current / ".git").exists():
            return current
        parent = current.parent
        if parent == current:
            return None
        current = parent


def resolve_project_identity(raw_cwd: str | None) -> tuple[str | None, str, str]:
    cwd = _normalize_cwd(raw_cwd)
    if cwd is None:
        return None, UNKNOWN_PROJECT_LABEL, "unknown"

    cwd_path = Path(cwd)
    git_root = _find_git_root(cwd_path)
    if git_root is not None:
        label = git_root.name or str(git_root)
        return str(git_root), label, "repo"

    label = cwd_path.name or str(cwd_path)
    return cwd, label, "cwd"
