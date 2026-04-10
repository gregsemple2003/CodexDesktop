# CodexDashboard Repo Guide

## Project Goal

Build a Windows-first desktop utility that reads local Codex session telemetry and shows recent token velocity through a hotkey-first overlay.

Current implementation target:

- Python 3.13
- Tkinter desktop UI
- SQLite persistence
- Win32 integration through `ctypes`

## Documentation Split

Use the shared workflow from:

- `C:\Users\gregs\.codex\AGENTS.md`
- `C:\Users\gregs\.codex\Orchestration\ORCHESTRATION.md`

Use repo-root docs here only for CodexDashboard-specific truth such as:

- implementation layout
- run commands
- regression lane definition
- the durable intervention canon and eval home under `C:\Users\gregs\.codex\Orchestration\Reports\Interventions\`

Keep task history and evidence under `Tracking/Task-<id>/`.
The explicit exception is:

- `C:\Users\gregs\.codex\Orchestration\Reports\Interventions\`
  - durable cross-task intervention canon, red-line, and eval-design home

Task bootstrap contracts, day-pass artifacts, and pass history should still start under the owning task until promoted.

## Implementation Layout

- `app/codex_dashboard/`
  - Python package for config, storage, ingest, aggregation, and later UI
- `C:\Users\gregs\.codex\Orchestration\Reports\Interventions\`
  - durable shared intervention canon, red-line rules, and test design
- `tests/`
  - focused unit coverage
- `Tracking/Task-0001/`
  - task-owned artifacts for the current workstream

## Build And Verification Commands

From repo root:

```powershell
python -m unittest discover -s tests -p "test_*.py" -v
python -m app.codex_dashboard --scan-once --print-summary
```

## Guardrails

- do not infer token usage from file counts or file sizes
- parse real `token_count` events from `C:\Users\gregs\.codex\sessions\`
- keep the first version Windows-only and hotkey-first
- treat repo-root `REGRESSION.md` as canonical for task-level regression in this repo
