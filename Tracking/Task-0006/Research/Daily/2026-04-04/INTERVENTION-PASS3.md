# INTERVENTION-PASS3

Source day: `2026-04-04`

## Source scope analyzed

- PASS2 event set in scope: `C01` through `C19`
- Day analyzed: `2026-04-04`
- This PASS3 stayed grounded in the PASS2 artifact's event records, boundary notes, cluster hints, and strongest human-model signals.
- No raw transcript reread was needed for PASS3 because the remaining cluster boundaries were clear enough from PASS2 to keep the principle statements evidence-bounded.

## PASS2 artifact used

- `C:\Agent\CodexDashboard\Tracking\Task-0006\Research\Daily\2026-04-04\INTERVENTION-PASS2.md`

## Candidate clusters considered

- `CL01` - Concrete task and plan contract before execution: `C01`, `C02`, `C03`, `C06`, `C08`, `C16`
  - Shared decision failure: continuing with product or plan ambiguity instead of collapsing the next approved state into one durable, reviewable, measurable slice.
  - Result: kept as `P01`.
- `CL02` - Proxy proof promoted into closure truth: `C04`, `C05`, `C11`, `C12`, `C13`, `C14`
  - Shared decision failure: calling work ready or done from smoke, structured state, stale process state, or degraded fallback instead of the real human-facing lane.
  - Result: merged with `CL03` into `P02`.
- `CL03` - Diagnostic surfaces outrunning their role: `C07`, `C10`
  - Shared decision failure: letting proof-only tooling or broad investigation continue after evidence should only be narrowing the seam.
  - Result: merged into `P02` because the stronger rule is that proxy evidence may guide branching but may not replace real-surface truth.
- `CL04` - Control-surface semantic mixing in the Jobs lane: `C15`, `C17`, `C18`, `C19`
  - Shared decision failure: making navigation clicks do action work, blurring hierarchy, and putting the wrong truth layer first.
  - Result: kept as `P03`.
- `CL05` - Immediate live blocker regression handling: `C09`
  - Shared decision failure: allowing an actively user-facing blocker to sit while other work continues.
  - Result: rejected as a standalone principle for this day because it is real but too local; it fits better as an enforcement example of `P02`'s live-surface authority rule.

## Final kept principles

### P01

- `principle_id`: `P01`
- `principle_statement`: `Before implementation or continued exploration, collapse the next task or plan state into one durable, falsifiable, mockup-ready outcome and smallest honest slice that a cold reviewer can evaluate.`
- `decision_point`: `When defining, rewriting, or re-grounding a task or plan for a human-facing slice.`
- `failure_signature`: `working from product superposition or an unwritten smallest slice`
- `why_this_is_durable`: `The same miss recurred in task harvesting, prompt hardening, and live task planning. The human repeatedly demanded reviewable specificity, measurability, concrete positive outcomes, and on-disk state rather than open-ended narration.`
- `supporting_events`: `C01`, `C02`, `C03`, `C06`, `C08`, `C16`
- `supporting_human_model_signals`: `C01` says the producer must collapse "superposition of states"; `C02` replaces softer language with "falsifiable or evaluat-able claim"; `C03` requires resolution wording clear enough for a hard mockup; `C16` says a good task must be specific and measurable enough to mark done; `C06` and `C08` show the same rule at planning and execution time when durable state or the smallest honest slice was still not on disk.`
- `counterfactual_prevention_claim`: `If this rule had been applied earlier, it likely would have prevented most of the repeated task-language rewrites and reduced the later anti-drift interventions that forced planning checkpoints and smallest-slice landings.`
- `scope_and_non_goals`: `This does not require a full design spec for every small bug. It requires enough concrete, positive, human-reviewable specificity that the next state and next slice are no longer guesswork.`
- `pre_action_question`: `Could a cold reviewer tell exactly what user-facing state should appear, what counts as done, and what smallest patch lands next from what I have written right now?`
- `operational_check`: `Plan-review check: reject any task or plan text that still permits multiple plausible end states or lacks a concrete next slice on disk.`
- `confidence`: `strong`

### P02

- `principle_id`: `P02`
- `principle_statement`: `Treat proxy evidence as a branch-narrowing tool only; do not call work ready or done until the real live path meets the exact human-facing bar defined by the task or repo.`
- `decision_point`: `When choosing proof, narrowing a defect seam, or deciding whether readiness, closure, or regression claims are allowed.`
- `failure_signature`: `promoting proxy proof into closure truth`
- `why_this_is_durable`: `This failure appeared across server restart readiness, visibility proof, stale-process verification, regression closure, proof-only camera use, and reference-pose fallback. The recurring correction is not "gather more evidence"; it is "stop mistaking proxy evidence for the real bar."`
- `supporting_events`: `C04`, `C05`, `C07`, `C10`, `C11`, `C12`, `C13`, `C14`
- `supporting_human_model_signals`: `C04` says overclaimed restart robustness should be recorded as inadequate or dishonest coverage; `C05` says structured runtime state alone is not enough; `C07` says the live gameplay camera remains truth and proof views stay bounded; `C10` says a discriminating run should collapse the defect seam instead of broadening the search; `C13` says repo-local regression means the real app-surface path, not smoke; `C14` says reference pose, blocked checkpoints, and other degraded fallbacks do not count as completion.`
- `counterfactual_prevention_claim`: `If this rule had been applied earlier, it likely would have prevented the reboot-readiness overclaim, blocked smoke-shaped regression closure, cut off reference-pose closure drift, and shortened proof-helper expansion once the real seam was isolated.`
- `scope_and_non_goals`: `This does not ban smoke checks, structured state, screenshots, or proof-only views. It limits their role: they can support diagnosis or branch narrowing, but they cannot substitute for the real operating surface or the named human-facing done bar.`
- `pre_action_question`: `Am I using this artifact only to narrow the next branch, or am I about to treat it as proof that the real live path already satisfies the true done bar?`
- `operational_check`: `Closure gate: name the exact live surface, interaction path, and human-facing proof required; block closure if the evidence is only smoke, structured state, stale-process output, or degraded fallback.`
- `confidence`: `strong`

