# INTERVENTION-PASS1

Source day: `2026-04-06`

Canonical note: promoted on `2026-04-09` from the explanatory rerun; the earlier day-local PASS1 is archived as `INTERVENTION-PASS1-SUPERSEDED-2026-04-09.md`.

## Source scope reviewed

- Source day: `2026-04-06`
- Transcript scope: all `rollout-*.jsonl` files under `C:\Users\gregs\.codex\sessions\2026\04\06` (`16` files total).
- Contracts consulted before reread:
  - `C:\Users\gregs\.codex\Orchestration\Reports\Incidents\README.md`
  - `C:\Users\gregs\.codex\Orchestration\Reports\Incidents\INCIDENT.schema.json`
  - `C:\Users\gregs\.codex\Orchestration\Prompts\INTERVENTION-PASS1.md`
  - `C:\Users\gregs\.codex\Orchestration\Prompts\INTERVENTION-PASS2.md`
  - `C:\Users\gregs\.codex\skills\codex-session-search\SKILL.md`
- Existing April 6 pass artifacts were not used as evidence for candidate discovery.
- Read strategy: whole-file chronological reread with emphasis on `event_msg` `user_message` and `agent_message` turns, then bounded rereads around each suspected correction arc.

## Transcript index used below

- `T01` `Implement task 22 tracking`
  - `C:\Users\gregs\.codex\sessions\2026\04\06\rollout-2026-04-06T19-28-46-019d6520-919c-7b60-94c2-422680147d6c.jsonl`
- `T02` `Own Task-0022 lifecycle`
  - `C:\Users\gregs\.codex\sessions\2026\04\06\rollout-2026-04-06T19-30-25-019d6522-1369-7692-bdba-b53998c88601.jsonl`
- `T03` `Run TASK-HARVESTER`
  - `C:\Users\gregs\.codex\sessions\2026\04\06\rollout-2026-04-06T19-54-03-019d6537-b945-7173-888e-667687b4d67c.jsonl`
- `T04` `Implement task 5 state`
  - `C:\Users\gregs\.codex\sessions\2026\04\06\rollout-2026-04-06T20-36-54-019d655e-f1c5-7463-af8a-779567b8955a.jsonl`
- `T05` `Lead Task-0005 lifecycle`
  - `C:\Users\gregs\.codex\sessions\2026\04\06\rollout-2026-04-06T20-39-30-019d6561-529c-75e0-9a57-ab9d9fd66d80.jsonl`
- `T06` `Trace divergence root cause`
  - `C:\Users\gregs\.codex\sessions\2026\04\06\rollout-2026-04-06T20-42-01-019d6563-9fe8-7741-ad48-79434c5386d7.jsonl`
- `T07` `Trace hierarchy mismatch`
  - `C:\Users\gregs\.codex\sessions\2026\04\06\rollout-2026-04-06T20-43-54-019d6565-591e-7831-b050-55f12d0b304f.jsonl`

## Total candidate intervention events found

- Total candidates: `19`
- Strong: `10`
- Medium: `9`
- Weak / ambiguous promoted as standalone candidates: `0`

## Strong candidates

### C02 - Make the Home progress bar mean time-to-done

- When / session: `2026-04-06T23:47Z`, `T01` plus relay into `T02`
- Primary refs: `T01#L242`, `T01#L245`, `T02#L252`, `T02#L255`
- AI course: the planning baseline still left room for a Home progress bar to act like generic upload activity.
- Human intervention: the human stated the decisive rule that if Home shows a progress bar, humans will read it as `how long until done`, and the long-lived owner had to be told to encode that rule explicitly.
- Better outcome forced: the plan had to reserve Home bars for sustained active uploads where remaining time is honest, and keep retained-local-copy state out of progress treatment.
- Why this is real: this was not routine preference input; it corrected a concrete plan already at the approval gate and required durable plan updates before implementation could proceed.
- Confidence: `strong`
- PASS1 read: `likely accepted incident candidate`

### C03 - Call out the remaining disconnect as a dropped ball

