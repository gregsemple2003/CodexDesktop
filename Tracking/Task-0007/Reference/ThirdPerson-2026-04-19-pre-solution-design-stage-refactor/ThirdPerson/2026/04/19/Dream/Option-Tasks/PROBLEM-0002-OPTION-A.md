# PROBLEM-0002 OPTION-A: Durable Wait-State Contract And Stop-Boundary Enforcement

## Title

Require a durable wait-state contract before a task-owning leader may stop, and enforce continuity plus pause-gate latching at the shared stop boundary.

## Summary

The burden packet shows three linked supervision failures:

- work dropped after failed increments or handoff boundaries even though ownership should have continued
- idle pauses happened without a durable explanation of what was blocking progress or what would happen next
- explicit human STOP or approval-only gates were not always treated as hard latches

This rewrite narrows Option A to the exact first slice justified by the local durable sources:

- extend the shared `TASK-STATE.json` contract so a task can represent explicit wait state, explicit continuation under ownership, and a cross-phase pause gate
- enforce those fields at the stop boundary in the shared lifecycle and prompt files that already own continuation, handoff, waiting, and human-gate behavior

This slice does **not** include a dashboard or UI surface.
The local sources justify shared state plus shared workflow enforcement.
They do not justify a first-slice UI surface honestly.

## Writeup Type

Concrete implementation task.

The local durable sources already justify the burden, the first enforcement seam, and the exact shared docs and prompt files that own task stopping, leader continuity, and human gates.

## Source Event IDs

```json
[
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-3447",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-3692",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-3929",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4015",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4590",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-4919",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5108",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-6558",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5272",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5300",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5310",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-7038"
]
```

## Burden Being Reduced

The human is currently forced to do three kinds of continuity-repair work by hand:

1. Wait-state policing.
   The human must infer whether the system is honestly waiting on a real human gate or external blocker, or whether it just drifted into an idle ambiguous stop.
2. Continuity policing.
   The human must restate `continue`, `do not stop`, or `keep working under your ownership` after a failed incremental attempt or a designed handoff boundary.
3. Pause-gate policing.
   The human must restate that STOP or approval-only instructions are real gates and must hold until explicitly released.

The deeper burden is exported supervision.
A later session cannot reliably tell whether the task is paused for a real reason, should already be continuing, or was improperly advanced past a human gate.

## Current Truth

The shared workflow already contains the right pieces in isolation:

- `ORCHESTRATION.md` already says standing instructions such as `continue until complete` survive delegated-leader rotation unless a real human gate or real blocker interrupts them
- `ORCHESTRATION.md` already says task pauses need a reliable resume point with an explicit next step and visible blockers
- `TASK-STATE.md` already records coarse task status, blocking issues, and `next_expected_artifacts`
- `TASK-LEADER.md`, `IMPLEMENTATION-LEADER.md`, and `REGRESSION-LEADER.md` already describe waiting at real human gates, preserving standing ownership instructions, and not `spawn and forget`

But `BD-002` shows that this is still not enough in practice.

The current system truth is therefore not merely `the agent should have kept going`.
The current truth is that the shared workflow lacks one durable, machine-readable stop contract at the exact point where a task-owning leader is about to stop.

That creates three different failure modes:

- a task can stop without making clear whether it is in a real wait state or just dropped
- a task can stop after a failed attempt even though the standing ownership instruction still implies continued work
- a human STOP or approval-only instruction can remain only in chat or prose and fail to act as a durable cross-phase latch

The current fallback truth is only that the human can wake the system back up, restate ownership, or restate the pause gate.
That is a rescue path, not a durable stop boundary.

## Target Truth

Whenever a task-owning leader is about to stop while the task remains open, a later reader should be able to inspect `TASK-STATE.json` and know immediately which of these is true:

- the task is in an explicit human-gate wait
- the task is in an explicit external-blocker wait
- the task is still expected to continue under existing ownership and the next continuation step is already named
- a pause gate is latched and no further work may proceed until the human releases it

A later session should not have to infer this from commentary or raw chat alone.

The durable state should answer:

- why the task is waiting, if it is waiting
- what exact condition releases that wait
- what next step is expected if the task should continue
- whether an explicit human pause gate is latched

When the standing instruction is still `keep going unless blocked`, the task should not honestly stop in an ambiguous middle state.

