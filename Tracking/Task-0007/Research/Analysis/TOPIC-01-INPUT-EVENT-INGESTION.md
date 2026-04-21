# Topic 01 Input And Event Ingestion
## Frontier Landscape
The frontier pattern is not “one giant prompt.” It is a harness plus an event stream plus retrieval and memory. OpenAI’s Codex App Server makes this explicit with durable `thread` / `turn` / `item` primitives and bidirectional progress events, which is the right shape for capturing human inputs without flattening them into chat history ([Unlocking the Codex harness](</c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-02-04-openai-unlocking-the-codex-harness.md>), [Research/Analysis](</c:/Agent/CodexDashboard/Tracking/Task-0007/RESEARCH.md>)). Harrison Chase’s context-engineering framing says the model needs the right information, tools, and format, while Deep Agents adds the practical stack: planning, filesystem context, sub-agents, and context management ([The rise of "context engineering"](</c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Thought-Leaders/2025-06-23-harrison-chase-rise-of-context-engineering.md>), [Deep Agents](</c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Thought-Leaders/2025-07-30-harrison-chase-deep-agents.md>)).

What looks frontier-common is:
- canonical event capture from chat, PR comments, issue comments, and agent/tool events
- a stable schema for source, author, repo, thread, timestamp, and raw text
- repository-local context retrieval before action
- durable memory of prior interventions and corrections
- explicit approval/sandbox boundaries for risky actions

What looks over-engineered for us is:
- multi-layer corporate governance pyramids before we have a stable intake spine ([Context Engineering: From Prompts to Corporate Multi-Agent Architecture](</c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2026-03-12-context-engineering-corporate-multi-agent-architecture.md>))
- heavy stakeholder-agent architectures like HADA when the immediate problem is local repo understanding, not enterprise decision alignment ([HADA](</c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2025-06-01-hada-human-ai-agent-decision-alignment.md>))
- preference-horizon propagation into imagined future states before we have reliable event capture and record shape ([Predictive Preference Learning from Human Interventions](</c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2025-10-02-predictive-preference-learning-human-interventions.md>))
- protocol inflation, where ingestion becomes a new standards project instead of a simple normalized log

The local corpus points to a narrower lesson: input ingestion should make the repository more legible to future agent runs. OpenAI’s harness engineering post says to make repository knowledge the system of record, keep `AGENTS.md` as a map instead of an encyclopedia, and encode recurring corrections into mechanical checks ([Harness engineering: leveraging Codex in an agent-first world](</c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-02-11-openai-harness-engineering.md>)). That is the strongest frontier signal here.

## Variables That Matter To Humans
The human does not care first about event plumbing. They care about whether the system understood the need without forcing another round trip.

- Time to acknowledgment matters: the human wants the system to notice the input quickly and not bury it in a queue.
- Truthfulness matters: the system should say what it knows, what it infers, and what it does not know.
- Local constraint respect matters: repo-specific norms should beat generic “best practice” when they legitimately differ.
- Low-friction correction matters: a human should be able to override, scope, or forget a pattern without fighting the system.
- Privacy matters: raw human inputs and repo artifacts should be retained only as much as needed to support later inference.
- Auditability matters: when the system acts, the human should be able to trace which input caused it and why.
- Calmness matters: the system should not spam clarifications or turn every message into a dramatic workflow.
- Useful follow-through matters: the human wants proposals, not just logs.

These are not just UX preferences. They are the actual failure surface in the corpus: Codex harness engineering emphasizes human time and attention as the scarce resource, while Devin’s scheduling and session-insight features show that the user values recurrence handling, continuity, and action over raw transcript volume ([Devin can now Schedule Devins](</c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-03-20-cognition-devin-can-now-schedule-devins.md>), [How Cognition Uses Devin to Build Devin](</c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-02-27-cognition-how-cognition-uses-devin-to-build-devin.md>)).

## Technique Likely To Work For Us
Use one append-only `jsonl` event spine with thin source adapters, then derive higher-level records from it.

This is the simplest technique that still matches the frontier:
- each human or tool surface writes one normalized event
- each event stores only the source facts and minimal context pointers
- a deterministic reducer turns events into intervention records
- a daily synthesis step turns intervention records into repo proposals

Why this is the right first move:
- it matches the App Server’s event-oriented shape without requiring the whole Codex protocol stack ([Unlocking the Codex harness](</c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-02-04-openai-unlocking-the-codex-harness.md>))
- it keeps ingestion separate from interpretation, which preserves truthfulness
- it avoids committing early to a heavy memory system or agent swarm
- it leaves room for later replay, backfill, and promotion into repo-local docs

The concrete approach is:
- write every input into a single append-only `jsonl` ledger
- include `event_id`, `source`, `repo`, `thread_id`, `author`, `timestamp`, `raw_text`, and `local_context_refs`
- derive an `InterventionRecord` only after retrieval of local repo docs and recent history
- keep daily reporting as a reducer over the ledger, not as the source of truth itself

