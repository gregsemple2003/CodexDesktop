# INTERVENTION-PASS2

Source day: `2026-04-06`

Canonical note: promoted on `2026-04-09` from the explanatory rerun; the earlier day-local PASS2 is archived as `INTERVENTION-PASS2-SUPERSEDED-2026-04-09.md`.

## Source scope analyzed

- PASS1 artifact reopened: `C:\Agent\CodexDashboard\Tracking\Task-0006\Research\Daily\2026-04-06\INTERVENTION-PASS1.md`
- Raw transcript scope reopened: the same April 6 day transcripts cited in PASS1, with bounded rereads around every candidate ref.
- Contract files used for classification: the incident corpus README, the incident schema, and the current PASS1/PASS2 prompt templates.
- Existing April 6 pass artifacts were still not used as event evidence. A short compare against them was done only at the end for the `previously missed` note below.

## Candidate ids analyzed

- `C01` through `C19`

## Boundary corrections relative to PASS1

- No major boundary corrections.
- `C16` and `C17` remain split. `C16` is the rejection of the assistant's closure model; `C17` is the forced conversion of that corrected model into a concrete service lane.
- `C18` and `C19` remain split. `C18` is the human-surface/product-read complaint; `C19` is the narrower restart-ownership correction.
- `C07` and `C08` also remain split. `C07` is the higher-level autonomy/persistence correction; `C08` is the later wrong-seam debugging redirection.

## Per-event analysis

### C01 - Reopen whether a Home progress bar belongs at all for small uploads

- `event_id`: `C01`
- `title`: `Reopen whether a Home progress bar belongs at all for small uploads`
- `session_or_thread`: `Implement task 22 tracking`
- `transcript_path`: `T01`
- `primary_refs`: `T01#L195`, `T01#L198`, `T01#L235`
- `ai_course`: The assistant had a plan-gate recommendation that still treated a progress bar as a reasonable default Home treatment for active upload.
- `human_intervention`: The human challenged that premise and forced designer consultation before approval.
- `adequate_outcome`: Re-open the Home-card design seam instead of approving a bar-forward plan on the first pass.
- `event_boundary_notes`: This is the first correction in the Task-0022 arc. It is weaker than the later time-to-done rule but still a real plan redirection.
- `human_model_signal`: Implicit signal made explicit by the question: single small uploads do not automatically justify a progress bar on Home.
- `failure_family_hypothesis`: `ui_semantics`
- `intervention_kind_hypothesis`: `reframe_problem`
- `human_cost_or_risk`: Plan approval could have ratified a misleading Home-card treatment.
- `local_lesson_hypothesis`: When a human questions the rationale for a visible status control, treat that as a prompt to re-open the semantic meaning of the control, not merely tweak wording.
- `cluster_hints`: `status-signal-meaning`, `ui-contract`
- `accepted_incident_likelihood`: `possible`
- `confidence`: `medium`
- `uncertainties`: The later `time-to-done` correction is stronger and may subsume this event in a narrower incident set.

### C02 - Make the Home progress bar mean time-to-done

- `event_id`: `C02`
- `title`: `Make the Home progress bar mean time-to-done`
- `session_or_thread`: `Implement task 22 tracking` plus `Own Task-0022 lifecycle`
- `transcript_path`: `T01`, `T02`
- `primary_refs`: `T01#L242`, `T01#L245`, `T02#L252`, `T02#L255`
- `ai_course`: The assistant still treated progress treatment as a looser indicator of upload-related activity than the human wanted.
- `human_intervention`: The human stated the decisive adequacy rule: if Home shows a progress bar, humans will read it as `how long until done`.
- `adequate_outcome`: A truthful planning baseline that makes time-to-done the meaning of a Home bar and removes bars from safe/retained states.
- `event_boundary_notes`: Clean event boundary. The correction appears in the parent thread and is then pushed into the long-lived owner.
- `human_model_signal`: Explicit human-world interpretation rule: a progress bar means remaining transfer time.
- `failure_family_hypothesis`: `usability_state_truth`
- `intervention_kind_hypothesis`: `tighten_contract`
- `human_cost_or_risk`: Without this rule the Home card would misread progress-like UI as unfinished work or remaining transfer.
- `local_lesson_hypothesis`: On a human-facing surface, treat a progress bar as a concrete promise about remaining time unless the human explicitly accepts a weaker meaning.
- `cluster_hints`: `status-signal-meaning`, `time-to-done`, `ui-contract`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: None material.