## Causal Claim

If the shared workflow forces every task-owning leader through a durable stop contract before it may stop, then ambiguous idle drops, continuity loss after failed increments, and pause-gate drift will all reduce materially.

The cause being addressed is not only weak reminders.
The cause is that the workflow still lacks a single fail-closed stop boundary that distinguishes:

- explicit wait
- explicit continuation under ownership
- explicit human pause latch

## Evidence

`BD-002` in [`../BURDEN-ANALYSIS.md`](../BURDEN-ANALYSIS.md) identifies all three burden slices directly:

- wake-up supervision after dropped work
- repeated `continue under your ownership` correction
- explicit STOP or approval-only gates that needed re-imposition

`../ORTHOGONAL-SOLUTIONS-MATRIX.md` already fixes the mechanism boundary for `PROBLEM-0002` at the stop or resume boundary:

- when the system is not actively executing, it must be in an explicit durable wait state or it must persist a concrete next step and resumption plan

The shared lifecycle and prompt docs already identify the exact enforcement surfaces that own this boundary:

- `ORCHESTRATION.md` defines standing-instruction continuity, active monitoring, and pause or resume expectations at handoff boundaries
- `TASK-STATE.md` and `TASK-STATE.schema.json` define the machine-readable current-state contract
- `TASK-LEADER.md`, `IMPLEMENTATION-LEADER.md`, and `REGRESSION-LEADER.md` own the task-level stop, wait, and continuation behavior for the main shared phases where the burden occurred

That is enough to justify a first shared state-and-enforcement task without inventing a UI surface.

## Why This Mechanism

The first durable intervention should harden the shared stop boundary itself.

This mechanism is chosen because it acts at the exact place where the burden currently escapes:

- right before a task-owning leader stops
- right before a designed handoff boundary is treated as a task pause
- right before a human STOP or approval-only instruction could be forgotten in a later phase

The missing piece is not a notification surface.
The missing piece is a durable, enforced contract for what counts as an honest stop.

## Scope Rationale

This rewrite intentionally combines three internal mechanisms in one first task:

1. explicit wait-state representation
2. explicit continuation-under-ownership representation
3. explicit pause-gate latch behavior

That merge is earned because all three mechanisms act at the same stop boundary and each is weak without the others:

- if explicit wait exists without continuity enforcement, the task can still stop after a failed increment even though no real gate exists
- if continuity enforcement exists without a durable wait contract, the difference between a real wait and a dropped task remains ambiguous
- if pause-gate behavior is left outside the same state and prompt boundary, STOP and approval-only instructions can still drift out of force across phase rotation

The rewrite narrows away dashboard or UI work.
The local sources do not justify a first-slice display surface.
They justify shared state plus shared prompt and workflow enforcement only.

## Goals

- require a durable explicit-wait representation before a task-owning leader may stop for a real human gate or external blocker
- require a durable next continuation step when standing ownership still implies continued work
- make human STOP or approval-only instructions durable cross-phase pause latches
- block ambiguous task stops at shared leader handoff boundaries
- reduce wake-up supervision and repeated `continue under your ownership` corrections

## Non-Goals

- building a watchdog or auto-resume service
- adding a dashboard or other UI surface in this first slice
- solving the underlying runtime or product defects that happened to trigger the pauses
- redefining repo-local operator lanes or regression lanes
- replacing task-owned `HANDOFF.md` or other narrative artifacts with state-only automation

## Implementation Home

Primary shared state-contract home:

- `C:\Users\gregs\.codex\Orchestration\TASK-STATE.md`
- `C:\Users\gregs\.codex\Orchestration\TASK-STATE.schema.json`

Shared lifecycle home:

- `C:\Users\gregs\.codex\Orchestration\ORCHESTRATION.md`

Shared prompt enforcement home:

- `C:\Users\gregs\.codex\Orchestration\Prompts\TASK-LEADER.md`
- `C:\Users\gregs\.codex\Orchestration\Prompts\IMPLEMENTATION-LEADER.md`
- `C:\Users\gregs\.codex\Orchestration\Prompts\REGRESSION-LEADER.md`

Task-owned durable artifact home:

- `Tracking/Task-<id>/TASK-STATE.json`
- `Tracking/Task-<id>/HANDOFF.md`

Dashboard or UI home in this first slice:

- none

