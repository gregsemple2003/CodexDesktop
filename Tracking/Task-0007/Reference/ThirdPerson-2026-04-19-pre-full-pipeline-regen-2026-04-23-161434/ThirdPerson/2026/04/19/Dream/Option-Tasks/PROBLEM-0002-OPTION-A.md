# PROBLEM-0002 OPTION-A: Shared Execution-State Pause/Ownership Gate (Durable STOP + Resume Semantics)

## Title

Add a shared, machine-checkable execution-state contract so STOP, idle gaps, and approval gates become explicit durable control flow instead of chat-only intent.

## Summary

The packet shows a repeated continuity failure pattern:

- the human issues an explicit STOP or pause gate and the system continues anyway
- the system stops in an ambiguous state (no explicit "waiting" contract) and later requires the human to "wake it up"
- ownership drifts: intermediate failure is handed back to the human as "homework" even when a concrete next narrowing step exists

This winner proposes a shared execution-state gate that makes stop/resume/approval semantics durable, inspectable, and therefore enforceable across sessions.

## Writeup Type

Concrete implementation task (burden-reduction proposal).

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

The human is currently forced to do repeated supervision labor that should be handled by a durable execution-state contract:

- Repeating STOP: the human has to say "STOP" multiple times and still cannot trust that STOP is latched.
- Wake-up supervision: after idle gaps or dropped increments, the human must re-initiate work and re-establish ownership.
- "Homework" export: when progress stalls, the system implicitly hands the work back to the human by stopping without a durable "waiting / blocked / next step" state.

This burden is costly because it is time- and attention-bound: it interrupts the human’s work and forces real-time monitoring.

## Current Truth

Today, the shared orchestration system has:

- a lifecycle model in [../../../../../../../../ORCHESTRATION.md](../../../../../../../../ORCHESTRATION.md)
- a compact machine-readable task state contract in [../../../../../../../../TASK-STATE.md](../../../../../../../../TASK-STATE.md) and [../../../../../../../../TASK-STATE.schema.json](../../../../../../../../TASK-STATE.schema.json)

But the existing `TASK-STATE.json` contract does not currently encode:

- who owns the next step (human vs agent)
- whether work is allowed to continue right now (a continuation latch)
- what explicit resume condition is required (approval vs "human said continue" vs "waiting on tool")

So STOP and approval gates can exist only as chat intent. That makes them easy to ignore, hard to audit, and hard for later sessions to obey consistently.

## Target Truth

After this change, a cold reader (human or agent) should be able to open `Tracking/Task-<id>/TASK-STATE.json` and answer:

- who currently owns the next step
- whether execution is active, waiting, blocked, or paused
- whether continuation is latched off due to STOP or an approval gate
- what explicit resume condition is required before work may proceed

If execution is latched off, later sessions should not proceed without first changing the durable state (and therefore acknowledging the gate explicitly).

## Causal Claim

If the shared task-state contract encodes owner, execution run-state, continuation latch, and resume condition, and the shared lifecycle rules treat those fields as binding control flow, then:

- STOP and approval gates stop being "messages" and become durable gates
- wake-up supervision decreases because "waiting" and "blocked" states are explicitly recorded instead of silently implied
- ownership drift decreases because the task’s durable state makes it explicit when the agent still owns continuation versus when the human is required

## Evidence

The evidence for this burden is concentrated in [../BURDEN-ANALYSIS.md](../BURDEN-ANALYSIS.md) (`BD-002`) and the stable `event_id` entries in [../../HumanInputEvents/INDEX.json](../../HumanInputEvents/INDEX.json), including:

- explicit STOP/pause gates (`...-3447`, `...-3692`, `...-5108`)
- explicit "continue under your existing ownership; do not stop after a failed attempt" supervision (`...-7038`)
- explicit wake-up / continuity nudges after a gap (`...-3929`, `...-4590`, `...-6558`)

## Why This Mechanism