### C03 - Call out the remaining disconnect as a dropped ball

- `event_id`: `C03`
- `title`: `Call out the remaining disconnect as a dropped ball`
- `session_or_thread`: `Implement task 22 tracking`
- `transcript_path`: `T01`
- `primary_refs`: `T01#L297`, `T01#L300`, `T01#L310`
- `ai_course`: After the first correction, the assistant still left the bar sounding like a generic `I'm busy` indicator and had not made the concrete alternatives explicit.
- `human_intervention`: The human explicitly labeled this a `dropped ball` because they still had to follow up and ask where the disconnect was.
- `adequate_outcome`: State the missing semantic rule plainly and give concrete alternative UI states.
- `event_boundary_notes`: This is a repeated intervention because the earlier correction did not fully stick.
- `human_model_signal`: Strong signal that the user values reduction of follow-up burden as part of adequacy.
- `failure_family_hypothesis`: `usability_state_truth`
- `intervention_kind_hypothesis`: `tighten_contract`
- `human_cost_or_risk`: Repeated explanatory burden at the plan gate and risk of shipping a still-misleading Home card.
- `local_lesson_hypothesis`: When a human says they still had to follow up, treat that explanatory burden itself as evidence that the semantic correction has not been internalized.
- `cluster_hints`: `status-signal-meaning`, `dropped-ball`, `repeated-teaching`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: None material.

### C04 - Correct the implied promise that install plus restart should clear the visible blocked banner

- `event_id`: `C04`
- `title`: `Correct the implied promise that install plus restart should clear the visible blocked banner`
- `session_or_thread`: `Run TASK-HARVESTER`
- `transcript_path`: `T03`
- `primary_refs`: `T03#L716`, `T03#L725`, `T03#L748`, `T03#L755`
- `ai_course`: The assistant's wording made it sound like the patch/install/restart step should materially help the visible phone banner.
- `human_intervention`: The human asked whether the assistant was surprised the phone was still stuck at `Needs attention`, or whether it had failed to restart the app.
- `adequate_outcome`: Explain that the patch prevents future bad rows but does not clear the already-stranded rows without a separate repair step.
- `event_boundary_notes`: Concrete but narrower than the broader charity/regression event later in the same thread.
- `human_model_signal`: Implicit human expectation: if the assistant says the push-to-phone should help, the visible user-facing state should change or the limitation should already have been called out plainly.
- `failure_family_hypothesis`: `usability_state_truth`
- `intervention_kind_hypothesis`: `reframe_problem`
- `human_cost_or_risk`: Confusion about whether the patch actually fixed the visible problem.
- `local_lesson_hypothesis`: When describing a partial repair, separate `prevents future corruption` from `clears current visible bad state` explicitly.
- `cluster_hints`: `visible-state-truth`, `partial-fix-framing`
- `accepted_incident_likelihood`: `possible`
- `confidence`: `medium`
- `uncertainties`: This could be folded into a broader `truthful repair framing` incident if later clustering prefers fewer bug-thread incidents.

### C05 - Record the user's regression report charitably instead of flattening it into optimization

- `event_id`: `C05`
- `title`: `Record the user's regression report charitably instead of flattening it into optimization`
- `session_or_thread`: `Run TASK-HARVESTER`
- `transcript_path`: `T03`
- `primary_refs`: `T03#L1431`, `T03#L1438`, `T03#L1450`
- `ai_course`: The assistant said the system looked healthy now and framed remaining slowness as a throughput-optimization problem.
- `human_intervention`: The human said it was not `simply an optimization problem`, reported that uploads were materially faster before the fixes, required durable notes, called out the dropped ball, and asked for a principle of charity.
- `adequate_outcome`: Treat the user's before/after report as real regression evidence and record it durably.
- `event_boundary_notes`: Clear event boundary with a direct human statement of the meta-lesson.
- `human_model_signal`: Explicit signals: user-reported regression relative to earlier behavior is real evidence, and the assistant should use a principle of charity when describing the claim.
- `failure_family_hypothesis`: `human_world`
- `intervention_kind_hypothesis`: `tighten_contract`
- `human_cost_or_risk`: Trust damage and loss of important regression evidence.
- `local_lesson_hypothesis`: When a user reports that a fix made reality worse in their lived experience, preserve that claim as first-class evidence before trying to classify it.
- `cluster_hints`: `principle-of-charity`, `user-report-evidence`, `dropped-ball`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: None material.

