# Solution Task 0003

## Title

Add a reply-shape preflight for direct questions

## Summary

The packet shows repeated `STOP` turns because explicit questions were not answered in the asked shape.
This task adds a lightweight preflight that forces the first line of a reply to answer direct questions before extra detail appears.

## Goals

- detect explicit direct questions in the active human turn
- require the opening line to answer those questions directly
- keep support detail optional and secondary

## Non-Goals

- rewriting all response style guidance
- banning longer explanations when they are useful
- building a separate Q-and-A appendix system

## Implementation Home

- shared response prompt files under `C:\Users\gregs\.codex\Orchestration\Prompts\`
- any shared reply-audit or send-time checklist used by the workflow
- `C:\Users\gregs\.codex\AGENTS.md` only if one small cross-repo rule needs to be documented there

## Constraints And Baseline

- explicit questions should get direct answers first, especially yes or no and short factual asks
- if the human asks for small words, the opening answer should stay plain
- support detail can still follow after the direct answer

## Proposed Changes

- define a `Direct Question Preflight` step in shared response prompts
- require the preflight to list active direct questions and the opening-line answer for each one
- add failure examples where the reply opens with context, summary, or defensiveness instead of the answer
- add passing examples for short yes or no, factual, and definitional questions

## Expected Resolution

Users should stop needing extra turns just to get the first sentence they asked for.
Longer support can still be present, but the main answer will land first.

## What Does Not Count

- a vague reminder to be concise
- a response that answers the question only halfway through the paragraph
- examples with no send-time preflight step

## Acceptance Criteria

- shared prompts define a `Direct Question Preflight` that runs before final reply output
- the preflight requires the opening line to answer explicit direct questions
- prompt examples include both failing and passing answer shapes
- the workflow can point to unanswered direct questions before the reply is sent

## Proof Plan

- test the preflight against April 19 question turns such as `Agree/disagree?` and `what's still wrong`
- confirm the check fails when the opening line is indirect
- confirm it passes when the first line answers the question directly

## References

- `../ORTHOGONAL-SOLUTIONS-MATRIX.md`
- `../Plans/PROBLEM-0002-OPTION-B-PLAN-0001.md`

## Plan Addendum

Start with only the direct-question cases that caused real packet friction.
The preflight should list the active questions, require an opening answer, and catch evasive openings before send time.
After that, add a few short passing and failing examples so the rule is easy to apply.
