# PROBLEM-0004 OPTION-B: Interactive Review UI For Approval Gates

## Title

Build a dashboard UI that renders approval packets as interactive diffs with context.

## Summary

If static approval packets remain insufficient, an interactive UI can make review even faster and reduce the chance of misreading changes. This option is heavier than OPTION-A and should follow after the standard packet shape is in place.

Writeup type: concrete implementation.

## Goals

- Reduce approval friction further by making diffs and context browsable.
- Make "what changed" visible without opening many separate files.

## Non-Goals

- Redefining regression lanes or task lifecycle rules.

## Implementation Home

- The dashboard UI and storage layer that can:
  - ingest the standard approval packet
  - render diffs with context
  - link back to stable task artifact paths

## Proposed Changes

- Add an "Approvals" panel that shows:
  - pending approval packet
  - per-file diffs with anchors
  - a one-click approval/reject action that records the decision durably
- Add a minimal schema for an approval packet in storage to support rendering.

## Acceptance Criteria

- A reviewer can complete an approval decision without leaving the UI.
- The rendered view includes the same "what changed" content as the standard approval packet, plus easier navigation.

## Source Event IDs

```json
[
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-3494",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4379",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4399",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4421",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-4755",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4512",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4522",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4571",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4581"
]
```

## Burden Reduced

The time cost of reviewing diff context across multiple files and tabs.

## Causal Claim

If approvals are rendered as interactive diffs with stable links, then approval decisions will become faster and more reliable than static packets alone.

## Evidence

- `BD-004` in [`../BURDEN-ANALYSIS.md`](../BURDEN-ANALYSIS.md)

## Why This Mechanism

Some approval friction comes from editor/file hopping. A single UI surface can reduce that cost.

## Human Relief If Successful

- Faster approvals.
- Less fatigue reconstructing changes.

## Remaining Uncertainty

- Building a good diff UI is non-trivial; this option should follow after OPTION-A establishes the canonical packet schema.

## Falsifier

- The UI is not materially faster than the static approval packet, or it cannot render key context reliably.

## What Does Not Count

- A UI panel that lists file names but still requires manual diff reconstruction elsewhere.

## References

- `BD-004` in [`../BURDEN-ANALYSIS.md`](../BURDEN-ANALYSIS.md)

