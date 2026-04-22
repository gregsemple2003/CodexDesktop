# Problem 0001 Option B Plan 0001

## Summary

Add a proof-surface gate that compares claimed closure evidence against repo regression rules before a task can claim regression success.

## What Changes

- define a small machine-readable proof-surface record for claimed lane, map, mode, pawn, and evidence type
- add a validator that reads repo-local `REGRESSION.md` and `TESTING.md` rules or their extracted task-owned equivalents
- block closeout text that labels supporting proof as regression closure

## Files Or Artifact Types

- shared orchestration validator or generator code
- repo-task regression-run artifact template
- shared testing or closeout prompt bundle

## Rollout

1. Define the claimed proof-surface fields that must exist in regression artifacts.
2. Teach the closeout generator to compare those fields against the repo's regression lane rules.
3. Emit a hard warning or stop-state when the claimed proof is headless, operator-only, or otherwise off-lane.
4. Test against April 19 ThirdPerson examples where supporting proof existed but closure was false.

## Success Checks

- a task cannot honestly claim default-lane regression closure without matching lane metadata
- supporting proof can still be preserved, but it is labeled as supporting only
- ThirdPerson-style false closures are blocked before handoff text is emitted
