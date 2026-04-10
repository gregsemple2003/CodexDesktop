# INTERVENTION-PASS1

Source day: `2026-04-07`

## Source scope reviewed

- Primary source scope only: `C:\Users\gregs\.codex\sessions\2026\04\07\rollout-*.jsonl`
- Broad scan completed across all `17` day-scoped rollout files at raw `event_msg` / `user_message` level.
- Raw-count note: the source-day folder currently contains `233` raw `user_message` records, while the official day CSV has `231` rows. This pass treated raw JSONL as source of truth and used the CSV only as a navigation/count sanity aid.
- Narrow rereads used for exact extraction:
- `T01` = [rollout-2026-04-07T10-35-51-019d685f-06df-71c3-be76-e7f1b05e6221.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/07/rollout-2026-04-07T10-35-51-019d685f-06df-71c3-be76-e7f1b05e6221.jsonl)
- `T02` = [rollout-2026-04-07T14-31-32-019d6936-ce54-7bd1-aa40-76b7d72614a5.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/07/rollout-2026-04-07T14-31-32-019d6936-ce54-7bd1-aa40-76b7d72614a5.jsonl)
- `T03` = [rollout-2026-04-07T14-44-11-019d6942-63b6-7c72-9abc-9ab209816f97.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/07/rollout-2026-04-07T14-44-11-019d6942-63b6-7c72-9abc-9ab209816f97.jsonl)
- `T04` = [rollout-2026-04-07T15-54-57-019d6983-2d79-7973-b95b-4c9088afdc3c.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/07/rollout-2026-04-07T15-54-57-019d6983-2d79-7973-b95b-4c9088afdc3c.jsonl)
- No real intervention boundary was found in the eight single-prompt automation runs under `C:\Users\gregs\.codex`; candidate activity concentrated in `Locate UE network thread`, `Implement tasks 22-27`, and `Open task 6 empty settings`, with one ambiguous scope-tightening exchange in `Check MassMover compatibility`.
- Chronology note: like the official day folder, this pass stays ordered by transcript position inside the `2026-04-07` source-day files even when later records spill into `2026-04-08` and `2026-04-09`.

## Total candidate intervention events found

- `20`

## Chronological candidate list

### C01 - Redirect from legacy profiler fallback to the newest Insights stack

- Session: `Locate UE network thread`
- Refs: `T01:L238`, `T01:L249`, `T01:L262`, `T01:L269`, `T01:L272`
- AI course: the assistant was deriving bandwidth from logs and then recommended the older `NETPROFILE` / `.nprof` path as the shortest path.
- Human intervention: the human rejected that seam with `No there's an insights tool...` and `The insights tooling. I specifically mean that because I want to use the newest stack.`
- Better outcome forced: give the actual Unreal Insights / NetTrace workflow instead of steering back to the older profiler path.
- Why real: this was not ordinary preference-setting; it rejected the active debugging path and forced the tool choice back onto the requested stack.
- Confidence: `medium`
- Triage: `intervention event but probably not an accepted incident`

### C02 - Inline, verifiable Epic quotes only

- Session: `Locate UE network thread`
- Refs: `T01:L1146`, `T01:L1153`, `T01:L1170`, `T01:L1177`, `T01:L1200`
- AI course: the assistant cited Epic docs without inlined grounding and then presented a release-note quote that turned out not to be safely live-verifiable from the current docs route.
- Human intervention: the human asked for the whole relevant quoted section inline and then challenged the unsupported quote with `did you hallucinate it?`
- Better outcome forced: only quote what can be directly verified now, and call out redirect/source uncertainty instead of laundering a search-snippet-style quote into a verified claim.
- Why real: the human explicitly rejected the evidentiary standard and source-truth bar of the answer.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C03 - Reframe the bandwidth note for executive leadership instead of tech leads

- Session: `Locate UE network thread`
- Refs: `T01:L1207`, `T01:L1210`, `T01:L1217`, `T01:L1219`
- AI course: the assistant answered the leadership request with a jargon-heavy technical writeup centered on internal engine names and sub-costs.
- Human intervention: the human said `they don't care about syncstatecollection or replicationproxy_simulated` and clarified that `leadership` meant executive leadership.
- Better outcome forced: return to the original task format and express the bandwidth point in executive-friendly terms.
- Why real: the human materially tightened the audience/meaning contract rather than making a cosmetic wording request.
- Confidence: `medium`
- Triage: `intervention event but probably not an accepted incident`

### C04 - Put the idle TASK-LEADER into passive watch mode

