# Problem 0001 Option C Plan 0001

## Planning Intent

This file turns Problem `0001`, Option `C. Claim gate` from [ORTHOGONAL-SOLUTIONS-MATRIX.md](../ORTHOGONAL-SOLUTIONS-MATRIX.md) into a bounded implementation sequence.

It is an alternative route, not the selected winner task.

## Summary

Add a narrow gate that blocks `passed`, `fixed`, or `closed` wording unless the closeout artifact explicitly says the proof came from the default lane and counts for closure.

## Fixed Defaults

- scope: shared workflow and closeout artifacts only
- canonical homes:
  - `C:\Users\gregs\.codex\Orchestration\Processes\TESTING.md`
  - `C:\Users\gregs\.codex\Orchestration\ORCHESTRATION.md`
  - `C:\Users\gregs\.codex\Orchestration\Prompts\TASK-LEADER.md`
  - `C:\Users\gregs\.codex\Orchestration\Prompts\REGRESSION-LEADER.md`
  - `C:\Users\gregs\.codex\Orchestration\Prompts\REGRESSION-TESTER.md`
- this option uses inline closeout fields instead of a separate manifest file
- required inline fields:
  - `Repo Default Lane`
  - `Executed Lane`
  - `Counts For Closure`
  - `Supporting Only Reason`

## Pass Plan

### Pass 0000 - Inline Claim Header Contract

Goal:

- define the minimum inline fields needed for a gate decision

Build:

- update `Processes/TESTING.md` to require the four inline fields in regression or closure closeouts
- update `ORCHESTRATION.md` to state that closure wording is blocked when `Counts For Closure=false`
- add one valid and one invalid closeout header example

Unit Proof:

- the two docs use the same field names
- the invalid example clearly shows why off-lane proof cannot close the task

Exit Bar:

- the gate inputs are defined before prompt enforcement starts

### Pass 0001 - Gate Adoption

Goal:

- enforce the gate at the prompt layer

Build:

- update `REGRESSION-LEADER.md`, `REGRESSION-TESTER.md`, and `TASK-LEADER.md` so they may not use closure language when `Counts For Closure=false`
- require `Supporting Only Reason` when the executed lane differs from the repo default lane

Unit Proof:

- all three prompts reference the same inline fields
- the prompts explicitly block closure wording on supporting-only proof

Exit Bar:

- false closure is stopped at the closeout boundary even without a separate manifest artifact

## Testing Strategy

- check that the closeout header fields are named identically in docs and prompts
- treat missing `Supporting Only Reason` on off-lane proof as a rollout failure

## Deferred Work

- a standalone machine-checkable claim manifest
- schema validation for claim artifacts
- richer closure dashboards
