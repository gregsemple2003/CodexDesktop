# Regression Run 0001

## Test Type

`regression test`

## Claimed Lane

`REG-001 Desktop Overlay Launch And Data Smoke`

## Result

`passed`

## Actual Flow Exercised

1. From repo root, ran the supporting scripted proof:

   ```powershell
   python -m unittest discover -s tests -p "test_*.py" -v
   python -m app.codex_dashboard --scan-once --print-summary
   ```

2. From repo root, launched the real desktop app surface in smoke capture mode:

   ```powershell
   python -m app.codex_dashboard --smoke-artifact-dir Tracking/Task-0002/Testing/REGRESSION-RUN-0001
   ```

3. While the real Tk overlay was visible, captured the live desktop surface to `desktop-overlay.png`.
4. Allowed the app to finish its smoke capture and exit cleanly.

## Why This Counts

- The run starts from the real Windows desktop app entrypoint, not from a parser-only or CLI-only helper.
- The app ingested the real local `C:\Users\gregs\.codex` telemetry tree during the run.
- The overlay was opened through the real hotkey path on this run:
  - `hotkey_triggered=True`
  - `overlay_fallback=False`
- The task-owned app-surface artifacts show the required lane claims:
  - selected interval: `15m`
  - visible bar data
  - weekly budget state and redline status

## Evidence

- `Tracking/Task-0002/Testing/REGRESSION-RUN-0001/desktop-overlay.png`
- `Tracking/Task-0002/Testing/REGRESSION-RUN-0001/overlay-summary.txt`
- `Tracking/Task-0002/Testing/REGRESSION-RUN-0001/overlay-chart.ps`

Observed runtime summary:

```text
Last ingest 22:06:17 | files 2/214 | events +1
interval=15m
weekly_budget=8000000
startup_enabled=False
7d_total=1.8B
projected=3.4B
current_bucket=5M
status=Redline
Codex advisory window: 50.0% used | reset epoch 1775638824
hotkey_triggered=True
overlay_fallback=False
```

Supporting scripted proof:

- `13` unit tests passed
- ingest core smoke printed a real summary after updating SQLite from the local Codex telemetry tree

## Disqualifiers / Limitations

- The regression run uses the app's smoke harness to automate opening and artifact capture; it does not prove a long manual operator session.
- The screenshot artifact is a full desktop capture with the overlay visible rather than a native per-window export.
- These limitations do not disqualify `REG-001` because the repo-root lane requires a real app-surface launch, one ingest cycle, overlay visibility, and a captured artifact showing interval data plus budget state; all of those claims are directly evidenced here.
