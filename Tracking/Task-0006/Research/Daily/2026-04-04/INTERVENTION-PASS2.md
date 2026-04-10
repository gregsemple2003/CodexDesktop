# INTERVENTION-PASS2

Source day: `2026-04-04`

## Source scope analyzed

- PASS1 artifact: `C:\Agent\CodexDashboard\Tracking\Task-0006\Research\Daily\2026-04-04\INTERVENTION-PASS1.md`
- Incident corpus contract reread:
  - `C:\Users\gregs\.codex\Orchestration\Reports\Incidents\README.md`
  - `C:\Users\gregs\.codex\Orchestration\Reports\Incidents\INCIDENT.schema.json`
- Raw transcripts reopened from the PASS1 refs:
  - `T04` = `C:\Users\gregs\.codex\sessions\2026\04\04\rollout-2026-04-04T01-22-17-019d56f1-261c-7e32-8a07-7323a20c471c.jsonl`
  - `T09` = `C:\Users\gregs\.codex\sessions\2026\04\04\rollout-2026-04-04T11-15-56-019d5910-a6c5-7373-bd7e-46d7a2a93cb9.jsonl`
  - `T14` = `C:\Users\gregs\.codex\sessions\2026\04\04\rollout-2026-04-04T11-18-43-019d5913-3060-70f1-a7d7-44a771efaef8.jsonl`
  - `T18` = `C:\Users\gregs\.codex\sessions\2026\04\04\rollout-2026-04-04T11-43-16-019d5929-ab66-7c21-9f8a-9dc6024325ab.jsonl`
  - `T19` = `C:\Users\gregs\.codex\sessions\2026\04\04\rollout-2026-04-04T11-44-43-019d592b-0085-7ad2-8cfd-22009eae6dfa.jsonl`
- Direct task-artifact checks pulled in only where the transcript hinged on expected-state wording:
  - `C:\Agent\ThirdPerson\Tracking\Task-0002\TASK.md`
  - `C:\Agent\CodexDashboard\REGRESSION.md`
  - `C:\Agent\CodexDashboard\Tracking\Task-0004\Design\STITCH-JOBS-TAB-0001\APPROVED-NOTES.md`
  - `C:\Users\gregs\.codex\Orchestration\Jobs\README.md`

## Candidate ids analyzed

- `C01` through `C19`

## Boundary corrections relative to PASS1

- No hard split or merge corrections were required from the PASS1 candidate set.
- `C03` remains separate from `C02`. The transcript supports a stricter second intervention: "hard-mockup-ready" is a stronger local bar than merely "falsifiable or evaluable."
- `C08` stays intentionally wide. It is one recurring anti-drift correction arc across `PASS-0000` and `PASS-0001`: the same "stop broad exploration and land the smallest honest patch now" rule had to be reasserted on the next pass.
- `C14` also stays intentionally wide. The initial rejection of reference-pose closure and the later reopen/continue-until-done restatements are the same closure-truth event recurring against renewed drift, not unrelated incidents.

## Per-event analysis records

### C01 - Reject open-ended narrow usability task definitions

- `event_id`: `C01`
- `title`: `Reject open-ended narrow usability task definitions`
- `session_or_thread`: `Harvest task context`
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\04\rollout-2026-04-04T01-22-17-019d56f1-261c-7e32-8a07-7323a20c471c.jsonl`
- `primary_refs`: `T04:L149`, `T04:L159`, `T04:L169`
- `ai_course`: The assistant had written narrow usability follow-up tasks as loose investigation-shaped artifacts with unresolved product branches, such as whether onboarding remained on the People surface.
- `human_intervention`: The human said the tasks were too open-ended, clarified that bug reports should guide direction without becoming the full design spec, and said the producer should drive consensus with interface/general design instead of leaving the product in a "superposition of states."
- `adequate_outcome`: Point-fix usability tasks should record one reviewable expected resolution so a human can agree, disagree, or redirect before implementation starts.
- `event_boundary_notes`: PASS1 boundary is accurate. `C01` is about consensus responsibility and rejecting the loose task shape. `C02` is the later wording refinement.
- `human_model_signal`: Explicit producer rule: if the task leaves the product in a "superposition of states," the producer has failed because the human cannot tell whether course correction is needed.
- `failure_family_hypothesis`: `workflow_orchestration`
- `intervention_kind_hypothesis`: `collapse_branches`
- `human_cost_or_risk`: The human cannot tell whether they agree with the task, so product direction stays unresolved until they step in.
- `local_lesson_hypothesis`: For narrow usability tasks, the producer must collapse likely product branches into one reviewable human-facing resolution before implementation begins.
- `cluster_hints`: `task-definition`, `product-superposition`, `consensus-building`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: none material

### C02 - Replace "declarative terms" with falsifiable or evaluable claims

- `event_id`: `C02`
- `title`: `Replace "declarative terms" with falsifiable or evaluable claims`
- `session_or_thread`: `Harvest task context`
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\04\rollout-2026-04-04T01-22-17-019d56f1-261c-7e32-8a07-7323a20c471c.jsonl`
- `primary_refs`: `T04:L179`, `T04:L220`, `T04:L230`
- `ai_course`: After accepting the need to collapse branches, the assistant encoded the durable rule in weaker "declarative terms" language, which still tolerated vague task definitions.
- `human_intervention`: The human explicitly rejected that phrasing and said usability tasks should "collapse the waveform to a falsifiable or evaluat-able claim."
- `adequate_outcome`: Task definitions must collapse into one or more falsifiable or evaluable claims about the expected user-facing resolution.
- `event_boundary_notes`: PASS1 boundary is accurate. This is a second intervention because the assistant's first durable rewrite still weakened the bar.
- `human_model_signal`: Explicit wording signal: "falsifiable or evaluat-able claim" is the right producer standard, and "declarative terms" is too weak.
- `failure_family_hypothesis`: `workflow_orchestration`
- `intervention_kind_hypothesis`: `tighten_contract`
- `human_cost_or_risk`: A task can be declarative and still vague, which keeps review mushy and lets weaker work slip through.
- `local_lesson_hypothesis`: When the human needs reviewable task language, use falsifiable or evaluable claims rather than softer labels like "declarative."
- `cluster_hints`: `task-definition`, `falsifiability`, `reviewability`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: none material

