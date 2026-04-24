# Task 0007 Handoff

## Current Status

`Task-0007` is complete and should now be treated as a closed research-and-reference task.

What this task now owns durably:

- the task-owned research, reference packets, and comparison material for Jarvis-style intervention analysis
- the recovered Dream-process design discussion and the reference packet backups
- the task-local research that shaped the current shared Dream writer and auditor workflow

Follow-on product work has been split out into separate tasks:

- [Task-0008](../Task-0008/TASK.md): the separate dispatch layer and durable execution-state contract
- [Task-0009](../Task-0009/TASK.md): the dashboard `Tasks` tab as the humane dispatch and monitoring surface
- [Task-0010](../Task-0010/TASK.md): the daily Dream run, digest, and option-task promotion flow

The earlier [Task-0006](../Task-0006/TASK.md) lane is now cancelled as superseded by this research-and-reference baseline.

This handoff remains valuable as historical context, but it is no longer the active implementation queue.

## Historical Objective At Close

The current problem is:

- regenerate the canonical April 19 `ThirdPerson` intervention packet with equal or better quality than the real reference packet at [pre-durable-replay](/c:/Agent/CodexDashboard/Tracking/Task-0007/Reference/ThirdPerson-2026-04-19-pre-durable-replay/ThirdPerson/2026/04/19/INDEX.md)

That is the reference.

Not this:

- `pre-clean-pass12-rerun`
- any later backup made during replay experiments

If a future session compares against the wrong reference, the comparison is invalid.

## Current Live State

Live packet root:

- [April 19 packet](/c:/Users/gregs/.codex/Orchestration/Reports/Interventions/ThirdPerson/2026/04/19/INDEX.md)

Current live rerun scope:

- `HumanInputEvents`
- `HumanNeeds`
- `HumanInterventionTime`

`Dream` was intentionally not rebuilt in the latest rerun and is currently absent from the live packet.

## What Was Learned

### 1. Clean rerun discipline is not optional

The human explicitly wants replay experiments to avoid confounding variables.

That means:

