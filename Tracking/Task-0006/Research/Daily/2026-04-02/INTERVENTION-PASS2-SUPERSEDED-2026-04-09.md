# INTERVENTION-PASS2

Source day: `2026-04-02`

## Source scope analyzed

- Zero-context backfill. Consulted only local artifacts derived directly from this source day.
- PASS1 artifact: `C:\Agent\CodexDashboard\Tracking\Task-0006\Research\Daily\2026-04-02\INTERVENTION-PASS1.md`
- Incident corpus contract reread:
  - `C:\Users\gregs\.codex\Orchestration\Reports\Incidents\README.md`
  - `C:\Users\gregs\.codex\Orchestration\Reports\Incidents\INCIDENT.schema.json`
- Raw transcripts reopened from the PASS1 refs:
  - `T01` = `C:\Users\gregs\.codex\sessions\2026\04\02\rollout-2026-04-02T03-03-04-019d4d00-b1a4-7d93-9c26-e06ce205db13.jsonl`
  - `T02` = `C:\Users\gregs\.codex\sessions\2026\04\02\rollout-2026-04-02T03-47-36-019d4d29-77ca-7c62-95d7-c2f2ed12cf09.jsonl`
  - `T03` = `C:\Users\gregs\.codex\sessions\2026\04\02\rollout-2026-04-02T12-41-19-019d4f12-19e2-7fd3-8896-f51d93edbf9f.jsonl`
  - `T04` = `C:\Users\gregs\.codex\sessions\2026\04\02\rollout-2026-04-02T16-54-20-019d4ff9-bedd-7470-8e3a-ce2946226a4b.jsonl`
  - `T05` = `C:\Users\gregs\.codex\sessions\2026\04\02\rollout-2026-04-02T16-59-35-019d4ffe-8bfa-7903-a5c4-7ad5eb0d6aef.jsonl`
  - `T06` = `C:\Users\gregs\.codex\sessions\2026\04\02\rollout-2026-04-02T19-43-23-019d5094-823f-7932-a8ca-f0c75e6f98c8.jsonl`
  - `T07` = `C:\Users\gregs\.codex\sessions\2026\04\02\rollout-2026-04-02T21-40-28-019d50ff-b51b-7c82-9c73-207d0d24abf6.jsonl`
- No prior accepted-incident JSONs, daily briefs, or parent-thread summaries were consulted.
- No extra repo/task artifact rereads were required for local PASS2 classification. In the few places where the assistant itself pulled repo truth from on-disk files inside the transcript, I treated that as part of the event evidence rather than as an independent source.
- Source-day note: this backfill follows the folder day `2026-04-02`. Several transcript timestamps inside those files extend into `2026-04-03`; I preserved the raw transcript timing as written.

## Candidate ids analyzed

- `C01` through `C18`

## Boundary corrections relative to PASS1

- No hard split or merge corrections were required from the PASS1 candidate set.
- `C03` stays separate from `C04`. `C03` is the local semantics and workflow-specialization correction around the `Connected` card. `C04` is the later rejection of workflow-only closure when the real task artifacts and emulator work had not been carried through.
- `C05` stays separate from `C06`. `C05` is the false closure on the visible Sync surface. `C06` is the later hardening of the implementation contract after that reopen still did not stick cleanly.
- `C11` stays intentionally wide. The transcript supports one recurring correction arc around the same real contract: background-safe PIE, correct repro map, and no focus theft while the human keeps working.
- `C18` also stays intentionally wide. The same user complaint keeps recurring against the investigation feature: session-summary output is not enough; the tool must answer trigger, mechanism, and avoidable mistake.

## Per-event analysis records

### C01 - Keep UI-fidelity ownership out of `AGENTS.md`

