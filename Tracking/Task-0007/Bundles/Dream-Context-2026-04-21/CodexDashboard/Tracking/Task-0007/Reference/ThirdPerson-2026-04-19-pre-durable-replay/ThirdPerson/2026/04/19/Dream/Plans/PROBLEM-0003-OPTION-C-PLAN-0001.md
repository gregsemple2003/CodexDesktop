# Problem 0003 Option C Plan 0001

## Planning Intent

This file turns Problem `0003`, Option `C. Dedicated review UI` from [ORTHOGONAL-SOLUTIONS-MATRIX.md](../ORTHOGONAL-SOLUTIONS-MATRIX.md) into a bounded implementation sequence.

It is an alternative route, not the selected winner task.

## Summary

Build a small read-only review surface in CodexDashboard that renders approval packets with direct file links, change summaries, and current pass context.

## Fixed Defaults

- scope: CodexDashboard product code plus approval packet artifacts
- canonical homes:
  - `C:\Agent\CodexDashboard\app\codex_dashboard\ui.py`
  - `C:\Agent\CodexDashboard\app\codex_dashboard\models.py`
  - `C:\Agent\CodexDashboard\app\codex_dashboard\storage.py`
  - `C:\Agent\CodexDashboard\tests\`
- packet substrate:
  - `Tracking/Task-<id>/APPROVAL-PACKET-PLAN.md`
  - `Tracking/Task-<id>/APPROVAL-PACKET-PASS-<NNNN>.md`
- UI scope for this rollout:
  - list packets
  - show changed artifact links
  - show old vs new summary
  - open linked files externally
- no inline editing or approval write-back in this rollout

## Pass Plan

### Pass 0000 - Packet Reader And Model

Goal:

- make approval packets parseable by the desktop app

Build:

- add a packet reader in the app layer
- add a small in-memory model for packet title, links, change summaries, and gate status
- add focused tests for packet parsing

Unit Proof:

- packet parser tests pass on plan and pass packet fixtures
- malformed packets fail clearly instead of rendering partial garbage

Exit Bar:

- the app can read approval packets without UI code needing to parse markdown ad hoc

### Pass 0001 - Review Surface

Goal:

- render one usable review pane for the parsed packet data

Build:

- add a packet list and detail pane to `ui.py`
- show current task, current pass or gate, changed artifact links, and old vs new summary
- add open-file actions for linked artifacts

Unit Proof:

- touched UI tests remain green if present
- one manual desktop smoke proves packet list and detail rendering on real packet fixtures

Exit Bar:

- a reviewer can inspect an approval packet without hunting through raw markdown first

### Pass 0002 - Operator Notes And Smoke Coverage

Goal:

- make the surface usable by someone other than the implementer

Build:

- add one short operator note to the repo README or task handoff
- save one manual smoke artifact that shows the review pane on a real packet

Unit Proof:

- the operator note matches the current UI entrypoint
- the smoke artifact uses real packet input

Exit Bar:

- the review surface is discoverable and honestly exercised

## Testing Strategy

- use parser tests for packet-model correctness
- use one manual UI smoke for the human-facing review flow

## Deferred Work

- inline approval actions
- threaded comments
- diff rendering beyond packet-provided summaries
