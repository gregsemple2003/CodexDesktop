# Formal Solution Development And Refinement: Preprint Packet

Retrieved: 2026-04-23

This packet captures recent primary preprints that treat software-agent solution design as an explicit artifact with structured refinement, verification, and recovery loops.

## 01. AgentForge: Execution-Grounded Multi-Agent LLM Framework for Autonomous Software Engineering
- Published: 2026-04-13
- URL: https://arxiv.org/abs/2604.13120
- Formal-solution signal: frames software engineering as an iterative decision process over repository states and makes sandboxed execution a mandatory gate before propagation.
- Task-0007 relevance: supports keeping solution artifacts auditable and refusing to promote a design choice until it survives a grounded verifier, not just persuasive prose.

## 02. TDAD: Test-Driven Agentic Development - Reducing Code Regressions in AI Coding Agents via Graph-Based Impact Analysis
- Published: 2026-03-18
- URL: https://arxiv.org/abs/2603.17973
- Formal-solution signal: shows that targeted impact analysis and regression-aware context beat generic "do TDD" instructions for refinement quality.
- Task-0007 relevance: each proposed winner should surface the impacted seams and falsifier set explicitly instead of relying on broad workflow slogans.

## 03. Test-Driven AI Agent Definition (TDAD): Compiling Tool-Using Agents from Behavioral Specifications
- Published: 2026-03-09
- URL: https://arxiv.org/abs/2603.08806
- Formal-solution signal: treats the agent definition as a compiled artifact produced from behavioral specifications, then iteratively refined until visible and hidden tests pass.
- Task-0007 relevance: this is the closest research analogue to `solution spec -> verifier -> refinement loop`, which argues for `2A` and `2B` owning a living solution doc before task drafting.

## 04. Agentic Rubrics as Contextual Verifiers for SWE Agents
- Published: 2026-01-07
- URL: https://arxiv.org/abs/2601.04171
- Formal-solution signal: generates repository-specific rubric checklists that score candidate patches without requiring full execution every time.
- Task-0007 relevance: strengthens the case for `2B` as an explicit contextual-audit stage with blocking questions, narrowings, and codebase-specific checks.

## 05. EvoDev: An Iterative Feature-Driven Framework for End-to-End Software Development with LLM-based Agents
- Published: 2025-11-04
- URL: https://arxiv.org/abs/2511.02399
- Formal-solution signal: centers a feature map DAG where business logic, design intent, and code context are propagated across dependent work.
- Task-0007 relevance: supports explicit intermediate design artifacts instead of jumping from burden statements directly into task drafts.

## 06. TDFlow: Agentic Workflows for Test Driven Software Engineering
- Published: 2025-10-27
- URL: https://arxiv.org/abs/2510.23761
- Formal-solution signal: forces separation between proposing, debugging, revising, and test generation so no single stage silently owns the entire refinement loop.
- Task-0007 relevance: aligns with a split where `2A` designs options, `2B` audits and narrows them, and later stages consume that decision rather than redesigning it ad hoc.

## 07. Towards Engineering Multi-Agent LLMs: A Protocol-Driven Approach
- Published: 2025-10-14
- URL: https://arxiv.org/abs/2510.12120
- Formal-solution signal: argues that multi-agent SWE systems fail when they lack explicit behavioral contract modeling, structured messaging, and lifecycle-guided verification.
- Task-0007 relevance: implies the Dream packet should preserve contract-level distinctions between option design, winner synthesis, and downstream task drafting.

## 08. SWE-Debate: Competitive Multi-Agent Debate for Software Issue Resolution
- Published: 2025-07-31
- URL: https://arxiv.org/abs/2507.23348
- Formal-solution signal: explicitly creates multiple competing fault traces and forces them through debate before patch generation.
- Task-0007 relevance: supports restoring stronger exploration pressure such as three materially distinct options per problem before winner synthesis closes.

## 09. Magentic-One: A Generalist Multi-Agent System for Solving Complex Tasks
- Published: 2024-11-07
- URL: https://arxiv.org/abs/2411.04468
- Formal-solution signal: makes the orchestrator responsible for plan tracking and re-planning while specialists execute bounded work.
- Task-0007 relevance: supports a clean boundary where the refinement owner changes the solution contract deliberately rather than letting downstream workers mutate it implicitly.
