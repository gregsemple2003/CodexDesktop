# Topic 07 Daily Repo Report Design
## Frontier Landscape
The local corpus converges on a few durable ideas. The first is that the report should be driven by repository-local evidence, not by chat memory or a giant instruction blob. OpenAI's harness writeups emphasize that repository knowledge is the system of record, that `AGENTS.md` is more like a map than an encyclopedia, and that durable thread/turn/item history belongs in explicit artifacts rather than in a model's head ([Harness engineering](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-02-11-openai-harness-engineering.md), [Unlocking the Codex harness](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-02-04-openai-unlocking-the-codex-harness.md)). The same pattern shows up in the SDK and harness docs from OpenAI, LangChain, and Simon Willison: keep the context explicit, keep the tools visible, keep the loop inspectable, and keep the artifact small enough to maintain ([OpenAI Agents SDK](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Open-Source/openai-agents-python-README.md), [Deep Agents](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Open-Source/langchain-deepagents-README.md), [How coding agents work](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Thought-Leaders/2026-03-01-simon-willison-how-coding-agents-work.md), [The rise of context engineering](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Thought-Leaders/2025-06-23-harrison-chase-rise-of-context-engineering.md)).

The second common thread is that recurring work should be made repeatable, not merely remembered. Cognition's Devin posts describe daily audits, session insights, scheduled recurring work, and backlog cleanup as a natural use case for agentic systems ([How Cognition Uses Devin to Build Devin](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-02-27-cognition-how-cognition-uses-devin-to-build-devin.md), [Devin can now Schedule Devins](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-03-20-cognition-devin-can-now-schedule-devins.md)). That fits the task framing here: the daily repo report is not just a status note, it is a recurring conversion layer from repeated intervention into durable next action. The relevant research is also consistent with that direction: interventions can be propagated into future behavior if the record keeps the observed event separate from the inferred preference, and verification should be grounded in repository context rather than generic heuristics ([Predictive Preference Learning from Human Interventions](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2025-10-02-predictive-preference-learning-human-interventions.md), [Agentic Rubrics as Contextual Verifiers for SWE Agents](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2026-01-07-agentic-rubrics-contextual-verifiers-swe-agents.md)).

What looks frontier-common:
- Explicit repo-local artifacts, not implicit memory.
- Small, stable map files and short instruction surfaces.
- Repeated corrections turned into durable repo learning.
- Daily or scheduled audits that generate concrete follow-up.
- Grounded verification and provenance for claims.
- Human-in-the-loop escalation when uncertainty or risk is real.

What looks over-engineered for us:
- Multi-layer enterprise pyramids before the basic report loop works.
- Many specialized agent roles before the report shape is stable.
- Broad memory systems that blur evidence, interpretation, and policy.
- Huge contract stacks that are more ceremony than leverage.
- Complex governance layers that bury the human under taxonomy.

The heavy architecture papers are useful as warning signs, not templates. `HADA` and the context-engineering pyramid are the clearest examples: they show how quickly a system can turn context, intent, and specification into a deep enterprise stack, but that is more machinery than this task needs right now ([HADA](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2025-06-01-hada-human-ai-agent-decision-alignment.md), [Context Engineering: From Prompts to Corporate Multi-Agent Architecture](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2026-03-12-context-engineering-corporate-multi-agent-architecture.md)). For this repo, the frontier value is in the discipline of explicit context and durable recurrence, not in the size of the architecture.

## Variables That Matter To Humans
- Whether the report stops the same correction from recurring.
- Whether the report is truthful about observation, inference, and uncertainty.
- Whether local repo constraints stay intact instead of getting flattened into global defaults.
- Whether a human can skim the report and know what matters in under a minute.
- Whether the top next steps are concrete enough to act on.
- Whether the report avoids blame, mind-reading, and inflated certainty.
- Whether the output is small enough that people will actually keep using it.
- Whether the report advances the repo's real purpose, not just system metrics.

These are the human variables that keep showing up across the corpus. `RESEARCH-ANALYSIS.md` and `2026-04-19-JARVIS-MODEL.md` both insist on explicit separation between explicit, implied, and speculative claims, because the human cares about honesty more than narrative polish ([RESEARCH-ANALYSIS.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/RESEARCH-ANALYSIS.md), [2026-04-19 Jarvis Model Capture](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Conversations/2026-04-19-JARVIS-MODEL.md)). `TOPIC-05-REPO-CONTRACT-FILES.md` also points to the same human concerns from a different angle: the file set must stay small, precedence must stay obvious, and accepted corrections must become durable instead of fading away ([TOPIC-05 Repo Contract Files](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Analysis/TOPIC-05-REPO-CONTRACT-FILES.md)).

## Technique Likely To Work For Us
Use a single daily Markdown report per repo, generated from a rolling intervention ledger. One record per human intervention, with the report grouping those records by repo, theme, and recurrence. That gives us one concrete technique instead of a menu: write the daily report as a deterministic rollup of explicit evidence plus a small amount of bounded inference.

This is the right size for this repo. It matches the harness lesson that the system of record should stay small and legible, and it fits the daily-audit pattern from Devin without importing its heavier product stack. It also leaves room for later automation, because the report can be backed by structured records without requiring a big agent council or a corporate governance layer before the basics work.

The report should answer three questions only:
- What did the system have to learn from today's interventions?
- What repo proposal would reduce repetition next time?
- What local constraint or value changed the shape of that proposal?

That is enough to make the report useful. Anything broader risks turning it into a dashboard of interesting facts instead of a decision artifact.

