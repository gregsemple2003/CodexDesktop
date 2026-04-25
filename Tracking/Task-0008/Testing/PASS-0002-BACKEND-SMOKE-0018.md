# PASS-0002 Backend Smoke 0018

## Type

Unit test / server-only smoke / live manual backend proof.

## Date

`2026-04-24`

## Scope

Proof for the bounded repair slice on the current live Task-0008 workload-execution path:

- repair the stale Task-0008 owned-lane mutation recipe so the workload step no longer stalls at `workload_execution_failed`
- retarget the owned-lane existing-file edit to the current `service.go` baseline
- prove the repaired path advances to a new Task-0008-specific reason code and writes a workload result artifact again

## Commands

```powershell
cd C:\Agent\CodexDashboard\backend\orchestration
gofmt -w .\internal\taskexec\taskexec.go .\internal\taskexec\taskexec_test.go
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
for ($i = 0; $i -lt 120; $i++) {
  Start-Sleep -Milliseconds 500
  $candidate = Invoke-RestMethod -Uri "$base/api/v1/task-runs/taskrun--Task-0008--active"
  $run = $candidate
  if ($candidate.state_envelope.reason_code -eq 'task_0008_workload_failure_attention_escalated') { break }
  if ($candidate.state_envelope.reason_code -eq 'workload_execution_failed') { break }
}

$result = Get-Content -Raw $run.repo_lane.workload_result_path | ConvertFrom-Json
$probe = Get-Content -Raw $result.behavior_probe_path | ConvertFrom-Json
$code = Get-Content -Raw $run.repo_lane.workload_code_path
$taskAfter = Invoke-RestMethod -Uri "$base/api/v1/tasks/Task-0008"
```

## Result

Pass.

`go test ./...` passed, including the repaired Task-0008 owned-lane execution unit coverage.

The live manual backend proof on the clean `127.0.0.1:15318` listener then proved:

1. `POST /api/v1/tasks/Task-0008/dispatch` still returned:
   - `reason_code = owned_lane_bootstrapped`
2. The active run no longer stalled at `workload_execution_failed`.
3. Instead, the run advanced automatically to:
   - `reason_code = task_0008_workload_failure_attention_escalated`
   - `state_summary = Run validated Task-0008 and changed blocked-run recovery attention in an existing implementation file.`
4. The run exposed:
   - `repo_lane.workload_result_path`
   - `repo_lane.workload_code_path`
5. The workload result artifact recorded:
   - `execution_summary = Executed Task-0008 backend validation and changed blocked-run recovery attention in an existing implementation file.`
   - `behavior_probe_path`
6. The behavior-probe artifact recorded:
   - `blocked_attention_level = urgent`
   - `proof_test_path`
   - `go_test_passed = true`
7. The owned-lane edited file contained the new blocked-run recovery attention change:
   - `AttentionUrgent`
   - sort key `18-blocked_recovery`
8. The next task read still showed:
   - `current_story.status = active_run`

## Why This Counts

This slice repaired the real live Task-0008 workload-execution path.

The owned-lane mutation recipe had drifted behind the current repo baseline and was causing a legitimate compile failure during the owned-lane `go test` step. After this repair, the workload step completes again, writes its result artifact again, and advances to a Task-0008-specific runtime reason code instead of stalling in `workload_execution_failed`.

That is a real recovery of the backend-driven execution path, not another placeholder marker.

## Limitations

- This slice repaired the owned-lane execution recipe and advanced it to a new bounded recovery-policy change, but it did not yet add a new real backend action for recovering an actual `workload_execution_failed` run.
- The owned-lane proof for this slice still uses a generated package-local test file to validate the changed behavior inside the owned checkout.
- The default validation lane on `14318` was not used for final proof; final proof used a clean manual listener on `15318`.
- Direct validation-compose runs still require the explicit validation port overrides so Postgres does not collide with the service lane on `5432`.
- After resetting the disposable validation Temporal volume, the manual listener still needs a short warm-up delay before startup or Temporal client init can fail with `error reading server preface: EOF`.
- This is still backend proof, not dashboard regression proof.
