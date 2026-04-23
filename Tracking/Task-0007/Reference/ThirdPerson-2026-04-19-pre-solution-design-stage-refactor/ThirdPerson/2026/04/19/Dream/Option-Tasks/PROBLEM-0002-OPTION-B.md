# PROBLEM-0002 OPTION-B: Watchdog To Detect Dropped Work And Trigger Resume

## Title

Add a watchdog that detects dropped work (no explicit wait state + idle) and triggers a resume or escalates to a clear operator notification.

## Summary

Some dropped-work cases are not prevented purely by policy. This option adds an active watchdog that monitors for idle gaps without an explicit wait state and takes an action: auto-resume, or alert the human with a precise "stopped without a wait gate" message.

Writeup type: concrete implementation.

## Goals

- Reduce wake-up supervision and stall-loss time by detecting drops automatically.
- Convert ambiguous idle into an explicit next action (resume) or explicit escalation (notification).

## Non-Goals

- Making proof claims truthful (handled by PROBLEM-0001).
- Fixing pass content; this is a lifecycle/continuity mechanism.

## Implementation Home

- Orchestration backend / dashboard runtime that can:
  - observe durable task state
  - observe last meaningful AI activity timestamps
  - schedule a resume action or send an operator notification

## Proposed Changes

- Define "dropped" heuristics:
  - idle gap exceeds a threshold AND no explicit wait gate is set AND there is a valid next step recorded
- Add actions:
  - auto-resume the task, or
  - send a notification that includes task id, pass id, last known next step, and which stop contract was violated
- Record watchdog actions durably to avoid hidden behavior.

## Acceptance Criteria

- In a comparable long-running pass, a dropped-work condition triggers exactly one of:
  - a successful auto-resume that continues from the recorded next step, or
  - a notification that makes the stop-state violation explicit
- Wake-up events (manual "continue") are reduced materially.

## Source Event IDs

```json
[
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-3447",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-3692",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-3929",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4015",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4590",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-4919",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5108",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-6558",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5272",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5300",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5310",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-7038"
]
```

## Burden Reduced

Human time spent prodding a dropped task back into motion.

## Causal Claim

If dropped work is detected and resumed automatically (or escalated clearly), then wake-up interventions will drop even when the underlying agent forgets to set an explicit wait gate.

## Evidence

- `BD-002` in [`../BURDEN-ANALYSIS.md`](../BURDEN-ANALYSIS.md)
- Wake-up events and stall loss in `../../HumanInterventionTime/SUMMARY.json`

## Why This Mechanism

This mechanism is active: it does not rely on the agent complying with policy perfectly.

## Human Relief If Successful

- Fewer "continue now" interventions.
- Clearer notifications that are actionable instead of silent stalls.

## Remaining Uncertainty

- Auto-resume can be risky if the environment changed; the watchdog must prefer explicit escalation when a safe resume is unclear.

## Falsifier

- Wake-up events remain necessary because the watchdog misclassifies or cannot trigger a safe action.

## What Does Not Count

- Auto-resume that causes the task to run past a pause gate or approval-only constraint.

## References

- `BD-002` in [`../BURDEN-ANALYSIS.md`](../BURDEN-ANALYSIS.md)

