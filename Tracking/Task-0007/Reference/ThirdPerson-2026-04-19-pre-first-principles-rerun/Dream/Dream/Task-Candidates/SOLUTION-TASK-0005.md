# Solution Task 0005

## Title

Generate a human approval packet before asking for plan or pass approval

## Summary

April 19 approval requests repeatedly failed because the human did not get a usable diff summary, contextual links, or clear pass framing.
This task adds a small approval packet that makes the requested decision reviewable.

## Goals

- give every approval ask a concise review packet
- show changed files with useful links and plain-language summaries
- state exactly what approval is being requested and under which pass framing

## Non-Goals

- replacing raw diffs when they are still useful
- changing the shared task artifact structure
- adding a large review UI

## Implementation Home

- `C:\Users\gregs\.codex\Orchestration\Processes\TASK-CREATE.md`
- `C:\Users\gregs\.codex\Orchestration\ORCHESTRATION.md`
- shared approval or handoff prompt files under `C:\Users\gregs\.codex\Orchestration\Prompts\`

## Constraints And Baseline

- approval packets should be small and fast to review
- links need enough context to be useful
- reopened work must state whether it is a new pass, a reopened pass, or an audit-only change

## Proposed Changes

- define an `Approval Packet` shape with `Decision Requested`, `Pass Framing`, `Changed Files`, `Diff Summary`, and `Open Risks`
- require the packet before plan approval or pass-structure approval asks
- include file links plus one-line summaries of what changed and why it matters
- add at least one example showing a reopened-pass approval case

## Expected Resolution

The human should be able to decide yes or no without reconstructing the work from scattered files.
Approval turns should get shorter and calmer because the review surface is explicit.

## What Does Not Count

- a bare file list with no change summary
- raw diffs only
- approval asks that do not say what decision is needed

## Acceptance Criteria

- shared workflow docs define the required approval packet sections
- approval prompts require the packet before approval is requested
- changed files in the packet include contextual links and one-line summaries
- examples include a reopened-pass approval case with correct pass framing

## Proof Plan

- replay the April 19 plan-approval friction that lacked diff context
- confirm the new packet gives enough context to approve or reject quickly
- confirm a reopened-pass case makes the pass framing explicit

## References

- `../ORTHOGONAL-SOLUTIONS-MATRIX.md`
- `../Plans/PROBLEM-0004-OPTION-B-PLAN-0001.md`
- `../../TASK-CREATE.md`

## Plan Addendum

Define the approval packet as a small front door, not a new archive.
It should state the decision needed, the pass framing, the changed files, and a short summary of what moved.
Generate it before the approval ask so the human does not have to reconstruct the review state by hand.
