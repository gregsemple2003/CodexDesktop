# PROBLEM-0005 OPTION-A: Enforce First-Disagreement Debugging And Durable Bug Routing

## Title

Enforce the first-disagreement debugging gate and require a durable bug narrative before a failed regression can drift into tweak mode or false closure.

## Summary

The packet shows repeated human correction of two linked failures:

- the system stayed in bounded-tweak mode instead of switching into controlled narrowing from a first concrete disagreement
- confirmed or blocked regression failures were not being carried forward with an honest `BUG-<NNNN>.md` and honest `REGRESSION-RUN-<NNNN>.md` narrative soon enough

This rewrite narrows Option A to the exact first slice justified by the local sources:

- strengthen the shared debugging contract around first-disagreement narrowing and durable bug-note structure
- enforce that contract in the shared task, regression, and debugging prompts that route work after a required regression failure

This slice does **not** include standalone lint tooling, because the local sources do not name an exact shared tooling home for that mechanism honestly.

## Writeup Type

Concrete implementation task.

The local sources already justify the failure class, the first intervention boundary, and the exact shared workflow artifacts that own debug routing after a regression failure.

## Source Event IDs

```json
[
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4768",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-6105",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5119",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5129",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5154",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-6755",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5340",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-7446"
]
```

## Burden Being Reduced

The human is currently forced to do three kinds of repeated debugging-repair work by hand:

1. Mode-switch policing.
   The human must repeatedly say `stop tweaking and trace the first disagreement upstream`.
2. Bug-narrative policing.
   The human must demand a real `BUG-<NNNN>.md` narrative instead of letting the issue live only in ad hoc commentary or scattered rerun notes.
3. Regression-honesty policing.
   The human must stop a failed or blocked regression lane from drifting back into implied progress or implied closure without an honest rerun artifact and explicit bug route.

The deeper burden is loss of trust in the narrowing path.
A reviewer cannot rely on the system to preserve what failed, what disagreement was observed, what branch was ruled out, or why the task is still open.

## Current Truth

The shared workflow already contains the right pieces in isolation:

- `Processes/DEBUGGING.md` already says debugging is controlled narrowing, starts from the first concrete disagreement, and requires immediate `BUG-<NNNN>.md` creation when proof becomes defect tracking
- `Processes/TESTING.md` already says a required regression-lane failure or block must be preserved in `BUG-<NNNN>.md` before routing into debugging
- `ORCHESTRATION.md` already says a failed or blocked required regression lane must create or update `BUG-<NNNN>.md` before leaving regression
- `REGRESSION-LEADER.md`, `DEBUG-LEADER.md`, and `DEBUG-WORKER.md` already describe the expected routing and first-disagreement behavior

But `BD-005` shows that this is still not enough in practice.

The current system truth is therefore not merely `the bug note could be better`.
The current truth is that, after a confirmed runtime failure:

- the system can keep proposing bounded tweaks instead of switching into first-disagreement narrowing
- the bug narrative can remain too weak or too late to preserve the actual disagreement path
- a failed or blocked regression lane can still drift toward implicit progress without an explicit bug route and honest rerun record

The current fallback truth is only that the human can intervene and restate the debugging method, reopen the issue, and demand a real bug note.
That is a rescue path, not a durable gate.

## Target Truth

After a required regression failure or comparable runtime defect is confirmed, the workflow should force one honest transition:

- the failing or blocked lane is preserved in `REGRESSION-RUN-<NNNN>.md`
- the task has an active `BUG-<NNNN>.md`
- the bug note names the first concrete disagreement, the current upstream boundary, and the narrowing path so far
- subsequent debug work stays anchored on checks chosen for narrowing power rather than drifting into tweak mode

A later reviewer should be able to inspect the durable artifacts and answer:

- what lane failed or blocked
- what exact disagreement was first observed
- what branch the current debugging effort is tracing upstream
- what evidence supports or weakens the current hypothesis
- what contradictory evidence remains
- why the task is still in debugging instead of closure

A failed or blocked regression lane may still lead to a fix later, but it should not be able to bypass the durable bug narrative or pretend the issue is already understood.

## Causal Claim

