# Task 0006 Human Course Correction Incident Contract

## Purpose

For `Task-0006`, an incident is not "anything that went wrong." It is a `human course correction incident`.

The record should preserve:

- the concrete event that triggered disagreement
- the human intervention itself
- the `why_chains` that progressively generalize what target state the human was actually protecting

Incident writing is allowed to happen in two steps:

- `first-pass capture`
  - preserve the event, the intervention, the evidence, and the grounded human-facing target state quickly
- `second-pass root-cause refinement`
  - revisit the accepted incident and tighten it against the actual causal chain so the record does not stop at symptom language when stronger diagnosis is available

The key simplification is this:

- `expected_state` and `actual_state` describe the concrete event
- `why_chains` is where the abstraction starts

## Qualification Gate

A record qualifies only if it can answer these questions from durable evidence:

1. What concrete state or active course existed before the human stepped in?
2. What did the human explicitly reject, tighten, or redirect?
3. What concrete event should have happened instead?

If those are not preserved, the record is probably a bug note, review note, or task input instead of a `Task-0006` incident.

Screenshot-visible distortion can still qualify as a real incident boundary when the human is correcting what the surface reads like or what quality it communicates, for example through typography, spacing, icon fidelity, hierarchy, clipping, or component-semantic mismatch against the intended mockup or surface contract.

## Required Core Fields

Every incident should preserve:

- `incident_kind`
- `incident_id`
- `title`
- `source_date`
- `primary_family`
- `expected_state`
- `actual_state`
- `human_intervention`
- `human_costs`
- `feedback_quotes`
- `evidence_refs`
- `verbatim_transcript_windows`
- `why_chains`

Accepted official incidents should be heavyweight:

- keep the concrete event and the distilled `why_chains`
- but also keep the verbatim transcript context inside the incident JSON itself

The later actionable principles distilled from incidents should be lightweight.
Do not make the incident lightweight by stripping out the raw evidence that gives it warrant.

Incidents may also preserve an optional second-pass `analysis` block when a human-readable causal read would help:

- `analysis.deeper_miss`
  - short grounded lines explaining the deeper miss or causal stack
- `analysis.correction`
  - short grounded lines explaining how the incident was actually corrected

Keep those analysis lines short enough to read comfortably in raw JSON without horizontal scrolling.

## Daily Artifact Home

Keep the contract and examples separate from the dated capture corpus.

- top-level `Tracking/Task-0006/Research/`
  - schema
  - contract notes
  - hand-built examples such as `INCIDENT-EXAMPLE-0001.json`
- task-local bootstrap capture home
  - `Tracking/Task-0006/Research/Daily/YYYY-MM-DD/`
