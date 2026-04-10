# INTERVENTION-PASS2

Source day: `2026-04-06`

## Source scope analyzed

- PASS1 artifact: `C:\Agent\CodexDashboard\Tracking\Task-0006\Research\Daily\2026-04-06\INTERVENTION-PASS1.md`
- Incident corpus contract reread:
  - `C:\Users\gregs\.codex\Orchestration\Reports\Incidents\README.md`
  - `C:\Users\gregs\.codex\Orchestration\Reports\Incidents\INCIDENT.schema.json`
- Raw transcripts reopened from the PASS1 refs:
  - `T01` = `C:\Users\gregs\.codex\sessions\2026\04\06\rollout-2026-04-06T18-59-50-019d6506-147a-7c00-972d-2498ff61ff81.jsonl`
  - `T02` = `C:\Users\gregs\.codex\sessions\2026\04\06\rollout-2026-04-06T19-28-46-019d6520-919c-7b60-94c2-422680147d6c.jsonl`
  - `T03` = `C:\Users\gregs\.codex\sessions\2026\04\06\rollout-2026-04-06T19-30-25-019d6522-1369-7692-bdba-b53998c88601.jsonl`
  - `T04` = `C:\Users\gregs\.codex\sessions\2026\04\06\rollout-2026-04-06T19-54-03-019d6537-b945-7173-888e-667687b4d67c.jsonl`
  - `T05` = `C:\Users\gregs\.codex\sessions\2026\04\06\rollout-2026-04-06T20-36-54-019d655e-f1c5-7463-af8a-779567b8955a.jsonl`
  - `T06` = `C:\Users\gregs\.codex\sessions\2026\04\06\rollout-2026-04-06T20-39-30-019d6561-529c-75e0-9a57-ab9d9fd66d80.jsonl`
  - `T07` = `C:\Users\gregs\.codex\sessions\2026\04\06\rollout-2026-04-06T20-43-54-019d6565-591e-7831-b050-55f12d0b304f.jsonl`
- Direct task-artifact checks pulled in only where the transcript explicitly hinged on expected-state wording:
  - `C:\Agent\CodexDashboard\Tracking\Task-0005\TASK.md`
  - `C:\Agent\Crystallize\Tracking\Task-0022\PLAN.md`

## Candidate ids analyzed

- `C01` through `C26`

## Boundary corrections relative to PASS1

- `C10` needed a real boundary correction. PASS1 titled it as a meta-correction about inferred human meaning, but the cited refs `T05:L116` and `T05:L123` only support a much smaller gate-clear event: the human removed future human gates and told the task to continue until closure or blocker. The broader coaching about filling in human meaning appears later in the `C12`/`C13` arc, not at the cited `C10` refs.
- `C12` and `C13` remain split. `C12` is the rejection of the assistant's closure model; `C13` is the forced conversion of that corrected model into a new concrete service lane.
- `C23` and `C24` remain split. `C23` forces transition into durable closeout once proof looks done; `C24` is a second intervention because explanatory churn and extra follow-up still remained after the first closeout order.

## Per-event analysis records

### C01 - Re-scope away from Task-0004 closure

- `event_id`: `C01`
- `title`: `Re-scope away from Task-0004 closure`
- `session_or_thread`: `Harvest next task candidates`
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\06\rollout-2026-04-06T18-59-50-019d6506-147a-7c00-972d-2498ff61ff81.jsonl`
- `primary_refs`: `T01:L223`, `T01:L230`
- `ai_course`: The assistant framed Task-0004 as the only real open task, described the exact remaining regression steps, and demoted Task-0005 to "not a real open task yet."
- `human_intervention`: The human explicitly closed Task-0004 as out of scope, changed direction to the new backend/orchestration workstream, and inserted a Chrome-export prerequisite into Task-0005 research before anything else.
- `adequate_outcome`: Stop optimizing for Task-0004 closure and pivot immediately into Task-0005 bootstrap work in the new product direction.
- `event_boundary_notes`: Single-step course correction. PASS1 boundary is accurate.
- `human_model_signal`: The human treated active-task truth as changed reality, not a suggestion: Task-0004 was no longer the right closure target once the new backend/orchestration direction was chosen.
- `failure_family_hypothesis`: `workflow_orchestration`
- `intervention_kind_hypothesis`: `reframe_problem`
- `human_cost_or_risk`: Time would have been spent closing the wrong task while the actual new workstream stayed unbootstrapped.
- `local_lesson_hypothesis`: When the human changes product direction and task priority, rewrite the active-task model immediately instead of defending the previous queue.
- `cluster_hints`: `closure-truth`, `workflow-reframe`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: none material

### C02 - Make the Home progress bar mean time-to-done

- `event_id`: `C02`
- `title`: `Make the Home progress bar mean time-to-done`
- `session_or_thread`: `Implement task 22 tracking`
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\06\rollout-2026-04-06T19-28-46-019d6520-919c-7b60-94c2-422680147d6c.jsonl`
- `primary_refs`: `T02:L235`, `T02:L242`
- `ai_course`: The assistant had already moved away from using a bar in the safe state, but it still explained the bar in generic "visible progress" terms rather than in explicit human expectation terms.
- `human_intervention`: The human stated the decisive adequacy rule: if Home shows a progress bar, humans will read it as "how long until done."
- `adequate_outcome`: The task plan and any explanation of the Home card must treat a progress bar as a time-to-done promise, not generic upload-related activity.
- `event_boundary_notes`: Tight local correction of the UI meaning model. PASS1 boundary is accurate.
- `human_model_signal`: Explicit human-world interpretation rule: a progress bar means "how long until done."
- `failure_family_hypothesis`: `ui_semantics`
- `intervention_kind_hypothesis`: `tighten_contract`
- `human_cost_or_risk`: A misleading progress signal would make a safe or tiny upload read as unfinished work.
- `local_lesson_hypothesis`: On a human-facing surface, treat a progress bar as a commitment about remaining time unless the human explicitly accepts a weaker meaning.
- `cluster_hints`: `status-signal-meaning`, `time-to-done`, `ui-contract`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: none material

### C03 - Call out the remaining disconnect as a dropped ball

