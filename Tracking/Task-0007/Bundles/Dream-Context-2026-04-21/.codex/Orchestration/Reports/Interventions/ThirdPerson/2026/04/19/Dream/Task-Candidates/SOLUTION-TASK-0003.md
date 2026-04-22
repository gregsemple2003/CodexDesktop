# Solution Task 0003

## Title

Block premature stops with a stop-eligibility gate and explicit blocker record

## Summary

The packet's clearest measured stall loss came from work that stopped before the human's done bar and then needed wake-up messages.
This task adds one stop-eligibility gate so active work cannot stop cleanly while the next justified action is known and no real external blocker exists.

## Goals

- Reduce wake-up and resume-prod messages on open tasks.
- Make the reason for every stop or pause explicit in task state and closeout artifacts.
- Prevent checkpoint-style stops from being treated as legitimate closure.

## Non-Goals

- Adding a menu of next steps for the human to pick from.
- Auto-resuming every long-running task regardless of blocker state.
- Replacing real approval stops or external dependency stops.

## Constraints And Baseline

Current truth:

- The packet records `1647.027` seconds and `452.282` seconds of charged stall loss on resume-prod events.
- The human had to say `Continue ... now` because the system stopped at its own checkpoint instead of the real done bar.

Hard constraints:

- a valid stop must name a real external block or completed done bar
- known next action plus no blocker means the task should continue
- the gate must treat wake-up turns as failure telemetry, not normal coordination

## Implementation Home

- `C:\Users\gregs\.codex\Orchestration\ORCHESTRATION.md`
- `C:\Users\gregs\.codex\Orchestration\TASK-STATE.md`
- `C:\Users\gregs\.codex\Orchestration\TASK-STATE.schema.json`
- `C:\Users\gregs\.codex\Orchestration\PASS-CHECKLIST.md`
- `C:\Users\gregs\.codex\Orchestration\PASS-CHECKLIST.schema.json`
- `C:\Users\gregs\.codex\Orchestration\Prompts\TASK-LEADER.md`
- `C:\Users\gregs\.codex\Orchestration\Prompts\IMPLEMENTATION-LEADER.md`
- `C:\Users\gregs\.codex\Orchestration\Prompts\AUDITOR.md`

## Proposed Changes

- Add these stop-governance fields to task state:
  - `next_action`
  - `blocking_reason`
  - `blocking_owner`
  - `stop_eligibility`
  - `last_stop_decision`
- Define one stop-eligibility rule in shared orchestration: if the task is still open, the next justified action is known, and there is no real external block, the agent must continue instead of stopping.
- Require pass-closeout state to record whether the stop was valid and, if valid, which blocker justified it.
- Update task and implementation leader prompts so milestone summaries cannot masquerade as legitimate stops.

## Expected Resolution

When work is still open, the system should keep moving until it hits a real block or the actual done bar.
If it must stop, the human should be able to see exactly why the stop was allowed without reconstructing the situation from chat.

## What Does Not Count

- A milestone summary that ends with "next steps" but records no blocker.
- A pause justified only by "good stopping point" or "waiting" with no named dependency.
- A stop that asks the human to wake the work back up when the next action was already known.

## Acceptance Criteria

- `TASK-STATE.md` and `TASK-STATE.schema.json` define and validate `next_action`, `blocking_reason`, `blocking_owner`, and `stop_eligibility`.
- `PASS-CHECKLIST.md` and `PASS-CHECKLIST.schema.json` require each stop or pause to record whether it was stop-eligible and why.
- `ORCHESTRATION.md`, `TASK-LEADER.md`, and `IMPLEMENTATION-LEADER.md` state that open work with a known next action must continue unless a real external blocker exists.
- On a pilot task with an open `ThirdPerson` default-lane defect and a known next action, the workflow cannot end as a clean stop without a blocker record.
- On a pilot task with a real external dependency, the workflow can stop and the recorded blocker is visible in both task state and the closeout artifact.

## Proof Plan

- Re-run one wake-up-prone task flow in a dry-run or pilot mode and verify the gate blocks one premature stop.
- Record one valid externally blocked stop and verify the blocker propagates into both state and checklist artifacts.
- Compare the stop record against the human-facing summary to confirm the same reason is surfaced in both places.

## References

- `..\BURDEN-ANALYSIS.md`
- `..\ORTHOGONAL-SOLUTIONS-MATRIX.md`
- `..\Plans\PROBLEM-0005-OPTION-A-PLAN-0001.md`
- `C:\Agent\ThirdPerson\AGENTS.md`

## Plan Addendum

Chosen plan: add a stop-eligibility gate so open work does not stop while the next justified action is known and no real external block exists.

Implementation notes from the selected plan:

- evaluate every proposed stop against open-task status, known next action, and external blockers
- block checkpoint-style stops that would only force the human to wake the work back up
- require an explicit blocker record when stopping before the human's done bar

Rollout:

1. Define the stop-eligibility test.
2. Add it before milestone summaries and pause decisions.
3. Make blocker recording mandatory when the gate says the task cannot honestly stop cleanly.
4. Pilot it on task flows with measured wake-up loss like the `ThirdPerson` packet.
5. Promote the rule after it proves it reduces resume-prod events.