- Session: `Implement tasks 22-27`
- Refs: `T02:L531`, `T02:L534`, `T02:L540`, `T02:L560`
- AI course: the delegated TASK-LEADER was still sitting as a latent worker while local implementation continued, creating needless conflict/race risk.
- Human intervention: the human said it would be better to `nudge it to keep watch instead of just sitting there and causing more conflicts / races`.
- Better outcome forced: keep ownership but stop repo-changing activity unless explicitly tasked.
- Why real: the human rejected a live orchestration behavior that was judged inadequate and risky.
- Confidence: `medium`
- Triage: `intervention event but probably not an accepted incident`

### C05 - Do not stop at batch closure; roll directly into QA and task mining

- Session: `Implement tasks 22-27`
- Refs: `T02:L1282`, `T02:L1289`, `T02:L1292`, `T02:L1304`
- AI course: after closing Tasks 22 through 27, the assistant treated the explicit batch boundary as a stopping point and idled.
- Human intervention: `Why are you just sitting there?`
- Better outcome forced: move into QA-first and then mine follow-up tasks from the product vision rather than waiting.
- Why real: this is a direct correction of the assistant's stopping point and ownership model.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C06 - Use the app like a human, not only via QA proxies

- Session: `Implement tasks 22-27`
- Refs: `T02:L1386`, `T02:L1393`, `T02:L1402`, `T02:L1405`
- AI course: the assistant had corrected into broader QA and task-mining work, but was still operating through harnesses, artifacts, and task proposals rather than the visible app surface itself.
- Human intervention: the human ordered a click-through `like a human would`, grounded only in what was visible on screen.
- Better outcome forced: a human-facing UX review rather than another proxy proof pass.
- Why real: this rejects automated/test/task-mining proxies as sufficient for the next step.
- Confidence: `medium`
- Triage: `unclear and needs transcript reread`

### C07 - Never do destructive things on the human's phone without asking

- Session: `Implement tasks 22-27`
- Refs: `T02:L1810`, `T02:L1851`, `T02:L1858`, `T02:L1861`
- AI course: connected Android tests were run on the physical phone and wiped the debug app's saved server URL and related phone-local state.
- Human intervention: `Never do destructive things on my phone without first asking.`
- Better outcome forced: hard trust-boundary rule that the physical phone is off-limits unless the exact operation is explicitly authorized.
- Why real: the human is rejecting a completed action as unsafe and unacceptable, not merely expressing frustration.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C08 - Fix the phone now, not just the future policy

- Session: `Implement tasks 22-27`
- Refs: `T02:L1868`, `T02:L1871`, `T02:L1878`, `T02:L1891`, `T02:L1898`, `T02:L1901`
- AI course: after admitting the phone mistake, the assistant moved into future policy and guardrail design rather than immediate remediation of the current phone state.
- Human intervention: the human re-centered on immediate recovery and then escalated to `You fix it.`
- Better outcome forced: concrete recovery work for the current device before abstract process hardening.
- Why real: this is a repeated correction because the first answer did not satisfy the real-world recovery demand.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C09 - Give the exact WAV count and back up the corpus before anything else

- Session: `Implement tasks 22-27`
- Refs: `T02:L2046`, `T02:L2059`, `T02:L2076`, `T02:L2080`, `T02:L2093`, `T02:L2100`, `T02:L2112`
- AI course: the assistant answered the `2 clips safe on server` scare with DB/artifact counts and an interpretation of what the number meant, not the exact raw-WAV inventory or a concrete safety step.
- Human intervention: the human demanded the exact WAV count, then required server shutdown plus a backup of all `7785` raw WAVs and a manifest before proceeding.
- Better outcome forced: direct inventory plus protective backup of the raw corpus.
- Why real: the human had to force the assistant off an interpretive answer and onto explicit recovery-risk mitigation.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C10 - Restore the pre-incident phone-visible vault state, not just explain product limits

- Session: `Implement tasks 22-27`
- Refs: `T02:L2202`, `T02:L2215`, `T02:L2222`, `T02:L2235`, `T02:L2251`
- AI course: the assistant answered `hook me up` by explaining current app limitations and the difference between raw corpus access and processed review surfaces.
- Human intervention: the human clarified that the real target was the prior phone state where the app believed `7000+` clips were already safe on server.
- Better outcome forced: diagnose and recover the lost phone-local clip ledger or otherwise restore the prior visible state, not retreat into a feature-gap explanation.
- Why real: the human rejected scope narrowing and forced the assistant back onto the actual harm they were trying to undo.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C11 - The incident plan/contract cannot stop at surface summaries

