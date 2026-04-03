# Task 0002 Plan

## Planning Verdict

This plan is approved for execution under the standing user instruction to auto-approve plan, pass, and phase gates that do not expand scope, spend money, delete data, or violate workflow rules.

## Implementation Strategy

Reshape the existing Tk overlay into the Stitch-inspired composition while preserving the product-intent requirements from `Design/GENERAL-DESIGN.md`.

Implementation should:

- keep the work centered in `app/codex_dashboard/ui.py`
- preserve the current hotkey, ingest refresh, and chart bucket logic
- compress the visible surface into:
  - a utility header with title, interval controls, and dismiss path
  - a compact metrics strip
  - a single focal token-velocity chart
  - a lower operational rail that still carries budget editing, startup state, advisory context, and action controls

## Pass Order

### PASS-0000 Stitch-Aligned Overlay Redesign

Objective:

- implement the redesign of the real Tk overlay surface without widening task scope beyond UI work

Expected work:

- restyle and restructure the overlay shell in `ui.py`
- keep the interval controls, main chart, seven-day total, and redline semantics
- retain budget editing, startup toggle, and advisory context in a more compact operational layout
- update any supporting docs or tests needed to keep the redesign honest

Verification:

- run `python -m unittest discover -s tests -p "test_*.py" -v`
- run `python -m app.codex_dashboard --scan-once --print-summary`
- capture pass-local UI evidence if needed for the audit

Exit bar:

- the real overlay structure visibly matches the Stitch composition more closely than the current baseline
- the general-design requirements are still present
- the redesign remains hotkey-first and dismissible
- pass-local proof passes

### PASS-0001 Regression, Handoff, And Closure

Objective:

- execute honest task-level proof and close the task if the redesigned overlay passes the repo-root lane

Expected work:

- execute repo-root `REG-001 Desktop Overlay Launch And Data Smoke`
- write the task-owned regression artifact
- finalize handoff and closure state

Verification:

- run the repo-root regression lane from `REGRESSION.md`

Exit bar:

- the acceptance criteria are satisfied or any gap is explicit
- the task-owned regression artifact names the actual lane run
- `HANDOFF.md` and `TASK-STATE.json` reflect the honest final state
