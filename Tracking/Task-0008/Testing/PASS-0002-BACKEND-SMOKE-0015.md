# PASS-0002 Backend Smoke 0015

## Type

Unit test / server-only smoke / live manual backend proof.

## Date

`2026-04-24`

## Scope

Proof for the next worker-applied existing-file change on a later Task-0008 runtime path inside the owned lane:

- after dispatch, preflight, workload-step preparation, Task-0008 validation, and owned-lane brief generation, the workflow now edits the existing implementation file:
  - `backend/orchestration/internal/taskrun/service.go`
- that edit shortens the interrupt-review follow-up window in the owned lane from:
  - `24 hours`
  to:
  - `2 hours`
- after the edit, the workflow runs an owned-lane behavior probe that exercises `InterruptRun` against the edited owned lane
- task readback advances automatically to:
  - `reason_code = task_0008_interrupt_review_window_changed`

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
for ($i = 0; $i -lt 100; $i++) {
  Start-Sleep -Milliseconds 500
  $candidate = Invoke-RestMethod -Uri "$base/api/v1/task-runs/taskrun--Task-0008--active"
  $run = $candidate
  if ($candidate.state_envelope.reason_code -eq 'task_0008_interrupt_review_window_changed') { break }
}

$taskAfter = Invoke-RestMethod -Uri "$base/api/v1/tasks/Task-0008"
$artifact = Get-Content -Raw $run.repo_lane.workload_result_path | ConvertFrom-Json
$probe = Get-Content -Raw $artifact.behavior_probe_path | ConvertFrom-Json
$code = Get-Content -Raw $run.repo_lane.workload_code_path
```

## Result

Pass.

`go test ./...` passed with the updated coverage for:

- editing an existing owned-lane implementation file on the later interrupt path
- proving that edit changes interrupt-review follow-up timing rather than only a bootstrap-time field
- recording a behavior-probe artifact after the owned-lane edit

The live manual backend proof on the clean `127.0.0.1:15318` listener then proved:

1. `POST /api/v1/tasks/Task-0008/dispatch` still returned:
   - `reason_code = owned_lane_bootstrapped`
2. A follow-up `GET /api/v1/task-runs/taskrun--Task-0008--active` then advanced automatically to:
   - `state_envelope.state = running`
   - `reason_code = task_0008_interrupt_review_window_changed`
   - `state_summary = Run validated Task-0008 and changed a later interrupt-review behavior in an existing owned-lane implementation file.`
3. The run exposed:
   - `repo_lane.workload_code_path`
   - `repo_lane.workload_result_path`
4. The edited file was:
   - `backend/orchestration/internal/taskrun/service.go`
5. The edited file contained the real code change:
   - `DueAt:       now.Add(2 * time.Hour),`
6. The workload result artifact recorded:
   - `execution_summary = Executed Task-0008 backend validation and changed a later interrupt-review behavior in an existing owned-lane implementation file.`
   - `behavior_probe_path`
7. The owned-lane behavior probe recorded:
   - `reason_code = interrupt_requested`
   - `follow_up_kind = interrupt_review`
   - `due_window_hours = 2`
   - `reset_status = restored`
8. The next task read returned:
   - `current_story.status = active_run`

## Why This Counts

This is the next worker-applied existing-file change on a later Task-0008 runtime path inside the owned lane.

The runtime now:

- validates Task-0008 backend packages inside the owned lane
- writes a bounded implementation brief for worker context
- edits the existing owned-lane `service.go` on the later interrupt path
- proves the changed interrupt-review follow-up timing with an owned-lane runtime probe

without any manual `POST /state` mutation.

That is the smallest honest step from the earlier bootstrap-window change into a later Task-0008 runtime path in an existing implementation file.

## Limitations

- The behavior change is still bounded to the interrupt-review follow-up window; it does not yet alter a broader recovery or redispatch policy beyond that timing.
- The default validation lane on `14318` was not used for final proof; final proof used a clean manual listener on `15318`.
- Direct validation-compose runs still require the explicit validation port overrides so Postgres does not collide with the service lane on `5432`.
- After resetting the disposable validation Temporal volume, the manual listener still needs a short warm-up delay before startup or Temporal client init can fail with `error reading server preface: EOF`.
- This is still backend proof, not dashboard regression proof.
