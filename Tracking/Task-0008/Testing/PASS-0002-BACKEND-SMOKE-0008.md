# PASS-0002 Backend Smoke 0008

## Type

Unit test / server-only smoke / live manual backend proof.

## Date

`2026-04-24`

## Scope

Proof for the next bounded `PASS-0002` workload-execution slice:

- after dispatch and preflight, the task-run workflow prepares the first actual backend-driven workload step inside the owned lane
- the workload step writes a backend-owned execution packet under `.codex-taskrun/` inside the owned checkout
- the run advances automatically to `workload_step_prepared`
- task readback reflects the active run story at that workload-prepared state

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
  if ($candidate.state_envelope.reason_code -eq 'workload_step_prepared') { break }
}

$taskAfter = Invoke-RestMethod -Uri "$base/api/v1/tasks/Task-0008"
$artifact = Get-Content -Raw $run.repo_lane.workload_step_path | ConvertFrom-Json
```

## Result

Pass.

`go test ./...` passed with the added coverage for:

- writing the first backend workload step packet under `.codex-taskrun/` inside the owned lane
- preserving the current commit and preflight linkage in that workload packet

The live manual backend proof on the clean `127.0.0.1:15318` listener then proved:

1. `POST /api/v1/tasks/Task-0008/dispatch` still returned the initial dispatch bootstrap:
   - `reason_code = owned_lane_bootstrapped`
2. A follow-up `GET /api/v1/task-runs/taskrun--Task-0008--active` then advanced automatically to:
   - `state_envelope.state = running`
   - `reason_code = workload_step_prepared`
   - `next_owner = backend_worker`
   - `last_progress_summary = Prepared the first backend workload step inside the owned lane.`
3. The run exposed `repo_lane.workload_step_path`, and that file existed inside the owned checkout under:
   - `.codex-taskrun/taskrun--Task-0008--active/workload-step-0001.json`
4. The workload packet recorded:
   - a backend workload instruction
   - the current commit
   - linkage back to the execution-preflight artifact
5. The next task read returned:
   - `state_envelope.state = running`
   - `reason_code = workload_step_prepared`
   - `current_story.status = active_run`

## Why This Counts

This is the first slice that prepares a real backend-owned workload artifact inside the owned checkout rather than only inspecting the owned lane or writing external run artifacts.

The workflow now advances through:

- owned-lane bootstrap
- execution preflight
- first workload-step preparation

without relying on a manual `POST /state` mutation.

That is the smallest honest step from preparation-only behavior toward actual backend-driven execution inside the owned lane.

## Limitations

- The workload step currently prepares the first execution packet rather than running the full task workload.
- Because Task-0008 uses a fixed active workflow id, the disposable validation lane must still be reset with `docker compose ... down -v` after workflow-shape changes or Temporal will correctly reject replay with deterministic mismatch.
- This is still backend proof, not dashboard regression proof.
