# Problem 0002 Option C Plan 0001

## Planning Intent

This file turns Problem `0002`, Option `C. Unanswered-question validator` from [ORTHOGONAL-SOLUTIONS-MATRIX.md](../ORTHOGONAL-SOLUTIONS-MATRIX.md) into a bounded implementation sequence.

It is an alternative route, not the selected winner task.

## Summary

Add one shared validation step that detects direct-question forms and blocks a reply unless the first response block satisfies the requested answer shape.

## Fixed Defaults

- scope: shared workflow plus prompt enforcement
- canonical homes:
  - `C:\Users\gregs\.codex\Orchestration\ANSWER-SHAPE-WORKFLOW.md`
  - `C:\Users\gregs\.codex\Orchestration\Prompts\TASK-HARVESTER.md`
  - `C:\Users\gregs\.codex\Orchestration\Prompts\TASK-LEADER.md`
  - `C:\Users\gregs\.codex\Orchestration\Prompts\RESEARCH-LEADER.md`
  - `C:\Users\gregs\.codex\Orchestration\Prompts\IMPLEMENTATION-LEADER.md`
  - `C:\Users\gregs\.codex\Orchestration\Prompts\REGRESSION-LEADER.md`
  - `C:\Users\gregs\.codex\Orchestration\Prompts\DEBUG-LEADER.md`
- validator checks:
  - detect explicit question form
  - infer required first-block shape
  - block process-preface before the answer
  - require the answer block to fully satisfy the question type

## Pass Plan

### Pass 0000 - Validator Contract

Goal:

- define the validation rules clearly enough that prompts can use them consistently

Build:

- add `ANSWER-SHAPE-WORKFLOW.md` with a detection table for direct-question forms
- define first-block requirements for `yes_no_reason`, `where_is_it`, `what_changed`, and `do_you_agree`
- add one failure example for `I asked you a question` style misses

Unit Proof:

- each question form maps to one required first-block shape
- the failure example is obviously invalid under the rule

Exit Bar:

- the validator logic is concrete enough to be adopted as a prompt gate

### Pass 0001 - Prompt Gate Adoption

Goal:

- make the validator a real pre-send requirement in the leader prompts

Build:

- update the leader prompts so they must run the answer-shape check before final output
- require the first response block to land before any extra framing or status
- add one line that unanswered-question failures must be corrected before sending

Unit Proof:

- the prompts all reference the same validation rule
- none of the prompts allow process narration to satisfy the question check

Exit Bar:

- direct-question repair turns are prevented at the first response boundary

### Pass 0002 - Packet-Backed Evals

Goal:

- tie the validator to the real April 19 failure pattern

Build:

- preserve two or three packet-backed examples as regression-style answer-shape evals
- note what first block would have counted for each one

Unit Proof:

- each eval maps to a concrete question form and answer shape
- the expected first block is short and testable

Exit Bar:

- a later reviewer can tell whether the validator would have caught the known repair turns

## Testing Strategy

- compare the question-form table to the prompt gate wording
- treat missing expected first-block shapes as rollout failures

## Deferred Work

- richer structured reply modes beyond direct questions
- machine-side telemetry on answer misses
- UI review tools for answer-shape audits
