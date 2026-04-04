# Task 0004 Handoff

## Current Status

`Task-0004` is reopened and in regression.

The original closure was invalidated by the live `Jobs` interaction regression and the mockup-fidelity miss. The current repo state contains a candidate fix set with fresh supporting proof, but the task is not honestly closed again until the live overlay readback is positive.

## Current Baseline

The current candidate baseline includes:

- managed jobs registry bootstrap under `C:\Users\gregs\.codex\Orchestration\codex-jobs-registry.json`
- `Jobs` tab clicks that only switch surfaces
- explicit `Refresh` and `Force Reconcile` actions for jobs-state work
- hidden child PowerShell windows for explicit jobs actions
- a separate primary-nav strip with tab-owned controls below it
- stronger tab emphasis and a quieter header closer to the approved Stitch direction
- a redline label treatment that matches the approved chart language more closely than the old right-edge utility annotation

## Validation Status

Executed successfully:

- `python -m unittest discover -s tests -p "test_*.py" -v`
- `python -m app.codex_dashboard --scan-once --print-summary`
- `python -m app.codex_dashboard --smoke-artifact-dir Tracking/Task-0004/Testing/PASS-0002-REG-001-0004`
- `python -m app.codex_dashboard --smoke-artifact-dir Tracking/Task-0004/Testing/PASS-0002-JOBS-SMOKE-0005 --smoke-tab jobs`

Current evidence:

- [PASS-0002-REG-001-0004/desktop-overlay.png](/c:/Agent/CodexDashboard/Tracking/Task-0004/Testing/PASS-0002-REG-001-0004/desktop-overlay.png)
- [PASS-0002-REG-001-0004/overlay-summary.txt](/c:/Agent/CodexDashboard/Tracking/Task-0004/Testing/PASS-0002-REG-001-0004/overlay-summary.txt)
- [PASS-0002-JOBS-SMOKE-0005/desktop-overlay.png](/c:/Agent/CodexDashboard/Tracking/Task-0004/Testing/PASS-0002-JOBS-SMOKE-0005/desktop-overlay.png)
- [PASS-0002-JOBS-SMOKE-0005/overlay-summary.txt](/c:/Agent/CodexDashboard/Tracking/Task-0004/Testing/PASS-0002-JOBS-SMOKE-0005/overlay-summary.txt)

## Next Step

Run the live reopened regression readback on the actual overlay:

1. `Ctrl+Alt+Space`
2. click `Jobs`
3. confirm there is no hitch and no extra windows
4. click `Refresh`
5. confirm the rendered `Jobs` surface is close enough to the approved mockup direction

If that readback is positive, update the reopened pass artifacts one more time and close the task honestly.

## Watchouts

- keep `Jobs` refresh and reconcile off the tab click path
- keep repo-local [REGRESSION.md](/c:/Agent/CodexDashboard/REGRESSION.md) aligned with the actual live interaction model
- keep visible copy human-facing and keep raw Windows plumbing behind details
- do not let future shell controls drift back into the primary nav strip

## References

- [TASK.md](/c:/Agent/CodexDashboard/Tracking/Task-0004/TASK.md)
- [PLAN.md](/c:/Agent/CodexDashboard/Tracking/Task-0004/PLAN.md)
- [BUG-0001.md](/c:/Agent/CodexDashboard/Tracking/Task-0004/BUG-0001.md)
- [PASS-0002-AUDIT.md](/c:/Agent/CodexDashboard/Tracking/Task-0004/Testing/PASS-0002-AUDIT.md)
