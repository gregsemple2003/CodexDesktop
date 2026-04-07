from __future__ import annotations

import json
import os
from collections import Counter
from datetime import datetime
from typing import Any
from urllib import error, parse, request

DEFAULT_JOBS_BACKEND_URL = "http://127.0.0.1:4318"
JOBS_BACKEND_URL_ENV = "CODEX_DASHBOARD_JOBS_BACKEND_URL"
REQUEST_TIMEOUT_SECONDS = 10.0


class JobsBackendError(RuntimeError):
    pass


def configured_jobs_backend_url() -> str:
    configured = os.environ.get(JOBS_BACKEND_URL_ENV, "").strip()
    if configured:
        return configured
    return DEFAULT_JOBS_BACKEND_URL


def fetch_jobs_snapshot(base_url: str = DEFAULT_JOBS_BACKEND_URL) -> dict[str, object]:
    payload = _request_json("GET", _join_url(base_url, "/api/v1/jobs"))
    return map_state_view_to_jobs_snapshot(payload)


def sync_jobs_snapshot(
    base_url: str = DEFAULT_JOBS_BACKEND_URL,
) -> tuple[dict[str, object], dict[str, Any]]:
    report = _request_json("POST", _join_url(base_url, "/sync"))
    state = report.get("state", {})
    if not isinstance(state, dict):
        raise JobsBackendError("Sync response did not include a valid state payload.")
    return map_state_view_to_jobs_snapshot(state), report


def start_job_run(job_id: str, base_url: str = DEFAULT_JOBS_BACKEND_URL) -> dict[str, Any]:
    encoded_job_id = parse.quote(job_id, safe="")
    payload = _request_json("POST", _join_url(base_url, f"/api/v1/jobs/{encoded_job_id}/run"))
    if not isinstance(payload, dict):
        raise JobsBackendError("Run-now response was not a JSON object.")
    return payload


def jobs_backend_error_snapshot(message: str) -> dict[str, object]:
    return {
        "last_reconciled_at": None,
        "summary": {"blocked": 1},
        "jobs": [
            {
                "job_id": "jobs-backend-blocked",
                "label": "Jobs backend",
                "kind": "orchestration_backend",
                "mechanism_label": "Backend",
                "desired_state": "enabled",
                "desired_label": "Enabled",
                "observed_label": "Blocked",
                "status": "blocked",
                "reason": message,
                "definition": {"error": message},
                "supports_run_now": False,
            }
        ],
    }


def map_state_view_to_jobs_snapshot(payload: dict[str, Any]) -> dict[str, object]:
    jobs_payload = payload.get("jobs", [])
    if not isinstance(jobs_payload, list):
        raise JobsBackendError("Jobs backend payload did not contain a jobs list.")

    jobs = [map_backend_job(job_payload) for job_payload in jobs_payload if isinstance(job_payload, dict)]
    summary = Counter(str(job.get("status", "unknown")) for job in jobs)
    last_sync = payload.get("last_sync", {})
    last_reconciled_at = None
    if isinstance(last_sync, dict):
        last_reconciled_at = _text(last_sync.get("last_success_at"))
    return {
        "generated_at": _text(payload.get("generated_at")),
        "last_reconciled_at": last_reconciled_at,
        "summary": dict(summary),
        "jobs": jobs,
    }


def map_backend_job(job_payload: dict[str, Any]) -> dict[str, object]:
    status = _text(job_payload.get("status"), default="unknown")
    desired_state = _text(job_payload.get("desired_state"), default="enabled")
    triggers = _dict_list(job_payload.get("triggers"))
    schedules = _dict_list(job_payload.get("schedules"))
    recent_runs = _dict_list(job_payload.get("recent_runs"))
    latest_run = recent_runs[0] if recent_runs else _latest_schedule_run(schedules)
    next_run_at = _next_schedule_time(schedules)
    supports_run_now = any(_text(trigger.get("type")) == "manual" for trigger in triggers)

    return {
        "job_id": _text(job_payload.get("job_id"), default="unknown-job"),
        "label": _text(job_payload.get("label"), default="Unnamed job"),
        "kind": "orchestration_backend",
        "mechanism_label": trigger_label(triggers),
        "desired_state": desired_state,
        "desired_label": "Enabled" if desired_state == "enabled" else "Disabled",
        "observed_label": observed_status_label(status),
        "status": status,
        "reason": reason_text(
            status=status,
            next_run_at=next_run_at,
            latest_run=latest_run,
            supports_run_now=supports_run_now,
            schedules=schedules,
        ),
        "definition": dict(job_payload),
        "supports_run_now": supports_run_now,
    }