- `event_id`: `C03`
- `title`: `Call out the remaining disconnect as a dropped ball`
- `session_or_thread`: `Implement task 22 tracking`
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\06\rollout-2026-04-06T19-28-46-019d6520-919c-7b60-94c2-422680147d6c.jsonl`
- `primary_refs`: `T02:L290`, `T02:L297`
- `ai_course`: The assistant said the revised docs encoded the correct rule, but its explanation still sounded to the human like the bar just meant "I'm busy" with no concrete alternative.
- `human_intervention`: The human explicitly labeled this a "dropped ball" because they still had to follow up and ask where the disconnect was.
- `adequate_outcome`: Explain the corrected rule clearly enough that the human does not need a second repair pass to recover what changed.
- `event_boundary_notes`: Distinct from `C02`. `C02` corrected the rule; `C03` corrected the assistant's still-ambiguous explanation of that rule.
- `human_model_signal`: Explicit adequacy bar: the assistant's explanation should eliminate follow-up burden, not leave the human to chase the missing interpretation.
- `failure_family_hypothesis`: `ui_semantics`
- `intervention_kind_hypothesis`: `tighten_contract`
- `human_cost_or_risk`: Extra follow-up burden and reduced trust in whether the corrected design rule was actually understood.
- `local_lesson_hypothesis`: After a design rule is corrected, restate it in the human's meaning model, not in a watered-down paraphrase that reopens ambiguity.
- `cluster_hints`: `status-signal-meaning`, `explanation-adequacy`, `follow-up-burden`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: none material

### C04 - Re-ground Task-0022 planning with the clarified no-bar rule

- `event_id`: `C04`
- `title`: `Re-ground Task-0022 planning with the clarified no-bar rule`
- `session_or_thread`: `Own Task-0022 lifecycle`
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\06\rollout-2026-04-06T19-30-25-019d6522-1369-7692-bdba-b53998c88601.jsonl`
- `primary_refs`: `T03:L245`, `T03:L252`
- `ai_course`: The assistant had Task-0022 at the plan-approval gate with a safe-state fix, but the planning baseline still needed the human's clarified time-to-done rule and the explicit "safe state never shows a bar" rule.
- `human_intervention`: The human supervisory message forced a re-ground from disk and required the task artifacts to encode that safe state never shows a bar, retained local copies must not read like remaining transfer time, and only sustained active uploads can justify progress treatment.
- `adequate_outcome`: A truthful planning baseline that makes the time-to-done rule explicit before plan approval resumes.
- `event_boundary_notes`: This is a durable-artifact correction, not just another conversational opinion. PASS1 boundary is accurate.
- `human_model_signal`: Explicit model of adequacy: Home progress treatment is valid only when it honestly helps a human judge time-to-done; safe or already-safe states should use calmer text or chips.
- `failure_family_hypothesis`: `ui_semantics`
- `intervention_kind_hypothesis`: `tighten_contract`
- `human_cost_or_risk`: Without the durable re-grounding, the planning baseline could still authorize a misleading Home state.
- `local_lesson_hypothesis`: When the human clarifies a human-facing meaning rule during plan review, put that rule into the durable planning artifacts before asking for approval again.
- `cluster_hints`: `status-signal-meaning`, `durable-contract`, `time-to-done`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: none material

### C05 - Force the move from source-level fix to live phone patching

- `event_id`: `C05`
- `title`: `Force the move from source-level fix to live phone patching`
- `session_or_thread`: `Run TASK-HARVESTER`
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\06\rollout-2026-04-06T19-54-03-019d6537-b945-7173-888e-667687b4d67c.jsonl`
- `primary_refs`: `T04:L644`, `T04:L651`
- `ai_course`: The assistant had root cause, tests, build output, and task notes, but it stopped short of actually installing the patched APK to the phone and validating the live surface.
- `human_intervention`: The human said to patch the phone app and restart it now.
- `adequate_outcome`: Move from source-level correctness to on-device validation on the real failing surface.
- `event_boundary_notes`: Small but real nudge past an early stop point. It is weaker than the later phone-repair interventions.
- `human_model_signal`: none explicit
- `failure_family_hypothesis`: `verification_proof`
- `intervention_kind_hypothesis`: `reject_outcome`
- `human_cost_or_risk`: The repo would look repaired while the actual phone remained unchanged.
- `local_lesson_hypothesis`: When the remaining honest step is live-device validation, do it instead of waiting for the human to push past the source-only milestone.
- `cluster_hints`: `proof-scope`, `real-device-validation`
- `accepted_incident_likelihood`: `unlikely`
- `confidence`: `medium`
- `uncertainties`: The transcript does not prove whether the assistant would have installed the build without the prompt; it only shows that it had not done so yet.

### C06 - Correct the implied expectation around restart and force a bounded manual repair

- `event_id`: `C06`
- `title`: `Correct the implied expectation around restart and force a bounded manual repair`
- `session_or_thread`: `Run TASK-HARVESTER`
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\06\rollout-2026-04-06T19-54-03-019d6537-b945-7173-888e-667687b4d67c.jsonl`
- `primary_refs`: `T04:L709`, `T04:L716`, `T04:L725`, `T04:L748`, `T04:L755`
- `ai_course`: The assistant installed and restarted the patched app, but its earlier wording had implied that pushing the patch to the phone might clear the current blocked banner.
- `human_intervention`: The human challenged that implication, asked whether the assistant was surprised or had failed to restart, and then authorized a careful manual repair limited to the exact `23` rows once the assistant admitted the earlier wording implied too much.
- `adequate_outcome`: An honest explanation that install plus restart prevents new false rows but does not heal already-stranded rows, followed by a narrowly bounded repair.
- `event_boundary_notes`: Multi-turn correction arc. PASS1 boundary is accurate and should stay intact.
- `human_model_signal`: Two explicit signals: do not imply a restart should fix the visible state when it will not, and any manual write must be bounded to the exact `23` affected rows.
- `failure_family_hypothesis`: `verification_proof`
- `intervention_kind_hypothesis`: `reframe_problem`
- `human_cost_or_risk`: Misleading expectations about what the patch changed, plus risk of an overly broad on-device write.
- `local_lesson_hypothesis`: Distinguish "prevents future recurrence" from "repairs current visible state" before claiming what install and restart will accomplish.
- `cluster_hints`: `real-world-completion-mismatch`, `repair-scope-discipline`, `status-signal-meaning`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: none material

