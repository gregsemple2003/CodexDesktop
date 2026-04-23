# PROBLEM-0003 OPTION-A: Shared Direct-Answer-First Rule (Consensus On Carve-Outs)

## Title

Decide and write a shared "direct answer first" interaction contract so short-answer questions receive a first-sentence answer before framing.

## Summary

The packet shows repeated direct-answer failures: the human asks a short, concrete question (often yes/no or "what is still wrong?") and the system responds with framing or analysis without answering. This forces the human to stop the system and re-ask, increasing time cost and irritation.

The frozen winner is a shared rule, not a ThirdPerson-only patch: the burden is conversational and cross-repo.

However, winner synthesis explicitly notes that the *exact carve-outs* still require a bounded decision artifact before an enqueue-ready implementation can be considered complete.

## Writeup Type

Consensus task (burden-reduction proposal).

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

The human is currently forced to do repeated conversational repair work:

- stopping the system mid-response to demand the actual answer first
- asking the same question multiple times because the first response did not answer it
- clarifying the question shape ("use small words", "yes/no", "stop evading") to force a direct response

This is not merely a tone problem. It is a throughput problem: it delays technical work because the human must first fight for an answer.

## Current Truth

The shared `.codex` system defines many structural workflow norms, but it does not currently contain a crisp, enforceable "answer-first" contract in its primary shared interaction home:

- [../../../../../../../../../AGENTS.md](../../../../../../../../../AGENTS.md)

So when the human asks a direct question, the system has no durable shared rule that says:

- answer in the first sentence, then (optionally) add framing

## Target Truth

The shared system should have one durable, reviewable interaction rule such that:

- when the human asks a short-answer question (yes/no, agree/disagree, "what is still wrong", "define this term"), the response begins with a direct first-sentence answer
- if the answer is uncertain, the first sentence is direct uncertainty (for example: "I don’t know yet; I need X to answer.")
- framing, caveats, and supporting context are allowed only after the first-sentence answer

The contract must also include bounded carve-outs so it is not weaponized into unsafe or misleading behavior (for example: refusing to answer when the question is ill-posed, or requiring clarification when multiple mutually exclusive interpretations exist).

## Causal Claim

If the shared contract explicitly requires "direct answer first" for specific question classes and names explicit carve-outs, then:

- future interactions will require fewer STOP/re-ask interventions
- the human will spend less time coercing shape and more time on the actual work

## Evidence

Evidence is concentrated in [../BURDEN-ANALYSIS.md](../BURDEN-ANALYSIS.md) (`BD-003`) and the stable event previews in [../../HumanInputEvents/INDEX.json](../../HumanInputEvents/INDEX.json), including:

- explicit STOP because a question was not answered (`...-3392`)
- repeated "use small words" and "answer my question" prompts (`...-5001` through `...-5090`)
- repeated "what is still wrong" / "what does that mean here" clarification demands (`...-5263`, `...-5281`, `...-5290`)

## Why This Mechanism

This proposal chooses a shared, durable contract because:

- the failure happens before repo content matters
- it appears across multiple interaction contexts (debugging, review, planning)
- a repo-local patch would incorrectly localize a generic burden

The narrowest honest shared home for an interaction contract is the shared `.codex` front door:

- `AGENTS.md`

## Scope Rationale

Shared scope is earned here:

- The burden is conversational, not ThirdPerson-specific.
- A single repo-local rule would not help the human in other repos or in `.codex` maintenance flows.

## Goals

- Produce a clear, durable answer-first contract that is specific enough to be enforced and audited.
- Decide the carve-outs so the rule does not create new failure modes (false certainty, unsafe compliance, or answer-without-clarification).
- Reduce repeated "STOP / answer my question" human interventions.

## Non-Goals

- Changing repo-local process docs as a proxy; the home is shared.
- Adding tooling or transcript scoring in this first step.
- Forcing "one-sentence answers always"; the rule is about *first sentence directness*, not total brevity.

## Implementation Home

- Shared interaction contract home:
  - [../../../../../../../../../AGENTS.md](../../../../../../../../../AGENTS.md)

## Implementation Home Rationale

`AGENTS.md` is already the shared front door for cross-repo norms about behavior, artifact structure, and workflow precedence. It is the narrowest durable shared home that can own an interaction rule without dragging it into repo-local docs or prompt bundles.

## Proposed Changes

These are the concrete, reviewable surfaces this consensus task changes.

1. **Add a new shared interaction rule section**
   - File: [../../../../../../../../../AGENTS.md](../../../../../../../../../AGENTS.md)
   - Add a new section, for example `## Direct Answer First`, containing:
     - a short statement of the rule ("first sentence answers the question directly")
     - a definition of covered question classes (yes/no, agree/disagree, short status, "what is still wrong", "define this term")
     - explicit carve-outs that require clarification instead of forced yes/no
     - explicit instruction for uncertainty ("answer uncertainty directly in the first sentence")
2. **Add 3-5 short examples**
   - Same file/section
   - Include paired examples of:
     - bad: framing-first evasion
     - good: direct first sentence + optional framing

## Acceptance Criteria

- The shared `AGENTS.md` contains a clearly labeled "Direct Answer First" rule with:
  - defined covered question classes
  - defined carve-outs/clarification triggers
  - an explicit uncertainty rule
  - examples that make misinterpretation harder
- A cold reader can apply the contract without additional oral tradition.

## Expected Resolution

Human-facing outcome:

- When the human asks a short-answer question, the response starts with the answer instead of forcing a STOP/re-ask loop.
- The human experiences fewer "shape fights" during technical work.

## Human Relief If Successful

- Less time spent coercing response shape.
- Lower frustration cost during debugging and review.
- Faster iteration when a single concrete answer would unblock the next step.

## Internal Mechanism Map

1. Decide the rule boundary (what question types trigger answer-first).
2. Decide carve-outs (when the system must clarify instead of answering).
3. Write the durable contract in shared `AGENTS.md`.
4. Use examples to reduce drift and to make auditing feasible later.

## Rival Explanations Considered

- "This was just frustration; answer shape doesn’t matter."
  - Rejected: the packet contains multiple separate events where lack of direct answering directly delayed work.
- "This is ThirdPerson-specific because of its high-friction debugging context."
  - Rejected: the burden is generic interaction behavior; nothing in the evidence ties it uniquely to ThirdPerson content.

## Rival Mechanisms Considered

- `P-003 / Option B` (repo-local direct-answer clause in ThirdPerson docs):
  - Rejected as winner: would incorrectly localize a shared conversational burden and would not help outside ThirdPerson.
- Post-hoc transcript auditing/scoring:
  - Rejected for first slice: detects failure after the human already paid the cost; does not prevent the first evasion.

## Tradeoffs

- Risk of oversimplification:
  - a rigid answer-first rule can create false certainty if carve-outs are missing or unclear
- Risk of adversarial framing:
  - a human can ask an ill-posed yes/no; the rule must explicitly allow clarification when multiple interpretations exist

## Shared Substrate

- Shared doc precedence and glossary: [../../../../../../../../../AGENTS.md](../../../../../../../../../AGENTS.md)
- Shared task-writing standards for later implementation tasks that might enforce this rule more mechanically: [../../../../../../../../Processes/TASK-CREATE.md](../../../../../../../../Processes/TASK-CREATE.md)

## Not Solved Here

- Implementing enforcement in prompts, UIs, or validators.
- Measuring compliance automatically.
- Repo-local debugging/proof gates (those are separate winners).

## What Does Not Count

- Adding vague prose like "be concise" without an explicit first-sentence answer rule.
- Adding a rule with no carve-outs, such that it forces unsafe or misleading yes/no answers.
- Treating "a direct answer somewhere in the paragraph" as compliant; the goal is first-sentence directness for covered question classes.

## Remaining Uncertainty

- The minimal carve-out set that prevents bad behavior without weakening the contract into mush.
- Whether some question classes (for example "why") should be explicitly excluded or handled as "answer the direct part first, then explain."

## Falsifier

This proposal is falsified if, after the shared contract is adopted:

- the human still frequently has to repeat the question because the first sentence does not answer it directly, and the situation is not explained by an explicit carve-out/clarification trigger

## Proof Plan

1. Draft the `AGENTS.md` section text and examples in a single diff.
2. Have a reviewer attempt to apply the rule to a handful of transcripts (including the packet excerpts) and identify ambiguous cases.
3. Revise carve-outs until:
   - the rule prevents framing-first evasion for covered question classes
   - the carve-outs protect against forced false certainty

## Open Questions

- Should "agree/disagree" be treated as a strict first-token requirement ("Agree." / "Disagree.") or only a first-sentence requirement?
- Should the contract explicitly require an "answer token" (`Yes.` / `No.` / `I don’t know yet.`) for certain question classes to reduce drift further?

## Decision To Make

Freeze the exact shared "direct answer first" contract, including:

- covered question classes
- carve-outs / clarification triggers
- uncertainty handling
- example set

## Inputs

- Evidence: [../BURDEN-ANALYSIS.md](../BURDEN-ANALYSIS.md) (`BD-003`) and the listed `Source Event IDs`
- Current shared interaction contract home: [../../../../../../../../../AGENTS.md](../../../../../../../../../AGENTS.md)
- Winner boundary: [../WINNER-SYNTHESIS.md](../WINNER-SYNTHESIS.md#w-003-shared-direct-answer-first-rule)

## Options To Compare

These are *within-winner* options to decide carve-outs and strictness, not re-choosing the winner.

1. **Strict token-first for yes/no classes**
   - first token must be `Yes.` / `No.` / `Agree.` / `Disagree.` / `Unknown.` (or equivalent), then framing
2. **First-sentence directness only**
   - first sentence must directly answer, but token may vary
3. **Hybrid**
   - strict token-first for yes/no + agree/disagree
   - first-sentence directness for "what is still wrong" and "define this"

## Decision Output

One durable artifact that freezes the contract:

- a committed edit to [../../../../../../../../../AGENTS.md](../../../../../../../../../AGENTS.md) adding a new `## Direct Answer First` section with the chosen option and carve-outs

## References

- Burden driver: [../BURDEN-ANALYSIS.md](../BURDEN-ANALYSIS.md) (`BD-003`)
- Designed options: [../SOLUTION-DESIGN.md](../SOLUTION-DESIGN.md#p-003-direct-answer-first-contract)
- Frozen winner boundary: [../WINNER-SYNTHESIS.md](../WINNER-SYNTHESIS.md#w-003-shared-direct-answer-first-rule)
- Final matrix row: [../ORTHOGONAL-SOLUTIONS-MATRIX.md](../ORTHOGONAL-SOLUTIONS-MATRIX.md#p-003-direct-answer-first-contract)

