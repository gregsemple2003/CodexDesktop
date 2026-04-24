# Task 0011 Plan

## Planning Verdict

This plan is ready for explicit human approval.

Human gate:

- do not start implementation passes until the human agrees that `Review` is the canonical intake surface for incoming asks while `Tasks` remains the canonical surface for committed work

## Planning Basis

This plan is grounded in:

- the repo design anchor at [Design/GENERAL-DESIGN.md](../../Design/GENERAL-DESIGN.md)
- the committed-work split preserved in [Task-0009](../Task-0009/TASK.md)
- the Dream generation and promotion split preserved in [Task-0010](../Task-0010/TASK.md)
- the review-surface design briefs under [Design/](./Design)
- the task-local survey at [Research/2026-04-23-REVIEW-SURFACE-SURVEY.md](./Research/2026-04-23-REVIEW-SURFACE-SURVEY.md)

## PASS-0000 Lock The Product Split

### Objective

Turn the moved review docs into one stable first-release contract for `Review`.

### Implementation Notes

- freeze the separation between `Review` and `Tasks`
- decide the first-release ask taxonomy
- decide the summary model and grouped ask-stream structure
- decide which dispositions are real in the first slice and which remain disabled or deferred

### Verification

- the task-local design briefs and survey agree on the first-release screen shape
- the first-release scope is explicit enough to implement without inventing behavior ad hoc

### Exit Bar

- the team can build the first `Review` tab without rediscovering what human job it is for

## PASS-0001 Build The Read-Only Review Surface

### Objective

Add the first real `Review` tab UI with trustworthy read-only intake visibility.

### Implementation Notes

- add the new tab shell and navigation entry
- render the chosen top summary regions
- render the grouped ask stream and persistent detail surface
- support the main non-destructive states:
  - empty
  - loading
  - populated
  - stale
  - backend unavailable

### Verification

- the dashboard can render `Review` without live dispositions enabled
- the surface communicates useful intake state without requiring email-first or packet-first investigation

### Exit Bar

- `Review` is already useful as an intake-reading surface before live actions land

## PASS-0002 Add Dispositions And Downstream Handoffs

### Objective

Integrate the humane intake surface with the real backend actions and downstream product surfaces.

### Implementation Notes

- consume the promotion contract delivered by [Task-0010](../Task-0010/TASK.md)
- expose only the bounded dispositions whose contracts are real and explainable
- preserve provenance when work moves from `Review` into `Tasks` or other downstream lanes

### Verification

- at least one real candidate ask can be promoted through the single backend promotion contract
- the surface makes downstream consequences explicit before the human acts
- reviewed items do not lose their origin story after handoff

### Exit Bar

- `Review` is useful for real decision-making, not only passive reading

## PASS-0003 Polish, Audit, And Regression

### Objective

Make the tab trustworthy under real human use instead of only technically complete.

### Implementation Notes

- tighten copy, hierarchy, grouping, and placeholder truth
- verify that `Review` stays calm and source-aware under mixed ask types
- capture repo-root regression proof for the new human-facing lane

### Verification

- unit tests cover any new ask-state mapping logic
- regression proves the real dashboard surface, not only backend readback
- task-owned audit confirms that `Review` still behaves like a humane decision surface instead of an inbox dump

### Exit Bar

- the first `Review` tab slice can be used as the canonical intake surface without apologizing for it

## Task-Level Validation

Expected validation before closure:

- intake-surface unit coverage where mapping or formatting logic exists
- repo-root regression for the real dashboard surface
- task-owned audit against the review design briefs and copy semantics

## Watchouts

- do not let `Review` become a raw inbox or packet browser
- do not let `Review` silently redefine the backend actions owned elsewhere
- do not let incoming asks look like already-committed tasks
- do not bury high-attention approvals or anomalies under speculative candidate work
