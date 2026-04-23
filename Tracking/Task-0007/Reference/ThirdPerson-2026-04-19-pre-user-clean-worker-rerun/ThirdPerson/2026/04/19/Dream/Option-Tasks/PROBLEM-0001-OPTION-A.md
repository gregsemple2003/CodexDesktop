## Title
Default-lane regression proof gate (contract + checklist)

## Summary
Make it mechanically hard to claim regression closure without naming the lane and providing a default-lane proof artifact.

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
- Require selecting a lane id from ThirdPerson/REGRESSION.md for any 'regression' claim.
- Require emitting a durable artifact that includes lane id + evidence summary before claiming regression closure.
- Add a checklist item: if lane != human default lane, the proof must be labeled non-regression.

## Burden Reduction Contract
- Burden being reduced: repeated human corrections about lane/proof/boundaries/answers.
- Causal claim: without durable gates and templates, the workflow drifts and the human must intervene to restate constraints.
- Evidence:
  - EXCERPT-0001 (default-lane constraint)
  - hie-...-3325 (human observed invisible pawn in regression)
- Why this mechanism: it changes the default workflow so the agent cannot easily 'slide' into off-lane proof or continue past stops.
- Human relief if successful: fewer STOP/answer/proof interventions; faster approvals; higher trust in proof claims.
- Remaining uncertainty: Exact enforcement surface depends on which orchestration prompt bundle is used to make regression claims.
- Falsifier: New packets still show repeated 'regression must use the human default lane' interventions.

## Acceptance Criteria
- Regression claims include an explicit lane id and artifact reference.
- Off-lane runs cannot be presented as regression closure without failing the checklist/gate.

## What Does Not Count
- A prose reminder with no enforceable gate or template.
- Off-lane proof relabeled as regression without lane id + artifact.

## References
- [../APPROACH.md](../APPROACH.md)
- [../BURDEN-ANALYSIS.md](../BURDEN-ANALYSIS.md)
- [../ORTHOGONAL-SOLUTIONS-MATRIX.md](../ORTHOGONAL-SOLUTIONS-MATRIX.md)
- [../SessionExcerpts/INDEX.json](../SessionExcerpts/INDEX.json)
