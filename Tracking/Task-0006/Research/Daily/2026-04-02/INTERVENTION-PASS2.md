# INTERVENTION-PASS2

Source day: `2026-04-02`

Canonical note: promoted on `2026-04-09` from the April 2 rerun; the earlier day-local PASS2 is archived as `INTERVENTION-PASS2-SUPERSEDED-2026-04-09.md`.

## Source scope analyzed

- PASS1 artifact:
  - `C:\Agent\CodexDashboard\Tracking\Task-0006\Research\Daily\2026-04-02\INTERVENTION-PASS1.md`
- Incident corpus contract reread:
  - `C:\Users\gregs\.codex\Orchestration\Reports\Incidents\README.md`
  - `C:\Users\gregs\.codex\Orchestration\Reports\Incidents\INCIDENT.schema.json`
- Prompt contracts reread:
  - `C:\Users\gregs\.codex\Orchestration\Prompts\INTERVENTION-PASS1.md`
  - `C:\Users\gregs\.codex\Orchestration\Prompts\INTERVENTION-PASS2.md`
  - `C:\Users\gregs\.codex\skills\codex-session-search\SKILL.md`
- Raw transcripts reopened from the PASS1 refs:
  - `T01` = `C:\Users\gregs\.codex\sessions\2026\04\02\rollout-2026-04-02T03-03-04-019d4d00-b1a4-7d93-9c26-e06ce205db13.jsonl`
  - `T02` = `C:\Users\gregs\.codex\sessions\2026\04\02\rollout-2026-04-02T03-47-36-019d4d29-77ca-7c62-95d7-c2f2ed12cf09.jsonl`
  - `T03` = `C:\Users\gregs\.codex\sessions\2026\04\02\rollout-2026-04-02T12-41-19-019d4f12-19e2-7fd3-8896-f51d93edbf9f.jsonl`
  - `T04` = `C:\Users\gregs\.codex\sessions\2026\04\02\rollout-2026-04-02T16-59-35-019d4ffe-8bfa-7903-a5c4-7ad5eb0d6aef.jsonl`
  - `T05` = `C:\Users\gregs\.codex\sessions\2026\04\02\rollout-2026-04-02T20-24-39-019d50ba-4bb6-74b2-8614-bb4a41483025.jsonl`
  - `T06` = `C:\Users\gregs\.codex\sessions\2026\04\02\rollout-2026-04-02T21-40-28-019d50ff-b51b-7c82-9c73-207d0d24abf6.jsonl`

## Candidate ids analyzed

- `C01` through `C19`

## Boundary corrections relative to PASS1

- No merge or split corrections were needed after transcript reread.
- `C08`, `C09`, and `C10` stay separate because they mark three distinct interventions in the same Sync-shell repair arc:
  - visible structure still wrong after a done-claim
  - false fix impression caused by an old installed build
  - remaining fidelity misses after a real rebuild
- `C11` and `C12` stay separate because the human first called out legibility and then had to keep pushing on spacing and rhythm after an initial typography pass.
- `C18` and `C19` stay separate because the first intervention rejected summary-only investigation, while the second rejected the still-insufficient improved version for failing to explain the causal user action.

## Per-event analysis records

### C01 - Move UI-fidelity responsibility out of `AGENTS.md` and onto the real owner

- `event_id`: `C01`
- `title`: `Move UI-fidelity responsibility out of AGENTS.md and onto the real owner`
- `session_or_thread`: `Sync docs and design responsibility`
- `transcript_path`: `T01`
- `primary_refs`: `T01:L371`, `T01:L442`, `T01:L449`, `T01:L459`, `T01:L462`
- `ai_course`: The assistant encoded the UI-fidelity rule into `AGENTS.md`, `TESTING.md`, and `GENERAL-DESIGN.md`, treating broad doc reinforcement as the right fix.
- `human_intervention`: The human said `AGENTS.md` was bloated, rejected stuffing more there, and redirected responsibility to the actual role or prompt that owns UI judgment.
- `adequate_outcome`: Keep the fidelity rule durable, but attach it to the real decision owner instead of spamming generic front-door docs.
- `event_boundary_notes`: Single event with a short correction arc. PASS1 boundary holds.
- `human_model_signal`: Explicit responsibility-boundary rule: durable guidance should live with the actor that must apply it, not in overgrown generic docs.
- `failure_family_hypothesis`: `workflow_orchestration`
- `intervention_kind_hypothesis`: `reframe_problem`
- `human_cost_or_risk`: Doc bloat plus continued misses because the real owner still would not see the rule at the moment of judgment.
- `local_lesson_hypothesis`: Put corrective UI-fidelity guidance on the role or workflow seam that actually makes the call.
- `cluster_hints`: `operator-boundary`, `durable-contract`, `responsibility-routing`
- `accepted_incident_likelihood`: `possible`
- `confidence`: `strong`
- `uncertainties`: This is clearly a real intervention, but it is more about durable routing than a direct user-surface failure.