### C03 - Make usability resolutions hard-mockup-ready

- `event_id`: `C03`
- `title`: `Make usability resolutions hard-mockup-ready`
- `session_or_thread`: `Harvest task context`
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\04\rollout-2026-04-04T01-22-17-019d56f1-261c-7e32-8a07-7323a20c471c.jsonl`
- `primary_refs`: `T04:L452`, `T04:L539`
- `ai_course`: Even after the falsifiable/evaluable rewrite, `Task-0022` still used a negative-only claim about removing the incomplete progress bar while leaving the replacement safe-state UI underspecified.
- `human_intervention`: The human said the sentence still failed because the resolution was not clear enough to produce a hard mockup, then required prompt iteration plus fresh subagent retesting.
- `adequate_outcome`: Task language must state the positive replacement UI clearly enough to draw, review, and reproduce with cold agents, not merely say what disappears.
- `event_boundary_notes`: PASS1 boundary is accurate. `T04:L539` is the carry-forward confirmation that the stronger mockup-ready rule should be applied to the remaining tasks and to the task-harvester itself.
- `human_model_signal`: Explicit rule: the resolution should be "clear enough to produce a hard mockup," and saying only what the UI will not be is not sufficient.
- `failure_family_hypothesis`: `workflow_orchestration`
- `intervention_kind_hypothesis`: `tighten_contract`
- `human_cost_or_risk`: Even a "falsifiable" task can stay too vague for human review, forcing another correction pass.
- `local_lesson_hypothesis`: For narrow UI tasks, do not stop at negative or minimal claims; state the replacement surface concretely enough that a reviewer could sketch it.
- `cluster_hints`: `task-definition`, `mockup-ready`, `falsifiability`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: none material

### C04 - Reject reboot-robustness overclaim on the server lane

- `event_id`: `C04`
- `title`: `Reject reboot-robustness overclaim on the server lane`
- `session_or_thread`: `Harvest task context`
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\04\rollout-2026-04-04T01-22-17-019d56f1-261c-7e32-8a07-7323a20c471c.jsonl`
- `primary_refs`: `T04:L843`, `T04:L951`, `T04:L1008`
- `ai_course`: The assistant had to explain why the always-on server was not running after a machine restart, and its initial next-step framing still treated the gap partly as a follow-up process/product question rather than first as an overclaimed readiness defect.
- `human_intervention`: The human restated the original bar as restart robustness, then explicitly required a bug on the task that overclaimed readiness and said the cause should be logged as "not adequate or honest test coverage" before another restart test.
- `adequate_outcome`: Reclassify the issue as durable proof failure, file it as a bug, reopen task truth on disk, and do readiness due diligence before the next restart retest.
- `event_boundary_notes`: PASS1 boundary is accurate. `T04:L1008` is aftermath showing the restart retest remains gated on honest readiness, not a separate incident.
- `human_model_signal`: Explicit adequacy rule: overclaiming readiness should be recorded as "not adequate or honest test coverage," not softened into a generic follow-up note.
- `failure_family_hypothesis`: `verification_proof`
- `intervention_kind_hypothesis`: `reject_outcome`
- `human_cost_or_risk`: False confidence about restart robustness and risk of rebooting into another false positive.
- `local_lesson_hypothesis`: Do not claim a durability lane is ready until the exact real-world persistence behavior, here reboot survival, has been exercised honestly.
- `cluster_hints`: `closure-truth`, `proof-scope`, `durable-bug-truth`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: none material

### C05 - Add screenshot-backed proof as a required Task-0002 bar