- `event_id`: `C01`
- `title`: `Keep UI-fidelity ownership out of AGENTS.md`
- `session_or_thread`: `Crystallize doc-seam follow-up`
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\02\rollout-2026-04-02T03-03-04-019d4d00-b1a4-7d93-9c26-e06ce205db13.jsonl`
- `primary_refs`: `T01:L449`, `T01:L452`, `T01:L459`, `T01:L462`
- `ai_course`: After agreeing that the task leader should own final UI fidelity sign-off, the assistant had already spread the new rule across `GENERAL-DESIGN.md`, `TESTING.md`, and `AGENTS.md`.
- `human_intervention`: The human rejected the `AGENTS.md` placement, said the front-door repo doc should stay light, and redirected the explicit ownership rule into the task-leader prompt layer while leaving verification detail in `TESTING.md`.
- `adequate_outcome`: Keep repo-specific UI verification detail in `TESTING.md`, keep `AGENTS.md` light, and put task-leader responsibility in the prompt source that actually governs leader behavior.
- `event_boundary_notes`: PASS1 boundary is accurate. This is a real correction of documentation seam and ownership layer, not a wording-only preference.
- `human_model_signal`: Explicit doc-layer rule: `AGENTS.md` should not grow by carrying every edge case, and the task leader's responsibility should not be left as a `free floating constraint`.
- `failure_family_hypothesis`: `workflow_orchestration`
- `intervention_kind_hypothesis`: `reframe_problem`
- `human_cost_or_risk`: Front-door doc bloat plus a weaker ownership signal, because the actual responsible role might not receive the rule in its operative prompt.
- `local_lesson_hypothesis`: Put durable behavioral responsibility in the prompt layer that owns the actor, and keep front-door repo docs light enough to stay legible.
- `cluster_hints`: `doc-seam`, `prompt-vs-front-door`, `responsibility-ownership`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: none material

### C02 - Redirect token-loss handling into session forensics

- `event_id`: `C02`
- `title`: `Redirect token-loss handling into session forensics`
- `session_or_thread`: `Task-0016 token-forensics thread`
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\02\rollout-2026-04-02T03-47-36-019d4d29-77ca-7c62-95d7-c2f2ed12cf09.jsonl`
- `primary_refs`: `T02:L281`, `T02:L292`, `T02:L299`, `T02:L308`, `T02:L311`
- `ai_course`: The assistant had corrected its used-vs-remaining quota read and was still close to acting on live-thread handling, including offering to stop the worker thread.
- `human_intervention`: The human explicitly said `No don't touch anything`, described a sudden local loss of session history and backup state after reinstall/reclone churn, and asked the assistant to look at everything because they could not see what was consuming tokens.
- `adequate_outcome`: Treat the problem as local session forensics first: inspect session logs, thread index, and recent local state before touching code or thread topology.
- `event_boundary_notes`: PASS1 boundary is accurate. The key local move is from possible corrective action to broad local diagnosis.
- `human_model_signal`: Explicit diagnostic-scope rule: `No don't touch anything ... please look at everything` because the human could not see what was burning tokens.
- `failure_family_hypothesis`: `workflow_orchestration`
- `intervention_kind_hypothesis`: `redirect_debugging`
- `human_cost_or_risk`: More action or thread churn could have obscured the real cause while the human was already missing local history.
- `local_lesson_hypothesis`: When the human reports unexplained resource loss and missing local history, freeze action and do local forensics before making more changes.
- `cluster_hints`: `diagnostic-scope`, `local-forensics`, `do-not-churn`
- `accepted_incident_likelihood`: `unlikely`
- `confidence`: `strong`
- `uncertainties`: The intervention is real even if a later pass ultimately explains the underlying quota story more narrowly.

### C03 - Treat `Connected` card copy as ambiguous human semantics

- `event_id`: `C03`
- `title`: `Treat Connected card copy as ambiguous human semantics`
- `session_or_thread`: `Task-0016 Home semantics and interface-designer setup`
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\02\rollout-2026-04-02T03-47-36-019d4d29-77ca-7c62-95d7-c2f2ed12cf09.jsonl`
- `primary_refs`: `T02:L2678`, `T02:L2692`, `T02:L2699`, `T02:L2720`, `T02:L2727`, `T02:L2760`
- `ai_course`: The assistant had compared the emulator to the Home mockup and identified structural gaps, but the shipped `Connected` card language still read like upload/server state instead of people/profile readiness.
- `human_intervention`: The human stopped on the card meaning itself, asked what `Connected`, `Active now`, and `Me, Dad, 2 others ready` were supposed to mean, then said `Not ready` was too vague for humans and reframed the fix as a workflow specialization problem that needed an `INTERFACE-DESIGNER` role.
- `adequate_outcome`: Human-readable card semantics that say what setup is missing, plus a workflow that explicitly pressure-tests labels and iconography for human interpretation instead of only structural resemblance.
- `event_boundary_notes`: PASS1 boundary is accurate. This is the semantic/human-meaning event, not yet the later reopen over premature workflow-only closure.
- `human_model_signal`: Explicit human-meaning rule: `not ready` is `too vague` and leaves the state `too open for interpretation`.
- `failure_family_hypothesis`: `ui_semantics`
- `intervention_kind_hypothesis`: `reframe_problem`
- `human_cost_or_risk`: Ordinary human reading would misclassify the state story on the Home surface.
- `local_lesson_hypothesis`: If a status label can plausibly be read as the wrong domain, rename it around the real concept and name the missing setup explicitly.
- `cluster_hints`: `status-signal-meaning`, `human-semantics`, `interface-designer`
- `accepted_incident_likelihood`: `possible`
- `confidence`: `strong`
- `uncertainties`: This may cluster into the larger Task-0016 workflow/fidelity family rather than stand alone in the accepted set.

### C04 - Reopen Task-0016 because prompt hardening was not real closure

- `event_id`: `C04`
- `title`: `Reopen Task-0016 because prompt hardening was not real closure`
- `session_or_thread`: `Task-0016 UI follow-up`
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\02\rollout-2026-04-02T03-47-36-019d4d29-77ca-7c62-95d7-c2f2ed12cf09.jsonl`
- `primary_refs`: `T02:L2805`, `T02:L2833`, `T02:L2840`, `T02:L2855`
- `ai_course`: After adding and testing `INTERFACE-DESIGNER.md`, the assistant reported the workflow change as complete and explicitly noted that no app code changed in that step.
- `human_intervention`: The human asked whether any artifacts had actually been changed, clarified that they wanted the assistant to `take the ball for awhile`, iterate prompt and design artifacts, and come back with a working emulator and the issues fixed.
- `adequate_outcome`: Reopen the task honestly, update task/design artifacts, then continue through real emulator-visible UI correction rather than stopping at orchestration hardening.
- `event_boundary_notes`: PASS1 boundary is accurate. This is the first strong rejection of the assistant's premature stopping point.
- `human_model_signal`: Explicit real-world completion rule: `take the ball for awhile` and return with `a working emulator with all issues fixed`, not just a workflow patch.
- `failure_family_hypothesis`: `workflow_orchestration`
- `intervention_kind_hypothesis`: `reject_outcome`
- `human_cost_or_risk`: Workflow-layer changes could be mistaken for product progress while the actual UI and task state stayed stale.
- `local_lesson_hypothesis`: Prompt or workflow hardening is not completion when the real task artifacts and running surface still need to be carried through.
- `cluster_hints`: `closure-truth`, `workflow-fix-vs-product-fix`, `take-the-ball`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: none material

