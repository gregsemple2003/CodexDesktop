# PROBLEM-0001 OPTION-A: Default-lane regression proof gate

## Title

Default-lane regression proof gate

## Summary

Require a named repo regression lane plus lane-valid artifacts before any regression closure claim.

## Goals

- Reduce the burden described under PROBLEM-0001 in the April 19 ThirdPerson intervention packet.
- Convert the burden into a durable contract, gate, or artifact rather than a wording-only reminder.
- Preserve the packet's truth, compassion, and tolerance constraints.

## Non-Goals

- Rewriting the packet itself as a closure substitute.
- Solving unrelated repo defects outside the burden described here.
- Pretending missing transcript context is available when it is not.

## Implementation Home

Shared orchestration docs and/or repo-local ThirdPerson docs/artifacts, depending on where the durable enforcement boundary belongs.

## Burden Being Reduced

The human has to restate what counts as valid regression closure on the default runtime lane and reject adjacent proof that should never have counted as closure.

## Causal Claim

The packet suggests this burden persists because extend the shared regression closeout contract so thirdperson closure fails unless the artifact names lane a or another repo-canonical lane from regression.md and links the required runtime proof.

## Evidence

- Packet evidence ids: hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-3325, hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-3431, hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4909, hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-3814
- Dream burden analysis: [../BURDEN-ANALYSIS.md](../BURDEN-ANALYSIS.md)
- Matrix entry: [../ORTHOGONAL-SOLUTIONS-MATRIX.md](../ORTHOGONAL-SOLUTIONS-MATRIX.md)

## Why This Mechanism

Extend the shared regression closeout contract so ThirdPerson closure fails unless the artifact names Lane A or another repo-canonical lane from REGRESSION.md and links the required runtime proof.

## Human Relief If Successful

The human should no longer have to carry this burden repeatedly inside one day. The system should either block the weak behavior up front or present a review surface that makes the next decision cheap and trustworthy.

## Remaining Uncertainty

- Missing transcript neighborhoods mean some assistant-side failure mechanics remain only partially observed.
- A contract without enforcement hooks could still be gamed unless the acceptance criteria name the actual gate.

## Falsifier

If the same burden still recurs in future packets after the new gate or artifact is in place, the causal claim is incomplete or the enforcement boundary is too weak.

## Proposed Changes

- Add or revise the durable docs, prompts, schemas, or closeout artifacts needed to enforce this option's mechanism.
- Make the enforcement boundary explicit so weak closure does not still count.
- Update any required packet or task-owned artifacts that prove the mechanism exists.

## Acceptance Criteria

- A later reviewer can point to the exact durable gate or artifact that implements this option.
- The resulting contract blocks the weak closure path described in the burden analysis.
- The proof path for this option does not depend on chat-memory reconstruction.
