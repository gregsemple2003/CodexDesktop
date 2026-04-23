# PROBLEM-0001 OPTION-A: Shared Proof Contract And Closure Gate For Default-Lane Claims

## Title

Enforce a shared proof contract for default-lane regression claims so wrong-lane or visually invalid proof cannot be closed as clean regression.

## Summary

The burden packet shows repeated human correction of two linked failures:

- supporting or non-default lanes were treated as if they satisfied the repo-defined human default lane
- cropped or incomplete evidence was treated as if it proved the disputed visual fact

This rewrite narrows Option A to the exact first slice justified by the local sources:

- strengthen the shared `REGRESSION-RUN-<NNNN>.md` proof contract
- enforce that contract in the shared regression and closure workflow

This slice does **not** include a standalone linter or other shared tooling, because the local sources do not name an exact shared tooling home for that mechanism honestly.

## Writeup Type

Concrete implementation task.

The durable local sources already justify the failure class, the shared artifact boundary, and the exact prompt surfaces that own regression and closure claims.

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

The human is currently forced to do three kinds of repeated proof-repair work by hand:

1. Lane policing.
   The human must restate that regression proof has to come from the repo-defined human default lane rather than a supporting lane.
2. Visual-evidence policing.
   The human must inspect whether the cited screenshots or artifacts actually show the disputed surface fully enough to support the claim.
3. Closure policing.
   The human must stop malformed proof from being summarized as a clean regression pass or clean closeout.

The deeper burden is trust loss. A reviewer cannot rely on the current artifact boundary without reconstructing what lane actually ran and whether the evidence really proves the claimed visual outcome.

## Current Truth

The shared workflow already contains pieces of the right boundary:

- `Processes/TESTING.md` already requires `Claimed lane`, `Actual flow exercised`, `Why this counts`, and `Disqualifiers / limitations`
- `ORCHESTRATION.md` already requires honest closure preflight for regression claims
- `TASK-LEADER.md` already says a `REGRESSION-RUN-<NNNN>.md` does not count as closure evidence unless it names the claimed lane, actual flow, why it counts, and disqualifiers

But the burden packet shows that this is still not enough in practice.

The current system truth is therefore not merely `missing screenshots` or `weak wording`.
The current truth is that the proof artifact and closure boundary still allow claims such as:

- default-lane proof even though the actual proof came from a supporting or alternate lane
- visual-fix proof even though the disputed surface is cropped, absent, or otherwise not fully shown
- closure-ready proof even though the artifact itself discloses limitations that should downgrade the claim

The current fallback truth is only that the human can manually reopen the work and restate what does not count.
That is a rescue path, not a durable gate.

## Target Truth

A future regression or closure reviewer should be able to read one `REGRESSION-RUN-<NNNN>.md` artifact and answer, without reconstructing raw logs:

- what repo-root lane is being claimed
- what real flow was actually exercised
- why that flow counts for the claimed lane instead of merely supporting it
- what exact evidence artifacts support the claim
- for a disputed visual claim, whether the contested surface is actually shown fully enough to support the claim
- what disqualifiers or limitations remain
- whether the run is honestly closure-ready, only supporting, partial, blocked, or insufficient

A wrong-lane run or a run with cropped or missing contested-surface evidence may still exist as diagnostic evidence, but it should not be representable as a clean passed default-lane regression claim.

## Causal Claim

If the shared regression-run contract is strengthened and the shared regression and closure prompts are required to treat wrong-lane or visually disqualified evidence as result-affecting, malformed proof will fail before it becomes repeated human review burden.

The cause being addressed is not only missing reminders.
The cause is that the current claim boundary still allows incomplete or wrong-lane proof to travel too far toward closure.

## Evidence

`BD-001` in [`../BURDEN-ANALYSIS.md`](../BURDEN-ANALYSIS.md) identifies the failure pattern directly:

