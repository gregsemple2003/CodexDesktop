# Pass 0002 Audit

## Scope

`PASS-0002` remains reopened after the live `Jobs` regression reported in [BUG-0001](/c:/Agent/CodexDashboard/Tracking/Task-0004/BUG-0001.md).

This latest reopened slice fixes the remaining jobs-surface usability issue:

- the `Jobs` content area is now scrollable
- declared-job details and the jobs list now live inside that scroll container
- mousewheel scrolling is wired for the real `Jobs` surface

Previously preserved in the same reopened pass:

- declared jobs are the primary durable state under `.codex/Orchestration/Jobs`
- `Jobs` shows declared jobs first
- `Details` shows the declared job payload
- `Refresh` remains the explicit actual-versus-declared diff step

## Verification

Executed from repo root:

```powershell
python -m unittest discover -s tests -p "test_*.py" -v
python -m app.codex_dashboard --scan-once --print-summary
python -m app.codex_dashboard --smoke-artifact-dir Tracking/Task-0004/Testing/PASS-0002-JOBS-SMOKE-0009 --smoke-tab jobs
```

Observed result:

- full unit coverage passed: `52` tests
- the supporting `--scan-once --print-summary` path still completed successfully against the live telemetry tree
- refreshed jobs runtime evidence was captured in:
  - [desktop-overlay.png](/c:/Agent/CodexDashboard/Tracking/Task-0004/Testing/PASS-0002-JOBS-SMOKE-0009/desktop-overlay.png)
  - [overlay-summary.txt](/c:/Agent/CodexDashboard/Tracking/Task-0004/Testing/PASS-0002-JOBS-SMOKE-0009/overlay-summary.txt)

## Requirement Mapping

| Requirement | Evidence | Result |
| --- | --- | --- |
| `Jobs` surface is scrollable on the real overlay | [ui.py](/c:/Agent/CodexDashboard/app/codex_dashboard/ui.py); [desktop-overlay.png](/c:/Agent/CodexDashboard/Tracking/Task-0004/Testing/PASS-0002-JOBS-SMOKE-0009/desktop-overlay.png) | Passed |
| Mousewheel input scrolls the jobs surface | [ui.py](/c:/Agent/CodexDashboard/app/codex_dashboard/ui.py); [test_desktop_support.py](/c:/Agent/CodexDashboard/tests/test_desktop_support.py) | Passed |
| Declared-job details and rows share the scrollable surface | [ui.py](/c:/Agent/CodexDashboard/app/codex_dashboard/ui.py) | Passed |
| Full unit coverage remains green after the scrolling pass | `python -m unittest discover -s tests -p "test_*.py" -v` | Passed |

## Caveat

The reopened task is still not closed.

Final closeout still depends on live readback from the actual overlay that the `Jobs` panel now scrolls correctly and that the overall declared-jobs workflow feels right when used directly.

## Verdict

`ready_with_caveats`
