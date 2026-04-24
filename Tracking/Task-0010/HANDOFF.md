# Task 0010 Handoff

## Current Status

`Task-0010` is newly created and currently pending planning approval.

This task owns the bridge from Dream output into normal task lifecycle:

- daily Dream execution
- digest email
- collapsible option-task sections
- promote-to-task
- visibility to the future `Review` and `Tasks` tabs

The canonical promotion rule is:

- email and the future `Review` surface both use one shared backend promotion contract
- this task owns that contract and the generated task-skeleton output
- the system should not grow separate promotion semantics in different surfaces

## Current Objective

Turn Dream from a strong but manual research pipeline into a daily product workflow that produces reviewable candidate work and a real promotion path.

## Current Baseline

The Dream pipeline now exists durably through the work captured under [Task-0007](../Task-0007/TASK.md) and shared `.codex` docs.

What is missing is the product bridge:

- daily execution
- human triage email
- promotion into real tasks

## Next Recommended Step

Lock the daily output contract before implementation starts, especially:

- what the email shows by default
- how collapsible sections behave
- how promotion preserves provenance

## Watchouts

- do not confuse candidate work with approved work
- do not let promotion become lossy copy-paste
- do not bypass the future `Review`, `Tasks`, and dispatch surfaces

## References

- [TASK.md](./TASK.md)
- [PLAN.md](./PLAN.md)
- [Task-0005](../Task-0005/TASK.md)
- [Task-0007](../Task-0007/TASK.md)
- [Task-0011](../Task-0011/TASK.md)
- [Task-0009](../Task-0009/TASK.md)
- [Task-0008](../Task-0008/TASK.md)
