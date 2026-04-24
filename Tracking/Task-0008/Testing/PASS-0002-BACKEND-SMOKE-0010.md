# PASS-0002 Backend Smoke 0010

## Type

Unit test / server-only smoke / live manual backend proof.

## Date

`2026-04-24`

## Scope

Proof for the first real task-specific backend execution slice in `PASS-0002`:

- after dispatch, preflight, workload-step preparation, and generic workload execution scaffolding, the workflow now runs a concrete Task-0008-specific backend command inside the owned lane
- the command is focused backend validation for Task-0008-owned packages:
  - `go test ./internal/taskexec ./internal/taskrun`
- the run advances automatically to the task-specific reason code:
  - `task_0008_backend_validation_complete`
- the result artifact records the execution kind, command, stdout path, and stderr path durably

## Commands

```powershell
cd C:\Agent\CodexDashboard\backend\orchestration
go test ./...

$env:CODEX_ORCH_POSTGRES_PORT='15432'
$env:CODEX_ORCH_TEMPORAL_PORT='17233'
$env:CODEX_ORCH_TEMPORAL_UI_PORT='18080'
docker compose --project-name codex-orchestration-validation -f .\dev\docker-compose.temporal-postgres.yml down -v
docker compose --project-name codex-orchestration-validation -f .\dev\docker-compose.temporal-postgres.yml up -d

Start-Sleep -Seconds 8

$manualRoot = Join-Path $env:LOCALAPPDATA 'CodexDashboard\orchestration-validation-manual-15318'
New-Item -ItemType Directory -Force -Path $manualRoot | Out-Null
go build -o (Join-Path $manualRoot 'controlplane-manual.exe') .\cmd\controlplane

$envMap = @{
  CODEX_ORCHESTRATION_BIND_ADDRESS    = '127.0.0.1:15318'
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
for ($i = 0; $i -lt 80; $i++) {
  Start-Sleep -Milliseconds 500
  $candidate = Invoke-RestMethod -Uri "$base/api/v1/task-runs/taskrun--Task-0008--active"
  $run = $candidate
  if ($candidate.state_envelope.reason_code -eq 'task_0008_backend_validation_complete') { break }
}

$taskAfter = Invoke-RestMethod -Uri "$base/api/v1/tasks/Task-0008"
$artifact = Get-Content -Raw $run.repo_lane.workload_result_path | ConvertFrom-Json
```

## Result

Pass.

`go test ./...` passed with the added coverage for:

- writing a Task-0008-specific execution plan into the owned-lane workload packet
- running that task-specific validation command inside a disposable owned lane
- persisting stdout and stderr paths in the workload result artifact

The live manual backend proof on the clean `127.0.0.1:15318` listener then proved:

1. `POST /api/v1/tasks/Task-0008/dispatch` still returned the initial bootstrap state:
   - `reason_code = owned_lane_bootstrapped`
2. A follow-up `GET /api/v1/task-runs/taskrun--Task-0008--active` advanced automatically to:
   - `state_envelope.state = running`
   - `reason_code = task_0008_backend_validation_complete`
   - `state_summary = Run completed Task-0008 backend validation inside the owned lane.`
3. The run recorded a real task-specific command:
   - `go test ./internal/taskexec ./internal/taskrun`
4. The run exposed:
   - `repo_lane.workload_result_path`
   - result-artifact `stdout_path`
   - result-artifact `stderr_path`
5. The workload result artifact under the owned lane recorded:
   - `execution_kind = task_0008_backend_validation`
   - the exact command
   - an execution summary
6. The next task read returned:
   - `state_envelope.state = running`
   - `reason_code = task_0008_backend_validation_complete`
   - `current_story.status = active_run`

## Why This Counts

This is the first slice where Task-0008's workflow executes a concrete task-specific backend command inside the owned lane instead of only recording generic workload packets and results.

The runtime now performs:

- owned-lane bootstrap
- execution preflight
- workload-step preparation
- workload-step execution
- Task-0008-specific backend validation

without any manual `POST /state` mutation.

That is the smallest honest step from generic runtime progression into real task-specific backend execution behavior.

## Limitations

- The task-specific command is still a bounded validation step, not full agent-driven implementation work inside the owned lane.
- After a fresh validation-volume reset, the manual listener may need a short Temporal warm-up delay before startup or it can fail with `error reading server preface: EOF`.
- This is still backend proof, not dashboard regression proof.
