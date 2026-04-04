# CodexDashboard Regression Checklist

This file is the canonical regression matrix for CodexDashboard.

## Canonical Rule

Task-level regression in this repo starts from the real desktop app surface.

Supporting parser or CLI proof is useful, but it does not replace the required app-surface lane for closure.

When current human-facing functionality changes, this repo-local `REGRESSION.md` must be updated so the changed interaction is covered by a named case. Do not treat a new clickable surface as covered by an older case unless the written steps and expected result explicitly include that interaction.

## Default Regression Lane

### REG-001 Desktop Overlay Launch And Data Smoke

Goal:

Confirm the real desktop app can ingest Codex telemetry, open the real overlay, and render interval data plus budget state.

Steps:

1. Launch the real app from repo root.
2. Point it at the real `C:\Users\gregs\.codex` tree or a fixture tree with real `token_count` events.
3. Allow at least one ingest cycle to complete.
4. Trigger the real overlay path.
5. Capture an artifact from the running app surface that shows:
   - the selected interval
   - bar data
   - weekly budget state
6. Exit cleanly.

Expected result:

- the app starts without crashing
- the ingest loop persists real token events
- the overlay becomes visible
- interval data appears
- budget and redline state appear

## Required Current Cases

### REG-002 Jobs Tab Interaction And Status Surface

Goal:

Confirm the real desktop app can switch from the default `Usage` surface to the `Jobs` surface through the real tab click path, then explicitly refresh managed-job state without disruptive side effects.

Steps:

1. Launch the real app from repo root.
2. Point it at the real `C:\Users\gregs\.codex` tree or a fixture tree with real `token_count` events.
3. Allow at least one ingest cycle to complete.
4. Trigger the real overlay path.
5. Click the `Jobs` tab in the running overlay.
6. Verify the tab switch completes immediately and does not trigger jobs refresh or reconcile work on its own.
7. Click `Refresh` in the `Jobs` surface and observe the jobs state load to completion.
8. Capture an artifact from the running app surface that shows:
   - the `Jobs` tab as the active surface
   - job summary counts
   - per-job status rows
9. Exit cleanly.

Expected result:

- the tab switch completes without hitching the overlay
- no extra console or app windows are spawned by the interaction
- clicking `Jobs` alone does not trigger refresh or reconcile work
- the `Jobs` surface renders declared jobs, in-sync count, needs-attention count, and per-job rows
- visible copy remains human-facing

## Supporting Smoke

### SMOKE-001 Ingest Core

Run:

```powershell
python -m app.codex_dashboard --scan-once --print-summary
```

Expected result:

- the app reads real-shaped telemetry
- SQLite persistence is updated
- a human-readable summary prints

Interpretation:

- this is supporting proof only
- it does not replace `REG-001`
