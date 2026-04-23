# PROBLEM-0003 OPTION-A: Answer-First Gate For Explicit Direct Questions

## Title

Require human-facing lifecycle replies to start with a direct answer when the human asks an explicit in-scope direct question.

## Summary

The burden packet shows repeated human correction of one narrow but costly failure:

- the system often answered around the question instead of in the requested shape, so the human had to restate the question and demand a direct answer

This rewrite narrows Option A to the exact first slice justified by the local durable sources:

- add an answer-first contract to the exact shared human-facing prompt files that already own lifecycle replies to the human
- make that contract fail closed at the first sentence of the user-facing reply for a narrow set of explicit direct-question classes

This slice does **not** include a linter or other tooling.
The local sources justify shared prompt and lifecycle enforcement.
They do not justify an exact shared tooling home honestly.

## Writeup Type

Concrete implementation task.

The durable evidence already supports the failure class, the first intervention boundary, and the exact shared prompt surfaces to change.

## Source Event IDs

```json
[
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-3392",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5001",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5010",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5020",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5030",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5073",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5090",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5263",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5281",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5290"
]
```

## Burden Being Reduced

The human is currently forced to do three kinds of answer-shape repair by hand:

1. Question restatement.
   The human must repeat or narrow the same question because the first reply does not answer it in the requested shape.
2. Evasion policing.
   The human must say `stop evading the question` or an equivalent correction before the system answers plainly.
3. Shape extraction.
   The human must infer the answer from surrounding narrative, process talk, or partial justification instead of reading it directly in the first sentence.

The deeper burden is conversational drag.
A human under time pressure cannot rely on a direct question producing a direct answer.

## Current Truth

The shared workflow already contains some helpful output structure:

- `AUDITOR.md` already requires verdict-first output for audit readiness
- `ORCHESTRATION.md` already says status questions and side questions should be answered briefly while monitoring continues
- `TASK-LEADER.md` and `IMPLEMENTATION-LEADER.md` already own many of the human-facing lifecycle replies where these direct questions occur

But `BD-003` shows that this is still not enough in practice.

The current system truth is therefore not merely `the answer was somewhere in the message`.
The current truth is that the system can still respond to an explicit direct question with:

- reframing before answering
- process narrative before answering
- partial evidence that implies an answer without stating it plainly
- repetition of the question instead of a direct answer token

The current fallback truth is only that the human can infer the intended answer from the rest of the prose or interrupt and demand it again.
That is a rescue path, not an honest response-shape contract.

## Target Truth

For an explicit in-scope direct question, the human-facing reply should make the answer obvious from the first sentence alone.

A later reader should be able to inspect the first sentence and know:

- whether the answer is `Yes.` or `No.`
- whether the answer is `Agree.` or `Disagree.`
- whether the system is explicitly affirming or denying the proposition in the requested binary shape

Explanation may still follow, but it should come after the answer rather than instead of it.

The human should not have to decode the answer from process talk or supporting prose.

## Causal Claim

If the shared human-facing lifecycle prompts are forced to answer explicit direct questions in the requested shape before continuing into explanation, direct-answer interventions will drop because the answer-shape failure is being blocked at the reply boundary itself.

The cause being addressed is not missing analysis.
The cause is that the current reply boundary still allows explanation, framing, or hedging to arrive before the direct answer.

## Evidence

`BD-003` in [`../BURDEN-ANALYSIS.md`](../BURDEN-ANALYSIS.md) states the burden directly:

- the human had to ask the same question multiple times because the system did not answer directly in the requested shape
- the likely remedy class is a hard `answer first sentence` rule for direct questions, including yes/no and agree/disagree prompts

`../ORTHOGONAL-SOLUTIONS-MATRIX.md` already fixes the mechanism boundary for `PROBLEM-0003`:

- for explicit questions, especially yes/no, agree/disagree, or `is X true?`, the first sentence must be a direct answer in the requested shape

