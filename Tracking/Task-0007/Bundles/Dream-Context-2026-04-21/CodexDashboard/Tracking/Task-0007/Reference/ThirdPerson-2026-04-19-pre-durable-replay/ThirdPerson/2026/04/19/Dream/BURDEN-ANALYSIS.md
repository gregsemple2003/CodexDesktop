# Burden Analysis

Updated: 2026-04-21

## Core Thesis

The human was not mainly frustrated by one hard Unreal bug.

The larger frustration was repeated burden transfer:

- the system optimized easier-to-prove proxies instead of the human's real closure bar
- the system often answered adjacent questions instead of the asked question
- the system made the human reconstruct approval context, pass structure, and evidence validity
- repeated corrections did not become durable fast enough

## Observed Burden Pattern

Observed from [HumanInterventionTime/SUMMARY.json](../HumanInterventionTime/SUMMARY.json):

- `event_count = 62`
- `correction = 22`
- `boundary_reset = 13`
- `answer_to_question = 15`
- `wake_up = 2`

Observed from the packet content:

- runtime and default-lane disputes recur throughout the day
- approval and pass-structure friction recur throughout the day
- direct-answer requests recur throughout the day
- the human repeatedly has to invalidate evidence that does not match the claimed lane or does not show the claimed defect

Inference:

- the human's time cost came from repeated cleanup of system framing, not only from the existence of a defect

## Frustration Driver 1

### The System Chased The Wrong Closure Bar

Observed:

- the human explicitly said the invisible pawn was seen on the default lane and asked why regression missed it in [HumanInputEvents/INDEX.json](../HumanInputEvents/INDEX.json)
- the human explicitly restated that regression must use the human default lane in the same packet
- later messages explicitly rejected non-runtime or off-lane evidence as closure evidence

Observed examples:

- invisible pawn failure
- insistence that regression must use the human default lane
- repeated rejection of source, export, or non-runtime evidence

Inference:

- the human did not need more proof in general
- the human needed proof on the same surface they actually cared about

Why this increased intervention time:

- once the system claimed success on the wrong lane, every later step became suspect
- the human had to keep reasserting the target surface before any new evidence could be trusted

## Frustration Driver 2

### The System Often Answered Nearby Instead Of Directly

Observed examples from [HumanInputEvents/INDEX.json](../HumanInputEvents/INDEX.json):

- `STOP. I asked you a question.`
- `Agree/disagree?`
- `Short answers.`
- `use small words`
- `Now answer my questions.`
- `Stop evading the question.`
- `Why are you messaging me?`

Inference:

- the human kept paying extra time to get the system back into the answer shape they wanted
- this was not only a tone problem
- it was a turn-structure problem

Why this increased intervention time:

- each unanswered or over-expanded answer created one or more repair turns
- the human had to compress, restate, and narrow the channel manually

## Frustration Driver 3

### The Approval Surface Was Not Usable Enough

Observed examples from [HumanInputEvents/INDEX.json](../HumanInputEvents/INDEX.json):

- `I don't have a diff for plan.md so how am i supposed to approve. Do better.`
- `give me links ffs`
- `Put all new work under a new pass.`
- `where is pass 6?`
- `WHERE ARE THE NEW PASSES`
- `No links. Links without context. Links to headers with no indication of what changed.`

Observed:

- the human had to recover basic approval context that the system should have packaged
- pass numbering and pass ownership were not obvious enough at the point of approval

Inference:

- the human was being asked to act as a parser and navigator instead of as an approver

Why this increased intervention time:

- approval stopped being a lightweight gate
- it turned into manual archaeology through plan history and links

## Frustration Driver 4

### Evidence Kept Failing Validity Checks

Observed examples:

- cropped feet invalidated evidence
- non-runtime pictures were rejected as answers to runtime questions
- source or export claims were rejected as substitutes for same-lane runtime proof

Inference:

- the human was not mainly demanding perfection
- the human was demanding that evidence actually correspond to the claim being made

Why this increased intervention time:

- every invalid artifact forced a new capture loop
- each mismatch weakened trust in later claims even when they were closer to valid

## Frustration Driver 5

### The System Kept Handing Work Back Midstream

Observed examples:

- `Do not take over work from the subagent. Be extremely generous with time; at least 15 minutes of doing nothing before intervention.`
- `I'm trying to discern how I can just get you to keep working without throwing your homework at my feet.`
- `Do not stop after a failed incremental attempt and do not hand intermediate failure back to the user.`

Observed from [HumanInterventionTime/SUMMARY.json](../HumanInterventionTime/SUMMARY.json):

- two wake-up events carried the highest stall losses in the packet

Inference:

- the human did not want frequent progress pings that shifted triage work back onto them
- the human wanted the system to keep ownership until a real blocker or real closure

Why this increased intervention time:

- the human had to restart momentum after pauses
- the human had to decide whether a stop was real or merely agent caution

## Frustration Driver 6

### Repeated Corrections Did Not Become Durable Fast Enough

Observed:

- default-lane and runtime-only rules had to be restated multiple times
- the debugging target had to be restated from cosmetic tweak to concrete disagreement and upstream trace
- plan structure and pass boundaries had to be restated after being misunderstood

Inference:

- the human experienced each repeated correction as a new tax
- the largest emotional load likely came from watching the same lessons fail to stick

Why this increased intervention time:

- every repeated restatement consumed time and reduced trust
- later instructions had to be longer and more explicit because earlier ones had not stabilized behavior

## What The Human Seems To Have Needed

Observed:

- truthful lane claims
- direct answers
- approval artifacts that are easy to use
- valid evidence
- continuity of ownership
- root-cause discipline instead of repeated bounded tweaks

Inferred:

- the human wanted to supervise outcomes, not reconstruct context
- the human wanted the system to carry more of the orchestration burden itself
- the human wanted the same correction to stick the second time, not the sixth

## Most Important Conclusion

The best way to reduce future human intervention time is not to make the system more verbose or more apologetic.

It is to stop these specific burden transfers:

- wrong closure surface
- wrong answer shape
- unusable approval surface
- invalid evidence
- premature hand-back
- non-durable learning