### C05 - Reject PASS-0007 closure while Sync shell gaps were still visible

- `event_id`: `C05`
- `title`: `Reject PASS-0007 closure while Sync shell gaps were still visible`
- `session_or_thread`: `Task-0016 Sync / Status follow-up`
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\02\rollout-2026-04-02T03-47-36-019d4d29-77ca-7c62-95d7-c2f2ed12cf09.jsonl`
- `primary_refs`: `T02:L3684`, `T02:L3704`, `T02:L3711`, `T02:L3714`, `T02:L3725`
- `ai_course`: The assistant declared `PASS-0007` done, cited the pushed checkpoint and proof artifacts, and presented the Sync / Status surface as closed.
- `human_intervention`: The human immediately said there were still several `dropped balls`, named the missing persistent top and bottom bars as the biggest miss, called out missing summary regions and an out-of-place Back button, and asked for a probe into why the interface-designer let those mismatches through.
- `adequate_outcome`: Reopen the pass, diagnose why the delegated reviewer passed an obviously incomplete shell, and continue iteration until the live screen matches the approved structure.
- `event_boundary_notes`: PASS1 boundary is accurate. This is the live-surface closure rejection; `C06` is the subsequent prompt hardening response.
- `human_model_signal`: Explicit shell rule: the top and bottom bars `should always be in-view when the app is up`.
- `failure_family_hypothesis`: `verification_proof`
- `intervention_kind_hypothesis`: `reject_outcome`
- `human_cost_or_risk`: A UI pass would close on a visibly incomplete surface that the human could invalidate at a glance.
- `local_lesson_hypothesis`: Do not close a UI pass on partial semantic progress when major visible shell structure still diverges from the mockup.
- `cluster_hints`: `closure-truth`, `mockup-completeness`, `visible-shell`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: none material

### C06 - Add post-implementation self-audit to `INTERFACE-DESIGNER`

- `event_id`: `C06`
- `title`: `Add post-implementation self-audit to INTERFACE-DESIGNER`
- `session_or_thread`: `Task-0016 Sync / Status follow-up`
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\02\rollout-2026-04-02T03-47-36-019d4d29-77ca-7c62-95d7-c2f2ed12cf09.jsonl`
- `primary_refs`: `T02:L3962`, `T02:L3976`, `T02:L3982`, `T02:L3985`, `T02:L3992`
- `ai_course`: Even after the false closure was caught, the assistant was still treating the remaining gap as ordinary implementation cleanup and piecemeal correction.
- `human_intervention`: The human explicitly asked for a prompt revision hardened around the exact implementation gaps being witnessed, then ordered the prompt patched and relayed back to the same agent for another try.
- `adequate_outcome`: A delegated implementation role must re-audit its own patch against the mockup and cannot call a pass done while any formerly blocking discrepancy remains open.
- `event_boundary_notes`: PASS1 boundary is accurate. This is a second intervention because the first reopen did not yet change the delegated implementation bar enough.
- `human_model_signal`: Explicit contract rule: harden the prompt around `gaps that you're witnessing now`, not around abstract best practices.
- `failure_family_hypothesis`: `workflow_orchestration`
- `intervention_kind_hypothesis`: `tighten_contract`
- `human_cost_or_risk`: The same role could critique correctly in review mode and still ship a `close enough` implementation.
- `local_lesson_hypothesis`: If a delegated role misses blocking discrepancies after implementing, harden the implementation-side self-audit contract, not just the critique language.
- `cluster_hints`: `prompt-hardening`, `self-audit`, `implementation-vs-review`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: none material

