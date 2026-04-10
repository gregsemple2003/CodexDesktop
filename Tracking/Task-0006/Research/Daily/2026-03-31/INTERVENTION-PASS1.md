# INTERVENTION-PASS1 Candidate Report

Source day: `2026-03-31`

## Source Scope Reviewed

- Reviewed all `11` raw JSONL transcripts under [2026-03-31 session folder](/c:/Users/gregs/.codex/sessions/2026/03/31/).
- Read the current incident corpus contract first: [README.md](/c:/Users/gregs/.codex/Orchestration/Reports/Incidents/README.md), [INCIDENT.schema.json](/c:/Users/gregs/.codex/Orchestration/Reports/Incidents/INCIDENT.schema.json), and [OUTBOUND-MESSAGE-REVIEW.schema.json](/c:/Users/gregs/.codex/Orchestration/Reports/Incidents/OUTBOUND-MESSAGE-REVIEW.schema.json).
- I intentionally did not use same-day incident artifacts such as `README.md` or `OUTBOUND-MESSAGE-REVIEW.csv` for `2026-03-31`; this pass is grounded from raw transcripts plus the downstream contract only.
- Direct-human candidate-bearing root sessions:
  - [rollout-2026-03-31T09-49-34-019d4428-21ce-7d20-a144-c9cb83098d70.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/31/rollout-2026-03-31T09-49-34-019d4428-21ce-7d20-a144-c9cb83098d70.jsonl)
  - [rollout-2026-03-31T13-34-08-019d44f5-bd48-7432-baa3-01bd52b4bd13.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/31/rollout-2026-03-31T13-34-08-019d44f5-bd48-7432-baa3-01bd52b4bd13.jsonl)
- Screened but not counted as direct-human correction evidence:
  - two carryover child threads forked from a `2026-03-30` Crystallize parent session: [rollout-2026-03-31T00-47-19-019d4237-b2a9-7dc1-a855-86d25f8d4a6a.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/31/rollout-2026-03-31T00-47-19-019d4237-b2a9-7dc1-a855-86d25f8d4a6a.jsonl) and [rollout-2026-03-31T00-56-24-019d4240-0168-76b2-9999-7e82d5c71227.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/31/rollout-2026-03-31T00-56-24-019d4240-0168-76b2-9999-7e82d5c71227.jsonl)
  - three `.codex` `source=exec` runs with no human dialogue: [rollout-2026-03-31T05-00-05-019d431f-1d54-71c3-b47f-2dc30fc08bc1.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/31/rollout-2026-03-31T05-00-05-019d431f-1d54-71c3-b47f-2dc30fc08bc1.jsonl), [rollout-2026-03-31T05-00-05-019d431f-1d54-7292-a98b-c39ce93cde99.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/31/rollout-2026-03-31T05-00-05-019d431f-1d54-7292-a98b-c39ce93cde99.jsonl), and [rollout-2026-03-31T05-00-05-019d431f-1d54-7e83-bf46-6229437a6822.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/31/rollout-2026-03-31T05-00-05-019d431f-1d54-7e83-bf46-6229437a6822.jsonl)
  - delegated worker or child-leader threads whose `user_message` rows were supervisor prompts or status pings rather than fresh direct human corrections: [rollout-2026-03-31T09-35-33-019d441b-4e85-7bd1-8f3d-e5f532e5be53.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/31/rollout-2026-03-31T09-35-33-019d441b-4e85-7bd1-8f3d-e5f532e5be53.jsonl), [rollout-2026-03-31T12-01-21-019d44a0-ca4f-74b3-a47a-d3f9534f36a6.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/31/rollout-2026-03-31T12-01-21-019d44a0-ca4f-74b3-a47a-d3f9534f36a6.jsonl), [rollout-2026-03-31T13-08-00-019d44dd-cf67-7280-ab66-d452aacf5e95.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/31/rollout-2026-03-31T13-08-00-019d44dd-cf67-7280-ab66-d452aacf5e95.jsonl), and [rollout-2026-03-31T13-42-03-019d44fc-fbcf-75a2-90e9-ca4e3306baca.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/31/rollout-2026-03-31T13-42-03-019d44fc-fbcf-75a2-90e9-ca4e3306baca.jsonl)
- Important provenance note: several child-thread files preserve imported parent history from `2026-03-30`, and several child-thread `user_message` rows are clearly parent-agent workflow instructions. I did not count those as direct human intervention evidence for this pass.

## Total Candidate Intervention Events Found

- Total: `4`
- `strong`: `2`
- `medium`: `1`
- `weak`: `1`

## Chronological Candidate List

### C01. Camera-fix proposal crossed into engine edits until the human pushed it back to project code

