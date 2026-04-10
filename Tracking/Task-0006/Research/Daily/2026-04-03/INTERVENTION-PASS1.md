# INTERVENTION-PASS1

Source day: `2026-04-03`

## Source scope reviewed

- Primary source scope only: `C:\Users\gregs\.codex\sessions\2026\04\03\rollout-*.jsonl`
- Broad scan completed across all `28` rollout files in that source-day folder, covering `271` outbound human messages.
- Exact extraction aliases used for the candidate list:
  - `T10` = `C:\Users\gregs\.codex\sessions\2026\04\03\rollout-2026-04-03T10-11-16-019d53af-14f1-7a80-987d-cda22ee69505.jsonl`
  - `T11` = `C:\Users\gregs\.codex\sessions\2026\04\03\rollout-2026-04-03T11-35-01-019d53fb-c3db-7e61-8adf-a4ff171c0de9.jsonl`
  - `T12` = `C:\Users\gregs\.codex\sessions\2026\04\03\rollout-2026-04-03T12-28-10-019d542c-6ab7-7dc2-833c-5acc5a904748.jsonl`
  - `T20` = `C:\Users\gregs\.codex\sessions\2026\04\03\rollout-2026-04-03T20-51-56-019d55f9-a32a-7241-bb7b-2fa9afd3b849.jsonl`
  - `T22` = `C:\Users\gregs\.codex\sessions\2026\04\03\rollout-2026-04-03T22-52-48-019d5668-49a1-7563-88c0-47b562ff8120.jsonl`
- Confirmatory duplicate-read only, not primary refs:
  - `T23` = `C:\Users\gregs\.codex\sessions\2026\04\03\rollout-2026-04-03T23-13-44-019d567b-736f-7190-92fd-89e2bee70457.jsonl`
  - `T23` restated the same ThirdPerson planning and implementation corrections inside a spawned task-leader thread, but the primary refs below stay on the originating human turns in `T22`.
- Scope note: several rollout files in the `2026\04\03` source-day folder continue into later absolute timestamps on `2026-04-04`. This pass stayed folder-bounded, not strict-timestamp-bounded.
- Existing accepted-incident JSON, daily README notes, and legacy CSV artifacts were not used as source-of-truth inputs for candidate extraction.

## Total candidate intervention events found

- `13`

## Chronological candidate list

### C01 - Deleting the Windows sandbox block created a blocker instead of a safe Windows answer

- Session: `019d53af-14f1-7a80-987d-cda22ee69505` in `c:\Agent\Crystallize`
- Transcripts: `T10`
- Refs: `T10:L214`, `T10:L228`, `T10:L235`, `T10:L238`
- AI course: the assistant concluded the UAC/helper annoyance came from `[windows] sandbox = "elevated"` and recommended removing the block to stop the prompt path.
- Human intervention: after trying that, the human said they had to put it back because it was blocking later prompts and asked the assistant to research the real behavior instead of `run[ning them] into a blocker`.
- Better outcome forced: investigate the actual Windows agent-mode contract before changing a setting that can hard-block agent use.
- Why real: the assistant's current action made the workflow worse, and the human had to undo it and redirect the investigation.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C02 - The repaired sandbox advice still optimized for the wrong Windows mode

- Session: `019d53af-14f1-7a80-987d-cda22ee69505` in `c:\Agent\Crystallize`
- Transcripts: `T10`
- Refs: `T10:L349`, `T10:L383`, `T10:L458`, `T10:L465`
- AI course: after learning that deleting the block was wrong, the assistant steered toward `unelevated` or WSL as the safer answer.
- Human intervention: the human clarified the real bar was shared `C:\Users\gregs\.codex` logs and state for all Codex usage, no hidden split home, and then explicitly said `No run the elevated one` because that level of Windows self-sufficiency still mattered on this machine.
- Better outcome forced: test the mode that matches the human's actual workflow constraints rather than drifting to the cleaner alternative.
- Why real: the first repair still optimized for the wrong target, so the human had to restate the real contract.
- Confidence: `medium`
- Triage: `intervention event but probably not an accepted incident`

