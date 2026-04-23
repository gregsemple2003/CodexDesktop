# Dream Approach

## Method

This rerun follows the current Dream workflow only.

1. Read the durable Dream/process docs first.
2. Read the current April 19 packet inputs.
3. Validate that `HumanInputEvents/INDEX.json` is strictly chronological before analysis.
4. Use `HumanInputEvents/INDEX.json` and `HumanNeeds/*` to find recurring burden clusters.
5. Follow those clusters into:
   - `HumanInputEvents/SOURCE-PACKET.jsonl`
   - `HumanInputEvents/SOURCE-SESSIONS.json`
   - the ordered-parts transcript for the lead session where the packet exposes it
6. Treat the first session as packet-traced only where `transcript_surface = "missing"` and avoid ad hoc transcript discovery.
7. Keep exact `source_event_ids` visible in every Dream artifact that makes a causal or prioritization claim.

## Scope

This annex is limited to the April 19, 2026 `ThirdPerson` intervention packet and the current durable Dream/process docs.

Excluded on purpose:

- prior Dream outputs
- prior reruns
- reference folders
- inherited chat context outside the packet and its cited source surfaces

## Evidence Base

Primary packet inputs:

- [../INDEX.md](../INDEX.md)
- [../DAY-MANIFEST.json](../DAY-MANIFEST.json)
- [../HumanInputEvents/INDEX.json](../HumanInputEvents/INDEX.json)
- [../HumanInputEvents/SOURCE-PACKET.jsonl](../HumanInputEvents/SOURCE-PACKET.jsonl)
- [../HumanInputEvents/SOURCE-SESSIONS.json](../HumanInputEvents/SOURCE-SESSIONS.json)
- [../HumanNeeds/PACKET-TRIAGE.json](../HumanNeeds/PACKET-TRIAGE.json)
- [../HumanNeeds/PACKET-RECORD.json](../HumanNeeds/PACKET-RECORD.json)
- [../HumanInterventionTime/EVENT-MEASUREMENTS.jsonl](../HumanInterventionTime/EVENT-MEASUREMENTS.jsonl)
- [../HumanInterventionTime/SUMMARY.json](../HumanInterventionTime/SUMMARY.json)

Source-following notes:

- The packet exposes one session with `transcript_surface = "ordered_parts"` and one with `transcript_surface = "missing"`.
- For the ordered-parts session, this rerun used the packet-named April 19 chunks directly.
- For the missing-surface session, this rerun stayed with the packet-traced user turns from `SOURCE-PACKET.jsonl`.

## Intended Deliverables

Pass 1 owns:

- [./BURDEN-ANALYSIS.md](./BURDEN-ANALYSIS.md)
- [./SessionExcerpts/INDEX.json](./SessionExcerpts/INDEX.json)

Pass 2 will own:

- [./ORTHOGONAL-SOLUTIONS-MATRIX.md](./ORTHOGONAL-SOLUTIONS-MATRIX.md)
- [./Option-Tasks/INDEX.md](./Option-Tasks/INDEX.md)
- [./README.md](./README.md)

## Later Pass Sequence

Pass 2 should reread the raw packet surfaces and the same source neighborhoods before deciding winners.

The key pass-2 obligations are:

1. keep burden drivers separate unless one mechanism and one proof bar honestly cover the merge
2. turn every matrix option into a task-shaped proposal
3. carry exact `source_event_ids` into each downstream problem row and burden-reduction task
4. prefer enforcement boundaries and durable artifacts over reminder-only fixes when the burden is structural

## Raw Surfaces To Reread In Pass 2

Pass 2 should reread these before selecting winners:

- [../HumanInputEvents/INDEX.json](../HumanInputEvents/INDEX.json)
- [../HumanInputEvents/SOURCE-PACKET.jsonl](../HumanInputEvents/SOURCE-PACKET.jsonl)
- [../HumanInputEvents/SOURCE-SESSIONS.json](../HumanInputEvents/SOURCE-SESSIONS.json)
- [./SessionExcerpts/INDEX.json](./SessionExcerpts/INDEX.json)
- [./BURDEN-ANALYSIS.md](./BURDEN-ANALYSIS.md)

The most decision-relevant source neighborhoods are:

- the missing-surface first-session user corrections around default-lane proof, approval surface, direct answers, and runtime-evidence honesty
- the ordered-parts lead-session turns around lines `4429-4433`, `4684-4756`, `4919-4920`, `6555-6559`, `6755-6756`, `7035-7039`, and `7446-7447`

## Validation Result

`HumanInputEvents/INDEX.json` was rechecked before writing and is strictly ascending by `captured_at_local`.

