# Problem 0001 Option B Plan 0001

## Intent

Expand unattended default-lane runtime capture so real in-lane evidence is easier to gather than proxy evidence.

## What Changes

- Add a reusable runtime-capture path for the `ThirdPerson` default lane.
- Standardize artifact capture for viewport images, lane metadata, and default-lane run summaries.
- Make default-lane proof the path of least resistance for future task work.

## Files Or Artifact Types That Move

- Repo-local automation scripts under `C:\Agent\ThirdPerson\Tracking\Task-*\Research\` or a shared repo-owned automation area.
- Default-lane proof artifacts under task `Testing/` or `Research/`.
- Small wrapper docs or templates that label captures as human-default-lane runtime evidence.

## Rollout

1. Identify the current unattended surfaces that can launch the default lane without engine mods.
2. Package the working launch, capture, and exit sequence into one repo-local runner.
3. Emit a fixed artifact bundle for every run: lane metadata, runtime screenshots, and pass/fail notes.
4. Test the runner against at least one known-good and one known-bad lane outcome.
5. Make the runner the default evidence path for later runtime regressions.

## Success Check

- Engineers can gather default-lane runtime evidence without ad hoc proof setup.
- The artifact set is strong enough that non-runtime images are no longer the easiest available evidence.
- The path stays repo-local and uses public Unreal surfaces only.

## Burden Reduction Under Directional Context

`Truth`: more real runtime evidence reduces substitution pressure.

`Compassion`: the human spends less time asking for re-runs on the real lane.

`Tolerance`: the system can satisfy the lane bar even when the human only says "prove it on the default lane."