- default-lane truth had to be restated
- invalid evidence had to be replaced before the claim could stand
- cropped screenshots and non-default lanes were explicitly named as disqualifiers

`../ORTHOGONAL-SOLUTIONS-MATRIX.md` already fixes the mechanism boundary for `PROBLEM-0001` at the regression-claim boundary:

- the system cannot claim regression passed unless it restates the repo-defined lane and attaches an auditable evidence pack with explicit disqualifiers

The shared orchestration docs already identify the durable surfaces where this boundary lives:

- `Processes/TESTING.md` defines the shared proof-quality section for `REGRESSION-RUN-<NNNN>.md`
- `ORCHESTRATION.md` defines regression and closure honesty rules
- `Prompts/TASK-LEADER.md`, `Prompts/REGRESSION-LEADER.md`, and `Prompts/REGRESSION-TESTER.md` own the prompts that write, judge, and route regression claims
- `Exemplars/REGRESSION-RUN-0001.md` shows the intended run-artifact structure

That is enough to justify a first shared contract-and-enforcement task without inventing a tooling home.

## Why This Mechanism

The first durable intervention should be a stronger proof contract plus prompt and closure enforcement at the regression-claim boundary.

This mechanism is chosen because it acts at the exact point where the bad claim currently escapes:

- the task-owned `REGRESSION-RUN-<NNNN>.md` artifact
- the shared prompts and closure rules that decide whether that artifact counts as regression proof

That is the boundary where supporting evidence is currently mistaken for valid regression proof.
That is where the correction burden should be blocked first.

## Scope Rationale

This rewrite intentionally narrows the earlier draft.

It keeps two linked internal mechanisms:

1. a stronger shared proof contract for `REGRESSION-RUN-<NNNN>.md`
2. shared prompt and closure enforcement that makes the contract result-affecting

Those two belong in one first task because each is weak without the other:

- if the contract exists without enforcement, the artifact can still be ignored or summarized loosely at closeout
- if the prompts enforce a gate without a stronger contract, the gate still lacks a precise shared shape for what must be present

This rewrite removes the earlier standalone linter idea from the first slice.
The local sources justify the shared docs, shared prompts, and task-owned artifact boundary, but they do not name an exact shared tooling home for a linter honestly.

This rewrite also removes the contested-visual checklist as a separate third mechanism.
Contested-visual handling remains inside the proof contract itself:

- either the artifact shows the disputed surface fully enough to support the claim
- or the artifact records that limitation as a disqualifier and cannot be treated as clean closure proof

## Goals

- make default-lane regression claims auditable from the task-owned run artifact alone
- require visual proof claims to say whether the disputed surface is actually shown or disqualified
- make wrong-lane and visually disqualified proof result-affecting at regression and closure time
- block clean closure claims that rely only on supporting lanes or incomplete contested-surface evidence
- reduce repeated human restatement of what lane and what evidence count

## Non-Goals

- building a full unattended default-lane runner
- adding a standalone shared linter or validator in this first slice
- guaranteeing fully automatic pixel-level visual validation on day one
- redefining repo-root `REGRESSION.md`
- fixing the underlying ThirdPerson animation, camera, or rendering defect being measured

## Implementation Home

Primary shared contract home:

- `C:\Users\gregs\.codex\Orchestration\Processes\TESTING.md`
- `C:\Users\gregs\.codex\Orchestration\Exemplars\REGRESSION-RUN-0001.md`

Shared enforcement home:

- `C:\Users\gregs\.codex\Orchestration\ORCHESTRATION.md`
- `C:\Users\gregs\.codex\Orchestration\Prompts\TASK-LEADER.md`
- `C:\Users\gregs\.codex\Orchestration\Prompts\REGRESSION-LEADER.md`
- `C:\Users\gregs\.codex\Orchestration\Prompts\REGRESSION-TESTER.md`

Repo-local source of lane truth:

- each repo's canonical `REGRESSION.md`

Task-owned proof artifact home:

