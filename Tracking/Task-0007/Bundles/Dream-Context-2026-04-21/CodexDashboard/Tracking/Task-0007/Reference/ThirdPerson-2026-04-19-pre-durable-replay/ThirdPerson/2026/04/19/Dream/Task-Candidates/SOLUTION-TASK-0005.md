# Solution Task 0005

## Title

Require a disagreement trace before any root-cause claim.

## Summary

The packet shows debugging drift where the human had to pull the work back from symptom tweaks and restate that the task was root-cause tracing.

This is a concrete implementation task. It will change the shared debugging workflow and debug prompts so `root cause found` is only allowed when the bug artifact names the exact bad state, the first concrete disagreement, and the upstream writer chain that produced it.

## Goals

- make `root cause found` mean one specific thing
- force debug work to preserve the disagreement being traced
- keep symptom relief and root-cause proof clearly separate

## Non-Goals

- banning exploratory debugging
- requiring the permanent fix in the same task that proves the root cause
- building a formal proof system

## Constraints And Baseline

- shared debugging already prefers narrowing over random tweaking
- exploratory work still matters, but it should not be mislabeled as root-cause closure
- the rule must work with `BUG-<NNNN>.md` artifacts

## Proposed Changes

- update `C:\Users\gregs\.codex\Orchestration\Processes\DEBUGGING.md`
- update `C:\Users\gregs\.codex\Orchestration\Prompts\DEBUG-LEADER.md`
- update `C:\Users\gregs\.codex\Orchestration\Prompts\DEBUG-WORKER.md`
- update [TASK-LEADER.md](../../../../../../../../Prompts/TASK-LEADER.md)
- require every active `BUG-<NNNN>.md` that is pursuing root cause to include these exact sections: `Expected Runtime State`, `Observed Runtime State`, `First Concrete Disagreement`, `Upstream Writer Chain`, `Evidence`, `Rejected Symptom Paths`, and `Root Cause Claim Status`
- define these exact `Root Cause Claim Status` values: `hypothesis`, `narrowed`, `writer_chain_found`, and `verified_root_cause`
- add a gate rule that `root cause found`, `root cause fixed`, and similar wording are only allowed when `Root Cause Claim Status=verified_root_cause`
- require `Upstream Writer Chain` to name each writer step with code or artifact links, not just a prose summary

## Expected Resolution

- bug artifacts preserve the exact disagreement being traced
- symptom-only work stays labeled as hypothesis or narrowing
- root-cause claims name the writer chain that explains the bad state

## What Does Not Count

- a symptom description with no disagreement
- a likely fix with no writer chain
- a bug note that says `root cause found` while `Root Cause Claim Status` is still `hypothesis` or `narrowed`
- a writer chain with no links back to code or evidence

## Implementation Home

Implement the shared root-cause contract under:

- `C:\Users\gregs\.codex\Orchestration\Processes\DEBUGGING.md`
- `C:\Users\gregs\.codex\Orchestration\Prompts\DEBUG-LEADER.md`
- `C:\Users\gregs\.codex\Orchestration\Prompts\DEBUG-WORKER.md`
- [TASK-LEADER.md](../../../../../../../../Prompts/TASK-LEADER.md)

## Proof Plan

- check that `DEBUGGING.md` names the required `BUG` sections and status values
- check that the debug prompts enforce the same sections and status rule
- check that root-cause wording is blocked unless the status is `verified_root_cause`

## Acceptance Criteria

- `DEBUGGING.md`, `DEBUG-LEADER.md`, and `DEBUG-WORKER.md` all require the `BUG-<NNNN>.md` sections listed in `Proposed Changes`
- the shared debugging contract defines the status values `hypothesis`, `narrowed`, `writer_chain_found`, and `verified_root_cause`
- the debug prompts say `root cause found` or `root cause fixed` wording is only allowed when `Root Cause Claim Status=verified_root_cause`
- the contract requires `Upstream Writer Chain` to include links to code or durable artifacts for each hop
- the written design makes clear that symptom relief without a disagreement trace is not root-cause closure

## References

- [TASK-CREATE.md](../../../../../../../../Processes/TASK-CREATE.md)
- [BURDEN-ANALYSIS.md](../BURDEN-ANALYSIS.md)
- [ORTHOGONAL-SOLUTIONS-MATRIX.md](../ORTHOGONAL-SOLUTIONS-MATRIX.md)
