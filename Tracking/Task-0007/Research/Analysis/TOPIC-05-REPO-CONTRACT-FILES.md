# Topic 05 Repo Contract Files

## Frontier Landscape

The local corpus points to a narrow frontier consensus: repo contract files work best when they separate purpose, procedure, and autonomy limits. `AGENTS.md` gives the machine the working rules for the repo; a purpose file such as `HUMAN-DESIRE.md` gives the machine the human reason the repo exists; and a small constraints layer captures the durable local exceptions that must not be inferred away. That shape is consistent with the task framing in [TASK.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/TASK.md), the research decomposition in [RESEARCH.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/RESEARCH.md), and the integration spec captures in [2026-04-19-SYSTEM-INTEGRATION-SPEC-52.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Conversations/2026-04-19-SYSTEM-INTEGRATION-SPEC-52.md) and [2026-04-19-SYSTEM-INTEGRATION-SPEC-54.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Conversations/2026-04-19-SYSTEM-INTEGRATION-SPEC-54.md).

What looks frontier-common is not a giant policy stack. It is a small instruction hierarchy, explicit context boundaries, sandboxed execution, and traceable verification. The OpenAI harness writeup emphasizes durable thread/turn/item structure and a stable app server boundary; the Agents SDK README emphasizes instructions, tools, guardrails, handoffs, sessions, and tracing; Anthropic’s sandboxing note emphasizes filesystem and network isolation; Simon Willison’s writing emphasizes codebase patterns, tests, and repeated correction as a learning signal. Those all support the same conclusion: contract files should be few, named for distinct jobs, and easy for both humans and agents to inspect. See [2026-02-04-openai-unlocking-the-codex-harness.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-02-04-openai-unlocking-the-codex-harness.md), [openai-agents-python-README.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Open-Source/openai-agents-python-README.md), [2026-01-21-anthropic-claude-code-sandboxing.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-01-21-anthropic-claude-code-sandboxing.md), and [2026-03-14-simon-willison-pragmatic-summit-agentic-engineering.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Thought-Leaders/2026-03-14-simon-willison-pragmatic-summit-agentic-engineering.md).

What looks over-engineered for us is the enterprise tendency to turn every nuance into another role, policy tier, or architecture layer before the base contract has proven useful. The context-engineering paper’s pyramid, HADA’s stakeholder architecture, and several multi-agent frameworks all point in the same direction: rich modeling of intent, values, and specifications can be useful, but it is easy to bury the human under a taxonomy. For this repo, that would be premature. We do not yet need a corporate architecture of contract files; we need a small and legible local contract. Relevant anchors are [2026-03-12-context-engineering-corporate-multi-agent-architecture.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2026-03-12-context-engineering-corporate-multi-agent-architecture.md), [2025-06-01-hada-human-ai-agent-decision-alignment.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2025-06-01-hada-human-ai-agent-decision-alignment.md), [microsoft-autogen-README.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Open-Source/microsoft-autogen-README.md), and [openhands-README.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Open-Source/openhands-README.md).

## Variables That Matter To Humans

The human does not mainly care about taxonomy. The human cares about not having to repeat the same correction, not having their local constraints flattened into global defaults, and not having the system pretend certainty where only inference exists. The human also cares that the contract is reviewable, short enough to maintain, and specific enough to shape behavior in this repo rather than in some abstract “best practice” universe.

So the important variables are:

- whether the repo can infer the right thing without asking again
- whether explicit human desire stays separate from inferred desire
- whether hard local constraints are honored before generic optimization
- whether the file set stays small enough to maintain
- whether the precedence order is obvious when files conflict
- whether accepted corrections become durable instead of fading back into chat history
- whether the repo can explain itself truthfully when it does not know

That matches the task-level framing in [RESEARCH-ANALYSIS.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/RESEARCH-ANALYSIS.md) and the conversation captures in [2026-04-19-JARVIS-MODEL.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Conversations/2026-04-19-JARVIS-MODEL.md). It also fits the preference and intervention literature in [2025-10-02-PREDICTIVE-PREFERENCE-LEARNING-HUMAN-INTERVENTIONS.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2025-10-02-PREDICTIVE-PREFERENCE-LEARNING-HUMAN-INTERVENTIONS.md) and [2025-03-05-HUMAN-PREFERENCES-CONSTRUCTIVE-INTERACTIONS.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2025-03-05-HUMAN-PREFERENCES-CONSTRUCTIVE-INTERACTIONS.md), which both treat human intervention and human preference as signals that should be made more legible, not buried.

## Technique Likely To Work For Us

Use a three-file contract stack with an explicit promotion gate:

`AGENTS.md` says how to work here.
`HUMAN-DESIRE.md` says why this repo exists and what outcomes humans care about.
`HUMAN-CONSTRAINTS.md` exists only for durable hard overrides that actually recur.

All one-off corrections stay in task-local research artifacts until they repeat enough to justify promotion. That gives the system a simple rule for what is policy, what is purpose, and what is just evidence. It also prevents the contract layer from becoming a dumping ground for every observation the model ever made.

This is the concrete technique that best fits the corpus: keep the contract files minimal, keep the precedence order explicit, and promote only repeated, high-confidence corrections into repo-level policy. The frontier sources support the general pattern of instruction layering, tracing, and sandboxed autonomy; the task docs support the need to keep local human constraints explicit and to separate explicit, implied, and speculative claims.

