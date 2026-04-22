# Intervention Reports Process

Use this doc for the shared structure and promotion rules for canonical intervention day packets.

Keep shared process docs in `Processes/`.
Keep shared prompt bundles in `Prompts/`.
Keep output packets only in `Reports/Interventions/`.

## Purpose

Use the intervention report flow to produce one canonical repo-day packet that captures:

- raw human input evidence
- derived human-needs analysis
- derived human-intervention-time analysis

Use the optional `Dream` annex only when the human wants deeper solution design for one packet.

## Shared Homes

- packet outputs:
  - `C:\Users\gregs\.codex\Orchestration\Reports\Interventions\<repo>\<YYYY>\<MM>\<DD>\`
- shared packet prompt bundles:
  - [../Prompts/Interventions/README.md](../Prompts/Interventions/README.md)
- optional Dream process:
  - [./DREAMING.md](./DREAMING.md)
- optional Dream prompt set:
  - [../Prompts/Dream/WORKFLOW.md](../Prompts/Dream/WORKFLOW.md)
  - [../Prompts/Dream/PROMPT-PASS1.md](../Prompts/Dream/PROMPT-PASS1.md)

## Canonical Packet Shape

```text
<InterventionsRoot>/<repo>/<YYYY>/<MM>/<DD>/
  HumanInputEvents/
  HumanNeeds/
  HumanInterventionTime/
  DAY-MANIFEST.json
  INDEX.md
```

The canonical repo-day packet should:

- be day-first for browsing
- be repo-scoped for reuse across repos
- be canonical rather than run-id indexed
- overwrite the prior canonical packet on rerun for the same repo-day

## Harvest Seed Rule

Canonical reruns should prefer an authoritative task-owned seed.

Preferred seed home:

- `Tracking/Task-<id>/HumanInputEvents/<MM>/<DD>/`
  - `SOURCE-PACKET.jsonl`
  - `SOURCE-SESSIONS.json`

Rules:

- if an authoritative task-owned seed exists, use it exactly
- do not widen or narrow the packet by doing a fresh raw-transcript search during promotion
- if no authoritative seed exists, create it under the owning task first, then promote the canonical day packet
- do not improvise a promoted packet directly from a broad raw-transcript search

Bootstrap-harvest rules when a seed must be created:

- start from explicit session ids or session files named by durable task artifacts when possible
- if raw transcript harvest is still required, use `response_item` records where:
  - `payload.type = message`
  - `payload.role = user`
- treat matching `event_msg.user_message` records as duplicate transport echoes, not separate events
- if event inclusion still depends on incident-scope judgment beyond the resolved local date, stop and capture that judgment in the task-owned seed before promotion

The canonical day packet is the promoted output.
The task-owned seed is the durable bootstrap input.

## Domain Folders

### `HumanInputEvents/`

Canonical contents:

- `SOURCE-PACKET.jsonl`
- `SOURCE-SESSIONS.json`
- `INDEX.json`
- `PROMPT-INPUT.md`
- `LOCAL-CONTEXT.json` when needed

Explicit decisions:

- do not preserve `EVENT-*.json` in the promoted canonical packet
- do not preserve per-event Markdown browse files in the promoted canonical packet

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

Keep the reusable recipe outside the day packet.

Use this split:

- shared `PROMPT.md` or `WORKFLOW.md`
  - source-controlled recipe in `Prompts/Interventions/`
- packet-local `PROMPT-INPUT.md`
  - exact batch-specific invocation record inside the canonical day packet

## Repo-Local Context Rule

When any packet pass needs repo-local docs or local constraints:

- use the packet repo, not the current workspace repo
- resolve the target repo from durable packet inputs such as `repo_ref`, `SOURCE-PACKET.jsonl`, or `SOURCE-SESSIONS.json`
- prefer these repo-local docs when they exist:
  - `<repo_ref>/AGENTS.md`
  - `<repo_ref>/REGRESSION.md`
  - `<repo_ref>/TESTING.md`
- if one of those docs is missing, record the absence explicitly instead of substituting a different repo's docs

## Promotion Rules

- keep task-local bootstrap work under the owning task until the shape is stable
- promote only canonical day-packet outputs into `Reports/Interventions/`
- promote shared packet process rules into `Processes/`
- promote shared packet prompt bundles into `Prompts/Interventions/`
- do not keep shared prompt bundles or shared workflow docs under `Reports/Interventions/`

## Intervention-Time Reading

Current interpretation rules:

- `InterventionTime` is the default rough proxy for human cost
- it is a first approximation for prioritization, not a complete utility function
- keep event harvest, needs analysis, and intervention-time measurement as separate passes unless a later workflow proves the merge is both truer and easier to audit

## Current Example

The first promoted canonical example is:

- [../Reports/Interventions/ThirdPerson/2026/04/19/INDEX.md](../Reports/Interventions/ThirdPerson/2026/04/19/INDEX.md)
