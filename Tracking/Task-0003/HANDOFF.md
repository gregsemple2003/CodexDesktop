# Task 0003 Handoff

## Current Status

`Task-0003` is complete and ready for normal use.

## Baseline

The dashboard now supports a top-level metric toggle:

- `Total` keeps the existing raw total-token chart behavior
- `Norm` uses a cost-weighted proxy:
  - uncached input at `1.0x`
  - cached input at `0.1x`
  - output plus reasoning output at `6.0x`

The current product baseline also preserves these guardrails:

- weekly budget, projected burn, headroom, and status remain on the raw total-token model
- the chart hides the raw `BUDGET LINE` when `Norm` is selected
- repo mode uses the selected metric mode for stacked bar totals
- right-click investigations still use raw bucket context so the brief does not lie about “total tokens”

## Validation Status

Supporting scripted proof passed:

- `python -m unittest discover -s tests -p "test_*.py" -v`
- `python -m app.codex_dashboard --scan-once --print-summary`

Repo-root regression lane passed:

- `REG-001 Desktop Overlay Launch And Data Smoke`
- evidence in `Tracking/Task-0003/Testing/REGRESSION-RUN-0001.md`

## Watchouts

- `Norm` is a cost proxy, not a claimed OpenAI billing number.
- The raw weekly budget is still the only budget model in the product, so keep redline semantics tied to `Total`.
- If future work adds a normalized budget line, it needs its own honest source of truth rather than reusing the raw weekly total budget.

## References

- `Tracking/Task-0003/TASK.md`
- `Tracking/Task-0003/PLAN.md`
- `Tracking/Task-0003/Testing/PASS-0000-AUDIT.md`
- `Tracking/Task-0003/Testing/PASS-0000-AUDIT.json`
- `Tracking/Task-0003/Testing/REGRESSION-RUN-0001.md`
- `Tracking/Task-0003/Testing/REGRESSION-RUN-0001/desktop-overlay.png`
- `Tracking/Task-0003/Testing/REGRESSION-RUN-0001/overlay-summary.txt`
- `Tracking/Task-0003/Testing/REGRESSION-RUN-0001/overlay-chart.ps`
