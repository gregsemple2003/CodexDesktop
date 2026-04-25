# PASS-0002 Backend Smoke 0017

## Type

Unit test / server-only smoke / live manual backend proof.

## Date

`2026-04-24`

## Scope

Proof for the next broader recovery-policy slice over the owned lane:

- resolving `interrupt_review` now immediately releases the resolved owned lane instead of deferring cleanup until a later redispatch
- the resolved run now persists released-lane metadata:
  - `repo_lane.reset_status = released`
  - `repo_lane.last_reset_target_commit`
  - cleared live `repo_lane.owned_repo_root`
- Windows owned-lane cleanup now removes git worktrees with `core.longpaths=true` so review-resolution cleanup does not fail on long artifact paths

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
  $run = Invoke-RestMethod -Uri "$base/api/v1/task-runs/taskrun--Task-0008--active"
  if ($run.state_envelope.reason_code -ne 'owned_lane_bootstrapped') { break }
}

$interrupted = Invoke-RestMethod -Method Post -Uri "$base/api/v1/task-runs/taskrun--Task-0008--active/interrupt"
$resolveBody = @{ decision = 'redispatch_ready'; summary = 'Human review approved another dispatch attempt.'; resolved_by = 'human' } | ConvertTo-Json
$resolved = Invoke-RestMethod -Method Post -Uri "$base/api/v1/task-runs/taskrun--Task-0008--active/resolve-interrupt-review" -ContentType 'application/json' -Body $resolveBody
$taskAfter = Invoke-RestMethod -Uri "$base/api/v1/tasks/Task-0008"
```

## Result

Pass.

`go test ./...` passed, including the new taskrun coverage for:

- immediate owned-lane release on `redispatch_ready`
- immediate owned-lane release on `keep_closed`
- redispatch no longer being responsible for removing a previously resolved lane

The live manual backend proof on the clean `127.0.0.1:15318` listener then proved:

1. `POST /api/v1/tasks/Task-0008/dispatch` still returned:
   - `reason_code = owned_lane_bootstrapped`
2. On this proof lane, the active run later settled at:
   - `reason_code = workload_execution_failed`
3. `POST /api/v1/task-runs/taskrun--Task-0008--active/interrupt` still returned:
   - `reason_code = interrupt_requested`
4. `POST /api/v1/task-runs/taskrun--Task-0008--active/resolve-interrupt-review` now returned:
   - `reason_code = interrupt_review_resolved_redispatch_ready`
   - `state_summary = Interrupt review approved the run for redispatch and backend released the prior owned lane.`
   - `repo_lane.reset_status = released`
   - an empty live `repo_lane.owned_repo_root`
   - `repo_lane.last_reset_target_commit`
5. The previously interrupted owned lane was already gone on disk immediately after resolution:
   - `prior_owned_root_removed = true`
6. The next task read returned:
   - `dispatch_readiness.ready = true`
   - `actions.dispatch.allowed = true`
   - `current_story.status = no_active_run`

## Why This Counts

This slice materially advances recovery policy over the owned lane.

The backend no longer waits for a later redispatch to clean up a resolved interrupted lane. Once interrupt review is resolved, the owned lane is released immediately and readback reflects that release durably.

That is a real backend recovery-path improvement, not another synthetic marker-only transition.

## Limitations

- This proof recovered from a real run that had already reached `workload_execution_failed`; it did not depend on the older synthetic task-specific running marker.
- The current live Task-0008 workload-execution path still needs separate follow-up if the goal is to keep proving later owned-lane task-specific progress before interruption.
- The default validation lane on `14318` was not used for final proof; final proof used a clean manual listener on `15318`.
- Direct validation-compose runs still require the explicit validation port overrides so Postgres does not collide with the service lane on `5432`.
- After resetting the disposable validation Temporal volume, the manual listener still needs a short warm-up delay before startup or Temporal client init can fail with `error reading server preface: EOF`.
- This is still backend proof, not dashboard regression proof.
