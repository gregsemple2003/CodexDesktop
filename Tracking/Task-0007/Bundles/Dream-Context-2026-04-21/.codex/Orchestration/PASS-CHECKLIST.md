# Pass Checklist Contract

This file defines the shared shape and intended use of `Tracking/Task-<id>/Testing/PASS-<NNNN>-CHECKLIST.json`.

Use it together with:

- `C:\Users\gregs\.codex\AGENTS.md`
- `C:\Users\gregs\.codex\Orchestration\ORCHESTRATION.md`
- `C:\Users\gregs\.codex\Orchestration\PASS-CHECKLIST.schema.json`
- `C:\Users\gregs\.codex\Orchestration\Exemplars\PASS-0002-CHECKLIST.json`

## Purpose

`PASS-<NNNN>-CHECKLIST.json` is the machine-readable closeout state for one implementation pass.

It exists to answer:

- is the pass still in progress, blocked, ready for audit, ready for closeout, or fully closed
- which required closeout steps are already complete
- what is still blocking pass closure
- what the lead agent is expected to do next

It complements, but does not replace:

- `PASS-<NNNN>-AUDIT.md`
- `HANDOFF.md`
- `TASK-STATE.json`

## Canonical Location

Keep the file at:

- `Tracking/Task-<id>/Testing/PASS-<NNNN>-CHECKLIST.json`

## Ownership

The lead agent owns the final contents of `PASS-<NNNN>-CHECKLIST.json`.

Supporting agents may update fields inside their scope when explicitly assigned, but the lead is responsible for keeping the final checklist consistent with the real pass state.

## Update Rules

Update the checklist at durable pass-closeout transitions, for example:

- implementation landed
- unit proof completed
- pass audit written
- auditor verdict recorded
- handoff updated
- commit created
- push completed
- notification sent

Do not treat it like a scratchpad.

## Formatting Rules

- JSON keys: `snake_case`
- enum values: `snake_case`
- booleans: real JSON booleans
- no copied logs or prose explanations

## Required Fields

| Field | Type | Meaning |
| --- | --- | --- |
| `schema_version` | integer | Current schema version. Start at `1`. |
| `task_id` | string | Canonical task id, e.g. `Task-0002`. |
| `pass_id` | string | Canonical pass id, e.g. `PASS-0002`. |
| `checklist_status` | enum | Current closeout readiness for the pass. |
| `implemented` | boolean | Whether the intended pass implementation is present. |
| `build` | enum | Current build or compile status for the pass. |
| `unit_tests` | enum | Current pass-local unit-test status. |
| `pass_audit_written` | boolean | Whether the markdown pass audit exists and is current. |
| `audit_verdict` | enum | Latest auditor verdict for the pass. |
| `audit_result_ref` | string or `null` | Relative path to `PASS-<NNNN>-AUDIT.json` when used. |
| `handoff_updated` | boolean | Whether `HANDOFF.md` has been updated for the new baseline. |
| `commit_created` | boolean | Whether the pass closeout commit exists. |
| `pushed_upstream` | boolean | Whether that commit has been pushed. |
| `notify_sent` | boolean | Whether the final notification has been sent. |
| `blockers` | string array | Current blockers preventing full closeout. |
| `next_required_actions` | string array | The next remaining closeout steps. |
| `updated_at` | RFC 3339 timestamp string | Last checklist update time. |

## Enums

### `checklist_status`

- `in_progress`
- `ready_for_audit`
- `ready_for_closeout`
- `blocked`
- `closed`

### `build`

Use the shared validation-style values:

- `unknown`
- `pending`
- `passing`
- `failing`
- `not_run`
- `blocked`
- `not_applicable`

### `unit_tests`

Use the same enum values as `build`.

### `audit_verdict`

- `unknown`
- `ready`
- `ready_with_caveats`
- `not_ready`

## Relationship To Markdown Artifacts

Use a strict split:

- `PASS-<NNNN>-AUDIT.md` explains the proof and evidence
- `PASS-<NNNN>-CHECKLIST.json` records whether the pass-closeout gates are actually satisfied

Example split:

- the pass audit may already be written and the unit tests may already pass
- the checklist can still be `ready_for_closeout` if commit, push, or notify have not happened yet

## Validation

Validate the file against:

- `C:\Users\gregs\.codex\Orchestration\PASS-CHECKLIST.schema.json`
