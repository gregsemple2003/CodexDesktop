# INTERVENTION-PASS1 Candidate Report

Source day: `2026-03-28`

## Source Scope Reviewed

- Reviewed all `9` raw JSONL transcripts under [2026-03-28 session folder](/c:/Users/gregs/.codex/sessions/2026/03/28/).
- Read the current incident corpus contract first: [README.md](/c:/Users/gregs/.codex/Orchestration/Reports/Incidents/README.md), [INCIDENT.schema.json](/c:/Users/gregs/.codex/Orchestration/Reports/Incidents/INCIDENT.schema.json), and [OUTBOUND-MESSAGE-REVIEW.schema.json](/c:/Users/gregs/.codex/Orchestration/Reports/Incidents/OUTBOUND-MESSAGE-REVIEW.schema.json).
- I intentionally did not use same-day incident artifacts such as [2026-03-28 README.md](/c:/Users/gregs/.codex/Orchestration/Reports/Incidents/Daily/2026-03-28/README.md) or [OUTBOUND-MESSAGE-REVIEW.csv](/c:/Users/gregs/.codex/Orchestration/Reports/Incidents/Daily/2026-03-28/OUTBOUND-MESSAGE-REVIEW.csv); this pass is grounded from raw transcripts plus the downstream contract only.
- Direct-human candidate-bearing sessions:
  - [rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl)
  - [rollout-2026-03-28T09-45-14-019d34b1-1848-7ba2-99e0-73f1f7425302.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T09-45-14-019d34b1-1848-7ba2-99e0-73f1f7425302.jsonl)
  - [rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl)
  - [rollout-2026-03-28T14-08-52-019d35a2-7694-72f1-9a79-37ddd70cc3f9.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T14-08-52-019d35a2-7694-72f1-9a79-37ddd70cc3f9.jsonl)
  - [rollout-2026-03-28T14-55-59-019d35cd-9961-7b72-8320-494688ce63cb.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T14-55-59-019d35cd-9961-7b72-8320-494688ce63cb.jsonl)
  - [rollout-2026-03-28T21-50-41-019d3749-42f4-7d41-98e7-efa1d1f18771.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T21-50-41-019d3749-42f4-7d41-98e7-efa1d1f18771.jsonl)
- Screened but not promoted as direct intervention evidence:
  - [rollout-2026-03-28T09-34-42-019d34a7-714d-7db0-87bd-105f27a4d768.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T09-34-42-019d34a7-714d-7db0-87bd-105f27a4d768.jsonl): one human launch prompt and no later direct correction boundary.
  - [rollout-2026-03-28T14-40-56-019d35bf-d035-7343-90da-bd5d61cf1cfa.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T14-40-56-019d35bf-d035-7343-90da-bd5d61cf1cfa.jsonl): long debugger continuation with several human questions and requests, but I did not see a clean new "AI course was inadequate, do this instead" boundary that was not already better captured in the surrounding 10:47 and 21:50 sessions.
  - [rollout-2026-03-28T14-53-11-019d35cb-06a1-73c3-b7cf-90bb3402f4db.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T14-53-11-019d35cb-06a1-73c3-b7cf-90bb3402f4db.jsonl): short orchestration discussion about worker time horizons and checkpointing, but not a concrete correction of an already-inadequate AI course.

## Total Candidate Intervention Events Found

- Total: `17`
- `strong`: `11`
- `medium`: `6`
- `weak`: `0`

## Chronological Candidate List

### C01. Task grounding drifted onto `Task-0003` until the human reissued `task-0002`

Session or thread: `019d32b5-63ad-7a71-8239-5f9f0d4dddc0`

Transcript: [rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl)

