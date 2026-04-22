# Problem 0006 Option A Plan 0001

## Intent

Require a first-disagreement debugging gate: name the exact bad runtime value, then trace the writers upstream before implementing more fixes.

## What Changes

- Debugging cannot move into fix mode until the first concrete bad state is named with values.
- Each later note must preserve the chain from runtime disagreement to upstream writer.
- Symptom-level categories are not enough for closure.

## Files Or Artifact Types That Move

- Shared debugging workflow rules and task prompts.
- Task-local bug artifacts, regression notes, and research records that capture the disagreement seam.
- Repo-local runtime evidence bundles that support value tracing on the default lane.

## Rollout

1. Define the required disagreement record: frame or sample, bad value, expected value, lane, and evidence link.
2. Require a writer-trace section that names each upstream boundary checked next.
3. Add the gate to task workflows before code changes meant to fix a runtime defect.
4. Pilot it on a `ThirdPerson` runtime defect with default-lane evidence only.
5. Treat the debugging record as incomplete until the concrete writer is proven or the remaining branch is tightly bounded.

## Success Check

- Runtime defects are described with values, not only words like "looks wrong."
- Fix attempts cite a proven or tightly bounded writer path.
- Root-cause updates become more stable and less prone to being reopened.

## Burden Reduction Under Directional Context

`Truth`: it keeps the investigation pinned to reality on the actual runtime lane.

`Compassion`: it reduces human steering through repeated symptom-level loops.

`Tolerance`: it turns the human's rough defect report into a precise investigation record without blaming the roughness.
