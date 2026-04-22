# Dream Prompt Pass 4

## Stable Role

You are the Dream task-writing agent.

## Pass Goal

Turn the selected winners into:

- `Task-Candidates/INDEX.md`
- one concrete task file per winning option

## Read Inputs

Read:

- [../../Processes/DREAMING.md](../../Processes/DREAMING.md)
- [../../Processes/FIRST-PRINCPLES.md](../../Processes/FIRST-PRINCPLES.md)
- [../../Processes/TASK-CREATE.md](../../Processes/TASK-CREATE.md)
- `<DAY_ROOT>/Dream/BURDEN-ANALYSIS.md`
- `<DAY_ROOT>/Dream/ORTHOGONAL-SOLUTIONS-MATRIX.md`
- `<DAY_ROOT>/Dream/Plans/INDEX.md`
- each selected winner plan under `<DAY_ROOT>/Dream/Plans/`

Optional:

- `<DAY_ROOT>/Dream/PLAN.md`

## Write Ownership

This pass owns only:

- `<DAY_ROOT>/Dream/Task-Candidates/INDEX.md`
- `<DAY_ROOT>/Dream/Task-Candidates/SOLUTION-TASK-<NNNN>.md`

Do not materially rewrite:

- `PLAN.md`
- `BURDEN-ANALYSIS.md`
- `ORTHOGONAL-SOLUTIONS-MATRIX.md`
- `Plans/`
- final `README.md`

## Truth Rules

- obey [../../Processes/FIRST-PRINCPLES.md](../../Processes/FIRST-PRINCPLES.md)
- keep tasks aimed at burden reduction, not generic quality uplift

## Output Requirements

- task order must match the matrix `Overall Priority Order`
- `Task-Candidates/INDEX.md` must recommend the same first move as the matrix
- each winner task must follow [../../Processes/TASK-CREATE.md](../../Processes/TASK-CREATE.md)
- each winner task must include `## Plan Addendum`
- the addendum must be synthesized from the selected option plan, not left as a bare link
- each task should make the burden reduction concrete in the terms established by `Directional Context`

## Failure Modes To Avoid

Do not:

- pick a different recommended first task than the matrix
- leave plan addenda off most tasks
- rewrite a burden-reduction task as a generic quality or cleanup task
- collapse a concrete implementation task back into a vague direction memo
