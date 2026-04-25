# Task 0008 Handoff

## Current Status

`Task-0008` is complete.

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

The shipped backend contract now includes:

- durable dispatch and task-run persistence
- supervision, poke, interrupt, cleanup retry, and workload retry
- exclusive owned-lane execution with restore-to-commit semantics
- deep-context readback for operators and later clients
- declared-doc drift reconcile with explicit git-vs-runtime divergence reporting

Final closeout evidence is in [Testing/PASS-0004-BACKEND-SMOKE-0001.md](./Testing/PASS-0004-BACKEND-SMOKE-0001.md).

## Current Objective

No further implementation remains in Task-0008.

This handoff is now the final backend baseline for downstream consumers such as [Task-0009](../Task-0009/TASK.md).

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

`PASS-0002` now also has the first task-specific owned-lane code mutation:

- after Task-0008 validation and brief generation, the workflow writes:
  - `backend/orchestration/internal/taskexec/task0008_owned_lane_generated.go`
  inside the owned lane
- task readback now advances automatically to:
  - `reason_code = task_0008_owned_lane_code_written`
- run readback now exposes:
  - `repo_lane.workload_code_path`
- the workload result artifact now records scoped post-execution git status for only the owned-lane outputs this slice wrote
- live proof in [Testing/PASS-0002-BACKEND-SMOKE-0012.md](./Testing/PASS-0002-BACKEND-SMOKE-0012.md)

`PASS-0002` now also has the first worker-applied edit of an existing Task-0008 implementation file:

- after Task-0008 validation and brief generation, the workflow now edits the existing owned-lane implementation path:
  - `backend/orchestration/internal/taskexec/taskexec.go`
- the edited owned lane still passes:
  - `go test ./internal/taskexec ./internal/taskrun`
- task readback now advances automatically to:
  - `reason_code = task_0008_existing_file_edited`
- run readback continues to expose:
  - `repo_lane.workload_code_path`
  but the durable code path now points at an existing implementation file rather than a generated side file
- the workload result artifact records scoped owned-lane git status including:
  - `M backend/orchestration/internal/taskexec/taskexec.go`
  - `?? .codex-taskrun/`
  - `?? Tracking/Task-0008/OwnedLane/`
- live proof in [Testing/PASS-0002-BACKEND-SMOKE-0013.md](./Testing/PASS-0002-BACKEND-SMOKE-0013.md)

`PASS-0002` now also has the first worker-applied implementation change in an existing Task-0008 file that alters runtime behavior:

- after Task-0008 validation and brief generation, the workflow now edits the existing owned-lane implementation path:
  - `backend/orchestration/internal/taskexec/taskexec.go`
- that edit now changes the bootstrapped-run suspiciousness window in `InitialView` from:
  - `15 minutes`
  to:
  - `5 minutes`
- the workflow now runs an owned-lane behavior probe after the code edit and records:
  - `behavior_probe_path`
  in the workload result artifact
- task readback now advances automatically to:
  - `reason_code = task_0008_existing_file_behavior_changed`
- the live proof verified:
  - the owned-lane edited file contains the real code change
  - the owned-lane behavior probe reported `suspicious_window_minutes = 5`
- live proof in [Testing/PASS-0002-BACKEND-SMOKE-0014.md](./Testing/PASS-0002-BACKEND-SMOKE-0014.md)

`PASS-0002` now also has the next worker-applied existing-file change on a later Task-0008 runtime path:

- after Task-0008 validation and brief generation, the workflow now edits the existing owned-lane implementation path:
  - `backend/orchestration/internal/taskrun/service.go`
- that edit shortens the interrupt-review follow-up window in the owned lane from:
  - `24 hours`
  to:
  - `2 hours`
- the workflow now runs an owned-lane behavior probe that exercises `InterruptRun` and records:
  - `behavior_probe_path`
  in the workload result artifact
- task readback now advances automatically to:
  - `reason_code = task_0008_interrupt_review_window_changed`
- the live proof verified:
  - the owned-lane edited file contains the real `DueAt: now.Add(2 * time.Hour)` change
  - the owned-lane behavior probe reported:
    - `reason_code = interrupt_requested`
    - `follow_up_kind = interrupt_review`
    - `due_window_hours = 2`
    - `reset_status = restored`
- live proof in [Testing/PASS-0002-BACKEND-SMOKE-0015.md](./Testing/PASS-0002-BACKEND-SMOKE-0015.md)

`PASS-0002` now also has the next worker-applied existing-file change on a broader redispatch path:

