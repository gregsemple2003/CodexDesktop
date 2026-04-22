# Dreaming Process

`Dream` is an optional packet-local analysis annex for one canonical intervention day packet.

Use it when the human wants deeper solution design work that goes beyond the canonical `HumanInputEvents`, `HumanNeeds`, and `HumanInterventionTime` packet outputs.

Do not treat `Dream` as part of the minimum canonical day-packet shape.

## Shared Focus

Dream uses [FIRST-PRINCPLES.md](./FIRST-PRINCPLES.md) as shared directional context.

That file defines:

- the burden objective
- the prime-directive reading of direct human input
- `truth`, `compassion`, and `tolerance` as explicit governing principles
- the ranking rule for later passes

## Purpose

Use `Dream` to:

- explain why the human had to intervene
- separate burden drivers from one-off noise
- generate and compare orthogonal solutions
- turn the selected winners into concrete implementation tasks

Do not use `Dream` for:

- raw transcript storage
- generic research dumping
- long-lived scratch or temp output
- packet-local copies of shared prompt bundles

## Normative Structure

Required local shape:

```text
<DayRoot>/Dream/
  README.md
  PLAN.md
  BURDEN-ANALYSIS.md
  ORTHOGONAL-SOLUTIONS-MATRIX.md
  Plans/
    INDEX.md
    PROBLEM-0001-OPTION-A-PLAN-0001.md
  Task-Candidates/
    INDEX.md
    SOLUTION-TASK-0001.md
```

Required meanings:

- `README.md`
  - front door to the local Dream annex
  - finalized in the audit pass after the other artifacts exist
- `PLAN.md`
  - method, scope, evidence base, and intended deliverables
- `BURDEN-ANALYSIS.md`
  - the evidence-backed explanation of why the human had to intervene
  - must include `Directional Context` that states the burden-reduction objective, the first principles, and the prime-directive reading of direct human input
- `ORTHOGONAL-SOLUTIONS-MATRIX.md`
  - the problem set, three orthogonal options per problem, scoring, chosen winners, and rollout order
- `Plans/`
  - implementation plans for each solution option, not only the winners
  - name option plans as `PROBLEM-<NNNN>-OPTION-<A|B|C>-PLAN-0001.md`
- `Task-Candidates/`
  - concrete enqueue-ready task writeups for the selected winners
  - these writeups must follow [TASK-CREATE.md](./TASK-CREATE.md)
  - name winner task files as `SOLUTION-TASK-<NNNN>.md` in rollout order
  - every winner task must inline its chosen plan as `## Plan Addendum`

## Writing Contract

Keep the writing contract small and hard:

- separate `Observed` from `Inferred`
- use short words and short sentences
- name each burden in plain English
- optimize for reducing human input burden, not for making the artifact sound neat
- treat `InterventionTime` as the first approximation to human cost unless the packet gives a better grounded proxy
- treat direct human input as an error signal until the packet proves it was genuine novelty, required approval, or hidden external state
- keep `truth`, `compassion`, and `tolerance` explicit
- do not blame the human for boundary repair, frustration, or truth correction
- for each proposed fix, say:
  - what changes
  - what problem it removes
  - how a human would notice the improvement
- prefer concrete examples over abstract labels
- cut inflated wording that says less than the plain version

Do not create a packet-local writing-standards subtree unless the human explicitly asks for one as a temporary aid.

## Pass Ownership

Dream is a multi-pass workflow.

Pass ownership is hard:

- pass 1 owns `PLAN.md` and `BURDEN-ANALYSIS.md`
- pass 2 owns `ORTHOGONAL-SOLUTIONS-MATRIX.md`
- pass 3 owns `Plans/`
- pass 4 owns `Task-Candidates/`
- pass 5 owns `README.md` and the final conformance pass

Later passes may consume earlier outputs, but they should not silently rewrite earlier substantive conclusions.

If a later pass finds an objective structural problem:

- fix it in the audit pass
- keep the change narrow
- do not casually re-argue the burden analysis or rerank winners without saying why

## Workflow

