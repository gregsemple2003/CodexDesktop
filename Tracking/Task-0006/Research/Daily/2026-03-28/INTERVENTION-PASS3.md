# INTERVENTION-PASS3 Principle Report

Source day: `2026-03-28`

PASS2 artifact used: [INTERVENTION-PASS2.md](/c:/Agent/CodexDashboard/Tracking/Task-0006/Research/Daily/2026-03-28/INTERVENTION-PASS2.md)

## Source Scope Analyzed

- Source scope: `2026-03-28`
- PASS2 event set in scope: `C01` through `C17`
- PASS3 worked from [INTERVENTION-PASS2.md](/c:/Agent/CodexDashboard/Tracking/Task-0006/Research/Daily/2026-03-28/INTERVENTION-PASS2.md) only.
- I did not reopen [INTERVENTION-PASS1.md](/c:/Agent/CodexDashboard/Tracking/Task-0006/Research/Daily/2026-03-28/INTERVENTION-PASS1.md) or raw transcripts during PASS3 because PASS2 already preserved the boundary calls, repeated cluster hints, and strongest human-model signals needed to keep this pass honest.

## Candidate Clusters Considered

- `acceptance-lane grounding` (`C01`, `C16`, `C17`)
  Kept as `P01`.
- `prompt and launcher contract honesty` (`C02`, `C05`, `C06`, `C07`, `C08`, `C13`)
  Kept as `P02`.
- `proof-before-proxy` (`C09`, `C10`)
  Kept as `P03`.
- `artifact execution and evidence-packet adequacy` (`C11`, `C12`, `C14`, `C15`)
  Kept as `P04`.
- `single-source-of-truth / pointer minimization` (`C03`, `C04`)
  Not kept as a standalone principle. The support is real, but it is still one narrow cleanup arc around `CURRENT-TASK.json` rather than a multi-context recurring decision failure.

## Final Kept Principles

### P01

- `principle_id`: `P01`
- `principle_statement`: `Reground on the active task and real acceptance lane before choosing work; stale thread memory, local root-cause stories, and workaround convenience do not override the human's current goal.`
- `decision_point`: `At task start, after a relaunch/reset, and whenever debugging starts to redefine the goal or switch lanes.`
- `failure_signature`: `working a stale task or a convenient side lane instead of the currently named acceptance path`
- `why_this_is_durable`: `The same miss appeared as wrong-task drift, human-world goal drift, and premature workaround preference. The technical surface changed, but the missed decision was the same: a local frame outranked the currently stated lane.`
- `supporting_events`:
  - `C01`: wrong-task drift onto `Task-0003`
  - `C16`: emulator-internals explanation displaced the real app-plus-server regression goal
  - `C17`: workaround preference arrived before the direct no-code `injectAudio` lane was exhausted
- `supporting_human_model_signals`:
  - `C16`: "the problem we're trying to solve is automated end-to-end regression testing using the actual frontend android app and backend web service"
  - `C17`: "See if you can get the injectAudio path to work without code changes first."
  - `C01`: the human had to explicitly relaunch the agent on `task-0002`
- `counterfactual_prevention_claim`: `If applied earlier, C01 would likely have regrounded before reading the wrong task tree, C16 would have kept root-cause work subordinate to the real regression lane, and C17 would have tested the more faithful direct lane before recommending a workaround.`
- `scope_and_non_goals`: `This does not ban local root-cause work or workarounds. It requires that they remain subordinate to the named task and acceptance lane, and that a more faithful direct lane not be abandoned early without proof it has been exhausted.`
- `pre_action_question`: `What exact task and acceptance lane is active right now, and does this next step stay on that lane?`
- `operational_check`: `Audit question: before reading docs, reframing the problem, or proposing a workaround, did the agent restate the active task id and current human-world success criterion?`
- `confidence`: `medium`

### P02

- `principle_id`: `P02`
- `principle_statement`: `When shipping or editing a reusable prompt or worker handoff, make the operating contract explicit and honest: actor, generic-vs-instance scope, discoverable vs injected inputs, durable-capture duties, currently supported artifact shape, and what the worker does after consuming imported analysis.`
- `decision_point`: `Whenever a prompt template, launch contract, or handoff worker is being drafted or revised.`
- `failure_signature`: `template text implies hidden context, hides transport boundaries, or leaves worker purpose and supported workflow shape implicit`
- `why_this_is_durable`: `Six events across two prompt families repeated the same failure from different angles. Operators had to reassert who the prompt was for, what was local versus injected, what had to be saved, what naming shape was real, and whether the worker stopped at verdict or continued debugging.`
- `supporting_events`:
  - `C02`: generic gatherer template collapsed into one task instance and the wrong actor boundary
  - `C05`: imported external analysis was not durably captured
  - `C06`: inline vs local input truth stayed ambiguous after the first fix
  - `C07`: the only supported bug-scoped naming and injected slot remained too implicit
  - `C08`: worker purpose stopped at hypothesis evaluation instead of continuing debugging
  - `C13`: the prompt catalog hid discoverable vs injected context and the real launch contract