### C07 - Verify the installed build before trusting claimed UI fixes

- `event_id`: `C07`
- `title`: `Verify the installed build before trusting claimed UI fixes`
- `session_or_thread`: `Task-0016 Sync / Status follow-up`
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\02\rollout-2026-04-02T03-47-36-019d4d29-77ca-7c62-95d7-c2f2ed12cf09.jsonl`
- `primary_refs`: `T02:L4132`, `T02:L4135`, `T02:L4145`, `T02:L4152`
- `ai_course`: The assistant had the emulator up and the delegated agent had claimed the top and bottom bar persistence issues were fixed.
- `human_intervention`: The human said they were not seeing those changes at all, pointed to the still-missing bars and Back button, and pushed for a rebuild/rerun after the assistant confirmed the emulator was still showing the old installed build.
- `adequate_outcome`: Rebuild and reinstall the current pass, then only claim progress against the live surface the human is actually seeing.
- `event_boundary_notes`: PASS1 boundary is accurate. The important local distinction is between in-progress working-tree changes and the running artifact under review.
- `human_model_signal`: Explicit live-surface rule: if the human still sees the old UI, the fix does not count yet, whatever the source diff says.
- `failure_family_hypothesis`: `verification_proof`
- `intervention_kind_hypothesis`: `reject_outcome`
- `human_cost_or_risk`: Wasted human review and reduced trust because the assistant reported a fix that was not present on the visible surface.
- `local_lesson_hypothesis`: After UI changes, verify the actual installed/running build before asking the human to judge whether the fix landed.
- `cluster_hints`: `live-surface-truth`, `stale-build`, `build-install-verification`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: none material

### C08 - Keep pushing the live Rustfire client failure after a partial diagnosis

- `event_id`: `C08`
- `title`: `Keep pushing the live Rustfire client failure after a partial diagnosis`
- `session_or_thread`: `Rustfire client/server debugging`
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\02\rollout-2026-04-02T12-41-19-019d4f12-19e2-7fd3-8896-f51d93edbf9f.jsonl`
- `primary_refs`: `T03:L798`, `T03:L804`, `T03:L818`, `T03:L825`, `T03:L828`
- `ai_course`: The assistant had isolated the stale cooked-client crash, shown that the editor-game entrypoint avoided that specific failure, and was close to treating that slice as the main resolution.
- `human_intervention`: The human said `Its still having a problem` and redirected the assistant to leave the server running and keep iterating on the remaining live client issue until root cause was fixed.
- `adequate_outcome`: Do not treat a solved sub-cause as practical closure while the real client failure is still live.
- `event_boundary_notes`: PASS1 boundary is accurate. The important local correction is against premature decomposition, not against the later wrong-fix architecture claim in `C09`.
- `human_model_signal`: Explicit persistence rule: the problem is still the live client problem until that path is actually fixed.
- `failure_family_hypothesis`: `workflow_orchestration`
- `intervention_kind_hypothesis`: `redirect_debugging`
- `human_cost_or_risk`: A still-live defect could be partially parked because one upstream failure mode was explained.
- `local_lesson_hypothesis`: When the observed failure remains live, keep the diagnostic bar on that live path instead of mentally cashing out on the first solved slice.
- `cluster_hints`: `premature-segmentation`, `stay-on-live-failure`, `root-cause-persistence`
- `accepted_incident_likelihood`: `possible`
- `confidence`: `strong`
- `uncertainties`: A later clustering pass may merge this into `C09` as a lead-in rather than keep it separate.

### C09 - Reject a technically working but network-invalid Mover fix

