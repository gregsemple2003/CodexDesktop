# Task 0001 Handoff

## Current Status

`Task-0001` is complete locally. `PASS-0002` closed after the hotkey bug was fixed and the repo-root regression lane passed.

## Final Baseline

What this task delivered:

- a Windows-first Codex token-velocity dashboard under `app/codex_dashboard/`
- polling ingest over `C:\Users\gregs\.codex\sessions\`
- SQLite-backed cursor and token-event persistence
- interval aggregation for `1m`, `5m`, `15m`, `1h`, and `1d`
- a Tk overlay with bar-chart rendering, weekly budget editing, advisory usage context, and redline presentation
- Windows Startup-folder integration
- a working global hotkey toggle backed by a dedicated Win32 message-loop thread

Final proof state:

- `PASS-0000` delivered the ingest core and unit coverage
- `PASS-0001` delivered the real desktop surface and integration baseline
- `REGRESSION-RUN-0002.md` preserved the failing manual hotkey evidence
- `BUG-0001.md` captured the root cause and resolution
- `REGRESSION-RUN-0003.md` passed the repo-root `REG-001 Desktop Overlay Launch And Data Smoke` lane
- `PASS-0002-AUDIT.md` closed the final pass as `ready`

## Remaining Limits

- the final passing rerun used a synthetic system key event rather than a new physical-keyboard spot check after the fix
- the repo is intentionally still local-only; no upstream push was attempted by design
- startup behavior across a real Windows logon and long-running idle stability remain future hardening work, not this task's closure bar

## Recommended Next Step

No further work is required for this task to stay honest and complete locally.

If the prototype is prepared for upstream promotion later:

- run one quick physical-keyboard spot check of `Ctrl+Alt+Space`
- decide whether to keep the current Python/Tk prototype as-is or package it behind `pythonw` / a tray-first launcher
- then push the local git history upstream when you want the prototype published

## Watchouts

- keep the ingest model grounded in parsed `token_count` events, not file counts
- keep the advisory Codex weekly-window metadata labeled as advisory rather than exact quota math
- do not reintroduce Tk-main-thread hotkey polling; the dedicated message-loop thread is the current fix

## Key References

- `Tracking/Task-0001/TASK.md`
- `Tracking/Task-0001/PLAN.md`
- `Tracking/Task-0001/BUG-0001.md`
- `Tracking/Task-0001/TASK-STATE.json`
- `Tracking/Task-0001/Testing/PASS-0000-AUDIT.md`
- `Tracking/Task-0001/Testing/PASS-0000-AUDIT.json`
- `Tracking/Task-0001/Testing/PASS-0000-CHECKLIST.json`
- `Tracking/Task-0001/Testing/PASS-0001-AUDIT.md`
- `Tracking/Task-0001/Testing/PASS-0001-AUDIT.json`
- `Tracking/Task-0001/Testing/PASS-0001-CHECKLIST.json`
- `Tracking/Task-0001/Testing/PASS-0002-AUDIT.md`
- `Tracking/Task-0001/Testing/PASS-0002-AUDIT.json`
- `Tracking/Task-0001/Testing/PASS-0002-CHECKLIST.json`
- `Tracking/Task-0001/Testing/REGRESSION-RUN-0002.md`
- `Tracking/Task-0001/Testing/REGRESSION-RUN-0003.md`
- `Tracking/Task-0001/Testing/REGRESSION-UI-SMOKE-0002/overlay-chart.ps`
- `Tracking/Task-0001/Testing/REGRESSION-UI-SMOKE-0002/overlay-summary.txt`
- `Design/GENERAL-DESIGN.md`
