# INTERVENTION-PASS2 Analysis Report

Source date: `2026-04-08`

Pass scope: analysis-only. No accepted incident JSON files were written.

## 1. Source scope analyzed

- PASS1 source: [INTERVENTION-PASS1.md](/c:/Agent/CodexDashboard/Tracking/Task-0006/Research/Daily/2026-04-08/INTERVENTION-PASS1.md)
- Incident corpus contract read before classification: [README.md](/c:/Users/gregs/.codex/Orchestration/Reports/Incidents/README.md), [INCIDENT.schema.json](/c:/Users/gregs/.codex/Orchestration/Reports/Incidents/INCIDENT.schema.json), [2026-04-08 README.md](/c:/Users/gregs/.codex/Orchestration/Reports/Incidents/Daily/2026-04-08/README.md)
- Raw transcripts reopened from PASS1 refs:
  - [rollout-2026-04-08T12-24-37-019d6de8-f7df-7901-bca9-6ab0d656d5c1.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T12-24-37-019d6de8-f7df-7901-bca9-6ab0d656d5c1.jsonl)
  - [rollout-2026-04-08T13-03-23-019d6e0c-7467-7042-92f5-aa976e10a3fe.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T13-03-23-019d6e0c-7467-7042-92f5-aa976e10a3fe.jsonl)
  - [rollout-2026-04-08T15-59-13-019d6ead-6ee7-7201-a865-9dd3f168dd1f.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T15-59-13-019d6ead-6ee7-7201-a865-9dd3f168dd1f.jsonl)
  - [rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl)
  - [rollout-2026-04-08T21-38-33-019d6fe4-199d-7f50-bf08-8d68086d978e.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T21-38-33-019d6fe4-199d-7f50-bf08-8d68086d978e.jsonl)
- Candidate ids analyzed: `C01`, `C02`, `C03`, `C04`, `C05`, `C06`, `C07`, `C08`, `C09`, `C10`, `C11`
- Chronology note: inside [rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl), imported history preserves many same-second timestamps. For that file, line order was treated as the trustworthy chronology.

## 2. Candidate boundary corrections relative to PASS1

