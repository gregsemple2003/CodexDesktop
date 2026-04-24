# PASS-0002 Backend Smoke 0003

## Type

Unit test / server-only smoke / live validation-lane proof.

## Date

`2026-04-24`

## Scope

Proof for durable worker follow-up after `poke`:

- `poke` creates a pending backend-worker follow-up
- repeated `poke` is blocked while that follow-up is pending
- a later runtime update completes the follow-up durably

## Commands

```powershell
cd C:\Agent\CodexDashboard\backend\orchestration
go test ./...
powershell -ExecutionPolicy Bypass -File .\scripts\Start-ValidationLane.ps1

$body = @'
{
  "state": "sleeping_or_stalled",
  "reason_code": "progress_stale",
  "state_summary": "Run has gone quiet past its expected progress window.",
  "next_owner": "backend",
  "next_expected_event": "Poke or interrupt the run.",
  "follow_up": {},
  "last_progress_summary": "Prepared the run for worker follow-up proof."
}
'@

Invoke-RestMethod -Method Post `
  -Uri http://127.0.0.1:14318/api/v1/task-runs/taskrun--Task-0008--active/state `
  -ContentType 'application/json' `
  -Body $body | ConvertTo-Json -Depth 10

Invoke-RestMethod -Method Post http://127.0.0.1:14318/api/v1/task-runs/taskrun--Task-0008--active/poke | ConvertTo-Json -Depth 10

$resume = @'
{
  "state": "running",
  "reason_code": "worker_resumed",
  "state_summary": "Run resumed after worker follow-up.",
  "next_owner": "backend",
  "next_expected_event": "Execution worker records the next progress checkpoint.",
  "suspicious_after": "2026-04-24T22:00:00Z",
  "last_progress_summary": "Execution worker acknowledged the poke and resumed progress."
}
'@

Invoke-RestMethod -Method Post `
  -Uri http://127.0.0.1:14318/api/v1/task-runs/taskrun--Task-0008--active/state `
  -ContentType 'application/json' `
  -Body $resume | ConvertTo-Json -Depth 10

Invoke-RestMethod http://127.0.0.1:14318/api/v1/task-runs/taskrun--Task-0008--active | ConvertTo-Json -Depth 10
```

## Result

Pass.

`go test ./...` passed with the added coverage for:

- pending worker follow-up creation on `poke`
- automatic worker follow-up completion on a later runtime update
- overdue follow-up escalation
- explicit clearing of obsolete follow-up state

The live validation lane then proved:

1. A stalled run could be normalized into `sleeping_or_stalled` with obsolete follow-up cleared.
2. `POST /api/v1/task-runs/{run_id}/poke` returned:
   - `reason_code = poke_requested`
   - `follow_up.kind = poke_worker_check`
   - `follow_up.owner = backend_worker`
   - `follow_up.status = pending`
   - `actions.poke.block_reasons[0].code = follow_up_pending`
3. A later `POST /api/v1/task-runs/{run_id}/state` with fresh progress returned:
   - `state = running`
   - `reason_code = worker_resumed`
   - `follow_up.status = completed`
   - `follow_up.completed_at` populated
   - `actions.poke.block_reasons[0].code = run_not_suspicious_yet`

## Why This Counts

This closes the main worker-follow-up gap in PASS-0002.

`poke` is no longer just a recorded intervention. It now creates a durable worker obligation, blocks repeated `poke` while that obligation is pending, and completes that obligation when a fresh runtime update arrives.

That gives the backend a real follow-up loop instead of a write-only control surface.

## Limitations

- The worker follow-up is still completed by a later runtime update rather than a separate dedicated worker-ack endpoint.
- The validation-lane runner still showed intermittent stdout-log lock churn during restart, but the control plane served the proof requests correctly.
- This is not dashboard regression proof.
