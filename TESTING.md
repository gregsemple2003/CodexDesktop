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

Do not use the human's personal dashboard lane, app config, database, or live
Codex data for disposable backend proof, debugging, or regression work. Treat
the human's lane as off-limits unless the human explicitly authorizes touching
it for the current run.

Do not equate `default lane` with `human lane`. A default backend lane can be a
persistent product service lane; the human lane is whatever lane, config,
database, and data the human is actively using. Regression for task closure must
run on a separate isolated validation or task-specific regression lane.

Known lanes:

| Lane | Purpose | Backend URL | Temporal | Postgres | Data root | Default use |
| --- | --- | --- | --- | --- | --- | --- |
| service | persistent backend service lane | `http://127.0.0.1:4318` | `127.0.0.1:7233` | `5432` | `%LOCALAPPDATA%\CodexDashboard\orchestration-service-lane` | not a regression lane; do not use for task-closure regression unless explicitly authorized |
| validation | disposable agent proof lane | `http://127.0.0.1:14318` | `127.0.0.1:17233` | `15432` | `%LOCALAPPDATA%\CodexDashboard\orchestration-validation-lane` | default backend lane for agent smoke, debugging, and regression |

Any additional reusable lane must be documented here before relying on it for
task proof. The documentation must name its ports, runtime/data roots, start and
stop commands, and how it stays isolated from the human's service lane.

Use the validation lane under `backend/orchestration/scripts/` for backend smoke, regression, or debugging:

```powershell
cd C:\Agent\CodexDashboard\backend\orchestration
powershell -ExecutionPolicy Bypass -File .\scripts\Start-ValidationLane.ps1
```

When a desktop run should target that validation lane instead of any persistent service lane:

```powershell
$env:CODEX_DASHBOARD_JOBS_BACKEND_URL = "http://127.0.0.1:14318"
python -m app.codex_dashboard
Remove-Item Env:CODEX_DASHBOARD_JOBS_BACKEND_URL -ErrorAction SilentlyContinue
```

Do not run desktop regression against the human's dashboard config or database.
Create a task-owned or temp config that points at fixture telemetry and an
isolated SQLite file, then launch the app with that config:

```powershell
$root = "C:\Agent\CodexDashboard\Tracking\Task-<id>\Testing\Runtime"
New-Item -ItemType Directory -Force $root | Out-Null
@{
  codex_root = "$root\codex-fixture"
  db_path = "$root\dashboard-regression.db"
  polling_seconds = 5
  weekly_budget_tokens = 4000000000
  startup_enabled = $false
  hotkey = "Ctrl+Alt+Space"
} | ConvertTo-Json | Set-Content -Encoding UTF8 "$root\config.json"
$env:CODEX_DASHBOARD_JOBS_BACKEND_URL = "http://127.0.0.1:14318"
python -m app.codex_dashboard --config-path "$root\config.json"
Remove-Item Env:CODEX_DASHBOARD_JOBS_BACKEND_URL -ErrorAction SilentlyContinue
```

The fixture `codex_root` must contain only synthetic or task-owned
real-shaped telemetry. Do not point regression at `C:\Users\gregs\.codex`, the
human's active Codex data, or the human's active dashboard database unless the
human explicitly authorizes that run.

Tear the validation lane down after the proof run:

```powershell
cd C:\Agent\CodexDashboard\backend\orchestration
powershell -ExecutionPolicy Bypass -File .\scripts\Stop-ValidationLane.ps1
```

## Regression Adapter

Task-level regression for this repo is defined in repo-root `REGRESSION.md`.

The regression lane must start the real desktop app surface and exercise the real overlay behavior. Parser-only proof does not satisfy regression closure by itself.
