## Title
Explicit STOP/PAUSE protocol as durable execution state

## Summary
Make STOP/PAUSE and approval gates true execution state: while paused, tools are disallowed until an explicit resume.

## Goals
- Reduce repeated human interventions caused by lane/proof mismatch, boundary drift, and answer-shape mismatch.

## Non-Goals
- This task does not implement a gameplay fix inside the ThirdPerson codebase; it improves the workflow surface and proof semantics.

## Implementation Home
- Shared orchestration prompt bundles and/or ThirdPerson repo-local docs (REGRESSION.md, AGENTS.md) where applicable.

## Constraints And Baseline
- Treat ThirdPerson/REGRESSION.md as canonical for what counts as regression for this repo.
- Missing transcript for session_id 019da198-ea2f-7421-8adc-c566da0b6121 is a known fidelity gap; do not rely on assistant-side context from that session.

## Proposed Changes
- Add a STOP/PAUSE state machine to the orchestration prompts and worker behavior.
- Require a 'Paused' task_complete message when entering wait state, and require an explicit RESUME to continue.

## Burden Reduction Contract
- Burden being reduced: repeated human corrections about lane/proof/boundaries/answers.
- Causal claim: without durable gates and templates, the workflow drifts and the human must intervene to restate constraints.
- Evidence:
  - EXCERPT-0006 (pause gate + resume neighborhood)
  - hie-...-3447 (STOP only continue after fixed)
- Why this mechanism: it changes the default workflow so the agent cannot easily 'slide' into off-lane proof or continue past stops.
- Human relief if successful: fewer STOP/answer/proof interventions; faster approvals; higher trust in proof claims.
- Remaining uncertainty: Some worker environments may not support hard enforcement; still require durable 'paused' state and human confirmation to resume.
- Falsifier: Transcripts still show tool execution after explicit stop or pause instructions.

## Acceptance Criteria
- When STOP/PAUSE is issued, subsequent transcript shows no tool runs until RESUME.
- Intervention-time attribution can detect explicit wait gates and does not charge stall loss in those cases.

## What Does Not Count
- A prose reminder with no enforceable gate or template.
- Off-lane proof relabeled as regression without lane id + artifact.

## References
- [../APPROACH.md](../APPROACH.md)
- [../BURDEN-ANALYSIS.md](../BURDEN-ANALYSIS.md)
- [../ORTHOGONAL-SOLUTIONS-MATRIX.md](../ORTHOGONAL-SOLUTIONS-MATRIX.md)
- [../SessionExcerpts/INDEX.json](../SessionExcerpts/INDEX.json)
