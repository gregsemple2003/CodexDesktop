# Task 0009 Plan

## Planning Verdict

This plan is ready for explicit human approval.

Human gate:

- do not start implementation passes until the human agrees that this task owns the committed-work `Tasks` tab while [Task-0011](../Task-0011/TASK.md), [Task-0008](../Task-0008/TASK.md), and [Task-0010](../Task-0010/TASK.md) own the adjacent intake and backend subsystems

## Planning Basis

This plan is grounded in:

- the current repo design anchor at [Design/GENERAL-DESIGN.md](../../Design/GENERAL-DESIGN.md)
- the backend-and-dashboard runtime split delivered by [Task-0005](../Task-0005/TASK.md)
- the intervention and Dream analysis preserved in [Task-0007](../Task-0007/TASK.md)
- the review/intake split preserved in [Task-0011](../Task-0011/TASK.md)
- the task-local design brief at [Design/HUMAN-NEED-AND-TASKS-TAB-DIRECTION.md](./Design/HUMAN-NEED-AND-TASKS-TAB-DIRECTION.md)

## PASS-0000 Lock The Product Surface

### Objective

Turn the task-local design brief into a stable product-surface contract for the first `Tasks` tab slice.

### Implementation Notes

- decide the first-release information architecture
- decide the primary summary regions and task list model
- decide the default row status taxonomy and action affordances
- keep backend and orchestration leakage out of the default human-facing copy

### Verification

- the task-local design brief and `Stitch` prompt agree on the core surface
- the first-release screen structure is explicit enough to implement

### Exit Bar

- the team can build the first `Tasks` tab without rediscovering what the screen is for

## PASS-0001 Build The Read-Only Task Surface

### Objective

Add the first real `Tasks` tab UI with trustworthy read-only task visibility.

### Implementation Notes

- add the new tab shell and navigation entry
- render the chosen summary regions
- render the task list and detail surface
- support the main non-destructive states:
  - empty
  - loading
  - populated
  - stale
  - backend unavailable

### Verification

- the dashboard can render the `Tasks` tab without dispatch actions enabled
- the tab communicates useful task state without requiring markdown-first investigation

### Exit Bar

- the `Tasks` tab is a real high-level monitoring surface even before active dispatch controls land

## PASS-0002 Add Dispatch And Recovery Actions

### Objective

Integrate the humane task surface with the durable execution-state and dispatch layer.

### Implementation Notes

- consume the APIs and contracts delivered by [Task-0008](../Task-0008/TASK.md)
- add bounded task actions such as:
  - dispatch
  - poke
  - interrupt
  - open thread or working context
- keep each action specific enough that the human knows what will happen before clicking

### Verification

- a dispatched task reflects live durable state changes on the `Tasks` tab
- the surface can distinguish sleeping, blocked, running, and waiting-for-human honestly
- thread or session launch works from the task surface

### Exit Bar

- the `Tasks` tab is now useful for both monitoring and control, not only for read-only review

## PASS-0003 Integrate Reviewed And Promoted Work Provenance

### Objective

Show reviewed and promoted work on the committed-work surface without turning `Tasks` into the intake queue.

### Implementation Notes

- consume the promotion provenance delivered by [Task-0010](../Task-0010/TASK.md) and the intake split delivered by [Task-0011](../Task-0011/TASK.md)
- separate:
  - authored tasks
  - promoted tasks with durable provenance
  - enqueued tasks
  - active dispatched runs
- make promoted-task lineage and enqueue state legible without surfacing still-pending review items as committed work

### Verification

- a human can tell the difference between an authored task and a promoted committed task at a glance
- promoted tasks appear on the `Tasks` surface without ambiguous provenance or intake-state leakage

### Exit Bar

- the `Tasks` tab becomes the genuine top-level committed-work surface across authored and promoted tasks

## PASS-0004 Polish, Audit, And Regression

### Objective

Make the tab trustworthy under real human use instead of only technically complete.

### Implementation Notes

- run the interface against the task-local design brief
- tighten copy, hierarchy, placeholder truth, and visible status semantics
- capture repo-root regression proof for the new human-facing lane

### Verification

- unit tests cover any new task-state mapping logic
- regression proves the real dashboard surface, not only backend readback
- task-owned audit confirms the tab still behaves like a humane high-level surface

### Exit Bar

- the first `Tasks` tab slice can be used as the product heart without apologizing for it

## Task-Level Validation

Expected validation before closure:

- task-surface unit coverage where mapping or formatting logic exists
- repo-root regression for the real dashboard surface
- task-owned audit against the design brief and copy semantics

## Watchouts

- do not let the UI own the dispatch runtime
- do not let raw backend taxonomy become the default visible language
- do not hide uncertainty by collapsing sleeping, blocked, and waiting-for-human into one vague label
- do not make the `Tasks` tab depend on transcript reading for basic comprehension