- `event_id`: `C09`
- `title`: `Reject a technically working but network-invalid Mover fix`
- `session_or_thread`: `Rustfire client/server debugging`
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\02\rollout-2026-04-02T12-41-19-019d4f12-19e2-7fd3-8896-f51d93edbf9f.jsonl`
- `primary_refs`: `T03:L1060`, `T03:L1076`, `T03:L1079`, `T03:L1086`, `T03:L1109`, `T03:L1220`, `T03:L1223`
- `ai_course`: The assistant reported that the connection path was now working and framed the standalone-liaison rollback as the operative fix.
- `human_intervention`: The human objected that this could not be the right networked Mover answer, and then explicitly ordered the assistant to undo its changes and solve the problem correctly.
- `adequate_outcome`: Distinguish symptom relief from an architecture that is actually valid in networked play, undo the triage change, and continue from the correct seam.
- `event_boundary_notes`: PASS1 boundary is accurate. The transcript supports one event from first architectural objection through explicit undo order.
- `human_model_signal`: Explicit adequacy rule: a fix is not acceptable just because `that last one worked` if it is wrong for networked Mover behavior.
- `failure_family_hypothesis`: `verification_proof`
- `intervention_kind_hypothesis`: `reject_outcome`
- `human_cost_or_risk`: A technically working but architecturally invalid path could be normalized as the fix.
- `local_lesson_hypothesis`: Treat `it works now` as insufficient when the fix contradicts the runtime model the task actually has to satisfy.
- `cluster_hints`: `works-but-wrong`, `architecture-vs-symptom`, `networked-validity`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: none material

### C10 - Re-check the malformed server wrapper instead of narrating a relaunch

- `event_id`: `C10`
- `title`: `Re-check the malformed server wrapper instead of narrating a relaunch`
- `session_or_thread`: `Rustfire server relaunch thread`
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\02\rollout-2026-04-02T12-41-19-019d4f12-19e2-7fd3-8896-f51d93edbf9f.jsonl`
- `primary_refs`: `T03:L1965`, `T03:L1979`, `T03:L1985`, `T03:L1988`
- `ai_course`: The assistant said it was relaunching with a stricter `cmd /k` wrapper and verifying the live child process, implying the wrapper issue was already understood.
- `human_intervention`: The human immediately said `Your paths must be off`, rejecting the assistant's relaunch story while the syntax/path problem was still visible.
- `adequate_outcome`: Check the exact command line and path resolution before continuing the relaunch narrative.
- `event_boundary_notes`: PASS1 boundary is accurate. This reads as a real correction, but a local tooling one.
- `human_model_signal`: Explicit launch-truth rule: if the wrapper is still visibly malformed, the assistant does not yet understand the real launch path.
- `failure_family_hypothesis`: `workflow_orchestration`
- `intervention_kind_hypothesis`: `reject_outcome`
- `human_cost_or_risk`: Time lost on wrapper narration while the actual malformed launch path is still unresolved.
- `local_lesson_hypothesis`: When a wrapper launch still visibly errors, verify the exact command line and resolved paths before claiming the restart path is fixed.
- `cluster_hints`: `tooling-wrapper`, `launcher-verification`
- `accepted_incident_likelihood`: `unlikely`
- `confidence`: `strong`
- `uncertainties`: This still looks more like a local tooling slip than a durable accepted-incident candidate.

### C11 - Enforce a background-safe, CombatMap, no-focus PIE debugging contract

- `event_id`: `C11`
- `title`: `Enforce a background-safe, CombatMap, no-focus PIE debugging contract`
- `session_or_thread`: `Rustfire PIE debugging and throttle investigation`
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\02\rollout-2026-04-02T16-54-20-019d4ff9-bedd-7470-8e3a-ce2946226a4b.jsonl`
- `primary_refs`: `T04:L485`, `T04:L614`, `T04:L1075`, `T04:L1141`, `T04:L1925`, `T04:L1928`
- `ai_course`: The assistant initially planned to reproduce by activating the editor window and sending the PIE shortcut, then later validated on the default map and kept relying on focus-sensitive control paths even after learning the throttle root cause.
- `human_intervention`: The human kept reimposing the real contract: just start PIE after stabilization, turn foregrounding off, make it work as a background app because they are doing something else, run the validation on `CombatMap`, and stop depending on focus/keystrokes because they want to keep working.
- `adequate_outcome`: A background-safe PIE workflow on `CombatMap` that does not steal focus as an ordinary control path, because the human's live-work boundary is part of the repro contract.
- `event_boundary_notes`: PASS1 intentionally kept this wide, and the transcript supports that choice. The same control-boundary and correct-repro rule had to be restated multiple times against renewed drift.
- `human_model_signal`: Explicit control-boundary rule: `i'm deliberately backgrounding the ue window because i'm doing something` and `don't depend on focus at all i want to keep working here`.
- `failure_family_hypothesis`: `human_world`
- `intervention_kind_hypothesis`: `redirect_debugging`
- `human_cost_or_risk`: Focus theft and wrong-map validation could both invalidate the repro while interrupting the human's ongoing work.
- `local_lesson_hypothesis`: When the user states a control boundary and repro map explicitly, treat those as part of the bug contract, not as optional debugging convenience.
- `cluster_hints`: `control-boundary`, `background-safe-debugging`, `correct-repro`, `focus-theft`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: A later pass could split the map correction from the focus-boundary correction, but one local event still fits the transcript.

### C12 - Correct the literal scope of `push .codex`