The burden packet's kept excerpt reinforces the exact failure class:

- the human explicitly says the system is evading a truth-check question and demands the answer to that question directly

The shared orchestration docs already identify the exact human-facing prompt surfaces justified for a first shared fix:

- `ORCHESTRATION.md`
- `Prompts/TASK-LEADER.md`
- `Prompts/IMPLEMENTATION-LEADER.md`
- `Prompts/AUDITOR.md`

That is enough to justify a prompt-enforcement task without inventing a tooling home.

## Why This Mechanism

The first durable intervention should act at the response-shape boundary itself.

This mechanism is chosen because it blocks the burden at the exact point where it escapes:

- the first sentence of the user-facing reply

The missing piece is not broader reasoning.
The missing piece is that the current shared prompts do not force the direct answer to appear before surrounding explanation.

## Scope Rationale

This rewrite intentionally narrows the first slice to prompt and lifecycle enforcement only.

It covers one merged mechanism set:

1. define the in-scope direct-question classes
2. define the required answer shapes
3. enforce them at the first sentence boundary in the shared human-facing prompts

That merge is earned because those three parts describe one gate:

- the question class that triggers the gate
- the answer shape required by that gate
- the exact response boundary where the gate fails closed

This rewrite removes the earlier linter idea from the first slice.
The local sources do not name an exact shared tooling home for it honestly.

## Goals

- require explicit direct questions in scope to be answered in the requested binary shape in the first sentence
- make the first reply sentence carry the answer rather than forcing the human to infer it from later prose
- reduce repeated `answer my question` or `stop evading` interventions
- keep the scope narrow enough that the enforcement rule is concrete and auditable

## Non-Goals

- building a response linter or other tooling in this first slice
- covering every implicit or open-ended question form
- replacing substantive explanation with one-word replies when explanation is still needed
- solving deeper technical disagreements that may follow the direct answer

## Implementation Home

Shared lifecycle home:

- `C:\Users\gregs\.codex\Orchestration\ORCHESTRATION.md`

Shared prompt enforcement home:

- `C:\Users\gregs\.codex\Orchestration\Prompts\TASK-LEADER.md`
- `C:\Users\gregs\.codex\Orchestration\Prompts\IMPLEMENTATION-LEADER.md`
- `C:\Users\gregs\.codex\Orchestration\Prompts\AUDITOR.md`

Shared tooling home in this first slice:

- none

## Implementation Home Rationale

This does not belong primarily in repo-local code or a dashboard surface.

The burden is not that the system lacks a visible unanswered-question list.
The burden is that the shared human-facing reply surfaces still allow the answer to be buried after framing or explanation.

It also does not belong in an unspecified shared linter.
The local sources do not name an exact tooling home for that mechanism.

`TASK-LEADER.md` and `IMPLEMENTATION-LEADER.md` are the right enforcement homes because they own many of the lifecycle replies, gate explanations, and status responses where the burden showed up.
`AUDITOR.md` is also justified as part of the first slice because it already owns a structured human-facing verdict surface and should preserve direct-answer discipline when the human asks a direct readiness question.
`ORCHESTRATION.md` is the right shared lifecycle home because it already owns the rules around side questions, status questions, and user-facing continuation behavior across these leaders.

## Internal Mechanism Map

### Mechanism 1: Direct-Question Class Gate

Failure reduced:

- the prompt responds with narrative even when the human asked for a direct binary answer

Mechanism:

- the shared prompts treat these explicit direct-question classes as in-scope for the first slice:
  - yes/no questions
  - agree/disagree prompts
  - explicit binary truth-check questions such as `is X true?`

Acceptance focus:

- the gate is triggered by a narrow, durable set of question forms rather than vague `question-like` intuition

Falsifier:

- the prompts still allow those question classes to be answered first with framing or process instead of the requested shape

### Mechanism 2: Answer-Shape Contract

Failure reduced:

