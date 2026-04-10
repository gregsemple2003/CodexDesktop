# INTERVENTION-PASS1

Source day: `2026-04-04`

## Source scope reviewed

- Primary source scope only: `C:\Users\gregs\.codex\sessions\2026\04\04\rollout-*.jsonl`
- Broad scan completed across all `21` rollout files in that source-day folder, covering `141` outbound human messages.
- Exact extraction aliases used for the candidate list:
  - `T04` = `C:\Users\gregs\.codex\sessions\2026\04\04\rollout-2026-04-04T01-22-17-019d56f1-261c-7e32-8a07-7323a20c471c.jsonl`
  - `T09` = `C:\Users\gregs\.codex\sessions\2026\04\04\rollout-2026-04-04T11-15-56-019d5910-a6c5-7373-bd7e-46d7a2a93cb9.jsonl`
  - `T14` = `C:\Users\gregs\.codex\sessions\2026\04\04\rollout-2026-04-04T11-18-43-019d5913-3060-70f1-a7d7-44a771efaef8.jsonl`
  - `T18` = `C:\Users\gregs\.codex\sessions\2026\04\04\rollout-2026-04-04T11-43-16-019d5929-ab66-7c21-9f8a-9dc6024325ab.jsonl`
  - `T19` = `C:\Users\gregs\.codex\sessions\2026\04\04\rollout-2026-04-04T11-44-43-019d592b-0085-7ad2-8cfd-22009eae6dfa.jsonl`
- Confirmatory duplicate-read only, not primary refs:
  - `T06` copied the earlier `Harvest task context` usability-rule discussion into a fresh prompt thread.
  - `T20` copied the Task-0004 UI discussion into the screenshot-review thread.
- Scope note: some rollout files stored under the `2026\04\04` source-day folder continue into later absolute timestamps. This pass stayed bounded to the source-day folder, not to a stricter timestamp cutoff.
- The daily CSV was used only as a late navigation spot-check after the raw JSONL read. No candidate below was accepted or rejected from CSV rows alone.

## Total candidate intervention events found

- `19`

## Chronological candidate list

### C01 - Reject open-ended narrow usability task definitions

- Session: `Harvest task context`
- Transcripts: `T04`
- Refs: `T04:L149`, `T04:L159`, `T04:L169`
- AI course: narrow Crystallize usability tasks had been written with loose definition-of-done language and unresolved product branches.
- Human intervention: the human said the tasks were too open-ended, clarified that bug reports should guide direction without becoming the whole design spec, and said the producer should drive consensus with interface/general design instead of leaving product state in superposition.
- Better outcome forced: point-fix usability tasks should capture one likely agreed human-facing resolution clearly enough that a reviewer can pick a side before implementation.
- Why real: this directly rejects the adequacy of the AI's current task-writing bar, not just the wording style.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C02 - Replace `declarative terms` with falsifiable or evaluable claims

- Session: `Harvest task context`
- Transcripts: `T04`
- Refs: `T04:L179`, `T04:L220`, `T04:L230`
- AI course: after the first correction, the assistant restated the rule in weaker `declarative terms`, which still allowed vague task definitions.
- Human intervention: the human explicitly rejected that weaker phrasing and insisted that usability tasks collapse into `falsifiable or evaluat-able claim[s]`.
- Better outcome forced: task definitions must be concrete enough that a human can agree with them, reject them, or demand refinement.
- Why real: the assistant's first repair still weakened the bar, so the human had to tighten the contract again.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C03 - Make usability resolutions hard-mockup-ready

- Session: `Harvest task context`
- Transcripts: `T04`
- Refs: `T04:L452`, `T04:L539`
- AI course: even after the falsifiable-claim update, the rewritten `Task-0022` wording still said what the UI would not be while leaving the actual UI underspecified.
- Human intervention: the human said the sentence still failed because the resolution was not clear enough to produce a hard mockup, then required prompt iteration and fresh subagent retesting.
- Better outcome forced: narrow usability task wording should name the visible end state concretely enough to draw and review, not merely negate the old state.
- Why real: this is a second intervention after the first task-writing correction, because the correction had not stuck at the actual human review bar.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C04 - Reject reboot-robustness overclaim on the server lane

