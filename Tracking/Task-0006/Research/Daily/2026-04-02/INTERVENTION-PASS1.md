# INTERVENTION-PASS1

Source day: `2026-04-02`

## Source scope reviewed

- Primary source scope only: `C:\Users\gregs\.codex\sessions\2026\04\02\rollout-*.jsonl`
- Broad scan completed across all `35` source-day rollout files at raw `event_msg` / `user_message` level.
- Existing April 2 pass artifacts were not used as evidence.
- Worker transcripts that replay copied parent history were used only to detect duplicate boundaries; canonical refs point to the first raw source-day transcript where the intervention boundary is visible.
- Narrow rereads used for exact extraction:
- `T01` = `C:\Users\gregs\.codex\sessions\2026\04\02\rollout-2026-04-02T03-03-04-019d4d00-b1a4-7d93-9c26-e06ce205db13.jsonl`
- `T02` = `C:\Users\gregs\.codex\sessions\2026\04\02\rollout-2026-04-02T03-47-36-019d4d29-77ca-7c62-95d7-c2f2ed12cf09.jsonl`
- `T03` = `C:\Users\gregs\.codex\sessions\2026\04\02\rollout-2026-04-02T12-41-19-019d4f12-19e2-7fd3-8896-f51d93edbf9f.jsonl`
- `T04` = `C:\Users\gregs\.codex\sessions\2026\04\02\rollout-2026-04-02T16-59-35-019d4ffe-8bfa-7903-a5c4-7ad5eb0d6aef.jsonl`
- `T05` = `C:\Users\gregs\.codex\sessions\2026\04\02\rollout-2026-04-02T20-24-39-019d50ba-4bb6-74b2-8614-bb4a41483025.jsonl`
- `T06` = `C:\Users\gregs\.codex\sessions\2026\04\02\rollout-2026-04-02T21-40-28-019d50ff-b51b-7c82-9c73-207d0d24abf6.jsonl`
- Chronology note: some turns inside the `2026-04-02` file set continue past UTC midnight and even past local midnight. This pass stays bounded to the April 2 source-day files only.

## Total candidate intervention events found

- `19`

## Chronological candidate list

### C01 - Move UI-fidelity responsibility out of `AGENTS.md` and onto the real owner

- Session: `Crystallize doc and prompt rule update`
- Refs: `T01:L371`, `T01:L442`, `T01:L449`, `T01:L459`, `T01:L462`
- AI course: the assistant encoded the new emulator-vs-mockup fidelity rule in three places, including `AGENTS.md`.
- Human intervention: the human rejected that placement because `AGENTS.md` would become bloated, and pushed the explicit responsibility toward `TESTING.md` plus the task-leader prompt layer instead.
- Better outcome forced: keep repo front-door docs light while making task-leader ownership of UI fidelity explicit in the right durable contract.
- Why real: the human corrected both where the rule lives and who owns final mockup-fidelity sign-off.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C02 - Stop the quick subagent lever and demand full quota forensics

- Session: `Task-0016 quota forensics`
- Refs: `T02:L274`, `T02:L281`, `T02:L292`, `T02:L299`, `T02:L311`, `T02:L359`, `T02:L363`, `T02:L386`
- AI course: after misreading `used_percent` versus `remaining`, the assistant offered stopping Boyle as the main immediate token-saving lever.
- Human intervention: the human said `No don't touch anything` and demanded a wider diagnosis of everything burning quota, including reinstall side effects and missing session history.
- Better outcome forced: day-wide quota forensics rather than a quick child-session shutdown.
- Why real: the human halted an active operational suggestion and reframed the problem from one subagent to system-wide evidence recovery.
- Confidence: `strong`
- Triage: `intervention event but probably not an accepted incident`

### C03 - Rebuild the Home hero control from HTML composition instead of treating it like a loose icon

- Session: `Task-0016 main UI workflow`
- Refs: `T02:L1933`, `T02:L1936`, `T02:L1944`, `T02:L1951`
- AI course: the current Home pass was good overall, but still treated the record control too much like an icon-extraction problem.
- Human intervention: the human flagged the record button form factor as something that must preserve the mockup's graphical CTA feel, and accepted the assistant's shift toward rebuilding from the HTML composition.
- Better outcome forced: preserve the mockup's halo, glyph, spacing, and CTA form factor instead of approximating the control generically.
- Why real: the human corrected a screenshot-visible product-read issue, not merely a private implementation detail.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C04 - `Connected` reads like the wrong kind of status card

