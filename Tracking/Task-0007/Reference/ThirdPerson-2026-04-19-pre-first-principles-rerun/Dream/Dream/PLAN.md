# Dream Plan

## Method

Use the rebuilt April 19 packet as the main truth source.

Read the day packet first.
Read ThirdPerson repo docs only when the packet points there.
Keep `Observed` facts apart from `Inference`.
Rank problems by likely reduction in future human intervention time.
Prefer small fixes with hard enforcement over broad advice.

## Scope

This annex covers recurring burden in the April 19 ThirdPerson packet.

It is about:

- default-lane proof trust
- answer shape
- scope and pass boundaries
- approval friction
- runtime defect framing
- root-cause method
- continuation burden

It is not a rewrite of the base day packet.
It does not reopen the non-Dream packet outputs.
It does not propose engine changes because the packet adds a `no engine mods` constraint.

## Evidence Base

Primary packet inputs:

- `../INDEX.md`
- `../DAY-MANIFEST.json`
- `../HumanInputEvents/INDEX.json`
- `../HumanInputEvents/SOURCE-PACKET.jsonl`
- `../HumanNeeds/PACKET-TRIAGE.json`
- `../HumanNeeds/PACKET-RECORD.json`
- `../HumanNeeds/REPRESENTATIVE-EVENTS.json`
- `../HumanInterventionTime/SUMMARY.json`

Repo-local constraints used because the packet cites them:

- `C:\Agent\ThirdPerson\AGENTS.md`
- `C:\Agent\ThirdPerson\REGRESSION.md`
- `C:\Agent\ThirdPerson\TESTING.md`

Key packet facts:

- `62` human-input events
- `13` corrections
- `13` boundary resets
- `20` direct answers requested
- `23874.309` seconds total intervention time
- `2099.309` seconds charged stall loss

## Intended Deliverables

This Dream annex will produce:

- `BURDEN-ANALYSIS.md`
- `ORTHOGONAL-SOLUTIONS-MATRIX.md`
- `Plans/INDEX.md`
- `Plans/PROBLEM-<NNNN>-OPTION-<A|B|C>-PLAN-0001.md` for every option
- `Task-Candidates/INDEX.md`
- `Task-Candidates/SOLUTION-TASK-<NNNN>.md` for every winning option
- `README.md`

## Pass Sequence

Pass 1:

- write this file
- write `BURDEN-ANALYSIS.md`

Pass 2:

- write `ORTHOGONAL-SOLUTIONS-MATRIX.md`
- keep the same burden order unless a reorder is explained

Pass 3:

- write `Plans/INDEX.md`
- write one plan for every option in the matrix

Pass 4:

- write `Task-Candidates/INDEX.md`
- write one concrete task for each winning option
- inline the chosen plan under `## Plan Addendum`

Pass 5:

- write `README.md`
- audit numbering, links, completeness, and winner order

## Deliverable Bar

The annex is only complete if:

- burden drivers stay evidence-backed
- matrix problems match the burden numbering
- every problem has three real option types
- winner tasks stay in matrix priority order
- all links inside `Dream/` are relative
