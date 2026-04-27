from __future__ import annotations

import json
import os
import tempfile
import unittest
from unittest import mock

from app.codex_dashboard.tasks_backend import (
    TASKS_SNAPSHOT_PATH_ENV,
    configured_tasks_backend_url,
    fetch_tasks_snapshot,
    map_backend_tasks_snapshot,
    provenance_label,
)


class TasksBackendTests(unittest.TestCase):
    def test_configured_backend_url_uses_tasks_env_override(self) -> None:
        with mock.patch.dict(os.environ, {"CODEX_DASHBOARD_TASKS_BACKEND_URL": "http://127.0.0.1:14318"}):
            self.assertEqual(configured_tasks_backend_url(), "http://127.0.0.1:14318")

    def test_maps_committed_tasks_without_unpromoted_candidates_or_progress(self) -> None:
        snapshot = map_backend_tasks_snapshot(
            {
                "generated_at": "2026-04-26T21:30:00Z",
                "tasks": [
                    {
                        "task_id": "review-candidate-1",
                        "title": "Draft suggestion",
                        "kind": "candidate",
                        "status": "candidate",
                    },
                    {
                        "task_id": "Task-0009",
                        "title": "Build Tasks tab",
                        "meaning_summary": "Make committed work visible.",
                        "promoted_from": "review",
                        "latest_run": {"run_id": "taskrun--Task-0009--active"},
                        "state_envelope": {
                            "state": "waiting_for_human",
                            "state_summary": "Plan approval is needed.",
                            "next_expected_event": "Approve or redirect the plan.",
                        },
                        "attention": {"attention_level": "needs_attention"},
                        "actions": {
                            "interrupt": {"allowed": True},
                            "dispatch": {"allowed": False, "block_reasons": [{"summary": "Already running."}]},
                        },
                    },
                ],
            }
        )

        tasks = snapshot["tasks"]
        self.assertEqual(len(tasks), 1)
        task = tasks[0]
        self.assertEqual(task["task_id"], "Task-0009")
        self.assertEqual(task["provenance_label"], "Promoted from Review")
        self.assertEqual(task["summary_bucket"], "needs_you")
        self.assertEqual(task["latest_run_id"], "taskrun--Task-0009--active")
        self.assertEqual(snapshot["summary"]["needs_you"], 1)
        labels = [action["label"] for action in task["actions"]]
        self.assertIn("Pause", labels)
        self.assertNotIn("Interrupt", labels)
        pause_action = next(action for action in task["actions"] if action["label"] == "Pause")
        self.assertEqual(pause_action["backend_action"], "interrupt")
        self.assertEqual(pause_action["run_id"], "taskrun--Task-0009--active")
        self.assertNotIn("progress", json.dumps(task).lower())

    def test_promoted_candidate_source_does_not_display_candidate_label(self) -> None:
        self.assertEqual(
            provenance_label({"task_id": "Task-0012", "source": "candidate", "promoted": True}),
            "Promoted",
        )

    def test_fixture_path_loads_snapshot_without_http(self) -> None:
        with tempfile.TemporaryDirectory() as tmpdir:
            path = os.path.join(tmpdir, "tasks.json")
            with open(path, "w", encoding="utf-8") as handle:
                json.dump({"tasks": [{"task_id": "Task-0001", "title": "Fixture task"}]}, handle)
            with mock.patch.dict(os.environ, {TASKS_SNAPSHOT_PATH_ENV: path}):
                snapshot = fetch_tasks_snapshot("http://127.0.0.1:1")

        self.assertEqual(snapshot["tasks"][0]["task_id"], "Task-0001")


if __name__ == "__main__":
    unittest.main()