This mechanism is chosen because the burden is fundamentally a *durable control-flow* problem:

- a STOP gate is not advice; it is a latch that should prevent further action
- an idle gap is not "done"; it should be represented as waiting/blocked with a reason and a next resume condition
- ownership is not vibe; it is a contract about who must act next

Adding machine-checkable execution-state fields is the narrowest shared intervention that makes these truths durable and inspectable across sessions, tools, and delegated agents.

## Scope Rationale

This is intentionally shared scope because the failure mode is cross-repo and lifecycle-wide.

- Narrower, repo-local fixes do not prevent the same STOP/continuity failures in other repos.
- Broader fixes (a new global memory system or background scheduler) are not required for this first boundary; the packet does not justify that expansion.

## Goals

- Make STOP and approval gates durable, inspectable, and fail-closed across sessions.
- Make "waiting" and "blocked" states explicit so the human does not have to wake work up.
- Make ownership explicit so the system does not export intermediate failure back to the human without a durable contract.

## Non-Goals

- Rewriting all prompt templates or building a new orchestration runtime in this slice.
- Adding a notification-only alerting system; alerts without durable state do not prevent the STOP/continuity failure.
- Solving task-local "remember constraints" problems (that is `P-006`).

## Implementation Home

- Shared lifecycle rules:
  - [../../../../../../../../ORCHESTRATION.md](../../../../../../../../ORCHESTRATION.md)
- Shared task-state contract:
  - [../../../../../../../../TASK-STATE.md](../../../../../../../../TASK-STATE.md)
  - [../../../../../../../../TASK-STATE.schema.json](../../../../../../../../TASK-STATE.schema.json)
- Task-owned current-state artifacts that must reflect the contract:
  - `Tracking/Task-<id>/TASK-STATE.json`
  - `Tracking/Task-<id>/HANDOFF.md`

## Implementation Home Rationale

- The lifecycle meaning of STOP, approval gates, waiting, and resume is cross-repo truth and belongs in shared orchestration docs.
- The fields must be machine-checkable (schema + contract) to avoid drift and to make later sessions/tools able to interpret state.
- Task-owned artifacts (`TASK-STATE.json` + `HANDOFF.md`) are the durable per-task surfaces that actually carry the current state forward.

## Constraints And Baseline

- Keep the existing `TASK-STATE.json` purpose: compact current orchestration state, not narrative history.
- Preserve the split: markdown holds narrative/evidence; JSON holds current state.
- Keep the initial field set small: encode what is required to make STOP/ownership/wait/resume inspectable without turning `TASK-STATE.json` into a new journal.

## Proposed Changes

These are the concrete, reviewable surfaces this winner changes.

1. **Shared task-state schema expansion (new execution-state contract fields)**
   - Files:
     - [../../../../../../../../TASK-STATE.md](../../../../../../../../TASK-STATE.md)
     - [../../../../../../../../TASK-STATE.schema.json](../../../../../../../../TASK-STATE.schema.json)
   - Add an `execution` object (or equivalent top-level fields) that encodes, at minimum:
     - `owner`: who owns the next step (`human` or `agent`)
     - `run_state`: execution run-state (for example: `active`, `waiting`, `paused`)
     - `wait_reason`: short string for why we are not proceeding (required when `run_state != active`)
     - `continuation_latch`: whether continuation is allowed (`open`) or latched off (`stop_latched`, `approval_latched`)
     - `resume_condition`: explicit condition required to reopen continuation (for example: `human_continue`, `human_approve_plan`, `human_approve_pass`, `tool_finished`, `external_dependency`)
   - Update schema rules:
     - forbid silent "stopped" states: if `run_state != active`, require `wait_reason` and `resume_condition`
     - make latch semantics explicit: if `continuation_latch != open`, require `owner = human` and require an explicit `resume_condition` that is human-mediated
     - avoid "two truths" with existing task-state fields:
       - `status/phase/current_gate` remain task-level lifecycle truth (unchanged enums)
       - `execution.*` is run-control truth (stop/wait/resume); it must not claim `run_state = active` when the task is `status = blocked|complete|cancelled`
       - if the task is `status = blocked`, `execution.run_state` must be `waiting` or `paused`, and `execution.wait_reason` must explain the block
   - Versioning requirement:
     - either bump `schema_version` and support both versions via `oneOf`, or explicitly document the migration rule for existing tasks (do not silently invalidate all historical task state files).
