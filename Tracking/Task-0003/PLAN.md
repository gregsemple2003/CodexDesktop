# Task 0003 Plan

## Planning Verdict

This plan is approved for execution under the standing user instruction to auto-approve plan, pass, and phase gates that do not expand scope, spend money, delete data, or violate workflow rules.

## Pass Order

### PASS-0000 Metric Toggle And Aggregation Wiring

Objective:

- add the `Total` / `Norm` toggle and wire chart aggregation, repo stacks, and hover values to the selected metric mode

Expected work:

- add a normalized token helper and metric-mode-aware bucket builders
- add header controls for metric mode selection
- update chart rendering so normalized mode does not imply the raw budget line
- preserve honest raw context for the investigation path
- add or update focused unit coverage

Verification:

- run `python -m unittest discover -s tests -p "test_*.py" -v`
- run `python -m app.codex_dashboard --scan-once --print-summary`

Exit bar:

- the overlay supports both metric modes
- chart buckets, repo stacks, and hover values respect the selected mode
- investigations still use raw bucket context
- pass-local proof passes

### PASS-0001 Regression And Closure

Objective:

- execute the repo-root desktop overlay regression lane and close the task if the feature holds up in the real app surface

Expected work:

- run repo-root `REG-001 Desktop Overlay Launch And Data Smoke`
- write the task-owned regression artifact
- finalize handoff and task state

Verification:

- run the repo-root regression lane from `REGRESSION.md`

Exit bar:

- the task-owned regression artifact names the actual lane run
- `HANDOFF.md` and `TASK-STATE.json` reflect the honest final state