### C03 - Zero-touch self-hosting had to be forced into the durable Crystallize surface

- Session: `019d53fb-c3db-7e61-8adf-a4ff171c0de9` in `c:\Agent\Crystallize`
- Transcripts: `T11`
- Refs: `T11:L275`, `T11:L337`, `T11:L357`, `T11:L432`, `T11:L452`, `T11:L597`
- AI course: Task-0019 answers and docs described the server lane as directionally hands-off, but the durable product/task surface still left room for manual draining, repeated operator action, or an overcomplicated deployment path.
- Human intervention: across several turns the human kept asking whether the task surface really captured `deploy once, then don't touch the box`, added the Windows always-on requirement, and finally said there was `some dropping the ball` because they still had to clarify that the supported lane should be no more complicated than running one script.
- Better outcome forced: durable repo/task truth that supported self-hosting means one supported install action, reboot-safe always-on behavior, and no routine babysitting.
- Why real: the human had to restate the same adequacy bar several times because earlier durable wording still under-modeled the real-world expectation.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C04 - The humane-design fix had to move up to an explicit human model

- Session: `019d53fb-c3db-7e61-8adf-a4ff171c0de9` in `c:\Agent\Crystallize`
- Transcripts: `T11`
- Refs: `T11:L607`, `T11:L641`, `T11:L667`, `T11:L677`
- AI course: the assistant was improving humane-design and interface prompts mainly by adding more lower-level burden-budget and operator-lane heuristics.
- Human intervention: the human said the deeper problem was that `codex doesn't understand people by default` and suggested distilling `what humans are like` so designers can infer click-path, babysitting, and hidden-state rules from first principles.
- Better outcome forced: shared design prompts grounded in a compact human model rather than an accreting checklist of derived heuristics.
- Why real: this is a root-level reframing of the problem being solved, not a stylistic tweak to an already-adequate fix.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C05 - Rustfire debugging was redirected from intuition-first tuning to ruthless causal narrowing

- Session: `019d542c-6ab7-7dc2-833c-5acc5a904748` in `c:\EHG_GregS_main`
- Transcripts: `T12`
- Refs: `T12:L429`, `T12:L436`
- AI course: the active Rustfire investigation had already spent effort on broad causes like bandwidth caps and other global hypotheses.
- Human intervention: the human explicitly ordered `pursue root cause ruthlessly` and said the efficient path is callstacks, breakpoints, or verbose logging that trace the cause upstream from the concrete fault evidence instead of broad wandering.
- Better outcome forced: default to the narrowing step with the highest causal information gain from the first disagreement seam.
- Why real: the human was redirecting the active debugging method, not just asking for more output.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C06 - The updated debug docs were still one level too vague

- Session: `019d53fb-c3db-7e61-8adf-a4ff171c0de9` in `c:\Agent\Crystallize`
- Transcripts: `T11`
- Refs: `T11:L1330`, `T11:L1395`, `T11:L1405`, `T11:L1414`
- AI course: after updating the shared `DEBUG*` docs, the dry-run still answered at the `writer path vs reader path` level instead of tracing the exact fault variables named by the interpolation error.
- Human intervention: the human said `that's better` but still not the right level, asked whether the stronger inference was already valid from code-read, and then required an explicit rule to trace the named values from the message itself while pruning lower-level predicates that were now derivable.
- Better outcome forced: explicit variable-provenance tracing for fault/assert/log messages, with a leaner prompt that stays anchored on named values and their writers/updaters.
- Why real: the human had to tighten the repair because the first prompt rewrite still stopped one level too vague.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C07 - Task-0001 planning was rejected as too narrow for the real automation goal

- Session: `019d5668-49a1-7563-88c0-47b562ff8120` in `c:\Agent\ThirdPerson`
- Transcripts: `T22`
- Refs: `T22:L135`, `T22:L171`, `T22:L174`
- AI course: the initial Task-0001 plan treated the work as a relatively narrow module/observability slice, even after being expanded with more detail and an UnLua section.
- Human intervention: after the first revision, the human said the plan was `Scoped too narrowly`, said the task was to lay a foundation for automation, and made scripting, input, navigation or execution, observability, capture, and queryability non-negotiable.
- Better outcome forced: reframe Task-0001 as the broad automation-foundation task rather than a thin vertical slice.
- Why real: this is a direct rejection of the assistant's current planning model after the first repair had already landed.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C08 - The broadened Task-0001 plan still had to be corrected to Mover-first, minimal, and script-owned

