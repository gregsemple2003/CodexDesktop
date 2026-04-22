# Orthogonal Solutions Matrix

Updated: 2026-04-21

## Objective

Build on the earlier `Dream` task proposals by:

- framing each problem from first principles
- proposing three orthogonal solution types per problem
- scoring them with one explicit weighted matrix
- resolving the winner for this repo and this failure pattern

The sole north star is:

- reduce future human intervention time

## Exemplar Inputs

This brief uses the following as raw input:

- local burden analysis in [BURDEN-ANALYSIS.md](./BURDEN-ANALYSIS.md)
- shared process guidance in [AUTOIMPROVEMENT.md](../../../../../../../Processes/AUTOIMPROVEMENT.md) and [DEBUGGING.md](../../../../../../../Processes/DEBUGGING.md)
- exemplar comparison shape in [AUTOIMPROVEMENT-PROBLEM-0001.md](../../../../../../../Exemplars/AUTOIMPROVEMENT-PROBLEM-0001.md) and [AUTOIMPROVEMENT-BRIEF.md](../../../../../../../Exemplars/AUTOIMPROVEMENT-BRIEF.md)
- local research synthesis in [TOPIC-06-LOCAL-MEMORY-AND-LEARNING.md](../../../../../../../../../../../Agent/CodexDashboard/Tracking/Task-0007/Research/Analysis/TOPIC-06-LOCAL-MEMORY-AND-LEARNING.md), [TOPIC-08-AUTONOMY-AND-SAFETY-POLICY.md](../../../../../../../../../../../Agent/CodexDashboard/Tracking/Task-0007/Research/Analysis/TOPIC-08-AUTONOMY-AND-SAFETY-POLICY.md), [TOPIC-09-EVALUATION-LOOPS.md](../../../../../../../../../../../Agent/CodexDashboard/Tracking/Task-0007/Research/Analysis/TOPIC-09-EVALUATION-LOOPS.md), and [PRIORITIZATION.md](../../../../../../../../../../../Agent/CodexDashboard/Tracking/Task-0007/InterventionTime/PRIORITIZATION.md)
- best-practice source captures:
  - [2026-02-04-openai-unlocking-the-codex-harness.md](../../../../../../../../../../../Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-02-04-openai-unlocking-the-codex-harness.md)
  - [2026-02-11-openai-harness-engineering.md](../../../../../../../../../../../Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-02-11-openai-harness-engineering.md)
  - [2026-01-21-anthropic-claude-code-sandboxing.md](../../../../../../../../../../../Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-01-21-anthropic-claude-code-sandboxing.md)
  - [2026-03-20-cognition-devin-can-now-schedule-devins.md](../../../../../../../../../../../Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-03-20-cognition-devin-can-now-schedule-devins.md)
  - [2026-01-07-agentic-rubrics-contextual-verifiers-swe-agents.md](../../../../../../../../../../../Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2026-01-07-agentic-rubrics-contextual-verifiers-swe-agents.md)
  - [2025-10-27-tdflow-agentic-workflows-test-driven-development.md](../../../../../../../../../../../Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2025-10-27-tdflow-agentic-workflows-test-driven-development.md)

## First Principles

Human intervention time rises when the system exports unresolved coordination work to the human instead of resolving it internally.

The exported work on April 19 repeatedly took one of seven forms:

1. `target mismatch`
   - the system pursued the wrong closure surface or wrong lane
2. `answer mismatch`
   - the system did not answer the asked question in the requested shape
3. `approval mismatch`
   - the human could not approve because context was not packaged for approval
4. `evidence mismatch`
   - the artifact did not actually support the claim being made
5. `ownership mismatch`
   - the task stopped and the human had to restart supervision
6. `learning mismatch`
   - repeated corrections did not become durable quickly enough
7. `debugging mismatch`
   - the investigation drifted into symptom iteration instead of narrowing to a concrete disagreement and upstream writer

These failures all share the same deeper structure:

- the system knew too little, or preserved too little, at the point where a boundary had to be crossed

The relevant boundaries are:

- task target boundary
- answer boundary
- approval boundary
- evidence boundary
- handoff boundary
- memory boundary
- debugging boundary

Suboptimal outcomes happen when one of these is:

- implicit instead of explicit
- transient instead of durable
- advisory instead of enforced
- generic instead of repo-local
- unverifiable instead of checkable

This is why mere better wording is usually not enough. Wording can improve intent, but April 19 shows that the highest-cost failures were not missing prose alone. They were missing contracts, missing gates, and missing durable state.

