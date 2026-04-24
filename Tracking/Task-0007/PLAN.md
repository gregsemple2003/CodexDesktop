# Task 0007 Plan

## Planning Verdict

This is a bootstrap research plan for the new analysis task.

Human gate:

- do not treat this plan as approval to implement `Jarvis` or promote shared workflow changes yet

## Planning Basis

The plan is grounded in:

- the April 19, 2026 conversation that framed each direct human input as a system failure signal
- the earlier incident-capture artifacts preserved in [Task-0006](/c:/Agent/CodexDashboard/Tracking/Task-0006/TASK.md)
- current orchestration thinking in [2026-04-16-SWE-Orchestration.md](/c:/Users/gregs/.codex/reports/2026-04-16-SWE-Orchestration.md)

## PASS-0000 Bootstrap The Analysis Home

### Objective

Create the task-owned home for conversation captures and synthesized notes.

### Implementation Notes

- create `Research/Conversations/` for dated conversation captures
- create `Research/Analysis/` for synthesized notes
- seed the task with one initial conversation note and one initial analysis note
- keep the first artifacts local to this task rather than promoting them into shared `.codex` docs

### Verification

- the seeded conversation note exists
- the seeded analysis note exists
- the task root contains the minimum bootstrap artifacts

### Exit Bar

- the task can now hold follow-on analysis without rediscovering why it exists

## PASS-0001 Define The Inference-Failure Taxonomy

### Objective

Separate the useful core of the model from the parts that would overclaim inference.

### Implementation Notes

- define the main intervention categories, such as:
  - missing memory
  - missing situational inference
  - missing repo understanding
  - missing initiative
  - genuine novelty or changed human desire
- make the explicit versus implied versus speculative split durable
- define how local human constraints should bound later task proposals

### Verification

- `RESEARCH-ANALYSIS.md` names the main categories and the honesty split

### Exit Bar

- later work can analyze interventions without pretending every input was equally predictable

## PASS-0002 Define The Daily Repo Report And `HUMAN-DESIRE.md` Shape

### Objective

Describe the minimum useful daily output for a Jarvis-style analysis loop.

### Implementation Notes

- define the daily report questions
- define how repo-level task proposals should be expressed
- describe the intended role of `HUMAN-DESIRE.md`
- make clear what evidence should be treated as explicit, repeated, or speculative

### Verification

- the research artifacts describe a daily report shape that could be reviewed by a human

### Exit Bar

- the task leaves behind a concrete report shape instead of a vague desire for smarter orchestration

## PASS-0003 Decide Promotion Boundaries And Follow-On Work

### Objective

Decide what should stay task-local and what should later move into shared orchestration docs or follow-on implementation tasks.

### Implementation Notes

- identify which outputs are still exploratory
- identify which outputs are stable enough to promote later
- identify whether a follow-on task should own implementation, scheduling, or repo integration

### Verification

- the handoff states what remains local versus what is a candidate for promotion

### Exit Bar

- future work can continue without confusing analysis, shared workflow, and implementation

## Task-Level Validation

This is an analysis and artifact task.

Expected validation for closure:

- artifact review of task, plan, handoff, and research notes
- no build, unit-test, or regression claim unless later work adds real code or runnable workflows

## Watchouts

- do not turn the task into immediate automation
- do not overfit on the idea that every intervention was equally inferable
- do not collapse humane values into unsupported preference guesses
- do not promote unstable conclusions into `.codex` too early
