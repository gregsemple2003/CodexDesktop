# Solution Task 0002

## Title

Add a boundary manifest for live repo, pass, lane, and hard-constraint state

## Summary

April 19 work kept drifting after the human added new hard rules such as repo-local ownership, no engine mods, new pass framing, and debugging-method requirements.
This task makes those boundaries durable state instead of chat memory.

## Goals

- capture active repo, lane, pass, ownership, and hard constraints in one durable manifest
- refresh that manifest when the human adds or tightens a hard rule
- use the manifest to block out-of-bound actions before edits, proof claims, or handoff text

## Non-Goals

- replacing repo-local docs
- inventing new repo policy
- adding a ThirdPerson-only special case instead of a reusable boundary object

## Implementation Home

- `C:\Users\gregs\.codex\Orchestration\TASK-STATE.md`
- `C:\Users\gregs\.codex\Orchestration\TASK-STATE.schema.json`
- `C:\Users\gregs\.codex\Orchestration\ORCHESTRATION.md`
- shared prompt files under `C:\Users\gregs\.codex\Orchestration\Prompts\`

## Constraints And Baseline

- the manifest must reflect repo-local truth from the packet repo
- a new hard rule should update durable state immediately, not at closeout
- closed pass history must remain historical rather than being silently rewritten

## Proposed Changes

- extend task state guidance and schema with an `active_boundaries` object
- require fields for repo, active lane, active pass, ownership lane, hard constraints, prohibited surfaces, and last boundary update reason
- add workflow steps that refresh `active_boundaries` when the human states new rules like `no engine mods` or `put all new work under a new pass`
- add a pre-action check that requires later plan, proof, or edit steps to cite a manifest mismatch when they refuse or redirect an action

## Expected Resolution

The current constraint set becomes visible in durable state and stays live across the task.
Reviewers should be able to see why an action is allowed or blocked without replaying the full chat.

## What Does Not Count

- a prose reminder with no durable fields
- burying new constraints only in handoff text
- adding a manifest that is created once and never refreshed

## Acceptance Criteria

- `TASK-STATE.md` and its schema define `active_boundaries` with the required live-boundary fields
- shared workflow steps say that a new hard rule updates `active_boundaries` immediately
- prompts or validators can flag an out-of-bound action by pointing to a concrete manifest field
- pass rewrites and repo-crossing drift can be explained against manifest state instead of freehand chat recall

## Proof Plan

- replay the April 19 ThirdPerson sequence that added `no engine mods` and new pass framing
- confirm the manifest changes as soon as those rules appear
- confirm a later off-bound action can be rejected with a concrete manifest mismatch

## References

- `../ORTHOGONAL-SOLUTIONS-MATRIX.md`
- `../Plans/PROBLEM-0003-OPTION-B-PLAN-0001.md`
- `../../AGENTS.md`

## Plan Addendum

Keep the manifest small.
It only needs the fields that later work can actually check.
Populate it at task start, refresh it whenever the human adds a hard rule, and make later proof, planning, and edit steps compare themselves against it before they proceed.
