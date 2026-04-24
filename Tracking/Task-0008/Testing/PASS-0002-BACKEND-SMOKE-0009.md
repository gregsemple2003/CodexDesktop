# PASS-0002 Backend Smoke 0009

## Type

Unit test / server-only smoke / live manual backend proof.

## Date

`2026-04-24`

## Scope

Proof for the next bounded `PASS-0002` workload-execution slice:

- after dispatch, preflight, and workload-step preparation, the task-run workflow actually executes the first prepared backend workload step inside the owned lane
- workload execution writes a durable result artifact next to the prepared workload packet inside `.codex-taskrun/`
- the run advances automatically to `workload_step_executed`
- task readback reflects the executed active-run story

## Commands

```powershell
cd C:\Agent\CodexDashboard\backend\orchestration
go test ./...

$env:CODEX_ORCH_POSTGRES_PORT='15432'
$env:CODEX_ORCH_TEMPORAL_PORT='17233'
$env:CODEX_ORCH_TEMPORAL_UI_PORT='18080'
docker compose --project-name codex-orchestration-validation -f .\dev\docker-compose.temporal-postgres.yml down -v
docker compose --project-name codex-orchestration-validation -f .\dev\docker-compose.temporal-postgres.yml up -d

$manualRoot = Join-Path $env:LOCALAPPDATA 'CodexDashboard\orchestration-validation-manual-15318'
New-Item -ItemType Directory -Force -Path $manualRoot | Out-Null
go build -o (Join-Path $manualRoot 'controlplane-manual.exe') .\cmd\controlplane

Start-Process powershell.exe -ArgumentList @(
  '-NoProfile',
  '-Command',
  "$env:CODEX_ORCHESTRATION_BIND_ADDRESS='127.0.0.1:15318'; " +
  "$env:CODEX_ORCHESTRATION_TEMPORAL_ADDRESS='127.0.0.1:17233'; " +
  "$env:CODEX_ORCHESTRATION_WORKTREE_ROOT='C:\Agent\CodexDashboard'; " +
  "$env:CODEX_ORCHESTRATION_RUNS_ROOT='C:\Users\gregs\AppData\Local\CodexDashboard\orchestration-runs-validation'; " +
  "& '" + (Join-Path $manualRoot 'controlplane-manual.exe') + "'"
) -WorkingDirectory 'C:\Agent\CodexDashboard\backend\orchestration'

$base = 'http://127.0.0.1:15318'
$dispatch = Invoke-RestMethod -Method Post -Uri "$base/api/v1/tasks/Task-0008/dispatch"

$run = $null
for ($i = 0; $i -lt 20; $i++) {
  Start-Sleep -Milliseconds 500
  $candidate = Invoke-RestMethod -Uri "$base/api/v1/task-runs/taskrun--Task-0008--active"
  $run = $candidate
  if ($candidate.state_envelope.reason_code -eq 'workload_step_executed') { break }
}

$taskAfter = Invoke-RestMethod -Uri "$base/api/v1/tasks/Task-0008"
$artifact = Get-Content -Raw $run.repo_lane.workload_result_path | ConvertFrom-Json
```

## Result

Pass.

`go test ./...` passed with the added coverage for:

- consuming the prepared workload packet
- writing a workload execution result artifact inside the owned lane
- surfacing the execution result path durably on the run

The live manual backend proof on the clean `127.0.0.1:15318` listener then proved:

1. `POST /api/v1/tasks/Task-0008/dispatch` still returned the initial bootstrap state:
   - `reason_code = owned_lane_bootstrapped`
2. A follow-up `GET /api/v1/task-runs/taskrun--Task-0008--active` then advanced automatically to:
   - `state_envelope.state = running`
   - `reason_code = workload_step_executed`
   - `next_owner = backend_worker`
   - `last_progress_summary = Executed the first backend workload step inside the owned lane.`
3. The run exposed:
   - `repo_lane.workload_step_path`
   - `repo_lane.workload_result_path`
4. The result artifact existed on disk next to the workload packet inside the owned lane under:
   - `.codex-taskrun/taskrun--Task-0008--active/workload-step-0001.result.json`
5. The workload result recorded:
   - the executed workload instruction
   - the current commit
   - an explicit execution summary
6. The next task read returned:
   - `state_envelope.state = running`
   - `reason_code = workload_step_executed`
   - `current_story.status = active_run`

## Why This Counts

This is the first slice where the backend not only prepares a workload packet inside the owned lane, but also executes that prepared step and records the result there.

The runtime now advances through:

- owned-lane bootstrap
- execution preflight
- workload-step preparation
- workload-step execution

without any manual `POST /state` mutation.

That is the smallest honest step from prepared backend work toward richer real execution behavior inside the owned lane.

## Limitations

- The executed step is still a bounded backend-owned runtime step, not full task fulfillment.
- The disposable validation lane still must be reset with `docker compose ... down -v` after workflow-shape changes because the active workflow id is intentionally reused.
- This is still backend proof, not dashboard regression proof.
