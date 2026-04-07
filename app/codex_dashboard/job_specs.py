from __future__ import annotations

import json
from pathlib import Path
from typing import Any

from .paths import job_specs_root

API_VERSION = "codex.jobs/v1"
DESIRED_STATE_DISABLED = "disabled"
DESIRED_STATE_ENABLED = "enabled"
EXECUTOR_TYPE_CODEX_EXEC = "codex_exec"
TRIGGER_TYPE_MANUAL = "manual"
TRIGGER_TYPE_SCHEDULE = "schedule"
TRIGGER_TYPE_WEBHOOK = "webhook"


def iter_job_spec_paths(codex_root: Path | None = None) -> list[Path]:
    specs_root = job_specs_root(codex_root)
    if not specs_root.exists():
        return []
    return sorted(
        path
        for path in specs_root.rglob("*.json")
        if path.is_file() and not path.name.endswith(".schema.json")
    )


def load_job_spec(path: Path) -> dict[str, Any]:
    return json.loads(path.read_text(encoding="utf-8"))


def validate_job_spec(spec: dict[str, Any]) -> None:
    required = {
        "api_version",
        "job_id",
        "label",
        "description",
        "desired_state",
        "triggers",
        "executor",
        "runtime",
    }
    missing = sorted(required.difference(spec))
    if missing:
        raise ValueError(f"Missing required job spec fields: {', '.join(missing)}")

    if spec["api_version"] != API_VERSION:
        raise ValueError(f"Unsupported api_version: {spec['api_version']}")

    if spec["desired_state"] not in {DESIRED_STATE_ENABLED, DESIRED_STATE_DISABLED}:
        raise ValueError(f"Unsupported desired_state: {spec['desired_state']}")

    triggers = spec["triggers"]
    if not isinstance(triggers, list) or not triggers:
        raise ValueError("Job spec must define at least one trigger")

    for trigger in triggers:
        trigger_type = trigger.get("type")
        if trigger_type == TRIGGER_TYPE_SCHEDULE:
            if not trigger.get("cron") or not trigger.get("timezone"):
                raise ValueError("Schedule triggers require cron and timezone")
        elif trigger_type == TRIGGER_TYPE_MANUAL:
            continue
        elif trigger_type == TRIGGER_TYPE_WEBHOOK:
            if not trigger.get("path"):
                raise ValueError("Webhook triggers require path")
        else:
            raise ValueError(f"Unsupported trigger type: {trigger_type}")

    executor = spec["executor"]
    if executor.get("type") != EXECUTOR_TYPE_CODEX_EXEC:
        raise ValueError(f"Unsupported executor type: {executor.get('type')}")
    if not executor.get("cwd") or not executor.get("entrypoint"):
        raise ValueError("codex_exec executor requires cwd and entrypoint")
    if not isinstance(executor.get("args", []), list):
        raise ValueError("codex_exec executor args must be a list")

    runtime = spec["runtime"]
    if not runtime.get("workflow_type") or not runtime.get("task_queue"):
        raise ValueError("runtime requires workflow_type and task_queue")


def load_validated_job_specs(codex_root: Path | None = None) -> list[dict[str, Any]]:
    specs: list[dict[str, Any]] = []
    seen_job_ids: set[str] = set()
    for path in iter_job_spec_paths(codex_root):
        spec = load_job_spec(path)
        validate_job_spec(spec)
        job_id = spec["job_id"]
        if job_id in seen_job_ids:
            raise ValueError(f"Duplicate job_id: {job_id}")
        seen_job_ids.add(job_id)
        specs.append(spec)
    return specs