### C02 - Stop the quick subagent lever and demand full quota forensics

- `event_id`: `C02`
- `title`: `Stop the quick subagent lever and demand full quota forensics`
- `session_or_thread`: `Boyle quota investigation`
- `transcript_path`: `T02`
- `primary_refs`: `T02:L274`, `T02:L281`, `T02:L292`, `T02:L299`, `T02:L311`, `T02:L359`, `T02:L363`, `T02:L386`
- `ai_course`: The assistant misread used versus remaining quota and jumped toward stopping Boyle early instead of investigating the full state.
- `human_intervention`: The human said `No don't touch anything`, demanded a full look at everything, and rejected acting on the first shallow interpretation.
- `adequate_outcome`: Investigate the real quota picture before changing behavior or shutting work down.
- `event_boundary_notes`: Clean intervention. PASS1 boundary holds.
- `human_model_signal`: Explicit adequacy rule: do not act on quota fears until the whole state has been inspected.
- `failure_family_hypothesis`: `verification_proof`
- `intervention_kind_hypothesis`: `redirect_debugging`
- `human_cost_or_risk`: A premature stop would have disrupted real work based on a false read.
- `local_lesson_hypothesis`: Separate first-pass interpretation from action when the underlying telemetry can still be checked directly.
- `cluster_hints`: `wrong-seam-debugging`, `state-truth`, `premature-action`
- `accepted_incident_likelihood`: `unlikely`
- `confidence`: `strong`
- `uncertainties`: Important locally, but it looks more like a contained diagnostic correction than a likely accepted incident.

### C03 - Rebuild the Home hero control from HTML composition instead of treating it like a loose icon

- `event_id`: `C03`
- `title`: `Rebuild the Home hero control from HTML composition instead of treating it like a loose icon`
- `session_or_thread`: `Record button fidelity`
- `transcript_path`: `T02`
- `primary_refs`: `T02:L1933`, `T02:L1936`, `T02:L1944`, `T02:L1951`
- `ai_course`: The assistant was treating the hero record control too much like an extracted icon asset instead of preserving the full mockup form factor.
- `human_intervention`: The human said the Home hero should preserve the mockup's shape and composition and pushed the assistant toward rebuilding it from HTML structure.
- `adequate_outcome`: Preserve the intended button/container geometry and visual read from the mockup, not just the icon glyph.
- `event_boundary_notes`: PASS1 boundary holds. This is a genuine fidelity intervention, but it sits close to ordinary live design iteration.
- `human_model_signal`: Explicit mockup-fidelity rule: the visible control's full composition matters, not only the icon.
- `failure_family_hypothesis`: `usability_state_truth`
- `intervention_kind_hypothesis`: `tighten_contract`
- `human_cost_or_risk`: The product would visibly drift from the intended hero interaction and read as lower fidelity on the main surface.
- `local_lesson_hypothesis`: When a mockup-visible hero control carries product meaning through shape and framing, translate the whole composition, not just the icon asset.
- `cluster_hints`: `mockup-fidelity`, `hero-control`, `visible-surface-truth`
- `accepted_incident_likelihood`: `possible`
- `confidence`: `medium`
- `uncertainties`: This may still be ordinary design iteration rather than a clearly durable incident, even though the human treated it as important.

### C04 - `Connected` reads like the wrong kind of status card

