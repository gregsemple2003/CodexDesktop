# Solution Task 0002 Plan 0001

## Planning Intent

This file turns [SOLUTION-TASK-0002.md](../Task-Candidates/SOLUTION-TASK-0002.md) into a bounded implementation sequence.

It describes the intended route to done, not review history.

## Summary

Create one shared approval-packet workflow so plan and change approvals always come with a compact, complete review packet.

## Fixed Defaults

- packet format: plain Markdown
- task-owned packet names:
  - `APPROVAL-PACKET-PLAN.md`
  - `APPROVAL-PACKET-PASS-<NNNN>.md`
- required packet sections:
  - `Decision Needed`
  - `Changed Files`
  - `Old Vs New`
  - `Why Each Change Exists`
  - `Current Pass Or Work Location`
  - `Risks Or Open Questions`
- no custom UI in this rollout

## Pass Plan

### Pass 0000 - Packet Contract And Naming

Goal:

- define the approval-packet artifact and its fixed section layout

Build:

- add `APPROVAL-PACKET-WORKFLOW.md`
- update `FILE-NAMING.md` with the packet filenames
- define the fixed section set and the rule that prose changes must show exact old-versus-new wording

Unit Proof:

- the workflow doc and naming doc agree on packet filenames
- the workflow doc includes all six required sections

Exit Bar:

- a reviewer can read the packet contract and know exactly what an approval packet must contain

### Pass 0001 - Leader Gating

Goal:

- make approval requests invalid unless the packet exists

Build:

- update `TASK-LEADER.md` to require a matching approval packet before asking for approval
- update `IMPLEMENTATION-LEADER.md` to do the same for plan approval and change approval

Unit Proof:

- both prompts say approval requests are incomplete without the packet
- both prompts point at the task-owned packet artifact rather than free-form summary prose

Exit Bar:

- the shared leaders cannot honestly ask for approval without the packet

### Pass 0002 - Examples And Review Guidance

Goal:

- make the approval surface easy to apply consistently

Build:

- add one compact packet example to `APPROVAL-PACKET-WORKFLOW.md`
- add one invalid example that shows why raw links do not count
- add a note on keeping packets compact rather than turning them into a second plan document

Unit Proof:

- the valid example uses every required section
- the invalid example demonstrates at least one rejected pattern from the task

Exit Bar:

- a later agent has a concrete model for how to build and present the packet

## Testing Strategy

- compare packet examples against the fixed section list
- verify both leader prompts block approval without the packet
- reject any rollout that still allows raw links with no old-versus-new wording for prose

## Deferred Work

Keep these out of this rollout unless the task expands intentionally:

- PR-style review UI
- automatic diff extraction tooling
- packet rendering beyond Markdown