- `Tracking/Task-<id>/Testing/REGRESSION-RUN-<NNNN>.md`

Standalone shared tooling home in this first slice:

- none

## Implementation Home Rationale

This does not belong primarily in one repo's product code.

The visible failure happened in ThirdPerson, but the failure class is shared:

- an agent can still move a wrong-lane or incomplete-evidence artifact too far toward closure

It also does not belong primarily in repo-root `REGRESSION.md`.
Repo-root `REGRESSION.md` should define what lane or case counts as regression for that repo.
It should not carry the full shared contract for how regression claims and closure claims are narrated and judged across repos.

`Processes/TESTING.md` is the right contract home because it already owns the shared proof-quality fields for `REGRESSION-RUN-<NNNN>.md`.
`Exemplars/REGRESSION-RUN-0001.md` is the right supporting home because it already demonstrates the intended run-artifact structure.
`ORCHESTRATION.md`, `TASK-LEADER.md`, `REGRESSION-LEADER.md`, and `REGRESSION-TESTER.md` are the right enforcement homes because they already own:

- closure claim honesty
- regression-phase routing
- writing the executed run artifact
- refusing weaker surrogate proof when the repo-root lane is still required

This rewrite does not name a shared tooling home because the local sources do not justify one exactly.
Inventing a validator path here would make the task look more concrete than the durable evidence supports.

## Internal Mechanism Map

### Mechanism 1: Stronger Regression-Run Proof Contract

Failure reduced:

- the artifact still leaves the reviewer to reconstruct which lane ran and whether the visual evidence actually shows the disputed fact

Mechanism:

- `Processes/TESTING.md` and `Exemplars/REGRESSION-RUN-0001.md` require the run artifact to say:
  - claimed lane
  - actual flow exercised
  - why this counts
  - disqualifiers / limitations
  - evidence gathered
  - for disputed visual claims, whether the contested surface is fully shown or instead disqualified

Acceptance focus:

- a reviewer can read the artifact and know whether the claim is default-lane proof, only supporting proof, or disqualified

Falsifier:

- a run artifact can still describe a visual-fix claim without making clear whether the disputed surface is actually visible

### Mechanism 2: Regression And Closure Enforcement

Failure reduced:

- the artifact may contain caveats, but the workflow can still summarize it as clean regression or clean closure

Mechanism:

- `ORCHESTRATION.md`, `TASK-LEADER.md`, `REGRESSION-LEADER.md`, and `REGRESSION-TESTER.md` treat wrong-lane proof and visually disqualified proof as blocking or downgrading the claim rather than decorative notes

Acceptance focus:

- a wrong-lane run or a run with cropped or missing contested-surface evidence cannot be treated as clean closure-ready regression proof

Falsifier:

- the workflow can still mark clean pass or clean closure while the artifact itself names wrong-lane proof or invalid visual evidence

## Proposed Changes

### 1. Strengthen the shared regression-run proof contract

Update `C:\Users\gregs\.codex\Orchestration\Processes\TESTING.md` so the shared proof-quality contract for `REGRESSION-RUN-<NNNN>.md` explicitly requires:

- `Claimed lane`
- `Actual flow exercised`
- `Why this counts`
- `Disqualifiers / limitations`
- `Evidence gathered`

For a disputed visual claim, the contract must also require the artifact to say whether the contested surface is actually shown by the cited evidence.
If it is cropped, absent, or otherwise not fully supportable, that must be recorded as a disqualifier rather than silently left to reviewer inference.

### 2. Align the shared regression-run exemplar with that contract

Update `C:\Users\gregs\.codex\Orchestration\Exemplars\REGRESSION-RUN-0001.md` so the exemplar demonstrates:

- the strengthened proof-quality section
- evidence gathered tied to the claim
- contested-visual coverage or explicit disqualification
- the difference between diagnostic evidence and closure-ready proof

### 3. Enforce the contract in the named shared workflow and prompt artifacts

