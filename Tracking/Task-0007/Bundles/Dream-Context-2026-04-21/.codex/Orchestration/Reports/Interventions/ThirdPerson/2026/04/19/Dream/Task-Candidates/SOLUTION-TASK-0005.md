# Solution Task 0005

## Title

Add an answer-first gate for direct questions in active task work

## Summary

The packet shows repeated repair turns where the human asked a direct question and got summary text, hedges, or side paths first.
This task adds one shared answer-first rule so direct questions are answered in line one before any supporting detail.

## Goals

- Reduce "I asked you a question" repair turns.
- Make yes or no, one-fact, and short scoped answers land before explanation.
- Preserve the option to add useful detail after the direct answer.

## Non-Goals

- Forcing every reply into a one-line format.
- Adding a clarification turn before obvious direct questions.
- Removing context from replies that genuinely need follow-up detail.

## Constraints And Baseline

Current truth:

- The packet records repeated human corrections about answer shape and requests for short answers or small words.
- The burden was not lack of prose; it was lack of a direct answer at the start.

Hard constraints:

- the rule must trigger on direct questions during active task work
- extra detail may follow, but only after the answer line lands
- blunt, partial, or impatient phrasing must still count as a real question

## Implementation Home

- `C:\Users\gregs\.codex\Orchestration\Prompts\TASK-LEADER.md`
- `C:\Users\gregs\.codex\Orchestration\Prompts\IMPLEMENTATION-LEADER.md`
- `C:\Users\gregs\.codex\Orchestration\Prompts\IMPLEMENTER.md`
- `C:\Users\gregs\.codex\Orchestration\Prompts\DEBUG-LEADER.md`
- `C:\Users\gregs\.codex\Orchestration\Prompts\DEBUG-WORKER.md`
- `C:\Users\gregs\.codex\Orchestration\Prompts\REGRESSION-LEADER.md`
- `C:\Users\gregs\.codex\Orchestration\Prompts\REGRESSION-TESTER.md`
- `C:\Users\gregs\.codex\Orchestration\Prompts\AUDITOR.md`

## Proposed Changes

- Add one shared answer-first rule to the human-facing active-work prompts above:
  - detect direct-question turns
  - classify the expected answer shape as `yes_no`, `one_fact`, or `short_scoped`
  - require the first response line to contain that answer
- Add packet-backed examples that cover:
  - direct accountability questions
  - "what is still wrong"
  - "agree or disagree"
  - "small words" or short-answer requests
- Add an audit check that flags when a direct answer is buried below summary text or preamble.

## Expected Resolution

When the human asks a direct question, the first line should answer it plainly.
If more explanation helps, it can follow after the answer instead of blocking it.

## What Does Not Count

- A shorter summary that still delays the answer until later lines.
- A clarification question before answering an obvious yes or no request.
- A line-one hedge such as "it depends" when the packet context supports a direct answer.

## Acceptance Criteria

- Each prompt in the implementation home includes the answer-first rule or an explicit reference to one shared rule with the same behavior.
- The prompt set includes examples for `yes_no`, `one_fact`, and `short_scoped` answers.
- The audit rule flags a response as noncompliant when line one does not answer a direct packet-style question.
- In prompt tests or worked examples, questions such as "agree or disagree" and "what is still wrong" are answered directly in line one before any explanation.
- The rule still allows a short follow-up explanation after the answer line.

## Proof Plan

- Convert at least four packet-backed direct questions into prompt examples and verify the line-one output shape.
- Review one simulated noncompliant answer and verify the audit rule flags it.
- Review one compliant answer with a short follow-up paragraph and verify it still passes.

## References

- `..\BURDEN-ANALYSIS.md`
- `..\ORTHOGONAL-SOLUTIONS-MATRIX.md`
- `..\Plans\PROBLEM-0002-OPTION-A-PLAN-0001.md`

## Plan Addendum

Chosen plan: add an answer-first response gate so direct questions get a direct answer in the first line before any extra detail.

Implementation notes from the selected plan:

- detect direct questions in the active turn
- force the first response line to answer the question directly
- allow supporting detail only after the first-line answer is complete

Rollout:

1. Define the direct-question patterns that trigger the gate.
2. Write a minimal answer format for yes or no, factual, and scoped short-answer requests.
3. Add the gate to the shared response workflow used for active task work.
4. Test it against packet examples such as direct accountability questions and "small words" requests.
5. Keep the longer explanation path available only after the answer line lands first.