## Simple Structure That Maximizes The Variables

The simplest structure that still maximizes the human variables is:

1. Root `AGENTS.md` for procedure, tooling, and do-not rules.
2. Root `HUMAN-DESIRE.md` for core need, desired outcomes, anti-goals, and examples of good and bad inference.
3. Optional root `HUMAN-CONSTRAINTS.md` for durable local exceptions that should override generic defaults.
4. Task-local research artifacts for the evidence trail, observed corrections, and unresolved inference questions.

That structure works because it keeps purpose and procedure separate, which is the main thing humans need when they want the system to stop guessing. It also keeps the repo maintainable: one file for how to operate, one file for why the repo exists, and one file only if a hard constraint truly needs its own home.

The right `HUMAN-DESIRE.md` should stay short and stable. It should carry the repo’s core need, the values the human actually cares about, the local constraints that matter, and a few concrete examples of what counts as a good inference versus an overreach. The research corpus suggests that this beats a larger architecture because it is easy to inspect, easy to diff, and hard to misuse.

## Watchouts

- Do not turn `HUMAN-DESIRE.md` into a values manifesto with no operational effect.
- Do not use `HUMAN-CONSTRAINTS.md` unless there is a durable local override worth preserving.
- Do not promote speculative inference into repo policy.
- Do not add files just because a frontier system had more files.
- Do not confuse an evidence trail with a contract.
- Do not let local corrections vanish into chat history instead of becoming durable when repeated.
- Do not let a generic global policy outrank an explicit local desire unless safety, law, or privacy requires it.

## Source Anchors

- Task framing: [TASK.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/TASK.md), [RESEARCH.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/RESEARCH.md), [RESEARCH-PLAN.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/RESEARCH-PLAN.md), [RESEARCH-ANALYSIS.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/RESEARCH-ANALYSIS.md), [PLAN.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/PLAN.md), [HANDOFF.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/HANDOFF.md)
- Conversation captures: [2026-04-19-JARVIS-MODEL.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Conversations/2026-04-19-JARVIS-MODEL.md), [2026-04-19-SYSTEM-INTEGRATION-SPEC-52.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Conversations/2026-04-19-SYSTEM-INTEGRATION-SPEC-52.md), [2026-04-19-SYSTEM-INTEGRATION-SPEC-54.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Conversations/2026-04-19-SYSTEM-INTEGRATION-SPEC-54.md)
- Thought leadership: [2025-01-07-chip-huyen-agents.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Thought-Leaders/2025-01-07-chip-huyen-agents.md), [2025-06-23-harrison-chase-rise-of-context-engineering.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Thought-Leaders/2025-06-23-harrison-chase-rise-of-context-engineering.md), [2025-07-30-harrison-chase-deep-agents.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Thought-Leaders/2025-07-30-harrison-chase-deep-agents.md), [2026-03-01-simon-willison-how-coding-agents-work.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Thought-Leaders/2026-03-01-simon-willison-how-coding-agents-work.md), [2026-03-14-simon-willison-pragmatic-summit-agentic-engineering.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Thought-Leaders/2026-03-14-simon-willison-pragmatic-summit-agentic-engineering.md)
- Open-source exemplars: [openai-agents-python-README.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Open-Source/openai-agents-python-README.md), [langchain-deepagents-README.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Open-Source/langchain-deepagents-README.md), [microsoft-autogen-README.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Open-Source/microsoft-autogen-README.md), [openhands-README.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Open-Source/openhands-README.md)
- Industry and harness sources: [2026-02-04-openai-unlocking-the-codex-harness.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-02-04-openai-unlocking-the-codex-harness.md), [2026-01-21-anthropic-claude-code-sandboxing.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-01-21-anthropic-claude-code-sandboxing.md), [anthropic-claude-code-product.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/anthropic-claude-code-product.md), [2026-02-27-cognition-how-cognition-uses-devin-to-build-devin.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-02-27-cognition-how-cognition-uses-devin-to-build-devin.md), [2026-03-20-cognition-devin-can-now-schedule-devins.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-03-20-cognition-devin-can-now-schedule-devins.md)
- ArXiv sources: [2025-03-05-human-preferences-constructive-interactions.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2025-03-05-human-preferences-constructive-interactions.md), [2025-06-01-hada-human-ai-agent-decision-alignment.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2025-06-01-hada-human-ai-agent-decision-alignment.md), [2025-10-02-predictive-preference-learning-human-interventions.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2025-10-02-predictive-preference-learning-human-interventions.md), [2025-10-27-tdflow-agentic-workflows-test-driven-development.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2025-10-27-tdflow-agentic-workflows-test-driven-development.md), [2026-01-07-agentic-rubrics-contextual-verifiers-swe-agents.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2026-01-07-agentic-rubrics-contextual-verifiers-swe-agents.md), [2026-03-12-context-engineering-corporate-multi-agent-architecture.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2026-03-12-context-engineering-corporate-multi-agent-architecture.md), [2026-04-13-agentforge-execution-grounded-multi-agent-swe.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2026-04-13-agentforge-execution-grounded-multi-agent-swe.md)