- `event_id`: `C12`
- `title`: `Correct the literal scope of push .codex`
- `session_or_thread`: `Rustfire PIE debugging and .codex push-scope thread`
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\02\rollout-2026-04-02T16-54-20-019d4ff9-bedd-7470-8e3a-ce2946226a4b.jsonl`
- `primary_refs`: `T04:L1303`, `T04:L1313`, `T04:L1325`, `T04:L1332`, `T04:L1361`, `T04:L1364`
- `ai_course`: The assistant committed the `.codex` skill locally in `EHG_GregS_main`, but the commit widened into already staged repo changes beyond the intended scope and then reported that commit before the meaning of the instruction was corrected.
- `human_intervention`: The human clarified that when they say `push .codex` they mean `JUST .codex`, not a widened staged-scope commit in the game repo, then had the assistant soft-reset the local commit back into staging.
- `adequate_outcome`: Respect the named target literally and do not widen repo or staged scope beyond what the instruction names.
- `event_boundary_notes`: PASS1 boundary is accurate. The local repair arc matters because the user had to manage recovery from the wrong repo/scope action, not just restate preference for the future.
- `human_model_signal`: Explicit literal-scope rule: `when i say push .codex it means JUST .codex`.
- `failure_family_hypothesis`: `workflow_orchestration`
- `intervention_kind_hypothesis`: `tighten_contract`
- `human_cost_or_risk`: The human had to reason through git recovery for a commit that landed in the wrong repo scope and pulled in unrelated staged work.
- `local_lesson_hypothesis`: Treat named repo/path targets literally; do not `helpfully` widen action scope when the human names a narrower boundary.
- `cluster_hints`: `repo-boundary`, `commit-scope`, `named-target-literalism`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: none material

### C13 - Invalidate Sync-screen pass status when bottom-bar `Home` is dead

- `event_id`: `C13`
- `title`: `Invalidate Sync-screen pass status when bottom-bar Home is dead`
- `session_or_thread`: `Task-0016 PASS-0008 closeout`
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\02\rollout-2026-04-02T16-59-35-019d4ffe-8bfa-7903-a5c4-7ad5eb0d6aef.jsonl`
- `primary_refs`: `T05:L991`, `T05:L998`, `T05:L1001`, `T05:L1008`, `T05:L1011`, `T05:L1023`
- `ai_course`: The assistant had just said the cold interface-designer judged the screen `pass` with no blocking issues and was ready to stop polishing the Sync screen itself.
- `human_intervention`: The human asked to see Home and immediately found that clicking the bottom-bar `Home` button did nothing, which forced the assistant to admit the bottom bar was still only a visual shell.
- `adequate_outcome`: A screen with a visible navigation affordance cannot be treated as effectively done while that affordance is dead.
- `event_boundary_notes`: PASS1 boundary is accurate. The human invalidates the `pass` state by exercising the actual visible surface.
- `human_model_signal`: Explicit live-surface rule: if the visible bottom-bar action does nothing, the screen is not truly at pass/closure quality.
- `failure_family_hypothesis`: `human_world`
- `intervention_kind_hypothesis`: `reject_outcome`
- `human_cost_or_risk`: The surface could be called done while a plainly visible navigation affordance still fails on the actual app.
- `local_lesson_hypothesis`: For a human-facing surface, a cold `pass` judgment does not survive a dead primary affordance.
- `cluster_hints`: `live-affordance-truth`, `pass-vs-real-nav`, `human-visible-bar`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: none material

### C14 - Keep supervising until the delegated task leader reaches a real endpoint

- `event_id`: `C14`
- `title`: `Keep supervising until the delegated task leader reaches a real endpoint`
- `session_or_thread`: `Task-0017 supervision thread`
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\02\rollout-2026-04-02T19-43-23-019d5094-823f-7932-a8ca-f0c75e6f98c8.jsonl`
- `primary_refs`: `T06:L182`, `T06:L193`, `T06:L201`, `T06:L208`, `T06:L211`, `T06:L223`
- `ai_course`: The assistant was still monitoring the delegated task leader, but it emitted a concise status/final-style stopping point before the child leader had actually finished.
- `human_intervention`: The human explicitly asked `You stopped?`, clarified that the intent was for the assistant to continue working until the task leader was done, and asked for the disconnect.
- `adequate_outcome`: Keep supervision alive until the delegated leader reaches a real endpoint, blocker, or human gate; status summaries are not a substitute for ongoing ownership.
- `event_boundary_notes`: PASS1 boundary is accurate. The core event is premature supervision exit, not the child leader's procedural hiccup.
- `human_model_signal`: Explicit supervision rule: `continue working until the task leader is done`.
- `failure_family_hypothesis`: `workflow_orchestration`
- `intervention_kind_hypothesis`: `reject_outcome`
- `human_cost_or_risk`: The parent supervisor can silently drop ownership while the delegated work is still live.
- `local_lesson_hypothesis`: In supervisory delegation, a concise status is not completion; completion is when the delegated worker actually reaches a real stop condition.
- `cluster_hints`: `supervision-liveness`, `ownership-boundary`, `premature-final`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: none material

### C15 - Stop UI polish and reconcile the impossible budget number

- `event_id`: `C15`
- `title`: `Stop UI polish and reconcile the impossible budget number`
- `session_or_thread`: `CodexDashboard overlay and semantics thread`
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\02\rollout-2026-04-02T21-40-28-019d50ff-b51b-7c82-9c73-207d0d24abf6.jsonl`
- `primary_refs`: `T07:L687`, `T07:L703`, `T07:L717`, `T07:L720`, `T07:L742`
- `ai_course`: The assistant was mid-way through a cosmetic UI patch that changed the palette and metric card while leaving an `8M` weekly budget number on screen.
- `human_intervention`: The human stopped the polish pass and said `8M budget` made no sense, could not be consistent with the displayed percentage used, and might need to be `8B`.
- `adequate_outcome`: Reconcile local config budget, advisory telemetry, and displayed scale before returning to cosmetic polishing.
- `event_boundary_notes`: PASS1 boundary is accurate. The assistant itself pivots immediately from cosmetic change to state-story math.
- `human_model_signal`: Explicit state-truth rule: if the budget number `doesn't make any sense` relative to the quota story, the semantic math problem outranks the cosmetic patch.
- `failure_family_hypothesis`: `usability_state_truth`
- `intervention_kind_hypothesis`: `reject_outcome`
- `human_cost_or_risk`: The dashboard could tell an impossible quota story on the actual surface.
- `local_lesson_hypothesis`: Fix impossible on-screen numbers before polishing the surrounding visual treatment.
- `cluster_hints`: `state-story-truth`, `metric-scale`, `quota-semantics`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: none material

