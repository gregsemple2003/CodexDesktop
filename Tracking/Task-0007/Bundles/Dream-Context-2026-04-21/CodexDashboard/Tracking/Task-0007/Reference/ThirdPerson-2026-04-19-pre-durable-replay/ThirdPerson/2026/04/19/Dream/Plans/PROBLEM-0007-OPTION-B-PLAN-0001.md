# Problem 0007 Option B Plan 0001

## Planning Intent

This file turns Problem `0007`, Option `B. Required disagreement template` from [ORTHOGONAL-SOLUTIONS-MATRIX.md](../ORTHOGONAL-SOLUTIONS-MATRIX.md) into a bounded implementation sequence.

It is an alternative route, not the selected winner task.

## Summary

Add one required disagreement trace artifact so debug work has to name the exact seam between expected state and observed state before root-cause claims can proceed.

## Fixed Defaults

- scope: shared debugging contract plus task-owned trace artifacts
- canonical homes:
  - `C:\Users\gregs\.codex\Orchestration\DISAGREEMENT-TRACE.md`
  - `C:\Users\gregs\.codex\Orchestration\FILE-NAMING.md`
  - `C:\Users\gregs\.codex\Orchestration\Processes\DEBUGGING.md`
  - `C:\Users\gregs\.codex\Orchestration\Prompts\DEBUG-LEADER.md`
  - `C:\Users\gregs\.codex\Orchestration\Prompts\DEBUG-WORKER.md`
- task-owned artifact name:
  - `Tracking/Task-<id>/DISAGREEMENT-TRACE-<NNNN>.md`
- required sections:
  - expected state
  - observed state
  - disagreement seam
  - next upstream writer

## Pass Plan

### Pass 0000 - Trace Contract And Naming

Goal:

- define the disagreement trace format once

Build:

- add `DISAGREEMENT-TRACE.md`
- update `FILE-NAMING.md` with `DISAGREEMENT-TRACE-<NNNN>.md`
- update `Processes/DEBUGGING.md` to require the trace for active root-cause work

Unit Proof:

- the doc and naming rule use the same artifact name
- the required sections are concrete and finite

Exit Bar:

- debug work now has one canonical artifact for the first concrete disagreement

### Pass 0001 - Prompt Adoption

Goal:

- make trace creation part of live debug workflow

Build:

- update `DEBUG-LEADER.md` and `DEBUG-WORKER.md` to require a disagreement trace before root-cause language
- require the next upstream writer to be named, not implied

Unit Proof:

- both prompts point at the same trace artifact
- neither prompt allows a root-cause claim with no disagreement seam

Exit Bar:

- the debug path starts from a concrete disagreement instead of drifting through symptoms

### Pass 0002 - Example Trace

Goal:

- make the artifact shape easy to copy

Build:

- add one short example disagreement trace from expected runtime state to upstream writer
- add one anti-pattern where the trace talks about symptoms only

Unit Proof:

- the example trace uses all required sections
- the anti-pattern shows exactly what a non-narrowing trace looks like

Exit Bar:

- future tasks can create a disagreement trace without inventing their own structure

## Testing Strategy

- compare the trace section names across the shared contract and the prompt rules
- reject debug guidance that still allows root-cause language with no trace artifact

## Deferred Work

- root-cause claim verification gate
- automated trace linting
- dashboard trace browser
