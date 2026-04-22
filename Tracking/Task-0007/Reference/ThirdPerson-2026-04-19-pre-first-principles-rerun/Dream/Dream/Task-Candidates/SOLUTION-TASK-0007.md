# Solution Task 0007

## Title

Define a checkpoint and continuation protocol for active passes

## Summary

The packet shows active work stopping too early, durable state lagging behind the real work, and the human having to resume or coordinate tasks that should have continued.
This task defines when work keeps going, when state must refresh, and when messaging is actually allowed.

## Goals

- define durable milestone types for active passes
- require state refresh at those milestones
- narrow the allowed reasons for sending an interrupting progress message

## Non-Goals

- building a background heartbeat service
- preventing user-requested progress reports
- replacing normal handoff and closeout artifacts

## Implementation Home

- `C:\Users\gregs\.codex\Orchestration\ORCHESTRATION.md`
- `C:\Users\gregs\.codex\Orchestration\TASK-STATE.md`
- `C:\Users\gregs\.codex\Orchestration\PASS-CHECKLIST.md`
- shared active-work prompt files under `C:\Users\gregs\.codex\Orchestration\Prompts\`

## Constraints And Baseline

- active work should continue through the next durable milestone unless there is an approval gate, a real external block, or an explicit user request for status
- state refresh should happen before optional commentary messages
- the protocol must preserve existing task-owned history rather than hiding failed attempts

## Proposed Changes

- define milestone types such as `pass start`, `meaningful checkpoint`, `blocking failure`, `approval gate`, and `pass closeout`
- require task or pass state to record the last reached milestone and the next expected milestone
- define allowed message reasons during active work and make all other cases continue silently
- add a resume rule that says non-terminal failed attempts move to the next justified step instead of pausing by default

## Expected Resolution

Users should spend less time acting as the workflow scheduler.
Durable task state should stay closer to the real state of the work, and needless resume prompts should fall.

## What Does Not Count

- a generic `keep going` reminder with no milestone rules
- progress messages that still appear for normal non-terminal failures
- state fields that exist but are never required during active passes

## Acceptance Criteria

- shared workflow docs define the active-pass milestone types and allowed message reasons
- task-state or pass-checklist guidance records the last and next milestone for active work
- prompts require state refresh before optional progress commentary
- the protocol explicitly says that a non-terminal failed attempt continues to the next justified step unless an allowed stop reason exists

## Proof Plan

- test the protocol against the April 19 charged stall and resume cases
- confirm the workflow now requires a durable checkpoint before a progress message
- confirm ordinary failed attempts no longer default to `stop and ask what next`

## References

- `../ORTHOGONAL-SOLUTIONS-MATRIX.md`
- `../Plans/PROBLEM-0007-OPTION-B-PLAN-0001.md`

## Plan Addendum

Start with a short list of milestone types and allowed message reasons.
Tie those milestones to existing task-state and pass-closeout artifacts so the protocol does not become a separate shadow system.
Then add the rule that non-terminal failures continue to the next justified step unless a real gate or block exists.