2. **Shared lifecycle enforcement rule for STOP / approval gates**
   - File: [../../../../../../../../ORCHESTRATION.md](../../../../../../../../ORCHESTRATION.md)
   - Add a new section (or extend existing gate rules) that requires:
     - when the human issues STOP (or an explicit "approval gate only"), the lead must persist a `TASK-STATE.json` transition that latches continuation off before any further action
     - when a delegated agent is running, the supervising leader must either remain in an explicit `waiting` run-state (with a polling loop) or record a `waiting/paused` run-state with a `wait_reason` and `resume_condition`; "silence" is not a valid state
     - resume is only allowed after the durable latch is reopened (human says continue / approval granted / explicit resume condition satisfied)
3. **Task-owned artifact updates at gate boundaries**
   - Artifacts:
     - `Tracking/Task-<id>/TASK-STATE.json`
     - `Tracking/Task-<id>/HANDOFF.md`
   - Require that the human-readable handoff contains a brief, matching execution-state summary whenever the task is latched off or waiting (so the human does not have to infer state from chat history).
4. **Shared exemplar refresh for the new contract**
   - File:
     - [../../../../../../../../Exemplars/TASK-STATE.json](../../../../../../../../Exemplars/TASK-STATE.json)
   - Update the exemplar to demonstrate the new execution fields and a STOP-latched or approval-latched state.

## Acceptance Criteria

- After an idle gap, a cold reader can open `TASK-STATE.json` and see:
  - who owns the next step
  - whether execution is waiting/paused (and whether the task is blocked)
  - why it is not proceeding
  - what explicit condition must happen next to resume
- After STOP, continuation is latched off in durable state and later sessions do not proceed without explicitly reopening the latch.
- The human no longer needs to send "wake up" or repeated ownership reminders for the common stop/resume paths covered by this contract.

## Expected Resolution

Human-facing outcome:

- STOP and approval gates feel trustworthy: the system visibly "latches" and later sessions respect the latch.
- Pauses become understandable: state says *why* it is waiting and *what must happen next*, reducing human supervision.
- Ownership becomes explicit: the system does not export intermediate failure as "homework" without first recording what it is waiting for and what the next step is.

## Human Relief If Successful

- Fewer interruptive "status check" and "why are you messaging me" moments caused by ambiguous stop/wait states.
- Less repeated STOP prompting.
- Less time spent reconstructing whether a delegated agent is actually being supervised versus "spawn and forget."

## Internal Mechanism Map

1. Add explicit execution-state fields to the shared task-state schema.
2. Make shared lifecycle rules treat those fields as control flow, not passive metadata.
3. Require durable state updates at STOP, approval gates, and intentional waiting.
4. Use the task handoff as the human-facing mirror of the same execution truth.

## Rival Explanations Considered

- "This is just a one-off subagent stall."
  - Rejected: the packet contains multiple separate interventions around STOP, wake-up, and ownership drift.
- "The human should just tolerate monitoring."
  - Rejected: the workflow is human-facing; exporting supervision labor is exactly the burden to reduce.

## Rival Mechanisms Considered

- `P-002 / Option B` (workflow/prompt discipline without schema expansion):
  - Closest dangerous rival: cheaper and acts directly on worker behavior.
  - Rejected as winner: without durable state fields, later sessions/tools still cannot tell whether continuation is allowed, and STOP semantics remain easy to lose or misapply.
- Notification-only alerting:
  - Rejected: it detects failure after the human already paid the wake-up cost and does not make STOP fail closed.