If the shared workflow forces confirmed regression failures through a first-disagreement debugging gate and a durable bug-note route, the system will stop oscillating between tweaks and will preserve the narrowing path honestly enough that the human no longer has to re-impose the debugging method by hand.

The cause being addressed is not just missing notes.
The cause is that the workflow still allows a regression-exposed defect to continue without a strong enough mode switch from `proof gathering` to `controlled defect narrowing`.

## Evidence

`BD-005` in [`../BURDEN-ANALYSIS.md`](../BURDEN-ANALYSIS.md) identifies both halves of the failure directly:

- the human had to demand a first-disagreement, upstream-tracing debugging method with concrete values
- the human had to demand durable defect tracking through `BUG-...` and honest regression-run artifacts rather than implicit closure

`../ORTHOGONAL-SOLUTIONS-MATRIX.md` already sets the mechanism boundary for `PROBLEM-0005`:

- after a defect survives one fix attempt or after a regression failure is confirmed, switch to the shared debugging method
- define the first concrete disagreement with values
- trace those exact values upstream one boundary at a time
- preserve contradictions in a durable bug narrative

The shared orchestration docs already identify the exact durable surfaces where this gate lives:

- `Processes/DEBUGGING.md`
- `Processes/TESTING.md`
- `ORCHESTRATION.md`
- `Prompts/TASK-LEADER.md`
- `Prompts/DEBUG-LEADER.md`
- `Prompts/DEBUG-WORKER.md`
- `Prompts/REGRESSION-LEADER.md`
- `Prompts/REGRESSION-TESTER.md`

That is enough to justify a first shared process-and-prompt enforcement task without inventing a tooling home.

## Why This Mechanism

The first durable intervention should be a debugging gate plus bug-routing enforcement at the exact point where regression proof turns into concrete defect tracking.

This mechanism is chosen because it acts at the boundary where the burden currently escapes:

- the moment a required regression lane fails or blocks
- the task-owned `BUG-<NNNN>.md` and `REGRESSION-RUN-<NNNN>.md` artifacts that should preserve the issue
- the shared prompts that decide whether the task is now debugging, still in regression, or drifting toward false closure

That is the moment when tweak mode should stop and controlled narrowing should begin.
That is where the correction burden should be blocked first.

## Scope Rationale

This rewrite intentionally narrows the earlier draft.

It keeps two linked mechanisms:

1. a stronger first-disagreement debugging contract for `BUG-<NNNN>.md`
2. shared routing and prompt enforcement that makes bug-note creation and honest regression carry-forward mandatory after confirmed failure

Those belong in one first task because each is weak without the other:

- if the bug-note contract exists without routing enforcement, the system can still keep tweaking or leave the bug narrative too weak to matter
- if prompts route into debugging without a stronger bug-note contract, the resulting bug artifact can still fail to preserve the actual disagreement path

This rewrite removes optional lint tooling from the first slice.
The local sources justify the shared docs, exact shared prompts, and task-owned artifact boundary, but they do not identify an exact shared tooling home honestly.

It also keeps regression-proof honesty inside the task rather than treating it as a separate sibling mechanism.
That link is already explicit in the local sources:

- failed or blocked required regression must create a bug note
- the rerun artifact must stay honest about what actually failed, blocked, or was not rerun

## Goals

- force a mode switch from tweak-driven debugging to first-disagreement narrowing after confirmed regression failure
- require an active `BUG-<NNNN>.md` before a failed or blocked required regression lane can leave regression honestly
- make the bug note preserve the first disagreement, the current upstream boundary, and the narrowing path
- keep `REGRESSION-RUN-<NNNN>.md` honest about the failing or blocked lane while the bug is active
- reduce repeated human restatement of the debugging method and bug-note requirement

## Non-Goals

- building a runtime disagreement capture harness
- adding standalone shared lint tooling in this first slice
- replacing repo-root `REGRESSION.md`
- fixing the underlying runtime defect itself
- redefining the shared debugging philosophy beyond the existing contract

## Implementation Home

Primary shared contract home:

- `C:\Users\gregs\.codex\Orchestration\Processes\DEBUGGING.md`
- `C:\Users\gregs\.codex\Orchestration\Processes\TESTING.md`

