# PASS-0002 Backend Smoke 0006

## Type

Unit test / server-only smoke / live manual backend proof.

## Date

`2026-04-24`

## Scope

Proof for the next bounded `PASS-0002` execution-bootstrap slice:

- dispatch performs real backend-owned progress over the owned lane without requiring `POST /state`
- dispatch captures and persists the owned lane `current_commit`
- dispatch writes a real bootstrap artifact for the run under the backend run-artifact root
- the created run starts in backend-produced `running` state rather than waiting for a manual state mutation
- task readback reflects the active run story immediately after dispatch

## Commands

```powershell
cd C:\Agent\CodexDashboard\backend\orchestration
go test ./...

$env:CODEX_ORCH_POSTGRES_PORT='15432'
$env:CODEX_ORCH_TEMPORAL_PORT='17233'
$env:CODEX_ORCH_TEMPORAL_UI_PORT='18080'
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
$before = Invoke-RestMethod -Uri "$base/api/v1/tasks/Task-0008"
$dispatch = Invoke-RestMethod -Method Post -Uri "$base/api/v1/tasks/Task-0008/dispatch"
$taskAfter = Invoke-RestMethod -Uri "$base/api/v1/tasks/Task-0008"
$bootstrap = Get-Content -Raw $dispatch.repo_lane.bootstrap_artifact_path | ConvertFrom-Json
git -C $dispatch.repo_lane.owned_repo_root rev-parse HEAD
```

## Result

Pass.

`go test ./...` passed with the added coverage for:

- dispatch bootstrap populating `current_commit` and bootstrap-artifact paths before the run starts
- initial Temporal task-run view starting in `running` state when owned-lane bootstrap evidence already exists
- dispatch service proof that the bootstrap artifact exists and records the owned lane and commit honestly

The live manual backend proof on the clean `127.0.0.1:15318` listener then proved:

1. `GET /api/v1/tasks/Task-0008` returned `dispatch_readiness.ready = true` before dispatch.
2. `POST /api/v1/tasks/Task-0008/dispatch` returned:
   - `state_envelope.state = running`
   - `reason_code = owned_lane_bootstrapped`
   - `next_owner = backend_worker`
   - `repo_lane.current_commit = repo_lane.baseline_commit`
   - `repo_lane.bootstrap_artifact_path` pointing at a real file under the backend run-artifact root
3. The bootstrap artifact existed on disk and recorded:
   - the same `current_commit`
   - the same `owned_repo_root`
   - the declared task snapshot identifiers used for dispatch
4. `git -C <owned_repo_root> rev-parse HEAD` matched the persisted `current_commit`.
5. The next task read returned:
   - `state_envelope.state = running`
   - `current_story.status = active_run`
   - `actions.interrupt.allowed = true`

## Why This Counts

This is the first slice where Task-0008 moves a run forward over the owned lane without any manual runtime state mutation.

Dispatch is no longer only:

- provision owned checkout
- persist durable run

It now also performs a real backend-owned bootstrap step that:

- reads the owned checkout
- captures the current commit
- writes a run artifact
- starts the run in a backend-produced active state

That is the smallest honest step from manual `/state` updates toward real execution behavior over the owned lane.

## Limitations

- This slice bootstraps the owned lane but does not yet run the actual task workload inside it.
- The default validation-lane runner on `127.0.0.1:14318` still should not be trusted for fresh-proof work when stale binaries or log-lock failures are suspected. This proof used the clean manual `127.0.0.1:15318` path again.
- This is still backend proof, not dashboard regression proof.
