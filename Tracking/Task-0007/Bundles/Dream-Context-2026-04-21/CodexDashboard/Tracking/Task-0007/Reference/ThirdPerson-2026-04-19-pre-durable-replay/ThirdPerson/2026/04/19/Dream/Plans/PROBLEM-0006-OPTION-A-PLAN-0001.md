# Problem 0006 Option A Plan 0001

## Planning Intent

This file turns Problem `0006`, Option `A. Manual memory upkeep` from [ORTHOGONAL-SOLUTIONS-MATRIX.md](../ORTHOGONAL-SOLUTIONS-MATRIX.md) into a bounded implementation sequence.

It is an alternative route, not the selected winner task.

## Summary

Create one lightweight manual lesson-capture workflow so repeated corrections are intentionally patched into docs or prompts instead of being left in chat by accident.

## Fixed Defaults

- scope: shared intervention and autoimprovement docs only
- canonical homes:
  - `C:\Users\gregs\.codex\Orchestration\Processes\AUTOIMPROVEMENT.md`
  - `C:\Users\gregs\.codex\Orchestration\Processes\INTERVENTION-REPORTS.md`
  - `C:\Users\gregs\.codex\Orchestration\Prompts\TASK-HARVESTER.md`
- manual trigger rule:
  - after the same correction appears twice, the reviewer must choose a durable home or explicitly record why not
- no ledger, schema, or hidden profile in this rollout

## Pass Plan

### Pass 0000 - Manual Lesson-Capture Rule

Goal:

- define when a repeated correction must stop living only in chat

Build:

- update `AUTOIMPROVEMENT.md` with the two-repeat trigger rule
- update `Processes/INTERVENTION-REPORTS.md` with the manual promotion decision points
- add one short note that `leave it in chat` is not a default

Unit Proof:

- both docs use the same trigger rule
- the manual decision points are concrete enough to follow

Exit Bar:

- repeated corrections have a named manual capture path instead of pure good intentions

### Pass 0001 - Harvester Adoption And Examples

Goal:

- make manual capture part of the intervention review workflow

Build:

- update `TASK-HARVESTER.md` to ask whether a repeated correction needs durable promotion
- add one example where a correction becomes a prompt rule and one where it becomes a repo doc change

Unit Proof:

- the harvester prompt points at the same trigger rule as the shared docs
- the examples show two distinct durable homes

Exit Bar:

- future reviewers are more likely to promote stable lessons without needing new machinery

## Testing Strategy

- compare the trigger rule across the touched docs
- reject wording that still leaves the promotion decision as optional by default

## Deferred Work

- append-only lesson ledger
- automatic promotion thresholds
- hidden preference storage
