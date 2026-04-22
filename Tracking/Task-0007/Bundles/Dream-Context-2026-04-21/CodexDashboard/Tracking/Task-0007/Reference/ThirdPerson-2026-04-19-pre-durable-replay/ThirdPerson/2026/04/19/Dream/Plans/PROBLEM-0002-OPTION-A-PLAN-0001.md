# Problem 0002 Option A Plan 0001

## Planning Intent

This file turns Problem `0002`, Option `A. Answer-first wording rule` from [ORTHOGONAL-SOLUTIONS-MATRIX.md](../ORTHOGONAL-SOLUTIONS-MATRIX.md) into a bounded implementation sequence.

It is an alternative route, not the selected winner task.

## Summary

Revise the shared prompt wording so explicit questions get a direct answer first, before any process talk or extra framing.

## Fixed Defaults

- scope: prompt wording only
- canonical homes:
  - `C:\Users\gregs\.codex\Orchestration\Prompts\TASK-HARVESTER.md`
  - `C:\Users\gregs\.codex\Orchestration\Prompts\TASK-LEADER.md`
  - `C:\Users\gregs\.codex\Orchestration\Prompts\RESEARCH-LEADER.md`
  - `C:\Users\gregs\.codex\Orchestration\Prompts\IMPLEMENTATION-LEADER.md`
  - `C:\Users\gregs\.codex\Orchestration\Prompts\REGRESSION-LEADER.md`
  - `C:\Users\gregs\.codex\Orchestration\Prompts\DEBUG-LEADER.md`
- no new workflow doc or validator in this rollout

## Pass Plan

### Pass 0000 - Shared Direct-Question Rule

Goal:

- make `answer first` an explicit shared instruction in the main leader prompts

Build:

- add one compact direct-question rule to the six prompt templates
- require the first response block to answer the asked question before process narration
- add one anti-pattern line that bans leading with status chatter when the user asked a direct question

Unit Proof:

- all touched prompts carry the same direct-question rule
- none of the prompts leave room for `process first, answer later`

Exit Bar:

- the shared prompt layer now says direct questions must get direct answers first

### Pass 0001 - Examples And Edge Cases

Goal:

- make the rule easy to apply consistently

Build:

- add one short `yes/no` example and one short `where is it` example to the most natural shared prompt home
- add one note that short answers can be followed by detail only after the answer block lands

Unit Proof:

- the examples are brief and obviously satisfy the rule
- the edge-case note still allows follow-on detail after the direct answer

Exit Bar:

- a reviewer can tell what compliant answer shape looks like without inference

## Testing Strategy

- check that the touched prompts use the same `answer first` phrase
- reject wording that still allows process narration before the answer block

## Deferred Work

- structured reply modes
- explicit answer-shape validation
- machine-checkable output checks
