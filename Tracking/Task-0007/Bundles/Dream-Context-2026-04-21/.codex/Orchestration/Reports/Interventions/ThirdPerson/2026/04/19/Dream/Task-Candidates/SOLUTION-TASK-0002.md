# Solution Task 0002

## Title

Add an active boundary ledger to task state and enforce it before action

## Summary

The packet shows repeated lane, pass, ownership, and method resets.
This task adds one durable boundary ledger to task state and a pre-action boundary check so later steps cannot drift away from already-stated `ThirdPerson` constraints.

## Goals

- Keep active boundaries visible after the turn where the human states them.
- Force each next action to check lane, pass, ownership, wait rule, and hard constraints against the current ledger.
- Reduce repeated human restatement of the same `ThirdPerson` boundaries.

## Non-Goals

- Replacing task plans with chat memory.
- Asking the human to re-confirm every boundary after it is written once.
- Solving boundary drift only by shrinking task scope.

## Constraints And Baseline

Current truth:

- The packet shows repeated resets around default-lane proof, pass ownership, wait behavior, `no engine mods`, and root-cause method.
- `C:\Agent\ThirdPerson\AGENTS.md` and repo-local testing rules already make some of those boundaries durable, but the active task flow did not retain them.

Hard constraints:

- the ledger must record repo-local rules from `C:\Agent\ThirdPerson`
- short corrections must become durable state updates
- the check must happen before action, not only during audit

## Implementation Home

- `C:\Users\gregs\.codex\Orchestration\ORCHESTRATION.md`
- `C:\Users\gregs\.codex\Orchestration\TASK-STATE.md`
- `C:\Users\gregs\.codex\Orchestration\TASK-STATE.schema.json`
- `C:\Users\gregs\.codex\Orchestration\Prompts\TASK-LEADER.md`
- `C:\Users\gregs\.codex\Orchestration\Prompts\IMPLEMENTATION-LEADER.md`
- `C:\Users\gregs\.codex\Orchestration\Prompts\DEBUG-LEADER.md`
- `C:\Users\gregs\.codex\Orchestration\Prompts\REGRESSION-LEADER.md`
- pilot state under `C:\Agent\ThirdPerson\Tracking/Task-<id>\TASK-STATE.json`

## Proposed Changes

- Add a required `boundary_ledger` object to task state with these fields:
  - `lane_rules`
  - `active_pass`
  - `ownership_rule`
  - `wait_rule`
  - `hard_constraints`
  - `method_constraints`
  - `last_human_update`
- Add a shared pre-action rule that compares the proposed next action against the active ledger before the agent proceeds.
- Require prompts for task, implementation, debug, and regression leadership to update the ledger when the human narrows or corrects a boundary.
- Require human-facing review artifacts to show which ledger version governed the step that was taken.

## Expected Resolution

A human should no longer need to restate the same lane, pass, ownership, or method boundary after it has already been written into task state.
If a proposed action conflicts with the ledger, the system should stop and say which boundary it would violate.

## What Does Not Count

- Repeating the boundary in one chat reply without writing it to task state.
- A general reminder to "be careful" with no concrete ledger fields.
- Narrower work batches with no durable pre-action check.
- A ledger that exists only for audit after the wrong step already happened.

## Acceptance Criteria

- `TASK-STATE.md` and `TASK-STATE.schema.json` define and validate the `boundary_ledger` object and its required fields.
- `ORCHESTRATION.md`, `TASK-LEADER.md`, `IMPLEMENTATION-LEADER.md`, `DEBUG-LEADER.md`, and `REGRESSION-LEADER.md` require a pre-action check against the active ledger.
- On a `ThirdPerson` pilot, the ledger can store and preserve boundaries such as human-default-lane proof, `no engine mods`, pass ownership, and wait-before-intervening rules.
- On the same pilot, a proposed action that violates one of those stored boundaries is blocked before execution and names the conflicting field.
- Short human corrections update the ledger immediately instead of waiting for pass closeout.

## Proof Plan

- Populate a `ThirdPerson` pilot task state with at least four packet-backed boundaries.
- Attempt one conflicting next action and verify the ledger blocks it before work proceeds.
- Verify one short corrective human message updates the ledger without additional confirmation turns.

## References

- `..\BURDEN-ANALYSIS.md`
- `..\ORTHOGONAL-SOLUTIONS-MATRIX.md`
- `..\Plans\PROBLEM-0003-OPTION-A-PLAN-0001.md`
- `C:\Agent\ThirdPerson\AGENTS.md`
- `C:\Agent\ThirdPerson\REGRESSION.md`

## Plan Addendum

Chosen plan: keep an active boundary ledger for the task so lane, pass, ownership, and method rules remain visible and binding during later work.

Implementation notes from the selected plan:

- create a live boundary record for the active task
- require every next action to check against that record before execution
- update the record when the human adds or narrows a boundary

Rollout:

1. Define a small ledger shape for lane, pass, ownership, wait rule, and hard constraints.
2. Add a pre-action check that compares proposed work against the active ledger.
3. Write back new human boundaries immediately when they appear.
4. Pilot the ledger on a task with pass framing and repo-local lane rules.
5. Expand to shared orchestration use after the repo-local pattern is stable.
