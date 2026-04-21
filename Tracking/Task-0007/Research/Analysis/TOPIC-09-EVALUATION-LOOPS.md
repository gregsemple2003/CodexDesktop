# Topic 09 Evaluation Loops
## Frontier Landscape
The local corpus points to a clear frontier pattern: evaluation is not a one-shot score, it is a loop that turns traces, failure modes, and outcomes into the next round of context. The strongest recurring ideas are tracing, rubric-based verification, execution-grounded checks, and session or repo cleanup that feeds back into the next run. That is consistent across [TOPIC-01 input and event ingestion](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Analysis/TOPIC-01-INPUT-EVENT-INGESTION.md), [TOPIC-02 intervention record shape](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Analysis/TOPIC-02-INTERVENTION-RECORD-SHAPE.md), [TOPIC-04 deliberation record shape](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Analysis/TOPIC-04-DELIBERATION-RECORD-SHAPE.md), and [TOPIC-06 local memory and learning](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Analysis/TOPIC-06-LOCAL-MEMORY-AND-LEARNING.md).

What looks frontier-common in the source packet:
- failure-mode evaluation and explicit metrics, as in Chip Huyen’s breakdown of planning, tool, and efficiency failures in [2025-01-07-chip-huyen-agents.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Thought-Leaders/2025-01-07-chip-huyen-agents.md)
- trace-first agent engineering, as in OpenAI’s [Unlocking the Codex harness](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-02-04-openai-unlocking-the-codex-harness.md) and [Harness engineering](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-02-11-openai-harness-engineering.md)
- completed-session analysis and recurrence-aware improvement, as in Cognition’s [How Cognition Uses Devin to Build Devin](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-02-27-cognition-how-cognition-uses-devin-to-build-devin.md), [Devin can now Schedule Devins](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-03-20-cognition-devin-can-now-schedule-devins.md), and Devin Review / Session Insights pattern
- rubric or verifier layers that score candidate work without needing a full human reread, as in [Agentic Rubrics](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2026-01-07-agentic-rubrics-contextual-verifiers-swe-agents.md)
- execution-grounded iteration, as in [TDFlow](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2025-10-27-tdflow-agentic-workflows-test-driven-development.md) and [AgentForge](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2026-04-13-agentforge-execution-grounded-multi-agent-swe.md)
- trace and session primitives, as in [OpenAI Agents SDK](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Open-Source/openai-agents-python-README.md), [AutoGen](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Open-Source/microsoft-autogen-README.md), [Deep Agents](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Open-Source/langchain-deepagents-README.md), and [OpenHands](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Open-Source/openhands-README.md)

What looks over-engineered for us right now:
- a heavyweight evaluation council that tries to score every output against a broad values ontology before the local loop is stable
- large benchmark suites or leaderboard chasing before we know which repeated interventions actually matter in this repo
- multiple evaluation tiers, scorer agents, and dashboard layers that cannot be traced back to one concrete intervention or one concrete repo outcome
- machine-generated certainty about “improvement” when the only evidence is that the model sounded better

The local corpus is especially consistent on one point: the evaluator has to be grounded in a real workflow. OpenAI’s harness writing says repository knowledge should be the system of record, not a giant instruction blob; Anthropic’s sandboxing note says the boundaries must be real at the filesystem and network level; Chip Huyen says tool and planning failures should be measured directly; and TDFlow shows that test-first or execution-first loops work best when they are tightly scoped to the actual failure surface ([2026-02-11-openai-harness-engineering.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-02-11-openai-harness-engineering.md), [2026-01-21-anthropic-claude-code-sandboxing.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-01-21-anthropic-claude-code-sandboxing.md), [2025-01-07-chip-huyen-agents.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Thought-Leaders/2025-01-07-chip-huyen-agents.md), [2025-10-27-tdflow-agentic-workflows-test-driven-development.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2025-10-27-tdflow-agentic-workflows-test-driven-development.md)).

## Variables That Matter To Humans
The human does not mainly care about scores. They care about whether the system stops making them repeat the same correction and whether the proposals it emits actually make the repo easier to work in.

- Repetition burden: did the same correction stop recurring?
- Proposal usefulness: did the suggested task or instruction change actually help, or did it create more cleanup?
- Truthfulness: did the evaluation distinguish evidence, inference, and speculation?
- Local fit: did the evaluation respect repo-specific constraints instead of flattening them into a generic score?
- Review friction: did the output reduce human babysitting, or did it add another layer to review?
- Tone safety: did the system avoid blame drift and keep the framing repair-oriented?
- Scope control: did the evaluator keep a bad lesson local instead of turning it into a universal rule?
- Time-to-correction: did the loop find the problem quickly enough to matter?

The most important human variable is not raw accuracy, but whether the loop produces less repeated intervention over time. A good evaluation loop should let the human say, “that pattern is gone,” or “that proposal made the repo worse,” without digging through a separate analytics system to figure it out.

## Technique Likely To Work For Us
Use one append-only intervention ledger plus a daily delta report with three outcome labels: `helped`, `neutral`, and `harmed`.