- Session: `Task-0016 main UI workflow`
- Refs: `T02:L2692`, `T02:L2699`, `T02:L2720`
- AI course: the assistant had just described remaining Home gaps, but the visible card still read ambiguously enough that the human interpreted it as possible server connectivity.
- Human intervention: the human asked what `Connected`, `Active now`, and `Me, Dad, 2 others ready` were even supposed to mean.
- Better outcome forced: make the card read as people/profile readiness rather than upload/server connectivity.
- Why real: the human had to correct the visible product read of the surface, not just wordsmith copy.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C05 - Missing rounded peer card, vague readiness copy, and icon drift force a new interface-designer lane

- Session: `Task-0016 main UI workflow`
- Refs: `T02:L2727`, `T02:L2730`, `T02:L2753`
- AI course: the current implementation and general workflow kept missing rounded-background parity, specific readiness wording, and faithful icon treatment.
- Human intervention: the human explicitly called out those misses and told the assistant to create and iterate an `INTERFACE-DESIGNER.md` role/prompt around them.
- Better outcome forced: a specialized workflow lane for mockup-visible UI fidelity and communication problems.
- Why real: the human corrected both the visible UI result and the workflow that kept letting those misses through.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C06 - Reject prompt-only progress and reopen the actual task until the UI loop is fixed

- Session: `Task-0016 main UI workflow`
- Refs: `T02:L2833`, `T02:L2840`, `T02:L2855`, `T02:L2870`
- AI course: the assistant reported small prompt-hardening wins as though the fidelity problem was meaningfully solved.
- Human intervention: the human asked whether any artifacts had actually changed and explicitly told the assistant to take the ball, revise the workflow/docs, and come back with a working emulator and the issues fixed.
- Better outcome forced: real reopened UI work rather than a prompt-only success story.
- Why real: the human had to reject premature closure on process tweaking alone.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C07 - Undo the Rustfire triage fix so the real networked cause can be solved

- Session: `Rustfire client/server debug`
- Refs: `T03:L1079`, `T03:L1086`, `T03:L1109`, `T03:L1213`, `T03:L1220`
- AI course: the assistant declared the connection path fixed after moving non-player pawns back to the standalone mover liaison.
- Human intervention: the human challenged that fix on engine/network semantics, got the assistant to admit the objection was correct, and then ordered the rollback so the real fix could be found.
- Better outcome forced: remove the architecturally wrong workaround and continue from the real seam.
- Why real: the human explicitly rejected a claimed fix and redirected debugging toward the right causal layer.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C08 - Call out the dropped Sync shell structure and force prompt diagnosis

- Session: `Task-0016 main UI workflow`
- Refs: `T02:L3704`, `T02:L3711`, `T02:L3725`, `T02:L3731`, `T02:L3770`
- AI course: `PASS-0007` was reported as done even though the screen still dropped top and bottom bars, chip wording, summary regions, and introduced a back button that the mockup did not have.
- Human intervention: the human called those `dropped balls` and required a probe of why the interface designer missed each one before the prompt was revised.
- Better outcome forced: treat shell structure, chip semantics, and missing mockup sections as blocking fidelity issues.
- Why real: the human rejected the visible surface after a done-claim and forced a diagnostic loop on the miss.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C09 - Catch the false `fixed` impression when the emulator is still showing the old build

- Session: `Task-0016 main UI workflow`
- Refs: `T02:L4125`, `T02:L4132`, `T02:L4145`
- AI course: the in-progress Sync pass had supposedly fixed persistence and back-button issues, but the installed emulator build had not actually been rebuilt or reinstalled from that pass.
- Human intervention: the human said nothing appeared changed and explicitly pointed to the still-missing shell bars and lingering back button.
- Better outcome forced: visible proof from the actual rebuilt app before claiming the fix exists.
- Why real: the human had to reject a claimed improvement because the real surface in front of them was still old.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C10 - Recover remaining Sync color, iconography, alignment, and cramped-button misses after the real build

