# Topic 08 Autonomy And Safety Policy
## Frontier Landscape
The local corpus points to a fairly stable frontier pattern: the winning systems do not chase “maximum autonomy” in the abstract. They build bounded autonomy inside a harness, with explicit tools, explicit session history, explicit approval points, and explicit traces. That shows up in [Unlocking the Codex harness](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-02-04-openai-unlocking-the-codex-harness.md), [Harness engineering](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-02-11-openai-harness-engineering.md), [OpenAI Agents SDK](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Open-Source/openai-agents-python-README.md), [Deep Agents](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Open-Source/langchain-deepagents-README.md), and [Claude Code sandboxing](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-01-21-anthropic-claude-code-sandboxing.md).

What looks frontier-common is:
- sandboxed execution with filesystem and network boundaries
- human-in-the-loop approval for risky actions
- sessions, threads, and tracing so the system can be resumed and audited
- explicit subagents or role splits instead of one opaque controller
- repository knowledge as the system of record, not chat memory
- recurring audits, cleanup, and verification loops

That pattern is reinforced by [How coding agents work](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Thought-Leaders/2026-03-01-simon-willison-how-coding-agents-work.md), [Harness engineering](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-02-11-openai-harness-engineering.md), [Devin can now Schedule Devins](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-03-20-cognition-devin-can-now-schedule-devins.md), and [Agentic Rubrics](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2026-01-07-agentic-rubrics-contextual-verifiers-swe-agents.md). The common move is not “let the model decide everything.” It is “make the next action legible, bounded, and reviewable.”

What looks over-engineered for us is the large-governance end of the frontier: multi-layer corporate pyramids, broad intent ontologies, stakeholder councils, and universal policy stacks before the local autonomy loop is trustworthy. [Context Engineering: From Prompts to Corporate Multi-Agent Architecture](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2026-03-12-context-engineering-corporate-multi-agent-architecture.md) and [HADA](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2025-06-01-hada-human-ai-agent-decision-alignment.md) are useful as ceilings, but they are too heavy for this repo if we have not yet stabilized the local policy boundary.

## Variables That Matter To Humans
The human/operator/user does not mainly care about how many roles exist in the agent stack. They care about whether the system is safe to trust, whether it stays out of the way when it should, and whether it stops forcing the same correction over and over.
The policy also has to be truthful, compassionate in framing, and tolerant of legitimate local variation, or the human will either stop trusting it or stop using it.

- No surprise writes, deletes, or external side effects.
- No prompt-injection or data exfiltration surprises.
- Low approval fatigue for routine work.
- Fast execution for low-risk tasks without babysitting.
- Clear escalation when a task is ambiguous, risky, or meaning-sensitive.
- Local repo and team norms respected before generic defaults.
- Reversible changes and a visible recovery path.
- Honest uncertainty when the agent is inferring rather than observing.
- Outputs that feel faithful to the repo’s actual goals, not generic automation.
- Privacy protection for sensitive context and user data.

These are the real failure surfaces in the corpus. Anthropic’s sandboxing post is about reducing permission prompts without losing control. Simon Willison’s “lethal trifecta” framing is about preventing access to private data plus malicious instructions plus an exfiltration path. OpenAI’s harness post is about keeping repository knowledge and runtime state legible. Cognition’s Devin posts are about using scheduling, reviews, and session insight to cut repeated human intervention. The human variables are therefore trust, interruption cost, local fit, and recoverability.

## Technique Likely To Work For Us
Use a single risk-tiered autonomy gate with sandbox-first execution.

This is the concrete approach that fits the corpus and stays simple:
- Treat read-only work as the default.
- Allow low-risk work to run autonomously in a sandbox and open a PR when it is done.
- Require explicit human approval for medium-risk and meaning-sensitive work.
- Block high-risk work from automatic delegation entirely.

The practical cutoff is not “what the model feels confident about.” The cutoff is the action class:
- Low risk: docs, tests, lint, formatting, small reversible refactors, mechanical cleanup.
- Medium risk: behavior changes, dependency updates, API shape changes, migrations, performance-sensitive rewrites, privacy-sensitive or security-sensitive edits.
- High risk: mission or policy changes, destructive data operations, credential handling, legal or interpersonal judgments.