### C06 - Bound PASS-0001 and shut down open-ended research

- `event_id`: `C06`
- `title`: `Bound PASS-0001 and shut down open-ended research`
- `session_or_thread`: `Lead Task-0005 lifecycle`
- `transcript_path`: `T05`
- `primary_refs`: `T05#L718`, `T05#L721`, `T05#L748`, `T05#L751`
- `ai_course`: The assistant was still in SDK/doc lookup mode after the runtime was already healthy enough for bounded implementation.
- `human_intervention`: The human bounded the exact slice and explicitly said `No more open-ended SDK/doc research`.
- `adequate_outcome`: Move immediately into code/tests on the named PASS-0001 slice.
- `event_boundary_notes`: Strong local correction, but it reads more like execution discipline than a durable incident unless similar drift recurs often.
- `human_model_signal`: Implied autonomy rule: once the environment is ready and the slice is already shapeable, research should collapse into implementation.
- `failure_family_hypothesis`: `workflow_orchestration`
- `intervention_kind_hypothesis`: `collapse_branches`
- `human_cost_or_risk`: Lost time and delayed source edits.
- `local_lesson_hypothesis`: After a bounded slice and live runtime are available, additional broad research should need explicit justification.
- `cluster_hints`: `boundedness`, `anti-research-drift`
- `accepted_incident_likelihood`: `unlikely`
- `confidence`: `medium`
- `uncertainties`: None material.

### C07 - Resume deterministic root-cause pursuit instead of stopping at a plausible upstream producer

- `event_id`: `C07`
- `title`: `Resume deterministic root-cause pursuit instead of stopping at a plausible upstream producer`
- `session_or_thread`: `Trace divergence root cause`
- `transcript_path`: `T06`
- `primary_refs`: `T06#L1286`, `T06#L1293`, `T06#L1296`
- `ai_course`: The assistant isolated an upstream producer and then stopped, treating that as enough to enter handoff mode.
- `human_intervention`: The human asked why it had stopped and restated the original requirement to pursue root cause deterministically until the root cause was explicit.
- `adequate_outcome`: Keep tracing through the current seam until the exact structural cause is named or a hard blocker is proven.
- `event_boundary_notes`: Strong and clean. This is a higher-level autonomy correction, not just a seam tweak.
- `human_model_signal`: Explicit autonomy model: `pursue root cause deterministically` means do not stop at the first plausible upstream producer.
- `failure_family_hypothesis`: `workflow_orchestration`
- `intervention_kind_hypothesis`: `redirect_debugging`
- `human_cost_or_risk`: Premature handoff and unresolved defect narrative.
- `local_lesson_hypothesis`: For debugging tasks with an explicit deterministic-root-cause contract, `earliest plausible producer` is not enough; the stopping condition is the exact structural cause or an honest blocker.
- `cluster_hints`: `autonomy-persistence`, `debug-stop-too-early`, `root-cause-contract`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: None material.

### C08 - Redirect the debug branch from the Head-hierarchy story to the real Hips-coercion seam

- `event_id`: `C08`
- `title`: `Redirect the debug branch from the Head-hierarchy story to the real Hips-coercion seam`
- `session_or_thread`: `Trace hierarchy mismatch`
- `transcript_path`: `T07`
- `primary_refs`: `T07#L599`, `T07#L606`, `T07#L623`
- `ai_course`: The worker had a current explanation centered on the retained Manny control rig and `Head` mismatch.
- `human_intervention`: The human redirected the branch to the narrower IKRig question: why the target `Spine` chain gets forced back to `Hips`.
- `adequate_outcome`: Investigate the exact saved-IKRig structural cause instead of lingering on the earlier seam.
- `event_boundary_notes`: Real wrong-seam redirection, but more local and technical than `C07`.
- `human_model_signal`: Implicit signal: the assistant should track the most current high-signal seam, not keep polishing the previous best explanation.
- `failure_family_hypothesis`: `workflow_orchestration`
- `intervention_kind_hypothesis`: `redirect_debugging`
- `human_cost_or_risk`: Time spent on a seam that no longer matched the current highest-value divergence.
- `local_lesson_hypothesis`: Once the remaining deterministic disagreement is narrower, move to that exact seam instead of continuing to refine the previous layer's story.
- `cluster_hints`: `wrong-seam-debugging`, `debug-scope`
- `accepted_incident_likelihood`: `possible`
- `confidence`: `medium`
- `uncertainties`: This may later be folded into `C07` if the accepted corpus prefers one broader debugging-autonomy incident.

