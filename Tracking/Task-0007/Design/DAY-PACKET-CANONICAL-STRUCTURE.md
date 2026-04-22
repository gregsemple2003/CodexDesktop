# Day Packet Canonical Structure

## Status

Approved on April 20, 2026 as the next target structure for generalizing intervention artifacts across repos and days.

This note records the approved shape. The April 19, 2026 `ThirdPerson` packet has now been promoted into the shared intervention reports root. This note does not backfill older packets or historical days.

## Intent

Use one canonical repo-day packet as the authoritative home for:

- raw human input evidence
- derived human-needs analysis
- derived human-intervention-time analysis

The packet should be:

- day-first for browsing
- repo-scoped for cross-repo generalization
- canonical rather than run-id indexed
- auditable without duplicating prompt recipes into every day folder

## Canonical Day Layout

Target shape:

```text
<InterventionsRoot>/<repo>/<YYYY>/<MM>/<DD>/
  HumanInputEvents/
  HumanNeeds/
  HumanInterventionTime/
  DAY-MANIFEST.json
  INDEX.md
```

The canonical repo-day packet should not contain a run id in its path.

Rerunning the packet for the same repo-day should overwrite the previous canonical outputs rather than leaving multiple sibling candidates.

## Domain Folders

### `HumanInputEvents/`

Canonical contents:

- `SOURCE-PACKET.jsonl`
- `SOURCE-SESSIONS.json`
- `INDEX.json`
- `PROMPT-INPUT.md`
- `LOCAL-CONTEXT.json` when needed for the batch

Explicit decision:

- do not preserve `EVENT-*.json` as a required browse/cache layer in the generalized structure

### `HumanNeeds/`

Canonical contents:

- `PACKET-TRIAGE.json`
- `PACKET-RECORD.json`
- `REPRESENTATIVE-EVENTS.json`
- `PROMPT-INPUT.md`
- `LOCAL-CONTEXT.json`

### `HumanInterventionTime/`

Canonical contents:

- `EVENT-MEASUREMENTS.jsonl`
- `STALL-EVENTS.json`
- `SUMMARY.json`
- `PROMPT-INPUT.md`
- `LOCAL-CONTEXT.json`

## Prompt Split

The reusable recipe should not be duplicated into every day packet.

Use this split:

- `PROMPT.md`
  - canonical, source-controlled recipe
  - stored in the shared intervention prompt home
- `PROMPT-INPUT.md`
  - exact instantiated input used to generate the current canonical batch
  - stored inside the repo-day packet

`PROMPT-INPUT.md` is the historical invocation record for the current batch.

## Shared Prompt Home

Keep reusable recipe bundles under the shared intervention prompt home:

```text
C:\Users\gregs\.codex\Orchestration\Prompts\Interventions\
  HumanInputEvents\
  HumanNeeds\
  HumanInterventionTime\
```

These prompt homes may contain more than just the final prompt text. They can also hold:

- supporting notes
- workflow contracts
- schemas
- stable examples

## Day-Root Files

### `DAY-MANIFEST.json`

Keep this small and machine-oriented.

Minimum intended role:

- identify the repo and local date
- identify the three domain folders
- identify which canonical outputs are present
- identify when the packet was last regenerated

### `INDEX.md`

Keep this human-oriented.

It should act as the day-root browse entry for:

- `HumanInputEvents/`
- `HumanNeeds/`
- `HumanInterventionTime/`
- the canonical outputs that matter most for a quick review

## Explicit Decisions

Approved decisions at this stage:

- use a canonical repo-day packet rather than a run-id-indexed path
- store reusable recipes under shared `Prompts`
- store batch-specific invocation text as `PROMPT-INPUT.md`
- do not add `GENERATION.json` for now
- do not keep `EVENT-*.json` in the generalized structure
- do not backfill older repo-days until the structure has been validated on current data

## Deferred

Still deferred:

- exact `DAY-MANIFEST.json` field contract beyond the current v0 fields
- whether any additional repo-day metadata should be required before broader rollout