## Tradeoffs

- Schema churn and migration cost:
  - This is real and must be handled explicitly (versioning/migration), not hand-waved.
- Risk of over-encoding:
  - If too many fields are added, `TASK-STATE.json` may drift into narrative journaling; the contract must stay compact.

## Shared Substrate

- Existing lifecycle model: [../../../../../../../../ORCHESTRATION.md](../../../../../../../../ORCHESTRATION.md)
- Existing task-state contract and schema: [../../../../../../../../TASK-STATE.md](../../../../../../../../TASK-STATE.md), [../../../../../../../../TASK-STATE.schema.json](../../../../../../../../TASK-STATE.schema.json)
- Task handoff concept and structure (supporting reference): [../../../../../../../../Exemplars/HANDOFF.md](../../../../../../../../Exemplars/HANDOFF.md)

## Not Solved Here

- Task-local constraint retention (`P-006`).
- Approval-packet review surface (`P-004`).
- Direct-answer-first conversational contract (`P-003`).

## What Does Not Count

- A chat message saying "paused" without a durable state update.
- Saying "I’m watching the subagent" while not in an explicit wait/polling loop and without durable state reflecting the wait.
- Treating STOP as "a suggestion" and continuing to act.
- Encoding stop/wait state only in prose in `HANDOFF.md` while leaving `TASK-STATE.json` ambiguous.

## Remaining Uncertainty

- Exact field naming and minimal enum set that will be stable across repos (for example `paused` vs `waiting`).
- Migration policy: how existing tasks with `schema_version = 1` will be treated by validators and tools without breaking history.

## Falsifier

This proposal is falsified if, after implementation:

- STOP or approval gates are still ignored in practice, even when the durable state indicates continuation should be latched off.
- The human still needs to "wake up" work after idle gaps because durable state does not make waiting/blocked truth explicit.
- Later sessions still cannot tell whether work may proceed without rereading chat history.

## Proof Plan

1. Update `TASK-STATE.md` and `TASK-STATE.schema.json`, including explicit versioning/migration rules.
2. Update `ORCHESTRATION.md` to treat the new fields as binding control flow.
3. Add two exemplar `TASK-STATE.json` fixtures (in a real implementing task folder) that demonstrate:
   - STOP-latched state: `continuation_latch != open`, `owner = human`, explicit resume condition
   - waiting-on-tool state: `run_state = waiting`, `wait_reason` + `resume_condition = tool_finished`
4. Validate those fixtures against the updated JSON schema.
5. Run a small rehearsal scenario on a real task:
   - human issues STOP
   - leader persists state latch
   - later session attempts to proceed and is forced (by process) to reopen latch before acting

## Open Questions

- Should "approval latch" be a distinct latch type from STOP, or should both be represented as the same continuation-latch mechanism with different resume conditions?
- Should a delegated-agent supervision "polling loop" state be represented explicitly (for example `execution.run_state = waiting` with `resume_condition = agent_finished`) or treated as a separate mechanism?

## References

- Burden driver: [../BURDEN-ANALYSIS.md](../BURDEN-ANALYSIS.md) (`BD-002`)
- Designed options: [../SOLUTION-DESIGN.md](../SOLUTION-DESIGN.md#p-002-ownership-continuity-and-pause-latch-control)
- Frozen winner boundary: [../WINNER-SYNTHESIS.md](../WINNER-SYNTHESIS.md#w-002-shared-execution-state-pauseownership-gate)
- Final matrix row: [../ORTHOGONAL-SOLUTIONS-MATRIX.md](../ORTHOGONAL-SOLUTIONS-MATRIX.md#p-002-ownership-continuity-and-pause-latch-control)
- Shared orchestration lifecycle: [../../../../../../../../ORCHESTRATION.md](../../../../../../../../ORCHESTRATION.md)
- Shared task-state contract: [../../../../../../../../TASK-STATE.md](../../../../../../../../TASK-STATE.md)
