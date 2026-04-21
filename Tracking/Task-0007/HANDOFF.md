# Task 0007 Handoff

## Current Status

`Task-0007` has been created as the task-owned home for analysis and conversation artifacts about the proposed `Jarvis` intervention model.

The current baseline is:

- the model treats direct human input as evidence that the system failed to infer something it should have inferred
- the recurring postmortem question is `How could the system have inferred the need for this input?`
- the work is still in the analysis phase
- no product implementation or job wiring has been started, but shared prompt promotion and one canonical repo-day packet promotion have now been completed

The task now has:

- a durable task definition in [TASK.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/TASK.md)
- a bootstrap plan in [PLAN.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/PLAN.md)
- seed research notes under [Research](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/README.md)
- task-local source transcripts for the Jarvis x Codex integration thread:
  - [2026-04-19 System Integration Spec 52](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Conversations/2026-04-19-SYSTEM-INTEGRATION-SPEC-52.md)
  - [2026-04-19 System Integration Spec 54](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Conversations/2026-04-19-SYSTEM-INTEGRATION-SPEC-54.md)
- a task-local source packet covering overlap papers, practitioner writing, open-source exemplars, and official product pages under `Research/Sources/`
- a promoted canonical repo-day packet for April 19 `ThirdPerson` under [ThirdPerson/2026/04/19](/c:/Users/gregs/.codex/Orchestration/Reports/Interventions/ThirdPerson/2026/04/19/INDEX.md)
- a separate intervention-time workflow packet under [InterventionTime](/c:/Agent/CodexDashboard/Tracking/Task-0007/InterventionTime/README.md)
- an approved target structure for canonical repo-day intervention packets in [DAY-PACKET-CANONICAL-STRUCTURE.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/Design/DAY-PACKET-CANONICAL-STRUCTURE.md)
- a topic-by-topic frontier pass under `Research/Analysis/`:
  - [TOPIC-01 input and event ingestion](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Analysis/TOPIC-01-INPUT-EVENT-INGESTION.md)
  - [TOPIC-02 intervention record shape](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Analysis/TOPIC-02-INTERVENTION-RECORD-SHAPE.md)
  - [TOPIC-03 inference](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Analysis/TOPIC-03-INFERENCE.md)
  - [TOPIC-04 deliberation record shape](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Analysis/TOPIC-04-DELIBERATION-RECORD-SHAPE.md)
  - [TOPIC-05 repo contract files](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Analysis/TOPIC-05-REPO-CONTRACT-FILES.md)
  - [TOPIC-06 local memory and learning](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Analysis/TOPIC-06-LOCAL-MEMORY-AND-LEARNING.md)
  - [TOPIC-07 daily repo report design](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Analysis/TOPIC-07-DAILY-REPO-REPORT-DESIGN.md)
  - [TOPIC-08 autonomy and safety policy](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Analysis/TOPIC-08-AUTONOMY-AND-SAFETY-POLICY.md)
  - [TOPIC-09 evaluation loops](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Analysis/TOPIC-09-EVALUATION-LOOPS.md)

## Next Step

Refine the topic pass into a small number of trialable local designs and then generalize the promoted intervention packet structure:

- a first event and intervention schema
- a bounded inference-honesty taxonomy
- a minimum repo contract stack around `HUMAN-DESIRE.md`
- a first daily per-repo Jarvis report shape
- a promotion rule for when recurring interventions become durable memory or repo policy
- a durable intervention-time loop that can be run beside harvest without contaminating needs analysis
- a shared-prompt plus `PROMPT-INPUT.md` split for canonical repo-day packets
- a day-root structure around `HumanInputEvents`, `HumanNeeds`, and `HumanInterventionTime`
- no historical backfill until the approved structure has been validated on current packets

## Watchouts

- do not confuse this task with immediate implementation approval
- do not claim the system should have inferred genuine novelty
- keep explicit, implied, and speculative conclusions visibly separate
- do not treat the topic docs as final truth; they are the first full-corpus frontier pass and still need convergence work
- keep the relationship to [Task-0006](/c:/Agent/CodexDashboard/Tracking/Task-0006/TASK.md) explicit instead of duplicating incident-capture work
- keep further shared `.codex` generalization gated on stable local analysis baselines and validated repo-day packets