- `event_id`: `C05`
- `title`: `Add screenshot-backed proof as a required Task-0002 bar`
- `session_or_thread`: `Implement task 2 task / Lead Task-0002 workflow`
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\04\rollout-2026-04-04T11-15-56-019d5910-a6c5-7373-bd7e-46d7a2a93cb9.jsonl`
- `primary_refs`: `T09:L221`, `T14:L268`
- `ai_course`: Task-0002's plan-approval summary emphasized structured presentation-state evidence for the live pawn path and had not yet made screenshot proof a hard requirement.
- `human_intervention`: The human approved the plan only with the caveat that screenshot functionality must prove the pawn's existence and that structured runtime state alone is not enough.
- `adequate_outcome`: PASS-0001 must preserve screenshot-based proof of the live pawn in addition to machine-legible presentation-state evidence.
- `event_boundary_notes`: PASS1 boundary is accurate. This is a small preemptive proof-bar tightening and remains distinct from the later wrong-seam debugging redirect in `C07`.
- `human_model_signal`: Explicit rule: "Structured runtime state alone is not enough."
- `failure_family_hypothesis`: `verification_proof`
- `intervention_kind_hypothesis`: `tighten_contract`
- `human_cost_or_risk`: A visibility task could close on machine-legible state while still lacking human-visible proof.
- `local_lesson_hypothesis`: When the task is about visible presence, pair runtime state with live screenshot evidence before claiming proof.
- `cluster_hints`: `proof-scope`, `screenshot-proof`, `human-visible-bar`
- `accepted_incident_likelihood`: `possible`
- `confidence`: `strong`
- `uncertainties`: This may end up as a proof-bar caveat rather than a standalone accepted incident.

### C06 - Force the first durable Task-0004 planning checkpoint

- `event_id`: `C06`
- `title`: `Force the first durable Task-0004 planning checkpoint`
- `session_or_thread`: `Update Task-0004 state`
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\04\rollout-2026-04-04T11-44-43-019d592b-0085-7ad2-8cfd-22009eae6dfa.jsonl`
- `primary_refs`: `T19:L160`
- `ai_course`: The delegated Task-0004 leader had kept broad re-grounding and still had not produced the first durable planning checkpoint.
- `human_intervention`: The human ordered immediate `TASK-STATE.json` and `PLAN.md` creation with the honest current state and an explicit plan-approval gate.
- `adequate_outcome`: Durable planning state on disk before any more exploration or status narration.
- `event_boundary_notes`: PASS1 boundary is accurate.
- `human_model_signal`: Explicit durable-state rule: if the work is already planning-ready, that phase truth should exist on disk before more narration.
- `failure_family_hypothesis`: `workflow_orchestration`
- `intervention_kind_hypothesis`: `tighten_contract`
- `human_cost_or_risk`: Task ownership and lifecycle truth remain implicit, making review and handoff weaker than they should be.
- `local_lesson_hypothesis`: When research is already planning-ready, checkpoint planning state first instead of continuing broad re-grounding.
- `cluster_hints`: `durable-state-first`, `planning-gate`, `analysis-drift`
- `accepted_incident_likelihood`: `possible`
- `confidence`: `strong`
- `uncertainties`: none material

### C07 - Bound proof-view debugging to one discriminating check

