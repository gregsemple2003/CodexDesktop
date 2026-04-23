# PROBLEM-0001 OPTION-A: ThirdPerson Regression Claim Gate (Default-Lane Proof + Evidence Disqualifiers)

## Title

Make ThirdPerson regression claims fail closed unless they cite the repo-defined default lane and provide non-disqualifying evidence for the disputed surface.

## Summary

The packet shows a repeated failure mode: work is summarized as "regression passed" even when the evidence is from a supporting (non-default) lane or the cited evidence is incomplete for the disputed visual fact. This forces the human to do repeated lane policing and evidence policing, and it causes trust loss around closeout claims.

This winner proposes a narrow first intervention boundary:

- any claim that `TP-REG-001`, `TP-REG-002`, or `TP-REG-003` passed must cite repo-local lane truth from ThirdPerson’s `REGRESSION.md` and a task-owned `REGRESSION-RUN-<NNNN>.md` whose *own contents* block substitute-lane or cropped-evidence overclaims

This is not a proposal to fix the underlying gameplay/animation defects. It is a proposal to make closure claims auditable and fail-closed at the point where wrong claims currently escape.

## Writeup Type

Concrete implementation task (burden-reduction proposal).

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

## Burden Being Reduced

The human is currently forced to repeatedly do work the system should be preventing:

- Lane policing: restating that ThirdPerson regression proof must use the repo’s human default lane, and that a supporting lane does not substitute as regression proof.
- Evidence policing: inspecting whether a cited screenshot or artifact actually shows the disputed surface fully (for example, feet contact/hover, visible pawn) rather than allowing cropped or partial artifacts to stand as proof.
- Closure policing: stopping a claim from being summarized as "regression passed" when it is at best "supporting evidence gathered" or "partial/blocked run."

The deeper burden is exported trust reconstruction: a reviewer cannot rely on the task’s stated regression outcome without re-running or re-deriving what lane was actually exercised and whether the evidence actually covers the claim.

## Current Truth

Current durable truths already exist, but are not enforced as a fail-closed gate:

- ThirdPerson defines its human default lane and explicitly states that supporting lanes do not substitute for that lane in [ThirdPerson/REGRESSION.md](/c:/Agent/ThirdPerson/REGRESSION.md).
- The shared testing contract already requires `REGRESSION-RUN-<NNNN>.md` artifacts to contain a `Regression Claim` section with `Claimed lane`, `Actual flow exercised`, `Why this counts`, and `Disqualifiers / limitations` (see the shared [../../../../../../../../Processes/TESTING.md](../../../../../../../../Processes/TESTING.md)).

Despite this, the packet shows the human still repeatedly had to correct:

- substitute-lane proof being treated as regression proof
- evidence artifacts being treated as proving a disputed visual fact even when the artifact itself was incomplete or misleading

## Target Truth

A future ThirdPerson "regression passed" claim should be auditable from two durable anchors:

1. The repo’s lane truth in `REGRESSION.md` (what lane and case ids are actually required).
2. A task-owned `REGRESSION-RUN-<NNNN>.md` whose own content makes it impossible to honestly overclaim:
   - it names the claimed lane and case ids
   - it describes the actual flow exercised
   - it explains why the run counts for the claimed lane (or explicitly says it is supporting-only)
   - it links evidence artifacts that cover the disputed surface
   - it lists explicit disqualifiers that downgrade the claim

Supporting-lane or partial evidence can still exist as debugging support, but it should not be representable as "passed default-lane regression" at closeout.

## Causal Claim

If ThirdPerson’s `REGRESSION.md` and the shared `REGRESSION-RUN` proof contract explicitly define substitute-lane and incomplete-evidence disqualifiers, and the ThirdPerson repo requires regression claims to cite those two anchors, then:

- a future implementer cannot honestly summarize "TP-REG-001..003 passed" while only providing Lane C evidence or cropped/partial visual proof
- the human’s repeated lane/evidence policing burden will decrease because the claim artifact fails closed before it reaches an approval or closeout surface

## Evidence

The evidence for this burden and mechanism boundary is concentrated in the default-lane truth and runtime evidence honesty clusters from [../BURDEN-ANALYSIS.md](../BURDEN-ANALYSIS.md) (`BD-001`) and the packet’s stable `event_id` entries.

