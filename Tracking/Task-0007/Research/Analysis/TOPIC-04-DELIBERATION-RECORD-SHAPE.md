# Topic 04 Deliberation Record Shape

The frontier pattern is not “write down everything the model thought.” The frontier pattern is to keep a durable, structured record of what the system saw, what it inferred, what it did not know, and what it decided to do next. In this task, that record needs to support honest inference about human intervention without becoming a fake chain-of-thought dump.

## Frontier Landscape

The local corpus shows a few frontier-common moves that are relevant here:

- `Unlocking the Codex harness` frames the conversation as `thread`, `turn`, and `item` primitives with persisted event history, which is the right baseline for any durable deliberation record. [Research/Sources/Industry/2026-02-04-openai-unlocking-the-codex-harness.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-02-04-openai-unlocking-the-codex-harness.md)
- `Harness engineering` makes repository knowledge the system of record and treats repeated human corrections as something to encode back into instructions, tests, and feedback loops. [Research/Sources/Industry/2026-02-11-openai-harness-engineering.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-02-11-openai-harness-engineering.md)
- `Context engineering` says the hard part is getting the right information, tools, and format into the model, not producing a giant prompt. [Research/Sources/Thought-Leaders/2025-06-23-harrison-chase-rise-of-context-engineering.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Thought-Leaders/2025-06-23-harrison-chase-rise-of-context-engineering.md)
- `Deep Agents` and the Simon Willison writeup both point toward a compact agent harness with planning, subagents, file system state, and a detailed prompt, rather than a monolithic brain dump. [Research/Sources/Thought-Leaders/2025-07-30-harrison-chase-deep-agents.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Thought-Leaders/2025-07-30-harrison-chase-deep-agents.md) [Research/Sources/Thought-Leaders/2026-03-01-simon-willison-how-coding-agents-work.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Thought-Leaders/2026-03-01-simon-willison-how-coding-agents-work.md)
- `Agentic Rubrics` and `TDFlow` both show that structured verifiers and workflow decomposition are more useful than freeform reflection when the goal is dependable action. [Research/Sources/Arxiv/2026-01-07-agentic-rubrics-contextual-verifiers-swe-agents.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2026-01-07-agentic-rubrics-contextual-verifiers-swe-agents.md) [Research/Sources/Arxiv/2025-10-27-tdflow-agentic-workflows-test-driven-development.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2025-10-27-tdflow-agentic-workflows-test-driven-development.md)
- `Predictive Preference Learning` shows a useful precedent for turning human intervention into structured future-facing signals instead of just the immediate correction. [Research/Sources/Arxiv/2025-10-02-predictive-preference-learning-human-interventions.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2025-10-02-predictive-preference-learning-human-interventions.md)

What looks frontier-common:

- event-sourced histories
- explicit turn/decision boundaries
- separate public output and private internal state
- evidence-led summaries instead of raw free association
- repeated corrections feeding back into repo instructions or tests
- verifier-style scoring of whether a proposed action is justified

What looks over-engineered for us:

- a universal human-values ontology
- a large multi-agent council for every input
- full hidden chain-of-thought storage as the record
- many bespoke record types before the core evidence shape is stable
- trying to solve inference, governance, memory, and reporting in one schema

## Variables That Matter To Humans

The human/operator/user cares about a small set of outcomes, not about the system’s internal theatrics:

- Did the system understand the need correctly?
- Did it preserve local constraints instead of flattening them into a global norm?
- Did it avoid overclaiming certainty?
- Did it choose the right next move: ask, draft, act, or defer?
- Did it reduce the need for the same correction again?
- Did it stay truthful when the need was only inferred?
- Did it keep the tone respectful and non-blaming?

The central human variable is trust. Trust here means the system can explain what it inferred, what it did not know, and why the route it chose was the least-bad option under local constraints.

## Technique Likely To Work For Us

Use a compact, evidence-first deliberation record with one record per human input and one selected route.

This is the simplest technique that still matches the frontier pattern:

- normalize the input into an event
- attach repo and local-context references
- generate a few candidate interpretations
- choose one route
- record why that route was chosen
- keep the record private by default
- emit a short human-facing summary separately

This is better for us than a sprawling deliberation transcript because it preserves the key learning signal without creating a maintenance burden or pretending to be more certain than the evidence allows.

## Simple Structure That Maximizes The Variables

Use a fixed record shape with a small, stable set of fields:

- `event_ref`
- `input_summary`
- `interpreted_need`
- `evidence`
- `local_constraints`
- `candidate_interpretations`
- `selected_route`
- `confidence`
- `what_was_missed`
- `how_the_system_could_have_inferred_this`
- `follow_up_action`

Why this structure works:

- `input_summary` keeps the record anchored to the real human event.
- `evidence` and `local_constraints` force honesty and local fit.
- `candidate_interpretations` preserves uncertainty without overfitting to one guess.
- `selected_route` keeps the record operational.
- `what_was_missed` turns each intervention into a learning signal.
- `how_the_system_could_have_inferred_this` preserves the terminal question that this task is built around.

If we need a public summary, it should be derived from this record, not a second reasoning system.

## Watchouts

- Do not turn the record into a hidden chain-of-thought archive.
- Do not encode “human values” as a single global preference function.
- Do not confuse one correction with a universal rule.
- Do not add many subtypes before the basic shape is useful.
- Do not let the record replace human override or local judgment.
- Do not store more detail than improves action, review, or learning.

## Source Anchors

- [TASK.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/TASK.md)
- [RESEARCH.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/RESEARCH.md)
- [RESEARCH-ANALYSIS.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/RESEARCH-ANALYSIS.md)
- [2026-04-19 Jarvis model capture](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Conversations/2026-04-19-JARVIS-MODEL.md)
- [2026-04-19 System Integration Spec 52](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Conversations/2026-04-19-SYSTEM-INTEGRATION-SPEC-52.md)
- [2026-04-19 System Integration Spec 54](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Conversations/2026-04-19-SYSTEM-INTEGRATION-SPEC-54.md)
- [2026-04-19 bootstrap analysis](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Analysis/2026-04-19-BOOTSTRAP-ANALYSIS.md)
- [Unlocking the Codex harness](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-02-04-openai-unlocking-the-codex-harness.md)
- [Harness engineering](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-02-11-openai-harness-engineering.md)
- [Claude Code sandboxing](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Industry/2026-01-21-anthropic-claude-code-sandboxing.md)
- [Context engineering](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Thought-Leaders/2025-06-23-harrison-chase-rise-of-context-engineering.md)
- [Deep Agents](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Thought-Leaders/2025-07-30-harrison-chase-deep-agents.md)
- [How coding agents work](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Thought-Leaders/2026-03-01-simon-willison-how-coding-agents-work.md)
- [Agentic Rubrics](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2026-01-07-agentic-rubrics-contextual-verifiers-swe-agents.md)
- [TDFlow](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2025-10-27-tdflow-agentic-workflows-test-driven-development.md)
- [Predictive Preference Learning from Human Interventions](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Sources/Arxiv/2025-10-02-predictive-preference-learning-human-interventions.md)
