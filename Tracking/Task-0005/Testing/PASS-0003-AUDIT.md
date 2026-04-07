# Pass 0003 Audit

## Scope

`PASS-0003` turned the Tk `Jobs` surface into a backend client instead of a local Windows reconciler:

- the `Jobs` tab now reads backend state from the local orchestration service
- the row model now shows backend trigger labels, desired-vs-runtime state, recent-run context, and backend details
- the UI exposes bounded backend actions for `Sync now` and `Run now`
- smoke capture now records real overlay PNG artifacts for both `Usage` and `Jobs`
- repo-root regression coverage was updated for the backend-backed `Jobs` lane

This pass does not change the backend runtime contract itself. It consumes the backend delivered in earlier passes.

## Verification

Executed from repo root unless noted otherwise:

```powershell
& 'C:\Program Files\Go\bin\go.exe' test ./...
python -m unittest discover -s tests -p "test_*.py" -v
python -m app.codex_dashboard --scan-once --print-summary
Invoke-WebRequest http://127.0.0.1:4318/healthz | Select-Object -ExpandProperty Content
python -m app.codex_dashboard --smoke-artifact-dir 'C:\Agent\CodexDashboard\Tracking\Task-0005\Testing\REGRESSION-RUN-0001-usage-smoke' --smoke-tab usage
python -m app.codex_dashboard --smoke-artifact-dir 'C:\Agent\CodexDashboard\Tracking\Task-0005\Testing\REGRESSION-RUN-0001-jobs-smoke' --smoke-tab jobs
```

Observed result:

- Go backend tests still passed after the UI integration checkpoint
- repo-root Python unit coverage passed: `62` tests
- the supporting ingest smoke still completed successfully against the live telemetry tree
- `/healthz` still reported `status: ok` against the live backend
- the `Usage` smoke produced:
  - [overlay.png](/c:/Agent/CodexDashboard/Tracking/Task-0005/Testing/REGRESSION-RUN-0001-usage-smoke/overlay.png)
  - [overlay-summary.txt](/c:/Agent/CodexDashboard/Tracking/Task-0005/Testing/REGRESSION-RUN-0001-usage-smoke/overlay-summary.txt)
  - [overlay-chart.ps](/c:/Agent/CodexDashboard/Tracking/Task-0005/Testing/REGRESSION-RUN-0001-usage-smoke/overlay-chart.ps)
- the `Jobs` smoke produced:
  - [overlay.png](/c:/Agent/CodexDashboard/Tracking/Task-0005/Testing/REGRESSION-RUN-0001-jobs-smoke/overlay.png)
  - [overlay-summary.txt](/c:/Agent/CodexDashboard/Tracking/Task-0005/Testing/REGRESSION-RUN-0001-jobs-smoke/overlay-summary.txt)
- the `Jobs` overlay artifact showed:
  - `03` declared jobs
  - `03` in sync
  - `00` needs attention
  - `Sync now` and `Refresh` controls
  - visible `Run now` buttons only for manual-capable jobs, with the webhook-only job disabled
- the `Jobs` smoke status line recorded `Jobs sync completed. 0 schedule changes.`, proving the real app surface invoked the backend sync action during the capture

## Requirement Mapping

| Requirement | Evidence | Result |
| --- | --- | --- |
| Dashboard reads backend jobs state instead of local Windows reconciliation | [jobs_backend.py](/c:/Agent/CodexDashboard/app/codex_dashboard/jobs_backend.py); [ui.py](/c:/Agent/CodexDashboard/app/codex_dashboard/ui.py); [overlay.png](/c:/Agent/CodexDashboard/Tracking/Task-0005/Testing/REGRESSION-RUN-0001-jobs-smoke/overlay.png) | Passed |
| Jobs surface remains bounded and operator-clear | [overlay.png](/c:/Agent/CodexDashboard/Tracking/Task-0005/Testing/REGRESSION-RUN-0001-jobs-smoke/overlay.png); [overlay-summary.txt](/c:/Agent/CodexDashboard/Tracking/Task-0005/Testing/REGRESSION-RUN-0001-jobs-smoke/overlay-summary.txt) | Passed |
| Bounded backend control actions are exposed through the dashboard | `Sync now` widget invoked during jobs smoke; `Run now` button wiring covered by [test_desktop_support.py](/c:/Agent/CodexDashboard/tests/test_desktop_support.py) and visible in [overlay.png](/c:/Agent/CodexDashboard/Tracking/Task-0005/Testing/REGRESSION-RUN-0001-jobs-smoke/overlay.png) | Passed |
| Repo-root regression requirements are updated honestly for the backend-backed Jobs lane | [REGRESSION.md](/c:/Agent/CodexDashboard/REGRESSION.md); [REGRESSION-RUN-0001.md](/c:/Agent/CodexDashboard/Tracking/Task-0005/Testing/REGRESSION-RUN-0001.md) | Passed |
| Focused automated proof exists for backend snapshot mapping and UI control wiring | [test_jobs_backend.py](/c:/Agent/CodexDashboard/tests/test_jobs_backend.py); [test_desktop_support.py](/c:/Agent/CodexDashboard/tests/test_desktop_support.py); `python -m unittest discover -s tests -p "test_*.py" -v` | Passed |

## Caveat

The app-surface regression used the real widget command path in smoke automation rather than a physical mouse click.

That is still honest regression evidence here because the real Tk app and overlay were running, the live `Jobs` widget tree was rendered, the real `Sync now` button command was invoked through the widget, and the resulting live window was captured to PNG. The `Run now` button was not executed during `PASS-0003` regression because `PASS-0002` had already proven the paid/email-backed manual run path end to end and repeating it from the UI would have added unnecessary side effects.

## Verdict

`ready`
