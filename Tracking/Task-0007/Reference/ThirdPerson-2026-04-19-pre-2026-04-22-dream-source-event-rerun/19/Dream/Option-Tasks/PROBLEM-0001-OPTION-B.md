# Title

Generate default-lane runtime proof packets from GameAutomation

## Summary

Make the human default runtime lane emit a reusable proof packet that already carries lane metadata, runtime-only evidence, and required framing, then make closeout consume that packet instead of reconstructing proof from prose.

## Source Event IDs

```json
[
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-3325",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-3431",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-3473",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-3618",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-3702",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-3814",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4114",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-4202",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4909",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5100",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-6451"
]
```

## Burden Being Reduced

The human keeps having to invalidate evidence and restate the true proof surface because the runtime lane does not emit a standardized closing packet on its own.

## Causal Claim

The burden persists because the system reconstructs closure from mixed artifacts after the fact instead of generating a lane-bound proof packet at the runtime boundary where the truth is observed.

## Evidence

- [../BURDEN-ANALYSIS.md](../BURDEN-ANALYSIS.md) identifies proof truth drift as the main burden driver.
- [../SessionExcerpts/INDEX.json](../SessionExcerpts/INDEX.json) shows both the early default-lane correction and the later rejection of cropped or non-runtime evidence.
- [../../HumanInterventionTime/SUMMARY.json](../../HumanInterventionTime/SUMMARY.json) shows how much human time the day spent on correction rather than on new work.

## Goals

- emit one reusable proof packet from the default runtime lane
- include lane metadata, runtime-only evidence declarations, artifact completeness markers, and full-body visual captures when body geometry is disputed
- make regression closeout and bug artifacts consume that packet directly

## Non-Goals

- solving the runtime defect itself
- replacing manual exploratory debugging artifacts
- generalizing every repo at once

## Implementation Home

- `C:\Agent\ThirdPerson\Source\GameAutomation\`
- `C:\Agent\ThirdPerson\Content\Script\GameAutomation\`
- `C:\Agent\ThirdPerson\AUTOMATION.md`
- repo-local regression artifacts under `C:\Agent\ThirdPerson\Tracking\Task-<id>\Testing\`
- any repo-local validator that consumes the proof packet

## Proposed Changes

- extend the default-lane runtime automation lane to emit a proof-packet artifact root with:
  - lane identity
  - runtime-surface declaration
  - captured viewpoints
  - framing coverage markers
  - links or paths to the exact images and measurements
- make regression-run and bug artifacts reference the emitted packet instead of restating the evidence informally
- fail closure when the required packet is absent or when the packet shows a non-closing lane

## Expected Resolution

The repo's human default lane becomes the place where closing proof is created, not just a thing later prose claims to have exercised.

## What Does Not Count

- a packet emitted only for some runs while closeout can still proceed without it
- a packet that does not distinguish closing proof from supporting diagnostics
- a runtime packet that still allows cropped geometry disputes to pass without full-body coverage

## Why This Mechanism

This intercepts the failure at the earliest honest boundary: the runtime lane itself. That is stronger than asking later artifacts to label mixed evidence correctly after the fact.

## Human Relief If Successful

The human can inspect one packet from the right lane instead of checking whether scattered artifacts were from the right lane, cropped, or mislabeled.

## Remaining Uncertainty

This option is heavier than a manifest-only gate and depends on the existing GameAutomation lane being stable enough to emit the packet reliably.

## Falsifier

If the proof packet exists but later closeout still requires ad hoc explanation about lane identity or image validity, the packet is not carrying enough truth.

## Acceptance Criteria

- the default runtime lane emits a durable proof packet before closeout can claim success
- the proof packet records lane identity, runtime surface, and required capture set
- regression and bug artifacts consume the packet by reference instead of restating mixed evidence
- closure fails or stays open when the packet is missing, incomplete, or not from the human default lane

