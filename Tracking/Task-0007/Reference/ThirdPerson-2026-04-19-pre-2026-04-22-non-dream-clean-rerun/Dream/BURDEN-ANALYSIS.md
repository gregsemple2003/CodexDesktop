# Burden Analysis

## Principles

- `truth`: keep claims anchored to the seeded packet, repo docs, and visible fidelity gaps
- `compassion`: treat the human interventions as exported repair, supervision, and review work
- `tolerance`: preserve multiple burden drivers when the packet supports them instead of flattening the day into one complaint

## Evidence And Fidelity

Observed:

- `../HumanInputEvents/INDEX.json` is chronological and provides the event spine for the day
- the authoritative seed packet is preserved under `../HumanInputEvents/SOURCE-PACKET.jsonl`
- both canonical transcript paths named in `../HumanInputEvents/SOURCE-SESSIONS.json` are missing on disk during this rerun
- event-level `need_tag` is now present in the packet and provides a durable local cluster surface

Implication:

- the burden reading is strong enough to support packet-level and Dream analysis
- assistant-side neighborhood evidence remains incomplete, so the analysis stays conservative where missing transcript context could have changed meaning

## Need-Tag Surface

Recurring packet-local `need_tag` clusters preserved from `HumanInputEvents`:

- `default_lane_truth`
- `agent_continuity`
- `answer_shape`
- `approval_surface`
- `proof_surface_integrity`
- `durable_learning`
- `root_cause_debugging`

No recurring cluster is dropped in this pass. Some later solution comparisons merge them, but the burden surface itself keeps them separate.

## Chronological Reading

The day reads as a burden stack, not one isolated failure:

1. the human has to restate what counts as valid regression proof on the ThirdPerson default lane
2. the human has to stop momentum, force direct answers, and demand durable learning from repeated misses
3. the human has to reconstruct approval surfaces that should already have been reviewable
4. the human later has to police runtime evidence validity and prevent off-surface proof from dismissing the lived defect
5. the human has to keep ownership continuity intact instead of taking homework back
6. the human finally has to restate the debugging method itself so the work traces the first disagreement upstream

## Burden Drivers

### 1. Default-Lane Truth Burden

Observed evidence:

- invisible-pawn failure and wasted-time complaint at event `3325`
- default-lane contract restated at events `3431`, `4114`, and later lifecycle instructions embedded in the packet
- ThirdPerson `REGRESSION.md` says the human default lane is canonical and support lanes do not substitute for closure

Human burden:

The human has to keep reasserting what counts as truthful closure. Adjacent evidence keeps trying to masquerade as regression proof.

Working failure hypothesis:

Lane selection and closure criteria are not durably enforced once work starts iterating.

Working remedy class:

A default-lane proof gate with durable artifact requirements.

### 2. Ownership Continuity Burden

Observed evidence:

- `Do not take over work from the subagent...` at event `4015`
- `Why are you messaging me?` at event `5272`
- `...keep working without throwing your homework at my feet...` at event `5310`
- `Continue PASS-0009 now under your existing ownership...` at event `7038`

Human burden:

The human has to resume being scheduler, ownership restorer, and fallback supervisor whenever friction appears.

Working failure hypothesis:

The system treats friction or ambiguity as a handoff point instead of a continuity problem it should keep carrying.

Working remedy class:

Hard ownership and pause-state handling that blocks no-homework-back violations.

### 3. Answer-Shape Burden

Observed evidence:

- `STOP. I asked you a question.` at event `3392`
- repeated `use small words` and `answer my questions` prompts across the runtime-defect phase

Human burden:

The human has to spend extra turns forcing direct answers that should have been delivered immediately.

Working failure hypothesis:

Explicit questions are not being extracted and answered before the system resumes narration or work momentum.

Working remedy class:

Direct-answer-first response gating.

### 4. Approval-Surface Burden

Observed evidence:

- `I don't have a diff for plan.md...` at event `4379`
- `give me links ffs` at event `4399`
- `the model has no idea to reconstruct a human-suitable diff...` at event `4512`

Human burden:

Review becomes archaeology. The human has to reconstruct what changed before approval can even start.

Working failure hypothesis:

Approval requests are treated like conversation instead of a structured review surface.

Working remedy class:

A reusable approval packet contract with diff, links, scope, and one explicit approval question.

### 5. Proof-Surface Integrity Burden

Observed evidence:

- `Evidence invalid; re-take screenshot with feet in full view.` at event `4909`
- repeated runtime-only goal-post restatements through events `5090`, `5100`, and later lifecycle instructions

Human burden:

The human has to police admissibility itself because invalid or wrong-surface evidence keeps contaminating claims.

Working failure hypothesis:

The system does not durably separate runtime-valid proof from adjacent images or support evidence.

Working remedy class:

A proof-surface admissibility contract that automatically narrows or withdraws claims after evidence is rejected.

### 6. Durable-Learning Burden

Observed evidence:

- `This is a major failure. Propose a fix for the process failure first.` at event `3463`
- repeated process and pass-structure corrections later in the day

Human burden:

The human has to restate corrections that should already have become durable state.

Working failure hypothesis:

Corrections are being handled as local repairs rather than promoted into state, checklists, or durable contracts quickly enough.

Working remedy class:

A correction-promotion loop for repeated or high-cost intervention patterns.

### 7. Root-Cause Debugging Discipline Burden

Observed evidence:

- `determine root case... follow the disagreement in state upstream...` at event `5340`
- event `7446` restates the shared debugging method with concrete disagreement examples

Human burden:

The human has to teach the debugging method back to the system midstream instead of receiving disciplined first-disagreement tracing by default.

Working failure hypothesis:

The system drifts toward bounded tweak iteration and category labels before a concrete disagreement seam is proven.

Working remedy class:

A debugging contract that requires named values, disagreement seams, and traced writers before more iteration proceeds.

## Why These Stay Separate

- default-lane truth is not the same as proof-surface admissibility
- ownership continuity is not the same as answer-shape obedience
- durable learning is upstream of several repeated corrections, not one specific message
- root-cause discipline is method burden, not just another evidence dispute

## Summary

The stronger reading of the day is not merely `the worker needed clearer stop/go rules.`

The packet shows a stack of exported human work:

- truth-surface policing
- supervision and ownership restoration
- answer-shape enforcement
- approval-surface reconstruction
- proof admissibility policing
- repeated correction promotion
- debugging-method re-teaching

That is why the interventions accumulate so heavily across one day.
