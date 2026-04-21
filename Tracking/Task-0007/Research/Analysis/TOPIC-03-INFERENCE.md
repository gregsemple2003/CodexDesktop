# Topic 03 Inference

## Frontier Landscape

The frontier corpus points to a fairly stable pattern: the winning systems do not try to “magically infer” everything. They build a context pipeline, a role split, an execution loop, and a verification layer, then make the remaining uncertainty explicit.

In the local task record, that shows up first in the Jarvis framing itself: direct human intervention is treated as a signal that the system boundary was insufficient, but the analysis must separate explicit desire, strong implication, and speculation instead of collapsing them together ([TASK.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/TASK.md), [RESEARCH-ANALYSIS.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/RESEARCH-ANALYSIS.md)).

Across the source packet, the frontier-common pieces are:

- Context engineering as the main lever, not prompt cleverness alone. That is the dominant message in the practitioner writing and the newer corporate architecture paper: the model needs the right information, formatted well, with provenance and isolation ([2025-06-23-harrison-chase-rise-of-context-engineering.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Thought-Leaders/2025-06-23-harrison-chase-rise-of-context-engineering.md), [2026-03-12-context-engineering-corporate-multi-agent-architecture.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2026-03-12-context-engineering-corporate-multi-agent-architecture.md)).
- Explicit role decomposition. The common pattern is planner / coder / tester / debugger / critic, or some close variant of that split ([2025-07-30-harrison-chase-deep-agents.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Thought-Leaders/2025-07-30-harrison-chase-deep-agents.md), [2026-04-13-agentforge-execution-grounded-multi-agent-swe.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2026-04-13-agentforge-execution-grounded-multi-agent-swe.md), [2025-10-27-tdflow-agentic-workflows-test-driven-development.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2025-10-27-tdflow-agentic-workflows-test-driven-development.md)).
- Execution-grounded verification. The frontier is converging on “observe real behavior, then decide,” whether that is tests, rubrics, sandboxed execution, or manual validation artifacts ([2026-01-07-agentic-rubrics-contextual-verifiers-swe-agents.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2026-01-07-agentic-rubrics-contextual-verifiers-swe-agents.md), [2026-04-13-agentforge-execution-grounded-multi-agent-swe.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2026-04-13-agentforge-execution-grounded-multi-agent-swe.md)).
- Human feedback treated as a data source. The preference and intervention papers both treat human correction as something to propagate forward, not as a one-off fix ([2025-03-05-human-preferences-constructive-interactions.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2025-03-05-human-preferences-constructive-interactions.md), [2025-10-02-predictive-preference-learning-human-interventions.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2025-10-02-predictive-preference-learning-human-interventions.md), [2025-06-01-hada-human-ai-agent-decision-alignment.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2025-06-01-hada-human-ai-agent-decision-alignment.md)).

The frontier-common move is to make the system more legible and more constrained. The over-engineered move, for our purposes, is to jump straight to a corporate-grade agent operating system with layered intent/specification pyramids, large role taxonomies, or elaborate governance abstractions before the basic inference loop is trustworthy. That is the shape of [2026-03-12-context-engineering-corporate-multi-agent-architecture.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2026-03-12-context-engineering-corporate-multi-agent-architecture.md), but it is too heavy for this task unless we already had a much larger operating surface.

## Variables That Matter To Humans

The human/operator/user cares about whether the system gets the need right without making them clean up after it.

- Truthfulness: does the system clearly separate evidence, inference, and speculation?
- Honesty about uncertainty: does it say “I think” when it is inferring, and “I do not know” when the signal is thin?
- Local constraint respect: does it honor repo-specific and human-specific limits instead of imposing a global default?
- Compassionate framing: does it treat intervention as unmet need and system miss, not as blame?
- Tolerance for real variation: does it allow one repo, team, or moment to differ from another?
- Actionability: does the output tell the human what to do next, instead of adding another conceptual layer?
- Repetition reduction: does it actually reduce the number of times the human has to correct the same thing?
- Novelty handling: does it ask or defer when the request is genuinely new, rather than pretending it was predictable?

The source material strongly supports these variables. The constructive-interaction paper shows that people prefer reasoning, nuance, and respect, but also that preferences vary and can be skewed by toxic or value-guided prompting, which is exactly why inference must stay evidence-bound ([2025-03-05-human-preferences-constructive-interactions.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2025-03-05-human-preferences-constructive-interactions.md)). The HADA paper and the corporate context-engineering paper both reinforce that local values, policies, and constraints must be explicit and scoped, not assumed universal ([2025-06-01-hada-human-ai-agent-decision-alignment.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2025-06-01-hada-human-ai-agent-decision-alignment.md), [2026-03-12-context-engineering-corporate-multi-agent-architecture.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2026-03-12-context-engineering-corporate-multi-agent-architecture.md)).

## Technique Likely To Work For Us

Use a single **event-sourced inference ledger** with a constrained three-way classification:

- `explicit`
- `strongly implied`
- `speculative`

For each human input, write one compact inference record that captures:

- the raw event
- the inferred need
- the evidence
- the local constraints that applied
- the confidence level
- the recommended next action
- the exact question: `How could the system have inferred the need for this input?`

