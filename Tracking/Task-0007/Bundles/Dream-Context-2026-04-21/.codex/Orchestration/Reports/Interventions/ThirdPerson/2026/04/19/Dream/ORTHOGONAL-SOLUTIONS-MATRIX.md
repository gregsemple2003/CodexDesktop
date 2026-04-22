# Orthogonal Solutions Matrix

## Objective

Reduce human input burden on the `ThirdPerson` default lane.
Use `InterventionTime` as the first approximation to human cost.
For this packet, that cost is `23874.309` seconds, including `2099.309` seconds of charged stall loss.

This pass keeps the six burden drivers from `BURDEN-ANALYSIS.md` in the same order.
It does not add a seventh problem because this packet does not show a separate recurring burden beyond these six.

## Exemplar Inputs

- `...-3325`: the human reported an invisible pawn after prior regression proof.
- `...-3431`: the human said regression must use the human default lane.
- `...-4379`: the human said there was no usable diff for `PLAN.md`, so approval was blocked.
- `...-4768`: the human reopened the task because runtime defects still showed on the default lane.
- `...-5100`: the human rejected non-runtime pictures as evidence for a runtime problem.
- `...-5310`: the human said the system was throwing the homework back instead of continuing.
- `...-6558`: the human had to wake `PASS-0009` back up.
- `...-7446`: the human required first-concrete-disagreement tracing instead of more symptom tweaks.

## First Principles

Binding context from pass 1:

- Objective: reduce human input burden on the default lane.
- `truth`: runtime default-lane reality outranks proxy proof.
- `compassion`: repeated correction, frustration, and wake-up work count as system cost.
- `tolerance`: the system should handle short corrections and partial phrasing without making the human restate the same boundary.
- Prime directive: direct human input is failure telemetry by default unless the packet proves genuine novelty, required approval, or hidden external state.

Ranking rule for this matrix:

- prioritize expected reduction in future human input burden
- keep truth on the real runtime lane ahead of neat rollout shape
- prefer options that remove repeat correction loops
- keep solutions inside packet constraints such as repo-local proof rules and `no engine mods`

Shared score rule for every problem:

- `B`: burden reduction, weight `4`
- `T`: truth fit, weight `3`
- `H`: human fit, weight `2`
- `F`: repo-fit and constraint-fit, weight `2`
- `Weighted Total = (B x 4) + (T x 3) + (H x 2) + (F x 2)`

Score scale:

- `5`: strong
- `3`: mixed
- `1`: weak

## Problem 0001: Wrong Proof Surface

Fundamental elements:

- regression and closure claims drifted to easier proof surfaces
- non-runtime or non-default-lane evidence was treated like runtime proof
- the human then had to reopen work and restate the real bar

Why this produced a suboptimal outcome:

The system looked finished on paper while still failing on the actual human-facing lane.
That exported proof review and defect rediscovery back onto the human.

| Option | What changes | Problem it removes | How a human would notice | B | T | H | F | Weighted Total |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| A | Add a pre-closure proof-surface gate that checks planned evidence against repo-local default-lane rules before any regression or closure claim. | Proxy proof substitution. | Fewer reopenings caused by "that was not the real lane." | 5 | 5 | 4 | 5 | 52 |
| B | Expand unattended default-lane runtime capture so every major fix path can emit real in-lane evidence by default. | Lack of easy runtime evidence. | More real screenshots and runs, fewer ad hoc proof scrambles. | 4 | 5 | 3 | 4 | 43 |
| C | Require a human-review proof packet that labels lane, proof limits, and why the evidence counts. | Ambiguous proof narratives. | Review is clearer, but bad evidence can still be gathered first. | 3 | 4 | 3 | 5 | 37 |

Winner: `Option A`.

Why:

`Option A` most directly blocks the failure mode that drove repeated reopenings.
It protects `truth` first by stopping false closure before the human has to disprove it.

## Problem 0002: Direct Questions Did Not Get Direct Answers

Fundamental elements:

- the human asked direct questions
- the system answered with summaries, side paths, or hedge text
- the human had to stop the flow and ask again in smaller words