- `event_id`: `C07`
- `title`: `Bound proof-view debugging to one discriminating check`
- `session_or_thread`: `Implement task 2 task / Lead Task-0002 workflow`
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\04\rollout-2026-04-04T11-15-56-019d5910-a6c5-7373-bd7e-46d7a2a93cb9.jsonl` and `C:\Users\gregs\.codex\sessions\2026\04\04\rollout-2026-04-04T11-18-43-019d5913-3060-70f1-a7d7-44a771efaef8.jsonl`
- `primary_refs`: `T09:L543`, `T09:L580`, `T14:L607`, `T14:L751`
- `ai_course`: The assistant was explaining an automation-only proof view and risked expanding capture helpers or mutating the live gameplay camera path to satisfy screenshot proof.
- `human_intervention`: The human said the default third-person camera was probably fine and the mesh itself was likely the problem, then later bounded proof view to one automation-only discriminating check with the real gameplay camera left untouched.
- `adequate_outcome`: Run one bounded proof-view capture only; if it also fails, treat that as strong evidence of a real visibility/render defect and stop expanding automation helpers.
- `event_boundary_notes`: PASS1 boundary is accurate. `T09` captures the human's skepticism, and `T14` converts it into a durable PASS-0001 decision rule.
- `human_model_signal`: Two explicit rules: the live gameplay camera remains the truth surface, and if a proof-only angle also fails then the branch should move into the real defect rather than more tooling.
- `failure_family_hypothesis`: `verification_proof`
- `intervention_kind_hypothesis`: `redirect_debugging`
- `human_cost_or_risk`: Debugging could drift into proof-surface expansion rather than the real mesh visibility seam.
- `local_lesson_hypothesis`: Use proof-only viewpoints only as a bounded discriminator; if they fail too, redirect immediately into the product defect instead of extending the proof surface.
- `cluster_hints`: `wrong-seam-debugging`, `proof-scope`, `live-camera-truth`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: none material

### C08 - Stop Task-0004 exploration drift and land concrete patches

- `event_id`: `C08`
- `title`: `Stop Task-0004 exploration drift and land concrete patches`
- `session_or_thread`: `Update Task-0004 state`
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\04\rollout-2026-04-04T11-44-43-019d592b-0085-7ad2-8cfd-22009eae6dfa.jsonl`
- `primary_refs`: `T19:L306`, `T19:L484`
- `ai_course`: After the first durable checkpoint, the work kept slipping back into read-side grounding and broad exploration instead of landing concrete product-file edits. First the backend slice drifted, then the first Jobs UI slice drifted.
- `human_intervention`: At `T19:L306` the human forced a minimal concrete `PASS-0000` backend patch. At `T19:L484` they had to intervene again and force the smallest honest `PASS-0001` Jobs-lane UI patch instead of more exploration.
- `adequate_outcome`: Concrete backend code first, then concrete additive Jobs UI, with no more broad exploration in place of the approved smallest honest slices.
- `event_boundary_notes`: PASS1 keeps this as one multi-stage anti-drift event, and that still fits the transcript. The same "patch the product now" correction had to be repeated on the next pass.
- `human_model_signal`: Explicit boundedness rule: "Do not do more broad exploration" and "patch the product."
- `failure_family_hypothesis`: `workflow_orchestration`
- `intervention_kind_hypothesis`: `collapse_branches`
- `human_cost_or_risk`: Time lost in read/status churn with no product-file movement, plus risk that the task stalls before any real slice is proven.
- `local_lesson_hypothesis`: Once a pass has an approved smallest honest slice, land that code before reopening analysis or design space.
- `cluster_hints`: `analysis-drift`, `smallest-honest-slice`, `durable-state-first`
- `accepted_incident_likelihood`: `possible`
- `confidence`: `medium`
- `uncertainties`: A later clustering pass could split the backend and UI anti-drift moments if it wants pass-by-pass granularity.

### C09 - Interrupt the live Task-0004 hotkey regression immediately

- `event_id`: `C09`
- `title`: `Interrupt the live Task-0004 hotkey regression immediately`
- `session_or_thread`: `Implement task 4 task`
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\04\rollout-2026-04-04T11-43-16-019d5929-ab66-7c21-9f8a-9dc6024325ab.jsonl`
- `primary_refs`: `T18:L794`
- `ai_course`: After Jobs-lane work, the live dashboard no longer closed on `Ctrl Alt Space` while closure work was still in progress.
- `human_intervention`: The human flagged it as a blocker because the broken overlay was still in their face and demanded an immediate fix.
- `adequate_outcome`: Restore baseline live-overlay close behavior before treating the pass as healthy again.
- `event_boundary_notes`: PASS1 boundary is accurate.
- `human_model_signal`: Explicit human-world priority rule: a live regression obstructing the current user workflow becomes the blocker immediately.
- `failure_family_hypothesis`: `human_world`
- `intervention_kind_hypothesis`: `reject_outcome`
- `human_cost_or_risk`: Immediate disruption on a surface the human is actively using.
- `local_lesson_hypothesis`: When a task breaks a baseline control the user is actively relying on, stop and fix that regression before continuing closeout work.
- `cluster_hints`: `live-regression`, `blocking-surface`, `user-burden`
- `accepted_incident_likelihood`: `possible`
- `confidence`: `strong`
- `uncertainties`: This may be too local for the accepted set unless later clustering wants "blocking live regression" exemplars.

### C10 - Narrow the ZeroMale defect boundary to the animation path

- `event_id`: `C10`
- `title`: `Narrow the ZeroMale defect boundary to the animation path`
- `session_or_thread`: `Lead Task-0002 workflow`
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\04\rollout-2026-04-04T11-18-43-019d5913-3060-70f1-a7d7-44a771efaef8.jsonl`
- `primary_refs`: `T14:L1373`
- `ai_course`: After the no-animation/reference-pose run, the branch could still have stayed broad around generic visibility or render uncertainty.
- `human_intervention`: The human said this run was the first decisive success and that the defect seam should now be treated as `/Game/Experimental/ZeroMale/Animations/ZM_ABP_Unarmed`, not generic mesh visibility.
- `adequate_outcome`: Collapse the live defect boundary to the animation path and proceed from that narrower seam.
- `event_boundary_notes`: PASS1 boundary is accurate. This remains distinct from `C14`, which later rejects treating the fallback as closure.
- `human_model_signal`: Explicit debugging rule: once a discriminating run isolates the seam, the main uncertainty is no longer broad evidence collection.
- `failure_family_hypothesis`: `verification_proof`
- `intervention_kind_hypothesis`: `redirect_debugging`
- `human_cost_or_risk`: Continued broad render diagnostics after the seam was already narrowed.
- `local_lesson_hypothesis`: Once a discriminating run isolates the failing component, rewrite the defect boundary and stop broadening the search.
- `cluster_hints`: `wrong-seam-debugging`, `defect-boundary`, `discriminating-test`
- `accepted_incident_likelihood`: `possible`
- `confidence`: `strong`
- `uncertainties`: none material

