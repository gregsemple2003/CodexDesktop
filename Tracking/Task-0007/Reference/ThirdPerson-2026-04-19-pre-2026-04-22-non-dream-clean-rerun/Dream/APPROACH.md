# Approach

## Method

Rebuild Dream from the canonical day packet, then follow burden-driving events back to the authoritative seed packet rather than widening scope with a fresh search.

This run stays inside:

- the current durable intervention and Dream workflow docs
- the authoritative task-owned `SOURCE-PACKET.jsonl` and `SOURCE-SESSIONS.json`
- repo-local docs from `C:/Agent/ThirdPerson`

## Scope

This pass owns only:

- `APPROACH.md`
- `SessionExcerpts/INDEX.json`
- `BURDEN-ANALYSIS.md`

It does not score options or write task proposals.

## Evidence Base

Primary packet inputs:

- `../INDEX.md`
- `../DAY-MANIFEST.json`
- `../HumanInputEvents/INDEX.json`
- `../HumanInputEvents/SOURCE-PACKET.jsonl`
- `../HumanInputEvents/SOURCE-SESSIONS.json`
- `../HumanNeeds/PACKET-TRIAGE.json`
- `../HumanNeeds/PACKET-RECORD.json`
- `../HumanInterventionTime/EVENT-MEASUREMENTS.jsonl`
- `../HumanInterventionTime/SUMMARY.json`

Repo-local context:

- `../HumanNeeds/LOCAL-CONTEXT.json`
- `C:/Agent/ThirdPerson/AGENTS.md`
- `C:/Agent/ThirdPerson/REGRESSION.md`
- `C:/Agent/ThirdPerson/TESTING.md`

## Fidelity Notes

- `../HumanInputEvents/INDEX.json` is chronological and provides the stable event spine for the day.
- The authoritative packet now carries event-level `need_tag`, so Dream can preserve those recurring clusters instead of reconstructing them ad hoc.
- Both canonical session transcript paths named in `../HumanInputEvents/SOURCE-SESSIONS.json` are missing on disk during this rerun. Dream therefore grounds excerpt text in the seeded packet's event lines and does not claim assistant-side neighborhood re-reads that did not happen.
- The Dream workflow references `Processes/FIRST-PRINCIPLES.md`, but that file is absent on disk during this rerun. The run therefore grounds `truth`, `compassion`, and `tolerance` from `DREAMING.md` rather than a missing companion file.

## Intended Deliverables

- one evidence-backed burden analysis that preserves the recurring `need_tag` clusters unless evidence supports an explicit merge
- one packet-local excerpt index that keeps the burden-driving event lines close at hand for pass 2

## Later Pass Sequence

Pass 2 should reread the packet evidence, keep the burden distinctions honest, compare orthogonal remedy classes, and write one task-shaped proposal per matrix option.

## Raw Evidence Surfaces To Reread

Later work should reread:

- default-lane truth and proof-surface events in `../HumanInputEvents/SOURCE-PACKET.jsonl`
- ownership and no-homework-back events in the same packet
- the packet-level need and intervention-time outputs
- the ThirdPerson repo docs embedded in `../HumanNeeds/LOCAL-CONTEXT.json`
