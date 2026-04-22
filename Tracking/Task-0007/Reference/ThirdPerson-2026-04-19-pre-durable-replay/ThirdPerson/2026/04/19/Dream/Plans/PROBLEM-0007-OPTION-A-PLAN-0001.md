# Problem 0007 Option A Plan 0001

## Planning Intent

This file turns Problem `0007`, Option `A. Better debugging prose` from [ORTHOGONAL-SOLUTIONS-MATRIX.md](../ORTHOGONAL-SOLUTIONS-MATRIX.md) into a bounded implementation sequence.

It is an alternative route, not the selected winner task.

## Summary

Strengthen the shared debugging guidance so investigations must start at the first concrete disagreement and keep narrowing instead of iterating symptoms.

## Fixed Defaults

- scope: shared debugging prose and prompts only
- canonical homes:
  - `C:\Users\gregs\.codex\Orchestration\Processes\DEBUGGING.md`
  - `C:\Users\gregs\.codex\Orchestration\Prompts\DEBUG-LEADER.md`
  - `C:\Users\gregs\.codex\Orchestration\Prompts\DEBUG-WORKER.md`
- required method terms:
  - expected state
  - observed state
  - disagreement seam
  - next upstream writer
- no template artifact or verifier in this rollout

## Pass Plan

### Pass 0000 - Shared Narrowing Method

Goal:

- define the debugging method once in shared guidance

Build:

- update `Processes/DEBUGGING.md` with the four-step narrowing method
- add one anti-pattern showing why symptom tweaking is not narrowing

Unit Proof:

- the narrowing method uses concrete state-comparison language
- the anti-pattern clearly contrasts with the intended method

Exit Bar:

- the shared debug process tells people how to narrow instead of only telling them to be rigorous

### Pass 0001 - Prompt Adoption And Examples

Goal:

- make the same narrowing method show up in live debug work

Build:

- update `DEBUG-LEADER.md` and `DEBUG-WORKER.md` with the same four-step method
- add one short example that traces from bad runtime state to upstream writer

Unit Proof:

- both prompts use the same method terms as `DEBUGGING.md`
- the example is concrete enough to copy

Exit Bar:

- debugging guidance is stronger and more consistent even without a hard verifier

## Testing Strategy

- compare the method terms across the process doc and the debug prompts
- reject any wording that still treats `bounded tweak` as equivalent to `narrowing`

## Deferred Work

- required disagreement artifacts
- root-cause claim gating
- debug-closeout schema validation
