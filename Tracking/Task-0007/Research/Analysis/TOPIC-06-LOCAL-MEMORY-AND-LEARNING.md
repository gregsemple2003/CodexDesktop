# Topic 06 Local Memory And Learning

The corpus points to one durable conclusion: local memory should live in inspectable repo artifacts, and learning should happen by promoting repeated interventions into those artifacts, not by hiding everything in an opaque model state.

## Frontier Landscape

The frontier is converging on the same basic shape from different directions. Chip Huyen separates short-term, long-term, and hybrid memory and treats planning, reflection, and error correction as core agent capabilities. Harrison Chase frames this as context engineering: the system has to assemble the right information, in the right format, at the right time. Simon Willison and the OpenAI harness/App Server material push the same idea harder: the repo, the thread history, and the protocol trace are the real memory surface, not a vague prompt blob. Cognition adds the operational version of that with Ask Devin, DeepWiki, Playbooks, Session Insights, and scheduled Devins that carry state forward across runs. DeepAgents, AutoGen, and OpenAI Agents SDK all converge on the same practical ingredients: file-backed working state, sub-agents or handoffs, traceability, and explicit human-in-the-loop checkpoints.

The strongest frontier-common pattern is this: memory is not a single vector store or a giant summary; it is a system of sources, traces, and rules that can be inspected and updated. The less convincing frontier pattern is the “big agent mesh” answer to everything. Broad protocol stacks, universal value ontologies, or multi-layer orchestration graphs are useful when the organization is large enough to need them, but they are over-engineered for this topic if the goal is simply to stop repeating the same local correction and to make future behavior better.

The research side of the corpus also matters. PPL shows the cleanest learning pattern here: human interventions become preference data that can influence future rollouts instead of staying local to the moment. HADA shows a higher-level governance version of the same idea: keep stakeholder intent, values, and auditability connected over time. TDFlow shows the workflow side: break the work into isolated components, keep context narrow, and let the system accumulate reports instead of one sprawling session. Those are all variants of the same insight: durable learning comes from explicit records plus a controlled promotion path.

## Variables That Matter To Humans

The human does not mainly care about the internal representation. The human cares about whether the system stops making them repeat the same correction, whether it respects local norms, and whether it can explain what it learned without pretending to know more than it does.

The variables that matter most are:

- Repetition burden: how often the same correction has to be restated.
- Local correctness: whether the memory reflects the actual repo, team, or directory rule instead of a generic global guess.
- Trustworthiness: whether the system distinguishes evidence, inference, and speculation.
- Recoverability: whether stale or bad learning can be demoted or forgotten cleanly.
- Inspectability: whether a human can see why something was remembered and where it came from.
- Scope control: whether a lesson applies only where it should.
- Attention cost: whether the system reduces review and babysitting work instead of moving it to a new place.
- Behavioral impact: whether the remembered lesson actually changes future outputs.

These are human-facing variables, not storage-engine variables. A clever memory backend that increases surprise, ambiguity, or cleanup work is a failure even if it is technically sophisticated.

## Technique Likely To Work For Us

Use an append-only intervention ledger with explicit promotion gates.

Every correction, repeated request, or notable intervention becomes a durable local record with source links, scope, observed signal, inferred need, confidence, and whether the lesson is explicit, implied, or speculative. The ledger stays append-only. A daily or per-task synthesis reads that ledger, groups repeated patterns, and proposes only a small number of candidate updates. Those candidates are then promoted into the right repo artifact only when the evidence is strong enough: `HUMAN-DESIRE.md`, `AGENTS.md`, tests, validators, or task proposals.

This is the smallest approach that matches the frontier without importing its complexity. It works because the learning path is visible, reviewable, and reversible. It also matches the local corpus directly: interventions become future preference data, the repo becomes the system of record, and recurring patterns become explicit rules rather than invisible habits.

## Simple Structure That Maximizes The Variables

The simplest structure is:

1. Capture the input.
2. Record the source and scope.
3. Classify the lesson as explicit, implied, or speculative.
4. Keep the raw record append-only.
5. Let a daily report extract only repeated or high-value patterns.
6. Promote stable patterns into repo docs or checks.
7. Demote stale patterns with an explicit newer record, not silent deletion.

