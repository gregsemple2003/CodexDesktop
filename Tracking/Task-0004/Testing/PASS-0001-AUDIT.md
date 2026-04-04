# Pass 0001 Audit

## Scope

`PASS-0001` added the first honest `Jobs` lane to the existing overlay while keeping `Usage` as the default surface.

Implemented in this pass:

- added additive overlay tab state for:
  - `Usage`
  - `Jobs`
  - inert `Logs` and `Terminal` placeholders
- added a `Jobs` summary area with:
  - declared jobs
  - in-sync jobs
  - needs-attention count
  - last reconciliation time
- added per-job rows showing:
  - label
  - mechanism
  - desired versus observed state
  - drift status
  - human-facing reason text
- wired bounded actions to the backend jobs model for:
  - refresh state
  - reconcile supported drift
- extended smoke support so the real startup path can capture the `Jobs` lane explicitly
- added focused unit coverage for jobs-lane helper behavior and registry mutation support

## Verification

Executed from repo root:

```powershell
python -m unittest discover -s tests -p "test_*.py" -v
python -m app.codex_dashboard --smoke-artifact-dir Tracking/Task-0004/Testing/PASS-0001-UI-SMOKE-0002 --smoke-tab jobs
```

Observed result:

- full unit coverage passed: `43` tests
- the real app startup path rendered the overlay and captured the `Jobs` lane
- smoke summary confirms:
  - `active_tab=jobs`
  - refresh copy rendered from the live jobs backend
- smoke artifacts were written under:
  - `Tracking/Task-0004/Testing/PASS-0001-UI-SMOKE-0002/overlay-summary.txt`
  - `Tracking/Task-0004/Testing/PASS-0001-UI-SMOKE-0002/overlay-chart.ps`

## Requirement Mapping

| Requirement | Evidence | Result |
| --- | --- | --- |
| Keep `Usage` as the default tab | `app/codex_dashboard/ui.py`; `DashboardApp.active_tab = "usage"` | Passed |
| Add a visible `Jobs` lane without replacing the existing overlay | `app/codex_dashboard/ui.py` | Passed |
| Show job summary counts and last reconciliation time | `app/codex_dashboard/ui.py`; `format_jobs_timestamp`; smoke summary | Passed |
| Show per-job status and human-facing reason text | `app/codex_dashboard/ui.py`; backend labels from `app/codex_dashboard/jobs.py` | Passed |
| Wire bounded `Refresh` and `Reconcile` actions to the backend | `app/codex_dashboard/ui.py`; `reconcile_registry`; `apply_registry` | Passed |
| Keep `Logs` and `Terminal` out of scope as inert placeholders | `app/codex_dashboard/ui.py` | Passed |
| Add focused validation for jobs-lane behavior | `tests/test_desktop_support.py`; `tests/test_jobs.py` | Passed |

## Notes

- this pass intentionally stops short of task-level regression closure
- repo-canonical regression `REG-001` still belongs to `PASS-0002`
- the smoke capture is supporting UI proof, not final regression closure

## Verdict

`ready`
