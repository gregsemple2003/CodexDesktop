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

`PASS-0002` now also has the first backend-owned execution-bootstrap slice:

- dispatch now bootstraps the owned checkout before the run starts
- dispatch captures `repo_lane.current_commit`
- dispatch writes an owned-lane bootstrap artifact under the backend run-artifact root
- the created run starts in backend-produced `running` state with `reason_code = owned_lane_bootstrapped`
- live proof in [Testing/PASS-0002-BACKEND-SMOKE-0006.md](./Testing/PASS-0002-BACKEND-SMOKE-0006.md)

`PASS-0002` now also has the first worker-side execution preflight slice:

- after dispatch, the Temporal task-run workflow runs an execution preflight activity against the owned checkout
- the preflight resolves the owned task root inside the owned lane
- the preflight writes `execution-preflight.json` under the backend run-artifact root
- the run advances automatically to `reason_code = execution_preflight_complete` without `POST /state`
- live proof in [Testing/PASS-0002-BACKEND-SMOKE-0007.md](./Testing/PASS-0002-BACKEND-SMOKE-0007.md)

`PASS-0002` now also has the first actual workload-step preparation slice:

- after preflight, the task-run workflow writes the first backend workload packet inside the owned lane under `.codex-taskrun/`
- the run advances automatically to `reason_code = workload_step_prepared`
- task readback reflects that workload-prepared active story without `POST /state`
- live proof in [Testing/PASS-0002-BACKEND-SMOKE-0008.md](./Testing/PASS-0002-BACKEND-SMOKE-0008.md)

`PASS-0002` now also has the first workload-step execution slice:

- after workload-step preparation, the task-run workflow executes that prepared step inside the owned lane
- execution writes `workload-step-0001.result.json` next to the prepared workload packet
- the run advances automatically to `reason_code = workload_step_executed`
- task readback reflects the executed active-run story without `POST /state`
- live proof in [Testing/PASS-0002-BACKEND-SMOKE-0009.md](./Testing/PASS-0002-BACKEND-SMOKE-0009.md)

`PASS-0002` now also has the first real task-specific backend execution slice:

- the prepared workload packet for `Task-0008` now carries a concrete backend command plan instead of only a generic instruction
- the task-run workflow now executes focused Task-0008 backend validation from the owned lane:
  - `go test ./internal/taskexec ./internal/taskrun`
- task readback now advances automatically to:
  - `reason_code = task_0008_backend_validation_complete`
- the workload result artifact records:
  - `execution_kind = task_0008_backend_validation`
  - the exact command
  - stdout and stderr artifact paths
- live proof in [Testing/PASS-0002-BACKEND-SMOKE-0010.md](./Testing/PASS-0002-BACKEND-SMOKE-0010.md)

`PASS-0002` now also has the first task-specific owned-lane repo-state mutation:

- after Task-0008 validation, the workflow writes:
  - `Tracking/Task-0008/OwnedLane/IMPLEMENTATION-BRIEF.md`
  inside the owned lane
- task readback now advances automatically to:
  - `reason_code = task_0008_owned_lane_brief_written`
- run readback now exposes:
  - `repo_lane.workload_output_path`
- the workload result artifact records the owned-lane git status after that mutation
- live proof in [Testing/PASS-0002-BACKEND-SMOKE-0011.md](./Testing/PASS-0002-BACKEND-SMOKE-0011.md)

What is still missing is moving from owned-lane task-artifact mutation into the first task-specific owned-lane code mutation or worker-applied implementation change.

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

- move from owned-lane task-artifact mutation into the first task-specific owned-lane code mutation or worker-applied implementation change
- keep task and run readback aligned with the declared-doc ingest and reconcile model
- prepare the runtime shape that later pass work can drive through real execution and recovery events
- keep [CONSTRAINTS.md](./CONSTRAINTS.md) current if the human adds new constraints

## Watchouts

- do not treat silence as success
- do not let context recovery remain a manual search workflow
- do not split runtime truth between backend and any client memory
- validation-lane runner restarts on `14318` can serve stale binaries or fail on stdout-log locks; use a clean manual listener when live proof needs trustworthy current code
- when replaying the fixed active task-run id after workflow-shape changes, reset the disposable validation Temporal volume or the proof lane will correctly fail on old workflow history
- after a fresh validation-volume reset, a clean manual listener may need a short Temporal warm-up delay before backend startup or it can fail with `error reading server preface: EOF`
- do not mistake owned-lane task-artifact mutation for finished implementation work; it is only the first repo-state change in the bounded task-specific worker path
- do not broaden this task into dashboard implementation work

## References

- [TASK.md](./TASK.md)
- [PLAN.md](./PLAN.md)
- [Task-0005](../Task-0005/TASK.md)
- [Task-0009](../Task-0009/TASK.md)
- [Task-0010](../Task-0010/TASK.md)
