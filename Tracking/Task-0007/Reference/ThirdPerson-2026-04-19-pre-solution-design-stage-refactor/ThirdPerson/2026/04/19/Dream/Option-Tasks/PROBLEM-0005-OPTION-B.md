# PROBLEM-0005 OPTION-B: Runtime Disagreement Capture Harness (Feet/Ground/Root)

## Title

Add a repo-local harness that captures the first-disagreement values at runtime (feet world Z vs ground, pelvis roll, etc.) and writes a durable JSON artifact.

## Summary

The debugging contract demands concrete disagreement values. This option makes capturing those values cheap and repeatable by adding a small runtime telemetry harness for the default lane.

Writeup type: concrete implementation.

## Goals

- Make first-disagreement data capture fast and durable.
- Reduce guesswork and evidence drift in runtime animation/debugging work.

## Non-Goals

- Replacing manual default-lane proof for human-facing validation.

## Implementation Home

- Repo-local code and scripts under `C:\\Agent\\ThirdPerson`.
- Task-owned artifacts under `Tracking/Task-<id>/Research/` or `Tracking/Task-<id>/Testing/`.

## Proposed Changes

- Implement a harness that, on the default lane:
  - records per-frame (or sampled) values:
    - foot/toe world Z, ground/capsule bottom reference Z
    - pelvis/root roll and lateral offset
    - component-space foot transforms
  - writes `runtime-disagreement-<timestamp>.json` with:
    - lane description
    - commit under test
    - capture parameters
    - sampled values
- Add a small summarizer that:
  - highlights frames where disagreement exceeds thresholds (hover, rocking)
  - prints the top disagreements and their values for the bug note

## Acceptance Criteria

- Running the harness on the default lane produces a durable JSON artifact.
- A bug note can cite concrete values from this artifact as the first disagreement seam.
- The capture does not require engine modifications under `C:\\Agent\\UnrealEngine`.

## Source Event IDs

```json
[
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4768",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-6105",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5119",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5129",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5154",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-6755",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5340",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-7446"
]
```

## Burden Reduced

Human time spent arguing about categories or visual impressions when the needed truth is a numeric disagreement seam.

## Causal Claim

If the system can capture the first-disagreement values durably on the default lane, then debugging will converge faster and will require less human intervention to keep it grounded.

## Evidence

- Debugging seam examples demanded explicitly in `EXCERPT-0008` in [`../SessionExcerpts/INDEX.json`](../SessionExcerpts/INDEX.json)

## Why This Mechanism

This option operationalizes the shared debugging method by lowering the cost of capturing the required evidence.

## Human Relief If Successful

- Less repeated clarification about what disagreement matters.
- Faster narrowing because values are captured and comparable over time.

## Remaining Uncertainty

- Capturing ground-contact references robustly in all maps/lanes can be tricky; the initial scope should target the repo default lane only.

## Falsifier

- The harness outputs data that is not actually useful for narrowing (does not align with human-observed hovering/rocking frames).

## What Does Not Count

- Telemetry captured on a non-default lane and then used to dismiss default-lane defects.

## References

- `BD-005` in [`../BURDEN-ANALYSIS.md`](../BURDEN-ANALYSIS.md)

