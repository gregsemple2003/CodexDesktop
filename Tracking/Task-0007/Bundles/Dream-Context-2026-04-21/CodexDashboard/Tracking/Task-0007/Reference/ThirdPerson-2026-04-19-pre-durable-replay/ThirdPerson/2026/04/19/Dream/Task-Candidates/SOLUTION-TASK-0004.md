# Solution Task 0004

## Title

Add a claim manifest and block off-lane closure claims.

## Summary

The packet shows false closure pressure where proof from the wrong lane or wrong evidence type was treated as if it counted for regression or task closure.

This is a concrete implementation task. It will add one claim manifest artifact and one gate that prevents `passed`, `fixed`, or `closed` claims when the manifest says the proof is only supporting evidence.

## Goals

- make every regression or closure claim name its lane and proof type
- preserve supporting proof while preventing it from pretending to be closure proof
- keep repo-root `REGRESSION.md` as the real source of truth

## Non-Goals

- changing repo-root regression expectations
- banning supporting proof
- solving screenshot framing or region-of-interest problems beyond the claim boundary

## Constraints And Baseline

- repo-root `REGRESSION.md` stays authoritative
- task-owned artifacts may record proof but may not redefine the lane
- the gate must work with existing `REGRESSION-RUN-<NNNN>.md` and closure flow

## Proposed Changes

- add `C:\Users\gregs\.codex\Orchestration\CLAIM-MANIFEST.md`
- add `C:\Users\gregs\.codex\Orchestration\CLAIM-MANIFEST.schema.json`
- update [FILE-NAMING.md](../../../../../../../../FILE-NAMING.md) with `Tracking/Task-<id>/Testing/CLAIM-MANIFEST-<NNNN>.json`
- update `C:\Users\gregs\.codex\Orchestration\Processes\TESTING.md`
- update [ORCHESTRATION.md](../../../../../../../../ORCHESTRATION.md)
- update `C:\Users\gregs\.codex\Orchestration\Prompts\REGRESSION-LEADER.md`
- update `C:\Users\gregs\.codex\Orchestration\Prompts\REGRESSION-TESTER.md`
- update [TASK-LEADER.md](../../../../../../../../Prompts/TASK-LEADER.md)
- define these exact manifest fields: `claim_kind`, `repo_default_lane`, `executed_lane`, `is_default_lane`, `evidence_type`, `is_runtime_evidence`, `artifact_refs`, `counts_for_claim`, and `supporting_only_reason`
- define these exact `claim_kind` values: `regression_pass`, `closure_ready`, and `supporting_only`
- add a gate rule that `counts_for_claim=false` blocks `passed`, `fixed`, or `closed` wording in regression and closure output

## Expected Resolution

- every closure claim has a claim manifest
- off-lane proof is still allowed but is labeled `supporting_only`
- the system stops claiming success on proof that does not satisfy the repo lane

## What Does Not Count

- a prose reminder with no manifest
- a regression run that implies the lane but does not name it
- relabeling off-lane proof as closure proof
- a manifest that leaves `counts_for_claim` or `supporting_only_reason` blank

## Implementation Home

Implement the shared claim contract under:

- `C:\Users\gregs\.codex\Orchestration\CLAIM-MANIFEST.md`
- `C:\Users\gregs\.codex\Orchestration\CLAIM-MANIFEST.schema.json`
- [FILE-NAMING.md](../../../../../../../../FILE-NAMING.md)
- `C:\Users\gregs\.codex\Orchestration\Processes\TESTING.md`
- [ORCHESTRATION.md](../../../../../../../../ORCHESTRATION.md)
- `C:\Users\gregs\.codex\Orchestration\Prompts\REGRESSION-LEADER.md`
- `C:\Users\gregs\.codex\Orchestration\Prompts\REGRESSION-TESTER.md`
- [TASK-LEADER.md](../../../../../../../../Prompts/TASK-LEADER.md)

## Proof Plan

- check that the shared doc, schema, and naming rule all exist
- check that the manifest fields and enum values match the list above
- check that regression and closure prompts now block claim wording when `counts_for_claim=false`

## Acceptance Criteria

- `CLAIM-MANIFEST.md` and `CLAIM-MANIFEST.schema.json` exist and define every field listed in `Proposed Changes`
- [FILE-NAMING.md](../../../../../../../../FILE-NAMING.md) includes `CLAIM-MANIFEST-<NNNN>.json`
- `Processes\TESTING.md`, `REGRESSION-LEADER.md`, `REGRESSION-TESTER.md`, and [TASK-LEADER.md](../../../../../../../../Prompts/TASK-LEADER.md) all require a claim manifest for regression-pass or closure-ready claims
- the claim manifest design includes the enum values `regression_pass`, `closure_ready`, and `supporting_only`
- the gate rule says `counts_for_claim=false` blocks `passed`, `fixed`, and `closed` wording
- the written design makes clear that repo-root `REGRESSION.md` still decides what lane counts

## References

- [TASK-CREATE.md](../../../../../../../../Processes/TASK-CREATE.md)
- [BURDEN-ANALYSIS.md](../BURDEN-ANALYSIS.md)
- [ORTHOGONAL-SOLUTIONS-MATRIX.md](../ORTHOGONAL-SOLUTIONS-MATRIX.md)
