# Run HumanInputEvents

## Objective

Prepare or reuse the authoritative `HumanInputEvents` seed for one canonical repo-day packet, then run the shared ingest prompt without changing packet scope.

Use these placeholders:

- `<DAY_ROOT>`
  - canonical repo-day root such as `C:\Users\gregs\.codex\Orchestration\Reports\Interventions\<repo>\<YYYY>\<MM>\<DD>`
- `<TASK_ROOT>`
  - owning task root such as `C:\Agent\CodexDashboard\Tracking\Task-0007`

## Two Modes

### 1. Seeded Rebuild

Use this when an authoritative task-owned seed already exists.

Preferred seed home:

- `<TASK_ROOT>\HumanInputEvents\<MM>\<DD>\SOURCE-PACKET.jsonl`
- `<TASK_ROOT>\HumanInputEvents\<MM>\<DD>\SOURCE-SESSIONS.json`

Rules:

- if these seed files exist, treat them as authoritative
- do not re-harvest from raw transcripts just because the promoted packet was deleted
- copy or recreate the seed files into `<DAY_ROOT>\HumanInputEvents\` first
- then run [PROMPT.md](./PROMPT.md) to enrich `SOURCE-PACKET.jsonl` and build `INDEX.json`

### 2. Bootstrap Harvest

Use this only when no authoritative seed exists yet.

Write the new seed under the owning task first, not directly under the promoted reports path.

Allowed durable bootstrap sources:

- task-owned handoff or research artifacts that name the exact sessions or local date
- prior task-owned harvest directories
- prior worker session logs that print the authoritative `SOURCE-SESSIONS.json` or `SOURCE-PACKET.jsonl`
- the original session transcripts named by those artifacts
- `C:\Users\gregs\.codex\session_index.jsonl`

Rules:

- start from explicit session ids or session files named by durable artifacts when possible
- if you must read raw transcripts, harvest only `response_item` records where:
  - `payload.type = message`
  - `payload.role = user`
- treat matching `event_msg.user_message` records as duplicate echoes, not separate events
- preserve the authoritative line number for the chosen `response_item` record
- sort events by ascending `captured_at_local`
- stop if packet scope still depends on a new judgment call beyond the resolved local date
- if you had to make a scope choice, record it in the task-owned seed notes before any promotion

## Promotion Rules

After the seed exists:

1. place `SOURCE-PACKET.jsonl` and `SOURCE-SESSIONS.json` in `<DAY_ROOT>\HumanInputEvents\`
2. run [PROMPT.md](./PROMPT.md)
3. keep only the promoted canonical outputs in the day packet:
   - `SOURCE-PACKET.jsonl`
   - `SOURCE-SESSIONS.json`
   - `INDEX.json`
   - `PROMPT-INPUT.md`
   - optional `LOCAL-CONTEXT.json`

Do not promote:

- `EVENT-*.json`
- `EVENT-*.md`
- task-owned bootstrap notes that belong under `Tracking/Task-<id>/`

## Final Check

Before you stop, confirm:

1. the promoted packet scope came from an authoritative seed rather than a fresh broad search
2. no duplicate `event_msg.user_message` records were counted as separate events
3. the promoted day packet does not contain `EVENT-*.json`