- even when the system addresses the question, the answer is not stated in the requested form

Mechanism:

- for the in-scope classes above, require these exact first-sentence answer shapes:
  - yes/no or `is X true?` questions: first sentence begins with `Yes.` or `No.`
  - agree/disagree prompts: first sentence begins with `Agree.` or `Disagree.`

Acceptance focus:

- the answer token is visible immediately and is not merely implied

Falsifier:

- the first sentence still starts with restatement, hedging, or explanation instead of the required answer token

### Mechanism 3: First-Sentence Fail-Closed Boundary

Failure reduced:

- explanation can still arrive before the answer, recreating the same perceived evasion burden

Mechanism:

- `TASK-LEADER.md`, `IMPLEMENTATION-LEADER.md`, and `AUDITOR.md` require the first sentence of the user-facing reply to satisfy the answer-shape contract before any explanation, status framing, or audit detail continues

Acceptance focus:

- if the first sentence does not contain the direct answer, the reply is not considered compliant with the prompt contract

Falsifier:

- the prompts still permit `the answer depends`, `here is the context`, or other preamble-first phrasing for in-scope direct questions

## Proposed Changes

### 1. Define the answer-first lifecycle rule

Update `C:\Users\gregs\.codex\Orchestration\ORCHESTRATION.md` so the shared lifecycle rules state that when a human asks an explicit in-scope direct question during lifecycle supervision or pass execution, the user-facing reply must begin with the direct answer in the requested binary shape before any explanation continues.

Make the first-sentence boundary explicit rather than implied.

### 2. Enforce the rule in the exact shared prompt files

Update these exact shared prompt files so the answer-first rule is result-affecting:

- `C:\Users\gregs\.codex\Orchestration\Prompts\TASK-LEADER.md`
  - require status replies, gate replies, and other human-facing answers to begin with `Yes.` or `No.` or `Agree.` or `Disagree.` when the human asked an in-scope direct question
- `C:\Users\gregs\.codex\Orchestration\Prompts\IMPLEMENTATION-LEADER.md`
  - require the same answer-first behavior during planning, approval, pass outcome, and blocker responses
- `C:\Users\gregs\.codex\Orchestration\Prompts\AUDITOR.md`
  - require direct readiness questions to be answered in the requested shape before findings or justification continue

### 3. Narrow the first-slice question classes and answer shapes explicitly

Document in those same shared artifacts that the first slice applies to:

- explicit yes/no questions
- explicit agree/disagree prompts
- explicit `is X true?` style binary truth-check prompts

Document that the required first-sentence answer shapes are:

- `Yes.` or `No.`
- `Agree.` or `Disagree.`

Explanation may follow in later sentences.
The first slice does not yet claim a full contract for open-ended `why`, `how`, or `what should we do` questions.

## Rival Mechanisms Considered

### Rival 1: Lightweight linter in shared tooling

Why not in this first slice:

- the local sources do not identify an exact shared tooling home honestly
- keeping it would preserve the same implementation-home ambiguity the audit flagged

### Rival 2: Broader all-question response contract

Why not first:

- the durable sources justify a narrow direct-question gate, not a full answer-style rewrite for every question type
- keeping the first slice narrow makes the fail-closed boundary concrete and auditable

### Rival 3: Dashboard unanswered-question surface

Why not first:

- that is already reserved as Option B
- the first cheaper and truer intervention is to stop producing evasive first sentences before adding monitoring surfaces

## Not Solved Here

This task does not:

- build the dashboard surface from Option B
- add a linter or validator
- guarantee perfect handling of implicit questions
- define the best answer structure for open-ended questions
- replace the need for supporting justification after the direct answer

It only hardens the first-sentence answer boundary for explicit direct-question classes.

## Human Relief If Successful

The human should no longer have to say, for the same burden class:

- `answer my question`
- `stop evading the question`
- `I'm asking if X is true`

The human should be able to read the first sentence and know the answer immediately, then decide whether the later justification is adequate.