- When / session: `2026-04-06T23:50Z`, `T01`
- Primary refs: `T01#L297`, `T01#L300`, `T01#L310`
- AI course: after the time-to-done correction, the assistant still described the bar too much like a generic `I'm busy` signal and had not made the concrete alternative explicit enough.
- Human intervention: the human said this counted as a `dropped ball` because they still had to follow up and ask where the disconnect was.
- Better outcome forced: the assistant had to explicitly define the bar as a time-to-done promise and spell out the non-bar alternatives for short upload, safe state, and retained-local-copy state.
- Why this is real: the first correction did not fully stick, so the human had to reteach the semantic rule and the concrete UI alternatives.
- Confidence: `strong`
- PASS1 read: `likely accepted incident candidate`

### C05 - Record the user's regression report charitably instead of flattening it into optimization

- When / session: `2026-04-07T01:44Z`, `T03`
- Primary refs: `T03#L1438`, `T03#L1441`, `T03#L1450`
- AI course: the assistant treated slower upload drain after a fix as `a throughput optimization problem now`, even after the human had reported that it used to be much faster before the fix.
- Human intervention: the human said it was not simply an optimization problem, insisted that their report be recorded durably, called out the dropped ball, and explicitly asked for a principle of charity.
- Better outcome forced: the user report had to be treated as first-class regression evidence rather than explained away.
- Why this is real: the human had to spend effort teaching both evidentiary standards and conversational stance so they would not need to fight to have their lived regression report taken seriously.
- Confidence: `strong`
- PASS1 read: `likely accepted incident candidate`

### C07 - Resume deterministic root-cause pursuit instead of stopping at a plausible upstream producer

- When / session: `2026-04-07T01:48Z`, `T06`
- Primary refs: `T06#L1286`, `T06#L1293`, `T06#L1296`
- AI course: the assistant stopped after isolating an upstream producer and shifted into handoff mode even though the standing instruction was to keep tracing until the exact structural root cause was explicit.
- Human intervention: the human asked `Why did you stop working?` and restated the deterministic root-cause instruction.
- Better outcome forced: continue tracing through the authoring data rather than treating the first plausible upstream producer as the endpoint.
- Why this is real: this is a direct autonomy-standard correction, not ordinary steering. The assistant had already declared a stopping point; the human had to override that stopping rule.
- Confidence: `strong`
- PASS1 read: `likely accepted incident candidate`

### C10 - Contain an unsafe worker startup that released all scheduled jobs

- When / session: `2026-04-07T14:32Z`, `T05`
- Primary refs: `T05#L1249`, `T05#L1252`
- AI course: the assistant started the worker-backed control plane for one proof run, which released all three schedule-triggered jobs at once instead of staying within proof scope.
- Human intervention: the supervisor contained the runtime, killed the spawned jobs, documented the evidence, and explicitly redirected the assistant to recover honestly as a concrete bug or blocker.
- Better outcome forced: treat the launch as an unsafe runtime event, not as acceptable proof progress.
- Why this is real: the human had to actively stop an unsafe course and reframe it as a bounded bug-handling path.
- Confidence: `strong`
- PASS1 read: `likely accepted incident candidate`

### C12 - Enforce durable-state-first before more debugging

- When / session: `2026-04-07T14:35Z`, `T05`
- Primary refs: `T05#L1361`, `T05#L1364`
- AI course: after the unsafe launch, the assistant was still moving into deeper runtime exploration without first making the durable task state honestly blocked/debugging with the new bug recorded.
- Human intervention: the human explicitly stopped further exploration and prescribed the exact durable-state updates that had to happen first.
- Better outcome forced: truthful blocked/debugging state, updated handoff, and concrete bug record before any more investigation.
- Why this is real: this is not just documentation polish. The human had to reteach the task-lifecycle rule that pass artifacts and state must stay honest before deeper work continues.
- Confidence: `strong`
- PASS1 read: `likely accepted incident candidate`

### C16 - Reject the proof-only closure model for a `long-running backend`

