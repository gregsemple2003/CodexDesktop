# Task 0002 Research Plan

## Research Goal

Turn the Stitch overlay mockup into an implementable Tk overlay plan without losing the product-intent requirements from `Design/GENERAL-DESIGN.md`.

## Decision-Shaping Problems

### Problem 0001 Overlay Structure Mapping

Determine how the Stitch composition maps onto the current Tk overlay and where the repo-root general design intentionally overrides the mockup.

### Problem 0002 Tk Implementation Fit

Determine which parts of the current `ui.py` surface can stay structurally intact and which parts should move from `ttk` styling toward more explicit `tk` composition for tighter visual control.

### Problem 0003 Proof Strategy

Determine what pass-local proof is realistic for a Tk redesign before the repo-root desktop regression lane runs.

## Research Inputs

- `Tracking/Task-0002/TASK.md`
- `Tracking/Task-0002/HANDOFF.md`
- `Tracking/Task-0001/HANDOFF.md`
- `Design/GENERAL-DESIGN.md`
- `Design/Mockups/stitch/DESIGN.md`
- `Design/Mockups/stitch/code.html`
- `Design/Mockups/stitch/screen.png`
- `app/codex_dashboard/ui.py`
- repo-root `REGRESSION.md`
- repo-root `TESTING.md`

## Intended Outputs

- `RESEARCH-ANALYSIS.md`
- `RESEARCH.md`
- a planning-ready `PLAN.md`

## Exit Bar

- the Stitch structure is mapped to the current overlay honestly
- the retained general-design requirements are explicit
- the Tk implementation direction is concrete enough to execute without exploratory churn
- the verification plan distinguishes pass-local proof from task-level regression
