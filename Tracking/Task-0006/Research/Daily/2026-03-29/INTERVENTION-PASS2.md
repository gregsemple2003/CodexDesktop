# INTERVENTION-PASS2 Analysis Report

Source day: `2026-03-29`

## Source Scope Analyzed

- Primary input artifact: [INTERVENTION-PASS1.md](/c:/Agent/CodexDashboard/Tracking/Task-0006/Research/Daily/2026-03-29/INTERVENTION-PASS1.md)
- Incident corpus contract reread for classification legibility:
  - [README.md](/c:/Users/gregs/.codex/Orchestration/Reports/Incidents/README.md)
  - [INCIDENT.schema.json](/c:/Users/gregs/.codex/Orchestration/Reports/Incidents/INCIDENT.schema.json)
- Raw transcript rereads reopened from PASS1:
  - [rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl)
  - [rollout-2026-03-29T16-17-38-019d3b3e-b474-75a2-a2d2-254c407cd900.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T16-17-38-019d3b3e-b474-75a2-a2d2-254c407cd900.jsonl)
  - [rollout-2026-03-29T23-59-34-019d3ce5-9eef-7c61-922f-bb45f3f5a5d4.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T23-59-34-019d3ce5-9eef-7c61-922f-bb45f3f5a5d4.jsonl)
- I also spot-checked current accepted-incident examples to keep family and intervention-kind labels aligned with the existing corpus:
  - [INCIDENT-0001.json](/c:/Users/gregs/.codex/Orchestration/Reports/Incidents/Daily/2026-04-04/INCIDENT-0001.json)
  - [INCIDENT-0002.json](/c:/Users/gregs/.codex/Orchestration/Reports/Incidents/Daily/2026-04-04/INCIDENT-0002.json)
  - [INCIDENT-0001.json](/c:/Users/gregs/.codex/Orchestration/Reports/Incidents/Daily/2026-04-06/INCIDENT-0001.json)

## Candidate IDs Analyzed

- `C01`
- `C02`
- `C03`
- `C04`
- `C05`
- `C06`
- `C07`
- `C08`
- `C09`
- `C10`
- `C11`
- `C12`
- `C13`
- `C14`

## Candidate Boundary Corrections Relative To PASS1

- No candidate ids needed to be merged or split.
- `C01` is one staged correction arc, not two incidents. The first human pushback rejects the worksheet shape, and the later pushback sharpens that into the stricter "facts on the ground + 3 max open-ended questions" standard.
- `C03` should be read as assistant self-diagnosis plus human authorization to pivot. The decisive human move is the approval to regenerate under the coherence-first philosophy, not the diagnosis by itself.
- `C08` starts with the new checkpointing request, but the decisive intervention is later, when the human explicitly questions whether `commit` / `push` / `notify` belong in canonical task state at all.
- `C11` and `C12` remain distinct. `C11` is the rejection of the already-claimed Task 0004 closure outcome. `C12` is the follow-on authority and precedence hardening that explains why the bad closure call was even possible.
- `C13` and `C14` remain distinct. `C13` reclassifies the CA trust step from blocker to acceptable development friction. `C14` hardens the future workflow rule for what must happen when the required regression lane fails or is blocked.
- `C14` needs one factual tightening relative to PASS1: in this exact task, the assistant had already opened [BUG-0001.md](/c:/Agent/Crystallize/Tracking/Task-0004/BUG-0001.md) before the human's prompt hardening request. The event is still real, but the local miss is "the workflow default was not yet hardened" rather than "this run stopped with no bug artifact at all."

## Per-Event Analysis Records

### C01. The Task-0003 handoff brief still boxed the stronger external model into a worksheet

