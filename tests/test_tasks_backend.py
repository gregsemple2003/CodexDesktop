from __future__ import annotations

import json
import os
import tempfile
import unittest
from pathlib import Path
from unittest import mock

from app.codex_dashboard.tasks_backend import (
    TASKS_SNAPSHOT_PATH_ENV,
    configured_tasks_backend_url,
    fetch_tasks_snapshot,
    map_backend_tasks_snapshot,
    provenance_detail,
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

    def test_terminal_tasks_are_filtered_from_active_tasks_surface(self) -> None:
        snapshot = map_backend_tasks_snapshot(
            {
                "tasks": [
                    {
                        "task_id": "Task-0001",
                        "title": "Done task",
                        "state_envelope": {"state": "completed"},
                    },
                    {
                        "task_id": "Task-0004",
                        "title": "Cancelled task",
                        "status": "cancelled",
                    },
                    {
                        "task_id": "Task-0010",
                        "title": "Needs plan approval",
                        "state_envelope": {"state": "waiting_for_human"},
                    },
                ]
            }
        )

        self.assertEqual([task["task_id"] for task in snapshot["tasks"]], ["Task-0010"])
        self.assertEqual(snapshot["summary"]["needs_you"], 1)
        self.assertEqual(snapshot["summary"]["running"], 0)
        self.assertEqual(snapshot["summary"]["blocked"], 0)
        self.assertNotIn("Task-0001", json.dumps(snapshot))
        self.assertNotIn("Task-0004", json.dumps(snapshot))

    def test_terminal_top_level_state_overrides_stale_latest_run(self) -> None:
        snapshot = map_backend_tasks_snapshot(
            {
                "tasks": [
                    {
                        "task_id": "Task-0008",
                        "title": "Completed task with stale run",
                        "state_envelope": {"state": "completed"},
                        "latest_run": {
                            "run_id": "stale-run",
                            "state_envelope": {"state": "running"},
                            "actions": {"interrupt": {"allowed": True}},
                        },
                    }
                ]
            }
        )

        self.assertEqual(snapshot["tasks"], [])
        self.assertEqual(snapshot["summary"]["running"], 0)

    def test_blocked_backend_owned_task_is_not_waiting_on_you(self) -> None:
        snapshot = map_backend_tasks_snapshot(
            {
                "tasks": [
                    {
                        "task_id": "Task-0002",
                        "title": "Backend review",
                        "state_envelope": {
                            "state": "blocked",
                            "state_summary": "Task needs backend review before dispatch.",
                            "next_owner": "backend",
                        },
                        "attention": {"attention_level": "needs_attention"},
                    }
                ]
            }
        )

        task = snapshot["tasks"][0]
        self.assertEqual(task["state_label"], "Blocked")
        self.assertEqual(task["summary_bucket"], "blocked")
        self.assertEqual(snapshot["summary"]["needs_you"], 0)
        self.assertEqual(snapshot["summary"]["blocked"], 1)

    def test_no_active_run_hides_run_control_actions(self) -> None:
        snapshot = map_backend_tasks_snapshot(
            {
                "tasks": [
                    {
                        "task_id": "Task-0002",
                        "title": "No active run",
                        "state_envelope": {"state": "blocked"},
                        "actions": {
                            "dispatch": {"allowed": False, "block_reasons": [{"summary": "Blocked."}]},
                            "interrupt": {"allowed": False},
                            "poke": {"allowed": False},
                        },
                    }
                ]
            }
        )

        labels = [action["label"] for action in snapshot["tasks"][0]["actions"]]
        self.assertIn("Dispatch", labels)
        self.assertNotIn("Pause", labels)
        self.assertNotIn("Poke", labels)

    def test_task_artifact_launch_targets_get_distinct_action_labels(self) -> None:
        snapshot = map_backend_tasks_snapshot(
            {
                "tasks": [
                    {
                        "task_id": "Task-0011",
                        "title": "Planning task",
                        "state_envelope": {"state": "waiting_for_human"},
                        "actions": {"dispatch": {"allowed": False}},
                        "deep_context": {
                            "preferred_launch_target": {
                                "kind": "task_artifact",
                                "label": "Task folder",
                                "uri": "file:///C:/Agent/CodexDashboard/Tracking/Task-0011",
                                "command": ["code", "C:\\Agent\\CodexDashboard\\Tracking\\Task-0011"],
                            },
                            "launch_targets": [
                                {
                                    "kind": "task_artifact",
                                    "label": "Task handoff",
                                    "uri": "file:///C:/Agent/CodexDashboard/Tracking/Task-0011/HANDOFF.md",
                                    "command": ["code", "C:\\Agent\\CodexDashboard\\Tracking\\Task-0011\\HANDOFF.md"],
                                },
                                {
                                    "kind": "task_artifact",
                                    "label": "Task plan",
                                    "uri": "file:///C:/Agent/CodexDashboard/Tracking/Task-0011/PLAN.md",
                                    "command": ["code", "C:\\Agent\\CodexDashboard\\Tracking\\Task-0011\\PLAN.md"],
                                },
                            ],
                        },
                    }
                ]
            }
        )

        labels = [action["label"] for action in snapshot["tasks"][0]["actions"]]
        self.assertEqual(labels.count("Open Task"), 1)
        self.assertIn("Open Handoff", labels)
        self.assertIn("Open Plan", labels)

    def test_promotion_provenance_maps_dream_source_and_refs(self) -> None:
        snapshot = map_backend_tasks_snapshot(
            {
                "tasks": [
                    {
                        "task_id": "Task-0012",
                        "title": "Promoted Dream task",
                        "meaning_summary": "A promoted option task.",
                        "promotion_provenance": {
                            "source": "dream",
                            "source_packet": "2026-04-26 Dream packet",
                            "source_problem": "Problem 0003",
                            "source_winner": "Winner B",
                            "source_option_task": "Option task 2",
                            "source_packet_uri": "C:/reports/dream.md",
                        },
                    }
                ]
            }
        )

        task = snapshot["tasks"][0]
        self.assertEqual(task["provenance_label"], "Promoted from Dream")
        self.assertIn("Source packet: 2026-04-26 Dream packet", task["provenance_detail"])
        self.assertIn(
            {"label": "Source packet", "uri": "C:/reports/dream.md", "kind": "provenance"},
            task["artifacts"],
        )

    def test_unpromoted_review_candidate_is_filtered_even_with_provenance_object(self) -> None:
        snapshot = map_backend_tasks_snapshot(
            {
                "tasks": [
                    {
                        "task_id": "review-42",
                        "title": "Pending review",
                        "kind": "task_candidate",
                        "promotion_provenance": {"source": "review", "promotion_status": "candidate"},
                    }
                ]
            }
        )

        self.assertEqual(snapshot["tasks"], [])

    def test_provenance_detail_falls_back_to_label(self) -> None:
        self.assertEqual(provenance_detail({"task_id": "Task-0001", "source": "authored"}), "Authored")

    def test_fixture_path_loads_snapshot_without_http(self) -> None:
        with tempfile.TemporaryDirectory() as tmpdir:
            path = os.path.join(tmpdir, "tasks.json")
            with open(path, "w", encoding="utf-8") as handle:
                json.dump({"tasks": [{"task_id": "Task-0001", "title": "Fixture task"}]}, handle)
            with mock.patch.dict(os.environ, {TASKS_SNAPSHOT_PATH_ENV: path}):
                snapshot = fetch_tasks_snapshot("http://127.0.0.1:1")

        self.assertEqual(snapshot["tasks"][0]["task_id"], "Task-0001")

    def test_tasks_product_surface_fixture_covers_task_8_and_9_rules(self) -> None:
        fixture_path = Path(__file__).parent / "fixtures" / "tasks_product_surface.json"
        with mock.patch.dict(os.environ, {TASKS_SNAPSHOT_PATH_ENV: str(fixture_path)}):
            snapshot = fetch_tasks_snapshot("http://127.0.0.1:1")

        tasks = {task["task_id"]: task for task in snapshot["tasks"]}
        self.assertNotIn("review-candidate-1", tasks)
        self.assertIn("Task-0008", tasks)
        self.assertNotIn("Task-0009", tasks)

        task8 = tasks["Task-0008"]
        task8_labels = [action["label"] for action in task8["actions"]]
        self.assertIn("Pause", task8_labels)
        self.assertIn("Open Live Thread", task8_labels)
        self.assertIn("Open Working Context", task8_labels)
        self.assertNotIn("Interrupt", task8_labels)
        self.assertEqual(task8["summary_bucket"], "running")
        self.assertEqual(task8["provenance_label"], "Authored")

        provenance_labels = {task["task_id"]: task["provenance_label"] for task in snapshot["tasks"]}
        self.assertEqual(provenance_labels["Task-0012"], "Promoted from Dream")
        for label in provenance_labels.values():
            self.assertNotIn("Candidate", label)
        self.assertNotIn("Prov: Candidate", json.dumps(snapshot))
        self.assertNotIn("progress_bar", json.dumps(snapshot).lower())


if __name__ == "__main__":
    unittest.main()
