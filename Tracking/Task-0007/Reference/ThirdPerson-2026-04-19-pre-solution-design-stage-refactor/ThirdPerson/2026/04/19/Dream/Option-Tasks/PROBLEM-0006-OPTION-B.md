# PROBLEM-0006 OPTION-B: Global Constraint Memory Store With Auto-Injection

## Title

Add a global constraint memory store that persists constraints across sessions and auto-injects them into prompts.

## Summary

A task-local ledger reduces loss within a task, but constraints can still be dropped across sessions or when task context is not loaded. This option adds a global memory store keyed by repo + task (and optionally by pass) and automatically injects active constraints into prompts.

Writeup type: research -> concrete implementation (two-stage).

## Goals

- Persist constraints across session boundaries without relying on manual re-reading of task artifacts.
- Reduce repeated durable-learning interventions even when the context window changes.

## Non-Goals

- Replacing task-local durable artifacts; the global store should point at the source-of-truth ledger.

## Implementation Home

- Orchestration storage layer (SQLite or equivalent) that already persists session telemetry.
- Prompt assembly layer that can inject constraints into the active worker prompt.

## Proposed Changes

Stage 1 (research):

- Define the minimal schema:
  - `constraint_id`, `scope`, `status`, `source_event_id`, `source_artifact_path`
  - last-updated timestamp
- Define injection rules:
  - inject only `new|applied` constraints
  - cap size and collapse duplicates

Stage 2 (concrete implementation):

- Implement storage and an API for:
  - adding/updating constraints by `event_id`
  - fetching active constraints for a repo/task
- Implement prompt auto-injection of active constraints for each new pass boundary.

## Acceptance Criteria

- Constraints stated in one session are present automatically in the next session's prompt context.
- The injected constraints point back to the task-local ledger (source of truth) and include stable `event_id` provenance.

## Source Event IDs

```json
[
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-3094",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-3197",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-3463",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-3482",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-3549",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-3585",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4294",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-4684"
]
```

## Burden Reduced

Constraint loss across sessions and context boundaries.

## Causal Claim

If active constraints are persisted globally and auto-injected into prompts, then constraint re-statement and violation will drop even when task artifacts are not reloaded manually.

## Evidence

- `BD-006` in [`../BURDEN-ANALYSIS.md`](../BURDEN-ANALYSIS.md)

## Why This Mechanism

The packet suggests constraints are not retained durably enough for continuity. Auto-injection makes the "remember constraints" step automatic at pass boundaries.

## Human Relief If Successful

- Fewer "new constraint" restatements.
- Less time spent policing "do not touch" boundaries.

## Remaining Uncertainty

- Global stores can become a second source of truth; the design must ensure the ledger remains canonical.

## Falsifier

- Constraints are injected but still not followed (meaning enforcement is still missing).

## What Does Not Count

- A store that accumulates stale constraints and injects them without status/scoping, increasing confusion.

## References

- `BD-006` in [`../BURDEN-ANALYSIS.md`](../BURDEN-ANALYSIS.md)