### C11 - Reject the stale-process stop point because the live surface still lacked `Jobs`

- `event_id`: `C11`
- `title`: `Reject the stale-process stop point because the live surface still lacked Jobs`
- `session_or_thread`: `Implement task 4 task`
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\04\rollout-2026-04-04T11-43-16-019d5929-ab66-7c21-9f8a-9dc6024325ab.jsonl`
- `primary_refs`: `T18:L967`
- `ai_course`: After fixing the hotkey close regression, the assistant had effectively stopped while the running dashboard instance still showed old code with no `Jobs` tab.
- `human_intervention`: The human pointed out that `Ctrl Alt Space` still showed no `Jobs` tab and asked why the task had stopped.
- `adequate_outcome`: Restart the actual live dashboard process and verify the running surface, not just the repo state or current build artifacts.
- `event_boundary_notes`: PASS1 boundary is accurate. This is a stale-process/live-surface truth correction, separate from the deeper regression miss in `C12`.
- `human_model_signal`: Implicit but clear rule: the running human-facing surface is authoritative; stale code in memory does not count as delivery.
- `failure_family_hypothesis`: `verification_proof`
- `intervention_kind_hypothesis`: `reject_outcome`
- `human_cost_or_risk`: Misleading stop point while the user still sees old behavior.
- `local_lesson_hypothesis`: After UI changes, verify the actual running process shows them before pausing or closing a pass.
- `cluster_hints`: `live-surface-truth`, `stale-process`, `closure-truth`
- `accepted_incident_likelihood`: `unlikely`
- `confidence`: `strong`
- `uncertainties`: Smaller and more local than the surrounding regression-doctrine events.

### C12 - Reject Task-0004 regression closure after real `Jobs`-tab failures

- `event_id`: `C12`
- `title`: `Reject Task-0004 regression closure after real Jobs-tab failures`
- `session_or_thread`: `Implement task 4 task`
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\04\rollout-2026-04-04T11-43-16-019d5929-ab66-7c21-9f8a-9dc6024325ab.jsonl`
- `primary_refs`: `T18:L993`
- `ai_course`: The assistant had treated smoke/startup proof and artifact capture as enough to re-close Task-0004 without deeply exercising the real `Jobs` click path.
- `human_intervention`: The human reported live failures when clicking `Jobs`: hitching, a crazy number of windows, and mockup fidelity drift, effectively rejecting the closure.
- `adequate_outcome`: Reopen Task-0004, reproduce the actual click path in the live app, and fix both behavior and fidelity before claiming closure again.
- `event_boundary_notes`: PASS1 boundary is accurate and should stay split from `C13`. `C12` is the live evidence rejection; `C13` is the later durable regression rule tightened from that miss.
- `human_model_signal`: Explicit user-evidence rule that the real clicked surface is the closure bar for a new clickable feature.
- `failure_family_hypothesis`: `verification_proof`
- `intervention_kind_hypothesis`: `reject_outcome`
- `human_cost_or_risk`: A shipped clickable surface that behaves badly despite a "done" story.
- `local_lesson_hypothesis`: Do not let smoke/startup proof stand in for the actual interacted click path when the new risk lives in that interaction.
- `cluster_hints`: `real-click-path`, `closure-truth`, `mockup-fidelity`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: none material

### C13 - Tighten regression doctrine around repo-local click paths

