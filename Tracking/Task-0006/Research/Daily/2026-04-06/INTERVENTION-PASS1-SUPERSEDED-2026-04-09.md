# INTERVENTION-PASS1

Source day: `2026-04-06`

## Source scope reviewed

- Primary source scope only: `C:\Users\gregs\.codex\sessions\2026\04\06\rollout-*.jsonl`
- Broad scan completed across all 16 day-scoped rollout files at user-message index level.
- Narrow rereads used for exact extraction:
- `T01` = `C:\Users\gregs\.codex\sessions\2026\04\06\rollout-2026-04-06T18-59-50-019d6506-147a-7c00-972d-2498ff61ff81.jsonl`
- `T02` = `C:\Users\gregs\.codex\sessions\2026\04\06\rollout-2026-04-06T19-28-46-019d6520-919c-7b60-94c2-422680147d6c.jsonl`
- `T03` = `C:\Users\gregs\.codex\sessions\2026\04\06\rollout-2026-04-06T19-30-25-019d6522-1369-7692-bdba-b53998c88601.jsonl`
- `T04` = `C:\Users\gregs\.codex\sessions\2026\04\06\rollout-2026-04-06T19-54-03-019d6537-b945-7173-888e-667687b4d67c.jsonl`
- `T05` = `C:\Users\gregs\.codex\sessions\2026\04\06\rollout-2026-04-06T20-36-54-019d655e-f1c5-7463-af8a-779567b8955a.jsonl`
- `T06` = `C:\Users\gregs\.codex\sessions\2026\04\06\rollout-2026-04-06T20-39-30-019d6561-529c-75e0-9a57-ab9d9fd66d80.jsonl`
- `T07` = `C:\Users\gregs\.codex\sessions\2026\04\06\rollout-2026-04-06T20-43-54-019d6565-591e-7831-b050-55f12d0b304f.jsonl`
- Chronology note: some same-day transcripts reference local-time events on `2026-04-07`. This pass stays ordered by transcript position inside the `2026-04-06` source-day files.

## Total candidate intervention events found

- `26`

## Chronological candidate list

### C01 - Re-scope away from Task-0004 closure

- Session: `Harvest next task candidates`
- Refs: `T01:L223`, `T01:L230`
- AI course: Task-0004 was framed as the only real open task and the next required move.
- Human intervention: the human closed Task-0004 as out of scope, pivoted to the new backend/orchestration direction, and inserted a Chrome-export prerequisite for Task-0005.
- Better outcome forced: stop optimizing for Task-0004 closure and bootstrap the new workstream instead.
- Why real: active task and closure bar changed.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C02 - Make the Home progress bar mean time-to-done

- Session: `Implement task 22 tracking`
- Refs: `T02:L235`, `T02:L242`
- AI course: the plan had tightened, but the explanation still left room for a generic `busy` meaning.
- Human intervention: the human stated the decisive rule that a progress bar means `how long until done`.
- Better outcome forced: the plan must treat a bar as a remaining-time promise, not generic activity.
- Why real: this tightened the UI contract materially.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C03 - Call out the remaining disconnect as a dropped ball

- Session: `Implement task 22 tracking`
- Refs: `T02:L290`, `T02:L297`
- AI course: the assistant said the rule was encoded, but its explanation still sounded ambiguous.
- Human intervention: the human said this counted as a `dropped ball` because they still had to follow up and ask where the disconnect was.
- Better outcome forced: explain the corrected rule clearly enough that the human does not need another repair pass.
- Why real: explicit follow-up burden caused by inadequate explanation.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C04 - Re-ground Task-0022 planning with the clarified no-bar rule