## Scoring Matrix

Weighted criteria:

- `IR` = expected reduction in future human intervention time, weight `5`
- `Prev` = how much the option prevents the failure before the human sees it, weight `4`
- `Safe` = truthfulness, safety, and false-completion resistance, weight `4`
- `Fit` = local fit with current orchestration and repo artifacts, weight `3`
- `Speed` = implementation speed and rollout simplicity, weight `2`
- `Eval` = ease of measuring whether it actually helped, weight `2`

Scoring method:

- each criterion is scored `1-5`
- weighted total = `IR*5 + Prev*4 + Safe*4 + Fit*3 + Speed*2 + Eval*2`
- max score = `100`

Why these weights:

- the user explicitly asked that human intervention time be the sole ranking criterion
- prevention is weighted next because cleanup after the fact is expensive
- truthfulness and safety remain high because false closure creates more future intervention time, not less

## Problem 1

### Wrong Closure Surface Or Wrong Lane

Fundamental elements:

- a task has a real closure surface
- evidence has a lane and type
- a closure claim is only honest if it is bound to the right lane and evidence type

Why this produced suboptimal outcomes:

- once the system claimed progress on the wrong lane, the human had to spend time re-establishing the goal post before any later proof could be trusted

Orthogonal solutions:

- `A. Policy rule`
  - strengthen prompts and repo docs to restate that closure must be on the human default lane
- `B. Structured claim manifest`
  - require every closure or regression claim to carry explicit lane, evidence type, runtime flag, and closure-scope fields
- `C. Claim gate`
  - block `passed` or `closed` claims unless the manifest and artifact set satisfy repo-local lane and runtime rules

Scores:

| Option | IR | Prev | Safe | Fit | Speed | Eval | Total |
| --- | --- | --- | --- | --- | --- | --- | --- |
| A | 2 | 2 | 3 | 5 | 5 | 2 | 59 |
| B | 4 | 4 | 5 | 5 | 4 | 5 | 89 |
| C | 5 | 5 | 5 | 4 | 3 | 5 | 93 |

Winner:

- `C. Claim gate`

Reason:

- April 19 already had prose and still failed
- the winning move is enforcement, not restatement
- `B` is a required substrate, but `C` is the intervention-time winner because it stops false closure before it escapes the boundary

## Problem 2

### Wrong Answer Shape For Explicit Questions

Fundamental elements:

- a question creates an answer contract
- the contract includes both content and answer shape
- if the first response block misses the contract, the human has to spend a repair turn to restore it

Why this produced suboptimal outcomes:

- many short corrective turns existed only because the system answered nearby, expanded too much, or continued process narration before answering

Orthogonal solutions:

- `A. Answer-first wording rule`
  - revise prompts and docs to tell the system to answer direct questions directly and briefly
- `B. Structured reply modes`
  - introduce explicit reply modes such as `yes_no_reason`, `agree_disagree_reason`, and `short_answer`
- `C. Unanswered-question validator`
  - detect explicit question forms and block output unless the first response block satisfies the requested answer shape

Scores:

| Option | IR | Prev | Safe | Fit | Speed | Eval | Total |
| --- | --- | --- | --- | --- | --- | --- | --- |
| A | 3 | 2 | 3 | 5 | 5 | 2 | 64 |
| B | 4 | 3 | 4 | 4 | 4 | 4 | 76 |
| C | 5 | 5 | 4 | 4 | 3 | 5 | 89 |

Winner:

- `C. Unanswered-question validator`

Reason:

- answer-shape failures are frequent, cheap to detect, and expensive to keep leaking
- `B` improves consistency, but `C` is what stops `I asked you a question` repair turns from recurring

## Problem 3

### Approval Surface Forced Manual Archaeology

Fundamental elements:

- approval is a human review boundary
- a good boundary packages delta, reason, and location
- if the human must reconstruct any of those manually, the approval surface is failing

Why this produced suboptimal outcomes:

- approval turned into navigation and reconstruction work instead of a lightweight decision

Orthogonal solutions:

- `A. Approval checklist`
  - strengthen docs that say approval requests must include diffs, links, and pass context
- `B. Approval packet generator`
  - auto-generate compact approval packets with changed sections, old vs new wording, clickable links, one-line reasons, and current pass ownership
- `C. Dedicated review UI`
  - build a richer approval panel or PR-style review surface for plans and pass artifacts

Scores:

