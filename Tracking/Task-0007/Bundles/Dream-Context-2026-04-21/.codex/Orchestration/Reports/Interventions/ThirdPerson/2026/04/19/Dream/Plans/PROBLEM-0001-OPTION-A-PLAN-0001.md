# Problem 0001 Option A Plan 0001

## Intent

Add a pre-closure proof-surface gate that blocks regression or closure claims when the planned evidence does not match the repo's human default lane.

## What Changes

- Add a check that reads repo-local proof rules before any closure or regression claim is emitted.
- Require each proof attempt to declare lane, runtime surface, and evidence limits.
- Fail the closeout path when the evidence is from a non-default lane, non-runtime surface, or supporting-only artifact set.

## Files Or Artifact Types That Move

- Shared orchestration process or prompt logic that decides whether closure can be claimed.
- Task or pass closeout artifacts that record proof lane, proof type, and gate result.
- Repo-local rule readers that pull the proof bar from `C:\Agent\ThirdPerson\REGRESSION.md` and `C:\Agent\ThirdPerson\TESTING.md`.

## Rollout

1. Define a small proof-surface record with fields for lane, runtime status, evidence roots, and closure eligibility.
2. Build a repo-local rule reader that classifies the `ThirdPerson` human default lane and marks supporting-only lanes.
3. Insert the gate before milestone closeout, regression claims, and approval asks.
4. Pilot it on one `ThirdPerson` task flow that already has mixed proof surfaces.
5. Promote the rule to the shared orchestration layer only after the repo-local behavior is stable.

## Success Check

- A closeout attempt using non-default or non-runtime proof is blocked before the human has to reject it.
- The resulting artifact says exactly why the proof did not count.
- A valid default-lane runtime proof path still passes without extra human restatement.
- The gate works without engine changes.

## Burden Reduction Under Directional Context

`Truth`: it stops tidy but false closure on the wrong lane.

`Compassion`: it removes reopenings, proof rejection, and rediscovery work from the human.

`Tolerance`: it handles normal proof ambiguity with a clear failure reason instead of making the human restate the lane rule.