### C09 - Clarify that the already-authenticated CLI was not an extra-money gate

- `event_id`: `C09`
- `title`: `Clarify that the already-authenticated CLI was not an extra-money gate`
- `session_or_thread`: `Implement task 5 state`
- `transcript_path`: `T04`
- `primary_refs`: `T04#L1297`, `T04#L1304`, `T04#L1307`, `T04#L1320`
- `ai_course`: The assistant described a `spend gate` around one live `codex exec` proof run even though the local CLI was already authenticated with the user's ChatGPT-backed account.
- `human_intervention`: The human asked for clarity and then explicitly said that from their perspective there was no `additional money` here.
- `adequate_outcome`: Frame the gate as permission to use the already-authenticated account and the job's side effects, not as a new-money handoff.
- `event_boundary_notes`: Important false-blocker correction. The assistant's underlying caution was understandable, but the framing was still wrong enough to require human teaching.
- `human_model_signal`: Explicit operator-boundary signal: pre-authenticated local tooling is not the same as asking the human to newly spend money.
- `failure_family_hypothesis`: `workflow_orchestration`
- `intervention_kind_hypothesis`: `reframe_problem`
- `human_cost_or_risk`: Confusion about whether task progress was blocked on a human payment action that did not actually exist.
- `local_lesson_hypothesis`: Distinguish `this uses an account-backed hosted action` from `the human must newly provide money`; otherwise the assistant can invent false blockers.
- `cluster_hints`: `approval-model`, `false-blocker`, `operator-boundary`
- `accepted_incident_likelihood`: `possible`
- `confidence`: `medium`
- `uncertainties`: A narrower accepted set may prefer the repeated-leader correction `C11` or the unsafe-launch incident `C10` over this parent-thread precursor.

### C10 - Contain an unsafe worker startup that released all scheduled jobs

- `event_id`: `C10`
- `title`: `Contain an unsafe worker startup that released all scheduled jobs`
- `session_or_thread`: `Lead Task-0005 lifecycle`
- `transcript_path`: `T05`
- `primary_refs`: `T05#L1249`, `T05#L1252`
- `ai_course`: The assistant launched the worker-backed control plane for one proof run, but startup immediately released all three schedule-triggered jobs.
- `human_intervention`: The supervisor injected a containment summary, recorded the evidence, and required honest recovery from the contained state as a concrete control problem.
- `adequate_outcome`: Treat the launch as a bounded safety/proof-scope failure and recover from that state honestly.
- `event_boundary_notes`: Local boundary is very clear and the human intervention is decisive.
- `human_model_signal`: Explicit proof-scope rule: usage must stay bounded to the intended proof; an unsafe launch must not be normalized away.
- `failure_family_hypothesis`: `workflow_orchestration`
- `intervention_kind_hypothesis`: `redirect_debugging`
- `human_cost_or_risk`: Unintended hosted job execution and trust damage around proof containment.
- `local_lesson_hypothesis`: For live hosted proofs, the assistant must prove scope fences before starting a worker-backed runtime that can wake preexisting schedules.
- `cluster_hints`: `proof-scope-control`, `runtime-containment`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: None material.

### C11 - Repeat the spend-gate clarification in the long-lived leader thread

