# Task 0009 Handoff

## Current Status

`Task-0009` is complete.

Current lifecycle state:

- phase: `closure`
- current pass: `null`
- last completed pass: `PASS-0004`
- current gate: `closure`

This task owns the human-facing `Tasks` tab for committed work in CodexDashboard:

- the high-level view
- committed-task triage
- dispatch intent
- monitoring
- recovery affordances
- deep-context launch

This task does not own intake review, the promotion contract, the dispatch runtime, or the daily Dream automation.

Those are split out into:

- [Task-0011](../Task-0011/TASK.md)
- [Task-0008](../Task-0008/TASK.md)
- [Task-0010](../Task-0010/TASK.md)

If the `Tasks` tab later shows work that was promoted out of `Review`, it should consume durable provenance from the single promotion contract owned by [Task-0010](../Task-0010/TASK.md). This task does not define candidate review or promotion semantics.

## Current Objective

No further implementation remains in Task-0009.

The design work for that lives in:

- [Design/HUMAN-NEED-AND-TASKS-TAB-DIRECTION.md](./Design/HUMAN-NEED-AND-TASKS-TAB-DIRECTION.md)
- [Design/STITCH-PROMPT.md](./Design/STITCH-PROMPT.md)
- [Mockup/stitch_task_tab/screen.png](./Mockup/stitch_task_tab/screen.png)
- [Mockup/stitch_task_tab/code.html](./Mockup/stitch_task_tab/code.html)

The review-surface design and research now live under [Task-0011](../Task-0011/HANDOFF.md).

The product split is now explicit:

- `Review` for incoming asks
- `Tasks` for committed work

## Binding Implementation Constraints

Validation and regression for this task must not touch the human's live lane.

Use an isolated Task-0009 validation/regression lane with agent-owned data and ports. The human's dashboard service lane, service ports, and app data are off-limits unless the human gives explicit later instructions.

If isolated validation or regression cannot run, record the blocker in task-owned testing artifacts instead of falling back to the human's lane.

Implementation must stay aligned with the active generated mockup and style references while overriding the known semantic drift in [PLAN.md](./PLAN.md):

- no unpromoted candidates on `Tasks`
- no `Candidate` / `Prov: Candidate` provenance for committed work
- promoted work uses durable source provenance such as `Promoted from Dream` or `Promoted from Review`
- user-facing stop or hold action is `Pause`, even if backend internals use `interrupt`
- `Open Live Thread` is the preferred instruction path when available
- no AI task-run progress bars unless a future backend contract proves trustworthy bounded progress

## Completed Passes

### PASS-0000 Lock The Product Surface

Completed in this checkpoint.

Primary output:

- [Design/PASS-0000-PRODUCT-SURFACE-CONTRACT.md](./Design/PASS-0000-PRODUCT-SURFACE-CONTRACT.md)

Audit:

- [Testing/PASS-0000-AUDIT.md](./Testing/PASS-0000-AUDIT.md)

Result:

- the first-release information architecture is explicit
- the active mockup is accepted as a style reference, with candidate/progress-bar drift overridden
- action copy and instruction flow are frozen for implementation
- validation and regression lane isolation is explicit

### PASS-0001 Build The Read-Only Task Surface

Completed in this checkpoint.

Primary outputs:

- [app/codex_dashboard/tasks_backend.py](../../app/codex_dashboard/tasks_backend.py)
- [app/codex_dashboard/tasks_tab.py](../../app/codex_dashboard/tasks_tab.py)
- `Tasks` tab integration in [app/codex_dashboard/ui.py](../../app/codex_dashboard/ui.py)
- `--smoke-tab tasks` support in [app/codex_dashboard/__main__.py](../../app/codex_dashboard/__main__.py)
- task-owned fixture data at [Testing/Fixtures/tasks-snapshot.json](./Testing/Fixtures/tasks-snapshot.json)

Audit and proof:

- [Testing/PASS-0001-AUDIT.md](./Testing/PASS-0001-AUDIT.md)
- [Testing/PASS-0001-SMOKE-0001/overlay.png](./Testing/PASS-0001-SMOKE-0001/overlay.png)
- [Testing/PASS-0001-SMOKE-0001/overlay-summary.txt](./Testing/PASS-0001-SMOKE-0001/overlay-summary.txt)

Result:

- the dashboard has a real `Tasks` tab in top navigation
- the tab renders committed-task summary counts, grouped rows, and a persistent detail pane
- unpromoted candidates are filtered out
- promoted committed work uses promoted provenance rather than candidate labels
- backend `interrupt` availability is rendered as visible `Pause`
- no AI-run progress bars are used
- the desktop smoke used isolated Task-0009 config, SQLite, fixture data, and validation-lane backend URLs

### PASS-0002 Add Dispatch And Recovery Actions

Completed in this checkpoint.

Primary outputs:

- backend action clients in [app/codex_dashboard/tasks_backend.py](../../app/codex_dashboard/tasks_backend.py)
- action and launch handling in [app/codex_dashboard/ui.py](../../app/codex_dashboard/ui.py)
- action-aware fixture data in [Testing/Fixtures/tasks-snapshot.json](./Testing/Fixtures/tasks-snapshot.json)

