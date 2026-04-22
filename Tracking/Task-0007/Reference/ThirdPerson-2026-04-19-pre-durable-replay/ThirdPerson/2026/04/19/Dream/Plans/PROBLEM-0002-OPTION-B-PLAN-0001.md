# Problem 0002 Option B Plan 0001

## Planning Intent

This file turns Problem `0002`, Option `B. Structured reply modes` from [ORTHOGONAL-SOLUTIONS-MATRIX.md](../ORTHOGONAL-SOLUTIONS-MATRIX.md) into a bounded implementation sequence.

It is an alternative route, not the selected winner task.

## Summary

Add one shared answer-shape workflow that defines a small set of reply modes for common direct-question patterns.

## Fixed Defaults

- scope: shared workflow plus prompt adoption
- canonical homes:
  - `C:\Users\gregs\.codex\Orchestration\ANSWER-SHAPE-WORKFLOW.md`
  - `C:\Users\gregs\.codex\Orchestration\Prompts\TASK-LEADER.md`
  - `C:\Users\gregs\.codex\Orchestration\Prompts\RESEARCH-LEADER.md`
  - `C:\Users\gregs\.codex\Orchestration\Prompts\IMPLEMENTATION-LEADER.md`
- required modes:
  - `short_answer`
  - `yes_no_reason`
  - `agree_disagree_reason`
  - `list_requested_items`
  - `quote_then_reason`
- this option defines modes but does not block noncompliant output automatically

## Pass Plan

### Pass 0000 - Mode Contract

Goal:

- define the reply modes once in a shared durable home

Build:

- add `ANSWER-SHAPE-WORKFLOW.md`
- define when each mode should be chosen
- define the required first block for each mode with one short example

Unit Proof:

- each mode has a trigger and a first-block shape
- the examples are short enough to be usable as live guidance

Exit Bar:

- the repo has one canonical answer-shape vocabulary instead of prompt-local phrasing

### Pass 0001 - Prompt Adoption

Goal:

- make the main leader prompts use the same answer-shape vocabulary

Build:

- update the leader prompts to select a reply mode when the user asks a direct question
- require the first response block to match the chosen mode before extra detail

Unit Proof:

- prompt wording matches the mode names from the workflow doc
- no prompt invents a conflicting local mode name

Exit Bar:

- direct-question handling becomes more consistent without needing a full validator yet

## Testing Strategy

- verify that mode names are identical across workflow and prompts
- treat missing examples or ambiguous trigger rules as rollout failures

## Deferred Work

- validator logic that blocks the wrong first block
- dashboard telemetry on answer-shape failures
- repo-local mode extensions beyond the shared set
