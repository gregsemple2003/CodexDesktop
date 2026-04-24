# PASS-0000 Backend Smoke 0001

## Type

Unit test / server-only smoke.

## Date

`2026-04-24`

## Scope

Proof for the first `PASS-0000` backend contract slice:

- Task-0008 task-readback contract types exist in code.
- The backend can parse declared task docs from `Tracking/Task-<id>/`.
- `GET /api/v1/tasks`
- `GET /api/v1/tasks/{task_id}`

## Commands

```powershell
cd C:\Agent\CodexDashboard\backend\orchestration
go test ./...
```

## Result

Pass.

The suite succeeded after adding:

- task-run contract types under `internal/taskrun/`
- declared-doc snapshot and task readback parsing
- HTTP coverage for `GET /api/v1/tasks` and `GET /api/v1/tasks/{task_id}`

## Why This Counts

This is real evidence that the chosen Task-0008 contract is no longer only planning prose. The backend now has executable contract types and task-readback routes with unit coverage.

## Limitations

- This does not prove durable task-run dispatch or Temporal-backed run persistence yet.
- This does not prove owned-checkout allocation, restore-baseline capture, poke, or interrupt behavior.
- This is not task-level regression proof.
