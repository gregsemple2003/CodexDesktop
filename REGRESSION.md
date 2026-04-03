# CodexDashboard Regression Checklist

This file is the canonical regression matrix for CodexDashboard.

## Canonical Rule

Task-level regression in this repo starts from the real desktop app surface.

Supporting parser or CLI proof is useful, but it does not replace the required app-surface lane for closure.

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