### C07 - Reject the "phone is repaired" claim because uploads were still stalled

- `event_id`: `C07`
- `title`: `Reject the "phone is repaired" claim because uploads were still stalled`
- `session_or_thread`: `Run TASK-HARVESTER`
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\06\rollout-2026-04-06T19-54-03-019d6537-b945-7173-888e-667687b4d67c.jsonl`
- `primary_refs`: `T04:L1012`, `T04:L1019`
- `ai_course`: The assistant declared "The phone is repaired" because the `23` blocked rows were cleared and the banner symptom it had targeted was gone.
- `human_intervention`: The human rejected that closure because the actual surface still said "waiting to upload" with roughly `3400` clips remaining and ordered the assistant to keep fixing until it was truly fixed.
- `adequate_outcome`: Treat "fixed" as resumed real upload behavior, not merely one repaired metadata symptom.
- `event_boundary_notes`: Clean rejection of premature closure. The slower-throughput framing later in `C08` is a separate event.
- `human_model_signal`: Explicit human-world adequacy rule: if uploads are still not progressing, the phone is not repaired regardless of one cleared sub-symptom.
- `failure_family_hypothesis`: `human_world`
- `intervention_kind_hypothesis`: `reject_outcome`
- `human_cost_or_risk`: Premature closure on a still-broken real-world workflow.
- `local_lesson_hypothesis`: Do not claim a device is repaired until the end-user-visible behavior that motivated the fix is actually working.
- `cluster_hints`: `closure-truth`, `real-world-completion-mismatch`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: none material

### C08 - Reject the "optimization problem" framing and require durable notes plus charity

- `event_id`: `C08`
- `title`: `Reject the "optimization problem" framing and require durable notes plus charity`
- `session_or_thread`: `Run TASK-HARVESTER`
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\06\rollout-2026-04-06T19-54-03-019d6537-b945-7173-888e-667687b4d67c.jsonl`
- `primary_refs`: `T04:L1431`, `T04:L1438`
- `ai_course`: The assistant used a narrow log-based reading to say the queue looked healthy and the remaining issue was a throughput optimization problem, not the stuck-upload bug.
- `human_intervention`: The human said it was not "simply an optimization problem," reported that uploads had been materially faster before the assistant's fixes, required durable notes about that user report, called out the dropped ball, and asked for a principle of charity in future framing.
- `adequate_outcome`: Treat the user's report as live regression evidence and record it durably, rather than explaining it away as post-bug optimization chatter.
- `event_boundary_notes`: Distinct from `C09`. `C08` is about framing the problem and preserving the user's report charitably; `C09` raises the evidentiary bar further.
- `human_model_signal`: Explicit signals: user-reported regression relative to earlier behavior is real evidence, and the assistant should use a principle of charity when describing the user's claim.
- `failure_family_hypothesis`: `verification_proof`
- `intervention_kind_hypothesis`: `reframe_problem`
- `human_cost_or_risk`: Trust damage and loss of important regression evidence if the assistant minimizes what the user is observing.
- `local_lesson_hypothesis`: When the human says behavior regressed after your fix, record that as first-class evidence before downgrading it to optimization.
- `cluster_hints`: `user-report-as-evidence`, `charity-framing`, `real-world-completion-mismatch`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: none material

### C09 - Force objective speed analysis before disconnect

- `event_id`: `C09`
- `title`: `Force objective speed analysis before disconnect`
- `session_or_thread`: `Run TASK-HARVESTER`
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\06\rollout-2026-04-06T19-54-03-019d6537-b945-7173-888e-667687b4d67c.jsonl`
- `primary_refs`: `T04:L1450`, `T04:L1457`
- `ai_course`: The assistant had updated the bug notes to preserve the user's regression report but had not yet produced objective comparative speed analysis from the existing logs.
- `human_intervention`: The human required a numbers-backed speed analysis "to server as objective data" based on the logs already captured and asked the assistant to update the bug while the phone was being disconnected.
- `adequate_outcome`: Move from durable note-taking to objective comparative evidence.
- `event_boundary_notes`: Narrow evidence-standard correction. Smaller than `C08`, but still real.
- `human_model_signal`: Explicit adequacy rule: if logs already exist, back the regression claim with objective numbers rather than leaving it as a subjective report.
- `failure_family_hypothesis`: `verification_proof`
- `intervention_kind_hypothesis`: `tighten_contract`
- `human_cost_or_risk`: The bug record would remain under-evidenced and easier to wave away later.
- `local_lesson_hypothesis`: When the human flags regression and the logs already exist, quantify the comparison proactively.
- `cluster_hints`: `user-report-as-evidence`, `objective-proof`, `diagnostic-scope`
- `accepted_incident_likelihood`: `possible`
- `confidence`: `medium`
- `uncertainties`: This may be too small for the accepted incident set even though it tightened the evidentiary standard.

### C10 - Boundary correction: auto-approve Task-0005 human gates

- `event_id`: `C10`
- `title`: `Boundary correction: auto-approve Task-0005 human gates`
- `session_or_thread`: `Implement task 5 state`
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\06\rollout-2026-04-06T20-36-54-019d655e-f1c5-7463-af8a-779567b8955a.jsonl`
- `primary_refs`: `T05:L116`, `T05:L123`
- `ai_course`: The assistant paused at the explicit planning gate and asked for plan approval before continuing the task.
- `human_intervention`: The human removed future human gates by saying "Implement, auto-approval for all human gates. Continue until the task is closed or blocker reached."
- `adequate_outcome`: Carry the task forward under explicit durable auto-approval instead of stopping at that plan gate.
- `event_boundary_notes`: PASS1's original title and summary do not fit these refs. The transcript at these refs supports only a gate-clear event, not the later broader coaching about inferred human meaning.
- `human_model_signal`: none explicit beyond a desire to eliminate avoidable human gate chasing
- `failure_family_hypothesis`: `workflow_orchestration`
- `intervention_kind_hypothesis`: `collapse_branches`
- `human_cost_or_risk`: Unnecessary pause and extra gate traffic on a task the human wanted carried through autonomously.
- `local_lesson_hypothesis`: Once the human explicitly grants auto-approval for remaining gates, stop re-requesting local gate clearance.
- `cluster_hints`: `gate-removal`, `autonomy-boundary`
- `accepted_incident_likelihood`: `unlikely`
- `confidence`: `medium`
- `uncertainties`: If a later pass wants the stronger "infer likely human meaning so the human intervenes less" coaching event, it needs different refs than the ones PASS1 cited here.

