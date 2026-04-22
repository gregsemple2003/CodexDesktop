# Solution Task 0001

## Title

Enforce proof-surface checks before any task claims regression closure

## Summary

April 19 ThirdPerson work repeatedly claimed progress from supporting proof that did not satisfy the repo's human default-lane regression bar.
This task adds a shared proof-surface gate so closeout text must name the claimed lane and cannot label off-lane proof as regression closure.

## Goals

- require regression artifacts to name the claimed proof surface
- compare the claimed surface against the repo's regression rules before closeout
- force off-lane proof to stay labeled as supporting proof only

## Non-Goals

- changing any repo's regression definition
- requiring video capture for every repo
- replacing repo-local `REGRESSION.md`

## Implementation Home

- `C:\Users\gregs\.codex\Orchestration\Processes\TESTING.md`
- `C:\Users\gregs\.codex\Orchestration\Processes\DEBUGGING.md` if shared wording needs the same proof-surface terms
- shared prompt or generator files under `C:\Users\gregs\.codex\Orchestration\Prompts\`
- any shared validator or artifact-emission code that writes regression closeout text

## Constraints And Baseline

- repo-root `REGRESSION.md` stays the canonical source for repo regression lanes
- supporting diagnostics and operator lanes may remain useful, but they must not close a human default-lane regression claim
- the task must work with repo-local docs from the packet repo, not the current workspace repo

## Proposed Changes

- add a required `Claimed Proof Surface` block to shared regression-run guidance with fields for lane, map, game mode, pawn, evidence type, and whether the run counts as regression or supporting proof
- add a shared check that compares the claimed block against repo-local regression rules before a closeout or handoff can say `passed regression`
- add explicit output text for mismatch cases such as `supporting proof only` and `off-lane for closure`
- add one ThirdPerson-style example that shows why headless or operator-lane evidence cannot stand in for default-lane closure

## Expected Resolution

A human reading a regression claim can tell exactly which lane was exercised and whether that lane actually satisfies the repo's regression bar.
Wrong-lane proof can still be kept, but it will stop masquerading as closure.

## What Does Not Count

- stronger prose that still allows off-lane proof to be called regression
- a reminder with no required proof-surface fields
- storing screenshots without saying what lane they came from

## Acceptance Criteria

- shared testing guidance defines the required `Claimed Proof Surface` block and its fields
- the closeout flow has a durable check that can mark a run `supporting proof only` when the claimed lane does not satisfy repo-local regression rules
- shared prompts or generators stop emitting `passed regression` when the claimed surface is off-lane
- the new examples include a ThirdPerson-style human default-lane case and an off-lane supporting-proof case

## Proof Plan

- run the gate against the April 19 ThirdPerson evidence shape where supporting proof existed but closure was false
- confirm the output labels the off-lane proof as supporting only
- confirm a matching default-lane run shape can still pass

## References

- `../ORTHOGONAL-SOLUTIONS-MATRIX.md`
- `../Plans/PROBLEM-0001-OPTION-B-PLAN-0001.md`
- `../../REGRESSION.md`
- `../../TESTING.md`

## Plan Addendum

Define a small proof-surface record first.
It should capture lane, map, game mode, pawn, and evidence type.
Then wire one shared check into the closeout path so regression claims are compared against repo-local lane rules before they are emitted.
If the check finds a mismatch, the workflow should preserve the evidence but label it as supporting proof only.