- Session: `Own Task-0022 lifecycle`
- Refs: `T03:L245`, `T03:L252`
- AI course: Task-0022 was back at the approval gate, but the new semantics were not yet explicit enough in the planning baseline.
- Human intervention: the supervisory prompt ordered the owner to re-ground and update planning artifacts so safe state never shows a bar and retained local copies do not look like remaining transfer time.
- Better outcome forced: a truthful planning baseline before approval.
- Why real: the human had to force the clarified design contract into the durable artifacts.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C05 - Force the move from source-level fix to live phone patching

- Session: `Run TASK-HARVESTER`
- Refs: `T04:L644`, `T04:L651`
- AI course: the assistant had root cause, tests, build output, and docs, but had not yet installed the patch to the actual phone.
- Human intervention: `Go ahead and patch the phone app and restart.`
- Better outcome forced: on-device proof, not a source-only repair story.
- Why real: the human had to push the agent past an early stop point.
- Confidence: `medium`
- Triage: `intervention event but probably not an accepted incident`

### C06 - Correct the implied expectation around restart and force a bounded manual repair

- Session: `Run TASK-HARVESTER`
- Refs: `T04:L709`, `T04:L716`, `T04:L725`, `T04:L748`, `T04:L755`
- AI course: the assistant installed and restarted the patched app, but its earlier wording had implied that install plus restart might clear the blocked state.
- Human intervention: the human pushed on that implication, asked whether the assistant was surprised or had failed to restart, then authorized a careful write limited to the exact 23 rows.
- Better outcome forced: an honest explanation of what restart could and could not fix, followed by a narrowly bounded repair.
- Why real: the human corrected a misleading implication and tightened the safety bar.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C07 - Reject the `phone is repaired` claim because uploads were still stalled

- Session: `Run TASK-HARVESTER`
- Refs: `T04:L1012`, `T04:L1019`
- AI course: the assistant said `The phone is repaired` and framed the remaining issue as prevention rather than current behavior.
- Human intervention: the human reported the actual surface still said `waiting to upload` with roughly 3400 clips remaining and ordered the agent to keep fixing until it was truly fixed.
- Better outcome forced: real upload recovery, not a premature closure based on one repaired symptom.
- Why real: the human rejected the assistant's adequacy claim on direct real-world evidence.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C08 - Reject the `optimization problem` framing and require durable notes plus charity

- Session: `Run TASK-HARVESTER`
- Refs: `T04:L1431`, `T04:L1438`
- AI course: the assistant said the queue looked healthy and the remaining issue was throughput optimization, not the stuck-upload bug.
- Human intervention: the human said it had been much faster before the assistant's fixes, insisted that their report be recorded durably, called out the dropped ball, and asked for a more charitable framing.
- Better outcome forced: treat the slowdown as live regression evidence, not as a dismissible optimization follow-up.
- Why real: the human had to correct the assistant's framing of the problem.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C09 - Force objective speed analysis before disconnect

- Session: `Run TASK-HARVESTER`
- Refs: `T04:L1450`, `T04:L1457`
- AI course: the assistant had only recorded the user's regression report in the task docs.
- Human intervention: the human required a numbers-backed speed analysis from the captured logs and asked the bug to be updated while they disconnected the phone.
- Better outcome forced: objective comparison data, not just a note that it felt slower.
- Why real: the human had to ask for a stronger evidentiary standard.
- Confidence: `medium`
- Triage: `intervention event but probably not an accepted incident`

### C10 - Meta-correct the inference gap at the Task-0005 plan gate

- Session: `Implement task 5 state`
- Refs: `T05:L116`, `T05:L123`
- AI course: the assistant treated the written approval gate as the full contract and stopped there.
- Human intervention: the human explained that the more the agent can infer likely human meaning, the less the human has to intervene and chase agents.
- Better outcome forced: use implied intent more aggressively instead of only literal artifact text.
- Why real: it explicitly reframes the adequacy bar for future behavior.
- Confidence: `medium`
- Triage: `unclear and needs transcript reread`

### C11 - Correct the spend-gate framing on the already-authenticated Codex CLI

