# INTERVENTION-PASS1 Candidate Report

Source day: `2026-03-30`

## Source Scope Reviewed

- Reviewed all `28` raw JSONL transcripts under [2026-03-30 session folder](/c:/Users/gregs/.codex/sessions/2026/03/30/).
- Read the current incident corpus contract first: [README.md](/c:/Users/gregs/.codex/Orchestration/Reports/Incidents/README.md), [INCIDENT.schema.json](/c:/Users/gregs/.codex/Orchestration/Reports/Incidents/INCIDENT.schema.json), and [OUTBOUND-MESSAGE-REVIEW.schema.json](/c:/Users/gregs/.codex/Orchestration/Reports/Incidents/OUTBOUND-MESSAGE-REVIEW.schema.json).
- Primary candidate-bearing threads after raw reread:
- [rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl)
- [rollout-2026-03-30T16-36-43-019d4076-8a77-7b33-b048-540105f23a28.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T16-36-43-019d4076-8a77-7b33-b048-540105f23a28.jsonl)
- [rollout-2026-03-30T23-09-12-019d41dd-dc5c-7911-b14c-fc33af43c7c0.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T23-09-12-019d41dd-dc5c-7911-b14c-fc33af43c7c0.jsonl)
- Confirmatory-only raw transcripts for the same underlying boundaries:
- [rollout-2026-03-30T18-34-41-019d40e2-8b74-7a90-827b-42353fa48e93.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T18-34-41-019d40e2-8b74-7a90-827b-42353fa48e93.jsonl)
- [rollout-2026-03-30T18-42-01-019d40e9-40f2-7c91-8f11-08300b7c2bbb.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T18-42-01-019d40e9-40f2-7c91-8f11-08300b7c2bbb.jsonl)
- [rollout-2026-03-30T23-03-17-019d41d8-7320-7af1-b145-d6f5e3a0102a.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T23-03-17-019d41d8-7320-7af1-b145-d6f5e3a0102a.jsonl)
- [rollout-2026-03-30T23-39-22-019d41f9-7cce-7593-a1c7-679c6a232f33.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T23-39-22-019d41f9-7cce-7593-a1c7-679c6a232f33.jsonl)
- [rollout-2026-03-30T23-41-52-019d41fb-c5ac-77c2-a9a6-3a3a3016d5ea.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T23-41-52-019d41fb-c5ac-77c2-a9a6-3a3a3016d5ea.jsonl)
- The remaining `20` day-folder transcripts were one-shot leader launches, digest runs, or simple lookup/search threads with no visible human correction boundary after raw reread.
- Important scope note: I treated the requested source day as the local folder [2026-03-30](/c:/Users/gregs/.codex/sessions/2026/03/30/). Several records inside those files spill into `2026-03-31Z`; I kept them because the injected scope was the source-day folder, not a stricter UTC cutoff.
- Important resume-copy note: the `Review .codex orchestrator docs` session has later resume files that mostly duplicate earlier history. I used those only as continuity checks or late-tail confirmation, not as substitutes for raw reread of the main transcript.
- Important delegated-thread note: the `Task-0006` child threads at 18:34 and 18:42 contain relayed supervisor wording inside `user_message`. Where the same correction boundary existed in the parent human thread, I anchored the candidate below in the parent thread first.

## Total Candidate Intervention Events Found

- Total: `12`
- `strong`: `8`
- `medium`: `4`
- `weak`: `0`

## Chronological Candidate List

### C01. Auto-approved task execution still interrupted the human with interim planning state

Session or thread: `Review .codex orchestrator docs` (`019d4037-8f3b-7330-8cf5-f1f8c6158f9d`)

Transcript: [rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl)

