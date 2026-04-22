# Problem 0006 Option B Plan 0001

## Planning Intent

This file turns Problem `0006`, Option `B. Intervention ledger plus promotion loop` from [ORTHOGONAL-SOLUTIONS-MATRIX.md](../ORTHOGONAL-SOLUTIONS-MATRIX.md) into a bounded implementation sequence.

It is the matrix option that fed [SOLUTION-TASK-0003.md](../Task-Candidates/SOLUTION-TASK-0003.md).

## Summary

Add one append-only lesson ledger and one promotion workflow so repeated human corrections become inspectable, scoped, and durable instead of recurring as chat-only tax.

## Fixed Defaults

- scope: intervention canon plus autoimprovement workflow
- canonical homes:
  - `C:\Users\gregs\.codex\Orchestration\Processes\INTERVENTION-PROMOTION.md`
  - `C:\Users\gregs\.codex\Orchestration\Intervention-Lessons\LESSONS-LEDGER.jsonl`
  - `C:\Users\gregs\.codex\Orchestration\Intervention-Lessons\LESSONS-LEDGER-ENTRY.schema.json`
  - `C:\Users\gregs\.codex\Orchestration\Processes\AUTOIMPROVEMENT.md`
  - `C:\Users\gregs\.codex\Orchestration\Processes\INTERVENTION-REPORTS.md`
- required ledger fields:
  - correction summary
  - source packet refs
  - scope
  - confidence
  - suggested durable home
  - promotion status

## Pass Plan

### Pass 0000 - Ledger Contract

Goal:

- create one inspectable home for repeated intervention lessons

Build:

- add `INTERVENTION-PROMOTION.md`
- add `LESSONS-LEDGER.jsonl`
- add `LESSONS-LEDGER-ENTRY.schema.json`

Unit Proof:

- the workflow doc and schema define the same ledger fields
- the ledger is append-only by contract

Exit Bar:

- repeated corrections can now be captured without inventing a new local format each time

### Pass 0001 - Promotion Loop Adoption

Goal:

- make the ledger part of the daily self-improvement loop

Build:

- update `AUTOIMPROVEMENT.md` so daily review must append new lessons and check promotion thresholds
- update `Processes/INTERVENTION-REPORTS.md` so the canon explains where lessons live and how promotion works
- define the smallest-correct-home rule for promotion

Unit Proof:

- the two docs agree on when a ledger item should be promoted
- promotion status values are concrete and finite

Exit Bar:

- stable lessons stop dying at the end of the chat window

### Pass 0002 - Examples And Promotion Cases

Goal:

- make the workflow easy to apply and audit

Build:

- add one example ledger entry that promotes into a shared prompt
- add one example ledger entry that promotes into a repo-root doc
- add one example that stays unpromoted because confidence is too low

Unit Proof:

- the example entries validate against the schema
- the three examples show distinct promotion outcomes

Exit Bar:

- a reviewer can see how a correction turns into a durable rule without guessing

## Testing Strategy

- validate that example ledger entries match the schema
- reject promotion rules that cannot explain why a lesson did or did not move to a durable home

## Deferred Work

- hidden model-side learning
- automatic prompt rewriting
- cross-repo deduplication beyond the local ledger