- Session: `Implement task 5 state`
- Refs: `T05:L1297`, `T05:L1304`, `T05:L1307`, `T05:L1320`
- AI course: the assistant described a `spend gate` around one live `codex exec` run using the already logged-in account.
- Human intervention: the human pushed for a clearer explanation and said there was no `additional money` issue because the assistant was already using their subscription.
- Better outcome forced: stop treating this as fresh payment collection and frame it as permission to use the existing authenticated account.
- Why real: the human had to repair the gating model the assistant was using.
- Confidence: `medium`
- Triage: `likely accepted incident candidate`

### C12 - Reject human-world closure where the backend is not actually left running

- Session: `Implement task 5 state`
- Refs: `T05:L2641`, `T05:L2648`, `T05:L2650`, `T05:L2657`
- AI course: the assistant had closed Task-0005 after proving the backend and then tearing down the runtime.
- Human intervention: the human said that was not what a real human would interpret as `done`, stressed that they were not the operator, and explained that scheduled jobs needed to keep running on the local computer.
- Better outcome forced: a usable always-on system, not merely validated code plus cleanup.
- Why real: the human explicitly rejected the assistant's closure interpretation on human-facing outcome grounds.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C13 - Force creation of the separate always-on service lane

- Session: `Implement task 5 state`
- Refs: `T05:L2660`, `T05:L2667`
- AI course: the assistant articulated the higher-level correction but had not yet built the durable service lane implied by it.
- Human intervention: the human told it to make that service lane now and document everything according to the new understanding.
- Better outcome forced: actual persistent runtime and updated artifacts, not just reflection.
- Why real: the human had to convert self-critique into concrete implementation work.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C14 - Call out the dropped ball on the dashboard surface itself

- Session: `Implement task 5 state`
- Refs: `T05:L3020`, `T05:L3027`
- AI course: the assistant answered `Force Reconcile` versus `Refresh` using backend/docs truth but had not updated the dashboard or proactively offered that fix.
- Human intervention: the human called this `another dropped ball` and said humans should not need docs or CLI for that surface.
- Better outcome forced: fix the human-facing UI, not just provide accurate documentation.
- Why real: the human had to redirect the assistant from backend-explanation sufficiency to surface adequacy.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C15 - Force the assistant to restart the app instead of handing that burden back

- Session: `Implement task 5 state`
- Refs: `T05:L3055`, `T05:L3062`
- AI course: after patching the UI copy, the assistant told the human to restart the dashboard to see the change.
- Human intervention: `You restart it; you should know I don't have a restart button.`
- Better outcome forced: the agent owns the restart step on the human's behalf.
- Why real: the human had to reject needless operator burden on a human-facing surface.
- Confidence: `medium`
- Triage: `intervention event but probably not an accepted incident`

### C16 - Nudge the task leader out of silence and back into concrete work

- Session: `Lead Task-0005 lifecycle`
- Refs: `T06:L692`, `T06:L718`
- AI course: after the live runtime/tooling context was ready, the assistant had gone quiet and there were no repo writes past setup.
- Human intervention: the human asked for the current blocker or next concrete implementation step and told the agent to continue immediately.
- Better outcome forced: visible implementation momentum instead of silent drift.
- Why real: the human had to re-activate a stalled course.
- Confidence: `medium`
- Triage: `intervention event but probably not an accepted incident`

### C17 - Bound PASS-0001 and shut down open-ended research

- Session: `Lead Task-0005 lifecycle`
- Refs: `T06:L724`, `T06:L748`
- AI course: the assistant was still in SDK or method lookup mode after saying it would move into code.
- Human intervention: the human bounded the exact PASS-0001 slice and said `No more open-ended SDK/doc research`.
- Better outcome forced: ship the smallest honest PASS-0001 now.
- Why real: the human had to stop an inadequate research-heavy course after implementation was already the right next move.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C18 - Urgently contain the accidental release of all scheduled jobs