### C11 - Correct the spend-gate framing on the already-authenticated Codex CLI

- `event_id`: `C11`
- `title`: `Correct the spend-gate framing on the already-authenticated Codex CLI`
- `session_or_thread`: `Implement task 5 state`
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\06\rollout-2026-04-06T20-36-54-019d655e-f1c5-7463-af8a-779567b8955a.jsonl`
- `primary_refs`: `T05:L1297`, `T05:L1304`, `T05:L1307`, `T05:L1320`
- `ai_course`: The assistant described a "spend gate" around one live `codex exec` proof run even though the local CLI was already authenticated with the user's ChatGPT-backed account.
- `human_intervention`: The human pushed for clearer wording and explicitly said that, from their point of view, there was no "additional money" issue because the assistant was already using their subscription.
- `adequate_outcome`: Frame the decision as permission to use the already-authenticated account and separately name any real side effects such as email, rather than implying fresh payment collection.
- `event_boundary_notes`: PASS1 boundary is accurate. This is a local approval-model correction, not yet the fuller blocker correction later in `C19`.
- `human_model_signal`: Explicit signal: in an already-authenticated subscription context, "additional money" is the wrong framing.
- `failure_family_hypothesis`: `workflow_orchestration`
- `intervention_kind_hypothesis`: `reframe_problem`
- `human_cost_or_risk`: False blocker framing and unnecessary confusion around what permission is actually being requested.
- `local_lesson_hypothesis`: Separate "uses the already-authenticated account" from "requires fresh spending" when asking permission for a live proof run.
- `cluster_hints`: `approval-model`, `operator-boundary`, `false-blocker`
- `accepted_incident_likelihood`: `possible`
- `confidence`: `medium`
- `uncertainties`: It is not fully clear from this local event alone whether the spend framing was a durable blocker or a one-off explanation miss.

### C12 - Reject human-world closure where the backend is not actually left running

- `event_id`: `C12`
- `title`: `Reject human-world closure where the backend is not actually left running`
- `session_or_thread`: `Implement task 5 state`
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\06\rollout-2026-04-06T20-36-54-019d655e-f1c5-7463-af8a-779567b8955a.jsonl`
- `primary_refs`: `T05:L2641`, `T05:L2648`, `T05:L2650`, `T05:L2657`
- `ai_course`: The assistant explained that it had interpreted Task-0005 as "build and prove a long-running backend exists" rather than "leave that backend running as the delivered state," so it shut the runtime down after proof.
- `human_intervention`: The human rejected that closure model, explicitly said this was not what a human would interpret as done, stressed that they were not the operator, and said the real need was for jobs to be scheduled on the local computer.
- `adequate_outcome`: Closure must leave a usable always-on system, not merely validated code plus teardown.
- `event_boundary_notes`: This is the rejection of the assistant's human-world closure model. The follow-on order to build the new lane now is captured separately in `C13`.
- `human_model_signal`: Strong explicit model: "we exist in the real world," "I'm not trying to be an operator here," and "long-running backend" implies usable scheduled jobs on the local machine after task closure.
- `failure_family_hypothesis`: `human_world`
- `intervention_kind_hypothesis`: `reject_outcome`
- `human_cost_or_risk`: The task could be declared done while leaving the human with no actually running scheduler and unexpected operator burden.
- `local_lesson_hypothesis`: For a service described as long-running, default closure to "usable after I'm gone," not merely "implemented and proven once."
- `cluster_hints`: `closure-truth`, `operator-boundary`, `real-world-completion-mismatch`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: none material

### C13 - Force creation of the separate always-on service lane

- `event_id`: `C13`
- `title`: `Force creation of the separate always-on service lane`
- `session_or_thread`: `Implement task 5 state`
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\06\rollout-2026-04-06T20-36-54-019d655e-f1c5-7463-af8a-779567b8955a.jsonl`
- `primary_refs`: `T05:L2660`, `T05:L2667`
- `ai_course`: After acknowledging the higher-level correction, the assistant had articulated the right model but had not yet turned it into concrete implementation work.
- `human_intervention`: The human told the assistant to make that service lane now and document everything according to the new understanding.
- `adequate_outcome`: Create the actual persistent service lane and its documentation, not just an abstract reflection on what should have happened.
- `event_boundary_notes`: Kept separate from `C12` because it converts a corrected interpretation into a concrete new work order.
- `human_model_signal`: Explicit operating-model signal: the autonomous company-of-agents outcome should have a stable always-on service lane distinct from a disposable validation lane.
- `failure_family_hypothesis`: `workflow_orchestration`
- `intervention_kind_hypothesis`: `tighten_contract`
- `human_cost_or_risk`: The assistant could stop at self-critique instead of delivering the corrected operating model.
- `local_lesson_hypothesis`: Once the human names the correct operating model, convert it into implementation and documentation immediately rather than treating the insight as sufficient.
- `cluster_hints`: `operator-boundary`, `service-vs-validation-lane`, `closure-truth`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: none material

### C14 - Call out the dropped ball on the dashboard surface itself

- `event_id`: `C14`
- `title`: `Call out the dropped ball on the dashboard surface itself`
- `session_or_thread`: `Implement task 5 state`
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\06\rollout-2026-04-06T20-36-54-019d655e-f1c5-7463-af8a-779567b8955a.jsonl`
- `primary_refs`: `T05:L3020`, `T05:L3027`
- `ai_course`: The assistant answered the `Refresh` versus `Force Reconcile` question correctly using backend and docs truth, but it left the human-facing dashboard itself unchanged and did not proactively offer that fix.
- `human_intervention`: The human called this "another dropped ball," said humans should not have to read docs or use the CLI for that surface, and said the form factor burden rests with the assistant.
- `adequate_outcome`: Put the meaning directly on the dashboard or at least proactively offer that surface fix.
- `event_boundary_notes`: Local boundary is clear. The restart-burden correction that follows is separate and captured as `C15`.
- `human_model_signal`: Strong explicit signal: on a human-facing control surface, meaning belongs on the surface itself, not hidden in docs or CLI.
- `failure_family_hypothesis`: `information_architecture`
- `intervention_kind_hypothesis`: `reframe_problem`
- `human_cost_or_risk`: The human has to translate backend truth into UI meaning manually in a time-sensitive surface.
- `local_lesson_hypothesis`: When a human-facing control surface is ambiguous, treat docs and CLI explanation as inadequate and move the meaning onto the surface itself.
- `cluster_hints`: `form-factor-protection`, `status-signal-meaning`, `docs-vs-surface`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: none material