Shared lifecycle and enforcement home:

- `C:\Users\gregs\.codex\Orchestration\ORCHESTRATION.md`
- `C:\Users\gregs\.codex\Orchestration\Prompts\TASK-LEADER.md`
- `C:\Users\gregs\.codex\Orchestration\Prompts\DEBUG-LEADER.md`
- `C:\Users\gregs\.codex\Orchestration\Prompts\DEBUG-WORKER.md`
- `C:\Users\gregs\.codex\Orchestration\Prompts\REGRESSION-LEADER.md`
- `C:\Users\gregs\.codex\Orchestration\Prompts\REGRESSION-TESTER.md`

Task-owned artifact home:

- `Tracking/Task-<id>/BUG-<NNNN>.md`
- `Tracking/Task-<id>/Testing/REGRESSION-RUN-<NNNN>.md`

Standalone shared tooling home in this first slice:

- none

## Implementation Home Rationale

This does not belong primarily in product code for one repo.

The visible failure may happen in one task or one runtime surface, but the failure class is shared:

- a regression-exposed defect can still drift into tweak mode without a durable narrowing narrative

It also does not belong only in repo-root `REGRESSION.md`.
`REGRESSION.md` defines what lane counts as regression.
It should not also carry the shared cross-repo contract for how confirmed failures become bug narratives and debugging phases.

`Processes/DEBUGGING.md` is the right contract home because it already owns:

- first concrete disagreement
- controlled narrowing
- immediate bug-note creation
- bug-note structure

`Processes/TESTING.md` is also part of the contract home because it already owns the rule that required regression-lane failure or block must be preserved in a `BUG-<NNNN>.md`.

`ORCHESTRATION.md` and the exact prompt files above are the right enforcement homes because they already own:

- routing from regression into debugging
- the requirement not to leave a failed lane with only a regression note
- the canonical `BUG-<NNNN>.md` writer during debugging
- the first-disagreement behavior for debug branches

This rewrite does not name a shared tooling home because the local sources do not justify one exactly.
Inventing a validator path here would recreate the same implementation-home ambiguity the audit flagged.

## Internal Mechanism Map

### Mechanism 1: First-Disagreement Bug Contract

Failure reduced:

- the system keeps debugging at the level of categories or tweaks instead of a concrete disagreement with values

Mechanism:

- `DEBUGGING.md` and the debugging prompts require the active `BUG-<NNNN>.md` to preserve:
  - the first concrete disagreement
  - the current upstream boundary
  - the narrowing path so far
  - contradictory evidence
  - the next branch chosen for narrowing power

Acceptance focus:

- a reviewer can see that debugging is anchored on a concrete disagreement rather than symptom-level churn

Falsifier:

- the bug note still names only broad categories or tweak intentions without a concrete disagreement and narrowing path

### Mechanism 2: Failed-Regression Bug Routing And Proof Honesty

Failure reduced:

- a required regression failure can still drift forward without an active bug narrative or with misleading rerun claims

Mechanism:

- `TESTING.md`, `ORCHESTRATION.md`, `TASK-LEADER.md`, `REGRESSION-LEADER.md`, and `REGRESSION-TESTER.md` require a failed or blocked required lane to produce:
  - an honest `REGRESSION-RUN-<NNNN>.md`
  - an active `BUG-<NNNN>.md`
  - explicit routing into debugging or a human/environment gate

Acceptance focus:

- a confirmed failed or blocked regression lane cannot be treated as ordinary progress, partial closure, or silent carry-forward without a bug route

Falsifier:

- the workflow can still leave a failed required lane with only a rerun note, no active bug note, or a misleading implication that the task is closure-ready

## Proposed Changes

### 1. Strengthen the shared debugging contract

Update `C:\Users\gregs\.codex\Orchestration\Processes\DEBUGGING.md` so the shared bug-note contract explicitly requires:

- the first concrete disagreement with values when the seam exposes them
- the current upstream boundary being traced
- the narrowing path as a timeline of discriminating steps
- contradictory evidence, not only supporting evidence
- explicit linkage to any actually rerun `REGRESSION-RUN-<NNNN>.md`

### 2. Strengthen the shared failure-to-debugging route