- `event_id`: `C04`
- `title`: `Connected reads like the wrong kind of status card`
- `session_or_thread`: `Sync card product read`
- `transcript_path`: `T02`
- `primary_refs`: `T02:L2692`, `T02:L2699`, `T02:L2720`
- `ai_course`: The assistant presented a `Connected` card whose meaning was unclear enough that the human could not tell what real product state it conveyed.
- `human_intervention`: The human directly asked what `Connected` actually means, signaling that the visible card semantics were wrong or underspecified.
- `adequate_outcome`: A status card whose meaning is self-evident to a human without opening logs or repo context.
- `event_boundary_notes`: Tight event. PASS1 boundary holds.
- `human_model_signal`: Explicit product-read rule: visible status cards should tell a human what state they are in, not force interpretive work.
- `failure_family_hypothesis`: `ui_semantics`
- `intervention_kind_hypothesis`: `reframe_problem`
- `human_cost_or_risk`: Human confusion on a core visible state surface.
- `local_lesson_hypothesis`: If a visible status card triggers `what does that even mean`, its semantics are not yet done.
- `cluster_hints`: `status-signal-meaning`, `product-read`, `visible-surface-truth`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: none material

### C05 - Missing rounded peer card, vague readiness copy, and icon drift force a new interface-designer lane

- `event_id`: `C05`
- `title`: `Missing rounded peer card, vague readiness copy, and icon drift force a new interface-designer lane`
- `session_or_thread`: `Sync card parity and ownership`
- `transcript_path`: `T02`
- `primary_refs`: `T02:L2727`, `T02:L2730`, `T02:L2753`
- `ai_course`: The assistant had visible peer-card parity, copy quality, and iconography drift problems but no crisp owner for catching them.
- `human_intervention`: The human called out the missing rounded companion card, weak `Not ready yet` copy, and icon mismatch, then forced creation of an `INTERFACE-DESIGNER` lane.
- `adequate_outcome`: Treat screenshot-visible parity, copy clarity, and iconography fidelity as owned quality gates, not incidental polish.
- `event_boundary_notes`: PASS1 boundary holds. This is partly a role-routing fix, but it was triggered by concrete visible misses.
- `human_model_signal`: Explicit rule: screenshot-visible fidelity problems count as real misses that deserve a dedicated reviewer/owner.
- `failure_family_hypothesis`: `usability_state_truth`
- `intervention_kind_hypothesis`: `tighten_contract`
- `human_cost_or_risk`: Repeated visible drift and continued dependence on the human to notice basic product-read problems.
- `local_lesson_hypothesis`: When the same class of screenshot-visible UI misses recurs, create an explicit lane that owns fidelity and product-read review.
- `cluster_hints`: `mockup-fidelity`, `owner-gap`, `iconography`, `copy-clarity`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: none material

### C06 - Reject prompt-only progress and reopen the actual task until the UI loop is fixed

- `event_id`: `C06`
- `title`: `Reject prompt-only progress and reopen the actual task until the UI loop is fixed`
- `session_or_thread`: `Task closure versus real UI work`
- `transcript_path`: `T02`
- `primary_refs`: `T02:L2833`, `T02:L2840`, `T02:L2855`, `T02:L2870`
- `ai_course`: The assistant had moved work into prompts and role docs, but the human-facing UI loop was still not actually fixed.
- `human_intervention`: The human asked whether artifacts had changed, rejected prompt-only motion as closure, and told the assistant to take the ball until the emulator-side loop was really repaired.
- `adequate_outcome`: Reopen the concrete task and drive the visible UI work forward instead of treating prompt edits as sufficient.
- `event_boundary_notes`: PASS1 boundary holds. This is adjacent to `C01` but later and more outcome-focused.
- `human_model_signal`: Explicit closure rule: durable prompt changes do not count if the real visible workflow is still broken.
- `failure_family_hypothesis`: `workflow_orchestration`
- `intervention_kind_hypothesis`: `reject_outcome`
- `human_cost_or_risk`: False closure and stalled product repair.
- `local_lesson_hypothesis`: Treat prompt and doc changes as support artifacts; closure still depends on the real surface being fixed.
- `cluster_hints`: `closure-truth`, `artifact-vs-surface`, `real-work-ownership`
- `accepted_incident_likelihood`: `possible`
- `confidence`: `strong`
- `uncertainties`: It is clearly important, but it may collapse later into a broader closure-truth cluster rather than stand alone as an accepted incident.

### C07 - Undo the Rustfire triage fix so the real networked cause can be solved