- Session: `019d5668-49a1-7563-88c0-47b562ff8120` in `c:\Agent\ThirdPerson`
- Transcripts: `T22`
- Refs: `T22:L246`, `T22:L297`, `T22:L338`
- AI course: even after the scope correction, the plan still left room for base Character/CMC assumptions, heavy snapshot ownership in C++, and rigid early per-pass regression structure.
- Human intervention: the human required `Use Mover for movement`, said not to use base Characters because they impose CMC, clarified that snapshots should be optional script-composed output where possible, and pushed the plan back toward minimality without losing generalizability.
- Better outcome forced: a Mover-first plan with thinner C++ ownership and more script-side composition.
- Why real: the human tightened multiple architecture boundaries because the broadened plan was still drifting toward heavier or wrong-layer solutions.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C09 - Repeated git account-picker popups forced an auth-path detour

- Session: `019d5668-49a1-7563-88c0-47b562ff8120` in `c:\Agent\ThirdPerson`
- Transcripts: `T22`
- Refs: `T22:L1685`, `T22:L1688`
- AI course: implementation work kept invoking git commands that triggered a Git Credential Manager account-picker popup requiring manual account selection.
- Human intervention: the human paused implementation and said the popup was `obviously not scalable`, forcing investigation of the git auth path before more task work.
- Better outcome forced: a deterministic repo-local push/auth path without repeated human account selection.
- Why real: the human had to interrupt active implementation because the assistant's control path imposed recurring operator burden.
- Confidence: `medium`
- Triage: `intervention event but probably not an accepted incident`

### C10 - PASS-0004 implementation drifted back into pawn/class ownership and CMC assumptions

- Session: `019d5668-49a1-7563-88c0-47b562ff8120` in `c:\Agent\ThirdPerson`
- Transcripts: `T22`
- Refs: `T22:L2708`, `T22:L3060`, `T22:L3302`, `T22:L3315`
- AI course: PASS-0004 implementation had drifted toward `GameAutomation` owning pawn structure and toward Character/CMC-style seams in both `GameAutomation` and `ThirdPerson`.
- Human intervention: the human said the automation module should not enforce actor or pawn structure, that `GameAutomation` does not support CMC in v1, that the project module should not use CMC either, and then renamed the probe pawn to `AThirdPersonPawn`.
- Better outcome forced: game-owned actors with component assumptions only, plus Mover-only execution boundaries at both module layers.
- Why real: this is a repeated intervention after planning-phase architecture corrections had already set these boundaries; the correction still had not stuck in implementation.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C11 - The Android launcher icon was visibly too zoomed in against the approved concept

- Session: `019d55f9-a32a-7241-bb7b-2fa9afd3b849` in `c:\Agent\ModelGeneration`
- Transcripts: `T20`
- Refs: `T20:L645`, `T20:L648`
- AI course: the shipped Android launcher icon rendered with the wrong proportions relative to the icon border and approved concept art.
- Human intervention: the human said it looked `way too zoomed in` on the phone and required a screenshot-backed fix so the phone result matched the concept PNG's proportions.
- Better outcome forced: treat the real launcher render as the proof bar for icon proportion instead of trusting the source asset alone.
- Why real: this directly rejects the human-visible outcome on the device.
- Confidence: `medium`
- Triage: `intervention event but probably not an accepted incident`

### C12 - A claimed icon update still did not look different on the actual phone