- `C01`: keep as one weak event, but the correction arc has two stages, not one. The human first says to pass on the `ReadPrev` side issue at [L684](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T12-24-37-019d6de8-f7df-7901-bca9-6ab0d656d5c1.jsonl#L684), then later narrows the work again to a minimal production-grade crash fix at [L716](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T12-24-37-019d6de8-f7df-7901-bca9-6ab0d656d5c1.jsonl#L716).
- `C05`: extend the boundary beyond the first "this reads like a bug report" challenge. The same correction arc continues into event-level grounding for `expected_state` and `actual_state` at [L1143](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L1143).
- `C07`: extend the boundary through the explicit over-capture rule at [L2252](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L2252). The day-folder and CSV change alone misses the human's adequacy rule.
- `C10`: keep as one event, but treat it as a serial representation reset from [L82](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T21-38-33-019d6fe4-199d-7f50-bf08-8d68086d978e.jsonl#L82) through [L310](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T21-38-33-019d6fe4-199d-7f50-bf08-8d68086d978e.jsonl#L310), not a one-off complaint about a single section label.
- `C11`: no split or merge correction, but PASS2 clarifies the decisive provenance fact: the suspicious child-thread `role:"user"` text was a parent `send_input` relay, not the human's own wording, as established at [L563](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T21-38-33-019d6fe4-199d-7f50-bf08-8d68086d978e.jsonl#L563).

## 3. Per-event analysis records

### C01. Scope snapped back from `ReadPrev` analysis to the presentation-buffer crash fix

- `event_id`: `C01`
- `title`: Scope snapped back from `ReadPrev` analysis to the presentation-buffer crash fix
- `session_or_thread`: `019d6de8-f7df-7901-bca9-6ab0d656d5c1`
- `transcript_path`: [rollout-2026-04-08T12-24-37-019d6de8-f7df-7901-bca9-6ab0d656d5c1.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T12-24-37-019d6de8-f7df-7901-bca9-6ab0d656d5c1.jsonl)
- `primary_refs`: [L584](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T12-24-37-019d6de8-f7df-7901-bca9-6ab0d656d5c1.jsonl#L584), [L684](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T12-24-37-019d6de8-f7df-7901-bca9-6ab0d656d5c1.jsonl#L684), [L709](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T12-24-37-019d6de8-f7df-7901-bca9-6ab0d656d5c1.jsonl#L709), [L716](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T12-24-37-019d6de8-f7df-7901-bca9-6ab0d656d5c1.jsonl#L716)
- `ai_course`: After explaining the downstream `ReadPrevPresentationSyncState()` issue, the assistant kept treating that side seam and broader engine-hardening questions as the active lane, even after the resize-remap crash fix was the confirmed target.
- `human_intervention`: The human said to pass on the `ReadPrev` issue for now, stay focused on the presentation sync buffer resize, and later asked for only a minimal production-grade patch for the crash fix.
- `adequate_outcome`: Close the confirmed crash seam first and separate the `ReadPrev` issue and other hardening work into later follow-up.
- `event_boundary_notes`: One event with two scope-tightening pushes: the redirect at [L684](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T12-24-37-019d6de8-f7df-7901-bca9-6ab0d656d5c1.jsonl#L684) and the minimal-patch narrowing at [L716](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T12-24-37-019d6de8-f7df-7901-bca9-6ab0d656d5c1.jsonl#L716).
- `human_model_signal`: "Stay focused on the presentation sync buffer resize." / "Let's do only the presentation buffer crash fix."
- `failure_family_hypothesis`: `workflow_orchestration`; local flavor `wrong-seam debugging`
- `intervention_kind_hypothesis`: `redirect_debugging`
- `human_cost_or_risk`: Extra debug time and a broader patch than the user wanted for review.
- `local_lesson_hypothesis`: Once the requested seam is confirmed, adjacent findings should become separate follow-up unless the human explicitly broadens scope.
- `cluster_hints`: `wrong-seam debugging`, `minimal-patch discipline`
- `accepted_incident_likelihood`: `unlikely`
- `confidence`: `weak`
- `uncertainties`: This is close to ordinary collaborative narrowing. The redirect is real, but the accepted-set durability is marginal.

### C02. Digest success had to be redefined around real end-state proof

- `event_id`: `C02`
- `title`: Digest success had to be redefined around real end-state proof
- `session_or_thread`: `019d6e0c-7467-7042-92f5-aa976e10a3fe`
- `transcript_path`: [rollout-2026-04-08T13-03-23-019d6e0c-7467-7042-92f5-aa976e10a3fe.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T13-03-23-019d6e0c-7467-7042-92f5-aa976e10a3fe.jsonl)
- `primary_refs`: [L141](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T13-03-23-019d6e0c-7467-7042-92f5-aa976e10a3fe.jsonl#L141), [L171](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T13-03-23-019d6e0c-7467-7042-92f5-aa976e10a3fe.jsonl#L171), [L178](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T13-03-23-019d6e0c-7467-7042-92f5-aa976e10a3fe.jsonl#L178), [L190](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T13-03-23-019d6e0c-7467-7042-92f5-aa976e10a3fe.jsonl#L190)
- `ai_course`: The assistant verified that the manual digest run existed, noted the `rg` failure, and began reasoning about wrappers and execution paths, but it was still treating runner behavior as the primary success question.
- `human_intervention`: The human reset the success criterion to: failure email if the real end state was missing, specifically no email sent and no report saved.
- `adequate_outcome`: The orchestration layer should verify observable end-state artifacts and emit a failure alert when those artifacts are missing, not rely on raw runner behavior alone.
- `event_boundary_notes`: The intervention is the contract reset at [L178](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T13-03-23-019d6e0c-7467-7042-92f5-aa976e10a3fe.jsonl#L178), with the repair path starting immediately at [L190](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T13-03-23-019d6e0c-7467-7042-92f5-aa976e10a3fe.jsonl#L190).
- `human_model_signal`: "I want a failure email if it didn't reach the end state - email sent, or report saved."
- `failure_family_hypothesis`: `workflow_orchestration`; local flavor `closure-truth`
- `intervention_kind_hypothesis`: `tighten_contract`
- `human_cost_or_risk`: Silent or misleading success would force manual checking and weaken trust in scheduled digests.
- `local_lesson_hypothesis`: Human-facing jobs should define success by observable end-state artifacts, not by whether an intermediate runner looked active.
- `cluster_hints`: `closure-truth`, `end-state verification`, `human-facing completion proof`
- `accepted_incident_likelihood`: `possible`
- `confidence`: `medium`
- `uncertainties`: The transcript shows a contract correction more clearly than a concrete missed end state that day, so the accepted-set case is weaker than the lesson itself.

### C03. Same-day build intent was violated by a stale server fallback

- `event_id`: `C03`
- `title`: Same-day build intent was violated by a stale server fallback
- `session_or_thread`: `019d6ead-6ee7-7201-a865-9dd3f168dd1f`
- `transcript_path`: [rollout-2026-04-08T15-59-13-019d6ead-6ee7-7201-a865-9dd3f168dd1f.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T15-59-13-019d6ead-6ee7-7201-a865-9dd3f168dd1f.jsonl)
- `primary_refs`: [L477](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T15-59-13-019d6ead-6ee7-7201-a865-9dd3f168dd1f.jsonl#L477), [L497](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T15-59-13-019d6ead-6ee7-7201-a865-9dd3f168dd1f.jsonl#L497), [L545](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T15-59-13-019d6ead-6ee7-7201-a865-9dd3f168dd1f.jsonl#L545), [L558](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T15-59-13-019d6ead-6ee7-7201-a865-9dd3f168dd1f.jsonl#L558)
- `ai_course`: The assistant launched a packaged server, declared it up, and gave the user a client connect command. Only after the connection mismatch showed up did it surface that the running server binary was still from April 2, 2026.
- `human_intervention`: The human explicitly called the fallback a dropped ball and restated the contract: if the server build was requested on April 8, the launched server had to come from April 8 output.
- `adequate_outcome`: Treat the failed same-day server build as a hard blocker and say that no matching server exists yet, rather than silently falling back to an older package.
- `event_boundary_notes`: The event starts with the "server is up" claim at [L477](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T15-59-13-019d6ead-6ee7-7201-a865-9dd3f168dd1f.jsonl#L477) and closes once the assistant admits the same-day contract at [L558](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T15-59-13-019d6ead-6ee7-7201-a865-9dd3f168dd1f.jsonl#L558).
- `human_model_signal`: "If I asked you to build the server right? ... the expectation would be that you launch the same server from today right?"
- `failure_family_hypothesis`: `workflow_orchestration`; local flavor `real-world completion mismatch`
- `intervention_kind_hypothesis`: `reject_outcome`
- `human_cost_or_risk`: Wasted test cycle, misleading launch guidance, and trust damage around artifact lineage.
- `local_lesson_hypothesis`: When a request implies artifact lineage and same-day parity, an older fallback is a contract violation unless the human explicitly approves it.
- `cluster_hints`: `closure-truth`, `artifact lineage`, `stale fallback`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: None material. The transcript makes the stale-fallback miss and the human's adequacy rule explicit.

### C04. The Task-0006 plan had to be rebuilt around falsifiable upstream reasoning

- `event_id`: `C04`
- `title`: The Task-0006 plan had to be rebuilt around falsifiable upstream reasoning
- `session_or_thread`: `Task-0006 contract thread (session id 019d6942-63b6-7c72-9abc-9ab209816f97)`
- `transcript_path`: [rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl)
- `primary_refs`: [L729](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L729), [L747](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L747), [L854](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L854), [L869](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L869)
- `ai_course`: The assistant first marked the plan ready for review using high-level incident themes, then shifted into `reason_stack` and `goal_stack` terminology without giving the task a falsifiable contract.
- `human_intervention`: The human rejected the plan as surface-level and hand-wavy, demanded upstream reason tracing that stayed falsifiable, and then forced the work into a concrete schema plus worked examples.
- `adequate_outcome`: An incident contract that can be inspected and disproved, with machine-checkable structure and examples rather than only conceptual prose.
- `event_boundary_notes`: One event. The later "nail down the incident schema and give me a couple of examples" request at [L854](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L854) is the same miss expressed more concretely.
- `human_model_signal`: "Trace the incident upstream, keeping the interpretation falsifiable" and "Your plan isn't falsifiable, too much hand-waviness."
- `failure_family_hypothesis`: `workflow_orchestration`; local flavor `falsifiable contract`
- `intervention_kind_hypothesis`: `tighten_contract`
- `human_cost_or_risk`: Review would be impossible because the task would preserve themes, not inspectable incident reasoning.
- `local_lesson_hypothesis`: A task about durable incident evidence needs a falsifiable artifact contract before it needs more examples or more terminology.
- `cluster_hints`: `falsifiable contract`, `machine-checkable artifact`, `upstream-reason tracing`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: None material.

### C05. Incident examples had to stop reading like bug reports and regain a concrete correction boundary

- `event_id`: `C05`
- `title`: Incident examples had to stop reading like bug reports and regain a concrete correction boundary
- `session_or_thread`: `Task-0006 contract thread (session id 019d6942-63b6-7c72-9abc-9ab209816f97)`
- `transcript_path`: [rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl)
- `primary_refs`: [L925](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L925), [L950](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L950), [L1114](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L1114), [L1143](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L1143)
- `ai_course`: The assistant produced example incident records that overclassified product bugs as incidents, then kept the record shape abstract enough that a reader could not see the concrete correction event from `expected_state` and `actual_state` alone.
- `human_intervention`: The human asked why the example qualified at all, reframed the corpus as "human course correction incident" rather than generic incident, and then restricted `expected_state` and `actual_state` to the concrete event rather than the abstraction layer.
- `adequate_outcome`: Only keep examples where the AI's prior course, the human's correction, and the event-level expected/actual state are recoverable from the record itself.
- `event_boundary_notes`: PASS1 caught the overclassification challenge; the same event continues into the event-grounding rule at [L1143](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L1143).
- `human_model_signal`: "These are really 'human course correction incident', not just 'any incident' in the abstract." / "`expected_state` ... `actual_state` ... should reference the concrete event."
- `failure_family_hypothesis`: `workflow_orchestration`; local flavor `incident-boundary overreach`
- `intervention_kind_hypothesis`: `reframe_problem`
- `human_cost_or_risk`: The corpus would preserve bug reports and abstractions instead of actual correction events, making later synthesis unreliable.
- `local_lesson_hypothesis`: Before polishing schema language, prove that the record actually captures an AI correction event and ground the top-level fields in that event.
- `cluster_hints`: `qualification gate`, `event grounding`, `incident-vs-bug boundary`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: None material.

### C06. The strict `why_chain` rule did not stick until the shape changed to `why_chains`

- `event_id`: `C06`
- `title`: The strict `why_chain` rule did not stick until the shape changed to `why_chains`
- `session_or_thread`: `Task-0006 contract thread (session id 019d6942-63b6-7c72-9abc-9ab209816f97)`
- `transcript_path`: [rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl)
- `primary_refs`: [L1359](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L1359), [L1529](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L1529), [L1764](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L1764), [L1794](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L1794)
- `ai_course`: The assistant kept treating `goal_stack` and then `why_chain` as a place to preserve every useful abstraction, including paraphrases, reusable rules, and sibling reasons, even after the human explained the recursive rule multiple times.
- `human_intervention`: The human restated that each frame should be a progressive generalization of the last, challenged over-splitting, asked for the exact disconnect, and then changed the data shape to `why_chains` so sibling reasons could stay separate without weakening the recursive rule.
- `adequate_outcome`: One or more strict linear rationale chains where each next entry answers why the previous one mattered to the human, with sibling reasons represented as sibling chains.
- `event_boundary_notes`: PASS1 is correct to keep this as one event. The transcript shows repeated non-sticky corrections from [L1359](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L1359) through [L1794](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L1794).
- `human_model_signal`: "Each frame should be a progressive generalization of the last." / "How about why chains?"
- `failure_family_hypothesis`: `workflow_orchestration`; local flavor `structural non-compliance`
- `intervention_kind_hypothesis`: `tighten_contract`
- `human_cost_or_risk`: Repeated re-explanation, contract churn, and examples that still failed the rule even after the docs claimed it.
- `local_lesson_hypothesis`: When the human defines a strict recursive data structure, do not smuggle sibling reasons or distilled rules into it; if the evidence really branches, change the data shape instead.
- `cluster_hints`: `strict-why-structure`, `sibling-reasons`, `structural discipline over semantic packing`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: None material.

### C07. Daily incident review had to widen into per-day outbound-message capture and over-capture scoring

- `event_id`: `C07`
- `title`: Daily incident review had to widen into per-day outbound-message capture and over-capture scoring
- `session_or_thread`: `Task-0006 contract thread (session id 019d6942-63b6-7c72-9abc-9ab209816f97)`
- `transcript_path`: [rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl)
- `primary_refs`: [L2146](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L2146), [L2160](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L2160), [L2219](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L2219), [L2252](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L2252)
- `ai_course`: The assistant was turning reviewed examples into the main candidate layer, then adding scoring and per-day storage. That still left too much boundary judgment embedded inside the harvester's own filtering.
- `human_intervention`: The human asked for context around which outbound human statements count as corrective, one CSV per day, one folder per day, scoring-rule documentation, stronger phrase signals, and explicit bias toward over-capture instead of early filtering.
- `adequate_outcome`: A reviewable day-level candidate layer that preserves every outbound human message with provenance, heuristic sureness, and enough context for later human boundary correction.
- `event_boundary_notes`: The core event starts at [L2146](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L2146) and keeps tightening through the scoring explanation question at [L2219](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L2219) and the explicit over-capture rule at [L2252](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L2252).
- `human_model_signal`: "We don't filter at this phase just because it's not clear ... it's more important to capture the incident and understand it later."
- `failure_family_hypothesis`: `verification_proof`; local flavor `boundary-reviewability gap`
- `intervention_kind_hypothesis`: `tighten_contract`
- `human_cost_or_risk`: Real interventions can disappear before review, and boundary correction becomes harder because the raw candidate layer is too thin.
- `local_lesson_hypothesis`: When incident boundaries are uncertain, preserve a broad dated candidate layer with soft scoring before narrowing to accepted incidents.
- `cluster_hints`: `over-capture`, `reviewable boundary`, `recall-before-judgment`
- `accepted_incident_likelihood`: `possible`
- `confidence`: `medium`
- `uncertainties`: This may be too process-oriented for the accepted incident set even though the human-model signal is strong.

### C08. Accepted incidents needed a second pass for mechanism honesty and readable analysis

- `event_id`: `C08`
- `title`: Accepted incidents needed a second pass for mechanism honesty and readable analysis
- `session_or_thread`: `Task-0006 incident corpus thread (same imported contract thread family)`
- `transcript_path`: [rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl)
- `primary_refs`: [L2679](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L2679), [L2747](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L2747), [L2945](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L2945), [L2976](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L2976)
- `ai_course`: After reopening the accepted April 6 incidents, the assistant first judged them "mostly, but not fully" correct as human-course-correction records, then treated deeper causal refinement as a follow-up enhancement rather than part of adequacy. It later updated `actual_state` to reflect the real mechanism and proposed an `analysis` block so the refined miss was readable inside the incident itself.
- `human_intervention`: The human did not accept a surface-correct incident record as finished. The follow-up push was to preserve the fast correction, then do a second pass that makes the mechanism honest in `actual_state` and keeps the deeper read reviewable without relying on surrounding chat context.
- `adequate_outcome`: Accepted incidents that still preserve the original correction event, but also state the deeper mechanism truthfully and include a readable analysis payload for later reviewers.
- `event_boundary_notes`: One event. The need for a deeper causal reread appears in the question sequence around [L2747](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L2747), then becomes concrete in the `actual_state` correction at [L2945](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L2945) and the human-readable `analysis` addition at [L2976](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L2976).
- `human_model_signal`: Fast preservation is not enough if the stored incident still hides the real mechanism or requires transcript archaeology to understand the correction.
- `failure_family_hypothesis`: `workflow_orchestration`; local flavor `second-pass mechanism honesty`
- `intervention_kind_hypothesis`: `tighten_contract`
- `human_cost_or_risk`: The corpus would keep reviewable correction events but still mislead later synthesis about what actually failed and why.
- `local_lesson_hypothesis`: When an incident is accepted quickly for boundary reasons, run a second pass that makes the mechanism honest and adds a short readable explanation block instead of assuming the first preserved shape is sufficient.
- `cluster_hints`: `second-pass refinement`, `mechanism honesty`, `readable in-record analysis`
- `accepted_incident_likelihood`: `possible`
- `confidence`: `medium`
- `uncertainties`: This may be better treated as schema/process hardening around accepted incidents rather than a standalone accepted incident.

### C09. "Server is up" still used the wrong lane and hid the missing 51-mob outcome

- `event_id`: `C09`
- `title`: "Server is up" still used the wrong lane and hid the missing 51-mob outcome
- `session_or_thread`: `019d6ead-6ee7-7201-a865-9dd3f168dd1f`
- `transcript_path`: [rollout-2026-04-08T15-59-13-019d6ead-6ee7-7201-a865-9dd3f168dd1f.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T15-59-13-019d6ead-6ee7-7201-a865-9dd3f168dd1f.jsonl)
- `primary_refs`: [L861](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T15-59-13-019d6ead-6ee7-7201-a865-9dd3f168dd1f.jsonl#L861), [L876](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T15-59-13-019d6ead-6ee7-7201-a865-9dd3f168dd1f.jsonl#L876), [L976](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T15-59-13-019d6ead-6ee7-7201-a865-9dd3f168dd1f.jsonl#L976), [L1007](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T15-59-13-019d6ead-6ee7-7201-a865-9dd3f168dd1f.jsonl#L1007)
- `ai_course`: The assistant first closed on server-health evidence: map browse, socket bind, PID, and process path. That proved a runnable server process, but it did not prove the user's requested gameplay outcome and it stayed on the stale-content launch lane.
- `human_intervention`: The human redirected the pass into live server-debugging against the real target state, which was not "port 7777 is open" but whether the correct lane loaded live content and reached the expected 51-mob behavior. The assistant then proved the prior lane was booting stale cooked content and switched to the corrected editor-hosted lane.
- `adequate_outcome`: A verified lane that uses live project content and reaches the expected population outcome, with explicit evidence that the previous lane was wrong.
- `event_boundary_notes`: One event. The premature "server is up" closure at [L861](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T15-59-13-019d6ead-6ee7-7201-a865-9dd3f168dd1f.jsonl#L861) is corrected by the live-debug framing at [L876](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T15-59-13-019d6ead-6ee7-7201-a865-9dd3f168dd1f.jsonl#L876), the stale cooked-content proof at [L976](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T15-59-13-019d6ead-6ee7-7201-a865-9dd3f168dd1f.jsonl#L976), and the corrected 51-AI outcome at [L1007](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T15-59-13-019d6ead-6ee7-7201-a865-9dd3f168dd1f.jsonl#L1007).
- `human_model_signal`: A healthy socket and loaded map do not satisfy a server-debug task when the user is asking whether the live-content gameplay behavior actually happens.
- `failure_family_hypothesis`: `verification_proof`; local flavor `wrong lane / wrong success proxy`
- `intervention_kind_hypothesis`: `redirect_debugging`
- `human_cost_or_risk`: The task could have been closed as fixed while still running the wrong content path and failing the actual gameplay expectation.
- `local_lesson_hypothesis`: For live-debug requests, prove the requested gameplay or simulation outcome in the correct lane, not just process startup and network bind.
- `cluster_hints`: `closure-truth`, `wrong-lane debugging`, `outcome over startup proxy`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: None material.

### C10. The daily brief had to be turned into a repeatable first-principles evidence packet

- `event_id`: `C10`
- `title`: The daily brief had to be turned into a repeatable first-principles evidence packet
- `session_or_thread`: `019d6fe4-199d-7f50-bf08-8d68086d978e ("Boyle" child thread)`
- `transcript_path`: [rollout-2026-04-08T21-38-33-019d6fe4-199d-7f50-bf08-8d68086d978e.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T21-38-33-019d6fe4-199d-7f50-bf08-8d68086d978e.jsonl)
- `primary_refs`: [L82](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T21-38-33-019d6fe4-199d-7f50-bf08-8d68086d978e.jsonl#L82), [L100](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T21-38-33-019d6fe4-199d-7f50-bf08-8d68086d978e.jsonl#L100), [L267](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T21-38-33-019d6fe4-199d-7f50-bf08-8d68086d978e.jsonl#L267), [L310](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T21-38-33-019d6fe4-199d-7f50-bf08-8d68086d978e.jsonl#L310)
- `ai_course`: The assistant initially treated the daily brief as a summary artifact and used projected or rewritten material inside sections labeled as raw transcript. It then shifted toward source-bounded extraction, verbatim transcript windows, and verbatim JSON incident bodies after the contract problem was made explicit.
- `human_intervention`: The human reframed the problem as prompt and packet design, not just this one bad brief. The adequacy target became a repeatable first-principles evidence packet that forces separation among raw evidence, minimal reconstruction, and derived interpretation.
- `adequate_outcome`: A durable daily brief contract that preserves literal source windows where required and inlines incident bodies verbatim instead of turning them into another layer of prose summary.
- `event_boundary_notes`: One event spanning three tightenings: the initial complaint that the "raw transcript windows" block was actually summarized projection at [L82](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T21-38-33-019d6fe4-199d-7f50-bf08-8d68086d978e.jsonl#L82), the prompt-contract reframing at [L100](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T21-38-33-019d6fe4-199d-7f50-bf08-8d68086d978e.jsonl#L100), and the conversion to verbatim JSON incident bodies at [L267](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T21-38-33-019d6fe4-199d-7f50-bf08-8d68086d978e.jsonl#L267) through [L310](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T21-38-33-019d6fe4-199d-7f50-bf08-8d68086d978e.jsonl#L310).
- `human_model_signal`: "How do we prompt a model so that we get a repeatable first-principles packet..." The durable need is a source-bounded extraction contract, not a one-off cleanup pass.
- `failure_family_hypothesis`: `workflow_orchestration`; local flavor `source-bounded evidence packet`
- `intervention_kind_hypothesis`: `reframe_problem`
- `human_cost_or_risk`: Daily reports would look polished but lose audit value because raw, reconstructed, and derived content were being blended together.
- `local_lesson_hypothesis`: When a review artifact is meant to survive reinspection, design the packet so each layer has a hard provenance rule instead of relying on the model to "summarize carefully."
- `cluster_hints`: `first-principles packet`, `source-bounded extraction`, `raw-vs-derived separation`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: The event spans several concrete edits, but they all serve one tighter adequacy model for the same packet.

### C11. Child-thread `role:user` text was rejected as false human provenance

- `event_id`: `C11`
- `title`: Child-thread `role:user` text was rejected as false human provenance
- `session_or_thread`: `019d6fe4-199d-7f50-bf08-8d68086d978e ("Boyle" child thread)`
- `transcript_path`: [rollout-2026-04-08T21-38-33-019d6fe4-199d-7f50-bf08-8d68086d978e.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T21-38-33-019d6fe4-199d-7f50-bf08-8d68086d978e.jsonl)
- `primary_refs`: [L402](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T21-38-33-019d6fe4-199d-7f50-bf08-8d68086d978e.jsonl#L402), [L428](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T21-38-33-019d6fe4-199d-7f50-bf08-8d68086d978e.jsonl#L428), [L563](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T21-38-33-019d6fe4-199d-7f50-bf08-8d68086d978e.jsonl#L563), [L588](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T21-38-33-019d6fe4-199d-7f50-bf08-8d68086d978e.jsonl#L588)
- `ai_course`: The assistant first treated text present in the child thread's `role:user` payload as if it might count as direct human wording. On reread, it separated the child thread from its parent, recognized the relay path, and concluded that child-thread `role:user` content was provenance-insufficient for "raw transcript" claims.
- `human_intervention`: The human kept pressing on whether the suspect phrase was truly verbatim human wording or a paraphrase, and then asked for a hard rule. That forced the process to distinguish actual human-authored turns, relayed supervisor instructions, and agent paraphrases instead of collapsing them under a generic user-role label.
- `adequate_outcome`: Only direct parent-thread human-authored turns count as raw human transcript for provenance-sensitive reporting; relayed or subagent-spawned text must be marked separately.
- `event_boundary_notes`: One event. The initial provenance check starts at [L402](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T21-38-33-019d6fe4-199d-7f50-bf08-8d68086d978e.jsonl#L402), the relay-path discovery is explicit at [L563](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T21-38-33-019d6fe4-199d-7f50-bf08-8d68086d978e.jsonl#L563), and the hard-rule request is answered at [L588](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T21-38-33-019d6fe4-199d-7f50-bf08-8d68086d978e.jsonl#L588).
- `human_model_signal`: Direct human wording should outrank relayed supervision text even when both appear under `role:user` in different thread contexts.
- `failure_family_hypothesis`: `source_attribution`; local flavor `false human provenance from child thread`
- `intervention_kind_hypothesis`: `tighten_contract`
- `human_cost_or_risk`: A reporting pipeline could falsely quote relayed or model-generated text as direct human guidance, corrupting both incident evidence and later principle extraction.
- `local_lesson_hypothesis`: In spawned child threads, `role:user` is not enough to establish human provenance; resolve the parent-thread source before claiming verbatim human text.
- `cluster_hints`: `provenance truth`, `child-thread relay`, `direct-human-over-relay`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: None material.

## 4. Likely accepted incidents

- `C03`: Same-day artifact lineage had to be respected instead of normalized into a general config complaint.
- `C04`: The Task-0006 plan had to be rebuilt around falsifiable upstream reasoning.
- `C05`: Incident examples had to stop reading like bug reports and regain a concrete correction boundary.
- `C06`: The strict `why_chain` rule did not stick until the shape changed to `why_chains`.
- `C09`: "Server is up" still used the wrong lane and hid the missing 51-mob outcome.
- `C10`: The daily brief had to be turned into a repeatable first-principles evidence packet.
- `C11`: Child-thread `role:user` text was rejected as false human provenance.

## 5. Likely non-incident but still important intervention events

- `C01`: Wrong-seam debugging around the Android progress card and server-side ownership model.
- `C02`: Closure language outran the real upload outcome until the human restated the human-facing completion bar.
- `C07`: Daily review broadened into over-capture and per-day outbound-message storage.
- `C08`: Accepted incidents needed a second pass for mechanism honesty and readable analysis.

## 6. Repeated cluster hints noticed across the analyzed set

- `closure-truth` / `real-world completion mismatch`: `C02`, `C03`, `C09`
- `falsifiable-contract discipline`: `C04`, `C05`, `C06`, `C08`
- `boundary-reviewability` / `over-capture before pruning`: `C05`, `C07`, `C10`, `C11`
- `source and provenance truth`: `C10`, `C11`
- `wrong-seam or wrong-lane debugging`: `C01`, `C09`

These are only local hints from the April 8 set, not cross-corpus principles.

## 7. Strongest human-model signals to carry into later clustering or principle work

- `C03` at [L545](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L545): same-day artifact lineage matters when deciding whether a correction belongs to the same incident.
- `C04` at [L729](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L729): incident reasoning should be inspectable and falsifiable, not thematic.
- `C05` at [L950](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L950): the qualification gate is "human course correction incident," not generic bug or defect reporting.
- `C06` at [L1359](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L1359) and [L1784](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L1784): each why-step should progressively generalize the previous one; branch when the evidence branches.
- `C07` at [L2252](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L2252): over-capture is preferable to early filtering when boundary judgment is still uncertain.
- `C10` at [L100](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T21-38-33-019d6fe4-199d-7f50-bf08-8d68086d978e.jsonl#L100): the durable need is a repeatable first-principles packet, not a nicer summary.
- `C11` at [L588](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T21-38-33-019d6fe4-199d-7f50-bf08-8d68086d978e.jsonl#L588): direct human wording should outrank relayed supervisor text as a hard provenance rule.

## 8. Events That Still Need a Wider Reread

- `C01`: Likely important, but the transcript slice is still narrow enough that a fuller upstream/downstream reread could sharpen whether this belongs in the accepted incident set.
- `C02`: The adequacy signal is strong, but a wider reread would help separate closure-language failure from the underlying Android/server implementation miss.
- `C10`: Strong candidate, but it spans several brief/prompt edits and could benefit from a wider reread before any later clustering work tries to reduce it too aggressively.