1. Ground in the canonical day packet.
2. Run pass 1.
   - write `PLAN.md`
   - write `BURDEN-ANALYSIS.md`
   - ground the pass in `FIRST-PRINCPLES.md`
   - write `Directional Context` inside `BURDEN-ANALYSIS.md`
   - state that the objective is to reduce human input burden, with `InterventionTime` as the first approximation
   - state `truth`, `compassion`, and `tolerance` explicitly
   - state that direct human input is failure telemetry by default
   - identify recurring burden drivers
   - keep `Observed`, `Inferred`, and `Why this increased intervention time` separate
3. Run pass 2.
   - ground the pass in `FIRST-PRINCPLES.md`
   - write `ORTHOGONAL-SOLUTIONS-MATRIX.md`
   - keep `4` to `7` burden problems unless the packet honestly supports fewer
   - if the packet supports all seven default burden types as distinct patterns, keep all seven separate
   - for each problem, write three genuinely different options
   - use one explicit scoring matrix
   - choose one winner per problem
   - include explicit `Objective`, `Exemplar Inputs`, `First Principles`, `Resolved Winners`, and `Overall Priority Order` sections
4. Run pass 3.
   - ground the pass in `FIRST-PRINCPLES.md`
   - write `Plans/`
   - write one implementation plan for every option, not just the winners
5. Run pass 4.
   - ground the pass in `FIRST-PRINCPLES.md`
   - write `Task-Candidates/`
   - turn the selected winners into concrete implementation tasks
   - write those tasks to the shared standard in [TASK-CREATE.md](./TASK-CREATE.md)
   - inline the chosen plan as `## Plan Addendum` in every winner task
   - keep the task order aligned with the matrix `Overall Priority Order`
6. Run pass 5.
   - ground the audit in `FIRST-PRINCPLES.md`
   - write or refresh `README.md`
   - audit numbering, links, pass conformance, and output completeness
7. Prune the annex.
   - keep the folder narrow
   - delete scaffolding that is no longer needed

## Cleanup Rules

In a durable Dream annex, do not keep:

- `Temp/`
- packet-local `Writing-Standards/`
- packet-local prompt bundles once the shared prompt set exists under `Orchestration/Prompts/Dream/`
- superseded exploratory notes whose content now lives in `BURDEN-ANALYSIS.md`, `ORTHOGONAL-SOLUTIONS-MATRIX.md`, `Plans/`, or `Task-Candidates/`

## Completion Checks

Before a Dream annex is considered complete, confirm:

1. `ORTHOGONAL-SOLUTIONS-MATRIX.md` numbers the problems in the same order used by `Plans/` and `Task-Candidates/`
2. `BURDEN-ANALYSIS.md` includes `Directional Context` and that context is consistent with the later matrix, plans, and tasks
3. `Directional Context` names `truth`, `compassion`, and `tolerance` explicitly
4. `ORTHOGONAL-SOLUTIONS-MATRIX.md` includes `Objective`, `Exemplar Inputs`, `First Principles`, and `Overall Priority Order`
5. `Plans/` contains exactly three option plans per selected problem
6. `Task-Candidates/` contains exactly one concrete task per winning option
7. every winning task includes `## Plan Addendum`
8. the recommended first move in `Task-Candidates/INDEX.md` matches the first item in `ORTHOGONAL-SOLUTIONS-MATRIX.md`
9. if any of the default seven burden types were merged or dropped, the matrix says which ones and why
10. repo-local constraints were taken from the packet repo, not the current workspace repo

## Shared Prompt Set

The shared Codex-consumable prompt set for Dream lives here:

- [Prompts/Dream/WORKFLOW.md](../Prompts/Dream/WORKFLOW.md)
- [Prompts/Dream/PROMPT-PASS1.md](../Prompts/Dream/PROMPT-PASS1.md)
- [Prompts/Dream/PROMPT-PASS2.md](../Prompts/Dream/PROMPT-PASS2.md)
- [Prompts/Dream/PROMPT-PASS3.md](../Prompts/Dream/PROMPT-PASS3.md)
- [Prompts/Dream/PROMPT-PASS4.md](../Prompts/Dream/PROMPT-PASS4.md)
- [Prompts/Dream/PROMPT-PASS5.md](../Prompts/Dream/PROMPT-PASS5.md)
