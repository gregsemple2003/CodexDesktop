# PASS-0002 Backend Smoke 0014

## Type

Unit test / server-only smoke / live manual backend proof.

## Date

`2026-04-24`

## Scope

Proof for the first worker-applied implementation change in an existing Task-0008 file that alters runtime behavior inside the owned lane:

- after dispatch, preflight, workload-step preparation, Task-0008 validation, and owned-lane brief generation, the workflow now edits the existing implementation file:
  - `backend/orchestration/internal/taskexec/taskexec.go`
- that edit changes the bootstrapped-run suspiciousness window in `InitialView` from:
  - `15 minutes`
  to:
  - `5 minutes`
- after the edit, the workflow runs an owned-lane behavior probe and records its result in the workload artifact
- task readback advances automatically to:
  - `reason_code = task_0008_existing_file_behavior_changed`

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
  if ($candidate.state_envelope.reason_code -eq 'task_0008_existing_file_behavior_changed') { break }
}

$taskAfter = Invoke-RestMethod -Uri "$base/api/v1/tasks/Task-0008"
$artifact = Get-Content -Raw $run.repo_lane.workload_result_path | ConvertFrom-Json
$probe = Get-Content -Raw $artifact.behavior_probe_path | ConvertFrom-Json
$code = Get-Content -Raw $run.repo_lane.workload_code_path
```

## Result

Pass.

`go test ./...` passed with the updated coverage for:

- editing an existing owned-lane implementation file
- proving that edit changes runtime behavior rather than only leaving a comment marker
- recording a behavior-probe artifact after the owned-lane edit

The live manual backend proof on the clean `127.0.0.1:15318` listener then proved:

1. `POST /api/v1/tasks/Task-0008/dispatch` still returned:
   - `reason_code = owned_lane_bootstrapped`
2. A follow-up `GET /api/v1/task-runs/taskrun--Task-0008--active` then advanced automatically to:
   - `state_envelope.state = running`
   - `reason_code = task_0008_existing_file_behavior_changed`
   - `state_summary = Run validated Task-0008 and changed runtime behavior in an existing owned-lane implementation file.`
3. The run exposed:
   - `repo_lane.workload_code_path`
   - `repo_lane.workload_result_path`
4. The edited file was still:
   - `backend/orchestration/internal/taskexec/taskexec.go`
5. The edited file contained the real code change:
   - `suspiciousAfter = request.DispatchRequestedAt.Add(5 * time.Minute)`
6. The workload result artifact recorded:
   - `execution_summary = Executed Task-0008 backend validation and changed runtime behavior in an existing owned-lane implementation file.`
   - `behavior_probe_path`
7. The owned-lane behavior probe recorded:
   - `reason_code = owned_lane_bootstrapped`
   - `suspicious_window_minutes = 5`
8. The next task read returned:
   - `current_story.status = active_run`

## Why This Counts

This is the first Task-0008-specific owned-lane execution slice that edits an existing backend implementation file and proves a changed runtime behavior in that file.

The runtime now:

- validates Task-0008 backend packages inside the owned lane
- writes a bounded implementation brief for worker context
- edits the existing owned-lane `taskexec.go`
- proves the changed behavior with an owned-lane runtime probe

without any manual `POST /state` mutation.

That is the smallest honest step from a comment-level existing-file edit into a worker-applied implementation change that affects runtime behavior.

## Limitations

- The behavior change is still bounded to the bootstrapped-run suspiciousness window in `InitialView`; it does not yet alter a broader or later Task-0008 runtime path.
- The default validation lane on `14318` was not used for final proof; final proof used a clean manual listener on `15318`.
- When starting the disposable validation compose stack directly from `backend/orchestration`, explicit validation-lane port overrides are required or Postgres can collide with the service lane on `5432`.
- After resetting the disposable validation Temporal volume, the manual listener still needs a short warm-up delay before startup or Temporal client init can fail with `error reading server preface: EOF`.
- This is still backend proof, not dashboard regression proof.