- Session: `Task-0016 main UI workflow`
- Refs: `T02:L4204`, `T02:L4211`, `T02:L4227`, `T02:L4233`
- AI course: the rebuilt Sync screen was much better, but still diverged in top-bar color treatment, missing chip icons, CTA alignment, and cramped unreadable button text.
- Human intervention: the human named those visible divergences and asked for another designer-probe plus prompt revision cycle.
- Better outcome forced: continue until the real rebuilt screen reads like the mockup in color, iconography, alignment, and control legibility.
- Why real: these were screenshot-visible product-read misses, not abstract polish notes.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C11 - Treat Sync typography, chip sizing, and line spacing as legibility, not polish

- Session: `Task-0016 late UI workflow continuation`
- Refs: `T04:L811`, `T04:L833`, `T04:L852`, `T04:L858`, `T04:L861`
- AI course: the prior Sync pass improved state-story fidelity, but the visible type was still too thin and too small relative to the mockup, especially on a real phone form factor.
- Human intervention: the human explicitly called out legibility, thin weights, too-small chip text, and squished spacing, then asked for durable rules grounded in mobile standards and the HTML's relative type ratios.
- Better outcome forced: durable translation rules that treat typography and spacing as mockup fidelity and readability work.
- Why real: the human rejected the rendered text as hard to read for a human, not just aesthetically off.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C12 - Keep iterating because summary spacing, header rhythm, and bold weights still read cluttered

- Session: `Task-0016 late UI workflow continuation`
- Refs: `T04:L889`, `T04:L896`, `T04:L909`, `T04:L920`
- AI course: one typography pass had landed, but summary sections still felt dense and crowded.
- Human intervention: the human said `STORAGE USED` / `LAST SYNCED` still felt cluttered, wanted more space after headers, and explicitly asked for bolder bold text.
- Better outcome forced: match the mockup's spacing rhythm and text density, not just its approximate font sizes.
- Why real: the human had to intervene again because the first readability correction did not fully stick.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C13 - Reject the shared-row Home experiment and restore three distinct cards in the right order

- Session: `Task-0016 late UI workflow continuation`
- Refs: `T04:L1244`, `T04:L1251`, `T04:L1254`, `T04:L1268`
- AI course: the Home pass paired `Recent Activity` and `People setup` into a shared horizontal cluster to keep the layout compact.
- Human intervention: the human said the screens should not share the same horizontal area, required three separate cards, and fixed the order as `Recent Activity` then `People Setup` then `Waiting to upload`.
- Better outcome forced: a Home layout whose structure reads like the intended phone mockup, not a compact generic cluster.
- Why real: the human corrected the visible hierarchy and structural product read of the Home screen.
- Confidence: `medium`
- Triage: `intervention event but probably not an accepted incident`

### C14 - Reopen Task-0001 because the real hotkey did nothing on the actual machine

- Session: `Task-0001 hotkey regression`
- Refs: `T05:L580`, `T05:L587`, `T05:L632`
- AI course: the task state still treated the main remaining issue as a regression gate / confirmation problem.
- Human intervention: the human supplied direct manual evidence that the dashboard launched but `Ctrl+Alt+Space` did nothing.
- Better outcome forced: durable state and debugging must treat this as a real hotkey bug on the human-facing surface.
- Why real: new real-machine behavior contradicted the current blocker story and forced task reopening.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C15 - Stop the dashboard UI pass and fix the false budget math first

- Session: `Task-0002 dashboard build`
- Refs: `T06:L684`, `T06:L717`, `T06:L720`, `T06:L742`, `T06:L750`
- AI course: the assistant was in the middle of styling improvements while the dashboard still claimed an `8M` weekly budget that made no sense relative to the visible `% used`.
- Human intervention: the human halted the cosmetic work and called out the math as inconsistent and probably wrong by orders of magnitude.
- Better outcome forced: truthful budget semantics before more polish.
- Why real: the human rejected a visible numeric story as nonsense and redirected the work to state/story truth.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C16 - Footer clipping and control clutter still make the dashboard hard to read