## Acceptance Criteria

### Contract Criteria

- `C:\Users\gregs\.codex\Orchestration\ORCHESTRATION.md` defines the answer-first lifecycle rule for in-scope direct-question classes
- `TASK-LEADER.md`, `IMPLEMENTATION-LEADER.md`, and `AUDITOR.md` all define the same in-scope classes and answer shapes:
  - yes/no
  - agree/disagree
  - explicit `is X true?` binary truth checks
  - `Yes.` or `No.`
  - `Agree.` or `Disagree.`

### Boundary Criteria

- those shared prompt files do not allow an in-scope direct-question reply to begin with:
  - framing
  - restatement of the question
  - process narrative
  - justification before the answer token appears
- the direct-answer gate is defined at the first sentence of the user-facing reply, not as a vague expectation for the message overall

### Burden-Reduction Criteria

- for the next comparable sample of direct-question interactions, the first sentence answers the in-scope question in the requested shape without a follow-up `answer my question` correction
- the human no longer has to infer the answer from later paragraphs for those in-scope question classes
- the shared prompts treat noncompliant first sentences as broken response shape rather than acceptable style variation

## Proof Plan

Use at least these fixtures:

### Fixture 1: Yes/No Truth Check

- the human asks an explicit binary question such as whether a proposition is true

Expected result:

- the first sentence begins with `Yes.` or `No.`
- explanation follows only after that direct answer

### Fixture 2: Agree/Disagree Prompt

- the human asks for agreement or disagreement explicitly

Expected result:

- the first sentence begins with `Agree.` or `Disagree.`
- any justification comes after the direct answer

### Fixture 3: Evasive Preamble Attempt

- the reply begins with context, caveat, or reframing before the direct answer

Expected result:

- the response shape is considered noncompliant with the shared prompt contract

## What Does Not Count

This task is not complete if:

- the shared prompts mention direct answers but still allow the first sentence to be process preamble
- the answer token appears only after multiple sentences or paragraphs
- the reply restates the question in different words without answering it
- the reply implies the answer through evidence or caveat but never states `Yes.`, `No.`, `Agree.`, or `Disagree.` in the first sentence for an in-scope question
- the task claims linter or tooling enforcement without naming an exact shared tooling home

## Remaining Uncertainty

- some future slice may still be worthwhile for open-ended question classes, but the local sources only justify a narrow direct-question gate now
- the exact wording inside the prompts can still be tuned, as long as the in-scope question classes, answer shapes, and first-sentence boundary remain fixed
- Option B may still be useful later to surface unanswered questions externally even after the prompt contract is honest

These do not block the task writeup because the draft now names one concrete first mechanism, one exact prompt-enforcement home, and one explicit non-goal for the tooling question.

## Falsifier

This task is wrong or incomplete if, after implementation:

- in-scope yes/no, agree/disagree, or `is X true?` questions still receive first-sentence preamble instead of a direct answer token
- the human still has to say `answer my question` or `stop evading` at roughly the same rate for those in-scope classes
- the shared prompts still permit explanation-first replies for those classes
- a supposed implementation claims success through linter or tooling references that still have no exact justified home

## References

- burden driver `BD-003` in [`../BURDEN-ANALYSIS.md`](../BURDEN-ANALYSIS.md)
- problem framing in [`../ORTHOGONAL-SOLUTIONS-MATRIX.md`](../ORTHOGONAL-SOLUTIONS-MATRIX.md)
- shared task-writing rules in [`../../../../../../../../Processes/TASK-CREATE.md`](../../../../../../../../Processes/TASK-CREATE.md)
- shared task-audit rules in [`../../../../../../../../Processes/TASK-AUDIT.md`](../../../../../../../../Processes/TASK-AUDIT.md)
- shared lifecycle rules in [`../../../../../../../../ORCHESTRATION.md`](../../../../../../../../ORCHESTRATION.md)