- Session: `Open task 6 empty settings`
- Refs: `T03:L722`, `T03:L729`, `T03:L788`, `T03:L846`, `T03:L853`, `T03:L868`
- AI course: the assistant produced a plan and then a renamed `goal_stack` contract, but both still read as surface-level and non-falsifiable.
- Human intervention: the human rejected the plan and then escalated to `Nail down the incident schema and give me a couple of examples.`
- Better outcome forced: a falsifiable incident contract with concrete schema/examples rather than thematic planning prose.
- Why real: this is a real intervention boundary, and it happened twice because the first repair did not stick.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C12 - Tighten the definition from generic bugs to human course-correction incidents

- Session: `Open task 6 empty settings`
- Refs: `T03:L917`, `T03:L924`, `T03:L937`, `T03:L949`, `T03:L968`
- AI course: the first example was really a product bug/task-splitting example, but the assistant had let it masquerade as an incident.
- Human intervention: the human challenged the classification and said these are `human course correction incident`, not just `any incident`.
- Better outcome forced: only preserve cases with AI course, human intervention, and corrected target state.
- Why real: the human materially tightened the core semantic boundary of the task.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C13 - Keep `expected_state` and `actual_state` concrete at the event level

- Session: `Open task 6 empty settings`
- Refs: `T03:L1113`, `T03:L1142`, `T03:L1145`, `T03:L1178`
- AI course: incident examples were too abstract to tell what the concrete triggering event actually was from the record alone.
- Human intervention: the human required `expected_state` and `actual_state` to refer to the concrete incident event, with abstraction only in the later chain.
- Better outcome forced: event-level grounding in the record itself.
- Why real: this is a direct contract correction for reviewability and grounding, not a stylistic preference.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C14 - Replace the overcomplicated goal stack with an explicit-human-only `why_chain`

- Session: `Open task 6 empty settings`
- Refs: `T03:L1185`, `T03:L1358`, `T03:L1369`, `T03:L1420`, `T03:L1438`, `T03:L1441`, `T03:L1451`
- AI course: the assistant's `goal_stack` mixed concrete event, workflow tactic, and inferred higher-level principles in a way that was not a clean progressive generalization.
- Human intervention: the human simplified the model to a `why_chain`, required each step to progressively generalize from the concrete, and capped the chain at the highest reason explicitly stated by the human in evidence.
- Better outcome forced: one strict recursive chain from event to explicit human reason.
- Why real: the human rejected the current abstraction model and forced a narrower, more evidence-bound structure.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C15 - The `why_chain` rule still was not sticking; split sibling reasons into `why_chains`

- Session: `Open task 6 empty settings`
- Refs: `T03:L1518`, `T03:L1521`, `T03:L1528`, `T03:L1538`, `T03:L1562`, `T03:L1667`, `T03:L1678`, `T03:L1743`, `T03:L1746`, `T03:L1753`, `T03:L1763`, `T03:L1776`, `T03:L1783`, `T03:L1786`, `T03:L1793`
- AI course: even after the earlier correction, the assistant kept over-splitting reasons, mixing rule extraction into the chain, and violating the documented `entry N+1 answers why entry N mattered to the human` rule.
- Human intervention: the human forced the rules into the schema/docs, challenged repeated noncompliance directly, and steered the structure toward multiple sibling `why_chains` instead of one branchy chain.
- Better outcome forced: strict linear rationale paths, with multiple chains when there are multiple independent upstream reasons.
- Why real: this is a repeated corrective event because the prior `why_chain` correction did not hold.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C16 - Inherited context invalidated the official isolated run

- Session: `Open task 6 empty settings`
- Refs: `T03:L3237`, `T03:L3275`, `T03:L3284`, `T03:L3287`
- AI course: the assistant used a subagent launch shape that inherited the parent thread context into what was supposed to be an isolated incident-harvester test.
- Human intervention: the human declared the April 6 work invalid, ordered cleanup and relocation into the reports home, and required an actually isolated rerun.
- Better outcome forced: context-free validation against durable instructions only.
- Why real: the human rejected the run's epistemic validity, not merely its formatting.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C17 - `DAILY-BRIEF` must truly be self-contained

