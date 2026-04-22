# Solution Task 0001

## Title

Enforce a default-lane proof-surface gate before regression or closure claims

## Summary

The packet's highest-cost loop was false closure on the wrong proof surface.
This task adds one shared proof-surface gate so orchestration cannot claim regression, close a pass, or ask for approval with non-runtime or non-default-lane evidence when the repo requires human-default-lane proof.

## Goals

- Stop closure claims that rely on headless diagnostics, alternate lanes, or unlabeled screenshots when `C:\Agent\ThirdPerson\REGRESSION.md` requires the human default lane.
- Make the proof lane, runtime surface, and evidence limits explicit in pass-closeout state.
- Give the human a clear failure reason before bad proof reaches an approval or closure ask.

## Non-Goals

- Building new runtime capture tooling.
- Rewriting the `ThirdPerson` regression matrix.
- Treating supporting diagnostics as regression proof.

## Constraints And Baseline

Current truth:

- `C:\Agent\ThirdPerson\REGRESSION.md` says regression proof must pass the human default lane.
- `C:\Agent\ThirdPerson\TESTING.md` allows headless diagnostics only as supporting proof.
- The packet shows repeated reopenings because easier proof surfaces were accepted first and rejected later.

Hard constraints:

- use repo-local proof rules from the packet repo, not the current workspace repo
- keep `no engine mods` intact
- preserve valid human-default-lane runtime proof paths without adding a human restatement step

## Implementation Home

- `C:\Users\gregs\.codex\Orchestration\Processes\TESTING.md`
- `C:\Users\gregs\.codex\Orchestration\Prompts\REGRESSION-LEADER.md`
- `C:\Users\gregs\.codex\Orchestration\Prompts\REGRESSION-TESTER.md`
- `C:\Users\gregs\.codex\Orchestration\Prompts\IMPLEMENTATION-LEADER.md`
- `C:\Users\gregs\.codex\Orchestration\PASS-CHECKLIST.md`
- `C:\Users\gregs\.codex\Orchestration\PASS-CHECKLIST.schema.json`
- pilot proof artifacts under `C:\Agent\ThirdPerson\Tracking/Task-<id>/Testing\`

## Proposed Changes

- Add one required `proof_surface` record to pass-closeout state with these fields:
  - `repo_rule_source`
  - `claimed_lane`
  - `runtime_surface`
  - `evidence_roots`
  - `supporting_only_artifacts`
  - `closure_eligible`
  - `rejection_reason`
- Update the shared testing process and regression prompts so they must read the packet repo's regression rules before claiming closure.
- Require implementation-leader approval asks to surface the gate result instead of hiding it in narrative text.
- Pilot the gate on a `ThirdPerson` task flow with both human-default-lane and supporting-only evidence in the tree.

## Expected Resolution

A human reviewing a `ThirdPerson` closeout can see, in one place, whether the claim used the real default lane or only supporting proof.
If the evidence is off-lane or non-runtime, the system stops itself before the human has to reopen the task.

## What Does Not Count

- Stronger wording with no explicit gate result.
- A screenshot bundle that does not say which lane produced it.
- Headless Unreal Python output labeled as regression.
- An approval ask that links files but omits the proof-surface record.

## Acceptance Criteria

- `PASS-CHECKLIST.md` and `PASS-CHECKLIST.schema.json` define and validate the required `proof_surface` record and `closure_eligible` outcome.
- `TESTING.md`, `REGRESSION-LEADER.md`, and `REGRESSION-TESTER.md` require the gate before regression claims or pass closure.
- `IMPLEMENTATION-LEADER.md` requires approval asks to show the gate result and rejection reason when `closure_eligible=false`.
- On a `ThirdPerson` pilot, a closeout attempt that uses only headless diagnostics or another non-default lane is blocked with an explicit reason that cites the repo-local proof bar.
- On the same pilot, a Lane A human-default-lane runtime proof path can pass the gate without extra human clarification.

## Proof Plan

- Run one `ThirdPerson` pilot with mixed proof surfaces and record one blocked closeout.
- Run one `ThirdPerson` pilot with valid Lane A runtime proof and record one passing closeout.
- Verify the same gate result appears in the checklist state and the human-facing approval ask.

## References

- `..\BURDEN-ANALYSIS.md`
- `..\ORTHOGONAL-SOLUTIONS-MATRIX.md`
- `..\Plans\PROBLEM-0001-OPTION-A-PLAN-0001.md`
- `C:\Agent\ThirdPerson\REGRESSION.md`
- `C:\Agent\ThirdPerson\TESTING.md`

## Plan Addendum

Chosen plan: add a pre-closure proof-surface gate that blocks regression or closure claims when the planned evidence does not match the repo's human default lane.

Implementation notes from the selected plan:

- add a check that reads repo-local proof rules before any closure or regression claim is emitted
- require each proof attempt to declare lane, runtime surface, and evidence limits
- fail the closeout path when the evidence is from a non-default lane, non-runtime surface, or supporting-only artifact set

Rollout:

1. Define a small proof-surface record with fields for lane, runtime status, evidence roots, and closure eligibility.
2. Build a repo-local rule reader that classifies the `ThirdPerson` human default lane and marks supporting-only lanes.
3. Insert the gate before milestone closeout, regression claims, and approval asks.
4. Pilot it on one `ThirdPerson` task flow that already has mixed proof surfaces.
5. Promote the rule to the shared orchestration layer only after the repo-local behavior is stable.
