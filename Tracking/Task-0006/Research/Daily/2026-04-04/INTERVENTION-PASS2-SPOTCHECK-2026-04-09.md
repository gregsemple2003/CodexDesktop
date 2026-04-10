# INTERVENTION-PASS2 Spot-Check

Run date: 2026-04-09

Source day: 2026-04-04

Purpose: bounded PASS2-style investigation of the April 4 spot-check candidate set, with special attention to UI/product-read, mockup fidelity, navigation semantics, and explanatory-labor events.

This pass reopens only the bounded PASS1 spot-check candidates from raw transcript refs. It is not a full accepted-incident writeup pass.

## Source Scope Analyzed

- PASS1 spot-check candidate set from:
  - `C:\Users\gregs\.codex\sessions\2026\04\04\rollout-2026-04-04T01-22-17-019d56f1-261c-7e32-8a07-7323a20c471c.jsonl`
  - `C:\Users\gregs\.codex\sessions\2026\04\04\rollout-2026-04-04T11-44-43-019d592b-0085-7ad2-8cfd-22009eae6dfa.jsonl`
  - `C:\Users\gregs\.codex\sessions\2026\04\04\rollout-2026-04-04T13-23-58-019d5985-dbb1-77d0-a53d-debab8831378.jsonl`
- Candidate ids analyzed: `SC-01` through `SC-06`

## Candidate Boundary Corrections Relative To PASS1

- `SC-04` is treated as beginning at `L968`, not just `L994`, because the stale-process confusion is part of the same reopened real-world check: the assistant had already claimed fixed closure, the human reopened the real surface, first saw missing current work, then immediately discovered the deeper click-path and fidelity miss.
- No other boundary changes from PASS1.

## Per-Event Analysis

### SC-01

- `event_id`: `SC-01`
- `title`: Negative-only "Safe on server" wording rejected as not hard-mockup-ready
- `session_or_thread`: Crystallize main thread, task-writing follow-up
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\04\rollout-2026-04-04T01-22-17-019d56f1-261c-7e32-8a07-7323a20c471c.jsonl`
- `primary_refs`: `L452`, `L455-L532`
- `ai_course`: The assistant had already elevated the task-writing bar to "falsifiable or evaluable" and was treating the usability-task refinement as mostly complete.
- `human_intervention`: The human said the first sentence was still underspecified, explicitly contrasted "what it won't be" with "the exact UI," introduced the hard-mockup bar, and required prompt iteration plus fresh subagent testing.
- `adequate_outcome`: A task statement that positively specifies the safe-state UI composition of the Home card.
- `event_boundary_notes`: This is a clean boundary. The assistant had just summarized a durable rule change, and the human immediately rejected the remaining concrete task wording as still below the intended bar.
- `human_model_signal`: The human explicitly stated that a usability-task resolution must be clear enough "to produce a hard mockup." Negative-only UI deltas are not enough.
- `failure_family_hypothesis`: `usability_state_truth`
- `intervention_kind_hypothesis`: `tighten_contract`
- `human_cost_or_risk`: The task would still permit a product-distorting implementation because the visible safe state remained underdescribed.
- `local_lesson_hypothesis`: When the issue is human-facing UI/product-read, task closure needs positive surface description, not only the removal of a bad element.
- `cluster_hints`: `mockup-read`, `negative-only-state-story`, `surface-specification`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: None material inside this bounded reread.

### SC-02

- `event_id`: `SC-02`
- `title`: Human forces concrete PASS-0000 backend writes after narrative-only progress
- `session_or_thread`: CodexDashboard Task-0004 supervision thread
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\04\rollout-2026-04-04T11-44-43-019d592b-0085-7ad2-8cfd-22009eae6dfa.jsonl`
- `primary_refs`: `L281-L321`, decisive intervention at `L306`
- `ai_course`: The assistant described PASS-0000 direction and likely implementation shape, but the human still saw no product-file writes on disk.
- `human_intervention`: The human cut off further exploration, listed the minimum backend deliverable and write set, and required actual patching before more narrative commentary.
- `adequate_outcome`: A real backend slice on disk for jobs registry/discovery/bootstrap/tests.
- `event_boundary_notes`: This is narrower than the earlier planning-gate stall. I am not treating the earlier checkpoint nudge as part of this same event because here the human is correcting a later implementation-phase failure mode.
- `human_model_signal`: Disk truth matters more than status prose. "Use apply_patch and produce actual code edits before any further narrative update" is the explicit standard.
- `failure_family_hypothesis`: `workflow_orchestration`
- `intervention_kind_hypothesis`: `collapse_branches`
- `human_cost_or_risk`: Lost time and degraded trust due to progress narration without actual product movement.
- `local_lesson_hypothesis`: When the human is already signaling impatience with narrative drift, the next honest step is a bounded patch, not more exploratory framing.
- `cluster_hints`: `narrative-vs-disk-truth`, `ownership-discipline`, `explanatory-labor`
- `accepted_incident_likelihood`: `unlikely`
- `confidence`: `strong`
- `uncertainties`: A full-day reread could still reveal earlier adjacent messages that strengthen or weaken whether this should be grouped with another workflow-stall event.

