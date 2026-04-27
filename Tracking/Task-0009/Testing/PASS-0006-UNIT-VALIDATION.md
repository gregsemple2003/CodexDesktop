# PASS-0006 Unit Validation

Status: passing

Date: 2026-04-27

## Scope

Validation for the REG-004 remediation pass before rerunning the actual
dashboard-surface regression.

## Commands

```powershell
python -m unittest tests.test_tasks_backend -v
```

Result: passing, 12 tests.

```powershell
go test ./...
```

Working directory: `backend/orchestration`

Result: passing.

```powershell
python -m unittest discover -s tests -p "test_*.py" -v
```

Result: passing, 101 tests.

## Coverage Notes

- Backend task-state mapping treats durable `status: complete` as terminal.
- Backend task-state mapping treats durable `status: cancelled` as terminal.
- Terminal durable task state suppresses stale active runtime state.
- Frontend task mapping filters terminal tasks from the active Tasks surface.
- Frontend task mapping does not turn generic `needs_attention` into
  `Waiting on you`.
- Frontend task mapping hides run controls when no active run id exists.

REG-004 app-surface validation remains required before closure.