- Session: `Harvest task context`
- Transcripts: `T04`
- Refs: `T04:L843`, `T04:L951`, `T04:L1008`
- AI course: live troubleshooting had focused on the current upload blockage, while earlier task artifacts had overclaimed that the server lane was already robust to machine restarts.
- Human intervention: the human restated the original bar as restart robustness, then required a bug that explicitly records overclaimed readiness and inadequate or dishonest proof before doing a retest.
- Better outcome forced: honest durable defect tracking around reboot persistence, not just a one-off live repair attempt.
- Why real: the human rejected the earlier `ready` story as inadequate in real-world terms and forced the failure into durable bug evidence.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C05 - Add screenshot-backed proof as a required Task-0002 bar

- Session: `Implement task 2 task` / `Lead Task-0002 workflow`
- Transcripts: `T09`, `T14`
- Refs: `T09:L221`, `T14:L268`
- AI course: Task-0002 planning had been approved on structured runtime evidence without requiring visible screenshot proof of the pawn.
- Human intervention: the human added a plan-approval caveat that screenshot functionality must prove pawn existence and that structured state alone is not enough.
- Better outcome forced: screenshot-backed proof carried durably into implementation, rather than state-only proof.
- Why real: this materially tightens what counts as acceptable proof for a human-facing result.
- Confidence: `medium`
- Triage: `intervention event but probably not an accepted incident`

### C06 - Force the first durable Task-0004 planning checkpoint

- Session: `Update Task-0004 state`
- Transcripts: `T19`
- Refs: `T19:L160`
- AI course: the delegated Task-0004 leader had not produced the initial durable checkpoint and was still broadly re-grounding.
- Human intervention: the human ordered immediate `TASK-STATE.json` and `PLAN.md` creation with the honest planning state and an explicit plan-approval gate.
- Better outcome forced: durable planning state on disk before more exploration.
- Why real: the human had to interrupt a nonproductive course and force concrete state ownership.
- Confidence: `medium`
- Triage: `intervention event but probably not an accepted incident`

### C07 - Bound proof-view debugging to one discriminating check

- Session: `Implement task 2 task` / `Lead Task-0002 workflow`
- Transcripts: `T09`, `T14`
- Refs: `T09:L543`, `T09:L580`, `T14:L607`, `T14:L751`
- AI course: the assistant was explaining an automation proof view and risked expanding camera-path experimentation instead of preserving the real gameplay camera and using proof view only as a discriminator.
- Human intervention: the human said the default third-person camera likely was not the problem, ordered the gameplay camera path left alone, and bounded proof-view work to one discriminating check.
- Better outcome forced: one narrow proof-view test that either closes the proof bar honestly or redirects into a real visibility/render defect.
- Why real: this is a direct redirect away from the wrong debugging seam plus a hard bound on further proof-helper expansion.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C08 - Stop Task-0004 exploration drift and land concrete patches

- Session: `Update Task-0004 state`
- Transcripts: `T19`
- Refs: `T19:L306`, `T19:L484`
- AI course: after the PASS-0000 checkpoint, the Task-0004 work kept rereading, statusing, and broadly exploring UI directions without landing product-file patches.
- Human intervention: the human first prescribed the minimal backend slice for PASS-0000, then had to intervene again and prescribe the smallest honest PASS-0001 Jobs-lane UI slice.
- Better outcome forced: concrete backend and UI patches instead of further implementation drift.
- Why real: the human had to intervene twice because the first anti-drift correction did not fully stick.
- Confidence: `medium`
- Triage: `intervention event but probably not an accepted incident`

