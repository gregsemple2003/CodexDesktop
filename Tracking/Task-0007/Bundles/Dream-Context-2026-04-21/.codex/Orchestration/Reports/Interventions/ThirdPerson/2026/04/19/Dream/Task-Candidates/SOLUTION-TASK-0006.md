# Solution Task 0006

## Title

Emit one-screen approval packets with exact file links and state delta

## Summary

The packet shows that approval turns were blocked because the human could not see a usable diff shape, pass framing, or state delta in one review surface.
This task adds one standard approval packet so approval asks arrive with a concise change summary, exact links, pass id, and durable state change.

## Goals

- Make approval asks reviewable without tree hunting.
- Show content changes and state changes together.
- Keep pass framing visible when multiple pass artifacts changed.

## Non-Goals

- Replacing full diffs or git tools.
- Creating a separate running changelog for every task session.
- Treating raw links as a sufficient approval surface.

## Constraints And Baseline

Current truth:

- The packet records approval blockage because there was no usable diff for `PLAN.md`.
- The human later asked where the new passes were and said links without context were not enough.
- Durable task state also lagged behind live work.

Hard constraints:

- the approval surface must fit on one screen in common cases
- exact links and pass framing must be present together
- state delta must sit beside content change summary, not in a separate hunt

## Implementation Home

- `C:\Users\gregs\.codex\Orchestration\TASK-STATE.md`
- `C:\Users\gregs\.codex\Orchestration\TASK-STATE.schema.json`
- `C:\Users\gregs\.codex\Orchestration\PASS-CHECKLIST.md`
- `C:\Users\gregs\.codex\Orchestration\PASS-CHECKLIST.schema.json`
- `C:\Users\gregs\.codex\Orchestration\Prompts\TASK-LEADER.md`
- `C:\Users\gregs\.codex\Orchestration\Prompts\IMPLEMENTATION-LEADER.md`
- task approval surface in `C:\Agent\ThirdPerson\Tracking/Task-<id>\HANDOFF.md`

## Proposed Changes

- Define one standard approval packet shape with these required fields:
  - `pass_id`
  - `approval_question`
  - `changed_files`
  - `change_summary`
  - `state_delta`
  - `open_risks`
- Require `IMPLEMENTATION-LEADER.md` and `TASK-LEADER.md` to emit that packet whenever they ask for approval.
- Add matching fields to pass-closeout state so the durable record and the human-facing packet stay aligned.
- Require `HANDOFF.md` to mirror the current approval packet under a dedicated section when approval is pending.

## Expected Resolution

A human should be able to decide "approve," "reject," or "ask for one correction" from one bounded packet that says what changed, where it changed, what pass it belongs to, and how task state moved.

## What Does Not Count

- A bundle of links with no summary.
- A prose summary with no exact file links.
- A state update that is not shown beside the content change.
- A packet so large that the human still has to reconstruct the pass by hand.

## Acceptance Criteria

- `PASS-CHECKLIST.md` and `PASS-CHECKLIST.schema.json` define and validate the required approval-packet fields.
- `TASK-STATE.md` and `TASK-STATE.schema.json` record the latest pending approval packet and its state delta.
- `TASK-LEADER.md` and `IMPLEMENTATION-LEADER.md` require approval asks to emit the standard packet instead of loose link dumps.
- `HANDOFF.md` on a `ThirdPerson` pilot task can mirror a pending approval ask with pass id, change summary, exact links, and state delta in one screen.
- A review ask that contains only links or only prose summary fails the gate.

## Proof Plan

- Pilot the packet on one `ThirdPerson` pass that changes more than one task artifact.
- Verify the same packet fields appear in the human-facing approval ask, `HANDOFF.md`, and pass-closeout state.
- Ask whether the packet can be reviewed from one screen without opening the tree first; revise only if the answer is no.

## References

- `..\BURDEN-ANALYSIS.md`
- `..\ORTHOGONAL-SOLUTIONS-MATRIX.md`
- `..\Plans\PROBLEM-0004-OPTION-A-PLAN-0001.md`
- `C:\Agent\ThirdPerson\AGENTS.md`

## Plan Addendum

Chosen plan: auto-build a review packet for approval turns with a concise diff summary, exact file links, pass framing, and state delta.

Implementation notes from the selected plan:

- every approval ask emits a short review packet instead of a loose link dump
- the packet explains what changed, where it changed, which pass it belongs to, and what needs approval
- durable state changes are shown beside content changes

Rollout:

1. Define the minimum packet fields: pass id, changed files, short change summary, state change, and approval ask.
2. Generate the packet automatically whenever the workflow reaches an approval gate.
3. Pilot the packet on a `ThirdPerson` planning pass that changes multiple task artifacts.
4. Refine the packet until one screen is enough to understand the approval ask.
