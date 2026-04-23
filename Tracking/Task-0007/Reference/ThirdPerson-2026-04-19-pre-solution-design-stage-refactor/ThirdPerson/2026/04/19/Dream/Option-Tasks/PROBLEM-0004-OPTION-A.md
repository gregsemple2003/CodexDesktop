# PROBLEM-0004 OPTION-A: Approval-Ready `PLAN.md` Gate With Diff And Context

## Title

Require an approval-ready packet inside the current `PLAN.md` before explicit human plan approval.

## Summary

The burden packet shows two linked approval-surface failures:

- the human had to reconstruct what changed because approval requests did not present diffs, links, and context in a reviewable shape
- reopened planning work could be presented as if the existing approval still covered it, even though the human needed to understand the new delta before approving

This rewrite narrows Option A to the exact first slice justified by the local durable sources:

- define the approval-packet contract for the one explicit shared human-approval boundary the workflow already names: approval of the current `Tracking/Task-<id>/PLAN.md`
- enforce that gate in the exact shared lifecycle docs, prompt files, and state contract that already own `PLAN.md` approval

This slice does **not** include a standalone approval-packet generator, a new approval-packet filename, or a broad pass-closeout approval system, because the local sources do not justify those implementation homes honestly.

## Writeup Type

Concrete implementation task.

The local durable sources already justify the burden, the exact human gate, the existing task-owned artifact being approved, and the shared prompt and state surfaces that control that gate.

## Source Event IDs

```json
[
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-3494",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4379",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4399",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4421",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-4755",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4512",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4522",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4571",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4581"
]
```

## Burden Being Reduced

The human is currently forced to do three kinds of approval-repair work by hand:

1. Diff reconstruction.
   The human must reopen multiple task files just to determine what changed and what decision is actually being requested.
2. Approval-boundary policing.
   The human must restate that explicit plan approval is a real gate and that materially revised planning work is not still approved by implication.
3. Review-surface reconstruction.
   The human must translate raw file references or vague summaries into a human-usable review packet before deciding whether to approve.

The deeper burden is trust loss at the approval surface.
A reviewer cannot rely on the current plan-approval handoff to show the new delta, the affected task artifacts, and the exact decision boundary in one place.

## Current Truth

The shared workflow already contains the right approval boundary in pieces:

- `ORCHESTRATION.md` says planning plus explicit human `PLAN.md` approval is Stage A of `IMPLEMENTATION-LEADER.md`
- `TASK-LEADER.md` says the lifecycle must route through explicit `PLAN.md` approval and must not bypass that human gate
- `IMPLEMENTATION-LEADER.md` says Stage B cannot begin before explicit approval of the current `PLAN.md`
- `IMPLEMENTATION-LEADER.md` also says a materially revised plan must not be treated as still approved
- `TASK-STATE.md` provides the durable `plan_approved` state boundary for that approval transition
- `TASK-CREATE.md` already names the desired acceptance shape: `IMPLEMENTATION-LEADER.md` requires a task-owned approval packet before asking for plan approval

But the burden packet shows that this is still not enough in practice.

The current system truth is therefore not merely `missing links`.
The current truth is that the explicit plan-approval gate exists, yet the shared workflow does not define an approval-ready packet shape for the human-facing review surface of that gate.

That leaves two failure modes:

- approval requests can still reach the human as vague summaries, raw path dumps, or low-context links
- materially changed planning work can still require manual reconstruction before the human can tell whether the current approval request is truly new, partially reopened, or already understood

The current fallback truth is only that the human can refuse the approval request and ask for diffs and context in a better shape.
That is a rescue path, not a durable review surface.

## Target Truth

Before the workflow asks the human to approve the current `Tracking/Task-<id>/PLAN.md`, the plan itself should carry an approval-ready packet that lets the reviewer answer, without reopening multiple files first:

- what exact decision is being requested
- what changed since the last approval-ready baseline
- which task-owned artifacts are relevant to that delta
- why those changes were made
- what remains out of scope or still unresolved
- whether the plan is genuinely new or materially revised enough to require fresh approval

