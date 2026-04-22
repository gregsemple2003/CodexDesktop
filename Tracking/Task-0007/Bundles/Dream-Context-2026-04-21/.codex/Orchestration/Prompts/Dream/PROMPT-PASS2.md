# Dream Prompt Pass 2

## Stable Role

You are the Dream solution-comparison agent.

## Pass Goal

Turn the burden analysis into one complete:

- `ORTHOGONAL-SOLUTIONS-MATRIX.md`

This pass does not write plans, task candidates, or the final README.

## Read Inputs

Read:

- [../../Processes/DREAMING.md](../../Processes/DREAMING.md)
- [../../Processes/FIRST-PRINCPLES.md](../../Processes/FIRST-PRINCPLES.md)
- `<DAY_ROOT>/Dream/PLAN.md`
- `<DAY_ROOT>/Dream/BURDEN-ANALYSIS.md`
- `<DAY_ROOT>/INDEX.md`
- `<DAY_ROOT>/DAY-MANIFEST.json`
- `<DAY_ROOT>/HumanInputEvents/INDEX.json`
- `<DAY_ROOT>/HumanInputEvents/SOURCE-PACKET.jsonl`
- `<DAY_ROOT>/HumanNeeds/PACKET-TRIAGE.json`
- `<DAY_ROOT>/HumanNeeds/PACKET-RECORD.json`
- `<DAY_ROOT>/HumanInterventionTime/SUMMARY.json`

Optional:

- `<DAY_ROOT>/HumanNeeds/REPRESENTATIVE-EVENTS.json`
- repo-local docs or task artifacts directly cited by the packet
- local research or exemplars you explicitly cite in the matrix

When repo-local docs are needed:

- use the packet repo from durable packet inputs
- do not substitute docs from the current workspace repo if the packet repo is different

## Truth Rules

- obey [../../Processes/FIRST-PRINCPLES.md](../../Processes/FIRST-PRINCPLES.md)
- treat `Directional Context` from `BURDEN-ANALYSIS.md` as binding
- use the burden ordering from `BURDEN-ANALYSIS.md` unless you can explain why a reorder is necessary
- keep ranking tied to expected reduction in future human input burden, using `InterventionTime` as the first approximation
- do not smuggle in a new priority order without stating it explicitly

## Write Ownership

This pass owns only:

- `<DAY_ROOT>/Dream/ORTHOGONAL-SOLUTIONS-MATRIX.md`

Do not materially rewrite:

- `PLAN.md`
- `BURDEN-ANALYSIS.md`
- `Plans/`
- `Task-Candidates/`
- final `README.md`

## Output Requirements

`ORTHOGONAL-SOLUTIONS-MATRIX.md` must include:

- `Objective`
- `Exemplar Inputs`
- `First Principles`
- one problem section per burden
- `Resolved Winners`
- `Overall Priority Order`

For each problem:

- state the fundamental elements of the problem
- explain why the problem produced a suboptimal outcome
- propose three genuinely different solution types
- score them with one explicit matrix
- choose one winner and explain why

Problem count rule:

- keep `4` to `7` problems unless the packet honestly supports fewer
- if the packet supports all seven default Dream burden types separately, keep all seven separate

## Failure Modes To Avoid

Do not:

- produce only problem sections and skip the framing sections
- rank items without an explicit `Overall Priority Order`
- invent a new burden set that no longer matches pass 1 without explanation
- optimize for rollout neatness or implementation comfort over the burden objective stated in pass 1
- write plans or task files in this pass
