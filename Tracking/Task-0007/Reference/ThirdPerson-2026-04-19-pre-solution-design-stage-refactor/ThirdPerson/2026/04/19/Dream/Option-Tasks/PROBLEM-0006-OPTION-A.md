# PROBLEM-0006 OPTION-A: Task-Local Constraint Ledger With Pre-Action Repo-Truth Checkpoints

## Title

Add a task-local constraint ledger that also carries repo-local truth into a required pre-action checkpoint before risky moves.

## Summary

The burden packet shows two linked failures:

- newly stated constraints were not retained as durable state, so the human had to restate them
- repo-local truths were not injected early enough into action decisions, so the system drifted into disallowed or invalid moves

This task adds one task-local mechanism for both failures:

- a durable `CONSTRAINTS.md` ledger that records active constraints and task-relevant repo truths
- a required pre-action checkpoint that must cite the relevant ledger entries before any move that could violate them

The intended result is not just that a ledger exists. The intended result is that a risky move fails closed unless the agent has already recorded and consulted the relevant constraint or repo-truth entry.

## Writeup Type

Concrete implementation task.

The burden analysis and matrix already identify the first intervention boundary: durable task-local constraint retention plus enforcement at the action boundary. This is not a research task, and it is not the global memory option.

## Source Event IDs

```json
[
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-3094",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-3197",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-3463",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-3482",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-3549",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-3585",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4294",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-4684"
]
```

## Burden Being Reduced

The human is currently forced to do two kinds of repeated repair work by hand:

1. Constraint restatement.
   The human must restate hard boundaries such as `no engine mods`, `do not touch REGRESSION.md`, and `stop at approval gate` because they do not become durable state soon enough.
2. Repo-truth restatement.
   The human must restate repo-local truths such as the human default lane definition and the rule that supporting lanes do not substitute for regression proof.

The deeper burden is that the human cannot trust the system to carry forward either newly stated constraints or already-known repo truth before acting.

## Current Truth

The current system truth is not merely that some constraints are missing from notes.

The current truth is that constraints and repo-local truths are treated too much like transient chat and not enough like enforced task state. As a result:

- newly stated hard constraints can be lost between steps
- repo-local truths are not injected reliably before a proof, edit, or approval-boundary move
- the human has to restate the same rules when the next risky move appears

The current fallback truth is only that the human can reassert the rule again. That is a rescue path, not a durable fix.

## Target Truth

After the task is implemented, a later agent working the same task should have one durable place to look for both:

- active hard constraints that the human has stated for the task
- task-relevant repo-local truths that must shape action decisions

Before a risky move, the workflow should require a short checkpoint that says, in effect:

- what action is about to happen
- which constraint or repo-truth entries govern that action
- whether the action is allowed, blocked, or requires human clarification

The human should not need to restate a recorded hard constraint or repo truth for the same task when the next comparable risky move occurs.

## Causal Claim

If the system must record constraints and repo-local truths in a task-local ledger, and must cite the relevant entries at a pre-action checkpoint before risky moves, then repeated re-learning should drop because the workflow will fail closed before a violating move or invalid substitution claim is made.

The cause being addressed is not weak wording alone.
The cause is the absence of a durable action-boundary check that carries recorded constraints and repo truth into the next move.

## Evidence

`BD-006` in [`../BURDEN-ANALYSIS.md`](../BURDEN-ANALYSIS.md) states that:

- the human had to introduce hard constraints and restate them when the system drifted
- repo-local workflow truths should have been loaded and retained
- the working hypothesis is that newly stated constraints were not treated as durable state
- the likely remedy class is a durable ledger plus a requirement to reference it before action

The matrix row for `PROBLEM-0006` in [`../ORTHOGONAL-SOLUTIONS-MATRIX.md`](../ORTHOGONAL-SOLUTIONS-MATRIX.md) fixes the mechanism boundary at the constraint-application boundary:

- capture new constraints durably when they appear
- reference the current constraint ledger explicitly when making a move that could violate it

