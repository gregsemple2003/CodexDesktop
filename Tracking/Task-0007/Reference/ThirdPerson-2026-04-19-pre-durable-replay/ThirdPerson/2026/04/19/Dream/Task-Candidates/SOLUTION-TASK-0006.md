# Solution Task 0006

## Title

Add a shared answer-shape workflow for direct questions.

## Summary

The packet shows many small repair turns where the human had to say `I asked you a question`, `short answer`, or `use small words` because the reply started with process narration instead of an answer.

This is a concrete implementation task. It will define one shared answer-shape workflow and wire it into the prompt roles that talk directly to the human.

## Goals

- answer explicit questions in the first response block
- support common short-answer shapes
- move extra process detail after the answer
- reduce repair turns caused only by wrong answer shape

## Non-Goals

- forcing every response to be one sentence
- banning detail when the human asks for depth
- solving approval, ownership, or proof problems inside this task

## Constraints And Baseline

- the rule must still allow nuance when nuance is requested
- the answer-first block must stay readable and plain
- the workflow must live in shared docs, not in one isolated prompt

## Proposed Changes

- add `C:\Users\gregs\.codex\Orchestration\ANSWER-SHAPE-WORKFLOW.md`
- update [TASK-HARVESTER.md](../../../../../../../../Prompts/TASK-HARVESTER.md)
- update [TASK-LEADER.md](../../../../../../../../Prompts/TASK-LEADER.md)
- update `C:\Users\gregs\.codex\Orchestration\Prompts\RESEARCH-LEADER.md`
- update `C:\Users\gregs\.codex\Orchestration\Prompts\IMPLEMENTATION-LEADER.md`
- update `C:\Users\gregs\.codex\Orchestration\Prompts\REGRESSION-LEADER.md`
- update `C:\Users\gregs\.codex\Orchestration\Prompts\DEBUG-LEADER.md`
- define these exact reply modes in `ANSWER-SHAPE-WORKFLOW.md`: `yes_no_reason`, `agree_disagree_reason`, `why_reason`, and `short_answer`
- define this first-block rule:
  - when the user asks a direct question, the first paragraph must answer it before any process narration
  - when the user asks for a short answer, the first block must be at most three sentences
  - extra workflow detail may only appear after the direct answer block

## Expected Resolution

- direct questions get direct answers first
- short-answer requests stay short
- workflow narration moves after the answer instead of before it

## What Does Not Count

- answering the question later in the message
- giving adjacent context while skipping the direct answer
- writing a long preamble before the answer
- saying `here is some nuance` when the user asked for a short answer

## Implementation Home

Implement the shared answer contract under:

- `C:\Users\gregs\.codex\Orchestration\ANSWER-SHAPE-WORKFLOW.md`
- [TASK-HARVESTER.md](../../../../../../../../Prompts/TASK-HARVESTER.md)
- [TASK-LEADER.md](../../../../../../../../Prompts/TASK-LEADER.md)
- `C:\Users\gregs\.codex\Orchestration\Prompts\RESEARCH-LEADER.md`
- `C:\Users\gregs\.codex\Orchestration\Prompts\IMPLEMENTATION-LEADER.md`
- `C:\Users\gregs\.codex\Orchestration\Prompts\REGRESSION-LEADER.md`
- `C:\Users\gregs\.codex\Orchestration\Prompts\DEBUG-LEADER.md`

## Proof Plan

- check that the workflow doc defines the four reply modes and the first-block rule
- check that each listed prompt references the shared workflow instead of inventing its own wording
- check that the short-answer rule names the three-sentence limit

## Acceptance Criteria

- `ANSWER-SHAPE-WORKFLOW.md` exists and defines the reply modes `yes_no_reason`, `agree_disagree_reason`, `why_reason`, and `short_answer`
- the workflow doc says a direct question must be answered in the first paragraph before any process narration
- the workflow doc says `short_answer` means the first block is at most three sentences
- each prompt listed in `Implementation Home` references or adopts the shared answer-shape rule
- the written design makes clear that eventually answering the question does not count if the first block missed it

## References

- [TASK-CREATE.md](../../../../../../../../Processes/TASK-CREATE.md)
- [BURDEN-ANALYSIS.md](../BURDEN-ANALYSIS.md)
- [ORTHOGONAL-SOLUTIONS-MATRIX.md](../ORTHOGONAL-SOLUTIONS-MATRIX.md)
