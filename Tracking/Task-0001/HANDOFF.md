# Task 0001 Handoff

## Current Status

Planning is complete, auto-approved, and `PASS-0000` is now the active implementation pass.

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

## Recommended Next Step

Execute `PASS-0000`:

- add repo-root docs for this new repo
- implement the ingest engine, SQLite store, aggregation logic, and weekly redline math
- add focused unit coverage for event parsing, dedupe, and bucket aggregation

## Watchouts

- do not build the first version around file counts; the signal should come from parsed `token_count` events
- do not promise exact provider-budget parity unless the telemetry proves it
- do not over-design the chart surface before the ingest and aggregation model is correct
- do not make the first version depend on a heavyweight foreground app workflow if the hotkey overlay can solve the check-in use case directly
- do not quietly revert back to Tauri unless the toolchain is actually bootstrapped and that change is reflected durably
- keep the implementation organized so a future Tauri port remains possible if the prototype succeeds

## Key References

- `Tracking/Task-0001/TASK.md`
- `Tracking/Task-0001/RESEARCH-PLAN.md`
- `Tracking/Task-0001/RESEARCH-ANALYSIS.md`
- `Tracking/Task-0001/RESEARCH.md`
- `Tracking/Task-0001/PLAN.md`
- `C:\Users\gregs\.codex\AGENTS.md`
- `C:\Users\gregs\.codex\Orchestration\ORCHESTRATION.md`
- sampled session telemetry under `C:\Users\gregs\.codex\sessions\`
