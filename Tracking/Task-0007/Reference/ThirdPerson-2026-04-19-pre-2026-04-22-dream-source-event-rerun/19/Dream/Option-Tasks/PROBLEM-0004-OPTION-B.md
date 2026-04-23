# Title

Require first-disagreement bug packets and runtime disagreement probes

## Summary

Make runtime defect work produce a bug packet that names the first disagreement with values, captures the supporting runtime probes, and traces the writer chain before closure or "improved" claims can land.

## Source Event IDs

```json
[
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4114",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4768",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5119",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5129",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5154",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5340",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-6105",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-6755",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-7038",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-7446"
]
```

## Burden Being Reduced

The human should not have to keep pushing runtime defects out of tweak mode and into real root-cause tracing.

## Causal Claim

The burden persists because runtime defect artifacts do not require the first disagreement, supporting measurements, and upstream writer chain before the task is allowed to describe progress or closure.

## Evidence

- [../BURDEN-ANALYSIS.md](../BURDEN-ANALYSIS.md) names debugging discipline as a distinct burden driver.
- [../SessionExcerpts/INDEX.json](../SessionExcerpts/INDEX.json) preserves the final first-disagreement directive.
- [../../HumanNeeds/PACKET-RECORD.json](../../HumanNeeds/PACKET-RECORD.json) already recommends a debugging gate that forbids symptom-only closure.

## Goals

- require first-disagreement values in bug artifacts
- capture runtime probes that support those values
- require an upstream writer chain or explicit contradictory branch before progress claims

## Non-Goals

- solving a specific ZeroMale defect
- replacing manual narrative in bugs entirely
- generalizing every runtime genre at once

## Implementation Home

- `C:\Users\gregs\.codex\Orchestration\Processes\DEBUGGING.md`
- repo-local `DEBUGGING.md`
- repo bug artifacts under `C:\Agent\ThirdPerson\Tracking\Task-<id>\BUG-<NNNN>.md`
- repo regression artifacts under `C:\Agent\ThirdPerson\Tracking\Task-<id>\Testing\`
- repo-local runtime probes under `C:\Agent\ThirdPerson\Source\GameAutomation\` or equivalent runtime automation surfaces

## Proposed Changes

- extend the bug-artifact contract with sections for:
  - first disagreement
  - exact values
  - probe source
  - writer chain
  - contradictory evidence
- add runtime probes that can capture the kinds of disagreements named in this packet, such as:
  - foot or toe world Z versus ground or capsule bottom
  - pelvis or root roll
  - unexpected component-space transforms while the capsule remains upright
- block "fixed" or "improved" conclusions when the required bug packet sections are empty

## Expected Resolution

Runtime defect work becomes falsifiable earlier, because every fix claim must point at a concrete disagreement and the writer chain that changed it.

## What Does Not Count

- a bug template that names the sections but allows them to stay blank
- probes that dump data without linking it to a first disagreement
- a progress note that says "looks better" while the first disagreement is still unknown

## Why This Mechanism

This is the smallest durable boundary that matches the human's explicit request: identify the disagreement with values, then trace upstream until the cause is proven or the contradictory branch is preserved honestly.

## Human Relief If Successful

The human no longer has to keep turning symptom reports into a debugging method; the artifacts themselves demand that method.

## Remaining Uncertainty

The first probe set may need iteration before it captures every disagreement class cheaply enough for repeated use.

## Falsifier

If later runtime bugs still progress through tweak notes without a populated first-disagreement section and probe-backed values, the gate is not real.

## Acceptance Criteria

- bug artifacts require first-disagreement, exact values, writer chain, and contradictory-evidence sections
- runtime probes can capture at least one disagreement class named in this packet from the default runtime lane
- progress or closure validators reject runtime bug updates when those sections are missing
- regression artifacts can reference the same probe outputs without inventing a second proof layer

