# INTERVENTION-PASS3 Principle Report

Source date: `2026-04-08`

Pass scope: principle extraction from PASS2 only. No prompts, schemas, or shared workflow docs were modified.

Result: three kept principles. This is the minimum honest set that covers the recurring April 8 decision failures without turning each event into its own rule.

## 1. Source scope analyzed

- Source window: April 8, 2026 intervention analyses from `C01` through `C11`.
- PASS2 artifact in scope: [INTERVENTION-PASS2.md](/c:/Agent/CodexDashboard/Tracking/Task-0006/Research/Daily/2026-04-08/INTERVENTION-PASS2.md)
- PASS2 sections relied on most:
  - per-event records at [INTERVENTION-PASS2.md#L30](/c:/Agent/CodexDashboard/Tracking/Task-0006/Research/Daily/2026-04-08/INTERVENTION-PASS2.md#L30)
  - repeated cluster hints at [INTERVENTION-PASS2.md#L278](/c:/Agent/CodexDashboard/Tracking/Task-0006/Research/Daily/2026-04-08/INTERVENTION-PASS2.md#L278)
  - strongest human-model signals at [INTERVENTION-PASS2.md#L288](/c:/Agent/CodexDashboard/Tracking/Task-0006/Research/Daily/2026-04-08/INTERVENTION-PASS2.md#L288)
  - likely accepted vs non-incident split at [INTERVENTION-PASS2.md#L261](/c:/Agent/CodexDashboard/Tracking/Task-0006/Research/Daily/2026-04-08/INTERVENTION-PASS2.md#L261) and [INTERVENTION-PASS2.md#L271](/c:/Agent/CodexDashboard/Tracking/Task-0006/Research/Daily/2026-04-08/INTERVENTION-PASS2.md#L271)

## 2. PASS2 artifact used

- Primary analysis artifact: [INTERVENTION-PASS2.md](/c:/Agent/CodexDashboard/Tracking/Task-0006/Research/Daily/2026-04-08/INTERVENTION-PASS2.md)
- PASS3 start point: PASS2 event analyses and PASS2-carried transcript refs only.
- PASS1 was not reopened during PASS3.
- Raw transcript windows were not reopened during PASS3 because PASS2 already preserved the needed boundary corrections, primary refs, and human-model signals for the kept principle boundaries.

## 3. Candidate clusters considered

- `CL01 outcome truth over proxy completion`
  - supporting events: `C02`, `C03`, `C09`
  - decision failure: claiming success from runner health, server startup, or stale fallback rather than the requested end state, lane, or artifact lineage
  - disposition: kept as `P01`

- `CL02 falsifiable, event-grounded incident contract`
  - supporting events: `C04`, `C05`
  - decision failure: preserving themes or bug-like abstractions instead of a concrete, inspectable correction event
  - disposition: merged into `P02`

- `CL03 structurally honest rationale representation`
  - supporting events: `C06`
  - decision failure: stuffing sibling reasons and distilled rules into a recursive structure that was supposed to stay linear and falsifiable
  - disposition: merged into `P02`

- `CL04 mechanism honesty after fast preservation`
  - supporting events: `C08`
  - decision failure: treating a surface-correct preserved incident as finished even when the stored mechanism was still misleading
  - disposition: merged into `P02`

- `CL05 reviewable capture before pruning`
  - supporting events: `C07`
  - decision failure: embedding boundary judgment too early and losing reviewable candidate evidence
  - disposition: merged into `P03`

- `CL06 source-bounded evidence packet`
  - supporting events: `C10`
  - decision failure: presenting reconstructed or summarized material as if it were raw source
  - disposition: merged into `P03`

- `CL07 direct-human provenance over relayed text`
  - supporting events: `C11`
  - decision failure: treating child-thread `role:user` text as direct human wording without resolving the source path
  - disposition: merged into `P03`

- `CL08 scope discipline after seam confirmation`
  - supporting events: `C01`
  - decision failure: continuing to widen into adjacent seams after the requested fix had already been narrowed
  - disposition: rejected as a standalone principle for this day because support is weak and too close to ordinary collaborative narrowing

## 4. Final kept principles

### P01

- `principle_id`: `P01`
- `principle_statement`: Do not call work done from proxy health checks or fallback artifacts when the request implies a specific end state, lane, or artifact lineage; verify the requested end state directly or surface the blocker.
- `decision_point`: Closing a debugging, build, launch, or automation task and deciding whether the current evidence is enough to claim success.
- `failure_signature`: `calling it done from proxy proof or stale fallback`
- `why_this_is_durable`: The same miss appeared across scheduled digest handling, same-day server launch, and live server debugging. The shared human standard was stable: success meant the requested observable outcome or artifact lineage, not intermediate runner activity or a merely healthy process.
- `supporting_events`:
  - `C02` digest success reset around real end-state proof: [INTERVENTION-PASS2.md#L51](/c:/Agent/CodexDashboard/Tracking/Task-0006/Research/Daily/2026-04-08/INTERVENTION-PASS2.md#L51)
    - core refs: [L178](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T13-03-23-019d6e0c-7467-7042-92f5-aa976e10a3fe.jsonl#L178), [L190](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T13-03-23-019d6e0c-7467-7042-92f5-aa976e10a3fe.jsonl#L190)
    - support weight: weaker but aligned
  - `C03` stale server fallback violated same-day build intent: [INTERVENTION-PASS2.md#L72](/c:/Agent/CodexDashboard/Tracking/Task-0006/Research/Daily/2026-04-08/INTERVENTION-PASS2.md#L72)
    - core refs: [L477](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T15-59-13-019d6ead-6ee7-7201-a865-9dd3f168dd1f.jsonl#L477), [L545](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T15-59-13-019d6ead-6ee7-7201-a865-9dd3f168dd1f.jsonl#L545), [L558](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T15-59-13-019d6ead-6ee7-7201-a865-9dd3f168dd1f.jsonl#L558)
  - `C09` server-health proof hid the wrong lane and missing gameplay outcome: [INTERVENTION-PASS2.md#L198](/c:/Agent/CodexDashboard/Tracking/Task-0006/Research/Daily/2026-04-08/INTERVENTION-PASS2.md#L198)
    - core refs: [L861](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T15-59-13-019d6ead-6ee7-7201-a865-9dd3f168dd1f.jsonl#L861), [L876](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T15-59-13-019d6ead-6ee7-7201-a865-9dd3f168dd1f.jsonl#L876), [L976](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T15-59-13-019d6ead-6ee7-7201-a865-9dd3f168dd1f.jsonl#L976), [L1007](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T15-59-13-019d6ead-6ee7-7201-a865-9dd3f168dd1f.jsonl#L1007)
- `supporting_human_model_signals`:
  - `C02`: "I want a failure email if it didn't reach the end state - email sent, or report saved." [INTERVENTION-PASS2.md#L62](/c:/Agent/CodexDashboard/Tracking/Task-0006/Research/Daily/2026-04-08/INTERVENTION-PASS2.md#L62)
  - `C03`: "If I asked you to build the server right? ... the expectation would be that you launch the same server from today right?" [INTERVENTION-PASS2.md#L83](/c:/Agent/CodexDashboard/Tracking/Task-0006/Research/Daily/2026-04-08/INTERVENTION-PASS2.md#L83)
  - `C09`: A healthy socket and loaded map do not satisfy a server-debug request when the asked-for live-content gameplay behavior is still unproven. [INTERVENTION-PASS2.md#L209](/c:/Agent/CodexDashboard/Tracking/Task-0006/Research/Daily/2026-04-08/INTERVENTION-PASS2.md#L209)
- `counterfactual_prevention_claim`: If this rule had been applied earlier, the digest lane would have checked end-state artifacts first, the stale April 2 fallback would have been surfaced as a blocker instead of normalized, and the server-debug lane would have switched sooner from startup proof to gameplay proof.
- `scope_and_non_goals`: This is not a demand for exhaustive end-to-end proof on every task. It applies when the request names a human-visible outcome, a specific live lane, or a specific artifact lineage. Process-health checks still matter, but they are supporting evidence, not closure by themselves.
- `pre_action_question`: Have I verified the exact outcome or artifact the user asked for, or am I about to declare success from a proxy?
- `operational_check`: `closure gate` Ask: if process-health evidence vanished, what direct evidence would still prove the requested end state, lane, or artifact lineage?
- `confidence`: `strong`

### P02

- `principle_id`: `P02`
- `principle_statement`: For durable review artifacts, store only reasoning that stays falsifiable from the concrete correction event; ground fields in the event, keep mechanism truth visible, and change the data shape when the evidence branches instead of packing extra abstractions into one field.
- `decision_point`: Defining or filling a plan, schema, incident record, rationale chain, or other durable artifact that later reviewers will inspect without the surrounding chat.
- `failure_signature`: `hand-wavy or overloaded structure that hides the actual correction`
- `why_this_is_durable`: The same adequacy rule recurred across plan design, incident qualification, recursive why-structure, and accepted-incident refinement. The human repeatedly rejected artifacts that sounded plausible but were not inspectable, concretely grounded, or structurally honest.
- `supporting_events`:
  - `C04` plan had to become falsifiable and machine-checkable: [INTERVENTION-PASS2.md#L93](/c:/Agent/CodexDashboard/Tracking/Task-0006/Research/Daily/2026-04-08/INTERVENTION-PASS2.md#L93)
    - core refs: [L729](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L729), [L854](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L854), [L869](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L869)
  - `C05` incident examples had to qualify as concrete human-course-correction events: [INTERVENTION-PASS2.md#L114](/c:/Agent/CodexDashboard/Tracking/Task-0006/Research/Daily/2026-04-08/INTERVENTION-PASS2.md#L114)
    - core refs: [L925](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L925), [L950](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L950), [L1143](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L1143)
  - `C06` strict why progression required sibling branches to become sibling chains: [INTERVENTION-PASS2.md#L135](/c:/Agent/CodexDashboard/Tracking/Task-0006/Research/Daily/2026-04-08/INTERVENTION-PASS2.md#L135)
    - core refs: [L1359](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L1359), [L1764](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L1764), [L1794](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L1794)
  - `C08` accepted incidents still needed mechanism-honest refinement: [INTERVENTION-PASS2.md#L177](/c:/Agent/CodexDashboard/Tracking/Task-0006/Research/Daily/2026-04-08/INTERVENTION-PASS2.md#L177)
    - core refs: [L2747](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L2747), [L2945](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L2945), [L2976](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L2976)
- `supporting_human_model_signals`:
  - `C04`: "Trace the incident upstream, keeping the interpretation falsifiable" and "Your plan isn't falsifiable, too much hand-waviness." [INTERVENTION-PASS2.md#L104](/c:/Agent/CodexDashboard/Tracking/Task-0006/Research/Daily/2026-04-08/INTERVENTION-PASS2.md#L104)
  - `C05`: "These are really 'human course correction incident', not just 'any incident' in the abstract." / "`expected_state` ... `actual_state` ... should reference the concrete event." [INTERVENTION-PASS2.md#L125](/c:/Agent/CodexDashboard/Tracking/Task-0006/Research/Daily/2026-04-08/INTERVENTION-PASS2.md#L125)
  - `C06`: "Each frame should be a progressive generalization of the last." / "How about why chains?" [INTERVENTION-PASS2.md#L146](/c:/Agent/CodexDashboard/Tracking/Task-0006/Research/Daily/2026-04-08/INTERVENTION-PASS2.md#L146)
  - `C08`: Fast preservation is not enough when the stored incident still hides the real mechanism or requires transcript archaeology. [INTERVENTION-PASS2.md#L188](/c:/Agent/CodexDashboard/Tracking/Task-0006/Research/Daily/2026-04-08/INTERVENTION-PASS2.md#L188)
- `counterfactual_prevention_claim`: If this rule had been applied earlier, the plan would have started with a falsifiable contract, incident examples would have stayed grounded in concrete correction events, sibling reasons would have been branched by structure sooner, and accepted incidents would not have needed late mechanism-honesty repair.
- `scope_and_non_goals`: This is not a command to over-formalize every scratch note or suppress abstraction entirely. It applies to durable artifacts that must survive later review and synthesis. Abstract summaries are allowed, but they cannot replace the inspectable event truth or overload fields that are supposed to hold one kind of evidence.
- `pre_action_question`: Can a later reviewer recover and challenge the concrete correction event and mechanism from this artifact alone, or am I hiding it behind themes, paraphrases, or overloaded fields?
- `operational_check`: `plan-review check` For each durable field, name the concrete event it represents; if sibling reasons exist, they must appear as sibling structures rather than being packed into one linear field.
- `confidence`: `strong`

### P03

- `principle_id`: `P03`
- `principle_statement`: When building evidence packets or candidate corpora, preserve provenance and reviewability before pruning or summarizing: over-capture uncertain candidates, keep raw, reconstructed, and derived layers separate, and only label text as direct human wording when the source path proves it.
- `decision_point`: Harvesting, quoting, summarizing, or packaging transcript-derived evidence for later review, incident selection, or prompt design.
- `failure_signature`: `summarized or misattributed evidence presented as raw source`
- `why_this_is_durable`: The same failure recurred in daily candidate capture, the daily brief packet, and provenance-sensitive quoting. The shared human standard was that later review must still be able to trust what was captured, what was inferred, and where each piece came from.
- `supporting_events`:
  - `C07` uncertain boundaries required over-capture before filtering: [INTERVENTION-PASS2.md#L156](/c:/Agent/CodexDashboard/Tracking/Task-0006/Research/Daily/2026-04-08/INTERVENTION-PASS2.md#L156)
    - core refs: [L2146](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L2146), [L2219](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L2219), [L2252](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L2252)
  - `C10` daily brief had to become a first-principles, source-bounded packet: [INTERVENTION-PASS2.md#L219](/c:/Agent/CodexDashboard/Tracking/Task-0006/Research/Daily/2026-04-08/INTERVENTION-PASS2.md#L219)
    - core refs: [L82](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T21-38-33-019d6fe4-199d-7f50-bf08-8d68086d978e.jsonl#L82), [L100](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T21-38-33-019d6fe4-199d-7f50-bf08-8d68086d978e.jsonl#L100), [L267](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T21-38-33-019d6fe4-199d-7f50-bf08-8d68086d978e.jsonl#L267), [L310](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T21-38-33-019d6fe4-199d-7f50-bf08-8d68086d978e.jsonl#L310)
  - `C11` child-thread `role:user` text failed direct-human provenance: [INTERVENTION-PASS2.md#L240](/c:/Agent/CodexDashboard/Tracking/Task-0006/Research/Daily/2026-04-08/INTERVENTION-PASS2.md#L240)
    - core refs: [L402](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T21-38-33-019d6fe4-199d-7f50-bf08-8d68086d978e.jsonl#L402), [L563](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T21-38-33-019d6fe4-199d-7f50-bf08-8d68086d978e.jsonl#L563), [L588](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T21-38-33-019d6fe4-199d-7f50-bf08-8d68086d978e.jsonl#L588)
- `supporting_human_model_signals`:
  - `C07`: "We don't filter at this phase just because it's not clear ... it's more important to capture the incident and understand it later." [INTERVENTION-PASS2.md#L167](/c:/Agent/CodexDashboard/Tracking/Task-0006/Research/Daily/2026-04-08/INTERVENTION-PASS2.md#L167)
  - `C10`: "How do we prompt a model so that we get a repeatable first-principles packet..." [INTERVENTION-PASS2.md#L230](/c:/Agent/CodexDashboard/Tracking/Task-0006/Research/Daily/2026-04-08/INTERVENTION-PASS2.md#L230)
  - `C11`: Direct human wording should outrank relayed supervisor text as a hard provenance rule. [INTERVENTION-PASS2.md#L251](/c:/Agent/CodexDashboard/Tracking/Task-0006/Research/Daily/2026-04-08/INTERVENTION-PASS2.md#L251)
- `counterfactual_prevention_claim`: If this rule had been applied earlier, the candidate layer would have remained broad and reviewable, the daily brief would not have mixed projected prose into "raw transcript" sections, and child-thread relay text would not have been treated as direct human wording.
- `scope_and_non_goals`: This is not a ban on summarization, heuristics, or later filtering. It requires those steps to remain visibly downstream of a provenance-preserving capture layer. It applies when the output is meant for reinspection, later incident selection, or downstream behavioral learning.
- `pre_action_question`: Am I preserving enough source-bounded material and provenance that a reviewer can still audit this later, or am I collapsing judgment into capture?
- `operational_check`: `audit question` Can each quoted or "raw" element be traced to a proven source class, and can uncertain candidates survive into review without being silently filtered out?
- `confidence`: `medium`

## 5. Rejected or merged principle candidates and why

- `candidate_statement`: Treat same-day artifact lineage as mandatory unless the human explicitly approves a fallback.
  - `status`: `merged`
  - `reason`: Concrete sub-case of `P01`; the stronger retained rule covers same-day lineage, live lane selection, and end-state verification with one closure principle.
  - `merged_into`: `P01`

- `candidate_statement`: For live-debug requests, prove the requested gameplay or simulation outcome instead of stopping at startup or socket health.
  - `status`: `merged`
  - `reason`: Same decision failure as `P01` in a different lane; keeping it separate would split one stronger closure rule into two local variants.
  - `merged_into`: `P01`

- `candidate_statement`: Every incident example must qualify as a human-course-correction event rather than a bug report.
  - `status`: `merged`
  - `reason`: Important, but it is one grounding test inside the broader `P02` rule about falsifiable, event-grounded durable artifacts.
  - `merged_into`: `P02`

- `candidate_statement`: If sibling reasons exist, replace one `why_chain` with sibling `why_chains` rather than weaken the recursive rule.
  - `status`: `merged`
  - `reason`: Structural honesty is the sharper abstraction; `P02` keeps the decision rule and avoids overfitting to one field name.
  - `merged_into`: `P02`

- `candidate_statement`: Run a second pass for mechanism honesty before calling an accepted incident finished.
  - `status`: `merged`
  - `reason`: Better treated as a closure subrule of `P02` than as a standalone principle for this day; the broader retained rule already says preserved artifacts are not done if they still hide the real mechanism.
  - `merged_into`: `P02`

- `candidate_statement`: When incident boundaries are uncertain, over-capture candidates before filtering.
  - `status`: `merged`
  - `reason`: Useful, but it is the upstream capture half of `P03`, which also keeps the equally important provenance and layer-separation rule.
  - `merged_into`: `P03`

- `candidate_statement`: Review packets must separate raw evidence, minimal reconstruction, and derived interpretation.
  - `status`: `merged`
  - `reason`: Core subrule of `P03`; keeping it separate would unnecessarily split packet design from provenance and capture discipline.
  - `merged_into`: `P03`

- `candidate_statement`: Only direct parent-thread human turns count as raw human transcript.
  - `status`: `merged`
  - `reason`: Hard provenance rule, but still a concrete sub-case of `P03` rather than a separate principle for this day-level set.
  - `merged_into`: `P03`

- `candidate_statement`: Once the requested seam is confirmed, stop widening into adjacent seams unless the human explicitly broadens scope.
  - `status`: `rejected`
  - `reason`: Supported mainly by `C01`, which PASS2 already rates as weak and close to ordinary collaborative narrowing. It is actionable, but not durable enough from this day alone to survive as one of the few kept principles.

## 6. The smallest recommended principle set for this scope

- `P01` closure truth over proxy proof
- `P02` falsifiable, event-grounded, structurally honest durable artifacts
- `P03` provenance-preserving evidence capture before pruning or summarizing

Why this is the minimum honest set:

- `P01` covers the April 8 completion and wrong-lane failures in `C02`, `C03`, and `C09`.
- `P02` covers the durable artifact and reasoning-shape failures in `C04`, `C05`, `C06`, and `C08`.
- `P03` covers the evidence-packet and provenance failures in `C07`, `C10`, and `C11`.
- Compressing further would blur distinct decision points: closure proof, artifact structure, and evidence provenance.

## 7. Principles still too weak and needing more days or more events

- Standalone `scope discipline after seam confirmation` needs more recurrence beyond `C01` before it should be kept as a durable principle instead of a local debugging lesson.
- Standalone `second-pass mechanism honesty` may deserve its own principle later, but on April 8 it is better treated as a subrule of `P02` unless more days show it recurring independently of general artifact falsifiability.
- Standalone `over-capture before pruning` may later separate from `P03`, but current support is strongest when kept tied to provenance-preserving packet design rather than elevated alone.

## 8. Transcript windows reopened during PASS3 and why

- None.
- Reason: PASS2 already preserved enough event boundaries, human-model signals, and transcript refs to keep the three retained principles evidence-grounded without outrunning the record.
