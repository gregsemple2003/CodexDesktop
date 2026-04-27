from __future__ import annotations

import json
import os
from collections import Counter
from datetime import datetime
from pathlib import Path
from typing import Any
from urllib import error, parse, request


DEFAULT_TASKS_BACKEND_URL = "http://127.0.0.1:4318"
TASKS_BACKEND_URL_ENV = "CODEX_DASHBOARD_TASKS_BACKEND_URL"
TASKS_SNAPSHOT_PATH_ENV = "CODEX_DASHBOARD_TASKS_SNAPSHOT_PATH"
REQUEST_TIMEOUT_SECONDS = 10.0


class TasksBackendError(RuntimeError):
    pass


def configured_tasks_backend_url() -> str:
    configured = os.environ.get(TASKS_BACKEND_URL_ENV, "").strip()
    if configured:
        return configured
    return DEFAULT_TASKS_BACKEND_URL


def fetch_tasks_snapshot(base_url: str = DEFAULT_TASKS_BACKEND_URL) -> dict[str, object]:
    fixture_path = os.environ.get(TASKS_SNAPSHOT_PATH_ENV, "").strip()
    if fixture_path:
        payload = _read_snapshot_fixture(Path(fixture_path))
    else:
        payload = _request_json("GET", _join_url(base_url, "/api/v1/tasks"))
    return map_backend_tasks_snapshot(payload)


def dispatch_task(task_id: str, base_url: str = DEFAULT_TASKS_BACKEND_URL) -> dict[str, Any]:
    encoded_task_id = parse.quote(task_id, safe="")
    return _request_json("POST", _join_url(base_url, f"/api/v1/tasks/{encoded_task_id}/dispatch"))


def poke_task_run(run_id: str, base_url: str = DEFAULT_TASKS_BACKEND_URL) -> dict[str, Any]:
    encoded_run_id = parse.quote(run_id, safe="")
    return _request_json("POST", _join_url(base_url, f"/api/v1/task-runs/{encoded_run_id}/poke"))


def pause_task_run(run_id: str, base_url: str = DEFAULT_TASKS_BACKEND_URL) -> dict[str, Any]:
    encoded_run_id = parse.quote(run_id, safe="")
    return _request_json("POST", _join_url(base_url, f"/api/v1/task-runs/{encoded_run_id}/interrupt"))


def retry_task_run_workload(run_id: str, base_url: str = DEFAULT_TASKS_BACKEND_URL) -> dict[str, Any]:
    encoded_run_id = parse.quote(run_id, safe="")
    return _request_json("POST", _join_url(base_url, f"/api/v1/task-runs/{encoded_run_id}/retry-workload"))


def tasks_backend_error_snapshot(message: str) -> dict[str, object]:
    return {
        "status": "backend_unavailable",
        "generated_at": None,
        "summary": empty_tasks_summary(),
        "tasks": [],
        "message": message,
    }


def empty_tasks_summary() -> dict[str, int]:
    return {
        "needs_you": 0,
        "sleeping": 0,
        "running": 0,
        "blocked": 0,
        "ready": 0,
    }


def map_backend_tasks_snapshot(payload: dict[str, Any]) -> dict[str, object]:
    tasks_payload = payload.get("tasks", payload.get("items", []))
    if not isinstance(tasks_payload, list):
        raise TasksBackendError("Tasks backend payload did not contain a tasks list.")

    tasks = [
        task
        for task in (map_backend_task(item) for item in tasks_payload if isinstance(item, dict))
        if task is not None
    ]
    summary = empty_tasks_summary()
    summary.update(Counter(str(task.get("summary_bucket", "ready")) for task in tasks))

    return {
        "status": _text(payload.get("status"), default="ok"),
        "generated_at": _text(payload.get("generated_at")),
        "last_refreshed_at": _text(payload.get("last_refreshed_at"), default=_text(payload.get("generated_at"))),
        "summary": summary,
        "tasks": tasks,
        "message": _text(payload.get("message"), default="Tasks state loaded from orchestration backend."),
    }


