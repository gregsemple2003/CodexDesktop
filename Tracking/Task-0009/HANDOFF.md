# Task 0009 Handoff

## Current Status

`Task-0009` is newly created and currently pending planning approval.

This task owns the human-facing `Tasks` tab for committed work in CodexDashboard:

- the high-level view
- committed-task triage
- dispatch intent
- monitoring
- recovery affordances
- deep-context launch

This task does not own intake review, the promotion contract, the dispatch runtime, or the daily Dream automation.

Those are split out into:

- [Task-0011](../Task-0011/TASK.md)
- [Task-0008](../Task-0008/TASK.md)
- [Task-0010](../Task-0010/TASK.md)

If the `Tasks` tab later shows work that was promoted out of `Review`, it should consume durable provenance from the single promotion contract owned by [Task-0010](../Task-0010/TASK.md). This task does not define candidate review or promotion semantics.

## Current Objective

Lock the first humane product surface for `Tasks` before implementation hardens around backend convenience.

The design work for that lives in:

- [Design/HUMAN-NEED-AND-TASKS-TAB-DIRECTION.md](./Design/HUMAN-NEED-AND-TASKS-TAB-DIRECTION.md)
- [Design/STITCH-PROMPT.md](./Design/STITCH-PROMPT.md)

The review-surface design and research now live under [Task-0011](../Task-0011/HANDOFF.md).

The product split is now explicit:

- `Review` for incoming asks
- `Tasks` for committed work

## Current Baseline

The repo currently has:

- the token-velocity overlay
- the backend-backed `Jobs` tab
- task-owned markdown artifacts under `Tracking/`

What it does not yet have is one humane surface that ties together:

- real tasks
- dispatch state
- sleeping runs
- waiting-for-human committed work
- promoted-task provenance

That missing surface is why this task exists.

## Next Recommended Step

Review the task-local design brief and decide whether the first committed-work implementation slice should:

- start read-only
- or ship immediately with bounded dispatch actions once [Task-0008](../Task-0008/TASK.md) is ready enough

## Watchouts

- do not let the `Tasks` tab become a raw orchestration inspector
- do not hide task provenance
- do not let intake-review material leak back into `Tasks`; that belongs in [Task-0011](../Task-0011/TASK.md)
- do not make the human read markdown to understand the default screen

## References

- [TASK.md](./TASK.md)
- [PLAN.md](./PLAN.md)
- [Task-0005](../Task-0005/TASK.md)
- [Task-0007](../Task-0007/TASK.md)
- [Task-0008](../Task-0008/TASK.md)
- [Task-0010](../Task-0010/TASK.md)
- [Task-0011](../Task-0011/TASK.md)
- [Design/GENERAL-DESIGN.md](../../Design/GENERAL-DESIGN.md)