Update these exact shared enforcement artifacts so they require the stronger proof contract and treat wrong-lane or visually disqualified proof as result-affecting:

- `C:\Users\gregs\.codex\Orchestration\ORCHESTRATION.md`
  - clarify that closure-ready regression proof must not rely on wrong-lane evidence or on a run whose disputed visual claim is disqualified by cropped or missing evidence
- `C:\Users\gregs\.codex\Orchestration\Prompts\TASK-LEADER.md`
  - require the closing leader to reject `REGRESSION-RUN-<NNNN>.md` as closure evidence when disqualifiers show wrong-lane proof or incomplete contested-surface proof
- `C:\Users\gregs\.codex\Orchestration\Prompts\REGRESSION-LEADER.md`
  - require the regression leader to route such runs as partial, blocked, or insufficient rather than clean pass
- `C:\Users\gregs\.codex\Orchestration\Prompts\REGRESSION-TESTER.md`
  - require the regression tester to record contested-visual limitations explicitly and not silently treat supporting proof as default-lane proof

## Rival Mechanisms Considered

### Rival 1: Repo-local `REGRESSION.md` rewrite only

Why not first:

- repo-root `REGRESSION.md` defines lane truth, but the shared burden is that malformed proof can still be narrated and judged badly after that truth is already known

### Rival 2: Standalone linter in shared tooling

Why not in this first slice:

- the local sources do not identify an exact shared tooling home honestly
- inventing that home would leave the task in the same implementation-home ambiguity the audit flagged

### Rival 3: Contested-visual checklist as a separate mechanism

Why not:

- the actual burden is not `missing checklist text`
- the burden is that disputed-surface evidence is not made result-affecting at the regression and closure boundary

### Rival 4: Unattended runner first

Why not:

- a runner may later make evidence capture cheaper, but the proof contract still needs to define what counts as valid proof and what remains disqualified

## Not Solved Here

This task does not:

- build the unattended evidence-capture lane from Option B
- guarantee that every repo can validate contested visuals automatically
- replace repo-local lane truth
- prove the underlying product is fixed

It only hardens the proof and closure boundary for regression claims.

## Human Relief If Successful

The human should no longer have to say, for the same class of failure:

- `that was not the default lane`
- `that screenshot does not show the disputed surface`
- `that may be useful evidence, but it is not closure proof`
- `you cannot call that clean regression`

The reviewer should be able to inspect one durable run artifact and decide quickly whether it is valid regression proof, only supporting evidence, or still disqualified.

## Acceptance Criteria

### Contract Criteria

- `C:\Users\gregs\.codex\Orchestration\Processes\TESTING.md` defines the strengthened proof-quality contract for `REGRESSION-RUN-<NNNN>.md`
- `C:\Users\gregs\.codex\Orchestration\Exemplars\REGRESSION-RUN-0001.md` demonstrates that same contract in the shared exemplar shape
- for disputed visual claims, the contract requires the artifact either to identify evidence that shows the contested surface or to record the missing or cropped view as a disqualifier

### Enforcement Criteria

- `C:\Users\gregs\.codex\Orchestration\ORCHESTRATION.md`, `TASK-LEADER.md`, `REGRESSION-LEADER.md`, and `REGRESSION-TESTER.md` all align on the stronger proof contract
- those artifacts do not allow a clean regression or clean closure claim when the run artifact discloses:
  - supporting or alternate lane proof for a default-lane claim
  - cropped, absent, or otherwise incomplete contested-surface evidence for a disputed visual claim
- the shared workflow requires such runs to remain supporting, partial, blocked, or otherwise not closure-ready rather than treating the disqualifier as decorative

### Artifact Criteria

- the next comparable `REGRESSION-RUN-<NNNN>.md` for a default-lane claim states:
  - claimed lane
  - actual flow exercised
  - why this counts
  - disqualifiers / limitations
  - evidence gathered