def map_backend_task(task_payload: dict[str, Any]) -> dict[str, object] | None:
    if _is_unpromoted_candidate(task_payload):
        return None

    top_level_state_envelope = _dict(task_payload.get("state_envelope"))
    top_level_state = _text(
        top_level_state_envelope.get("state"),
        default=_text(task_payload.get("status")),
    )
    if top_level_state in {"completed", "cancelled"}:
        return None

    latest_run = _dict(task_payload.get("latest_run"))
    state_envelope = _dict(latest_run.get("state_envelope")) or top_level_state_envelope
    attention = _dict(latest_run.get("attention")) or _dict(task_payload.get("attention"))
    wait_contract = _dict(latest_run.get("wait_contract"))
    deep_context = _dict(latest_run.get("deep_context")) or _dict(task_payload.get("deep_context"))
    state = _text(state_envelope.get("state"), default=_text(latest_run.get("status"), default="ready"))
    if state in {"completed", "cancelled"}:
        return None
    actions = _dict(latest_run.get("actions")) or _dict(task_payload.get("actions"))

    task_id = _text(task_payload.get("task_id"), default="unknown-task")
    title = _text(task_payload.get("title"), default=task_id)
    meaning_summary = (
        _text(latest_run.get("meaning_summary"))
        or _text(task_payload.get("meaning_summary"))
        or _text(state_envelope.get("state_summary"))
        or "Committed task is known to the backend."
    )
    recent_change = (
        _text(latest_run.get("last_progress_summary"))
        or _text(task_payload.get("last_progress_summary"))
        or _text(state_envelope.get("state_summary"))
        or "No recent progress summary was provided."
    )
    next_expected_event = (
        _text(state_envelope.get("next_expected_event"))
        or _text(wait_contract.get("resume_when"))
        or _text(task_payload.get("next_expected_event"))
        or "No next expected event was provided."
    )

    artifacts = _artifact_refs(task_payload, latest_run, state_envelope, wait_contract, deep_context)
    launch_targets = _launch_targets(deep_context)

    return {
        "task_id": task_id,
        "latest_run_id": _text(latest_run.get("run_id"), default=""),
        "title": title,
        "meaning_summary": meaning_summary,
        "why": _task_why(task_payload, meaning_summary),
        "state": state,
        "state_label": state_label(state, attention),
        "summary_bucket": summary_bucket(state, attention),
        "stream_group": stream_group(state, attention),
        "reason": _reason_line(state_envelope, wait_contract, latest_run),
        "recent_change": recent_change,
        "next_expected_event": next_expected_event,
        "freshness_label": freshness_label(latest_run, task_payload),
        "provenance_label": provenance_label(task_payload),
        "provenance_detail": provenance_detail(task_payload),
        "artifacts": artifacts,
        "actions": visible_actions(actions, task_payload, latest_run, launch_targets),
        "launch_targets": launch_targets,
        "backend_payload": dict(task_payload),
    }


def state_label(state: str, attention: dict[str, Any] | None = None) -> str:
    labels = {
        "ready": "Ready",
        "queued": "Queued",
        "dispatching": "Dispatching",
        "running": "Running",
        "waiting_for_human": "Waiting on you",
        "blocked": "Blocked",
        "sleeping_or_stalled": "Sleeping",
        "interrupted": "Paused",
        "completed": "Completed",
        "cancelled": "Cancelled",
        "failed": "Failed",
    }
    return labels.get(state, state.replace("_", " ").title())


def summary_bucket(state: str, attention: dict[str, Any] | None = None) -> str:
    if state == "waiting_for_human":
        return "needs_you"
    if state in {"running", "queued", "dispatching"}:
        return "running"
    if state in {"blocked", "failed"}:
        return "blocked"
    if state == "sleeping_or_stalled":
        return "sleeping"
    return "ready"


def stream_group(state: str, attention: dict[str, Any] | None = None) -> str:
    bucket = summary_bucket(state, attention)
    groups = {
        "needs_you": "Needs Attention",
        "running": "Running",
        "ready": "Ready to Dispatch",
        "blocked": "Waiting or Blocked",
        "sleeping": "Sleeping / Stalled",
    }
    return groups.get(bucket, "Ready to Dispatch")


def provenance_label(task_payload: dict[str, Any]) -> str:
    provenance = _promotion_provenance(task_payload)
    raw = (
        _text(task_payload.get("provenance_label"))
        or _text(task_payload.get("source_provenance"))
        or _text(task_payload.get("source"))
        or _text(task_payload.get("origin"))
        or _text(task_payload.get("kind"))
        or _text(provenance.get("source"))
        or _text(provenance.get("source_type"))
        or _text(provenance.get("source_surface"))
    )
    promoted_from = (
        _text(task_payload.get("promoted_from"))
        or _text(provenance.get("promoted_from"))
        or _text(provenance.get("source"))
        or _text(provenance.get("source_type"))
        or _text(provenance.get("source_surface"))
    )
    if promoted_from:
        return f"Promoted from {promoted_from.replace('_', ' ').title()}"
    if not raw:
        return "Authored"
    normalized = raw.lower().replace("_", " ").replace("-", " ")
    if "promoted" in normalized and "dream" in normalized:
        return "Promoted from Dream"
    if "promoted" in normalized and "review" in normalized:
        return "Promoted from Review"
    if normalized in {"dream", "review"}:
        return f"Promoted from {normalized.title()}"
    if "candidate" in normalized:
        return "Promoted"
    if normalized in {"system", "system authored"}:
        return "System-authored"
    return raw.replace("_", " ").title()