| Option | IR | Prev | Safe | Fit | Speed | Eval | Total |
| --- | --- | --- | --- | --- | --- | --- | --- |
| A | 2 | 2 | 3 | 5 | 5 | 3 | 61 |
| B | 5 | 5 | 4 | 5 | 4 | 5 | 94 |
| C | 4 | 4 | 4 | 3 | 1 | 4 | 71 |

Winner:

- `B. Approval packet generator`

Reason:

- it directly targets the approval burden cluster without requiring a full product-surface build
- it follows the exemplar pattern of keeping repository knowledge legible and reviewable in-repo instead of inventing a larger platform first

## Problem 4

### Evidence Did Not Support The Claim

Fundamental elements:

- an artifact is useful only if it binds to a claim, lane, and region of interest
- evidence should be invalidated as early as possible
- if the human is the primary validator, the system is exporting proof hygiene work

Why this produced suboptimal outcomes:

- invalid evidence forced recapture loops and eroded trust in later artifacts

Orthogonal solutions:

- `A. Capture guidance`
  - add stronger capture checklists for runtime, region visibility, and proof-lane discipline
- `B. Evidence manifest`
  - require every proof artifact set to declare claim, lane, region of interest, runtime status, and allowed use
- `C. Evidence linter`
  - automatically reject or flag proof bundles that are off-lane, wrong-type, or missing declared subject/region coverage before presenting them

Scores:

| Option | IR | Prev | Safe | Fit | Speed | Eval | Total |
| --- | --- | --- | --- | --- | --- | --- | --- |
| A | 2 | 2 | 3 | 5 | 5 | 3 | 61 |
| B | 4 | 4 | 5 | 5 | 4 | 5 | 89 |
| C | 4 | 5 | 4 | 4 | 3 | 4 | 83 |

Winner:

- `B. Evidence manifest`

Reason:

- full automated validation is attractive, but without a stable manifest it risks brittle guesses
- the first high-leverage move is to make evidence claims explicit and machine-checkable
- `C` becomes stronger after `B` exists

## Problem 5

### Manual Restart Supervision

Fundamental elements:

- active work needs an owner
- stopping is a state transition, not a vibe
- after a recoverable failure, the system should either continue or declare a narrowly justified block

Why this produced suboptimal outcomes:

- the human repeatedly had to become the scheduler again by deciding whether work should resume, pause, or stay owned

Orthogonal solutions:

- `A. Stronger standing instruction wording`
  - emphasize `keep working`, `do not stop`, and `do not hand homework back`
- `B. Persisted continuity contract`
  - store ownership, standing instructions, allowed stop reasons, blocked status, and next planned attempt in durable task state
- `C. Background watcher/scheduler`
  - add heartbeat, timeout recovery, and automatic wake/resume tooling around long-running or delegated work

Scores:

| Option | IR | Prev | Safe | Fit | Speed | Eval | Total |
| --- | --- | --- | --- | --- | --- | --- | --- |
| A | 3 | 2 | 3 | 5 | 5 | 2 | 64 |
| B | 5 | 5 | 5 | 5 | 4 | 5 | 98 |
| C | 4 | 4 | 4 | 3 | 2 | 4 | 73 |

Winner:

- `B. Persisted continuity contract`

Reason:

- this is the smallest change that actually changes who has to remember what
- `C` is useful later, but without `B` a watcher merely automates ambiguity
- the right first move is durable ownership and stop-state semantics

## Problem 6

### Repeated Corrections Did Not Stick

Fundamental elements:

- repeated human corrections are preference and policy data
- useful learning must stay inspectable, scoped, and reversible
- if learning remains only in chat context, the same tax will recur

Why this produced suboptimal outcomes:

- the human had to restate stable repo-local constraints multiple times because the promotion path from correction to durable rule was weak

Orthogonal solutions:

- `A. Manual memory upkeep`
  - rely on people to patch docs or prompts when they notice a repeated problem
- `B. Intervention ledger plus promotion loop`
  - capture corrections append-only, classify them by confidence and scope, and promote repeated patterns into the smallest correct durable home
- `C. Hidden learned profile`
  - keep a model-side memory or preference profile that changes future behavior without an explicit artifact trail

Scores:

| Option | IR | Prev | Safe | Fit | Speed | Eval | Total |
| --- | --- | --- | --- | --- | --- | --- | --- |
| A | 3 | 2 | 4 | 4 | 5 | 3 | 67 |
| B | 5 | 4 | 5 | 5 | 4 | 5 | 94 |
| C | 4 | 4 | 2 | 2 | 2 | 1 | 56 |

