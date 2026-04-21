# Topic 02 Intervention Record Shape
## Frontier Landscape
Frontier-common patterns across [Unlocking the Codex harness](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-02-04-openai-unlocking-the-codex-harness.md), [Harness engineering](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-02-11-openai-harness-engineering.md), [Beyond permission prompts](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-01-21-anthropic-claude-code-sandboxing.md), [Deep Agents](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Open-Source/langchain-deepagents-README.md), and [Agentic Rubrics](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2026-01-07-agentic-rubrics-contextual-verifiers-swe-agents.md) are consistent:
- harnesses are the real product boundary, not the prompt
- the system of record is the repo-local artifact trail, not chat memory
- context must be explicit, local, and retrievable
- verification is increasingly layered, but execution, rubrics, and review are all downstream of a clean record

Over-engineered for us, right now:
- multi-role councils and enterprise alignment pyramids before we have a stable event ledger
- execution-heavy repair loops before we know what intervention evidence we actually need
- broad memory systems that blur raw input, inference, and follow-up action

[Predictive Preference Learning from Human Interventions](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2025-10-02-predictive-preference-learning-human-interventions.md) is the useful caution here: one intervention can inform future behavior, but only if the record keeps the observed event separate from the inferred preference. [HADA](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2025-06-01-hada-human-ai-agent-decision-alignment.md) adds the governance lesson: alignment metadata matters, but the control layer is only valuable if it stays auditable and role-bound.

## Variables That Matter To Humans
- Did the system avoid asking for the same thing again?
- Is the record truthful about what was observed versus what was inferred?
- Did it respect local repo constraints and human authority?
- Can a human skim it and immediately understand what happened?
- Does it preserve the raw input, the interpretation, and the outcome separately?
- Can it drive a daily repo report and concrete follow-up proposal?
- Does it avoid blame, mind-reading, and inflated certainty?
- Does it stay small enough that people will actually maintain it?

## Technique Likely To Work For Us
Use a single append-only intervention ledger, one record per human input, with a strict split between event, interpretation, constraints, action, and outcome. That gives us the minimum shape that can support daily reporting, repo learning, and later automation without committing to a large agent stack too early.

This fits the lessons in [Unlocking the Codex harness](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-02-04-openai-unlocking-the-codex-harness.md) and [Harness engineering](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-02-11-openai-harness-engineering.md): keep the system of record small, versioned, and legible, then let other layers consume it.

## Simple Structure That Maximizes The Variables
1. Event
- source, repo, thread, author, timestamp, raw input, and artifact refs
2. Interpretation
- inferred need, confidence, evidence, alternative readings, and explicit versus implied versus speculative split
3. Constraints
- local repo rules, human desires, and any hard stop conditions that applied
4. Action
- what the system did or proposed, and whether it asked, answered, deferred, or delegated
5. Outcome
- what changed, whether the same need recurred, and what follow-up record or task was created

This structure keeps the high-value variables visible while avoiding an overbuilt ontology. It is also easy to render as Markdown, store as JSONL later if needed, and roll up into a daily report without losing the source trail.

## Watchouts
- Do not collapse raw evidence into a summary sentence.
- Do not turn a single intervention into a universal rule.
- Do not let confidence become a substitute for proof.
- Do not store private reasoning when a concise rationale is enough.
- Do not add subagents or memory hierarchies before the ledger shape is stable.
- Do not let the record become policy unless recurrence actually justifies promotion.

## Source Anchors
- [Task 0007 Research](/c:/Agent/CodexDashboard/Tracking/Task-0007/RESEARCH.md)
- [Task 0007 Bootstrap Analysis](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Analysis/2026-04-19-BOOTSTRAP-ANALYSIS.md)
- [Jarvis Model Capture](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Conversations/2026-04-19-JARVIS-MODEL.md)
- [OpenAI App Server](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-02-04-openai-unlocking-the-codex-harness.md)
- [OpenAI Harness Engineering](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-02-11-openai-harness-engineering.md)
- [Anthropic Sandboxing](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-01-21-anthropic-claude-code-sandboxing.md)
- [Deep Agents](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Open-Source/langchain-deepagents-README.md)
- [Agentic Rubrics](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2026-01-07-agentic-rubrics-contextual-verifiers-swe-agents.md)
