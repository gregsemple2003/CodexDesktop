# Task 0006 Plan

## Planning Verdict

This plan is ready for explicit human approval.

Human gate:

- do not start `PASS-0000` until the human approves this `PLAN.md`

This revision addresses the rejected first plan: incidents must preserve grounded `why_chains`, not just a surface-level interpretation of what looked wrong.

This revision also simplifies the contract: the incident should preserve the grounded pre-correction state, the human intervention itself, the concrete event, and the progressively generalized target states the human was actually protecting.

The longer-horizon consumer is not just a generic critic. The likely downstream shape is a counterpoint agent that estimates `I think the human will want X` and advises the producer, usually without veto authority. That makes over-capture of daily human-interest context valuable even when a given row is not itself an incident.

The current incident set suggests at least four recurring help modes that later workflow should protect:

- state-story truth
- real-world done
- human-facing form factor
- control-boundary ownership

## Planning Basis

The plan is grounded in:

- `Task-0006` scope narrowing under [TASK.md](/c:/Agent/CodexDashboard/Tracking/Task-0006/TASK.md)
- the seed transcript in [Scaling-Agent-Orchestration.md](/c:/Agent/CodexDashboard/Tracking/Task-0006/Research/Scaling-Agent-Orchestration.md)
- the incident contract in [INCIDENT-GOAL-STACK.md](/c:/Agent/CodexDashboard/Tracking/Task-0006/Research/INCIDENT-GOAL-STACK.md)
- prompt review pressure from [VISION-HARVESTER.md](/c:/Users/gregs/.codex/Orchestration/Prompts/VISION-HARVESTER.md), [GENERAL-DESIGNER.md](/c:/Users/gregs/.codex/Orchestration/Prompts/GENERAL-DESIGNER.md), and [INTERFACE-DESIGNER.md](/c:/Users/gregs/.codex/Orchestration/Prompts/INTERFACE-DESIGNER.md)
- transcript-backed research recorded in [RESEARCH.md](/c:/Agent/CodexDashboard/Tracking/Task-0006/RESEARCH.md)

## Required Seed Incident Coverage

`PASS-0001` must backfill five incidents from April 3, 2026 through April 7, 2026.

Those five incidents are not arbitrary. They should deliberately cover usability, UI semantics, and general misunderstandings. More importantly, each one must preserve real `why_chains` that climb from the concrete event toward the highest-level reason the human explicitly stated in evidence.

Each candidate must also clear the qualification gate:

- preserve the grounded pre-correction state, including any active course when relevant
- preserve the `human_intervention`

And the base incident fields must stay event-specific:

- `expected_state` describes what should have happened in that concrete incident
- `actual_state` describes what actually happened in that same concrete incident
- broader abstractions belong in `why_chains`, not in those two fields

The recommended starting set is documented in [SEED-INCIDENTS-LAST-5D.md](/c:/Agent/CodexDashboard/Tracking/Task-0006/Research/SEED-INCIDENTS-LAST-5D.md):

1. Home upload state story mismatch.
   Expected content: not just the contradictory surface states, but clean `why_chains` that explain the target state the human was protecting.
2. People surface behaves like a setup detour.
   Expected content: not just the naming and CTA mismatch, but `why_chains` from the concrete route correction toward the broader human principle behind it.
3. Jobs tab row taxonomy is too operator-heavy.
   Expected content: not just the jargon labels, but `why_chains` that progressively explain the target state the human expected from the surface.
4. Producer left the resolution too loose.
   Expected content: not just the loose task wording, but `why_chains` that show the concrete task target, the agreement-check need behind it, and the highest usability-task rule the human actually stated.
5. Proof-view experiment drifted away from the real defect.
   Expected content: not just the drifting proof work, but `why_chains` that start at the bounded check and climb toward the broader proxy-discipline principle.

The point of the five-sample set is breadth. Do not substitute five near-duplicates from one UI surface.

The seed note is now explicitly candidate-only. Any item that cannot clear the qualification gate must be replaced rather than forced through as an incident.

For incident writing itself, use a two-step maturity model:

- `first-pass capture`
  - get the incident artifact in place with honest event grounding and `why_chains`
- `second-pass root-cause refinement`
  - revisit accepted incidents and tighten them against the real mechanism when later diagnosis makes that possible

