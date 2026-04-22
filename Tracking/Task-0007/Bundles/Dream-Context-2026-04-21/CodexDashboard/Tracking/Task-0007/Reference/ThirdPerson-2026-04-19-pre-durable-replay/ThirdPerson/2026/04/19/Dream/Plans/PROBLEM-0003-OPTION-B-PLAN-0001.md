# Problem 0003 Option B Plan 0001

## Planning Intent

This file turns Problem `0003`, Option `B. Approval packet generator` from [ORTHOGONAL-SOLUTIONS-MATRIX.md](../ORTHOGONAL-SOLUTIONS-MATRIX.md) into a bounded implementation sequence.

It is an alternative route, not the selected winner task.

## Summary

Add one durable approval packet workflow so plan and pass approval requests always arrive with the changed text, links, and reason already packaged.

## Fixed Defaults

- scope: shared workflow plus task-owned packet artifacts
- canonical homes:
  - `C:\Users\gregs\.codex\Orchestration\APPROVAL-PACKET-WORKFLOW.md`
  - `C:\Users\gregs\.codex\Orchestration\FILE-NAMING.md`
  - `C:\Users\gregs\.codex\Orchestration\Processes\TASK-CREATE.md`
  - `C:\Users\gregs\.codex\Orchestration\Prompts\TASK-LEADER.md`
  - `C:\Users\gregs\.codex\Orchestration\Prompts\IMPLEMENTATION-LEADER.md`
- task-owned packet names:
  - `Tracking/Task-<id>/APPROVAL-PACKET-PLAN.md`
  - `Tracking/Task-<id>/APPROVAL-PACKET-PASS-<NNNN>.md`
- each packet must include changed sections, old vs new wording, file links, one-line reasons, and current ownership or gate state

## Pass Plan

### Pass 0000 - Packet Contract And Naming

Goal:

- define one stable approval packet shape and one stable file name scheme

Build:

- add `APPROVAL-PACKET-WORKFLOW.md`
- update `FILE-NAMING.md` with the two packet names
- define required packet sections for plan approval and pass approval

Unit Proof:

- the workflow doc and naming doc use the same packet names
- both packet variants have explicit required sections

Exit Bar:

- approval packets have one canonical home and one canonical shape

### Pass 0001 - Leader Adoption

Goal:

- make packet creation mandatory before approval asks

Build:

- update `Processes/TASK-CREATE.md` to require a plan approval packet for enqueue-ready tasks
- update `TASK-LEADER.md` and `IMPLEMENTATION-LEADER.md` so they may not ask for approval unless the matching packet exists and is linked

Unit Proof:

- all touched docs treat missing packets as an incomplete approval request
- the prompts require direct links to the packet file

Exit Bar:

- approval requests become lightweight review events instead of archaeology sessions

### Pass 0002 - Examples And Packet-Backed Checks

Goal:

- make the packet shape obvious to future authors

Build:

- add one short example `APPROVAL-PACKET-PLAN.md`
- add one short example `APPROVAL-PACKET-PASS-0001.md`
- note one anti-pattern where links exist but the old vs new wording is missing

Unit Proof:

- the example packets use the required sections
- the anti-pattern shows exactly why the packet would still be unusable

Exit Bar:

- a later task can copy the packet shape without guesswork

## Testing Strategy

- check that packet names and sections match across naming, workflow, and prompts
- treat approval requests without packet links as rollout failures

## Deferred Work

- a richer desktop review surface
- schema validation for packet contents
- threaded approval comments