Representative evidence surfaces (see the `preview` text in [../../HumanInputEvents/INDEX.json](../../HumanInputEvents/INDEX.json)):

- The human explicitly corrects the substitute-lane failure mode and demands short, direct agreement about the default lane rule (e.g. `...-3431`, `...-3473`).
- The human provides concrete findings tying the default lane to the broken asset pair while repaired assets lived on an automation-only lane (`...-3814`).
- The human rejects proof that does not meet the visible, human-facing bar (e.g. claims that omit the disputed surface or use proof-only views) and forces continued work (`...-4114`, `...-4432`).

## Why This Mechanism

This proposal chooses the regression-claim boundary as the first enforcement point because that is where the burden is paid:

- The human cost is not just "bad evidence exists." The cost is "bad evidence is allowed to become a closure claim."
- A fail-closed claim gate prevents exported review labor by making the artifact itself refuse to represent a substitute lane or disqualified evidence as passed regression.

This mechanism is deliberately narrower than building new tooling. It prefers durable contracts and artifact shape over a new validator or generator surface that the packet did not explicitly justify.

## Scope Rationale

This is intentionally scoped to ThirdPerson’s repo-local regression claim boundary, while reusing the shared proof-quality contract fields where they already exist.

- Narrower (task-only) scope is insufficient: a single task-specific reminder does not protect the next ThirdPerson task from repeating the same substitute-lane overclaim.
- Broader (fully shared, UE-generic) scope is rejected by durable packet evidence: the human explicitly objected to moving ThirdPerson-specific lane semantics into generic shared orchestration prose.

## Goals

- Make it impossible to honestly overclaim passed ThirdPerson regression cases using supporting-only lanes.
- Make it impossible to honestly overclaim a disputed visual fix using evidence that does not actually cover the disputed surface.
- Reduce repeated human lane policing and evidence policing by moving the "no, that does not count" logic into the claim artifact boundary.

## Non-Goals

- Fixing the underlying ThirdPerson defects (invisible pawn, animation defects, foot hover/tilt) directly.
- Building a standalone proof-pack linter or validator in this first slice.
- Redefining ThirdPerson’s regression lanes; lane truth stays owned by the repo’s `REGRESSION.md`.

## Implementation Home

- Repo-local lane truth home:
  - [ThirdPerson/REGRESSION.md](/c:/Agent/ThirdPerson/REGRESSION.md)
- Shared proof-quality contract home:
  - [../../../../../../../../Processes/TESTING.md](../../../../../../../../Processes/TESTING.md)
- Task-owned executed proof artifact home:
  - `Tracking/Task-<id>/Testing/REGRESSION-RUN-<NNNN>.md`

## Implementation Home Rationale

- Repo-local `REGRESSION.md` is the only honest home for "what lane is required" and "what case ids mean passing" (by the shared testing rule).
- The shared `Processes/TESTING.md` is the right place to define the generic proof-quality claim fields, because those fields should exist regardless of repo content.
- The executed `REGRESSION-RUN-<NNNN>.md` artifact is the right enforcement boundary because it is the durable, reviewable object that later closure claims cite.

## Constraints And Baseline

- Preserve ThirdPerson’s canonical lane definitions and disqualifiers as already written in `REGRESSION.md`.
- Do not move ThirdPerson-specific lane semantics into generic shared orchestration prose.
- Keep this slice contract-first: make the claim auditable and fail-closed before adding any new automation or linting tools.

## Proposed Changes

These are the concrete, reviewable surfaces this winner changes.

1. **Repo-local: strengthen ThirdPerson regression claim gate language**
   - File: [ThirdPerson/REGRESSION.md](/c:/Agent/ThirdPerson/REGRESSION.md)
   - Add a new section (or extend an existing one) that explicitly requires:
     - any claim of `TP-REG-001`, `TP-REG-002`, or `TP-REG-003` "PASS" must link to a task-owned `Tracking/Task-<id>/Testing/REGRESSION-RUN-<NNNN>.md`
     - the claim must name the *claimed lane* (Lane A/B/C) as defined in `REGRESSION.md`
     - any non-default lane (`Lane C`) is labeled **supporting-only** unless `REGRESSION.md` explicitly declares it the human default (it currently does not)
     - for disputed visual claims, the evidence link(s) must be described as covering the disputed surface (or the run must mark the claim disqualified)
