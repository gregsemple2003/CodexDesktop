# INTERVENTION-PASS1 Candidate Report

Source day: `2026-04-08`

## Source Scope Reviewed

- Reviewed all `23` raw JSONL transcripts under [2026-04-08 session folder](/c:/Users/gregs/.codex/sessions/2026/04/08/).
- Read the current incident corpus contract first: [README.md](/c:/Users/gregs/.codex/Orchestration/Reports/Incidents/README.md), [INCIDENT.schema.json](/c:/Users/gregs/.codex/Orchestration/Reports/Incidents/INCIDENT.schema.json), [OUTBOUND-MESSAGE-REVIEW.schema.json](/c:/Users/gregs/.codex/Orchestration/Reports/Incidents/OUTBOUND-MESSAGE-REVIEW.schema.json), and [2026-04-08 day README](/c:/Users/gregs/.codex/Orchestration/Reports/Incidents/Daily/2026-04-08/README.md).
- Used [2026-04-08 OUTBOUND-MESSAGE-REVIEW.csv](/c:/Users/gregs/.codex/Orchestration/Reports/Incidents/Daily/2026-04-08/OUTBOUND-MESSAGE-REVIEW.csv) only as navigation aid. Candidate judgments below are grounded from raw transcript reads.
- Candidate-bearing or candidate-screened multi-turn sessions:
  - [rollout-2026-04-08T12-13-20-019d6dde-a292-76a3-a431-c3e72fa9c102.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T12-13-20-019d6dde-a292-76a3-a431-c3e72fa9c102.jsonl)
  - [rollout-2026-04-08T12-24-37-019d6de8-f7df-7901-bca9-6ab0d656d5c1.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T12-24-37-019d6de8-f7df-7901-bca9-6ab0d656d5c1.jsonl)
  - [rollout-2026-04-08T13-03-23-019d6e0c-7467-7042-92f5-aa976e10a3fe.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T13-03-23-019d6e0c-7467-7042-92f5-aa976e10a3fe.jsonl)
  - [rollout-2026-04-08T15-59-13-019d6ead-6ee7-7201-a865-9dd3f168dd1f.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T15-59-13-019d6ead-6ee7-7201-a865-9dd3f168dd1f.jsonl)
  - [rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl)
  - [rollout-2026-04-08T21-38-33-019d6fe4-199d-7f50-bf08-8d68086d978e.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T21-38-33-019d6fe4-199d-7f50-bf08-8d68086d978e.jsonl)
- The remaining `17` day-folder transcripts were one-shot digest launches or one-shot spawned worker runs with no follow-up human correction boundary visible in raw transcript.
- Important scope note: I treated the requested source day as the local corpus day-folder `C:\Users\gregs\.codex\sessions\2026\04\08\*.jsonl`. Some late-evening local activity appears in-record as `2026-04-09T...Z`; I kept those files because they live under the requested `2026-04-08` source-day folder.
- Important ordering note: [rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl) appears to preserve imported thread history with many same-second timestamps; inside that transcript, line order is the trustworthy chronology.

## Total Candidate Intervention Events Found

- Total: `11`
- `strong`: `7`
- `medium`: `3`
- `weak`: `1`

## Chronological Candidate List

### C01. Debug pass drifted from the confirmed crash seam into a side issue

Session or thread: `019d6de8-f7df-7901-bca9-6ab0d656d5c1`

Transcript: [rollout-2026-04-08T12-24-37-019d6de8-f7df-7901-bca9-6ab0d656d5c1.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T12-24-37-019d6de8-f7df-7901-bca9-6ab0d656d5c1.jsonl)

