# Intervention Time

This folder owns task-local workflows and evidence for estimating human intervention cost separately from needs extraction.

Current packet:

- [2026-04-19 Day Root](/c:/Users/gregs/.codex/Orchestration/Reports/Interventions/ThirdPerson/2026/04/19/INDEX.md)
- [2026-04-19 HumanInterventionTime Summary](/c:/Users/gregs/.codex/Orchestration/Reports/Interventions/ThirdPerson/2026/04/19/HumanInterventionTime/SUMMARY.json)
- [Shared HumanInterventionTime Prompt](/c:/Users/gregs/.codex/Orchestration/Reports/Interventions/Prompts/HumanInterventionTime/PROMPT.md)
- [Shared HumanInterventionTime Workflow](/c:/Users/gregs/.codex/Orchestration/Reports/Interventions/Prompts/HumanInterventionTime/WORKFLOW.md)
- [Prioritization Note](/c:/Agent/CodexDashboard/Tracking/Task-0007/InterventionTime/PRIORITIZATION.md)

Rules:

- keep intervention-time measurement separate from needs analysis
- preserve raw human input evidence and transcript provenance
- treat `typing_seconds` and `stall_loss_seconds` as separate components
- keep the workflow dispatchable from the same daily source-discovery root as harvest
- use `InterventionTime` as a rough prioritization proxy rather than a complete all-in-one utility score