- `event_id`: `C07`
- `title`: `Undo the Rustfire triage fix so the real networked cause can be solved`
- `session_or_thread`: `Rustfire liaison regression`
- `transcript_path`: `T03`
- `primary_refs`: `T03:L1079`, `T03:L1086`, `T03:L1109`, `T03:L1213`, `T03:L1220`
- `ai_course`: The assistant claimed a fix by moving pawns back to a standalone liaison setup, effectively dodging the real networked behavior.
- `human_intervention`: The human objected that this was the wrong semantic fix for the actual networked problem, and the assistant eventually admitted it was wrong and rolled it back.
- `adequate_outcome`: Solve the networked cause instead of redefining the scenario until the bug disappears.
- `event_boundary_notes`: PASS1 boundary holds.
- `human_model_signal`: Explicit rule: a workaround that changes the target behavior is not an adequate fix for a networked bug.
- `failure_family_hypothesis`: `verification_proof`
- `intervention_kind_hypothesis`: `redirect_debugging`
- `human_cost_or_risk`: Time lost on the wrong seam and risk of declaring victory over a non-equivalent scenario.
- `local_lesson_hypothesis`: Keep the repro's behavior contract intact while debugging; do not "fix" by moving the problem into a different mode.
- `cluster_hints`: `wrong-seam-debugging`, `closure-truth`, `scenario-equivalence`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: none material

### C08 - Call out the dropped Sync shell structure and force prompt diagnosis

- `event_id`: `C08`
- `title`: `Call out the dropped Sync shell structure and force prompt diagnosis`
- `session_or_thread`: `Sync shell fidelity after done-claim`
- `transcript_path`: `T02`
- `primary_refs`: `T02:L3704`, `T02:L3711`, `T02:L3725`, `T02:L3731`, `T02:L3770`
- `ai_course`: The assistant had claimed the screen was done, but the visible shell still lacked bars and summary regions and still had chip-text and back-button drift.
- `human_intervention`: The human called these `dropped balls`, asked why the designer lane missed them, and forced diagnosis of the prompt/role failure.
- `adequate_outcome`: Treat missing visible shell structure and obvious fidelity drift as real task failures that must be caught before done-claims.
- `event_boundary_notes`: Distinct from `C09` and `C10`; this is the first strong rejection of the visible done-claim.
- `human_model_signal`: Explicit UI-fidelity rule: screenshot-visible missing bars, layout regions, and basic icon/text structure are not polish.
- `failure_family_hypothesis`: `usability_state_truth`
- `intervention_kind_hypothesis`: `reject_outcome`
- `human_cost_or_risk`: False confidence, extra human review burden, and continued visible product drift.
- `local_lesson_hypothesis`: If the screenshot still reads structurally incomplete, the work is not done no matter how much invisible progress occurred.
- `cluster_hints`: `mockup-fidelity`, `closure-truth`, `visible-surface-truth`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: none material

### C09 - Catch the false `fixed` impression when the emulator is still showing the old build

- `event_id`: `C09`
- `title`: `Catch the false fixed impression when the emulator is still showing the old build`
- `session_or_thread`: `Sync rebuild verification`
- `transcript_path`: `T02`
- `primary_refs`: `T02:L4125`, `T02:L4132`, `T02:L4145`
- `ai_course`: The assistant thought the UI had been fixed, but the visible emulator state had not changed because the old build was still installed.
- `human_intervention`: The human said nothing had changed, forcing the assistant to confront the difference between source changes and the actually installed surface.
- `adequate_outcome`: Verify the real build state before claiming visible UI fixes.
- `event_boundary_notes`: Distinct from `C08` because this is specifically about proof on the installed surface.
- `human_model_signal`: Explicit proof rule: if the human cannot see the claimed fix on the actual build, it is not yet a real fix.
- `failure_family_hypothesis`: `verification_proof`
- `intervention_kind_hypothesis`: `reject_outcome`
- `human_cost_or_risk`: False fix claims and wasted review cycles against the wrong installed artifact.
- `local_lesson_hypothesis`: Close the loop on deployed state before narrating visual success.
- `cluster_hints`: `proof-scope`, `installed-surface-truth`, `closure-truth`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: none material

