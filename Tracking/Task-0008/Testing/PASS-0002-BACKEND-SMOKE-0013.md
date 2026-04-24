# PASS-0002 Backend Smoke 0013

## Type

Unit test / server-only smoke / live manual backend proof.

## Date

`2026-04-24`

## Scope

Proof for the first worker-applied edit of an existing Task-0008 implementation file inside the owned lane:

- after dispatch, preflight, workload-step preparation, Task-0008 validation, and owned-lane brief generation, the workflow now edits an existing implementation file in the owned checkout
- the owned-lane code mutation is:
  - `backend/orchestration/internal/taskexec/taskexec.go`
- run readback exposes that durable path as:
  - `repo_lane.workload_code_path`
- task readback advances automatically to:
  - `reason_code = task_0008_existing_file_edited`
- the result artifact records scoped post-execution git status for the existing-file edit plus the owned-lane support outputs

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
for ($i = 0; $i -lt 80; $i++) {
  Start-Sleep -Milliseconds 500
  $candidate = Invoke-RestMethod -Uri "$base/api/v1/task-runs/taskrun--Task-0008--active"
  $run = $candidate
  if ($candidate.state_envelope.reason_code -eq 'task_0008_existing_file_edited') { break }
}

$taskAfter = Invoke-RestMethod -Uri "$base/api/v1/tasks/Task-0008"
$artifact = Get-Content -Raw $run.repo_lane.workload_result_path | ConvertFrom-Json
$brief = Get-Content -Raw $run.repo_lane.workload_output_path
$code = Get-Content -Raw $run.repo_lane.workload_code_path
```

## Result

Pass.

`go test ./...` passed with the updated owned-lane execution coverage for:

- surfacing `repo_lane.workload_code_path`
- editing an existing owned-lane implementation file instead of writing a generated side file
- validating that the edited owned lane still passes:
  - `go test ./internal/taskexec ./internal/taskrun`

The live manual backend proof on the clean `127.0.0.1:15318` listener then proved:

1. `POST /api/v1/tasks/Task-0008/dispatch` still returned:
   - `reason_code = owned_lane_bootstrapped`
2. A follow-up `GET /api/v1/task-runs/taskrun--Task-0008--active` then advanced automatically to:
   - `state_envelope.state = running`
   - `reason_code = task_0008_existing_file_edited`
   - `state_summary = Run validated Task-0008 and edited an existing owned-lane implementation file.`
3. The run exposed:
   - `repo_lane.workload_output_path`
   - `repo_lane.workload_code_path`
4. The code mutation existed at:
   - `backend/orchestration/internal/taskexec/taskexec.go`
5. The edited file contained:
   - `Task0008OwnedLaneEditNote:`
   - `Task0008OwnedLaneEditTargets:`
6. The workload result artifact recorded:
   - `execution_summary = Executed Task-0008 backend validation and edited an existing owned-lane implementation file.`
   - scoped `git_status_short_after` showing only:
     - `M backend/orchestration/internal/taskexec/taskexec.go`
     - `?? .codex-taskrun/`
     - `?? Tracking/Task-0008/OwnedLane/`
7. The next task read returned:
   - `state_envelope.state = running`
   - `reason_code = task_0008_existing_file_edited`
   - `current_story.status = active_run`

## Why This Counts

This is the first Task-0008-specific owned-lane execution slice that edits an existing backend implementation file rather than only writing generated side files or task artifacts.

The runtime now:

- validates Task-0008 backend packages inside the owned lane
- writes a bounded implementation brief for worker context
- applies a bounded edit to the existing owned-lane `taskexec.go`
- validates the edited owned lane through the same focused package tests

without any manual `POST /state` mutation.

That is the smallest honest step from owned-lane scaffold generation into a worker-applied edit of an existing implementation file.

## Limitations

- The existing-file edit is still a bounded comment-level mutation, not yet a worker-applied implementation change that alters Task-0008 runtime behavior.
- The result artifact intentionally records scoped owned-lane git status instead of whole-worktree status, so proof stays focused on the files this slice writes.
- The default validation lane on `14318` was not trusted for final proof because stale listeners and startup/log-lock failures had already been observed there; final proof used a clean manual listener on `15318`.
- After resetting the disposable validation Temporal volume, the manual listener still needs a short warm-up delay before startup or Temporal client init can fail with `error reading server preface: EOF`.
- This is still backend proof, not dashboard regression proof.