The same row distinguishes Option A from Option B:

- Option A is task-local ledger plus enforcement
- Option B is global constraint memory across sessions

That means repo-truth injection is part of Option A, but only in the task-local, pre-action form.

## Why This Mechanism

The first durable intervention should be a task-local ledger plus pre-action checkpoint, not a reminder and not a global memory store.

This mechanism is chosen because it acts at the exact boundary where the burden escapes today:

- the moment just before an action, proof claim, or approval move that could violate a known constraint or repo truth

That boundary is where the system currently forgets, substitutes, or drifts.
That is where the correction burden should be blocked.

## Scope Rationale

This task intentionally combines two internal mechanisms at one boundary:

1. durable recording of task-local constraints and repo truths
2. a pre-action checkpoint that injects those entries into risky moves

This merge is earned because the burden analysis describes them as one failure chain:

- the rule is not retained durably
- therefore it is not present when the next risky move happens

A ledger without a checkpoint can still be ignored.
A checkpoint without a ledger can only cite memory or transient chat.

The task does not widen to global cross-session memory because the matrix reserves that as Option B.

## Goals

- record newly stated hard constraints durably within the active task
- record task-relevant repo-local truths durably within the same task artifact
- require a pre-action checkpoint before risky moves that could violate those entries
- block risky moves when no relevant checkpoint exists
- reduce repeated human restatement of the same constraint or repo truth within a task

## Non-Goals

- solving long-term cross-session memory or automatic global injection of constraints
- redefining repo-local truth sources such as a repo's canonical regression definition
- replacing the explicit wait-state and continuity contract from `PROBLEM-0002`
- replacing the regression-proof gate from `PROBLEM-0001`
- guaranteeing that every possible action can be classified automatically on day one

## Implementation Home

Primary task-owned artifact home:

- `Tracking/Task-<id>/CONSTRAINTS.md`

Shared enforcement home:

- `C:\Users\gregs\.codex\Orchestration\ORCHESTRATION.md`
- `C:\Users\gregs\.codex\Orchestration\Prompts\TASK-LEADER.md`
- `C:\Users\gregs\.codex\Orchestration\Prompts\IMPLEMENTATION-LEADER.md`
- `C:\Users\gregs\.codex\Orchestration\Prompts\REGRESSION-LEADER.md`

## Implementation Home Rationale

The source of truth for task-specific active constraints and task-relevant repo truths should live with the task, because the burden described in `BD-006` is repeated re-learning inside the active workstream.

This does not belong only in shared prompt wording, because wording alone cannot show which exact constraints are active for one task.

It does not belong only in repo-local docs, because repo-local docs define durable repo truth but do not record newly stated task-specific constraints such as `stop at approval gate`.

It does not belong in `PROBLEM-0002`'s wait-state contract, because that sibling task governs whether stopping and resuming are explicit. This task governs what truths must be consulted before acting.

The correct split is therefore:

- `Tracking/Task-<id>/CONSTRAINTS.md` stores the active ledger
- `C:\Users\gregs\.codex\Orchestration\ORCHESTRATION.md` owns the shared lifecycle rule that risky moves must consult the ledger
- `C:\Users\gregs\.codex\Orchestration\Prompts\TASK-LEADER.md` enforces that rule when routing phases and judging closure readiness
- `C:\Users\gregs\.codex\Orchestration\Prompts\IMPLEMENTATION-LEADER.md` enforces that rule before approval-boundary and pass-execution moves
- `C:\Users\gregs\.codex\Orchestration\Prompts\REGRESSION-LEADER.md` enforces that rule before regression-proof and closure-adjacent claims

## Internal Mechanism Map

### Mechanism 1: Durable Constraint And Repo-Truth Ledger

Failure reduced:

- constraints and repo truths disappear between steps and have to be re-taught

Mechanism:

- `CONSTRAINTS.md` stores active hard constraints and task-relevant repo truths as durable entries with status and enforcement notes

Acceptance focus:

- the next risky move can cite the governing entries from durable task state rather than from memory

