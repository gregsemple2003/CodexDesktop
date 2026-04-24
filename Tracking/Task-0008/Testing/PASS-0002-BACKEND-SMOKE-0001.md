# PASS-0002 Backend Smoke 0001

## Type

Unit test / server-only smoke / live validation-lane proof.

## Date

`2026-04-24`

## Scope

Proof for the first `PASS-0002` supervision and intervention slice:

- read-through supervision that marks stale active runs as `sleeping_or_stalled`
- `POST /api/v1/task-runs/{run_id}/interrupt`
- owned-checkout restore-to-commit cleanup on interrupt
- task-level readback that releases story ownership after a terminal run
- redispatch after interrupt using the same active workflow id
- `POST /api/v1/task-runs/{run_id}/poke`

## Commands

```powershell
cd C:\Agent\CodexDashboard\backend\orchestration
go test ./...
powershell -ExecutionPolicy Bypass -File .\scripts\Start-ValidationLane.ps1
Invoke-RestMethod http://127.0.0.1:14318/api/v1/tasks/Task-0008 | ConvertTo-Json -Depth 10
Invoke-RestMethod -Method Post http://127.0.0.1:14318/api/v1/task-runs/taskrun--Task-0008--active/interrupt | ConvertTo-Json -Depth 10
Invoke-RestMethod http://127.0.0.1:14318/api/v1/tasks/Task-0008 | ConvertTo-Json -Depth 10
Invoke-RestMethod -Method Post http://127.0.0.1:14318/api/v1/tasks/Task-0008/dispatch | ConvertTo-Json -Depth 10

$body = @'
{
  "state": "running",
  "reason_code": "worker_started",
  "state_summary": "Run is actively executing in the owned checkout.",
  "next_owner": "backend",
  "next_expected_event": "Execution worker records the next progress checkpoint.",
  "suspicious_after": "2026-04-24T21:10:00Z",
  "last_progress_summary": "Run started backend execution."
}
'@

Invoke-RestMethod -Method Post `
  -Uri http://127.0.0.1:14318/api/v1/task-runs/taskrun--Task-0008--active/state `
  -ContentType 'application/json' `
  -Body $body | ConvertTo-Json -Depth 10

Invoke-RestMethod http://127.0.0.1:14318/api/v1/task-runs/taskrun--Task-0008--active | ConvertTo-Json -Depth 10
Invoke-RestMethod http://127.0.0.1:14318/api/v1/tasks/Task-0008 | ConvertTo-Json -Depth 10
Invoke-RestMethod -Method Post http://127.0.0.1:14318/api/v1/task-runs/taskrun--Task-0008--active/poke | ConvertTo-Json -Depth 10
```

## Result

Pass.

`go test ./...` passed with the new service and route coverage for:

- stale-run supervision into `sleeping_or_stalled`
- interrupt-driven owned-checkout restore
- task readback that stops treating terminal runs as the current live story
- `poke` and `interrupt` HTTP entrypoints

The live validation lane then proved the end-to-end backend behavior:

1. An existing active run in `waiting_for_human` could be interrupted.
2. The interrupt response showed:
   - `status = interrupted`
   - `reason_code = interrupt_requested`
   - `repo_lane.reset_status = restored`
3. The next `GET /api/v1/tasks/Task-0008` showed:
   - `current_story.status = no_active_run`
   - `dispatch_readiness.ready = true`
4. A fresh `POST /api/v1/tasks/Task-0008/dispatch` succeeded on the same workflow id and produced a new Temporal execution run id.
5. After setting the run to `running` with an already-expired `suspicious_after`, `GET /api/v1/task-runs/{run_id}` read through supervision and durably returned:
   - `state = sleeping_or_stalled`
   - `reason_code = progress_stale`
   - `actions.poke.allowed = true`
   - `actions.interrupt.allowed = true`
6. `POST /api/v1/task-runs/{run_id}/poke` then durably recorded:
   - `reason_code = poke_requested`
   - a refreshed `next_expected_event`
   - a new progress checkpoint summary

## Why This Counts

This is the first real proof that Task-0008 can supervise and intervene instead of only storing manual state updates.

The backend now owns:

- stale-progress detection
- honest interruption with owned-lane restore
- release of task story ownership after terminal runs
- redispatch after interruption
- durable poke semantics for stalled runs

That moves the runtime from passive persistence into the first recoverable supervision loop.

## Limitations

- The stale human-wait path is not yet escalated beyond urgent readback and still needs a fuller intervention story.
- `poke` records a backend intervention request, but there is not yet a worker-side response path that turns that intervention into fresh execution progress automatically.
- Cleanup failure surfacing still needs a stronger live proof for the blocked-reset path.
- This is not dashboard regression proof.
