# Audit Result Contract

This file defines the shared shape and intended use of `Tracking/Task-<id>/Testing/PASS-<NNNN>-AUDIT.json`.

Use it together with:

- `C:\Users\gregs\.codex\AGENTS.md`
- `C:\Users\gregs\.codex\Orchestration\ORCHESTRATION.md`
- `C:\Users\gregs\.codex\Orchestration\AUDIT-RESULT.schema.json`
- `C:\Users\gregs\.codex\Orchestration\Exemplars\PASS-0002-AUDIT.json`
- `C:\Users\gregs\.codex\Orchestration\Prompts\AUDITOR.md`

## Purpose

`PASS-<NNNN>-AUDIT.json` is the machine-readable result of an audit gate.

It exists to capture:

- what transition was being audited
- the auditor verdict
- blocking and non-blocking findings
- missing evidence
- what evidence the auditor actually reviewed
- the suggested next step

It is intentionally smaller and more structured than a markdown narrative.

## Canonical Location

Keep the file at:

- `Tracking/Task-<id>/Testing/PASS-<NNNN>-AUDIT.json`

## Ownership

The lead agent owns the final persisted file.

The auditor provides the structured verdict content, but the auditor may remain read-only in the workflow. In that case the lead agent materializes the result faithfully.

## Update Rules

Write or replace the file when a durable audit verdict is available.

If a later rerun produces a new verdict for the same pass, update that pass's canonical `PASS-<NNNN>-AUDIT.json` instead of minting a second pass artifact.

Use git history to preserve the chronology of reruns and verdict changes.

## Formatting Rules

- JSON keys: `snake_case`
- enum values: `snake_case`
- findings should stay short and structured
- keep the file factual, not narrative

## Required Fields

| Field | Type | Meaning |
| --- | --- | --- |
| `schema_version` | integer | Current schema version. Start at `1`. |
| `task_id` | string | Canonical task id, e.g. `Task-0002`. |
| `pass_id` | string | Canonical pass id, e.g. `PASS-0002`. |
| `transition_target` | enum | The state transition being audited. |
| `verdict` | enum | Final audit verdict. |
| `blocking_findings` | finding array | Findings that prevent the transition. |
| `non_blocking_findings` | finding array | Findings that do not block the transition. |
| `missing_evidence` | string array | Evidence the auditor expected but did not have. |
| `open_questions` | string array | Open questions that remain after the audit. |
| `evidence_reviewed` | string array | Concrete docs, diffs, or results the auditor reviewed. |
| `suggested_next_step` | string | The next recommended action after the verdict. |
| `updated_at` | RFC 3339 timestamp string | Last audit-result update time. |

## Enums

### `transition_target`

- `pass_closeout`
- `commit`
- `push`
- `handoff`
- `task_closure`

### `verdict`

- `ready`
- `ready_with_caveats`
- `not_ready`

## Finding Object

Each finding should include:

- `summary`
- `why_it_matters`
- `file_refs`

Keep findings short enough that they can be consumed programmatically or shown in a compact review UI.

## Relationship To Markdown Artifacts

Use a strict split:

- markdown audits and handoffs explain the story and evidence in full
- `PASS-<NNNN>-AUDIT.json` records the structured verdict and findings

## Validation

Validate the file against:

- `C:\Users\gregs\.codex\Orchestration\AUDIT-RESULT.schema.json`
