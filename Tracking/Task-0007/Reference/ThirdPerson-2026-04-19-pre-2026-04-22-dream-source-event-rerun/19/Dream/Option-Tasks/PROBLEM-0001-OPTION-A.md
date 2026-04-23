# Title

Require explicit default-lane proof manifests before regression closure

## Summary

Add a repo-local proof-manifest contract so a regression closeout cannot quietly treat non-default or non-runtime evidence as final proof for the human default runtime lane.

## Source Event IDs

```json
[
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-3325",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-3431",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-3473",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-3618",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-3702",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4909",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5100",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-6451"
]
```

## Burden Being Reduced

The human has to keep reasserting that regression truth lives on the repo's default runtime lane and that cropped, off-lane, or non-runtime evidence does not close a runtime defect.

## Causal Claim

The burden persists because closeout artifacts do not have an enforced proof manifest that declares lane, runtime surface, artifact completeness, and whether a cited artifact is closing proof or supporting diagnosis only.

## Evidence

- [../BURDEN-ANALYSIS.md](../BURDEN-ANALYSIS.md) identifies proof-truth drift as the first burden driver.
- [../SessionExcerpts/INDEX.json](../SessionExcerpts/INDEX.json) keeps the first-session correction sequence and the later runtime-evidence rejections.
- [../../HumanNeeds/PACKET-RECORD.json](../../HumanNeeds/PACKET-RECORD.json) already concludes that default-lane truthful proof is the top packet need.

## Goals

- make lane class, runtime surface, and artifact completeness explicit in every closing regression artifact
- force non-default or non-runtime artifacts to be labeled supporting only
- make a reviewer able to reject weak closure from the manifest alone

## Non-Goals

- building a new runtime capture pipeline
- deciding whether a specific animation is believable by itself
- replacing repo-root regression policy

## Implementation Home

- `C:\Agent\ThirdPerson\REGRESSION.md`
- `C:\Agent\ThirdPerson\TESTING.md`
- repo-local regression artifact templates under `C:\Agent\ThirdPerson\Tracking\Task-<id>\Testing\`
- any repo-local validator or schema used to check regression closeout packets

## Proposed Changes

- add a proof-manifest section or companion JSON that declares:
  - lane class
  - lane label
  - runtime versus non-runtime
  - closing-proof versus supporting-only
  - required visual coverage such as full-body framing when body geometry is in dispute
- update repo-local testing docs so closeout artifacts must carry or reference that manifest
- add a validator that fails or downgrades closeout when the manifest proves the evidence is off-lane or incomplete

## Expected Resolution

A later Task-0006-style closeout can no longer sound complete while silently depending on the wrong lane or on invalid runtime evidence.

## What Does Not Count

- stronger prose with no manifest
- a manifest field that exists but is not checked
- a note that says "supporting only" in prose while downstream closeout still treats it as closure

## Why This Mechanism

This is the smallest durable mechanism that separates closing proof from helpful nearby evidence without requiring a new runtime exporter.

## Human Relief If Successful

The human no longer has to police whether the cited artifact is even from the right lane before evaluating the substantive claim.

## Remaining Uncertainty

This option still depends on humans or scripts filling the manifest honestly after the run; it does not by itself guarantee the right artifacts were produced.

## Falsifier

If a later closeout still passes while citing an automation-only surrogate, a cropped screenshot, or a non-runtime image for a runtime claim, this option is inadequate.

## Acceptance Criteria

- repo-local regression closeout artifacts require a proof manifest or manifest reference before closure
- the manifest distinguishes closing proof from supporting proof and distinguishes runtime from non-runtime artifacts
- validator behavior blocks or downgrades closure when the manifest proves the evidence is off-lane or incomplete
- repo-local docs state that the manifest, not prose convenience, controls closure truth