Falsifier:

- the human still has to restate a rule that was already recorded in the same task

### Mechanism 2: Pre-Action Checkpoint Gate

Failure reduced:

- the system acts first and only later discovers that the move violated a known boundary

Mechanism:

- before a risky move, the workflow must emit a checkpoint that names the pending action, the consulted ledger entries, and the decision: `allowed`, `blocked`, or `needs_clarification`

Acceptance focus:

- risky moves fail closed when the checkpoint is missing or when the cited entries block the move

Falsifier:

- a risky move can still proceed without citing the relevant recorded entries

## Proposed Changes

### 1. Define a task-local `CONSTRAINTS.md` contract

Require `Tracking/Task-<id>/CONSTRAINTS.md` to contain at least these sections:

- `Active Constraints`
- `Repo-Local Truth`
- `Checkpoint Log`

Require each active entry to include at least:

- entry id
- entry type: `constraint` or `repo_truth`
- statement
- scope
- source
- status: `active | resolved | superseded`
- enforcement note

### 2. Define what counts as a checkpoint

A checkpoint is a durable log entry recorded in `CONSTRAINTS.md` before a risky move.

Each checkpoint entry must include at least:

- pending action
- relevant entry ids consulted
- decision: `allowed | blocked | needs_clarification`
- reason
- next step if blocked or unclear

### 3. Attach the checkpoint to the action boundary

Require the workflow to create a checkpoint before any move that could violate a recorded entry, including at minimum:

- edits that could cross a recorded `do not touch` boundary
- changes that could cross an approval gate
- proof or regression claims that depend on recorded repo-local lane truth

If no relevant checkpoint exists, the move must not proceed.

### 4. Require immediate capture of new constraints

When the human states a new hard constraint or clarifies a repo truth that materially affects the task, the next durable checkpoint must record the new entry before further risky work continues.

### 5. Enforce the rule in the named shared orchestration artifacts

Update these exact shared artifacts so agents are required to consult `CONSTRAINTS.md`, cite the relevant entry ids in the checkpoint, and stop instead of proceeding when the ledger is missing, stale, or blocking:

- `C:\Users\gregs\.codex\Orchestration\ORCHESTRATION.md`
  - add the shared lifecycle rule that risky edits, approval-boundary moves, and regression-proof claims must consult the task-local ledger first
- `C:\Users\gregs\.codex\Orchestration\Prompts\TASK-LEADER.md`
  - require the top-level leader to check for blocking or missing ledger entries before dispatching the next phase or claiming closure readiness
- `C:\Users\gregs\.codex\Orchestration\Prompts\IMPLEMENTATION-LEADER.md`
  - require the implementation leader to checkpoint recorded constraints before crossing approval gates or taking pass moves that could violate a `do not touch` or similar boundary
- `C:\Users\gregs\.codex\Orchestration\Prompts\REGRESSION-LEADER.md`
  - require regression-proof and closure-adjacent claims to checkpoint any recorded repo-truth entry that governs which lane or proof claim is allowed

## Human Relief If Successful

The human should no longer have to repeat, within the same task:

- `no engine mods`
- `do not touch REGRESSION.md`
- `stop at approval gate`
- `supporting lanes do not substitute for the repo's default-lane proof`

Instead, the workflow should surface the recorded entry and either keep the move inside bounds or block it before the human has to intervene again.

## Acceptance Criteria

### Ledger Contract Criteria

- the task-owned `CONSTRAINTS.md` contract defines `Active Constraints`, `Repo-Local Truth`, and `Checkpoint Log`
- entries distinguish task-specific hard constraints from repo-local truths instead of merging them into vague notes
- the contract requires status and enforcement notes, so the ledger is actionable rather than archival

### Checkpoint Gate Criteria

- a checkpoint is required before a risky move that could violate a recorded entry
- the checkpoint must record the pending action, cited entry ids, and the decision `allowed`, `blocked`, or `needs_clarification`
- the workflow fails closed when a risky move is attempted without a relevant checkpoint

