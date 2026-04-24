# Task 0008 Handoff

## Current Status

`Task-0008` is in implementation with `PASS-0002` active.

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

`PASS-0001` has now landed the first durable dispatch slice:

- `POST /api/v1/tasks/{task_id}/dispatch`
- `GET /api/v1/task-runs/{run_id}`
- Temporal-backed task-run workflow registration and query shape
- owned worktree allocation for simple task dispatch
- baseline-commit capture and initial restore-baseline capture
- backend smoke evidence in [Testing/PASS-0001-BACKEND-SMOKE-0001.md](./Testing/PASS-0001-BACKEND-SMOKE-0001.md)

`PASS-0001` now also has the first richer runtime-state slice after dispatch:

- `POST /api/v1/task-runs/{run_id}/state`
- Temporal signal handling for post-dispatch task-run state mutation
- task-level readback that reflects live active-run state updates
- live validation-lane proof in [Testing/PASS-0001-BACKEND-SMOKE-0002.md](./Testing/PASS-0001-BACKEND-SMOKE-0002.md)

`PASS-0001` is now complete enough for the next honest step to shift into supervision and intervention behavior.

What is still missing is supervision, poke, interrupt, cleanup behavior, and real task execution over those durable runs.

`PASS-0002` now has its first real supervision and intervention slice:

- read-through supervision that marks stale active runs as `sleeping_or_stalled`
- `POST /api/v1/task-runs/{run_id}/poke`
- `POST /api/v1/task-runs/{run_id}/interrupt`
- owned-lane restore-to-commit cleanup on interrupt
- task readback that releases live-story ownership after terminal runs
- live proof in [Testing/PASS-0002-BACKEND-SMOKE-0001.md](./Testing/PASS-0002-BACKEND-SMOKE-0001.md)

`PASS-0002` now also has the next supervision refinement slice:

- stale human-wait escalation into `human_wait_stale`
- explicit `waiting_for_human_stale` poke rejection while keeping interrupt allowed
- cleanup-blocked interrupt readback with:
  - `repo_lane.reset_status = cleanup_blocked`
  - `repo_lane.last_reset_target_commit`
  - `repo_lane.reset_failure_summary`
  - run-level `failure_summary`
- live proof in [Testing/PASS-0002-BACKEND-SMOKE-0002.md](./Testing/PASS-0002-BACKEND-SMOKE-0002.md)

`PASS-0002` now also has the first durable worker-follow-up loop for `poke`:

- `poke` creates a pending backend-worker follow-up on the run
- repeated `poke` is blocked while that follow-up is pending
- a fresh runtime update completes the follow-up durably
- live proof in [Testing/PASS-0002-BACKEND-SMOKE-0003.md](./Testing/PASS-0002-BACKEND-SMOKE-0003.md)

`PASS-0002` now also has the first cleanup-repair follow-up slice:

- `POST /api/v1/task-runs/{run_id}/retry-cleanup`
- cleanup-blocked runs can retry owned-lane restore through the backend instead of staying readback-only truth
- successful cleanup retry converts the run into terminal `interrupted` state with a pending `interrupt_review` follow-up
- cleanup repair and interrupt review now use the same durable `follow_up` envelope already used by `poke`
- live proof in [Testing/PASS-0002-BACKEND-SMOKE-0004.md](./Testing/PASS-0002-BACKEND-SMOKE-0004.md)

`PASS-0002` now also has the interrupt-review resolution slice:

- `POST /api/v1/task-runs/{run_id}/resolve-interrupt-review`
- pending `interrupt_review` blocks task-level redispatch until the review is resolved
- resolved interrupt review records an explicit `resolution` on the run
- resolved interrupt review returns the task to `dispatch_readiness.ready = true`
- Temporal-backed run updates now fall back to the closed workflow result when a successful resolution closes the workflow before the immediate query returns
- live proof in [Testing/PASS-0002-BACKEND-SMOKE-0005.md](./Testing/PASS-0002-BACKEND-SMOKE-0005.md)

What is still missing is real task execution over those durable runs and richer backend-owned worker behavior beyond manual state updates.

## Current Gate

Implementation is active under the approved backend-only runtime split:

- backend-owned task-run workflow and API contract
- first proof allowed through Codex or direct backend interactions before frontend work
- explicit inclusion of the backend capabilities the later [Task-0009](../Task-0009/TASK.md) `Tasks` tab depends on
- explicit exclusion of the human's shared primary worktree as the normal simple-execution lane
- explicit reset-to-recorded-commit cleanup semantics for owned execution lanes

## Next Recommended Step

Continue with `PASS-0002` by deepening the supervision surface before any frontend work starts.

The next implementation slice should:

- move beyond manual runtime state mutation toward real worker-side execution or backend-driven runtime progress
- keep task and run readback aligned with the declared-doc ingest and reconcile model
- prepare the runtime shape that later pass work can drive through real execution and recovery events
- keep [CONSTRAINTS.md](./CONSTRAINTS.md) current if the human adds new constraints

## Watchouts

- do not treat silence as success
- do not let context recovery remain a manual search workflow
- do not split runtime truth between backend and any client memory
- validation-lane runner restarts on `14318` can serve stale binaries or fail on stdout-log locks; use a clean manual listener when live proof needs trustworthy current code
- do not broaden this task into dashboard implementation work

## References

- [TASK.md](./TASK.md)
- [PLAN.md](./PLAN.md)
- [Task-0005](../Task-0005/TASK.md)
- [Task-0009](../Task-0009/TASK.md)
- [Task-0010](../Task-0010/TASK.md)