- `event_id`: `C11`
- `title`: `Repeat the spend-gate clarification in the long-lived leader thread`
- `session_or_thread`: `Lead Task-0005 lifecycle`
- `transcript_path`: `T05`
- `primary_refs`: `T05#L1316`, `T05#L1319`
- `ai_course`: The leader was still treating spend approval as a plausible blocker and had not fully collapsed the environment confusion into the real runtime blocker.
- `human_intervention`: The supervisor explicitly removed the `additional-money` blocker again, named runtime behavior as the real blocker, and called out stale PATH as local shell noise.
- `adequate_outcome`: Keep the task focused on the actual contained-proof failures.
- `event_boundary_notes`: This is repeated human teaching because the earlier clarification in `C09` did not fully propagate to the owner thread.
- `human_model_signal`: Strong signal that false blocker framing is itself costly when it survives past an earlier clarification.
- `failure_family_hypothesis`: `workflow_orchestration`
- `intervention_kind_hypothesis`: `reframe_problem`
- `human_cost_or_risk`: Repeated explanatory burden and delayed focus on the real blocker.
- `local_lesson_hypothesis`: When a blocker correction lands in a parent/sibling thread, the long-lived owner thread needs that correction carried through immediately.
- `cluster_hints`: `approval-model`, `false-blocker`, `repeated-teaching`
- `accepted_incident_likelihood`: `possible`
- `confidence`: `medium`
- `uncertainties`: Might be folded into `C09` or `C12` in a tighter accepted set.

### C12 - Enforce durable-state-first before more debugging

- `event_id`: `C12`
- `title`: `Enforce durable-state-first before more debugging`
- `session_or_thread`: `Lead Task-0005 lifecycle`
- `transcript_path`: `T05`
- `primary_refs`: `T05#L1361`, `T05#L1364`
- `ai_course`: The assistant was still moving into deeper fence-fix work while the durable task state still read cleaner than the actual situation justified.
- `human_intervention`: The human stopped exploration and prescribed the exact `TASK-STATE`, `HANDOFF`, and bug updates that had to exist first.
- `adequate_outcome`: Make the task truthfully blocked/debugging before any more exploration.
- `event_boundary_notes`: Strong local event with a clear human-stated workflow principle.
- `human_model_signal`: Explicit task-lifecycle rule: do not let artifacts claim a cleaner state than the real proof boundary supports.
- `failure_family_hypothesis`: `workflow_orchestration`
- `intervention_kind_hypothesis`: `tighten_contract`
- `human_cost_or_risk`: Dishonest or stale artifacts would make later review and handoff unsafe.
- `local_lesson_hypothesis`: After a proof-scope failure, durable-state truth is part of the recovery itself, not optional bookkeeping.
- `cluster_hints`: `durable-state-first`, `proof-scope-control`, `bug-truth`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: None material.

### C13 - Supply exact local runtime facts and CLI semantics so the retry can proceed

- `event_id`: `C13`
- `title`: `Supply exact local runtime facts and CLI semantics so the retry can proceed`
- `session_or_thread`: `Lead Task-0005 lifecycle`
- `transcript_path`: `T05`
- `primary_refs`: `T05#L1513`, `T05#L1516`, `T05#L1538`, `T05#L1541`
- `ai_course`: The assistant was still resolving compose path, installed `codex.exe`, and sandbox semantics from incomplete environment inference.
- `human_intervention`: The human supplied the exact local paths and the real `codex exec --help` semantics from the installed binary.
- `adequate_outcome`: Take the bounded retry from concrete local facts instead of more environment guesswork.
- `event_boundary_notes`: Real explanatory labor, but likely better treated as an important supervisory event than a durable accepted incident.
- `human_model_signal`: Implicit rule: when the machine can reveal exact paths and tool semantics, the assistant should ground there rather than hedge from abstractions.
- `failure_family_hypothesis`: `workflow_orchestration`
- `intervention_kind_hypothesis`: `tighten_contract`
- `human_cost_or_risk`: Repeated human time spent teaching the local environment to the assistant.
- `local_lesson_hypothesis`: For local-runtime debugging, prefer direct machine-fact recovery early; repeated inference from stale PATH or assumed CLI behavior creates avoidable human follow-up.
- `cluster_hints`: `local-grounding`, `operator-facts`, `explanatory-burden`
- `accepted_incident_likelihood`: `unlikely`
- `confidence`: `medium`
- `uncertainties`: Could be folded into `C11` or `C14` if later clustering wants fewer Task-0005 supervisory events.

### C14 - Force an explicit bounded fork instead of more analysis