## Implementation Home Rationale

This does not belong primarily in a dashboard or monitoring UI.

The burden is not that the state is hidden on a screen.
The burden is that the workflow can still stop ambiguously before any durable state exists to display.

It also does not belong only in one prompt file.
The burden crosses the top-level task leader, the implementation leader, and the regression leader.
The shared state contract and the shared lifecycle doc must define the stop boundary once, and the phase-owning prompts must enforce it consistently.

`TASK-STATE.md` and `TASK-STATE.schema.json` are the right contract home because they already own machine-readable current orchestration state.
`ORCHESTRATION.md` is the right lifecycle home because it already owns continuity across leader rotation, active monitoring, task pauses, and resume expectations.
`TASK-LEADER.md`, `IMPLEMENTATION-LEADER.md`, and `REGRESSION-LEADER.md` are the right enforcement home because they already own:

- when a leader waits
- when a leader preserves standing ownership instructions
- when a leader stops at a designed phase boundary
- when a leader must surface a human gate instead of guessing

`HANDOFF.md` remains a supporting task-owned resume artifact, but it is not the machine-readable contract.
That contract belongs in `TASK-STATE.json`.

## Internal Mechanism Map

### Mechanism 1: Explicit Wait-State Contract

Failure reduced:

- a task stops without making clear whether it is truly waiting on a human gate or external blocker

Mechanism:

- `TASK-STATE.md` and `TASK-STATE.schema.json` add explicit stop-state fields:
  - `explicit_wait_state` boolean
  - `wait_reason` enum: `none`, `human_gate`, `external_blocker`
  - `wait_detail` string or `null`
  - `resume_condition` string or `null`

Acceptance focus:

- a real wait is machine-readable and says why the task is waiting and what releases it

Falsifier:

- a task can still stop with `status: in_progress` and no explicit wait reason or release condition

### Mechanism 2: Continuation-Under-Ownership Contract

Failure reduced:

- the task stops after a failed attempt or designed handoff boundary even though the standing instruction still implies continued work

Mechanism:

- `TASK-STATE.md` and `TASK-STATE.schema.json` add:
  - `next_continuation_step` string or `null`
- `ORCHESTRATION.md` and the leader prompts require that when no real wait gate exists, the leader must keep ownership active and persist the next continuation step instead of stopping ambiguously

Acceptance focus:

- a later reader can tell what the leader is expected to do next when work should continue without more human prompting

Falsifier:

- a task still falls idle after a failed increment with no durable next continuation step even though no real blocker or human gate exists

### Mechanism 3: Pause-Gate Latch

Failure reduced:

- STOP or approval-only instructions can be lost across phase changes or resumed work

Mechanism:

- `TASK-STATE.md` and `TASK-STATE.schema.json` add:
  - `pause_gate_latched` boolean
- `ORCHESTRATION.md`, `TASK-LEADER.md`, `IMPLEMENTATION-LEADER.md`, and `REGRESSION-LEADER.md` treat that latch as a cross-phase stop condition until the human explicitly releases it

Acceptance focus:

- once latched, work does not continue through ordinary handoff logic or standing continuity instructions

Falsifier:

- a leader still resumes ordinary work while the pause gate is latched and no explicit human release exists

## Proposed Changes

### 1. Extend the shared task-state contract for the stop boundary

Update `C:\Users\gregs\.codex\Orchestration\TASK-STATE.md` and `C:\Users\gregs\.codex\Orchestration\TASK-STATE.schema.json` to add these exact fields:

- `explicit_wait_state` boolean
- `wait_reason` enum with:
  - `none`
  - `human_gate`
  - `external_blocker`
- `wait_detail` string or `null`
- `resume_condition` string or `null`
- `next_continuation_step` string or `null`
- `pause_gate_latched` boolean

Keep the existing fields such as `status`, `current_gate`, `blockers`, and `next_expected_artifacts`.
Do not replace them.

Define the field contract so that:

- `explicit_wait_state: true` requires `wait_reason != none`
- `wait_reason: human_gate` requires `wait_detail` to name the blocking human decision, question, or approval release condition
- `wait_reason: external_blocker` requires `wait_detail` to name the blocker and `resume_condition` to state how resolution will be detected
- `explicit_wait_state: false` requires `wait_reason: none`
- when `status: in_progress` and `explicit_wait_state: false`, `next_continuation_step` must be populated
- `pause_gate_latched: true` means ordinary continuation is forbidden until a human explicitly releases the gate

