# INTERVENTION-PASS2 Analysis Report

Source day: `2026-03-28`

Primary input artifact: [INTERVENTION-PASS1.md](/c:/Agent/CodexDashboard/Tracking/Task-0006/Research/Daily/2026-03-28/INTERVENTION-PASS1.md)

## Source Scope Analyzed

- Re-read the `INTERVENTION-PASS1` candidate list at [INTERVENTION-PASS1.md](/c:/Agent/CodexDashboard/Tracking/Task-0006/Research/Daily/2026-03-28/INTERVENTION-PASS1.md).
- Re-read the current downstream incident corpus contract first: [README.md](/c:/Users/gregs/.codex/Orchestration/Reports/Incidents/README.md) and [INCIDENT.schema.json](/c:/Users/gregs/.codex/Orchestration/Reports/Incidents/incident.schema.json).
- Re-opened bounded raw transcript windows for every cited PASS1 candidate across these six transcripts:
- [rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl)
- [rollout-2026-03-28T09-45-14-019d34b1-1848-7ba2-99e0-73f1f7425302.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T09-45-14-019d34b1-1848-7ba2-99e0-73f1f7425302.jsonl)
- [rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl)
- [rollout-2026-03-28T14-08-52-019d35a2-7694-72f1-9a79-37ddd70cc3f9.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T14-08-52-019d35a2-7694-72f1-9a79-37ddd70cc3f9.jsonl)
- [rollout-2026-03-28T14-55-59-019d35cd-9961-7b72-8320-494688ce63cb.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T14-55-59-019d35cd-9961-7b72-8320-494688ce63cb.jsonl)
- [rollout-2026-03-28T21-50-41-019d3749-42f4-7d41-98e7-efa1d1f18771.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T21-50-41-019d3749-42f4-7d41-98e7-efa1d1f18771.jsonl)
- I did not use accepted-incident JSON files, same-day daily briefs, or parent-thread summaries.

## Candidate IDs Analyzed

- `C01`, `C02`, `C03`, `C04`, `C05`, `C06`, `C07`, `C08`, `C09`, `C10`, `C11`, `C12`, `C13`, `C14`, `C15`, `C16`, `C17`

## Boundary Corrections Relative To PASS1

- No PASS1 candidate was dropped after transcript reread.
- I kept `C03` and `C04` separate even though they are one `CURRENT-TASK.json` cleanup arc, because the human had to intervene a second time after the first simplification was already presented as complete.
- I kept `C05`, `C06`, `C07`, and `C08` separate even though they form one follow-up-template cluster, because the human corrected four different contract dimensions: durability, inline-vs-local input truth, bug-scoped naming plus injected-slot explicitness, and the follow-up worker's stopping condition.
- I kept `C14` separate from `C15`. `C14` is the first "full callstack, stop abbreviating" correction after the first rerun; `C15` is the later multi-stage tightening arc around source priority, measurable budget pressure, and shortfall explanation.
- I kept `C16` and `C17` separate. `C16` redefines the real goal at the human-world level; `C17` redirects the next search order and later materially changes what gets proven.
- I did not split `C15` further. The later source-priority, line-trade simplicity, `80-100 KB`, and hard `100 KB` corrections all answer the same local miss: the prompt still was not producing a source-heavy evidence packet that spent the available budget.

## Per-Event Analysis

### C01. Task grounding drifted onto `Task-0003` until the human reissued `task-0002`