Update `C:\Users\gregs\.codex\Orchestration\Processes\TESTING.md` and `C:\Users\gregs\.codex\Orchestration\ORCHESTRATION.md` so they align explicitly on this gate:

- when a required regression lane fails or is blocked, the task cannot leave regression honestly without:
  - an honest `REGRESSION-RUN-<NNNN>.md`
  - an active `BUG-<NNNN>.md`
  - explicit routing into debugging or a real human/environment gate

### 3. Enforce the gate in the named shared prompts

Update these exact shared prompt files so the first-disagreement gate and bug-routing rule are result-affecting:

- `C:\Users\gregs\.codex\Orchestration\Prompts\TASK-LEADER.md`
  - require an active bug note before the task can leave a failed or blocked required regression lane
- `C:\Users\gregs\.codex\Orchestration\Prompts\REGRESSION-LEADER.md`
  - require failed or blocked required-lane runs to route into `BUG-<NNNN>.md` plus debugging, not soft closure
- `C:\Users\gregs\.codex\Orchestration\Prompts\REGRESSION-TESTER.md`
  - require honest failing-lane and blocker capture, and bug-note creation when `DEBUGGING.md` says the failure is a real investigation note
- `C:\Users\gregs\.codex\Orchestration\Prompts\DEBUG-LEADER.md`
  - require the canonical bug note to stay anchored on the first concrete disagreement and current upstream boundary
- `C:\Users\gregs\.codex\Orchestration\Prompts\DEBUG-WORKER.md`
  - require each bounded debug branch to preserve the disagreement seam, update the bug note honestly, and avoid broadening away from that seam without an honest block

## Rival Mechanisms Considered

### Rival 1: Shared docs only

Why not first:

- the local sources already show the shared docs exist, yet the burden persisted
- the missing piece is that the shared prompts and lifecycle routing must make the gate result-affecting

### Rival 2: Optional lint tooling

Why not in this first slice:

- the local sources do not identify an exact shared tooling home honestly
- keeping it would leave the draft in hidden superposition between process enforcement and unspecified tooling

### Rival 3: Runtime disagreement capture harness first

Why not:

- the matrix already reserves that as Option B
- the first cheaper, truer intervention is to force the workflow to preserve and route the disagreement correctly before adding new capture machinery

## Not Solved Here

This task does not:

- build the runtime disagreement capture harness from Option B
- guarantee that every disagreement value can be captured automatically
- fix the underlying product defect
- replace repo-root regression truth

It only hardens the debugging and failure-routing boundary for regression-exposed defects.

## Human Relief If Successful

The human should no longer have to say, for the same failure class:

- `stop tweaking and name the first disagreement`
- `trace that exact bad state upstream`
- `open the bug note now`
- `do not call that progress without an honest rerun artifact`

A later reviewer should be able to inspect the bug note and rerun artifact and understand exactly why the task is in debugging, what branch is active, and what has already been ruled out.

## Acceptance Criteria

### Contract Criteria

- `C:\Users\gregs\.codex\Orchestration\Processes\DEBUGGING.md` defines the strengthened bug-note contract around first concrete disagreement, current upstream boundary, narrowing-path timeline, and contradictory evidence
- `C:\Users\gregs\.codex\Orchestration\Processes\TESTING.md` and `C:\Users\gregs\.codex\Orchestration\ORCHESTRATION.md` align on the rule that a failed or blocked required regression lane must have an active `BUG-<NNNN>.md` before the task leaves regression

### Prompt Enforcement Criteria

- `TASK-LEADER.md`, `REGRESSION-LEADER.md`, `REGRESSION-TESTER.md`, `DEBUG-LEADER.md`, and `DEBUG-WORKER.md` all align on the stronger debugging gate
- those prompt files do not allow a confirmed failed or blocked required regression lane to drift into:
  - tweak-driven follow-up without an active `BUG-<NNNN>.md`
  - a debug branch that lacks a first concrete disagreement and current upstream boundary
  - soft closure or misleading progress language unsupported by the rerun artifact

### Artifact Criteria

