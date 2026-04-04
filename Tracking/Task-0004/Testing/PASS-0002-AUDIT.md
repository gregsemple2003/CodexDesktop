# Pass 0002 Audit

## Scope

`PASS-0002` remains reopened after the live `Jobs` regression reported in [BUG-0001](/c:/Agent/CodexDashboard/Tracking/Task-0004/BUG-0001.md).

This latest reopened slice applies the clarified jobs contract:

- declared jobs are the primary durable state
- declared jobs now live under `.codex/Orchestration/Jobs/declared-jobs.json`
- a sibling schema now lives at `.codex/Orchestration/Jobs/declared-jobs.schema.json`
- clicking `Jobs` shows declared jobs first without triggering refresh
- `Refresh` computes the actual-versus-declared diff
- `Details` shows the declared job payload rather than observed Windows state

## Verification

Executed from repo root:

```powershell
python -m unittest discover -s tests -p "test_*.py" -v
python -m app.codex_dashboard --scan-once --print-summary
python -m app.codex_dashboard --smoke-artifact-dir Tracking/Task-0004/Testing/PASS-0002-REG-001-0007
python -m app.codex_dashboard --smoke-artifact-dir Tracking/Task-0004/Testing/PASS-0002-JOBS-SMOKE-0008 --smoke-tab jobs
```

Observed result:

- full unit coverage passed: `51` tests
- the supporting `--scan-once --print-summary` path still completed successfully against the live telemetry tree
- the live `.codex` tree now contains:
  - `C:\Users\gregs\.codex\Orchestration\Jobs\declared-jobs.json`
  - `C:\Users\gregs\.codex\Orchestration\Jobs\declared-jobs.schema.json`
- refreshed `Usage` runtime evidence was captured in:
  - [desktop-overlay.png](/c:/Agent/CodexDashboard/Tracking/Task-0004/Testing/PASS-0002-REG-001-0007/desktop-overlay.png)
  - [overlay-summary.txt](/c:/Agent/CodexDashboard/Tracking/Task-0004/Testing/PASS-0002-REG-001-0007/overlay-summary.txt)
  - [overlay-chart.ps](/c:/Agent/CodexDashboard/Tracking/Task-0004/Testing/PASS-0002-REG-001-0007/overlay-chart.ps)
- refreshed `Jobs` runtime evidence was captured in:
  - [desktop-overlay.png](/c:/Agent/CodexDashboard/Tracking/Task-0004/Testing/PASS-0002-JOBS-SMOKE-0008/desktop-overlay.png)
  - [overlay-summary.txt](/c:/Agent/CodexDashboard/Tracking/Task-0004/Testing/PASS-0002-JOBS-SMOKE-0008/overlay-summary.txt)

## Requirement Mapping

| Requirement | Evidence | Result |
| --- | --- | --- |
| Declared jobs live under `.codex/Orchestration/Jobs` | [paths.py](/c:/Agent/CodexDashboard/app/codex_dashboard/paths.py); live `.codex` tree | Passed |
| Declared jobs have a sibling schema file | [jobs.py](/c:/Agent/CodexDashboard/app/codex_dashboard/jobs.py); live `.codex` tree | Passed |
| `Jobs` tab shows declared jobs before refresh | [ui.py](/c:/Agent/CodexDashboard/app/codex_dashboard/ui.py); [test_desktop_support.py](/c:/Agent/CodexDashboard/tests/test_desktop_support.py) | Passed |
| `Refresh` remains the explicit diff step | [ui.py](/c:/Agent/CodexDashboard/app/codex_dashboard/ui.py); [REGRESSION.md](/c:/Agent/CodexDashboard/REGRESSION.md) | Passed |
| `Details` shows the declared job payload | [ui.py](/c:/Agent/CodexDashboard/app/codex_dashboard/ui.py); [test_desktop_support.py](/c:/Agent/CodexDashboard/tests/test_desktop_support.py) | Passed |
| Full unit coverage remains green after the clarified jobs pass | `python -m unittest discover -s tests -p "test_*.py" -v` | Passed |

## Caveat

The reopened task is still not closed.

Final closeout still depends on live readback from the actual overlay that this clarified declared-jobs contract matches the intended UX when used directly through `Ctrl+Alt+Space`, `Jobs`, `Details`, and `Refresh`.

## Verdict

`ready_with_caveats`
