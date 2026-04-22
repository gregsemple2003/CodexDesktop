# Problem 0005 Option B Plan 0001

## Summary

Add a runtime defect ledger that locks the exact lane, symptom list, invalidated evidence, and proof bar before deeper debugging continues.

## What Changes

- define a small bug companion artifact for runtime-only defect framing
- record agreed symptoms separately instead of blending them
- record invalidated evidence so rejected screenshots or off-lane artifacts do not sneak back in
- require the ledger to be refreshed when the human clarifies or splits defects

## Files Or Artifact Types

- task-owned bug or testing companion template
- debugging workflow prompts
- optional schema or checklist for required fields

## Rollout

1. Define fields for lane, defect list, valid evidence, invalid evidence, and current proof bar.
2. Require the ledger before broad root-cause work resumes.
3. Add update rules for new clarifications like `this is a separate runtime-only defect`.
4. Trial on a live runtime defect task.

## Success Checks

- the active runtime defect list is stable and easy to read
- rejected evidence stays explicitly rejected
- later analysis can point back to one current defect ledger instead of chat memory