### C16 - Reject the claimed layout fix while the dashboard was still clipped

- `event_id`: `C16`
- `title`: `Reject the claimed layout fix while the dashboard was still clipped`
- `session_or_thread`: `CodexDashboard overlay and semantics thread`
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\02\rollout-2026-04-02T21-40-28-019d50ff-b51b-7c82-9c73-207d0d24abf6.jsonl`
- `primary_refs`: `T07:L874`, `T07:L886`, `T07:L893`, `T07:L896`, `T07:L956`, `T07:L971`
- `ai_course`: The assistant had already said the layout fix was in, restarted the app, and explained the projection math and footer behavior.
- `human_intervention`: The human said the bottom was still cut off, then escalated from specific footer/button removals to `Remove entire bottom bar` and finally `remove bottom bar and leave room for graph`.
- `adequate_outcome`: Fix the live window the human is looking at, not the code-level explanation of why the footer should now fit.
- `event_boundary_notes`: PASS1 boundary is accurate. This is one visible-fidelity correction arc from `still cut off at bottom` through the stronger `remove bottom bar` directive.
- `human_model_signal`: Explicit visible-fit rule: `remove bottom bar and leave room for graph`.
- `failure_family_hypothesis`: `verification_proof`
- `intervention_kind_hypothesis`: `reject_outcome`
- `human_cost_or_risk`: The assistant could claim the layout was fixed while the live overlay remained clipped and visually wasteful.
- `local_lesson_hypothesis`: When the human still sees clipping, the fix is not done, whatever the layout rationale says on paper.
- `cluster_hints`: `live-surface-truth`, `footer-clipping`, `claimed-fix-vs-visible-result`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: none material

### C17 - Self-test the investigation launch path before asking the human again

- `event_id`: `C17`
- `title`: `Self-test the investigation launch path before asking the human again`
- `session_or_thread`: `CodexDashboard investigation launch debugging`
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\02\rollout-2026-04-02T21-40-28-019d50ff-b51b-7c82-9c73-207d0d24abf6.jsonl`
- `primary_refs`: `T07:L1586`, `T07:L1604`, `T07:L1612`, `T07:L1615`, `T07:L1631`
- `ai_course`: The assistant was still describing what should happen and was close to handing the next check back to the human after determining that no brief or PowerShell process had appeared.
- `human_intervention`: The human explicitly said `test it yourself by dispatching any command first before making me test again`.
- `adequate_outcome`: Prove the launch path end to end on the local machine before asking the human to retry the feature.
- `event_boundary_notes`: PASS1 boundary is accurate. The assistant responds by switching to machine-local end-to-end proof.
- `human_model_signal`: Explicit verification rule: `test it yourself ... before making me test again`.
- `failure_family_hypothesis`: `verification_proof`
- `intervention_kind_hypothesis`: `tighten_contract`
- `human_cost_or_risk`: The operator burden would bounce back to the human before the assistant had even proven the basic path locally.
- `local_lesson_hypothesis`: Do the first end-to-end launch proof yourself before using the human as the next debug step.
- `cluster_hints`: `operator-burden`, `self-verification-first`, `launch-path-proof`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: none material

### C18 - Reframe investigations from session summary to causal root cause

