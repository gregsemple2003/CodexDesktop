from __future__ import annotations

import os
import unittest
from unittest import mock
from urllib import error

from app.codex_dashboard.jobs_backend import (
    JobsBackendError,
    configured_jobs_backend_url,
    fetch_jobs_snapshot,
    jobs_backend_error_snapshot,
    map_state_view_to_jobs_snapshot,
    trigger_label,
)


class JobsBackendTests(unittest.TestCase):
    def test_configured_jobs_backend_url_uses_environment_override(self) -> None:
        with mock.patch.dict(
            os.environ,
            {"CODEX_DASHBOARD_JOBS_BACKEND_URL": "http://127.0.0.1:14318"},
            clear=False,
        ):
            self.assertEqual(
                configured_jobs_backend_url(),
                "http://127.0.0.1:14318",
            )

    def test_trigger_label_combines_supported_trigger_types(self) -> None:
        self.assertEqual(
            trigger_label(
                [
                    {"type": "schedule"},
                    {"type": "manual"},
                    {"type": "webhook"},
                ]
            ),
            "Schedule + Manual + Webhook",
        )

    def test_map_state_view_to_jobs_snapshot_normalizes_backend_jobs(self) -> None:
        snapshot = map_state_view_to_jobs_snapshot(
            {
                "generated_at": "2026-04-07T15:10:36Z",
                "last_sync": {"last_success_at": "2026-04-07T14:49:45Z"},
                "jobs": [
                    {
                        "job_id": "codex-daily-agentic-swe-digest",
                        "label": "Codex Daily Agentic SWE Digest",
                        "desired_state": "enabled",
                        "status": "in_sync",
                        "triggers": [{"type": "schedule"}, {"type": "manual"}],
                        "schedules": [
                            {
                                "schedule_id": "codex-job--codex-daily-agentic-swe-digest--00",
                                "next_action_times": ["2026-04-08T08:00:00Z"],
                                "recent_runs": [
                                    {
                                        "actual_time": "2026-04-07T08:00:00Z",
                                        "workflow_id": "codex-daily-agentic-swe-digest/schedule/00-2026-04-07T08:00:00Z",
                                    }
                                ],
                            }
                        ],
                        "recent_runs": [],
                    }
                ],
            }
        )

        self.assertEqual(snapshot["last_reconciled_at"], "2026-04-07T14:49:45Z")
        self.assertEqual(snapshot["summary"], {"in_sync": 1})
        self.assertEqual(len(snapshot["jobs"]), 1)
        job = snapshot["jobs"][0]
        self.assertEqual(job["mechanism_label"], "Schedule + Manual")
        self.assertEqual(job["observed_label"], "In sync")
        self.assertTrue(job["supports_run_now"])
        self.assertIn("Next run", job["reason"])
        self.assertIn("Run now available", job["reason"])

    def test_jobs_backend_error_snapshot_marks_backend_blocked(self) -> None:
        snapshot = jobs_backend_error_snapshot("connection refused")

        self.assertEqual(snapshot["summary"], {"blocked": 1})
        self.assertEqual(snapshot["jobs"][0]["status"], "blocked")
        self.assertEqual(snapshot["jobs"][0]["observed_label"], "Blocked")

    def test_fetch_jobs_snapshot_reports_actionable_local_backend_error(self) -> None:
        connection_error = error.URLError(
            OSError(10061, "No connection could be made because the target machine actively refused it")
        )
        with mock.patch(
            "app.codex_dashboard.jobs_backend.request.urlopen",
            side_effect=connection_error,
        ):
            with self.assertRaises(JobsBackendError) as raised:
                fetch_jobs_snapshot("http://127.0.0.1:4318")

        message = str(raised.exception)
        self.assertIn("the jobs backend is not reachable at http://127.0.0.1:4318", message)
        self.assertIn("Start the orchestration service lane", message)
        self.assertIn("CODEX_DASHBOARD_JOBS_BACKEND_URL", message)