- Session: `019d55f9-a32a-7241-bb7b-2fa9afd3b849` in `c:\Agent\ModelGeneration`
- Transcripts: `T20`
- Refs: `T20:L981`, `T20:L1018`, `T20:L1021`, `T20:L1050`
- AI course: after an icon/highlight update, the assistant still believed the new icon work was represented on the device.
- Human intervention: the human said the icon `doesn't look any different`, forcing the assistant to stop guessing, inspect the actual home-screen state, and use the safe cache-refresh path without touching app data.
- Better outcome forced: validate launcher-icon fixes against the live phone surface and launcher caching behavior.
- Why real: the human had to reject the assistant's visible-outcome claim because the device still contradicted it.
- Confidence: `medium`
- Triage: `intervention event but probably not an accepted incident`

### C13 - The on-device product name still surfaced as `MemoryRecorder`

- Session: `019d55f9-a32a-7241-bb7b-2fa9afd3b849` in `c:\Agent\ModelGeneration`
- Transcripts: `T20`
- Refs: `T20:L1301`, `T20:L1304`
- AI course: renaming work had treated the application-level label as sufficient proof that the product now surfaced as `Crystallize`.
- Human intervention: the human reported that search still only showed `MemoryRecorder`, forcing a launcher-activity-level fix instead of a repo-side assumption that the rename was already complete.
- Better outcome forced: the actual launchable surface name on the phone matches the intended product name.
- Why real: the assistant's completion story was still not adequate on the live device surface the human actually uses.
- Confidence: `medium`
- Triage: `intervention event but probably not an accepted incident`

## Confidence buckets

### Strong candidates

- `C01` Deleting the Windows sandbox block created a blocker instead of a safe Windows answer
- `C03` Zero-touch self-hosting had to be forced into the durable Crystallize surface
- `C04` The humane-design fix had to move up to an explicit human model
- `C05` Rustfire debugging was redirected from intuition-first tuning to ruthless causal narrowing
- `C06` The updated debug docs were still one level too vague
- `C07` Task-0001 planning was rejected as too narrow for the real automation goal
- `C08` The broadened Task-0001 plan still had to be corrected to Mover-first, minimal, and script-owned
- `C10` PASS-0004 implementation drifted back into pawn/class ownership and CMC assumptions

### Medium candidates

- `C02` The repaired sandbox advice still optimized for the wrong Windows mode
- `C09` Repeated git account-picker popups forced an auth-path detour
- `C11` The Android launcher icon was visibly too zoomed in against the approved concept
- `C12` A claimed icon update still did not look different on the actual phone
- `C13` The on-device product name still surfaced as `MemoryRecorder`

### Weak or ambiguous promoted candidates

- None promoted into the main list at `weak` confidence.

## Which candidates look like likely accepted incidents

- `C01` Deleting the Windows sandbox block created a blocker instead of a safe Windows answer
- `C03` Zero-touch self-hosting had to be forced into the durable Crystallize surface
- `C04` The humane-design fix had to move up to an explicit human model
- `C05` Rustfire debugging was redirected from intuition-first tuning to ruthless causal narrowing
- `C06` The updated debug docs were still one level too vague
- `C07` Task-0001 planning was rejected as too narrow for the real automation goal
- `C08` The broadened Task-0001 plan still had to be corrected to Mover-first, minimal, and script-owned
- `C10` PASS-0004 implementation drifted back into pawn/class ownership and CMC assumptions

## Which candidates are real interventions but probably belong outside the accepted incident set

- `C02` The repaired sandbox advice still optimized for the wrong Windows mode
- `C09` Repeated git account-picker popups forced an auth-path detour
- `C11` The Android launcher icon was visibly too zoomed in against the approved concept
- `C12` A claimed icon update still did not look different on the actual phone
- `C13` The on-device product name still surfaced as `MemoryRecorder`

## Any ambiguous boundaries that need a second read

- `C02` is real, but it may fold into `C01` in a stricter accepted-incident pass because both belong to the same sandbox-config correction arc.
- `C03` and `C04` are distinct in this pass, but they sit very close: `C03` is the repo-local product/task contract miss, while `C04` is the shared prompt/root-model fix that the human used to explain why the miss kept recurring.
- `C09` is a genuine operator-burden intervention, but it may be too task-local for the narrowed accepted corpus.
- `C11` through `C13` are all real visible-surface corrections on the phone, but a stricter accepted pass may keep them outside the final incident set as ordinary iterative design/device review rather than deeper recurring process failures.
