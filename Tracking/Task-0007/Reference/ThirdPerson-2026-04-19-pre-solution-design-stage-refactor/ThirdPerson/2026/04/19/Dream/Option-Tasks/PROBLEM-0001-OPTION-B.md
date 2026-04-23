# PROBLEM-0001 OPTION-B: Unattended Default-Lane Runner With Evidence Pack

## Title

Build an unattended runner that executes the repo-defined human default lane and emits a durable evidence pack by construction.

## Summary

The packet repeatedly shows that "we ran regression" is not trustworthy unless it is anchored to the exact default lane and produces reviewable artifacts. This task proposes building a repo-local unattended runner for the default lane, so evidence is not dependent on manual screenshots or surrogate lanes.

Writeup type: research -> concrete implementation (two-stage).

## Goals

- Make default-lane proof reproducible and durable.
- Reduce human time spent re-running the default lane manually just to validate claims.
- Emit an evidence pack that is reviewable and includes contested visuals fully in-frame.

## Non-Goals

- Changing engine source under `C:\\Agent\\UnrealEngine` (explicitly disallowed by packet constraint).
- Replacing the repo's canonical `REGRESSION.md` or redefining lanes.

## Implementation Home

- Repo-local tooling in `C:\\Agent\\ThirdPerson` (under repo-owned scripts and automation surfaces).
- Proof artifacts under `Tracking/Task-<id>/Testing/` (or repo-documented equivalent).

## Proposed Changes

Stage 1 (research):

- Confirm the minimal Unreal API surface that can drive:
  - launching PIE in the correct mode (Selected Viewport on the default pawn path)
  - starting/stopping play
  - capturing durable screenshots (and optionally short video)
  - capturing runtime logs/telemetry for the default pawn
- Validate whether the engine-exposed approach referenced in the packet (`ULevelEditorSubsystem::EditorRequestBeginPlay`) can be used without engine modifications.

Stage 2 (concrete implementation):

- Implement a single entrypoint script or command that:
  - runs the default lane
  - captures evidence with contested surfaces fully visible (feet, ground contact)
  - writes a deterministic evidence pack folder
  - writes or updates a `REGRESSION-RUN-<NNNN>.md` that points at the evidence pack

## Acceptance Criteria

- Stage 1: a short research writeup exists naming the chosen API seam, limitations, and the exact repo-local implementation plan.
- Stage 2: running the entrypoint produces:
  - a durable evidence pack folder
  - at least one runtime screenshot with feet fully visible in-frame (when feet are disputed)
  - a `REGRESSION-RUN-<NNNN>.md` that passes the proof-quality gate from PROBLEM-0001 OPTION-A
- A human can review the evidence pack and confirm the lane, without rerunning PIE manually.

## Source Event IDs

```json
[
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-3325",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-3649",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-3431",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-3473",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-3618",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-3702",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-3814",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-4202",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4114",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-4432",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4909",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-6451",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5040",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5050",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5100"
]
```

## Burden Reduced

Manual re-running and evidence disputes due to the system being unable to reliably produce default-lane proof artifacts.

## Causal Claim

If default-lane proof is produced by an unattended runner that *only* runs the repo-defined lane and writes a durable evidence pack, then the proof surface will converge and false closure due to lane drift will drop.

## Evidence

- `EXCERPT-0001` in [`../SessionExcerpts/INDEX.json`](../SessionExcerpts/INDEX.json) references using editor PIE automation while keeping proof on the human default lane.
- `EXCERPT-0002` shows evidence-quality disputes (cropped screenshots) that an unattended runner can avoid via consistent framing.

## Why This Mechanism

A proof gate (OPTION-A) prevents false closure but does not generate proof. This option builds the missing production capability for default-lane evidence.

## Human Relief If Successful

- Default lane can be re-proven quickly without human intervention.
- Evidence disputes become rarer because framing and artifact capture are consistent.

## Remaining Uncertainty

- Unreal editor automation may not be fully equivalent to manual PIE in Selected Viewport in all cases; this must be validated against the repo's `REGRESSION.md` definitions.

## Falsifier

- The unattended runner cannot faithfully reproduce the repo-defined default lane, or its artifacts still fail review (missing disputed surfaces, wrong camera path).

## What Does Not Count

- A runner that executes a non-default map or pawn path.
- A runner that depends on engine-source edits under `C:\\Agent\\UnrealEngine`.

## References

- `BD-001` in [`../BURDEN-ANALYSIS.md`](../BURDEN-ANALYSIS.md)
- `PROBLEM-0001` in [`../ORTHOGONAL-SOLUTIONS-MATRIX.md`](../ORTHOGONAL-SOLUTIONS-MATRIX.md)

