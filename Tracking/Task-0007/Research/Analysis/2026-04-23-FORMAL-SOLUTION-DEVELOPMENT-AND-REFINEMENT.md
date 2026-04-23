# Formal Solution Development And Refinement

Date: 2026-04-23

This note synthesizes a targeted source packet on formal solution development and refinement for Task 0007. The packet draws from recent preprints, open-source frameworks, and industry engineering writing:

- [Preprint packet](../Sources/Arxiv/2026-04-23-FORMAL-SOLUTION-DEVELOPMENT-AND-REFINEMENT.md)
- [Open-source packet](../Sources/Open-Source/2026-04-23-FORMAL-SOLUTION-DEVELOPMENT-AND-REFINEMENT.md)
- [Industry packet](../Sources/Industry/2026-04-23-FORMAL-SOLUTION-DEVELOPMENT-AND-REFINEMENT.md)

Corpus size: 30 primary sources normalized into Markdown packets above.

## What The Frontier Converges On

The strongest cross-source pattern is that serious solution development is not treated as one prompt that somehow "becomes" code. The best systems insert explicit intermediate artifacts between problem framing and execution. Those artifacts take different names across the corpus, but they play the same role:

- a specification, feature map, or protocol contract that says what the solution is trying to guarantee
- a verifier layer that forces the proposed solution through tests, rubrics, impact analysis, or guardrails
- a downstream execution stage that consumes the refined contract rather than redefining it casually

That structure appears in research systems such as TDAD, Agentic Rubrics, EvoDev, TDFlow, and SEMAP, in open-source workflows such as Spec Kit and the Agents SDK, and in industry practice from OpenAI, Anthropic, GitHub, and Cognition. The shared claim is simple: if the solution artifact is not explicit, the refinement loop is weak, and downstream stages silently redesign the work.

## What Counts As "Formal" In Practice

The frontier is not converging on heavyweight theorem proving for day-to-day agentic SWE. It is converging on operational formality:

- explicit contracts instead of implied intent
- typed stage boundaries instead of one blended loop
- verifier artifacts instead of persuasive prose
- observable failure modes instead of vague dissatisfaction
- reopen conditions instead of pretending a winner is final forever

In other words, the formalism is mostly workflow and artifact formalism. The solution is made legible enough that later passes can challenge it, execute it, or reject it without reconstructing the author's private reasoning.

## Strongest Patterns From The Corpus

### 1. The solution spec must be a first-class artifact

GitHub Spec Kit, EvoDev, and OpenAI's harness writing all make the same move: the specification or plan is the system of record, and downstream work remains derived from it. This is the clearest support for treating `SOLUTION-DESIGN.md` as the authoritative solution artifact instead of a temporary scratchpad.

### 2. Refinement needs its own stage

Agentic Rubrics, TDAD, TDFlow, and GitHub Agentic Workflows all separate proposal from verification. The winning pattern is not "propose, then immediately draft implementation tasks." It is "propose, evaluate under explicit criteria, narrow, then downstream the result." This supports `2B` as an explicit refinement and adversarial-audit pass rather than a soft synthesis note.

### 3. Exploration pressure matters

SWE-Debate and the broader debate-and-selection family show that multiple materially distinct candidates improve localization and planning quality. This matches the earlier Task-0007 reference packet more than the current two-option rhythm. The frontier evidence favors at least three distinct options when the design seam is important enough to matter later.

### 4. The verifier must be contextual

TDAD's impact analysis and Agentic Rubrics both show that generic refinement advice underperforms context-grounded verification. Good refinement asks which seams are touched, what breaks if the solution is wrong, and what evidence should falsify it in this repository. That is much closer to a local rubric than a generic "does this seem right?" review.

### 5. Long-horizon systems preserve refinement history

Anthropic's harness and context-engineering writing, Magentic-One, and Cognition's Devin posts all treat memory, checkpoints, and replayable state as essential. A winner is rarely just "chosen"; it is chosen with a trail of why, under what constraints, and when it should be reopened.

## Implications For Dream `2A` / `2B` / `2C`

The research points to a clean ownership split:

### `2A` should own the living solution specification

`2A` is where the work should become explicit: burden-to-problem mapping, competing options, boundaries, acceptance tests, falsifiers, impacted seams, and implementation homes. If that information is not durable at the end of `2A`, the refinement stage has nothing formal to refine.

### `2B` should own solution refinement, not just winner announcement

`2B` is where the frontier expects contextual verification:

- challenge whether the option set is sufficiently diverse
- force narrowings and splits when one option hides two seams
- record blocking questions when the current artifact is not decision-ready
- patch the solution spec when the synthesis proves the option design wrong
- emit a winner artifact only after those corrections are made

This is the strongest answer to the current Task-0007 question. The explicit iterated solution doc belongs to `2A` and `2B`, not to `2C`.

### `2C` should consume a refined winner set and draft execution artifacts

`2C` should behave more like Spec Kit's `/tasks` or GitHub Agentic Workflows' compiled execution surface. It should draft task files, sequence rollout work, and expose downstream execution plans. If drafting reveals a genuine flaw in the upstream solution contract, it should patch that contract upstream first rather than silently mutating the design at the task layer.

## Recommended Tightening For Task 0007

Based on the corpus, the current Dream flow would be stronger if it adopted five concrete rules:

1. `2A` must produce at least three materially distinct options for any seam that changes architecture, enforcement, or operator burden.
2. `2B` must emit explicit blocking questions, forced narrowings, or "ready with caveats"; an unconditional clean bill of health should be rare.
3. Every winning option should name its impacted seams and falsifier set, not just its mechanism and intended benefit.
4. `WINNER-SYNTHESIS.md` should include reopen conditions and a short refinement ledger that records what changed between `2A` and `2B`.
5. `2C` may draft tasks, but if task drafting changes winner boundaries, homes, or falsifiers, it must first patch `SOLUTION-DESIGN.md` and `WINNER-SYNTHESIS.md`.

## Bottom Line

The frontier literature and tooling do not support treating formal solution development as a late task-drafting exercise. They support a sequence where:

1. a solution spec is made explicit
2. that spec is refined under contextual verifiers
3. only then are downstream tasks or execution plans produced

For Task 0007, that means the explicit iterated solution doc belongs in `2A` and `2B`. `2C` is downstream consumption with upstream correction rights, not the primary home of solution design.