- when the claim depends on a disputed visual fact, the artifact explicitly says whether the contested surface is shown by the cited evidence

### Burden-Reduction Criteria

- for the next comparable regression review, the reviewer can tell from the run artifact itself whether the claim is valid or disqualified without reconstructing the lane and evidence meaning from raw logs
- wrong-lane proof and cropped or missing contested-surface proof are blocked from being presented as clean closure-ready regression proof
- the human does not need to restate the same default-lane or contested-surface rule for the same claim class

## Proof Plan

Use at least these fixtures:

### Fixture 1: Wrong-Lane Proof

- the artifact claims the repo's default lane
- the actual flow exercised is only a supporting or alternate lane

Expected workflow result:

- the artifact records that honestly
- the shared prompts do not allow the run to be treated as clean default-lane regression proof or clean closure evidence

### Fixture 2: Cropped Contested Visual

- the artifact claims a disputed visual fix
- the cited evidence does not show the contested surface fully enough to support the claim

Expected workflow result:

- the artifact records that limitation explicitly as a disqualifier
- the run is not treated as clean closure-ready proof for that visual claim

### Fixture 3: Honest Closure-Ready Regression Artifact

- the default lane actually ran
- the cited evidence supports the disputed visual claim
- no unresolved disqualifier remains for the claimed proof surface

Expected workflow result:

- the run artifact is readable without reconstruction
- the shared prompts and closure rules may treat it as closure-ready regression proof, subject to any other repo-local requirements

## What Does Not Count

This task is not complete if:

- `TESTING.md` lists the proof-quality fields but the shared prompts still allow clean closure from weaker prose
- the artifact names screenshots but never says whether the disputed surface is actually shown
- cropped or missing contested-surface evidence is recorded but still treated as clean pass or clean closure
- supporting-lane proof is carefully documented but still accepted as default-lane regression proof
- the task claims linter enforcement without naming an exact shared tooling home

## Remaining Uncertainty

- some repos may need small repo-local adapter wording to define which visual surfaces are contested for specific regression cases
- this first slice intentionally excludes standalone shared tooling because the local sources do not justify an exact shared linter or validator home yet
- Option B may still be useful later to make evidence capture cheaper and more repeatable once the proof contract is already honest

These do not block the task writeup because the current draft now names one concrete first mechanism set, one exact enforcement home, and one explicit non-goal for the tooling question.

## Falsifier

This task is wrong or incomplete if, after implementation:

- a run can still be treated as clean default-lane closure proof even though it used only a supporting or alternate lane
- a run can still be treated as clean closure proof even though the artifact itself discloses cropped, absent, or incomplete contested-surface evidence
- the reviewer still has to reconstruct raw logs or screenshots just to know whether the artifact proves the disputed visual claim
- the shared workflow still treats the strengthened proof contract as advisory wording rather than a real gate

## References

- burden driver `BD-001` in [`../BURDEN-ANALYSIS.md`](../BURDEN-ANALYSIS.md)
- problem framing in [`../ORTHOGONAL-SOLUTIONS-MATRIX.md`](../ORTHOGONAL-SOLUTIONS-MATRIX.md)
- shared task-writing rules in [`../../../../../../../../Processes/TASK-CREATE.md`](../../../../../../../../Processes/TASK-CREATE.md)
- shared task-audit rules in [`../../../../../../../../Processes/TASK-AUDIT.md`](../../../../../../../../Processes/TASK-AUDIT.md)
- shared testing contract in [`../../../../../../../../Processes/TESTING.md`](../../../../../../../../Processes/TESTING.md)
- shared lifecycle rules in [`../../../../../../../../ORCHESTRATION.md`](../../../../../../../../ORCHESTRATION.md)
- shared regression-run exemplar in [`../../../../../../../../Exemplars/REGRESSION-RUN-0001.md`](../../../../../../../../Exemplars/REGRESSION-RUN-0001.md)