- `event_id`: `C13`
- `title`: `Tighten regression doctrine around repo-local click paths`
- `session_or_thread`: `Implement task 4 task`
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\04\rollout-2026-04-04T11-43-16-019d5929-ab66-7c21-9f8a-9dc6024325ab.jsonl`
- `primary_refs`: `T18:L1050`, `T18:L1063`, `T18:L1104`
- `ai_course`: Even after admitting the click-path miss, the assistant's model still leaned on the existing smoke-shaped lane instead of the repo-local real-surface requirement. The reread of `C:\Agent\CodexDashboard\REGRESSION.md` confirms the repo already required real app-surface cases for changed functionality.
- `human_intervention`: The human paused implementation to focus on the expectation disconnect, required the explanation to anchor on repo-local `REGRESSION.md`, and explicitly rejected any wording suggesting smoke could ever count as regression.
- `adequate_outcome`: Durable doctrine that regression in this repo means real app-surface cases, and new clickable surfaces must be covered by named repo-local cases.
- `event_boundary_notes`: PASS1 boundary is accurate.
- `human_model_signal`: Explicit rule that repo-local `REGRESSION.md` is authoritative and regression, by definition, is app-surface interaction rather than smoke.
- `failure_family_hypothesis`: `workflow_orchestration`
- `intervention_kind_hypothesis`: `tighten_contract`
- `human_cost_or_risk`: Future tasks would keep closing on the wrong proof lane.
- `local_lesson_hypothesis`: When a repo adds or changes a clickable human-facing surface, map closure to the repo's actual regression cases instead of inheriting a smoke-shaped habit.
- `cluster_hints`: `regression-doctrine`, `real-click-path`, `repo-local-authority`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: none material

### C14 - Reject reference-pose closure and reopen Task-0002 under the real human bar

- `event_id`: `C14`
- `title`: `Reject reference-pose closure and reopen Task-0002 under the real human bar`
- `session_or_thread`: `Implement task 2 task / Lead Task-0002 workflow`
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\04\rollout-2026-04-04T11-15-56-019d5910-a6c5-7373-bd7e-46d7a2a93cb9.jsonl` and `C:\Users\gregs\.codex\sessions\2026\04\04\rollout-2026-04-04T11-18-43-019d5913-3060-70f1-a7d7-44a771efaef8.jsonl`
- `primary_refs`: `T09:L1239`, `T09:L1259`, `T14:L1470`, `T14:L1567`, `T14:L2415`, `T14:L3475`
- `ai_course`: Task-0002 kept drifting toward closure on proxy outcomes such as "mesh is provably on screen," "the seam is isolated," reference pose, or blocked-pass evidence, even though the tightened task north star in `C:\Agent\ThirdPerson\Tracking\Task-0002\TASK.md` requires a visibly animated and controllable pawn in real PIE.
- `human_intervention`: The human said closure on a T-pose/reference-pose fallback would be a "dreadful narrowing of scope," reopened the task from disk, invalidated the stale plan, repeatedly re-approved continuation only under the tightened north star, and finally restated that blocked checkpoints and degraded fallbacks do not count as completion.
- `adequate_outcome`: Keep Task-0002 open until the promoted ZeroMale pawn is visibly animated and controllable in real PIE, with screenshot-backed proof on the live path.
- `event_boundary_notes`: PASS1 keeps a deliberately wide boundary, and the transcript supports that. The later reopen and continue-until-done directives are the same closure-truth correction recurring after renewed drift.
- `human_model_signal`: Several explicit human-world signals matter here: "closing out a task designed to fix visibility when the resolution is 'just leave it in tpose' would clearly be a dreadful narrowing of scope," and the human-facing bar is a visibly animated, controllable embodied avatar on the real playable path.
- `failure_family_hypothesis`: `human_world`
- `intervention_kind_hypothesis`: `reject_outcome`
- `human_cost_or_risk`: A visibly broken playable character could be normalized as "done" through proxy proof or degraded fallback.
- `local_lesson_hypothesis`: For human-facing visibility tasks, interpret completion as the real playable experience, not a proxy proof or fallback pose.
- `cluster_hints`: `closure-truth`, `human-visible-bar`, `proxy-vs-experience`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: A later pass could split the initial closure rejection from the later standing-restatement chain, but the local lesson is already clear without doing so.

### C15 - Stop reconcile on tab entry and force fidelity repair on the live surface

- `event_id`: `C15`
- `title`: `Stop reconcile on tab entry and force fidelity repair on the live surface`
- `session_or_thread`: `Implement task 4 task`
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\04\rollout-2026-04-04T11-43-16-019d5929-ab66-7c21-9f8a-9dc6024325ab.jsonl`
- `primary_refs`: `T18:L1144`, `T18:L1210`
- `ai_course`: Even after the regression-doctrine fix, the live `Jobs` tab still hitched because tab entry triggered reconcile work, and the surface still read visually "off" relative to the approved direction.
- `human_intervention`: The human ordered reconcile removed from the tab click path, required fidelity repair against the mockup, and asked for interface-designer review of the screenshot.
- `adequate_outcome`: Tab click should only switch surfaces; refresh/reconcile should be explicit actions; fidelity should be checked against the approved direction with second-opinion review when needed.
- `event_boundary_notes`: PASS1 boundary is accurate. `T18:L1210` belongs to the same event because it requests the fidelity review needed to resolve the "off" live surface.
- `human_model_signal`: Explicit interaction rule that navigation should not trigger heavy work, plus explicit design rule that if the surface reads "off," design review is part of the fix rather than optional polish.
- `failure_family_hypothesis`: `information_architecture`
- `intervention_kind_hypothesis`: `reject_outcome`
- `human_cost_or_risk`: Hitching, unwanted work on navigation click, and lower trust in the Jobs surface.
- `local_lesson_hypothesis`: Keep navigation side-effect free and treat mockup fidelity as part of the real fix bar, not as post-hoc polish.
- `cluster_hints`: `navigation-vs-action`, `mockup-fidelity`, `real-click-path`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: none material

### C16 - Force TASK-HARVESTER prime directives for concrete measurable tasks

- `event_id`: `C16`
- `title`: `Force TASK-HARVESTER prime directives for concrete measurable tasks`
- `session_or_thread`: `Implement task 2 task`
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\04\rollout-2026-04-04T11-15-56-019d5910-a6c5-7373-bd7e-46d7a2a93cb9.jsonl`
- `primary_refs`: `T09:L1363`, `T09:L1403`, `T09:L1469`, `T09:L1540`, `T09:L1590`
- `ai_course`: Even after Task-0002's closure miss was acknowledged, the task-writing repair work still left the highest-priority producer rule too implicit, too broad, or buried under lower-priority prompt details.
- `human_intervention`: The human repeatedly pushed for clearer human-aware task language, requested research on best practices for human-readable task writeups, then insisted on a prime directive that a good task must be specific and measurable enough to mark done, and finally ordered Task-0002 updated as the new north star under that rule.
- `adequate_outcome`: Make task specificity/measurability a prime-directive-level rule that trumps weaker or conflicting prompt details, then prove it by rewriting a live task under that standard.
- `event_boundary_notes`: PASS1 boundary is accurate. This is one escalating prompt-hardening event from principle to research to prime-directive placement to concrete task rewrite.
- `human_model_signal`: Explicit prime directive: "a good task should be specific and measurable enough to mark done, it doesn't leave the outcome up for guessing."
- `failure_family_hypothesis`: `workflow_orchestration`
- `intervention_kind_hypothesis`: `tighten_contract`
- `human_cost_or_risk`: Broad task language lets agents "smuggle in inferior work" and makes closure guessable instead of reviewable.
- `local_lesson_hypothesis`: Put task specificity and measurability at the highest prompt priority, then validate it on a live task instead of leaving it as a lower-level heuristic.
- `cluster_hints`: `task-definition`, `prime-directives`, `human-aware-language`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: none material

