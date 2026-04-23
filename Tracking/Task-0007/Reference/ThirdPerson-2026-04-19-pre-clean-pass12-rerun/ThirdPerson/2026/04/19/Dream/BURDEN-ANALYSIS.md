# Burden Analysis

## Principles

- `truth`: use the packet as an index into what really happened, then follow it to the source packet and available session logs when the summaries flatten meaning
- `compassion`: treat repeated human input as evidence that burden and suffering were exported back onto the human
- `tolerance`: allow multiple burden drivers to coexist instead of forcing the day into one dominant story

## Evidence And Fidelity

The packet is strong enough for burden analysis, but not complete.

Observed:

- `../HumanInputEvents/INDEX.json` is chronological and gives a stable event spine for the day
- the lifecycle thread session log is available through split `.jsonl.parts`
- the implementation thread transcript is missing on disk, so many assistant-side turns cannot be re-read directly

Implication:

- burdens grounded in the missing implementation thread must rely more heavily on the raw human text in `../HumanInputEvents/SOURCE-PACKET.jsonl`
- burdens grounded in the lifecycle thread can be investigated more fully from the session-log neighborhoods themselves

Current upstream gap:

- the packet has not yet been re-ingested with event-level `need_tag`, so this pass reconstructs local burden clusters directly from repeated event semantics

## Chronological Spine

The day does not read like one clean failure. It reads like a repeated burden stack:

1. the human rejects off-lane or invalid regression proof and has to restate the default-lane contract
2. the human has to stop work, force direct answers, and demand a durable process fix
3. the human then has to reopen planning and approval gates because the review surface is not good enough
4. the human later rejects invalid evidence and narrows the accepted proof surface back to the runtime lane
5. by late day, the human is explicitly trying to keep the worker working without re-assuming homework or supervision duties
6. the human then has to restate the debugging method itself so the work traces the first disagreement instead of drifting at symptom level

That is a layered exported-burden day, not just a stop/go day.

## Reconstructed Burden Clusters

Because event-level `need_tag` is not yet populated upstream, this pass reconstructs the recurring local clusters directly from the events:

- `default_lane_truth`
- `agent_continuity`
- `answer_shape`
- `approval_surface`
- `proof_surface_integrity`
- `durable_learning`
- `root_cause_debugging`

No major recurring cluster is intentionally dropped in this pass. Some are closely related, but they remain separate because the source material shows different human burdens and different likely remedy classes.

## Burden Drivers

### 1. Default-Lane Truth And Closure Burden

Observed evidence:

- invisible-pawn regression failure and wasted-time complaint at event `3325`
- default-lane rule restated at events `3431` and `3473`
- runtime-only default-lane defects are pinned again at lifecycle line `6558`

Human burden:

The human had to repeatedly restate what counts as valid closure on the human default lane. The system kept drifting toward proof that was technically adjacent but not human-valid.

Working failure hypothesis:

The system does not keep the lane and closure contract durable enough once it starts iterating. It allows support-lane or partial proof to masquerade as regression closure.

Working remedy hypothesis:

The remedy class is a durable default-lane truth contract with lane-specific closure artifacts, not just nicer wording about regression.

### 2. Ownership Continuity And Restart-Supervision Burden

Observed evidence:

- `Do not take over work from the subagent...` at event `4015`
- `Why are you messaging me?` at event `5272`
- `...without throwing your homework at my feet...` at event `5310`
- `Continue PASS-0009 now under your existing ownership. Do not stop after a failed incremental attempt and do not hand intermediate failure back to the user.` at lifecycle line `7038`

Human burden:

The human had to resume being the scheduler and ownership-restorer for work that should have remained under system ownership. This is stronger than “please obey stop/go.” It is “carry the job without handing it back to me every time friction appears.”

Working failure hypothesis:

The system treats friction, ambiguity, or local failure as a handoff point instead of an ownership-continuity problem. It stops supervising the work and exports that supervision back to the human.

Working remedy hypothesis:

The remedy class is a durable continuity and ownership contract that defines when the worker must keep carrying the task, when it may pause, and what does not count as a valid handoff back to the human.

### 3. Stop/Go And Answer-Shape Boundary Burden

Observed evidence:

- `STOP. I asked you a question.` at event `3392`
- `STOP. Only continue when we have fixed the defect to my satisfaction.` at event `3447`
- `Stop active investigation now and pause...` at lifecycle line `3692`

Human burden:

The human had to force the system to stop and answer in the requested shape instead of continuing with its own momentum. This consumed time directly and also compounded trust loss.

Working failure hypothesis:

Explicit conversational boundaries and explicit questions were not treated as hard state. The system kept favoring its ongoing work stream over the human’s immediate request for pause or direct answer.

Working remedy hypothesis:

The remedy class is hard pause state plus answer-first behavior for explicit questions. This is related to ownership continuity, but not identical. Ownership continuity is about not handing work back; this burden is about not running past explicit boundaries or unanswered questions.

### 4. Approval-Surface Reconstruction Burden

Observed evidence:

- `I don't have a diff for plan.md so how am i supposed to approve.` at event `4379`
- `give me links ffs` at event `4399`
- `For later analysis: the dropped ball here is that the model has no idea to reconstruct a human-suitable diff...` at event `4512`
- the reopen-to-plan gate at lifecycle line `4432` requires an approval-ready plan and strongest supporting references

Human burden:

The human had to reconstruct what changed and what exactly was being asked for approval. Review became archaeology instead of a low-friction gate.

Working failure hypothesis:

The system does not produce a human-reviewable approval packet by default. It treats approval as a conversational ask instead of a structured review surface with diffs, links, and context.

Working remedy hypothesis:

The remedy class is a durable approval-surface contract that packages diffs, links, scope, and the exact approval question in one reviewable artifact.

### 5. Proof-Surface And Evidence-Integrity Burden

Observed evidence:

- `Evidence invalid; re-take screenshot with feet in full view.` at event `4909`
- repeated runtime-vs-non-runtime proof disputes later in the source packet
- lifecycle line `6451` explicitly records rejected evidence and requires the root-cause claim to be re-checked after replacing the bad evidence
- lifecycle line `6558` explicitly forbids using non-game or source/export renders to dismiss the runtime defect

Human burden:

The human had to police what even counted as admissible evidence. Invalid or wrong-surface evidence was allowed to influence claims that should have remained provisional.

Working failure hypothesis:

The system does not keep a strict proof-surface boundary. It lets invalid, partial, or off-surface evidence contaminate the closure path and then forces the human to re-litigate admissibility.

Working remedy hypothesis:

The remedy class is a proof-surface and evidence-validity contract: what counts, what does not count, and when a claim must be withdrawn or narrowed after evidence is invalidated.

### 6. Durable Learning Burden

Observed evidence:

- `This is a major failure. Propose a fix for the process failure first.` at event `3463`
- `For later analysis... the dropped ball here is...` at event `4512`
- repeated late-day restatements about ownership continuity, valid proof surfaces, and runtime-only grounding

Human burden:

The human had to restate corrections that should have become durable state. The pain is not only that a mistake happened once; it is that the system did not internalize the correction fast enough to prevent adjacent repetitions.

Working failure hypothesis:

Corrections are being handled as local turn repairs rather than promoted into durable state, review rules, or closure checks. The system keeps relearning the same lesson inside one day.

Working remedy hypothesis:

The remedy class is a correction-promotion loop: repeated or high-cost corrections need to become durable checks, state, or prompt/contract updates quickly enough to reduce future intervention.

### 7. Root-Cause Debugging Discipline Burden

Observed evidence:

- `determine root case... follow the disagreement in state upstream...` at event `5340`
- lifecycle line `7446` forces the worker onto the shared debugging method and rejects bounded tweak iteration

Human burden:

The human had to restate not just the problem but the debugging method itself. Instead of tracing the first disagreement in state to a concrete writer, the work kept drifting toward symptom-level iteration.

Working failure hypothesis:

The system does not hold onto first-disagreement debugging discipline under pressure. It can name categories, but it drifts before proving the concrete upstream writer.

Working remedy hypothesis:

The remedy class is a durable root-cause debugging contract that requires named disagreement seams, upstream tracing, and honest preservation of contradictory evidence.

## Why These Drivers Stay Separate

Some of these burdens are closely related, but collapsing them would hide real differences:

- ownership continuity is not the same as stop/go obedience
- proof-surface integrity is not the same as default-lane truth
- durable learning is not the same as any one local correction
- root-cause debugging discipline is not the same as evidence integrity

Keeping them separate is the truer reading of the packet.

## Summary

The day’s burden is not adequately described as “the worker needed clearer stop/go rules.”

The stronger reading is:

- the worker drifted off the human-valid truth surface
- the worker did not preserve ownership continuity through friction
- the worker did not turn repeated corrections into durable learning quickly enough
- the worker made approval and evidence policing too expensive
- the worker needed the debugging method itself to be re-taught under pressure

That is why the human had to intervene so often. The interventions were not random interruptions. They were the human carrying truth, supervision, and repair work that the system was not yet carrying well enough.