- `supporting_human_model_signals`:
  - `C02`: the prompt is "supposed to be generic across repos" and "you aren't prompting the superbrain"
  - `C06`: "Is this true? I thought we were doing inline."
  - `C08`: "take the superbrain's response and attempt to prove/disprove the hypothesis then carry on with your debugging efforts"
  - `C13`: "I thought it was all just the prompt" and "Context To Inject"
- `counterfactual_prevention_claim`: `If this rule had been applied at first draft, the follow-up template chain would likely have avoided most of C05 through C08, and the prompt-catalog correction in C13 would have been largely unnecessary.`
- `scope_and_non_goals`: `This is for reusable workflow contracts, not every one-off response. It does not require verbose prose; it requires that operator-critical boundaries be explicit and truthful.`
- `pre_action_question`: `Can a fresh operator tell from this prompt alone what is local, what is injected, what gets saved, what artifact shape is supported, and what the worker does next?`
- `operational_check`: `Prompt rule: every reusable prompt must expose actor, local inputs, injected inputs, durable outputs, and continuation behavior in named sections before case details.`
- `confidence`: `strong`

### P03

- `principle_id`: `P03`
- `principle_statement`: `When a failure explanation or debugging branch depends on proof, collect the requested evidence artifact or state the blocker before hardening labels or switching to surrogate reasoning.`
- `decision_point`: `When naming the failure mode, selecting the next debug seam, or answering a direct request for proof.`
- `failure_signature`: `crash labels or root-cause stories appear before the concrete evidence object that would distinguish them`
- `why_this_is_durable`: `The same error surfaced both as an evidence-light failure label and as source/frame reasoning substituted for requested symbolicated stacks. In both events, explanation got ahead of proof.`
- `supporting_events`:
  - `C09`: "emulator crash" framing outran the evidence required to justify that label
  - `C10`: source-correlation reasoning replaced the requested symbolicated callstacks
- `supporting_human_model_signals`:
  - `C09`: "if you're saying theres an emulator crash, there should be evidence that supports that hypothesis"
  - `C10`: "hang on i asked you to get symbolicated callstacks"
- `counterfactual_prevention_claim`: `If applied earlier, C09 would have kept crash language provisional until process-level evidence existed, and C10 would have either delivered the symbolicated stack or explicitly named the blocker before drifting into source inference.`
- `scope_and_non_goals`: `This does not ban provisional hypotheses. It bans presenting them as the main answer or next seam when the proving artifact is still missing and no blocker has been surfaced.`
- `pre_action_question`: `Am I about to name a failure mode or root cause without the evidence object that would actually prove it?`
- `operational_check`: `Review gate: if a response contains a failure label or root-cause claim, it must also point to the proving artifact or explicitly mark the claim as unproven plus name the next proof step.`
- `confidence`: `medium`

### P04

- `principle_id`: `P04`
- `principle_statement`: `When the human names a concrete artifact target and shape, execute that exact deliverable directly and fill it with the requested evidence density; do not substitute drafts, cleaner decompositions, abbreviated proof, or underfilled packets without approval.`
- `decision_point`: `When editing a named file or assembling a handoff brief or evidence packet.`
- `failure_signature`: `responding with drafts, folders, summaries, or thin packets instead of the named file and requested evidence payload`
- `why_this_is_durable`: `Four events show the same delivery-time substitution habit. The agent preferred copy-paste drafts, multi-file neatness, abbreviated callstacks, or soft budget usage over the explicitly requested artifact and adequacy bar.`
- `supporting_events`:
  - `C11`: direct bug-doc edit request became copy-paste draft text
  - `C12`: requested single raw brief became a multi-file folder until corrected
  - `C14`: central callstack stayed abbreviated and the packet remained too thin
  - `C15`: repeated reruns were needed before source-first, near-cap evidence density was enforced
