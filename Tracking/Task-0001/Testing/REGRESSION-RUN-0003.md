# Regression Run 0003

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
- artifact directory:
  - `Tracking/Task-0001/Testing/REGRESSION-UI-SMOKE-0002/`

## Flow Exercised

Executed from repo root:

```powershell
python -m unittest discover -s tests -p "test_*.py" -v
python -m app.codex_dashboard --smoke-artifact-dir Tracking/Task-0001/Testing/REGRESSION-UI-SMOKE-0002
```

What the run actually exercised:

- real desktop app launch
- real ingest against the local `.codex` session tree
- real global hotkey registration through `RegisterHotKey`
- real `WM_HOTKEY` handling on the dedicated message-loop thread
- real overlay toggle from the hotkey path
- real chart artifact export
- real summary artifact export
- clean smoke-process exit

## Result

`PASS`

## Evidence

- `Tracking/Task-0001/Testing/REGRESSION-UI-SMOKE-0002/overlay-chart.ps`
- `Tracking/Task-0001/Testing/REGRESSION-UI-SMOKE-0002/overlay-summary.txt`
- `overlay-summary.txt` reported:
  - `Last ingest 20:59:50 | files 2/209 | events +2`
  - `1,753,211,106` local tokens in the last 7 days
  - `Projected burn: 11,294,040,072 REDLINE`
  - `Codex advisory window: 46.0% used | reset epoch 1775638824`
  - `hotkey_triggered=True`
  - `overlay_fallback=False`
- supporting unit proof:
  - `8` tests passed via `python -m unittest discover -s tests -p "test_*.py" -v`

## Why This Run Counts

This run starts the real desktop app, ingests live Codex data, and reaches the visible overlay through the actual registered hotkey path rather than opening the overlay directly.

That matches the repo-root regression lane for this repo:

- the app starts
- the ingest loop updates the local store
- the overlay becomes visible
- the chart and budget state are captured from the running app surface

## What This Run Does Not Prove

- a fresh human physical-keyboard spot check after the fix
- startup-folder behavior across a real Windows logon
- long-running background stability over hours or days

## Closure Claim

This run satisfies the required repo-root regression lane for task closure because it exercises `REG-001 Desktop Overlay Launch And Data Smoke` on the real app surface and proves the hotkey path now opens the overlay without smoke fallback.
