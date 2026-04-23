# Dream Annex README (ThirdPerson 2026-04-19)

This Dream annex now includes winner-task drafting and task audit.

The current annex therefore covers:

- pass `1`
- pass `2A`
- pass `2B`
- pass `2C`

## Main Reading Path

1. [SOLUTION-DESIGN.md](./SOLUTION-DESIGN.md)
2. [WINNER-SYNTHESIS.md](./WINNER-SYNTHESIS.md)
3. [Option-Tasks/INDEX.md](./Option-Tasks/INDEX.md)
4. [BURDEN-ANALYSIS.md](./BURDEN-ANALYSIS.md)
5. [ORTHOGONAL-SOLUTIONS-MATRIX.md](./ORTHOGONAL-SOLUTIONS-MATRIX.md)
6. [APPROACH.md](./APPROACH.md)

## File Guide

- [APPROACH.md](./APPROACH.md)
  - pass-`1` scope, method, and evidence base
- [SessionExcerpts/INDEX.json](./SessionExcerpts/INDEX.json)
  - packet-local transcript neighborhoods with explicit `source_event_ids`
- [BURDEN-ANALYSIS.md](./BURDEN-ANALYSIS.md)
  - evidence-backed burden drivers from pass `1`
- [SOLUTION-DESIGN.md](./SOLUTION-DESIGN.md)
  - corrected pass-`2A` problem rows, option set, exact homes, enforcement boundaries, acceptance tests, and falsifiers
- [WINNER-SYNTHESIS.md](./WINNER-SYNTHESIS.md)
  - frozen winner set plus the pass-`2C` task-audit corrections that the audited tasks forced
- [Option-Tasks/INDEX.md](./Option-Tasks/INDEX.md)
  - one audited task-shaped proposal per frozen winner
- [ORTHOGONAL-SOLUTIONS-MATRIX.md](./ORTHOGONAL-SOLUTIONS-MATRIX.md)
  - compact final decision surface linking burden drivers, designed options, and winners

## Current Scope Boundary

- This annex includes winner tasks.
- The winner-task set was re-audited against `TASK-CREATE.md` and `TASK-AUDIT.md`.
- `P-003` remains an explicitly `consensus` task (the decision carve-outs are not frozen enough for honest enqueue-ready implementation without a bounded decision artifact).

## Corrections Forced In `2C`

- `P-001` now requires "bad/good" `REGRESSION-RUN` fixtures to live under the implementing task's own `Tracking/Task-<id>/Testing/` home (no implementer-chosen scratch location).
- `P-002` now names `execution.run_state` (not `execution.status`) and defines explicit consistency rules with existing task-level `status/phase/current_gate`, to avoid "two truths" about whether work may proceed.
- `P-004` now points its template/proof work explicitly at the shared exemplars (`C:\Users\gregs\.codex\Orchestration\Exemplars\PLAN.md` and `C:\Users\gregs\.codex\Orchestration\Exemplars\HANDOFF.md`) and removes the "pick any template home" escape hatch.

## Home Split

- repo-local
  - ThirdPerson lane truth and ThirdPerson runtime-debugging specifics
- shared
  - lifecycle rules, proof-quality rules, and any cross-repo decision that survives audit
- task-local
  - approval packets, current-state constraints, bug notes, and executed regression artifacts

## Notes

- the task-audit lane used only durable packet inputs, packet-cited source material, repo-local docs, shared process docs, and the current on-disk Dream artifacts
- no reference packet or hidden writer context was used