A reviewer should be able to read the current `PLAN.md`, follow a small number of targeted links, and decide quickly whether to approve, request revision, or reject.

The human gate should remain explicit, but the review burden to reach that gate should be low.

## Causal Claim

If the shared workflow forces the current `PLAN.md` to carry an approval-ready packet before the plan-approval request is made, the human will no longer need to reconstruct the approval delta by hand and materially revised planning work will stop slipping through as if earlier approval still covered it.

The cause being addressed is not only missing prose quality.
The cause is that the workflow has a real approval boundary without a strong enough review-surface contract attached to that boundary.

## Evidence

`BD-004` in [`../BURDEN-ANALYSIS.md`](../BURDEN-ANALYSIS.md) identifies both halves of the burden directly:

- the human had to reconstruct what changed because diffs and links were missing or unusable
- the human had to enforce pass-structure or approval-boundary correctness so reopened work was not silently treated as already settled

`../ORTHOGONAL-SOLUTIONS-MATRIX.md` already sets the mechanism boundary for `PROBLEM-0004`:

- before asking for approval, the system must emit an approval packet with diffs and contextual links that a human can review quickly

The exact shared human-approval boundary justified by the local sources is the `PLAN.md` approval gate:

- `TASK-CREATE.md` names a task-owned approval packet before asking for plan approval
- `ORCHESTRATION.md` defines planning plus explicit human `PLAN.md` approval as a real lifecycle transition
- `TASK-LEADER.md` and `IMPLEMENTATION-LEADER.md` own routing into and through that gate
- `TASK-STATE.md` defines the durable approval state with `plan_approved`

That is enough to justify a first shared contract-and-enforcement task without inventing a new approval-packet filename or tooling home.

## Why This Mechanism

The first durable intervention should act at the exact point where the burden currently escapes:

- the task-owned artifact being approved, `Tracking/Task-<id>/PLAN.md`
- the explicit human approval request for that artifact
- the durable `plan_approved` state transition that follows

This mechanism is chosen because it hardens the existing approval gate rather than creating a second, parallel review system.

The missing piece is not another approval workflow.
The missing piece is that the current approval boundary lacks an exact shared review surface.

## Scope Rationale

This rewrite intentionally narrows the earlier draft.

It keeps two linked mechanisms:

1. a shared approval-packet contract for the current `PLAN.md`
2. shared lifecycle and prompt enforcement that makes that packet required before explicit human plan approval

Those belong in one first task because each is weak without the other:

- if the packet contract exists without enforcement, the human gate can still be reached through vague summaries
- if the prompts enforce a gate without a packet contract, the review surface remains underspecified and the burden just shifts back to the human

This rewrite removes the earlier standalone generator from the first slice.
The local sources do not name an exact shared tooling home for it honestly.

This rewrite also narrows away a broad pass-closeout approval or reopened-pass-history system.
The local sources clearly justify the explicit `PLAN.md` approval boundary and the rule that materially revised plans require fresh approval.
They do **not** yet justify a broader shared human-approval packet system for every pass-closeout artifact.

## Goals

- require an approval-ready review packet before explicit human approval of the current `PLAN.md`
- make materially revised plans visibly re-approval-worthy rather than silently carried by earlier approval
- let the human review the requested plan change through one durable task-owned artifact with targeted links
- reduce repeated human requests for diffs, context, and approval-boundary clarification
- make the approval-review surface result-affecting at the actual lifecycle gate

## Non-Goals

- building a standalone approval-packet generator in this first slice
- introducing a new shared approval-packet filename unsupported by the local sources
- redefining pass-closeout or commit/push audit as a human approval workflow
- creating a broad reopened-pass-history correction system outside the `PLAN.md` approval boundary
- replacing repo-local task content with shared policy prose

## Implementation Home

Primary shared contract home:

- `C:\Users\gregs\.codex\Orchestration\Processes\TASK-CREATE.md`

Shared lifecycle and approval-state home:

- `C:\Users\gregs\.codex\Orchestration\ORCHESTRATION.md`
- `C:\Users\gregs\.codex\Orchestration\TASK-STATE.md`