- When / session: `2026-04-07T16:13Z`, `T04`
- Primary refs: `T04#L2638`, `T04#L2641`, `T04#L2648`, `T04#L2657`, `T04#L2660`
- AI course: the assistant had closed `Task-0005` after building and proving the backend, but had shut the runtime down and left the human with no actual always-on scheduler.
- Human intervention: the human said that was not what a human would interpret as `done`, stressed `I'm not trying to be an operator here`, and explained that the local machine needed an unattended scheduled-jobs lane.
- Better outcome forced: closure truth had to include a real always-on operating model, not just proof that the code could run.
- Why this is real: the human had to explain higher-level meaning, operator boundaries, and implied product expectations so the assistant would stop optimizing for proof-only closure.
- Confidence: `strong`
- PASS1 read: `likely accepted incident candidate`

### C17 - Force creation of the separate always-on service lane

- When / session: `2026-04-07T16:19Z`, `T04`
- Primary refs: `T04#L2667`, `T04#L2670`
- AI course: after the closure-model correction, the assistant had articulated the higher-level lesson but had not yet turned it into the actual service lane the new model implied.
- Human intervention: the human told it to make that service lane now and document the task according to the new understanding.
- Better outcome forced: convert the corrected operating model into a concrete two-lane implementation, not just a verbal admission.
- Why this is real: the human had to intervene again because the assistant had reached the lesson without yet taking ownership of the concrete fix.
- Confidence: `strong`
- PASS1 read: `possible accepted incident candidate`

### C18 - Call out the dropped ball on the dashboard surface itself

- When / session: `2026-04-07T18:00Z`, `T04`
- Primary refs: `T04#L3027`, `T04#L3030`, `T04#L3055`
- AI course: the assistant correctly explained `Refresh` versus `Force Reconcile` using backend/docs truth, but had not updated the dashboard itself or proactively offered that human-facing fix.
- Human intervention: the human called it `another dropped ball`, said humans should not have to read docs or CLI for that surface, and said the form-factor burden rests with the assistant.
- Better outcome forced: put the meaning on the dashboard surface itself instead of leaving the human to reconstruct it from docs or backend knowledge.
- Why this is real: the human had to teach a durable human-surface rule, not just ask for copy changes.
- Confidence: `strong`
- PASS1 read: `likely accepted incident candidate`

### C19 - Reject the restart handoff back to the human

- When / session: `2026-04-07T18:03Z`, `T04`
- Primary refs: `T04#L3062`, `T04#L3065`
- AI course: after fixing the UI copy, the assistant told the human to restart the dashboard if they wanted to see it.
- Human intervention: the human replied `You restart it; you should know I don't have a restart button.`
- Better outcome forced: the assistant had to own the restart so the corrected surface was actually running.
- Why this is real: this is a small but concrete operator-boundary correction on a human-facing surface.
- Confidence: `strong`
- PASS1 read: `possible accepted incident candidate`

## Medium candidates

### C01 - Reopen whether a Home progress bar belongs at all for small uploads

- When / session: `2026-04-06T23:41Z`, `T01`
- Primary refs: `T01#L195`, `T01#L198`, `T01#L235`
- AI course: the assistant brought a plan to the gate that still treated a Home progress bar as the default visual treatment for active upload.
- Human intervention: the human questioned the reasoning for using a progress bar on what are usually small uploads and forced designer consultation before plan approval.
- Better outcome forced: re-open the design seam instead of waving the current plan through.
- Why this is real: the human redirected an active plan, not just asked an idle design question.
- Confidence: `medium`
- PASS1 read: `intervention event but probably not an accepted incident on its own`

### C04 - Correct the implied promise that install plus restart should clear the `Needs attention` banner

