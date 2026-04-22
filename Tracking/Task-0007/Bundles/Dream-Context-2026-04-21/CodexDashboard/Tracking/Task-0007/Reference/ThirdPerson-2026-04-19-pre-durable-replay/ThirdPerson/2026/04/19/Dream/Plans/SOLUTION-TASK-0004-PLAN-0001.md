# Solution Task 0004 Plan 0001

## Planning Intent

This file turns [SOLUTION-TASK-0004.md](../Task-Candidates/SOLUTION-TASK-0004.md) into a bounded implementation sequence.

It describes the intended route to done, not regression history.

## Summary

Add one claim manifest and one closure-claim gate so regression and closure wording cannot use off-lane proof as if it counted.

## Fixed Defaults

- manifest home:
  - `C:\Users\gregs\.codex\Orchestration\CLAIM-MANIFEST.md`
  - `C:\Users\gregs\.codex\Orchestration\CLAIM-MANIFEST.schema.json`
- task-owned manifest name:
  - `Tracking/Task-<id>/Testing/CLAIM-MANIFEST-<NNNN>.json`
- claim kinds:
  - `regression_pass`
  - `closure_ready`
  - `supporting_only`
- off-lane proof remains allowed as support only

## Pass Plan

### Pass 0000 - Manifest Contract And Naming

Goal:

- define the claim manifest fields and task-owned artifact name

Build:

- add `CLAIM-MANIFEST.md`
- add `CLAIM-MANIFEST.schema.json`
- update `FILE-NAMING.md` with the manifest artifact name

Unit Proof:

- the doc and schema use the same fields and enum values
- the naming doc uses the same filename the workflow expects

Exit Bar:

- a reviewer can tell exactly how a claim manifest should look on disk

### Pass 0001 - Testing And Closure Adoption

Goal:

- make regression and closure flows require the manifest

Build:

- update `Processes\TESTING.md`
- update `REGRESSION-LEADER.md`
- update `REGRESSION-TESTER.md`
- update `TASK-LEADER.md`

Unit Proof:

- each workflow doc names the claim manifest requirement
- each workflow doc blocks closure wording when `counts_for_claim=false`

Exit Bar:

- shared testing and closure workflows no longer allow lane ambiguity

### Pass 0002 - Supporting-Proof Examples

Goal:

- make the difference between closure proof and supporting proof easy to apply

Build:

- add one valid default-lane closure example
- add one valid supporting-only example
- add one invalid example where off-lane proof is rejected as closure proof

Unit Proof:

- examples map cleanly to `regression_pass`, `closure_ready`, or `supporting_only`
- the invalid example shows the gate behavior explicitly

Exit Bar:

- a later reviewer can tell why an off-lane artifact does or does not count

## Testing Strategy

- validate field and enum consistency across manifest doc, schema, and prompt adoption
- verify that supporting proof is still allowed but cannot claim closure
- reject any rollout that leaves lane meaning implicit

## Deferred Work

Keep these out of this rollout unless the task expands intentionally:

- screenshot-region validation
- evidence-quality linting
- richer closure UI or dashboards