- Session: `Lead Task-0005 lifecycle`
- Refs: `T06:L1230`, `T06:L1249`
- AI course: worker-backed startup released three schedule-triggered jobs when the intended proof was supposed to be bounded.
- Human intervention: the supervisor injected runtime containment evidence and reframed the situation as an urgent control problem, not a normal proof continuation.
- Better outcome forced: contain the unintended releases and debug the schedule semantics before doing more.
- Why real: the human had to interrupt an actively unsafe course.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C19 - Remove the false money blocker and refocus on the real runtime blocker

- Session: `Lead Task-0005 lifecycle`
- Refs: `T06:L1302`, `T06:L1316`
- AI course: the assistant was still talking as if spend approval could be the blocker, while also running with stale PATH assumptions.
- Human intervention: the supervisor explicitly said there was no separate additional-money gate, identified runtime behavior as the real blocker, and pointed the agent at explicit user-observed paths.
- Better outcome forced: drop the wrong blocker and ground decisions in the real environment state.
- Why real: the human had to correct both the gating model and the execution environment assumptions.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C20 - Enforce durable-state-first before more debugging

- Session: `Lead Task-0005 lifecycle`
- Refs: `T06:L1334`, `T06:L1361`
- AI course: the assistant planned to checkpoint but was still poised to keep investigating while the durable state did not yet fully reflect the new blocker reality.
- Human intervention: the supervisor stopped deeper exploration and required immediate TASK-STATE, HANDOFF, and bug updates first.
- Better outcome forced: durable truth on disk before any more debugging or retries.
- Why real: the human had to prevent the task artifacts from lagging behind the actual state.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C21 - Inject real CLI semantics before the next executor change

- Session: `Lead Task-0005 lifecycle`
- Refs: `T06:L1528`, `T06:L1538`
- AI course: the assistant was about to choose a command-plan change after its own local checks.
- Human intervention: the human supplied the exact `codex exec --help` semantics from the installed binary and told the agent to choose the smallest explicit flag change from real CLI behavior.
- Better outcome forced: flag choice grounded in the actual local binary, not inference or guesswork.
- Why real: the human had to inject authoritative tool truth to keep the next retry from drifting.
- Confidence: `medium`
- Triage: `intervention event but probably not an accepted incident`

### C22 - Force an explicit fork choice instead of more analysis drift

- Session: `Lead Task-0005 lifecycle`
- Refs: `T06:L1631`, `T06:L1643`, `T06:L1651`
- AI course: the assistant was still giving the failed bounded run a little more time before deciding whether to escalate or stop.
- Human intervention: the supervisor injected exact failure evidence and demanded an immediate durable choice: one final stronger flag or an honest blocker closeout.
- Better outcome forced: a bounded decision now, not more observation drift.
- Why real: the human had to stop indecision in the face of already-sufficient evidence.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C23 - Force immediate durable closeout after the bounded proof was already successful

- Session: `Lead Task-0005 lifecycle`
- Refs: `T06:L1853`, `T06:L1865`
- AI course: the assistant was still giving the successful run a short settle window and considering whether a remaining issue might need to be written as a bug.
- Human intervention: the supervisor said the final bounded proof already looked successful end to end and ordered immediate durable closeout.
- Better outcome forced: closeout work immediately once success evidence was already present.
- Why real: the human had to override a lingering wait-and-see posture after the core closure condition was already satisfied.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C24 - Stop nonessential follow-up and demand closeout-only status

- Session: `Lead Task-0005 lifecycle`
- Refs: `T06:L1916`, `T06:L1924`, `T06:L1930`
- AI course: the assistant was still narrating evidence capture and artifact updates after success was already confirmed.
- Human intervention: the supervisor ordered all nonessential follow-up to stop, prescribed the exact remaining closeout order, and then demanded a one-line status only.
- Better outcome forced: finish closure instead of continuing explanatory churn.
- Why real: the human had to intervene again because the earlier closeout push had not fully stuck.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C25 - Redirect the debug branch to the surviving `Head` hierarchy discrepancy

