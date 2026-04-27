from __future__ import annotations

from collections import OrderedDict
from typing import Iterable


TASK_SUMMARY_CARDS = (
    ("needs_you", "NEEDS YOU"),
    ("sleeping", "SLEEPING"),
    ("running", "RUNNING"),
    ("blocked", "BLOCKED"),
    ("ready", "READY"),
)

TASK_GROUP_ORDER = (
    "Needs Attention",
    "Running",
    "Ready to Dispatch",
    "Waiting or Blocked",
    "Sleeping / Stalled",
)

TASK_STATE_COLORS = {
    "needs_you": "#ffb45c",
    "running": "#16d9f5",
    "ready": "#5ee69a",
    "blocked": "#ff5a52",
    "sleeping": "#b78cff",
}


def group_tasks_for_stream(tasks: Iterable[dict[str, object]]) -> "OrderedDict[str, list[dict[str, object]]]":
    grouped: "OrderedDict[str, list[dict[str, object]]]" = OrderedDict((group, []) for group in TASK_GROUP_ORDER)
    for task in tasks:
        group = str(task.get("stream_group") or "Ready to Dispatch")
        grouped.setdefault(group, []).append(task)
    return OrderedDict((group, rows) for group, rows in grouped.items() if rows)


def task_state_color(task: dict[str, object]) -> str:
    return TASK_STATE_COLORS.get(str(task.get("summary_bucket") or "ready"), "#8fa8bb")


def task_detail_sections(task: dict[str, object]) -> list[tuple[str, str]]:
    return [
        ("Summary", str(task.get("meaning_summary") or "")),
        ("Why this task exists", str(task.get("why") or "")),
        ("Current state", f"{task.get('state_label', 'Unknown')} - {task.get('reason', '')}"),
        ("What changed recently", str(task.get("recent_change") or "")),
        ("Next expected step", str(task.get("next_expected_event") or "")),
    ]


def first_task_id(tasks: Iterable[dict[str, object]]) -> str | None:
    for task in tasks:
        task_id = str(task.get("task_id") or "").strip()
        if task_id:
            return task_id
    return None