### 2. Define the exact about-to-stop boundary in shared lifecycle rules

Update `C:\Users\gregs\.codex\Orchestration\ORCHESTRATION.md` to define `about to stop` as:

- a task-owning leader or top-level supervisor is about to end its active turn, stop monitoring delegated work, or hand control back at a designed boundary
- while the task is still open and the next honest lifecycle move is not already being actively executed by a delegated worker

At that boundary, the lifecycle must fail closed:

- if a real human gate or external blocker exists, persist the explicit wait state before stopping
- if no real gate exists and standing ownership still implies continuation, persist `next_continuation_step` and continue rather than stopping
- if `pause_gate_latched` is true, do not continue ordinary work until the human explicitly releases it

### 3. Enforce the stop contract in the named shared prompt files

Update these exact prompt files so the stop contract becomes result-affecting:

- `C:\Users\gregs\.codex\Orchestration\Prompts\TASK-LEADER.md`
  - require standing umbrella instructions to survive delegated rotation unless the state records a real wait or latched pause gate
  - forbid ambiguous top-level stop while `TASK-STATE.json` lacks the required stop or continuation fields
- `C:\Users\gregs\.codex\Orchestration\Prompts\IMPLEMENTATION-LEADER.md`
  - forbid stopping after a failed attempt, support-agent return, or designed pass boundary without either explicit wait state or explicit continuation state
  - require approval-only or STOP instructions to set `pause_gate_latched` before stopping
- `C:\Users\gregs\.codex\Orchestration\Prompts\REGRESSION-LEADER.md`
  - forbid ambiguous stop after blocked or incomplete regression slices
  - require the same explicit wait or continuation contract when regression stays open

## Rival Mechanisms Considered

### Rival 1: Dashboard or UI display first

Why not in this first slice:

- the local sources do not justify a shared dashboard or UI surface as the first intervention home
- a display surface would still be downstream of the real problem if the workflow can stop before durable state is written

### Rival 2: Prompt-only reminders without state changes

Why not first:

- the burden already persisted despite workflow prose about continuity and waiting
- the missing piece is a machine-readable stop contract that those prompts must enforce

### Rival 3: Watchdog or auto-resume service

Why not first:

- that is already reserved as Option B
- the first cheaper and truer intervention is to make ambiguous stops impossible to represent honestly before adding active detection machinery

## Not Solved Here

This task does not:

- build the watchdog or auto-resume path from Option B
- guarantee that every idle drop is recovered automatically
- eliminate the need for human approval gates
- solve the underlying runtime, regression, or debugging burdens that may create a legitimate blocker
- add a dashboard surface for continuity state

It only hardens the durable stop boundary and its enforcement.

## Human Relief If Successful

The human should no longer have to say, for the same burden class:

- `continue under your existing ownership`
- `do not stop after a failed attempt`
- `why did you stop`
- `STOP means stop`

A later session should be able to inspect durable task state and know immediately whether the task is truly waiting, truly continuing, or truly paused by a human gate.

## Acceptance Criteria

### Contract Criteria

- `C:\Users\gregs\.codex\Orchestration\TASK-STATE.md` and `TASK-STATE.schema.json` define the exact stop-boundary fields:
  - `explicit_wait_state`
  - `wait_reason`
  - `wait_detail`
  - `resume_condition`
  - `next_continuation_step`
  - `pause_gate_latched`
- the contract defines valid combinations for real wait, real continuation, and real pause-gate latch behavior

### Lifecycle Enforcement Criteria

- `C:\Users\gregs\.codex\Orchestration\ORCHESTRATION.md` defines the exact `about to stop` condition and requires the stop boundary to fail closed
- the lifecycle rules do not allow a task-owning leader to stop while the task remains open unless one of these is true:
  - a real explicit wait state is persisted
  - a real pause gate is latched
  - the next continuation step is persisted and work remains under active ownership rather than stopping

### Prompt Enforcement Criteria

- `TASK-LEADER.md`, `IMPLEMENTATION-LEADER.md`, and `REGRESSION-LEADER.md` all align on the same stop contract
- those prompt files do not allow:
  - stopping after a failed increment or designed handoff boundary with no real wait and no next continuation step
  - silently dropping standing `continue until complete` style instructions when no real blocker or human gate exists
  - continuing ordinary work while `pause_gate_latched` is still true

