# PASS-0001 Backend Smoke 0002

## Type

Unit test / server-only smoke / live validation-lane proof.

## Date

`2026-04-24`

## Scope

Proof for the richer `PASS-0001` runtime-state slice:

- `POST /api/v1/task-runs/{run_id}/state`
- Temporal-backed task-run state mutation after dispatch
- task read-through of active-run state into `GET /api/v1/tasks/{task_id}`
- Windows-safe owned-checkout path shortening for live dispatch lanes

## Commands

```powershell
cd C:\Agent\CodexDashboard\backend\orchestration
go test ./...
powershell -ExecutionPolicy Bypass -File .\scripts\Start-ValidationLane.ps1
Invoke-RestMethod http://127.0.0.1:14318/healthz | ConvertTo-Json -Depth 10
Invoke-RestMethod http://127.0.0.1:14318/api/v1/tasks/Task-0008 | ConvertTo-Json -Depth 10

$body = @'
{
  "state": "waiting_for_human",
  "reason_code": "approval_required",
  "state_summary": "Run is waiting for approval of the PASS-0001 runtime slice.",
  "next_owner": "human",
  "next_expected_event": "Approve or redirect the next backend runtime step.",
  "suspicious_after": "2026-04-24T21:40:00Z",
  "last_progress_summary": "Run recorded a fresh live validation checkpoint.",
  "wait_contract": {
    "waiting_on": "human_approval",
    "why_blocked": "The next backend runtime step should not proceed without explicit approval.",
    "resume_when": "The human approves or redirects the next runtime step.",
    "human_action_required": true,
    "human_action_target": {
      "kind": "approval_action",
      "label": "Approve PASS-0001 runtime slice",
      "uri": "approval://taskrun/Task-0008/pass-0001"
    },
    "stale_after": "2026-04-24T21:50:00Z"
  }
}
'@

Invoke-RestMethod -Method Post `
  -Uri http://127.0.0.1:14318/api/v1/task-runs/taskrun--Task-0008--active/state `
  -ContentType 'application/json' `
  -Body $body | ConvertTo-Json -Depth 10

Invoke-RestMethod http://127.0.0.1:14318/api/v1/task-runs/taskrun--Task-0008--active | ConvertTo-Json -Depth 10
Invoke-RestMethod http://127.0.0.1:14318/api/v1/tasks/Task-0008 | ConvertTo-Json -Depth 10
```

## Result

Pass.

`go test ./...` passed after adding the task-run state-update API, workflow signal handling, and richer run-state mutation tests.

The live validation lane was healthy on `http://127.0.0.1:14318`, accepted a fresh `POST /api/v1/task-runs/taskrun--Task-0008--active/state`, and then returned the updated `waiting_for_human` state through both:

- `GET /api/v1/task-runs/taskrun--Task-0008--active`
- `GET /api/v1/tasks/Task-0008`

The readback showed:

- `reason_code = approval_required`
- plain-language `state_summary`
- explicit `wait_contract`
- `human_action_target.kind = approval_action`
- `attention_level = needs_attention`
- dispatch still blocked while the active run owns the live story

The owned checkout for the live run also remained on the shortened Windows-safe path under `%LOCALAPPDATA%\\Temp\\cdxow\\...`, which avoided the earlier path-length failure during worktree creation.

## Why This Counts

This proves `PASS-0001` now has more than initial dispatch persistence. A live Temporal-backed task run can move into a richer human-facing state contract after dispatch, and the task API surfaces that same live story without inventing client-side heuristics.

This is the first live proof that Task-0008 can carry:

- state-specific summaries
- explicit wait contracts
- explicit human action targets
- attention priority
- post-dispatch progress checkpoints

over a real backend lane instead of only unit-test fakes.

## Limitations

- This proof updated an already-active validation-lane run rather than dispatching a brand-new run in the same command sequence.
- Poke, interrupt, cleanup reset, and real task execution inside the owned checkout are still unfinished.
- This is not dashboard regression proof.
