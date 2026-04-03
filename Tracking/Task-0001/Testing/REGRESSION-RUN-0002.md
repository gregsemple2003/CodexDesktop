# Regression Run 0002

## Test Type

`regression test`

## Claimed Lane

Repo-root `REGRESSION.md`:

- `REG-001 Desktop Overlay Launch And Data Smoke`

## Environment

- date: `2026-04-02`
- repo: `C:\Agent\CodexDashboard`
- runtime: local Python 3.13 on Windows
- codex root: `C:\Users\gregs\.codex`
- launcher:
  - `python -m app.codex_dashboard`

## Flow Exercised

Human-executed manual regression slice:

1. launch the real app from repo root
2. wait for the process to stay open
3. press `Ctrl+Alt+Space`
4. observe whether the overlay becomes visible

## Result

`FAIL`

## Evidence

- the app did open a console window and remained running
- pressing `Ctrl+Alt+Space` did nothing
- the overlay did not appear
- this was the first real manual keyboard confirmation attempt after regression run `0001`

## Why This Run Counts

This run uses the real desktop app surface and a real keyboard interaction on the target Windows host, so it directly exercises the regression claim that matters for closure.

## Failure Summary

The hotkey path failed in the live product:

- the app launched
- the process stayed alive
- the required hotkey chord did not show the overlay

That means the task is in debugging, not just waiting on extra regression proof.

## What This Run Does Not Prove

- whether the ingest loop had already completed a scan before the manual hotkey press
- whether `Escape` dismissal works once the overlay is visible
- startup-folder behavior across logon

## Next Step

Debug and fix the real hotkey path, then rerun `REG-001` with a fresh manual confirmation that:

1. `Ctrl+Alt+Space` shows the overlay
2. `Ctrl+Alt+Space` or `Escape` hides it