That structure maximizes the human variables because it lowers repetition, keeps local rules local, preserves trust, and makes forgetting visible. It also stays operationally simple: one write path, one review path, one promotion rule. No vector database is required to start. No universal human model is required. No hidden memory mutation is required.

## Watchouts

- Do not let summaries replace source records.
- Do not turn every repeated request into a global rule.
- Do not store memory so abstractly that it becomes uninspectable.
- Do not promote lessons that are still speculative.
- Do not let the memory layer become a second, harder-to-review policy engine.
- Do not confuse “we can store it” with “we should remember it.”
- Do not overbuild the retrieval stack before the promotion path is stable.

## Source Anchors

- Task framing: [TASK.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/TASK.md), [RESEARCH.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/RESEARCH.md), [RESEARCH-PLAN.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/RESEARCH-PLAN.md), [RESEARCH-ANALYSIS.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/RESEARCH-ANALYSIS.md), [HANDOFF.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/HANDOFF.md)
- Conversation captures: [2026-04-19-JARVIS-MODEL.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Conversations/2026-04-19-JARVIS-MODEL.md), [2026-04-19-SYSTEM-INTEGRATION-SPEC-52.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Conversations/2026-04-19-SYSTEM-INTEGRATION-SPEC-52.md), [2026-04-19-SYSTEM-INTEGRATION-SPEC-54.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Conversations/2026-04-19-SYSTEM-INTEGRATION-SPEC-54.md)
- Bootstrap analysis: [2026-04-19-BOOTSTRAP-ANALYSIS.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Analysis/2026-04-19-BOOTSTRAP-ANALYSIS.md)
- Arxiv corpus: [2025-03-05-human-preferences-constructive-interactions.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2025-03-05-human-preferences-constructive-interactions.md), [2025-06-01-hada-human-ai-agent-decision-alignment.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2025-06-01-hada-human-ai-agent-decision-alignment.md), [2025-10-02-predictive-preference-learning-human-interventions.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2025-10-02-predictive-preference-learning-human-interventions.md), [2025-10-27-tdflow-agentic-workflows-test-driven-development.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2025-10-27-tdflow-agentic-workflows-test-driven-development.md), [2026-01-07-agentic-rubrics-contextual-verifiers-swe-agents.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2026-01-07-agentic-rubrics-contextual-verifiers-swe-agents.md), [2026-03-12-context-engineering-corporate-multi-agent-architecture.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2026-03-12-context-engineering-corporate-multi-agent-architecture.md), [2026-04-13-agentforge-execution-grounded-multi-agent-swe.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2026-04-13-agentforge-execution-grounded-multi-agent-swe.md)
- Thought leaders: [2025-01-07-chip-huyen-agents.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Thought-Leaders/2025-01-07-chip-huyen-agents.md), [2025-06-23-harrison-chase-rise-of-context-engineering.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Thought-Leaders/2025-06-23-harrison-chase-rise-of-context-engineering.md), [2025-07-30-harrison-chase-deep-agents.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Thought-Leaders/2025-07-30-harrison-chase-deep-agents.md), [2026-03-01-simon-willison-how-coding-agents-work.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Thought-Leaders/2026-03-01-simon-willison-how-coding-agents-work.md), [2026-03-14-simon-willison-pragmatic-summit-agentic-engineering.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Thought-Leaders/2026-03-14-simon-willison-pragmatic-summit-agentic-engineering.md)
- Open-source exemplars: [langchain-deepagents-README.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Open-Source/langchain-deepagents-README.md), [microsoft-autogen-README.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Open-Source/microsoft-autogen-README.md), [openai-agents-python-README.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Open-Source/openai-agents-python-README.md), [openhands-README.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Open-Source/openhands-README.md)
- Industry exemplars: [2026-01-21-anthropic-claude-code-sandboxing.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-01-21-anthropic-claude-code-sandboxing.md), [2026-02-04-openai-unlocking-the-codex-harness.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-02-04-openai-unlocking-the-codex-harness.md), [2026-02-11-openai-harness-engineering.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-02-11-openai-harness-engineering.md), [2026-02-27-cognition-how-cognition-uses-devin-to-build-devin.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-02-27-cognition-how-cognition-uses-devin-to-build-devin.md), [2026-03-20-cognition-devin-can-now-schedule-devins.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-03-20-cognition-devin-can-now-schedule-devins.md), [anthropic-claude-code-product.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/anthropic-claude-code-product.md)