- Session: `Trace hierarchy mismatch`
- Refs: `T07:L375`, `T07:L382`
- AI course: the assistant had produced a solid checkpoint around the old extra-root defect and was ready to leave the branch resumable there.
- Human intervention: the human narrowed the next bounded branch to the current first-pass promoted lane and the live `Head` hierarchy mismatch that still collapsed Lane C.
- Better outcome forced: move from the now-solved old defect to the earliest remaining deterministic divergence.
- Why real: the human had to redirect debugging away from a solved seam and onto the still-relevant one.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C26 - Redirect again when the first narrowed explanation still was not upstream enough

- Session: `Trace hierarchy mismatch`
- Refs: `T07:L599`, `T07:L606`
- AI course: the assistant had identified Manny Control Rig hierarchy as the earliest remaining confirmed producer of the live `Head` warning.
- Human intervention: the human said the real seam to investigate was the target IKRig forcing `Spine` back to `Hips` and asked for the exact first-pass structural cause upstream.
- Better outcome forced: continue drilling until the structural cause of the retarget coercion is named, not just the first plausible surviving warning producer.
- Why real: the human had to redirect the debug seam again because the first correction did not yet land on the right upstream cause.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

## Which candidates look like likely accepted incidents

- `C01` Re-scope away from Task-0004 closure
- `C02` Make the Home progress bar mean time-to-done
- `C03` Call out the remaining disconnect as a dropped ball
- `C04` Re-ground Task-0022 planning with the clarified no-bar rule
- `C06` Correct the implied expectation around restart and force a bounded manual repair
- `C07` Reject the `phone is repaired` claim because uploads were still stalled
- `C08` Reject the `optimization problem` framing and require durable notes plus charity
- `C11` Correct the spend-gate framing on the already-authenticated Codex CLI
- `C12` Reject human-world closure where the backend is not actually left running
- `C13` Force creation of the separate always-on service lane
- `C14` Call out the dropped ball on the dashboard surface itself
- `C17` Bound PASS-0001 and shut down open-ended research
- `C18` Urgently contain the accidental release of all scheduled jobs
- `C19` Remove the false money blocker and refocus on the real runtime blocker
- `C20` Enforce durable-state-first before more debugging
- `C22` Force an explicit fork choice instead of more analysis drift
- `C23` Force immediate durable closeout after the bounded proof was already successful
- `C24` Stop nonessential follow-up and demand closeout-only status
- `C25` Redirect the debug branch to the surviving `Head` hierarchy discrepancy
- `C26` Redirect again when the first narrowed explanation still was not upstream enough

## Which candidates are real interventions but probably belong outside the accepted incident set

- `C05` Force the move from source-level fix to live phone patching
- `C09` Force objective speed analysis before disconnect
- `C15` Force the assistant to restart the app instead of handing that burden back
- `C16` Nudge the task leader out of silence and back into concrete work
- `C21` Inject real CLI semantics before the next executor change

## Any ambiguous boundaries that need a second read

- `C10` is a real meta-correction, but it sits on the boundary between a durable intervention event and a general coaching statement. It likely matters, but it would benefit from a fuller reread of the surrounding planning turns.
- Possible weak extra candidate not promoted in this pass: `T01:L373` to `T01:L380`, where the assistant said it lacked a concrete target and the human clarified `Cron System with JSON`. The intervention boundary looked weak in this pass.
- Possible weak extra candidate not promoted in this pass: `T02:L188` to `T02:L195`, where the human first asked why a progress bar would be used for small uploads before the stronger `time-to-done` correction landed.
- Possible weak extra candidate not promoted in this pass: `T07:L350` to `T07:L372`, where the human paused the branch for a status checkpoint before the later, clearer seam redirections at `C25` and `C26`.
