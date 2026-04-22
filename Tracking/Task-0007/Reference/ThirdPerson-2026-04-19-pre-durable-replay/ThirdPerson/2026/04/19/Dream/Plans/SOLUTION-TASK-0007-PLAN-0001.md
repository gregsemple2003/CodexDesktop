# Solution Task 0007 Plan 0001

## Planning Intent

This file turns [SOLUTION-TASK-0007.md](../Task-Candidates/SOLUTION-TASK-0007.md) into a bounded implementation sequence.

It describes the intended route to done, not proof history.

## Summary

Add one evidence manifest so proof bundles say what claim they support, what lane they came from, what they show, and whether they count as closure or only support.

## Fixed Defaults

- manifest home:
  - `C:\Users\gregs\.codex\Orchestration\EVIDENCE-MANIFEST.md`
  - `C:\Users\gregs\.codex\Orchestration\EVIDENCE-MANIFEST.schema.json`
- task-owned manifest name:
  - `Tracking/Task-<id>/Testing/EVIDENCE-MANIFEST-<NNNN>.json`
- capture surfaces:
  - `runtime`
  - `editor`
  - `source_export`
  - `log_only`
- allowed uses:
  - `closure`
  - `supporting_only`
  - `diagnostic_only`

## Pass Plan

### Pass 0000 - Manifest Contract And Naming

Goal:

- define the evidence-manifest fields, enums, and task-owned artifact name

Build:

- add `EVIDENCE-MANIFEST.md`
- add `EVIDENCE-MANIFEST.schema.json`
- update `FILE-NAMING.md` with the manifest artifact name

Unit Proof:

- the doc and schema use the same fields and enum values
- the naming doc matches the workflow contract

Exit Bar:

- a reviewer can tell what an evidence manifest must contain before proof is presented

### Pass 0001 - Testing And Debugging Adoption

Goal:

- make shared testing and debugging flows require the manifest

Build:

- update `Processes\TESTING.md`
- update `Processes\DEBUGGING.md`
- update `REGRESSION-LEADER.md`
- update `REGRESSION-TESTER.md`
- update `DEBUG-LEADER.md`

Unit Proof:

- each workflow doc requires the manifest for proof bundles
- each workflow doc preserves the difference between closure, supporting-only, and diagnostic-only use

Exit Bar:

- proof bundles in regression and debugging cannot stay lane- or subject-ambiguous

### Pass 0002 - Closure Rules And Examples

Goal:

- make the manifest easy to use and hard to game

Build:

- add one valid runtime-closure example
- add one supporting-only source-export example
- add one invalid example where `allowed_use=closure` is rejected because `runtime_visible=false`

Unit Proof:

- examples use the allowed enum values only
- the invalid example demonstrates the closure rule clearly

Exit Bar:

- a later reviewer can tell what an evidence manifest proves and what it does not prove

## Testing Strategy

- validate consistency across manifest doc, schema, naming, and prompt adoption
- reject any rollout that allows runtime closure with non-runtime capture surfaces
- reject any rollout that leaves subject or region-of-interest meaning implicit

## Deferred Work

Keep these out of this rollout unless the task expands intentionally:

- automated image-quality checks
- OCR or screenshot-region detection
- cross-manifest aggregation or dashboards
