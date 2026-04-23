# High-Impact Solution Generation Process Tweaks

Date: 2026-04-23

This note extracts the highest-leverage process changes for Dream-style solution generation when the goal is truth-seeking rather than neat convergence.

It builds on:

- the current shared solution-process contracts in [../../../../../../Users/gregs/.codex/Orchestration/Processes/SOLUTION-CREATE.md](../../../../../../Users/gregs/.codex/Orchestration/Processes/SOLUTION-CREATE.md) and [../../../../../../Users/gregs/.codex/Orchestration/Processes/SOLUTION-AUDIT.md](../../../../../../Users/gregs/.codex/Orchestration/Processes/SOLUTION-AUDIT.md)
- the current Dream pass contracts in [../../../../../../Users/gregs/.codex/Orchestration/Prompts/Dream/WORKFLOW.md](../../../../../../Users/gregs/.codex/Orchestration/Prompts/Dream/WORKFLOW.md), [../../../../../../Users/gregs/.codex/Orchestration/Prompts/Dream/PROMPT-PASS2A-SOLUTION-DESIGN.md](../../../../../../Users/gregs/.codex/Orchestration/Prompts/Dream/PROMPT-PASS2A-SOLUTION-DESIGN.md), and [../../../../../../Users/gregs/.codex/Orchestration/Prompts/Dream/PROMPT-PASS2B-WINNER-SYNTHESIS.md](../../../../../../Users/gregs/.codex/Orchestration/Prompts/Dream/PROMPT-PASS2B-WINNER-SYNTHESIS.md)
- the current Dream packet artifacts in [../../../../../../Users/gregs/.codex/Orchestration/Reports/Interventions/ThirdPerson/2026/04/19/Dream/SOLUTION-DESIGN.md](../../../../../../Users/gregs/.codex/Orchestration/Reports/Interventions/ThirdPerson/2026/04/19/Dream/SOLUTION-DESIGN.md) and [../../../../../../Users/gregs/.codex/Orchestration/Reports/Interventions/ThirdPerson/2026/04/19/Dream/WINNER-SYNTHESIS.md](../../../../../../Users/gregs/.codex/Orchestration/Reports/Interventions/ThirdPerson/2026/04/19/Dream/WINNER-SYNTHESIS.md)
- the stronger reference packet in [../Reference/ThirdPerson-2026-04-19-pre-durable-replay/ThirdPerson/2026/04/19/Dream/ORTHOGONAL-SOLUTIONS-MATRIX.md](../Reference/ThirdPerson-2026-04-19-pre-durable-replay/ThirdPerson/2026/04/19/Dream/ORTHOGONAL-SOLUTIONS-MATRIX.md) and [../Reference/ThirdPerson-2026-04-19-pre-durable-replay/ThirdPerson/2026/04/19/Dream/Task-Candidates/SOLUTION-TASK-0003.md](../Reference/ThirdPerson-2026-04-19-pre-durable-replay/ThirdPerson/2026/04/19/Dream/Task-Candidates/SOLUTION-TASK-0003.md)
- the frontier research synthesis in [2026-04-23-FORMAL-SOLUTION-DEVELOPMENT-AND-REFINEMENT.md](./2026-04-23-FORMAL-SOLUTION-DEVELOPMENT-AND-REFINEMENT.md)

## The Core Failure Pattern

The current Dream process already says the right things at a high level:

- challenge hidden merges
- challenge fake option diversity
- do not freeze winners early
- keep seams explicit

The problem is that the live refinement loop is still too deferential to the first coherent-looking idea.

In the current packet, `2B` returned a clean `ready` verdict with no blocking questions and no required narrowings. That is too easy for a packet with six burden rows and several architecture-level choices. The result is solution reverence: once a plausible option exists, the process tends to polish it rather than trying to break it.

The clearest example is the current durable-learning winner:

- current winner: task-state-oriented `Shared Constraint Ledger And Reread Gate`
- missing challenge: why is this not a repo-wide or shared intervention-learning ledger with explicit promotion rules?
- reference winner family: append-only intervention ledger plus promotion loop

