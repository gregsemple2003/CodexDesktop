# Solution Task 0003 Plan 0001

## Planning Intent

This file turns [SOLUTION-TASK-0003.md](../Task-Candidates/SOLUTION-TASK-0003.md) into a bounded implementation sequence.

It describes the intended route to done, not lesson history.

## Summary

Add one durable lesson ledger and one promotion workflow so repeated explicit corrections become shared, repo-local, or task-local rules instead of staying trapped in chat.

## Fixed Defaults

- durable ledger home:
  - `Intervention-Lessons\LESSONS-LEDGER.jsonl`
- schema home:
  - `Intervention-Lessons\LESSONS-LEDGER-ENTRY.schema.json`
- workflow home:
  - `Processes\INTERVENTION-PROMOTION.md`
- promotion threshold:
  - `lesson_kind=explicit`
  - `repeat_count>=2`
- red-line exception:
  - one entry may be enough when a shared rule is explicitly marked red-line

## Pass Plan

### Pass 0000 - Ledger Contract

Goal:

- define the ledger entry shape and promotion state machine

Build:

- add the JSONL ledger file
- add the entry schema
- define the exact fields, enums, and promotion-state meanings in the workflow doc

Unit Proof:

- schema and workflow doc name the same fields and enums
- promotion states have one clear meaning each

Exit Bar:

- a reviewer can tell how a lesson is logged, promoted, rejected, or superseded

### Pass 0001 - Process Integration

Goal:

- make the ledger part of the shared self-improvement loop

Build:

- update `AUTOIMPROVEMENT.md` so the daily sweep appends lessons and checks promotion thresholds
- update `Processes\INTERVENTION-REPORTS.md` so the intervention canon explains the ledger and promotion loop

Unit Proof:

- both docs point at the same ledger and workflow files
- both docs preserve the separation between explicit, strongly implied, and speculative lessons

Exit Bar:

- the promotion loop is part of the durable workflow rather than an optional side habit

### Pass 0002 - Promotion Examples

Goal:

- make the promotion rule concrete enough to apply consistently

Build:

- add one valid example for shared promotion
- add one valid example for repo-local promotion
- add one example that stays logged but does not get promoted

Unit Proof:

- each example maps to one allowed scope
- the non-promotion example shows why weak evidence does not count

Exit Bar:

- a reviewer can tell when a lesson belongs in the ledger only and when it must become a rule

## Testing Strategy

- validate the schema and workflow against the same field set
- verify `AUTOIMPROVEMENT.md` and `Processes\INTERVENTION-REPORTS.md` reference the durable ledger path
- reject any rollout that hides learning in undocumented runtime state

## Deferred Work

Keep these out of this rollout unless the task expands intentionally:

- automatic promotion bots
- model-side preference memory
- cross-packet ranking or analytics dashboards