### C09 - Interrupt the live Task-0004 hotkey regression immediately

- Session: `Implement task 4 task`
- Transcripts: `T18`
- Refs: `T18:L794`
- AI course: after the Jobs-lane work, the live dashboard no longer closed on `Ctrl Alt Space`.
- Human intervention: the human flagged it as a blocker bug because the broken overlay was still in their face and demanded an immediate fix.
- Better outcome forced: restore baseline live-overlay close behavior before treating the pass as healthy.
- Why real: this is an explicit rejection of a shipped live-surface regression, not a routine bug report without an AI boundary.
- Confidence: `medium`
- Triage: `intervention event but probably not an accepted incident`

### C10 - Narrow the ZeroMale defect boundary to the animation path

- Session: `Lead Task-0002 workflow`
- Transcripts: `T14`
- Refs: `T14:L1373`
- AI course: after the no-animation reference-pose run, the debugging lane could still have stayed broad around generic visibility/render uncertainty.
- Human intervention: the human said the no-anim run was the first decisive success and that the live defect boundary should now be treated as the `ZM_ABP_Unarmed` animation path, with artifacts updated accordingly.
- Better outcome forced: stop widening generic visibility debugging and narrow the product defect to the animation seam now supported by evidence.
- Why real: this is a concrete redirect to a narrower upstream cause, not mere commentary on the result.
- Confidence: `medium`
- Triage: `intervention event but probably not an accepted incident`

### C11 - Reject the stale-process stop point because the live surface still lacked `Jobs`

- Session: `Implement task 4 task`
- Transcripts: `T18`
- Refs: `T18:L967`
- AI course: after the hotkey-close hotfix, the assistant had effectively stopped while the running dashboard instance still showed no `Jobs` tab.
- Human intervention: the human pointed out that the live overlay still lacked the new tab and asked why work had stopped.
- Better outcome forced: restart the actual live surface and keep working until the visible overlay matches the claimed build state.
- Why real: the human had to intervene again because the first live fix did not yet produce the promised human-visible surface.
- Confidence: `medium`
- Triage: `intervention event but probably not an accepted incident`

### C12 - Reject Task-0004 regression closure after real `Jobs`-tab failures

- Session: `Implement task 4 task`
- Transcripts: `T18`
- Refs: `T18:L993`
- AI course: the assistant had treated smoke/startup proof as enough for closure even though the new human-facing `Jobs` click path had not been exercised honestly.
- Human intervention: the human reported real live failures on that click path: hitching, a crazy number of windows, and tab buttons that did not match the mockup.
- Better outcome forced: reopen the task, reproduce the actual click path, and fix both behavior and fidelity before claiming closure.
- Why real: the human explicitly rejected the assistant's regression claim on real-app evidence.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C13 - Tighten regression doctrine around repo-local click paths

- Session: `Implement task 4 task`
- Transcripts: `T18`
- Refs: `T18:L1050`, `T18:L1063`, `T18:L1104`
- AI course: the assistant's mental model still allowed the repo's smoke lane to count as sufficient regression for the new clickable `Jobs` surface.
- Human intervention: the human paused implementation, focused on the regression-expectation disconnect, and required explicit wording that ties the rule to the repo-local `REGRESSION.md`.
- Better outcome forced: future tasks that add or materially change clickable human-facing functionality must map to or update repo-local regression cases before closure.
- Why real: the human tightened the durable workflow contract because the assistant had interpreted it too weakly in practice.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C14 - Reject reference-pose closure and reopen Task-0002 under the real human bar

