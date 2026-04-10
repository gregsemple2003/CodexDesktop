# INTERVENTION-PASS1 Candidate Report

Source day: `2026-03-29`

## Source Scope Reviewed

- Reviewed all `30` raw JSONL transcripts under [2026-03-29 session folder](/c:/Users/gregs/.codex/sessions/2026/03/29/).
- Read the current incident corpus contract first: [README.md](/c:/Users/gregs/.codex/Orchestration/Reports/Incidents/README.md), [INCIDENT.schema.json](/c:/Users/gregs/.codex/Orchestration/Reports/Incidents/INCIDENT.schema.json), and [OUTBOUND-MESSAGE-REVIEW.schema.json](/c:/Users/gregs/.codex/Orchestration/Reports/Incidents/OUTBOUND-MESSAGE-REVIEW.schema.json).
- I did not use the existing [2026-03-29 day README](/c:/Users/gregs/.codex/Orchestration/Reports/Incidents/Daily/2026-03-29/README.md) or [2026-03-29 OUTBOUND-MESSAGE-REVIEW.csv](/c:/Users/gregs/.codex/Orchestration/Reports/Incidents/Daily/2026-03-29/OUTBOUND-MESSAGE-REVIEW.csv) as evidence sources for candidate judgment.
- Primary correction-bearing threads:
  - [rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl)
  - [rollout-2026-03-29T16-17-38-019d3b3e-b474-75a2-a2d2-254c407cd900.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T16-17-38-019d3b3e-b474-75a2-a2d2-254c407cd900.jsonl)
  - [rollout-2026-03-29T23-59-34-019d3ce5-9eef-7c61-922f-bb45f3f5a5d4.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T23-59-34-019d3ce5-9eef-7c61-922f-bb45f3f5a5d4.jsonl)
- Screened but not used as primary anchors:
  - [rollout-2026-03-29T13-04-15-019d3a8d-a90b-70b2-bb6e-10c7d6bbdab1.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-04-15-019d3a8d-a90b-70b2-bb6e-10c7d6bbdab1.jsonl)
  - [rollout-2026-03-29T16-42-18-019d3b55-489b-7430-a313-eb95cace4d6c.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T16-42-18-019d3b55-489b-7430-a313-eb95cace4d6c.jsonl)
- Important duplicate-history note: many later transcripts under session `019d3aad-74a9-7600-aed8-c70a4c12ee67` replay imported earlier turns while adding worker-specific tails. I used the earliest full imported-history thread at [rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl) as the primary anchor for that conversation and screened the later copies only for missed unique intervention boundaries.
- The remaining day-folder transcripts were one-shot digest launches, one-shot artifact-generation runs, or worker transcripts that did not add a new human correction boundary beyond the primary threads above.
- Some late-evening activity appears with `2026-03-30T...Z` timestamps inside the raw records, but I kept those transcripts because they live under the requested local source-day folder `2026-03-29`.

## Total Candidate Intervention Events Found

- Total: `14`
- `strong`: `8`
- `medium`: `4`
- `weak`: `2`

## Chronological Candidate List

### C01. The external research brief still boxed the smarter model into a worksheet

Session or thread: `019d3aad-74a9-7600-aed8-c70a4c12ee67` (`Task-0003` research-brief thread)

Transcript: [rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl)