That is exactly the kind of question a truth-seeking refinement loop should force before freezing a winner.

## What The Search Says Actually Improves Solution Generation

The search-backed read from [2026-04-23-FORMAL-SOLUTION-DEVELOPMENT-AND-REFINEMENT.md](./2026-04-23-FORMAL-SOLUTION-DEVELOPMENT-AND-REFINEMENT.md) is that the biggest gains do not come from "more ideation" in the abstract. They come from a small set of techniques that force solution ideas to become explicit, comparable, and breakable before downstream execution starts.

Across the preprint, open-source, and industry corpus, the most impactful techniques are:

### 1. Make The Solution Spec A First-Class Artifact

Documented by:

- TDAD
- EvoDev
- GitHub Spec Kit
- OpenAI harness engineering

Why it matters:

- the strongest systems do not jump from burden statement straight to tasks or implementation
- they create an explicit intermediate artifact that says what the proposed solution is, what it is trying to guarantee, and where it acts
- once that artifact exists, later stages can challenge it instead of reconstructing it from memory

Read for Dream:

- `SOLUTION-DESIGN.md` should be treated as the system of record for solution intent, not as a temporary design scratchpad on the way to `2C`

### 2. Separate Proposal From Verification And Give Verification Real Authority

Documented by:

- Agentic Rubrics
- TDAD
- TDFlow
- GitHub Agentic Workflows
- Anthropic's multi-agent research and harness writing

Why it matters:

- the high-performing workflows do not treat verification as commentary
- they give the verifier stage the power to reject, narrow, or force redesign before the result is allowed downstream
- the technique that keeps recurring is not "generate better ideas," but "make one stage explicitly responsible for trying to prove the current idea insufficient"

Read for Dream:

- `2B` should not mainly summarize winners
- `2B` should be the stage that is allowed to injure the design, force rewrites, and reopen rows

### 3. Increase Comparative Pressure By Generating Nearby Serious Alternatives

Documented by:

- SWE-Debate
- EvoDev
- protocol-driven multi-agent work such as SEMAP
- the stronger April 19 reference packet

Why it matters:

- the frontier repeatedly benefits from multiple competing hypotheses, traces, or plans
- the most useful comparison is not "good option versus weak backup"
- it is "good option versus the closest dangerous alternative that could plausibly win"

Read for Dream:

- the solution set needs more adversarial nearby options, especially on rows that change durable state, learning scope, or enforcement shape

### 4. Use Contextual Verifiers Instead Of Generic Quality Judgments

Documented by:

- Agentic Rubrics
- TDAD graph-based impact analysis
- OpenAI harness writing
- Anthropic context engineering and harness writing

Why it matters:

- generic questions like "does this seem right?" are weak
- the stronger systems ask repository-grounded questions:
  - what seam does this touch?
  - what proof would falsify it?
  - what nearby scope would be safer?
  - what would regress if this were wrong?
- this is what turns refinement from style review into design review

Read for Dream:

- every serious option should be challenged with repo- and packet-specific attacks, not generic plausibility checks

### 5. Preserve Refinement History And Reopen Conditions

Documented by:

- Magentic-One
- Anthropic effective harnesses and context engineering
- Cognition's Devin process writing

Why it matters:

- long-horizon systems do better when they preserve why a choice was made, what was challenged, and what would reopen it
- otherwise winner selection looks final even when it is only locally convenient

Read for Dream:

- a winner should not only say "why this wins"
- it should also say "what challenge it survived" and "what evidence would make us redesign it"

### 6. For Learning And Memory Ideas, Separate Capture From Promotion

Documented by:

- the reference intervention-ledger winner
- GitHub Spec Kit's spec-first downstreaming pattern
- OpenAI's repository-as-system-of-record framing
- Anthropic's durable context and memory patterns

Why it matters:

- many weak designs stop at capture: write the lesson down, reread it later
- stronger designs also define promotion:
  - when does a repeated correction move from task-local fact to repo-local rule?
  - when does a repo-local rule become shared workflow guidance?
  - how is the promotion reversible?