### Artifact Criteria

- for the next comparable human STOP or approval-only gate, `Tracking/Task-<id>/TASK-STATE.json` records `pause_gate_latched: true` before the leader stops
- for the next comparable external blocker pause, `TASK-STATE.json` records a real wait reason plus a resume condition
- for the next comparable continuing task with no real blocker, `TASK-STATE.json` records `next_continuation_step` rather than leaving the task in an ambiguous stop state

### Burden-Reduction Criteria

- the next comparable multi-hour pass does not require a human wake-up message merely to restore ongoing ownership when no real blocker or human gate exists
- explicit STOP or approval-only instructions hold across shared leader rotation until released
- the human does not need to restate both:
  - `continue under your ownership`
  - `STOP means stop`

## Proof Plan

Use at least these fixtures:

### Fixture 1: Real Human Pause Gate

- the human issues STOP or an approval-only instruction
- the task is still otherwise runnable

Expected workflow result:

- `pause_gate_latched` becomes true in `TASK-STATE.json`
- the leader stops honestly in an explicit wait state
- no ordinary continuation occurs until the human releases that gate

### Fixture 2: Failed Increment With Standing Continuation

- a pass or regression slice fails meaningfully
- there is no real human gate and no real external blocker
- the standing instruction still implies continued ownership

Expected workflow result:

- the leader does not stop ambiguously
- `next_continuation_step` is persisted
- the workflow continues under existing ownership rather than waiting for a human wake-up

### Fixture 3: Real External Blocker

- the task cannot proceed because of a named external dependency

Expected workflow result:

- `explicit_wait_state` is true
- `wait_reason` is `external_blocker`
- `wait_detail` names the blocker
- `resume_condition` states how the blocker resolution will be detected

## What Does Not Count

This task is not complete if:

- the schema adds fields but the shared prompts still allow ambiguous stop states
- `pause_gate_latched` exists but acts only as planning prose instead of a cross-phase stop latch
- a task can still stop with `status: in_progress` and no explicit wait reason, no explicit continuation step, and no active delegated execution
- standing continuity instructions still disappear at a designed handoff boundary without a real blocker or human gate
- a dashboard or UI surface is added while the stop contract and prompt enforcement remain weak

## Remaining Uncertainty

- the exact enum names can still be tightened during implementation if the final schema wording stays equivalent to the contract above
- some repos may want repo-local conventions for how detailed `next_continuation_step` should be, but the shared stop boundary itself is already justified
- Option B may still be useful later for active idle detection even after ambiguous stop states are made invalid

These do not block the task writeup because the draft now names:

- the exact shared state contract home
- the exact shared lifecycle and prompt enforcement homes
- the exact first-slice mechanism without any unsupported UI scope

## Falsifier

This task is wrong or incomplete if, after implementation:

- a task still falls idle and needs a human `continue` or `resume` message even though no real blocker or human gate existed
- a task still proceeds past a human STOP or approval-only instruction without an explicit release
- `TASK-STATE.json` can still represent an open task with no explicit wait, no explicit continuation step, and no active delegated execution
- the shared prompts still treat continuity and pause-gate behavior as advisory wording rather than a real stop boundary

## References

- burden driver `BD-002` in [`../BURDEN-ANALYSIS.md`](../BURDEN-ANALYSIS.md)
- problem framing in [`../ORTHOGONAL-SOLUTIONS-MATRIX.md`](../ORTHOGONAL-SOLUTIONS-MATRIX.md)
- shared task-writing rules in [`../../../../../../../../Processes/TASK-CREATE.md`](../../../../../../../../Processes/TASK-CREATE.md)
- shared task-audit rules in [`../../../../../../../../Processes/TASK-AUDIT.md`](../../../../../../../../Processes/TASK-AUDIT.md)
- shared lifecycle rules in [`../../../../../../../../ORCHESTRATION.md`](../../../../../../../../ORCHESTRATION.md)
- shared task-state contract in [`../../../../../../../../TASK-STATE.md`](../../../../../../../../TASK-STATE.md)
- shared task-state schema in [`../../../../../../../../TASK-STATE.schema.json`](../../../../../../../../TASK-STATE.schema.json)
