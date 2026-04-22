# Problem 0003 Option A Plan 0001

## Intent

Keep an active boundary ledger for the task so lane, pass, ownership, and method rules remain visible and binding during later work.

## What Changes

- Create a live boundary record for the active task.
- Require every next action to check against that record before execution.
- Update the record when the human adds or narrows a boundary.

## Files Or Artifact Types That Move

- Task-state or session-state boundary records.
- Shared workflow prompts that require a boundary check before action.
- Review artifacts that show which boundaries were in force when a step was taken.

## Rollout

1. Define a small ledger shape for lane, pass, ownership, wait rule, and hard constraints.
2. Add a pre-action check that compares proposed work against the active ledger.
3. Write back new human boundaries immediately when they appear.
4. Pilot the ledger on a task with pass framing and repo-local lane rules.
5. Expand to shared orchestration use after the repo-local pattern is stable.

## Success Check

- Previously stated boundaries remain visible later in the pass.
- The system stops drifting back into disallowed lanes.
- The human no longer needs to restate the same `ThirdPerson` boundaries multiple times.

## Burden Reduction Under Directional Context

`Truth`: the working state matches the latest human-stated reality.

`Compassion`: it prevents repeated re-scoping and stop commands.

`Tolerance`: the ledger treats short corrections as durable state updates, not as one-turn noise.