This approach works because it separates autonomy from authority. The agent can move quickly inside a bounded sandbox, but it cannot silently expand the scope of its own power. That matches [Claude Code sandboxing](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-01-21-anthropic-claude-code-sandboxing.md), [OpenAI Agents SDK](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Open-Source/openai-agents-python-README.md), and [OpenHands](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Open-Source/openhands-README.md) better than it matches the large governance papers.

## Simple Structure That Maximizes The Variables
Keep the policy structure to one small decision table plus one exception log.

1. Action class
- What kind of work is this: read, draft, low-risk write, medium-risk write, or high-risk write?
2. Default route
- Can the system do it directly, propose it, or must it stop?
3. Required boundary
- Sandbox, approval, or human review.
4. Recovery note
- How the change is rolled back or corrected if the first attempt is wrong.

That structure maximizes the human variables because it makes the routing rule obvious, keeps the number of autonomy states small, and makes exceptions reviewable. It also keeps the policy honest: if a task cannot be classified quickly, the system should not invent a route. It should escalate.

The same structure should apply across repo work, daily reports, and recurring routines. If the system can schedule itself, as in [Devin can now Schedule Devins](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-03-20-cognition-devin-can-now-schedule-devins.md), then the schedule itself must still be bounded by the same action classes and approval rules. Recurrence is not a license to become less careful.

## Watchouts
- Do not confuse sandboxing with safety if network egress or secret access is still open.
- Do not let “approval fatigue” become an excuse to remove meaningful review from risky actions.
- Do not let recurring routines drift into silent mission expansion.
- Do not turn a local policy into a universal policy when the repo explicitly needs local variation.
- Do not build a large autonomy ontology before the default route works reliably.
- Do not use high autonomy to hide low traceability.
- Do not ask humans to approve every trivial action; that just moves the burden instead of reducing it.

## Source Anchors
- [TASK.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/TASK.md)
- [RESEARCH.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/RESEARCH.md)
- [RESEARCH-ANALYSIS.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/RESEARCH-ANALYSIS.md)
- [2026-04-19 bootstrap analysis](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Analysis/2026-04-19-BOOTSTRAP-ANALYSIS.md)
- [2026-04-19 Jarvis model capture](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Conversations/2026-04-19-JARVIS-MODEL.md)
- [Unlocking the Codex harness](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-02-04-openai-unlocking-the-codex-harness.md)
- [Harness engineering](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-02-11-openai-harness-engineering.md)
- [Claude Code sandboxing](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-01-21-anthropic-claude-code-sandboxing.md)
- [Claude Code product page](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/anthropic-claude-code-product.md)
- [OpenAI Agents SDK](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Open-Source/openai-agents-python-README.md)
- [Deep Agents](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Open-Source/langchain-deepagents-README.md)
- [OpenHands](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Open-Source/openhands-README.md)
- [AutoGen](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Open-Source/microsoft-autogen-README.md)
- [How coding agents work](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Thought-Leaders/2026-03-01-simon-willison-how-coding-agents-work.md)
- [Pragmatic Summit agentic engineering](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Thought-Leaders/2026-03-14-simon-willison-pragmatic-summit-agentic-engineering.md)
- [The rise of "context engineering"](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Thought-Leaders/2025-06-23-harrison-chase-rise-of-context-engineering.md)
- [Deep Agents](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Thought-Leaders/2025-07-30-harrison-chase-deep-agents.md)
- [Agentic Rubrics](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2026-01-07-agentic-rubrics-contextual-verifiers-swe-agents.md)
- [TDFlow](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2025-10-27-tdflow-agentic-workflows-test-driven-development.md)
- [AgentForge](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2026-04-13-agentforge-execution-grounded-multi-agent-swe.md)
- [Context Engineering: From Prompts to Corporate Multi-Agent Architecture](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2026-03-12-context-engineering-corporate-multi-agent-architecture.md)
- [HADA](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2025-06-01-hada-human-ai-agent-decision-alignment.md)
