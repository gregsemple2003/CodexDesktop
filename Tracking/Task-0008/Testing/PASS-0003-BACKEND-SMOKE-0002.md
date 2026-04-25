# PASS-0003 Backend Smoke 0002

## Type

Unit test / server-only smoke / live manual backend proof.

## Date

`2026-04-24`

## Scope

Proof for the bounded declared-doc drift slice:

- dispatch an active Task-0008 run
- mutate a declared task doc in the repo worktree
- trigger read-through reconcile through task readback
- confirm the active run later exposes:
  - `doc_runtime_divergence_status = reconciled`
  - an old-to-new declared revision summary
  - the updated captured task snapshot
- keep the active run's live story intact while reconcile happens

## Commands

```powershell
cd C:\Agent\CodexDashboard\backend\orchestration
gofmt -w .\internal\taskexec\taskexec.go .\internal\taskrun\service_test.go .\internal\temporalbackend\backend.go .\internal\temporalbackend\backend_test.go
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
$handoffPath = 'C:\Agent\CodexDashboard\Tracking\Task-0008\HANDOFF.md'
$original = Get-Content -Raw $handoffPath

try {
  $dispatch = Invoke-RestMethod -Method Post -Uri "$base/api/v1/tasks/Task-0008/dispatch"

  $runBefore = $null
  for ($i = 0; $i -lt 120; $i++) {
    Start-Sleep -Milliseconds 500
    $candidate = Invoke-RestMethod -Uri "$base/api/v1/task-runs/taskrun--Task-0008--active"
    $runBefore = $candidate
    if ($candidate.state_envelope.reason_code) { break }
  }

  $oldRevision = $runBefore.captured_task_snapshot.declared_task_revision
  $newBody = $original.TrimEnd("`r","`n") + "`r`n`r`n<!-- drift-proof $(Get-Date -Format o) -->`r`n"
  Set-Content -Path $handoffPath -Value $newBody -NoNewline

  $taskAfter = Invoke-RestMethod -Uri "$base/api/v1/tasks/Task-0008"

  $runAfter = $null
  for ($i = 0; $i -lt 120; $i++) {
    Start-Sleep -Milliseconds 500
    $candidate = Invoke-RestMethod -Uri "$base/api/v1/task-runs/taskrun--Task-0008--active"
    $runAfter = $candidate
    if ($candidate.doc_runtime_divergence_status -eq 'reconciled') { break }
  }
}
finally {
  Set-Content -Path $handoffPath -Value $original -NoNewline
}

Get-NetTCPConnection -LocalPort 15318 -State Listen | ForEach-Object { Stop-Process -Id $_.OwningProcess -Force }
docker compose --project-name codex-orchestration-validation -f .\dev\docker-compose.temporal-postgres.yml down -v
```

## Result

Pass.

`go test ./...` passed, including:

- taskrun coverage for task read reconcile after declared-doc drift
- taskexec coverage for the richer reconcile summary
- temporalbackend coverage for the bounded wait-until-reconciled helper used by signal-based read-through reconcile

The live manual backend proof on the clean `127.0.0.1:15318` listener then proved:

1. A normal `POST /api/v1/tasks/Task-0008/dispatch` started the active run.
2. The active run initially carried one declared task revision in its captured snapshot.
3. A temporary edit to `Tracking/Task-0008/HANDOFF.md` changed the declared task revision in the repo worktree.
4. `GET /api/v1/tasks/Task-0008` triggered the active-run read-through reconcile path.
5. The active run later exposed:
   - `doc_runtime_divergence_status = reconciled`
   - `doc_runtime_divergence_summary = Runtime reconciled the active run from declared task revision <old> to <new> during task readback.`
6. The run's `captured_task_snapshot.declared_task_revision` changed to the new declared task revision.
7. The run stayed the active story during reconcile:
   - `current_story.status = active_run`
8. The run's live execution story remained intact after reconcile:
   - `reason_code = task_0008_workload_failure_attention_escalated`
   - `last_progress_summary = Reconciled declared task docs into runtime state.`

## Why This Counts

This slice closes the remaining declared-doc drift proof gap in the approved backend contract.

The backend now proves, end to end, that:

- declared task docs can change while a run is active
- the backend detects that drift through task readback
- runtime state is not forgotten
- the run's captured snapshot is reconciled durably
- readback reports the divergence honestly instead of silently pretending nothing changed

## Limitations

- This proof exercised bounded declared-doc drift by temporarily editing `HANDOFF.md`; it did not perform a literal git commit-and-rewind sequence.
- On an active workflow, the reconcile signal may not show up in the immediate task-read response while the current activity is still running; the proof therefore polls the active run until the reconciled state is visible.
- The default validation lane on `14318` was not used for final proof; final proof used a clean manual listener on `15318`.
- Direct validation-compose runs still require the explicit validation port overrides so Postgres does not collide with the service lane on `5432`.
- After resetting the disposable validation Temporal volume, the manual listener still needs a short warm-up delay before startup or Temporal client init can fail with `error reading server preface: EOF`.
- This is still backend proof, not dashboard regression proof.