Primary refs: [L584](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T12-24-37-019d6de8-f7df-7901-bca9-6ab0d656d5c1.jsonl#L584), [L677](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T12-24-37-019d6de8-f7df-7901-bca9-6ab0d656d5c1.jsonl#L677), [L684](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T12-24-37-019d6de8-f7df-7901-bca9-6ab0d656d5c1.jsonl#L684), [L716](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T12-24-37-019d6de8-f7df-7901-bca9-6ab0d656d5c1.jsonl#L716)

AI course, outcome, or stopping point: after the resize-remap crash fix was effectively confirmed, the assistant kept tracing the separate `ReadPrevPresentationSyncState()` downstream issue and expanded into a broader analysis path.

Human intervention summary: the human explicitly said to pass on that side issue for now and stay focused on the presentation sync buffer resize, then later narrowed the implementation target again to a minimal production-grade engine patch.

Concrete better outcome the human was forcing: close on the confirmed crash seam first, with a minimal patch scoped to the presentation buffer crash fix rather than continuing to widen the investigation.

Why this is a real intervention event rather than mere dissatisfaction: the transcript shows a live debugging course already underway on the wrong seam, followed by a direct redirect back to the seam the human actually wanted closed.

Confidence: `weak`

Accepted-set read: `intervention event but probably not an accepted incident`

### C02. Digest jobs needed failure-alert behavior tied to real end state, not just run invocation

Session or thread: `019d6e0c-7467-7042-92f5-aa976e10a3fe`

Transcript: [rollout-2026-04-08T13-03-23-019d6e0c-7467-7042-92f5-aa976e10a3fe.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T13-03-23-019d6e0c-7467-7042-92f5-aa976e10a3fe.jsonl)

Primary refs: [L141](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T13-03-23-019d6e0c-7467-7042-92f5-aa976e10a3fe.jsonl#L141), [L178](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T13-03-23-019d6e0c-7467-7042-92f5-aa976e10a3fe.jsonl#L178), [L187](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T13-03-23-019d6e0c-7467-7042-92f5-aa976e10a3fe.jsonl#L187), [L190](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T13-03-23-019d6e0c-7467-7042-92f5-aa976e10a3fe.jsonl#L190)

AI course, outcome, or stopping point: the assistant was validating whether a manual digest run was visible and running, but the user still did not have a hard contract that a failed run would alert them when the real end state was missing.

Human intervention summary: after seeing the run behavior, the human tightened the requirement to "failure email if it didn't reach the end state - email sent, or report saved."

Concrete better outcome the human was forcing: backend orchestration should verify the real human-facing end state and emit a failure signal when report-save or email-send proof is missing.

Why this is a real intervention event rather than mere dissatisfaction: the human was not just asking for a feature in the abstract; they were correcting the adequacy bar for what counts as a successful digest run.

Confidence: `medium`

Accepted-set read: `intervention event but probably not an accepted incident`

### C03. Server launch fell back to an April 2 packaged server after an April 8 build request

Session or thread: `019d6ead-6ee7-7201-a865-9dd3f168dd1f`

Transcript: [rollout-2026-04-08T15-59-13-019d6ead-6ee7-7201-a865-9dd3f168dd1f.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T15-59-13-019d6ead-6ee7-7201-a865-9dd3f168dd1f.jsonl)

Primary refs: [L477](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T15-59-13-019d6ead-6ee7-7201-a865-9dd3f168dd1f.jsonl#L477), [L497](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T15-59-13-019d6ead-6ee7-7201-a865-9dd3f168dd1f.jsonl#L497), [L529](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T15-59-13-019d6ead-6ee7-7201-a865-9dd3f168dd1f.jsonl#L529), [L536](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T15-59-13-019d6ead-6ee7-7201-a865-9dd3f168dd1f.jsonl#L536), [L545](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T15-59-13-019d6ead-6ee7-7201-a865-9dd3f168dd1f.jsonl#L545), [L548](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T15-59-13-019d6ead-6ee7-7201-a865-9dd3f168dd1f.jsonl#L548), [L565](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T15-59-13-019d6ead-6ee7-7201-a865-9dd3f168dd1f.jsonl#L565)

AI course, outcome, or stopping point: the assistant gave the user a client command line against a running server, but that server was an older packaged build from April 2 after the requested April 8 server package had failed.

Human intervention summary: the human explicitly called this a dropped ball and restated the contract: if the assistant was asked to build the server that day, the launched server needed to be from that same day.

Concrete better outcome the human was forcing: matching same-day server and client build outputs, with no fallback to older packaged artifacts unless explicitly approved.

Why this is a real intervention event rather than mere dissatisfaction: the assistant had already represented the lane as runnable; the human had to step in to reject the actual launch basis as contractually wrong.

Confidence: `strong`

Accepted-set read: `likely accepted incident candidate`

### C04. Task 0006 plan lacked a falsifiable incident contract and had to be rebuilt around upstream reason tracing

Session or thread: `019d6942-63b6-7c72-9abc-9ab209816f97`

Transcript: [rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl)

Primary refs: [L723](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L723), [L730](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L730), [L789](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L789), [L796](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L796), [L854](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L854), [L869](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L869)

AI course, outcome, or stopping point: the assistant produced a plan and then a renamed `goal_stack` concept, but the human read both as hand-wavy and not falsifiable enough to distinguish real incidents from surface summaries.

Human intervention summary: the human rejected the plan, demanded upstream reason tracing that stayed falsifiable, then pushed the work from vague planning into concrete schema and examples.

Concrete better outcome the human was forcing: an incident contract that could be inspected and disproved, with machine-checkable schema and worked examples rather than soft prose.

Why this is a real intervention event rather than mere dissatisfaction: the assistant had already declared the plan ready for review; the human explicitly rejected it as inadequate on the core definition of what the task was supposed to preserve.

Confidence: `strong`

Accepted-set read: `likely accepted incident candidate`

### C05. Incident examples drifted into bug-report territory and lost the concrete correction event

Session or thread: `019d6942-63b6-7c72-9abc-9ab209816f97`

Transcript: [rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl)

Primary refs: [L918](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L918), [L925](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L925), [L938](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L938), [L950](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L950), [L1114](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L1114), [L1143](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L1143), [L1186](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L1186)

AI course, outcome, or stopping point: the first schema-backed example still read as a product bug report, and the incident fields were abstract enough that the concrete correction event was hard to recover from the record itself.

Human intervention summary: the human questioned why the example qualified at all, reframed the corpus as "human course correction incident" rather than generic incident, then forced `expected_state` and `actual_state` back down to the concrete event level.

Concrete better outcome the human was forcing: only preserve records where the assistant’s prior course, the human intervention, and the corrected event boundary are all explicit and readable from the incident alone.

Why this is a real intervention event rather than mere dissatisfaction: the human was not merely polishing prose; they were rejecting the assistant’s qualification boundary and forcing a narrower contract for what counts.

Confidence: `strong`

Accepted-set read: `likely accepted incident candidate`

### C06. The `why_chain` contract kept failing the reflexive human-only rule until the shape itself changed

Session or thread: `019d6942-63b6-7c72-9abc-9ab209816f97`

Transcript: [rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl)

Primary refs: [L1359](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L1359), [L1370](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L1370), [L1439](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L1439), [L1529](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L1529), [L1539](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L1539), [L1570](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L1570), [L1615](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L1615), [L1668](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L1668), [L1744](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L1744), [L1754](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L1754), [L1764](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L1764), [L1784](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L1784), [L1794](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L1794)

AI course, outcome, or stopping point: even after multiple revisions, the assistant kept treating `why_chain` as a place to preserve useful abstraction instead of a strict recursive answer to "why did the previous frame matter to the human?"

Human intervention summary: the human repeatedly checked whether the contract actually matched the documented rule, forced explicit schema prose for the rule, removed overlapping fields, and ultimately changed the shape to sibling `why_chains` when one strict linear chain could not honestly represent multiple upstream reasons.

Concrete better outcome the human was forcing: a linearly recursive, explicit-human-only causal structure that stops before unsupported inference and cleanly separates sibling high-level reasons.

Why this is a real intervention event rather than mere dissatisfaction: the assistant had already claimed the contract matched the rule; the human had to intervene again and again because the correction did not stick.

Confidence: `strong`

Accepted-set read: `likely accepted incident candidate`

### C07. Daily incident review had to widen into per-day outbound-message capture with over-capture scoring

Session or thread: `019d6942-63b6-7c72-9abc-9ab209816f97`

Transcript: [rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl)

Primary refs: [L2138](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L2138), [L2146](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L2146), [L2160](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L2160), [L2219](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L2219), [L2243](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L2243), [L2252](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L2252)

AI course, outcome, or stopping point: the assistant had produced more incident examples, but the human still lacked a durable, reviewable boundary for which outbound human statements counted as corrective and why.

Human intervention summary: the human redirected the process toward one CSV per day containing every outbound human message plus heuristic corrective sureness, then explicitly told the assistant to bias toward over-capture and not filter too early.

Concrete better outcome the human was forcing: a reviewable day-level candidate layer that preserves boundary uncertainty instead of silently shrinking the incident set too soon.

Why this is a real intervention event rather than mere dissatisfaction: the human intervened to prevent the assistant from hiding correction-boundary ambiguity inside example selection alone.

Confidence: `medium`

Accepted-set read: `intervention event but probably not an accepted incident`

### C08. Accepted incidents needed an explicit second pass for root-cause refinement and readable analysis

Session or thread: `019d6942-63b6-7c72-9abc-9ab209816f97`

Transcript: [rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl)

Primary refs: [L2679](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L2679), [L2747](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L2747), [L2759](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L2759), [L2768](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L2768), [L2864](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L2864), [L2976](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L2976), [L2979](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T16-53-21-019d6edf-017a-78e2-bc77-0df081782b3f.jsonl#L2979)

AI course, outcome, or stopping point: the accepted incident records were good at preserving the correction event, but still too surface-level about real mechanism and too hard to read directly inside the JSON.

Human intervention summary: the human asked for deeper cause research on an accepted incident, then turned that into a durable process rule: quick first-pass capture, second-pass root-cause refinement, and a readable `analysis` block for the refined incident.

Concrete better outcome the human was forcing: incident records that remain true to the correction event while also becoming honest about mechanism and legible to a human reviewer without horizontal-scrolling archaeology.

Why this is a real intervention event rather than mere dissatisfaction: the human is explicitly correcting the assistant’s stopping point for incident quality, not simply requesting extra nice-to-have detail.

Confidence: `medium`

Accepted-set read: `intervention event but probably not an accepted incident`

### C09. "Server is up" still used the wrong lane and hid a missing 51-mob population result

Session or thread: `019d6ead-6ee7-7201-a865-9dd3f168dd1f`

Transcript: [rollout-2026-04-08T15-59-13-019d6ead-6ee7-7201-a865-9dd3f168dd1f.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T15-59-13-019d6ead-6ee7-7201-a865-9dd3f168dd1f.jsonl)

Primary refs: [L861](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T15-59-13-019d6ead-6ee7-7201-a865-9dd3f168dd1f.jsonl#L861), [L869](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T15-59-13-019d6ead-6ee7-7201-a865-9dd3f168dd1f.jsonl#L869), [L876](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T15-59-13-019d6ead-6ee7-7201-a865-9dd3f168dd1f.jsonl#L876), [L939](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T15-59-13-019d6ead-6ee7-7201-a865-9dd3f168dd1f.jsonl#L939), [L976](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T15-59-13-019d6ead-6ee7-7201-a865-9dd3f168dd1f.jsonl#L976), [L995](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T15-59-13-019d6ead-6ee7-7201-a865-9dd3f168dd1f.jsonl#L995), [L1007](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T15-59-13-019d6ead-6ee7-7201-a865-9dd3f168dd1f.jsonl#L1007)

AI course, outcome, or stopping point: the assistant restarted with a fresh uncooked server and declared it up on `CombatMap`, but the lane was still effectively using stale cooked content and failed the user’s actual population check.

Human intervention summary: the human explicitly rejected the claimed ready state because the server was not spawning the expected 51 mobs and pushed the assistant to iterate until the evidence matched editor behavior.

Concrete better outcome the human was forcing: use the server lane that actually exercised live project content and proved the expected spawn behavior, not just a process that bound to the right map/port.

Why this is a real intervention event rather than mere dissatisfaction: the assistant had already closed the launch question at a weaker proxy; the human had to re-open it around the real gameplay outcome.

Confidence: `strong`

Accepted-set read: `likely accepted incident candidate`

### C10. April 4 daily brief output was still summary-shaped instead of a repeatable first-principles evidence packet

Session or thread: `019d6fe4-199d-7f50-bf08-8d68086d978e` (`Boyle` child thread)

Transcript: [rollout-2026-04-08T21-38-33-019d6fe4-199d-7f50-bf08-8d68086d978e.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T21-38-33-019d6fe4-199d-7f50-bf08-8d68086d978e.jsonl)

Primary refs: [L74](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T21-38-33-019d6fe4-199d-7f50-bf08-8d68086d978e.jsonl#L74), [L82](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T21-38-33-019d6fe4-199d-7f50-bf08-8d68086d978e.jsonl#L82), [L100](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T21-38-33-019d6fe4-199d-7f50-bf08-8d68086d978e.jsonl#L100), [L110](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T21-38-33-019d6fe4-199d-7f50-bf08-8d68086d978e.jsonl#L110), [L167](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T21-38-33-019d6fe4-199d-7f50-bf08-8d68086d978e.jsonl#L167), [L176](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T21-38-33-019d6fe4-199d-7f50-bf08-8d68086d978e.jsonl#L176), [L267](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T21-38-33-019d6fe4-199d-7f50-bf08-8d68086d978e.jsonl#L267), [L310](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T21-38-33-019d6fe4-199d-7f50-bf08-8d68086d978e.jsonl#L310)

AI course, outcome, or stopping point: the assistant rewrote the daily brief, but the "raw transcript" section was still filtered/projection-heavy, and the packet format was still optimized like a summary rather than a repeatable first-principles review artifact.

Human intervention summary: the human stopped the run, reframed the actual problem as prompt/contract design, forced a new subagent pass, then kept tightening the output shape toward dialogue-only verbatim transcript, inline incident JSON, and transcript arrays embedded in the JSON.

Concrete better outcome the human was forcing: a source-bounded packet that another reviewer can inspect from first principles instead of a polished paraphrase that quietly mixes source and interpretation.

Why this is a real intervention event rather than mere dissatisfaction: the human explicitly rejected the assistant’s chosen representation as non-first-principles and forced a new contract for how the evidence packet should be produced at all.

Confidence: `strong`

Accepted-set read: `likely accepted incident candidate`

### C11. Relayed subagent text was treated as literal human speech until the human challenged the provenance

Session or thread: `019d6fe4-199d-7f50-bf08-8d68086d978e` (`Boyle` child thread)

Transcript: [rollout-2026-04-08T21-38-33-019d6fe4-199d-7f50-bf08-8d68086d978e.jsonl](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T21-38-33-019d6fe4-199d-7f50-bf08-8d68086d978e.jsonl)

Primary refs: [L402](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T21-38-33-019d6fe4-199d-7f50-bf08-8d68086d978e.jsonl#L402), [L421](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T21-38-33-019d6fe4-199d-7f50-bf08-8d68086d978e.jsonl#L421), [L428](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T21-38-33-019d6fe4-199d-7f50-bf08-8d68086d978e.jsonl#L428), [L446](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T21-38-33-019d6fe4-199d-7f50-bf08-8d68086d978e.jsonl#L446), [L474](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T21-38-33-019d6fe4-199d-7f50-bf08-8d68086d978e.jsonl#L474), [L514](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T21-38-33-019d6fe4-199d-7f50-bf08-8d68086d978e.jsonl#L514), [L524](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T21-38-33-019d6fe4-199d-7f50-bf08-8d68086d978e.jsonl#L524), [L563](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T21-38-33-019d6fe4-199d-7f50-bf08-8d68086d978e.jsonl#L563), [L570](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T21-38-33-019d6fe4-199d-7f50-bf08-8d68086d978e.jsonl#L570), [L588](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T21-38-33-019d6fe4-199d-7f50-bf08-8d68086d978e.jsonl#L588), [L608](/c:/Users/gregs/.codex/sessions/2026/04/08/rollout-2026-04-08T21-38-33-019d6fe4-199d-7f50-bf08-8d68086d978e.jsonl#L608)

AI course, outcome, or stopping point: the assistant initially answered "yes" to whether a suspicious sentence was exact human wording because the raw child transcript labeled it as `user`, even though the sentence did not read like something the human would actually say.

Human intervention summary: the human challenged the wording, pointed out that it looked like AI self-talk, asked for deeper provenance investigation, and then set the rule that the process must prefer the verbatim actual human message, not relayed supervisor text.

Concrete better outcome the human was forcing: provenance-aware transcript harvesting that distinguishes direct human-authored text from relayed subagent instructions or editorialized/injected `user` content.

Why this is a real intervention event rather than mere dissatisfaction: the assistant had already made a concrete source claim and defended it; the human had to intervene because that claim would have encoded the wrong human evidence into the durable corpus.

Confidence: `strong`

Accepted-set read: `likely accepted incident candidate`

## Which Candidates Look Like Likely Accepted Incidents

- `C03` stale April 2 server fallback after April 8 build request
- `C04` Task 0006 plan lacked falsifiable incident reasoning and concrete contract
- `C05` incident examples drifted into bug-report territory instead of human course correction
- `C06` `why_chain` rules did not stick and had to become explicit `why_chains`
- `C09` "server is up" still used the wrong lane and failed the 51-mob gameplay check
- `C10` daily brief output was not actually first-principles evidence
- `C11` relayed subagent text was treated as literal human speech

## Which Candidates Are Real Interventions But Probably Belong Outside The Accepted Incident Set

- `C02` digest jobs needed failure-alert behavior tied to real end state
- `C07` daily review widened into full outbound-message CSV capture and over-capture scoring
- `C08` accepted incidents needed explicit second-pass root-cause refinement and readable analysis

## Ambiguous Boundaries That Need A Second Read

- `C01` could be a real "wrong seam" debugging correction, but it is also close to ordinary collaborative narrowing inside an active debug pass. The line-level intervention boundary is real; the accepted-incident boundary is less certain.

