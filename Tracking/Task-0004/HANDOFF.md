# Task 0004 Handoff

## Current Status

`Task-0004` is active.

`PASS-0001` is complete and `PASS-0002` is now the active closure pass.

## Current Baseline

The repo now has both the backend jobs model and the first additive `Jobs` lane in the overlay.

Backend foundation:

- managed jobs registry bootstrap under `C:\Users\gregs\.codex\Orchestration\codex-jobs-registry.json`
- supported job kinds:
  - Startup-folder launcher
  - Scheduled Task
- reconciliation states for in-sync, missing, drifted, disabled, and blocked
- bounded apply behavior for the supported job kinds
- bootstrap wired into the existing CLI and app startup path

The live machine bootstrap already imported:

- the CodexDashboard startup launcher
- `Codex Daily Agentic SWE Digest`
- `Codex Daily Physical Agents Digest`
- `Codex Daily UE Determinism Digest`

Overlay surface now includes:

- `Usage` as the default tab
- additive `Jobs` tab
- summary cards for declared jobs, in-sync jobs, needs-attention count, and last reconciliation time
- per-job rows with mechanism, desired versus observed state, drift status, and visible reason text
- bounded `Refresh state` and `Reconcile supported drift` actions
- inert `Logs` and `Terminal` placeholders only

## Validation Status

Pass-local proof passed:

- `python -m unittest discover -s tests -p "test_*.py" -v`
- `python -m app.codex_dashboard --scan-once --print-summary`
- `python -m app.codex_dashboard --smoke-artifact-dir Tracking/Task-0004/Testing/PASS-0001-UI-SMOKE-0002 --smoke-tab jobs`

Task-level regression has not run yet.

## Next Step

Execute `PASS-0002`:

- run repo-canonical regression lane `REG-001` from the real desktop app surface
- capture honest closure evidence for the additive `Jobs` lane alongside the existing usage surface
- close task artifacts once validation and regression status are durable

## Watchouts

- preserve the hotkey-first overlay behavior and keep `Usage` as the default tab
- keep visible copy human-facing; do not expose operator acronyms on the default surface
- keep `Logs` and `Terminal` inactive unless a later task explicitly pulls them into scope
- treat smoke proof as supporting evidence only; do not substitute it for repo-canonical regression

## References

- `Tracking/Task-0004/TASK.md`
- `Tracking/Task-0004/PLAN.md`
- `Tracking/Task-0004/RESEARCH.md`
- `Tracking/Task-0004/Testing/PASS-0000-AUDIT.md`
- `Tracking/Task-0004/Testing/PASS-0001-AUDIT.md`
- `C:\Users\gregs\.codex\Orchestration\codex-jobs-registry.json`
