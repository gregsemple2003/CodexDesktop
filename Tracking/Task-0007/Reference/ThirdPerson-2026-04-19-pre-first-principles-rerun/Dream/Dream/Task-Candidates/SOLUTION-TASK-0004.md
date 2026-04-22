# Solution Task 0004

## Title

Require a first-disagreement worksheet before root-cause claims on runtime defects

## Summary

The April 19 runtime debugging loop kept risking broad cause labels and bounded tweaks instead of one measured bad state with an upstream writer chain.
This task adds a worksheet that makes `root cause` claims falsifiable.

## Goals

- require one concrete disagreement with values before a root-cause claim
- require boundary-by-boundary tracing from that disagreement to the writer that made it wrong
- preserve contradictory evidence instead of hiding it

## Non-Goals

- building deep always-on instrumentation first
- replacing repo-local bug notes with a new system
- restricting debugging to ThirdPerson only

## Implementation Home

- `C:\Users\gregs\.codex\Orchestration\Processes\DEBUGGING.md`
- shared debugging prompt files under `C:\Users\gregs\.codex\Orchestration\Prompts\`
- task-owned bug and testing artifact templates used during debug work

## Constraints And Baseline

- category-only claims like `single-node seam` are not enough by themselves
- the worksheet must work for real runtime-lane bugs where evidence may disagree
- the artifact should stay small enough to fill during active debugging

## Proposed Changes

- add a required worksheet section for `First Concrete Disagreement`
- require fields for lane, timestamp or frame, measured bad values, expected values, and artifact refs
- add a `Writer Chain` section that traces state one boundary at a time until the responsible writer is named or the remaining branch is tightly bounded
- add a `Contradictory Evidence` section that keeps evidence that does not fit the current hypothesis

## Expected Resolution

Reviewers should be able to tell whether a root-cause claim is proven, still branching, or just a category guess.
Runtime debugging should spend less time on tweak loops that never named the bad state precisely.

## What Does Not Count

- linking the debugging doc with no required worksheet fields
- a root-cause claim with no measured disagreement values
- hiding contradictory evidence because it complicates the story

## Acceptance Criteria

- `DEBUGGING.md` requires a `First Concrete Disagreement` section for runtime root-cause claims
- shared prompts and templates require a `Writer Chain` and `Contradictory Evidence` section
- category-only cause labels fail review unless they also include a measured disagreement and writer chain
- the worksheet can represent either confirmed cause or a tightly bounded remaining branch

## Proof Plan

- test the worksheet against the April 19 ThirdPerson hover and rocking examples from the packet
- confirm a category-only claim fails the new review bar
- confirm a measured seam with traced writers passes the structure check

## References

- `../ORTHOGONAL-SOLUTIONS-MATRIX.md`
- `../Plans/PROBLEM-0006-OPTION-B-PLAN-0001.md`
- `../../DEBUGGING.md`

## Plan Addendum

Keep the worksheet tied to one bad state.
Capture the values, the frame or timestamp, and the artifact refs first.
Then force the writer chain to move boundary by boundary until the responsible writer is named or the remaining uncertainty is tightly bounded and preserved.
