# Solution Task 0006 Plan 0001

## Planning Intent

This file turns [SOLUTION-TASK-0006.md](../Task-Candidates/SOLUTION-TASK-0006.md) into a bounded implementation sequence.

It describes the intended route to done, not reply history.

## Summary

Add one shared answer-shape workflow so explicit questions get a direct first answer block before process narration.

## Fixed Defaults

- workflow home:
  - `C:\Users\gregs\.codex\Orchestration\ANSWER-SHAPE-WORKFLOW.md`
- reply modes:
  - `yes_no_reason`
  - `agree_disagree_reason`
  - `why_reason`
  - `short_answer`
- first-block rule:
  - answer the direct question before process narration
- short-answer limit:
  - first block no longer than three sentences

## Pass Plan

### Pass 0000 - Answer Workflow Contract

Goal:

- define the shared reply modes and first-block rule

Build:

- add `ANSWER-SHAPE-WORKFLOW.md`
- define the four reply modes and the three-sentence short-answer rule
- define what does not count, such as answering later in the message

Unit Proof:

- the workflow doc names all four reply modes
- the workflow doc names the first-block rule and short-answer limit

Exit Bar:

- a reviewer can tell what a compliant direct-answer response looks like

### Pass 0001 - Prompt Adoption

Goal:

- make the shared answer rule apply across the roles that talk to the human

Build:

- update `TASK-HARVESTER.md`
- update `TASK-LEADER.md`
- update `RESEARCH-LEADER.md`
- update `IMPLEMENTATION-LEADER.md`
- update `REGRESSION-LEADER.md`
- update `DEBUG-LEADER.md`

Unit Proof:

- each prompt references or adopts the shared workflow
- none of the prompts define a competing answer-shape rule

Exit Bar:

- the human-facing roles all use the same direct-answer contract

### Pass 0002 - Examples And Boundary Cases

Goal:

- make the workflow easy to apply in real conversations

Build:

- add one example for each reply mode
- add one example showing truthful uncertainty without evasive process narration
- add one anti-example showing a failed first block

Unit Proof:

- examples cover all four modes
- the anti-example clearly fails the first-block rule

Exit Bar:

- a reviewer can judge whether a reply met the shared answer contract

## Testing Strategy

- compare the workflow doc and role prompts for consistent direct-answer language
- reject any rollout that leaves `short_answer` undefined
- reject any rollout that allows process narration before the answer in the direct-question path

## Deferred Work

Keep these out of this rollout unless the task expands intentionally:

- model-side output classification
- automatic transcript linting
- non-English answer-mode variants