### C17 - Separate primary navigation from tab content structurally

- `event_id`: `C17`
- `title`: `Separate primary navigation from tab content structurally`
- `session_or_thread`: `Implement task 4 task`
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\04\rollout-2026-04-04T11-43-16-019d5929-ab66-7c21-9f8a-9dc6024325ab.jsonl`
- `primary_refs`: `T18:L1243`, `T18:L1253`, `T18:L1263`
- `ai_course`: The Task-0004 shell still mixed primary navigation and tab-specific controls/content in one strip instead of separating "where am I?" from "what belongs to this tab?"
- `human_intervention`: The human specified that primary nav must be its own strip and that the area underneath must hold the tab content, while preserving the redline treatment from the original mockup.
- `adequate_outcome`: Structural separation between primary navigation and tab-owned content/actions, then one coherent implementation pass that preserves the approved exception.
- `event_boundary_notes`: PASS1 boundary is accurate. `T18:L1253` and `T18:L1263` remain inside the same structural correction arc because the human narrows one exception and then authorizes implementation of the combined layout pass.
- `human_model_signal`: Explicit hierarchy rule: primary navigation is its own strip; content belongs below.
- `failure_family_hypothesis`: `information_architecture`
- `intervention_kind_hypothesis`: `tighten_contract`
- `human_cost_or_risk`: Muddled hierarchy and a less legible overlay surface.
- `local_lesson_hypothesis`: Keep primary navigation structurally separate from tab-owned controls so hierarchy is legible at a glance.
- `cluster_hints`: `navigation-vs-content`, `visual-hierarchy`, `surface-structure`
- `accepted_incident_likelihood`: `possible`
- `confidence`: `strong`
- `uncertainties`: This may later cluster as design-review feedback rather than an accepted incident unless similar hierarchy misses recur elsewhere.

### C18 - Clarify declared jobs JSON plus schema as the primary durable jobs state

- `event_id`: `C18`
- `title`: `Clarify declared jobs JSON plus schema as the primary durable jobs state`
- `session_or_thread`: `Implement task 4 task`
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\04\rollout-2026-04-04T11-43-16-019d5929-ab66-7c21-9f8a-9dc6024325ab.jsonl`
- `primary_refs`: `T18:L1716`
- `ai_course`: The Jobs lane still read as though actual runtime state or immediate reconcile output might be the primary truth on tab entry, and the details surface was drifting toward actual-job plumbing.
- `human_intervention`: The human clarified that the primary durable state is the declared jobs JSON plus schema under `.codex/Orchestration/Jobs`; reconcile is actual-vs-intended diff "similar to terraform," and `Details` should show the declared job rather than the actual runtime object. The reread of `C:\Users\gregs\.codex\Orchestration\Jobs\README.md` supports that declared-vs-runtime split.
- `adequate_outcome`: Make declared jobs the first-class Jobs surface on click, with reconcile layered on top as a diff view rather than as the product's primary truth.
- `event_boundary_notes`: PASS1 boundary is accurate.
- `human_model_signal`: Explicit source-of-truth rule: declared jobs are the "intent of codex wrt jobs," and reconcile is secondary comparison rather than primary state.
- `failure_family_hypothesis`: `usability_state_truth`
- `intervention_kind_hypothesis`: `reframe_problem`
- `human_cost_or_risk`: The Jobs surface would misstate what the system is supposed to be doing and make the wrong layer feel authoritative.
- `local_lesson_hypothesis`: For desired-vs-runtime control surfaces, anchor the default view on durable declared intent and treat reconcile as a comparison layer instead of the primary product state.
- `cluster_hints`: `status-signal-meaning`, `declared-vs-actual`, `source-of-truth`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: none material

