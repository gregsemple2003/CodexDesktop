# Pass 0002 Audit

## Scope

`PASS-0002` remains reopened after the live `Jobs` regression reported in [BUG-0001](/c:/Agent/CodexDashboard/Tracking/Task-0004/BUG-0001.md).

This latest reopened slice adds the user-requested UI follow-up:

- move the budget editor onto the same top `Usage` content strip as the interval controls, right-justified
- remove the visible last-ingest line from the `Usage` surface
- restyle the chart threshold as a dotted blue `BUDGET` line
- remove inactive `LOGS` and `TERMINAL` placeholders from the primary nav strip

Previously fixed in the same reopened pass:

- `Jobs` tab clicks now only switch surfaces
- explicit jobs actions hide child PowerShell windows
- the nav/header split matches the approved structure more closely

## Verification

Executed from repo root:

```powershell
python -m unittest discover -s tests -p "test_*.py" -v
python -m app.codex_dashboard --scan-once --print-summary
python -m app.codex_dashboard --smoke-artifact-dir Tracking/Task-0004/Testing/PASS-0002-REG-001-0005
python -m app.codex_dashboard --smoke-artifact-dir Tracking/Task-0004/Testing/PASS-0002-JOBS-SMOKE-0006 --smoke-tab jobs
```

Observed result:

- full unit coverage passed: `47` tests
- the supporting `--scan-once --print-summary` path still completed successfully against the live telemetry tree
- refreshed `Usage` runtime evidence was captured in:
  - [desktop-overlay.png](/c:/Agent/CodexDashboard/Tracking/Task-0004/Testing/PASS-0002-REG-001-0005/desktop-overlay.png)
  - [overlay-summary.txt](/c:/Agent/CodexDashboard/Tracking/Task-0004/Testing/PASS-0002-REG-001-0005/overlay-summary.txt)
  - [overlay-chart.ps](/c:/Agent/CodexDashboard/Tracking/Task-0004/Testing/PASS-0002-REG-001-0005/overlay-chart.ps)
- refreshed `Jobs` runtime evidence was captured in:
  - [desktop-overlay.png](/c:/Agent/CodexDashboard/Tracking/Task-0004/Testing/PASS-0002-JOBS-SMOKE-0006/desktop-overlay.png)
  - [overlay-summary.txt](/c:/Agent/CodexDashboard/Tracking/Task-0004/Testing/PASS-0002-JOBS-SMOKE-0006/overlay-summary.txt)

## Requirement Mapping

| Requirement | Evidence | Result |
| --- | --- | --- |
| Budget editor sits on the top `Usage` content strip, right-justified | [ui.py](/c:/Agent/CodexDashboard/app/codex_dashboard/ui.py); refreshed usage screenshot | Passed |
| Visible last-ingest line is removed from the `Usage` surface | [ui.py](/c:/Agent/CodexDashboard/app/codex_dashboard/ui.py); refreshed usage screenshot | Passed |
| Chart threshold is a dotted blue `BUDGET` line | [ui.py](/c:/Agent/CodexDashboard/app/codex_dashboard/ui.py); refreshed usage screenshot | Passed |
| Inactive placeholder tabs are removed from primary nav | [ui.py](/c:/Agent/CodexDashboard/app/codex_dashboard/ui.py); refreshed usage screenshot | Passed |
| `Jobs` tab click path still does not trigger refresh or reconcile work | [ui.py](/c:/Agent/CodexDashboard/app/codex_dashboard/ui.py); [test_desktop_support.py](/c:/Agent/CodexDashboard/tests/test_desktop_support.py) | Passed |
| Full unit coverage remains green after the reopened UI pass | `python -m unittest discover -s tests -p "test_*.py" -v` | Passed |

## Caveat

The reopened task is still not closed.

Final closeout still depends on live readback from the real overlay that this latest surface is correct when used directly through `Ctrl+Alt+Space`, `Jobs`, and `Refresh`.

## Verdict

`ready_with_caveats`