- Session: `Open task 6 empty settings`
- Refs: `T03:L3687`, `T03:L3699`, `T03:L3706`, `T03:L3761`, `T03:L3768`, `T03:L3771`
- AI course: the first `DAILY-BRIEF` implementation still behaved as if a reviewer could click out to adjacent incident JSON and CSV files.
- Human intervention: the human said the brief had failed to understand what `self-contained` means and required the incident itself to be inlined.
- Better outcome forced: a review packet that stands alone for a second reviewer or upstream model.
- Why real: the human is rejecting the active artifact model as inadequate for the stated use.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C18 - The daily brief still cherry-picked and compressed too much

- Session: `Open task 6 empty settings`
- Refs: `T03:L3824`, `T03:L3827`, `T03:L3840`, `T03:L3847`, `T03:L3850`, `T03:L3857`, `T03:L3860`
- AI course: even after the self-contained fix, the brief still behaved like a curated summary optimized for compression and supporting the harvester's conclusion.
- Human intervention: the human called out cherry-picking and `navel-gazing`, then explicitly allowed a much larger evidence budget so the brief could support first-principles review.
- Better outcome forced: audit-packet behavior with more complete raw, relevant context and clearer separation between source material and interpretation.
- Why real: this is another repeated intervention because the first self-contained fix still missed the real review goal.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C19 - Even the raw windows were still too lossy; incident JSONs must carry heavyweight verbatim evidence

- Session: `Open task 6 empty settings`
- Refs: `T03:L3944`, `T03:L3947`, `T03:L3954`, `T03:L4011`, `T03:L4056`, `T03:L4059`, `T03:L4066`, `T03:L4112`, `T03:L4119`, `T03:L4122`
- AI course: the assistant kept moving toward `brief` / `evidence packet` language, but still selected and compressed transcript evidence in a way that did not reliably recover what the human had actually said.
- Human intervention: the human required verbatim-only evidence, memory-jogging source sections, and heavyweight transcript evidence inside the incident JSON itself rather than only in the brief.
- Better outcome forced: heavyweight incident records that carry their own epistemic warrant.
- Why real: the human explicitly rejected the assistant's ongoing compression/selection model as incompatible with first-principles review.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

### C20 - Backfill was declared complete while omitting the obvious phone-data incident

- Session: `Open task 6 empty settings`
- Refs: `T03:L4795`, `T03:L4802`, `T03:L4805`, `T03:L4815`, `T03:L4825`, `T03:L4832`, `T03:L4835`
- AI course: the assistant declared backfill complete and left April 7 effectively as a zero-incident day.
- Human intervention: the human pointed to the missing phone-data-loss incident as obvious and then stopped the pass, asking for a focused JSONL-mining handoff instead of continued promotion work.
- Better outcome forced: do not claim completeness while a high-salience intervention event is still missing from the corpus.
- Why real: the human is directly rejecting a closure/completeness claim.
- Confidence: `strong`
- Triage: `likely accepted incident candidate`

## Which candidates look like likely accepted incidents

- `C02` source-truth / quote-verification failure
- `C05` premature idle stop after batch closure
- `C07` destructive phone operations without permission
- `C08` policy-only answer when immediate recovery was required
- `C09` exact WAV-count / backup demand after perceived server-loss risk
- `C10` restore prior phone-visible vault state instead of narrowing to product limitations
- `C11` non-falsifiable incident plan/contract
- `C12` generic-bug example misclassified as incident
- `C13` abstract state fields that obscured the concrete event
- `C14` overcomplicated/inferential goal stack replaced by explicit-human-only `why_chain`
- `C15` repeated `why_chain` noncompliance and need for sibling `why_chains`
- `C16` inherited context invalidating an official isolated run
- `C17` daily brief not actually self-contained
- `C18` daily brief cherry-picking/compression undermining first-principles review
- `C19` incident/brief evidence still too lossy and lightweight
- `C20` false completion of backfill while obvious incident remained omitted

## Which candidates are real interventions but probably belong outside the accepted incident set

- `C01` tool-path correction toward the newest Insights stack
- `C03` executive-audience framing correction for the leadership brief
- `C04` passive-watch redirection for a delegated TASK-LEADER

## Ambiguous boundaries that need a second read

- `C06` feels real because it rejects QA/task-mining proxies in favor of a human-surface review, but it may also read as a fresh tasking turn rather than a strong enough incident on second pass.
- `T04:L99`, `T04:L385`, `T04:L437` in `Check MassMover compatibility` contain one plausible scope-tightening correction around keeping only behavior-changing changes and getting off the `NetDriver.cpp` workaround, but the current read treats it as ordinary implementation iteration rather than a clear human intervention event.
- `C04` could also be re-read as local coordination rather than incident-grade course correction if the accepted set wants to stay tighter on durable human-cost signal.