Do not hold first-pass capture hostage to perfect causal diagnosis. But do not let accepted incidents stay permanently surface-level when the deeper cause becomes clear.

## PASS-0000 Incident Contract, Why Chains, And Storage Home

### Objective

Define the first durable incident contract, the required `why_chains` shape, and where incident artifacts live before backfilling examples.

### Implementation Notes

- land [INCIDENT.schema.json](/c:/Agent/CodexDashboard/Tracking/Task-0006/Research/INCIDENT.schema.json) under `Tracking/Task-0006/Research/`
- land a dated storage note under `Tracking/Task-0006/Research/Daily/README.md`
- promote the stable shared intervention workflow into `C:\Users\gregs\.codex\Orchestration\Reports\Interventions\` once the shape is stable enough for a real run
- land exactly two validated examples:
  - [INCIDENT-EXAMPLE-0001.json](/c:/Agent/CodexDashboard/Tracking/Task-0006/Research/INCIDENT-EXAMPLE-0001.json)
  - [INCIDENT-EXAMPLE-0002.json](/c:/Agent/CodexDashboard/Tracking/Task-0006/Research/INCIDENT-EXAMPLE-0002.json)
- define the minimum incident fields needed to preserve both the surface miss and the orchestration context
- define the heavyweight-evidence fields required by accepted incidents:
  - `verbatim_transcript_windows`
- require the qualification triplet:
  - grounded `actual_state`
  - `human_intervention`
- define the required `why_chains` shape and its meaning
- define how each chain records:
  - `chain_label`
  - `entries`
- define how each chain entry records:
  - `principle_category`
  - `target_state`
  - `evidence_refs`
- decide the first storage home for incident records and evidence links
- make the daily capture split explicit:
  - one day folder per incident date under `Research/Daily/YYYY-MM-DD/`
  - same-day incident JSONs stored beside that day's pass artifacts
  - transcript-first `INTERVENTION-PASS1.md`, `INTERVENTION-PASS2.md`, and `INTERVENTION-PASS3.md` artifacts as the task-local recall and analysis layer
- define the promotion boundary from task-local bootstrap storage to `C:\Users\gregs\.codex\Orchestration\Reports\Interventions\`
- define explicitly that the task-local intervention passes preserve two things at once:
  - likely incident candidates for later acceptance work
  - broader human-interest and human-model evidence that may later feed the counterpoint-agent corpus
- define the derived `DAILY-BRIEF.md` artifact as a self-contained second-opinion packet assembled from accepted incidents and transcript context
- require `DAILY-BRIEF.md` to favor first-principles review over compression by carrying larger contiguous verbatim context, a chronological verbatim human-outbound layer when needed for memory recovery, and a strict separation between raw evidence and interpretation
- keep the first implementation task-local unless a shared `.codex` contract is already stable enough to promote
- make the shared-versus-task-local split explicit so future orchestration work knows what should move into `.codex`

### Verification

- `Test-Json -Path Tracking/Task-0006/Research/INCIDENT-EXAMPLE-0001.json -SchemaFile Tracking/Task-0006/Research/INCIDENT.schema.json` returns `True`
- `Test-Json -Path Tracking/Task-0006/Research/INCIDENT-EXAMPLE-0002.json -SchemaFile Tracking/Task-0006/Research/INCIDENT.schema.json` returns `True`
- the schema requires `incident_kind = human_course_correction`
- the schema requires grounded `actual_state` and `human_intervention`
- the official accepted-corpus schema requires `verbatim_transcript_windows`
- the schema requires one or more ordered `why_chains`
- each chain requires `chain_label` and ordered `entries`
- each chain entry requires `principle_category`, `target_state`, and `evidence_refs`
- `principle_category` is constrained to the schema enum for clustering rather than free-form labeling
- verify that the chosen storage home preserves evidence links to tasks, screenshots, and session transcripts
- verify that the daily storage note defines one day folder per incident date and the intended task-local pass artifacts
- verify that the daily storage note defines the purpose and non-source-of-truth status of `DAILY-BRIEF.md`
- verify that the daily storage note defines the no-cherry-picking rule, the verbatim-only evidence requirement, and the soft `200 KB` evidence budget for `DAILY-BRIEF.md`
- verify that the daily storage note explicitly says deprecated CSV review captures should not be recreated
- record any still-open design choices as bounded `Open Questions`, not as vague plan drift

### Exit Bar

- [INCIDENT.schema.json](/c:/Agent/CodexDashboard/Tracking/Task-0006/Research/INCIDENT.schema.json) exists
- [INCIDENT-EXAMPLE-0001.json](/c:/Agent/CodexDashboard/Tracking/Task-0006/Research/INCIDENT-EXAMPLE-0001.json) exists and validates against the schema
- [INCIDENT-EXAMPLE-0002.json](/c:/Agent/CodexDashboard/Tracking/Task-0006/Research/INCIDENT-EXAMPLE-0002.json) exists and validates against the schema
- the first artifact home for incidents is explicit
- the daily storage contract is explicit
- the role of `DAILY-BRIEF.md` as a derived review packet is explicit
- the accepted incident JSON is explicitly heavyweight rather than relying on the brief to reconstruct the evidence
- the promotion path into `C:\Users\gregs\.codex\Orchestration\Reports\Interventions\` is explicit
- the contract is narrow enough to exclude ordinary bugs and broad enough to preserve real human interventions
- the contract cannot validate a plain bug note that lacks the grounded pre-correction state and direct human intervention evidence
- the contract cleanly separates the concrete event from the progressively generalized `why_chains`

## PASS-0001 Backfill Five Recent Incidents

### Objective

Capture five incidents from April 3, 2026 through April 7, 2026 using the approved contract and full `why_chains`.

### Implementation Notes

- backfill five qualified incidents from the candidate set documented in `Research/SEED-INCIDENTS-LAST-5D.md`
- create one `INTERVENTION-PASS1.md`, `INTERVENTION-PASS2.md`, and `INTERVENTION-PASS3.md` per day for April 3, 2026 through April 7, 2026 under `Research/Daily/YYYY-MM-DD/`
- each daily pass stack must stay transcript-first and preserve candidate intervention evidence even when an event is not promoted into the accepted incident set
- record each incident in the folder matching its `source_date`, for example `Research/Daily/2026-04-03/INCIDENT-0001.json`
- when a shared rule, canon update, or eval design is ready for promotion, land it under `C:\Users\gregs\.codex\Orchestration\Reports\Interventions\`
- when outside review or second opinion is needed, generate one `DAILY-BRIEF.md` in that same day folder
- structure the brief so a second reviewer can reason from raw relevant evidence before reading the harvester's conclusions
- keep the set intentionally mixed across usability, UI semantics, and general misunderstandings
- replace any candidate that does not clear the qualification gate instead of weakening the schema
- preserve the expected state, actual outcome, and human intervention for each record
- preserve the full upward `why_chains` for each incident
- allow first-pass capture to land before full causal diagnosis is complete
- mark accepted incidents as eligible for second-pass root-cause refinement
- make each later entry in each chain answer why the prior entry mattered to the human
- stop each chain at the highest-level human-stated reason in evidence rather than inferring a broader motive
- use multiple sibling chains when the human expressed multiple independent upstream whys
- prefer task or transcript evidence that already contains the human disagreement in durable form

### Verification

- `Test-Json` returns `True` for each incident JSON created under the daily folders against `INCIDENT.schema.json`
- each day in the April 3-7 window has a full `INTERVENTION-PASS1.md` -> `INTERVENTION-PASS2.md` -> `INTERVENTION-PASS3.md` stack
- when a daily brief exists, it inlines the accepted incidents plus enough contiguous raw context for first-principles review without reconstructing the corpus manually
- each incident has at least one concrete transcript or task-artifact evidence reference
- the five-record set includes at least:
  - one `usability_state_truth`
  - one `information_architecture` or `ui_semantics`
  - one `workflow_orchestration` or `verification_proof`
- the five-record set is visibly non-redundant in coverage
- the task-local pass artifacts still retain important non-incident intervention or human-model evidence instead of narrowing the corpus to accepted incidents only

### Exit Bar

- the April 3-7 daily folders exist under `Research/Daily/`
- each day folder has the intervention-pass artifact stack
- five incident JSONs exist across those dated folders and validate against the schema
- at least one accepted official day can produce a self-contained `DAILY-BRIEF.md` for second-opinion review
- at least one stable shared intervention artifact can be materialized under `C:\Users\gregs\.codex\Orchestration\Reports\Interventions\`
- the records reflect good coverage across usability, UI, and general misunderstandings
- the records also reflect good coverage across different kinds of human principles rather than stopping at symptom summaries
- the seed set feels like a real starter corpus for later prompt, eval, or workflow work

## PASS-0002 Root-Cause Refinement On Accepted Incidents

### Objective

Revisit accepted incidents after capture and tighten them against the actual causal chain so the corpus preserves more than the surface objection.

### Implementation Notes

- review accepted incidents that now have stronger downstream evidence from task artifacts, code, screenshots, tests, or follow-on investigation
- refine `actual_state` when the first-pass version was honest but too shallow
- distinguish symptom from mechanism where the evidence supports it, such as:
  - product-model bug
  - renderer bug
  - workflow miss
  - prompt miss
  - operator-lane mismatch
- keep the incident anchored to the same human correction event rather than turning it into a separate bug record
- preserve the original `why_chains` unless the new evidence shows the incident was grounded against the wrong target state
- do not invent deeper causes that are not supported by durable evidence

### Verification

- refined incidents still validate against `INCIDENT.schema.json`
- refined incidents read more causally truthful than their first-pass versions
- `actual_state` reflects the stronger diagnosis when one exists
- the record still reads as a human course correction incident rather than a generic bug report

### Exit Bar

- at least one accepted incident has been tightened through second-pass root-cause refinement
- the corpus documents that first-pass capture and second-pass refinement are both expected parts of incident quality
- the refined incident is more useful for future orchestration learning than the purely surface-level version

## PASS-0003 Workflow Hook Points, Promotion Boundaries, And Escalation Capture

### Objective

Define how future incidents get captured at the moment of human intervention instead of only through later archaeology, and decide when recorded `why_chains` should later be promoted into reusable primitives or shared rules.

### Implementation Notes

- identify which orchestration roles should open or update incidents when the human steps in
- identify which role or helper is responsible for opening or refreshing the day-scoped intervention-pass artifacts
- identify which roles are responsible for filling in the concrete event fields and the `why_chains` honestly
- identify how later workflow should distinguish:
  - advisory human-interest signals for the counterpoint agent
  - stricter incident triggers that justify stronger intervention
- define the default authority split explicitly:
  - the counterpoint agent usually advises
  - the producer usually decides
  - stronger intervention or escalation only happens when the later workflow says the conflict is high-confidence or high-stakes
- define how incident records link to `TASK.md`, `PLAN.md`, `HANDOFF.md`, bug notes, or pass artifacts without replacing them
- define when an individual `why` path is still incident-local and should not yet be promoted
- define when a repeated `why` path pattern is strong enough to justify shared prompt or workflow changes
- promote only the durable shared rules into `.codex/Orchestration/`
- keep repo-specific or task-specific history under `Tracking/Task-0006/`

### Verification

- the workflow says who captures the incident, when it should happen, and what minimum evidence must be attached
- the workflow says who is allowed to promote a repeated `why` path pattern into a durable rule
- the workflow says clearly that the future counterpoint role is advisory by default rather than a generic veto layer
- the workflow does not confuse incidents with generic bugs, audits, or regression failures
- the shared-doc changes are still true north rules rather than task history

### Exit Bar

- future human disagreement can be turned into a durable incident record without guesswork
- the shared-versus-local split is documented honestly
- the promotion boundary between local incident, `why`-path distillation, and shared rule is documented honestly
- the task can move from research capture into durable orchestration changes without another scope reset

## Task-Level Validation

This is an orchestration and artifact task, not an app-feature task.

Expected validation for closure:

- artifact review of the contract, sample records, and workflow updates
- artifact review of whether each incident's `why_chains` are grounded, progressive, and scoped honestly
- schema or template sanity checks if a machine-readable companion lands
- no build, unit-test, or app regression claims unless the task later grows real implementation helpers that need them

## Watchouts

- do not widen the task into reviewer swarms, automated critics, or training pipelines
- do not capture only final fixes while losing the disagreement that made the incident valuable
- do not let the seed set collapse into five UI-copy incidents
- do not let any `why` path jump to inferred product philosophy, economics, or other unstated motives before the concrete event is clear
- do not promote task history into shared `.codex` docs
