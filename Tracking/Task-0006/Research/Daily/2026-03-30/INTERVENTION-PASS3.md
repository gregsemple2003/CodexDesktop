# INTERVENTION-PASS3

Source day: `2026-03-30`

## Source scope analyzed

- PASS2 artifact in scope: `C:\Agent\CodexDashboard\Tracking\Task-0006\Research\Daily\2026-03-30\INTERVENTION-PASS2.md`
- Candidate event set in scope: `C01` through `C12`
- PASS3 worked from the PASS2 event analyses, repeated cluster hints, likely accepted-incident calls, and strongest human-model signals.
- Additional artifact rereads during PASS3: none
- Transcript windows reopened during PASS3: none
- Why no reread was needed: PASS2 already preserved the decisive human standards clearly enough to keep the principle boundaries evidence-bound without re-opening raw transcript windows.

## PASS2 artifact used

- Primary input artifact: `C:\Agent\CodexDashboard\Tracking\Task-0006\Research\Daily\2026-03-30\INTERVENTION-PASS2.md`
- PASS2 sections relied on most heavily:
  - `Likely accepted incidents`
  - `Likely non-incident but still important intervention events`
  - `Repeated cluster hints noticed across the analyzed set`
  - `Strongest human-model signals worth carrying into later clustering or principle work`
- PASS3 did not import any parent-thread summary or non-local context outside what PASS2 already grounded.

## Candidate clusters considered

### CL01 - Human-time burden transfer after delegation or auto-approval

- Shared decision failure: the assistant treated delegation as launch-only or claimed supervision without actually carrying the timekeeping burden.
- Supporting events: `C01`, `C04`, `C08`, `C10`
- Why this is a real cluster: the same human standard repeats across unwanted chatter, spawn-and-forget delegation, broken completion promises, and misleading `watching` language.
- Disposition: kept as `P01`

### CL02 - Structural owner, phase, and scope mismatches

- Shared decision failure: the assistant reacted to structural workflow problems as if they were local artifact problems, or kept work in the wrong owner/phase/scope after the human clarified where the durable rule belonged.
- Supporting events: `C03`, `C05`, `C11`, `C12`
- Why this is a real cluster: the corrections all turn on where the durable control point lives, not on the surface topic of a specific brief or document.
- Disposition: kept as `P02`

### CL03 - Proxy evidence and over-narrowed proof

- Shared decision failure: the assistant accepted evidence or a diagnostic seam that was too narrow to honestly support the claim being made.
- Supporting events: `C02`, `C07`, `C09`
- Why this is a real cluster: each event is a miss on evidentiary adequacy, where the human had to push back toward the wider or more direct proving surface.
- Disposition: kept as `P03`

### CL04 - Requested artifact form is part of completion

- Shared decision failure: the assistant substituted an adjacent output for the form and placement the human actually intended to use.
- Supporting events: `C06`, `C12`
- Secondary related support: `C03`
- Why this is a real cluster: the intervention is not just that an artifact existed, but that the human needed it in a repo-usable form at the correct scope.
- Disposition: kept as `P04`

### CL05 - Standalone local-first research doctrine

- Shared decision failure: the standing workflow still centered external-model briefing after the human had already rejected that as the default.
- Supporting events: `C11`
- Why this was considered separately: it is a meaningful workflow redesign signal, but the direct evidence is still mostly one intervention arc.
- Disposition: not kept standalone; merged into `P02` as a narrower owner/phase/scope sub-case

## Final kept principles

### P01

- `principle_id`: `P01`
- `principle_statement`: `When the human has already delegated routine workflow or asked you to watch delegated work, absorb the supervision burden in human time: do not surface routine intermediate state, and do not say you are watching unless an active polling loop will carry the work to terminal-state reporting without another prompt.`
- `decision_point`: before posting status after delegation, and immediately after launching or inheriting responsibility for delegated work
- `failure_signature`: `routine-gate chatter`, `spawn-and-forget delegation`, or `"watching"` without a live completion loop
- `why_this_is_durable`: the same adequacy rule applies across orchestration, monitoring, and notification semantics; the human repeatedly treated the real issue as burden transfer, not as four unrelated mistakes
- `supporting_events`: `C01`, `C04`, `C08`, `C10`
- `supporting_human_model_signals`: `C01` says auto-approved gates should not bubble routine planning state back up; `C04` says directed subagents must be waited on and reported immediately; `C08` says a promise to notify on completion must survive until the work finishes; `C10` says `watching` only counts if the assistant is actively polling in the human's timeline
- `counterfactual_prevention_claim`: if this rule had been applied early, the human likely would not have needed to restate the auto-approve rule, ask whether the agent was watching, or ping for completions that had already happened
- `scope_and_non_goals`: this does not forbid all status updates; real blockers, material scope questions, and honestly final completion states should still surface
- `pre_action_question`: `Has the human already delegated this gate, and if so am I actually blocked, actually finished, or actively polling to completion right now?`
- `operational_check`: `Audit question: if I say "watching" or stay silent on a delegated task, can completion still be reported without the human prompting me again?`
- `confidence`: `strong`