- `event_id`: `C14`
- `title`: `Force an explicit bounded fork instead of more analysis`
- `session_or_thread`: `Lead Task-0005 lifecycle`
- `transcript_path`: `T05`
- `primary_refs`: `T05#L1643`, `T05#L1651`, `T05#L1654`
- `ai_course`: The assistant still had not chosen between one final bounded retry and an honest blocker declaration.
- `human_intervention`: The human said `Decision required now`, gave the two acceptable branches, and explicitly said not to stay in analysis.
- `adequate_outcome`: Collapse the branch space and take a bounded decision with durable state behind it.
- `event_boundary_notes`: Strong supervisory correction, though probably smaller than the unsafe-launch and durable-state events.
- `human_model_signal`: Explicit autonomy rule: once the branch set is small and evidence is sufficient, choose rather than keep analyzing.
- `failure_family_hypothesis`: `workflow_orchestration`
- `intervention_kind_hypothesis`: `collapse_branches`
- `human_cost_or_risk`: Extended debugging churn with no decision.
- `local_lesson_hypothesis`: After a bounded proof and explicit branch set exist, indecision becomes its own failure mode.
- `cluster_hints`: `boundedness`, `decision-discipline`
- `accepted_incident_likelihood`: `possible`
- `confidence`: `medium`
- `uncertainties`: None material.

### C15 - Force immediate closeout after proof success

- `event_id`: `C15`
- `title`: `Force immediate closeout after proof success`
- `session_or_thread`: `Lead Task-0005 lifecycle`
- `transcript_path`: `T05`
- `primary_refs`: `T05#L1865`, `T05#L1924`, `T05#L1930`, `T05#L1932`
- `ai_course`: Even after the proof looked successful and evidence existed, the assistant still had some nonessential follow-up momentum instead of closing out directly.
- `human_intervention`: The human prescribed the exact closeout order, said not to stop at summary, and then demanded a one-line status only.
- `adequate_outcome`: Immediate durable closeout and no more nonessential follow-up.
- `event_boundary_notes`: Concrete intervention, but it reads more like end-of-task supervision than a likely durable incident.
- `human_model_signal`: Explicit closure rule: once the evidence is already sufficient, the assistant should finish the task rather than drift into more narration or side work.
- `failure_family_hypothesis`: `workflow_orchestration`
- `intervention_kind_hypothesis`: `tighten_contract`
- `human_cost_or_risk`: Delayed closure and extra supervisory overhead.
- `local_lesson_hypothesis`: After explicit proof success, closeout work should dominate immediately; summary is not a substitute for closure.
- `cluster_hints`: `closeout-discipline`, `anti-drift`
- `accepted_incident_likelihood`: `unlikely`
- `confidence`: `medium`
- `uncertainties`: None material.

### C16 - Reject the proof-only closure model for a long-running backend

- `event_id`: `C16`
- `title`: `Reject the proof-only closure model for a long-running backend`
- `session_or_thread`: `Implement task 5 state`
- `transcript_path`: `T04`
- `primary_refs`: `T04#L2638`, `T04#L2641`, `T04#L2648`, `T04#L2657`, `T04#L2660`
- `ai_course`: The assistant interpreted `long-running backend` as implemented-and-proven rather than actually left running as the delivered state.
- `human_intervention`: The human rejected that closure model, said `we exist in the real world`, stressed `I'm not trying to be an operator here`, and explained the need for jobs to keep running on the local computer.
- `adequate_outcome`: Redefine done as a usable unattended operating model, not just successful proof.
- `event_boundary_notes`: Strong explanatory-labor event with explicit higher-level meaning, operator-boundary, and implied-expectation coaching.
- `human_model_signal`: Strong explicit model: the human is not the operator, and a `long-running backend` implies usable unattended scheduling after task closure.
- `failure_family_hypothesis`: `human_world`
- `intervention_kind_hypothesis`: `reject_outcome`
- `human_cost_or_risk`: Task could be called done while leaving the human with no actual scheduler and surprise operator burden.
- `local_lesson_hypothesis`: When task wording implies an operating model, optimize for the human-world delivered state unless the task explicitly narrows the bar to proof-only.
- `cluster_hints`: `closure-truth`, `operator-boundary`, `real-world-completion-mismatch`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: None material.

### C17 - Force creation of the separate always-on service lane

