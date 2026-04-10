# INTERVENTION-PASS3

## Source scope analyzed

- Source day: `2026-04-07`
- PASS3 prompt used: [INTERVENTION-PASS3.md](/c:/Users/gregs/.codex/Orchestration/Prompts/INTERVENTION-PASS3.md)
- PASS2 event set in scope: `C01` through `C20`
- Supporting local doc used during PASS3: [README.md](/c:/Users/gregs/.codex/Orchestration/Reports/Incidents/README.md)
- PASS1 reread during PASS3: none
- Raw transcripts reopened during PASS3: none
- PASS3 basis: cluster from the April 7 PASS2 artifact, not from independent day recall

## PASS2 artifact used

- Primary artifact: [INTERVENTION-PASS2.md](/c:/Agent/CodexDashboard/Tracking/Task-0006/Research/Daily/2026-04-07/INTERVENTION-PASS2.md)
- PASS2 sections used most heavily:
  - candidate boundary corrections
  - per-event analysis records
  - repeated cluster hints
  - strongest human-model signals
  - likely accepted incidents versus likely non-incident intervention events

## Candidate clusters considered

### CC01 - Exact human-requested operating frame over easier proxy

- Events: `C01`, `C03`, `C06`, `C10`
- Status: kept as `P01`
- Why: the recurring miss is silently solving a nearby easier target instead of the named stack, audience, visible surface, or user-visible harmed state.

### CC02 - Human-controlled assets are deny-by-default and recovery comes before abstraction

- Events: `C07`, `C08`, `C09`, `C10`
- Status: kept as `P02`
- Why: the human standard stayed consistent from permission boundary through remediation sequence: do not touch live human state casually, and if harm is possible, inventory, preserve, and restore before theorizing.

### CC03 - Ownership and closure state must stay explicit

- Events: `C04`, `C05`, `C20`
- Status: kept as `P03`
- Why: each event is a false-safe state around active ownership: latent delegated writer, idle stop after batch closure, or corpus completion claim with an obvious missing incident.

### CC04 - Review claims need local, sufficient, first-principles evidence

- Events: `C02`, `C13`, `C16`, `C17`, `C18`, `C19`
- Status: kept as `P04`
- Why: quoting, incident grounding, brief design, heavyweight incident shape, and isolated-run validity all hinge on the same adequacy rule: a later reviewer must be able to inspect the claim from local evidence and valid conditions.

### CC05 - Contract and rationale structures must stay concrete, falsifiable, and human-bounded

- Events: `C11`, `C12`, `C13`, `C14`, `C15`
- Status: kept as `P05`
- Why: the repeated corrections all push the same way: concrete event first, hard incident admission gate, then narrow progressive generalization that stops at the human's stated ceiling.

### CC06 - Standalone audience-calibration principle

- Events: `C03`
- Status: merged into `P01`
- Why: real but too narrow alone; it is one operating-frame subcase, not a stronger standalone rule.

### CC07 - Standalone context-isolation principle

- Events: `C16`
- Status: merged into `P04`
- Why: valid isolation is one form of reviewability and epistemic sufficiency, not a broader principle than that cluster.

### CC08 - Standalone `why_chains` data-shape principle

- Events: `C15`
- Status: merged into `P05`
- Why: the stronger rule is not the literal field name; it is keeping rationale structures evidence-bound and using sibling chains only when the human actually expressed sibling upstream reasons.

## Final kept principles

### Principle `P01`

- `principle_id`: `P01`
- `principle_statement`: Do not silently substitute a proxy operating frame for the one the human named. Work against the specified tool stack, audience, visible surface, or user-visible target state unless the human explicitly approves a substitution.
- `decision_point`: Choosing what target frame or proof surface to optimize around before answering, debugging, testing, or repairing.
- `failure_signature`: `solving a nearby easier problem from a proxy`
- `why_this_is_durable`: The same miss appears across tooling, writing, UX review, and recovery work. The surface changes, but the underlying failure is stable: the assistant quietly swaps in the frame it can prove most easily instead of the frame the human actually asked for.
- `supporting_events`: `C01`, `C03`, `C06`, `C10`
- `supporting_human_model_signals`:
  - `C01`: if the human names the current Insights stack, that tool choice is part of the task, not an optional preference.
  - `C03`: "executive leadership" excludes engine-internal framing even when it is technically correct.
  - `C06`: "Use the app, click around like a human would" and ground only in visible screens.
  - `C10`: the harmed state to recover was the phone-visible vault story, not merely backend truth.
