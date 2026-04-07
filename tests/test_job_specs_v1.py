from __future__ import annotations

import json
import tempfile
import unittest
from pathlib import Path

from app.codex_dashboard.job_specs import (
    API_VERSION,
    DESIRED_STATE_ENABLED,
    EXECUTOR_TYPE_CODEX_EXEC,
    TRIGGER_TYPE_MANUAL,
    TRIGGER_TYPE_SCHEDULE,
    TRIGGER_TYPE_WEBHOOK,
    iter_job_spec_paths,
    load_validated_job_specs,
    validate_job_spec,
)
from app.codex_dashboard.paths import job_spec_schema_path, job_specs_root


class JobSpecsV1Tests(unittest.TestCase):
    def setUp(self) -> None:
        self.temp_dir = tempfile.TemporaryDirectory()
        self.codex_root = Path(self.temp_dir.name) / ".codex"
        self.specs_root = job_specs_root(self.codex_root)
        self.specs_root.mkdir(parents=True, exist_ok=True)

    def tearDown(self) -> None:
        self.temp_dir.cleanup()

    def write_spec(self, name: str, payload: dict) -> Path:
        path = self.specs_root / name
        path.write_text(json.dumps(payload, indent=2) + "\n", encoding="utf-8")
        return path

    def sample_spec(self, job_id: str, trigger_type: str) -> dict:
        trigger: dict[str, object]
        if trigger_type == TRIGGER_TYPE_SCHEDULE:
            trigger = {"type": trigger_type, "cron": "0 4 * * *", "timezone": "America/Toronto"}
        elif trigger_type == TRIGGER_TYPE_WEBHOOK:
            trigger = {"type": trigger_type, "path": "digests/example"}
        else:
            trigger = {"type": trigger_type}
        return {
            "api_version": API_VERSION,
            "job_id": job_id,
            "label": "Example job",
            "description": "Example job for unit validation.",
            "desired_state": DESIRED_STATE_ENABLED,
            "triggers": [trigger],
            "executor": {
                "type": EXECUTOR_TYPE_CODEX_EXEC,
                "cwd": "C:\\Users\\gregs\\.codex",
                "entrypoint": "example-skill",
                "args": ["--days", "1"],
            },
            "runtime": {
                "workflow_type": "codex.exec.job",
                "task_queue": "codex-orchestration",
            },
        }

    def test_job_spec_paths_use_orchestration_jobs_specs_subtree(self) -> None:
        self.assertEqual(
            job_specs_root(self.codex_root),
            self.codex_root / "Orchestration" / "Jobs" / "specs",
        )
        self.assertEqual(
            job_spec_schema_path(self.codex_root),
            self.codex_root / "Orchestration" / "Jobs" / "job-spec.schema.json",
        )

    def test_validate_job_spec_accepts_each_supported_trigger_type(self) -> None:
        for index, trigger_type in enumerate(
            [TRIGGER_TYPE_SCHEDULE, TRIGGER_TYPE_MANUAL, TRIGGER_TYPE_WEBHOOK],
            start=1,
        ):
            spec = self.sample_spec(f"job-{index}", trigger_type)
            validate_job_spec(spec)

    def test_load_validated_job_specs_rejects_duplicate_job_ids(self) -> None:
        self.write_spec("one.json", self.sample_spec("duplicate-job", TRIGGER_TYPE_MANUAL))
        self.write_spec("two.json", self.sample_spec("duplicate-job", TRIGGER_TYPE_SCHEDULE))

        with self.assertRaisesRegex(ValueError, "Duplicate job_id"):
            load_validated_job_specs(self.codex_root)

    def test_iter_job_spec_paths_ignores_schema_files(self) -> None:
        self.write_spec("example.json", self.sample_spec("example-job", TRIGGER_TYPE_MANUAL))
        schema_path = self.specs_root / "ignored.schema.json"
        schema_path.write_text("{}", encoding="utf-8")

        paths = iter_job_spec_paths(self.codex_root)

        self.assertEqual(paths, [self.specs_root / "example.json"])
