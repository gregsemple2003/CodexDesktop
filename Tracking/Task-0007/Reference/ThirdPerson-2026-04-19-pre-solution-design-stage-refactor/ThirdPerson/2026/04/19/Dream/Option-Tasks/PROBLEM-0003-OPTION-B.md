# PROBLEM-0003 OPTION-B: Dashboard "Unanswered Questions" Surface

## Title

Add a dashboard surface that detects explicit questions in human inputs and flags them until a direct answer is given.

## Summary

Even with prompt guidance, unanswered questions can slip. This option adds a UI surface that extracts questions from the packet stream and marks them as unresolved until the system produces a direct answer.

Writeup type: concrete implementation.

## Goals

- Make unanswered questions visible without re-reading the full transcript.
- Reduce the need for the human to stop the system to demand an answer.

## Non-Goals

- Replacing prompt-based answer-first discipline (OPTION-A is still preferred as first move).

## Implementation Home

- The dashboard UI that consumes Codex session telemetry (question extraction + unresolved list + links to source events).

## Proposed Changes

- Add a question extractor that:
  - identifies explicit question patterns (question marks, "agree/disagree", "yes/no")
  - tags the originating `HumanInputEvents.event_id`
- Add an "Unanswered Questions" panel that:
  - lists active questions in chronological order
  - marks them resolved when a subsequent assistant message contains a direct answer signature

## Acceptance Criteria

- In a comparable session, the dashboard shows unanswered questions within seconds of the human input.
- After the system answers directly, the question is marked resolved without manual cleanup.

## Source Event IDs

```json
[
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-3392",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5001",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5010",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5020",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5030",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5073",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5090",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5263",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5281",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5290"
]
```

## Burden Reduced

Human effort spent noticing unanswered questions and re-asking them.

## Causal Claim

If unanswered questions are surfaced as an explicit UI list, then the system can correct itself earlier and the human will not need to intervene as often to demand answers.

## Evidence

- `BD-003` in [`../BURDEN-ANALYSIS.md`](../BURDEN-ANALYSIS.md)
- Direct-answer exemplar: `EXCERPT-0005` in [`../SessionExcerpts/INDEX.json`](../SessionExcerpts/INDEX.json)

## Why This Mechanism

The UI acts as a secondary enforcement surface when prompt-level discipline fails.

## Human Relief If Successful

- Faster detection of evasion/omission.
- Less repeated correction.

## Remaining Uncertainty

- Automatically detecting "answered" reliably is non-trivial; false positives must be avoided.

## Falsifier

- The UI is noisy or fails to track question resolution accurately, and does not reduce interventions.

## What Does Not Count

- A UI list that lacks stable links back to source `event_id` values.

## References

- `BD-003` in [`../BURDEN-ANALYSIS.md`](../BURDEN-ANALYSIS.md)

