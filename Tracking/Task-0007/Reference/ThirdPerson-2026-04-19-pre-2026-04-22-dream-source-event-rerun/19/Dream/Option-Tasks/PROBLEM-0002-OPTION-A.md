# Title

Add prompt-level continuity and constraint reminders

## Summary

Tighten shared prompts so workers are reminded to preserve ownership, pause state, and newly introduced constraints before stopping or replying.

## Source Event IDs

```json
[
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-3447",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-3692",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4015",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-4432",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-4684",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-4755",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5272",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5310",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-7038"
]
```

## Burden Being Reduced

The human keeps having to wake work back up, restate ownership, and restate constraints that should already have been durable.

## Causal Claim

Workers are not being reminded strongly enough to treat pause state, existing ownership, and new constraints as first-class facts before stopping.

## Evidence

- [../BURDEN-ANALYSIS.md](../BURDEN-ANALYSIS.md) identifies ownership-memory failure as a separate burden driver.
- [../SessionExcerpts/INDEX.json](../SessionExcerpts/INDEX.json) preserves the reopen-after-checkpoint neighborhoods that turned stop state into human work.
- [../../HumanInterventionTime/SUMMARY.json](../../HumanInterventionTime/SUMMARY.json) records three restart-supervision events with charged stall loss.

## Goals

- remind workers to restate active ownership and stop state before stopping
- remind workers to persist new constraints immediately
- reduce avoidable wake-up supervision

## Non-Goals

- changing task-state schema
- adding orchestration-level enforcement
- solving idle detection mechanically

## Implementation Home

- shared orchestration prompt bundles that govern execution and handoff behavior
- any shared stop-or-reply checklist used by those prompts

## Proposed Changes

- add prompt clauses that require:
  - ownership restatement before stop
  - explicit pause versus continue distinction
  - immediate capture of new constraints such as pass structure or no-engine-mods
- add reminder examples drawn from wake-up failures like those in this packet

## Expected Resolution

Workers are more likely to preserve the right state before they stop.

## What Does Not Count

- adding reminder language with no place to store the state being reminded about
- broader prose about "ownership" that does not mention pause, resume, and active constraints explicitly

## Why This Mechanism

This is the cheapest intervention and may reduce some obvious misses quickly.

## Human Relief If Successful

Some repeated restatements may disappear even before deeper orchestration work lands.

## Remaining Uncertainty

The packet already contains many strong instructions. That makes reminder-only improvement structurally weak.

## Falsifier

If later workers still need the human to restate ownership or constraints after the reminders are added, this option is insufficient.

## Acceptance Criteria

- the relevant shared prompts explicitly name ownership, stop state, resume trigger, and active constraints
- prompt examples cover pause gates, wake-up failure, and constraint persistence
- stop or handoff replies mention those fields when they are known

