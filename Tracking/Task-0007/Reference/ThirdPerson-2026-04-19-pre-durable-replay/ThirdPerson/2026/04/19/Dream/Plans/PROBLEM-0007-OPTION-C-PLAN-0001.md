# Problem 0007 Option C Plan 0001

## Planning Intent

This file turns Problem `0007`, Option `C. Root-cause claim verifier` from [ORTHOGONAL-SOLUTIONS-MATRIX.md](../ORTHOGONAL-SOLUTIONS-MATRIX.md) into a bounded implementation sequence.

It is the matrix option that fed [SOLUTION-TASK-0005.md](../Task-Candidates/SOLUTION-TASK-0005.md).

## Summary

Add a narrow verification rule that blocks `root cause found` or `root cause fixed` language unless the debug artifacts show the exact disagreement seam and the writer chain that produced the bad state.

## Fixed Defaults

- scope: shared debugging process and prompts
- canonical homes:
  - `C:\Users\gregs\.codex\Orchestration\Processes\DEBUGGING.md`
  - `C:\Users\gregs\.codex\Orchestration\Prompts\DEBUG-LEADER.md`
  - `C:\Users\gregs\.codex\Orchestration\Prompts\DEBUG-WORKER.md`
  - `C:\Users\gregs\.codex\Orchestration\Prompts\TASK-LEADER.md`
- required verification fields:
  - `Root Cause Claim Status`
  - `Disagreement Seam`
  - `Writer Chain`
  - `Why This Writer Produces The Bad State`
- only `verified_root_cause` allows root-cause claim wording

## Pass Plan

### Pass 0000 - Verification Contract

Goal:

- define the minimum evidence needed for a root-cause claim

Build:

- update `Processes/DEBUGGING.md` with the four verification fields
- define allowed `Root Cause Claim Status` values
- add one valid and one invalid root-cause closeout example

Unit Proof:

- the field names and status values are explicit and finite
- the invalid example clearly fails because it lacks a real writer chain

Exit Bar:

- root-cause language now has a concrete contract instead of relying on vibes

### Pass 0001 - Prompt Gate Adoption

Goal:

- enforce the verification contract where root-cause wording is generated

Build:

- update `DEBUG-LEADER.md`, `DEBUG-WORKER.md`, and `TASK-LEADER.md` so `root cause found` or equivalent wording is blocked unless `Root Cause Claim Status=verified_root_cause`
- require the disagreement seam and writer chain to be named directly in the closeout artifact

Unit Proof:

- all three prompts use the same field names and status gate
- none of the prompts allow root-cause language on symptom-only evidence

Exit Bar:

- premature root-cause closure is stopped at the debug-closeout boundary

### Pass 0002 - Packet-Backed Evals

Goal:

- tie the verifier to the real April 19 debugging drift

Build:

- preserve two packet-backed examples of symptom iteration vs verified narrowing
- note which verification field would have failed in each non-root-cause case

Unit Proof:

- each eval maps to a concrete missing field or wrong status
- the verified case shows the full writer chain

Exit Bar:

- a reviewer can judge whether the verifier would have caught the known drift pattern

## Testing Strategy

- keep verification field names aligned across process docs and prompts
- reject any rollout where `verified_root_cause` can be claimed without a writer chain

## Deferred Work

- separate disagreement-trace artifact requirement
- machine-checkable debug closeout schema
- dashboard root-cause review tooling