### P03

- `principle_id`: `P03`
- `principle_statement`: `Keep control-surface semantics honest: navigation should only navigate, hierarchy should visibly separate navigation from tab-owned content, and the default view should show declared intent before reconcile or runtime side effects.`
- `decision_point`: `When designing or patching a control surface that mixes navigation, status, and actions.`
- `failure_signature`: `entry clicks that both navigate and act while the surface blurs hierarchy or source of truth`
- `why_this_is_durable`: `The Jobs-lane corrections were not isolated polish notes. They repeatedly corrected the meaning of a click, the surface's primary truth layer, and the visual hierarchy needed for operators to understand where they are and what state they are seeing.`
- `supporting_events`: `C15`, `C17`, `C18`, `C19`
- `supporting_human_model_signals`: `C15` says tab entry should not trigger reconcile and that fidelity review is part of the real fix; `C17` says primary navigation must be its own strip with content below; `C18` says declared jobs JSON plus schema is the primary durable state and reconcile is a secondary diff layer; `C19` reinforces that a dense operational panel still has to remain operable through basic affordances such as scrolling.`
- `counterfactual_prevention_claim`: `If this principle had been applied before implementation, the Jobs tab likely would not have shipped with side-effectful entry, muddled hierarchy, the wrong primary truth layer, or inaccessible overflow.`
- `scope_and_non_goals`: `This does not forbid explicit refresh, reconcile, or runtime detail views. It forbids making them implicit on navigation or presenting them before the product's declared intent when declared intent is the primary state story.`
- `pre_action_question`: `On first click into this surface, am I only changing location and showing the intended state clearly, or am I also sneaking in action work or surfacing the wrong truth layer?`
- `operational_check`: `UI review gate: for every new tab or control surface, verify click meaning, primary truth layer, hierarchy separation, and basic operability such as scrolling before closure.`
- `confidence`: `medium`

## Rejected or merged principle candidates and why

- `candidate_statement`: `Write tasks in declarative terms.`
  - `status`: `rejected`
  - `reason`: `The human explicitly rejected "declarative" as too weak in `C02`; it does not force reviewable specificity or measurable closure.`
- `candidate_statement`: `Use falsifiable or evaluable claims for task definitions.`
  - `status`: `merged`
  - `reason`: `This is a necessary part of the stronger concrete-task rule but too narrow to keep apart from measurability, mockup readiness, and on-disk next-slice truth.`
  - `merged_into`: `P01`
- `candidate_statement`: `Make every usability resolution hard-mockup-ready.`
  - `status`: `merged`
  - `reason`: `This is the positive-form strengthening of the same task-definition failure and works best as one component of the broader task-and-plan contract rule.`
  - `merged_into`: `P01`
