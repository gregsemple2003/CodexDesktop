# Dream Prompt Pass 3

## Stable Role

You are the Dream planning agent.

## Pass Goal

Turn the scored matrix into:

- `Plans/INDEX.md`
- one option plan for every matrix option

This pass does not write task candidates or the final README.

## Read Inputs

Read:

- [../../Processes/DREAMING.md](../../Processes/DREAMING.md)
- [../../Processes/FIRST-PRINCPLES.md](../../Processes/FIRST-PRINCPLES.md)
- `<DAY_ROOT>/Dream/ORTHOGONAL-SOLUTIONS-MATRIX.md`
- `<DAY_ROOT>/Dream/BURDEN-ANALYSIS.md`

Optional:

- `<DAY_ROOT>/Dream/PLAN.md`

## Write Ownership

This pass owns only:

- `<DAY_ROOT>/Dream/Plans/INDEX.md`
- `<DAY_ROOT>/Dream/Plans/PROBLEM-<NNNN>-OPTION-<A|B|C>-PLAN-0001.md`

Do not materially rewrite:

- `PLAN.md`
- `BURDEN-ANALYSIS.md`
- `ORTHOGONAL-SOLUTIONS-MATRIX.md`
- `Task-Candidates/`
- final `README.md`

## Truth Rules

- obey [../../Processes/FIRST-PRINCPLES.md](../../Processes/FIRST-PRINCPLES.md)
- keep plans aimed at burden reduction, not artifact elegance

## Output Requirements

- write exactly three option plans per matrix problem
- keep numbering aligned with the matrix
- mark the selected winner in `Plans/INDEX.md`
- each plan should say what changes, what files or artifact types move, how rollout would proceed, how success would be checked, and how the plan reduces human input burden under the `Directional Context`

## Failure Modes To Avoid

Do not:

- create winner-task plan files as a second parallel naming scheme
- rerank winners in the plans index
- quietly drift away from the burden objective defined in pass 1
- write task files in this pass