- `counterfactual_prevention_claim`: Applying this earlier would likely have prevented the legacy-tool detour in `C01`, the wrong-audience memo in `C03`, the proxy QA answer in `C06`, and the backend-centric repair framing in `C10`.
- `scope_and_non_goals`: This does not ban clarifying questions or user-approved alternatives. It does ban unannounced substitution once the human has already made the operating frame explicit.
- `pre_action_question`: Am I working on the exact tool, audience, surface, or user-visible state the human named, or on a proxy that is just easier for me to prove?
- `operational_check`: `prompt rule` - If the response changes the requested tool family, audience, proof surface, or target state, quote the user instruction that authorizes that change; otherwise stop and ask.
- `confidence`: `medium`

### Principle `P02`

- `principle_id`: `P02`
- `principle_statement`: Treat human-controlled devices and live data as deny-by-default, and if you may have disturbed them, pivot immediately to exact inventory, preservation, and restoration before explanation or policy design.
- `decision_point`: Planning or responding to operations that can change a human-controlled device, personal data store, or human-facing ledger.
- `failure_signature`: `touching live human state as if it were disposable, then abstracting instead of restoring`
- `why_this_is_durable`: This is a recurring human-world control rule, not a one-off phone quirk. The same standard governs permission, blast-radius awareness, emergency inventory, backup order, and what counts as adequate remediation after harm.
- `supporting_events`: `C07`, `C08`, `C09`, `C10`
- `supporting_human_model_signals`:
  - `C07`: the physical phone is human data; destructive or stateful actions require explicit prior permission.
  - `C08`: after causing harm, fix the actual damage before talking about future guardrails.
  - `C09`: in a data-loss scare, answer the literal inventory question and back up the corpus before explanatory analysis.
  - `C10`: recovery has to target the destroyed phone-visible ledger, not just explain server-side truth.
- `counterfactual_prevention_claim`: Applying this earlier would likely have prevented the phone-state breach in `C07` and changed the post-incident sequence in `C08` through `C10` toward immediate protection and restoration.
- `scope_and_non_goals`: This is about live human-controlled state. It does not forbid safe read-only inspection or explicitly authorized destructive work with a clear blast radius and rollback plan.
- `pre_action_question`: Do I have exact permission for this live human state, and if something may already have changed, have I switched fully into inventory, preserve, and restore mode?
- `operational_check`: `closure gate` - Before any operation that can change human-controlled state, record approval, likely blast radius, rollback or backup path, and the exact user-visible state that must be restored if things go wrong.
- `confidence`: `strong`

### Principle `P03`

- `principle_id`: `P03`
- `principle_statement`: Do not leave ownership in a latent or falsely closed state. Before idling, handing off, or declaring completion, either park agents safely, continue the next obvious verification or mining step, or keep the pass open until the salient gap is closed.
- `decision_point`: Deciding whether work is safe to stop, safe to hand off, or honestly complete while ownership is still active.
- `failure_signature`: `latent actor or completion claim outrunning actual responsibility`
- `why_this_is_durable`: The same human standard appears in orchestration, QA continuity, and corpus backfill: active ownership must end in an explicit safe state, not in implied idleness or optimistic completeness language.
- `supporting_events`: `C04`, `C05`, `C20`
- `supporting_human_model_signals`:
  - `C04`: an idle delegated leader should keep watch instead of sitting there as a latent conflict source.
  - `C05`: closing the current batch does not authorize idling when QA and task mining are still obvious next steps.
  - `C20`: backfill cannot be called complete while the most salient missing incident is still absent.
- `counterfactual_prevention_claim`: Applying this earlier would likely have prevented the latent TASK-LEADER conflict risk in `C04`, the idle stop in `C05`, and the false completeness claim in `C20`.
- `scope_and_non_goals`: This is about continuing ownership, delegated-agent lanes, and completion claims. It does not require never pausing after a clearly completed one-shot user request.
- `pre_action_question`: If I stop or say complete now, what active owner, latent writer, next required check, or glaring gap would still remain?
- `operational_check`: `plan-review check` - Any stop, handoff, or completion note must name the active owner, whether any delegate remains parked or detached, the next verification or mining step if ownership continues, and the most salient known gap; if any slot is blank, do not call it closed.
- `confidence`: `medium`

### Principle `P04`

