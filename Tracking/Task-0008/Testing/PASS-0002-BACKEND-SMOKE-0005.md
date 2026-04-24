# PASS-0002 Backend Smoke 0005

## Type

Unit test / server-only smoke / live manual backend proof.

## Date

`2026-04-24`

## Scope

Proof for the bounded `PASS-0002` interrupt-review slice:

- `POST /api/v1/task-runs/{run_id}/resolve-interrupt-review`
- interrupted runs with pending `interrupt_review` block redispatch at the task level
- interrupt review resolution records a durable `resolution`
- resolved interrupt review releases the task back to `dispatch_readiness.ready = true`
- the Temporal-backed workflow must stay alive long enough to accept the resolution signal and return the resolved run view

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
  "& '" + (Join-Path $manualRoot 'controlplane-manual.exe') + "'"
) -WorkingDirectory 'C:\Agent\CodexDashboard\backend\orchestration'

$base = 'http://127.0.0.1:15318'
Invoke-RestMethod -Uri "$base/api/v1/tasks/Task-0008" | ConvertTo-Json -Depth 10
Invoke-RestMethod -Method Post -Uri "$base/api/v1/tasks/Task-0008/dispatch" | ConvertTo-Json -Depth 10
Invoke-RestMethod -Method Post -Uri "$base/api/v1/task-runs/taskrun--Task-0008--active/interrupt" | ConvertTo-Json -Depth 10
Invoke-RestMethod -Uri "$base/api/v1/tasks/Task-0008" | ConvertTo-Json -Depth 10

$resolveBody = @'
{
  "decision": "redispatch_ready",
  "summary": "Human review approved another dispatch attempt.",
  "resolved_by": "human"
}
'@

Invoke-RestMethod -Method Post `
  -Uri "$base/api/v1/task-runs/taskrun--Task-0008--active/resolve-interrupt-review" `
  -ContentType 'application/json' `
  -Body $resolveBody | ConvertTo-Json -Depth 10

Invoke-RestMethod -Uri "$base/api/v1/tasks/Task-0008" | ConvertTo-Json -Depth 10
Invoke-RestMethod -Uri "$base/api/v1/task-runs/taskrun--Task-0008--active" | ConvertTo-Json -Depth 10
```

## Result

Pass.

`go test ./...` passed with the added coverage for:

- keeping an interrupted workflow alive while `interrupt_review` is still pending
- falling back to the closed Temporal workflow result when a successful update closes the run before the immediate query returns
- task-level blocking while interrupt review is pending
- task-level redispatch readiness after interrupt review resolves
- HTTP route coverage for `POST /api/v1/task-runs/{run_id}/resolve-interrupt-review`

The live manual backend proof on the clean `127.0.0.1:15318` listener then proved:

1. `POST /api/v1/tasks/Task-0008/dispatch` returned a run in `dispatching`.
2. `POST /api/v1/task-runs/taskrun--Task-0008--active/interrupt` returned:
   - `state_envelope.state = interrupted`
   - `follow_up.kind = interrupt_review`
   - `follow_up.status = pending`
3. `GET /api/v1/tasks/Task-0008` then returned:
   - `state_envelope.state = waiting_for_human`
   - `reason_code = interrupt_review_pending`
   - `dispatch_readiness.ready = false`
   - `actions.dispatch.allowed = false`
   - `actions.dispatch.block_reasons[0].code = interrupt_review_pending`
4. `POST /api/v1/task-runs/taskrun--Task-0008--active/resolve-interrupt-review` returned:
   - `reason_code = interrupt_review_resolved_redispatch_ready`
   - `follow_up.status = completed`
   - `resolution.kind = interrupt_review`
   - `resolution.decision = redispatch_ready`
5. The next task read returned:
   - `dispatch_readiness.ready = true`
   - `state_envelope.state = ready`
   - `current_story.status = no_active_run`
6. `GET /api/v1/task-runs/taskrun--Task-0008--active` still returned the resolved interrupted run with:
   - `state_envelope.reason_code = interrupt_review_resolved_redispatch_ready`
   - `follow_up.status = completed`
   - `resolution.decision = redispatch_ready`

## Why This Counts

This closes the interrupt-review decision gap inside `PASS-0002`.

Interrupt no longer stops at a durable pending follow-up with no backend resolution path. The backend can now:

- block redispatch while review is unresolved
- record the human or supervisor decision explicitly
- preserve that decision in the run record even when the Temporal workflow closes immediately after resolution
- return the task to a dispatch-ready state without inventing client heuristics

That keeps the humane task story aligned with the runtime contract instead of leaving review resolution implicit or manual.

## Limitations

- The default validation-lane runner on `127.0.0.1:14318` was not trustworthy for this proof. Its logs showed `go build failed; reusing existing binary`, repeated stdout-log lock failures, and a stale listener serving older behavior. This slice was proved instead against a clean manual current binary on `127.0.0.1:15318`.
- The validation compose file defaults to service-lane ports when the validation port overrides are not set. For this proof it had to be started explicitly with `15432`, `17233`, and `18080`.
- This is still backend proof, not dashboard regression proof.
