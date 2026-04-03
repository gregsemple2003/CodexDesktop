# Regression Run 0001

## Test Type

`regression test`

## Claimed Lane

`REG-001 Desktop Overlay Launch And Data Smoke`

## Result

`passed`

## Actual Flow Exercised

1. From repo root, ran supporting scripted proof:

   ```powershell
   python -m unittest discover -s tests -p "test_*.py" -v
   python -m app.codex_dashboard --scan-once --print-summary
   ```

2. Stopped the already-running dashboard instance so the app-surface screenshot would reflect the fresh smoke run.
3. From repo root, launched the real desktop app surface in smoke capture mode:

   ```powershell
   python -m app.codex_dashboard --smoke-artifact-dir Tracking/Task-0003/Testing/REGRESSION-RUN-0001
   ```

4. While the smoke-run overlay was visible, captured a live full-desktop screenshot to `desktop-overlay.png`.
5. Allowed the smoke harness to export `overlay-chart.ps`, write `overlay-summary.txt`, and exit cleanly.

## Why This Counts

- The run starts from the real Windows desktop app entrypoint, not from a parser-only or CLI-only helper.
- The smoke summary proves the app completed a real ingest cycle during the run:
  - `Last ingest 01:17:36 | files 1/217 | events +2`
- The overlay was opened through the real hotkey path on this run:
  - `hotkey_triggered=True`
  - `overlay_fallback=False`
- The screenshot artifact shows the real overlay surface with:
  - visible interval controls
  - visible `Velocity` / `Repo` controls
  - visible `Total` / `Norm` controls
  - visible bar data and budget state

## Evidence

- `Tracking/Task-0003/Testing/REGRESSION-RUN-0001/desktop-overlay.png`
- `Tracking/Task-0003/Testing/REGRESSION-RUN-0001/overlay-summary.txt`
- `Tracking/Task-0003/Testing/REGRESSION-RUN-0001/overlay-chart.ps`

Observed smoke summary:

```text
Last ingest 01:17:36 | files 1/217 | events +2
interval=15m
metric_mode=total
weekly_budget=3550000000
startup_enabled=False
7d_total=1.9B
projected=3.1B
headroom=+644.9K
budget_line=5.3M
status=Operational
Codex advisory window: 59.0% used | reset in 5.2d
hotkey_triggered=True
overlay_fallback=False
```

Supporting feature-specific proof:

- the unit suite covers normalized bucket math and normalized repo-stack totals in `tests/test_ingest_core.py`
- the unit suite covers normalized chart-title wording in `tests/test_desktop_support.py`

## Disqualifiers / Limitations

- The screenshot was captured slightly before the smoke harness wrote its final summary line, so the screenshot and summary together make the full claim rather than either artifact alone.
- The app-surface regression lane was captured in default `Total` mode; `Norm` behavior is covered by focused unit tests instead of a separate second screenshot.
- These limitations do not disqualify `REG-001` because the repo-root lane requires a real app-surface launch, one ingest cycle, overlay visibility, and captured evidence of interval data plus budget state. Those claims are directly evidenced here.
