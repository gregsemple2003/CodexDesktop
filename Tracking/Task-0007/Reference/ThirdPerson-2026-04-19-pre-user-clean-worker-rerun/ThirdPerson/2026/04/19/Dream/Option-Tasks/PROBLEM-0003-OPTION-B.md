## Title
Question capture checklist for multi-question messages

## Summary
Extract questions into a checklist and answer each before proceeding, to prevent omissions.

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
- For multi-question messages, extract each explicit question into a short checklist.
- Answer each extracted question directly; then proceed to next steps.

## Burden Reduction Contract
- Burden being reduced: repeated human corrections about lane/proof/boundaries/answers.
- Causal claim: without durable gates and templates, the workflow drifts and the human must intervene to restate constraints.
- Evidence:
  - hie-...-5090 (Stop evading the question.)
  - hie-...-5272..5300 cluster (definition questions)
- Why this mechanism: it changes the default workflow so the agent cannot easily 'slide' into off-lane proof or continue past stops.
- Human relief if successful: fewer STOP/answer/proof interventions; faster approvals; higher trust in proof claims.
- Remaining uncertainty: Poor formatting could increase burden; implement with a strict small-checklist cap.
- Falsifier: The checklist becomes longer than the answers and increases human frustration.

## Acceptance Criteria
- Checklist extraction is grounded (no invented questions) and each question has a direct answer.
- Checklist stays short and does not become a new burden surface for simple questions.

## What Does Not Count
- A prose reminder with no enforceable gate or template.
- Off-lane proof relabeled as regression without lane id + artifact.

## References
- [../APPROACH.md](../APPROACH.md)
- [../BURDEN-ANALYSIS.md](../BURDEN-ANALYSIS.md)
- [../ORTHOGONAL-SOLUTIONS-MATRIX.md](../ORTHOGONAL-SOLUTIONS-MATRIX.md)
- [../SessionExcerpts/INDEX.json](../SessionExcerpts/INDEX.json)
