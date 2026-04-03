# Task 0002 Research Analysis

## Problem 0001 Overlay Structure Mapping

### Current Surface Facts

The existing overlay already contains most of the product ingredients required by `Task-0002`:

- a title row with hide and quit actions
- a live ingest status line
- interval buttons for `1m`, `5m`, `15m`, `1h`, and `1d`
- a single token-velocity bar chart
- a seven-day total
- projected weekly burn and redline state
- a weekly budget editor
- a startup toggle
- advisory Codex weekly-window context

### Current Drift From The Stitch Mockup

The current overlay is functionally correct but visually too loose and too form-like:

- the header is taller and more generic than the compressed Stitch strip
- interval controls live in their own row instead of the header utility zone
- metric presentation is label-plus-text rather than compact instrument cards
- the chart field has the right data but not the monolithic control-room framing
- the footer actions and live-state strip do not match the Stitch composition
- budget, startup, and advisory context are presented as a form row instead of a secondary operational rail

### Product-Intent Overrides

The repo-root general design remains authoritative where the mockup is too minimal:

- the redesign must keep a clear immediate hide path
- weekly burn and redline meaning must remain visible and honest
- weekly budget editing must stay present
- the startup toggle must stay present
- advisory Codex weekly-window context must stay present when available
- the product must remain a compact overlay, not widen into a console browser

## Problem 0002 Tk Implementation Fit

### Honest Repo Fit

The redesign should stay inside `app/codex_dashboard/ui.py` and reuse the current data-refresh and chart logic rather than rewriting the app architecture.

### Practical Implementation Direction

The current `ttk`-heavy layout is good enough structurally, but explicit Tk widgets will give tighter control over the visual system:

- use `tk.Frame`, `tk.Label`, `tk.Button`, `tk.Entry`, and `tk.Canvas` for the visible overlay surface
- keep `ttk` only where it still meaningfully helps, or remove it entirely from the overlay surface
- preserve the existing chart canvas and bucket rendering logic, but restyle the field, grid, threshold line, and label treatment
- compress the overlay into a tighter header, metrics strip, chart field, and footer rail
- move interval controls into the header utility area
- move budget editing, startup toggle, and advisory context into the lower operational rail instead of a standalone form block

This keeps the work bounded to interface implementation while preserving proven ingest, aggregation, hotkey, and startup behavior.

## Problem 0003 Proof Strategy

Pass-local proof should stay honest:

- unit tests can still cover format helpers and desktop-support helpers
- the scripted ingest smoke can still prove the underlying token model remains intact
- task-level closure still requires the repo-root desktop regression lane from `REGRESSION.md`

For this redesign pass, the realistic pass-local bar is:

- run the focused unit suite
- run the scripted ingest smoke
- capture pass-local UI evidence only if it helps the audit

## Research Verdict

Task research is planning-ready.

## Recommended Direction

- plan the task in two passes
- use `PASS-0000` for the Stitch-aligned overlay redesign and pass-local proof
- use `PASS-0001` for repo-root regression, final handoff, and closure
- keep the implementation concentrated in `ui.py` unless a small supporting helper or test update is genuinely needed
