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

## Regression Adapter

Task-level regression for this repo is defined in repo-root `REGRESSION.md`.

The regression lane must start the real desktop app surface and exercise the real overlay behavior. Parser-only proof does not satisfy regression closure by itself.
