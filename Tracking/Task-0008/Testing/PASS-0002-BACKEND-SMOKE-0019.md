# PASS-0002 Backend Smoke 0019

## Type

Unit test / server-only smoke / live manual backend proof.

## Date

`2026-04-24`

## Scope

Proof for the bounded recovery slice that adds a real backend action for actual `workload_execution_failed` runs:

- add `POST /api/v1/task-runs/{run_id}/retry-workload`
- require that retry is only valid from:
  - `state = blocked`
  - `reason_code = workload_execution_failed`
- release the failed owned lane
- provision and bootstrap a fresh owned lane
- signal the active Temporal task-run workflow to rerun the owned-lane workload path

## Commands

```powershell
cd C:\Agent\CodexDashboard\backend\orchestration
gofmt -w .\internal\httpapi\mux.go .\internal\httpapi\mux_test.go .\internal\taskexec\taskexec.go .\internal\taskrun\service.go .\internal\taskrun\service_test.go .\internal\taskrun\types.go .\internal\temporalbackend\backend.go
go test ./...

$env:CODEX_ORCH_POSTGRES_PORT='15432'
$env:CODEX_ORCH_TEMPORAL_PORT='17233'
$env:CODEX_ORCH_TEMPORAL_UI_PORT='18080'
docker compose --project-name codex-orchestration-validation -f .\dev\docker-compose.temporal-postgres.yml down -v
docker compose --project-name codex-orchestration-validation -f .\dev\docker-compose.temporal-postgres.yml up -d

Start-Sleep -Seconds 10

$manualRoot = Join-Path $env:LOCALAPPDATA 'CodexDashboard\orchestration-validation-manual-15318'
New-Item -ItemType Directory -Force -Path $manualRoot | Out-Null
go build -o (Join-Path $manualRoot 'controlplane-manual.exe') .\cmd\controlplane

$envMap = @{
  CODEX_ORCHESTRATION_BIND_ADDRESS     = '127.0.0.1:15318'
  CODEX_ORCHESTRATION_TEMPORAL_ADDRESS = '127.0.0.1:17233'
  CODEX_ORCHESTRATION_WORKTREE_ROOT    = 'C:\Agent\CodexDashboard'
  CODEX_ORCHESTRATION_RUNS_ROOT        = 'C:\Users\gregs\AppData\Local\CodexDashboard\orchestration-runs-validation'
}

Start-Process -FilePath (Join-Path $manualRoot 'controlplane-manual.exe') `
  -WorkingDirectory 'C:\Agent\CodexDashboard\backend\orchestration' `
  -RedirectStandardOutput (Join-Path $manualRoot 'manual.stdout.log') `
  -RedirectStandardError (Join-Path $manualRoot 'manual.stderr.log') `
  -Environment $envMap

$base = 'http://127.0.0.1:15318'
$dispatch = Invoke-RestMethod -Method Post -Uri "$base/api/v1/tasks/Task-0008/dispatch"

$run = $null
for ($i = 0; $i -lt 160; $i++) {
  Start-Sleep -Milliseconds 500
  $candidate = Invoke-RestMethod -Uri "$base/api/v1/task-runs/taskrun--Task-0008--active"
  $run = $candidate
  if ($candidate.state_envelope.reason_code -eq 'task_0008_workload_failure_attention_escalated') { break }
}

$oldOwnedRoot = $run.repo_lane.owned_repo_root
$seedBody = @{
  state               = 'blocked'
  reason_code         = 'workload_execution_failed'
  state_summary       = 'Run could not execute the prepared workload step inside the owned lane.'
  next_owner          = 'human_or_supervisor'
  next_expected_event = 'Retry the workload path with a fresh owned lane.'
} | ConvertTo-Json

$seeded = Invoke-RestMethod -Method Post -Uri "$base/api/v1/task-runs/taskrun--Task-0008--active/state" -ContentType 'application/json' -Body $seedBody
$retried = Invoke-RestMethod -Method Post -Uri "$base/api/v1/task-runs/taskrun--Task-0008--active/retry-workload"

$after = $null
for ($i = 0; $i -lt 160; $i++) {
  Start-Sleep -Milliseconds 500
  $candidate = Invoke-RestMethod -Uri "$base/api/v1/task-runs/taskrun--Task-0008--active"
  $after = $candidate
  if ($candidate.state_envelope.reason_code -eq 'task_0008_workload_failure_attention_escalated') { break }
}

$result = Get-Content -Raw $after.repo_lane.workload_result_path | ConvertFrom-Json
$taskAfter = Invoke-RestMethod -Uri "$base/api/v1/tasks/Task-0008"

Get-NetTCPConnection -LocalPort 15318 -State Listen | ForEach-Object { Stop-Process -Id $_.OwningProcess -Force }
docker compose --project-name codex-orchestration-validation -f .\dev\docker-compose.temporal-postgres.yml down -v
```

## Result

Pass.

`go test ./...` passed, including the new focused service and mux coverage for `retry-workload`.

The live manual backend proof on the clean `127.0.0.1:15318` listener then proved:

1. `POST /api/v1/tasks/Task-0008/dispatch` still returned:
   - `reason_code = owned_lane_bootstrapped`
2. The repaired live happy path still advanced automatically to:
   - `reason_code = task_0008_workload_failure_attention_escalated`
3. Seeding the active run through `POST /api/v1/task-runs/taskrun--Task-0008--active/state` produced:
   - `state = blocked`
   - `reason_code = workload_execution_failed`
4. `POST /api/v1/task-runs/taskrun--Task-0008--active/retry-workload` returned:
   - `reason_code = workload_retry_requested`
5. After the retry, the active run advanced back through the real owned-lane workload path to:
   - `reason_code = task_0008_workload_failure_attention_escalated`
6. The old failed owned lane was gone on disk:
   - `old_root_removed = true`
7. The retried run used a fresh owned lane:
   - `new_owned_root` differed from `old_owned_root`
8. The retried run exposed a new live workload result artifact at:
   - `repo_lane.workload_result_path`
9. The workload result artifact recorded:
   - `execution_summary = Executed Task-0008 backend validation and changed blocked-run recovery attention in an existing implementation file.`
10. The task readback after retry still showed:
    - `current_story.status = active_run`

## Why This Counts

This slice adds a real backend recovery action for actual `workload_execution_failed` runs instead of leaving those runs at truthful blocked readback only.

The backend now owns the recovery step:

- it rejects invalid retries
- it removes the failed owned lane
- it provisions and bootstraps a fresh owned lane
- it signals the active Temporal run to rerun the owned-lane workload path

That is a real backend recovery action, not just stronger blocked attention.

## Limitations

- The repaired normal live path no longer fails naturally, so this proof seeded `workload_execution_failed` through the backend `POST /state` path before exercising `retry-workload`.
- This slice proves backend-owned recovery for the active run, but it does not yet provide a naturally reproducible fault-injection hook for workload-execution failure.
- The default validation lane on `14318` was not used for final proof; final proof used a clean manual listener on `15318`.
- Direct validation-compose runs still require the explicit validation port overrides so Postgres does not collide with the service lane on `5432`.
- After resetting the disposable validation Temporal volume, the manual listener still needs a short warm-up delay before startup or Temporal client init can fail with `error reading server preface: EOF`.
- This is still backend proof, not dashboard regression proof.
