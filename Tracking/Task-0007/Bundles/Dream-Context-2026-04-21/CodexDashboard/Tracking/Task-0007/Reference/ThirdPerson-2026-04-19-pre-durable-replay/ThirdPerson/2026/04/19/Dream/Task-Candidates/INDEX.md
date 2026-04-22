# Task Candidate Index

Updated: 2026-04-21

These are concrete implementation task writeups for the seven winning solutions identified from the April 19 `ThirdPerson` intervention packet.

They are written to match the shared task-writing standards from:

- [TASK-CREATE.md](../../../../../../../../Processes/TASK-CREATE.md)
- [TASK-HARVESTER.md](../../../../../../../../Prompts/TASK-HARVESTER.md)
- [TASK-LEADER.md](../../../../../../../../Prompts/TASK-LEADER.md)
- [ORCHESTRATION.md](../../../../../../../../ORCHESTRATION.md)
- [TASK.md exemplar](../../../../../../../../Exemplars/TASK.md)
- [FILE-NAMING.md](../../../../../../../../FILE-NAMING.md)

Each task now names the chosen solution shape and the concrete docs, prompts, schemas, or task-owned artifacts that would change. If a task could not meet that bar, it should have been written as a consensus task or research task instead.

Plan packet:

- [Plans/INDEX.md](../Plans/INDEX.md)
  - durable implementation plans for all twenty-one matrix options plus the seven synthesized winner-task plans

Recommended order:

1. [SOLUTION-TASK-0001.md](./SOLUTION-TASK-0001.md)
   - save task ownership, stop reasons, and next-step state in `TASK-STATE.json`
2. [SOLUTION-TASK-0002.md](./SOLUTION-TASK-0002.md)
   - require a task-owned approval packet before any plan or change approval request
3. [SOLUTION-TASK-0003.md](./SOLUTION-TASK-0003.md)
   - add an append-only intervention lesson ledger and promotion rule
4. [SOLUTION-TASK-0004.md](./SOLUTION-TASK-0004.md)
   - add a claim manifest and block off-lane closure claims
5. [SOLUTION-TASK-0005.md](./SOLUTION-TASK-0005.md)
   - require a disagreement trace before any root-cause claim
6. [SOLUTION-TASK-0006.md](./SOLUTION-TASK-0006.md)
   - add a shared answer-shape workflow for direct questions
7. [SOLUTION-TASK-0007.md](./SOLUTION-TASK-0007.md)
   - require an evidence manifest for proof bundles

Recommended next task:

- [SOLUTION-TASK-0001.md](./SOLUTION-TASK-0001.md)
  - includes the selected plan as an inline addendum

Why this one is first:

- it removes the clearest direct burden from the packet: the human having to restart supervision and tell the system to keep going
- the other fixes still help, but they are less useful if the task keeps dropping ownership and falling back to the human at the wrong time

What should stay deferred until later:

- a larger UI layer for approval
- heavy background scheduler or watcher tooling
- hidden model-side memory that is not inspectable on disk