### SC-03

- `event_id`: `SC-03`
- `title`: Human defines the first honest Jobs-lane slice instead of accepting more UI exploration
- `session_or_thread`: CodexDashboard Task-0004 supervision thread
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\04\rollout-2026-04-04T11-44-43-019d592b-0085-7ad2-8cfd-22009eae6dfa.jsonl`
- `primary_refs`: `L452-L506`, decisive intervention at `L484`
- `ai_course`: The assistant was still reading design references and the existing UI seam after entering `PASS-0001`, without yet landing the first visible Jobs-lane product slice.
- `human_intervention`: The human explicitly stopped broad exploration and specified the surface contract for the first slice: tab state, Jobs summary area, per-job rows with visible status/reason, bounded actions, and `Usage` remaining default.
- `adequate_outcome`: A patch that makes the Jobs lane visible, navigable, and reviewable enough to iterate on.
- `event_boundary_notes`: This is a UI/product-read event, not just workflow pacing. The human is directly defining the visible surface and navigation semantics.
- `human_model_signal`: A first UI pass should land the smallest honest human-facing slice, not reopen design research once the design direction is already approved.
- `failure_family_hypothesis`: `ui_semantics`
- `intervention_kind_hypothesis`: `collapse_branches`
- `human_cost_or_risk`: Extra delay plus risk of shipping a UI lane whose semantics remain implied rather than explicit.
- `local_lesson_hypothesis`: Once the design direction is sufficiently approved, a UI pass should converge into concrete visible behavior that a human can read and judge.
- `cluster_hints`: `navigation-semantics`, `first-honest-slice`, `product-read`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: None material in the bounded window.

### SC-04

- `event_id`: `SC-04`
- `title`: Reopened closure fails on the actual Jobs click path and on visible mockup fidelity
- `session_or_thread`: CodexDashboard main thread, reopened Task-0004 review
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\04\rollout-2026-04-04T13-23-58-019d5985-dbb1-77d0-a53d-debab8831378.jsonl`
- `primary_refs`: `L968-L1009`
- `ai_course`: The assistant had already reported a hotfix, pushed state, and described the task as repaired, but the human reopened the live dashboard and encountered missing current UI, then the actual `Jobs` click-path hitch, window spam, and tab-fidelity drift.
- `human_intervention`: The human rejected the result in human-world terms, asking whether it had been regression tested and explicitly reporting both the click-path failure and the mockup-fidelity miss.
- `adequate_outcome`: A running dashboard that shows the current Jobs lane, switches to it without hitching or spawning console windows, and visually reads closer to the mockup.
- `event_boundary_notes`: The stale-running-process explanation at `L971-L987` matters because it shows the first attempted explanation did not close the human's real complaint. The decisive failure boundary becomes clear at `L994`.
- `human_model_signal`: Real regression has to cover the click path a human will use. Visible mockup drift is also a real miss, not generic polish.
- `failure_family_hypothesis`: `verification_proof`
- `intervention_kind_hypothesis`: `reject_outcome`
- `human_cost_or_risk`: Visible breakage in the live app, noisy shell windows, hitching, and degraded trust in claimed closure.
- `local_lesson_hypothesis`: For newly added clickable UI, launch smoke and artifact capture are weaker than the actual click path and should never be allowed to stand in for it.
- `cluster_hints`: `click-path-regression`, `closure-truth`, `mockup-fidelity`, `real-app-surface`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: A wider reread could show whether the stale-process complaint at `L968` deserves its own smaller event, but it is not needed to support the stronger reopened-failure boundary.

### SC-05

