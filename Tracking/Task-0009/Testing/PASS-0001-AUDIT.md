# PASS-0001 Audit

## Scope

PASS-0001 built the first read-only `Tasks` tab surface.

The pass intentionally did not wire destructive or state-changing backend control actions. It renders backend-provided action availability and launch targets as visible affordances, while bounded execution controls remain in later passes.

## Implemented

- Added `app/codex_dashboard/tasks_backend.py` to load and normalize `GET /api/v1/tasks` readback.
- Added `app/codex_dashboard/tasks_tab.py` for task stream grouping, summary card order, state colors, and detail-section semantics.
- Added the `Tasks` nav tab to the Tk dashboard shell.
- Added a committed-work summary strip:
  - `Needs you`
  - `Sleeping`
  - `Running`
  - `Blocked`
  - `Ready`
- Added grouped committed-task rows and a persistent selected-task detail pane.
- Added loading, empty, backend-unavailable, and populated read paths.
- Added isolated fixture loading through `CODEX_DASHBOARD_TASKS_SNAPSHOT_PATH` for validation without touching the human service lane.
- Added `--smoke-tab tasks` support.

## Semantic Guardrails Verified

- Unpromoted candidate rows are filtered out of the committed-work Tasks surface.
- Committed promoted work does not display `Candidate` or `Prov: Candidate`.
- Promoted fixture rows display provenance such as `Promoted from Review` and `Promoted from Dream`.
- Backend `interrupt` action is translated to visible `Pause` copy.
- `Open Live Thread` and working-context launch targets are visible when backend readback provides them.
- The UI does not draw AI-run progress bars.
- PASS-0001 does not kill or interrupt local processes from the UI.

## Validation

All validation used isolated Task-0009 data:

- task-owned fixture: [Fixtures/tasks-snapshot.json](./Fixtures/tasks-snapshot.json)
- task-owned disposable runtime root: `Tracking/Task-0009/Testing/Runtime/PASS-0001` (cleaned after smoke)
- task-owned smoke artifact: [PASS-0001-SMOKE-0001](./PASS-0001-SMOKE-0001)
- backend env vars pointed at the validation lane URL `http://127.0.0.1:14318`
- no `C:\Users\gregs\.codex` live data was used
- no default dashboard config or default dashboard SQLite DB was used

Commands:

```powershell
python -m py_compile app/codex_dashboard/tasks_backend.py app/codex_dashboard/tasks_tab.py app/codex_dashboard/ui.py app/codex_dashboard/__main__.py
python -m unittest discover -s tests -p "test_*.py" -v
python -m app.codex_dashboard --config-path Tracking\Task-0009\Testing\Runtime\PASS-0001\config.json --db-path Tracking\Task-0009\Testing\Runtime\PASS-0001\dashboard-smoke.db --codex-root Tracking\Task-0009\Testing\Runtime\PASS-0001\codex-fixture --smoke-tab tasks --smoke-artifact-dir Tracking\Task-0009\Testing\PASS-0001-SMOKE-0001
```

Results:

- compile check passed
- unit discovery passed: `82` tests
- desktop smoke passed and captured [PASS-0001-SMOKE-0001/overlay.png](./PASS-0001-SMOKE-0001/overlay.png)
- smoke summary recorded:
  - `active_tab=tasks`
  - `tasks_needs_you=01`
  - `tasks_ready=02`
  - `tasks_selected=Task-0009`

## Residual Risk

This pass proves the read-only surface and state mapping. The following remain for later approved passes:

- live backend dispatch and run-control POST actions
- actual editor/VSCodium launch behavior for deep-context targets
- backend-driven resume/continue behavior
- full final regression case closure after action wiring and polish