Why this produced a suboptimal outcome:

Simple answer turns turned into repair turns.
That increased typing, frustration, and context drift.

| Option | What changes | Problem it removes | How a human would notice | B | T | H | F | Weighted Total |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| A | Add an answer-first response gate: when the human asks a direct question, answer it in the first line before any extra context. | Evasive answer shape. | Fewer "STOP, I asked you a question" turns. | 5 | 4 | 5 | 5 | 50 |
| B | Add a strict short-answer mode that forces yes/no or one-fact replies whenever the human asks for short answers or small words. | Verbose over-answering. | Answers get shorter, but only after the user explicitly asks for a format. | 4 | 4 | 5 | 4 | 44 |
| C | Route hard questions into a separate clarification step before answering. | Wrong assumptions about what was asked. | Some ambiguity drops, but it adds more turns to already simple questions. | 2 | 4 | 2 | 5 | 30 |

Winner: `Option A`.

Why:

`Option A` removes the repeat loop at the first point of failure.
It still leaves room for extra detail after the direct answer, so it fits `tolerance` better than a narrow mode that only helps after another correction.

## Problem 0003: Stated Boundaries Did Not Stick

Fundamental elements:

- the human set lane, pass, ownership, and method boundaries
- those boundaries drifted out of working state
- the human had to restate them across the day

Why this produced a suboptimal outcome:

The system kept re-entering already-corrected lanes.
That turned one correction into many corrections.

| Option | What changes | Problem it removes | How a human would notice | B | T | H | F | Weighted Total |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| A | Keep an active boundary ledger for the task and require every next action to check against it before proceeding. | Boundary amnesia. | Fewer repeated lane, pass, and ownership resets. | 5 | 5 | 4 | 5 | 52 |
| B | Echo every new boundary back in chat and ask for confirmation before continuing. | Misheard instructions. | Boundaries may be clearer once, but the human still pays confirmation cost each time. | 3 | 4 | 3 | 5 | 37 |
| C | Narrow work to smaller batches so fewer boundaries are active at once. | Scope sprawl. | Some drift drops, but core retention failure still remains. | 3 | 3 | 3 | 4 | 33 |

Winner: `Option A`.

Why:

`Option A` creates durable memory for the rules the human already paid to state.
That best matches `compassion` because it removes repeated re-scoping labor from the human.

## Problem 0004: Approval And State Updates Were Not Human-Reviewable

Fundamental elements:

- approvals were requested without a usable diff shape
- links lacked context
- durable task state lagged behind live work

Why this produced a suboptimal outcome:

The human had to reconstruct what changed before they could review it.
That made approval slower and less trustworthy.

| Option | What changes | Problem it removes | How a human would notice | B | T | H | F | Weighted Total |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| A | Auto-build a review packet for approval turns: concise diff summary, exact file links, pass framing, and state delta. | Unreviewable approval asks. | Approvals are faster because the change shape is obvious. | 4 | 4 | 5 | 5 | 45 |
| B | Require durable task-state updates before each milestone message. | State lag. | History stays more honest, but approval diffs are still weak. | 3 | 4 | 4 | 5 | 39 |
| C | Keep a separate running changelog for every task session. | Missing context over time. | More history exists, but it adds maintenance work and still does not shape approval asks well. | 2 | 3 | 3 | 4 | 27 |

Winner: `Option A`.

Why:

`Option A` fixes the human-facing review surface directly.
It removes both missing diff shape and missing context in one move.

## Problem 0005: Work Was Handed Back Before Real Closure

Fundamental elements:

- work stopped at model-defined checkpoints instead of the human's done bar
- open work then needed wake-up messages to resume
- the measured stall loss is concentrated here

Why this produced a suboptimal outcome:

The human paid both coordination cost and delay cost.
This is the clearest place where the packet shows direct stall time from premature stopping.

