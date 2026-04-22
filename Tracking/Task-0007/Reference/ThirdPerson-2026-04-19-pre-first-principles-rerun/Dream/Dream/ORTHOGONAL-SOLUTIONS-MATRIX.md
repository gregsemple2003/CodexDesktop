# Orthogonal Solutions Matrix

## Objective

Reduce future human intervention time on ThirdPerson work by blocking the recurring April 19 failure modes before they become reopenings, boundary resets, or approval friction.

## Exemplar Inputs

- Event `3325`: the human reports a failed default-lane regression and asks why the prior regression missed it.
- Event `3431`: the human states the core lane rule.
- Event `4379`: the human cannot approve without a reviewable diff.
- Event `4768`: the task reopens on concrete runtime defects after earlier progress claims.
- Event `5100`: the human rejects non-runtime proof for a runtime claim.
- Event `5340`: the human asks for root cause by tracing disagreement upstream.
- Event `7446`: the human tightens that method into concrete disagreement seams with values.
- `REGRESSION.md` and `TESTING.md`: supporting or operator lanes do not replace human default-lane regression proof.

## First Principles

- The claimed proof surface must match the repo's real acceptance surface.
- A direct question should get a direct answer first.
- Hard constraints should become active gates, not soft reminders.
- Review friction is real work. If approval is hard to perform, the workflow is still broken.
- Symptom lists must be agreed before diagnosis branches widen.
- Root cause needs one measured bad state and its writer chain.
- Active work should continue until a real gate or block, with durable state kept current.

## Problem 0001: Closure kept using the wrong proof lane

### Fundamental Elements

- The repo defines the human default lane as the regression bar.
- Supporting proof still has value, but it cannot close the task by itself.
- The packet shows repeated substitution of easier proof surfaces for the required one.

### Why This Produced A Suboptimal Outcome

The system spent time gathering evidence that could not actually satisfy the human or the repo docs.
That made earlier closeout claims fragile and caused reopenings.

### Option Scoring

Scoring scale: `1` low, `5` high.
`Risk` means rollout safety and clarity, so `5` is safer.

| Option | Type | Time Saved | Enforcement | Repo Fit | Reuse | Risk | Total |
| --- | --- | ---: | ---: | ---: | ---: | ---: | ---: |
| A | Add stronger prompt wording about default-lane proof | 2 | 1 | 4 | 3 | 5 | 15 |
| B | Add a proof-surface gate that checks claimed closure evidence against `REGRESSION.md` and `TESTING.md` before closeout text is emitted | 5 | 5 | 5 | 5 | 4 | 24 |
| C | Require full runtime video capture for every regression closeout | 3 | 4 | 3 | 3 | 2 | 15 |

### Winner

Option `B`.

It creates a hard boundary at the place the bad closure claim is made.
It also matches the repo's existing lane split without forcing heavy new capture on every task.

## Problem 0002: Direct questions did not get direct answers first

### Fundamental Elements

- The packet contains many explicit questions.
- The human repeatedly asked for short answers and small words.
- Delayed answers caused extra turns before work could resume.

### Why This Produced A Suboptimal Outcome

Simple checks turned into repair loops.
The human spent time extracting the first sentence that should have been there already.

### Option Scoring

| Option | Type | Time Saved | Enforcement | Repo Fit | Reuse | Risk | Total |
| --- | --- | ---: | ---: | ---: | ---: | ---: | ---: |
| A | Add a style note that says to answer direct questions first | 2 | 1 | 5 | 4 | 5 | 17 |
| B | Add a reply-shape preflight that detects explicit questions and requires the first line to answer them directly before extra detail | 4 | 4 | 5 | 5 | 4 | 22 |
| C | Create a separate Q-and-A appendix for each pass update | 2 | 3 | 2 | 2 | 3 | 12 |

### Winner

Option `B`.

It changes live behavior at response time.
It also scales beyond this packet because the same failure mode shows up in many direct-question turns.

## Problem 0003: New boundaries did not stay live after the human stated them

### Fundamental Elements

- The packet adds and tightens constraints midstream.
- Examples include repo-local doc ownership, no engine mods, pass framing, and debugging method.
- These constraints needed to shape later work immediately.

### Why This Produced A Suboptimal Outcome

The human had to keep restating live boundaries.
That created avoidable re-planning and blocked trust in later steps.

### Option Scoring

| Option | Type | Time Saved | Enforcement | Repo Fit | Reuse | Risk | Total |
| --- | --- | ---: | ---: | ---: | ---: | ---: | ---: |
| A | Strengthen general docs and hope later work follows them | 2 | 1 | 4 | 4 | 5 | 16 |
| B | Add a boundary manifest that captures current repo, lane, pass, ownership, and hard constraints at start of work and checks later actions against it | 5 | 5 | 5 | 5 | 4 | 24 |
| C | Route all ThirdPerson work through a fixed specialty agent prompt | 3 | 3 | 3 | 3 | 3 | 15 |

### Winner

Option `B`.

This turns live constraints into a durable object instead of relying on memory.
It also helps later audits explain why a proposed action is out of bounds.

## Problem 0004: Approval surfaces were not reviewable enough for a fast human decision

### Fundamental Elements

- Approval requests need change shape, links, and pass framing.
- The packet shows that raw existence of edits was not enough.
- The human needed to understand what changed and what decision was being requested.

### Why This Produced A Suboptimal Outcome

Work paused while the human reconstructed diffs and file context.
That is approval burden the system should remove.

