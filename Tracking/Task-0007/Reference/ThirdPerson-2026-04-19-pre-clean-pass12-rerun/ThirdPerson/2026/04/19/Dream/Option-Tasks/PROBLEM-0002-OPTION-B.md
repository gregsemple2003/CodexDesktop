## Title
Gate-aware approval surface template (diff + links + explicit approval question)

## Summary
When asking for approval, always provide the review surface in one place to avoid human reconstruction work.

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
- Adopt a plan-approval template that includes: (1) concise diff summary, (2) concrete file links, (3) one explicit approval question.
- When no diff is possible, state why and provide an alternate review artifact (e.g., rendered diff excerpt).

## Burden Reduction Contract
- Burden being reduced: repeated human corrections about lane/proof/boundaries/answers.
- Causal claim: without durable gates and templates, the workflow drifts and the human must intervene to restate constraints.
- Evidence:
  - EXCERPT-0003 (no diff/links)
- Why this mechanism: it changes the default workflow so the agent cannot easily 'slide' into off-lane proof or continue past stops.
- Human relief if successful: fewer STOP/answer/proof interventions; faster approvals; higher trust in proof claims.
- Remaining uncertainty: Diff generation depends on git state; needs a defined fallback when the worktree is dirty.
- Falsifier: Humans still have to ask for links/diffs for approval gates in later packets.

## Acceptance Criteria
- Approval requests consistently include diff + links + one explicit approval question.
- Later transcripts stop showing 'give me links/diff' follow-ups for the same approval gate.

## What Does Not Count
- A prose reminder with no enforceable gate or template.
- Off-lane proof relabeled as regression without lane id + artifact.

## References
- [../APPROACH.md](../APPROACH.md)
- [../BURDEN-ANALYSIS.md](../BURDEN-ANALYSIS.md)
- [../ORTHOGONAL-SOLUTIONS-MATRIX.md](../ORTHOGONAL-SOLUTIONS-MATRIX.md)
- [../SessionExcerpts/INDEX.json](../SessionExcerpts/INDEX.json)
