# Solution Task 0006

## Title

Add a runtime defect ledger before deeper lane-specific debugging

## Summary

The April 19 runtime debugging flow lost time because the exact defect list, valid evidence set, and proof bar kept shifting in chat.
This task adds a small ledger artifact that pins those items before analysis widens.

## Goals

- record the active runtime lane for the defect
- keep agreed symptoms separate instead of blending them
- record invalidated evidence so rejected proof does not quietly return

## Non-Goals

- building an image classifier
- replacing full bug narratives
- turning every bug into a heavy template

## Implementation Home

- `C:\Users\gregs\.codex\Orchestration\Processes\DEBUGGING.md`
- task-owned bug templates or companion artifacts under `Tracking/Task-<id>/`
- shared debugging prompts under `C:\Users\gregs\.codex\Orchestration\Prompts\`

## Constraints And Baseline

- the ledger must work for runtime-only defects on the human default lane
- the ledger must be easy to update when the human splits or narrows symptoms
- rejected screenshots or off-lane evidence must stay explicitly rejected

## Proposed Changes

- define a `Runtime Defect Ledger` shape with `Claimed Lane`, `Agreed Symptoms`, `Valid Evidence`, `Invalid Evidence`, and `Current Proof Bar`
- require the ledger to be updated before broad root-cause work resumes on a runtime visual defect
- add update rules for cases where the human says a symptom is separate and must stay separate
- point later bug notes and regression reruns back to the current ledger instead of restating it ad hoc

## Expected Resolution

The debugging target should stay stable while deeper analysis happens.
Reviewers should be able to see the active symptom set and the rejected evidence set without reading long chat history.

## What Does Not Count

- a reminder to restate the bug with no durable artifact
- one merged defect line that hides separate runtime symptoms
- keeping invalid evidence only in chat

## Acceptance Criteria

- `DEBUGGING.md` or the related bug template defines the required runtime defect ledger fields
- runtime visual defects require the ledger before deeper root-cause work resumes
- the ledger has explicit `Valid Evidence` and `Invalid Evidence` sections
- later bug and test artifacts can point to one current ledger instead of freehand symptom restatement

## Proof Plan

- test the ledger on the April 19 ThirdPerson foot read, offset or rocking, and hover defect split
- confirm rejected non-runtime evidence remains in `Invalid Evidence`
- confirm later debug steps can cite the ledger instead of rebuilding the symptom list

## References

- `../ORTHOGONAL-SOLUTIONS-MATRIX.md`
- `../Plans/PROBLEM-0005-OPTION-B-PLAN-0001.md`
- `../../DEBUGGING.md`

## Plan Addendum

Keep the ledger narrow.
It should name the lane, the exact agreed symptoms, what evidence still counts, what evidence was rejected, and the current proof bar.
Update it as soon as the human splits or narrows the defect set so deeper analysis stays pinned to the same target.
