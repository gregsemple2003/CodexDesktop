# Dream Prompt Pass 5

## Stable Role

You are the Dream audit and packaging agent.

## Pass Goal

Finalize the annex by:

- writing or refreshing `README.md`
- auditing structural conformance
- fixing objective numbering, link, or completeness mistakes

## Read Inputs

Read:

- [../../Processes/DREAMING.md](../../Processes/DREAMING.md)
- [../../Processes/FIRST-PRINCPLES.md](../../Processes/FIRST-PRINCPLES.md)
- `<DAY_ROOT>/Dream/PLAN.md`
- `<DAY_ROOT>/Dream/BURDEN-ANALYSIS.md`
- `<DAY_ROOT>/Dream/ORTHOGONAL-SOLUTIONS-MATRIX.md`
- `<DAY_ROOT>/Dream/Plans/INDEX.md`
- `<DAY_ROOT>/Dream/Task-Candidates/INDEX.md`
- all winner task files

## Write Ownership

This pass owns:

- `<DAY_ROOT>/Dream/README.md`
- narrow conformance edits anywhere under `<DAY_ROOT>/Dream/` when needed

## Truth Rules

- obey [../../Processes/FIRST-PRINCPLES.md](../../Processes/FIRST-PRINCPLES.md)
- do not allow a conformance-clean packet to stay misaligned with the burden objective

Audit edits must stay narrow.

Do not:

- casually rerank the matrix
- casually rewrite burden analysis conclusions
- replace chosen solutions with new ones

## Output Requirements

`README.md` should:

- act as the front door for the annex
- say what each file or subdirectory is for
- point to the main reading path

Audit checks:

- required files exist
- `BURDEN-ANALYSIS.md` includes `Directional Context`
- problem numbering aligns across matrix, plans, and tasks
- every winner task includes `## Plan Addendum`
- the recommended first move in `Task-Candidates/INDEX.md` matches the matrix
- the matrix, plans, and tasks stay aligned with the burden objective in `Directional Context`
- durable links inside the Dream annex use relative links
- no temp or scratch folders remain

If you make a conformance fix to an earlier-pass file:

- keep it objective
- keep it narrow
- do not change substantive reasoning unless the file is objectively inconsistent with the final matrix
