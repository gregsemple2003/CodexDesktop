# PASS-0002 Backend Smoke 0020

## Type

Unit test / server-only smoke / live manual backend proof.

## Date

`2026-04-24`

## Scope

Proof for the bounded slice that removes the need to seed `workload_execution_failed` through `/state`:

- add `POST /api/v1/tasks/{task_id}/dispatch-workload-failure-exercise`
- for `Task-0008`, start a one-shot backend-owned workload-failure exercise through the normal workflow path
- let the workflow reach `workload_execution_failed` naturally
- verify the pending `workload_recovery` follow-up appears
- recover through `POST /api/v1/task-runs/{run_id}/retry-workload`

## Commands

```powershell
cd C:\Agent\CodexDashboard\backend\orchestration
gofmt -w .\internal\httpapi\mux.go .\internal\httpapi\mux_test.go .\internal\taskexec\taskexec.go .\internal\taskexec\taskexec_test.go .\internal\taskrun\service.go .\internal\taskrun\service_test.go .\internal\taskrun\types.go
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
$dispatch = Invoke-RestMethod -Method Post -Uri "$base/api/v1/tasks/Task-0008/dispatch-workload-failure-exercise"

$blocked = $null
for ($i = 0; $i -lt 200; $i++) {
  Start-Sleep -Milliseconds 500
  $candidate = Invoke-RestMethod -Uri "$base/api/v1/task-runs/taskrun--Task-0008--active"
  $blocked = $candidate
  if ($candidate.state_envelope.reason_code -eq 'workload_execution_failed') { break }
}

$step = Get-Content -Raw $blocked.repo_lane.workload_step_path | ConvertFrom-Json
$retried = Invoke-RestMethod -Method Post -Uri "$base/api/v1/task-runs/taskrun--Task-0008--active/retry-workload"

$after = $null
for ($i = 0; $i -lt 200; $i++) {
  Start-Sleep -Milliseconds 500
  $candidate = Invoke-RestMethod -Uri "$base/api/v1/task-runs/taskrun--Task-0008--active"
  $after = $candidate
  if ($candidate.state_envelope.reason_code -eq 'task_0008_workload_failure_attention_escalated') { break }
}

$result = Get-Content -Raw $after.repo_lane.workload_result_path | ConvertFrom-Json

Get-NetTCPConnection -LocalPort 15318 -State Listen | ForEach-Object { Stop-Process -Id $_.OwningProcess -Force }
docker compose --project-name codex-orchestration-validation -f .\dev\docker-compose.temporal-postgres.yml down -v
```

## Result

Pass.

`go test ./...` passed, including:

- taskrun coverage for the new failure-exercise dispatch directive
- mux coverage for the new task-level route
- taskexec coverage for the one-shot natural failure path

The live manual backend proof on the clean `127.0.0.1:15318` listener then proved:

1. `POST /api/v1/tasks/Task-0008/dispatch-workload-failure-exercise` still began with:
   - `reason_code = owned_lane_bootstrapped`
2. The run then reached:
   - `reason_code = workload_execution_failed`
   naturally through the workflow and activity path, without any `POST /state` mutation
3. The blocked run exposed:
   - `follow_up.kind = workload_recovery`
   - `follow_up.status = pending`
4. The workload packet recorded the bounded one-shot failure directive:
   - `failure_mode = task_0008_workload_execution_failure_once`
5. The blocked run's `failure_summary` pointed at the owned-lane failure-exercise test file path written by the activity before `go test` failed
6. `POST /api/v1/task-runs/taskrun--Task-0008--active/retry-workload` returned:
   - `reason_code = workload_retry_requested`
   - cleared `follow_up`
7. After retry, the run advanced back through the repaired happy path to:
   - `reason_code = task_0008_workload_failure_attention_escalated`
8. The retried run wrote a fresh workload result artifact again, and its `execution_summary` was:
   - `Executed Task-0008 backend validation and changed blocked-run recovery attention in an existing implementation file.`

## Why This Counts

This slice removes the remaining proof-honesty gap around workload-failure recovery.

The backend no longer needs a synthetic `POST /state` mutation to prove `workload_execution_failed` and `retry-workload` end to end. Instead:

- dispatch can start a bounded one-shot failure exercise
- the workflow reaches `workload_execution_failed` naturally
- readback exposes the expected `workload_recovery` follow-up
- retry runs the real backend recovery path back to success

That is a bounded backend-owned proof path that preserves the repaired happy path instead of regressing it.

## Limitations

- This failure-exercise dispatch is intentionally bounded to `Task-0008`; it is not a generalized fault-injection surface.
- The one-shot failure mode is a proof/debug hook, not a normal product path.
- The default validation lane on `14318` was not used for final proof; final proof used a clean manual listener on `15318`.
- Direct validation-compose runs still require the explicit validation port overrides so Postgres does not collide with the service lane on `5432`.
- After resetting the disposable validation Temporal volume, the manual listener still needs a short warm-up delay before startup or Temporal client init can fail with `error reading server preface: EOF`.
- This is still backend proof, not dashboard regression proof.
