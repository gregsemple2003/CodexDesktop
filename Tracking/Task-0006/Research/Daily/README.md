# Task 0006 Daily Capture Layout

Use one folder per source day for transcript-first incident mining artifacts.

This task-local folder is the bootstrap workspace for contract iteration, day-scoped recall, and later analysis passes.

The promoted shared intervention canon and eval layer now live at [Interventions](/c:/Users/gregs/.codex/Orchestration/Reports/Interventions/). Keep this task-local area focused on day-scoped recall, analysis, and synthesis lineage.

## Layout

```text
Tracking/Task-0006/Research/
  INCIDENT.schema.json
  INCIDENT-GOAL-STACK.md
  INCIDENT-EXAMPLE-0001.json
  INCIDENT-EXAMPLE-0002.json
  Daily/
    2026-04-03/
      INTERVENTION-PASS1.md
      INTERVENTION-PASS2.md
      INTERVENTION-PASS3.md
      DAILY-BRIEF.md
      INCIDENT-0001.json
      INCIDENT-0002.json
    2026-04-04/
      INTERVENTION-PASS1.md
      INTERVENTION-PASS2.md
      INTERVENTION-PASS3.md
      DAILY-BRIEF.md
      INCIDENT-0003.json
```

## Rules

- Keep schema files, contract notes, and hand-built examples at the top level of `Research/`.
- Keep dated capture artifacts under `Research/Daily/YYYY-MM-DD/`.
- Use the dated day folder for transcript-first investigation artifacts:
  - `INTERVENTION-PASS1.md` for exhaustive candidate recall
  - `INTERVENTION-PASS2.md` for per-event investigation
  - `INTERVENTION-PASS3.md` for clustering and local principle extraction
- Store each real incident JSON in the folder matching its `source_date`.
- A day folder may also contain one `DAILY-BRIEF.md` when a self-contained second-opinion packet is useful.
- Do not recreate deprecated `OUTBOUND-MESSAGE-REVIEW.csv` captures. Older bootstrap artifacts may still exist, but they are no longer part of the active contract.
- Within a day folder, `INTERVENTION-PASS1.md`, `INTERVENTION-PASS2.md`, and `INTERVENTION-PASS3.md` are always the authoritative files for later passes.
- If a later rerun becomes authoritative, rename it back to the canonical pass name and archive the replaced file as `INTERVENTION-PASS<N>-SUPERSEDED-YYYY-MM-DD.md`.
- Keep bounded verification artifacts that do not replace the canonical pass as auxiliary files such as `INTERVENTION-PASS1-SPOTCHECK-YYYY-MM-DD.md`.

## Canonical Status

Every day from `2026-03-28` through `2026-04-08` now has one authoritative `INTERVENTION-PASS3.md`.

| Day | Canonical status | Notes |
| --- | --- | --- |
| `2026-03-28` | original canonical kept | Current `INTERVENTION-PASS3.md` remains authoritative. |
| `2026-03-29` | original canonical kept | Current `INTERVENTION-PASS3.md` remains authoritative. |
| `2026-03-30` | original canonical kept | Current `INTERVENTION-PASS3.md` remains authoritative. |
| `2026-03-31` | original canonical kept | Current `INTERVENTION-PASS3.md` remains authoritative. |
| `2026-04-01` | original canonical kept | No local session directory existed for the day; the empty-stack `INTERVENTION-PASS3.md` is the honest canonical record. |
| `2026-04-02` | rerun promoted to canonical | Earlier pass files archived as `INTERVENTION-PASS{1,2,3}-SUPERSEDED-2026-04-09.md`. |
| `2026-04-03` | original canonical kept | Current `INTERVENTION-PASS3.md` remains authoritative. |
| `2026-04-04` | original canonical kept | Spot-check artifacts from `2026-04-09` are auxiliary only and did not replace the main stack. |
| `2026-04-05` | original canonical kept | Current `INTERVENTION-PASS3.md` remains authoritative. |
| `2026-04-06` | explanatory rerun promoted to canonical | Earlier pass files archived as `INTERVENTION-PASS{1,2,3}-SUPERSEDED-2026-04-09.md`. |
| `2026-04-07` | original canonical kept | Current `INTERVENTION-PASS3.md` remains authoritative. |
| `2026-04-08` | original canonical kept | Current `INTERVENTION-PASS3.md` remains authoritative. |

## Daily Brief

`DAILY-BRIEF.md` is a derived text report for second-opinion review.

Use it when you want one self-contained file that gathers:

- the accepted incidents for the day
- the nearby session context around those incidents
- the main boundary calls, weaknesses, or open doubts

Treat it as a review packet assembled from accepted incident JSON files and transcript evidence.

Self-contained means a reviewer may have only the brief. So the brief should inline the incident body in readable form, including the fields needed to judge whether the incident is well-formed, rather than assuming the reviewer will open adjacent JSON files.

The accepted incident JSON itself should already be heavyweight:

- `verbatim_transcript_windows`

The brief should reuse that heavy incident body verbatim instead of rebuilding a lighter summary from scratch.
Do not strip the incident JSON down and move the real transcript only into the brief; the brief supplements the heavyweight incident record.

For second-opinion use, optimize for epistemic warrant rather than compression:

- include larger contiguous and relevant transcript windows, not only selected supporting lines
- keep the brief's main evidence sections verbatim-only rather than paraphrased
- include a chronological verbatim timeline of the relevant human outbound messages when that helps memory recovery or incident-boundary review
- preserve enough producer trajectory, human correction, and immediate aftermath that another reviewer can reason from first principles
- separate raw source context from your own interpretation, adequacy judgment, or boundary notes
- use a soft size ceiling of about `200 KB`, but fill that budget with raw high-value evidence rather than a pre-digested summary
- if tradeoffs are needed under that cap, cut interpretation before cutting verbatim human messages or contiguous dialogue blocks

It should not become the source of truth or silently redefine the accepted set.

## Intent

This layout separates three layers cleanly:

- transcript-first day analysis in the task-local pass artifacts
- accepted incident records in schema-backed JSON
- optional self-contained review packaging in `DAILY-BRIEF.md`

That keeps the official corpus small and durable while still preserving enough local analysis to learn from near-misses, boundary corrections, and human-model signals.
