# Solution Task 0007

## Title

Require an evidence manifest for proof bundles.

## Summary

The packet shows proof artifacts that did not clearly say what they proved, what lane they came from, or what part of the problem they actually showed.

This is a concrete implementation task. It will add one evidence-manifest artifact for proof bundles and require regression and debug flows to use it before presenting proof as closure or defect evidence.

## Goals

- bind each proof bundle to one claim, one lane, and one subject
- show whether the artifact came from runtime, editor, source export, or logs
- keep closure proof distinct from supporting or diagnostic proof

## Non-Goals

- full automated screenshot analysis
- replacing proof artifacts with prose summaries
- solving answer-shape or approval problems inside this task

## Constraints And Baseline

- supporting and diagnostic proof are still useful
- the manifest must work for runtime captures, editor captures, source exports, and logs
- the contract must stay inspectable on disk

## Proposed Changes

- add `C:\Users\gregs\.codex\Orchestration\EVIDENCE-MANIFEST.md`
- add `C:\Users\gregs\.codex\Orchestration\EVIDENCE-MANIFEST.schema.json`
- update [FILE-NAMING.md](../../../../../../../../FILE-NAMING.md) with `Tracking/Task-<id>/Testing/EVIDENCE-MANIFEST-<NNNN>.json`
- update `C:\Users\gregs\.codex\Orchestration\Processes\TESTING.md`
- update `C:\Users\gregs\.codex\Orchestration\Processes\DEBUGGING.md`
- update `C:\Users\gregs\.codex\Orchestration\Prompts\REGRESSION-LEADER.md`
- update `C:\Users\gregs\.codex\Orchestration\Prompts\REGRESSION-TESTER.md`
- update `C:\Users\gregs\.codex\Orchestration\Prompts\DEBUG-LEADER.md`
- define these exact manifest fields: `claim_id`, `claim_summary`, `executed_lane`, `capture_surface`, `artifact_refs`, `subject`, `region_of_interest`, `runtime_visible`, `allowed_use`, and `limitations`
- define these exact `capture_surface` values: `runtime`, `editor`, `source_export`, and `log_only`
- define these exact `allowed_use` values: `closure`, `supporting_only`, and `diagnostic_only`
- add a rule that `allowed_use=closure` is only valid when `capture_surface=runtime` for runtime claims and `runtime_visible=true`

## Expected Resolution

- every proof bundle has a manifest
- reviewers can tell what the bundle proves and what it does not prove
- source-only, editor-only, or log-only artifacts stop pretending to be runtime closure proof

## What Does Not Count

- a screenshot folder with no manifest
- a manifest that omits subject or region of interest
- a runtime claim with `capture_surface=source_export`
- `allowed_use=closure` when `runtime_visible=false`

## Implementation Home

Implement the shared evidence contract under:

- `C:\Users\gregs\.codex\Orchestration\EVIDENCE-MANIFEST.md`
- `C:\Users\gregs\.codex\Orchestration\EVIDENCE-MANIFEST.schema.json`
- [FILE-NAMING.md](../../../../../../../../FILE-NAMING.md)
- `C:\Users\gregs\.codex\Orchestration\Processes\TESTING.md`
- `C:\Users\gregs\.codex\Orchestration\Processes\DEBUGGING.md`
- `C:\Users\gregs\.codex\Orchestration\Prompts\REGRESSION-LEADER.md`
- `C:\Users\gregs\.codex\Orchestration\Prompts\REGRESSION-TESTER.md`
- `C:\Users\gregs\.codex\Orchestration\Prompts\DEBUG-LEADER.md`

## Proof Plan

- check that the shared manifest doc, schema, and naming rule all exist
- check that the exact fields and enum values match the list above
- check that testing and debugging prompts require the manifest before claiming closure or defect proof

## Acceptance Criteria

- `EVIDENCE-MANIFEST.md` and `EVIDENCE-MANIFEST.schema.json` exist and define every field listed in `Proposed Changes`
- the manifest design defines `capture_surface` values `runtime`, `editor`, `source_export`, and `log_only`
- the manifest design defines `allowed_use` values `closure`, `supporting_only`, and `diagnostic_only`
- [FILE-NAMING.md](../../../../../../../../FILE-NAMING.md) includes `EVIDENCE-MANIFEST-<NNNN>.json`
- `Processes\TESTING.md`, `Processes\DEBUGGING.md`, `REGRESSION-LEADER.md`, `REGRESSION-TESTER.md`, and `DEBUG-LEADER.md` require the manifest for proof bundles
- the written rule says `allowed_use=closure` is invalid for runtime claims unless `capture_surface=runtime` and `runtime_visible=true`

## References

- [TASK-CREATE.md](../../../../../../../../Processes/TASK-CREATE.md)
- [BURDEN-ANALYSIS.md](../BURDEN-ANALYSIS.md)
- [ORTHOGONAL-SOLUTIONS-MATRIX.md](../ORTHOGONAL-SOLUTIONS-MATRIX.md)
