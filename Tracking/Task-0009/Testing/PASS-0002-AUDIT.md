# PASS-0002 Audit

## Scope

PASS-0002 wired the `Tasks` tab to bounded backend action contracts and deep-context launch targets.

Task-0008 still owns backend run-control semantics. Task-0009 now consumes those contracts through the humane dashboard surface.

## Implemented

- Added backend client calls for:
  - `POST /api/v1/tasks/{task_id}/dispatch`
  - `POST /api/v1/task-runs/{run_id}/poke`
  - `POST /api/v1/task-runs/{run_id}/interrupt`
  - `POST /api/v1/task-runs/{run_id}/retry-workload`
- Kept visible stop or hold copy as `Pause` while calling the backend `interrupt` endpoint.
- Added action metadata to mapped task rows:
  - `task_id`
  - `latest_run_id`
  - action label
  - backend action name
  - backend run id when needed
  - launch target when available
- Added launch handling for backend deep-context targets:
  - `Open Live Thread`
  - `Open Thread`
  - `Open Working Context`
  - `Open Task`
  - `Open Transcript`
- Bounded command-based launching to editor-style commands such as `code` and `codium`; otherwise the UI falls back to opening the target URI/path.
- Added smoke summary evidence for selected-task actions.

## Semantic Guardrails Verified

- The dashboard does not kill arbitrary local processes.
- `Pause` is a backend-backed action, not a UI-side process kill.
- `Open Live Thread` remains the preferred instruction path when the backend exposes it.
- `Resume` or `Continue` remains backend-gated; this pass does not invent a new unrelated run when backend readback does not expose a safe continue contract.
- Unpromoted candidates remain filtered from the `Tasks` tab.
- No AI-run progress bars were introduced.

## Validation

All validation used isolated Task-0009 data:

- task-owned fixture: [Fixtures/tasks-snapshot.json](./Fixtures/tasks-snapshot.json)
- task-owned smoke artifact: [PASS-0002-SMOKE-0001](./PASS-0002-SMOKE-0001)
- disposable runtime root: `Tracking/Task-0009/Testing/Runtime/PASS-0002` (cleaned after smoke)
- backend env vars pointed at validation-lane URL `http://127.0.0.1:14318`
- no `C:\Users\gregs\.codex` live data was used
- no human dashboard config or active dashboard SQLite DB was used

Commands:

```powershell
python -m py_compile app/codex_dashboard/tasks_backend.py app/codex_dashboard/ui.py
python -m unittest discover -s tests -p "test_*.py" -v
python -m app.codex_dashboard --config-path Tracking\Task-0009\Testing\Runtime\PASS-0002\config.json --db-path Tracking\Task-0009\Testing\Runtime\PASS-0002\dashboard-smoke.db --codex-root Tracking\Task-0009\Testing\Runtime\PASS-0002\codex-fixture --smoke-tab tasks --smoke-artifact-dir Tracking\Task-0009\Testing\PASS-0002-SMOKE-0001
```

Results:

- compile check passed
- unit discovery passed: `85` tests
- desktop smoke passed and captured [PASS-0002-SMOKE-0001/overlay.png](./PASS-0002-SMOKE-0001/overlay.png)
- smoke summary recorded:
  - `active_tab=tasks`
  - `tasks_needs_you=01`
  - `tasks_ready=02`
  - `tasks_selected=Task-0009`
  - `tasks_selected_actions=Dispatch,Pause,Open Live Thread,Open Working Context,Open Task`

## Residual Risk

This pass validates UI action wiring through focused unit tests and fixture-backed smoke. Final closure still needs regression through the real app surface under the isolated lane rules from [REGRESSION.md](../../REGRESSION.md).

Promotion-specific provenance from Task-0010 remains a later integration pass.