- Session: `Implement task 2 task` / `Lead Task-0002 workflow`
- Transcripts: `T09`, `T14`
- Refs: `T09:L1239`, `T09:L1259`, `T14:L1470`, `T14:L1567`, `T14:L2415`, `T14:L3475`
- AI course: Task-0002 closure and later checkpoints were still drifting toward `visible reference pose`, `defect seam isolated`, or blocked pass state as if those might count as enough.
- Human intervention: the human said that closing a visibility task with a T-pose or reference-pose fallback would be a dreadful narrowing of scope, reopened the task, and repeatedly restated that the real bar is a visibly animated and controllable pawn in real PIE.
- Better outcome forced: keep working until the promoted ZeroMale pawn is visibly animated and controllable with screenshot-backed proof on the live path.
- Why real: the human explicitly rejected the assistant's closure interpretation and had to restate the correction repeatedly as later checkpoints risked drifting again.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C15 - Stop reconcile on tab entry and force fidelity repair on the live surface

- Session: `Implement task 4 task`
- Transcripts: `T18`
- Refs: `T18:L1144`, `T18:L1210`
- AI course: even after the regression-doc fix, the live `Jobs` tab still hitched because tab entry triggered reconcile work, and visible fidelity was still off enough that the human requested a screenshot review.
- Human intervention: the human ordered reconcile removed from the tab click path, required fidelity work against the mockup, and requested an interface-designer screenshot review.
- Better outcome forced: tab clicks should only switch surfaces, while performance and fidelity remain part of the real fix bar.
- Why real: the human rejected a process-only repair and pulled the focus back to the live human-facing surface.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C16 - Force TASK-HARVESTER prime directives for concrete measurable tasks

- Session: `Implement task 2 task`
- Transcripts: `T09`
- Refs: `T09:L1363`, `T09:L1403`, `T09:L1469`, `T09:L1540`, `T09:L1590`
- AI course: even after the Task-0002 miss was acknowledged, the first TASK-HARVESTER repair still left the task-writing priority too implicit and too tolerant of broad language.
- Human intervention: the human demanded clearer human-aware task language, then insisted on an explicit prime directive that good tasks must be specific and measurable enough to mark done, and finally forced that rewrite to become the new Task-0002 north star.
- Better outcome forced: top-priority task-writing rules that make vague or proxy-friendly task bars harder to slip through later.
- Why real: the human had to intervene again because the first prompt repair still did not make the intended priority unmistakable.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C17 - Separate primary navigation from tab content structurally

- Session: `Implement task 4 task`
- Transcripts: `T18`
- Refs: `T18:L1243`, `T18:L1253`, `T18:L1263`
- AI course: the Task-0004 shell still mixed primary nav and tab-specific controls in one strip.
- Human intervention: the human specified that primary nav must be its own strip and that the area beneath it must hold the tab content, with only a limited render-line exception kept from the mockup.
- Better outcome forced: a clearer IA and visual hierarchy where navigation and per-tab controls are structurally separated.
- Why real: the human corrected the assistant's UI structure, not just optional polish.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C18 - Clarify declared jobs JSON plus schema as the primary durable jobs state

- Session: `Implement task 4 task`
- Transcripts: `T18`
- Refs: `T18:L1716`
- AI course: the Jobs lane still read as if actual runtime state or on-click reconcile state might be the primary truth being surfaced.
- Human intervention: the human clarified that Codex jobs should have a durable primary JSON declaration plus schema under `.codex/Orchestration/Jobs`.
- Better outcome forced: the Jobs UI should anchor on declared jobs first, with reconcile treated as an actual-vs-declared diff layer on top.
- Why real: the human corrected the architectural truth the surface was supposed to represent.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C19 - Require a scrollable `Jobs` panel

- Session: `Implement task 4 task`
- Transcripts: `T18`
- Refs: `T18:L1855`
- AI course: the `Jobs` surface shipped without scrollability.
- Human intervention: the human explicitly called out that the panel was not scrollable.
- Better outcome forced: a usable `Jobs` surface when content exceeds the visible panel.
- Why real: this is a direct human correction of a live usability miss.
- Confidence: `medium`
- Triage: `intervention event but probably not an accepted incident`

## Confidence buckets

### Strong candidates

