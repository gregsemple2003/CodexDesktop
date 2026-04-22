# Problem 0004 Option B Plan 0001

## Intent

Require durable task-state updates before each milestone message so task history stays current with live work.

## What Changes

- Move state updates earlier in the milestone sequence.
- Block milestone summaries when the durable state still reflects an old phase.
- Keep pass status, blockers, and current lane honest in task artifacts.

## Files Or Artifact Types That Move

- `TASK-STATE.json`, `HANDOFF.md`, pass checklists, and related task-owned state files.
- Shared lifecycle prompts that order state updates before milestone messages.
- Small conformance checks for state freshness.

## Rollout

1. Define which state fields must be current before a milestone can be sent.
2. Add the freshness check to the milestone workflow.
3. Pilot it on a task with frequent reopen and pass transitions.
4. Review whether state drift drops enough to justify the extra step.

## Success Check

- Milestone messages no longer get ahead of durable state.
- Reopened work is reflected in task artifacts before more work proceeds.
- The human can trust the task state without cross-checking chat history.

## Burden Reduction Under Directional Context

`Truth`: the durable task record matches the real task phase.

`Compassion`: it reduces the human's need to demand state sync after the fact.

`Tolerance`: the workflow turns short reopen commands into immediate state changes rather than letting them fade.
