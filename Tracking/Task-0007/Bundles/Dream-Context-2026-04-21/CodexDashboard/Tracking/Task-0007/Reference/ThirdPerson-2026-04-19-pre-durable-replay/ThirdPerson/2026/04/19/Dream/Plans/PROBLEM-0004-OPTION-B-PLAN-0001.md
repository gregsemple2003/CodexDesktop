# Problem 0004 Option B Plan 0001

## Planning Intent

This file turns Problem `0004`, Option `B. Evidence manifest` from [ORTHOGONAL-SOLUTIONS-MATRIX.md](../ORTHOGONAL-SOLUTIONS-MATRIX.md) into a bounded implementation sequence.

It is an alternative route, not the selected winner task.

## Summary

Add one evidence manifest contract so every proof bundle names the claim, lane, region of interest, runtime status, and allowed use in machine-checkable form.

## Fixed Defaults

- scope: shared orchestration contract plus task-owned evidence artifacts
- canonical homes:
  - `C:\Users\gregs\.codex\Orchestration\EVIDENCE-MANIFEST.md`
  - `C:\Users\gregs\.codex\Orchestration\EVIDENCE-MANIFEST.schema.json`
  - `C:\Users\gregs\.codex\Orchestration\FILE-NAMING.md`
  - `C:\Users\gregs\.codex\Orchestration\Processes\TESTING.md`
  - `C:\Users\gregs\.codex\Orchestration\Processes\DEBUGGING.md`
  - `C:\Users\gregs\.codex\Orchestration\Prompts\REGRESSION-LEADER.md`
  - `C:\Users\gregs\.codex\Orchestration\Prompts\REGRESSION-TESTER.md`
  - `C:\Users\gregs\.codex\Orchestration\Prompts\DEBUG-LEADER.md`
- task-owned manifest name:
  - `Tracking/Task-<id>/Testing/EVIDENCE-MANIFEST-<NNNN>.json`
- required fields:
  - `claim_ref`
  - `capture_surface`
  - `repo_lane`
  - `region_of_interest`
  - `runtime_visible`
  - `allowed_use`
  - `artifact_refs`

## Pass Plan

### Pass 0000 - Manifest Contract And Naming

Goal:

- define the evidence manifest once and make it machine-checkable

Build:

- add `EVIDENCE-MANIFEST.md`
- add `EVIDENCE-MANIFEST.schema.json`
- update `FILE-NAMING.md` with `EVIDENCE-MANIFEST-<NNNN>.json`

Unit Proof:

- the doc and schema use the same field names
- the naming rule matches the artifact name used in the contract

Exit Bar:

- evidence bundles can carry explicit proof metadata without local invention

### Pass 0001 - Workflow Adoption

Goal:

- make the manifest part of testing and debugging proof flow

Build:

- update `Processes/TESTING.md` and `Processes/DEBUGGING.md` to require an evidence manifest for proof bundles
- update `REGRESSION-LEADER.md`, `REGRESSION-TESTER.md`, and `DEBUG-LEADER.md` to require the manifest before closure or root-cause language

Unit Proof:

- the process docs and prompts refer to the same manifest
- `allowed_use` meaning is consistent across all touched docs

Exit Bar:

- later proof review becomes clearer because every bundle says what it proves and how far it counts

### Pass 0002 - Examples And Eval Hooks

Goal:

- make the manifest easy to apply and easy to audit

Build:

- add one example runtime-closure manifest
- add one example supporting-only manifest
- preserve one April 19 packet-backed eval where the manifest would have exposed invalid proof early

Unit Proof:

- the example manifests validate against the schema
- the eval note maps cleanly to the packet failure

Exit Bar:

- a future reviewer can inspect one manifest and know what the evidence does and does not prove

## Testing Strategy

- validate that example manifests only use schema-approved fields and values
- reject rollout if any touched prompt uses a different field name or `allowed_use` meaning

## Deferred Work

- automated evidence linting
- image-region detection
- dashboard proof review tooling