- Session: `Task-0002 dashboard build`
- Refs: `T06:L841`, `T06:L848`, `T06:L886`, `T06:L893`, `T06:L904`
- AI course: the assistant had already patched the overlay height and explained the projection math, but the bottom text was still clipped and extra controls were still crowding the surface.
- Human intervention: the human said the footer was still cut off, simplified the controls, and questioned the headroom number.
- Better outcome forced: readable footer, simpler surface, and headroom math that lines up with the displayed projection model.
- Why real: the first layout/readability fix did not resolve the human-facing clipping and clutter problem.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C17 - Repo-mode stacked colors are unreadable for a human

- Session: `Task-0002 dashboard build`
- Refs: `T06:L1379`, `T06:L1386`, `T06:L1397`, `T06:L1402`
- AI course: the assistant had already changed the palette once, but the stacked repo bars still read as visually fused and hard to parse.
- Human intervention: the human explicitly said the colors were unreadable for a human and told the assistant to apply interface-design thinking rather than just minor hue variation.
- Better outcome forced: perceptually distinct stacked bars with boundaries that read at a glance.
- Why real: this was a direct readability/product-read rejection on the actual surface.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C18 - Bucket investigation flow stops at session summaries instead of root cause

- Session: `Task-0002 dashboard build`
- Refs: `T06:L1738`, `T06:L1745`, `T06:L1772`, `T06:L1782`
- AI course: the new right-click investigation flow still mostly produced bucket summaries and session lists instead of causal analysis.
- Human intervention: the human said the output wasn't answering anything useful and pushed for a real root-cause workflow.
- Better outcome forced: investigations that explain what happened, not just where to look.
- Why real: the human rejected the feature's current product output as inadequate for its core purpose.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C19 - Even the improved investigation still doesn't tell the user what action caused the 100M burst

- Session: `Task-0002 dashboard build`
- Refs: `T06:L1872`, `T06:L1879`, `T06:L1900`, `T06:L1906`
- AI course: the first investigation repair improved attribution and report generation, but still stopped short of telling the user what they actually did to trigger the spike.
- Human intervention: the human explicitly said they did not want to open the session and read it themselves, and demanded the causal user action behind the burst.
- Better outcome forced: causal explanation of the runaway pattern, not better metadata around the same summary.
- Why real: the human had to intervene again because the first investigation correction still failed the real product need.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

## Which candidates look like likely accepted incidents

- `C01` Move UI-fidelity responsibility out of `AGENTS.md` and onto the real owner
- `C03` Rebuild the Home hero control from HTML composition instead of treating it like a loose icon
- `C04` `Connected` reads like the wrong kind of status card
- `C05` Missing rounded peer card, vague readiness copy, and icon drift force a new interface-designer lane
- `C06` Reject prompt-only progress and reopen the actual task until the UI loop is fixed
- `C07` Undo the Rustfire triage fix so the real networked cause can be solved
- `C08` Call out the dropped Sync shell structure and force prompt diagnosis
- `C09` Catch the false `fixed` impression when the emulator is still showing the old build
- `C10` Recover remaining Sync color, iconography, alignment, and cramped-button misses after the real build
- `C11` Treat Sync typography, chip sizing, and line spacing as legibility, not polish
- `C12` Keep iterating because summary spacing, header rhythm, and bold weights still read cluttered
- `C14` Reopen Task-0001 because the real hotkey did nothing on the actual machine
- `C15` Stop the dashboard UI pass and fix the false budget math first
- `C16` Footer clipping and control clutter still make the dashboard hard to read
- `C17` Repo-mode stacked colors are unreadable for a human
- `C18` Bucket investigation flow stops at session summaries instead of root cause
- `C19` Even the improved investigation still doesn't tell the user what action caused the 100M burst

## Which candidates are real interventions but probably belong outside the accepted incident set

- `C02` Stop the quick subagent lever and demand full quota forensics
- `C13` Reject the shared-row Home experiment and restore three distinct cards in the right order

## Any ambiguous boundaries that need a second read

- `C03` is a real screenshot-visible fidelity correction, but it sits near collaborative design iteration rather than an outright claimed-closure rejection. It still looks durable enough to analyze further.
- `C06` is a real ownership/closure intervention, but PASS2 should decide whether it stands on its own or is better treated as part of the larger Sync prompt-hardening arc around `C05` and `C08`.
- `C13` may be an ordinary late design-iteration correction rather than a durable accepted incident, even though the human is clearly correcting visible hierarchy and product read.