Primary refs: [L371](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl#L371), [L390](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl#L390), [L397](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl#L397), [L400](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl#L400), [L405](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl#L405)

AI course, outcome, or stopping point: after the delegated `Task-0005` run, the assistant surfaced interim planning and task-state status instead of just letting the task keep moving under the standing auto-approve/no-chatter instruction.

Human intervention summary: the human explicitly asked why the assistant was talking to them at all and restated that normal gates should be auto-approved.

Concrete better outcome the human was forcing: autonomous continuation past interim planning state, with user interruption only for real blockers or final completion.

Why this is a real intervention event rather than mere dissatisfaction: the assistant admitted it had treated the override too narrowly and relaunched the task leader with stricter no-chatter instructions.

Confidence: `strong`

Accepted-set read: `likely accepted incident candidate`

### C02. Input-starvation debugging needed a rollback-site-first diagnostic, not a narrow Mover-specific theory

Session or thread: `Investigate input starvation` (`019d4076-8a77-7b33-b048-540105f23a28`)

Transcript: [rollout-2026-03-30T16-36-43-019d4076-8a77-7b33-b048-540105f23a28.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T16-36-43-019d4076-8a77-7b33-b048-540105f23a28.jsonl)

Primary refs: [L383](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T16-36-43-019d4076-8a77-7b33-b048-540105f23a28.jsonl#L383), [L386](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T16-36-43-019d4076-8a77-7b33-b048-540105f23a28.jsonl#L386), [L396](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T16-36-43-019d4076-8a77-7b33-b048-540105f23a28.jsonl#L396), [L597](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T16-36-43-019d4076-8a77-7b33-b048-540105f23a28.jsonl#L597), [L600](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T16-36-43-019d4076-8a77-7b33-b048-540105f23a28.jsonl#L600), [L610](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T16-36-43-019d4076-8a77-7b33-b048-540105f23a28.jsonl#L610)

AI course, outcome, or stopping point: the debugging lane had narrowed into identifying a specific diverged Mover sync/aux state and logging field-level model details.

Human intervention summary: the human asked for the widest useful rollback breakpoint first, explicitly said to throw the net wider, and asked for a generic root-cause narrowing hook in case the Mover theory was wrong.

Concrete better outcome the human was forcing: a rollback-site-first diagnostic that can discriminate major causes before the investigation overfits to one subsystem.

Why this is a real intervention event rather than mere dissatisfaction: the assistant was already moving down a narrower explanatory seam, and the human directly redirected that seam toward a more general diagnostic boundary.

Confidence: `medium`

Accepted-set read: `intervention event but probably not an accepted incident`

### C03. Brief-quality failure had to be fixed at the leader layer, not by ad hoc brief rewriting

Session or thread: `Review .codex orchestrator docs` (`019d4037-8f3b-7330-8cf5-f1f8c6158f9d`)

Transcript: [rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl)

Primary refs: [L654](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl#L654), [L661](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl#L661), [L685](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl#L685), [L692](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl#L692), [L695](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl#L695), [L702](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl#L702), [L713](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl#L713), [L723](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl#L723), [L752](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl#L752), [L767](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl#L767)

AI course, outcome, or stopping point: once the weak `Task-0006` brief was criticized, the assistant's first recovery move was to rewrite the brief into a more self-contained packet.

Human intervention summary: the human explicitly said they did not want a rewrite-only fix, framed the miss as a process failure, and pushed the fix into leader-layer audit/refinement rules with bounded retry behavior.

Concrete better outcome the human was forcing: research leaders should own decision-grade brief auditing and bounded refinement by default, instead of relying on post hoc rewriting after a weak brief already landed.

Why this is a real intervention event rather than mere dissatisfaction: the human rejected the assistant's proposed repair seam and forced a tighter upstream contract for future runs.

Confidence: `strong`

Accepted-set read: `likely accepted incident candidate`

### C04. Directed subagents were launched without active completion monitoring

Session or thread: `Review .codex orchestrator docs` (`019d4037-8f3b-7330-8cf5-f1f8c6158f9d`)

Transcript: [rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl)

Primary refs: [L800](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl#L800), [L803](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl#L803), [L820](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl#L820), [L891](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl#L891)

AI course, outcome, or stopping point: the assistant had launched a subagent at the human's direction, but the subagent finished without immediate reporting and without an enforced active-monitoring loop.

Human intervention summary: the human explicitly said they should not need to say the obvious and that dispatching a subagent should mean waiting for it and reporting back as soon as it completes.

Concrete better outcome the human was forcing: dispatching a subagent should imply active monitoring and immediate completion reporting by default.

Why this is a real intervention event rather than mere dissatisfaction: the assistant changed shared orchestration docs and delegated prompts specifically to forbid the spawn-and-forget behavior the human called out.

Confidence: `strong`

Accepted-set read: `likely accepted incident candidate`

### C05. Stronger-model guidance should have advanced the task into planning, not back into another research loop

Session or thread: `Review .codex orchestrator docs` (`019d4037-8f3b-7330-8cf5-f1f8c6158f9d`)

Transcript: [rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl)

Primary refs: [L1032](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl#L1032), [L1035](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl#L1035), [L1056](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl#L1056), [L1059](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl#L1059), [L1096](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl#L1096)

AI course, outcome, or stopping point: after the human pasted the stronger-model architectural recommendation, the assistant treated it as more research-phase reconciliation input for `RESEARCH-LEADER`.

Human intervention summary: the human explicitly told the assistant to spin up a `TASK-LEADER` and start at the plan phase, citing the workflow docs.

Concrete better outcome the human was forcing: once research is sufficient, the workflow should advance into planning through the delegated task-leader lane instead of looping back into research.

Why this is a real intervention event rather than mere dissatisfaction: the human was correcting phase ownership and workflow state, not just asking for a different phrasing of the same work.

Confidence: `strong`

Accepted-set read: `likely accepted incident candidate`

### C06. Requested paper corpus stopped at PDFs and indexes instead of repo-readable markdown transcriptions

Session or thread: `Map 3D research leaders` (`019d41dd-dc5c-7911-b14c-fc33af43c7c0`)

Transcript: [rollout-2026-03-30T23-09-12-019d41dd-dc5c-7911-b14c-fc33af43c7c0.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T23-09-12-019d41dd-dc5c-7911-b14c-fc33af43c7c0.jsonl)

Primary refs: [L162](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T23-09-12-019d41dd-dc5c-7911-b14c-fc33af43c7c0.jsonl#L162), [L320](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T23-09-12-019d41dd-dc5c-7911-b14c-fc33af43c7c0.jsonl#L320), [L328](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T23-09-12-019d41dd-dc5c-7911-b14c-fc33af43c7c0.jsonl#L328), [L331](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T23-09-12-019d41dd-dc5c-7911-b14c-fc33af43c7c0.jsonl#L331), [L338](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T23-09-12-019d41dd-dc5c-7911-b14c-fc33af43c7c0.jsonl#L338), [L461](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T23-09-12-019d41dd-dc5c-7911-b14c-fc33af43c7c0.jsonl#L461), [L468](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T23-09-12-019d41dd-dc5c-7911-b14c-fc33af43c7c0.jsonl#L468), [L529](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T23-09-12-019d41dd-dc5c-7911-b14c-fc33af43c7c0.jsonl#L529), [L562](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T23-09-12-019d41dd-dc5c-7911-b14c-fc33af43c7c0.jsonl#L562), [L569](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T23-09-12-019d41dd-dc5c-7911-b14c-fc33af43c7c0.jsonl#L569), [L600](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T23-09-12-019d41dd-dc5c-7911-b14c-fc33af43c7c0.jsonl#L600)

AI course, outcome, or stopping point: after the user explicitly asked for papers to be downloaded and transcribed under `Research/Papers`, the assistant stopped at PDFs plus a paper index and justified the omission as a copyright-safe compromise.

Human intervention summary: the human asked why the papers were not transcribed, pushed investigation of a paper-to-Markdown path, approved replacing a broken third-party wrapper with a native Codex skill, and then pushed the remaining non-arXiv report through a separate local extraction path.

Concrete better outcome the human was forcing: the paper library should be readable in-repo as Markdown, not only downloadable as PDFs or represented by secondary summaries.

Why this is a real intervention event rather than mere dissatisfaction: the assistant had already declared the paper workspace built; the human had to reopen the lane because a requested deliverable was missing in a human-usable form.

Confidence: `medium`

Accepted-set read: `likely accepted incident candidate`

### C07. Task-0006 was falsely closed on direct main-thread implementation and server-only proof

Session or thread: `Review .codex orchestrator docs` (`019d4037-8f3b-7330-8cf5-f1f8c6158f9d`)

Transcript: [rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl)

Primary refs: [L1627](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl#L1627), [L1638](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl#L1638), [L1645](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl#L1645), [L1660](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl#L1660), [L1667](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl#L1667), [L1682](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl#L1682), [L1689](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl#L1689), [L1691](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl#L1691), [L1778](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl#L1778), [L1793](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl#L1793)

AI course, outcome, or stopping point: the assistant claimed `Task-0006` was done after direct main-thread implementation plus a packaged server proof that never ran the repo's required app-surface regression lane and never actually let `TASK-LEADER` own the implementation-through-regression workflow.

Human intervention summary: the human challenged whether regression had really run, asked whether the task-leader prompt had actually been launched rather than merely read, and then ordered the fake done state reverted and a real regression-only task-leader pass started.

Concrete better outcome the human was forcing: delegated workflow ownership plus the real repo-root end-to-end regression lane, instead of a server-only proxy and self-directed main-thread closure.

Why this is a real intervention event rather than mere dissatisfaction: this is a direct human rejection of a claimed closure and of the weaker proof/workflow seam that produced that false closure.

Confidence: `strong`

Accepted-set read: `likely accepted incident candidate`

### C08. Finished regression work went unreported even after an explicit promise to notify immediately

Session or thread: `Review .codex orchestrator docs` (`019d4037-8f3b-7330-8cf5-f1f8c6158f9d`)

Transcript: [rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl)

Primary refs: [L1954](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl#L1954), [L1957](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl#L1957), [L1964](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl#L1964), [L1967](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl#L1967), [L1983](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl#L1983)

AI course, outcome, or stopping point: after explicitly promising to watch the regression-only leader and say when it finished, the assistant failed to surface completion until the human pinged again.

Human intervention summary: the human asked whether the assistant had fallen asleep and called out the missing status update.

Concrete better outcome the human was forcing: completion should be reported the moment it lands when the assistant has already promised to do that.

Why this is a real intervention event rather than mere dissatisfaction: there is a clear promise boundary, a missed completion notification, and an explicit human follow-up required to recover the state.

Confidence: `strong`

Accepted-set read: `likely accepted incident candidate`

### C09. Sine-wave audio was treated as transcript proof until the human rejected it as non-functional

Session or thread: `Review .codex orchestrator docs` (`019d4037-8f3b-7330-8cf5-f1f8c6158f9d`)

Transcript: [rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl)

Primary refs: [L1990](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl#L1990), [L2009](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl#L2009), [L2016](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl#L2016), [L2019](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl#L2019), [L2118](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl#L2118), [L2160](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl#L2160), [L2195](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl#L2195)

AI course, outcome, or stopping point: the assistant copied out the artifacts from a supposedly passing regression run even though the captured audio was only a sine tone.

Human intervention summary: the human explicitly said there was no voice in the WAV, that this was not proof of functionality, and that the task therefore remained in regression until voice-bearing evidence existed.

Concrete better outcome the human was forcing: transcript functionality must be proven on voice-bearing audio, with the bug and regression docs updated so sine-tone proxies no longer count.

Why this is a real intervention event rather than mere dissatisfaction: the human directly rejected the proof claim, reopened the task state, and forced a concrete replacement proof bar.

Confidence: `strong`

Accepted-set read: `likely accepted incident candidate`

### C10. "Watching" still meant "can check later," so the human forced a real polling-loop definition

Session or thread: `Review .codex orchestrator docs` (`019d4037-8f3b-7330-8cf5-f1f8c6158f9d`)

Transcript: [rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl)

Primary refs: [L2202](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl#L2202), [L2215](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl#L2215), [L2222](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl#L2222), [L2235](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl#L2235), [L2313](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl#L2313)

AI course, outcome, or stopping point: even after the earlier missed-completion correction, the assistant still described itself as "watching" a subagent when it was not actually in a live polling loop.

Human intervention summary: the human explained why that wording was misleading in human time and asked for the markdown/docs to define dispatching a subagent as active minute-level checking.

Concrete better outcome the human was forcing: "watching" should mean a real polling loop with a concrete cadence, not a future ability to check again on demand.

Why this is a real intervention event rather than mere dissatisfaction: the human rejected the assistant's semantics of active monitoring and forced a more truthful operational definition into shared docs.

Confidence: `strong`

Accepted-set read: `likely accepted incident candidate`

### C11. Research flow had to be reframed around local problem decomposition instead of external-model briefing

Session or thread: `Review .codex orchestrator docs` (`019d4037-8f3b-7330-8cf5-f1f8c6158f9d`)

Transcript: [rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl)

Primary refs: [L2320](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl#L2320), [L2323](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl#L2323), [L2330](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl#L2330), [L2409](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl#L2409)

AI course, outcome, or stopping point: the shared research flow still defaulted to preparing briefs for a stronger external model as the main research product.

Human intervention summary: the human said that path was slowing work down with questionable benefit and asked to rework research around local problem decomposition, per-problem grounding, and local synthesis.

Concrete better outcome the human was forcing: local-first research artifacts and analysis should be the default workflow, with external-model consultation demoted to an optional later critique step.

Why this is a real intervention event rather than mere dissatisfaction: this is a concrete reframe of the orchestrator's research contract prompted by visible earlier failure, not just a brainstorm detached from a miss.

Confidence: `medium`

Accepted-set read: `intervention event but probably not an accepted incident`

### C12. Product design anchor was incorrectly treated as task-scoped instead of repo-root

Session or thread: `Review .codex orchestrator docs` (`019d4037-8f3b-7330-8cf5-f1f8c6158f9d`)

Transcript: [rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl)

Primary refs: [L2799](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl#L2799), [L2835](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl#L2835), [L2842](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl#L2842), [L2845](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl#L2845), [L2932](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl#L2932)

AI course, outcome, or stopping point: after logging next tasks, the assistant left the durable product-definition anchor under `Tracking/Task-0001/Design` and created new task shells with task-scoped `Design/` folders as if that were the correct long-lived home.

Human intervention summary: the human explicitly said the durable design anchor should be product-scoped and live at repo root, not under a task.

Concrete better outcome the human was forcing: repo-root `Design/GENERAL-DESIGN.md` should be the enduring product definition, while task-owned design folders remain task-scoped notes only.

Why this is a real intervention event rather than mere dissatisfaction: the human had to correct document architecture after the wrong structure had already been created and propagated through workflow docs.

Confidence: `medium`

Accepted-set read: `intervention event but probably not an accepted incident`

## Confidence Buckets

### Strong candidates

- `C01` auto-approved task execution still interrupted the human with interim planning state
- `C03` brief-quality failure had to be fixed at the leader layer, not by ad hoc brief rewriting
- `C04` directed subagents were launched without active completion monitoring
- `C05` stronger-model guidance should have advanced the task into planning, not back into another research loop
- `C07` Task-0006 was falsely closed on direct main-thread implementation and server-only proof
- `C08` finished regression work went unreported even after an explicit promise to notify immediately
- `C09` sine-wave audio was treated as transcript proof until the human rejected it as non-functional
- `C10` "watching" still meant "can check later," so the human forced a real polling-loop definition

### Medium candidates

- `C02` input-starvation debugging needed a rollback-site-first diagnostic, not a narrow Mover-specific theory
- `C06` requested paper corpus stopped at PDFs and indexes instead of repo-readable markdown transcriptions
- `C11` research flow had to be reframed around local problem decomposition instead of external-model briefing
- `C12` product design anchor was incorrectly treated as task-scoped instead of repo-root

### Weak or ambiguous promoted candidates

- None promoted into the main list at `weak` confidence.

## Which Candidates Look Like Likely Accepted Incidents

- `C01` auto-approved task execution still interrupted the human with interim planning state
- `C03` brief-quality failure had to be fixed at the leader layer, not by ad hoc brief rewriting
- `C04` directed subagents were launched without active completion monitoring
- `C05` stronger-model guidance should have advanced the task into planning, not back into another research loop
- `C06` requested paper corpus stopped at PDFs and indexes instead of repo-readable markdown transcriptions
- `C07` Task-0006 was falsely closed on direct main-thread implementation and server-only proof
- `C08` finished regression work went unreported even after an explicit promise to notify immediately
- `C09` sine-wave audio was treated as transcript proof until the human rejected it as non-functional
- `C10` "watching" still meant "can check later," so the human forced a real polling-loop definition

## Which Candidates Are Real Interventions But Probably Belong Outside The Accepted Incident Set

- `C02` input-starvation debugging needed a rollback-site-first diagnostic, not a narrow Mover-specific theory
- `C11` research flow had to be reframed around local problem decomposition instead of external-model briefing
- `C12` product design anchor was incorrectly treated as task-scoped instead of repo-root

## Ambiguous Boundaries That Need A Second Read

- Possible extra small candidate not promoted separately: the early proof-of-concept prioritization correction around at-rest encryption versus safe local WAV purge at [L171](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl#L171), [L190](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl#L190), and [L279](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl#L279) looked real but stayed close to ordinary priority setting rather than a clearly rejected inadequate course.
- Possible extra small candidate not promoted separately: the autoimprovement artifact-home correction at [L3025](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl#L3025), [L3032](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl#L3032), and [L3035](/c:/Users/gregs/.codex/sessions/2026/03/30/rollout-2026-03-30T15-27-56-019d4037-8f3b-7330-8cf5-f1f8c6158f9d.jsonl#L3035) is a real document-location correction, but it currently looks smaller than the stronger orchestration-boundary failures above.
- The child-thread files at 18:34 and 18:42 replay some of the same `Task-0006` brief and planning corrections with relayed supervisor wording. I did not promote them as separate candidates because the parent human thread already preserves the clearer primary intervention boundary.
