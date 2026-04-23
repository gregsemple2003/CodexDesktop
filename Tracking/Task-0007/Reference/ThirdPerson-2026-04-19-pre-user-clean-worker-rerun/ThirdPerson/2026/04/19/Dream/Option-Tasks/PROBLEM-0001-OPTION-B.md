## Title
Default-lane runner standardization (artifact-first proof)

## Summary
Provide a repeatable default-lane runner that produces canonical artifacts and makes lane overrides explicit.

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
- Standardize a default-lane runner entrypoint that archives evidence (screenshots/telemetry/logs) in a known artifact root.
- Embed lane id, map, pawn, and evidence identifiers in the artifact root manifest.

## Burden Reduction Contract
- Burden being reduced: repeated human corrections about lane/proof/boundaries/answers.
- Causal claim: without durable gates and templates, the workflow drifts and the human must intervene to restate constraints.
- Evidence:
  - EXCERPT-0004 (evidence invalidation)
  - EXCERPT-0001 (default-lane constraint)
- Why this mechanism: it changes the default workflow so the agent cannot easily 'slide' into off-lane proof or continue past stops.
- Human relief if successful: fewer STOP/answer/proof interventions; faster approvals; higher trust in proof claims.
- Remaining uncertainty: Some evidence (e.g., full-body screenshot framing) may require editor runtime surfaces and careful camera control.
- Falsifier: Runner artifacts still fail human evidence questions (e.g., truncated feet) in later packets.

## Acceptance Criteria
- A default-lane run produces a durable artifact root with lane id + evidence identifiers.
- The runner makes lane overrides explicit and non-ambiguous in outputs.

## What Does Not Count
- A prose reminder with no enforceable gate or template.
- Off-lane proof relabeled as regression without lane id + artifact.

## References
- [../APPROACH.md](../APPROACH.md)
- [../BURDEN-ANALYSIS.md](../BURDEN-ANALYSIS.md)
- [../ORTHOGONAL-SOLUTIONS-MATRIX.md](../ORTHOGONAL-SOLUTIONS-MATRIX.md)
- [../SessionExcerpts/INDEX.json](../SessionExcerpts/INDEX.json)