- after Task-0008 validation and brief generation, the workflow now edits the existing owned-lane implementation path:
  - `backend/orchestration/internal/taskrun/service.go`
- that edit changes redispatch behavior so a fresh dispatch releases the previous terminal owned lane before provisioning a new one
- the workflow now runs an owned-lane behavior probe that exercises redispatch after a resolved terminal run and records:
  - `behavior_probe_path`
  in the workload result artifact
- task readback now advances automatically to:
  - `reason_code = task_0008_redispatch_lane_released`
- the live proof verified:
  - the owned-lane edited file contains the real `releasePreviousOwnedLane` dispatch call
  - the owned-lane behavior probe reported:
    - `original_owned_root_removed = true`
    - `new_owned_root_exists = true`
    - `new_owned_root_differs = true`
    - `new_run_reason_code = owned_lane_bootstrapped`
- live proof in [Testing/PASS-0002-BACKEND-SMOKE-0016.md](./Testing/PASS-0002-BACKEND-SMOKE-0016.md)

`PASS-0002` now also has the next broader recovery-policy slice on the real backend interrupt-review path:

- resolving `interrupt_review` now immediately releases the resolved owned lane instead of deferring cleanup until a later redispatch
- resolved runs now persist:
  - `repo_lane.reset_status = released`
  - `repo_lane.last_reset_target_commit`
  - a cleared live `repo_lane.owned_repo_root`
- owned-lane cleanup now removes git worktrees with Windows long-path support enabled so review-resolution cleanup does not fail on long artifact paths
- live proof in [Testing/PASS-0002-BACKEND-SMOKE-0017.md](./Testing/PASS-0002-BACKEND-SMOKE-0017.md)

What is still missing is repairing or replacing the current live Task-0008 owned-lane workload-execution path so later progress does not stall at `workload_execution_failed` before interruption or recovery is exercised.

`PASS-0002` now also has the bounded repair slice on the current live Task-0008 workload-execution path:

- the stale Task-0008 owned-lane mutation recipe was repaired so the owned-lane `go test` step no longer stalls at `workload_execution_failed`
- the owned-lane existing-file edit is now retargeted to the current `backend/orchestration/internal/taskrun/service.go` baseline
- that owned-lane edit now escalates blocked-run recovery attention from `needs_attention` to `urgent`
- the repaired workload step now writes its result artifact again and advances automatically to:
  - `reason_code = task_0008_workload_failure_attention_escalated`
- live proof in [Testing/PASS-0002-BACKEND-SMOKE-0018.md](./Testing/PASS-0002-BACKEND-SMOKE-0018.md)

What is still missing is a real backend recovery action for actual `workload_execution_failed` runs, rather than only a repaired owned-lane recipe plus stronger blocked-run attention.

`PASS-0002` now also has that real backend recovery action for actual `workload_execution_failed` runs:

- `POST /api/v1/task-runs/{run_id}/retry-workload`
- workload retry is only allowed from:
  - `state = blocked`
  - `reason_code = workload_execution_failed`
- blocked workload-execution failures now expose:
  - `follow_up.kind = workload_recovery`
  - `follow_up.status = pending`
- retry releases the failed owned lane, provisions and bootstraps a fresh owned lane, and signals the active Temporal workflow to rerun the owned-lane execution path
- the active run advances through:
  - `reason_code = workload_retry_requested`
  before re-entering the repaired Task-0008-specific execution path
- retry clears the pending `workload_recovery` follow-up as the fresh owned lane starts
- live proof in [Testing/PASS-0002-BACKEND-SMOKE-0019.md](./Testing/PASS-0002-BACKEND-SMOKE-0019.md)

What is still missing is a less synthetic way to prove workload-execution failure recovery than seeding `workload_execution_failed` through the backend update path. The repaired normal live path no longer fails naturally, so later work should decide whether to preserve a bounded fault-injection hook or make a naturally failing recovery case reproducible without regressing the happy path.

`PASS-0002` now also removes that `/state` seeding dependency for proof:

- `POST /api/v1/tasks/{task_id}/dispatch-workload-failure-exercise`
- for `Task-0008`, that bounded dispatch path starts a one-shot backend-owned workload-failure exercise instead of relying on a synthetic `POST /state` mutation
- the workload packet records:
  - `failure_mode = task_0008_workload_execution_failure_once`
- the run now reaches:
  - `reason_code = workload_execution_failed`
  naturally through the workflow and activity path
- the run exposes:
  - `follow_up.kind = workload_recovery`
  - `follow_up.status = pending`
