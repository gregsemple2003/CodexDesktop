# PASS-0002 Backend Smoke 0011

## Type

Unit test / server-only smoke / live manual backend proof.

## Date

`2026-04-24`

## Scope

Proof for the next bounded Task-0008-specific owned-lane slice:

- after dispatch, preflight, workload-step preparation, and Task-0008 validation, the workflow now makes a real repo-state change inside the owned lane
- the task-specific execution path writes:
  - `Tracking/Task-0008/OwnedLane/IMPLEMENTATION-BRIEF.md`
- run readback exposes that durable path as:
  - `repo_lane.workload_output_path`
- the result artifact records the post-execution owned-lane git status so the repo-state change is explicit

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
  if ($candidate.state_envelope.reason_code -eq 'task_0008_owned_lane_brief_written') { break }
}

$taskAfter = Invoke-RestMethod -Uri "$base/api/v1/tasks/Task-0008"
$artifact = Get-Content -Raw $run.repo_lane.workload_result_path | ConvertFrom-Json
$brief = Get-Content -Raw $run.repo_lane.workload_output_path
git -C $run.repo_lane.owned_repo_root status --short
```

## Result

Pass.

`go test ./...` passed with the added coverage for:

- surfacing `repo_lane.workload_output_path`
- writing an owned-lane Task-0008 implementation brief
- recording post-execution git status after the repo-state change

The live manual backend proof on the clean `127.0.0.1:15318` listener then proved:

1. `POST /api/v1/tasks/Task-0008/dispatch` still returned:
   - `reason_code = owned_lane_bootstrapped`
2. A follow-up `GET /api/v1/task-runs/taskrun--Task-0008--active` then advanced automatically to:
   - `state_envelope.state = running`
   - `reason_code = task_0008_owned_lane_brief_written`
   - `state_summary = Run completed Task-0008 validation and wrote an owned-lane implementation brief.`
3. The run exposed:
   - `repo_lane.workload_result_path`
   - `repo_lane.workload_output_path`
4. The owned-lane output file existed at:
   - `Tracking/Task-0008/OwnedLane/IMPLEMENTATION-BRIEF.md`
5. The output file contained a concrete Task-0008-specific brief with:
   - task summary
   - handoff-derived next step
   - candidate backend file targets
6. The workload result artifact recorded:
   - `execution_kind = task_0008_backend_validation`
   - `execution_summary = Executed Task-0008 backend validation and wrote an owned-lane implementation brief.`
   - `git_status_short_after = ?? .codex-taskrun/` and `?? Tracking/Task-0008/OwnedLane/`
7. The next task read returned:
   - `state_envelope.state = running`
   - `reason_code = task_0008_owned_lane_brief_written`
   - `current_story.status = active_run`

## Why This Counts

This is the first Task-0008-specific execution slice that leaves behind a real repo-state mutation in the owned lane instead of only validation output or generic runtime markers.

The runtime now:

- validates Task-0008 backend packages inside the owned lane
- writes a task-specific implementation brief under the owned task directory
- records the resulting owned-lane git dirtiness durably

without any manual `POST /state` mutation.

That is the smallest honest step from task-specific validation into task-specific owned-lane execution that changes repo state.

## Limitations

- The repo-state change is still a bounded owned-lane task artifact, not yet a source-code mutation or agent-authored code patch.
- After resetting the disposable validation Temporal volume, the manual listener still needs a short warm-up delay before startup or Temporal client init can fail with `error reading server preface: EOF`.
- This is still backend proof, not dashboard regression proof.
