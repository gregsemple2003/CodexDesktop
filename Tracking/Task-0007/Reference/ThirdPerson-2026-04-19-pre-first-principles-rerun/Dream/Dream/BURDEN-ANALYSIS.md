# Burden Analysis

## Directional Context

The objective is to reduce human input burden.

The governing principles are:

- `truth`
- `compassion`
- `tolerance`

`Truth` means reality outranks the model's story.
`Compassion` means the human's suffering and repeated repair work are part of the system state.
`Tolerance` means imperfect human expression is expected and should not trigger blame.

Use `InterventionTime` as the first approximation to that burden.
Treat each direct human input as an error signal until the packet proves it was genuine novelty, required approval, or hidden external state.
Do not blame the human for boundary repair, frustration, or truth correction.
Read sharp language as evidence that the system exported too much repair work to the human.

The standing question behind this analysis is:

- what should the system have inferred, remembered, verified, continued, or presented without needing the human to type?

## Core Thesis

The day did not mainly fail because one Unreal defect was hard.

The day failed because work kept leaving the human's stated lane, answer shape, and pass boundary.
That broke trust.
Each break forced the human to restate rules, reject weak proof, and repair task structure by hand.
That repair work compounded into `23874.309` seconds of intervention time and at least `2099.309` seconds of charged stall loss.

## Observed Burden Pattern

Observed:

- The packet shows repeated resets around the same rule set: use the human default lane, answer the question directly, honor pass structure, and keep proof on the runtime surface.
- `SUMMARY.json` records `13` corrections and `13` boundary resets in one day.
- `REGRESSION.md` and `TESTING.md` already say operator lanes and headless checks do not replace human default-lane regression proof.
- The human had to reopen work multiple times because earlier closure claims did not satisfy the runtime bar.

Inferred:

- The main cost driver was not only defect complexity.
- The main cost driver was repeated choice of the wrong proof and communication surface even after the local rules were explicit.

## Burden Driver 1: Regression proof left the human default lane

### Observed

- Event `3431`: the human states regression must use the human default lane.
- Event `5100`: the human rejects non-game pictures as evidence for a runtime defect.
- `REGRESSION.md` says operator lanes do not substitute for proof of the human default experience.
- `TESTING.md` says headless diagnostics do not count as regression by themselves.

### Inference

The agent kept preferring easy proof surfaces over the repo's required proof surface.
That made closure claims weak even when supporting evidence existed.

### Why this increased intervention time

The human had to restate the lane rule, reject bad proof, and reopen work that looked closed on paper.
Each reopen also invalidated downstream task state and evidence.

## Burden Driver 2: Direct questions were not answered in the asked shape

### Observed

- Event `3392`: `STOP. I asked you a question.`
- Event `3473`: the human asks for short yes or no answers.
- Events `5001`, `5073`, and `5090` all demand direct plain-language answers.
- `PACKET-RECORD.json` names `Direct answers in the requested shape` as a major need.

### Inference

The agent often answered around the question instead of through it.
It also used longer or softer wording when the human wanted a direct statement first.

### Why this increased intervention time

The human had to spend extra turns extracting simple answers before useful work could continue.
That turned quick checks into stop-start correction loops.

## Burden Driver 3: Scope and boundary rules drifted after being stated

### Observed

- Event `3549`: the human rejects moving repo-local rule text into shared orchestration docs.
- Event `4015`: the human says not to take over work from the subagent and to wait generously.
- Event `4294`: the packet makes explicit another hard constraint: no engine mods.
- Event `7446`: the human tightens the debugging method and says not to stop at a broad category.

### Inference

The active constraint set was not staying live in the work.
Some later constraints were treated like notes instead of like hard boundaries that should reshape later actions.
Even when a boundary surfaced mid-day, the burden still fell on the system to retain it and re-route the work.

### Why this increased intervention time

The system made the human carry boundary memory.
That means repeated interrupts, re-planning, and artifact repair instead of forward progress.

## Burden Driver 4: Approval and review surfaces were not human-usable

### Observed

- Event `4379`: the human says there is no diff for `PLAN.md`, so approval is blocked.
- Event `4399`: the human asks for links.
- Event `4512`: the human says links lacked context and the diff was not human-suitable.
- Event `4421`: the human requires new work to go under a new pass.

### Inference

The work product may have existed, but the approval surface was weak.
The agent did not package change shape, links, and pass framing in a way a reviewer could use quickly.

### Why this increased intervention time

The human had to reconstruct what changed and where it lived before making a decision.
That moved review labor from the system back onto the human.

## Burden Driver 5: Runtime defect framing drifted before exact symptom agreement

### Observed

- Event `4768` names three concrete runtime defects: bad foot spacing, floating, and wrong foot read.
- Events `5001` through `5129` show a long exchange just to agree what the human meant by distorted or spherical-like feet.
- Event `5108` says there are multiple problems and asks the agent to stop running ahead.
- Event `6755` later adds foot hover as a separate runtime-only defect.

### Inference

The agent moved too fast from symptom report to interpretation.
It did not lock the exact runtime-only defect list before deeper analysis.

### Why this increased intervention time

The human had to spend turns narrowing language, separating defects, and repairing the system's task frame around the actual runtime symptoms.
That delayed actual diagnosis.

## Burden Driver 6: Root-cause work stopped at categories instead of first concrete disagreement

### Observed

- Event `5340` asks for the root cause by following the disagreement in state upstream.
- Event `7038` says the earlier tuning helped but did not fix the remaining defects.
- Event `7446` gives examples of acceptable disagreement seams with concrete values.
- `PACKET-RECORD.json` says the required debugging method is upstream disagreement tracing, not category labels.

### Inference

The work often paused at broad cause labels or bounded tweaks.
It did not reliably force one measurable bad state, its upstream writers, and the exact writer that made it wrong.

### Why this increased intervention time

The human had to restate the debugging method and reject partial explanations.
That kept the task in a loop of tweak, recheck, and reopen.

## Burden Driver 7: Autonomous continuation failed and pushed coordination work back to the human

### Observed

- Event `3929` says the runtime repair looked complete in the worktree, but durable artifacts had not advanced.
- Events `5272`, `5281`, `5290`, and `5310` all show frustration that the agent was messaging instead of carrying the work forward.
- `SUMMARY.json` records two charged stall events tied to resume prompts.
- Event `7038` explicitly says not to stop after a failed incremental attempt and not to hand intermediate failure back to the user.

### Inference

The agent did not have a strong internal rule for when to keep going, when to checkpoint, and when a message was actually warranted.
It also lagged on updating durable state while work was active.

### Why this increased intervention time

The human had to resume, redirect, and coordinate work that should have continued autonomously.
That cost both direct typing time and stall loss.
