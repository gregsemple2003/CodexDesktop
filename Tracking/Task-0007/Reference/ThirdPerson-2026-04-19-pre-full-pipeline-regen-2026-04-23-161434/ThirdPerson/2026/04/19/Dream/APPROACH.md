# Dream Approach (ThirdPerson 2026-04-19)

## Method

1. Read the canonical day packet and validate `HumanInputEvents/INDEX.json` chronology.
2. Use `HumanInputEvents` as an index into the packet evidence.
3. Follow `SOURCE-SESSIONS.json` to the authoritative transcript surface when available.
   - One seeded session is `transcript_surface = missing`, so those events stay grounded in `SOURCE-PACKET.jsonl`.
   - One seeded session is `transcript_surface = ordered_parts`, so excerpts are grounded in the listed `.parts` files.
4. Keep a packet-local excerpt index for the neighborhoods that clarify meaning.
5. Write an evidence-backed burden inventory keyed by stable `event_id` values.
6. (Pass 2) Convert burdens into an orthogonal solution matrix and concrete option tasks.

## Scope

This annex is scoped strictly to the `ThirdPerson` intervention packet for local date `2026-04-19`.

## Evidence Base (Primary Inputs)

Canonical day packet:

- [`../INDEX.md`](../INDEX.md)
- [`../DAY-MANIFEST.json`](../DAY-MANIFEST.json)
- [`../HumanInputEvents/INDEX.json`](../HumanInputEvents/INDEX.json)
- [`../HumanInputEvents/SOURCE-PACKET.jsonl`](../HumanInputEvents/SOURCE-PACKET.jsonl)
- [`../HumanInputEvents/SOURCE-SESSIONS.json`](../HumanInputEvents/SOURCE-SESSIONS.json)
- [`../HumanNeeds/PACKET-TRIAGE.json`](../HumanNeeds/PACKET-TRIAGE.json)
- [`../HumanNeeds/PACKET-RECORD.json`](../HumanNeeds/PACKET-RECORD.json)
- [`../HumanNeeds/REPRESENTATIVE-EVENTS.json`](../HumanNeeds/REPRESENTATIVE-EVENTS.json)
- [`../HumanNeeds/LOCAL-CONTEXT.json`](../HumanNeeds/LOCAL-CONTEXT.json)
- [`../HumanInterventionTime/EVENT-MEASUREMENTS.jsonl`](../HumanInterventionTime/EVENT-MEASUREMENTS.jsonl)
- [`../HumanInterventionTime/SUMMARY.json`](../HumanInterventionTime/SUMMARY.json)
- [`../HumanInterventionTime/LOCAL-CONTEXT.json`](../HumanInterventionTime/LOCAL-CONTEXT.json)

Dream evidence neighborhoods:

- [`./SessionExcerpts/INDEX.json`](./SessionExcerpts/INDEX.json)

## Intended Deliverables

Pass 1 outputs:

- [`./APPROACH.md`](./APPROACH.md)
- [`./SessionExcerpts/INDEX.json`](./SessionExcerpts/INDEX.json)
- [`./BURDEN-ANALYSIS.md`](./BURDEN-ANALYSIS.md)

Pass 2 outputs:

- `ORTHOGONAL-SOLUTIONS-MATRIX.md`
- `Option-Tasks/`
- `README.md`

## Pass Sequence

1. Pass 1: evidence + burden analysis
2. Pass 2: matrix + option tasks + final conformance

## Evidence Surfaces To Reread When Making Decisions

- `../HumanInputEvents/SOURCE-PACKET.jsonl` for the human-authored packet seed text.
- `../HumanInputEvents/SOURCE-SESSIONS.json` for the authoritative transcript surface contract.
- The ordered `.parts` transcript files listed there (for the non-missing session).
- `../HumanNeeds/*` for the current needs triage and record (as derived, not as final truth).
- `../HumanInterventionTime/*` for cost signals (as prioritization input, not a full utility function).