## Simple Structure That Maximizes The Variables
The structure should be four layers, not one:

- Raw event: the exact human input or tool-triggering signal.
- Normalized event: the same input with stable metadata and references to local context.
- Intervention record: the system’s bounded interpretation, confidence, evidence, local constraints, and proposed action.
- Daily repo report: the per-repo synthesis that turns repeated events into task proposals and instruction updates.

This structure maximizes the human variables because it keeps raw evidence intact, preserves inference honesty, and makes corrections cheap. It also fits the repo’s own doc layout: task-local conversations in `Research/Conversations/`, synthesis in `Research/Analysis/`, and durable source captures under `Research/Sources/` ([Research README](</c:/Agent/CodexDashboard/Tracking/Task-0007/Research/README.md>), [Task-0007](</c:/Agent/CodexDashboard/Tracking/Task-0007/TASK.md>)).

The minimal fields that matter most are:
- source and thread identity
- repo and subtree scope
- author and timestamp
- raw content
- local document refs
- inferred need hypothesis
- evidence and confidence
- local constraints applied
- proposed next action
- preventability score

This is enough to support daily report generation without requiring a more elaborate orchestration system up front.

## Watchouts
Do not collapse novelty into failure. The corpus is clear that “the system should have inferred it” is useful as a forcing function, but wrong if it erases genuinely new human intent ([2026-04-19 bootstrap analysis](</c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Analysis/2026-04-19-BOOTSTRAP-ANALYSIS.md>), [Task-0007](</c:/Agent/CodexDashboard/Tracking/Task-0007/TASK.md>)).

Do not treat a giant prompt or a giant instruction file as ingestion. OpenAI’s harness engineering writeup explicitly warns that one big `AGENTS.md` becomes stale and unreadable; the same failure applies to event capture if we overstuff it with policy text ([Harness engineering](</c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-02-11-openai-harness-engineering.md>)).

Do not let safety become ceremony. Sandboxing and network isolation are real boundary controls, but they should protect the ingestion spine, not replace it ([Making Claude Code more secure and autonomous with sandboxing](</c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-01-21-anthropic-claude-code-sandboxing.md>)).

Do not over-index on enterprise multi-agent formalisms before the local loop works. The paper stack around HADA, intent engineering, and corporate specification layers is useful as a ceiling, but it is not the right starting point for this repo’s immediate need ([HADA](</c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2025-06-01-hada-human-ai-agent-decision-alignment.md>), [Context Engineering: From Prompts to Corporate Multi-Agent Architecture](</c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2026-03-12-context-engineering-corporate-multi-agent-architecture.md>)).

Do not bury the human in raw records. The goal is fewer repeated interventions, not more transcript volume.

## Source Anchors
- [Task-0007](</c:/Agent/CodexDashboard/Tracking/Task-0007/TASK.md>)
- [Research README](</c:/Agent/CodexDashboard/Tracking/Task-0007/Research/README.md>)
- [RESEARCH.md](</c:/Agent/CodexDashboard/Tracking/Task-0007/RESEARCH.md>)
- [2026-04-19 bootstrap analysis](</c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Analysis/2026-04-19-BOOTSTRAP-ANALYSIS.md>)
- [2026-04-19 Jarvis model capture](</c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Conversations/2026-04-19-JARVIS-MODEL.md>)
- [OpenAI harness engineering](</c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-02-11-openai-harness-engineering.md>)
- [Unlocking the Codex harness](</c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-02-04-openai-unlocking-the-codex-harness.md>)
- [The rise of "context engineering"](</c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Thought-Leaders/2025-06-23-harrison-chase-rise-of-context-engineering.md>)
- [Deep Agents](</c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Thought-Leaders/2025-07-30-harrison-chase-deep-agents.md>)
- [How coding agents work](</c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Thought-Leaders/2026-03-01-simon-willison-how-coding-agents-work.md>)
- [Claude Code sandboxing](</c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-01-21-anthropic-claude-code-sandboxing.md>)
- [How Cognition Uses Devin to Build Devin](</c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-02-27-cognition-how-cognition-uses-devin-to-build-devin.md>)
- [Devin can now Schedule Devins](</c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-03-20-cognition-devin-can-now-schedule-devins.md>)
- [HADA](</c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2025-06-01-hada-human-ai-agent-decision-alignment.md>)
- [Predictive Preference Learning from Human Interventions](</c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2025-10-02-predictive-preference-learning-human-interventions.md>)
- [TDFlow](</c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2025-10-27-tdflow-agentic-workflows-test-driven-development.md>)
- [AgentForge](</c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2026-04-13-agentforge-execution-grounded-multi-agent-swe.md>)