### C10 - Recover remaining Sync color, iconography, alignment, and cramped-button misses after the real build

- `event_id`: `C10`
- `title`: `Recover remaining Sync color, iconography, alignment, and cramped-button misses after the real build`
- `session_or_thread`: `Sync fidelity cleanup after rebuild`
- `transcript_path`: `T02`
- `primary_refs`: `T02:L4204`, `T02:L4211`, `T02:L4227`, `T02:L4233`
- `ai_course`: After the real rebuild, the assistant still had top-color, iconography, alignment, and button/readability misses on the visible screen.
- `human_intervention`: The human called out the remaining visible problems directly rather than accepting the rebuilt screen as good enough.
- `adequate_outcome`: Keep iterating until the rebuilt screen also satisfies screenshot-visible fidelity and readability.
- `event_boundary_notes`: Distinct from `C08` because it happens after the real build is finally on-screen.
- `human_model_signal`: Explicit rule: a rebuilt surface can still be wrong if iconography, hierarchy, alignment, or readability drift remains visible.
- `failure_family_hypothesis`: `usability_state_truth`
- `intervention_kind_hypothesis`: `tighten_contract`
- `human_cost_or_risk`: Visible product-read degradation survives the technical deploy loop.
- `local_lesson_hypothesis`: Treat post-build screenshot review as a real acceptance gate, not a ceremonial last glance.
- `cluster_hints`: `mockup-fidelity`, `iconography`, `readability`, `visible-surface-truth`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: none material

### C11 - Treat Sync typography, chip sizing, and line spacing as legibility, not polish

- `event_id`: `C11`
- `title`: `Treat Sync typography, chip sizing, and line spacing as legibility, not polish`
- `session_or_thread`: `Sync typography legibility`
- `transcript_path`: `T04`
- `primary_refs`: `T04:L811`, `T04:L833`, `T04:L852`, `T04:L858`, `T04:L861`
- `ai_course`: The assistant's Sync screen still had thin or small type, illegible chip text, and squished line spacing.
- `human_intervention`: The human reframed these as legibility and human readability issues and asked for durable rules grounded in mobile standards and HTML ratios.
- `adequate_outcome`: Typography and spacing must preserve readability on the real mobile surface, not merely resemble a compact mockup.
- `event_boundary_notes`: PASS1 boundary holds.
- `human_model_signal`: Strong explicit rule: screenshot-visible typography and spacing that distort human readability count as real product failures.
- `failure_family_hypothesis`: `usability_state_truth`
- `intervention_kind_hypothesis`: `reframe_problem`
- `human_cost_or_risk`: A human-facing screen that is harder to read and therefore less trustworthy or usable.
- `local_lesson_hypothesis`: Translate mockup ratios into readable mobile typography, rather than copying density at the expense of legibility.
- `cluster_hints`: `typography`, `spacing`, `readability`, `human-facing-quality`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: none material

### C12 - Keep iterating because summary spacing, header rhythm, and bold weights still read cluttered

- `event_id`: `C12`
- `title`: `Keep iterating because summary spacing, header rhythm, and bold weights still read cluttered`
- `session_or_thread`: `Sync readability second pass`
- `transcript_path`: `T04`
- `primary_refs`: `T04:L889`, `T04:L896`, `T04:L909`, `T04:L920`
- `ai_course`: After an initial typography adjustment, the assistant still left the summary area cluttered through spacing, rhythm, and emphasis choices.
- `human_intervention`: The human asked for more space and stronger bolding, making clear that the screen still did not read cleanly enough.
- `adequate_outcome`: Keep working until hierarchy and spacing make the summary easy to scan.
- `event_boundary_notes`: Separate from `C11`; this is the second intervention after the first readability pass.
- `human_model_signal`: Explicit human-facing rule: readability depends on rhythm and visual hierarchy, not just font size.
- `failure_family_hypothesis`: `usability_state_truth`
- `intervention_kind_hypothesis`: `tighten_contract`
- `human_cost_or_risk`: Continued clutter on a summary surface intended for fast human comprehension.
- `local_lesson_hypothesis`: A readability fix is incomplete if hierarchy and spacing still make the surface feel packed or muddy.
- `cluster_hints`: `spacing`, `hierarchy`, `readability`, `summary-scanability`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: none material