- wipe the live packet first
- use one fresh clean worker for the rerun
- do not synthesize the packet from accumulated local context
- use only durable markdown docs under:
  - `C:\Users\gregs\.codex\Orchestration\Processes\`
  - `C:\Users\gregs\.codex\Orchestration\Prompts\Interventions\`
  - `C:\Users\gregs\.codex\Orchestration\Prompts\Dream\`
- preserve the authoritative task-owned seed
- do not widen scope with ad hoc local search

If a rerun preserves stale packet content in place, the packet becomes internally inconsistent and the run is corrupted.

### 2. The real reference packet was created on April 21

The first reference packet created under `Tracking/Task-0007/Reference/` was:

- [pre-durable-replay](/c:/Agent/CodexDashboard/Tracking/Task-0007/Reference/ThirdPerson-2026-04-19-pre-durable-replay/ThirdPerson/2026/04/19/INDEX.md)

That reference was created in session:

- [rollout-2026-04-20T23-46-50-019dae25-dd3d-7ee3-8888-613be8428412.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/20/rollout-2026-04-20T23-46-50-019dae25-dd3d-7ee3-8888-613be8428412.jsonl#L2489)

### 3. One session is missing; one survives through `.parts`

The packet depends on two source sessions:

- `019da198-ea2f-7421-8adc-c566da0b6121`
  - thread name: `Implement task 6 workflow`
  - raw transcript file is missing
  - extracted human events still exist in the task-owned seed packet
- `019da19b-19c6-7e20-bdd8-f03aea761fb9`
  - thread name: `Lead Task-0006 lifecycle`
  - raw full `.jsonl` is gone
  - ordered `.parts` files exist and are now usable durably

Important distinction:

- `019da198...` existed and is represented in the seed
- the raw transcript file itself is missing on disk
- so early event judgments still lack surrounding transcript context

### 4. Ordered `.parts` support is now durable

The replay system was updated so `SOURCE-SESSIONS.json` can explicitly encode:

- `transcript_surface = "session_file"`
- `transcript_surface = "ordered_parts"`
- `transcript_surface = "missing"`

The live and task-owned `SOURCE-SESSIONS.json` for April 19 now name the ordered `.parts` for `019da19b...`.

This means future clean reruns can consume the surviving lifecycle transcript deterministically without filesystem guessing.

### 5. `wake_up` had been suppressed incorrectly

The main human burden in the reference packet includes restart-supervision cost:

- the human had to wake work back up
- the human had to tell the system to continue supervising instead of handing homework back

The newer `HumanInterventionTime` rules had become too conservative:

- if a `continue` message also contained ownership, correction, or boundary language
- it got reclassified as `boundary_reset` or `correction`
- stall loss was zeroed out

That suppressed the key restart-supervision burden.

This has now been fixed durably in:

- [PROMPT.md](/c:/Users/gregs/.codex/Orchestration/Prompts/Interventions/HumanInterventionTime/PROMPT.md#L88)
- [WORKFLOW.md](/c:/Users/gregs/.codex/Orchestration/Prompts/Interventions/HumanInterventionTime/WORKFLOW.md#L92)

Current rule:

- if the human is clearly restarting dropped work after idle time
- and the prior AI state was not an explicit wait gate
- then `wake_up` wins even if the same message also carries correction, boundary, or ownership content

## Latest Clean Non-Dream Rerun

Latest clean worker:

- `Darwin`
- thread id: `019db7a2-c997-7441-85d6-3d3ae679f4db`

Latest live outputs:

- [INDEX.md](/c:/Users/gregs/.codex/Orchestration/Reports/Interventions/ThirdPerson/2026/04/19/INDEX.md#L1)
- [HumanNeeds/PACKET-TRIAGE.json](/c:/Users/gregs/.codex/Orchestration/Reports/Interventions/ThirdPerson/2026/04/19/HumanNeeds/PACKET-TRIAGE.json#L1)
- [HumanInterventionTime/SUMMARY.json](/c:/Users/gregs/.codex/Orchestration/Reports/Interventions/ThirdPerson/2026/04/19/HumanInterventionTime/SUMMARY.json#L1)

Key current numbers:

- `62` events
- `3` `wake_up` events
- `2422.637` stall-loss seconds
- `7866.387` total intervention seconds

Current non-Dream quality against the real reference:

- `HumanInputEvents`: better
- `HumanNeeds`: better
- `HumanInterventionTime`: still mixed

Why `HumanInterventionTime` is still mixed:

- it now captures restart-supervision burden better than before
- but it still does not match the reference packet’s total burden shape
- the missing `019da198...` transcript still limits early-day judgments
- the current baseline uses `typing_rate_chars_per_second = 4.0`, while the old reference used `1.0`

## Important Comparison Rules

When comparing current packet vs reference:

- compare against [pre-durable-replay](/c:/Agent/CodexDashboard/Tracking/Task-0007/Reference/ThirdPerson-2026-04-19-pre-durable-replay/ThirdPerson/2026/04/19/INDEX.md#L1)
- compare only like-for-like scope
- do not call a rerun “better” just because it is newer
- do not let stale `Dream` content contaminate the comparison

Meaning:

- if only non-Dream passes are rerun, compare non-Dream quality only
- if the whole packet is judged, all packet domains must be regenerated coherently

## Current Main Gaps

The next session should assume these gaps remain:

- session `019da198...` raw transcript is still missing
- image payload bytes do not survive in the seeded packet
- `need_tag` is still packet-local inference, not a shared ontology
- the time model still needs judgment about whether the current baseline reflects the human’s real burden well enough

## What The Human Cares About

Keep this frame explicit:

- truth-seeking
- compassion for the human burden
- tolerance for imperfect expression of frustration

And this metric frame:

- repeated direct human input is failure telemetry
- the system should minimize human intervention burden
- restart-supervision burden is real human cost

Do not frame the human as “changing the plan” when the better reading is:

- the system failed to hold the correct boundary
- the system claimed too much too early
- the human had to repair the task frame

## Recommended Next Step

The next honest step is:

1. compare the current non-Dream rerun against the reference non-Dream packet carefully
2. decide whether `HumanInterventionTime` is now good enough or still needs another durable rule change
3. only after non-Dream quality is acceptable, run a clean full-packet rebuild including `Dream`

If full-packet quality is judged:

- wipe the whole live packet first
- rebuild all domains cleanly
- do not preserve old `Dream/` in place

## Watchouts

- do not use `pre-clean-pass12-rerun` as the quality reference
- do not compare against stale mixed-state packets
- do not self-synthesize reruns from local memory
- do not say a raw transcript “exists” when only extracted seed events exist
- do not use jargon like `transcript neighborhoods`; say `surrounding transcript context`

## Dream Process Thread Lookup

This handoff now also records where the Dream pass split was actually designed.

The important distinction is:

- the design-origin discussion happened in `Review Task-0007 handoff`
- `Run Dream pass 2A and 2B` was a later spawned worker thread that executed under that parent

Primary historical thread:

- thread name: `Review Task-0007 handoff`
- session id: `019db7be-6d89-7682-9924-51f719d86d10`
- session index entry: [/c:/Users/gregs/.codex/session_index.jsonl](/c:/Users/gregs/.codex/session_index.jsonl#L398)
- transcript: [/c:/Users/gregs/.codex/sessions/2026/04/22/rollout-2026-04-22T20-30-03-019db7be-6d89-7682-9924-51f719d86d10.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/22/rollout-2026-04-22T20-30-03-019db7be-6d89-7682-9924-51f719d86d10.jsonl#L10)

Key turns inside that parent thread:

- the explicit `Pass 2A` and `Pass 2B` split proposal is at [/c:/Users/gregs/.codex/sessions/2026/04/22/rollout-2026-04-22T20-30-03-019db7be-6d89-7682-9924-51f719d86d10.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/22/rollout-2026-04-22T20-30-03-019db7be-6d89-7682-9924-51f719d86d10.jsonl#L2424)
- the human follow-up pushing the adversarial `SOLUTION-CREATE.md` and `SOLUTION-AUDIT.md` framing is at [/c:/Users/gregs/.codex/sessions/2026/04/22/rollout-2026-04-22T20-30-03-019db7be-6d89-7682-9924-51f719d86d10.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/22/rollout-2026-04-22T20-30-03-019db7be-6d89-7682-9924-51f719d86d10.jsonl#L2430)

Later worker thread:

- thread name: `Run Dream pass 2A and 2B`
- session id: `019db910-dc58-7d30-b0aa-bd67fe3617e0`
- transcript: [/c:/Users/gregs/.codex/sessions/2026/04/23/rollout-2026-04-23T02-39-43-019db910-dc58-7d30-b0aa-bd67fe3617e0.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/23/rollout-2026-04-23T02-39-43-019db910-dc58-7d30-b0aa-bd67fe3617e0.jsonl#L9)
- parent-thread linkage proving it was spawned under `Review Task-0007 handoff` is at [/c:/Users/gregs/.codex/sessions/2026/04/23/rollout-2026-04-23T02-39-43-019db910-dc58-7d30-b0aa-bd67fe3617e0.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/23/rollout-2026-04-23T02-39-43-019db910-dc58-7d30-b0aa-bd67fe3617e0.jsonl#L1)

What this means for future discussion:

- if the goal is to understand where the Dream split was designed, start from `Review Task-0007 handoff`
- if the goal is to inspect the later execution pass, use `Run Dream pass 2A and 2B`
- do not mistake the worker thread for the origin of the `2A` and `2B` design

Current continuation frame from this lookup thread:

- no durable Dream docs were changed in this lookup thread
- this thread only recovered the historical source for the Dream pass split and the adversarial solution-audit follow-up
- the next honest discussion can start from whether the current Dream workflow should stay split as `2A` and `2B` only, or remain the later `2A` / `2B` / `2C` design that separates solution design, winner synthesis, and winner-task drafting

Later recovery thread that wrote this lookup down:

- thread name: `Locate Dream process thread`
- session id: `019dbac2-19f9-75f0-b04b-3accc4fa08d3`
- session index entry: [/c:/Users/gregs/.codex/session_index.jsonl](/c:/Users/gregs/.codex/session_index.jsonl#L408)
- transcript: [/c:/Users/gregs/.codex/sessions/2026/04/23/rollout-2026-04-23T10-32-56-019dbac2-19f9-75f0-b04b-3accc4fa08d3.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/23/rollout-2026-04-23T10-32-56-019dbac2-19f9-75f0-b04b-3accc4fa08d3.jsonl#L6)
- purpose: recover the design-origin thread and distinguish it from the later worker thread
- no `spawn_agent`, `wait_agent`, `send_input`, `resume_agent`, or `close_agent` calls were found in that recovery transcript during this lookup, so there is no separate child agent from that recovery session to stop now