- `supporting_human_model_signals`:
  - `C11`: "I'm saying don't touch the directory. Touch the bug doc."
  - `C12`: "Put in a single raw brief md total file size 100KB"
  - `C14`: "it should reproduce the full callstack in its entirety"
  - `C15`: "source code related to the problem as the priority" and "you MUST fill it up to 100KB ... and report back"
- `counterfactual_prevention_claim`: `If applied earlier, C11 would have edited the requested bug document directly, C12 would have produced the right handoff shape immediately, and C14 through C15 would have reached the requested proof density with fewer reruns.`
- `scope_and_non_goals`: `This is not a ban on decomposition or summaries in general. It applies when the human has already specified the target artifact, transport shape, and evidence adequacy bar.`
- `pre_action_question`: `Has the human already specified the exact file, packet shape, or evidence bar, and am I about to substitute something cleaner for me but worse for them?`
- `operational_check`: `Closure gate: before presenting completion, compare the produced artifact against the named target file or shape, required raw evidence elements, and any numeric budget or shortfall-reporting rule.`
- `confidence`: `strong`

## Rejected Or Merged Principle Candidates

- `candidate_statement`: `Drop prior task identity whenever a new turn reissues a task id.`
  - `status`: `merged`
  - `reason`: `This is a reset-specific subcase of the stronger acceptance-lane grounding rule.`
  - `merged_into`: `P01`
- `candidate_statement`: `Before choosing a workaround, exhaust the most faithful no-code direct lane first.`
  - `status`: `merged`
  - `reason`: `Durable, but it is the debugging-specific expression of keeping work anchored to the real acceptance lane.`
  - `merged_into`: `P01`
- `candidate_statement`: `Shared prompts must stay generic across repos.`
  - `status`: `merged`
  - `reason`: `Too narrow alone. The durable rule is broader contract honesty about actor, scope, transport, and continuation.`
  - `merged_into`: `P02`
- `candidate_statement`: `Always inject imported analyzer responses inline and save them locally.`
  - `status`: `merged`
  - `reason`: `The durable lesson is not one transport mechanism in isolation; it is explicit and truthful handling of injected versus local inputs plus durable capture duties.`
  - `merged_into`: `P02`
- `candidate_statement`: `Workers that test a hypothesis must always continue debugging afterward.`
  - `status`: `merged`
  - `reason`: `This is one stop-condition clause within the larger reusable-prompt contract rule.`
  - `merged_into`: `P02`
- `candidate_statement`: `Never summarize callstacks.`
  - `status`: `rejected`
  - `reason`: `Too absolute. The durable rule is to preserve central callstacks in full when the requested proof packet requires them, captured under P04.`
- `candidate_statement`: `Always fill briefs to 100KB.`
  - `status`: `merged`
  - `reason`: `The numeric budget is local. The durable rule is to honor explicit evidence budgets and report shortfall instead of underfilling by preference.`
  - `merged_into`: `P04`
- `candidate_statement`: `Edit requested files directly instead of drafting copy-paste text.`
  - `status`: `merged`
  - `reason`: `This is the simplest expression of the broader exact-deliverable rule covering direct edits, artifact shape, and evidence density.`
  - `merged_into`: `P04`
- `candidate_statement`: `CURRENT-TASK.json should only contain a canonical state path and no duplicate identity fields.`
  - `status`: `rejected`
  - `reason`: `Supported only by the narrow C03-C04 cleanup arc so far. It looks real, but it is still too local and too same-implementation-specific to keep as one of the minimal cross-scope principles for this day.`

## Smallest Recommended Principle Set For This Scope

The smallest honest set is four principles:

- `P01`: stay grounded on the active task and real acceptance lane
- `P02`: make prompt and handoff contracts explicit and truthful
- `P03`: do not let theory outrun the requested proof artifact
- `P04`: execute the exact named deliverable in the exact requested artifact shape and evidence density

Collapsing further would blur two distinct decision points that mattered repeatedly on this day:

- prompt-contract truth versus delivery-time artifact execution
- acceptance-lane grounding versus proof-discipline during debugging

## Principles Still Too Weak For Standalone Use

- `single-source-of-truth / pointer minimization` from `C03` and `C04` still needs more recurrence outside the `CURRENT-TASK.json` cleanup arc before it deserves its own kept principle.
- No kept principle is too weak to use now, but `P01` and `P03` remain `medium` confidence and should be watched for reinforcement or narrowing on later days.

## Transcript Windows Reopened During PASS3

- None. PASS2 already provided the event boundaries, transcript lineage, repeated cluster hints, and strongest human-model signals needed for principle distillation on this source day.
