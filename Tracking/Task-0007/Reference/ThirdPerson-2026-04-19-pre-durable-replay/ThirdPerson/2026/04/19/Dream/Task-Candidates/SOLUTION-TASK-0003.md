# Solution Task 0003

## Title

Add an append-only intervention lesson ledger and promotion rule.

## Summary

The packet shows the human repeating stable corrections about lanes, proof, answer shape, and task ownership.

This is a concrete implementation task. It will add one structured ledger for durable lessons and one promotion rule that turns repeated explicit corrections into the smallest correct doc, prompt, or gate change.

## Goals

- capture repeated corrections in one inspectable place
- keep explicit, strongly implied, and speculative lessons separate
- make promotion thresholds explicit
- tie promoted rules back to source packet evidence

## Non-Goals

- hidden model-side memory
- promoting every complaint into a global rule
- replacing source packets with summaries
- solving approval, proof, or debugging gates inside this task

## Constraints And Baseline

- the ledger must stay in the shared intervention canon, not in private runtime state
- promotion must preserve scope so task-local lessons do not become shared rules by accident
- promoted lessons must stay reversible

## Proposed Changes

- add `C:\Users\gregs\.codex\Orchestration\Processes\INTERVENTION-PROMOTION.md`
- add `C:\Users\gregs\.codex\Orchestration\Intervention-Lessons\LESSONS-LEDGER.jsonl`
- add `C:\Users\gregs\.codex\Orchestration\Intervention-Lessons\LESSONS-LEDGER-ENTRY.schema.json`
- update `C:\Users\gregs\.codex\Orchestration\Processes\AUTOIMPROVEMENT.md` so the daily sweep must append new lessons and check promotion thresholds
- update `C:\Users\gregs\.codex\Orchestration\Processes\INTERVENTION-REPORTS.md` so the canon explains where lessons live and how promotion works
- define these exact ledger fields: `lesson_id`, `source_packet`, `source_event_ids`, `lesson_text`, `lesson_kind`, `scope`, `promotion_target`, `promotion_state`, `repeat_count`, and `supersedes`
- define these exact enum values:
  - `lesson_kind`: `explicit`, `strongly_implied`, `speculative`
  - `scope`: `shared`, `repo_local`, `task_local`
  - `promotion_state`: `logged`, `ready_for_promotion`, `promoted`, `rejected`, `superseded`
- define this promotion rule:
  - promote when `lesson_kind=explicit` and `repeat_count>=2`
  - also allow promotion after one entry when the lesson is marked as a red-line rule in the packet or shared docs

## Expected Resolution

- repeated corrections stop living only in chat
- strong lessons become durable rules in the right home
- weak or speculative lessons stay logged without being over-promoted

## What Does Not Count

- a hidden preference profile
- a summary doc with no structured ledger
- a ledger with no promotion rule
- promotion with no source packet or lesson id reference

## Implementation Home

Implement the shared learning and promotion path under:

- `C:\Users\gregs\.codex\Orchestration\Processes\INTERVENTION-PROMOTION.md`
- `C:\Users\gregs\.codex\Orchestration\Intervention-Lessons\LESSONS-LEDGER.jsonl`
- `C:\Users\gregs\.codex\Orchestration\Intervention-Lessons\LESSONS-LEDGER-ENTRY.schema.json`
- `C:\Users\gregs\.codex\Orchestration\Processes\AUTOIMPROVEMENT.md`
- `C:\Users\gregs\.codex\Orchestration\Processes\INTERVENTION-REPORTS.md`

## Proof Plan

- check that the ledger file, schema, and workflow doc all exist
- check that the workflow doc defines the exact fields and enums listed above
- check that `AUTOIMPROVEMENT.md` names the append-and-promote loop explicitly

## Acceptance Criteria

- `LESSONS-LEDGER.jsonl` exists and its entry schema defines every field listed in `Proposed Changes`
- `INTERVENTION-PROMOTION.md` defines the enum values `explicit`, `strongly_implied`, `speculative`, `shared`, `repo_local`, `task_local`, `logged`, `ready_for_promotion`, `promoted`, `rejected`, and `superseded`
- the workflow doc defines the promotion threshold `lesson_kind=explicit` and `repeat_count>=2`, plus the one-entry red-line exception
- `AUTOIMPROVEMENT.md` says the daily sweep must append lessons and check promotion state
- the workflow doc says every promoted rule must link back to a `lesson_id` and source packet
- the written design makes clear that hidden memory alone does not satisfy the task

## References

- [TASK-CREATE.md](../../../../../../../../Processes/TASK-CREATE.md)
- [TOPIC-06-LOCAL-MEMORY-AND-LEARNING.md](../../../../../../../../../../../../Agent/CodexDashboard/Tracking/Task-0007/Research/Analysis/TOPIC-06-LOCAL-MEMORY-AND-LEARNING.md)
- [ORTHOGONAL-SOLUTIONS-MATRIX.md](../ORTHOGONAL-SOLUTIONS-MATRIX.md)