Session or thread: `019d4428-21ce-7d20-a144-c9cb83098d70`

Transcript: [rollout-2026-03-31T09-49-34-019d4428-21ce-7d20-a144-c9cb83098d70.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/31/rollout-2026-03-31T09-49-34-019d4428-21ce-7d20-a144-c9cb83098d70.jsonl)

Primary refs: [L224](/c:/Users/gregs/.codex/sessions/2026/03/31/rollout-2026-03-31T09-49-34-019d4428-21ce-7d20-a144-c9cb83098d70.jsonl#L224), [L264](/c:/Users/gregs/.codex/sessions/2026/03/31/rollout-2026-03-31T09-49-34-019d4428-21ce-7d20-a144-c9cb83098d70.jsonl#L264), [L271](/c:/Users/gregs/.codex/sessions/2026/03/31/rollout-2026-03-31T09-49-34-019d4428-21ce-7d20-a144-c9cb83098d70.jsonl#L271), [L274](/c:/Users/gregs/.codex/sessions/2026/03/31/rollout-2026-03-31T09-49-34-019d4428-21ce-7d20-a144-c9cb83098d70.jsonl#L274)

AI course, outcome, or stopping point: the assistant converged on a recommended camera-smoothing fix that depended on editing `MoverComponent.h/.cpp` to add a new public presentation-state helper in engine code.

Human intervention summary: the human explicitly stopped that seam, said they did not want engine mods, and asked for an alternate approach.

Concrete better outcome the human was forcing: a project-only camera-presentation-anchor fix that reused the existing public Mover liaison API instead of changing engine source.

Why this is a real intervention event rather than mere dissatisfaction: the assistant had already proposed a concrete implementation shape and rationale; the human redirected the active solution seam after that proposal, not before it existed.

Confidence: `medium`

Accepted-set read: `intervention event but probably not an accepted incident`

### C02. Write-journal callstack capture had to be restored to a ready-to-use path before the next repro

Session or thread: `019d4428-21ce-7d20-a144-c9cb83098d70`

Transcript: [rollout-2026-03-31T09-49-34-019d4428-21ce-7d20-a144-c9cb83098d70.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/31/rollout-2026-03-31T09-49-34-019d4428-21ce-7d20-a144-c9cb83098d70.jsonl)

Primary refs: [L1539](/c:/Users/gregs/.codex/sessions/2026/03/31/rollout-2026-03-31T09-49-34-019d4428-21ce-7d20-a144-c9cb83098d70.jsonl#L1539), [L1552](/c:/Users/gregs/.codex/sessions/2026/03/31/rollout-2026-03-31T09-49-34-019d4428-21ce-7d20-a144-c9cb83098d70.jsonl#L1552), [L1559](/c:/Users/gregs/.codex/sessions/2026/03/31/rollout-2026-03-31T09-49-34-019d4428-21ce-7d20-a144-c9cb83098d70.jsonl#L1559), [L1582](/c:/Users/gregs/.codex/sessions/2026/03/31/rollout-2026-03-31T09-49-34-019d4428-21ce-7d20-a144-c9cb83098d70.jsonl#L1582)

AI course, outcome, or stopping point: the assistant clarified that the binary-callstack path still existed, but its earlier enablement change had made default `Start()` runs no longer capture callstacks without extra setup.

Human intervention summary: the human checked whether the old smooth callstack path still existed, then asked the assistant to re-enable callstacks before rerunning the repro.

Concrete better outcome the human was forcing: a ready-to-use write-journal capture path that preserved the existing binary-callstack and cached-symbol workflow for the next diagnostic pass.

Why this is a real intervention event rather than mere dissatisfaction: the transcript shows the assistant acknowledging its own prior enablement change and then restoring a more usable diagnostic state because the current default was not adequate for the next rerun.

Confidence: `weak`

Accepted-set read: `unclear and needs transcript reread`

### C03. Git push used a popup auth path and forced the human into the automation loop

Session or thread: `019d44f5-bd48-7432-baa3-01bd52b4bd13`

Transcript: [rollout-2026-03-31T13-34-08-019d44f5-bd48-7432-baa3-01bd52b4bd13.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/31/rollout-2026-03-31T13-34-08-019d44f5-bd48-7432-baa3-01bd52b4bd13.jsonl)

Primary refs: [L163](/c:/Users/gregs/.codex/sessions/2026/03/31/rollout-2026-03-31T13-34-08-019d44f5-bd48-7432-baa3-01bd52b4bd13.jsonl#L163), [L176](/c:/Users/gregs/.codex/sessions/2026/03/31/rollout-2026-03-31T13-34-08-019d44f5-bd48-7432-baa3-01bd52b4bd13.jsonl#L176), [L196](/c:/Users/gregs/.codex/sessions/2026/03/31/rollout-2026-03-31T13-34-08-019d44f5-bd48-7432-baa3-01bd52b4bd13.jsonl#L196), [L214](/c:/Users/gregs/.codex/sessions/2026/03/31/rollout-2026-03-31T13-34-08-019d44f5-bd48-7432-baa3-01bd52b4bd13.jsonl#L214), [L217](/c:/Users/gregs/.codex/sessions/2026/03/31/rollout-2026-03-31T13-34-08-019d44f5-bd48-7432-baa3-01bd52b4bd13.jsonl#L217)

AI course, outcome, or stopping point: the assistant treated the current git push path as a normal checkpoint step while cleaning the baseline and launching the next task workflow.

Human intervention summary: the human reported that the git command had triggered a popup window they had to click and explicitly called out that this would break later automation.

Concrete better outcome the human was forcing: a genuinely headless git/auth path for automation-sensitive flows, or at minimum an honest recognition that the current command shape was not automation-safe on this machine.

Why this is a real intervention event rather than mere dissatisfaction: the assistant’s chosen path already imposed manual operator work in the middle of an automation-oriented flow; the human had to intervene because the real human-world behavior contradicted the implied headless contract.

Confidence: `strong`

Accepted-set read: `likely accepted incident candidate`

### C04. Commit messages were too process-shaped until the human forced a leader-owned human-readable git narrative

Session or thread: `019d44f5-bd48-7432-baa3-01bd52b4bd13`

Transcript: [rollout-2026-03-31T13-34-08-019d44f5-bd48-7432-baa3-01bd52b4bd13.jsonl](/c:/Users/gregs/.codex/sessions/2026/03/31/rollout-2026-03-31T13-34-08-019d44f5-bd48-7432-baa3-01bd52b4bd13.jsonl)

Primary refs: [L411](/c:/Users/gregs/.codex/sessions/2026/03/31/rollout-2026-03-31T13-34-08-019d44f5-bd48-7432-baa3-01bd52b4bd13.jsonl#L411), [L426](/c:/Users/gregs/.codex/sessions/2026/03/31/rollout-2026-03-31T13-34-08-019d44f5-bd48-7432-baa3-01bd52b4bd13.jsonl#L426), [L433](/c:/Users/gregs/.codex/sessions/2026/03/31/rollout-2026-03-31T13-34-08-019d44f5-bd48-7432-baa3-01bd52b4bd13.jsonl#L433), [L496](/c:/Users/gregs/.codex/sessions/2026/03/31/rollout-2026-03-31T13-34-08-019d44f5-bd48-7432-baa3-01bd52b4bd13.jsonl#L496)

AI course, outcome, or stopping point: the workflow had already been producing generic, process-shaped commit comments that located the step mechanically but did not tell a human why the commit existed.

Human intervention summary: the human paused forward work, called the current commit comments overly generic, and pushed for both a human-readable summary bar and leader ownership of git narrative so workers would not carry that burden.

Concrete better outcome the human was forcing: a shared commit-message contract whose subjects carry both task-location context and a one-sentence human-readable summary, authored by leaders rather than by worker prompts.

Why this is a real intervention event rather than mere dissatisfaction: this was a direct rejection of already-produced workflow output and a durable reframing of the process contract, not a cosmetic preference note.

Confidence: `strong`

Accepted-set read: `likely accepted incident candidate`

## Which Candidates Look Like Likely Accepted Incidents

- `C03` git push path triggered a popup and forced manual human click-through in an automation-sensitive flow
- `C04` commit messages were too process-shaped until the human forced a leader-owned human-readable git narrative

## Which Candidates Are Real Interventions But Probably Belong Outside The Accepted Incident Set

- `C01` camera-fix proposal was redirected away from engine mods and back into project code

## Ambiguous Boundaries That Need A Second Read

- `C02` is a real usability correction boundary if the key fact is "the assistant’s prior change made callstack capture materially less ready-to-use." It is ordinary investigation collaboration if the key fact is only "the human asked to enable a diagnostic before the next rerun." The transcript supports both reads, so I kept it as a weak recall-first candidate.
- The root Rustfire framing note at [L91](/c:/Users/gregs/.codex/sessions/2026/03/31/rollout-2026-03-31T09-49-34-019d4428-21ce-7d20-a144-c9cb83098d70.jsonl#L91) may be a mild redirect away from over-attributing the issue to local changes, but I did not promote it because the assistant had not yet committed to a wrong diagnosis strongly enough to make the intervention boundary clear.