- `POST /api/v1/task-runs/{run_id}/retry-workload` then clears the follow-up and reruns the repaired happy path back to:
  - `reason_code = task_0008_workload_failure_attention_escalated`
- live proof in [Testing/PASS-0002-BACKEND-SMOKE-0020.md](./Testing/PASS-0002-BACKEND-SMOKE-0020.md)

The remaining gap is no longer proof honesty for workload-failure recovery. The next honest question is whether to keep this bounded failure-exercise path as the durable proof/debug hook, or to move on to the next broader recovery or execution-policy gap in `PASS-0002`.

`PASS-0003` now has its first real context-readback slice:

- task readback now exposes `deep_context` even when no run is active
- run readback now exposes `deep_context` with launchable targets for:
  - task folder
  - task handoff
  - owned checkout
  - run artifacts
  - active run API resource
- if the dispatching process can see a session id or transcript path, that best-effort provenance is captured into the run and surfaced in the same `deep_context` contract
- live proof in [Testing/PASS-0003-BACKEND-SMOKE-0001.md](./Testing/PASS-0003-BACKEND-SMOKE-0001.md)

`PASS-0003` now also has the declared-doc drift and divergence-reconcile slice:

- active task readback now performs read-through reconcile and records the newer declared snapshot durably in the active run
- run readback now exposes:
  - `doc_runtime_divergence_status = reconciled`
  - `doc_runtime_divergence_summary` with old-to-new declared task revisions
- the active run keeps its live runtime story while declared docs drift
- live proof in [Testing/PASS-0003-BACKEND-SMOKE-0002.md](./Testing/PASS-0003-BACKEND-SMOKE-0002.md)

`PASS-0004` closeout is now complete:

- full backend unit coverage passed again through `go test ./...`
- Task-0008 task-owned proof now covers dispatch, supervision, interrupt, retry, deep-context, and divergence behavior end to end
- repo-root regression is honestly `not_applicable` because this task changed backend orchestration only and did not change the desktop app surface
- final closeout evidence is in [Testing/PASS-0004-BACKEND-SMOKE-0001.md](./Testing/PASS-0004-BACKEND-SMOKE-0001.md)

## Current Gate

Closure is complete.

Task-0008 shipped the approved backend-only runtime split:

- backend-owned task-run workflow and API contract
- first proof allowed through Codex or direct backend interactions before frontend work
- explicit inclusion of the backend capabilities the later [Task-0009](../Task-0009/TASK.md) `Tasks` tab depends on
- explicit exclusion of the human's shared primary worktree as the normal simple-execution lane
- explicit reset-to-recorded-commit cleanup semantics for owned execution lanes

## Next Recommended Step

No more implementation remains in Task-0008.

The next honest work is downstream consumption of this contract, primarily in [Task-0009](../Task-0009/TASK.md), rather than more Task-0008-owned runtime expansion.

## Watchouts

- do not treat silence as success
- do not let context recovery remain a manual search workflow
- do not split runtime truth between backend and any client memory
- validation-lane runner restarts on `14318` can serve stale binaries or fail on stdout-log locks; use a clean manual listener when live proof needs trustworthy current code
- when replaying the fixed active task-run id after workflow-shape changes, reset the disposable validation Temporal volume or the proof lane will correctly fail on old workflow history
- after a fresh validation-volume reset, a clean manual listener may need a short Temporal warm-up delay before backend startup or it can fail with `error reading server preface: EOF`
- when starting the validation compose stack directly from `backend/orchestration`, set the validation-lane port overrides explicitly or Postgres can collide with the service lane on `5432`
- do not mistake owned-lane task-artifact mutation for finished implementation work; it is only the first repo-state change in the bounded task-specific worker path
- do not mistake a bounded owned-lane recovery improvement for finished implementation work; the next honest step is to fix the current live workload-execution gap before piling on more synthetic owned-lane proof edits
- the natural workload-failure proof path is intentionally bounded to `Task-0008` and uses a one-shot backend-owned execution directive rather than a generalized fault-injection surface
- declared-doc drift proof on an active run is signal-based; live verification should poll the active run until reconcile becomes visible rather than assuming the immediate task-read response already reflects it
- do not let Task-0008-owned mutation recipes drift behind the real repo baseline or the owned-lane validation step will correctly fail before later proof can run
- do not broaden this task into dashboard implementation work

## References

- [TASK.md](./TASK.md)
- [PLAN.md](./PLAN.md)
- [Task-0005](../Task-0005/TASK.md)
- [Task-0009](../Task-0009/TASK.md)
- [Task-0010](../Task-0010/TASK.md)
