# Dream Replay Confusion Analysis

## Purpose

This note records why the first clean replay of the April 19 ThirdPerson packet regressed, what signs of confusion appeared in the subagent session log, and which durable process changes were added before the next replay.

Reviewed session log:

- `C:\Users\gregs\.codex\sessions\2026\04\21\rollout-2026-04-21T15-02-06-019db16b-d16a-7392-89c5-48bdbb1b3381.jsonl`

## Observed Confusion Signs

### 1. The worker had no authoritative seeded input packet

Observed:

- The worker started by searching raw `.codex/sessions` for likely ThirdPerson transcripts instead of starting from a durable seed.
- It said it was "narrowing to the ThirdPerson transcripts by their `cwd` and thread metadata" at line 53.
- It said it was "sampling the session heads directly" at line 61.

Why this mattered:

- That pushed packet scope selection back into live judgment.
- It created room for extra events, missing events, and packet-shape drift.

Durable fix:

- `C:\Users\gregs\.codex\Orchestration\Processes\INTERVENTION-REPORTS.md` now requires authoritative task-owned harvest seeds before promotion.
- `C:\Users\gregs\.codex\Orchestration\Prompts\Interventions\HumanInputEvents\WORKFLOW.md` now distinguishes seeded rebuild from bootstrap harvest and says canonical reruns must reuse the task-owned seed when it exists.
- `Tracking/Task-0007/HumanInputEvents/04/19/SOURCE-PACKET.jsonl` and `SOURCE-SESSIONS.json` are now restored and treated as the authoritative seed for this packet.

### 2. The worker had to infer the raw-transcript harvest rule

Observed:

- The worker said the original packet "appears to key each event off the `response_item` user line, not the duplicate `event_msg user_message` line" at line 113.

Why this mattered:

- The worker was discovering a packet rule that should already have been durable.
- That risked duplicate counting and arbitrary packet rewrites.

Durable fix:

- `C:\Users\gregs\.codex\Orchestration\Processes\INTERVENTION-REPORTS.md` now says raw transcript harvest must use `response_item` user messages and treat matching `event_msg.user_message` records as duplicate transport echoes.
- `C:\Users\gregs\.codex\Orchestration\Prompts\Interventions\HumanInputEvents\WORKFLOW.md` now repeats that rule and blocks preserving `EVENT-*.json` in the promoted canonical packet.

### 3. The worker used the wrong repo-local docs

Observed:

- The worker opened `C:\Agent\CodexDashboard\AGENTS.md` at line 116.
- It opened `C:\Agent\CodexDashboard\TESTING.md` at line 117.
- It opened `C:\Agent\CodexDashboard\REGRESSION.md` at line 118.

Why this mattered:

- The packet being rebuilt belonged to `C:\Agent\ThirdPerson`, not the current workspace repo.
- That contaminated local-context reconstruction and regression-lane grounding.

Durable fix:

- `C:\Users\gregs\.codex\Orchestration\Processes\INTERVENTION-REPORTS.md` now has a packet-repo grounding rule.
- `C:\Users\gregs\.codex\Orchestration\Prompts\Interventions\HumanNeeds\WORKFLOW.md` and the Dream prompt/workflow docs now say repo-local constraints must come from the packet repo, not the current workspace repo.

### 4. The worker had to make a new scope judgment mid-run

Observed:

- The worker said two extra records "likely fall outside the original packet's intended ThirdPerson incident window" and was deciding whether they were "out-of-scope tail chatter or true packet events" at line 189.

Why this mattered:

- Canonical reruns should not invent fresh scope rulings when authoritative seeds already exist.
- This is how the replay drifted from `62` events to `64`.

Durable fix:

- `C:\Users\gregs\.codex\Orchestration\Processes\INTERVENTION-REPORTS.md` now says if incident-scope judgment is still needed, stop and create a task-owned seed first rather than improvising inside the promoted packet.
- This packet now has the restored `62`-event authoritative seed in `Tracking/Task-0007/HumanInputEvents/04/19/`.

### 5. Intervention-time calibration was under-specified

Observed:

- The replay eventually reported packet-specific reconstruction defaults with `typing_rate_chars_per_second = 4.0` and `stall_grace_seconds = 300`.

Why this mattered:

- Those defaults were not durably justified.
- They collapsed stall attribution and sharply reduced measured intervention time.

Durable fix:

- `C:\Users\gregs\.codex\Orchestration\Prompts\Interventions\HumanInterventionTime\WORKFLOW.md` now sets the shared baseline to `typing_rate_chars_per_second = 1.0` and `stall_grace_seconds = 60` unless a stronger durable local override exists.
- `C:\Users\gregs\.codex\Orchestration\Prompts\Interventions\HumanInterventionTime\PROMPT.md` now repeats that baseline.
- `Tracking/Task-0007/HumanInterventionTime/04/19/LOCAL-CONTEXT.json` now carries the normalized baseline.

### 6. Dream output shape was too loose

Observed:

- The replay collapsed the old seven-burden solution set to six problems.
- It also wrote ad hoc names for some plan and task artifacts instead of keeping the prior numbering scheme.

Why this mattered:

- That made the rerun structurally different even before quality was judged.
- It also made the task-candidate packet thinner than the known-good reference.

Durable fix:

- `C:\Users\gregs\.codex\Orchestration\Processes\DREAMING.md` now pins the expected stable naming scheme for problem-option plans and winner tasks.
- `C:\Users\gregs\.codex\Orchestration\Prompts\Dream\WORKFLOW.md` and `PROMPT-PASS1.md` now say:
  - keep all seven burden types separate when the packet supports them
  - use `PROBLEM-<NNNN>-OPTION-<A|B|C>-PLAN-0001.md`
  - use `SOLUTION-TASK-<NNNN>.md`
  - keep numbering aligned across matrix, plans, and tasks

## Pre-Rerun Checklist

Before the next clean replay:

1. Use the restored task-owned `SOURCE-PACKET.jsonl` and `SOURCE-SESSIONS.json` exactly.
2. Use the normalized task-owned intervention-time `LOCAL-CONTEXT.json`.
3. Ground repo-local context in `C:\Agent\ThirdPerson`, not `C:\Agent\CodexDashboard`.
4. Do not create or preserve `EVENT-*.json` in the promoted canonical packet.
5. Require the Dream rerun to preserve all seven burden types if the packet still supports them separately.

## Outcome Expected From The Next Replay

If the durable fixes above are enough, the next replay should:

- return to `62` human input events
- recover stall attribution closer to the prior packet
- stop using CodexDashboard repo-local rules for a ThirdPerson packet
- regenerate a seven-problem Dream matrix with stable plan and task naming
- stay closer to the known-good reference without copying its prose
