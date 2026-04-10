# Task 0006 Handoff

## Current Status

`Task-0006` is now framed as an incident-capture task rather than a broad "make orchestration smarter" bucket.

For this task, an incident is a divergence where the human had to step in and express disagreement over an AI-produced or AI-directed outcome to preserve the intended human-facing result.

The current task definition was distilled from the seed transcript in [Scaling-Agent-Orchestration.md](/c:/Agent/CodexDashboard/Tracking/Task-0006/Research/Scaling-Agent-Orchestration.md) and pressure-tested against:

- [VISION-HARVESTER.md](/c:/Users/gregs/.codex/Orchestration/Prompts/VISION-HARVESTER.md)
- [GENERAL-DESIGNER.md](/c:/Users/gregs/.codex/Orchestration/Prompts/GENERAL-DESIGNER.md)
- [INTERFACE-DESIGNER.md](/c:/Users/gregs/.codex/Orchestration/Prompts/INTERFACE-DESIGNER.md)

Research is now complete enough to support a real plan:

- [RESEARCH-PLAN.md](/c:/Agent/CodexDashboard/Tracking/Task-0006/RESEARCH-PLAN.md)
- [RESEARCH-ANALYSIS.md](/c:/Agent/CodexDashboard/Tracking/Task-0006/RESEARCH-ANALYSIS.md)
- [RESEARCH.md](/c:/Agent/CodexDashboard/Tracking/Task-0006/RESEARCH.md)
- [INCIDENT-GOAL-STACK.md](/c:/Agent/CodexDashboard/Tracking/Task-0006/Research/INCIDENT-GOAL-STACK.md)
- [INCIDENT.schema.json](/c:/Agent/CodexDashboard/Tracking/Task-0006/Research/INCIDENT.schema.json)
- [INCIDENT-EXAMPLE-0001.json](/c:/Agent/CodexDashboard/Tracking/Task-0006/Research/INCIDENT-EXAMPLE-0001.json)
- [INCIDENT-EXAMPLE-0002.json](/c:/Agent/CodexDashboard/Tracking/Task-0006/Research/INCIDENT-EXAMPLE-0002.json)
- [SEED-INCIDENTS-LAST-5D.md](/c:/Agent/CodexDashboard/Tracking/Task-0006/Research/SEED-INCIDENTS-LAST-5D.md)
- [PLAN.md](/c:/Agent/CodexDashboard/Tracking/Task-0006/PLAN.md)

## Current Baseline

The first planning draft was rejected because it stayed at surface-level incident interpretation and did not preserve a usable upward explanation of what target state the human was actually protecting.

The revised baseline now requires every incident to preserve:

- the grounded pre-correction state, including active course when relevant
- the human intervention itself
- the concrete event-level expected and actual state
- one or more ordered `why_chains` that progressively generalize from the concrete target state toward the broader human principle being expressed

The task now also has a concrete task-local schema and two validated example incidents, so `PASS-0000` is no longer a vague contract-design pass. It can be checked by file existence and schema validation.

The earlier Home-card example was dropped because it failed the qualification gate. It described a real bug, but the preserved evidence did not clearly show a human correcting an AI-produced outcome.

The revised plan still includes an explicit five-incident seed requirement covering:

- usability or state truth
- information architecture and CTA framing
- UI semantics and operator jargon
- orchestration or producer misunderstanding
- proxy-proof or debugging misunderstanding

The five recommended starter incidents all come from April 3, 2026 through April 7, 2026 history rather than from the older seed transcript alone.

## Next Step

Request explicit human approval on `PLAN.md` before starting `PASS-0000`.

## Watchouts

- do not widen the task into full critic automation, training, or autonomous remediation
- do not flatten all incidents into "taste" or "design" problems
- do not lose the human correction path while trying to normalize incidents
- do not let the `why_chains` drift into unsupported abstraction before the concrete event is clear
- shared orchestration and prompt rules belong in `.codex`; task-local evidence and history belong under `Tracking/Task-0006/`
- `PLAN.md` is not approved.
- no pass, audit, bug, or regression artifacts exist yet