- `event_id`: `C01`
- `title`: `Task grounding drifted onto Task-0003 until the human reissued task-0002`
- `session_or_thread`: `019d32b5-63ad-7a71-8239-5f9f0d4dddc0`
- `transcript_path`: [rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl)
- `primary_refs`: [L48](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl#L48), [L60](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl#L60), [L80](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl#L80), [L83](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl#L83)
- `ai_course`: At [L48](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl#L48) and [L60](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl#L60), the assistant says the acceptance lane is `Task 0003`, talks about missing HTTPS implementation, and starts reading the wrong task tree.
- `human_intervention`: At [L80](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl#L80), the human relaunches the agent explicitly as regression tester for `task-0002`; at [L83](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl#L83) the assistant regrounds on `0002`.
- `adequate_outcome`: Start from the actual active task and its real regression lane before reading docs or forming blocker stories.
- `event_boundary_notes`: PASS1's boundary is correct. The event begins once the assistant commits to the wrong task and ends once the human reissues the task id and the assistant regrounds.
- `human_model_signal`: `none explicit`; the human corrected scope by restating the task id rather than stating a broader rule.
- `failure_family_hypothesis`: `workflow_orchestration` - stale thread state overrode explicit task identity.
- `intervention_kind_hypothesis`: `other` - task re-grounding.
- `human_cost_or_risk`: Wasted work on the wrong acceptance lane, plus a false blocker story about missing `Task 0003` implementation.
- `local_lesson_hypothesis`: When a new turn explicitly reissues a task id, drop prior task identity immediately before reading repo artifacts.
- `cluster_hints`: `task-identity-grounding`, `stale-thread-state`, `wrong-acceptance-lane`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: The transcript shows the wrong task course clearly, but not whether the drift came from rollback residue or the assistant's own task inference.

### C02. The first `SUPERBRAIN` brief prompt came out task-specific and downstream-model-shaped until the human forced a generic local gatherer template

- `event_id`: `C02`
- `title`: `The first SUPERBRAIN brief prompt came out task-specific and downstream-model-shaped until the human forced a generic local gatherer template`
- `session_or_thread`: `019d32b5-63ad-7a71-8239-5f9f0d4dddc0`
- `transcript_path`: [rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl)
- `primary_refs`: [L843](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl#L843), [L850](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl#L850), [L853](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl#L853), [L872](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl#L872)
- `ai_course`: At [L843](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl#L843), the assistant ships a `SUPERBRAIN-DEBUG-BRIEF` prompt packed with current `Task-0002` Android/gRPC details and frames it as if it were prompting the stronger model directly.
- `human_intervention`: At [L850](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl#L850), the human says the template is supposed to be generic across repos, data-oriented, and aimed at a local Codex gatherer; the assistant commits to that rewrite at [L853](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl#L853) and reports the repaired role at [L872](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl#L872).
- `adequate_outcome`: A durable, repo-agnostic gatherer template that assembles local evidence for a stronger model instead of hard-coding one task's domain details.
- `event_boundary_notes`: This is one clean artifact-level correction from shipped prompt to repaired prompt.
- `human_model_signal`: Explicit: the prompt is "supposed to be generic across repos," should "keep it strictly data-oriented," and "you aren't prompting the superbrain."
- `failure_family_hypothesis`: `workflow_orchestration` - template role and abstraction level drifted from shared-orchestration use to one task instance.
- `intervention_kind_hypothesis`: `tighten_contract`
- `human_cost_or_risk`: A shared prompt would have encoded one repo's case and the wrong actor boundary, making future launches misleading.
- `local_lesson_hypothesis`: Shared orchestration prompts should declare actor, scope, and reuse boundary before case details are added.
- `cluster_hints`: `prompt-role-boundary`, `generic-template-vs-instance`, `gatherer-vs-analyzer`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: Low.

### C03. `CURRENT-TASK.json` duplicated task-document paths until the human stripped the extra intent

- `event_id`: `C03`
- `title`: `CURRENT-TASK.json duplicated task-document paths until the human stripped the extra intent`
- `session_or_thread`: `019d32b5-63ad-7a71-8239-5f9f0d4dddc0`
- `transcript_path`: [rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl)
- `primary_refs`: [L913](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl#L913), [L920](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl#L920), [L923](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl#L923), [L933](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl#L933)
- `ai_course`: At [L913](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl#L913), the assistant reports a new pointer file that still names `task_root`, `task_file`, `plan_file`, and `handoff_file`.
- `human_intervention`: At [L920](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl#L920), the human says those fields are "too much duplication of intent"; at [L923](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl#L923) the assistant trims the file, and [L933](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl#L933) reports the slimmer pointer.
- `adequate_outcome`: Keep the pointer file as a minimal anchor to canonical task state rather than a second place that restates task layout.
- `event_boundary_notes`: I kept this separate from `C04` because after this pass the assistant presented the file as cleaned up, yet a second redundant field still remained.
- `human_model_signal`: Explicit: "too much duplication of intent."
- `failure_family_hypothesis`: `information_architecture` - redundant sources of truth were added to a pointer contract.
- `intervention_kind_hypothesis`: `tighten_contract`
- `human_cost_or_risk`: Multiple duplicated task paths invite drift and make the pointer harder to trust.
- `local_lesson_hypothesis`: `CURRENT-TASK` should carry only the minimum anchor needed to discover canonical task artifacts.
- `cluster_hints`: `single-source-of-truth`, `pointer-minimization`, `contract-slimming`
- `accepted_incident_likelihood`: `unlikely`
- `confidence`: `medium`
- `uncertainties`: A later clustering pass may merge this with `C04`.

### C04. The current-task pointer still carried a redundant `task_id` after the first simplification pass

- `event_id`: `C04`
- `title`: `The current-task pointer still carried a redundant task_id after the first simplification pass`
- `session_or_thread`: `019d32b5-63ad-7a71-8239-5f9f0d4dddc0`
- `transcript_path`: [rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl)
- `primary_refs`: [L933](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl#L933), [L940](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl#L940), [L943](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl#L943), [L953](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl#L953)
- `ai_course`: At [L933](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl#L933), the assistant says the pointer now "just identifies `Task-0002` and points at the canonical task state file."
- `human_intervention`: At [L940](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl#L940), the human says "i guess also remove task_id"; the assistant agrees at [L943](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl#L943) and reports full removal at [L953](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl#L953).
- `adequate_outcome`: Infer task identity from the canonical state-file path rather than keeping another explicit `task_id`.
- `event_boundary_notes`: I kept this separate because it is a second human correction after the first cleanup was already presented as done.
- `human_model_signal`: Explicit but terse: "i guess also remove task_id."
- `failure_family_hypothesis`: `information_architecture` - one more redundant source of truth survived the first minimization pass.
- `intervention_kind_hypothesis`: `tighten_contract`
- `human_cost_or_risk`: Another redundant identity field undermines the whole point of a minimal pointer file.
- `local_lesson_hypothesis`: If the path already names the task unambiguously, do not add a second identity field without a demonstrated downstream need.
- `cluster_hints`: `single-source-of-truth`, `pointer-minimization`, `redundant-identity-field`
- `accepted_incident_likelihood`: `unlikely`
- `confidence`: `medium`
- `uncertainties`: A later pass may collapse `C03` and `C04` into one multi-step simplification incident.

### C05. The first follow-up template would have acted on an external-model response without preserving it durably

- `event_id`: `C05`
- `title`: `The first follow-up template would have acted on an external-model response without preserving it durably`
- `session_or_thread`: `019d34b1-1848-7ba2-99e0-73f1f7425302`
- `transcript_path`: [rollout-2026-03-28T09-45-14-019d34b1-1848-7ba2-99e0-73f1f7425302.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T09-45-14-019d34b1-1848-7ba2-99e0-73f1f7425302.jsonl)
- `primary_refs`: [L239](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T09-45-14-019d34b1-1848-7ba2-99e0-73f1f7425302.jsonl#L239), [L246](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T09-45-14-019d34b1-1848-7ba2-99e0-73f1f7425302.jsonl#L246), [L249](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T09-45-14-019d34b1-1848-7ba2-99e0-73f1f7425302.jsonl#L249)
- `ai_course`: At [L239](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T09-45-14-019d34b1-1848-7ba2-99e0-73f1f7425302.jsonl#L239), the assistant describes a follow-up prompt that takes `DEBUG-BRIEF-0001.md` plus the Superbrain analysis, but the response is still only assumed in launch context.
- `human_intervention`: At [L246](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T09-45-14-019d34b1-1848-7ba2-99e0-73f1f7425302.jsonl#L246), the human says the template "doesn't actually include that information" and suggests inline-plus-save handling; at [L249](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T09-45-14-019d34b1-1848-7ba2-99e0-73f1f7425302.jsonl#L249) the assistant agrees and distinguishes brief, analysis, and follow-up output.
- `adequate_outcome`: Inline the exact imported analysis at launch and save it durably as a task-owned artifact before continuing.
- `event_boundary_notes`: PASS1's boundary is right. This is the first time the human forces durability rather than relying on chat context.
- `human_model_signal`: Explicit: "Your template doesn't actually include that information" and "inline the superbrain response into the prompt, but also direct the debug agent to save that as a markdown file."
- `failure_family_hypothesis`: `workflow_orchestration` - imported analysis was going to be consumed without a durable local artifact.
- `intervention_kind_hypothesis`: `tighten_contract`
- `human_cost_or_risk`: Future agents and later sessions would not have the exact external analysis available in the task tree.
- `local_lesson_hypothesis`: Any external-model output that affects downstream local work should be durably captured before the local worker proceeds.
- `cluster_hints`: `durable-handoff`, `imported-analysis`, `context-vs-artifact`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: Low.

### C06. The follow-up prompt still misstated the inline-vs-local input boundary after the first durability fix

- `event_id`: `C06`
- `title`: `The follow-up prompt still misstated the inline-vs-local input boundary after the first durability fix`
- `session_or_thread`: `019d34b1-1848-7ba2-99e0-73f1f7425302`
- `transcript_path`: [rollout-2026-03-28T09-45-14-019d34b1-1848-7ba2-99e0-73f1f7425302.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T09-45-14-019d34b1-1848-7ba2-99e0-73f1f7425302.jsonl)
- `primary_refs`: [L318](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T09-45-14-019d34b1-1848-7ba2-99e0-73f1f7425302.jsonl#L318), [L325](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T09-45-14-019d34b1-1848-7ba2-99e0-73f1f7425302.jsonl#L325), [L328](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T09-45-14-019d34b1-1848-7ba2-99e0-73f1f7425302.jsonl#L328)
- `ai_course`: Even after the durability discussion, the prompt still described `DEBUG-BRIEF-<NNNN>.md` as the primary input and left the analyzer-response transport boundary ambiguous.
- `human_intervention`: At [L325](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T09-45-14-019d34b1-1848-7ba2-99e0-73f1f7425302.jsonl#L325), the human asks "Is this true? I thought we were doing inline."; at [L328](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T09-45-14-019d34b1-1848-7ba2-99e0-73f1f7425302.jsonl#L328) the assistant clarifies that only the analyzer response should be injected inline.
- `adequate_outcome`: State plainly that the brief is a local artifact input, while the imported analyzer response is an inline launch payload that must then be saved locally.
- `event_boundary_notes`: I kept this separate from `C05` because the first correction fixed durability, while this one fixes truthfulness about transport boundary.
- `human_model_signal`: Explicit: "Is this true? I thought we were doing inline."
- `failure_family_hypothesis`: `workflow_orchestration` - the prompt still misdescribed which inputs were local vs injected.
- `intervention_kind_hypothesis`: `tighten_contract`
- `human_cost_or_risk`: The launcher contract would remain misleading even after durability was fixed.
- `local_lesson_hypothesis`: When one input is local and another must be injected, the template should name that asymmetry directly instead of leaving it to implication.
- `cluster_hints`: `inline-vs-local-boundary`, `prompt-transport-honesty`, `imported-analysis`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: Low.

### C07. The follow-up trio still left bug-scoped naming and the injected analyzer slot too implicit

- `event_id`: `C07`
- `title`: `The follow-up trio still left bug-scoped naming and the injected analyzer slot too implicit`
- `session_or_thread`: `019d34b1-1848-7ba2-99e0-73f1f7425302`
- `transcript_path`: [rollout-2026-03-28T09-45-14-019d34b1-1848-7ba2-99e0-73f1f7425302.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T09-45-14-019d34b1-1848-7ba2-99e0-73f1f7425302.jsonl)
- `primary_refs`: [L338](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T09-45-14-019d34b1-1848-7ba2-99e0-73f1f7425302.jsonl#L338), [L345](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T09-45-14-019d34b1-1848-7ba2-99e0-73f1f7425302.jsonl#L345), [L355](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T09-45-14-019d34b1-1848-7ba2-99e0-73f1f7425302.jsonl#L355), [L383](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T09-45-14-019d34b1-1848-7ba2-99e0-73f1f7425302.jsonl#L383)
- `ai_course`: After the inline/local clarification, the prompt set still treated bug id, bug-scoped brief naming, and the analyzer-response insertion slot as too implicit.
- `human_intervention`: The human first invites more fixes at [L338](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T09-45-14-019d34b1-1848-7ba2-99e0-73f1f7425302.jsonl#L338); by [L355](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T09-45-14-019d34b1-1848-7ba2-99e0-73f1f7425302.jsonl#L355), the human says "For now we just have bug-scoped briefs" and the analyzer response "will always be injected into the template"; the assistant reports an explicit injected slot and bug-scoped naming at [L383](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T09-45-14-019d34b1-1848-7ba2-99e0-73f1f7425302.jsonl#L383).
- `adequate_outcome`: Encode the only currently supported workflow shape directly: bug-scoped brief naming plus an explicit inline analyzer-response slot.
- `event_boundary_notes`: I kept this separate from `C06` because it is about remaining prompt-shape ambiguity after the inline/local truth issue was already corrected.
- `human_model_signal`: Explicit: "For now we just have bug-scoped briefs" and "it will always be injected into the template."
- `failure_family_hypothesis`: `workflow_orchestration` - residual prompt ambiguity remained after the first repair.
- `intervention_kind_hypothesis`: `tighten_contract`
- `human_cost_or_risk`: Prompt users still had to infer filename scope and insertion location for imported analysis.
- `local_lesson_hypothesis`: Once a workflow only supports one real shape, the template should encode that shape directly rather than leave it as an optional convention.
- `cluster_hints`: `bug-scoped-artifacts`, `injected-slot-explicitness`, `prompt-contract-residue`
- `accepted_incident_likelihood`: `unlikely`
- `confidence`: `medium`
- `uncertainties`: A later clustering pass could collapse this into a broader `C05-C08` follow-up-template chain.

### C08. The follow-up prompt still aimed at evaluating a hypothesis rather than carrying on debugging

- `event_id`: `C08`
- `title`: `The follow-up prompt still aimed at evaluating a hypothesis rather than carrying on debugging`
- `session_or_thread`: `019d34b1-1848-7ba2-99e0-73f1f7425302`
- `transcript_path`: [rollout-2026-03-28T09-45-14-019d34b1-1848-7ba2-99e0-73f1f7425302.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T09-45-14-019d34b1-1848-7ba2-99e0-73f1f7425302.jsonl)
- `primary_refs`: [L392](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T09-45-14-019d34b1-1848-7ba2-99e0-73f1f7425302.jsonl#L392), [L399](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T09-45-14-019d34b1-1848-7ba2-99e0-73f1f7425302.jsonl#L399), [L402](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T09-45-14-019d34b1-1848-7ba2-99e0-73f1f7425302.jsonl#L402), [L412](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T09-45-14-019d34b1-1848-7ba2-99e0-73f1f7425302.jsonl#L412)
- `ai_course`: At [L392](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T09-45-14-019d34b1-1848-7ba2-99e0-73f1f7425302.jsonl#L392), the assistant points to where the injected analyzer response goes, but the prompt still reads like a one-shot evaluation worker.
- `human_intervention`: At [L399](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T09-45-14-019d34b1-1848-7ba2-99e0-73f1f7425302.jsonl#L399), the human says the directive should be closer to "take the superbrain's response and attempt to prove/disprove the hypothesis then carry on with your debugging efforts"; the assistant rewrites the runtime purpose at [L402](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T09-45-14-019d34b1-1848-7ba2-99e0-73f1f7425302.jsonl#L402) and reports the new continuation language at [L412](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T09-45-14-019d34b1-1848-7ba2-99e0-73f1f7425302.jsonl#L412).
- `adequate_outcome`: A continuation-debugging worker that treats imported analysis as a hypothesis to test and then keeps pushing the bug investigation forward.
- `event_boundary_notes`: This is separate from the prior prompt-shape fixes because it changes the worker's stopping condition.
- `human_model_signal`: Explicit: "take the superbrain's response and attempt to prove/disprove the hypothesis then carry on with your debugging efforts."
- `failure_family_hypothesis`: `workflow_orchestration` - the follow-up worker was framed as verdict generation instead of debugging continuation.
- `intervention_kind_hypothesis`: `tighten_contract`
- `human_cost_or_risk`: A downstream worker could stop at agreement or disagreement instead of making the bug investigation advance.
- `local_lesson_hypothesis`: When a worker consumes external analysis mid-debug, the prompt should say what happens after the hypothesis test, not just how the test is performed.
- `cluster_hints`: `continuation-not-judgment`, `follow-up-worker-role`, `imported-analysis`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: Low.

### C09. The emulator-crash read had to be backed by actual crash evidence instead of just an emerging theory

- `event_id`: `C09`
- `title`: `The emulator-crash read had to be backed by actual crash evidence instead of just an emerging theory`
- `session_or_thread`: `019d34e9-f366-78b0-ac5a-565f3fc0b222`
- `transcript_path`: [rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl)
- `primary_refs`: [L616](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl#L616), [L633](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl#L633), [L644](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl#L644), [L654](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl#L654)
- `ai_course`: At [L616](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl#L616), the assistant summarizes the situation using emulator-crash language and says `injectAudio` destabilizes the session itself.
- `human_intervention`: At [L633](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl#L633), the human says that if the claim is "emulator crash," there should be evidence supporting it and asks whether that should be investigated next; at [L644](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl#L644) the assistant agrees crash investigation is the highest-signal next branch.
- `adequate_outcome`: Treat crash language as a hypothesis that needs concrete process-level evidence before it becomes the leading explanation.
- `event_boundary_notes`: The event ends once the assistant explicitly converts the next pass into crash-cause collection.
- `human_model_signal`: Explicit: "if you're saying theres an emulator crash, there should be evidence that supports that hypothesis."
- `failure_family_hypothesis`: `verification_proof` - explanatory label outran directly cited evidence.
- `intervention_kind_hypothesis`: `redirect_debugging`
- `human_cost_or_risk`: Debug effort could branch around an overconfident label rather than a proven failure mode.
- `local_lesson_hypothesis`: Before hardening a failure-mode name, cite the concrete evidence that distinguishes it from weaker alternatives.
- `cluster_hints`: `proof-before-theory`, `evidence-before-label`, `crash-cause-branch`
- `accepted_incident_likelihood`: `possible`
- `confidence`: `medium`
- `uncertainties`: Later reviewers may read this as healthy scientific rigor rather than accepted-incident material because some crash-adjacent evidence already existed by the time of the correction.

### C10. The agent went into source/frame-IP reasoning when the ask was symbolicated callstacks

- `event_id`: `C10`
- `title`: `The agent went into source/frame-IP reasoning when the ask was symbolicated callstacks`
- `session_or_thread`: `019d34e9-f366-78b0-ac5a-565f3fc0b222`
- `transcript_path`: [rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl)
- `primary_refs`: [L2142](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl#L2142), [L2152](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl#L2152), [L2155](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl#L2155), [L2165](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl#L2165)
- `ai_course`: At [L2142](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl#L2142), the assistant claims a "critical source correlation" from frame `0` instruction/offset analysis and starts pinning exact source lines as the current best root cause.
- `human_intervention`: At [L2152](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl#L2152), the human stops the drift with "hang on i asked you to get symbolicated callstacks"; at [L2155](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl#L2155) the assistant admits it does not yet have a truly symbolicated vendor stack; at [L2165](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl#L2165) it returns to the build-from-source path.
- `adequate_outcome`: Produce the requested proof artifact or state the blocker first; do not substitute source inference for a missing symbolicated stack without approval.
- `event_boundary_notes`: PASS1's boundary is correct. This is a clear wrong-seam detour followed by an explicit admission and rerouting.
- `human_model_signal`: Explicit: "hang on i asked you to get symbolicated callstacks" and "stop and tell me why you're off in the weeds."
- `failure_family_hypothesis`: `verification_proof` - a requested proof object was replaced by surrogate reasoning.
- `intervention_kind_hypothesis`: `redirect_debugging`
- `human_cost_or_risk`: The assistant spent effort on a root-cause story before delivering the requested evidence artifact.
- `local_lesson_hypothesis`: When the human requests a specific proof artifact, do not substitute an inferred proxy without first naming the blocker and asking to change seams.
- `cluster_hints`: `wrong-seam-debugging`, `requested-deliverable-vs-surrogate`, `proof-artifact-discipline`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: Low.

### C11. The bug-note update kept turning into copy-paste drafts until the human forced the agent to do the edit

- `event_id`: `C11`
- `title`: `The bug-note update kept turning into copy-paste drafts until the human forced the agent to do the edit`
- `session_or_thread`: `019d34e9-f366-78b0-ac5a-565f3fc0b222`
- `transcript_path`: [rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl)
- `primary_refs`: [L2318](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl#L2318), [L2321](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl#L2321), [L2328](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl#L2328), [L2338](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl#L2338), [L2341](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl#L2341)
- `ai_course`: At [L2321](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl#L2321) and again at [L2331](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl#L2331), the assistant says it has not touched `BUG-0003.md` and offers draft sections for the human to paste later.
- `human_intervention`: The human request at [L2318](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl#L2318) was already to put next-step notes in bug `3`; after two draft responses, the human sharpens it at [L2338](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl#L2338): "I'm saying don't touch the directory. Touch the bug doc." The assistant finally commits to the edit at [L2341](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl#L2341).
- `adequate_outcome`: Edit the requested bug document directly instead of handing copy-paste work back to the human.
- `event_boundary_notes`: PASS1's boundary is right and the repeated draft responses matter; later compliance does not erase the earlier offload.
- `human_model_signal`: Explicit: "I'm saying don't touch the directory. Touch the bug doc."
- `failure_family_hypothesis`: `workflow_orchestration` - an execution request got downgraded into draft assistance.
- `intervention_kind_hypothesis`: `reject_outcome`
- `human_cost_or_risk`: The human had to police the control boundary and absorb avoidable copy-paste burden.
- `local_lesson_hypothesis`: If the request is to update a named artifact, default to doing the edit unless the human explicitly asks for draft-only text.
- `cluster_hints`: `do-the-work-not-draft`, `control-boundary`, `human-burden-shift`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: Low.

### C12. The evidence packet became a multi-file folder until the human forced a single raw brief under the size cap

- `event_id`: `C12`
- `title`: `The evidence packet became a multi-file folder until the human forced a single raw brief under the size cap`
- `session_or_thread`: `019d34e9-f366-78b0-ac5a-565f3fc0b222`
- `transcript_path`: [rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl)
- `primary_refs`: [L2480](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl#L2480), [L2508](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl#L2508), [L2511](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl#L2511), [L2566](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl#L2566)
- `ai_course`: At [L2480](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl#L2480), the assistant explicitly says it will package the callstack context as a small folder of separate files instead of one large markdown brief.
- `human_intervention`: At [L2508](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl#L2508), the human says "Put in a single raw brief md total file size 100KB. Keep stuffing context into it until it hits 100KB."; the assistant switches at [L2511](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl#L2511) and reports the single-file result at [L2566](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl#L2566).
- `adequate_outcome`: Produce one self-contained raw brief under the cap, shaped for downstream forwarding rather than a tidier local folder.
- `event_boundary_notes`: PASS1's boundary holds. The key local truth is that the assistant had already committed to the wrong artifact shape before the human overrode it.
- `human_model_signal`: Explicit: "Put in a single raw brief md total file size 100KB. Keep stuffing context into it until it hits 100KB."
- `failure_family_hypothesis`: `workflow_orchestration` - the artifact was shaped for internal neatness instead of downstream transport need.
- `intervention_kind_hypothesis`: `reject_outcome`
- `human_cost_or_risk`: A downstream handoff would have had the wrong artifact form and higher friction for the reviewer.
- `local_lesson_hypothesis`: When the human names the downstream artifact shape explicitly, optimize for that transport constraint instead of a cleaner internal decomposition.
- `cluster_hints`: `artifact-shape-for-handoff`, `single-file-packet`, `downstream-transport-needs`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: Low.

### C13. The prompt catalog hid the current-task and injection contract until the human forced a more honest model

- `event_id`: `C13`
- `title`: `The prompt catalog hid the current-task and injection contract until the human forced a more honest model`
- `session_or_thread`: `019d35a2-7694-72f1-9a79-37ddd70cc3f9`
- `transcript_path`: [rollout-2026-03-28T14-08-52-019d35a2-7694-72f1-9a79-37ddd70cc3f9.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T14-08-52-019d35a2-7694-72f1-9a79-37ddd70cc3f9.jsonl)
- `primary_refs`: [L134](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T14-08-52-019d35a2-7694-72f1-9a79-37ddd70cc3f9.jsonl#L134), [L141](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T14-08-52-019d35a2-7694-72f1-9a79-37ddd70cc3f9.jsonl#L141), [L152](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T14-08-52-019d35a2-7694-72f1-9a79-37ddd70cc3f9.jsonl#L152), [L162](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T14-08-52-019d35a2-7694-72f1-9a79-37ddd70cc3f9.jsonl#L162), [L169](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T14-08-52-019d35a2-7694-72f1-9a79-37ddd70cc3f9.jsonl#L169), [L209](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T14-08-52-019d35a2-7694-72f1-9a79-37ddd70cc3f9.jsonl#L209)
- `ai_course`: At [L134](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T14-08-52-019d35a2-7694-72f1-9a79-37ddd70cc3f9.jsonl#L134), the assistant explains that the debugger prompt does not know the current task automatically and treats current task identity as orchestrator-supplied context. At [L152](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T14-08-52-019d35a2-7694-72f1-9a79-37ddd70cc3f9.jsonl#L152), it acknowledges a real disconnect between `Context To Provide` and `Runtime Prompt`, but the catalog still encodes that split opaquely.
- `human_intervention`: The human asks at [L141](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T14-08-52-019d35a2-7694-72f1-9a79-37ddd70cc3f9.jsonl#L141) whether task id needs to be in the template and whether "Runtime Prompt" is misleading; at [L162](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T14-08-52-019d35a2-7694-72f1-9a79-37ddd70cc3f9.jsonl#L162) it presses on what channel actually transports `Context To Provide`; at [L169](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T14-08-52-019d35a2-7694-72f1-9a79-37ddd70cc3f9.jsonl#L169) it directs the prompts to tell agents to discover what they can from `CURRENT-TASK.json`; by [L209](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T14-08-52-019d35a2-7694-72f1-9a79-37ddd70cc3f9.jsonl#L209), the assistant is patching the catalog around discoverable vs injected context.
- `adequate_outcome`: Runtime prompts should teach agents how to discover stable local context; only true interpolation points should remain in an explicitly named injected-context section.
- `event_boundary_notes`: PASS1's boundary is correct. The later rename to `Context To Inject` belongs to the same correction arc rather than a separate event.
- `human_model_signal`: Explicit: "I thought it was all just the prompt," "just say see CURRENT-TASK.json," and "Instead of Context to Provide clarify its Context To Inject."
- `failure_family_hypothesis`: `workflow_orchestration` - the launcher/context transport contract was misleading and hid what the agent really knew at launch.
- `intervention_kind_hypothesis`: `tighten_contract`
- `human_cost_or_risk`: Prompt users could silently assume subagents had hidden context they did not actually have.
- `local_lesson_hypothesis`: Prompt templates should be explicit about discoverable local context versus true injected launch context, and should name the injection boundary honestly.
- `cluster_hints`: `prompt-transport-honesty`, `discoverable-vs-injected-context`, `current-task-contract`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: Low.

### C14. The first simplified brief prompt still abbreviated callstacks and stopped too early on evidence budget

- `event_id`: `C14`
- `title`: `The first simplified brief prompt still abbreviated callstacks and stopped too early on evidence budget`
- `session_or_thread`: `019d35cd-9961-7b72-8320-494688ce63cb`
- `transcript_path`: [rollout-2026-03-28T14-55-59-019d35cd-9961-7b72-8320-494688ce63cb.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T14-55-59-019d35cd-9961-7b72-8320-494688ce63cb.jsonl)
- `primary_refs`: [L293](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T14-55-59-019d35cd-9961-7b72-8320-494688ce63cb.jsonl#L293), [L300](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T14-55-59-019d35cd-9961-7b72-8320-494688ce63cb.jsonl#L300), [L303](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T14-55-59-019d35cd-9961-7b72-8320-494688ce63cb.jsonl#L303), [L333](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T14-55-59-019d35cd-9961-7b72-8320-494688ce63cb.jsonl#L333)
- `ai_course`: At [L293](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T14-55-59-019d35cd-9961-7b72-8320-494688ce63cb.jsonl#L293), the assistant reports that `BUG-0003-BRIEF-0004.md` is simplified and cleaner, but still only about `4 KB`, with the stack summarized and lots of budget left unused.
- `human_intervention`: At [L300](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T14-55-59-019d35cd-9961-7b72-8320-494688ce63cb.jsonl#L300), the human says the brief should not abbreviate callstacks and asks for prompt modification plus relaunch; the assistant updates the prompt at [L303](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T14-55-59-019d35cd-9961-7b72-8320-494688ce63cb.jsonl#L303) and reports improved but still thin output at [L333](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T14-55-59-019d35cd-9961-7b72-8320-494688ce63cb.jsonl#L333).
- `adequate_outcome`: Any relevant callstack should be reproduced in full, and the brief should spend far more of the evidence budget on raw proof.
- `event_boundary_notes`: This is the first post-simplification adequacy correction. I kept later source-priority and hard-budget ratcheting in `C15`.
- `human_model_signal`: Explicit: "it should reproduce the full callstack in its entirety."
- `failure_family_hypothesis`: `verification_proof` - a downstream proof packet was too compressed and too light on raw evidence.
- `intervention_kind_hypothesis`: `tighten_contract`
- `human_cost_or_risk`: The downstream brief was not audit-ready despite having plenty of room for raw stack and source evidence.
- `local_lesson_hypothesis`: If a callstack is central to the seam, preserve it verbatim before spending budget on explanatory prose.
- `cluster_hints`: `evidence-budget-discipline`, `full-callstack-not-summary`, `raw-packet-adequacy`
- `accepted_incident_likelihood`: `unlikely`
- `confidence`: `strong`
- `uncertainties`: Low.

### C15. Even after that rerun, the brief prompt still needed source-first and harder budget discipline before it behaved the way the human wanted

- `event_id`: `C15`
- `title`: `Even after that rerun, the brief prompt still needed source-first and harder budget discipline before it behaved the way the human wanted`
- `session_or_thread`: `019d35cd-9961-7b72-8320-494688ce63cb`
- `transcript_path`: [rollout-2026-03-28T14-55-59-019d35cd-9961-7b72-8320-494688ce63cb.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T14-55-59-019d35cd-9961-7b72-8320-494688ce63cb.jsonl)
- `primary_refs`: [L340](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T14-55-59-019d35cd-9961-7b72-8320-494688ce63cb.jsonl#L340), [L364](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T14-55-59-019d35cd-9961-7b72-8320-494688ce63cb.jsonl#L364), [L374](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T14-55-59-019d35cd-9961-7b72-8320-494688ce63cb.jsonl#L374), [L415](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T14-55-59-019d35cd-9961-7b72-8320-494688ce63cb.jsonl#L415), [L422](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T14-55-59-019d35cd-9961-7b72-8320-494688ce63cb.jsonl#L422), [L451](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T14-55-59-019d35cd-9961-7b72-8320-494688ce63cb.jsonl#L451)
- `ai_course`: Even after the first rerun, the assistant reports at [L333](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T14-55-59-019d35cd-9961-7b72-8320-494688ce63cb.jsonl#L333) that the brief is still far from the `100 KB` budget. The next iterations continue to rely on relatively soft prompt wording.
- `human_intervention`: At [L340](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T14-55-59-019d35cd-9961-7b72-8320-494688ce63cb.jsonl#L340), the human asks for more adjacent context and says source code related to the problem should be the priority; at [L374](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T14-55-59-019d35cd-9961-7b72-8320-494688ce63cb.jsonl#L374), it asks to keep instruction density constant while rerunning; at [L422](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T14-55-59-019d35cd-9961-7b72-8320-494688ce63cb.jsonl#L422), it hardens the budget to a near-must-fill `100 KB` rule with explanation if short; by [L451](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T14-55-59-019d35cd-9961-7b72-8320-494688ce63cb.jsonl#L451), the assistant reports the near-full `96,393` byte brief and the shortfall explanation.
- `adequate_outcome`: A simple prompt that makes failing-seam source code first, requires adjacent caller/callee/state/lifecycle context, and pushes the brief close to the full evidence budget with explicit shortfall reporting.
- `event_boundary_notes`: I kept this as one multi-stage tightening arc instead of splitting it. The human keeps correcting the same local miss: the prompt still was not coercing a sufficiently source-heavy, budget-filling evidence packet.
- `human_model_signal`: Explicit: "source code related to the problem as the priority," "If you add a line, remove another line," and "make the wording alot stronger like you MUST fill it up to 100KB ... and report back."
- `failure_family_hypothesis`: `workflow_orchestration` - prompt adequacy remained too soft even after the first repair.
- `intervention_kind_hypothesis`: `tighten_contract`
- `human_cost_or_risk`: Repeated reruns and supervision were needed before the prompt spent the available budget on the evidence the human actually wanted.
- `local_lesson_hypothesis`: When a prompt repeatedly produces thin packets, replace soft preference language with measurable fill rules and explicit evidence-priority ordering.
- `cluster_hints`: `evidence-budget-discipline`, `source-first-packetization`, `measurable-shortfall-rule`
- `accepted_incident_likelihood`: `possible`
- `confidence`: `medium`
- `uncertainties`: A later clustering pass may still want to split this into separate `source-priority` and `hard-budget-shortfall` events.

### C16. Root-cause analysis had to be pulled back toward the real goal: automated end-to-end regression through the real app and server

- `event_id`: `C16`
- `title`: `Root-cause analysis had to be pulled back toward the real goal: automated end-to-end regression through the real app and server`
- `session_or_thread`: `019d3749-42f4-7d41-98e7-efa1d1f18771`
- `transcript_path`: [rollout-2026-03-28T21-50-41-019d3749-42f4-7d41-98e7-efa1d1f18771.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T21-50-41-019d3749-42f4-7d41-98e7-efa1d1f18771.jsonl)
- `primary_refs`: [L210](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T21-50-41-019d3749-42f4-7d41-98e7-efa1d1f18771.jsonl#L210), [L217](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T21-50-41-019d3749-42f4-7d41-98e7-efa1d1f18771.jsonl#L217), [L227](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T21-50-41-019d3749-42f4-7d41-98e7-efa1d1f18771.jsonl#L227), [L230](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T21-50-41-019d3749-42f4-7d41-98e7-efa1d1f18771.jsonl#L230)
- `ai_course`: At [L210](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T21-50-41-019d3749-42f4-7d41-98e7-efa1d1f18771.jsonl#L210), the assistant gives a root-cause explanation about null dereference, forwarder registrants, and guest sound devices.
- `human_intervention`: After a clarifying goldfish-audio question at [L217](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T21-50-41-019d3749-42f4-7d41-98e7-efa1d1f18771.jsonl#L217), the human restates the actual problem at [L227](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T21-50-41-019d3749-42f4-7d41-98e7-efa1d1f18771.jsonl#L227): "automated end-to-end regression testing using the actual frontend android app and backend web service"; at [L230](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T21-50-41-019d3749-42f4-7d41-98e7-efa1d1f18771.jsonl#L230), the assistant pivots to looking for a workable path that keeps the real lane.
- `adequate_outcome`: Keep the investigation anchored on restoring the real app-plus-server regression lane rather than merely explaining emulator internals.
- `event_boundary_notes`: PASS1's boundary is good. This is a real problem-frame correction, not just a request for more detail.
- `human_model_signal`: Explicit: "the problem we're trying to solve is automated end-to-end regression testing using the actual frontend android app and backend web service."
- `failure_family_hypothesis`: `human_world` - local technical explanation had started to overshadow the human's actual success criterion.
- `intervention_kind_hypothesis`: `reframe_problem`
- `human_cost_or_risk`: Time could be spent on technically correct emulator analysis that still does not restore the acceptance lane the human actually needs.
- `local_lesson_hypothesis`: Root-cause work should stay subordinate to the human's real end-to-end success criterion, especially when the acceptance lane is user-shaped.
- `cluster_hints`: `real-goal-vs-local-root-cause`, `acceptance-target-drift`, `human-world-closure`
- `accepted_incident_likelihood`: `possible`
- `confidence`: `medium`
- `uncertainties`: A later pass may judge this as ordinary collaborative reprioritization rather than accepted-incident material.

### C17. The assistant proposed abandoning `injectAudio` for a virtual-mic workaround until the human forced a no-code rescue attempt first

- `event_id`: `C17`
- `title`: `The assistant proposed abandoning injectAudio for a virtual-mic workaround until the human forced a no-code rescue attempt first`
- `session_or_thread`: `019d3749-42f4-7d41-98e7-efa1d1f18771`
- `transcript_path`: [rollout-2026-03-28T21-50-41-019d3749-42f4-7d41-98e7-efa1d1f18771.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T21-50-41-019d3749-42f4-7d41-98e7-efa1d1f18771.jsonl)
- `primary_refs`: [L450](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T21-50-41-019d3749-42f4-7d41-98e7-efa1d1f18771.jsonl#L450), [L458](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T21-50-41-019d3749-42f4-7d41-98e7-efa1d1f18771.jsonl#L458), [L461](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T21-50-41-019d3749-42f4-7d41-98e7-efa1d1f18771.jsonl#L461), [L681](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T21-50-41-019d3749-42f4-7d41-98e7-efa1d1f18771.jsonl#L681)
- `ai_course`: At [L450](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T21-50-41-019d3749-42f4-7d41-98e7-efa1d1f18771.jsonl#L450), the assistant recommends abandoning emulator gRPC `injectAudio` as the primary lane and instead feeding known WAV audio through a virtual microphone on the normal host-mic path.
- `human_intervention`: At [L458](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T21-50-41-019d3749-42f4-7d41-98e7-efa1d1f18771.jsonl#L458), the human says "See if you can get the injectAudio path to work without code changes first."; the assistant accepts that directive at [L461](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T21-50-41-019d3749-42f4-7d41-98e7-efa1d1f18771.jsonl#L461), and by [L681](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T21-50-41-019d3749-42f4-7d41-98e7-efa1d1f18771.jsonl#L681) the assistant reports a real no-code `injectAudio` working path once the app is already recording.
- `adequate_outcome`: Exhaust the more direct/emulator-native lane before accepting a workaround that changes the operational sound-source mechanism.
- `event_boundary_notes`: PASS1's boundary is correct. The later success at [L681](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T21-50-41-019d3749-42f4-7d41-98e7-efa1d1f18771.jsonl#L681) is part of the same correction arc because it proves the human's redirect materially changed what got learned.
- `human_model_signal`: Explicit: "See if you can get the injectAudio path to work without code changes first."
- `failure_family_hypothesis`: `workflow_orchestration` - a workaround path was preferred before the more faithful direct lane had been fully exhausted.
- `intervention_kind_hypothesis`: `redirect_debugging`
- `human_cost_or_risk`: The task could have locked onto a less faithful workaround and missed a viable no-code salvage of the intended path.
- `local_lesson_hypothesis`: Before abandoning the direct lane, test the most faithful no-code salvage path if it is still plausibly available.
- `cluster_hints`: `direct-lane-before-workaround`, `no-code-salvage`, `real-lane-fidelity`
- `accepted_incident_likelihood`: `possible`
- `confidence`: `strong`
- `uncertainties`: Threshold depends on whether later synthesis treats strategy redirection as normal collaboration or as a durable "don't abandon the direct lane early" lesson.

## Likely Accepted Incidents

- `C01`: explicit wrong-task drift after a rollback/startover boundary
- `C02`: shared prompt template built at the wrong actor/scope level until the human corrected the role
- `C05`: imported external analysis would have been consumed without durable capture
- `C06`: the follow-up prompt still lied about inline-vs-local input shape after the first fix
- `C08`: the follow-up worker's stopping condition was wrong for the actual debugging workflow
- `C10`: requested symbolicated-callstack proof got replaced by source-correlation surrogate reasoning
- `C11`: requested artifact edit was pushed back to the human as a copy-paste draft
- `C12`: downstream handoff artifact shape was optimized for internal neatness instead of the human's stated transport need
- `C13`: the prompt catalog hid the true launch-context contract until the human forced a more honest model

## Likely Non-Incident But Still Important Intervention Events

- `C03`: the pointer file was overbuilt and needed slimming
- `C04`: one more redundant identity field survived the first slimming pass
- `C07`: prompt-shape cleanup around bug-scoped naming and injected-slot explicitness
- `C14`: first packet-quality correction after the simplified brief rerun

## Borderline / Threshold-Sensitive Events

- `C09`: strong local lesson about evidence-before-label, but incident weight depends on whether the earlier crash evidence is judged sufficient
- `C15`: durable prompt-adequacy lesson is real, but it may still be too intra-loop and tooling-specific for the accepted set
- `C16`: strong human-world reframing, but may read as normal collaborative reprioritization under a narrow incident threshold
- `C17`: later proof makes the redirection important, but some reviewers may still treat it as strategy coaching rather than incident-grade correction

## Repeated Cluster Hints Across The Set

- `prompt-transport-honesty`
  Prompt users kept encountering hidden or misleading launch-context assumptions across `C02`, `C05`, `C06`, `C07`, `C08`, and `C13`.
- `single-source-of-truth / contract-slimming`
  `C03` and `C04` both push against redundant pointer fields and duplicated intent.
- `proof-before-theory / wrong-seam-debugging`
  `C09` and `C10` are a direct pair: first evidence must support the failure-mode label, then the requested proof artifact must not be replaced by a surrogate seam.
- `do-the-work-not-draft / control-boundary`
  `C11` and, in a different form, `C12` both protect against shifting burden onto the human by choosing agent-convenient artifact shapes or outputs.
- `evidence-budget-discipline`
  `C12`, `C14`, and `C15` all show the same local adequacy bar: raw packets should spend the available budget on contiguous evidence, not on thin summary.
- `real-goal-vs-local-root-cause`
  `C16` and `C17` both insist that workaround or root-cause thinking stay subordinate to the real app/server regression goal.

## Strongest Human-Model Signals To Carry Forward

- `C03`: "too much duplication of intent"
- `C06`: "Is this true? I thought we were doing inline."
- `C10`: "hang on i asked you to get symbolicated callstacks"
- `C11`: "I'm saying don't touch the directory. Touch the bug doc."
- `C12`: "Put in a single raw brief md total file size 100KB"
- `C13`: "I thought it was all just the prompt" and "Context To Inject"
- `C16`: "the problem we're trying to solve is automated end-to-end regression testing using the actual frontend android app and backend web service"
- `C17`: "See if you can get the injectAudio path to work without code changes first"

## Events That Still Need A Wider Reread

- `C09` if a later pass wants to judge whether the crash hypothesis was already adequately evidenced before the human objected.
- `C15` if a later pass wants to split `source-priority` and `hard-budget-shortfall` into separate prompt-adequacy events.
- `C16` and `C17` if the accepted-set threshold later becomes strict enough that strategy redirection and human-world reframing need a second adjudication pass.
