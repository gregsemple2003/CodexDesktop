# Problem 0004 Option C Plan 0001

## Planning Intent

This file turns Problem `0004`, Option `C. Evidence linter` from [ORTHOGONAL-SOLUTIONS-MATRIX.md](../ORTHOGONAL-SOLUTIONS-MATRIX.md) into a bounded implementation sequence.

It is an alternative route, not the selected winner task.

## Summary

Add a small evidence-bundle linter that rejects proof bundles when their machine-readable metadata says they are off-lane, non-runtime, or missing required subject coverage.

## Fixed Defaults

- scope: shared orchestration tool plus prompt adoption
- canonical homes:
  - `C:\Users\gregs\.codex\Orchestration\Tools\lint-evidence-bundle.py`
  - `C:\Users\gregs\.codex\Orchestration\EVIDENCE-LINTER.md`
  - `C:\Users\gregs\.codex\Orchestration\Prompts\REGRESSION-TESTER.md`
  - `C:\Users\gregs\.codex\Orchestration\Prompts\DEBUG-LEADER.md`
- required linter inputs:
  - one machine-readable evidence record per bundle
  - declared claim ref
  - declared capture surface
  - declared region of interest
  - declared allowed use
- initial failure codes:
  - `off_lane`
  - `not_runtime_when_runtime_required`
  - `missing_region_of_interest`
  - `supporting_only_presented_as_closure`

## Pass Plan

### Pass 0000 - Linter Contract

Goal:

- define the linter input and failure outputs before writing the tool

Build:

- add `EVIDENCE-LINTER.md` with CLI contract, input fields, and failure codes
- define the minimum machine-readable evidence record the linter expects
- document that missing machine-readable metadata is itself a linter failure

Unit Proof:

- the contract names the same failure codes everywhere
- the required input fields are concrete enough to build against

Exit Bar:

- the linter behavior is specified before prompt or tool adoption begins

### Pass 0001 - Tool And Prompt Adoption

Goal:

- make the linter runnable and required before proof presentation

Build:

- add `Tools/lint-evidence-bundle.py`
- update `REGRESSION-TESTER.md` and `DEBUG-LEADER.md` so linter failures block evidence presentation as closure proof
- return structured failure output that can be copied into task artifacts

Unit Proof:

- linter fixtures pass and fail on the expected cases
- the prompts explicitly treat linter failure as a stop for closure proof

Exit Bar:

- invalid evidence bundles are rejected before they reach the human

### Pass 0002 - Eval Fixtures

Goal:

- tie the linter to the known April 19 evidence failures

Build:

- add small fixtures that model one valid runtime bundle and two invalid bundles
- preserve one packet-backed note showing which failure code would have fired

Unit Proof:

- each fixture triggers the expected linter result
- the packet-backed example maps to a specific failure code

Exit Bar:

- the linter can be judged against the real failure pattern instead of only synthetic examples

## Testing Strategy

- keep the failure-code table synchronized across doc, tool, and prompts
- treat any silent pass on missing metadata as a rollout failure

## Deferred Work

- visual computer-vision checks
- dashboard proof gallery
- human override workflow for disputed linter failures
