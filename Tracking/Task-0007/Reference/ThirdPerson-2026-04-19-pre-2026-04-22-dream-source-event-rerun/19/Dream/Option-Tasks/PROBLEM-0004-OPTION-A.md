# Title

Strengthen debugging prompts around first-disagreement tracing

## Summary

Revise shared and repo-local debugging prompts so runtime defect work more consistently asks for the first disagreement with values before applying tweaks.

## Source Event IDs

```json
[
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4768",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5119",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5154",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5340",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-6105",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-6755",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-7446"
]
```

## Burden Being Reduced

The human should not have to keep converting "improve the look" work back into first-disagreement debugging.

## Causal Claim

Workers are not being reminded clearly enough that runtime defect work must stay attached to a concrete disagreement and its writer chain.

## Evidence

- [../BURDEN-ANALYSIS.md](../BURDEN-ANALYSIS.md) identifies tweak-mode debugging as a separate burden driver.
- [../SessionExcerpts/INDEX.json](../SessionExcerpts/INDEX.json) preserves the final explicit debugging directive.

## Goals

- remind workers to identify the first disagreement with values
- remind workers to trace the writer chain upstream
- reduce symptom-only fix attempts

## Non-Goals

- adding runtime probes
- changing bug artifact schemas
- blocking closure mechanically

## Implementation Home

- `C:\Users\gregs\.codex\Orchestration\Processes\DEBUGGING.md`
- repo-local `DEBUGGING.md`
- shared debugging prompt bundles

## Proposed Changes

- add stronger examples of acceptable disagreement seams
- add a reject list for symptom-only closeout language
- restate that runtime-only claims require runtime-only evidence

## Expected Resolution

Workers reach for the right debugging method more often before applying retunes.

## What Does Not Count

- more exhortation without an explicit reject list
- examples that still allow "category only" explanations like "single-node seam"

## Why This Mechanism

It is the cheapest way to align the written method with what the human had to restate manually.

## Human Relief If Successful

Some tweak churn may disappear before heavier instrumentation lands.

## Remaining Uncertainty

This packet already cites the shared debugging method directly, so reminder-only gains may be small.

## Falsifier

If later runtime bugs still advance without a first disagreement named in values, wording-only changes are inadequate.

## Acceptance Criteria

- shared and repo-local debugging docs explicitly require first-disagreement values and writer-chain tracing
- docs explicitly reject symptom-only closure language
- runtime-only evidence rules remain explicit in the debugging flow

