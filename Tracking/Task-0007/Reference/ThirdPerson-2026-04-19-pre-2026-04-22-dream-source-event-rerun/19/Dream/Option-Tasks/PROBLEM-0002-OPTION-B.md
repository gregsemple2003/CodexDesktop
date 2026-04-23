# Title

Make stop state, ownership, and active constraints first-class task state

## Summary

Extend durable task state so workers must record ownership, stop state, resume trigger, and active constraints before they pause, hand back, or declare a checkpoint stop.

## Source Event IDs

```json
[
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-3447",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-3692",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4015",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-4432",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4294",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-4684",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-4755",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-4919",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5108",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5272",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5300",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5310",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-6558",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-7038"
]
```

## Burden Being Reduced

The human should not have to reactivate owned work, restate "no engine mods", restate the pass structure, or ask why they are being messaged after the system already knew those facts.

## Causal Claim

The burden persists because ownership and constraint facts live in prose and local context instead of in a task-state contract that later workers and stop logic must obey.

## Evidence

- [../BURDEN-ANALYSIS.md](../BURDEN-ANALYSIS.md) identifies ownership-memory failure as a separate burden driver.
- [../SessionExcerpts/INDEX.json](../SessionExcerpts/INDEX.json) preserves the checkpoint-to-reopen and PASS-0009 wake-up neighborhoods.
- [../../HumanInterventionTime/SUMMARY.json](../../HumanInterventionTime/SUMMARY.json) quantifies the wake-up cost instead of treating it as mere annoyance.

## Goals

- make stop state durable
- make current owner and expected next actor durable
- preserve new constraints immediately when introduced
- prevent hand-back of intermediate failure when the task is still owned work

## Non-Goals

- solving every orchestration scheduling problem
- replacing repo-local task artifacts
- inferring hidden external approvals automatically

## Implementation Home

- `C:\Users\gregs\.codex\Orchestration\TASK-STATE.md`
- `C:\Users\gregs\.codex\Orchestration\TASK-STATE.schema.json`
- any shared worker prompt or validator that reads task state before stop or handoff
- repo task artifacts that mirror the new fields in `TASK-STATE.json`

## Proposed Changes

- add explicit task-state fields for:
  - active owner
  - stop state
  - stop reason
  - resume trigger
  - active constraints
  - whether the worker may hand back intermediate failure
- require those fields to be updated before stop, reopen, or planning/implementation gate changes
- add validation that blocks worker stop messages when the state is incomplete or contradictory

## Expected Resolution

When the human says "planning gate only", "no engine mods", "continue under your existing ownership", or "do not stop after a failed incremental attempt", those facts survive as durable state instead of vanishing into local context.

## What Does Not Count

- storing the fields but not gating stop behavior with them
- putting constraints only in HANDOFF prose while task state stays generic
- a schema field named vaguely enough that workers can ignore its meaning

## Why This Mechanism

The packet's wake-up cost is a state-management failure. The fix should therefore live at the state boundary, not only in reminders.

## Human Relief If Successful

The human no longer has to keep work alive manually or restate already-decided constraints when the task crosses turns and passes.

## Remaining Uncertainty

This task still depends on workers and validators actually consuming the new state fields consistently.

## Falsifier

If later packets still show the human restating active constraints or existing ownership after those fields were already set, the state contract is too weak or not enforced.

## Acceptance Criteria

- shared task-state contract includes fields for ownership, stop state, resume trigger, and active constraints
- schema validation rejects missing or contradictory values for those fields
- worker stop or handoff behavior is gated on those fields being current
- repo task artifacts can mirror the new fields without rewriting closed pass history

