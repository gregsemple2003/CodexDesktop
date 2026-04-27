# BUG-0010 Validation

Status: ready for human retest

Date: 2026-04-27

## Scope

Validation for [BUG-0010](../BUG-0010.md), after the human retest proved that
`Open Handoff` and `Open Plan` were acceptable but `Open Task` still opened a
new VSCodium window without a document.

## Commands

```powershell
go test ./...
```

Working directory: `backend/orchestration`

Result: passing.

```powershell
python -m unittest tests.test_tasks_backend tests.test_desktop_support -v
```

Result: passing, 53 tests.

```powershell
python -m unittest discover -s tests -p "test_*.py" -v
```

Result: passing, 105 tests.

## Backend Readback

Validation lane:

- backend URL: `http://127.0.0.1:14318`
- Temporal: `127.0.0.1:17233`

`Task-0011` deep-context readback now reports:

```text
Open Task    -> C:\Agent\CodexDashboard\Tracking\Task-0011\TASK.md
Open Handoff -> C:\Agent\CodexDashboard\Tracking\Task-0011\HANDOFF.md
Open Plan    -> C:\Agent\CodexDashboard\Tracking\Task-0011\PLAN.md
```

## Human Retest

The separate validation dashboard was restarted from the working tree against
the validation backend.

- process: `pythonw.exe` PID `45788`
- config:
  `Tracking/Task-0009/Testing/Manual-TestDashboard/config.json`
- backend: `http://127.0.0.1:14318`
- hotkey: `Ctrl+Alt+Shift+Space`

Expected human-visible result:

- `Open Task` raises `TASK.md` in VSCodium
- `Open Handoff` raises `HANDOFF.md`
- `Open Plan` raises `PLAN.md`