### C13 - Reject the shared-row Home experiment and restore three distinct cards in the right order

- `event_id`: `C13`
- `title`: `Reject the shared-row Home experiment and restore three distinct cards in the right order`
- `session_or_thread`: `Home card layout experiment`
- `transcript_path`: `T04`
- `primary_refs`: `T04:L1244`, `T04:L1251`, `T04:L1254`, `T04:L1268`
- `ai_course`: The assistant had tried a shared-row treatment for Home content instead of the separate cards the human expected.
- `human_intervention`: The human rejected the experiment and specified the exact three-card order: `Recent Activity -> People Setup -> Waiting to upload`.
- `adequate_outcome`: Restore the intended information architecture and visual separation on Home.
- `event_boundary_notes`: PASS1 boundary holds, but this remains close to ordinary design iteration.
- `human_model_signal`: Explicit structure rule: card separation and ordering are part of the product read, not optional styling.
- `failure_family_hypothesis`: `information_architecture`
- `intervention_kind_hypothesis`: `tighten_contract`
- `human_cost_or_risk`: Home reads incorrectly and loses the intended priority/order of information.
- `local_lesson_hypothesis`: Preserve card separation and order when they carry product meaning on the landing surface.
- `cluster_hints`: `information-architecture`, `home-surface`, `product-read`
- `accepted_incident_likelihood`: `possible`
- `confidence`: `medium`
- `uncertainties`: This may be an ordinary late design correction rather than a durable incident unless paired with broader repeated evidence.

### C14 - Reopen Task-0001 because the real hotkey did nothing on the actual machine

- `event_id`: `C14`
- `title`: `Reopen Task-0001 because the real hotkey did nothing on the actual machine`
- `session_or_thread`: `CodexDashboard overlay regression`
- `transcript_path`: `T05`
- `primary_refs`: `T05:L580`, `T05:L587`, `T05:L632`
- `ai_course`: The task had effectively been treated as working, but on the user's real machine `Ctrl+Alt+Space` still did nothing.
- `human_intervention`: The human reported the real regression result and forced the task to reopen as a genuine bug.
- `adequate_outcome`: Treat live regression on the actual machine as authoritative over prior assumptions or partial proof.
- `event_boundary_notes`: PASS1 boundary holds.
- `human_model_signal`: Explicit proof rule: if the real human-facing entrypoint does nothing, the feature is not working regardless of supporting artifacts.
- `failure_family_hypothesis`: `verification_proof`
- `intervention_kind_hypothesis`: `reject_outcome`
- `human_cost_or_risk`: False closure on a broken primary interaction.
- `local_lesson_hypothesis`: Weight real-machine behavior above indirect proof when closing a user-triggered desktop feature.
- `cluster_hints`: `installed-surface-truth`, `closure-truth`, `real-machine-validation`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: none material

### C15 - Stop the dashboard UI pass and fix the false budget math first

- `event_id`: `C15`
- `title`: `Stop the dashboard UI pass and fix the false budget math first`
- `session_or_thread`: `CodexDashboard budget truth`
- `transcript_path`: `T06`
- `primary_refs`: `T06:L684`, `T06:L717`, `T06:L720`, `T06:L742`, `T06:L750`
- `ai_course`: The assistant was still doing dashboard UI work while the visible weekly budget story showed nonsense like `8M` against a visible percent-used concept.
- `human_intervention`: The human stopped cosmetic work and redirected the effort to fixing the truth of the displayed numbers.
- `adequate_outcome`: Make the visible budget math truthful before polishing the surrounding UI.
- `event_boundary_notes`: PASS1 boundary holds.
- `human_model_signal`: Explicit product rule: visible dashboard numbers must tell the truth before presentation polish matters.
- `failure_family_hypothesis`: `usability_state_truth`
- `intervention_kind_hypothesis`: `reframe_problem`
- `human_cost_or_risk`: Humans make decisions from false telemetry while the UI appears polished.
- `local_lesson_hypothesis`: On a dashboard, truth of the visible metric outranks cosmetic refinement.
- `cluster_hints`: `state-truth`, `dashboard-semantics`, `truth-before-polish`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: none material

### C16 - Footer clipping and control clutter still make the dashboard hard to read

