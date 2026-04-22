# Run Analysis

## Objective

Use the shared `HumanNeeds` prompt recipe to analyze one canonical repo-day packet and keep all outputs in the `HumanNeeds` folder for that day.

Use these placeholders:

- `<DAY_ROOT>`
  - canonical repo-day root such as `C:\Users\gregs\.codex\Orchestration\Reports\Interventions\<repo>\<YYYY>\<MM>\<DD>`

## Allowed Inputs Only

You may read only:

- `C:\Users\gregs\.codex\Orchestration\Prompts\Interventions\HumanNeeds\PROMPT.md`
- `<DAY_ROOT>\HumanNeeds\LOCAL-CONTEXT.json`
- `<DAY_ROOT>\HumanInputEvents\INDEX.json`
- `<DAY_ROOT>\HumanInputEvents\SOURCE-PACKET.jsonl`

Do not read any other files.
Do not search the repo.
Do not browse the web.
Do not infer missing context from memory.

When `LOCAL-CONTEXT.json` is prepared for this pass:

- source repo-local constraints from the packet repo
- do not substitute docs from the current workspace repo if the packet repo is different
- prefer `<repo_ref>\\AGENTS.md`, `<repo_ref>\\REGRESSION.md`, and `<repo_ref>\\TESTING.md` when they exist

## Output Directory

Write all outputs into:

- `<DAY_ROOT>\HumanNeeds\`

Do not write outside that directory.

## Required Outputs

1. `REPRESENTATIVE-EVENTS.json`

Schema:

```json
{
  "schema_version": "needs.packet.events.v1",
  "packet_label": "string",
  "selected_events": [
    {
      "event_id": "string",
      "captured_at_local": "string",
      "thread_name": "string",
      "explicit_request": "string",
      "selection_reason": "string",
      "priority_rank": 1
    }
  ]
}
```

Rules:

- choose 8 to 12 representative events
- prioritize events that best expose recurring need patterns, suffering patterns, or local-constraint conflicts
- keep `explicit_request` concise and grounded in the packet

2. `PACKET-TRIAGE.json`

Schema:

```json
{
  "schema_version": "needs.packet.triage.v1",
  "packet_label": "string",
  "explicit_request": "string",
  "candidate_needs": [],
  "suffering_signals": [],
  "highest_leverage_context_to_fetch": [],
  "clarifying_question": {},
  "risk_notes": []
}
```

Use the Prompt Pack C triage contract, adapted to the packet as a whole.

Rules:

- treat the packet as one day-level intervention surface
- include only context fetches that would materially change the packet-level understanding
- because the input set is already substantial, prefer `clarifying_question.needed = false` unless a truly blocking ambiguity remains

3. `PACKET-RECORD.json`

Schema:

```json
{
  "schema_version": "needs.packet.record.v1",
  "packet_label": "string",
  "explicit_request": "string",
  "selected_need": {},
  "need_distribution": [],
  "suffering_assessment": {},
  "evidence": [],
  "local_constraints_applied": [],
  "assumptions": [],
  "recommended_next_step": {},
  "self_sufficiency_improvements": [],
  "values_check": {},
  "how_could_the_system_have_inferred_the_need_for_the_input": {}
}
```

Use the Prompt Pack C fill contract, adapted to the packet as a whole.

Additional rules:

- `evidence` should cite representative events from the packet
- `local_constraints_applied` should cite `AGENTS.md`, `REGRESSION.md`, and `TESTING.md` excerpts from `LOCAL-CONTEXT.json` when used
- `how_could_the_system_have_inferred_the_need_for_the_input.answer` must be concrete, mechanistic, and repo-local
- `preventability_score` remains `0..4`

## General Rules

- Output JSON only for the three required outputs.
- Keep quotes or excerpts to 25 words or fewer.
- Separate OBSERVED vs INFERRED in notes where relevant.
- Do not claim tests were run.
- Do not pretend unavailable docs existed.
- Prefer least-committal explanations under the provided constraints.

## Final Response

At the end, report:

- exact output files written
- representative event count
- whether a clarifying question remained necessary
