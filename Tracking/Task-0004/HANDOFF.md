# Task 0004 Handoff

## Current Status

`Task-0004` is complete.

`PASS-0002` closed the task with passing validation and repo-canonical regression, and a same-task hotfix restored the `Ctrl+Alt+Space` close path before final checkpointing.

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
- explicit overlay visibility tracking so the global hotkey can reliably hide the borderless dashboard again

## Validation Status

Final proof passed:

- `python -m unittest discover -s tests -p "test_*.py" -v`
- `python -m app.codex_dashboard --scan-once --print-summary`
- `python -m app.codex_dashboard --smoke-artifact-dir Tracking/Task-0004/Testing/PASS-0002-REG-001-0003`
- `python -m app.codex_dashboard --smoke-artifact-dir Tracking/Task-0004/Testing/PASS-0002-JOBS-SMOKE-0004 --smoke-tab jobs`

Regression status:

- repo-canonical `REG-001` passed
- additive `Jobs` lane smoke passed on the real app path
- focused desktop-support coverage passed for the restored hotkey close path

## Next Step

No further task work is pending under `Task-0004`.

## Watchouts

- preserve the hotkey-first overlay behavior and keep `Usage` as the default tab
- keep the explicit overlay visibility flag aligned with any future show/hide path changes; do not fall back to Tk window-state checks for toggle truth
- keep visible copy human-facing; do not expose operator acronyms on the default surface
- keep `Logs` and `Terminal` inactive unless a later task explicitly pulls them into scope

## References

- `Tracking/Task-0004/TASK.md`
- `Tracking/Task-0004/PLAN.md`
- `Tracking/Task-0004/RESEARCH.md`
- `Tracking/Task-0004/Testing/PASS-0000-AUDIT.md`
- `Tracking/Task-0004/Testing/PASS-0001-AUDIT.md`
- `Tracking/Task-0004/Testing/PASS-0002-AUDIT.md`
- `C:\Users\gregs\.codex\Orchestration\codex-jobs-registry.json`
