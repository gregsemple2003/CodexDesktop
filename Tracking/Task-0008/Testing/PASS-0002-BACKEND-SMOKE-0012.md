# PASS-0002 Backend Smoke 0012

## Type

Unit test / server-only smoke / live manual backend proof.

## Date

`2026-04-24`

## Scope

Proof for the first Task-0008-specific owned-lane code mutation slice:

- after dispatch, preflight, workload-step preparation, and Task-0008 validation, the workflow now writes a real Go source file inside the owned lane
- the code mutation is:
  - `backend/orchestration/internal/taskexec/task0008_owned_lane_generated.go`
- run readback exposes that durable path as:
  - `repo_lane.workload_code_path`
- the result artifact records a scoped post-execution git status over just the owned-lane outputs

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
  if ($candidate.state_envelope.reason_code -eq 'task_0008_owned_lane_code_written') { break }
}

$taskAfter = Invoke-RestMethod -Uri "$base/api/v1/tasks/Task-0008"
$artifact = Get-Content -Raw $run.repo_lane.workload_result_path | ConvertFrom-Json
$brief = Get-Content -Raw $run.repo_lane.workload_output_path
$code = Get-Content -Raw $run.repo_lane.workload_code_path
```

## Result

Pass.

`go test ./...` passed with the added coverage for:

- surfacing `repo_lane.workload_code_path`
- writing the owned-lane Go source scaffold
- validating that the mutated owned lane still passes:
  - `go test ./internal/taskexec ./internal/taskrun`

The live manual backend proof on the clean `127.0.0.1:15318` listener then proved:

1. `POST /api/v1/tasks/Task-0008/dispatch` still returned:
   - `reason_code = owned_lane_bootstrapped`
2. A follow-up `GET /api/v1/task-runs/taskrun--Task-0008--active` then advanced automatically to:
   - `state_envelope.state = running`
   - `reason_code = task_0008_owned_lane_code_written`
   - `state_summary = Run validated Task-0008 and wrote an owned-lane code scaffold.`
3. The run exposed:
   - `repo_lane.workload_output_path`
   - `repo_lane.workload_code_path`
4. The code mutation existed at:
   - `backend/orchestration/internal/taskexec/task0008_owned_lane_generated.go`
5. The generated code contained:
   - `Task0008OwnedLaneGeneratedSummary`
   - `Task0008OwnedLaneGeneratedTargets`
6. The workload result artifact recorded:
   - `execution_summary = Executed Task-0008 backend validation and wrote an owned-lane code scaffold.`
   - scoped `git_status_short_after` showing only:
     - `.codex-taskrun/`
     - `Tracking/Task-0008/OwnedLane/`
     - `backend/orchestration/internal/taskexec/task0008_owned_lane_generated.go`
7. The next task read returned:
   - `state_envelope.state = running`
   - `reason_code = task_0008_owned_lane_code_written`
   - `current_story.status = active_run`

## Why This Counts

This is the first Task-0008-specific owned-lane execution slice that mutates a code file in the backend code tree rather than only writing task artifacts.

The runtime now:

- validates Task-0008 backend packages inside the owned lane
- writes a bounded implementation brief for worker context
- writes a bounded Go source scaffold in the owned code tree
- validates the mutated owned lane again through the same focused package tests

without any manual `POST /state` mutation.

That is the smallest honest step from owned-lane artifact mutation into owned-lane code mutation.

## Limitations

- The code mutation is still a bounded scaffold file, not yet a worker-applied edit to an existing production implementation file.
- The result artifact now intentionally records scoped owned-lane git status instead of whole-worktree status, so it stays focused on the files this slice writes.
- After resetting the disposable validation Temporal volume, the manual listener still needs a short warm-up delay before startup or Temporal client init can fail with `error reading server preface: EOF`.
- This is still backend proof, not dashboard regression proof.