def provenance_detail(task_payload: dict[str, Any]) -> str:
    provenance = _promotion_provenance(task_payload)
    if not provenance:
        return provenance_label(task_payload)
    parts: list[str] = []
    for label, key in (
        ("Source packet", "source_packet"),
        ("Problem", "source_problem"),
        ("Winner", "source_winner"),
        ("Option task", "source_option_task"),
        ("Promoted by", "promoted_by"),
    ):
        value = _text(provenance.get(key))
        if value:
            parts.append(f"{label}: {value}")
    promoted_at = _text(provenance.get("promoted_at"))
    if promoted_at:
        parts.append(f"Promoted at: {promoted_at}")
    if parts:
        return ". ".join(parts) + "."
    return provenance_label(task_payload)


def visible_actions(
    backend_actions: dict[str, Any],
    task_payload: dict[str, Any],
    latest_run: dict[str, Any],
    launch_targets: list[dict[str, object]],
) -> list[dict[str, object]]:
    actions: list[dict[str, object]] = []
    task_id = _text(task_payload.get("task_id"), default="")
    run_id = _text(latest_run.get("run_id"), default="")
    for backend_name, label in (
        ("dispatch", "Dispatch"),
        ("interrupt", "Pause"),
        ("poke", "Poke"),
        ("retry-workload", "Continue"),
        ("resume", "Resume"),
        ("continue", "Continue"),
    ):
        availability = _dict(backend_actions.get(backend_name))
        if availability:
            if backend_name in {"interrupt", "poke", "retry-workload", "resume", "continue"} and not run_id:
                continue
            actions.append(
                {
                    "label": label,
                    "backend_action": backend_name,
                    "task_id": task_id,
                    "run_id": run_id,
                    "allowed": bool(availability.get("allowed")),
                    "reason": _block_reason(availability),
                    "destructive": False,
                }
            )
    if not actions and bool(_dict(task_payload.get("dispatch_readiness")).get("ready")):
        actions.append(
            {
                "label": "Dispatch",
                "backend_action": "dispatch",
                "task_id": task_id,
                "run_id": run_id,
                "allowed": True,
                "reason": "",
            }
        )

    seen_targets: set[str] = set()
    for target in launch_targets:
        label = _launch_action_label(target)
        uri = _text(target.get("uri"))
        key = f"{label}:{uri}"
        if key in seen_targets:
            continue
        seen_targets.add(key)
        actions.append(
            {
                "label": label,
                "backend_action": "open",
                "task_id": task_id,
                "run_id": run_id,
                "allowed": bool(uri) or bool(target.get("command")),
                "reason": _text(target.get("label"), default=""),
                "target": target,
            }
        )

    if not any(action["label"] == "Open Task" for action in actions):
        task_root = _text(task_payload.get("declared_task_root")) or _text(_dict(latest_run.get("repo_lane")).get("run_artifact_root"))
        if task_root:
            actions.append(
                {
                    "label": "Open Task",
                    "backend_action": "open",
                    "task_id": task_id,
                    "run_id": run_id,
                    "allowed": True,
                    "reason": task_root,
                    "target": {"kind": "task_artifact", "label": "Open Task", "uri": task_root},
                }
            )
    return actions


def freshness_label(latest_run: dict[str, Any], task_payload: dict[str, Any]) -> str:
    raw = (
        _text(latest_run.get("last_progress_at"))
        or _text(task_payload.get("updated_at"))
        or _text(task_payload.get("declared_task_revision"))
    )
    if not raw:
        return "Freshness unknown"
    parsed = _parse_timestamp(raw)
    if parsed is None:
        return raw
    return parsed.astimezone().strftime("%b %d, %I:%M %p").replace(" 0", " ")


def _read_snapshot_fixture(path: Path) -> dict[str, Any]:
    try:
        payload = json.loads(path.read_text(encoding="utf-8"))
    except OSError as exc:
        raise TasksBackendError(f"Could not read tasks snapshot fixture {path}: {exc}") from exc
    except json.JSONDecodeError as exc:
        raise TasksBackendError(f"Tasks snapshot fixture {path} is invalid JSON: {exc}") from exc
    if not isinstance(payload, dict):
        raise TasksBackendError(f"Tasks snapshot fixture {path} must contain a JSON object.")
    return payload