- `C01` Reject open-ended narrow usability task definitions
- `C02` Replace `declarative terms` with falsifiable or evaluable claims
- `C03` Make usability resolutions hard-mockup-ready
- `C04` Reject reboot-robustness overclaim on the server lane
- `C07` Bound proof-view debugging to one discriminating check
- `C12` Reject Task-0004 regression closure after real `Jobs`-tab failures
- `C13` Tighten regression doctrine around repo-local click paths
- `C14` Reject reference-pose closure and reopen Task-0002 under the real human bar
- `C15` Stop reconcile on tab entry and force fidelity repair on the live surface
- `C16` Force TASK-HARVESTER prime directives for concrete measurable tasks
- `C17` Separate primary navigation from tab content structurally
- `C18` Clarify declared jobs JSON plus schema as the primary durable jobs state

### Medium candidates

- `C05` Add screenshot-backed proof as a required Task-0002 bar
- `C06` Force the first durable Task-0004 planning checkpoint
- `C08` Stop Task-0004 exploration drift and land concrete patches
- `C09` Interrupt the live Task-0004 hotkey regression immediately
- `C10` Narrow the ZeroMale defect boundary to the animation path
- `C11` Reject the stale-process stop point because the live surface still lacked `Jobs`
- `C19` Require a scrollable `Jobs` panel

### Weak or ambiguous promoted candidates

- None promoted into the main list at `weak` confidence. The questionable boundaries are called out below instead.

## Which candidates look like likely accepted incidents

- `C01` Reject open-ended narrow usability task definitions
- `C02` Replace `declarative terms` with falsifiable or evaluable claims
- `C03` Make usability resolutions hard-mockup-ready
- `C04` Reject reboot-robustness overclaim on the server lane
- `C07` Bound proof-view debugging to one discriminating check
- `C12` Reject Task-0004 regression closure after real `Jobs`-tab failures
- `C13` Tighten regression doctrine around repo-local click paths
- `C14` Reject reference-pose closure and reopen Task-0002 under the real human bar
- `C15` Stop reconcile on tab entry and force fidelity repair on the live surface
- `C16` Force TASK-HARVESTER prime directives for concrete measurable tasks
- `C17` Separate primary navigation from tab content structurally
- `C18` Clarify declared jobs JSON plus schema as the primary durable jobs state

## Which candidates are real interventions but probably belong outside the accepted incident set

- `C05` Add screenshot-backed proof as a required Task-0002 bar
- `C06` Force the first durable Task-0004 planning checkpoint
- `C08` Stop Task-0004 exploration drift and land concrete patches
- `C09` Interrupt the live Task-0004 hotkey regression immediately
- `C10` Narrow the ZeroMale defect boundary to the animation path
- `C11` Reject the stale-process stop point because the live surface still lacked `Jobs`
- `C19` Require a scrollable `Jobs` panel

## Any ambiguous boundaries that need a second read

- `C05` is real, but it may be better understood as the lead-in to `C07` rather than a durable incident on its own. The approval-caveat versus intervention boundary is thin.
- `C06` and `C08` are both authentic anti-drift interventions, but they sit close to normal task supervision. A stricter accepted-incident pass may keep them outside the final corpus even though the human clearly had to step in.
- Possible extra candidate not promoted separately: `T14:L1154`, where the human said the latest triangle artifact was enough and forced one concrete render-side fix trial instead of more diagnostics. It looked real, but in this pass it folds cleanly into the broader Task-0002 seam-correction chain.
- Possible extra candidate not promoted separately: `T18:L1518` and `T18:L1636`, where later live UI/layout/copy corrections kept coming after the main shell repair. They likely matter as surface feedback, but the boundary between real intervention and ordinary iterative design review is thinner than the earlier Task-0004 breakpoints.
- Possible extra candidate not promoted: `T17:L75`, where the human said `sorry ignore me, wrong project undo changes in this thread`. I did not promote it because the harmful AI course had barely started before the human corrected the routing mistake.