Primary refs: [L48](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl#L48), [L60](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl#L60), [L80](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl#L80), [L83](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl#L83)

AI course, outcome, or stopping point: the assistant was actively reasoning about `Task-0003`, missing HTTPS implementation, and the wrong acceptance lane.

Human intervention summary: the human reissued the regression-tester prompt with explicit `"(task-0002)"` after the assistant had already started on the wrong task.

Concrete better outcome the human was forcing: ground on the actual active task and its real regression lane instead of spending effort on the wrong workstream.

Why this is a real intervention event rather than mere dissatisfaction: the assistant had already committed to a concrete wrong course; the human had to correct task identity to get the work back onto the intended repo state.

Confidence: `strong`

Accepted-set read: `likely accepted incident candidate`

### C02. The first `SUPERBRAIN` brief prompt came out task-specific and downstream-model-shaped until the human forced a generic local gatherer template

Session or thread: `019d32b5-63ad-7a71-8239-5f9f0d4dddc0`

Transcript: [rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl)

Primary refs: [L843](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl#L843), [L850](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl#L850), [L853](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl#L853), [L872](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl#L872)

AI course, outcome, or stopping point: the assistant produced a prompt loaded with current Task-0002 domain detail and framed as if it were prompting the stronger "superbrain" directly.

Human intervention summary: the human explicitly said the template should be generic across repos, strictly data-oriented, and aimed at a local Codex gatherer rather than the downstream model.

Concrete better outcome the human was forcing: a durable orchestration template that gathers local evidence packets without baking in one repo's Android/gRPC case or confusing gatherer and analyzer roles.

Why this is a real intervention event rather than mere dissatisfaction: the first prompt artifact already existed and was wrong in role, scope, and abstraction level; the human had to redirect the implementation target.

Confidence: `strong`

Accepted-set read: `likely accepted incident candidate`

### C03. `CURRENT-TASK.json` duplicated task-document paths until the human stripped the extra intent

Session or thread: `019d32b5-63ad-7a71-8239-5f9f0d4dddc0`

Transcript: [rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl)

Primary refs: [L913](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl#L913), [L920](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl#L920), [L923](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl#L923), [L933](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl#L933)

AI course, outcome, or stopping point: the new repo-root current-task pointer carried `task_root`, `task_file`, `plan_file`, and `handoff_file`, restating what the task tree already encoded elsewhere.

Human intervention summary: the human called that "too much duplication of intent" and told the assistant to remove those fields.

Concrete better outcome the human was forcing: a lean pointer file that names only the canonical current-task anchor instead of copying task layout into another place.

Why this is a real intervention event rather than mere dissatisfaction: the assistant had already created and validated the heavier shape; the human intervened to tighten the contract after seeing the concrete artifact.

Confidence: `medium`

Accepted-set read: `intervention event but probably not an accepted incident`

### C04. The current-task pointer still carried a redundant `task_id` after the first simplification pass

Session or thread: `019d32b5-63ad-7a71-8239-5f9f0d4dddc0`

Transcript: [rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl)

Primary refs: [L933](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl#L933), [L940](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl#L940), [L943](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl#L943), [L953](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T00-30-41-019d32b5-63ad-7a71-8239-5f9f0d4dddc0.jsonl#L953)

AI course, outcome, or stopping point: after the first trim, the pointer still explicitly carried `task_id`, even though the path to the state file already implied it.

Human intervention summary: the human followed up with `i guess also remove task_id`.

Concrete better outcome the human was forcing: infer task identity from the canonical state-file path and avoid another redundant source of truth.

Why this is a real intervention event rather than mere dissatisfaction: the human had to intervene again because the first simplification did not fully stick.

Confidence: `medium`

Accepted-set read: `intervention event but probably not an accepted incident`

### C05. The first follow-up template would have acted on an external-model response without preserving it durably

Session or thread: `019d34b1-1848-7ba2-99e0-73f1f7425302`

Transcript: [rollout-2026-03-28T09-45-14-019d34b1-1848-7ba2-99e0-73f1f7425302.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T09-45-14-019d34b1-1848-7ba2-99e0-73f1f7425302.jsonl)

Primary refs: [L239](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T09-45-14-019d34b1-1848-7ba2-99e0-73f1f7425302.jsonl#L239), [L246](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T09-45-14-019d34b1-1848-7ba2-99e0-73f1f7425302.jsonl#L246), [L249](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T09-45-14-019d34b1-1848-7ba2-99e0-73f1f7425302.jsonl#L249)

AI course, outcome, or stopping point: the new follow-up prompt assumed the stronger model's response would simply be present in launch context and acted on, but it did not make that response durable as a task-owned artifact.

Human intervention summary: the human said the template "doesn't actually include that information" and pushed for inline-plus-save behavior.

Concrete better outcome the human was forcing: preserve the exact imported external analysis in the task directory before the local debugger continues.

Why this is a real intervention event rather than mere dissatisfaction: without this correction the handoff would have left a dropped ball in the durable evidence trail; the human intervened to keep the imported analysis from disappearing into chat context.

Confidence: `strong`

Accepted-set read: `likely accepted incident candidate`

### C06. The follow-up prompt still misstated the inline-vs-local input boundary after the first durability fix

Session or thread: `019d34b1-1848-7ba2-99e0-73f1f7425302`

Transcript: [rollout-2026-03-28T09-45-14-019d34b1-1848-7ba2-99e0-73f1f7425302.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T09-45-14-019d34b1-1848-7ba2-99e0-73f1f7425302.jsonl)

Primary refs: [L318](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T09-45-14-019d34b1-1848-7ba2-99e0-73f1f7425302.jsonl#L318), [L325](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T09-45-14-019d34b1-1848-7ba2-99e0-73f1f7425302.jsonl#L325), [L328](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T09-45-14-019d34b1-1848-7ba2-99e0-73f1f7425302.jsonl#L328)

AI course, outcome, or stopping point: even after the save-file discussion, the template still presented `DEBUG-BRIEF-<NNNN>.md` as the main primary input and left the inline analyzer-response boundary muddy.

Human intervention summary: the human explicitly said `Is this true? I thought we were doing inline`.

Concrete better outcome the human was forcing: a prompt that clearly distinguishes local brief artifact access from the injected analyzer-response payload that must be captured and saved.

Why this is a real intervention event rather than mere dissatisfaction: this is a repeated correction because the first fix did not fully repair the contract the human was aiming for.

Confidence: `strong`

Accepted-set read: `likely accepted incident candidate`

### C07. The follow-up trio still left bug-scoped naming and the injected analyzer slot too implicit

Session or thread: `019d34b1-1848-7ba2-99e0-73f1f7425302`

Transcript: [rollout-2026-03-28T09-45-14-019d34b1-1848-7ba2-99e0-73f1f7425302.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T09-45-14-019d34b1-1848-7ba2-99e0-73f1f7425302.jsonl)

Primary refs: [L338](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T09-45-14-019d34b1-1848-7ba2-99e0-73f1f7425302.jsonl#L338), [L345](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T09-45-14-019d34b1-1848-7ba2-99e0-73f1f7425302.jsonl#L345), [L355](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T09-45-14-019d34b1-1848-7ba2-99e0-73f1f7425302.jsonl#L355), [L383](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T09-45-14-019d34b1-1848-7ba2-99e0-73f1f7425302.jsonl#L383)

AI course, outcome, or stopping point: the prompt set still treated brief scope and analyzer injection as too optional and too generic for the workflow the human was actually running.

Human intervention summary: the human pushed for bug-prefixed brief naming and said the analyzer response `will always be injected into the template`.

Concrete better outcome the human was forcing: make the prompt trio assume bug-scoped briefs for now and give the analyzer response an explicit inline slot instead of leaving it to implication.

Why this is a real intervention event rather than mere dissatisfaction: the human had already corrected the handoff mechanics once; this was a further intervention because ambiguity remained in the actual prompt artifacts.

Confidence: `medium`

Accepted-set read: `intervention event but probably not an accepted incident`

### C08. The follow-up prompt still aimed at evaluating a hypothesis rather than carrying on debugging

Session or thread: `019d34b1-1848-7ba2-99e0-73f1f7425302`

Transcript: [rollout-2026-03-28T09-45-14-019d34b1-1848-7ba2-99e0-73f1f7425302.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T09-45-14-019d34b1-1848-7ba2-99e0-73f1f7425302.jsonl)

Primary refs: [L392](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T09-45-14-019d34b1-1848-7ba2-99e0-73f1f7425302.jsonl#L392), [L399](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T09-45-14-019d34b1-1848-7ba2-99e0-73f1f7425302.jsonl#L399), [L402](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T09-45-14-019d34b1-1848-7ba2-99e0-73f1f7425302.jsonl#L402), [L412](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T09-45-14-019d34b1-1848-7ba2-99e0-73f1f7425302.jsonl#L412)

AI course, outcome, or stopping point: the prompt language still risked a workflow that tests agreement/disagreement with the stronger model and then stops.

Human intervention summary: the human explicitly reframed the directive to `take the superbrain's response and attempt to prove/disprove the hypothesis then carry on with your debugging efforts`.

Concrete better outcome the human was forcing: analysis follow-up as a continuation-debugging worker, not a one-shot agreement check.

Why this is a real intervention event rather than mere dissatisfaction: the human was tightening the operational contract after the assistant's prompt wording still left a too-weak stopping point.

Confidence: `strong`

Accepted-set read: `likely accepted incident candidate`

### C09. The emulator-crash read had to be backed by actual crash evidence instead of just an emerging theory

Session or thread: `019d34e9-f366-78b0-ac5a-565f3fc0b222`

Transcript: [rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl)

Primary refs: [L616](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl#L616), [L633](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl#L633), [L644](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl#L644), [L654](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl#L654)

AI course, outcome, or stopping point: the assistant was already using `emulator crash` language as the leading explanatory frame.

Human intervention summary: the human pushed back that if there is an emulator crash hypothesis, there should be evidence supporting it and asked to investigate that next.

Concrete better outcome the human was forcing: promote crash language only after concrete crash evidence exists, then branch into explicit crash-cause investigation.

Why this is a real intervention event rather than mere dissatisfaction: the human redirected the next debug branch because the current explanation was not yet warranted strongly enough by the visible evidence.

Confidence: `medium`

Accepted-set read: `intervention event but probably not an accepted incident`

### C10. The agent went into source/frame-IP reasoning when the ask was symbolicated callstacks

Session or thread: `019d34e9-f366-78b0-ac5a-565f3fc0b222`

Transcript: [rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl)

Primary refs: [L2142](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl#L2142), [L2152](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl#L2152), [L2155](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl#L2155), [L2165](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl#L2165)

AI course, outcome, or stopping point: the assistant was correlating the raw instruction pointer and source layouts to make root-cause claims without yet having a truly symbolicated vendor stack.

Human intervention summary: the human explicitly stopped that drift with `hang on i asked you to get symbolicated callstacks`.

Concrete better outcome the human was forcing: return to the requested build-and-symbol path rather than substitute source inference for the missing symbolicated evidence.

Why this is a real intervention event rather than mere dissatisfaction: the assistant was on an active but different seam from the one the human had asked for; the human had to pull it back to the warranted deliverable.

Confidence: `strong`

Accepted-set read: `likely accepted incident candidate`

### C11. The bug-note update kept turning into copy-paste drafts until the human forced the agent to do the edit

Session or thread: `019d34e9-f366-78b0-ac5a-565f3fc0b222`

Transcript: [rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl)

Primary refs: [L2318](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl#L2318), [L2321](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl#L2321), [L2328](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl#L2328), [L2338](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl#L2338), [L2341](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl#L2341)

AI course, outcome, or stopping point: when asked to put next-step notes into `BUG-0003.md`, the assistant twice replied with "here's a draft section you can drop in later" instead of making the requested edit.

Human intervention summary: the human had to clarify `I'm saying don't touch the directory. Touch the bug doc.`

Concrete better outcome the human was forcing: keep responsibility on the agent side and update the actual bug note directly instead of bouncing copy-paste work back to the human.

Why this is a real intervention event rather than mere dissatisfaction: the assistant had converted an execution request into a handoff-to-human draft; the human intervened to force completion rather than documentation-only assistance.

Confidence: `strong`

Accepted-set read: `likely accepted incident candidate`

### C12. The evidence packet became a multi-file folder until the human forced a single raw brief under the size cap

Session or thread: `019d34e9-f366-78b0-ac5a-565f3fc0b222`

Transcript: [rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl)

Primary refs: [L2480](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl#L2480), [L2508](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl#L2508), [L2511](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl#L2511), [L2566](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T10-47-20-019d34e9-f366-78b0-ac5a-565f3fc0b222.jsonl#L2566)

AI course, outcome, or stopping point: the assistant packaged the crash packet as a folder of smaller artifacts because that felt cleaner for forwarding.

Human intervention summary: the human overrode that packaging choice with `Put in a single raw brief md total file size 100KB`.

Concrete better outcome the human was forcing: one self-contained evidence packet in the exact form the downstream handoff needed, not a cleaner-but-split folder layout.

Why this is a real intervention event rather than mere dissatisfaction: the assistant had already committed to a concrete artifact shape that missed the human's handoff need; the human had to redirect the format itself.

Confidence: `strong`

Accepted-set read: `likely accepted incident candidate`

### C13. The prompt catalog hid the current-task and injection contract until the human forced a more honest model

Session or thread: `019d35a2-7694-72f1-9a79-37ddd70cc3f9`

Transcript: [rollout-2026-03-28T14-08-52-019d35a2-7694-72f1-9a79-37ddd70cc3f9.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T14-08-52-019d35a2-7694-72f1-9a79-37ddd70cc3f9.jsonl)

Primary refs: [L134](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T14-08-52-019d35a2-7694-72f1-9a79-37ddd70cc3f9.jsonl#L134), [L141](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T14-08-52-019d35a2-7694-72f1-9a79-37ddd70cc3f9.jsonl#L141), [L152](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T14-08-52-019d35a2-7694-72f1-9a79-37ddd70cc3f9.jsonl#L152), [L162](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T14-08-52-019d35a2-7694-72f1-9a79-37ddd70cc3f9.jsonl#L162), [L169](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T14-08-52-019d35a2-7694-72f1-9a79-37ddd70cc3f9.jsonl#L169), [L209](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T14-08-52-019d35a2-7694-72f1-9a79-37ddd70cc3f9.jsonl#L209)

AI course, outcome, or stopping point: the new `DEBUGGER.md` and catalog language treated task identity and launch context as if they were obvious, while the actual transport model lived only implicitly in orchestration behavior.

Human intervention summary: the human asked how the prompt knows the current task, challenged the meaning of `Runtime Prompt`, and then forced the catalog to prefer discovery from `CURRENT-TASK.json` for anything hard-codable.

Concrete better outcome the human was forcing: prompt templates whose runtime blocks tell agents how to find the current task from local artifacts and whose context sections only carry truly injected values.

Why this is a real intervention event rather than mere dissatisfaction: the human is correcting a misleading workflow contract that would otherwise cause prompt users to assume the subagent has context it does not actually have.

Confidence: `strong`

Accepted-set read: `likely accepted incident candidate`

### C14. The first simplified brief prompt still abbreviated callstacks and stopped too early on evidence budget

Session or thread: `019d35cd-9961-7b72-8320-494688ce63cb`

Transcript: [rollout-2026-03-28T14-55-59-019d35cd-9961-7b72-8320-494688ce63cb.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T14-55-59-019d35cd-9961-7b72-8320-494688ce63cb.jsonl)

Primary refs: [L293](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T14-55-59-019d35cd-9961-7b72-8320-494688ce63cb.jsonl#L293), [L300](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T14-55-59-019d35cd-9961-7b72-8320-494688ce63cb.jsonl#L300), [L303](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T14-55-59-019d35cd-9961-7b72-8320-494688ce63cb.jsonl#L303), [L333](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T14-55-59-019d35cd-9961-7b72-8320-494688ce63cb.jsonl#L333)

AI course, outcome, or stopping point: after the first simplified prompt run, the generated brief was still short, still used only a fraction of the budget, and still summarized a callstack instead of packetizing it in full.

Human intervention summary: the human explicitly said not to abbreviate callstacks and told the assistant to modify the prompt template and relaunch the worker.

Concrete better outcome the human was forcing: full raw callstack reproduction plus stronger pressure to spend the evidence budget on real crash packet content.

Why this is a real intervention event rather than mere dissatisfaction: the assistant had already run the worker once and inspected the thin result; the human intervened to tighten the prompt because the output still missed a concrete evidence bar.

Confidence: `strong`

Accepted-set read: `intervention event but probably not an accepted incident`

### C15. Even after that rerun, the brief prompt still needed source-first and harder budget discipline before it behaved the way the human wanted

Session or thread: `019d35cd-9961-7b72-8320-494688ce63cb`

Transcript: [rollout-2026-03-28T14-55-59-019d35cd-9961-7b72-8320-494688ce63cb.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T14-55-59-019d35cd-9961-7b72-8320-494688ce63cb.jsonl)

Primary refs: [L340](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T14-55-59-019d35cd-9961-7b72-8320-494688ce63cb.jsonl#L340), [L364](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T14-55-59-019d35cd-9961-7b72-8320-494688ce63cb.jsonl#L364), [L374](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T14-55-59-019d35cd-9961-7b72-8320-494688ce63cb.jsonl#L374), [L415](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T14-55-59-019d35cd-9961-7b72-8320-494688ce63cb.jsonl#L415), [L422](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T14-55-59-019d35cd-9961-7b72-8320-494688ce63cb.jsonl#L422), [L451](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T14-55-59-019d35cd-9961-7b72-8320-494688ce63cb.jsonl#L451)

AI course, outcome, or stopping point: the next brief iterations were better, but the worker still underused the budget and still needed more adjacent source context and a more measurable "don't stop early" rule.

Human intervention summary: the human first pushed source-code-first adjacent context, then explicitly said the softer directives were unlikely to change behavior, then escalated to a near-must-fill `100 KB` rule with shortfall explanation.

Concrete better outcome the human was forcing: a simple but much harder prompt that reliably spends nearly the full evidence budget on source, stacks, logs, and harness context instead of stopping once the brief feels coherent.

Why this is a real intervention event rather than mere dissatisfaction: this is a repeated tightening after concrete output still failed the human's adequacy bar; the assistant's iterative prompt changes were not sufficient until the human forced a much stronger measurable rule.

Confidence: `medium`

Accepted-set read: `intervention event but probably not an accepted incident`

### C16. Root-cause analysis had to be pulled back toward the real goal: automated end-to-end regression through the real app and server

Session or thread: `019d3749-42f4-7d41-98e7-efa1d1f18771`

Transcript: [rollout-2026-03-28T21-50-41-019d3749-42f4-7d41-98e7-efa1d1f18771.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T21-50-41-019d3749-42f4-7d41-98e7-efa1d1f18771.jsonl)

Primary refs: [L210](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T21-50-41-019d3749-42f4-7d41-98e7-efa1d1f18771.jsonl#L210), [L217](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T21-50-41-019d3749-42f4-7d41-98e7-efa1d1f18771.jsonl#L217), [L227](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T21-50-41-019d3749-42f4-7d41-98e7-efa1d1f18771.jsonl#L227), [L230](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T21-50-41-019d3749-42f4-7d41-98e7-efa1d1f18771.jsonl#L230)

AI course, outcome, or stopping point: the assistant was deep in source-level crash explanation about forwarder callbacks, null voices, and guest audio devices.

Human intervention summary: the human explicitly restated the actual problem: automated end-to-end regression testing through the real Android frontend and backend web service, not just source-theory correctness.

Concrete better outcome the human was forcing: a practical path that keeps the real UI and real packaged server in play instead of staying at the level of emulator internals only.

Why this is a real intervention event rather than mere dissatisfaction: the human is reframing the work back to the human-world outcome after the assistant drifted into a lower-level explanatory seam.

Confidence: `medium`

Accepted-set read: `intervention event but probably not an accepted incident`

### C17. The assistant proposed abandoning `injectAudio` for a virtual-mic workaround until the human forced a no-code rescue attempt first

Session or thread: `019d3749-42f4-7d41-98e7-efa1d1f18771`

Transcript: [rollout-2026-03-28T21-50-41-019d3749-42f4-7d41-98e7-efa1d1f18771.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T21-50-41-019d3749-42f4-7d41-98e7-efa1d1f18771.jsonl)

Primary refs: [L450](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T21-50-41-019d3749-42f4-7d41-98e7-efa1d1f18771.jsonl#L450), [L458](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T21-50-41-019d3749-42f4-7d41-98e7-efa1d1f18771.jsonl#L458), [L461](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T21-50-41-019d3749-42f4-7d41-98e7-efa1d1f18771.jsonl#L461), [L681](/c:/Users/gregs/.codex/sessions/2026/03/28/rollout-2026-03-28T21-50-41-019d3749-42f4-7d41-98e7-efa1d1f18771.jsonl#L681)

AI course, outcome, or stopping point: the assistant's best-path recommendation was to stop depending on emulator gRPC `injectAudio` for the primary lane and instead use a virtual microphone on the normal host-mic path.

Human intervention summary: the human redirected that plan with `See if you can get the injectAudio path to work without code changes first`.

Concrete better outcome the human was forcing: exhaust the emulator-native, no-code `injectAudio` path before accepting a less direct workaround that changes the sound-source mechanism.

Why this is a real intervention event rather than mere dissatisfaction: the assistant had already proposed a concrete workaround direction; the human redirected the next pass to salvage the more faithful lane first, and the later session evidence shows that redirection materially changed what got proven.

Confidence: `strong`

Accepted-set read: `intervention event but probably not an accepted incident`

## Which Candidates Look Like Likely Accepted Incidents

- `C01` task identity drifted to the wrong task until the human reissued `task-0002`
- `C02` a shared orchestration prompt was produced as a task-specific downstream-model prompt until the human forced the correct generic gatherer role
- `C05` the follow-up prompt would have consumed external-model output without preserving it durably
- `C06` the follow-up prompt still got the inline-vs-local input contract wrong after the first correction
- `C08` the follow-up prompt still aimed at "evaluate the hypothesis" instead of "test it and keep debugging"
- `C10` the assistant drifted from the requested symbolicated-callstack path into unsupported source inference
- `C11` the assistant offloaded a requested bug-note edit back to the human as a copy-paste draft
- `C12` the assistant chose a cleaner multi-file packet when the human actually needed one raw single-file brief
- `C13` the prompt catalog hid the true current-task/injection contract until the human forced a more honest model

## Which Candidates Are Real Interventions But Probably Belong Outside The Accepted Incident Set

- `C03` current-task pointer duplicated task-doc paths until the human trimmed them
- `C04` current-task pointer still carried a redundant `task_id` after the first simplification
- `C07` the follow-up trio still needed bug-scoped naming and an explicit analyzer slot after earlier fixes
- `C09` the human forced stronger crash evidence before letting the crash hypothesis harden
- `C14` the brief prompt still abbreviated callstacks and underused evidence budget until the human tightened it
- `C15` the human kept ratcheting up source-priority and budget rules because prompt iterations were still too thin
- `C16` the human reframed source-level explanation back toward the actual end-to-end regression goal
- `C17` the human redirected the search away from a virtual-mic workaround and back toward a no-code `injectAudio` rescue attempt first

## Ambiguous Boundaries That Need A Second Read

- `C03` and `C04` may collapse into one broader "CURRENT-TASK pointer was overbuilt and had to be simplified in multiple passes" event on reread. I kept them separate because the human had to intervene twice after the first trim.
- `C09` is clearly corrective, but it may read as ordinary scientific rigor rather than a durable incident if a reviewer sees the assistant's preexisting evidence as already sufficient.
- `C15` is a real repeated tightening cluster, but a second read could split it into finer sub-events around source priority, measurable budget targets, and the final shortfall-report rule.
- `C16` and `C17` are both real redirects, but they sit close to normal collaborative strategy selection. On reread, a reviewer could keep one and drop the other if they want a narrower accepted-incident threshold.