Shared prompt enforcement home:

- `C:\Users\gregs\.codex\Orchestration\Prompts\TASK-LEADER.md`
- `C:\Users\gregs\.codex\Orchestration\Prompts\IMPLEMENTATION-LEADER.md`

Task-owned artifact home for the packet in this first slice:

- `Tracking/Task-<id>/PLAN.md`

Supporting task-owned artifacts that the packet may link when they explain the approval delta:

- `Tracking/Task-<id>/TASK.md`
- `Tracking/Task-<id>/RESEARCH.md`
- `Tracking/Task-<id>/HANDOFF.md`
- optional `Tracking/Task-<id>/TASK-STATE.json`

Standalone shared tooling home in this first slice:

- none

## Implementation Home Rationale

This does not belong primarily in one repo's product code.

The burden is a shared orchestration failure:

- a real human gate exists, but its review surface is underspecified

It also does not belong primarily in `PASS-CHECKLIST.md` or `AUDITOR.md`.
Those artifacts own pass-closeout state and read-only readiness review, not the explicit human `PLAN.md` approval boundary that the local sources name here.

`TASK-CREATE.md` is the right contract home because it already names the expected result:

- `IMPLEMENTATION-LEADER.md` requires a task-owned approval packet before asking for plan approval

`ORCHESTRATION.md`, `TASK-STATE.md`, `TASK-LEADER.md`, and `IMPLEMENTATION-LEADER.md` are the right enforcement homes because they already own:

- the explicit `PLAN.md` approval gate
- the rule that materially revised plans are not silently still approved
- the durable `plan_approved` transition
- routing from planning into implementation only after that approval

`Tracking/Task-<id>/PLAN.md` is the right task-owned artifact home because it is the existing artifact being approved.
The local sources do not justify inventing a new sibling approval-packet file, so this first slice should harden the review surface of the current `PLAN.md` itself.

## Internal Mechanism Map

### Mechanism 1: Approval-Ready `PLAN.md` Packet Contract

Failure reduced:

- the human must reconstruct what changed and what is being approved from vague summaries or low-context links

Mechanism:

- `TASK-CREATE.md` defines that the current `PLAN.md` must carry an approval-ready packet before explicit plan approval is requested
- that packet must include:
  - the exact approval decision being requested
  - the changed task-owned artifacts relevant to that decision
  - a concise summary of what changed in each relevant artifact
  - why those changes were made
  - targeted links that let the human inspect the delta quickly
  - what remains open or does not count as approved yet

Acceptance focus:

- a reviewer can read the current `PLAN.md` and understand the approval delta without reopening many files to reconstruct it manually

Falsifier:

- the approval request can still reach the human as `please approve the plan` with only raw path dumps, broad prose, or link spam

### Mechanism 2: Explicit Plan-Approval Gate Enforcement

Failure reduced:

- the workflow can still ask for or act on plan approval when the review surface is too weak or when the plan was materially revised after earlier approval

Mechanism:

- `ORCHESTRATION.md`, `TASK-STATE.md`, `TASK-LEADER.md`, and `IMPLEMENTATION-LEADER.md` make the packet result-affecting:
  - Stage A does not ask for plan approval until the current `PLAN.md` contains the required packet
  - materially revised plans are treated as needing fresh approval
  - `plan_approved` stays false until the human approves the packet-backed current plan
  - Stage B does not begin from an approval request that lacks the required review surface

Acceptance focus:

- the workflow cannot honestly move from planning into implementation on the basis of a low-context or stale approval request

Falsifier:

- a materially revised `PLAN.md` can still be treated as approved without a fresh packet-backed approval request

## Proposed Changes

### 1. Define the approval-packet contract for explicit `PLAN.md` approval

Update `C:\Users\gregs\.codex\Orchestration\Processes\TASK-CREATE.md` so the named `task-owned approval packet` is made concrete for this first slice:

