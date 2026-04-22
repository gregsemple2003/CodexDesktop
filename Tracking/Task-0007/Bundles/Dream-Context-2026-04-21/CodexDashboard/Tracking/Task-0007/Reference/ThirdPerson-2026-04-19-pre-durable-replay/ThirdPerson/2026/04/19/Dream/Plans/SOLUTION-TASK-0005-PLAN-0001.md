# Solution Task 0005 Plan 0001

## Planning Intent

This file turns [SOLUTION-TASK-0005.md](../Task-Candidates/SOLUTION-TASK-0005.md) into a bounded implementation sequence.

It describes the intended route to done, not debug history.

## Summary

Change the shared debugging contract so `root cause found` is only allowed when the bug artifact shows the concrete runtime disagreement and the traced writer chain behind it.

## Fixed Defaults

- workflow home:
  - `Processes\DEBUGGING.md`
- prompt homes:
  - `Prompts\DEBUG-LEADER.md`
  - `Prompts\DEBUG-WORKER.md`
  - `Prompts\TASK-LEADER.md`
- required bug sections:
  - `Expected Runtime State`
  - `Observed Runtime State`
  - `First Concrete Disagreement`
  - `Upstream Writer Chain`
  - `Evidence`
  - `Rejected Symptom Paths`
  - `Root Cause Claim Status`
- allowed statuses:
  - `hypothesis`
  - `narrowed`
  - `writer_chain_found`
  - `verified_root_cause`

## Pass Plan

### Pass 0000 - Debug Contract

Goal:

- define the required bug-artifact sections and status values

Build:

- update `DEBUGGING.md` with the section contract and the allowed statuses
- define the rule that root-cause wording is blocked until status reaches `verified_root_cause`

Unit Proof:

- every required section is named once in the shared contract
- each status has one clear meaning

Exit Bar:

- a reviewer can tell what must be present before a root-cause claim is allowed

### Pass 0001 - Prompt Enforcement

Goal:

- make the shared debug prompts enforce the contract

Build:

- update `DEBUG-LEADER.md`
- update `DEBUG-WORKER.md`
- update `TASK-LEADER.md`

Unit Proof:

- the prompts require the same bug sections and status rule as `DEBUGGING.md`
- the prompts do not allow symptom-only closure language

Exit Bar:

- root-cause claims are gated at the point where debug work is summarized or closed

### Pass 0002 - Examples And Anti-Examples

Goal:

- make the contract easy to audit in practice

Build:

- add one valid writer-chain example
- add one example that is still only `narrowed`
- add one anti-example showing why symptom relief is not root-cause proof

Unit Proof:

- the examples map cleanly to the allowed statuses
- the anti-example shows a rejected closure path

Exit Bar:

- a later reviewer can challenge vague root-cause language without rereading the packet

## Testing Strategy

- compare `DEBUGGING.md` and the debug prompts for the same section list and status values
- reject any rollout that allows `root cause found` before `verified_root_cause`
- reject any rollout that treats a symptom category as a writer chain

## Deferred Work

Keep these out of this rollout unless the task expands intentionally:

- automated code-path tracing
- custom debug UI
- broader evidence-manifest work outside the root-cause gate