Winner:

- `B. Intervention ledger plus promotion loop`

Reason:

- it matches the strongest local research conclusion: durable learning should live in inspectable repo artifacts
- it reduces repetition without hiding policy in opaque state

## Problem 7

### Debugging Drifted From Narrowing To Symptom Tweaks

Fundamental elements:

- debugging must start at the first concrete disagreement
- each step should narrow the plausible cause space
- root-cause claims need a trace from bad state to upstream writer

Why this produced suboptimal outcomes:

- the human had to restate that the work was not `another bounded tweak`, but a narrowing path from concrete runtime disagreement to upstream writer

Orthogonal solutions:

- `A. Better debugging prose`
  - strengthen prompts and repo docs that describe the narrowing method
- `B. Required disagreement template`
  - require bug notes and debug passes to name expected state, observed state, disagreement seam, and next upstream writer
- `C. Root-cause claim verifier`
  - block `root cause found` or equivalent debug closeout unless the artifact set contains a concrete disagreement seam and a traced writer chain

Scores:

| Option | IR | Prev | Safe | Fit | Speed | Eval | Total |
| --- | --- | --- | --- | --- | --- | --- | --- |
| A | 2 | 2 | 4 | 5 | 5 | 2 | 63 |
| B | 4 | 4 | 5 | 5 | 4 | 5 | 89 |
| C | 5 | 5 | 5 | 4 | 3 | 5 | 93 |

Winner:

- `C. Root-cause claim verifier`

Reason:

- the debugging prose already existed in shared guidance
- the failure persisted because the method was not enforced at the point where closure claims were made

## Resolved Winners

Per problem:

1. wrong closure surface
   - winner: `claim gate`
2. wrong answer shape
   - winner: `unanswered-question validator`
3. unusable approval surface
   - winner: `approval packet generator`
4. invalid evidence
   - winner: `evidence manifest`
5. manual restart supervision
   - winner: `persisted continuity contract`
6. non-durable learning
   - winner: `intervention ledger plus promotion loop`
7. debugging drift
   - winner: `root-cause claim verifier`

## Overall Priority Order

Plain-language version:

1. Save task ownership and stop rules in durable task state.
   - This reduces the biggest direct burden: you having to restart supervision and tell the system to keep going.
2. Build one clean approval bundle whenever approval is requested.
   - This removes the repeated burden of asking for diffs, links, and pass context before you can review anything.
3. Keep a log of repeated corrections and turn stable ones into docs or checks.
   - This stops the same correction from staying trapped in chat and having to be repeated later.
4. Block `done` or `passed` claims when the proof comes from the wrong lane.
   - This prevents false closure on evidence the human never agreed counted as closure.
5. Block `root cause found` claims unless the debug notes show the exact bad state and the code path that produced it.
   - This reduces the repeated need to pull debugging back from vague symptom tweaks to concrete narrowing.
6. Check whether the reply actually answers the user's explicit question first.
   - This removes many short but frequent repair turns such as `I asked you a question`.
7. Require each proof bundle to say what claim it supports and what it actually shows.
   - This makes later proof checking easier and clearer, but it depends less directly on the highest-cost failures above.

## Why These Winners

Across the seven problems, the same pattern repeats:

- policy-only options are cheap but too weak
- heavyweight UI or large-governance options are real but too expensive or too early
- the best fit is usually a repo-local contract or generator plus a narrow gate or verifier at the boundary where the failure leaks to the human

This is also the common best-practice pattern from the exemplar set:

- OpenAI harness materials favor threads, durable items, and repository knowledge as the system of record
- Anthropic sandboxing favors bounded autonomy that reduces approval fatigue without losing control
- Cognition scheduling favors continuity across sessions rather than restarts from scratch
- Agentic Rubrics favors repository-grounded verifiers over generic judging
- TDFlow favors execution-grounded iterative loops with anti-cheating constraints rather than vague exploration

The shared lesson is:

- the frontier move is not a more eloquent prompt
- it is a small number of explicit, durable, reviewable contracts at the boundaries where human time is currently being consumed

## Implementation Note

Several winners rely on a substrate:

- the `claim gate` needs structured claim metadata
- the `root-cause claim verifier` needs a required disagreement template
- the `evidence manifest` should precede more ambitious evidence linting

So the right rollout is not to treat all winners as independent product features.

It is to implement the enabling contracts first, then the narrow gates that enforce them.