### P02

- `principle_id`: `P02`
- `principle_statement`: `When new information reveals a structural workflow miss, move the work to the correct owner, phase, and scope. Fix durable problems where the rule lives instead of patching one artifact, repeating the wrong phase, or storing product-level intent in a task-local location.`
- `decision_point`: after critique, after a human clarifies workflow state, or when placing a durable artifact that will guide later work
- `failure_signature`: `treating a structural problem as a one-off rewrite`, `looping back into an earlier phase`, or `putting durable guidance at the wrong scope`
- `why_this_is_durable`: leader ownership, phase transitions, research structure, and design-anchor placement are all recurring orchestration decisions; the human corrections all point to the same meta-rule about durable ownership
- `supporting_events`: `C03`, `C05`, `C11`, `C12`
- `supporting_human_model_signals`: `C03` says weak brief quality should be fixed at the leader layer, not with ad hoc rewriting; `C05` says planning-ready input should advance into `TASK-LEADER` at the plan phase; `C11` says research should default to local decomposition and synthesis instead of briefing an external model by default; `C12` says the durable product definition belongs at repo-root `Design/GENERAL-DESIGN.md`
- `counterfactual_prevention_claim`: if this rule had been applied early, the assistant likely would have escalated brief repair to the right workflow owner, advanced planning-ready work into the planning phase, and avoided misplacing durable design intent under one task
- `scope_and_non_goals`: this is not a claim that every artifact defect is structural; small local defects can still be fixed locally when the owning process, phase, and scope are already correct
- `pre_action_question`: `Is this problem really about the file in front of me, or about the owner, phase, or scope that should govern it?`
- `operational_check`: `Plan-review check: name the workflow owner, current phase, and durable scope for the next action; if any of those are unclear, do not default to another local rewrite`
- `confidence`: `medium`

### P03

- `principle_id`: `P03`
- `principle_statement`: `Do not narrow a theory or close work from proxy evidence. Before steering the next debug step or declaring success, use the widest practical discriminating check first, and require evidence from the real human-facing lane for any completion claim.`
- `decision_point`: before committing to a debugging seam, before claiming proof, and before marking work complete
- `failure_signature`: `overfitting to one subsystem early`, `calling work done from server-only proof`, or `accepting artifacts that cannot directly demonstrate the claimed behavior`
- `why_this_is_durable`: the human corrections all reject convenient but inadequate evidence; the same discipline applies whether the miss is a rollback diagnosis, a regression lane, or a bogus transcript artifact
- `supporting_events`: `C02`, `C07`, `C09`
- `supporting_human_model_signals`: `C02` says to throw the net wider and start with the most general rollback discriminator; `C07` says reading workflow docs and passing a server-only lane does not close a task that still requires real delegated regression; `C09` says a sine-wave clip with no voice is not proof of transcript functionality and cannot close regression
- `counterfactual_prevention_claim`: if this rule had been applied early, the assistant likely would have started debugging from a broader discriminator, avoided the false `Task-0006` closure, and rejected the sine-wave transcript artifact before presenting it as proof
- `scope_and_non_goals`: this does not require the single most expensive test at every step; it requires the widest useful discriminator that can honestly validate or falsify the claim being made
- `pre_action_question`: `What exact claim am I about to make, and does my next check directly discriminate that claim on the real operating surface rather than through a proxy?`
- `operational_check`: `Closure gate: map each success claim to its required proof lane and reject evidence that cannot directly exercise the claimed user-visible behavior`
- `confidence`: `strong`

### P04