- `principle_id`: `P04`
- `principle_statement`: Before claiming a quote, artifact, or validation is reviewable, make the supporting evidence and test conditions locally inspectable and sufficient for first-principles review.
- `decision_point`: Preparing an evidence-bearing answer, durable artifact, or official validation run that another reviewer may need to audit later.
- `failure_signature`: `reviewability claimed from thin, curated, adjacent, or contaminated context`
- `why_this_is_durable`: The same adequacy rule spans live quote verification, incident grounding, self-contained daily briefs, heavyweight incident JSONs, and isolated prompt tests. The medium changes, but the requirement does not: a reviewer must be able to inspect the claim from what is present and from valid test conditions, not from trust in the assistant's curation.
- `supporting_events`: `C02`, `C13`, `C16`, `C17`, `C18`, `C19`
- `supporting_human_model_signals`:
  - `C02`: only quote what can be directly verified now, and inline enough evidence for review.
  - `C13`: top-level incident state must stay concrete enough that the event is legible from the incident alone.
  - `C16`: an inherited-context run is invalid for durable-instruction testing because the conditions are contaminated.
  - `C17`: "self-contained" means the important incident content itself must live in the artifact.
  - `C18`: second-opinion review must preserve enough raw, contiguous context for first-principles disagreement.
  - `C19`: heavyweight verbatim evidence belongs in the incident JSON itself, not only in a derived brief.
- `counterfactual_prevention_claim`: Applying this earlier would likely have prevented the unverified Epic quote in `C02`, the too-abstract incident grounding in `C13`, the invalid isolated-run result in `C16`, and the thin or curated brief and incident evidence problems in `C17` through `C19`.
- `scope_and_non_goals`: This does not mean every file must contain an entire transcript. It does mean the local artifact or validation context must carry enough direct evidence and valid conditions for the specific review or claim being made.
- `pre_action_question`: Could another reviewer verify this claim or rerun this validation from what is present here, without trusting my summary, adjacent files, search snippets, or inherited context?
- `operational_check`: `audit question` - What exact evidence and run conditions live locally in this artifact, and what missing adjacent context or inherited state am I still depending on?
- `confidence`: `strong`

### Principle `P05`

- `principle_id`: `P05`
- `principle_statement`: Design incident and rationale structures so the concrete event is explicit, the admission gate is falsifiable, and generalization stops at the highest reason the human actually stated.
- `decision_point`: Designing or filling durable capture structures for incidents, explanations, and higher-level rationale.
- `failure_signature`: `hand-wavy schema or overgrown rationale smuggling in assistant theory`
- `why_this_is_durable`: The April 7 contract corrections all push on the same structural weakness: when event facts, admission rules, inferred motives, and downstream principles get mixed together, the record stops preserving the human's actual model and stops being falsifiable.
- `supporting_events`: `C11`, `C12`, `C13`, `C14`, `C15`
- `supporting_human_model_signals`:
  - `C11`: when a contract is hand-wavy, move from prose planning to concrete schema and examples.
  - `C12`: a generic bug is not an incident unless it preserves human correction of AI course.
  - `C13`: `expected_state` and `actual_state` must describe the concrete event before later abstraction.
  - `C14`: use a progressive `why_chain` that stops at the highest refinement explicitly stated in evidence.
  - `C15`: if the human expressed sibling reasons, model them as sibling `why_chains` instead of forcing one overgrown ladder.
- `counterfactual_prevention_claim`: Applying this earlier would likely have prevented the vague contract planning in `C11`, the boundary drift in `C12`, the abstract top-level fields in `C13`, and the repeated overgrown rationale structure failures in `C14` and `C15`.
- `scope_and_non_goals`: This governs durable capture structures. It does not ban later clustering or principle extraction; it requires that those later abstractions stay visibly downstream from concrete event capture.
- `pre_action_question`: Can an outside reviewer see the concrete event and the qualifying human correction here, and have I stopped my rationale exactly where the human's stated reasoning stops?
- `operational_check`: `plan-review check` - Reject any schema or example whose top-level fields are too abstract to identify the event, whose incident boundary could admit a generic bug, or whose rationale chain mixes sibling reasons and later rule extraction into one ladder.
- `confidence`: `strong`

## Rejected or merged principle candidates and why

### Candidate 1

- `candidate_statement`: Always use the newest requested tool stack instead of a legacy fallback.
- `status`: `merged`
- `reason`: Too tool-specific. The stronger recurring rule is about honoring the exact operating frame the human named, not only modern-versus-legacy tooling.
- `merged_into`: `P01`

### Candidate 2

- `candidate_statement`: Write for executive leadership, not for engine specialists, when the user says leadership.
- `status`: `merged`
- `reason`: Real correction, but it is one audience-framing subcase of the larger proxy-frame problem.
- `merged_into`: `P01`

