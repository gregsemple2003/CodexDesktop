# Task 0004 Handoff

## Current Status

`Task-0004` is active.

`PASS-0000` is complete and `PASS-0001` is now the active implementation pass.

## Current Baseline

The repo now has a backend jobs model in place:

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

## Validation Status

Pass-local proof passed:

- `python -m unittest tests.test_jobs -v`
- `python -m unittest discover -s tests -p "test_*.py" -v`
- `python -m app.codex_dashboard --scan-once --print-summary`

Task-level regression has not run yet.

## Next Step

Execute `PASS-0001`:

- add the `Jobs` tab while keeping `Usage` as the default surface
- surface summary counts, per-job state, last reconciliation time, and bounded actions
- keep raw Windows plumbing behind explicit details
- run the required UI-fidelity review against `Tracking/Task-0004/Design/STITCH-JOBS-TAB-0001/`

## Watchouts

- preserve the hotkey-first overlay behavior and keep `Usage` as the default tab
- keep visible copy human-facing; do not expose operator acronyms on the default surface
- treat the approved Stitch assets as the binding UI direction for `PASS-0001`
- keep `Logs` and `Terminal` inactive unless a later task explicitly pulls them into scope

## References

- `Tracking/Task-0004/TASK.md`
- `Tracking/Task-0004/PLAN.md`
- `Tracking/Task-0004/RESEARCH.md`
- `Tracking/Task-0004/Testing/PASS-0000-AUDIT.md`
- `C:\Users\gregs\.codex\Orchestration\codex-jobs-registry.json`
