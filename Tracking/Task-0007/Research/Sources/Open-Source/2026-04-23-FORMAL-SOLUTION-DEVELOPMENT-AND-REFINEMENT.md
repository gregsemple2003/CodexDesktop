# Formal Solution Development And Refinement: Open-Source Packet

Retrieved: 2026-04-23

This packet captures open-source frameworks and docs that make planning, delegation, guardrails, and refinement first-class workflow objects.

## 01. OpenAI Agents SDK README
- Observed: 2026-04-23
- URL: https://github.com/openai/openai-agents-python
- Formal-solution signal: the core abstractions are explicit agents, tools, guardrails, handoffs, sessions, and tracing rather than one undifferentiated prompt loop.
- Task-0007 relevance: supports formalizing solution development as a typed artifact flow with explicit boundaries between design, delegation, and validation.

## 02. OpenAI Agents SDK: Sandbox Agents
- Observed: 2026-04-23
- URL: https://openai.github.io/openai-agents-python/sandbox_agents/
- Formal-solution signal: long-horizon work is anchored in a manifest-backed workspace and controlled execution environment instead of being treated as pure conversation.
- Task-0007 relevance: solution refinement should name the execution surface and verification environment, not just the desired intent.

## 03. OpenAI Agents SDK: Handoffs
- Observed: 2026-04-23
- URL: https://openai.github.io/openai-agents-python/handoffs/
- Formal-solution signal: delegation is modeled as an explicit handoff with bounded context rather than a fuzzy continuation.
- Task-0007 relevance: supports a durable contract where `2B` hands a frozen winner set downstream to `2C` instead of leaving ownership ambiguous.

## 04. OpenAI Agents SDK: Guardrails
- Observed: 2026-04-23
- URL: https://openai.github.io/openai-agents-python/guardrails/
- Formal-solution signal: validation happens at explicit boundaries with success and failure handling built into the workflow model.
- Task-0007 relevance: `2B` should act like a guardrail stage with explicit failure conditions, not a light editorial pass.

## 05. OpenHands README
- Observed: 2026-04-23
- URL: https://github.com/All-Hands-AI/OpenHands
- Formal-solution signal: agentic software development is treated as a reproducible environment-plus-workflow problem with sandboxes, task state, and reviewable actions.
- Task-0007 relevance: supports explicit artifact ownership and replayable refinement loops when solutions are contested or incomplete.

## 06. AutoGen: Magentic-One
- Observed: 2026-04-23
- URL: https://microsoft.github.io/autogen/stable/user-guide/agentchat-user-guide/magentic-one.html
- Formal-solution signal: the orchestrator owns progress tracking, replanning, and specialist coordination while the team works through bounded interfaces.
- Task-0007 relevance: reinforces that solution refinement needs an owner and an explicit point where re-plan is allowed.

## 07. LangChain Deep Agents README
- Observed: 2026-04-23
- URL: https://github.com/langchain-ai/deepagents
- Formal-solution signal: emphasizes decomposing deep work into structured subproblems with context management and longer-horizon execution support.
- Task-0007 relevance: supports splitting the Dream loop by function instead of asking one stage to both invent and validate the same solution.

## 08. GitHub Spec Kit README
- Observed: 2026-04-23
- URL: https://github.com/github/spec-kit
- Formal-solution signal: makes `/specify`, `/plan`, and `/tasks` separate artifacts so the specification remains the system of record while downstream work stays derived.
- Task-0007 relevance: this is the clearest open-source precedent for treating `SOLUTION-DESIGN.md` as the living source of truth and task drafts as downstream products.

## 09. TDAD Repository
- Observed: 2026-04-23
- URL: https://github.com/pepealonso95/TDAD
- Formal-solution signal: publishes a concrete regression-aware development loop with graph-based impact analysis, logs, and benchmark outputs.
- Task-0007 relevance: winner refinement should name the concrete impact surface and regression checks each option must survive.
