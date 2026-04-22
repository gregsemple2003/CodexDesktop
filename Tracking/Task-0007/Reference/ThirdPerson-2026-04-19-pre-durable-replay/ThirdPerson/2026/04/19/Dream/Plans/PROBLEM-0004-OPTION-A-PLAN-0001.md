# Problem 0004 Option A Plan 0001

## Planning Intent

This file turns Problem `0004`, Option `A. Capture guidance` from [ORTHOGONAL-SOLUTIONS-MATRIX.md](../ORTHOGONAL-SOLUTIONS-MATRIX.md) into a bounded implementation sequence.

It is an alternative route, not the selected winner task.

## Summary

Strengthen proof-capture guidance so runtime artifacts clearly show the intended subject, region, and lane before they reach the human.

## Fixed Defaults

- scope: shared process docs and prompts only
- canonical homes:
  - `C:\Users\gregs\.codex\Orchestration\Processes\TESTING.md`
  - `C:\Users\gregs\.codex\Orchestration\Processes\DEBUGGING.md`
  - `C:\Users\gregs\.codex\Orchestration\Prompts\REGRESSION-TESTER.md`
  - `C:\Users\gregs\.codex\Orchestration\Prompts\DEBUG-LEADER.md`
- required capture checklist items:
  - lane named
  - subject named
  - region of interest visible
  - runtime state explicit
  - debug-only substitutes labeled as support only
- no manifest or linter in this rollout

## Pass Plan

### Pass 0000 - Shared Capture Checklist

Goal:

- define one reusable proof-capture checklist

Build:

- update `Processes/TESTING.md` with the capture checklist for regression proof
- update `Processes/DEBUGGING.md` with the same checklist for debug artifacts
- add one anti-pattern example for a screenshot that does not actually show the subject

Unit Proof:

- both process docs use the same checklist terms
- the anti-pattern makes the evidence failure obvious

Exit Bar:

- the shared guidance makes proof hygiene explicit before capture starts

### Pass 0001 - Prompt Adoption

Goal:

- make the capture checklist part of live proof collection behavior

Build:

- update `REGRESSION-TESTER.md` and `DEBUG-LEADER.md` so the checklist is applied before presenting artifacts
- require supporting-only labeling when a debug-only or off-lane artifact is still useful

Unit Proof:

- both prompts reference the same checklist items
- neither prompt treats unlabeled substitute proof as closure proof

Exit Bar:

- poor proof hygiene is less likely to escape to the human even without machine checks

## Testing Strategy

- compare the checklist text across the two process docs and two prompts
- reject wording that leaves subject visibility or lane labeling implicit

## Deferred Work

- evidence manifests
- evidence linting
- automated region-of-interest checks