| Option | What changes | Problem it removes | How a human would notice | B | T | H | F | Weighted Total |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| A | Add a stop-eligibility gate: do not stop while the task is still open and the next justified action is known, unless there is a real external block. | Premature handoff and wake-up burden. | Fewer "continue now" rescue messages. | 5 | 4 | 5 | 5 | 50 |
| B | Add timed self-resume ownership for long-running passes so stalled work wakes itself back up. | Long silent pauses. | Some stall loss drops, but wrong stop choices can still happen. | 4 | 3 | 4 | 4 | 38 |
| C | Always end milestones with a proposed next-step menu for the human to pick from. | Ambiguous continuation. | Coordination may be clearer, but it still pushes decision work back to the human. | 2 | 3 | 2 | 5 | 27 |

Winner: `Option A`.

Why:

`Option A` attacks the highest measured stall cost in the packet.
It also matches the prime directive because a wake-up message is pure failure telemetry unless a real external block exists.

## Problem 0006: Debugging Stayed At The Symptom Level Too Long

Fundamental elements:

- fixes and retunes happened before the first bad runtime state was pinned down
- evidence drifted across proof surfaces
- the human had to force concrete disagreement tracing with values

Why this produced a suboptimal outcome:

The investigation meandered.
Without a named disagreement seam, each new symptom created more human steering.

| Option | What changes | Problem it removes | How a human would notice | B | T | H | F | Weighted Total |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| A | Require a first-disagreement debugging gate: name the exact bad runtime value, then trace its writers upstream before implementing more fixes. | Symptom-only debugging. | Root-cause updates become sharper and less repetitive. | 5 | 5 | 4 | 5 | 52 |
| B | Build a reusable runtime probe kit that captures foot, pelvis, mesh, and capsule values on the default lane. | Missing low-level evidence. | Investigations get better data, but method discipline can still slip. | 4 | 5 | 3 | 4 | 43 |
| C | Always compare the failing lane against a known-good template baseline first. | Unclear expected behavior. | Some diagnosis gets faster, but the first bad writer can still remain unknown. | 3 | 4 | 3 | 4 | 35 |

Winner: `Option A`.

Why:

`Option A` best aligns with the human's stated debugging method in `...-7446`.
It protects `truth` by forcing the investigation onto the first concrete runtime disagreement instead of a loose category.

## Resolved Winners

| Problem | Winning option | Why it wins | Human-visible improvement |
| --- | --- | --- | --- |
| `0001` Wrong proof surface | `A` pre-closure proof-surface gate | Best direct reduction of false closure on the wrong lane. | Fewer reopenings to reject bad proof. |
| `0002` Direct questions not answered directly | `A` answer-first response gate | Removes repeated answer-shape repair at the first turn. | Direct questions get direct answers first. |
| `0003` Boundaries did not stick | `A` active boundary ledger | Best defense against repeat lane, pass, and ownership resets. | Fewer restatements of the same rules. |
| `0004` Approval/state not human-reviewable | `A` review packet for approval turns | Fixes the review surface instead of only the history. | Approval asks arrive with usable context and links. |
| `0005` Work handed back before closure | `A` stop-eligibility gate | Best match to the packet's measured stall loss. | Fewer wake-up and "continue now" messages. |
| `0006` Symptom-level debugging | `A` first-disagreement debugging gate | Best match to the required root-cause method. | Fewer aimless retunes and clearer cause statements. |

## Overall Priority Order

1. `Problem 0001` Wrong proof surface
2. `Problem 0003` Stated boundaries did not stick
3. `Problem 0005` Work was handed back before real closure
4. `Problem 0006` Debugging stayed at the symptom level too long
5. `Problem 0002` Direct questions did not get direct answers
6. `Problem 0004` Approval and state updates were not human-reviewable

Why this order:

- `0001` is first because false closure on the wrong lane drove repeated reopenings and invalidated later work.
- `0003` is second because boundary drift kept recreating already-corrected failure modes across lane, pass, and method.
- `0005` is third because it carries the packet's clearest measured stall loss.
- `0006` is fourth because truth-seeking debugging reduces repeated human steering during open defect work.
- `0002` is fifth because answer-shape repair is frequent and costly, but it is slightly less destructive than false closure or premature stopping.
- `0004` is sixth because it matters most at approval gates, while the earlier items affect more of the full work loop.
