# Regression Run 0001

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
  - `Tracking/Task-0001/Testing/REGRESSION-UI-SMOKE-0001/`

## Flow Exercised

Executed from repo root:

```powershell
python -m app.codex_dashboard --smoke-artifact-dir Tracking/Task-0001/Testing/REGRESSION-UI-SMOKE-0001
python -m unittest discover -s tests -p "test_*.py" -v
```

What the run actually exercised:

- real desktop app launch
- real ingest against the local `.codex` session tree
- real Tk overlay render
- real chart artifact export
- real summary artifact export
- clean smoke-process exit

## Result

`BLOCKED`

## Evidence

- `Tracking/Task-0001/Testing/REGRESSION-UI-SMOKE-0001/overlay-chart.ps`
- `Tracking/Task-0001/Testing/REGRESSION-UI-SMOKE-0001/overlay-summary.txt`
- `overlay-summary.txt` reported:
  - `Last ingest 20:48:39 | files 4/209 | events +15`
  - `1,743,649,371` local tokens in the last 7 days
  - `Projected burn: 8,975,443,344 REDLINE`
  - `Codex advisory window: 44.0% used | reset epoch 1775638824`
  - `hotkey_triggered=False`
  - `overlay_fallback=True`
- supporting unit proof:
  - `7` tests passed via `python -m unittest discover -s tests -p "test_*.py" -v`

## Why This Run Counts

This run does count as a real app-surface regression slice because it starts the real desktop app and captures evidence from the visible overlay surface, not only from parser code or a CLI helper.

## Why It Is Blocked

The executed smoke did not prove that the actual global hotkey chord toggled the overlay:

- the synthetic hotkey trigger did not fire the registered `WM_HOTKEY` path
- the smoke harness fell back to opening the overlay directly so the chart artifact could still be captured

That means the run proves the real overlay surface and live data path, but not the final hotkey interaction claim.

## What This Run Does Not Prove

- that a real keyboard press of `Ctrl+Alt+Space` toggles the overlay on this machine
- that startup integration survives a real logon or reboot
- long-running background stability over hours or days

## Next Step

Run one human-confirmed regression slice for the hotkey interaction:

1. launch the app normally
2. press `Ctrl+Alt+Space`
3. confirm the overlay appears
4. press `Ctrl+Alt+Space` again or `Escape`
5. confirm the overlay hides

If that succeeds, update the regression artifact or add a follow-up regression run and close the task.