### Option Scoring

| Option | Type | Time Saved | Enforcement | Repo Fit | Reuse | Risk | Total |
| --- | --- | ---: | ---: | ---: | ---: | ---: | ---: |
| A | Add a reminder to include links and a summary | 2 | 1 | 5 | 4 | 5 | 17 |
| B | Auto-build a human approval packet with changed files, concise diff summary, pass framing, and the exact approval ask | 5 | 4 | 5 | 5 | 4 | 23 |
| C | Attach full raw diffs only | 2 | 2 | 3 | 3 | 3 | 13 |

### Winner

Option `B`.

This matches the human complaint directly.
It reduces review cost without hiding the actual file changes.

## Problem 0005: The exact runtime defect set was not locked before analysis widened

### Fundamental Elements

- The packet shows several distinct runtime-only defects.
- The human had to separate them and keep the goal post on the runtime lane.
- Evidence validity also changed during the day.

### Why This Produced A Suboptimal Outcome

Diagnosis drifted because the active symptom list was unstable.
That led to debate about what counted as the defect before root-cause work could settle.

### Option Scoring

| Option | Type | Time Saved | Enforcement | Repo Fit | Reuse | Risk | Total |
| --- | --- | ---: | ---: | ---: | ---: | ---: | ---: |
| A | Add a reminder to restate the defect list in small words | 2 | 1 | 5 | 4 | 5 | 17 |
| B | Add a runtime defect ledger artifact that records lane, agreed symptoms, invalidated evidence, and current proof bar before deeper iteration | 4 | 4 | 5 | 5 | 4 | 22 |
| C | Add automatic screenshot classification to detect foot-shape claims | 2 | 3 | 2 | 2 | 2 | 11 |

### Winner

Option `B`.

It keeps the analysis pinned to the exact runtime defects and the valid evidence set.
It is much lighter and more trustworthy than an image classifier.

## Problem 0006: Root-cause work did not force one concrete bad state and writer chain

### Fundamental Elements

- The human asked for first concrete disagreement tracing, not more symptom-level tuning.
- The packet even gives examples of acceptable measured seams.
- Without a measured seam, fixes can look active while still being guesswork.

### Why This Produced A Suboptimal Outcome

The work kept risking another tweak loop.
That delays real root-cause proof and makes closure hard to trust.

### Option Scoring

| Option | Type | Time Saved | Enforcement | Repo Fit | Reuse | Risk | Total |
| --- | --- | ---: | ---: | ---: | ---: | ---: | ---: |
| A | Link the debugging doc and rely on the writer to follow it | 2 | 1 | 5 | 4 | 5 | 17 |
| B | Add a first-disagreement debug worksheet that requires measured bad values, boundary-by-boundary writer tracing, and explicit stop rules before fix claims | 5 | 5 | 5 | 5 | 4 | 24 |
| C | Build always-on deep pose instrumentation before any future debugging | 4 | 4 | 3 | 4 | 2 | 17 |

### Winner

Option `B`.

It enforces the needed method now without waiting for a larger instrumentation project.
It also preserves contradictory evidence better than ad hoc notes.

## Problem 0007: Work stopped early and pushed coordination back onto the human

### Fundamental Elements

- The packet contains resume prompts and frustration about needless messaging.
- Durable state lagged active work.
- The human wanted the system to keep moving until a real gate or block.

### Why This Produced A Suboptimal Outcome

The human became the workflow scheduler.
That adds both immediate typing cost and stall loss.

### Option Scoring

| Option | Type | Time Saved | Enforcement | Repo Fit | Reuse | Risk | Total |
| --- | --- | ---: | ---: | ---: | ---: | ---: | ---: |
| A | Add a reminder to keep going unless blocked | 2 | 1 | 5 | 4 | 5 | 17 |
| B | Add a checkpoint protocol that defines when to keep working, when to update durable state, and when a user message is allowed during an active pass | 4 | 4 | 5 | 5 | 4 | 22 |
| C | Add an automated heartbeat and resume queue service | 4 | 5 | 2 | 4 | 2 | 17 |

### Winner

Option `B`.

It addresses the workflow failure directly with much lower rollout cost than a new service.
It also fits the packet's complaint about state lag and unnecessary interruption.

## Resolved Winners

| Problem | Winning Option | Chosen Direction |
| --- | --- | --- |
| `0001` | `B` | Proof-surface gate tied to repo regression docs |
| `0002` | `B` | Reply-shape preflight for direct questions |
| `0003` | `B` | Boundary manifest for live hard constraints |
| `0004` | `B` | Human approval packet generator |
| `0005` | `B` | Runtime defect ledger before deeper iteration |
| `0006` | `B` | First-disagreement debug worksheet |
| `0007` | `B` | Checkpoint and continuation protocol |

## Overall Priority Order

1. `Problem 0001 / Option B` because wrong-lane closure invalidates large amounts of later work.
2. `Problem 0003 / Option B` because live boundary drift causes repeated re-planning across the whole task.
3. `Problem 0002 / Option B` because direct-answer failures create constant high-frequency interruption.
4. `Problem 0006 / Option B` because root-cause method drift keeps runtime debugging in a tweak loop.
5. `Problem 0004 / Option B` because approval friction slows safe forward motion and creates more human review work.
6. `Problem 0005 / Option B` because agreed defect framing is needed before deeper runtime debugging can stay honest.
7. `Problem 0007 / Option B` because the continuation protocol matters, but it is safest after the harder gates above are defined.