### C15 - Force the assistant to restart the app instead of handing that burden back

- `event_id`: `C15`
- `title`: `Force the assistant to restart the app instead of handing that burden back`
- `session_or_thread`: `Implement task 5 state`
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\06\rollout-2026-04-06T20-36-54-019d655e-f1c5-7463-af8a-779567b8955a.jsonl`
- `primary_refs`: `T05:L3055`, `T05:L3062`
- `ai_course`: After patching the dashboard UI copy, the assistant told the human to restart the app to see the change.
- `human_intervention`: The human rejected that handoff and said the assistant should know they did not have a restart button.
- `adequate_outcome`: The agent owns the restart step on the human's behalf when the human cannot realistically perform it from the surface.
- `event_boundary_notes`: Small but concrete operator-boundary correction. It is narrower than `C14`.
- `human_model_signal`: Explicit control-boundary rule: do not hand the human a step they have no affordance to perform.
- `failure_family_hypothesis`: `human_world`
- `intervention_kind_hypothesis`: `reject_outcome`
- `human_cost_or_risk`: The fix would technically exist while still being unusable to the human.
- `local_lesson_hypothesis`: When the agent changes a human-facing surface, it should complete the activation step itself if the human lacks the necessary control.
- `cluster_hints`: `operator-boundary`, `control-surface`, `handoff-burden`
- `accepted_incident_likelihood`: `unlikely`
- `confidence`: `strong`
- `uncertainties`: This may be better carried as a local control-boundary example than as a standalone accepted incident.

### C16 - Nudge the task leader out of silence and back into concrete work

- `event_id`: `C16`
- `title`: `Nudge the task leader out of silence and back into concrete work`
- `session_or_thread`: `Lead Task-0005 lifecycle`
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\06\rollout-2026-04-06T20-39-30-019d6561-529c-75e0-9a57-ab9d9fd66d80.jsonl`
- `primary_refs`: `T06:L692`, `T06:L718`
- `ai_course`: The assistant said it was moving into the backend slice, but then there was no visible repo movement after the runtime and tooling setup.
- `human_intervention`: The human asked for a concise blocker-or-next-step status and explicitly told the assistant to continue immediately and stop researching if it was stuck there.
- `adequate_outcome`: Restore visible implementation momentum and make the current concrete next step explicit.
- `event_boundary_notes`: Real stall correction, but weaker than the later bounded PASS-0001 intervention.
- `human_model_signal`: none explicit beyond a requirement for visible momentum and concrete next-step communication
- `failure_family_hypothesis`: `workflow_orchestration`
- `intervention_kind_hypothesis`: `other`
- `human_cost_or_risk`: Silent drift after context compaction and unclear task ownership.
- `local_lesson_hypothesis`: After a compaction or setup phase, surface the next concrete implementation step quickly rather than going quiet.
- `cluster_hints`: `momentum-loss`, `research-drift`
- `accepted_incident_likelihood`: `unlikely`
- `confidence`: `medium`
- `uncertainties`: The transcript does not fully distinguish true drift from compaction latency; it only shows that the human had to re-activate the course.

### C17 - Bound PASS-0001 and shut down open-ended research

- `event_id`: `C17`
- `title`: `Bound PASS-0001 and shut down open-ended research`
- `session_or_thread`: `Lead Task-0005 lifecycle`
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\06\rollout-2026-04-06T20-39-30-019d6561-529c-75e0-9a57-ab9d9fd66d80.jsonl`
- `primary_refs`: `T06:L724`, `T06:L748`
- `ai_course`: Even after saying it would move into code, the assistant was still doing SDK or method lookup and had not yet started the bounded PASS-0001 implementation.
- `human_intervention`: The human bounded the exact PASS-0001 slice, explicitly said "No more open-ended SDK/doc research," and named the acceptable shortcuts and expected proofs.
- `adequate_outcome`: Ship the smallest honest PASS-0001 control-plane service and tests now instead of widening into framework research.
- `event_boundary_notes`: Strong branch-collapsing intervention. PASS1 boundary is accurate.
- `human_model_signal`: Explicit boundedness rule: the current `.codex` corpus is small, schedule IDs can be simple, recent runs can stay lightweight, and the task should ship a bounded control plane rather than design a full framework.
- `failure_family_hypothesis`: `workflow_orchestration`
- `intervention_kind_hypothesis`: `collapse_branches`
- `human_cost_or_risk`: Time lost to open-ended research and delayed source edits despite a ready runtime and a named slice.
- `local_lesson_hypothesis`: Once the smallest honest slice is known and the runtime is already live, stop researching and ship that slice.
- `cluster_hints`: `boundedness`, `research-drift`, `proof-scope`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: none material

### C18 - Urgently contain the accidental release of all scheduled jobs

- `event_id`: `C18`
- `title`: `Urgently contain the accidental release of all scheduled jobs`
- `session_or_thread`: `Lead Task-0005 lifecycle`
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\06\rollout-2026-04-06T20-39-30-019d6561-529c-75e0-9a57-ab9d9fd66d80.jsonl`
- `primary_refs`: `T06:L1230`, `T06:L1249`
- `ai_course`: The assistant launched the worker-backed control plane and then started inspecting the resulting artifacts, even though startup had released all three scheduled Codex jobs at once instead of staying within the intended proof scope.
- `human_intervention`: The human supervisor injected a runtime containment summary, recorded the evidence that all three jobs had been released, stated that the task-owned runtime and child Codex processes had been stopped, and told the assistant to recover honestly from this as a concrete PASS-0002 control problem.
- `adequate_outcome`: Treat the unintended multi-run release as an urgent proof-scope and runtime-control failure, update durable state, and choose a bounded safe retry path before doing anything else.
- `event_boundary_notes`: Local boundary is clear. This is the first hard containment intervention after an unsafe proof launch.
- `human_model_signal`: Explicit proof-scope rule: usage must stay bounded to the intended proof, and an unsafe launch must not be treated as a normal continuation.
- `failure_family_hypothesis`: `workflow_orchestration`
- `intervention_kind_hypothesis`: `reject_outcome`
- `human_cost_or_risk`: Unintended release of all scheduled jobs, uncontrolled Codex usage, and dishonest proof claims if the launch were normalized.
- `local_lesson_hypothesis`: When a supposedly bounded proof releases live schedules, stop the runtime, record the failure honestly, and re-bound the proof before continuing.
- `cluster_hints`: `proof-scope-control`, `runtime-containment`, `durable-state-first`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: none material

