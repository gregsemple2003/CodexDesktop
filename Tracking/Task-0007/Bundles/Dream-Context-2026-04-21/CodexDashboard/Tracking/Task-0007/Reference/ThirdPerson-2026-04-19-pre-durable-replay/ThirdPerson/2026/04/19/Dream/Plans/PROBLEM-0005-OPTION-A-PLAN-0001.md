# Problem 0005 Option A Plan 0001

## Planning Intent

This file turns Problem `0005`, Option `A. Stronger standing instruction wording` from [ORTHOGONAL-SOLUTIONS-MATRIX.md](../ORTHOGONAL-SOLUTIONS-MATRIX.md) into a bounded implementation sequence.

It is an alternative route, not the selected winner task.

## Summary

Strengthen shared prompt wording so active work stays owned until there is one clearly named reason to stop.

## Fixed Defaults

- scope: shared workflow prose and leader prompts only
- canonical homes:
  - `C:\Users\gregs\.codex\Orchestration\ORCHESTRATION.md`
  - `C:\Users\gregs\.codex\Orchestration\Prompts\TASK-LEADER.md`
  - `C:\Users\gregs\.codex\Orchestration\Prompts\IMPLEMENTATION-LEADER.md`
- required phrases:
  - keep ownership until a named stop reason applies
  - recoverable failure does not return ownership to the human
  - do not ask the human what to do next when the next bounded step is still available
- no task-state fields or watcher in this rollout

## Pass Plan

### Pass 0000 - Shared Ownership Wording

Goal:

- define the ownership rule once in shared workflow prose

Build:

- update `ORCHESTRATION.md` with one compact ownership rule and one anti-pattern
- define named stop reasons in prose even though they are not yet schema-backed

Unit Proof:

- the ownership rule and stop reasons live in one place
- the anti-pattern shows why recoverable failure is not a stop

Exit Bar:

- the shared workflow no longer treats stopping as a vague preference

### Pass 0001 - Prompt Adoption

Goal:

- make the wording show up at live execution boundaries

Build:

- update `TASK-LEADER.md` and `IMPLEMENTATION-LEADER.md` with the same ownership rule
- add one line that waking the human up to resume ordinary work is a failure mode

Unit Proof:

- both prompts reference the same stop reasons
- neither prompt allows an implicit pause after a failed attempt

Exit Bar:

- the prompt layer now reinforces the same ownership discipline as the workflow prose

## Testing Strategy

- compare the stop-reason wording across shared docs and prompts
- reject any wording that still frames pause vs continue as a casual judgment call

## Deferred Work

- durable task-state fields
- schema validation
- background monitoring
