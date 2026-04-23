# Orthogonal Solutions Matrix

## Problems
- PROBLEM-0001: Regression proof is not durably pinned to the human default lane.
- PROBLEM-0002: Ownership, stop/go, and approval gates are too soft.
- PROBLEM-0003: The system does not hold direct-answer and first-disagreement debugging discipline under pressure.

## Options
### PROBLEM-0001
- OPTION-A: Default-lane regression proof gate. Require a named repo regression lane plus lane-valid artifacts before any regression closure claim.
- OPTION-B: Canonical default-lane runner surface. Standardize one durable runner/output shape that makes override lanes explicit instead of silently adjacent.

### PROBLEM-0002
- OPTION-A: Hard STOP/PAUSE ownership state. Treat stop, pause, and no-homework-back boundaries as explicit execution state rather than conversational tone.
- OPTION-B: Approval packet surface. Make approval asks include diff, links, scope, and one explicit approval question by default.

### PROBLEM-0003
- OPTION-A: Direct-answer-first response rule. Force yes/no, agree/disagree, and simple factual questions to be answered on line one before any elaboration.
- OPTION-B: Debugging disagreement-seam contract. Require named disagreement values and upstream writer tracing before symptom iteration continues.

## Winners And Rollout Order
1. PROBLEM-0001: OPTION-A then OPTION-B.
2. PROBLEM-0002: OPTION-A then OPTION-B.
3. PROBLEM-0003: OPTION-A then OPTION-B.

## Mapping Back To Burden Clusters
- PROBLEM-0001 intercepts `default_lane_truth` and most of `proof_surface_integrity`.
- PROBLEM-0002 intercepts `agent_continuity`, `approval_surface`, and much of `durable_learning`.
- PROBLEM-0003 intercepts `answer_shape` and `root_cause_debugging`.

## Option Tasks
- [Option-Tasks/INDEX.md](./Option-Tasks/INDEX.md)