### Repo-Truth Injection Criteria

- the workflow requires task-relevant repo truths to be recorded in the ledger when they materially govern the current task
- when a risky move depends on repo-local truth, the checkpoint must cite the relevant repo-truth entry rather than relying on free-form memory or prompt reminders
- a supporting or alternate lane cannot be treated as satisfying a recorded default-lane truth unless the checkpoint explicitly says the move is not a regression-proof claim

### Burden-Reduction Criteria

- after a hard constraint is recorded once for the task, the next comparable risky move is blocked or redirected by the checkpoint without requiring the human to restate that same rule
- after a repo-local truth is recorded once for the task, the next comparable truth-sensitive move cites it at the checkpoint without requiring the human to reintroduce the repo truth
- the human can inspect `CONSTRAINTS.md` alone and see why a disallowed move did not proceed

## Proof Plan

Use at least these fixtures:

### Fixture 1: Mid-Task New Constraint

- the human states a new hard constraint such as `stop at approval gate`
- the next step would otherwise proceed past that gate

Expected result:

- the next checkpoint records the new constraint before the risky move
- the move is blocked or redirected without the human restating the rule again

### Fixture 2: Recorded Do-Not-Touch Boundary

- `CONSTRAINTS.md` already records a `do not touch REGRESSION.md` or equivalent boundary
- the next proposed move would edit that surface

Expected result:

- the checkpoint cites the relevant entry
- the move is marked `blocked` or `needs_clarification`
- the disallowed edit does not proceed

### Fixture 3: Repo-Truth-Sensitive Proof Move

- the ledger records a repo-local truth that the human default lane is the required regression lane
- the next proposed move tries to rely on a supporting lane for a claim that would read as regression proof

Expected result:

- the checkpoint cites the repo-truth entry
- the move is blocked as a regression-proof claim or is explicitly downgraded to supporting evidence only

## What Does Not Count

This task is not complete if:

- `CONSTRAINTS.md` exists but has no checkpoint log
- the ledger contains only raw notes with no status or enforcement note
- `ORCHESTRATION.md`, `TASK-LEADER.md`, `IMPLEMENTATION-LEADER.md`, or `REGRESSION-LEADER.md` merely say `remember constraints` without a fail-closed checkpoint
- repo-local truths still live only in reminders rather than in the task's durable ledger when they govern the task
- the workflow can still take a risky move without citing the relevant ledger entries
- pause-state wording is added, but the action-boundary checkpoint is still missing

## Remaining Uncertainty

- some repos may need small local conventions for deciding which repo truths are important enough to copy into the task-local ledger
- this first slice is justified as `CONSTRAINTS.md` plus enforcement in `ORCHESTRATION.md`, `TASK-LEADER.md`, `IMPLEMENTATION-LEADER.md`, and `REGRESSION-LEADER.md`; the local sources do not justify widening it further to a separate validator or global memory mechanism

These do not block the task writeup because the implementation home and first enforcement surfaces are now explicit.

## Falsifier

This task is wrong or incomplete if, after implementation:

- the human still has to restate the same recorded hard constraint during the same task before the next comparable risky move
- a risky move still proceeds without a checkpoint citing the relevant ledger entries
- repo-local truth is still absent from the checkpoint when the move clearly depends on it
- the ledger exists, but the burden remains because agents can ignore it without being blocked

## References

- burden driver `BD-006` in [`../BURDEN-ANALYSIS.md`](../BURDEN-ANALYSIS.md)
- problem framing and option boundary in [`../ORTHOGONAL-SOLUTIONS-MATRIX.md`](../ORTHOGONAL-SOLUTIONS-MATRIX.md)
- shared task-writing rules in [`../../../../../../../../Processes/TASK-CREATE.md`](../../../../../../../../Processes/TASK-CREATE.md)
- shared task-audit rules in [`../../../../../../../../Processes/TASK-AUDIT.md`](../../../../../../../../Processes/TASK-AUDIT.md)
