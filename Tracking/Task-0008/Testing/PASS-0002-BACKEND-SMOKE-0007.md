# PASS-0002 Backend Smoke 0007

## Type

Unit test / server-only smoke / live manual backend proof.

## Date

`2026-04-24`

## Scope

Proof for the next bounded `PASS-0002` worker-side execution slice:

- after dispatch, the Temporal task-run workflow performs an execution preflight activity
- the preflight runs against the owned checkout instead of relying on `POST /state`
- the preflight writes a durable `execution-preflight.json` artifact under the backend run-artifact root
- the run advances automatically from dispatch bootstrap into `execution_preflight_complete`
- task readback reflects the updated active-run story

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
  if ($candidate.state_envelope.reason_code -eq 'execution_preflight_complete') { break }
}

$taskAfter = Invoke-RestMethod -Uri "$base/api/v1/tasks/Task-0008"
$artifact = Get-Content -Raw $run.repo_lane.preflight_artifact_path | ConvertFrom-Json
```

## Result

Pass.

`go test ./...` passed with the added coverage for:

- owned-lane execution preflight artifact writing
- owned task-root resolution inside the owned checkout
- automatic worker-side progress after dispatch without a manual state mutation

The live manual backend proof on the clean `127.0.0.1:15318` listener then proved:

1. `POST /api/v1/tasks/Task-0008/dispatch` still returned the initial bootstrap state:
   - `reason_code = owned_lane_bootstrapped`
2. A follow-up `GET /api/v1/task-runs/taskrun--Task-0008--active` then advanced automatically to:
   - `state_envelope.state = running`
   - `reason_code = execution_preflight_complete`
   - `next_owner = backend_worker`
   - `last_progress_summary = Execution preflight inspected the owned task docs and recorded owned-lane readiness.`
3. The run exposed `repo_lane.preflight_artifact_path`, and that file existed on disk.
4. The `execution-preflight.json` artifact recorded:
   - the owned task root inside the owned checkout
   - the current commit
   - `TASK.md`, `PLAN.md`, `HANDOFF.md`, and `TASK-STATE.json` present inside the owned task root
5. The next task read returned:
   - `state_envelope.state = running`
   - `reason_code = execution_preflight_complete`
   - `current_story.status = active_run`

## Why This Counts

This is the first true worker-side workload step after dispatch.

The backend now does more than:

- provision an owned checkout
- bootstrap repo metadata

It also runs a Temporal-driven execution preflight against the owned task root and records the result durably without needing a manual `/state` call.

That is the next smallest honest step toward real backend-driven workload execution in the owned lane.

## Limitations

- The preflight still prepares execution rather than running the actual task workload.
- Because Task-0008 uses a fixed active workflow id, the disposable validation lane must be reset with `docker compose ... down -v` when replaying the same task-run id across workflow-shape changes. Without that reset, Temporal correctly reports deterministic mismatch on old history.
- This is still backend proof, not dashboard regression proof.
