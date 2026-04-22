# Dream Prompt Workflow

## Objective

Use the Dream prompt set to turn one canonical intervention day packet into one narrow Dream annex.

The workflow owns prompt ordering. Prompt files are numbered by pass order. If a future pass can run in parallel, note that in this workflow.

## Required Inputs

Read these first:

- [../../Processes/DREAMING.md](../../Processes/DREAMING.md)
- [../../Processes/FIRST-PRINCPLES.md](../../Processes/FIRST-PRINCPLES.md)
- [../../Processes/TASK-CREATE.md](../../Processes/TASK-CREATE.md)
- [PROMPT-PASS1.md](./PROMPT-PASS1.md)
- [PROMPT-PASS2.md](./PROMPT-PASS2.md)
- [PROMPT-PASS3.md](./PROMPT-PASS3.md)
- [PROMPT-PASS4.md](./PROMPT-PASS4.md)
- [PROMPT-PASS5.md](./PROMPT-PASS5.md)
- `<DAY_ROOT>/INDEX.md`
- `<DAY_ROOT>/DAY-MANIFEST.json`
- `<DAY_ROOT>/HumanInputEvents/INDEX.json`
- `<DAY_ROOT>/HumanInputEvents/SOURCE-PACKET.jsonl`
- `<DAY_ROOT>/HumanNeeds/PACKET-TRIAGE.json`
- `<DAY_ROOT>/HumanNeeds/PACKET-RECORD.json`
- `<DAY_ROOT>/HumanInterventionTime/SUMMARY.json`

Optional inputs:

- `<DAY_ROOT>/HumanNeeds/REPRESENTATIVE-EVENTS.json`
- `<DAY_ROOT>/HumanInputEvents/SOURCE-SESSIONS.json`
- the original session transcripts named there
- repo-local docs or task artifacts directly cited by the packet

When repo-local docs are needed:

- use docs from the packet repo, not the current workspace repo
- resolve that repo from durable packet inputs such as `repo_ref`, `SOURCE-PACKET.jsonl`, or `SOURCE-SESSIONS.json`

## Output Location

Write the outputs into:

- `<DAY_ROOT>/Dream/`

Keep the finished local structure aligned with [../../Processes/DREAMING.md](../../Processes/DREAMING.md).

## Passes

1. Run [PROMPT-PASS1.md](./PROMPT-PASS1.md).
   - write only `PLAN.md` and `BURDEN-ANALYSIS.md`
   - `BURDEN-ANALYSIS.md` must include binding `Directional Context`
   - obey [../../Processes/FIRST-PRINCPLES.md](../../Processes/FIRST-PRINCPLES.md)
2. Run [PROMPT-PASS2.md](./PROMPT-PASS2.md).
   - write only `ORTHOGONAL-SOLUTIONS-MATRIX.md`
   - consume pass 1 outputs
   - treat `Directional Context` from `BURDEN-ANALYSIS.md` as binding
   - obey [../../Processes/FIRST-PRINCPLES.md](../../Processes/FIRST-PRINCPLES.md)
3. Run [PROMPT-PASS3.md](./PROMPT-PASS3.md).
   - write only `Plans/`
   - consume the matrix from pass 2
   - keep plans aligned with pass 1 `Directional Context`
   - obey [../../Processes/FIRST-PRINCPLES.md](../../Processes/FIRST-PRINCPLES.md)
4. Run [PROMPT-PASS4.md](./PROMPT-PASS4.md).
   - write only `Task-Candidates/`
   - consume the matrix and selected plans
   - when the pass writes `Task-Candidates/`, those task writeups must follow [../../Processes/TASK-CREATE.md](../../Processes/TASK-CREATE.md)
   - keep tasks aligned with pass 1 `Directional Context`
   - obey [../../Processes/FIRST-PRINCPLES.md](../../Processes/FIRST-PRINCPLES.md)
5. Run [PROMPT-PASS5.md](./PROMPT-PASS5.md).
   - write or refresh `README.md`
   - audit numbering, links, and pass conformance
   - audit consistency with pass 1 `Directional Context`
   - obey [../../Processes/FIRST-PRINCPLES.md](../../Processes/FIRST-PRINCPLES.md)

Pass dependency rule:

- later passes may read earlier pass outputs
- passes 2 to 4 should not materially rewrite earlier-pass artifacts
- only pass 5 may patch earlier files, and only to resolve objective structure or conformance problems

## Acceptance Check Before You Stop

Before you finish, confirm:

1. `BURDEN-ANALYSIS.md` separates `Observed` from `Inferred`
2. `BURDEN-ANALYSIS.md` includes `Directional Context` with the human-burden objective and the first principles
3. every pass stayed aligned with [../../Processes/FIRST-PRINCPLES.md](../../Processes/FIRST-PRINCPLES.md)
4. `ORTHOGONAL-SOLUTIONS-MATRIX.md` has `4` to `7` problems and three orthogonal options per problem
5. `ORTHOGONAL-SOLUTIONS-MATRIX.md` includes `Objective`, `Exemplar Inputs`, `First Principles`, `Resolved Winners`, and `Overall Priority Order`
6. if the packet supports all seven default burden types separately, the matrix keeps all seven separate or explains the merge
7. `Plans/` includes one plan per option and uses `PROBLEM-<NNNN>-OPTION-<A|B|C>-PLAN-0001.md`
8. `Task-Candidates/` contains one task per winner and uses `SOLUTION-TASK-<NNNN>.md`
9. every winner task includes `## Plan Addendum`
10. the recommended first move in `Task-Candidates/INDEX.md` matches the first item in the matrix `Overall Priority Order`
11. the local Dream folder matches the normative structure from [../../Processes/DREAMING.md](../../Processes/DREAMING.md)