- `event_id`: `SC-05`
- `title`: Human repeatedly teaches that new clickable UI needs named repo-local regression coverage
- `session_or_thread`: CodexDashboard main thread, regression-expectation follow-up
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\04\rollout-2026-04-04T13-23-58-019d5985-dbb1-77d0-a53d-debab8831378.jsonl`
- `primary_refs`: `L1051-L1138`
- `ai_course`: After the reopened failure, the assistant explained why the miss happened, but still needed several turns of human questioning to sharpen the actual rule and integrate it into durable docs.
- `human_intervention`: The human asked why the Jobs click path had not been tested, pressed on the disconnect between smoke and regression, proposed adding language around new regression cases for changed functionality, and then corrected the wording again so repo-local `REGRESSION.md` stayed the authoritative app-surface contract.
- `adequate_outcome`: A durable rule that changed human-facing functionality must map to named repo-local regression coverage, plus explicit repo updates for the Jobs interaction.
- `event_boundary_notes`: This is a separate explanatory-labor event, not just the same bug as `SC-04`, because the human is now teaching the underlying rule and making the assistant restate and rewrite it multiple times.
- `human_model_signal`: Regression must exercise the real app surface. New clickable UI should either map to an existing named repo-root regression case or add one. Repo-local `REGRESSION.md` is the operative surface contract.
- `failure_family_hypothesis`: `verification_proof`
- `intervention_kind_hypothesis`: `tighten_contract`
- `human_cost_or_risk`: Several turns of repeated teaching to recover a rule the model should have applied during closure, plus doc-edit overhead.
- `local_lesson_hypothesis`: When a human-facing interaction is newly introduced, the regression contract has to name that interaction explicitly enough that smoke proof cannot masquerade as closure.
- `cluster_hints`: `repeated-human-teaching`, `regression-scope`, `repo-local-contract`, `surface-vs-proxy-proof`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: The explanatory arc is clear, but a full-day read could still show whether an earlier April 4 thread already seeded this same rule before the human had to reteach it here.

### SC-06

- `event_id`: `SC-06`
- `title`: Human restates that tab clicks are navigation only and the tabs still read unlike the mockup
- `session_or_thread`: CodexDashboard main thread, post-repair follow-up
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\04\rollout-2026-04-04T13-23-58-019d5985-dbb1-77d0-a53d-debab8831378.jsonl`
- `primary_refs`: `L1145-L1223`
- `ai_course`: Even after the assistant reopened the task and acknowledged the regression miss, the tab click still triggered reconcile work and the tabs still did not meet the mockup fidelity bar.
- `human_intervention`: The human said "Stop doing the reconcile when i clicked on the tab," moved that behavior behind another button, reasserted the mockup-fidelity requirement, and told the assistant to use design reviewers if the intended read was still unclear.
- `adequate_outcome`: Clicking `Jobs` only changes the visible surface; heavier operations stay behind explicit controls; tab styling shifts toward the mockup's text-plus-underline read.
- `event_boundary_notes`: This is a repeated intervention after repair, which is why it stays separate from `SC-04`. The human is not just re-reporting the same bug; they are clarifying the intended navigation contract and product read.
- `human_model_signal`: Tabs are navigation semantics, not implicit side-effect triggers. Mockup fidelity is part of the intended outcome, and uncertainty should trigger design consultation rather than generic control defaults.
- `failure_family_hypothesis`: `ui_semantics`
- `intervention_kind_hypothesis`: `tighten_contract`
- `human_cost_or_risk`: Ongoing hitch, visible shell churn, and a surface that still reads wrong even after a repair attempt.
- `local_lesson_hypothesis`: Navigation controls should default to cheap surface switching unless the contract explicitly says the control performs heavier work.
- `cluster_hints`: `navigation-semantics`, `mockup-fidelity`, `button-vs-tab-read`, `repair-did-not-stick`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: The screenshot-review follow-up at `L1211-L1223` might warrant a wider reread if later acceptance work wants a larger verbatim fidelity window.

## Likely Accepted Incidents

- `SC-01`
- `SC-03`
- `SC-04`
- `SC-05`
- `SC-06`

## Likely Non-Incident But Still Important Intervention Events

- `SC-02` because it is a real course-correction and explanatory-burden event, but in this bounded read it is more valuable as support for drift diagnosis than as a clear standalone accepted incident.

## Repeated Cluster Hints Across The Analyzed Set

- `mockup-fidelity` or `product-read` is recurring, not incidental. It appears in both Crystallize task-writing and CodexDashboard live UI review.
- `surface-vs-proxy-proof` recurs: the assistant keeps drifting toward smoke, summaries, or implied state rather than the actual human-facing surface.
- `navigation-semantics` recurs: the human has to specify what a tab click means and what it should not do.
- `repeated-human-teaching` recurs: the human spends multiple turns re-explaining adequacy rules that should already have constrained closure behavior.

## Strongest Human-Model Signals Worth Carrying Forward

- A UI/task resolution is inadequate if it only says what disappears and still does not say what the human-facing safe state actually looks like.
- Regression means real app-surface exercise, not smoke-plus-confidence.
- New clickable UI needs named repo-local regression coverage rather than implied inheritance from an older smoke-shaped lane.
- Tabs are primarily navigation semantics. They should not silently do heavier work unless the contract clearly says so.
- Mockup fidelity is part of the outcome bar when the human says the surface should read like an approved mockup.

## Events That Still Need A Wider Reread

- `SC-04` through `SC-06` if later acceptance work wants to decide whether the late-day Task-0004 arc should become one accepted incident or multiple.
- `SC-01` only if later clustering wants a larger transcript window around how the assistant initially defined the weaker "falsifiable/evaluable" bar before the hard-mockup correction landed.

## PASS2 Spot-Check Judgment

The bounded reread confirms that April 4 contains multiple widened-rule events and a real pattern, especially around product-read, click-path proof, navigation semantics, and repeated human teaching.

The bounded reread recovered not just one extra event but a pattern:

- one clear Crystallize hard-mockup correction that upgrades product-read expectations
- one strong CodexDashboard UI-slice correction where the human had to define visible Jobs-lane semantics directly
- a multi-step late-day CodexDashboard arc where the human had to reject weak regression proof, reteach click-path testing expectations, and then restate tab navigation semantics and fidelity requirements after the repair still missed the mark

That is enough to say a raw-only read of April 4 would flag meaningful widened-rule material.

This PASS2 judgment is still transcript-first and does not yet compare against the existing April 4 pass artifacts.
