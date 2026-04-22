# Task State Contract

This file defines the shared shape and intended use of `Tracking/Task-<id>/TASK-STATE.json`.

Use it together with:

- `C:\Users\gregs\.codex\AGENTS.md`
- `C:\Users\gregs\.codex\Orchestration\ORCHESTRATION.md`
- `C:\Users\gregs\.codex\Orchestration\TASK-STATE.schema.json`
- `C:\Users\gregs\.codex\Orchestration\Exemplars\TASK-STATE.json`

## Purpose

`TASK-STATE.json` is the machine-readable current state of a task.

It exists to give agents, auditors, and a future orchestration framework a compact source of truth for:

- where the task is in the lifecycle
- which pass is active
- which gate is currently in front of the task
- what the latest validation and audit outcomes are
- what artifacts are expected next

It does **not** replace the markdown artifacts.

Keep using:

- `TASK.md` for scope and acceptance
- `PLAN.md` for pass sequencing
- `HANDOFF.md` for human-readable current baseline
- audits, bug notes, and regression runs for evidence and history

`TASK-STATE.json` should only capture the current orchestration state.

## Canonical Location

When a task uses structured orchestration state, keep the file at:

- `Tracking/Task-<id>/TASK-STATE.json`

## Ownership

The lead agent owns the final contents of `TASK-STATE.json`.

Supporting agents may recommend updates, but they should not become the final authority on task state unless explicitly assigned.

## Update Rules

Update `TASK-STATE.json` only at durable state transitions, for example:

- task created
- plan approved
- pass started
- gate changed from implementation to unit testing
- audit verdict recorded
- task blocked
- task completed or cancelled

Do not rewrite the file for every intermediate thought or experiment.

When the task state changes materially, update `TASK-STATE.json` in the same change as the related task artifact updates when practical.

Leader checkpoint rule:

- when a leader-owned workflow persists a durable `TASK-STATE.json` transition that changes `phase`, `current_pass`, or `current_gate`, treat that update as a checkpoint
- include the related task-artifact updates in the same checkpoint when practical
- commit that checkpoint and push it upstream before proceeding to the next workflow stage
- this rule belongs to leader-owned workflows such as research and implementation leaders, not to supporting non-leader agents by default
- apply this rule to the real persisted checkpoint that moves the task forward, not to synthetic bookkeeping edits

## Formatting Rules

- JSON keys: `snake_case`
- enum values: `snake_case`
- booleans: real JSON booleans
- no narrative prose blocks
- no copied test logs
- no duplicated markdown history

## Required Fields

| Field | Type | Meaning |
| --- | --- | --- |
| `schema_version` | integer | Current schema version. Start at `1`. |
| `task_id` | string | Canonical task id, e.g. `Task-0002`. |
| `status` | enum | High-level task status. |
| `phase` | enum | Current task-level lifecycle phase. |
| `plan_approved` | boolean | Whether the current `PLAN.md` is approved for execution. |
| `current_pass` | string or `null` | The pass currently being worked, e.g. `PASS-0002`. |
| `last_completed_pass` | string or `null` | The most recently completed pass. |
| `current_gate` | enum | The current fine-grained gate or subphase in front of the task. |
| `latest_audit_verdict` | enum | The most recent auditor verdict. |
| `validation` | object | Compact validation state for build, unit tests, and regression. |
| `blockers` | string array | Current blocking issues that prevent progress. |
| `next_expected_artifacts` | string array | Relative task paths expected next. |
| `updated_at` | RFC 3339 timestamp string | Last state update time. |

## Enums

### `status`

- `pending`
- `in_progress`
- `blocked`
- `complete`
- `cancelled`

### `phase`

`phase` is coarse task lifecycle state, not pass-local subphase state.

Use it to answer "what part of the task lifecycle are we in overall?"

- `creation`
- `research`
- `planning`
- `implementation`
- `regression`
- `closure`

### `current_gate`

`current_gate` is the fine-grained step inside the current phase.

Use it to answer "what is the task waiting on right now?"

- `none`
- `research`
- `planning`
- `implementation`
- `unit_testing`
- `audit`
- `handoff`
- `regression`
- `debugging`
- `closure`

### `latest_audit_verdict`

- `unknown`
- `ready`
- `ready_with_caveats`
- `not_ready`

### `validation.*`

Use the same enum values for `build`, `unit_tests`, and `regression`:

- `unknown`
- `pending`
- `passing`
- `failing`
- `not_run`
- `blocked`
- `not_applicable`

## Relationship To Markdown Artifacts

Use a strict split:

- markdown files hold narrative, rationale, evidence, and history
- `TASK-STATE.json` holds current machine-readable orchestration state only

Example split:

- `phase: "implementation"` with `current_gate: "unit_testing"` means the task is still in implementation overall, but the current pass is in its unit-proof subphase
- `phase: "regression"` with `current_gate: "debugging"` means the task-level regression phase is active and the current problem is being debugged before the next rerun

If the JSON and markdown ever disagree, fix the inconsistency in the same follow-up change instead of letting both drift.

## Validation

Validate `TASK-STATE.json` against:

- `C:\Users\gregs\.codex\Orchestration\TASK-STATE.schema.json`

Use the exemplar at:

- `C:\Users\gregs\.codex\Orchestration\Exemplars\TASK-STATE.json`

when you need a concrete starting point.
