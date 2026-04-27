# REGRESSION-RUN-0001

## Case

Repo-root [REG-003 Tasks Tab Committed-Work Surface](../../../REGRESSION.md#reg-003-tasks-tab-committed-work-surface).

## Lane

Isolated Task-0009 regression lane.

This run did not use the human's personal dashboard lane, human dashboard config, active dashboard database, or live Codex data.

Inputs:

- task-owned telemetry fixture: [Fixtures/codex-root/sessions/2026/04/26/rollout-task-0009-regression.jsonl](./Fixtures/codex-root/sessions/2026/04/26/rollout-task-0009-regression.jsonl)
- task-owned tasks fixture: [Fixtures/tasks-snapshot.json](./Fixtures/tasks-snapshot.json)
- task readback fixture env: `CODEX_DASHBOARD_TASKS_SNAPSHOT_PATH`
- backend URL envs: `http://127.0.0.1:14318`
- disposable config and SQLite runtime under `Tracking/Task-0009/Testing/Runtime/PASS-0004-REGRESSION`

The disposable runtime was removed after the run.

## Command

```powershell
python -m app.codex_dashboard --config-path Tracking\Task-0009\Testing\Runtime\PASS-0004-REGRESSION\config.json --db-path Tracking\Task-0009\Testing\Runtime\PASS-0004-REGRESSION\dashboard-regression.db --codex-root Tracking\Task-0009\Testing\Fixtures\codex-root --smoke-tab tasks --smoke-artifact-dir Tracking\Task-0009\Testing\PASS-0004-REGRESSION-0001
```

## Result

Passed.

Evidence:

- [PASS-0004-REGRESSION-0001/overlay.png](./PASS-0004-REGRESSION-0001/overlay.png)
- [PASS-0004-REGRESSION-0001/overlay-summary.txt](./PASS-0004-REGRESSION-0001/overlay-summary.txt)

Key summary lines:

```text
active_tab=tasks
Last ingest 21:57:03 | files 1/1 | events +1
tasks_needs_you=01
tasks_ready=02
tasks_selected=Task-0009
tasks_selected_actions=Dispatch,Pause,Open Live Thread,Open Working Context,Open Task
tasks_provenance_labels=Promoted from Review,Authored,Promoted from Dream
```

The captured surface shows the `Tasks` tab selected, committed-work summary cards, grouped task rows, selected-task detail, promoted provenance, visible `Pause` copy, and no AI-run progress bars.
