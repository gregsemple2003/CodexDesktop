# Task 0001 Handoff

## Current Status

`PASS-0001` remains complete, but the task is now in active debugging during regression because a real manual hotkey run failed.

## Current Baseline

What is already grounded:

- the source files live under `C:\Users\gregs\.codex`
- recent sampled session files contain frequent `token_count` events
- those events provide both:
  - per-event token deltas through `last_token_usage.total_tokens`
  - rolling 7-day limit metadata through `rate_limits.secondary`
- the product direction should favor a background app plus summonable overlay instead of a normal foreground-heavy desktop window
- the local machine does not currently expose the Tauri toolchain on `PATH`
- the local machine does expose Python 3.13 with standard-library `tkinter`
- research therefore recommends a Windows-first Python prototype as the first working version
- the approved pass order is now:
  - `PASS-0000`: repo bootstrap and ingest core
  - `PASS-0001`: overlay UI and desktop integration
  - `PASS-0002`: regression, handoff, and closure
- `PASS-0000` delivered:
  - repo-root docs for `AGENTS`, `TESTING`, and `REGRESSION`
  - the Python package baseline under `app/codex_dashboard/`
  - SQLite-backed file cursors and token event persistence
  - interval aggregation and weekly redline math
  - five focused unit tests
  - a real-session ingest smoke against the live `.codex` tree
- `PASS-0001` delivered:
  - the repo-root `Design/GENERAL-DESIGN.md` anchor
  - the real Tk overlay surface
  - hotkey parsing and registration
  - startup integration
  - budget editing and redline presentation
  - an app-surface smoke artifact bundle under `Testing/UI-SMOKE-0001/`
- regression run `0001` proved:
  - real app launch
  - real live-data ingest
  - real overlay render
  - real artifact export
  - clean smoke exit
- regression run `0001` did not prove:
  - a real keyboard-triggered `Ctrl+Alt+Space` overlay toggle
- regression run `0002` now proves:
  - a human launched `python -m app.codex_dashboard`
  - the process stayed open
  - pressing `Ctrl+Alt+Space` did nothing
  - the overlay did not appear

## Recommended Next Step

Debug the real hotkey path in the Windows app, then rerun regression:

- inspect `RegisterHotKey` success and the live `WM_HOTKEY` polling loop
- fix the failing message-delivery path
- rerun unit tests
- rerun the desktop-app regression lane with a real keyboard confirmation

## Watchouts

- do not build the first version around file counts; the signal should come from parsed `token_count` events
- do not promise exact provider-budget parity unless the telemetry proves it
- do not over-design the chart surface before the ingest and aggregation model is correct
- do not make the first version depend on a heavyweight foreground app workflow if the hotkey overlay can solve the check-in use case directly
- do not quietly revert back to Tauri unless the toolchain is actually bootstrapped and that change is reflected durably
- keep the implementation organized so a future Tauri port remains possible if the prototype succeeds
- do not treat the hotkey issue as a harness-only limitation anymore; the live app has now failed manual regression

## Key References

- `Tracking/Task-0001/TASK.md`
- `Tracking/Task-0001/RESEARCH-PLAN.md`
- `Tracking/Task-0001/RESEARCH-ANALYSIS.md`
- `Tracking/Task-0001/RESEARCH.md`
- `Tracking/Task-0001/PLAN.md`
- `Tracking/Task-0001/Testing/PASS-0000-AUDIT.md`
- `Tracking/Task-0001/Testing/PASS-0000-AUDIT.json`
- `Tracking/Task-0001/Testing/PASS-0000-CHECKLIST.json`
- `Tracking/Task-0001/Testing/PASS-0001-AUDIT.md`
- `Tracking/Task-0001/Testing/PASS-0001-AUDIT.json`
- `Tracking/Task-0001/Testing/PASS-0001-CHECKLIST.json`
- `Tracking/Task-0001/Testing/UI-SMOKE-0001/overlay-chart.ps`
- `Tracking/Task-0001/Testing/UI-SMOKE-0001/overlay-summary.txt`
- `Tracking/Task-0001/Testing/REGRESSION-RUN-0001.md`
- `Tracking/Task-0001/Testing/REGRESSION-RUN-0002.md`
- `Tracking/Task-0001/BUG-0001.md`
- `Tracking/Task-0001/Testing/REGRESSION-UI-SMOKE-0001/overlay-chart.ps`
- `Tracking/Task-0001/Testing/REGRESSION-UI-SMOKE-0001/overlay-summary.txt`
- `Design/GENERAL-DESIGN.md`
- `C:\Users\gregs\.codex\AGENTS.md`
- `C:\Users\gregs\.codex\Orchestration\ORCHESTRATION.md`
- sampled session telemetry under `C:\Users\gregs\.codex\sessions\`
