# PASS-0003 Audit

## Scope

PASS-0003 tightened the committed-work provenance contract for the `Tasks` tab.

Task-0010 still owns the final backend promotion mechanism. Task-0009 now accepts backend-shaped promotion provenance without treating unpromoted candidates as committed work.

## Implemented

- Added promotion provenance parsing for backend fields such as:
  - `promotion_provenance`
  - `promotion`
  - `provenance`
- Added support for source fields such as:
  - `source`
  - `source_type`
  - `source_surface`
  - `promoted_from`
- Added durable source detail rendering for:
  - source packet
  - source problem
  - source winner
  - source option task
  - promoter and promotion timestamp when present
- Added promotion evidence refs into task artifacts.
- Tightened unpromoted candidate filtering across:
  - `kind`
  - `status`
  - `lifecycle_state`
  - `promotion_status`
  - `review_status`
  - provenance object status fields
- Added smoke summary provenance labels for all rendered committed tasks.

## Semantic Guardrails Verified

- Unpromoted review candidates stay off the committed-work `Tasks` surface.
- Promoted Dream work displays as `Promoted from Dream`.
- Promoted Review work displays as `Promoted from Review`.
- A committed task never displays `Candidate` or `Prov: Candidate` as its provenance label.
- Candidate/intake work remains owned by Task-0011 until promotion creates a committed task.
- Task-0009 does not define a second promotion mechanism.

## Validation

All validation used isolated Task-0009 data:

- task-owned fixture: [Fixtures/tasks-snapshot.json](./Fixtures/tasks-snapshot.json)
- task-owned smoke artifact: [PASS-0003-SMOKE-0001](./PASS-0003-SMOKE-0001)
- disposable runtime root: `Tracking/Task-0009/Testing/Runtime/PASS-0003` (cleaned after smoke)
- backend env vars pointed at validation-lane URL `http://127.0.0.1:14318`
- no `C:\Users\gregs\.codex` live data was used
- no human dashboard config or active dashboard SQLite DB was used

Commands:

```powershell
python -m py_compile app/codex_dashboard/ui.py app/codex_dashboard/tasks_backend.py app/codex_dashboard/tasks_tab.py
python -m unittest discover -s tests -p "test_*.py" -v
python -m app.codex_dashboard --config-path Tracking\Task-0009\Testing\Runtime\PASS-0003\config.json --db-path Tracking\Task-0009\Testing\Runtime\PASS-0003\dashboard-smoke.db --codex-root Tracking\Task-0009\Testing\Runtime\PASS-0003\codex-fixture --smoke-tab tasks --smoke-artifact-dir Tracking\Task-0009\Testing\PASS-0003-SMOKE-0001
```

Results:

- compile check passed
- unit discovery passed: `88` tests
- desktop smoke passed and captured [PASS-0003-SMOKE-0001/overlay.png](./PASS-0003-SMOKE-0001/overlay.png)
- smoke summary recorded:
  - `active_tab=tasks`
  - `tasks_needs_you=01`
  - `tasks_ready=02`
  - `tasks_provenance_labels=Promoted from Review,Authored,Promoted from Dream`

## Residual Risk

This pass supports likely Task-0010 provenance shapes but cannot prove the final live promotion endpoint until Task-0010 ships that contract.

Final closure still needs polish and repo-root regression under the isolated lane rules from [REGRESSION.md](../../REGRESSION.md).
