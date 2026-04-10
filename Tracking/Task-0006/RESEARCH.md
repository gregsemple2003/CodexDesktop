# Task 0006 Research Summary

## Research Verdict

Task research is planning-ready.

## Key Decisions

- treat an incident as explicit human disagreement over a human-facing outcome, not as generic model uncertainty
- preserve both the surface miss and the correction path that repaired it
- require a qualification gate so plain bugs do not masquerade as incidents
- treat an incident as incomplete unless it also preserves one or more ordered `why_chains` that climb from the concrete target state toward the broader human principle
- allow incident quality to mature in two steps: fast first-pass capture, then second-pass root-cause refinement
- start with task-owned markdown records and transcript or task evidence links before introducing a shared index or store
- require the first backfill set to cover usability, UI semantics, and general misunderstandings
- use the seed set to refine the contract before promoting cross-repo workflow changes
- keep the longer-horizon consumer explicit: this corpus is meant to support an advisory counterpoint agent that models the human's interests, not just a generic critic

## Why This Is The Honest Fit

- the task scope is capture and classification, not autonomous resolution
- the last five days of history already provide enough real interventions to define the contract against actual misses
- the named prompts push toward human-facing truth, dignity, and interface clarity, which only helps if the incidents preserve those failures explicitly
- the seed transcript explicitly recommends `raw feedback -> primitive -> boundary -> falsifier -> eval -> training artifact`, which means incident capture has to preserve upstream causal structure rather than just downstream symptom notes

## Minimum Incident Record Shape

The first useful record should preserve:

- `incident_kind`
- `incident_id`
- `title`
- `source_date`
- `primary_family`
- `expected_state`
- `actual_state`
- `human_intervention`
- `human_costs`
- `evidence_refs`
- `why_chains`
  - one or more linear rationale paths
  - each path ordered from concrete to abstract
  - each entry names the target state at that level
  - each entry carries a schema-constrained clustering category

See [INCIDENT-GOAL-STACK.md](/c:/Agent/CodexDashboard/Tracking/Task-0006/Research/INCIDENT-GOAL-STACK.md) for the working contract.

## Seed Coverage Set

The recommended seed set is recorded in `Research/SEED-INCIDENTS-LAST-5D.md` and intentionally spans:

- truthful status communication
- people-surface information architecture
- operator-heavy UI language
- producer consensus failure
- proxy-proof debugging drift

Each seed incident should now carry compressed `why_chains` so the sample set tests whether the incident can climb from event grounding toward the broader human principle, not just whether it names the right topic.

The seed list is candidate-only. Any item that lacks direct evidence of a grounded pre-correction state plus human correction must be replaced during backfill instead of being forced through as an incident.

## Planning Recommendation

Plan the work in three passes:

- `PASS-0000`: define the incident contract, `why_chains`, and storage home
- `PASS-0001`: backfill five incidents from April 3, 2026 through April 7, 2026 using the approved contract and full `why_chains`
- `PASS-0002`: revisit accepted incidents for root-cause refinement so they preserve mechanism, not only surface objection
- `PASS-0003`: wire incident capture into the relevant orchestration workflow checkpoints so future human interventions stop disappearing into chat history