This is the concrete approach that seems most likely to work here because it matches the corpus in two ways. First, the frontier systems keep converging on structured context, explicit roles, and execution feedback rather than opaque global reasoning ([2025-07-30-harrison-chase-deep-agents.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Thought-Leaders/2025-07-30-harrison-chase-deep-agents.md), [2026-01-07-agentic-rubrics-contextual-verifiers-swe-agents.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2026-01-07-agentic-rubrics-contextual-verifiers-swe-agents.md)). Second, the local task already wants the durable distinction between evidence and speculation, which means a ledger is a better fit than a prose-only narrative ([RESEARCH-ANALYSIS.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/RESEARCH-ANALYSIS.md)).

I would not make this a multi-agent judgment stack yet. That would be frontier-common in the abstract, but over-engineered for us at this stage. We need trustworthy inference before we need a fleet of inference roles.

## Simple Structure That Maximizes The Variables

Keep the structure to three layers:

1. `Event record`
2. `Inference record`
3. `Daily synthesis`

The event record should be mechanically simple: source, timestamp, actor, thread, raw text, and repo context. The inference record should be the only place where the system is allowed to interpret. The daily synthesis should only promote recurring patterns into task proposals or repo contract updates.

That structure maximizes the human variables because it:

- keeps evidence close to the source
- makes uncertainty visible
- prevents inference from leaking into policy too early
- allows local exceptions without rewriting the whole system
- keeps the daily report actionable instead of encyclopedic

The simplest useful daily output is:

- what happened
- what need was inferred
- what evidence supports that inference
- what remains uncertain
- what task proposal or contract change follows

This is aligned with the repo docs that already separate conversation capture, analysis, and promotion boundaries ([Research/README.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/README.md), [PLAN.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/PLAN.md), [HANDOFF.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/HANDOFF.md)).

## Watchouts

- Do not confuse high-frequency intervention with high-confidence inference. Repetition helps, but it does not make a guess true.
- Do not let “truthfulness, compassion, tolerance” become a license to infer private motives. The corpus repeatedly warns against mind-reading and value overreach.
- Do not turn the ledger into a policy engine. Promotion to repo rules should happen only after a pattern is stable.
- Do not overfit on enterprise architecture. The big-pyramid literature is useful as a warning, but most of it is too heavy for this task.
- Do not store hidden reasoning as if it were a durable artifact. Persist the evidence and conclusion, not a theatrical transcript.
- Do not expand the taxonomy faster than you can validate it against real interventions.

## Source Anchors

- [TASK.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/TASK.md)
- [RESEARCH.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/RESEARCH.md)
- [RESEARCH-ANALYSIS.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/RESEARCH-ANALYSIS.md)
- [2026-04-19 Jarvis model capture](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Conversations/2026-04-19-JARVIS-MODEL.md)
- [2026-04-19 System Integration Spec 52](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Conversations/2026-04-19-SYSTEM-INTEGRATION-SPEC-52.md)
- [2026-04-19 System Integration Spec 54](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Conversations/2026-04-19-SYSTEM-INTEGRATION-SPEC-54.md)
- [2025-03-05 human preferences constructive interactions](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2025-03-05-human-preferences-constructive-interactions.md)
- [2025-06-01 HADA](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2025-06-01-hada-human-ai-agent-decision-alignment.md)
- [2025-10-02 predictive preference learning](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2025-10-02-predictive-preference-learning-human-interventions.md)
- [2025-10-27 TDFlow](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2025-10-27-tdflow-agentic-workflows-test-driven-development.md)
- [2026-01-07 agentic rubrics](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2026-01-07-agentic-rubrics-contextual-verifiers-swe-agents.md)
- [2026-03-12 context engineering corporate multi-agent architecture](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2026-03-12-context-engineering-corporate-multi-agent-architecture.md)
- [2026-04-13 AgentForge](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2026-04-13-agentforge-execution-grounded-multi-agent-swe.md)
- [2025-01-07 Chip Huyen agents](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Thought-Leaders/2025-01-07-chip-huyen-agents.md)
- [2025-06-23 Harrison Chase context engineering](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Thought-Leaders/2025-06-23-harrison-chase-rise-of-context-engineering.md)
- [2025-07-30 Harrison Chase deep agents](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Thought-Leaders/2025-07-30-harrison-chase-deep-agents.md)
- [2026-03-01 Simon Willison coding agents](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Thought-Leaders/2026-03-01-simon-willison-how-coding-agents-work.md)
- [2026-03-14 Simon Willison pragmatic summit](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Thought-Leaders/2026-03-14-simon-willison-pragmatic-summit-agentic-engineering.md)
- [OpenAI harness engineering](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-02-11-openai-harness-engineering.md)
- [OpenAI unlocking the Codex harness](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-02-04-openai-unlocking-the-codex-harness.md)
- [Anthropic Claude Code sandboxing](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-01-21-anthropic-claude-code-sandboxing.md)
- [Anthropic Claude Code product](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/anthropic-claude-code-product.md)
- [Cognition how Cognition uses Devin to build Devin](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-02-27-cognition-how-cognition-uses-devin-to-build-devin.md)
- [Cognition Devin can now schedule Devins](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-03-20-cognition-devin-can-now-schedule-devins.md)
- [OpenHands README](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Open-Source/openhands-README.md)
- [OpenAI Agents SDK README](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Open-Source/openai-agents-python-README.md)
- [AutoGen README](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Open-Source/microsoft-autogen-README.md)
- [Deep Agents README](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Open-Source/langchain-deepagents-README.md)
