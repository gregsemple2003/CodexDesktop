# Problem 0005 Option B Plan 0001

## Planning Intent

This file turns Problem `0005`, Option `B. Persisted continuity contract` from [ORTHOGONAL-SOLUTIONS-MATRIX.md](../ORTHOGONAL-SOLUTIONS-MATRIX.md) into a bounded implementation sequence.

It is the matrix option that fed [SOLUTION-TASK-0001.md](../Task-Candidates/SOLUTION-TASK-0001.md).

## Summary

Extend shared task state and leader rules so active work stays owned until one named stop reason applies.

## Fixed Defaults

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

## Pass Plan

### Pass 0000 - Durable State Contract

Goal:

- define the continuity contract in shared state so ownership and stopping are visible on disk

Build:

- extend `TASK-STATE.md` with ownership, standing-instruction, stop-contract, progress, and continuation fields
- extend `TASK-STATE.schema.json` with matching keys and enum values
- add one compact example state block for active work, approval gate, and hard block

Unit Proof:

- the prose contract and schema name the same keys and enum values
- the example state blocks only use allowed enum values

Exit Bar:

- a reviewer can read the state contract alone and tell who owns the task, why it may stop, and what the next step is after a recoverable failure

### Pass 0001 - Leader And Workflow Adoption

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

- shared prompt and workflow prose enforce the same continuity contract that the schema defines

### Pass 0002 - Rollout Examples And Eval Hooks

Goal:

- make the contract easy to apply and easy to audit after rollout

Build:

- add one valid restart-supervision-saving example and one invalid blocked-state example
- preserve the April 19 packet-backed eval cases as rollout checks in the most natural shared location

Unit Proof:

- the final docs include at least one valid example and one invalid example
- the eval notes still map to the April 19 failure pattern

Exit Bar:

- a later reviewer can check a live task state and tell whether the continuity contract is being followed honestly

## Testing Strategy

- validate consistency across the prose contract, the schema, and the leader prompts
- treat mismatched enum names or missing required fields as rollout failures

## Deferred Work

- background polling or auto-resume daemons
- hidden model-side memory
- approval-surface improvements