- `event_id`: `C17`
- `title`: `Force creation of the separate always-on service lane`
- `session_or_thread`: `Implement task 5 state`
- `transcript_path`: `T04`
- `primary_refs`: `T04#L2667`, `T04#L2670`
- `ai_course`: After accepting the higher-level critique, the assistant had not yet turned that corrected model into the concrete always-on service lane it implied.
- `human_intervention`: The human told it to create that service lane now and document the task according to the new understanding.
- `adequate_outcome`: Build the actual service lane and make the task artifacts match the corrected operating model.
- `event_boundary_notes`: Distinct from `C16`: this is not the rejection of the closure model; it is the demand for the concrete structural fix.
- `human_model_signal`: Explicit operating-model signal: a company-of-agents outcome needs a stable service lane separate from disposable validation infrastructure.
- `failure_family_hypothesis`: `workflow_orchestration`
- `intervention_kind_hypothesis`: `reframe_problem`
- `human_cost_or_risk`: Assistant could stop at the lesson and still leave the product unchanged.
- `local_lesson_hypothesis`: Once a higher-level operating-model miss is explicit, the assistant should immediately convert that lesson into the missing concrete lane or mechanism.
- `cluster_hints`: `service-vs-validation-lane`, `closure-truth`, `operator-boundary`
- `accepted_incident_likelihood`: `possible`
- `confidence`: `strong`
- `uncertainties`: A tighter accepted set may prefer only `C16` as the durable incident and treat `C17` as the enacted repair step.

### C18 - Call out the dropped ball on the dashboard surface itself

- `event_id`: `C18`
- `title`: `Call out the dropped ball on the dashboard surface itself`
- `session_or_thread`: `Implement task 5 state`
- `transcript_path`: `T04`
- `primary_refs`: `T04#L3020`, `T04#L3027`, `T04#L3030`, `T04#L3055`
- `ai_course`: The assistant answered `Refresh` versus `Force Reconcile` correctly from backend/docs truth but left the human-facing dashboard wording unchanged and did not proactively offer that fix.
- `human_intervention`: The human called this `another dropped ball`, said humans should not need docs or CLI for that surface, and said the form-factor burden rests with the assistant.
- `adequate_outcome`: Put the meaning directly on the dashboard surface and stop relying on docs/CLI interpretation for a quick-use control.
- `event_boundary_notes`: Strong human-facing surface event with explicit human-world reasoning.
- `human_model_signal`: Explicit rule: for quick human-facing controls, the right question is what the human would prefer the assistant to do, not whether the human will notice the omission.
- `failure_family_hypothesis`: `information_architecture`
- `intervention_kind_hypothesis`: `reframe_problem`
- `human_cost_or_risk`: Hidden meaning and operator burden on a surface meant for quick use.
- `local_lesson_hypothesis`: If the dashboard is the human-facing control surface, meaning must live there first; docs and CLI are support, not the primary form factor.
- `cluster_hints`: `control-surface`, `form-factor-protection`, `dropped-ball`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: None material.

### C19 - Reject the restart handoff back to the human

- `event_id`: `C19`
- `title`: `Reject the restart handoff back to the human`
- `session_or_thread`: `Implement task 5 state`
- `transcript_path`: `T04`
- `primary_refs`: `T04#L3055`, `T04#L3062`, `T04#L3065`
- `ai_course`: Even after accepting the human-surface critique, the assistant still ended with `restart the dashboard app so it reloads the updated UI code`.
- `human_intervention`: The human rejected that handoff and said the assistant should know they did not have a restart button.
- `adequate_outcome`: The assistant should own the restart and leave the corrected surface actually running.
- `event_boundary_notes`: Smaller than `C18`, but still a real repeated-boundary correction because the first human-surface lesson did not fully stick.
- `human_model_signal`: Explicit operator-boundary signal: if the assistant can restart the app, it should not leave that burden on the human by default.
- `failure_family_hypothesis`: `workflow_orchestration`
- `intervention_kind_hypothesis`: `tighten_contract`
- `human_cost_or_risk`: The human could be left unable to realize the surface fix the assistant just claimed to have made.
- `local_lesson_hypothesis`: After fixing a human-facing surface, own the last-mile action needed to make that fix live when the human lacks an obvious control.
- `cluster_hints`: `operator-boundary`, `last-mile-ownership`
- `accepted_incident_likelihood`: `possible`
- `confidence`: `strong`
- `uncertainties`: Many accepted sets would keep this as a small follow-on to `C18` rather than a separate durable incident.

## Likely accepted incidents