def _request_json(method: str, url: str) -> dict[str, Any]:
    req = request.Request(url, method=method, headers={"Accept": "application/json"})
    try:
        with request.urlopen(req, timeout=REQUEST_TIMEOUT_SECONDS) as response:
            raw = response.read().decode("utf-8")
    except error.HTTPError as exc:
        body = exc.read().decode("utf-8", errors="replace").strip()
        raise TasksBackendError(f"{method} {url} failed with HTTP {exc.code}: {body or exc.reason}") from exc
    except OSError as exc:
        raise TasksBackendError(_format_request_os_error(method, url, exc)) from exc

    try:
        payload = json.loads(raw)
    except json.JSONDecodeError as exc:
        raise TasksBackendError(f"{method} {url} returned invalid JSON: {exc}") from exc
    if not isinstance(payload, dict):
        raise TasksBackendError(f"{method} {url} returned an unexpected JSON payload.")
    return payload


def _is_unpromoted_candidate(task_payload: dict[str, Any]) -> bool:
    task_id = _text(task_payload.get("task_id"), default="")
    durable_state = " ".join(
        filter(
            None,
            (
                _text(task_payload.get("kind")),
                _text(task_payload.get("status")),
                _text(task_payload.get("lifecycle_state")),
                _text(task_payload.get("promotion_status")),
                _text(task_payload.get("review_status")),
                _text(task_payload.get("provenance_label")),
                _text(task_payload.get("source_provenance")),
                _text(task_payload.get("source")),
            ),
        )
    ).lower()
    provenance = _promotion_provenance(task_payload)
    provenance_state = " ".join(
        filter(
            None,
            (
                _text(provenance.get("status")),
                _text(provenance.get("promotion_status")),
                _text(provenance.get("source")),
                _text(provenance.get("source_type")),
            ),
        )
    ).lower()
    durable_state = f"{durable_state} {provenance_state}".strip()
    promoted = bool(task_payload.get("promoted")) or "promoted" in durable_state
    committed = bool(task_payload.get("committed")) or "committed" in durable_state or task_id.startswith("Task-")
    return "candidate" in durable_state and not promoted and not committed


def _task_why(task_payload: dict[str, Any], fallback: str) -> str:
    return (
        _text(task_payload.get("why"))
        or _text(task_payload.get("task_goal"))
        or _text(task_payload.get("intent"))
        or fallback
    )


def _reason_line(
    state_envelope: dict[str, Any],
    wait_contract: dict[str, Any],
    latest_run: dict[str, Any],
) -> str:
    return (
        _text(wait_contract.get("why_blocked"))
        or _text(state_envelope.get("state_summary"))
        or _text(latest_run.get("failure_summary"))
        or _text(state_envelope.get("reason_code"))
        or "Backend did not provide a detailed reason yet."
    )


def _artifact_refs(
    task_payload: dict[str, Any],
    latest_run: dict[str, Any],
    state_envelope: dict[str, Any],
    wait_contract: dict[str, Any],
    deep_context: dict[str, Any],
) -> list[dict[str, object]]:
    refs: list[dict[str, object]] = []
    for container in (state_envelope, wait_contract):
        for ref in _dict_list(container.get("evidence_refs")):
            refs.append(
                {
                    "label": _text(ref.get("label"), default=_text(ref.get("type"), default="Evidence")),
                    "uri": _text(ref.get("uri"), default=""),
                    "kind": _text(ref.get("type"), default="evidence"),
                }
            )
    for target in _launch_targets(deep_context):
        refs.append({"label": target.get("label", "Context"), "uri": target.get("uri", ""), "kind": target.get("kind", "context")})
    provenance = _promotion_provenance(task_payload)
    for label, key in (
        ("Source packet", "source_packet_uri"),
        ("Source problem", "source_problem_uri"),
        ("Source winner", "source_winner_uri"),
        ("Source option task", "source_option_task_uri"),
    ):
        value = _text(provenance.get(key))
        if value:
            refs.append({"label": label, "uri": value, "kind": "provenance"})
    for ref in _dict_list(provenance.get("refs")) + _dict_list(provenance.get("evidence_refs")):
        refs.append(
            {
                "label": _text(ref.get("label"), default=_text(ref.get("type"), default="Provenance")),
                "uri": _text(ref.get("uri"), default=""),
                "kind": _text(ref.get("type"), default="provenance"),
            }
        )
    for label, key in (("Task folder", "declared_task_root"), ("Worktree", "declared_worktree_root")):
        value = _text(task_payload.get(key))
        if value:
            refs.append({"label": label, "uri": value, "kind": "path"})
    run_artifact_root = _text(_dict(latest_run.get("repo_lane")).get("run_artifact_root"))
    if run_artifact_root:
        refs.append({"label": "Run artifacts", "uri": run_artifact_root, "kind": "path"})
    return refs[:8]


