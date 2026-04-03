# Pass 0000 Audit

## Scope

`PASS-0000 Metric Toggle And Aggregation Wiring`

## Result

`ready_with_caveats`

## What Changed

- added a `Total` / `Norm` header toggle to the real overlay surface
- added metric-mode-aware aggregation so bucket totals and repo stacks can switch between raw totals and a normalized cost-weighted proxy
- kept weekly budget projection, headroom, and status on the raw total-token model
- hid the raw `BUDGET LINE` when the chart is in normalized mode to avoid a misleading comparison
- preserved honest right-click investigation context by keeping bucket investigation ranges tied to raw bucket totals rather than normalized display values

## Proof Reviewed

Commands run from repo root:

```powershell
python -m unittest discover -s tests -p "test_*.py" -v
python -m app.codex_dashboard --scan-once --print-summary
python -m app.codex_dashboard --smoke-artifact-dir Tracking/Task-0003/Testing/REGRESSION-RUN-0001
```

Evidence reviewed:

- `app/codex_dashboard/aggregation.py`
- `app/codex_dashboard/ui.py`
- `tests/test_desktop_support.py`
- `tests/test_ingest_core.py`
- `Tracking/Task-0003/Testing/REGRESSION-RUN-0001/desktop-overlay.png`
- `Tracking/Task-0003/Testing/REGRESSION-RUN-0001/overlay-summary.txt`
- `Tracking/Task-0003/Testing/REGRESSION-RUN-0001/overlay-chart.ps`

## Acceptance Check

| Requirement | Evidence | Status |
| --- | --- | --- |
| Header exposes a `Total` / `Norm` toggle | `app/codex_dashboard/ui.py`; `desktop-overlay.png` | Passed |
| Chart buckets switch on selected metric mode | `app/codex_dashboard/aggregation.py`; `tests/test_ingest_core.py` | Passed |
| Repo mode uses the selected metric mode consistently | `app/codex_dashboard/aggregation.py`; `tests/test_ingest_core.py` | Passed |
| Hover values stay aligned with displayed bucket values | `app/codex_dashboard/ui.py`; existing tooltip tests plus metric-mode bucket wiring | Passed |
| Investigation stays honest about raw bucket context | `app/codex_dashboard/ui.py`; `app/codex_dashboard/investigation.py` | Passed |
| Norm mode does not imply a raw budget comparison | `app/codex_dashboard/ui.py`; smoke artifact path | Passed |

## Caveat

- The app-surface regression artifact is captured in default `Total` mode. The `Norm` math and repo-stack behavior are covered by focused unit tests rather than a second app-surface screenshot.

## Verdict

The pass is ready for closeout. The one caveat is non-blocking because the task’s risky behavior is the aggregation math and investigation honesty, and those are covered directly by unit tests plus the real app-surface launch lane.
