# Intervention Packet Prompt Bundles

This folder contains the shared prompt bundles for canonical intervention day packets.

Use the shared packet structure and promotion rules from:

- [../../Processes/INTERVENTION-REPORTS.md](../../Processes/INTERVENTION-REPORTS.md)

Use these prompt bundles:

- [HumanInputEvents/WORKFLOW.md](./HumanInputEvents/WORKFLOW.md)
  - seed-preparation and promotion rules for canonical `HumanInputEvents` rebuilds
- [HumanInputEvents/PROMPT.md](./HumanInputEvents/PROMPT.md)
  - enrich one seeded `HumanInputEvents` packet as canonical JSON artifacts
- [HumanNeeds/PROMPT.md](./HumanNeeds/PROMPT.md)
  - shared inference lens for packet-level human-needs analysis
- [HumanNeeds/WORKFLOW.md](./HumanNeeds/WORKFLOW.md)
  - execution notes for the `HumanNeeds` packet pass
- [HumanInterventionTime/PROMPT.md](./HumanInterventionTime/PROMPT.md)
  - shared metric and output contract for intervention-time measurement
- [HumanInterventionTime/WORKFLOW.md](./HumanInterventionTime/WORKFLOW.md)
  - execution notes for the intervention-time pass
- [HumanInterventionTime/DISPATCH-PATTERN.md](./HumanInterventionTime/DISPATCH-PATTERN.md)
  - dispatch order and separation from the other packet passes

Keep these bundles reusable and packet-agnostic.
Batch-specific invocation text belongs in each packet-local `PROMPT-INPUT.md`.