- promoted shared intervention home
  - `C:\Users\gregs\.codex\Orchestration\Reports\Interventions\`

Each dated folder should own:

- any transcript-first intervention analysis artifacts for that day
  - `INTERVENTION-PASS1.md`
  - `INTERVENTION-PASS2.md`
  - `INTERVENTION-PASS3.md`
  - these are task-local recall, investigation, and principle-extraction artifacts rather than accepted-corpus records
- any narrower accepted-evidence packaging the task still wants to keep locally
  - only if it is honestly derived from the pass stack
  - accepted incident JSONs should carry the heavyweight verbatim evidence for that incident, not only references out to the transcript
- optional review packaging when a second-opinion packet is helpful
  - `DAILY-BRIEF.md`

Older CSV review captures from the bootstrap phase are deprecated and should not be recreated.

This split keeps examples and schema work stable while letting task-local backfill stay close to the original transcript day and letting accepted official runs promote into the durable reports corpus.

## Event Grounding Rule

`expected_state` and `actual_state` must stay grounded in the concrete incident.

- `expected_state`
  - what should have happened in this specific task, artifact, pass, or message
- `actual_state`
  - what actually happened in that same specific task, artifact, pass, or message
  - if the incident is about ongoing drift, include that active course here as part of the same grounded description

Do not use those fields to summarize the broader lesson. They are for the event itself.

If first-pass capture only has the surface event phrasing, that is acceptable as long as it is honest.

But once later evidence shows the deeper mechanism clearly, use the second-pass refinement to make `actual_state` more causally truthful. For example:

- distinguish a product-model bug from a renderer bug
- distinguish a workflow miss from a prompt miss
- distinguish a real implementation defect from a merely ambiguous plan

## Required Why Chains

The incident must also preserve one or more ordered `why_chains`.

Treat each chain as a reflexive `why` structure:

- entry `N+1` should answer: "why did entry `N` matter to the human?"
- each later entry should be more generalized than the one before it
- the chain should stop at the highest-level reason the human explicitly stated in the evidence for that path
- do not infer a broader product, economic, or philosophical principle yet

Operationally, that means:

- do not force extra layers if the evidence only supports one broader `why`
- do not split the same `why` into two adjacent synonym frames
- do not switch from `why this mattered` into `what process should exist`
- do not turn a later distilled rule into a chain entry unless that rule is itself the human's stated reason for the previous entry
- treat `principle_category` as a constrained clustering label, not as proof of a hierarchy
- `principle_category` values should come from the schema enum exactly; if no category fits yet, use `other`
- if the human expressed multiple independent upstream whys, use multiple chains instead of forcing them into one sequence

Each chain should usually start near the concrete target state and end near the highest-level human-stated reason being expressed for that path.

Each chain must include:

- `chain_label`
  - short label that distinguishes this path from sibling paths in the same incident
- `entries`
  - ordered linear entries for that path

Each entry must include:

- `principle_category`
  - constrained category for the kind of principle being expressed
  - intended mainly for later cross-incident clustering
- `target_state`
  - the target state at that level of generalization
- `evidence_refs`

## What The Why Chains Are For

The `why_chains` are not a blame tree and not yet a reusable primitive library rule.

They are a progressive explanation of where the human is coming from.

Good chains usually move through shapes like:

- the immediate task or artifact target
- the next broader target state that made that event matter
- the next broader human-stated reason behind that target state
- the highest broader reason explicitly stated by the human in the evidence

Use multiple chains when the human expressed multiple sibling rationales that do not recursively answer each other.

Bad chains usually fail in one of two ways:

- they switch from a reason chain into a process prescription
- they jump past the evidence into an inferred broader motive like product cost, time, or philosophy

Another common failure mode:

- they overfit the shape and invent three or four layers when the evidence really only supports one or two broader reasons

## Allowed Categories

The current schema intentionally constrains `principle_category` to:

- `product_resolution`
- `reviewability`
- `evaluability`
- `falsifiability`
- `truthfulness`
- `boundedness`
- `diagnostic_scope`
- `proxy_discipline`
- `directness`
- `consensus_building`
- `human_dignity`
- `operator_clarity`
- `human_control`
- `other`

This list is for clustering discipline, not to force a false theory of the incident.

## Non-Example

A product defect by itself is not a qualified incident.

Example:

- "The Home card says `Safe on server` while still showing unfinished progress" is a valid bug statement.
- It only becomes a `Task-0006` incident if the artifact also preserves an AI-produced framing, proposal, or continued course that the human had to correct.

## Minimal Skeleton

```json
{
  "incident_kind": "human_course_correction",
  "incident_id": "INC-0001",
  "title": "",
  "source_date": "",
  "primary_family": "",
  "expected_state": "",
  "actual_state": "",
  "human_intervention": {
    "kind": "other",
    "summary": "",
    "evidence_refs": []
  },
  "human_costs": [],
  "feedback_quotes": [],
  "evidence_refs": [],
  "verbatim_transcript_windows": [
    {
      "window_label": "",
      "source_ref": "",
      "dialogue": []
    }
  ],
  "analysis": {
    "deeper_miss": [],
    "correction": []
  },
  "why_chains": [
    {
      "chain_label": "",
      "entries": [
        {
          "principle_category": "",
          "target_state": "",
          "evidence_refs": []
        },
        {
          "principle_category": "",
          "target_state": "",
          "evidence_refs": []
        }
      ]
    }
  ]
}
```

## Working Rule

Ground the record in one concrete event first.

Then use `why_chains` to progressively generalize from:

- what should have happened here
- to what target state the human was really protecting
- to the highest-level reason the human explicitly stated in the evidence

Stop before inferring a broader unstated motive or turning it into a reusable primitive rule. That comes later.

Store the real capture artifacts by day so the human can review the raw message boundary and the incident boundary together.