### C19 - Remove the false money blocker and refocus on the real runtime blocker

- `event_id`: `C19`
- `title`: `Remove the false money blocker and refocus on the real runtime blocker`
- `session_or_thread`: `Lead Task-0005 lifecycle`
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\06\rollout-2026-04-06T20-39-30-019d6561-529c-75e0-9a57-ab9d9fd66d80.jsonl`
- `primary_refs`: `T06:L1302`, `T06:L1316`
- `ai_course`: The assistant had correctly identified the schedule-flood runtime problem, but it was still carrying stale spend-gate framing and a stale-shell execution model into its next-step story.
- `human_intervention`: The human supervisor explicitly removed the "additional money" blocker, identified the real blocker as the runtime behavior from the contained proof, pointed out the stale PATH issue, and required the assistant to ground further work in the actual local environment state.
- `adequate_outcome`: Drop the false cost blocker and make runtime behavior plus explicit executable resolution the operative state.
- `event_boundary_notes`: This sharpens the earlier `C11` permission-framing correction into a concrete blocker correction inside PASS-0002.
- `human_model_signal`: Explicit rule: once the real blocker is known, do not keep reporting a false blocker just because it used to be in the story.
- `failure_family_hypothesis`: `workflow_orchestration`
- `intervention_kind_hypothesis`: `reframe_problem`
- `human_cost_or_risk`: Debugging and state updates would be grounded in the wrong blocker model.
- `local_lesson_hypothesis`: After a failed proof attempt, aggressively delete stale blocker narratives and realign on the actual environment and runtime evidence.
- `cluster_hints`: `false-blocker`, `environment-truth`, `proof-scope-control`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: none material

### C20 - Enforce durable-state-first before more debugging

- `event_id`: `C20`
- `title`: `Enforce durable-state-first before more debugging`
- `session_or_thread`: `Lead Task-0005 lifecycle`
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\06\rollout-2026-04-06T20-39-30-019d6561-529c-75e0-9a57-ab9d9fd66d80.jsonl`
- `primary_refs`: `T06:L1334`, `T06:L1361`
- `ai_course`: The assistant said it was writing a durable checkpoint and then moving into the runtime fence fix, but the human supervisor still judged the task as violating the durable-state rule.
- `human_intervention`: The human stopped further exploration and explicitly ordered the durable writes first: honest blocked/debugging state in `TASK-STATE.json`, updated `HANDOFF.md`, and a bug record that captured both the schedule release and the `CreateProcessAsUserW failed: 1920` failure.
- `adequate_outcome`: Make the task files tell the truth on disk before any more debugging or retries.
- `event_boundary_notes`: Strong contract-enforcement event. PASS1 boundary is accurate.
- `human_model_signal`: Explicit adequacy rule: if the task has turned into a concrete defect investigation, durable artifacts must lead, not trail, the live state.
- `failure_family_hypothesis`: `workflow_orchestration`
- `intervention_kind_hypothesis`: `tighten_contract`
- `human_cost_or_risk`: Future work could proceed from stale or dishonest task artifacts, making handoff and review unsafe.
- `local_lesson_hypothesis`: When a bounded proof turns into a real defect, update durable state before continuing the debugging branch.
- `cluster_hints`: `durable-state-first`, `proof-scope-control`, `bug-truth`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: none material

### C21 - Inject real CLI semantics before the next executor change

