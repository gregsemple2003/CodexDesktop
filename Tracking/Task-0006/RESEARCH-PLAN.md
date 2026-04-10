# Task 0006 Research Plan

## Research Goal

Turn the incident-capture framing in `Task-0006` into a planning-ready implementation path that can preserve real human interventions as durable orchestration evidence without widening into a full critic stack.

## Decision-Shaping Problems

### Problem 0001 Incident Contract Boundary And Why Chains

Determine the minimum record shape for an `incident` so the system captures explicit human disagreement over an AI-produced or AI-directed outcome, stays grounded in the concrete event, and climbs upward through progressive `why_chains` instead of flattening everything into generic bugs or overbuilt root-cause trees.

Also determine how the record matures in two steps:

- first-pass capture to preserve the correction quickly
- second-pass root-cause refinement to tighten the incident against the real mechanism once stronger evidence exists

### Problem 0002 Storage And Workflow Fit

Determine where incidents should live first, how they should link back to task artifacts and session evidence, and which parts belong under shared `.codex` orchestration docs versus task-local history.

### Problem 0003 Seed Backfill And Coverage Standard

Determine what a useful starter incident set looks like by backfilling five incidents from the last five days of history and checking that the set covers usability, UI semantics, and general misunderstandings rather than five near-duplicates.

## Research Inputs

- `Tracking/Task-0006/TASK.md`
- `Tracking/Task-0006/HANDOFF.md`
- `Tracking/Task-0006/Research/Scaling-Agent-Orchestration.md`
- `Tracking/Task-0006/Research/INCIDENT-GOAL-STACK.md`
- `C:\Users\gregs\.codex\Orchestration\Prompts\VISION-HARVESTER.md`
- `C:\Users\gregs\.codex\Orchestration\Prompts\GENERAL-DESIGNER.md`
- `C:\Users\gregs\.codex\Orchestration\Prompts\INTERFACE-DESIGNER.md`
- recent Codex session history from April 3, 2026 through April 7, 2026 under `C:\Users\gregs\.codex\sessions\`

## Intended Outputs

- `RESEARCH-ANALYSIS.md`
- `RESEARCH.md`
- `Research/SEED-INCIDENTS-LAST-5D.md`
- a planning-ready `PLAN.md`

## Exit Bar

- the incident contract is narrow enough to exclude ordinary bugs and broad enough to preserve the real correction path
- the incident contract requires the AI pre-correction outcome and the human intervention
- the incident contract includes explicit `why_chains` ordered from concrete target state to broader human principle
- the research output makes the first-pass capture versus second-pass refinement split explicit
- the first storage and linking strategy is explicit
- the seed set includes five recent incidents with deliberate coverage across usability, UI, and general misunderstandings
- the plan names a realistic pass order for contract definition, seed backfill, root-cause refinement, and workflow integration
