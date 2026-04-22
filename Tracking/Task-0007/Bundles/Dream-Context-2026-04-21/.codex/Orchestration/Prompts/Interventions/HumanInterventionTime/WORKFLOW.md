# Run Intervention Time v0

## Objective

Measure `InterventionTime` for one canonical repo-day packet and keep all outputs in the `HumanInterventionTime` folder for that day.

Use these placeholders:

- `<DAY_ROOT>`
  - canonical repo-day root such as `C:\Users\gregs\.codex\Orchestration\Reports\Interventions\<repo>\<YYYY>\<MM>\<DD>`

This workflow is intentionally separate from:

- human-input harvest
- needs extraction

## Local Context Seed

Preferred source for packet-local `LOCAL-CONTEXT.json`:

- a task-owned seed such as `Tracking/Task-<id>/HumanInterventionTime/<MM>/<DD>/LOCAL-CONTEXT.json`

Rules:

- if a task-owned local-context seed exists, use it exactly
- otherwise create the packet-local `LOCAL-CONTEXT.json` from the shared baseline in this workflow
- do not invent packet-specific measurement parameters without durable justification

## Dispatch Intent

This packet is designed to be launched as its own worker from the same daily source-discovery root as harvest.

The worker is safe to run independently because it:

- reads immutable input artifacts and raw session transcripts
- writes only inside this folder
- does not modify `HumanInputEvents/`
- does not write into `HumanNeeds/`

## Allowed Inputs Only

You may read only:

- `C:\Users\gregs\.codex\Orchestration\Prompts\Interventions\HumanInterventionTime\PROMPT.md`
- `<DAY_ROOT>\HumanInterventionTime\LOCAL-CONTEXT.json`
- `<DAY_ROOT>\HumanInputEvents\SOURCE-PACKET.jsonl`
- `<DAY_ROOT>\HumanInputEvents\SOURCE-SESSIONS.json`
- each `session_file` path explicitly listed in `SOURCE-SESSIONS.json`

Do not read any other files.
Do not search the repo.
Do not browse the web.
Do not infer missing context from memory.

## Output Directory

Write all outputs into:

- `<DAY_ROOT>\HumanInterventionTime\`

Do not write outside that directory.

## Required Outputs

Create:

- `EVENT-MEASUREMENTS.jsonl`
- `STALL-EVENTS.json`
- `SUMMARY.json`

Follow the contracts in [PROMPT.md](./PROMPT.md).

## Additional Rules

- preserve the event order from `SOURCE-PACKET.jsonl`
- use `parsed_envelope.request_text` when present to estimate `manual_chars`
- use the configured `typing_rate_chars_per_second` and `stall_grace_seconds` from `LOCAL-CONTEXT.json`
- if `LOCAL-CONTEXT.json` is being created for this packet and no task-specific override is available, use the shared baseline:
  - `typing_rate_chars_per_second = 1.0`
  - `stall_grace_seconds = 60`
- do not invent faster typing baselines or looser stall grace windows without durable justification in `LOCAL-CONTEXT.json`
- charge `stall_loss_seconds` only when the human input is best classified as `wake_up`
- if the last assistant-visible state was a pause gate, approval gate, or other explicit wait state, do not classify the event as `wake_up`
- if a human message mixes a prod with substantial new work content, prefer `correction` or `new_request` and set `stall_loss_seconds = 0` unless the idle-drop evidence is overwhelming

## Classification Review

Before you finalize the outputs, re-check `intervention_kind` using this decision ladder:

1. `wake_up`
   - the human message is primarily a prod to resume or continue after AI inactivity
   - the prior AI state was not an explicit wait gate
2. `answer_to_question`
   - the primary human burden is that the agent did not answer an explicit question directly or in the requested answer shape
3. `boundary_reset`
   - the primary human burden is that the agent crossed or ignored a lane, scope, method, ownership, or stop/go boundary
4. `correction`
   - the primary human burden is that the agent made a wrong claim, showed wrong evidence, misread the state, or otherwise needs factual or behavioral repair
5. `new_request`
   - the message primarily opens a new work item or approves the next distinct step
6. `other`
   - use only when none of the above are the simplest defensible fit

When a message matches more than one label:

- prefer `answer_to_question` over `correction` when the core burden is `you did not answer the question`
- prefer `boundary_reset` over `correction` when the core burden is `you violated the allowed lane, scope, method, or ownership boundary`
- prefer `correction` over `new_request` when the message mainly repairs the agent's mistake rather than opening fresh work
- prefer `wake_up` only when resume pressure is the main point and the message does not mostly function as one of the other labels

Before you stop:

- if `other` exceeds `10%` of events, re-review the classifications
- if `other` is larger than either `correction` or `answer_to_question`, re-review the classifications
- if you keep a large `other` bucket anyway, explain why in `SUMMARY.json`

## Final Response

At the end, report:

- exact output files written
- total event count measured
- stall event count
- total estimated intervention seconds