### Candidate 3

- `candidate_statement`: Use the app like a human instead of relying on QA proxies.
- `status`: `merged`
- `reason`: Strong but narrower than the general rule against switching from the requested operating surface to an easier proof surface.
- `merged_into`: `P01`

### Candidate 4

- `candidate_statement`: Never run destructive phone operations without explicit approval.
- `status`: `merged`
- `reason`: Kept in stronger form with the post-harm response sequence, because the human standard covered both permission and remediation.
- `merged_into`: `P02`

### Candidate 5

- `candidate_statement`: After causing harm, fix it before discussing future policy.
- `status`: `merged`
- `reason`: Same control-boundary cluster as Candidate 4; weaker on its own than the combined live-state permission and recovery rule.
- `merged_into`: `P02`

### Candidate 6

- `candidate_statement`: In a data-loss scare, count exact files and back them up before analysis.
- `status`: `merged`
- `reason`: This is the emergency-recovery subcase of `P02`, not a broader standalone principle.
- `merged_into`: `P02`

### Candidate 7

- `candidate_statement`: Idle delegated agents must be parked in passive watch mode.
- `status`: `merged`
- `reason`: Too orchestration-specific alone. The stronger rule is explicit ownership and closure truth across latent actors, idle stops, and premature completion claims.
- `merged_into`: `P03`

### Candidate 8

- `candidate_statement`: Official prompt-behavior tests must be context-isolated.
- `status`: `merged`
- `reason`: Strong local rule, but best treated as one validity check inside the larger reviewability and epistemic-sufficiency principle.
- `merged_into`: `P04`

### Candidate 9

- `candidate_statement`: `DAILY-BRIEF` must inline accepted incidents and enough raw evidence to support first-principles review.
- `status`: `merged`
- `reason`: This is one artifact-specific expression of the broader local-evidence rule.
- `merged_into`: `P04`

### Candidate 10

- `candidate_statement`: Heavyweight verbatim evidence belongs inside the incident JSON.
- `status`: `merged`
- `reason`: Strong explicit human rule, but it is still a placement-specific subcase of the reviewability principle.
- `merged_into`: `P04`

### Candidate 11

- `candidate_statement`: Use `why_chains` instead of `goal_stack` or a single overgrown `why_chain`.
- `status`: `merged`
- `reason`: The literal field choice is not the durable rule. The durable rule is to keep rationale structures concrete, recursive, and capped by the human's stated reasons.
- `merged_into`: `P05`

### Candidate 12

- `candidate_statement`: Be more self-contained.
- `status`: `rejected`
- `reason`: Too vague to guide a future decision. It hides the real standard, which is local evidence sufficiency for first-principles review.

### Candidate 13

- `candidate_statement`: Communicate more clearly with the human.
- `status`: `rejected`
- `reason`: Too broad and not operationalizable enough. It would collapse distinct failure modes that need different checks.

## The smallest recommended principle set for this scope

- `P01` - exact human-requested operating frame over proxy substitution
- `P02` - deny-by-default for live human state, with recovery-first response if disturbed
- `P03` - explicit ownership and closure truth
- `P04` - local, sufficient evidence and valid conditions for reviewability
- `P05` - concrete, falsifiable, human-bounded incident and rationale structures

Five principles is the smallest set I can defend for this day without blurring distinct decision points. Compressing below five would force at least one harmful merge:

- `P01` and `P02` would collapse target-frame fidelity into human-world safety and remediation.
- `P03` would disappear into generic "be responsible" language.
- `P04` and `P05` could be merged, but that would erase the distinction between evidence sufficiency and schema or rationale design discipline.

## Principles still too weak for standalone promotion

- `Dedicated audience-calibration principle`
  - Current support is mainly `C03`.
  - It is better carried as a subcase of `P01` unless more days show the same audience-mismatch pattern.

- `Dedicated passive-watch delegation principle`
  - Current support is mainly `C04`, with weaker echoes in `C05`.
  - It is better kept inside `P03` unless more delegated-agent incidents appear.

- `Dedicated context-isolation principle`
  - `C16` is explicit and important, but still single-event heavy for standalone principle status.
  - It is better kept inside `P04` unless more invalid official runs recur.

## Transcript windows reopened during PASS3 and why

- None.
- PASS2 already carried clear event boundaries, repeated cluster hints, and explicit human-model signals for the five kept principles.
- No principle statement required a raw transcript reread to avoid outrunning the evidence.
