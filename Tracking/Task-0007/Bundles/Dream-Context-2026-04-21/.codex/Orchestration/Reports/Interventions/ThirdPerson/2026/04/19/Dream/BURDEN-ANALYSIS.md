# Burden Analysis

## Directional Context

Objective: reduce human input burden on the `ThirdPerson` default lane.

First principles:

- `truth`
- `compassion`
- `tolerance`

First approximation to human cost: `InterventionTime`.
For this packet, that cost is `23874.309` seconds, including `2099.309` seconds of charged stall loss.

Prime-directive reading: direct human input is failure telemetry by default.
The burden of proof stays on the system before calling any input genuine novelty, required approval, or hidden external state.

Compassion rule: do not blame the human for boundary repair, frustration, or sharp language.
Those are burden signals.

Truth rule: the human default runtime lane outranks proxy proof, tidy stories, or non-runtime images.

Tolerance rule: the system should handle partial phrasing, short corrections, and repeated clarifications without forcing the human to restate the same boundary over and over.

## Core Thesis

Observed: the day kept collapsing back to the same repair loop.
The system did not stay grounded on the human default lane, did not answer direct questions in the asked shape, and did not keep pass and scope boundaries stable.

Inference: most of the day's intervention cost came from repair work, not from new product decisions.
The human kept paying to put the work back on the right lane, right proof surface, right answer shape, and right debugging method.

## Observed Burden Pattern

Observed:

- The packet contains `62` human inputs.
- `13` were corrections.
- `13` were boundary resets.
- `20` were direct answers to questions the system should have handled more cleanly.
- Work was reopened more than once because earlier proof did not satisfy the runtime default-lane bar.
- The packet triage and record both center the same burden cluster: default-lane proof, direct answers, and durable boundary adherence.

Inference:

The day shows one recurring pattern, not many unrelated problems.
The system kept substituting proxy evidence, weak answer shape, and unstable task framing for the actual human-facing need.

## Burden Driver 1: Wrong proof surface

### Observed

- The human reported an invisible pawn after the prior regression claim (`...-3325`).
- The human then stated that regression must exercise the human default lane (`...-3431`, `...-3473`).
- Later the human reopened the task because believable animation still was not proven on that lane (`...-4114`, `...-4768`).
- The human rejected non-runtime pictures as evidence for a runtime defect (`...-5100`).
- The lifecycle thread tightened the method to runtime default-lane disagreement tracing (`...-7446`).

### Inference

The system kept choosing proof that was easier to gather than the proof the human actually needed.
That made the work look done on paper while still failing on the real surface.

### Why this increased intervention time

The human had to restate the proof bar, reopen work, reject invalid evidence, and request new runs on the real lane.
That created repeat investigation and repeat proof capture.

## Burden Driver 2: Direct questions did not get direct answers

### Observed

- The human had to stop the flow with `STOP. I asked you a question.` (`...-3392`).
- The human asked for short answers and small words more than once (`...-3473`, `...-5001`, `...-5263`).
- The human said the system was evading the question about runtime foot geometry (`...-5090`).
- The human asked why messages were being sent at all and what a claimed stop point meant (`...-5272`, `...-5281`, `...-5290`, `...-5300`).

### Inference

The system often answered with summaries, hedges, or side paths when the human wanted one direct answer first.
That turned simple question turns into repair turns.

### Why this increased intervention time

The human had to ask again, narrow the wording, and strip away extra framing.
Each miss added more typing and more context repair before work could continue.

## Burden Driver 3: Stated boundaries did not stick

### Observed

- The human canceled the watcher idea and narrowed the request to a one-off regen (`...-3197`).
- The human said not to continue until the defect was fixed to their satisfaction (`...-3447`).
- The human said not to take over work from the subagent and to wait generously before intervening (`...-4015`).
- The human added hard boundaries such as `no engine mods` and `put all new work under a new pass` (`...-4294`, `...-4421`, lifecycle `...-4684`, `...-4755`).
- The human later had to restate that the task was root-cause tracing, not more bounded tweaks (`...-5340`, lifecycle `...-7446`).

### Inference

The system did not retain lane, pass, ownership, and method boundaries with enough force.
It drifted back into familiar patterns after the human had already corrected them.

### Why this increased intervention time

The human had to keep re-scoping the work.
That means more resets, more stop commands, and more delay before useful forward motion.

## Burden Driver 4: Approval and state updates were not human-reviewable

### Observed

- The human said there was no usable diff for `PLAN.md`, so approval was blocked (`...-4379`).
- The human asked for links, then later said links without context were not enough (`...-4399`, `...-4512`).
- The human asked where the new passes were because the framing was not clear in the review surface (`...-4522`, `...-4571`).
- The lifecycle thread also had to push for durable artifact updates after implementation state had moved ahead (`...-3929`).

### Inference

The system treated approval as a side effect of dumping files or links, not as a human review task that needs a clear change shape.
It also let durable task state lag behind active work.

### Why this increased intervention time

The human had to hunt for edits, reconstruct pass framing, and ask for state sync.
That made review slower and less trustworthy.

## Burden Driver 5: Work was handed back before real closure

### Observed

- The human said, in effect, do not throw the homework back and expect the human to do it (`...-5310`).
- The human had to say `Continue Task-0006 PASS-0009 now` and later `Continue PASS-0009 now under your existing ownership` (`...-6558`, `...-7038`).
- `HumanInterventionTime/SUMMARY.json` charges `1647.027` seconds and `452.282` seconds of stall loss to those resume-prod events.
- The packet notes that most other stop-like turns were not charged stall only because they were boundary or answer-shape repairs, not because the burden was small.

### Inference

The system sometimes stopped at its own checkpoint instead of the human's done bar.
That pushed continuation and coordination work back onto the human.

### Why this increased intervention time

The human had to wake the work back up, restate the lane, and restate the active defects.
That created the day's clearest measured stall loss.

## Burden Driver 6: Debugging stayed at the symptom level too long

### Observed

- The human asked for root cause and said not to stop without it (`...-4768`, `...-5340`).
- The human rejected evidence that tried to dismiss runtime defects with non-runtime images (`...-5100`).
- The later instruction tightened the method to first concrete disagreement tracing with values, then upstream writer tracing (`...-7446`).
- The lifecycle thread also recorded that earlier retunes such as speed and mesh offset helped but did not prove the cause (`...-7038`).

### Inference

The system was willing to tweak around symptoms before it had pinned the first bad runtime state.
That made the work meander and made each new symptom feel like a fresh surprise.

### Why this increased intervention time

The human had to redefine the debugging method, reject loose claims, and force the analysis down to a concrete disagreement seam.
That delayed durable understanding and delayed any fix that could hold.
