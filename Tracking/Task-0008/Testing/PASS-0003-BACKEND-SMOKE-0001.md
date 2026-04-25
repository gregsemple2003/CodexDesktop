# PASS-0003 Backend Smoke 0001

## Type

Unit test / server-only smoke / live manual backend proof.

## Date

`2026-04-24`

## Scope

Proof for the first `PASS-0003` context-readback slice:

- task readback exposes `deep_context` before any active run exists
- run readback exposes `deep_context` launch targets after dispatch
- task readback switches to the active run's `deep_context` once a run owns the live story

## Commands

```powershell
cd C:\Agent\CodexDashboard\backend\orchestration
gofmt -w .\internal\taskrun\types.go .\internal\taskrun\service.go .\internal\taskrun\service_test.go .\internal\taskexec\taskexec.go .\internal\httpapi\mux_test.go
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
$taskBefore = Invoke-RestMethod -Uri "$base/api/v1/tasks/Task-0008"
$dispatch = Invoke-RestMethod -Method Post -Uri "$base/api/v1/tasks/Task-0008/dispatch"

$run = $null
for ($i = 0; $i -lt 120; $i++) {
  Start-Sleep -Milliseconds 500
  $candidate = Invoke-RestMethod -Uri "$base/api/v1/task-runs/taskrun--Task-0008--active"
  $run = $candidate
  if ($candidate.deep_context -and $candidate.deep_context.launch_targets.Count -gt 0) { break }
}

$taskAfter = Invoke-RestMethod -Uri "$base/api/v1/tasks/Task-0008"

Get-NetTCPConnection -LocalPort 15318 -State Listen | ForEach-Object { Stop-Process -Id $_.OwningProcess -Force }
docker compose --project-name codex-orchestration-validation -f .\dev\docker-compose.temporal-postgres.yml down -v
```

## Result

Pass.

`go test ./...` passed, including:

- taskrun coverage for task-level and run-level deep-context population
- mux coverage that the read APIs expose the new shape
- taskexec coverage that run views carry the deep-context contract from workflow start

The live manual backend proof on the clean `127.0.0.1:15318` listener then proved:

1. `GET /api/v1/tasks/Task-0008` before dispatch exposed `deep_context` with:
   - preferred launch target kind `task_artifact`
   - launch targets:
     - `Task folder`
     - `Task handoff`
     - `Task plan`
2. `POST /api/v1/tasks/Task-0008/dispatch` still began the normal backend run path.
3. `GET /api/v1/task-runs/taskrun--Task-0008--active` then exposed `deep_context` with launch targets:
   - `Task folder`
   - `Task handoff`
   - `Owned checkout`
   - `Run artifacts`
   - `Active run API resource`
4. The run readback still showed a real active runtime state:
   - `reason_code = workload_step_prepared`
5. `GET /api/v1/tasks/Task-0008` after dispatch switched the task-level `deep_context` to the active run's launch-target set instead of leaving the operator at markdown-only context.

## Why This Counts

This slice materially advances the approved operator contract.

The backend now tells the caller what to open next instead of forcing manual reconstruction from:

- task markdown
- owned-lane paths
- run-artifact paths
- raw run ids

That is a real backend context-readback capability, not just another state label.

## Limitations

- This proof exercised launch-target readback, not a richer transcript or session deep-link.
- If the dispatching process does not have a visible session id or transcript path in its environment, `deep_context` still falls back to launch targets rather than inventing transcript provenance.
- The task still needs explicit proof that git rollback or task-doc drift is reported honestly through divergence fields while runtime truth stays preserved.
- The default validation lane on `14318` was not used for final proof; final proof used a clean manual listener on `15318`.
- Direct validation-compose runs still require the explicit validation port overrides so Postgres does not collide with the service lane on `5432`.
- After resetting the disposable validation Temporal volume, the manual listener still needs a short warm-up delay before startup or Temporal client init can fail with `error reading server preface: EOF`.
- This is still backend proof, not dashboard regression proof.