- the packet lives in the current `Tracking/Task-<id>/PLAN.md`
- it is required before asking the human to approve that plan
- it must state:
  - the exact decision being requested
  - which task-owned artifacts changed and matter to that decision
  - a concise per-artifact `what changed` summary
  - why the change matters for the approval request
  - targeted links to the relevant task artifacts
  - what is still open, deferred, or not covered by this approval request

### 2. Align the lifecycle and durable approval-state boundary

Update these exact shared lifecycle artifacts so the approval packet becomes part of the actual gate:

- `C:\Users\gregs\.codex\Orchestration\ORCHESTRATION.md`
  - clarify that the explicit `PLAN.md` approval transition is not ready to present to the human until the current `PLAN.md` carries the required approval packet
- `C:\Users\gregs\.codex\Orchestration\TASK-STATE.md`
  - clarify that `plan_approved` remains false until the human approves the current packet-backed plan, and that materially revised plans require returning to the unapproved state rather than carrying prior approval implicitly

### 3. Enforce the packet in the named shared prompt files

Update these exact shared prompt artifacts so the missing packet blocks the approval request rather than being treated as optional presentation quality:

- `C:\Users\gregs\.codex\Orchestration\Prompts\IMPLEMENTATION-LEADER.md`
  - require the current `PLAN.md` to contain the approval-ready packet before asking for explicit human approval
  - require materially revised plans to seek fresh approval with a refreshed packet
  - forbid entering Stage B from a plan-approval request that lacks the required packet
- `C:\Users\gregs\.codex\Orchestration\Prompts\TASK-LEADER.md`
  - require the task leader to challenge or route back any plan-approval request that lacks the packet-backed review surface before sending it to the human

## Rival Mechanisms Considered

### Rival 1: Standalone approval-packet generator

Why not in this first slice:

- the local sources do not identify an exact shared tooling home honestly
- keeping it would leave the draft with the same implementation-home ambiguity the audit flagged

### Rival 2: New dedicated approval-packet artifact file

Why not first:

- the local sources name a `task-owned approval packet` but do not name a durable standalone filename for it
- the existing artifact being approved is `Tracking/Task-<id>/PLAN.md`, so the first honest slice is to harden that existing review surface

### Rival 3: Broad pass-closeout approval system

Why not first:

- the exact shared human gate justified by the local sources here is explicit `PLAN.md` approval
- `PASS-CHECKLIST.md` and `AUDITOR.md` govern pass-closeout state and audit readiness, not a parallel human approval packet system

## Not Solved Here

This task does not:

- create a generator or validator for approval packets
- define a new shared approval-packet filename
- solve every repo-specific review-surface problem
- standardize human approval for pass closeout, audit verdicts, or commit readiness
- replace the need for concise writing inside task-owned artifacts

It only hardens the review surface of the existing explicit `PLAN.md` approval gate.

## Human Relief If Successful

The human should no longer have to say, for the same approval class:

- `what changed`
- `show me the diff in a usable shape`
- `which files am I actually approving`
- `this was revised, so why are you acting like it is already approved`

The reviewer should be able to inspect the current `PLAN.md`, follow a few targeted links, and decide whether to approve without reconstructing the delta from scratch.

## Acceptance Criteria

### Contract Criteria

- `C:\Users\gregs\.codex\Orchestration\Processes\TASK-CREATE.md` defines the approval-packet contract for the current `Tracking/Task-<id>/PLAN.md`
- that contract requires the current plan to state:
  - the decision being requested
  - the relevant changed task-owned artifacts
  - what changed and why
  - targeted review links
  - what remains open or outside the approval request

### Lifecycle And State Criteria

- `C:\Users\gregs\.codex\Orchestration\ORCHESTRATION.md` and `TASK-STATE.md` align on the rule that the plan-approval gate is not honestly ready while the current `PLAN.md` lacks the required packet
- `TASK-STATE.md` makes clear that materially revised plans cannot remain implicitly approved and must return through explicit approval of the current packet-backed plan

### Prompt Enforcement Criteria

- `TASK-LEADER.md` and `IMPLEMENTATION-LEADER.md` both align on the stronger gate
- those prompt files do not allow:
  - asking the human to approve a plan that lacks the required packet
  - treating a materially revised plan as still approved without a fresh approval request
  - entering Stage B from a low-context approval request that still forces the human to reconstruct the delta manually

