# Problem 0006 Option B Plan 0001

## Intent

Build a reusable runtime probe kit that captures foot, pelvis, mesh, and capsule values on the default lane.

## What Changes

- Add repo-local probes for the main runtime state that humans keep questioning.
- Save machine-readable values beside runtime screenshots and logs.
- Use the probe kit to accelerate later disagreement tracing.

## Files Or Artifact Types That Move

- Repo-local probe scripts or instrumentation under `C:\Agent\ThirdPerson`.
- Runtime evidence bundles with numeric captures.
- Bug or regression artifacts that cite the probe outputs.

## Rollout

1. Choose the minimum useful state set for `ThirdPerson`: foot world Z, mesh offset, capsule bottom, and similar runtime values.
2. Capture those values on the default lane only, with no engine mods.
3. Emit the probe outputs beside the usual runtime proof artifacts.
4. Test the kit on one failing run and one corrected run.
5. Fold the probe kit into future default-lane investigations.

## Success Check

- Investigators can inspect runtime values without ad hoc instrumentation each time.
- Probe outputs line up with the screenshot or video evidence from the same run.
- The kit works with repo-owned code and public APIs only.

## Burden Reduction Under Directional Context

`Truth`: it makes hidden runtime state easier to inspect.

`Compassion`: it can shorten the path from human complaint to machine-checkable evidence.

`Tolerance`: it helps translate a rough visual complaint into precise state capture.