- `event_id`: `C16`
- `title`: `Footer clipping and control clutter still make the dashboard hard to read`
- `session_or_thread`: `CodexDashboard visible layout pass`
- `transcript_path`: `T06`
- `primary_refs`: `T06:L841`, `T06:L848`, `T06:L886`, `T06:L893`, `T06:L904`
- `ai_course`: The assistant had improved the dashboard but still left the bottom clipped and the controls crowded enough to look wrong.
- `human_intervention`: The human said it was still cut off at the bottom and still cluttered, rejecting the current visible state.
- `adequate_outcome`: A dashboard that fits the window and reads clearly without clipping or crowded control density.
- `event_boundary_notes`: PASS1 boundary holds.
- `human_model_signal`: Explicit readability rule: clipping and crowding are real usability-state failures on the visible surface.
- `failure_family_hypothesis`: `usability_state_truth`
- `intervention_kind_hypothesis`: `reject_outcome`
- `human_cost_or_risk`: Hard-to-read UI and reduced trust in the polish and accuracy of the tool.
- `local_lesson_hypothesis`: Treat clipping and control crowding as acceptance failures, not cleanup chores.
- `cluster_hints`: `clipping`, `readability`, `visible-surface-truth`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: none material

### C17 - Repo-mode stacked colors are unreadable for a human

- `event_id`: `C17`
- `title`: `Repo-mode stacked colors are unreadable for a human`
- `session_or_thread`: `CodexDashboard repo mode readability`
- `transcript_path`: `T06`
- `primary_refs`: `T06:L1379`, `T06:L1386`, `T06:L1397`, `T06:L1402`
- `ai_course`: The assistant's repo-mode color treatment stacked colors in a way that the human found unreadable.
- `human_intervention`: The human explicitly said the colors were unreadable for a human.
- `adequate_outcome`: Use contrast and layering that keep repo-mode information human-readable.
- `event_boundary_notes`: PASS1 boundary holds.
- `human_model_signal`: Strong explicit rule: human readability is the bar, and stacked-color styling can fail that bar even if technically rendered correctly.
- `failure_family_hypothesis`: `usability_state_truth`
- `intervention_kind_hypothesis`: `reframe_problem`
- `human_cost_or_risk`: Reduced readability on a dense information surface.
- `local_lesson_hypothesis`: Evaluate color treatments by whether a human can actually read the information, not by whether the palette is internally consistent.
- `cluster_hints`: `readability`, `color-contrast`, `human-facing-quality`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: none material

### C18 - Bucket investigation flow stops at session summaries instead of root cause

- `event_id`: `C18`
- `title`: `Bucket investigation flow stops at session summaries instead of root cause`
- `session_or_thread`: `CodexDashboard bucket investigation`
- `transcript_path`: `T06`
- `primary_refs`: `T06:L1738`, `T06:L1745`, `T06:L1772`, `T06:L1782`
- `ai_course`: The assistant's investigation output summarized sessions and buckets but did not explain the root cause of the huge burst.
- `human_intervention`: The human rejected the summary-only output and asked for actual causal investigation.
- `adequate_outcome`: Explain why the bucket spiked, not merely which sessions appear in it.
- `event_boundary_notes`: PASS1 boundary holds. Distinct from `C19`, which is the next failed repair.
- `human_model_signal`: Explicit investigation rule: a causal analysis tool fails if it stops at summaries rather than root cause.
- `failure_family_hypothesis`: `human_world`
- `intervention_kind_hypothesis`: `redirect_debugging`
- `human_cost_or_risk`: More human detective work to recover meaning from telemetry that should already explain itself.
- `local_lesson_hypothesis`: Investigation surfaces should answer the user's causal question directly before presenting supporting summaries.
- `cluster_hints`: `diagnostic-scope`, `causal-explanation`, `analysis-surface`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: none material

### C19 - Even the improved investigation still does not say what action caused the 100M burst