2. **Shared: expand the proof-quality contract to include evidence-coverage disqualifiers**
   - File: [../../../../../../../../Processes/TESTING.md](../../../../../../../../Processes/TESTING.md)
   - Extend the `Regression Claim Requirement` section to add an explicit sub-requirement:
     - `Evidence coverage` (or equivalent): when the claim is about a disputed visual surface, the run must state whether the evidence artifact actually shows the disputed surface fully enough to support the claim
     - `Disqualifiers` must explicitly include "supporting-only lane" and "incomplete/cropped evidence for disputed surface" as examples of reasons a run cannot support a clean pass claim
3. **Task-owned artifact convention: require a stable section shape inside `REGRESSION-RUN-<NNNN>.md`**
   - Artifact: `Tracking/Task-<id>/Testing/REGRESSION-RUN-<NNNN>.md`
   - Require a stable set of headings/fields inside the `Regression Claim` section:
     - `Claimed lane / case ids` (explicitly naming `TP-REG-001..003` when applicable)
     - `Actual flow exercised` (explicitly stating whether it was "launch project normally + PIE Selected Viewport" versus "Lane C headless GameAutomation" or other)
     - `Why this counts` (explicitly tying back to the repo’s lane truth in `REGRESSION.md`)
     - `Evidence coverage for disputed surfaces` (explicitly stating what artifact shows the disputed surface)
     - `Disqualifiers / limitations` (explicitly stating why this run is supporting-only, partial, or disqualified if applicable)

## Acceptance Criteria

- A future ThirdPerson task cannot honestly claim `TP-REG-001`, `TP-REG-002`, or `TP-REG-003` passed while citing only a supporting lane (e.g. Lane C) without also passing the repo-defined human default lane.
- A future ThirdPerson task cannot honestly claim a disputed visual fix is proven if the cited evidence artifact is incomplete for the disputed surface; the `REGRESSION-RUN` must mark it disqualified or partial instead.
- A cold reviewer can follow two links (repo `REGRESSION.md` + task `REGRESSION-RUN-<NNNN>.md`) and answer:
  - what lane was claimed
  - what lane was actually exercised
  - why it counts (or why it does not)
  - what evidence supports the disputed claim (or why it is disqualified)

## Expected Resolution

Human-facing outcome:

- Approval and closeout reviews become faster because the claim surface itself contains the "what counts" and "what does not count" information the human currently has to reconstruct by hand.
- Trust increases: "regression passed" becomes a more reliable statement because substitute-lane and disqualified-evidence claims fail closed.

## Human Relief If Successful

- Fewer repeated prompts demanding "default lane only" and "short answers" about whether evidence counts.
- Fewer reopening cycles caused by an overclaimed proof artifact.
- Lower cognitive load at approval time: the reviewer does not have to reconstruct lane identity or inspect whether the evidence was cropped just to know whether the claim is valid.

## Internal Mechanism Map

1. Repo-local lane truth stays authoritative in ThirdPerson `REGRESSION.md`.
2. Shared testing contract defines the minimum proof-quality fields for any `REGRESSION-RUN`.
3. ThirdPerson regression claim gate requires that any "PASS" claim for `TP-REG-001..003` cite the `REGRESSION-RUN` and explicitly mark supporting-only evidence as supporting-only.
4. The `REGRESSION-RUN` artifact’s own disqualifiers make overclaiming non-viable.

## Rival Explanations Considered

- "This was just one bad run and won’t recur."
  - Rejected: the packet contains repeated corrections of the same substitute-lane and evidence-coverage failure shape.
- "The human is being overly strict about visuals."
  - Rejected: the repo’s regression lane explicitly targets the human-facing playable surface (visible, controllable pawn); evidence coverage is the ordinary requirement for proving a disputed visual.

## Rival Mechanisms Considered

- `P-001 / Option B` (shared proof-claim contract only):
  - Kept as a seam: the generic proof-quality fields belong in shared testing docs.
  - Rejected as the sole mechanism: it cannot carry repo-specific lane semantics and disqualifiers without reintroducing the UE-generic prose problem the human rejected.