Read for Dream:

- this is the exact technique the current ledger winner missed
- a task-state reread gate captures constraints, but a truth-seeking design also asks how repeated corrections become durable scope-correct improvements

## What Seems Most Impactful Overall

If the search corpus is compressed down to the highest-signal answer, the most impactful solution-generation techniques are:

1. create a living solution spec before tasking
2. give refinement its own adversarial stage with authority to force redesign
3. generate several serious nearby alternatives, including the closest dangerous rival
4. verify with contextual, repository-grounded attacks instead of vibe checks
5. preserve what changed during refinement and what would reopen the choice

Those five techniques appear more consistently, and with more practical force, than heavier additions such as extra orchestration layers, richer dashboards, or more elaborate prose templates.

## Highest-Impact Tweaks

These are ranked by expected improvement in explored solution-space quality, not by ease of implementation.

### 1. Raise The Default Option Count To Three For Consequential Rows

Current weakness:

- `2A` defaults to two options, which makes it too easy to compare one preferred design against one weak foil.

Required change:

- for any row that changes architecture, durable state, enforcement, learning, or operator burden, require at least three materially distinct options
- one of the three must be the `closest dangerous alternative`, not a strawman

Why this matters:

- two-option design tends to produce `my idea` versus `obviously weaker backup`
- three-option design forces a more honest search over nearby shapes
- the reference packet's stronger exploration pressure came from preserving three orthogonal types per problem

Exact insertion points:

- `SOLUTION-CREATE.md`
- `PROMPT-PASS2A-SOLUTION-DESIGN.md`

### 2. Add A Mandatory Scope Ladder Challenge For Every Winner Candidate

Current weakness:

- options are not forced to defend their scope honestly
- a design can land as task-local, repo-local, or shared without the process asking whether it is sitting in the smallest or largest correct home

Required change:

- every consequential option must answer:
  - why is the home task-local?
  - why is it not repo-local?
  - why is it not shared?
  - what evidence would promote or demote its scope?

Why this matters:

- this is the missing question that would have challenged the task-state ledger idea
- repeated-correction and durable-learning mechanisms are especially vulnerable to underscoping
- truth-seeking requires challenging not just the mechanism, but the jurisdiction of the mechanism

Exact insertion points:

- `SOLUTION-CREATE.md`
- `PROMPT-PASS2A-SOLUTION-DESIGN.md`
- `WINNER-SYNTHESIS.md` field contract

### 3. Make `2B` Run Explicit "What If This Assumption Is Wrong?" Attacks

Current weakness:

- current audit language is adversarial in principle, but the actual pass can still stop at "sounds plausible"
- there is no mandatory counterfactual attack pass over the provisional winners

Required change:

- for each provisional winner, `2B` must run at least these attacks:
  - what if the chosen home is one layer too low or too high?
  - what if this seam is actually two seams?
  - what if the burden repeats outside the current task?
  - what if the enforcement boundary is advisory in practice?
  - what if the closest runner-up is actually safer under recurrence?
- if any attack changes the winner boundary, patch `SOLUTION-DESIGN.md` before freezing

Why this matters:

- this is the direct antidote to solution reverence
- the designer has to defend the choice against plausible nearby worlds, not just against obviously weak alternatives

Exact insertion points:

- `SOLUTION-AUDIT.md`
- `PROMPT-PASS2B-WINNER-SYNTHESIS.md`

### 4. Make A Clean `ready` Verdict Rare

Current weakness:

- `2B` can return `ready` with no blockers, no narrowings, and no mishandled seams even on a packet with multiple high-leverage rows

Required change:

- for non-trivial packets, `ready` should require one of:
  - at least one real blocking question was answered during audit and the design was patched
  - at least one real narrowing or split was forced
  - the auditor explicitly records why no blocking attack survived despite trying the required attack set
- a zero-blocker, zero-change audit should be exceptional and defended

Why this matters:

- otherwise the audit stage becomes ceremonial
- truth-seeking refinement should leave evidence that a strong challenge was attempted and defeated