## Simple Structure That Maximizes The Variables
Use a fixed one-page structure:

1. Repo header
- repo name, date, source window, and a short confidence note
2. Core need
- one sentence stating the repo's core need
- explicit if it came from `HUMAN-DESIRE.md`, otherwise marked as inferred
3. Observed interventions
- 3 to 5 bullets
- each bullet names the source, the human ask, and the system miss
4. Task proposals
- 3 to 5 ranked proposals
- each proposal includes impact, evidence, local constraints, and next step
5. Instruction updates
- only when a repeated pattern justifies a repo-level change
6. Open questions
- at most 2
- only when the report cannot honestly close the loop

This structure maximizes the human variables because it makes the important distinctions visible without making the user page through a lot of machinery. It also keeps the report operationally simple: a small fixed schema, a small number of proposals, and a small number of places where uncertainty is allowed to appear.

The strongest constraint on structure is that the report should stay readable as a report, not become a research essay. The daily artifact should make the next action obvious. If it takes more than a quick skim to find the repo need, the repeated interventions, and the ranked proposals, the structure is too heavy.

## Watchouts
- Do not collapse evidence and inference into the same sentence.
- Do not turn one intervention into a universal rule.
- Do not let a report become a taxonomy competition.
- Do not let the number of sections grow faster than the value they create.
- Do not over-automate before the report shape is stable.
- Do not substitute a values slogan for a concrete next step.
- Do not make the report so long that no one reads it daily.

## Source Anchors
- Task framing: [TASK.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/TASK.md), [PLAN.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/PLAN.md), [HANDOFF.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/HANDOFF.md), [RESEARCH.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/RESEARCH.md), [RESEARCH-PLAN.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/RESEARCH-PLAN.md), [RESEARCH-ANALYSIS.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/RESEARCH-ANALYSIS.md)
- Conversation captures: [2026-04-19-JARVIS-MODEL.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Conversations/2026-04-19-JARVIS-MODEL.md), [2026-04-19-SYSTEM-INTEGRATION-SPEC-52.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Conversations/2026-04-19-SYSTEM-INTEGRATION-SPEC-52.md), [2026-04-19-SYSTEM-INTEGRATION-SPEC-54.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Conversations/2026-04-19-SYSTEM-INTEGRATION-SPEC-54.md)
- Analysis notes: [2026-04-19-BOOTSTRAP-ANALYSIS.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Analysis/2026-04-19-BOOTSTRAP-ANALYSIS.md), [TOPIC-02-INTERVENTION-RECORD-SHAPE.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Analysis/TOPIC-02-INTERVENTION-RECORD-SHAPE.md), [TOPIC-05-REPO-CONTRACT-FILES.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Analysis/TOPIC-05-REPO-CONTRACT-FILES.md)
- Open-source exemplars: [openai-agents-python-README.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Open-Source/openai-agents-python-README.md), [langchain-deepagents-README.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Open-Source/langchain-deepagents-README.md), [microsoft-autogen-README.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Open-Source/microsoft-autogen-README.md), [openhands-README.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Open-Source/openhands-README.md)
- Thought leaders: [2025-01-07-chip-huyen-agents.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Thought-Leaders/2025-01-07-chip-huyen-agents.md), [2025-06-23-harrison-chase-rise-of-context-engineering.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Thought-Leaders/2025-06-23-harrison-chase-rise-of-context-engineering.md), [2025-07-30-harrison-chase-deep-agents.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Thought-Leaders/2025-07-30-harrison-chase-deep-agents.md), [2026-03-01-simon-willison-how-coding-agents-work.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Thought-Leaders/2026-03-01-simon-willison-how-coding-agents-work.md), [2026-03-14-simon-willison-pragmatic-summit-agentic-engineering.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Thought-Leaders/2026-03-14-simon-willison-pragmatic-summit-agentic-engineering.md)
- Industry captures: [2026-02-04-openai-unlocking-the-codex-harness.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-02-04-openai-unlocking-the-codex-harness.md), [2026-02-11-openai-harness-engineering.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-02-11-openai-harness-engineering.md), [2026-01-21-anthropic-claude-code-sandboxing.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-01-21-anthropic-claude-code-sandboxing.md), [anthropic-claude-code-product.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/anthropic-claude-code-product.md), [2026-02-27-cognition-how-cognition-uses-devin-to-build-devin.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-02-27-cognition-how-cognition-uses-devin-to-build-devin.md), [2026-03-20-cognition-devin-can-now-schedule-devins.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-03-20-cognition-devin-can-now-schedule-devins.md)
- ArXiv sources: [2025-03-05-human-preferences-constructive-interactions.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2025-03-05-human-preferences-constructive-interactions.md), [2025-06-01-hada-human-ai-agent-decision-alignment.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2025-06-01-hada-human-ai-agent-decision-alignment.md), [2025-10-02-predictive-preference-learning-human-interventions.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2025-10-02-predictive-preference-learning-human-interventions.md), [2025-10-27-tdflow-agentic-workflows-test-driven-development.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2025-10-27-tdflow-agentic-workflows-test-driven-development.md), [2026-01-07-agentic-rubrics-contextual-verifiers-swe-agents.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2026-01-07-agentic-rubrics-contextual-verifiers-swe-agents.md), [2026-03-12-context-engineering-corporate-multi-agent-architecture.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2026-03-12-context-engineering-corporate-multi-agent-architecture.md), [2026-04-13-agentforge-execution-grounded-multi-agent-swe.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2026-04-13-agentforge-execution-grounded-multi-agent-swe.md)
