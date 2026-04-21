# Task 0007 Research

`Task-0007` is a holding task for the Jarvis analysis thread.

The current synthesis is:

- the useful core is treating intervention as evidence that the current system boundary was insufficient
- the central question is `How could the system have inferred the need for this input?`
- the main risk is overclaiming inference where the human was expressing novelty, changed desire, or information the system did not actually have
- the current frontier pass is now split into topic-owned analysis docs that read the full local markdown corpus with one topic in mind
- the next honest step is refinement across topics into a small set of schemas, contracts, and operating loops that can actually be tried locally

## Research Decomposition

For now, the research is narrowed to these topics:

1. Input and event ingestion.
   Questions:
   - which human surfaces should be canonical first
   - how GitHub comments, chat inputs, hooks, and app-server events should enter one normalized intake path
   - whether local Codex prompt capture is advisory, canonical, or supplemental

2. Intervention record shape.
   Questions:
   - what the durable equivalent of an `InterventionInferenceRecord` should be
   - which fields are mandatory for evidence, confidence, local constraints, and preventability
   - how event-level expected state, actual state, intervention, and outcome should be preserved

3. Inference.
   Questions:
   - how inference honesty should be encoded explicitly
   - how to separate missing memory, missing repo understanding, missing initiative, and genuine novelty
   - how to distinguish preventable misses from non-preventable new information
   - how to keep explicit, implied, and speculative inference visibly separate

4. Deliberation record shape.
   Questions:
   - what a durable structured reasoning summary should contain
   - how to preserve the required terminal question `How could the system have inferred the need for this input?`
   - how to capture reasoning artifacts without storing unsafe or low-value hidden chain-of-thought

5. Repo contract files.
   Questions:
   - what the minimum useful shape of `HUMAN-DESIRE.md` is
   - whether `HUMAN-CONSTRAINTS.md` is needed as a distinct artifact
   - how `HUMAN-DESIRE.md`, `AGENTS.md`, and accepted corrections should interact and which are hard constraints

6. Local memory and learning.
   Questions:
   - how accepted corrections become durable local memory
   - when recurring interventions should update repo instructions, tests, validators, or task proposals
   - how a human can scope, forget, or demote a learned pattern

7. Daily repo report design.
   Questions:
   - what one daily report per repo should answer
   - how task proposals should be structured
   - how the report should separate evidence, inference, uncertainty, and recommended action

8. Autonomy and safety policy.
   Questions:
   - which classes of work can be auto-delegated
   - which changes require explicit human approval
   - how truthfulness, compassion, and tolerance should constrain routing, wording, and action

9. Evaluation loops.
   Questions:
   - how to measure reduced repeated intervention
   - how to measure proposal usefulness, merge quality, and rework
   - how to detect unsupported inference, blame drift, or false generalization

Deferred for now:

- raw storage format details such as `jsonl` versus SQLite versus hybrid persistence
- the precise Jarvis versus Codex execution boundary
- artifact promotion and shared-doc boundary details

## Immediate Priorities

The immediate priorities are now:

1. input and event ingestion
2. intervention record shape
3. inference honesty and taxonomy
4. deliberation record shape
5. repo contract files
6. local memory and learning
7. daily repo report design
8. autonomy and safety policy
9. evaluation loops

## Topic Analyses

These topic docs are the current frontier review set. Each one was written from a full read of the task-local markdown corpus with a specific topic in mind, not from a keyword skim.

- [TOPIC-01 input and event ingestion](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Analysis/TOPIC-01-INPUT-EVENT-INGESTION.md)
- [TOPIC-02 intervention record shape](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Analysis/TOPIC-02-INTERVENTION-RECORD-SHAPE.md)
- [TOPIC-03 inference](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Analysis/TOPIC-03-inference.md)
- [TOPIC-04 deliberation record shape](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Analysis/TOPIC-04-DELIBERATION-RECORD-SHAPE.md)
- [TOPIC-05 repo contract files](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Analysis/TOPIC-05-REPO-CONTRACT-FILES.md)
- [TOPIC-06 local memory and learning](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Analysis/TOPIC-06-LOCAL-MEMORY-AND-LEARNING.md)
- [TOPIC-07 daily repo report design](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Analysis/TOPIC-07-DAILY-REPO-REPORT-DESIGN.md)
- [TOPIC-08 autonomy and safety policy](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Analysis/TOPIC-08-AUTONOMY-AND-SAFETY-POLICY.md)
- [TOPIC-09 evaluation loops](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Analysis/TOPIC-09-EVALUATION-LOOPS.md)

Supporting notes:

- [RESEARCH-PLAN.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/RESEARCH-PLAN.md)
- [RESEARCH-ANALYSIS.md](/c:/Agent/CodexDashboard/Tracking/Task-0007/RESEARCH-ANALYSIS.md)
- [2026-04-19 Jarvis model capture](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Conversations/2026-04-19-JARVIS-MODEL.md)
- [2026-04-19 System Integration Spec 52](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Conversations/2026-04-19-SYSTEM-INTEGRATION-SPEC-52.md)
- [2026-04-19 System Integration Spec 54](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Conversations/2026-04-19-SYSTEM-INTEGRATION-SPEC-54.md)
- [2026-04-19 bootstrap analysis](/c:/Agent/CodexDashboard/Tracking/Task-0007/Research/Analysis/2026-04-19-BOOTSTRAP-ANALYSIS.md)
