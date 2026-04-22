# Problem 0001 Option A Plan 0001

## Planning Intent

This file turns Problem `0001`, Option `A. Policy rule` from [ORTHOGONAL-SOLUTIONS-MATRIX.md](../ORTHOGONAL-SOLUTIONS-MATRIX.md) into a bounded implementation sequence.

It is an alternative route, not the selected winner task.

## Summary

Strengthen shared docs and leader prompts so closure language always names the human default lane and rejects casual lane drift in prose.

## Fixed Defaults

- scope: shared docs and prompts only
- canonical homes:
  - `C:\Users\gregs\.codex\Orchestration\ORCHESTRATION.md`
  - `C:\Users\gregs\.codex\Orchestration\Processes\TESTING.md`
  - `C:\Users\gregs\.codex\Orchestration\Prompts\TASK-LEADER.md`
  - `C:\Users\gregs\.codex\Orchestration\Prompts\REGRESSION-LEADER.md`
  - `C:\Users\gregs\.codex\Orchestration\Prompts\REGRESSION-TESTER.md`
- repo-root `REGRESSION.md` remains the authoritative lane definition in each repo
- no new schema, manifest, or gate in this rollout

## Pass Plan

### Pass 0000 - Define The Lane Rule

Goal:

- make the default-lane requirement explicit anywhere closure or regression language is defined

Build:

- add one shared rule in `ORCHESTRATION.md` that closure claims must bind to the repo's default lane
- update `Processes/TESTING.md` to distinguish supporting proof from closure proof in plain language
- add one short anti-pattern example for off-lane proof

Unit Proof:

- the two shared docs use the same lane language
- the anti-pattern example clearly shows what does not count

Exit Bar:

- a reader can tell from shared docs alone that wrong-lane proof is not closure proof

### Pass 0001 - Prompt Adoption

Goal:

- make the policy show up at the moments where closure wording is produced

Build:

- update `TASK-LEADER.md`, `REGRESSION-LEADER.md`, and `REGRESSION-TESTER.md` to require lane naming before `passed`, `fixed`, or `closed`
- add one short phrase to each prompt that supporting proof must be labeled as supporting only

Unit Proof:

- each prompt names the same closure rule
- none of the prompts allow lane ambiguity at closeout

Exit Bar:

- the shared prose and prompt layer now say the same thing about lane-bound closure

## Testing Strategy

- check that the same lane terms appear in shared docs and prompts
- treat wording that allows implied lane assumptions as a rollout failure

## Deferred Work

- structured claim metadata
- automated closure gating
- repo-local claim manifests