def _promotion_provenance(task_payload: dict[str, Any]) -> dict[str, Any]:
    for key in ("promotion_provenance", "promotion", "provenance"):
        value = task_payload.get(key)
        if isinstance(value, dict):
            return value
    return {}


def _launch_targets(deep_context: dict[str, Any]) -> list[dict[str, object]]:
    targets: list[dict[str, object]] = []
    preferred = _dict(deep_context.get("preferred_launch_target"))
    if preferred:
        targets.append(preferred)
    targets.extend(_dict_list(deep_context.get("launch_targets")))
    transcript_path = _text(deep_context.get("transcript_path"))
    if transcript_path:
        targets.append({"kind": "transcript", "label": "Transcript", "uri": transcript_path})

    deduped: list[dict[str, object]] = []
    seen: set[tuple[str, str]] = set()
    for target in targets:
        uri = _text(target.get("uri"), default="")
        label = _text(target.get("label"), default="Context")
        key = (label, uri)
        if key in seen:
            continue
        seen.add(key)
        command = target.get("command")
        if not isinstance(command, list):
            command = []
        deduped.append(
            {
                "kind": _text(target.get("kind"), default="context"),
                "label": label,
                "uri": uri,
                "command": [str(part) for part in command],
            }
        )
    return deduped


def _launch_action_label(target: dict[str, object]) -> str:
    kind = _text(target.get("kind"), default="context")
    label = _text(target.get("label"), default="")
    combined = f"{kind} {label}".lower()
    if "handoff" in combined:
        return "Open Handoff"
    if "plan" in combined:
        return "Open Plan"
    if "task-state" in combined or "task state" in combined:
        return "Open Task State"
    if "live" in combined and "thread" in combined:
        return "Open Live Thread"
    if "thread" in combined or "session" in combined:
        return "Open Thread"
    if "transcript" in combined:
        return "Open Transcript"
    if "work" in combined or "repo" in combined or "context" in combined:
        return "Open Working Context"
    if "task folder" in combined or "task.md" in combined or "task artifact" in combined:
        return "Open Task"
    return "Open Task"


def _block_reason(availability: dict[str, Any]) -> str:
    if bool(availability.get("allowed")):
        return ""
    reasons = _dict_list(availability.get("block_reasons"))
    if reasons:
        return _text(reasons[0].get("summary"), default=_text(reasons[0].get("code"), default="Blocked"))
    return "Backend says this action is not available now."


def _format_request_os_error(method: str, url: str, exc: OSError) -> str:
    detail = str(exc)
    normalized = detail.lower()
    origin = _request_origin(url)
    if "10061" in detail or "actively refused it" in normalized or "connection refused" in normalized:
        return (
            f"{method} {url} failed: the tasks backend is not reachable at {origin}. "
            f"Start the orchestration backend or set {TASKS_BACKEND_URL_ENV} to a running isolated lane."
        )
    if "timed out" in normalized:
        return f"{method} {url} failed: the tasks backend at {origin} timed out."
    return f"{method} {url} failed: {exc}"


def _join_url(base_url: str, path: str) -> str:
    return base_url.rstrip("/") + path


def _request_origin(url: str) -> str:
    parsed = parse.urlsplit(url)
    if parsed.scheme and parsed.netloc:
        return f"{parsed.scheme}://{parsed.netloc}"
    return url


def _dict(value: Any) -> dict[str, Any]:
    return value if isinstance(value, dict) else {}


def _dict_list(value: Any) -> list[dict[str, Any]]:
    if not isinstance(value, list):
        return []
    return [item for item in value if isinstance(item, dict)]


def _parse_timestamp(raw_value: str) -> datetime | None:
    try:
        return datetime.fromisoformat(raw_value.replace("Z", "+00:00"))
    except ValueError:
        return None


def _text(value: Any, default: str | None = None) -> str | None:
    if value is None:
        return default
    text = str(value).strip()
    if text == "":
        return default
    return text
