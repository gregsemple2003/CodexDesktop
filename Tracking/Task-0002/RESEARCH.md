# Task 0002 Research Summary

## Research Verdict

Task research is planning-ready.

## Key Decisions

- redesign the overlay in place inside `app/codex_dashboard/ui.py`
- use the Stitch mockup as the primary composition reference
- keep `Design/GENERAL-DESIGN.md` authoritative for budget editing, startup toggle, weekly burn/redline semantics, and advisory context
- keep a single token-velocity chart and avoid widening into any multi-pane browser
- shift the visible overlay surface toward explicit Tk composition for tighter control over spacing, color, typography, and footer/header layout

## Why This Is The Honest Repo Fit

- the current app already has the required data model and interaction model
- most task scope is present semantically already; the gap is visual grouping and control-room density
- reusing the current refresh, ingest, and chart logic keeps the task bounded to UI implementation

## Carry-Forward Constraints

- preserve hotkey-first overlay behavior
- preserve immediate dismissability
- keep the dark operational palette and redline emphasis
- keep budget editing, startup toggle, and advisory context visible
- do not change ingest, persistence, or aggregation semantics

## Planning Recommendation

Plan the work in two passes:

- `PASS-0000`: Stitch-aligned overlay redesign, supporting tests, and pass closeout
- `PASS-0001`: repo-root regression, final handoff, and closure

## Remaining Open Questions

- whether the real app-surface smoke for the redesign pass should stay on the existing postscript artifact path or add a fuller overlay capture later

This does not block planning or implementation.