- `event_id`: `C21`
- `title`: `Inject real CLI semantics before the next executor change`
- `session_or_thread`: `Lead Task-0005 lifecycle`
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\06\rollout-2026-04-06T20-39-30-019d6561-529c-75e0-9a57-ab9d9fd66d80.jsonl`
- `primary_refs`: `T06:L1528`, `T06:L1538`
- `ai_course`: The assistant was about to choose the next executor flag change based on its own current understanding after the failed bounded retry.
- `human_intervention`: The human supplied the exact `codex exec --help` semantics from the installed binary and pointed the assistant to the smallest explicit flag change that matched real CLI behavior.
- `adequate_outcome`: Ground the next executor change in the local binary's actual semantics rather than inference or memory.
- `event_boundary_notes`: Useful tool-truth injection, but smaller than the surrounding proof-scope incidents.
- `human_model_signal`: Explicit rule: use real installed CLI semantics, not guessed semantics, when changing executor behavior under failure.
- `failure_family_hypothesis`: `verification_proof`
- `intervention_kind_hypothesis`: `tighten_contract`
- `human_cost_or_risk`: The next retry could be based on a wrong flag model and waste the remaining bounded proof budget.
- `local_lesson_hypothesis`: Before changing executor flags in a failing runtime path, consult the actual local binary help and choose the smallest justified change.
- `cluster_hints`: `tool-truth`, `bounded-retry`, `diagnostic-scope`
- `accepted_incident_likelihood`: `unlikely`
- `confidence`: `medium`
- `uncertainties`: The event is likely too narrow for the accepted incident set even though the human had to inject authoritative tool truth.

### C22 - Force an explicit fork choice instead of more analysis drift

- `event_id`: `C22`
- `title`: `Force an explicit fork choice instead of more analysis drift`
- `session_or_thread`: `Lead Task-0005 lifecycle`
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\06\rollout-2026-04-06T20-39-30-019d6561-529c-75e0-9a57-ab9d9fd66d80.jsonl`
- `primary_refs`: `T06:L1631`, `T06:L1643`, `T06:L1651`
- `ai_course`: The assistant was still giving the bounded manual run a little more time and had not yet committed to either the final executor change or an honest blocker closeout.
- `human_intervention`: The human supervisor supplied the relevant evidence from the still-pending run and demanded an immediate explicit fork: one final bounded executor change or blocked closeout, with durable state either way.
- `adequate_outcome`: Stop analysis drift, choose the final bounded branch now, and preserve ownership through the result.
- `event_boundary_notes`: Strong branch-collapse event. PASS1 boundary is accurate.
- `human_model_signal`: Explicit boundedness rule: once the option space has collapsed and the evidence is sufficient, analysis must stop and the next durable choice must be made.
- `failure_family_hypothesis`: `workflow_orchestration`
- `intervention_kind_hypothesis`: `collapse_branches`
- `human_cost_or_risk`: Endless analysis while the task remained in limbo and the remaining bounded proof budget was being wasted.
- `local_lesson_hypothesis`: When only one bounded retry or an honest blocker closeout remains, choose explicitly instead of waiting for more drift.
- `cluster_hints`: `boundedness`, `proof-scope-control`, `analysis-drift`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: none material

### C23 - Force immediate durable closeout after the bounded proof was already successful

- `event_id`: `C23`
- `title`: `Force immediate durable closeout after the bounded proof was already successful`
- `session_or_thread`: `Lead Task-0005 lifecycle`
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\06\rollout-2026-04-06T20-39-30-019d6561-529c-75e0-9a57-ab9d9fd66d80.jsonl`
- `primary_refs`: `T06:L1853`, `T06:L1865`
- `ai_course`: The assistant acknowledged that the hard part was already proven but still wanted to give the process a short settle window and maybe open another bug if the process stayed pending.
- `human_intervention`: The human said the bounded proof already looked successful end to end and ordered immediate durable closeout: update task state and handoff, record the environment fixes honestly, commit, push, clean up runtime if appropriate, and report final status.
- `adequate_outcome`: Transition immediately into closeout once the proof bar is met.
- `event_boundary_notes`: Kept separate from `C24` because the human had to intervene again after this order.
- `human_model_signal`: Explicit adequacy rule: once the bounded proof has already demonstrated the target outcome, more settling time is not the right next move.
- `failure_family_hypothesis`: `workflow_orchestration`
- `intervention_kind_hypothesis`: `collapse_branches`
- `human_cost_or_risk`: Delayed closeout and continued ambiguity after the core closure condition was already satisfied.
- `local_lesson_hypothesis`: When the final bounded proof already clears the closure bar, switch to durable closeout immediately instead of waiting for perfect stillness.
- `cluster_hints`: `closure-truth`, `boundedness`, `analysis-drift`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: none material

### C24 - Stop nonessential follow-up and demand closeout-only status

- `event_id`: `C24`
- `title`: `Stop nonessential follow-up and demand closeout-only status`
- `session_or_thread`: `Lead Task-0005 lifecycle`
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\06\rollout-2026-04-06T20-39-30-019d6561-529c-75e0-9a57-ab9d9fd66d80.jsonl`
- `primary_refs`: `T06:L1916`, `T06:L1924`, `T06:L1930`
- `ai_course`: Even after the proof was confirmed, the assistant was still narrating evidence capture and artifact plans instead of purely executing the closeout order.
- `human_intervention`: The human stopped all nonessential follow-up, prescribed the exact remaining closeout order, and then required an immediate one-line status only.
- `adequate_outcome`: Closeout-only behavior and minimal response surface until the task is finished.
- `event_boundary_notes`: This is the second intervention because the earlier closeout push in `C23` had not fully stuck.
- `human_model_signal`: Explicit rule: after success is verified and closeout order is given, explanatory churn is itself the wrong move.
- `failure_family_hypothesis`: `workflow_orchestration`
- `intervention_kind_hypothesis`: `collapse_branches`
- `human_cost_or_risk`: More churn and delay in a task that should have been closing out.
- `local_lesson_hypothesis`: After a closeout-only order, keep the response surface minimal and finish the ordered steps rather than continuing to narrate.
- `cluster_hints`: `closeout-discipline`, `boundedness`, `analysis-drift`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: none material

### C25 - Redirect the debug branch to the surviving `Head` hierarchy discrepancy

