# Pass 0002 Audit

## Scope

`PASS-0002` was reopened after the live `Jobs` regression reported in [BUG-0001](/c:/Agent/CodexDashboard/Tracking/Task-0004/BUG-0001.md).

This reopened pass now includes:

- removing jobs refresh and reconcile work from the `Jobs` tab click path
- hiding child PowerShell windows during explicit jobs actions
- recomposing the overlay so the primary nav is its own top strip and `Usage` controls live in the tab content area
- tightening tab/header fidelity toward the approved Stitch direction
- restyling the chart redline label closer to the approved mockup treatment

## Verification

Executed from repo root:

```powershell
python -m unittest discover -s tests -p "test_*.py" -v
python -m app.codex_dashboard --scan-once --print-summary
python -m app.codex_dashboard --smoke-artifact-dir Tracking/Task-0004/Testing/PASS-0002-REG-001-0004
python -m app.codex_dashboard --smoke-artifact-dir Tracking/Task-0004/Testing/PASS-0002-JOBS-SMOKE-0005 --smoke-tab jobs
```

Observed result:

- full unit coverage passed: `47` tests
- the supporting `--scan-once --print-summary` path still completed successfully against the live telemetry tree
- refreshed `Usage` smoke evidence was captured in:
  - [desktop-overlay.png](/c:/Agent/CodexDashboard/Tracking/Task-0004/Testing/PASS-0002-REG-001-0004/desktop-overlay.png)
  - [overlay-summary.txt](/c:/Agent/CodexDashboard/Tracking/Task-0004/Testing/PASS-0002-REG-001-0004/overlay-summary.txt)
  - [overlay-chart.ps](/c:/Agent/CodexDashboard/Tracking/Task-0004/Testing/PASS-0002-REG-001-0004/overlay-chart.ps)
- refreshed `Jobs` smoke evidence was captured in:
  - [desktop-overlay.png](/c:/Agent/CodexDashboard/Tracking/Task-0004/Testing/PASS-0002-JOBS-SMOKE-0005/desktop-overlay.png)
  - [overlay-summary.txt](/c:/Agent/CodexDashboard/Tracking/Task-0004/Testing/PASS-0002-JOBS-SMOKE-0005/overlay-summary.txt)
- the new jobs smoke explicitly refreshes jobs state after selecting the `Jobs` surface, so the captured summary reflects the new button-driven interaction model instead of an implicit tab-click refresh

## Requirement Mapping

| Requirement | Evidence | Result |
| --- | --- | --- |
| `Jobs` tab click path must not trigger refresh or reconcile work | `app/codex_dashboard/ui.py`; `tests/test_desktop_support.py` | Passed |
| Explicit jobs actions must not surface extra shell windows | `app/codex_dashboard/jobs.py`; `tests/test_jobs.py` | Passed |
| Primary nav must be its own strip with tab-owned controls below | `app/codex_dashboard/ui.py`; refreshed desktop overlay screenshots | Passed |
| Supporting runtime proof must reflect the current `.codex\\Orchestration` jobs location | [overlay-summary.txt](/c:/Agent/CodexDashboard/Tracking/Task-0004/Testing/PASS-0002-JOBS-SMOKE-0005/overlay-summary.txt) | Passed |
| Full unit coverage remains green after the reopened UI/debug pass | `python -m unittest discover -s tests -p "test_*.py" -v` | Passed |

## Caveat

The reopened pass is not closed yet.

The repo-local `REG-002` definition now expects the real live flow to be:

1. open the overlay
2. click `Jobs`
3. verify there is no hitch and no extra windows
4. click `Refresh`
5. verify the rendered jobs state and fidelity

The current artifacts are strong supporting runtime proof, but final reopened closeout still needs live readback on that real click path.

## Verdict

`ready_with_caveats`