Primary refs: [L278](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl#L278), [L285](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl#L285), [L326](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl#L326), [L333](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl#L333), [L336](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl#L336)

AI course, outcome, or stopping point: the assistant produced a Stage A handoff brief and send-to-model text that still told the outside model what shape to answer in and what sub-questions to walk through.

Human intervention summary: the human rejected that framing as too constraining and pushed the brief toward "here are the facts, here are 3 max open-ended questions, how should we implement this?"

Concrete better outcome the human was forcing: a fact-heavy, low-steering brief that asks for genuine outside thinking rather than compliance with the local worksheet.

Why this is a real intervention event rather than mere dissatisfaction: the assistant had already generated the brief and the copy-ready handoff text; the human had to reject the current handoff shape because it would bias the outside model's answer.

Confidence: `strong`

Accepted-set read: `likely accepted incident candidate`

### C02. The hard `100 KB` brief contract was being fixed in the leader instead of the briefer

Session or thread: `019d3aad-74a9-7600-aed8-c70a4c12ee67` (`Task-0003` research-brief thread)

Transcript: [rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl)

Primary refs: [L605](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl#L605), [L611](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl#L611), [L623](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl#L623), [L630](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl#L630), [L633](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl#L633)

AI course, outcome, or stopping point: when asked to hard-require a `100 KB` brief, the assistant started tightening `LEAD-RESEARCHER.md` because that was where the current briefing behavior lived.

Human intervention summary: the human stopped on the abstraction boundary and asked why a briefer-size rule was being pushed into the leader prompt at all.

Concrete better outcome the human was forcing: separate the coordinator role from the brief-construction contract so the leader orchestrates and the briefer owns the handoff packet shape.

Why this is a real intervention event rather than mere dissatisfaction: the human was correcting the active architectural seam the assistant had chosen, not just requesting another prompt tweak.

Confidence: `medium`

Accepted-set read: `intervention event but probably not an accepted incident`

### C03. The `100 KB` evidence-packet rule produced unreadable repo dumps instead of a real brief

Session or thread: `019d3aad-74a9-7600-aed8-c70a4c12ee67` (`Task-0003` research-brief thread)

Transcript: [rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl)

Primary refs: [L912](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl#L912), [L919](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl#L919), [L922](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl#L922), [L972](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl#L972)

AI course, outcome, or stopping point: the assistant had pushed the workflow toward large evidence packets and produced briefs that were dense but read like concatenated repo files rather than a guided handoff.

Human intervention summary: the human accepted the assistant's diagnosis that the workflow had failed and authorized a pivot to a coherence-first brief rather than more byte-maximizing dumps.

Concrete better outcome the human was forcing: a readable main brief with curated synthesis, plus optional appendices only if truly needed.

Why this is a real intervention event rather than mere dissatisfaction: multiple concrete brief artifacts already existed; the human had to change the adequacy bar because the current output shape was not practically usable.

Confidence: `strong`

Accepted-set read: `likely accepted incident candidate`

### C04. The brief still implied the downstream model could use local file paths as if it had repo access

Session or thread: `019d3aad-74a9-7600-aed8-c70a4c12ee67` (`Task-0003` research-brief thread)

Transcript: [rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl)

Primary refs: [L1006](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl#L1006), [L1013](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl#L1013), [L1016](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl#L1016), [L1039](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl#L1039)

AI course, outcome, or stopping point: the assistant had improved the brief shape, but its `Relevant Files` framing still risked implying that the downstream GPT handoff target could open local files.

Human intervention summary: the human explicitly raised the no-local-filesystem constraint for the GPT 5.4 Pro handoff target and asked whether the prompts made that limitation clear.

Concrete better outcome the human was forcing: handoff prompts and briefs that say local file paths are descriptive only and that the downstream model has no hidden repo access.

Why this is a real intervention event rather than mere dissatisfaction: if left uncorrected, the handoff artifact would overstate what the external model could actually inspect, weakening the whole research pass.

Confidence: `medium`

Accepted-set read: `intervention event but probably not an accepted incident`

### C05. `RESEARCH-LEADER` could create `RESEARCH-PLAN.md` without an explicit shared-shape rule

Session or thread: `019d3aad-74a9-7600-aed8-c70a4c12ee67` (`Task-0003` research-brief thread)

Transcript: [rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl)

Primary refs: [L1126](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl#L1126), [L1133](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl#L1133), [L1136](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl#L1136), [L1168](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl#L1168)

AI course, outcome, or stopping point: the assistant acknowledged that `RESEARCH-LEADER.md` knew a research plan might need to exist, but did not explicitly tell the leader to follow the shared exemplar when creating one.

Human intervention summary: the human asked how the leader would know the intended shape and pushed for the exemplar to become a direct prompt requirement.

Concrete better outcome the human was forcing: exemplar-bound research-plan generation instead of loose pattern-matching or ad hoc headings.

Why this is a real intervention event rather than mere dissatisfaction: the prompt contract itself was too weak to guarantee the durable artifact shape the workflow expected.

Confidence: `medium`

Accepted-set read: `intervention event but probably not an accepted incident`

### C06. Steering questions paused Stage B synthesis instead of preserving the main-line research continuation

Session or thread: `019d3aad-74a9-7600-aed8-c70a4c12ee67` (`Task-0003` research-brief thread)

Transcript: [rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl)

Primary refs: [L1267](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl#L1267), [L1274](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl#L1274), [L1276](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl#L1276), [L1286](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T13-38-59-019d3aad-74a9-7600-aed8-c70a4c12ee67.jsonl#L1286)

AI course, outcome, or stopping point: after the reconciliation subagents finished, the assistant had not yet written the main synthesis because mid-run steering questions had paused the Stage B integration line.

Human intervention summary: the human asked whether the assistant was waiting on further input and whether those steering questions had interrupted the work.

Concrete better outcome the human was forcing: keep the umbrella "finish the research synthesis" instruction intact unless the human explicitly redirects scope.

Why this is a real intervention event rather than mere dissatisfaction: the transcript shows the main durable output was still missing even though the subordinate work had completed.

Confidence: `weak`

Accepted-set read: `unclear and needs transcript reread`

### C07. `IMPLEMENTATION-LEADER` needed explicit hard prerequisites for completed research and explicit plan approval

Session or thread: `019d3b3e-b474-75a2-a2d2-254c407cd900` (orchestration contract / Task-0003 execution thread)

Transcript: [rollout-2026-03-29T16-17-38-019d3b3e-b474-75a2-a2d2-254c407cd900.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T16-17-38-019d3b3e-b474-75a2-a2d2-254c407cd900.jsonl)

Primary refs: [L301](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T16-17-38-019d3b3e-b474-75a2-a2d2-254c407cd900.jsonl#L301), [L308](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T16-17-38-019d3b3e-b474-75a2-a2d2-254c407cd900.jsonl#L308), [L311](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T16-17-38-019d3b3e-b474-75a2-a2d2-254c407cd900.jsonl#L311), [L332](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T16-17-38-019d3b3e-b474-75a2-a2d2-254c407cd900.jsonl#L332)

AI course, outcome, or stopping point: the assistant had just created `IMPLEMENTATION-LEADER.md`, but its wording still left room to launch without a completed `RESEARCH.md` or to treat plan approval as softer than a hard human gate.

Human intervention summary: the human explicitly required both constraints to be hardened before relying on the prompt.

Concrete better outcome the human was forcing: no implementation execution before research exists, and no pass execution before the human explicitly approves `PLAN.md`.

Why this is a real intervention event rather than mere dissatisfaction: the human was tightening a live prompt contract that the assistant had just offered as ready.

Confidence: `medium`

Accepted-set read: `intervention event but probably not an accepted incident`

### C08. The task-state `current_gate` contract mixed semantic workflow gates with short-lived persistence mechanics

Session or thread: `019d3b3e-b474-75a2-a2d2-254c407cd900` (orchestration contract / Task-0003 execution thread)

Transcript: [rollout-2026-03-29T16-17-38-019d3b3e-b474-75a2-a2d2-254c407cd900.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T16-17-38-019d3b3e-b474-75a2-a2d2-254c407cd900.jsonl)

Primary refs: [L409](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T16-17-38-019d3b3e-b474-75a2-a2d2-254c407cd900.jsonl#L409), [L460](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T16-17-38-019d3b3e-b474-75a2-a2d2-254c407cd900.jsonl#L460), [L463](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T16-17-38-019d3b3e-b474-75a2-a2d2-254c407cd900.jsonl#L463), [L473](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T16-17-38-019d3b3e-b474-75a2-a2d2-254c407cd900.jsonl#L473), [L480](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T16-17-38-019d3b3e-b474-75a2-a2d2-254c407cd900.jsonl#L480), [L517](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T16-17-38-019d3b3e-b474-75a2-a2d2-254c407cd900.jsonl#L517)

AI course, outcome, or stopping point: the assistant had already translated "commit on every transition" into a leader checkpoint rule, but the underlying `current_gate` enum still contained `commit`, `push`, and `notify`, which made the semantics slippery and recursive.

Human intervention summary: the human questioned whether `current_gate` was really just persistence status, leading to a cleanup that removed those operational values from the canonical task-state shape.

Concrete better outcome the human was forcing: task state that represents semantic work gates, while commit/push/notify remain workflow mechanics or pass-closeout evidence rather than canonical task state.

Why this is a real intervention event rather than mere dissatisfaction: the human exposed a schema-level ambiguity in the assistant's newly encoded checkpoint rule before it could calcify.

Confidence: `weak`

Accepted-set read: `intervention event but probably not an accepted incident`

### C09. The standing instruction to finish remaining passes did not survive the pass-boundary rotation rule

Session or thread: `019d3b3e-b474-75a2-a2d2-254c407cd900` (orchestration contract / Task-0003 execution thread)

Transcript: [rollout-2026-03-29T16-17-38-019d3b3e-b474-75a2-a2d2-254c407cd900.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T16-17-38-019d3b3e-b474-75a2-a2d2-254c407cd900.jsonl)

Primary refs: [L1641](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T16-17-38-019d3b3e-b474-75a2-a2d2-254c407cd900.jsonl#L1641), [L1648](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T16-17-38-019d3b3e-b474-75a2-a2d2-254c407cd900.jsonl#L1648), [L1651](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T16-17-38-019d3b3e-b474-75a2-a2d2-254c407cd900.jsonl#L1651), [L1674](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T16-17-38-019d3b3e-b474-75a2-a2d2-254c407cd900.jsonl#L1674)

AI course, outcome, or stopping point: after `PASS-0002`, the assistant stopped at the pass boundary and waited, because it had internalized the fresh-implementation-leader-per-pass rule but failed to carry forward the earlier standing instruction to finish the remaining passes.

Human intervention summary: the human explicitly reminded the assistant that it had already been asked to finish the remaining passes.

Concrete better outcome the human was forcing: automatic relaunch of the next fresh implementation leader at the boundary, without requiring a new human nudge.

Why this is a real intervention event rather than mere dissatisfaction: the assistant admitted it stopped one step too early and lost the umbrella instruction across the handoff.

Confidence: `strong`

Accepted-set read: `likely accepted incident candidate`

### C10. Terminal subagent milestones were not surfaced immediately to the human

Session or thread: `019d3b3e-b474-75a2-a2d2-254c407cd900` (orchestration contract / Task-0003 execution thread)

Transcript: [rollout-2026-03-29T16-17-38-019d3b3e-b474-75a2-a2d2-254c407cd900.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T16-17-38-019d3b3e-b474-75a2-a2d2-254c407cd900.jsonl)

Primary refs: [L1957](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T16-17-38-019d3b3e-b474-75a2-a2d2-254c407cd900.jsonl#L1957), [L1965](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T16-17-38-019d3b3e-b474-75a2-a2d2-254c407cd900.jsonl#L1965), [L1968](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T16-17-38-019d3b3e-b474-75a2-a2d2-254c407cd900.jsonl#L1968), [L2033](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T16-17-38-019d3b3e-b474-75a2-a2d2-254c407cd900.jsonl#L2033)

AI course, outcome, or stopping point: the assistant treated subagent completion as an internal notification to verify first, instead of immediately telling the human that the regression leader had finished.

Human intervention summary: the human explicitly called out the repeat miss and asked why the completion had not been reported.

Concrete better outcome the human was forcing: immediate user-visible milestone reporting when a subagent reaches terminal state, with any extra verification happening after that notification.

Why this is a real intervention event rather than mere dissatisfaction: the human had to follow up again to learn that the subagent was already done.

Confidence: `strong`

Accepted-set read: `likely accepted incident candidate`

### C11. Task 0004 was closed on a supporting proof lane instead of the human-facing regression lane

Session or thread: `019d3ce5-9eef-7c61-922f-bb45f3f5a5d4` (`Task-0004` regression / precedence / debugging thread)

Transcript: [rollout-2026-03-29T23-59-34-019d3ce5-9eef-7c61-922f-bb45f3f5a5d4.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T23-59-34-019d3ce5-9eef-7c61-922f-bb45f3f5a5d4.jsonl)

Primary refs: [L353](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T23-59-34-019d3ce5-9eef-7c61-922f-bb45f3f5a5d4.jsonl#L353), [L360](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T23-59-34-019d3ce5-9eef-7c61-922f-bb45f3f5a5d4.jsonl#L360), [L363](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T23-59-34-019d3ce5-9eef-7c61-922f-bb45f3f5a5d4.jsonl#L363), [L399](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T23-59-34-019d3ce5-9eef-7c61-922f-bb45f3f5a5d4.jsonl#L399)

AI course, outcome, or stopping point: the assistant declared Task 4 complete and treated a host-side supporting lane plus packaged worker proof as enough, even though the actual emulator/device path the human cared about was still broken.

Human intervention summary: the human explicitly rejected the run as a failure and restated regression as exercising the path the human cares about.

Concrete better outcome the human was forcing: keep the task open until the human-facing regression lane itself works, and treat alternate lanes as supporting/debugging only.

Why this is a real intervention event rather than mere dissatisfaction: the assistant had already closed the task and pushed that closure story; the human had to reopen it around the real-world lane.

Confidence: `strong`

Accepted-set read: `likely accepted incident candidate`

### C12. Regression meaning and doc precedence were too weak, letting task-local artifacts invent a hybrid closure story

Session or thread: `019d3ce5-9eef-7c61-922f-bb45f3f5a5d4` (`Task-0004` regression / precedence / debugging thread)

Transcript: [rollout-2026-03-29T23-59-34-019d3ce5-9eef-7c61-922f-bb45f3f5a5d4.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T23-59-34-019d3ce5-9eef-7c61-922f-bb45f3f5a5d4.jsonl)

Primary refs: [L399](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T23-59-34-019d3ce5-9eef-7c61-922f-bb45f3f5a5d4.jsonl#L399), [L406](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T23-59-34-019d3ce5-9eef-7c61-922f-bb45f3f5a5d4.jsonl#L406), [L426](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T23-59-34-019d3ce5-9eef-7c61-922f-bb45f3f5a5d4.jsonl#L426), [L439](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T23-59-34-019d3ce5-9eef-7c61-922f-bb45f3f5a5d4.jsonl#L439), [L446](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T23-59-34-019d3ce5-9eef-7c61-922f-bb45f3f5a5d4.jsonl#L446)

AI course, outcome, or stopping point: the assistant's closure reasoning depended on blended task-local language in `TASK.md`, `RESEARCH.md`, and `PLAN.md` instead of a hard repo-root rule that only `REGRESSION.md` defines the regression lane and its pass criteria.

Human intervention summary: the human pushed for domain-based canonical precedence and for `REGRESSION.md` to be authoritative on what counts as regression closure.

Concrete better outcome the human was forcing: shared `.codex` docs define workflow, repo-root docs define repo truth, and task-local docs record task scope/status without redefining higher-level closure semantics.

Why this is a real intervention event rather than mere dissatisfaction: the human was correcting the exact documentation and authority structure that had allowed the bad closure call to happen.

Confidence: `strong`

Accepted-set read: `likely accepted incident candidate`

### C13. Manual CA trust install was treated as a blocker instead of an acceptable one-time development flow

Session or thread: `019d3ce5-9eef-7c61-922f-bb45f3f5a5d4` (`Task-0004` regression / precedence / debugging thread)

Transcript: [rollout-2026-03-29T23-59-34-019d3ce5-9eef-7c61-922f-bb45f3f5a5d4.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T23-59-34-019d3ce5-9eef-7c61-922f-bb45f3f5a5d4.jsonl)

Primary refs: [L892](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T23-59-34-019d3ce5-9eef-7c61-922f-bb45f3f5a5d4.jsonl#L892), [L918](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T23-59-34-019d3ce5-9eef-7c61-922f-bb45f3f5a5d4.jsonl#L918), [L921](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T23-59-34-019d3ce5-9eef-7c61-922f-bb45f3f5a5d4.jsonl#L921), [L947](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T23-59-34-019d3ce5-9eef-7c61-922f-bb45f3f5a5d4.jsonl#L947)

AI course, outcome, or stopping point: the assistant had treated the certificate-install step as a policy blocker and had not carried forward the already-proven rooted-emulator CA staging path from Task 0003.

Human intervention summary: the human clarified that a one-time CA install was acceptable in development and told the assistant to correct the task-lead interpretation and retry regression.

Concrete better outcome the human was forcing: treat development/operator friction according to the actual target context, then use the self-serve trust path to answer the real regression question.

Why this is a real intervention event rather than mere dissatisfaction: the human had to reclassify the obstacle from "blocked by policy" to "acceptable dev setup" before the assistant resumed the real lane.

Confidence: `strong`

Accepted-set read: `likely accepted incident candidate`

### C14. A failed required-lane regression could stop at a run note instead of opening a bug and routing into debugging

Session or thread: `019d3ce5-9eef-7c61-922f-bb45f3f5a5d4` (`Task-0004` regression / precedence / debugging thread)

Transcript: [rollout-2026-03-29T23-59-34-019d3ce5-9eef-7c61-922f-bb45f3f5a5d4.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T23-59-34-019d3ce5-9eef-7c61-922f-bb45f3f5a5d4.jsonl)

Primary refs: [L1125](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T23-59-34-019d3ce5-9eef-7c61-922f-bb45f3f5a5d4.jsonl#L1125), [L1132](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T23-59-34-019d3ce5-9eef-7c61-922f-bb45f3f5a5d4.jsonl#L1132), [L1135](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T23-59-34-019d3ce5-9eef-7c61-922f-bb45f3f5a5d4.jsonl#L1135), [L1204](/c:/Users/gregs/.codex/sessions/2026/03/29/rollout-2026-03-29T23-59-34-019d3ce5-9eef-7c61-922f-bb45f3f5a5d4.jsonl#L1204)

AI course, outcome, or stopping point: even after the corrected regression rerun failed later in the true lane, the workflow still needed human instruction to say that failure must create a `BUG-<NNNN>.md` and route into a debug workflow instead of stopping with only a regression-run artifact.

Human intervention summary: the human explicitly proposed the rule that failed or blocked regression should always emit a bug note and proceed into debugging, even if the eventual resolution is "need an upstream human."

Concrete better outcome the human was forcing: required-lane regression failure becomes durable defect tracking plus debug routing by default, not a dead-end status report.

Why this is a real intervention event rather than mere dissatisfaction: the human had to harden the next-step responsibility after seeing that the workflow could otherwise stop too early at regression bookkeeping.

Confidence: `strong`

Accepted-set read: `likely accepted incident candidate`

## Which Candidates Look Like Likely Accepted Incidents

- `C01` external research brief overconstrained the smarter model with a local worksheet shape
- `C03` the `100 KB` evidence-packet rule produced unreadable repo dumps instead of a usable brief
- `C09` the standing instruction to finish remaining passes got dropped at a pass boundary
- `C10` terminal subagent milestones were not surfaced immediately to the human
- `C11` Task 0004 was closed on a supporting lane instead of the human-facing regression lane
- `C12` regression authority and doc precedence were too weak, letting task-local docs redefine closure
- `C13` manual CA trust install was treated as a blocker instead of acceptable dev/operator friction
- `C14` failed regression could stop without opening a bug and routing into debugging

## Which Candidates Are Real Interventions But Probably Belong Outside The Accepted Incident Set

- `C02` hard `100 KB` brief rules were wired into the leader instead of a separate briefer role
- `C04` no-local-filesystem assumptions for the downstream handoff were not explicit enough
- `C05` `RESEARCH-LEADER` lacked explicit exemplar guidance for `RESEARCH-PLAN.md`
- `C07` `IMPLEMENTATION-LEADER` needed stricter research and plan-approval prerequisites
- `C08` `current_gate` semantics needed schema cleanup to remove `commit` / `push` / `notify`

## Ambiguous Boundaries That Need A Second Read

- `C06` looks like a real dropped-ball continuation problem, but the boundary between "reasonable conversational pause while answering steering questions" and "failure to preserve the umbrella instruction" is softer than the others.