- `principle_id`: `P04`
- `principle_statement`: `Treat the requested output form and location as part of the requirement. If the human asked for repo-readable or product-scoped artifacts, produce that exact form or surface the blocker before saying the workspace is complete.`
- `decision_point`: before declaring an artifact set, workspace, or documentation change complete
- `failure_signature`: `substituting an adjacent artifact for the requested usable form`
- `why_this_is_durable`: many human requests are about downstream usability, not just existence; output form and placement determine whether the artifact can actually support the next step in the repo
- `supporting_events`: `C06`, `C12`
- `supporting_human_model_signals`: `C06` explicitly asked for papers to be downloaded and transcribed under `Research/Papers`, and the later correction reinforced that Codex-native tooling was needed to deliver that form; `C12` explicitly says the enduring product definition belongs at repo-root rather than inside a task folder
- `counterfactual_prevention_claim`: if this rule had been applied early, the assistant likely would have either produced the requested Markdown paper corpus and correct design-anchor placement, or surfaced the blocker before claiming the relevant workspace was complete
- `scope_and_non_goals`: this does not mean every formatting preference is immutable; it applies when the requested form or location is part of how the human intends to use the artifact or when the workflow already defines the authoritative home
- `pre_action_question`: `Am I delivering the artifact in the form and location the human will actually use next, or just something technically adjacent?`
- `operational_check`: `Output review check: compare the requested noun plus format/location against the actual repo artifacts before using completion language`
- `confidence`: `medium`

## Rejected or merged principle candidates and why

- `candidate_statement`: `Never interrupt the parent thread after auto-approval.`
  - `status`: `merged`
  - `reason`: narrower than the stronger burden-transfer rule; the real issue is not only interruption, but failure to carry delegated supervision through completion
  - `merged_into`: `P01`

- `candidate_statement`: `Always monitor subagents until they finish and report immediately.`
  - `status`: `merged`
  - `reason`: this is a direct sub-case of human-time burden transfer and would duplicate `P01`
  - `merged_into`: `P01`

- `candidate_statement`: `Do not say "watching" unless polling every minute.`
  - `status`: `merged`
  - `reason`: the durable rule is about active human-time monitoring, not one exact cadence; the one-minute loop is an operationalization, not the principle itself
  - `merged_into`: `P01`

- `candidate_statement`: `Fix weak research briefs only by making leader-layer audit and bounded refinement the default.`
  - `status`: `merged`
  - `reason`: this is one strong example of the broader owner/phase/scope rule captured in `P02`
  - `merged_into`: `P02`

- `candidate_statement`: `Always advance planning-ready work into TASK-LEADER instead of another research pass.`
  - `status`: `merged`
  - `reason`: accurate but too phase-specific to stand alone once `P02` already captures moving work to the correct owner and phase
  - `merged_into`: `P02`

- `candidate_statement`: `Default research to local-first decomposition and avoid external-model briefing.`
  - `status`: `merged`
  - `reason`: important but currently supported mainly by `C11`; it is better carried as a narrower workflow-shape instance inside `P02` until more days show it recurring independently
  - `merged_into`: `P02`

- `candidate_statement`: `Always start debugging at the widest rollback breakpoint before investigating any subsystem theory.`
  - `status`: `merged`
  - `reason`: too debugging-specific as written; the broader durable rule is to avoid proxy or over-narrow evidence and use the widest practical discriminator first
  - `merged_into`: `P03`

- `candidate_statement`: `Never close a transcript task with a sine-wave audio file.`
  - `status`: `merged`
  - `reason`: this is a concrete manifestation of the broader anti-proxy proof rule in `P03`
  - `merged_into`: `P03`

- `candidate_statement`: `Never stop at PDFs when asked for paper transcriptions.`
  - `status`: `merged`
  - `reason`: too artifact-specific; the stronger durable rule is that requested output form is part of completion
  - `merged_into`: `P04`

- `candidate_statement`: `Keep the durable product design anchor at repo root.`
  - `status`: `merged`
  - `reason`: correct but narrower than the combined scope and artifact-form rules already captured in `P02` and `P04`
  - `merged_into`: `P02`

## Smallest recommended principle set for this scope

- `P01` - absorb delegated supervision burden in human time
- `P02` - move structural misses to the correct owner, phase, and scope
- `P03` - reject proxy evidence and use the real proof lane before narrowing or closing
- `P04` - treat requested artifact form and location as part of completion

Four principles is the smallest honest set for this source day. Collapsing `P04` into `P03` would make `completion truth` too broad and would hide the separate recurring rule that output form and artifact placement are themselves part of adequacy.

## Principles that are still too weak and need more days or more events

- `Prefer local-first research over external-model briefing as the default research doctrine.`
  - Current support is mainly `C11`; it is useful now as a sub-rule inside `P02`, but it needs more recurrence before it should stand alone as a durable cross-day principle.

- `Keep enduring product-definition artifacts at repo-root product scope.`
  - `C12` supports this clearly, but today it still looks more like a specific document-architecture rule than a top-tier standalone principle.

- `Start debugging from the widest practical discriminator before investing in a subsystem theory.`
  - `C02` is a strong local example, but more examples would help show whether this should stay inside `P03` or become its own principle later.

## Transcript windows reopened during PASS3 and why

- None.
- PASS3 stayed within PASS2 because the cluster boundaries and human standards were already explicit enough to support conservative principle extraction.
