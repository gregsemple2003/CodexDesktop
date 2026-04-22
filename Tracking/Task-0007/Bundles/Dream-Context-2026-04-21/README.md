# Dream Context Bundle

Date: 2026-04-21

This bundle is a self-contained context packet for external review of the
April 19 `ThirdPerson` intervention packet and Dream regeneration flow.

It includes:

- the current canonical packet
- the preserved reference packet used as the quality bar
- the shared process docs and prompt docs that shaped packet rebuild and Dream
  regeneration
- the shared task-writing docs used when Dream generated concrete tasks
- the repo and shared `AGENTS.md` files that could influence execution
- two task-owned research notes that directly influenced the latest prompt and
  process changes
- a ready-to-use review prompt for diagnosing why the current packet is worse
  than the reference

## Top-Level Contents

- `.codex/`
  - shared workflow, process, prompt, exemplar, and `AGENTS.md` context
- `CodexDashboard/`
  - repo-local `AGENTS.md`
  - task-owned reference packet copy
  - task-owned research notes that influenced the durable process rewrite

## Current Canonical Packet

The live packet being judged is here inside the bundle:

- `.codex/Orchestration/Reports/Interventions/ThirdPerson/2026/04/19/`

This includes:

- `HumanInputEvents/`
- `HumanNeeds/`
- `HumanInterventionTime/`
- `Dream/`
- `DAY-MANIFEST.json`
- `INDEX.md`

## Preserved Reference Packet

The earlier stronger packet used as the comparison bar is here:

- `CodexDashboard/Tracking/Task-0007/Reference/ThirdPerson-2026-04-19-pre-durable-replay/ThirdPerson/2026/04/19/`

## Shared Process And Prompt Inputs

These are the durable docs that shaped the process:

- `.codex/AGENTS.md`
- `.codex/Orchestration/ORCHESTRATION.md`
- `.codex/Orchestration/FILE-NAMING.md`
- `.codex/Orchestration/Prompts/README.md`
- `.codex/Orchestration/Processes/INTERVENTION-REPORTS.md`
- `.codex/Orchestration/Processes/FIRST-PRINCPLES.md`
- `.codex/Orchestration/Processes/DREAMING.md`
- `.codex/Orchestration/Processes/TASK-CREATE.md`
- `.codex/Orchestration/TASK-STATE.md`
- `.codex/Orchestration/TASK-STATE.schema.json`
- `.codex/Orchestration/PASS-CHECKLIST.md`
- `.codex/Orchestration/PASS-CHECKLIST.schema.json`
- `.codex/Orchestration/AUDIT-RESULT.md`
- `.codex/Orchestration/AUDIT-RESULT.schema.json`
- `.codex/Orchestration/Prompts/Interventions/README.md`
- `.codex/Orchestration/Prompts/Interventions/HumanInputEvents/WORKFLOW.md`
- `.codex/Orchestration/Prompts/Interventions/HumanInputEvents/PROMPT.md`
- `.codex/Orchestration/Prompts/Interventions/HumanNeeds/WORKFLOW.md`
- `.codex/Orchestration/Prompts/Interventions/HumanNeeds/PROMPT.md`
- `.codex/Orchestration/Prompts/Interventions/HumanInterventionTime/WORKFLOW.md`
- `.codex/Orchestration/Prompts/Interventions/HumanInterventionTime/PROMPT.md`
- `.codex/Orchestration/Prompts/Interventions/HumanInterventionTime/DISPATCH-PATTERN.md`
- `.codex/Orchestration/Prompts/Dream/WORKFLOW.md`
- `.codex/Orchestration/Prompts/Dream/PROMPT-PASS1.md`
- `.codex/Orchestration/Prompts/Dream/PROMPT-PASS2.md`
- `.codex/Orchestration/Prompts/Dream/PROMPT-PASS3.md`
- `.codex/Orchestration/Prompts/Dream/PROMPT-PASS4.md`
- `.codex/Orchestration/Prompts/Dream/PROMPT-PASS5.md`
- `.codex/Orchestration/Prompts/TASK-HARVESTER.md`
- `.codex/Orchestration/Prompts/TASK-LEADER.md`
- `.codex/Orchestration/Prompts/IMPLEMENTATION-LEADER.md`
- `.codex/Orchestration/Exemplars/TASK.md`
- `.codex/Orchestration/Exemplars/PLAN.md`
- `CodexDashboard/AGENTS.md`

## Supplemental Notes

These are not execution prompts, but they influenced the durable rewrite:

- `CodexDashboard/Tracking/Task-0007/Research/Analysis/2026-04-21-DREAM-REPLAY-CONFUSION-ANALYSIS.md`
- `CodexDashboard/Tracking/Task-0007/Research/Analysis/2026-04-21-HUMAN-INPUT-BURDEN-AND-PRIME-DIRECTIVE.md`

## Review Prompt

Use this file to ask an external model for the root cause of the current
packet's weaker results:

- `ROOT-CAUSE-REVIEW-PROMPT.md`

## Intended Review Order

1. Read `FIRST-PRINCPLES.md`.
2. Read `INTERVENTION-REPORTS.md`.
3. Read the Dream process and pass prompts.
4. Read the reference packet.
5. Read the current packet.
6. Compare the Dream outputs, rankings, plans, and task candidates.

## Notes

- Relative links were preserved where possible by copying docs into matching
  subtree layouts.
- This bundle is for external review only. It does not modify the canonical
  packet or the shared prompt set.