### C19 - Require a scrollable Jobs panel

- `event_id`: `C19`
- `title`: `Require a scrollable Jobs panel`
- `session_or_thread`: `Implement task 4 task`
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\04\rollout-2026-04-04T11-43-16-019d5929-ab66-7c21-9f8a-9dc6024325ab.jsonl`
- `primary_refs`: `T18:L1855`
- `ai_course`: The `Jobs` surface shipped without its own scrollable panel even though the content can exceed the viewport.
- `human_intervention`: The human explicitly called out that the panel must be scrollable.
- `adequate_outcome`: Make the `Jobs` panel itself scroll so all intended rows and details remain reachable.
- `event_boundary_notes`: PASS1 boundary is accurate.
- `human_model_signal`: `none explicit`
- `failure_family_hypothesis`: `information_architecture`
- `intervention_kind_hypothesis`: `reject_outcome`
- `human_cost_or_risk`: Inaccessible content on a dense operational surface.
- `local_lesson_hypothesis`: If a panel can exceed its viewport, scrolling is part of basic usability rather than optional polish.
- `cluster_hints`: `surface-ergonomics`, `viewport-fit`
- `accepted_incident_likelihood`: `unlikely`
- `confidence`: `strong`
- `uncertainties`: This is closer to a straightforward live usability fix than a durable accepted-incident candidate.

## Likely accepted incidents

- `C01` - Reject open-ended narrow usability task definitions
- `C02` - Replace "declarative terms" with falsifiable or evaluable claims
- `C03` - Make usability resolutions hard-mockup-ready
- `C04` - Reject reboot-robustness overclaim on the server lane
- `C07` - Bound proof-view debugging to one discriminating check
- `C12` - Reject Task-0004 regression closure after real Jobs-tab failures
- `C13` - Tighten regression doctrine around repo-local click paths
- `C14` - Reject reference-pose closure and reopen Task-0002 under the real human bar
- `C15` - Stop reconcile on tab entry and force fidelity repair on the live surface
- `C16` - Force TASK-HARVESTER prime directives for concrete measurable tasks
- `C18` - Clarify declared jobs JSON plus schema as the primary durable jobs state

## Likely non-incident but still important intervention events

- `C05` - Real proof-bar tightening, but closer to a plan-approval caveat than to a standalone accepted incident.
- `C06` - Clear durable-state correction, but probably too close to ordinary task supervision for the accepted set.
- `C08` - Authentic anti-drift intervention with two stages, but likely better treated as a workflow-smell example than as a top-tier accepted incident.
- `C09` - Real live regression interrupt, but probably too local unless later clustering wants blocker-surface exemplars.
- `C10` - Strong seam-narrowing redirect, but more of a debugging-boundary exemplar than an accepted incident on its own.
- `C11` - Useful live-surface truth correction, but smaller than the surrounding click-path and regression-doctrine incidents.
- `C17` - Real hierarchy correction, but it may later read more as design-review feedback than as an accepted incident.
- `C19` - Direct live usability fix, but likely too small and straightforward for the accepted set.

## Repeated cluster hints noticed across the analyzed set

- `task-definition`
  - `C01`, `C02`, `C03`, `C16`
- `closure-truth`
  - `C04`, `C11`, `C12`, `C14`
- `proof-scope`
  - `C05`, `C07`, `C12`, `C14`
- `wrong-seam-debugging`
  - `C07`, `C10`
- `real-click-path`
  - `C12`, `C13`, `C15`
- `navigation-vs-content`
  - `C15`, `C17`, `C19`
- `declared-vs-actual`
  - `C18`

## Strongest human-model signals worth carrying into later clustering or principle work

- `C01`, `C02`, `C03`, and `C16`: For narrow usability tasks, the producer must drive consensus and write falsifiable/evaluable, mockup-ready claims. Broad language leaves the product in superposition and lets weaker work "smuggle in."
- `C04`: If readiness was overclaimed, record that as inadequate or dishonest test coverage rather than softening it into a vague follow-up story.
- `C05`, `C07`, and `C14`: Human-facing visibility tasks need human-visible proof and human-world closure. Structured state, proof-only views, reference pose, or blocked proxy checkpoints are not completion.
- `C13`: Repo-local `REGRESSION.md` is authoritative. Regression means the real app-surface click path, not smoke.
- `C15` and `C17`: Navigation should be side-effect free, and hierarchy should be structurally legible rather than mixing navigation and content in one strip.
- `C18`: For desired-vs-runtime control surfaces, declared intent is the primary state story and reconcile is a secondary diff layer.

## Events that still need a wider reread

- `C08` only if a later pass wants to split the backend anti-drift moment from the later UI anti-drift moment.
- `C14` only if a later pass wants to separate the initial rejection of reference-pose closure from the later standing-restatement chain after the task was reopened.
- No other candidate needs a wider reread for local PASS2 classification. The reopened transcript windows are sufficient for event-level diagnosis.