- Proof-pack linting / automation-only validation:
  - Deferred follow-on: helpful later, but the packet does not justify freezing an exact shared tooling home for it in this first slice.

## Tradeoffs

- Increased ceremony in `REGRESSION-RUN` artifacts (more explicit fields).
  - Tradeoff is intentional: it buys auditability and removes exported reviewer labor.
- Some runs will become "PARTIAL PASS" more often instead of being summarized as "pass."
  - This is a feature: it keeps truth and avoids proxy closure.

## Shared Substrate

- Shared testing model and naming: [../../../../../../../../Processes/TESTING.md](../../../../../../../../Processes/TESTING.md)
- ThirdPerson repo regression matrix and lane definitions: [ThirdPerson/REGRESSION.md](/c:/Agent/ThirdPerson/REGRESSION.md)
- Executed regression-run exemplar shape (supporting reference): [../../../../../../../../Exemplars/REGRESSION-RUN-0001.md](../../../../../../../../Exemplars/REGRESSION-RUN-0001.md)

## Not Solved Here

- Fixing the underlying ThirdPerson defects being measured by regression.
- Making Lane C (automation) sufficient as the human default lane; ThirdPerson explicitly treats it as supplemental today.
- Adding an automated evidence verifier, screenshot linter, or video-based proof generator.

## What Does Not Count

- "The mesh exists in Content Browser" without a real default-lane PIE run.
- A proof-view-only screenshot that is not the human-facing playable surface.
- A supporting-lane run that is treated as satisfying Lane A without explicit justification rooted in `REGRESSION.md`.
- A cropped screenshot that omits the disputed surface while still being cited as proof of that surface.

## Remaining Uncertainty

- The exact minimum "evidence coverage" wording that is strict enough to prevent overclaiming without forcing excessive duplication across repos.
- Whether ThirdPerson wants a canonical "disputed visual checklist" subsection inside `REGRESSION-RUN` artifacts, or whether a single `Evidence coverage` field is sufficient.

## Falsifier

This proposal is falsified if, after implementation, any of these remain true in durable artifacts:

- A ThirdPerson task can still cite only a supporting lane (Lane C or other) and successfully present `TP-REG-001..003` as passed regression without also passing the human default lane.
- A disputed visual claim can still survive review after the cited evidence is shown to be incomplete for the disputed surface.

## Proof Plan

1. Update the shared and repo-local docs in the proposed homes.
2. Create two fixture `REGRESSION-RUN` examples under a real implementing task folder (to prove the gate is implementable in practice):
   - a "bad" fixture that attempts to claim `TP-REG-003 PASS` using Lane C only and/or incomplete evidence, and is forced by the template/contract to mark itself supporting-only or disqualified
   - a "good" fixture that claims Lane A and cites evidence that covers the disputed surface
3. Run a real ThirdPerson regression slice and write a new `REGRESSION-RUN-<NNNN>.md` using the new shape.
4. Verify that a reviewer can reject an invalid claim without needing to reconstruct lane semantics or hunt for missing evidence context.

## Open Questions

- Should ThirdPerson require a screenshot (or equivalent) for each of `TP-REG-001..003` when the case is visual, or only when a disputed visual claim exists?
- Should the repo require a short "supporting-lane evidence index" section when Lane C is used for debugging, to prevent it from being conflated with regression proof later?

## References

- Burden driver: [../BURDEN-ANALYSIS.md](../BURDEN-ANALYSIS.md) (`BD-001`)
- Designed options: [../SOLUTION-DESIGN.md](../SOLUTION-DESIGN.md#p-001-default-lane-proof-and-evidence-gate)
- Frozen winner boundary: [../WINNER-SYNTHESIS.md](../WINNER-SYNTHESIS.md#w-001-thirdperson-regression-claim-gate)
- Final matrix row: [../ORTHOGONAL-SOLUTIONS-MATRIX.md](../ORTHOGONAL-SOLUTIONS-MATRIX.md#p-001-default-lane-proof-and-evidence-gate)
- ThirdPerson regression lanes: [ThirdPerson/REGRESSION.md](/c:/Agent/ThirdPerson/REGRESSION.md)
- Shared testing contract: [../../../../../../../../Processes/TESTING.md](../../../../../../../../Processes/TESTING.md)

