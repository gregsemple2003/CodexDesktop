# Problem 0006 Option C Plan 0001

## Intent

Always compare the failing default lane against a known-good baseline before deeper root-cause work.

## What Changes

- Add a baseline comparison step early in the investigation.
- Use a known-good template or approved prior lane to show what normal looks like.
- Focus later tracing on the state that diverges first.

## Files Or Artifact Types That Move

- Baseline run artifacts and comparison notes.
- Task research or bug records that cite the first observed divergence.
- Small repo-local helpers that load or compare baseline evidence.

## Rollout

1. Choose the valid baseline for `ThirdPerson` under the no-engine-mods rule.
2. Capture the same evidence set on both failing and baseline runs.
3. Compare the two runs to find the earliest visible or numeric divergence.
4. Feed that divergence into a deeper root-cause investigation.

## Success Check

- Investigations start with a clear picture of what "good" looks like.
- The first divergence is easier to identify than in a single-run analysis.
- Baseline use stays supportive and does not replace writer tracing.

## Burden Reduction Under Directional Context

`Truth`: it reduces confusion about expected runtime behavior.

`Compassion`: it can shorten the search space before the human needs to intervene again.

`Tolerance`: it helps the system ground vague human observations against a stable comparison target.
