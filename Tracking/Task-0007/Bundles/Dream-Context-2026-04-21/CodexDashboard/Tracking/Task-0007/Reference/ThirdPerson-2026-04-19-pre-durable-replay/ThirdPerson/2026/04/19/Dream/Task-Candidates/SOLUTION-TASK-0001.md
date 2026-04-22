# Solution Task 0001

## Title

Save task ownership, stop reasons, and next-step state in `TASK-STATE.json`.

## Summary

The April 19 packet shows repeated cases where active work stopped and the human had to say `continue`, restate ownership, or explain that a failed step did not justify a pause.

This is a concrete implementation task. It will extend shared task state, the task-state schema, and the leader prompts so active work stays owned until one named stop reason applies.

## Goals

- make ownership and stop state durable on disk
- make recoverable failure route to the next bounded step instead of back to the human
- make pause, approval, and hard-block states explicit
- make it possible for a later agent to resume from task state without rereading the whole chat

## Non-Goals

- building a background watcher or scheduler
- solving wrong-lane proof, bad evidence, or wrong-answer-shape problems inside this task
- hiding continuity behavior in model-only memory

## Constraints And Baseline

- the shared task-state contract already lives under `C:\Users\gregs\.codex\Orchestration\`
- the change must be visible in shared docs and prompt rules, not only in reasoning
- the state must preserve real human gates such as explicit approval, explicit stop, and real missing input

## Proposed Changes

- extend [TASK-STATE.md](../../../../../../../../TASK-STATE.md) with these exact fields: `task_owner_mode`, `active_owner`, `ownership_started_at`, `ownership_basis`, `standing_instructions`, `standing_instruction_sources`, `standing_instruction_last_confirmed_at`, `allowed_stop_reasons`, `blocked_status`, `blocker_summary`, `blocker_evidence`, `last_failed_attempt`, `next_planned_attempt`, `progress_update_mode`, `last_progress_update_at`, `progress_update_threshold`, `continuation_state`, `resume_condition`, `wake_up_needed`, and `wake_up_reason`
- extend [TASK-STATE.schema.json](../../../../../../../../TASK-STATE.schema.json) with these exact enums:
  - `task_owner_mode`: `system_owns_until_block`, `human_holding_gate`, `paused_by_user`, `closed`
  - `allowed_stop_reasons`: `explicit_user_stop`, `explicit_approval_gate`, `real_external_blocker`, `real_missing_required_input`, `task_closed`
  - `blocked_status`: `not_blocked`, `soft_friction`, `recoverable_failure`, `hard_blocked`
  - `continuation_state`: `continue_autonomously`, `awaiting_user_gate`, `paused_until_condition`, `closed`
- update [TASK-LEADER.md](../../../../../../../../Prompts/TASK-LEADER.md) and `IMPLEMENTATION-LEADER.md` so they must write these fields when work starts, pauses, blocks, resumes, or reaches approval
- update [ORCHESTRATION.md](../../../../../../../../ORCHESTRATION.md) so the shared lifecycle names the same four continuation states and does not treat stopping as an implicit vibe

## Expected Resolution

- an active task shows who owns it and why
- a recoverable failure records the failed attempt and the next planned attempt
- a pause records the exact resume condition
- the human only needs to step in for a named approval gate, a real blocker, or an explicit stop

## What Does Not Count

- stronger wording with no new task-state fields
- a prompt hint that says `keep going` but leaves no durable state on disk
- a `blocked` field with no blocker evidence
- progress commentary that still makes the human decide whether work should continue

## Implementation Home

Keep task-owned planning, handoff, and testing artifacts under the eventual `Tracking/Task-<id>/`.

Implement the shared contract under:

- [TASK-STATE.md](../../../../../../../../TASK-STATE.md)
- [TASK-STATE.schema.json](../../../../../../../../TASK-STATE.schema.json)
- [ORCHESTRATION.md](../../../../../../../../ORCHESTRATION.md)
- [TASK-LEADER.md](../../../../../../../../Prompts/TASK-LEADER.md)
- `C:\Users\gregs\.codex\Orchestration\Prompts\IMPLEMENTATION-LEADER.md`

## Proof Plan

- check that the new keys and enum values exist in both the prose contract and schema
- check that the leader prompts say when each field must be written
- check that the shared workflow now distinguishes recoverable failure from hard block

## Acceptance Criteria

- [TASK-STATE.md](../../../../../../../../TASK-STATE.md) defines every field listed in `Proposed Changes`
- [TASK-STATE.schema.json](../../../../../../../../TASK-STATE.schema.json) validates every enum listed in `Proposed Changes`
- [TASK-LEADER.md](../../../../../../../../Prompts/TASK-LEADER.md) and `IMPLEMENTATION-LEADER.md` require `last_failed_attempt` plus `next_planned_attempt` after a recoverable failure
- [ORCHESTRATION.md](../../../../../../../../ORCHESTRATION.md) names the four continuation states `continue_autonomously`, `awaiting_user_gate`, `paused_until_condition`, and `closed`
- the prompt rules say that only `explicit_user_stop`, `explicit_approval_gate`, `real_external_blocker`, `real_missing_required_input`, and `task_closed` may stop owned work
- a reviewer can tell from the written contract when the system should keep going and when the human should re-enter the loop

## References

- [TASK-CREATE.md](../../../../../../../../Processes/TASK-CREATE.md)
- [BURDEN-ANALYSIS.md](../BURDEN-ANALYSIS.md)
- [ORTHOGONAL-SOLUTIONS-MATRIX.md](../ORTHOGONAL-SOLUTIONS-MATRIX.md)

## Plan Addendum

This task is the selected first implementation task from the current solution set.

Matrix source:

- Problem `0005`
- Option `B. Persisted continuity contract`

Canonical standalone plan:

- [SOLUTION-TASK-0001-PLAN-0001.md](../Plans/SOLUTION-TASK-0001-PLAN-0001.md)

### Planning Intent

This addendum turns the task into a bounded implementation sequence.

It describes the intended route to done, not pass history.

### Summary

Extend shared orchestration state and leader rules so active work stays owned until one named stop reason applies.

The implementation must land in three places together:

- the durable state contract
- the machine-checkable schema
- the leader and workflow rules that require those fields to be used

### Fixed Defaults

- scope: shared orchestration docs and prompts only
- canonical state home:
  - `C:\Users\gregs\.codex\Orchestration\TASK-STATE.md`
  - `C:\Users\gregs\.codex\Orchestration\TASK-STATE.schema.json`
- canonical workflow home:
  - `C:\Users\gregs\.codex\Orchestration\ORCHESTRATION.md`
  - `C:\Users\gregs\.codex\Orchestration\Prompts\TASK-LEADER.md`
  - `C:\Users\gregs\.codex\Orchestration\Prompts\IMPLEMENTATION-LEADER.md`
- stop reasons are limited to:
  - `explicit_user_stop`
  - `explicit_approval_gate`
  - `real_external_blocker`
  - `real_missing_required_input`
  - `task_closed`
- no background watcher or scheduler in this rollout

### Pass Plan

#### Pass 0000 - Durable State Contract

Goal:

- define the continuity contract in shared state so ownership and stopping are visible on disk

Build:

- extend `TASK-STATE.md` with the ownership, standing-instruction, stop-contract, progress, and continuation fields from the task
- extend `TASK-STATE.schema.json` with the matching keys and enum values
- add one compact example state block for:
  - active owned work
  - approval gate
  - hard block

Unit Proof:

- the prose contract and schema name the same keys and enum values
- the example state blocks only use allowed enum values

Exit Bar:

- a reviewer can read the state contract alone and tell who owns the task, why it may stop, and what the next step is after a recoverable failure

#### Pass 0001 - Leader And Workflow Adoption

Goal:

- make the continuity fields mandatory at the points where work actually pauses, resumes, blocks, or waits for approval

Build:

- update `TASK-LEADER.md` to require the continuity fields during lifecycle transitions
- update `IMPLEMENTATION-LEADER.md` to require the same fields during planning and pass execution
- update `ORCHESTRATION.md` so the shared lifecycle names the continuation states and allowed stop reasons directly

Unit Proof:

- each of the three docs names the same continuation states
- each of the three docs treats recoverable failure as continue-or-next-step, not implicit pause

Exit Bar:

- shared prompt and workflow prose now enforce the same continuity contract that the schema defines

#### Pass 0002 - Rollout Examples And Eval Hooks

Goal:

- make the contract easy to apply and easy to audit after rollout

Build:

- add one short example of a valid restart-supervision-saving transition to `TASK-STATE.md` or `ORCHESTRATION.md`
- add one short anti-pattern example showing what does not count, such as a `blocked` state with no blocker evidence
- preserve the April 19 packet-backed eval cases as rollout checks in the most natural shared location

Unit Proof:

- the final docs include at least one valid example and one invalid example
- the eval notes still map to the April 19 failure pattern

Exit Bar:

- a later reviewer can check a live task state and tell whether the continuity contract is being followed honestly

### Testing Strategy

- validate consistency across the prose contract, the schema, and the leader prompts
- treat mismatched enum names or missing required fields as rollout failures
- do not treat the existence of new fields as sufficient unless the leader docs also say when to write them

### Deferred Work

Keep these out of this rollout unless the task expands intentionally:

- background polling or auto-resume daemons
- hidden model-side memory
- approval-surface improvements
- answer-shape changes
- claim or evidence gating beyond the continuity contract
