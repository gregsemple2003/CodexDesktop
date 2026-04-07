# Testing Guide

CodexDashboard follows the shared lifecycle and glossary from:

- `C:\Users\gregs\.codex\AGENTS.md`
- `C:\Users\gregs\.codex\Orchestration\ORCHESTRATION.md`
- `C:\Users\gregs\.codex\Orchestration\Processes\TESTING.md`

This repo-local file only defines CodexDashboard-specific commands and evidence expectations.

## Unit Test Commands

From repo root:

```powershell
python -m unittest discover -s tests -p "test_*.py" -v
```

Use unit coverage for:

- line parsing
- cursor handling
- dedupe
- interval aggregation
- weekly redline math

## Supporting Scripted Smoke

From repo root:

```powershell
python -m app.codex_dashboard --scan-once --print-summary
```

Use this to prove the ingest core can read real-shaped `.jsonl` telemetry and persist token events into SQLite.

## Backend Lane Discipline

Do not use the always-on service lane for disposable backend proof work.

Use the validation lane under `backend/orchestration/scripts/` for backend smoke, regression, or debugging:

```powershell
cd C:\Agent\CodexDashboard\backend\orchestration
powershell -ExecutionPolicy Bypass -File .\scripts\Start-ValidationLane.ps1
```

When a desktop run should target that validation lane instead of the default service lane:

```powershell
$env:CODEX_DASHBOARD_JOBS_BACKEND_URL = "http://127.0.0.1:14318"
python -m app.codex_dashboard
Remove-Item Env:CODEX_DASHBOARD_JOBS_BACKEND_URL -ErrorAction SilentlyContinue
```

Tear the validation lane down after the proof run:

```powershell
cd C:\Agent\CodexDashboard\backend\orchestration
powershell -ExecutionPolicy Bypass -File .\scripts\Stop-ValidationLane.ps1
```

## Regression Adapter

Task-level regression for this repo is defined in repo-root `REGRESSION.md`.

The regression lane must start the real desktop app surface and exercise the real overlay behavior. Parser-only proof does not satisfy regression closure by itself.