- `event_id`: `C19`
- `title`: `Even the improved investigation still does not say what action caused the 100M burst`
- `session_or_thread`: `CodexDashboard bucket investigation second rejection`
- `transcript_path`: `T06`
- `primary_refs`: `T06:L1872`, `T06:L1879`, `T06:L1900`, `T06:L1906`
- `ai_course`: After the first correction, the assistant improved the investigation output but still did not tell the human what user action or workflow caused the spike.
- `human_intervention`: The human said that is what Codex is for, making clear that causal attribution to a concrete action is the real deliverable.
- `adequate_outcome`: Name the action or workflow pattern that caused the 100M burst.
- `event_boundary_notes`: Separate from `C18` because the assistant attempted a repair and still missed the product goal.
- `human_model_signal`: Strong explicit product rule: investigation tooling should recover the actionable human-world cause, not stop at prettier summaries.
- `failure_family_hypothesis`: `human_world`
- `intervention_kind_hypothesis`: `tighten_contract`
- `human_cost_or_risk`: The human still has to perform the core interpretation manually.
- `local_lesson_hypothesis`: A telemetry investigation experience is incomplete until it can explain the causal action a human actually cares about.
- `cluster_hints`: `diagnostic-scope`, `causal-explanation`, `human-world-meaning`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: none material

## Likely accepted incidents

- `C04` - visible status-card semantics were not self-explanatory
- `C05` - repeated screenshot-visible card/copy/icon drift forced a dedicated fidelity lane
- `C07` - wrong-seam Rustfire fix changed the target behavior instead of solving it
- `C08` - done-claim collapsed under obvious screenshot-visible shell omissions
- `C09` - claimed fix was not present on the actually installed build
- `C10` - real rebuild still left visible color, iconography, alignment, and readability misses
- `C11` - typography and spacing were corrected as legibility failures, not polish
- `C12` - summary spacing and hierarchy still required explicit human rescue
- `C14` - real-machine hotkey regression invalidated assumed closure
- `C15` - visible budget math was false while UI work continued
- `C16` - clipping and crowding still broke the visible dashboard surface
- `C17` - repo-mode color treatment was unreadable for a human
- `C18` - investigation flow stopped at summaries instead of root cause
- `C19` - even after repair, investigation still failed to name the causal user action

## Likely non-incident but still important intervention events

- `C02` - premature quota action from a shallow state read
- `C03` - hero-control mockup fidelity correction
- `C06` - rejection of prompt-only progress in place of real UI repair
- `C13` - Home card-order and separation correction

## UI-fidelity recovery check

- Screenshot-visible or mockup-visible UI fidelity issues were clearly recovered as PASS1 candidate events.
- The strongest recovered fidelity/readability candidates are:
  - `C03`, `C04`, `C05`, `C08`, `C09`, `C10`, `C11`, `C12`, `C13`, `C16`, `C17`
- Among those, the events that look `likely` accepted incidents are:
  - `C04`, `C05`, `C08`, `C09`, `C10`, `C11`, `C12`, `C16`, `C17`
- The main fidelity candidates that still feel more like ordinary live design iteration than durable accepted incidents are:
  - `C03`
  - `C13`

## Repeated cluster hints noticed across the analyzed set

- `mockup-fidelity` / `visible-surface-truth`:
  - `C03`, `C04`, `C05`, `C08`, `C10`, `C13`
- `typography` / `spacing` / `readability`:
  - `C11`, `C12`, `C16`, `C17`
- `closure-truth` / `installed-surface-truth`:
  - `C06`, `C08`, `C09`, `C14`
- `state-truth` / `dashboard-semantics`:
  - `C04`, `C15`
- `wrong-seam-debugging`:
  - `C02`, `C07`
- `diagnostic-scope` / `causal-explanation`:
  - `C18`, `C19`

## Strongest human-model signals worth carrying forward

- Screenshot-visible typography, spacing, iconography, card shell, and clipping can materially distort the intended human-facing outcome and therefore count as real state-truth failures.
- Mockup fidelity should be translated into readable mobile/UI ratios, not flattened into asset extraction or density for density's sake.
- If the human cannot see the claimed fix on the actually installed or running surface, the fix is not yet real.
- Status cards and investigation views should explain human-meaningful state directly, not force the human to decode them from logs, docs, or session summaries.
- Dashboard and telemetry surfaces must tell the truth before surrounding presentation polish matters.

## Events that still need a wider reread

- `C02` because it is clearly a real intervention but may be too small for the accepted set.
- `C03` because its importance is obvious, but the line between durable incident and ordinary design iteration remains fuzzy.
- `C13` because information architecture and card-order corrections often sit close to normal live design iteration unless there is broader repeated evidence.
