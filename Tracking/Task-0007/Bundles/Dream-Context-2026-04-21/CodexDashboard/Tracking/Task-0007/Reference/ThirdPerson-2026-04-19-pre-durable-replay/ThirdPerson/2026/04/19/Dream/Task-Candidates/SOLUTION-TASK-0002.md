# Solution Task 0002

## Title

Require a task-owned approval packet before any plan or change approval request.

## Summary

The packet shows repeated approval failures where the human had to ask for diffs, links, and file locations before they could review anything.

This is a concrete implementation task. It will define one approval-packet format, one task-owned file naming rule, and one prompt rule that says an approval request is incomplete until that packet exists.

## Goals

- make approval a one-stop review surface
- show exact changed text, exact links, and the reason for each change
- make plan approval and later change approval follow the same packet shape
- remove the need for manual document archaeology

## Non-Goals

- building a full review UI
- changing the substance of the plan or pass being reviewed
- solving task continuity or wrong-lane proof inside this task

## Constraints And Baseline

- the approval packet must work in plain Markdown
- the packet must be readable from disk without assuming a fresh editor buffer
- the packet must be compact enough to review quickly

## Proposed Changes

- add `C:\Users\gregs\.codex\Orchestration\APPROVAL-PACKET-WORKFLOW.md`
- update [FILE-NAMING.md](../../../../../../../../FILE-NAMING.md) with these task-owned packet names: `Tracking/Task-<id>/APPROVAL-PACKET-PLAN.md` and `Tracking/Task-<id>/APPROVAL-PACKET-PASS-<NNNN>.md`
- define these required packet sections in `APPROVAL-PACKET-WORKFLOW.md`: `Decision Needed`, `Changed Files`, `Old Vs New`, `Why Each Change Exists`, `Current Pass Or Work Location`, and `Risks Or Open Questions`
- update [TASK-LEADER.md](../../../../../../../../Prompts/TASK-LEADER.md) and `IMPLEMENTATION-LEADER.md` so they may not ask for approval unless the matching packet file exists and is linked in the request
- require `Old Vs New` to show exact old wording and exact new wording for prose artifacts instead of only linking a file

## Expected Resolution

- every approval request points to one packet file
- the packet shows exactly what changed and why
- the human can approve or reject without asking where the change lives

## What Does Not Count

- a raw file link with no packet
- a packet that lists files but not the changed text
- telling the human to compare local files by hand
- a packet that omits the pass id or work location

## Implementation Home

Keep task-owned planning, handoff, and testing artifacts under the eventual `Tracking/Task-<id>/`.

Implement the shared approval contract under:

- `C:\Users\gregs\.codex\Orchestration\APPROVAL-PACKET-WORKFLOW.md`
- [FILE-NAMING.md](../../../../../../../../FILE-NAMING.md)
- [TASK-LEADER.md](../../../../../../../../Prompts/TASK-LEADER.md)
- `C:\Users\gregs\.codex\Orchestration\Prompts\IMPLEMENTATION-LEADER.md`

## Proof Plan

- check that the workflow doc names the exact packet files and sections
- check that the leader prompts treat approval requests without a packet as invalid
- check that the packet shape includes both links and old-vs-new wording

## Acceptance Criteria

- `APPROVAL-PACKET-WORKFLOW.md` exists and defines both `APPROVAL-PACKET-PLAN.md` and `APPROVAL-PACKET-PASS-<NNNN>.md`
- the workflow doc requires the sections `Decision Needed`, `Changed Files`, `Old Vs New`, `Why Each Change Exists`, `Current Pass Or Work Location`, and `Risks Or Open Questions`
- [FILE-NAMING.md](../../../../../../../../FILE-NAMING.md) lists the new approval-packet artifact names
- [TASK-LEADER.md](../../../../../../../../Prompts/TASK-LEADER.md) and `IMPLEMENTATION-LEADER.md` say an approval request is incomplete without the matching packet file
- the approval packet rules require exact old-versus-new wording for prose changes, not just links
- a reviewer could tell from the written contract how this change would remove `where is the diff?` and `where does this live?` approval turns

## References

- [TASK-CREATE.md](../../../../../../../../Processes/TASK-CREATE.md)
- [BURDEN-ANALYSIS.md](../BURDEN-ANALYSIS.md)
- [ORTHOGONAL-SOLUTIONS-MATRIX.md](../ORTHOGONAL-SOLUTIONS-MATRIX.md)
