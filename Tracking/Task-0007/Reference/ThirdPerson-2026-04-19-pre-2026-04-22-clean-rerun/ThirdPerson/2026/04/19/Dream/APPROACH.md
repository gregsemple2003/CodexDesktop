# Approach

## Method

Rebuild pass 1 from the canonical day packet, then follow the burden-driving events into the source packet and the available session logs themselves.

This pass does not treat packet JSON as sufficient understanding. It uses:

- packet summaries for scope and chronology
- raw human event text from `../HumanInputEvents/SOURCE-PACKET.jsonl`
- available session-log neighborhoods from the lifecycle thread parts
- packet-level measurement summaries from `../HumanInterventionTime/SUMMARY.json`

## Scope

This pass is limited to:

- `APPROACH.md`
- `SessionExcerpts/INDEX.json`
- `BURDEN-ANALYSIS.md`

It does not write solutions, option tasks, or the final Dream front door.

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

Source-following used in this pass:

- lifecycle-thread session-log parts for session `019da19b-19c6-7e20-bdd8-f03aea761fb9`
- raw source-packet event text for session `019da198-ea2f-7421-8adc-c566da0b6121`, whose transcript is missing on disk

Fidelity notes:

- `HumanInputEvents/INDEX.json` is in chronological order.
- The main implementation-thread transcript `019da198-ea2f-7421-8adc-c566da0b6121` is still missing on disk, so assistant-side context for 46/62 events cannot be re-read directly.
- The current `HumanInputEvents` packet predates the new event-level `need_tag` contract. This pass therefore reconstructs recurring burden clusters directly from raw event text and source-log neighborhoods rather than from populated `need_tag` fields.

## Intended Deliverables

- one packet-local excerpt set that preserves the key burden neighborhoods
- one burden analysis that keeps the major burden drivers separate instead of flattening them into a few neat buckets

## Later Pass Sequence

Pass 2 should:

- reread the raw packet evidence and the relevant session-log neighborhoods directly
- preserve the burden distinctions established here unless the source material proves they should be merged
- test candidate remedies against the actual circumstances of the events, not just the labels in this pass

## Raw Evidence Surfaces To Reread

Later work should reread:

- ownership / restart-supervision events from `../HumanInputEvents/SOURCE-PACKET.jsonl`
  - especially `4015`, `5272`, `5310`
- lifecycle continuity and no-homework-back instructions from the lifecycle session-log parts
  - especially original lines `6558` and `7038`
- evidence-integrity and proof-surface disputes
  - especially source-packet `4909` and lifecycle line `6451`
- root-cause debugging discipline
  - especially source-packet `5340` and lifecycle line `7446`