Audit and proof:

- [Testing/PASS-0002-AUDIT.md](./Testing/PASS-0002-AUDIT.md)
- [Testing/PASS-0002-SMOKE-0001/overlay.png](./Testing/PASS-0002-SMOKE-0001/overlay.png)
- [Testing/PASS-0002-SMOKE-0001/overlay-summary.txt](./Testing/PASS-0002-SMOKE-0001/overlay-summary.txt)

Result:

- `Dispatch`, `Poke`, visible `Pause`, and workload `Continue` paths call bounded backend endpoints when backend readback exposes those actions
- visible `Pause` calls backend `interrupt`; the UI does not kill arbitrary processes
- deep-context launch buttons use backend-provided targets such as `Open Live Thread`, `Open Working Context`, and `Open Task`
- command launches are bounded to editor-style commands such as `code` and `codium`
- the smoke summary records selected actions for the fixture-backed selected task
- the desktop smoke again used isolated Task-0009 config, SQLite, fixture data, and validation-lane backend URLs

### PASS-0003 Integrate Reviewed And Promoted Work Provenance

Completed in this checkpoint.

Primary outputs:

- promotion provenance parsing in [app/codex_dashboard/tasks_backend.py](../../app/codex_dashboard/tasks_backend.py)
- source provenance detail rendering through [app/codex_dashboard/tasks_tab.py](../../app/codex_dashboard/tasks_tab.py)
- action/provenance smoke summary additions in [app/codex_dashboard/ui.py](../../app/codex_dashboard/ui.py)
- action-aware and provenance-aware fixture data in [Testing/Fixtures/tasks-snapshot.json](./Testing/Fixtures/tasks-snapshot.json)

Audit and proof:

- [Testing/PASS-0003-AUDIT.md](./Testing/PASS-0003-AUDIT.md)
- [Testing/PASS-0003-SMOKE-0001/overlay.png](./Testing/PASS-0003-SMOKE-0001/overlay.png)
- [Testing/PASS-0003-SMOKE-0001/overlay-summary.txt](./Testing/PASS-0003-SMOKE-0001/overlay-summary.txt)

Result:

- unpromoted review candidates remain filtered from `Tasks`
- promoted Dream and Review work gets explicit committed-task provenance labels
- source packet/problem/winner/option-task details can appear in the selected-task detail pane
- Task-0009 consumes likely Task-0010 provenance shapes without defining a second promotion mechanism

### PASS-0004 Polish, Audit, And Regression

Completed in this checkpoint.

Primary outputs:

- final scrollbar polish in [app/codex_dashboard/ui.py](../../app/codex_dashboard/ui.py)
- task-owned regression telemetry fixture at [Testing/Fixtures/codex-root/sessions/2026/04/26/rollout-task-0009-regression.jsonl](./Testing/Fixtures/codex-root/sessions/2026/04/26/rollout-task-0009-regression.jsonl)

Audit and proof:

- [Testing/PASS-0004-AUDIT.md](./Testing/PASS-0004-AUDIT.md)
- [Testing/REGRESSION-RUN-0001.md](./Testing/REGRESSION-RUN-0001.md)
- [Testing/PASS-0004-REGRESSION-0001/overlay.png](./Testing/PASS-0004-REGRESSION-0001/overlay.png)
- [Testing/PASS-0004-REGRESSION-0001/overlay-summary.txt](./Testing/PASS-0004-REGRESSION-0001/overlay-summary.txt)

Result:

- compile validation passed
- unit discovery passed with `88` tests
- repo-root REG-003 regression passed using isolated Task-0009 data and runtime
- disposable runtime was cleaned after regression

## Current Baseline

The repo currently has:

- the token-velocity overlay
- the backend-backed `Jobs` tab
- task-owned markdown artifacts under `Tracking/`

The first action-wired, provenance-aware humane surface now exists.

Follow-on product work remains intentionally split:

- Task-0010 owns the final live promotion endpoint and promotion writer.
- Task-0011 owns the provisional intake `Review` tab.
- Task-0008 owns backend runtime semantics.

## Next Recommended Step

No Task-0009 implementation step remains. Future work should continue in the adjacent owning tasks rather than reopening this task's scope.

## Watchouts

- do not let the `Tasks` tab become a raw orchestration inspector
- do not hide task provenance
- do not let intake-review material leak back into `Tasks`; that belongs in [Task-0011](../Task-0011/TASK.md)
- do not make the human read markdown to understand the default screen
- do not use the human's dashboard lane, ports, or app data for validation/regression
- do not copy generated mockup candidate labels or AI progress bars into implementation

## References

- [TASK.md](./TASK.md)
- [PLAN.md](./PLAN.md)
- [Task-0005](../Task-0005/TASK.md)
- [Task-0007](../Task-0007/TASK.md)
- [Task-0008](../Task-0008/TASK.md)
- [Task-0010](../Task-0010/TASK.md)
- [Task-0011](../Task-0011/TASK.md)
- [Design/GENERAL-DESIGN.md](../../Design/GENERAL-DESIGN.md)
