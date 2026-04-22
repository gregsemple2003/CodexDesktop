# Problem 0003 Option A Plan 0001

## Planning Intent

This file turns Problem `0003`, Option `A. Approval checklist` from [ORTHOGONAL-SOLUTIONS-MATRIX.md](../ORTHOGONAL-SOLUTIONS-MATRIX.md) into a bounded implementation sequence.

It is an alternative route, not the selected winner task.

## Summary

Strengthen the approval boundary by requiring one small checklist in every approval request so the human does not have to reconstruct what changed.

## Fixed Defaults

- scope: shared workflow and prompt rules only
- canonical homes:
  - `C:\Users\gregs\.codex\Orchestration\Processes\TASK-CREATE.md`
  - `C:\Users\gregs\.codex\Orchestration\ORCHESTRATION.md`
  - `C:\Users\gregs\.codex\Orchestration\Prompts\TASK-LEADER.md`
  - `C:\Users\gregs\.codex\Orchestration\Prompts\IMPLEMENTATION-LEADER.md`
- required checklist items:
  - changed artifact links
  - old vs new summary
  - one-line reason per change
  - current pass or gate
  - exact approval ask
- no generated packet file in this rollout

## Pass Plan

### Pass 0000 - Checklist Contract

Goal:

- define the minimum approval checklist once

Build:

- update `Processes/TASK-CREATE.md` with the checklist for plan or pass approval
- update `ORCHESTRATION.md` so approval requests are incomplete when the checklist is missing
- add one short good example and one bad example

Unit Proof:

- both docs list the same checklist items
- the bad example clearly shows missing review context

Exit Bar:

- approval expectations are durable and explicit even without a generated packet artifact

### Pass 0001 - Prompt Adoption

Goal:

- make the checklist show up when approval is actually requested

Build:

- update `TASK-LEADER.md` and `IMPLEMENTATION-LEADER.md` so approval requests must include the checklist
- require that approval messages link the changed artifacts directly

Unit Proof:

- both prompts require the same checklist
- the prompts do not allow `please approve` with no linked change context

Exit Bar:

- the human can review an approval ask without doing manual archaeology in chat

## Testing Strategy

- check that checklist item names match across workflow and prompts
- reject approval guidance that still leaves `what changed` or `where is it` implicit

## Deferred Work

- generated approval packets
- dedicated review UI
- packet-level schema validation