- `event_id`: `C01`
- `title`: The Task-0003 handoff brief still boxed the stronger external model into a worksheet
- `session_or_thread`: `019d3aad-74a9-7600-aed8-c70a4c12ee67` (`Task-0003` research-brief thread)
- `transcript_path`: [rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl)
- `primary_refs`: [L279](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl#L279), [L284](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl#L284), [L327](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl#L327), [L332](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl#L332), [L337](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl#L337)
- `ai_course`: The assistant generated a send-to-model handoff that explicitly told the external model to answer using the exact requested answer shape, then after one revision still kept a fairly guided "here are the key questions, give us the implementation recommendation" frame.
- `human_intervention`: The human rejected the worksheet-like framing twice: first by saying the stronger model should get facts plus key questions rather than a constrained answer shape, then by sharpening that into "facts on the ground" plus `3 MAX` very open-ended questions.
- `adequate_outcome`: A fact-heavy, low-steering brief that exposes the real world-state and asks a few open questions, without boxing the downstream model into the local task list or a preset answer template.
- `event_boundary_notes`: This is one staged correction arc. The later "still too specific" intervention deepens the earlier correction rather than forming a second event.
- `human_model_signal`: Strong explicit signal. The human says the brief should give "all the facts I think are relevant" and ask `3 MAX very open-ended questions`; "we want its thoughts not a prepackaged template."
- `failure_family_hypothesis`: `information_architecture` with some `workflow_orchestration` overlap. The miss is mainly about the structure and steering properties of the handoff artifact.
- `intervention_kind_hypothesis`: `reject_outcome`
- `human_cost_or_risk`: The stronger external model would be pushed toward compliance with the local worksheet instead of offering independent thinking. The human had to reread the artifact and restate the actual philosophy of the research handoff.
- `local_lesson_hypothesis`: When handing off to a stronger outside model, default to world-state plus a few open questions. Do not impose answer shapes unless the human explicitly asks for that.
- `cluster_hints`: `external-handoff-overconstraint`, `brief-as-evidence-packet`, `open-ended-question-ceiling`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: The assistant quickly agreed with the human, so the transcript mostly preserves the human's standard rather than a contested back-and-forth. That is enough for this event.

### C02. The hard 100 KB rule was being fixed in the leader instead of a separate briefer

- `event_id`: `C02`
- `title`: The hard 100 KB rule was being fixed in the leader instead of a separate briefer
- `session_or_thread`: `019d3aad-74a9-7600-aed8-c70a4c12ee67` (`Task-0003` research-brief thread)
- `transcript_path`: [rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl)
- `primary_refs`: [L610](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl#L610), [L615](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl#L615), [L629](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl#L629), [L634](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl#L634)
- `ai_course`: After being asked to use the `SUPERBRAIN-DEBUG-BRIEF` style for the `100 KB` target, the assistant updated `LEAD-RESEARCHER.md` and related plan text because that was where briefing behavior currently lived.
- `human_intervention`: The human stopped on the abstraction seam and asked why the hard brief-size rule was being pushed into the lead at all, since what he was advocating for was a briefer-specific contract.
- `adequate_outcome`: A cleaner split where the leader coordinates and a dedicated briefer owns the hard `100 KB` constraint and evidence-packet rules.
- `event_boundary_notes`: Short, self-contained architecture-seam correction. It stays distinct from `C01` and `C03`, which are about brief quality and handoff shape rather than role placement.
- `human_model_signal`: Explicit. "What I'm advocating for is asking the briefer subagent for a hard 100KB brief size, but you're making changes in the lead."
- `failure_family_hypothesis`: `workflow_orchestration`
- `intervention_kind_hypothesis`: `reframe_problem`
- `human_cost_or_risk`: Prompt roles drift together, making later maintenance harder and forcing the human to clean up architecture rather than just evaluate the brief behavior.
- `local_lesson_hypothesis`: Put a rule on the component that actually owns the behavior. Do not encode a briefer contract into the coordinator just because that is the current implementation seam.
- `cluster_hints`: `role-boundary-drift`, `wrong-seam-fix`, `leader-vs-briefer-contract`
- `accepted_incident_likelihood`: `possible`
- `confidence`: `medium`
- `uncertainties`: The architecture was in flux during the thread, so part of the confusion was transitional. The intervention is still real because the human corrected the chosen seam explicitly.

### C03. The 100 KB evidence-packet philosophy produced repo-dump briefs instead of a readable handoff

- `event_id`: `C03`
- `title`: The 100 KB evidence-packet philosophy produced repo-dump briefs instead of a readable handoff
- `session_or_thread`: `019d3aad-74a9-7600-aed8-c70a4c12ee67` (`Task-0003` research-brief thread)
- `transcript_path`: [rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl)
- `primary_refs`: [L913](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl#L913), [L918](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl#L918), [L923](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl#L923), [L973](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl#L973), [L991](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl#L991)
- `ai_course`: The assistant had optimized the briefer around size and evidence density, which yielded a large artifact that read like concatenated repo files rather than a coherent brief.
- `human_intervention`: After the assistant diagnosed the failure mode, the human authorized the pivot: "let's try it your way" and asked for one regenerated brief link instead of more oversized packet iteration.
- `adequate_outcome`: A readable main brief with editorial synthesis, concrete seams, and optional appendices only if truly needed.
- `event_boundary_notes`: PASS1's event stands, but the decisive human move is the authorization to regenerate under the smaller coherence-first philosophy. The diagnosis alone is not the intervention.
- `human_model_signal`: `none explicit` beyond the human's acceptance of the coherence-first pivot. The strongest explanation in this window is the assistant's diagnosis, not a fresh human principle statement.
- `failure_family_hypothesis`: `information_architecture`
- `intervention_kind_hypothesis`: `reject_outcome`
- `human_cost_or_risk`: The external handoff would consume attention and tokens without being pleasant or useful to read. The human had to spend another round evaluating whether the artifact was even sendable.
- `local_lesson_hypothesis`: Size and evidence density need an explicit readability gate. A brief can be context-rich and still fail if the main body does not curate, de-duplicate, and synthesize.
- `cluster_hints`: `readability-vs-dump`, `evidence-packet-overfit`, `curate-dont-transcribe`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `medium`
- `uncertainties`: If a later accepted-incident writer needs stronger human-owned wording for the failure, a slightly wider reread just before [L913](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl#L913) could help recover the human's dissatisfaction with the earlier brief in more direct terms.

### C04. The handoff still risked implying that the downstream ChatGPT target could open local files

- `event_id`: `C04`
- `title`: The handoff still risked implying that the downstream ChatGPT target could open local files
- `session_or_thread`: `019d3aad-74a9-7600-aed8-c70a4c12ee67` (`Task-0003` research-brief thread)
- `transcript_path`: [rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl)
- `primary_refs`: [L1007](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl#L1007), [L1012](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl#L1012), [L1017](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl#L1017), [L1040](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl#L1040)
- `ai_course`: The assistant added a `Relevant Files` map to the brief, but the no-filesystem assumption for the downstream ChatGPT handoff target was only implied, not stated as a hard limit.
- `human_intervention`: The human explicitly raised the real target model context: GPT 5.4 Pro in ChatGPT has no local file access. He asked whether that was clear in the prompts.
- `adequate_outcome`: Prompt stack and live brief explicitly state that the downstream model has no local filesystem, no repo checkout, and that local paths are descriptive only.
- `event_boundary_notes`: Distinct from `C01`. `C01` is about answer-shape steering. `C04` is about truthful representation of the recipient's actual capabilities.
- `human_model_signal`: Strong explicit signal. "The model that I'm sending it to doesn't have access to local files."
- `failure_family_hypothesis`: `human_world`
- `intervention_kind_hypothesis`: `tighten_contract`
- `human_cost_or_risk`: The handoff could quietly overstate the evidence available to the downstream model, producing fake certainty and a weaker external answer.
- `local_lesson_hypothesis`: Whenever a handoff artifact names local paths for an external model, say explicitly whether those paths are actionable or descriptive only.
- `cluster_hints`: `access-truth`, `downstream-capability-clarity`, `real-world-handoff-assumptions`
- `accepted_incident_likelihood`: `possible`
- `confidence`: `strong`
- `uncertainties`: None material.

### C05. RESEARCH-LEADER could create RESEARCH-PLAN.md without being told to use the shared exemplar

- `event_id`: `C05`
- `title`: RESEARCH-LEADER could create `RESEARCH-PLAN.md` without being told to use the shared exemplar
- `session_or_thread`: `019d3aad-74a9-7600-aed8-c70a4c12ee67` (`Task-0003` research-brief thread)
- `transcript_path`: [rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl)
- `primary_refs`: [L1127](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl#L1127), [L1132](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl#L1132), [L1137](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl#L1137), [L1142](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl#L1142), [L1155](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl#L1155), [L1169](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl#L1169)
- `ai_course`: The assistant explained that `RESEARCH-LEADER` knew it may need to create or refresh `RESEARCH-PLAN.md`, but conceded that the prompt did not explicitly tell the leader to use the shared exemplar for artifact shape.
- `human_intervention`: The human asked how the leader would know the plan's shape, agreed that explicit exemplar binding was needed, then asked for other unboundedness issues before approving a fix.
- `adequate_outcome`: When `RESEARCH-LEADER` creates or refreshes `RESEARCH-PLAN.md`, it must read and follow the shared exemplar, with some bounded definition of when a plan is stale.
- `event_boundary_notes`: I keep this event focused on artifact-shape guidance. The broader unboundedness review is context, not a separate candidate in PASS1.
- `human_model_signal`: Explicit. "So how would it know what shape the research plan should take?"
- `failure_family_hypothesis`: `workflow_orchestration`
- `intervention_kind_hypothesis`: `tighten_contract`
- `human_cost_or_risk`: A leader that can generate durable artifacts without exemplar guidance will drift artifact shape over time, forcing the human to restate structure manually.
- `local_lesson_hypothesis`: If a prompt may mint a canonical artifact, it needs an explicit pointer to the canonical exemplar, not just the artifact filename.
- `cluster_hints`: `artifact-shape-drift`, `exemplar-binding`, `leader-unboundedness`
- `accepted_incident_likelihood`: `possible`
- `confidence`: `strong`
- `uncertainties`: The transcript also surfaces other research-leader unboundedness issues, but only the exemplar-guidance fix is directly accepted and closed inside this event.

### C06. Steering questions paused Stage B synthesis instead of the leader preserving the main line

- `event_id`: `C06`
- `title`: Steering questions paused Stage B synthesis instead of the leader preserving the main line
- `session_or_thread`: `019d3aad-74a9-7600-aed8-c70a4c12ee67` (`Task-0003` research-brief thread)
- `transcript_path`: [rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl)
- `primary_refs`: [L1268](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl#L1268), [L1273](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl#L1273), [L1277](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl#L1277), [L1282](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl#L1282), [L1287](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl#L1287)
- `ai_course`: All reconciliation subagents had completed, but the durable Stage B synthesis into `RESEARCH.md` and `HANDOFF.md` had not happened because the assistant treated mid-run steering questions as a real pause in the main work.
- `human_intervention`: The human asked directly whether the assistant was waiting on him and whether steering questions interrupt the work. After the explanation, he explicitly told it to continue and write everything out.
- `adequate_outcome`: Clarifying or steering questions should not silently cancel the umbrella instruction to finish the research synthesis unless the human explicitly changes scope.
- `event_boundary_notes`: This is the softest candidate in the set. The transcript does show that the main artifact was still unwritten after all subordinate work had finished, but the conversational interruption makes the failure boundary less crisp than the other events.
- `human_model_signal`: Partial but real. The human tests an adequacy rule about continuity: "You were working ... does that interrupt you?"
- `failure_family_hypothesis`: `workflow_orchestration`
- `intervention_kind_hypothesis`: `tighten_contract`
- `human_cost_or_risk`: Completed subagent work sat unsynthesized until the human asked about it. The human had to manage continuation semantics instead of trusting the leader to resume.
- `local_lesson_hypothesis`: Unless the human explicitly redirects scope, answer steering questions and then resume the active main-line task automatically.
- `cluster_hints`: `standing-instruction-persistence`, `interruption-semantics`, `main-line-continuation`
- `accepted_incident_likelihood`: `possible`
- `confidence`: `weak`
- `uncertainties`: This still needs the most caution. A later pass could reasonably decide it was only a conversational pause rather than a durable workflow miss.

### C07. IMPLEMENTATION-LEADER needed hard prerequisites for completed research and explicit plan approval

- `event_id`: `C07`
- `title`: IMPLEMENTATION-LEADER needed hard prerequisites for completed research and explicit plan approval
- `session_or_thread`: `019d3b3e-b474-75a2-a2d2-254c407cd900` (orchestration contract / Task-0003 execution thread)
- `transcript_path`: [rollout-2026-03-29T16-17-38-019d3b3e-b474-75a2-a2d2-254c407cd900.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T16-17-38-019d3b3e-b474-75a2-a2d2-254c407cd900.jsonl)
- `primary_refs`: [L302](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T16-17-38-019d3b3e-b474-75a2-a2d2-254c407cd900.jsonl#L302), [L307](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T16-17-38-019d3b3e-b474-75a2-a2d2-254c407cd900.jsonl#L307), [L312](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T16-17-38-019d3b3e-b474-75a2-a2d2-254c407cd900.jsonl#L312), [L333](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T16-17-38-019d3b3e-b474-75a2-a2d2-254c407cd900.jsonl#L333)
- `ai_course`: The assistant had just written `IMPLEMENTATION-LEADER.md` as ready-to-use, but the contract still left room to launch without a completed `RESEARCH.md` or to treat `PLAN.md` approval as softer than an explicit human gate.
- `human_intervention`: The human explicitly required both constraints to be hardened in the prompt before trusting it.
- `adequate_outcome`: `RESEARCH.md` is a hard prerequisite, and Stage B implementation cannot start without explicit human approval of the current `PLAN.md`.
- `event_boundary_notes`: Clean prompt-hardening event. It does not need to be merged with later task-state semantics work in `C08`.
- `human_model_signal`: Strong explicit signals: "`RESEARCH.md` must exist is a hard requirement" and "Clarify the human approval of `PLAN.md` is required before implementation begins."
- `failure_family_hypothesis`: `workflow_orchestration`
- `intervention_kind_hypothesis`: `tighten_contract`
- `human_cost_or_risk`: Implementation could start on incomplete research or inferred approval, forcing the human to police preconditions manually.
- `local_lesson_hypothesis`: Hard launch prerequisites and human gates should be encoded as explicit blockers, not implied expectations.
- `cluster_hints`: `hard-gate-clarity`, `prerequisite-truth`, `human-approval-gate`
- `accepted_incident_likelihood`: `possible`
- `confidence`: `strong`
- `uncertainties`: None material.

### C08. current_gate mixed semantic workflow state with commit/push/notify mechanics

- `event_id`: `C08`
- `title`: `current_gate` mixed semantic workflow state with `commit` / `push` / `notify` mechanics
- `session_or_thread`: `019d3b3e-b474-75a2-a2d2-254c407cd900` (orchestration contract / Task-0003 execution thread)
- `transcript_path`: [rollout-2026-03-29T16-17-38-019d3b3e-b474-75a2-a2d2-254c407cd900.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T16-17-38-019d3b3e-b474-75a2-a2d2-254c407cd900.jsonl)
- `primary_refs`: [L405](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T16-17-38-019d3b3e-b474-75a2-a2d2-254c407cd900.jsonl#L405), [L410](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T16-17-38-019d3b3e-b474-75a2-a2d2-254c407cd900.jsonl#L410), [L459](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T16-17-38-019d3b3e-b474-75a2-a2d2-254c407cd900.jsonl#L459), [L464](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T16-17-38-019d3b3e-b474-75a2-a2d2-254c407cd900.jsonl#L464), [L469](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T16-17-38-019d3b3e-b474-75a2-a2d2-254c407cd900.jsonl#L469), [L474](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T16-17-38-019d3b3e-b474-75a2-a2d2-254c407cd900.jsonl#L474), [L479](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T16-17-38-019d3b3e-b474-75a2-a2d2-254c407cd900.jsonl#L479), [L518](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T16-17-38-019d3b3e-b474-75a2-a2d2-254c407cd900.jsonl#L518)
- `ai_course`: The assistant translated a new checkpointing rule into leader behavior while the shared `current_gate` enum still included `commit`, `push`, and `notify`, which made the field partly semantic state and partly persistence bookkeeping.
- `human_intervention`: The human first wondered whether `current_gate` was basically persistence status, then sharpened that into the direct question: do `commit` / `push` / `notify` make sense as task state at all?
- `adequate_outcome`: `current_gate` should represent semantic work gates. Commit/push/notify should remain workflow mechanics or pass-closeout evidence, not canonical task state.
- `event_boundary_notes`: PASS1's window is right, but the decisive intervention is the later semantic cleanup at [L469](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T16-17-38-019d3b3e-b474-75a2-a2d2-254c407cd900.jsonl#L469) to [L479](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T16-17-38-019d3b3e-b474-75a2-a2d2-254c407cd900.jsonl#L479), not the initial checkpointing request by itself.
- `human_model_signal`: Strong explicit signal: "current_gate is really just the persistence status basically?" followed by "do commit / push / notify make sense as task state?"
- `failure_family_hypothesis`: `workflow_orchestration`
- `intervention_kind_hypothesis`: `reframe_problem`
- `human_cost_or_risk`: The task-state contract became mentally slippery and recursive, making automation and human reasoning over state less trustworthy.
- `local_lesson_hypothesis`: Canonical state fields should encode durable semantic state, not ephemeral mechanics for persisting that state.
- `cluster_hints`: `status-signal-meaning`, `state-vs-mechanics`, `schema-semantics`
- `accepted_incident_likelihood`: `possible`
- `confidence`: `medium`
- `uncertainties`: The confusion arose at the intersection of two reasonable ideas: durable checkpointing and semantic task-state tracking. The event is still real because the human had to untangle the boundary explicitly.

### C09. The standing instruction to finish remaining passes was dropped at the pass boundary

- `event_id`: `C09`
- `title`: The standing instruction to finish remaining passes was dropped at the pass boundary
- `session_or_thread`: `019d3b3e-b474-75a2-a2d2-254c407cd900` (orchestration contract / Task-0003 execution thread)
- `transcript_path`: [rollout-2026-03-29T16-17-38-019d3b3e-b474-75a2-a2d2-254c407cd900.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T16-17-38-019d3b3e-b474-75a2-a2d2-254c407cd900.jsonl)
- `primary_refs`: [L1642](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T16-17-38-019d3b3e-b474-75a2-a2d2-254c407cd900.jsonl#L1642), [L1647](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T16-17-38-019d3b3e-b474-75a2-a2d2-254c407cd900.jsonl#L1647), [L1652](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T16-17-38-019d3b3e-b474-75a2-a2d2-254c407cd900.jsonl#L1652), [L1665](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T16-17-38-019d3b3e-b474-75a2-a2d2-254c407cd900.jsonl#L1665), [L1670](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T16-17-38-019d3b3e-b474-75a2-a2d2-254c407cd900.jsonl#L1670), [L1675](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T16-17-38-019d3b3e-b474-75a2-a2d2-254c407cd900.jsonl#L1675)
- `ai_course`: After `PASS-0002`, the assistant treated the pass boundary as a natural stopping point and waited, even though the task still had `PASS-0003` left and the human had already asked to finish the remaining passes.
- `human_intervention`: The human reminded the assistant that he had already asked for the remaining passes to be finished, then pressed for the real cause instead of accepting the first explanation.
- `adequate_outcome`: The fresh-agent-per-pass rule should rotate the agent internally, not pause the overall task. A standing instruction to finish remaining passes should survive the handoff.
- `event_boundary_notes`: One event with two stages: first the human catches the premature stop, then he rejects a too-narrow explanation and forces the assistant to name the real miss.
- `human_model_signal`: Very strong explicit signal. "Even with the fresh impl rule (or not) the ask was to launch a fresh agent. So what was the real cause?"
- `failure_family_hypothesis`: `workflow_orchestration`
- `intervention_kind_hypothesis`: `reject_outcome`
- `human_cost_or_risk`: The human had to police pass boundaries and restate a higher-level standing instruction that should have persisted automatically.
- `local_lesson_hypothesis`: Boundary rotation rules are subordinate to standing umbrella instructions unless a real blocker or explicit human gate intervenes.
- `cluster_hints`: `standing-instruction-persistence`, `baton-drop`, `agent-rotation-continuity`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: None material.

### C10. Terminal subagent completion was not surfaced immediately to the human

- `event_id`: `C10`
- `title`: Terminal subagent completion was not surfaced immediately to the human
- `session_or_thread`: `019d3b3e-b474-75a2-a2d2-254c407cd900` (orchestration contract / Task-0003 execution thread)
- `transcript_path`: [rollout-2026-03-29T16-17-38-019d3b3e-b474-75a2-a2d2-254c407cd900.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T16-17-38-019d3b3e-b474-75a2-a2d2-254c407cd900.jsonl)
- `primary_refs`: [L1958](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T16-17-38-019d3b3e-b474-75a2-a2d2-254c407cd900.jsonl#L1958), [L1961](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T16-17-38-019d3b3e-b474-75a2-a2d2-254c407cd900.jsonl#L1961), [L1964](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T16-17-38-019d3b3e-b474-75a2-a2d2-254c407cd900.jsonl#L1964), [L1969](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T16-17-38-019d3b3e-b474-75a2-a2d2-254c407cd900.jsonl#L1969), [L1996](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T16-17-38-019d3b3e-b474-75a2-a2d2-254c407cd900.jsonl#L1996), [L2034](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T16-17-38-019d3b3e-b474-75a2-a2d2-254c407cd900.jsonl#L2034)
- `ai_course`: The assistant received the regression leader's completion notification but treated it as an internal signal to verify first instead of immediately telling the human that a terminal milestone had been reached.
- `human_intervention`: The human called out the repeat miss directly ("You didn't tell me the regression leader is done again"), then later agreed the rule should be made explicit in the shared orchestration docs and prompt stack.
- `adequate_outcome`: Any subagent terminal state should trigger an immediate user-visible update first; any extra verification belongs after that notification.
- `event_boundary_notes`: This is one staged correction arc. The originating intervention is the human's complaint about the missed notification; the later prompt/doc hardening is the durable repair.
- `human_model_signal`: Strong explicit signal. "You didn't tell me the regression leader is done again. Why?"
- `failure_family_hypothesis`: `workflow_orchestration`
- `intervention_kind_hypothesis`: `tighten_contract`
- `human_cost_or_risk`: The human has to ask redundant status questions to discover that work is already done. Real milestones become opaque instead of confidence-building.
- `local_lesson_hypothesis`: Terminal subagent completion is itself a user-facing milestone and should never stay internal while the supervisor does extra checks.
- `cluster_hints`: `terminal-state-reporting`, `status-signal-meaning`, `human-notification-timing`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: None material.

### C11. Task 0004 was closed on a supporting proof lane instead of the human-facing regression lane

- `event_id`: `C11`
- `title`: Task 0004 was closed on a supporting proof lane instead of the human-facing regression lane
- `session_or_thread`: `019d3ce5-9eef-7c61-922f-bb45f3f5a5d4` (`Task-0004` regression / precedence / debugging thread)
- `transcript_path`: [rollout-2026-03-29T23-59-34-019d3ce5-9eef-7c61-922f-bb45f3f5a5d4.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T23-59-34-019d3ce5-9eef-7c61-922f-bb45f3f5a5d4.jsonl)
- `primary_refs`: [L354](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T23-59-34-019d3ce5-9eef-7c61-922f-bb45f3f5a5d4.jsonl#L354), [L359](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T23-59-34-019d3ce5-9eef-7c61-922f-bb45f3f5a5d4.jsonl#L359), [L400](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T23-59-34-019d3ce5-9eef-7c61-922f-bb45f3f5a5d4.jsonl#L400)
- `ai_course`: The assistant declared Task 4 complete based on host-side full-stack automation plus packaged worker proof, while openly admitting that the preferred emulator-visible lane was still blocked.
- `human_intervention`: The human flatly rejected that closure call and restated what regression meant to him: exercise the path the human cares about, not a supporting surrogate lane.
- `adequate_outcome`: The task remains open until the required human-facing regression lane passes. Supporting lanes may inform diagnosis but do not answer the closure question.
- `event_boundary_notes`: Distinct from `C12`. This event is the rejection of the claimed successful outcome itself.
- `human_model_signal`: Very strong explicit signal. "It won't be working when the human uses it" and "regression - exercise the paths the human cares about."
- `failure_family_hypothesis`: `verification_proof`
- `intervention_kind_hypothesis`: `reject_outcome`
- `human_cost_or_risk`: A task can be marked complete on surrogate evidence even though the actual user-facing lane is still broken. The human then has to reopen supposedly closed work.
- `local_lesson_hypothesis`: Supporting proof cannot satisfy a regression ask or closure claim when the human-facing lane remains blocked or failing.
- `cluster_hints`: `regression-means-human-lane`, `supporting-proof-substitution`, `premature-closure`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: None material.

### C12. Regression meaning drifted because task-local docs were allowed to outrun REGRESSION.md and shared precedence rules

- `event_id`: `C12`
- `title`: Regression meaning drifted because task-local docs were allowed to outrun `REGRESSION.md` and shared precedence rules
- `session_or_thread`: `019d3ce5-9eef-7c61-922f-bb45f3f5a5d4` (`Task-0004` regression / precedence / debugging thread)
- `transcript_path`: [rollout-2026-03-29T23-59-34-019d3ce5-9eef-7c61-922f-bb45f3f5a5d4.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T23-59-34-019d3ce5-9eef-7c61-922f-bb45f3f5a5d4.jsonl)
- `primary_refs`: [L405](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T23-59-34-019d3ce5-9eef-7c61-922f-bb45f3f5a5d4.jsonl#L405), [L410](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T23-59-34-019d3ce5-9eef-7c61-922f-bb45f3f5a5d4.jsonl#L410), [L415](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T23-59-34-019d3ce5-9eef-7c61-922f-bb45f3f5a5d4.jsonl#L415), [L420](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T23-59-34-019d3ce5-9eef-7c61-922f-bb45f3f5a5d4.jsonl#L420), [L425](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T23-59-34-019d3ce5-9eef-7c61-922f-bb45f3f5a5d4.jsonl#L425), [L430](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T23-59-34-019d3ce5-9eef-7c61-922f-bb45f3f5a5d4.jsonl#L430), [L435](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T23-59-34-019d3ce5-9eef-7c61-922f-bb45f3f5a5d4.jsonl#L435), [L440](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T23-59-34-019d3ce5-9eef-7c61-922f-bb45f3f5a5d4.jsonl#L440)
- `ai_course`: The assistant's closure reasoning depended on task-local documents and ad hoc lane terminology, which left room for a hybrid closure story outside repo-root `REGRESSION.md`.
- `human_intervention`: The human pushed for a cleaner rule: regression is the lane the human cares about, it is defined only in `REGRESSION.md`, and task-local artifacts must not redefine it. He then pushed for explicit precedence rules to encode that authority structure.
- `adequate_outcome`: Shared `.codex` docs are authoritative for workflow/process, repo-root `REGRESSION.md` is authoritative for regression lanes and pass criteria, and task-local artifacts may reference required cases but may not redefine or substitute them.
- `event_boundary_notes`: Follow-on to `C11`, but distinct. `C11` corrects the wrong closure outcome. `C12` explains and hardens the documentation authority model that should have prevented it.
- `human_model_signal`: Very strong explicit signal. The human asks whether "the regression lane is defined as the lane the human cares about and is documented only in REGRESSION.md," then pushes on precedence rules for canonical docs.
- `failure_family_hypothesis`: `workflow_orchestration`
- `intervention_kind_hypothesis`: `tighten_contract`
- `human_cost_or_risk`: Without an authority rule, task-local docs can silently relax closure bars and leads have no hard standard for spotting drift.
- `local_lesson_hypothesis`: Encode canonical source authority by domain, then require task-local artifacts and leads to conform to it rather than improvise local closure semantics.
- `cluster_hints`: `doc-precedence`, `regression-lane-authority`, `canonical-source-discipline`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: The human briefly floated stronger blanket `.codex` supremacy language, but the durable accepted rule in-thread was the domain-based precedence model, not an absolute "all .codex beats all repo docs" rule.

### C13. Manual CA trust install was treated as a blocker instead of acceptable one-time development friction

- `event_id`: `C13`
- `title`: Manual CA trust install was treated as a blocker instead of acceptable one-time development friction
- `session_or_thread`: `019d3ce5-9eef-7c61-922f-bb45f3f5a5d4` (`Task-0004` regression / precedence / debugging thread)
- `transcript_path`: [rollout-2026-03-29T23-59-34-019d3ce5-9eef-7c61-922f-bb45f3f5a5d4.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T23-59-34-019d3ce5-9eef-7c61-922f-bb45f3f5a5d4.jsonl)
- `primary_refs`: [L893](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T23-59-34-019d3ce5-9eef-7c61-922f-bb45f3f5a5d4.jsonl#L893), [L898](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T23-59-34-019d3ce5-9eef-7c61-922f-bb45f3f5a5d4.jsonl#L898), [L907](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T23-59-34-019d3ce5-9eef-7c61-922f-bb45f3f5a5d4.jsonl#L907), [L912](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T23-59-34-019d3ce5-9eef-7c61-922f-bb45f3f5a5d4.jsonl#L912), [L917](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T23-59-34-019d3ce5-9eef-7c61-922f-bb45f3f5a5d4.jsonl#L917), [L922](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T23-59-34-019d3ce5-9eef-7c61-922f-bb45f3f5a5d4.jsonl#L922), [L948](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T23-59-34-019d3ce5-9eef-7c61-922f-bb45f3f5a5d4.jsonl#L948)
- `ai_course`: The assistant treated the CA trust-store step as a policy blocker because the current docs described a manual Settings flow and the lead chose not to improvise another operator story, even though Task 0003 had already demonstrated a rooted-emulator CA staging path.
- `human_intervention`: The human first asked for the manual Settings flow in plain language, then tested whether that was something normal users should do, and finally clarified the real intent: a one-time manual CA flow is acceptable for development, so the task lead should align to that and retry regression.
- `adequate_outcome`: Development regression can use the acceptable CA-trust setup path and then continue answering the real required-lane regression question instead of stopping at the policy interpretation layer.
- `event_boundary_notes`: Distinct from `C14`. This event changes the interpretation of the CA step from blocker to acceptable setup cost.
- `human_model_signal`: Strong explicit signal. "That's fine I can acccept a one-time manual flow for development."
- `failure_family_hypothesis`: `human_world`
- `intervention_kind_hypothesis`: `reframe_problem`
- `human_cost_or_risk`: The agent stalled on an overly strict interpretation of acceptable operator burden, leaving the real regression question unanswered until the human clarified context.
- `local_lesson_hypothesis`: Judge setup friction against the actual development or home-lab context before treating it as a blocker. Do not import generic consumer-app expectations into a development-only lane without checking human intent.
- `cluster_hints`: `operator-burden-calibration`, `development-vs-product-context`, `acceptable-manual-flow`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: This event establishes acceptability of the one-time CA flow in development. It does not, by itself, settle the long-term preferred operator story for general users.

### C14. Required-lane regression failure still needed an explicit default route into BUG tracking and debugging

- `event_id`: `C14`
- `title`: Required-lane regression failure still needed an explicit default route into `BUG` tracking and debugging
- `session_or_thread`: `019d3ce5-9eef-7c61-922f-bb45f3f5a5d4` (`Task-0004` regression / precedence / debugging thread)
- `transcript_path`: [rollout-2026-03-29T23-59-34-019d3ce5-9eef-7c61-922f-bb45f3f5a5d4.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T23-59-34-019d3ce5-9eef-7c61-922f-bb45f3f5a5d4.jsonl)
- `primary_refs`: [L1126](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T23-59-34-019d3ce5-9eef-7c61-922f-bb45f3f5a5d4.jsonl#L1126), [L1131](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T23-59-34-019d3ce5-9eef-7c61-922f-bb45f3f5a5d4.jsonl#L1131), [L1136](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T23-59-34-019d3ce5-9eef-7c61-922f-bb45f3f5a5d4.jsonl#L1136), [L1160](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T23-59-34-019d3ce5-9eef-7c61-922f-bb45f3f5a5d4.jsonl#L1160), [L1205](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T23-59-34-019d3ce5-9eef-7c61-922f-bb45f3f5a5d4.jsonl#L1205)
- `ai_course`: In this concrete rerun, the assistant had already opened `BUG-0001.md`, updated `TASK-STATE.json` to debugging, and described the post-CA regression failure honestly. But the shared task-leader and regression-leader flow still lacked a hard rule making that route mandatory whenever a required regression lane fails or is blocked.
- `human_intervention`: The human explicitly asked to harden the prompt flow so any failed regression documents the problem in a `BUG` artifact and then proceeds into a debug workflow, even if the near-term resolution is "need an upstream human."
- `adequate_outcome`: Required-lane regression failure or blockage automatically creates or updates a task-owned `BUG-<NNNN>.md` and routes into debugging instead of living only in a regression-run note.
- `event_boundary_notes`: This is the clearest PASS1 correction. The local event is not "the run stopped with no bug note at all" because the transcript shows the assistant had already created one in this task. The actual intervention is the human noticing that the workflow default still depended on discretionary judgment and hardening it into a durable rule.
- `human_model_signal`: Strong explicit signal. "When regression fails for whatever reason it should document that in a BUG ... Then it should proceed to some debug workflow."
- `failure_family_hypothesis`: `workflow_orchestration`
- `intervention_kind_hypothesis`: `redirect_debugging`
- `human_cost_or_risk`: Without a hardened default, future required-lane failures can stall in bookkeeping or ambiguous pause states unless the human manually insists on defect tracking and debug routing.
- `local_lesson_hypothesis`: Required-lane regression failure should automatically become durable defect tracking plus debug workflow entry, not an optional follow-on judgment.
- `cluster_hints`: `regression-failure-routing`, `bug-before-pause`, `required-lane-failure-discipline`
- `accepted_incident_likelihood`: `likely`
- `confidence`: `strong`
- `uncertainties`: The event is about default workflow hardening, not a missing bug note in this exact rerun.

## Likely Accepted Incidents

- `C01` because the human had to reject a completed handoff artifact that overconstrained the stronger external model's thinking.
- `C03` because the briefing workflow overfit to evidence-packet size and produced unreadable repo dumps until the philosophy changed.
- `C09` because a standing umbrella instruction did not survive pass-boundary rotation and the assistant stopped one step too early.
- `C10` because terminal subagent completion was not surfaced as an immediate user-visible milestone until the human corrected it.
- `C11` because Task 0004 was explicitly closed on supporting proof even though the required human-facing regression lane was still not working.
- `C12` because regression authority and documentation precedence were weak enough to let task-local artifacts drift into a hybrid closure story.
- `C13` because acceptable development/operator friction was misread as a blocker, preventing the real regression question from being answered.
- `C14` because required-lane regression failure still needed human instruction to become a hardened default route into bug tracking and debugging.

## Likely Non-Incident But Still Important Intervention Events

- `C02` real architecture-seam correction, but narrower and more prompt-design-local than the stronger accepted candidates.
- `C04` important reality check about downstream model capabilities, but likely better understood as a contract hardening than a durable accepted incident.
- `C05` important exemplar-binding and boundedness fix for `RESEARCH-LEADER`, but still more prompt-scaffolding than corpus-defining incident.
- `C06` real continuation concern, but still the softest boundary in the set.
- `C07` direct prompt hardening around prerequisites and approval gates; important, but more like contract tightening than accepted incident.
- `C08` useful schema semantics cleanup; durable, but lower human-cost intensity than the strongest events.

## Repeated Cluster Hints Across The Set

- `external-handoff-shaping`: `C01`, `C03`, `C04`, `C05`
- `role-boundary-drift`: `C02`, `C05`, `C07`, `C08`, `C14`
- `standing-instruction-persistence`: `C06`, `C09`
- `terminal-state-reporting`: `C10`
- `status-signal-meaning`: `C08`, `C10`, `C12`
- `regression-means-human-lane`: `C11`, `C12`, `C13`, `C14`
- `real-world-completion-mismatch`: `C04`, `C11`, `C13`
- `required-lane-failure-routing`: `C11`, `C14`

## Strongest Human-Model Signals To Carry Into Later Clustering

- External research handoffs should give "facts on the ground" and at most `3 MAX` open-ended questions. The human wants genuine outside thinking, not compliance with a local worksheet.
- If a prompt may create a durable artifact, it needs explicit shape authority. Naming `RESEARCH-PLAN.md` is not enough; the leader must be told to follow the exemplar.
- A standing instruction like "finish the remaining passes" survives internal agent rotation. Pass boundaries are internal relays unless a real blocker or human gate appears.
- Terminal subagent completion is itself a user-visible milestone. The human should hear about it immediately, before extra verification work.
- Regression means the lane the human cares about. Supporting proof can narrow diagnosis, but it cannot answer the regression question or close the task.
- `REGRESSION.md` should be the authoritative definition of regression lanes and pass criteria in a repo, with task-local artifacts only naming relevant cases and reporting current status.
- Development/operator burden is contextual. A one-time manual CA install can be acceptable in development even if it would be wrong as a general-user expectation.
- When the required regression lane fails or is blocked, the workflow should not stop at a run note. It should create or update a bug artifact and route into debugging by default.

## Events That Still Need A Wider Reread

- `C06` still has the weakest boundary and would benefit from a wider reread if a later pass wants to decide whether it belongs in the accepted set or should stay as a softer continuation-semantics lesson.
- `C03` may need a slightly wider reread immediately before [L913](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl#L913) if a later accepted-incident writer wants stronger human-owned dissatisfaction wording rather than relying mainly on the assistant's diagnosis plus the human's authorization to pivot.