Exact insertion points:

- `SOLUTION-AUDIT.md`
- `PROMPT-PASS2B-WINNER-SYNTHESIS.md`

### 5. Require A "Closest Rival Lost Because..." Defense

Current weakness:

- winners often explain why they are good, but do not always explain why the nearest serious alternative is wrong now

Required change:

- every frozen winner must name:
  - the closest rival
  - the strongest argument for that rival
  - why that rival still loses
  - what evidence would reverse the ranking

Why this matters:

- this forces real comparative reasoning instead of isolated winner prose
- it makes the hidden assumptions visible
- it reduces the chance that the process quietly selected a winner before the runner-up was fully understood

Exact insertion points:

- `SOLUTION-CREATE.md`
- `WINNER-SYNTHESIS.md`

### 6. Add A Special Rule For Learning, Ledger, Memory, And Constraint Options

Current weakness:

- these options are unusually prone to fake sufficiency because they often sound good at any scope
- "ledger" can mean a task checklist, a repo policy surface, a shared cross-repo learning loop, or all three

Required change:

- any option involving memory, ledgering, durable learning, promotion, or retained constraints must also state:
  - `Capture Home`
  - `Promotion Home`
  - `Promotion Trigger`
  - `Scope Guardrails`
  - `Reversal Path`

Why this matters:

- it forces the design to answer whether the system is just storing information or actually improving from it
- it would have exposed the difference between a task-only reread gate and a shared intervention-promotion loop

Exact insertion points:

- `SOLUTION-CREATE.md`
- `PROMPT-PASS2A-SOLUTION-DESIGN.md`

### 7. Add A Durable Refinement Ledger Between `2A` And `2B`

Current weakness:

- the current process can patch the solution doc, but it does not require a crisp record of what the audit changed

Required change:

- `WINNER-SYNTHESIS.md` should carry a short refinement ledger:
  - what `2A` believed initially
  - what `2B` challenged
  - what changed
  - what still remains unresolved

Why this matters:

- it proves the refinement stage actually refined something
- it lets later readers see whether the design survived challenge or just skipped it

Exact insertion points:

- `SOLUTION-AUDIT.md`
- `PROMPT-PASS2B-WINNER-SYNTHESIS.md`
- `WINNER-SYNTHESIS.md` exemplar contract

## Worked Example: Why The Ledger Idea Should Have Been Challenged Harder

Current winner shape:

- shared constraint ledger
- reread before action
- home anchored in task-state/process docs and worker-start/resume prompts

Why that is incomplete:

- it solves "do not forget this constraint"
- it does not fully solve "how do repeated corrections become durable improvements at the correct scope?"

The missing adversarial questions were:

1. If the same class of correction repeats across tasks, why is the primary home still task-state?
2. What promotes a repeated task-local lesson into repo-local or shared policy?
3. How is this different from a task checklist with better prompting?
4. Where is the append-only learning surface that lets later audits see recurrence?
5. What prevents the same lesson from being rediscovered independently in multiple tasks?

The reference packet asked the better question and landed on the stronger shape:

- append-only intervention lesson ledger
- explicit scope values
- promotion thresholds
- durable promotion targets
- reversible promoted rules

That is not just a better task. It is evidence that the current refinement loop did not attack the current winner hard enough.

## Recommended First Pass On Process Tightening

If only four changes are made first, they should be:

1. require three options for consequential rows
2. add the mandatory scope ladder challenge
3. require explicit counterfactual attacks in `2B`
4. make clean `ready` verdicts rare and defended

Those four changes would do the most to break the current pattern of "find one coherent idea, then protect it."

## Bottom Line

The main problem is not that the process lacks solution fields. It already has many of them.

The main problem is that the refinement loop is still too polite.

A truth-seeking solution process should behave more like this:

1. generate several serious nearby options
2. force each one to defend its scope, home, and proof bar
3. attack the provisional winner with realistic counterfactuals
4. redesign when those attacks expose a better shape
5. freeze the winner only after that attack history is durable

That is the direction that best addresses the current failure mode.
