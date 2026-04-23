# Orthogonal Solutions Matrix (ThirdPerson 2026-04-19)

This file is the compact final decision surface for the Dream annex after adversarial winner synthesis.

## Final Winner Resolution

## P-001 Default-Lane Proof And Evidence Gate

- `Source Burden Drivers`: `BD-001`
- `Mechanism Boundary`: any claimed ThirdPerson regression pass or closure citation
- `Acceptance Test`: a supporting lane or cropped visual cannot be written up as passing `TP-REG-001`, `TP-REG-002`, or `TP-REG-003`
- `Falsifier`: an off-lane or invalid-evidence run still closes as passed regression
- `Designed Options`:
  - [`P-001 / Option A`](./SOLUTION-DESIGN.md#option-a-repo-local-regression-claim-gate-with-evidence-disqualifiers)
  - [`P-001 / Option B`](./SOLUTION-DESIGN.md#option-b-shared-proof-claim-contract-in-shared-testing-docs)
- `Final Winner`: [`W-001`](./WINNER-SYNTHESIS.md#w-001-thirdperson-regression-claim-gate)
- `Seam Resolution`: shared proof-quality claim shape is `merged_with_justification`; proof-pack linting is `deferred_follow_on`

## P-002 Ownership Continuity And Pause-Latch Control

- `Source Burden Drivers`: `BD-002`
- `Mechanism Boundary`: any stop, idle, approval gate, or resume transition during a live task pass
- `Acceptance Test`: after a gap or STOP, durable state already shows owner, waiting or block reason, and whether continuation is latched off
- `Falsifier`: the human still has to wake work up or STOP or approval gates are ignored
- `Designed Options`:
  - [`P-002 / Option A`](./SOLUTION-DESIGN.md#option-a-shared-execution-state-gate-with-pause-latch)
  - [`P-002 / Option B`](./SOLUTION-DESIGN.md#option-b-shared-workflow-enforcement-without-schema-expansion)
- `Final Winner`: [`W-002`](./WINNER-SYNTHESIS.md#w-002-shared-execution-state-pauseownership-gate)
- `Seam Resolution`: owner continuity and STOP or approval latch are `merged_with_justification`; notification-only alerts are `rejected_with_reason`

## P-003 Direct-Answer-First Contract

- `Source Burden Drivers`: `BD-003`
- `Mechanism Boundary`: any explicit short-answer question
- `Acceptance Test`: future transcripts no longer require restated questions to get a first-sentence answer
- `Falsifier`: the user still has to repeat the question because the first sentence does not answer it
- `Designed Options`:
  - [`P-003 / Option A`](./SOLUTION-DESIGN.md#option-a-shared-answer-first-rule)
  - [`P-003 / Option B`](./SOLUTION-DESIGN.md#option-b-repo-local-direct-answer-clause-in-thirdperson-docs)
- `Final Winner`: [`W-003`](./WINNER-SYNTHESIS.md#w-003-shared-direct-answer-first-rule)
- `Seam Resolution`: yes or no, agree or disagree, and short status questions are `merged_with_justification`; ThirdPerson-only wording is `rejected_with_reason`

## P-004 Approval-Gate Review Packet

- `Source Burden Drivers`: `BD-004`
- `Mechanism Boundary`: any human approval request for plan, pass, or reopen changes
- `Acceptance Test`: the human can approve or reject without asking for links, diff context, or whether reopened work changed pass history
- `Falsifier`: approval still requires manual reconstruction or reopened work is silently folded into closed pass history
- `Designed Options`:
  - [`P-004 / Option A`](./SOLUTION-DESIGN.md#option-a-shared-approval-gate-packet-contract)
  - [`P-004 / Option B`](./SOLUTION-DESIGN.md#option-b-task-local-approval-template-only)
- `Final Winner`: [`W-004`](./WINNER-SYNTHESIS.md#w-004-approval-gate-packet-contract)
- `Seam Resolution`: diff/context/link packaging and pass-history disclosure are `merged_with_justification`; IDE stale-buffer behavior is `rejected_with_reason`

## P-005 First-Disagreement Debugging Gate

- `Source Burden Drivers`: `BD-005`
- `Mechanism Boundary`: before the next fix or closure claim on a reopened or surviving default-lane runtime defect
- `Acceptance Test`: the next such defect records a concrete disagreement and upstream writer chain before more fixing is justified
- `Falsifier`: the task still proceeds on symptom-only or category-only diagnosis
- `Designed Options`:
  - [`P-005 / Option A`](./SOLUTION-DESIGN.md#option-a-repo-local-first-disagreement-gate-in-bug-and-regression-artifacts)
  - [`P-005 / Option B`](./SOLUTION-DESIGN.md#option-b-automation-owned-runtime-checker-with-quantitative-fail-thresholds)
- `Final Winner`: [`W-005`](./WINNER-SYNTHESIS.md#w-005-repo-local-first-disagreement-debugging-gate)
- `Seam Resolution`: bug narrative discipline and regression rerun honesty are `merged_with_justification`; checker hardening is `deferred_follow_on`

## P-006 Durable Constraint Ledger

- `Source Burden Drivers`: `BD-006`
- `Mechanism Boundary`: any edit, build, or plan transition after a new hard constraint, approval gate, or repo-truth correction
- `Acceptance Test`: constraints and repo-truth refs appear durably before work resumes and are consulted later without human restatement
- `Falsifier`: a stated constraint or truth must be repeated later or is violated despite the durable record
- `Designed Options`:
  - [`P-006 / Option A`](./SOLUTION-DESIGN.md#option-a-task-local-active-constraints-ledger-in-current-state-artifacts)
  - [`P-006 / Option B`](./SOLUTION-DESIGN.md#option-b-repo-local-preflight-loader-for-stable-truth-only)
- `Final Winner`: [`W-006`](./WINNER-SYNTHESIS.md#w-006-task-local-active-constraints-ledger)
- `Seam Resolution`: transient hard constraints and stable repo-truth refs are `merged_with_justification`; repo-preflight-only reread is `rejected_with_reason`

## Final Rollout Order

1. [`W-001 ThirdPerson Regression Claim Gate`](./WINNER-SYNTHESIS.md#w-001-thirdperson-regression-claim-gate)
2. [`W-005 Repo-Local First-Disagreement Debugging Gate`](./WINNER-SYNTHESIS.md#w-005-repo-local-first-disagreement-debugging-gate)
3. [`W-006 Task-Local Active-Constraints Ledger`](./WINNER-SYNTHESIS.md#w-006-task-local-active-constraints-ledger)
4. [`W-004 Approval-Gate Packet Contract`](./WINNER-SYNTHESIS.md#w-004-approval-gate-packet-contract)
5. [`W-002 Shared Execution-State Pause/Ownership Gate`](./WINNER-SYNTHESIS.md#w-002-shared-execution-state-pauseownership-gate)
6. [`W-003 Shared Direct-Answer-First Rule`](./WINNER-SYNTHESIS.md#w-003-shared-direct-answer-first-rule)

This order reflects the attacker lane's final synthesis after forcing narrower boundaries and exact homes back into `SOLUTION-DESIGN.md`.