- `event_id`: `C18`
- `title`: `Reframe investigations from session summary to causal root cause`
- `session_or_thread`: `CodexDashboard investigation output tightening`
- `transcript_path`: `C:\Users\gregs\.codex\sessions\2026\04\02\rollout-2026-04-02T21-40-28-019d50ff-b51b-7c82-9c73-207d0d24abf6.jsonl`
- `primary_refs`: `T07:L1745`, `T07:L1772`, `T07:L1879`, `T07:L1906`, `T07:L1951`, `T07:L1962`, `T07:L1969`, `T07:L1972`
- `ai_course`: The assistant improved the investigation machinery enough to build briefs and recover more context, but the resulting output still mostly named sessions and repo buckets instead of answering what user action caused the spike and what mistake to avoid next time.
- `human_intervention`: The human first said `this doesn't seem like its answering anything`, then later said `i want to know what i did to cause the 100M token burst`, and finally tightened the contract to `root cause, no bullshit`.
- `adequate_outcome`: Investigation output must reconstruct the triggering operator action, mechanism, and avoidable mistake, not stop at session-summary evidence.
- `event_boundary_notes`: PASS1 intentionally kept this wide, and the transcript supports that. The same product-miss recurs across the brief shape, generated output, and exec prompt.
- `human_model_signal`: Explicit adequacy rule: the tool should answer what the human `did to cause the 100M token burst` and give `root cause, no bullshit`.
- `failure_family_hypothesis`: `usability_state_truth`
- `intervention_kind_hypothesis`: `reframe_problem`
- `human_cost_or_risk`: The investigation surface could look informative while still failing the human's actual causal question.
- `local_lesson_hypothesis`: If the product's job is investigation, require a causal answer about trigger, mechanism, and avoidable mistake; session summaries are supporting evidence, not the product.
- `cluster_hints`: `root-cause-not-summary`, `causal-investigation`, `operator-action-timeline`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: A later pass could split the shallow-evidence-pack problem from the still-too-report-y exec prompt, but one local event is sufficient for PASS2.

## Likely accepted incidents

- `C01` - Keep UI-fidelity ownership out of `AGENTS.md`
- `C04` - Reopen Task-0016 because prompt hardening was not real closure
- `C05` - Reject PASS-0007 closure while Sync shell gaps were still visible
- `C06` - Add post-implementation self-audit to `INTERFACE-DESIGNER`
- `C07` - Verify the installed build before trusting claimed UI fixes
- `C09` - Reject a technically working but network-invalid Mover fix
- `C11` - Enforce a background-safe, CombatMap, no-focus PIE debugging contract
- `C12` - Correct the literal scope of `push .codex`
- `C13` - Invalidate Sync-screen pass status when bottom-bar `Home` is dead
- `C14` - Keep supervising until the delegated task leader reaches a real endpoint
- `C15` - Stop UI polish and reconcile the impossible budget number
- `C16` - Reject the claimed layout fix while the dashboard was still clipped
- `C17` - Self-test the investigation launch path before asking the human again
- `C18` - Reframe investigations from session summary to causal root cause

## Likely non-incident but still important intervention events

- `C02` - Real diagnostic-scope correction, but it currently reads more like local forensics redirection than a durable accepted incident.
- `C03` - Strong human-semantics intervention, but it likely folds into the larger Task-0016 interface-designer and closure-truth cluster.
- `C08` - Real persistence correction against premature segmentation, but the more durable accepted candidate is `C09`.
- `C10` - Clear launcher-path correction, but it still looks like a local tooling slip rather than a top-tier accepted incident.

## Repeated cluster hints noticed across the analyzed set

- `closure-truth`
  - `C04`, `C05`, `C07`, `C13`, `C14`, `C16`
- `prompt-hardening`
  - `C03`, `C04`, `C05`, `C06`
- `live-surface-truth`
  - `C05`, `C07`, `C13`, `C15`, `C16`
- `control-boundary`
  - `C11`, `C14`, `C17`
- `repo-boundary`
  - `C01`, `C12`
- `root-cause-not-summary`
  - `C02`, `C08`, `C18`
- `status-signal-meaning`
  - `C03`, `C15`, `C18`

## Strongest human-model signals worth carrying into later clustering or principle work

- `C01`: Keep `AGENTS.md` light. Behavioral ownership belongs in the prompt layer that governs the responsible role, while verification detail can live in the repo testing doc.
- `C03`: Status language must say what is missing in ordinary human terms. Vague labels like `Not ready` leave too much room for wrong interpretation.
- `C04`, `C05`, `C07`, and `C16`: Workflow edits, code diffs, or narrated fixes do not satisfy closure if the human-visible emulator/app state still disagrees.
- `C11`: The user's control boundary is part of the bug contract. If they need the editor backgrounded so they can keep working, focus theft is not an acceptable control path.
- `C12`: Named scope words are literal. `push .codex` means `.codex`, not a widened commit of whatever else happens to be staged.
- `C14`: Supervision ends at a real delegated endpoint, not at the first concise status summary.
- `C15`: On-screen numbers must tell a coherent state story. If a metric is impossible relative to the rest of the surface, semantic truth outranks cosmetic polish.
- `C17`: The assistant should carry the first proof burden on local launch/debug paths before asking the human to retest.
- `C18`: An investigation feature should answer trigger, mechanism, and avoidable mistake. Session-summary output is evidence, not the product.

## Events that still need a wider reread

- `C11` only if a later pass wants to separate the wrong-map correction from the broader background-safe, no-focus control-boundary event.
- `C18` only if a later pass wants to separate the shallow-evidence-pack problem from the still-too-report-y exec prompt problem.
- No other candidate needs a wider reread for local PASS2 classification. The reopened transcript windows are sufficient for event-level diagnosis.