- `C02` Make the Home progress bar mean time-to-done.
- `C03` Call out the remaining disconnect as a dropped ball.
- `C05` Record the user's regression report charitably instead of flattening it into optimization.
- `C07` Resume deterministic root-cause pursuit instead of stopping at a plausible upstream producer.
- `C10` Contain an unsafe worker startup that released all scheduled jobs.
- `C12` Enforce durable-state-first before more debugging.
- `C16` Reject the proof-only closure model for a long-running backend.
- `C18` Call out the dropped ball on the dashboard surface itself.

## Likely non-incident but still important intervention events

- `C01` Home-bar rationale had to be reopened before plan approval.
- `C06` PASS-0001 needed boundedness and anti-research discipline.
- `C13` Human supplied exact local runtime facts and CLI semantics.
- `C15` Human forced immediate closeout after the proof was already sufficient.

## Repeated cluster hints across the analyzed set

- `status-signal-meaning`
  - `C01`, `C02`, `C03`
- `user-report-evidence` / `principle-of-charity`
  - `C04`, `C05`
- `autonomy-persistence` / `wrong-seam-debugging`
  - `C07`, `C08`
- `approval-model` / `false-blocker`
  - `C09`, `C11`, `C13`
- `proof-scope-control` / `durable-state-first`
  - `C10`, `C12`, `C14`, `C15`
- `closure-truth` / `operator-boundary`
  - `C16`, `C17`
- `control-surface` / `last-mile-ownership`
  - `C18`, `C19`

## Strongest human-model signals worth carrying forward

- A Home progress bar means `how long until done`; if that promise is not true, do not use the bar.
- When the human says they still had to follow up, that explanatory burden is itself evidence of a miss.
- A user report that `it was faster before your fix` is real regression evidence, not something to explain away as mere optimization without comparative proof.
- `Pursue root cause deterministically` means keep tracing until the exact structural cause is explicit, not until a plausible upstream producer is found.
- A pre-authenticated local hosted tool is not automatically an `additional money` blocker.
- After a proof-scope failure, durable-state truth is part of the recovery path.
- `Long-running backend` in a local-first system implies a usable unattended operating model, not just a successful demo.
- The human is not the operator by default.
- If the dashboard is the human-facing control surface, meaning must live there, and the assistant should own last-mile actions like restart when the human lacks the affordance.

## Explanatory-labor / repeated-human-teaching events

- Best fit as explanatory-labor / human-teaching: `C01`, `C02`, `C03`, `C05`, `C06`, `C07`, `C08`, `C09`, `C11`, `C12`, `C13`, `C14`, `C15`, `C16`, `C17`, `C18`, `C19`.
- Strongest explanatory-labor events: `C02`, `C03`, `C05`, `C07`, `C12`, `C16`, `C18`.
- Explanatory-labor events that also look like likely accepted incidents: `C02`, `C03`, `C05`, `C07`, `C12`, `C16`, `C18`.

## Previously missed higher-level coaching events in this rerun

- Yes. The clearest newly surfaced higher-level coaching event is `C07`.
- Short compare note only: the existing April 6 PASS1/PASS2 artifacts already appear to capture most of the major `Task-0022` and `Task-0005` explanatory-labor events, including the Home time-to-done rule, the charity/regression note, the service-lane / non-operator critique, the dashboard-surface dropped ball, the anti-research correction, the spend-gate clarification cluster, and the durable-state-first rule.
- What this rerun adds most clearly is the `ThirdPerson` autonomy-standard event: the human had to reteach that `deterministic root-cause pursuit` means do not stop at a plausible upstream producer. `C08` is a related but narrower newly clearer wrong-seam redirection.

## Events that still need a wider reread

- None obvious.
- The remaining judgment calls are mostly clustering choices:
  - whether `C07` and `C08` should stay separate
  - whether `C16` and `C17` should stay separate in a durable incident layer
  - whether `C18` and `C19` are best modeled as one human-surface incident with a repeated follow-up or as two neighboring events

## Bottom line

- The widened criteria do recover explanatory-labor cleanly on April 6.
- The day does not look like it is still missing an obvious high-signal explanatory-labor event after this rerun.
- The largest remaining modeling choice is not recall but how aggressively to merge adjacent `Task-0005` supervisory events when narrowing from local pass analysis into accepted incidents.
