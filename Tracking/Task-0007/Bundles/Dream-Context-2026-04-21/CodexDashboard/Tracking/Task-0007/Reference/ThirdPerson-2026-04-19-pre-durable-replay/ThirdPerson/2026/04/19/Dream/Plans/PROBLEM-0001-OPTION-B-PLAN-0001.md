# Problem 0001 Option B Plan 0001

## Planning Intent

This file turns Problem `0001`, Option `B. Structured claim manifest` from [ORTHOGONAL-SOLUTIONS-MATRIX.md](../ORTHOGONAL-SOLUTIONS-MATRIX.md) into a bounded implementation sequence.

It is an alternative route, not the selected winner task.

## Summary

Add one explicit claim manifest so every closure or regression claim names its lane, proof type, runtime status, and closure scope in machine-checkable form.

## Fixed Defaults

- scope: shared orchestration contract plus task-owned claim artifacts
- canonical homes:
  - `C:\Users\gregs\.codex\Orchestration\CLAIM-MANIFEST.md`
  - `C:\Users\gregs\.codex\Orchestration\CLAIM-MANIFEST.schema.json`
  - `C:\Users\gregs\.codex\Orchestration\FILE-NAMING.md`
  - `C:\Users\gregs\.codex\Orchestration\Processes\TESTING.md`
  - `C:\Users\gregs\.codex\Orchestration\Prompts\REGRESSION-LEADER.md`
- task-owned manifest name:
  - `Tracking/Task-<id>/Testing/CLAIM-MANIFEST-<NNNN>.json`
- this option adds explicit metadata but does not block claims automatically

## Pass Plan

### Pass 0000 - Manifest Contract

Goal:

- define the claim manifest fields and file name once

Build:

- add `CLAIM-MANIFEST.md` with required fields for lane, evidence type, runtime status, artifact refs, and closure scope
- add `CLAIM-MANIFEST.schema.json` with matching keys and enum values
- update `FILE-NAMING.md` with `CLAIM-MANIFEST-<NNNN>.json`

Unit Proof:

- the prose contract and schema define the same fields
- the file naming rule matches the contract

Exit Bar:

- a task can attach a claim manifest without inventing local field names

### Pass 0001 - Workflow Adoption

Goal:

- make the manifest part of normal testing and closure flow

Build:

- update `Processes/TESTING.md` so regression or closure claims should carry a claim manifest
- update `REGRESSION-LEADER.md` to require the manifest when closing a regression run
- add one compact example manifest for default-lane proof and one for supporting-only proof

Unit Proof:

- the testing doc and prompt both require the same manifest
- the two example manifests validate against the schema

Exit Bar:

- claim metadata becomes durable and inspectable before any gating exists

## Testing Strategy

- validate that example manifests only use schema-approved values
- treat mismatched field names across docs and schema as rollout failures

## Deferred Work

- blocking off-lane closure claims automatically
- richer approval UI around claim review
- evidence linting
