# Task 0008 Handoff

## Current Status

`Task-0008` is in implementation with `PASS-0001` active.

This task is a backend-only runtime task:

- dispatch
- durable task-run state
- supervision
- poke
- interrupt
- deep-context provenance readback
- exclusive repo-lane ownership for simple execution
- restore-to-commit cleanup semantics for proof and reset

The first honest operator path can be Codex or direct backend interaction.

This task does not own frontend controls or dashboard wiring.

## Current Objective

Turn `task dispatch` from a vague future UI action into a real durable runtime capability with a state contract strong enough to support humane monitoring.

The initial contract note is here:

- [Design/DURABLE-EXECUTION-STATE-CONTRACT.md](./Design/DURABLE-EXECUTION-STATE-CONTRACT.md)

The runtime shape is intentionally frozen here:

- task dispatch is a separate backend-owned task-run workflow and API contract
- it is not modeled as another Git-tracked recurring job spec

Task-0009's `Tasks`-tab design has now been inspected as a downstream consumer of this runtime.

That consumer expects Task-0008 to provide backend truth for:

- task and active-run readback
- durable state distinctions plus reason inputs
- freshness and staleness signals
- next expected step and owner
- bounded action gating for `Dispatch`, `Poke`, and `Interrupt`
- deep-context launch provenance for `Open Thread`
- an exclusive backend-owned execution checkout or equivalent isolated repo lane
- a recorded useful restore commit for unit proof and execution cleanup

## Current Baseline

The repo already has:

- Temporal-backed job orchestration from [Task-0005](../Task-0005/TASK.md)
- a planned `Tasks` tab surface from [Task-0009](../Task-0009/TASK.md)

`PASS-0000` landed the first executable task-readback slice:

- task-run contract types under `backend/orchestration/internal/taskrun/`
- declared-doc parsing for `Tracking/Task-<id>/`
- `GET /api/v1/tasks`
- `GET /api/v1/tasks/{task_id}`
- unit coverage plus backend smoke evidence in [Testing/PASS-0000-BACKEND-SMOKE-0001.md](./Testing/PASS-0000-BACKEND-SMOKE-0001.md)

What is still missing is durable dispatch and Temporal-backed task-run persistence that tie readback to real owned execution.

## Current Gate

Implementation is active under the approved backend-only runtime split:

- backend-owned task-run workflow and API contract
- first proof allowed through Codex or direct backend interactions before frontend work
- explicit inclusion of the backend capabilities the later [Task-0009](../Task-0009/TASK.md) `Tasks` tab depends on
- explicit exclusion of the human's shared primary worktree as the normal simple-execution lane
- explicit reset-to-recorded-commit cleanup semantics for owned execution lanes

## Next Recommended Step

Implement `PASS-0001` and turn readback into a real runtime capability before any frontend work starts.

The next implementation slice should:

- add real dispatch entrypoints
- persist task runs durably through Temporal-backed runtime state
- capture an owned checkout or owned-lane identity for simple execution
- capture the dispatch baseline commit before the run mutates that owned lane
- keep task and run readback aligned with the declared-doc ingest and reconcile model
- keep [CONSTRAINTS.md](./CONSTRAINTS.md) current if the human adds new constraints

## Watchouts

- do not treat silence as success
- do not let context recovery remain a manual search workflow
- do not split runtime truth between backend and any client memory
- do not broaden this task into dashboard implementation work

## References

- [TASK.md](./TASK.md)
- [PLAN.md](./PLAN.md)
- [Task-0005](../Task-0005/TASK.md)
- [Task-0009](../Task-0009/TASK.md)
- [Task-0010](../Task-0010/TASK.md)