- `candidate_statement`: `Checkpoint planning state on disk before more narration.`
  - `status`: `merged`
  - `reason`: `On this day it was part of the same broader failure to turn the next approved state into a durable, reviewable object before more exploration.`
  - `merged_into`: `P01`
- `candidate_statement`: `Once the smallest honest slice is approved, patch the product before more exploration.`
  - `status`: `merged`
  - `reason`: `This is the execution-side manifestation of the same concrete-next-state rule and did not need a separate principle yet.`
  - `merged_into`: `P01`
- `candidate_statement`: `Require screenshot-backed proof for visibility tasks.`
  - `status`: `merged`
  - `reason`: `Screenshot proof mattered because the live human-facing path was the real bar; the stronger general rule is about proxy evidence versus real closure truth.`
  - `merged_into`: `P02`
- `candidate_statement`: `Bound proof-only views to one discriminating check.`
  - `status`: `merged`
  - `reason`: `The watch-for is not proof views themselves but promoting them past their branch-narrowing role.`
  - `merged_into`: `P02`
- `candidate_statement`: `Once a discriminating run isolates the seam, stop broad diagnostics.`
  - `status`: `merged`
  - `reason`: `This is the debugging form of the same stronger proxy-evidence rule kept in `P02`.`
  - `merged_into`: `P02`
- `candidate_statement`: `Regression always means real app-surface interaction per repo-local docs.`
  - `status`: `merged`
  - `reason`: `Important and explicit, but on this day it operated as one instance of the broader requirement to prove the real live path named by the task or repo.`
  - `merged_into`: `P02`
- `candidate_statement`: `Fix live blocker regressions immediately when the user is actively hit.`
  - `status`: `rejected`
  - `reason`: `C09` is real but still too local and single-example-heavy to keep as its own principle for this day; it is better treated as a high-severity enforcement case of live-surface truth.`
- `candidate_statement`: `Navigation must be side-effect free.`
  - `status`: `merged`
  - `reason`: `Strong and explicit, but it became more useful when combined with the paired hierarchy and source-of-truth rules for control surfaces.`
  - `merged_into`: `P03`
- `candidate_statement`: `Primary navigation must be structurally separate from tab content.`
  - `status`: `merged`
  - `reason`: `This is one concrete expression of the broader control-surface semantic honesty rule.`
  - `merged_into`: `P03`
- `candidate_statement`: `Declared intent must be the primary Jobs truth and reconcile must be secondary.`
  - `status`: `merged`
  - `reason`: `This is the source-of-truth branch of the same control-surface rule and did not need to stand alone for this day.`
  - `merged_into`: `P03`
- `candidate_statement`: `Panels that can overflow must scroll.`
  - `status`: `merged`
  - `reason`: `Useful but too local as a standalone principle; it functions here as a basic operability check inside the broader control-surface rule.`
  - `merged_into`: `P03`
- `candidate_statement`: `When local docs define truth or closure, always prefer them over generic habits.`
  - `status`: `rejected`
  - `reason`: `This is probably durable, but in the April 4 evidence it appears only through stronger, more actionable rules about real-surface closure (`P02`) and declared-intent-first control surfaces (`P03`). It would be premature to keep it separately without more days.`

## The smallest recommended principle set for this scope

- `P01` - Collapse the next task or plan state into a durable, measurable, mockup-ready, on-disk next slice before more work.
- `P02` - Treat proxy evidence as branch-narrowing only; closure requires the real live path and named human-facing bar.
- `P03` - Keep control-surface semantics honest: navigation-only entry, legible hierarchy, declared-intent-first default state.

## Any principles that are still too weak and need more days or more events

- `P03` is useful and operational now, but it is still mostly supported by one Jobs-lane workstream, so it should stay `medium` until similar control-surface misses recur elsewhere.
- A possible future standalone principle about `durable-state-first / smallest-honest-slice-before-more-exploration` may deserve to split away from `P01` if more days show the same anti-drift execution failure outside task-definition work.
- A possible future standalone principle about `prefer repo-local authority over generic habits` needs more cross-task evidence before it becomes stronger than the concrete `P02` and `P03` rules it currently feeds.

## Any transcript windows reopened during PASS3 and why

- None.
- PASS3 stayed within PASS2 because the recorded boundary corrections, repeated cluster hints, strongest human-model signals, and uncertainty notes were already sufficient to:
  - merge `C07` and `C10` into the stronger `P02` proxy-evidence rule without outrunning the evidence
  - keep `C15`, `C17`, `C18`, and `C19` as one narrower `P03` control-surface principle with honest `medium` confidence