- When / session: `2026-04-07T00:46Z`, `T03`
- Primary refs: `T03#L716`, `T03#L725`, `T03#L748`, `T03#L755`
- AI course: the assistant's earlier wording made it sound like pushing the patch to the phone should help the visible blocked-state banner.
- Human intervention: the human followed up because the phone still said `Needs attention` and asked whether the assistant was surprised or had failed to restart it.
- Better outcome forced: explain honestly that the patch prevents new bad rows but does not automatically heal the already-stranded rows.
- Why this is real: the human had to follow up because the assistant had implicitly promised more visible improvement than the patch could actually deliver.
- Confidence: `medium`
- PASS1 read: `possible accepted incident candidate`

### C06 - Bound PASS-0001 and shut down open-ended research

- When / session: `2026-04-07T01:31Z`, `T05`
- Primary refs: `T05#L718`, `T05#L721`, `T05#L748`, `T05#L751`
- AI course: the assistant was still spending time on SDK/doc lookup after the runtime was ready and the implementation slice could already be bounded concretely.
- Human intervention: the human explicitly said `No more open-ended SDK/doc research` and bounded the exact PASS-0001 slice that was acceptable.
- Better outcome forced: source edits and tests had to start immediately on the named bounded slice.
- Why this is real: this is a clear intervention against a drifting active course, though it reads more like local supervisory correction than a likely durable incident.
- Confidence: `medium`
- PASS1 read: `intervention event but probably not an accepted incident`

### C08 - Redirect the debug branch from the `Head` warning story to the real `Hips` coercion seam

- When / session: `2026-04-07T01:49Z`, `T07`
- Primary refs: `T07#L599`, `T07#L606`, `T07#L623`
- AI course: the worker had a current explanation centered on the retained Manny control rig and `Head` mismatch.
- Human intervention: the human redirected the branch to a narrower IKRig question: why the target `Spine` chain gets forced back to `Hips`.
- Better outcome forced: investigate the exact saved-IKRig structural cause instead of lingering on the earlier seam.
- Why this is real: the human had to move the debugging effort off a no-longer-sufficient explanation and onto the remaining deterministic divergence.
- Confidence: `medium`
- PASS1 read: `possible accepted incident candidate`

### C09 - Clarify that the already-authenticated CLI was not an extra-money gate

- When / session: `2026-04-07T14:25Z`, `T04`
- Primary refs: `T04#L1297`, `T04#L1304`, `T04#L1307`, `T04#L1320`
- AI course: the assistant framed the live run as a `spend gate` even though the local CLI was already authenticated via the user's existing ChatGPT-backed account.
- Human intervention: the human asked for a clearer explanation and then explicitly said there was no separate `additional money` in their model.
- Better outcome forced: reframe the gate as permission to use the existing account-backed run, not as if new money had to change hands.
- Why this is real: the human had to explain the practical meaning of the already-authenticated subscription boundary because the assistant was tracking the wrong blocker.
- Confidence: `medium`
- PASS1 read: `possible accepted incident candidate`

### C11 - Repeat the spend-gate clarification in the long-lived leader thread

- When / session: `2026-04-07T14:33Z`, `T05`
- Primary refs: `T05#L1316`, `T05#L1319`
- AI course: even after the main-thread clarification, the leader was still talking as if extra-money approval could be the blocker and was also behaving as if PATH uncertainty were still the main issue.
- Human intervention: the supervisor explicitly removed the `additional-money` blocker again, named runtime behavior as the real blocker, and pointed at stale PATH as a local shell issue rather than a human gate.
- Better outcome forced: move the debugging frame onto the real contained-proof failures.
- Why this is real: this is repeated human teaching because the earlier blocker correction did not fully stick.
- Confidence: `medium`
- PASS1 read: `possible accepted incident candidate`

### C13 - Supply exact local runtime facts and CLI semantics so the retry can proceed

- When / session: `2026-04-07T14:40Z` to `14:42Z`, `T05`
- Primary refs: `T05#L1513`, `T05#L1516`, `T05#L1538`, `T05#L1541`
- AI course: the assistant was still resolving compose-file path, installed `codex.exe` location, and sandbox flag semantics from inference and partial environment state.
- Human intervention: the human supplied the exact compose path, the observed installed binary path, and the real `codex exec --help` sandbox semantics.
- Better outcome forced: make the next retry from concrete local facts instead of more environmental guesswork.
- Why this is real: the human had to spend explanatory effort teaching local operator facts and tool semantics because the assistant was not grounding aggressively enough from the actual machine.
- Confidence: `medium`
- PASS1 read: `intervention event but probably not an accepted incident`