- for the next comparable runtime defect, the task has an active `BUG-<NNNN>.md` as soon as the required regression lane is confirmed failed or blocked
- that bug note preserves:
  - the first concrete disagreement
  - the current upstream boundary
  - the debugging path so far
  - contradictory evidence
  - links or references to the relevant `REGRESSION-RUN-<NNNN>.md`
- the corresponding rerun artifact remains honest about what failed, blocked, or was not rerun

### Burden-Reduction Criteria

- after a comparable confirmed regression failure, the workflow routes into a durable bug narrative without the human having to restate that requirement
- the next debug step is chosen and recorded for narrowing power rather than as another bounded tweak
- the human does not need to restate both:
  - the mode switch into first-disagreement debugging
  - the requirement for bug-note creation plus honest regression carry-forward

## Proof Plan

Use at least these fixtures:

### Fixture 1: Confirmed Required-Lane Failure

- a required regression lane fails or blocks honestly
- the task is not yet in debugging

Expected workflow result:

- the executed lane is preserved in `REGRESSION-RUN-<NNNN>.md`
- an active `BUG-<NNNN>.md` exists before the task leaves regression
- the task routes into debugging or a real human/environment gate rather than implied progress

### Fixture 2: Tweak-Mode Drift Attempt

- a defect survives one attempted fix
- the next proposed work is another bounded tweak without naming a first concrete disagreement

Expected workflow result:

- the workflow keeps the task in the debugging gate until the bug note names the disagreement and current upstream boundary
- the task does not honestly count as disciplined debugging yet

### Fixture 3: Honest Debug Narrative

- the bug note names a concrete disagreement with values
- one branch traces the disagreement upstream and records what it ruled in or out
- contradictory evidence is preserved

Expected workflow result:

- a later session can continue the narrowing path without rediscovering the same facts
- the task may proceed through debugging and later rerun regression honestly

## What Does Not Count

This task is not complete if:

- `DEBUGGING.md` sounds stronger but the shared prompts still allow tweak mode after confirmed failure
- a `BUG-<NNNN>.md` exists but lacks the first disagreement, current upstream boundary, or narrowing path
- a failed or blocked required regression lane can still end with only `REGRESSION-RUN-<NNNN>.md` and no active bug note
- the rerun artifact hides blocked or failed status while the bug is still open
- the task claims lint or validator enforcement without naming an exact shared tooling home

## Remaining Uncertainty

- some repos may need small repo-local adapter wording for which runtime values are most useful to name at the first disagreement seam
- this first slice intentionally excludes standalone shared tooling because the local sources do not justify an exact linter or validator home yet
- Option B may still be useful later to make disagreement capture cheaper and more repeatable once the workflow gate is already honest

These do not block the task writeup because the current draft now names one concrete first mechanism set, one exact enforcement home, and one explicit non-goal for the tooling question.

## Falsifier

This task is wrong or incomplete if, after implementation:

- a confirmed failed or blocked required regression lane can still leave regression without an active `BUG-<NNNN>.md`
- the system can still continue in tweak mode after confirmed failure without naming a first concrete disagreement and current upstream boundary
- the bug note exists but still fails to preserve the narrowing path or contradictory evidence well enough for a later session to continue honestly
- the rerun artifact and bug note can still combine into misleading progress or closure claims while the defect remains under active debugging

## References

- burden driver `BD-005` in [`../BURDEN-ANALYSIS.md`](../BURDEN-ANALYSIS.md)
- problem framing in [`../ORTHOGONAL-SOLUTIONS-MATRIX.md`](../ORTHOGONAL-SOLUTIONS-MATRIX.md)
- shared task-writing rules in [`../../../../../../../../Processes/TASK-CREATE.md`](../../../../../../../../Processes/TASK-CREATE.md)
- shared task-audit rules in [`../../../../../../../../Processes/TASK-AUDIT.md`](../../../../../../../../Processes/TASK-AUDIT.md)
- shared debugging process in [`../../../../../../../../Processes/DEBUGGING.md`](../../../../../../../../Processes/DEBUGGING.md)
- shared testing process in [`../../../../../../../../Processes/TESTING.md`](../../../../../../../../Processes/TESTING.md)
- shared lifecycle rules in [`../../../../../../../../ORCHESTRATION.md`](../../../../../../../../ORCHESTRATION.md)