def trigger_label(triggers: list[dict[str, Any]]) -> str:
    labels: list[str] = []
    trigger_types = {_text(trigger.get("type")) for trigger in triggers}
    if "schedule" in trigger_types:
        labels.append("Schedule")
    if "manual" in trigger_types:
        labels.append("Manual")
    if "webhook" in trigger_types:
        labels.append("Webhook")
    if not labels:
        return "Backend"
    return " + ".join(labels)


def observed_status_label(status: str) -> str:
    labels = {
        "in_sync": "In sync",
        "drifted": "Drifted",
        "missing": "Missing",
        "disabled": "Disabled",
        "blocked": "Blocked",
        "unknown": "Unknown",
    }
    return labels.get(status, status.replace("_", " ").title())


def reason_text(
    *,
    status: str,
    next_run_at: str | None,
    latest_run: dict[str, Any] | None,
    supports_run_now: bool,
    schedules: list[dict[str, Any]],
) -> str:
    if status == "blocked":
        note = _first_schedule_note(schedules)
        return note or "The orchestration backend reported a blocked state."
    if status == "missing":
        return "Desired job is missing from the orchestration runtime."
    if status == "drifted":
        note = _first_schedule_note(schedules)
        if note:
            return f"Runtime drift detected. {note}"
        return "Runtime drift detected between Git desired state and Temporal."
    if status == "disabled":
        return "Desired state is disabled."

    parts: list[str] = []
    if next_run_at:
        parts.append(f"Next run {local_clock_label(next_run_at)}")
    latest_run_at = _text((latest_run or {}).get("actual_time")) or _text((latest_run or {}).get("schedule_time"))
    if latest_run_at:
        parts.append(f"Last run {local_clock_label(latest_run_at)}")
    if supports_run_now:
        parts.append("Run now available")
    if parts:
        return ". ".join(parts) + "."
    return "Backend state is current."


def local_clock_label(raw_value: str) -> str:
    parsed = _parse_timestamp(raw_value)
    return parsed.astimezone().strftime("%I:%M %p").lstrip("0")


def _latest_schedule_run(schedules: list[dict[str, Any]]) -> dict[str, Any] | None:
    for schedule in schedules:
        recent_runs = _dict_list(schedule.get("recent_runs"))
        if recent_runs:
            return recent_runs[0]
    return None


def _next_schedule_time(schedules: list[dict[str, Any]]) -> str | None:
    for schedule in schedules:
        next_action_times = schedule.get("next_action_times", [])
        if isinstance(next_action_times, list) and next_action_times:
            return _text(next_action_times[0])
    return None


def _first_schedule_note(schedules: list[dict[str, Any]]) -> str | None:
    for schedule in schedules:
        note = _text(schedule.get("note"))
        if note:
            return note
    return None


def _request_json(method: str, url: str) -> dict[str, Any]:
    req = request.Request(
        url,
        method=method,
        headers={"Accept": "application/json"},
    )
    try:
        with request.urlopen(req, timeout=REQUEST_TIMEOUT_SECONDS) as response:
            raw = response.read().decode("utf-8")
    except error.HTTPError as exc:
        body = exc.read().decode("utf-8", errors="replace").strip()
        raise JobsBackendError(f"{method} {url} failed with HTTP {exc.code}: {body or exc.reason}") from exc
    except OSError as exc:
        raise JobsBackendError(f"{method} {url} failed: {exc}") from exc

    try:
        payload = json.loads(raw)
    except json.JSONDecodeError as exc:
        raise JobsBackendError(f"{method} {url} returned invalid JSON: {exc}") from exc
    if not isinstance(payload, dict):
        raise JobsBackendError(f"{method} {url} returned an unexpected JSON payload.")
    return payload


def _join_url(base_url: str, path: str) -> str:
    return base_url.rstrip("/") + path


def _dict_list(value: Any) -> list[dict[str, Any]]:
    if not isinstance(value, list):
        return []
    return [item for item in value if isinstance(item, dict)]


def _parse_timestamp(raw_value: str):
    return datetime.fromisoformat(raw_value.replace("Z", "+00:00"))


def _text(value: Any, default: str | None = None) -> str | None:
    if value is None:
        return default
    text = str(value).strip()
    if text == "":
        return default
    return text