### Artifact Criteria

- for the next comparable plan-approval request, the current `Tracking/Task-<id>/PLAN.md` itself contains an approval-ready packet
- that packet names the changed task-owned artifacts and summarizes their delta in a reviewable shape
- a reviewer can decide whether to approve without reopening many files just to discover what changed

### Burden-Reduction Criteria

- the next comparable plan-approval request reaches the human with enough diff and context to review quickly
- materially revised planning work is clearly presented as a fresh approval request rather than being silently carried by prior approval
- the human does not need to restate both:
  - `show me the changes in a usable review shape`
  - `this revised plan still needs explicit approval`

## Proof Plan

Use at least these fixtures:

### Fixture 1: First Plan Approval Request

- `PLAN.md` is newly created or materially refreshed
- the task is still in planning

Expected workflow result:

- the current `PLAN.md` includes the required approval packet
- the human gate is presented through that packet
- `plan_approved` does not flip until the human explicitly approves the current plan

### Fixture 2: Material Plan Revision After Earlier Approval

- a previously approved `PLAN.md` is materially revised
- the workflow attempts to continue on the strength of earlier approval

Expected workflow result:

- the current plan is treated as needing fresh approval
- the updated `PLAN.md` includes a refreshed approval packet that shows the new delta
- Stage B does not proceed until the human approves that revised plan

### Fixture 3: Low-Context Approval Attempt

- the approval request consists only of broad prose, raw file paths, or untargeted links

Expected workflow result:

- the shared prompts refuse to treat that as an honest plan-approval request
- the plan returns to packet completion rather than proceeding into implementation

## What Does Not Count

This task is not complete if:

- `TASK-CREATE.md` says `approval packet` but still leaves the current `PLAN.md` review surface unspecified
- the approval request is only `please approve the plan` plus raw file paths
- links are provided, but the human still has to open many files just to discover what changed
- a materially revised `PLAN.md` is treated as still approved because an older plan version was approved earlier
- the task claims a generator, validator, or new packet filename without a locally justified home
- pass-closeout or audit artifacts are renamed as `approval packets` even though the local sources do not justify that broader gate here

## Remaining Uncertainty

- some repos may still need small repo-local guidance for what kinds of plan deltas deserve the most emphasis in the approval packet
- this first slice intentionally excludes a standalone generator or validator because the local sources do not justify an exact shared tooling home
- broader pass-closeout approval-surface standardization may still be valuable later, but the local sources here only justify the explicit `PLAN.md` approval boundary as the first shared task

These do not block the task writeup because the draft now names:

- the exact first human gate
- the exact existing task-owned artifact that carries the packet
- the exact shared docs, state contract, and prompts that must enforce it

## Falsifier

This task is wrong or incomplete if, after implementation:

- a plan-approval request can still reach the human without the current `PLAN.md` carrying the required approval packet
- a materially revised plan can still be treated as approved without a fresh packet-backed approval request
- the human still has to reconstruct the approval delta manually from scattered task artifacts before deciding whether to approve
- the shared workflow still treats the approval packet as presentation advice rather than a real gate on entering Stage B

## References

- burden driver `BD-004` in [`../BURDEN-ANALYSIS.md`](../BURDEN-ANALYSIS.md)
- problem framing in [`../ORTHOGONAL-SOLUTIONS-MATRIX.md`](../ORTHOGONAL-SOLUTIONS-MATRIX.md)
- shared task-writing rules in [`../../../../../../../../Processes/TASK-CREATE.md`](../../../../../../../../Processes/TASK-CREATE.md)
- shared task-audit rules in [`../../../../../../../../Processes/TASK-AUDIT.md`](../../../../../../../../Processes/TASK-AUDIT.md)
- shared lifecycle rules in [`../../../../../../../../ORCHESTRATION.md`](../../../../../../../../ORCHESTRATION.md)
- shared task-state contract in [`../../../../../../../../TASK-STATE.md`](../../../../../../../../TASK-STATE.md)
