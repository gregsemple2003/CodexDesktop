# PASS-0001 Backend Smoke 0001

## Type

Unit test / server-only smoke.

## Date

`2026-04-24`

## Scope

Proof for the first `PASS-0001` dispatch and persistence slice:

- `POST /api/v1/tasks/{task_id}/dispatch`
- `GET /api/v1/task-runs/{run_id}`
- Temporal-backed task-run workflow registration and query shape
- owned-checkout allocation for simple task dispatch
- baseline-commit capture and initial restore baseline capture

## Commands

```powershell
cd C:\Agent\CodexDashboard\backend\orchestration
go test ./...
```

## Result

Pass.

The passing suite now covers:

- task dispatch service behavior with owned worktree provisioning in a temporary git repo
- HTTP route coverage for task dispatch and task-run detail
- Temporal backend task-run start and query wiring at compile and unit-test level

## Why This Counts

This is real proof that Task-0008 now has a working backend dispatch slice instead of only passive task readback. Dispatch can allocate an isolated repo lane, capture a baseline commit, and create a durable task-run object shape that later passes can supervise.

## Limitations

- This does not yet prove live end-to-end dispatch against a running Temporal lane.
- This does not yet prove poke, interrupt, or cleanup reset behavior.
- This does not yet execute real task work inside the owned checkout.
- This is not task-level regression proof.