Each observed intervention, proposal, or follow-up outcome gets one record. The next day, the report compares the last seven days against the previous seven days and answers a small set of questions: which patterns recurred, which proposals reduced recurrence, which proposals created rework, and which inferences were unsupported. Only patterns that show repeated reduction in human intervention should be promoted into repo-level guidance.

This is the simplest technique that still matches the frontier. It borrows the tracing and session-memory ideas from [Unlocking the Codex harness](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-02-04-openai-unlocking-the-codex-harness.md), the recurrence discipline from [Devin can now Schedule Devins](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-03-20-cognition-devin-can-now-schedule-devins.md), and the failure-mode discipline from [Agents](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Thought-Leaders/2025-01-07-chip-huyen-agents.md). It also fits the local task structure: the ledger can stay task-local until a pattern is stable enough to promote into [RESEARCH.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/RESEARCH.md), [TASK.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/TASK.md), or a repo contract file.

## Simple Structure That Maximizes The Variables
Keep the evaluation loop to four steps:

1. Capture the input or task proposal.
2. Record the immediate action and the observed outcome.
3. Check whether the same need recurred, whether the human had to repair the result, and whether the inference was supported.
4. Promote only the patterns that repeatedly help.

That structure maximizes the human variables because it keeps the loop legible and cheap to review. The report should be a short Markdown artifact with a recurrence table, a short list of helped or harmed patterns, and a small set of next proposals. The point is to make the human’s actual experience visible: less repetition, less rework, less blame, more durable improvement.

The most useful counters are simple:
- repeated intervention count
- accepted proposal count
- proposal rework count
- unsupported inference count
- tone or blame corrections
- local-constraint violations

That is enough to learn whether the system is actually getting better without turning evaluation into a separate platform project.

## Watchouts
- Do not optimize only for benchmark score or pass rate.
- Do not treat one successful fix as proof of a stable rule.
- Do not let the evaluator become a policy engine or a blame machine.
- Do not confuse a prettier report with a better loop.
- Do not add separate scorer hierarchies before the basic recurrence check works.
- Do not turn evaluation into bureaucracy that is harder to maintain than the thing being evaluated.
- Do not promote unsupported inferences just because they were repeated.
- Do not erase local human constraints in the name of consistency.

## Source Anchors
- [TASK.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/TASK.md)
- [RESEARCH.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/RESEARCH.md)
- [RESEARCH-ANALYSIS.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/RESEARCH-ANALYSIS.md)
- [2026-04-19 bootstrap analysis](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Analysis/2026-04-19-BOOTSTRAP-ANALYSIS.md)
- [2026-04-19 Jarvis model capture](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Conversations/2026-04-19-JARVIS-MODEL.md)
- [TOPIC-01 input and event ingestion](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Analysis/TOPIC-01-INPUT-EVENT-INGESTION.md)
- [TOPIC-02 intervention record shape](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Analysis/TOPIC-02-INTERVENTION-RECORD-SHAPE.md)
- [TOPIC-03 inference](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Analysis/TOPIC-03-inference.md)
- [TOPIC-04 deliberation record shape](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Analysis/TOPIC-04-DELIBERATION-RECORD-SHAPE.md)
- [TOPIC-06 local memory and learning](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Analysis/TOPIC-06-LOCAL-MEMORY-AND-LEARNING.md)
- [2025-01-07 Chip Huyen agents](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Thought-Leaders/2025-01-07-chip-huyen-agents.md)
- [2025-06-23 Harrison Chase context engineering](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Thought-Leaders/2025-06-23-harrison-chase-rise-of-context-engineering.md)
- [2025-07-30 Harrison Chase deep agents](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Thought-Leaders/2025-07-30-harrison-chase-deep-agents.md)
- [2026-03-01 Simon Willison coding agents](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Thought-Leaders/2026-03-01-simon-willison-how-coding-agents-work.md)
- [2026-03-14 Simon Willison agentic engineering](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Thought-Leaders/2026-03-14-simon-willison-pragmatic-summit-agentic-engineering.md)
- [2026-02-04 OpenAI unlocking the Codex harness](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-02-04-openai-unlocking-the-codex-harness.md)
- [2026-02-11 OpenAI harness engineering](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-02-11-openai-harness-engineering.md)
- [2026-01-21 Anthropic Claude Code sandboxing](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-01-21-anthropic-claude-code-sandboxing.md)
- [2026-02-27 Cognition Devin uses Devin](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-02-27-cognition-how-cognition-uses-devin-to-build-devin.md)
- [2026-03-20 Cognition Devin schedules Devins](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-03-20-cognition-devin-can-now-schedule-devins.md)
- [2026-01-07 Agentic Rubrics](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2026-01-07-agentic-rubrics-contextual-verifiers-swe-agents.md)
- [2025-10-27 TDFlow](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2025-10-27-tdflow-agentic-workflows-test-driven-development.md)
- [2026-04-13 AgentForge](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2026-04-13-agentforge-execution-grounded-multi-agent-swe.md)
- [OpenAI Agents SDK README](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Open-Source/openai-agents-python-README.md)
- [AutoGen README](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Open-Source/microsoft-autogen-README.md)
- [Deep Agents README](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Open-Source/langchain-deepagents-README.md)
- [OpenHands README](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Open-Source/openhands-README.md)