- `event_id`: `C25`
- `title`: `Redirect the debug branch to the surviving Head hierarchy discrepancy`
- `session_or_thread`: `Trace hierarchy mismatch`
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\06\rollout-2026-04-06T20-43-54-019d6565-591e-7831-b050-55f12d0b304f.jsonl`
- `primary_refs`: `T07:L375`, `T07:L382`
- `ai_course`: The assistant had a solid checkpoint around the now-solved extra-root defect and was prepared to leave the branch resumable there.
- `human_intervention`: The human narrowed the branch to the current promoted first-pass lane and told the assistant to chase the earliest remaining deterministic divergence upstream of the live `Head` hierarchy mismatch.
- `adequate_outcome`: Move off the solved historical defect and onto the current surviving seam.
- `event_boundary_notes`: PASS1 boundary is accurate. This is a seam-redirection event, not a closure claim.
- `human_model_signal`: Explicit debugging rule: stay on the current failing lane and target the earliest remaining deterministic disagreement, not the already-solved older defect.
- `failure_family_hypothesis`: `verification_proof`
- `intervention_kind_hypothesis`: `redirect_debugging`
- `human_cost_or_risk`: Debug effort could stay attached to a solved cause while the live lane still fails.
- `local_lesson_hypothesis`: Once one root cause is repaired, re-anchor debugging on the earliest still-live divergence rather than lingering on the solved seam.
- `cluster_hints`: `wrong-seam-debugging`, `earliest-deterministic-divergence`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: none material

### C26 - Redirect again when the first narrowed explanation still was not upstream enough

- `event_id`: `C26`
- `title`: `Redirect again when the first narrowed explanation still was not upstream enough`
- `session_or_thread`: `Trace hierarchy mismatch`
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\06\rollout-2026-04-06T20-43-54-019d6565-591e-7831-b050-55f12d0b304f.jsonl`
- `primary_refs`: `T07:L599`, `T07:L606`
- `ai_course`: The assistant had a plausible explanation for the surviving `Head` warning: the retained Manny Control Rig hierarchy inside the first-pass anim blueprint.
- `human_intervention`: The human said the live seam to investigate was farther upstream: the target IKRig forcing the `Spine` chain back to `Hips`, despite logs showing `Spine -> Spine2`, and ordered the assistant to stop only when it could name the exact structural cause.
- `adequate_outcome`: Continue debugging until the first structural cause upstream of the saved `Spine -> Hips` coercion is named.
- `event_boundary_notes`: Distinct from `C25`. `C25` moved off the old solved defect; `C26` tightens the seam again because the first narrowed explanation was still downstream.
- `human_model_signal`: Explicit debugging adequacy rule: a first plausible surviving warning producer is not enough if it is not yet the exact structural fact forcing the bad saved state.
- `failure_family_hypothesis`: `verification_proof`
- `intervention_kind_hypothesis`: `redirect_debugging`
- `human_cost_or_risk`: The branch could stop at a downstream producer and miss the actual structural cause.
- `local_lesson_hypothesis`: In structural debugging, keep drilling upstream until you can name the exact fact that forces the bad saved configuration, not just the first live warning producer.
- `cluster_hints`: `wrong-seam-debugging`, `earliest-deterministic-divergence`, `upstream-cause`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: The transcript does not by itself prove whether the IKRig coercion is the only remaining cause of full Lane C collapse; it only establishes that the human judged the Manny Control Rig explanation as still downstream.

## Likely accepted incidents

- `C01` - Re-scope away from Task-0004 closure
- `C02` - Make the Home progress bar mean time-to-done
- `C03` - Call out the remaining disconnect as a dropped ball
- `C04` - Re-ground Task-0022 planning with the clarified no-bar rule
- `C06` - Correct the implied expectation around restart and force a bounded manual repair
- `C07` - Reject the "phone is repaired" claim because uploads were still stalled
- `C08` - Reject the "optimization problem" framing and require durable notes plus charity
- `C12` - Reject human-world closure where the backend is not actually left running
- `C13` - Force creation of the separate always-on service lane
- `C14` - Call out the dropped ball on the dashboard surface itself
- `C17` - Bound PASS-0001 and shut down open-ended research
- `C18` - Urgently contain the accidental release of all scheduled jobs
- `C19` - Remove the false money blocker and refocus on the real runtime blocker
- `C20` - Enforce durable-state-first before more debugging
- `C22` - Force an explicit fork choice instead of more analysis drift
- `C23` - Force immediate durable closeout after the bounded proof was already successful
- `C24` - Stop nonessential follow-up and demand closeout-only status
- `C25` - Redirect the debug branch to the surviving Head hierarchy discrepancy
- `C26` - Redirect again when the first narrowed explanation still was not upstream enough

## Likely non-incident but still important intervention events

- `C05` - A real push from source-level fix to live-device validation, but probably too small for the accepted incident set on its own.
- `C09` - A useful upgrade from subjective regression note to objective numbers, but probably better treated as evidence-discipline tightening than as a standalone accepted incident.
- `C10` - Boundary-corrected gate-clear event rather than the broader meta-coaching summary PASS1 originally attached to it.
- `C15` - Clear control-boundary correction on restart ownership, but likely best carried as a local example rather than a standalone accepted incident.
- `C16` - Real reactivation of stalled momentum, but probably too generic for the accepted set.
- `C21` - Strong tool-truth injection before a retry, but likely too narrow for the accepted set.

## Repeated cluster hints noticed across the analyzed set

- `status-signal-meaning`
  - `C02`, `C03`, `C04`, `C14`
- `real-world-completion-mismatch`
  - `C06`, `C07`, `C12`, `C13`, `C23`, `C24`
- `operator-boundary`
  - `C12`, `C13`, `C14`, `C15`
- `proof-scope-control`
  - `C05`, `C17`, `C18`, `C19`, `C20`, `C22`, `C23`, `C24`
- `user-report-as-evidence`
  - `C08`, `C09`
- `wrong-seam-debugging`
  - `C25`, `C26`

## Strongest human-model signals worth carrying into later clustering or principle work

- `C02` and `C04`: A Home progress bar means "how long until done." If that promise is not true or useful, use calmer text or chips instead.
- `C08`: A user-reported regression after the assistant's change is first-class evidence and should be handled with charity, not minimized into "just optimization."
- `C12` and `C13`: The human is not the operator. "Done" for a long-running backend means a usable always-on service lane remains running on the local computer, distinct from disposable validation infrastructure.
- `C14`: Meaning for a human-facing control surface belongs on the surface itself, not buried in docs or CLI explanations.
- `C20`: Once a bounded proof becomes a concrete defect, durable task state must be updated before more debugging.
- `C22`, `C23`, and `C24`: When the option space or closure bar is already clear, stop analysis drift and move into the bounded next action or closeout.
- `C25` and `C26`: Debugging should stay on the earliest remaining deterministic divergence upstream, not on solved history or the first merely plausible downstream producer.

## Events that still need a wider reread

- `C10` only, and only if a later pass wants to recover the broader "infer likely human meaning so the human intervenes less" coaching theme. The cited PASS1 refs do not capture that broader statement; they capture only the gate-clear event.
- No other candidate needs a wider reread for local PASS2 classification. The remaining boundaries are sufficiently clear from the reopened transcript windows.