### C14 - Force an explicit bounded fork instead of more analysis

- When / session: `2026-04-07T14:47Z`, `T05`
- Primary refs: `T05#L1643`, `T05#L1651`, `T05#L1654`
- AI course: after the bounded retry still failed, the assistant was still in investigation mode instead of choosing the last bounded retry or recording an honest blocker.
- Human intervention: the human said `Decision required now`, gave the two acceptable branches, and explicitly said `Do not stay in analysis`.
- Better outcome forced: choose one bounded next move and keep state durable.
- Why this is real: the human had to impose decision discipline on an active stalled course.
- Confidence: `medium`
- PASS1 read: `possible accepted incident candidate`

### C15 - Force immediate closeout after proof success

- When / session: `2026-04-07T14:59Z` to `15:03Z`, `T05`
- Primary refs: `T05#L1865`, `T05#L1868`, `T05#L1924`, `T05#L1930`, `T05#L1932`
- AI course: after the final bounded proof looked successful, the assistant was still doing nonessential follow-up instead of finishing the durable closeout.
- Human intervention: the human prescribed the closeout order, said `Do not stop at summary; finish the task`, and then required a one-line status only.
- Better outcome forced: immediate durable closure rather than more open-ended follow-up.
- Why this is real: this is a real autonomy correction, though it is more likely task-supervision friction than a durable accepted incident.
- Confidence: `medium`
- PASS1 read: `intervention event but probably not an accepted incident`

## Likely accepted incident candidates

- `C02` Make the Home progress bar mean time-to-done.
- `C03` Call out the remaining disconnect as a dropped ball.
- `C05` Record the user's regression report charitably instead of flattening it into optimization.
- `C07` Resume deterministic root-cause pursuit instead of stopping at a plausible upstream producer.
- `C10` Contain an unsafe worker startup that released all scheduled jobs.
- `C12` Enforce durable-state-first before more debugging.
- `C16` Reject the proof-only closure model for a `long-running backend`.
- `C18` Call out the dropped ball on the dashboard surface itself.

## Real interventions but probably outside the accepted incident set

- `C01` Reopen whether a Home progress bar belongs at all for small uploads.
- `C04` Correct the implied promise that install plus restart should clear the `Needs attention` banner.
- `C06` Bound PASS-0001 and shut down open-ended research.
- `C13` Supply exact local runtime facts and CLI semantics so the retry can proceed.
- `C15` Force immediate closeout after proof success.

## Ambiguous boundaries that would need a wider reread

- `C08` could be merged with `C07` if later synthesis wants one broader `debugging stopped too early / wrong seam` incident, but the transcript supports keeping them separate locally because `C07` is a stopping-rule correction and `C08` is a seam redirection.
- `C16` and `C17` are adjacent. They stay split here because `C16` is the rejection of the assistant's closure model, while `C17` is the demand for the concrete always-on lane.
- `C18` and `C19` are adjacent. They stay split because `C18` is the form-factor/surface-burden complaint, while `C19` is the narrower `do not hand restart back to the human` correction.

## PASS1 takeaways for PASS2

- The widened criteria do recover explanatory-labor events cleanly on this day. The highest-signal ones are `C02`, `C03`, `C05`, `C07`, `C11`, `C12`, `C16`, `C17`, `C18`, and `C19`.
- The strongest repeated-human-teaching cluster is `Task-0005`: false blocker framing, durable-state honesty, bounded decision discipline, non-operator operating model, and human-surface burden.
- The clearest same-day repeated-teaching pattern is that once the assistant accepted a higher-level rule verbally, the human still often had to intervene again to make the assistant encode or enact it concretely.
