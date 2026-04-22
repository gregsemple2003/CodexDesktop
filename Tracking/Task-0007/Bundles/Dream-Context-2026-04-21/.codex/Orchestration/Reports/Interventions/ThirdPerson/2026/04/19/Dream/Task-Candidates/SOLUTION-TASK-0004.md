# Solution Task 0004

## Title

Require first-disagreement records before runtime fixes or root-cause claims

## Summary

The packet shows symptom-level retunes and loose proof loops before the first bad runtime state was pinned down.
This task adds one shared debugging gate that requires a concrete runtime disagreement record and writer trace before code changes or root-cause claims can advance.

## Goals

- Keep runtime debugging pinned to the first concrete disagreement on the real default lane.
- Make root-cause updates falsifiable by naming observed and expected values plus the evidence link.
- Reduce repeated human steering through symptom-only loops.

## Non-Goals

- Building a full new runtime probe framework in the same task.
- Replacing the human-default-lane regression run with headless diagnostics.
- Declaring root cause from visual descriptions alone.

## Constraints And Baseline

Current truth:

- The packet ends with a sharper method requirement: find the first concrete disagreement, then trace its writers upstream.
- `C:\Agent\ThirdPerson\REGRESSION.md` says the human default lane is the real proof surface.
- `C:\Agent\ThirdPerson\TESTING.md` keeps headless diagnostics in a supporting lane.

Hard constraints:

- runtime defects must stay grounded on default-lane evidence
- the debugging record must survive beyond chat prose
- symptom labels such as "looks wrong" are not enough for closure

## Implementation Home

- `C:\Users\gregs\.codex\Orchestration\Processes\DEBUGGING.md`
- `C:\Users\gregs\.codex\Orchestration\Prompts\DEBUG-LEADER.md`
- `C:\Users\gregs\.codex\Orchestration\Prompts\DEBUG-WORKER.md`
- `C:\Users\gregs\.codex\Orchestration\Prompts\IMPLEMENTATION-LEADER.md`
- task bug artifacts under `C:\Agent\ThirdPerson\Tracking/Task-<id>\BUG-<NNNN>.md`
- task testing evidence under `C:\Agent\ThirdPerson\Tracking/Task-<id>\Testing\`

## Proposed Changes

- Add one required `First Runtime Disagreement` section to runtime-defect bug artifacts with these fields:
  - `lane`
  - `evidence_link`
  - `sample_or_frame`
  - `observed_value`
  - `expected_value`
  - `first_bad_boundary`
- Add one required `Writer Trace` section that records each upstream writer or boundary checked next and the result.
- Update shared debug prompts so they cannot treat symptom descriptions as root cause and cannot move into fix mode until the disagreement record exists.
- Require implementation-leader summaries to cite the disagreement record when they claim a root cause or recommend a fix.

## Expected Resolution

A human reviewing a runtime defect should see the exact bad state, what value was expected, where the evidence lives, and how the writer trace narrowed the cause.
Fix work should start from that record instead of from repeated visual guesses.

## What Does Not Count

- "Root cause" text with no observed and expected values.
- A fix attempt justified only by screenshots or prose summaries.
- A headless-only record for a defect that the repo requires to be proven on the human default lane.
- A bug note that lists hypotheses but no writer trace.

## Acceptance Criteria

- `DEBUGGING.md`, `DEBUG-LEADER.md`, and `DEBUG-WORKER.md` require `First Runtime Disagreement` and `Writer Trace` sections before a runtime fix or root-cause claim can proceed.
- `IMPLEMENTATION-LEADER.md` requires fix summaries and root-cause claims to cite the disagreement record directly.
- On a `ThirdPerson` pilot bug, the artifact records the default-lane evidence link, observed value, expected value, and first bad boundary.
- On the same pilot, at least one writer-trace step is recorded before any fix is presented as the chosen path.
- A symptom-only debug note without those fields fails the gate and cannot be used as closure proof.

## Proof Plan

- Apply the new sections to one open or recreated `ThirdPerson` runtime defect.
- Verify that the first fix recommendation cites the disagreement record and writer trace.
- Attempt one symptom-only summary and verify the gate rejects it.

## References

- `..\BURDEN-ANALYSIS.md`
- `..\ORTHOGONAL-SOLUTIONS-MATRIX.md`
- `..\Plans\PROBLEM-0006-OPTION-A-PLAN-0001.md`
- `C:\Agent\ThirdPerson\REGRESSION.md`
- `C:\Agent\ThirdPerson\TESTING.md`

## Plan Addendum

Chosen plan: require a first-disagreement debugging gate that names the exact bad runtime value, then traces the writers upstream before more fixes are attempted.

Implementation notes from the selected plan:

- debugging cannot move into fix mode until the first concrete bad state is named with values
- each later note must preserve the chain from runtime disagreement to upstream writer
- symptom-level categories are not enough for closure

Rollout:

1. Define the required disagreement record: frame or sample, bad value, expected value, lane, and evidence link.
2. Require a writer-trace section that names each upstream boundary checked next.
3. Add the gate to task workflows before code changes meant to fix a runtime defect.
4. Pilot it on a `ThirdPerson` runtime defect with default-lane evidence only.
5. Treat the debugging record as incomplete until the concrete writer is proven or the remaining branch is tightly bounded.
