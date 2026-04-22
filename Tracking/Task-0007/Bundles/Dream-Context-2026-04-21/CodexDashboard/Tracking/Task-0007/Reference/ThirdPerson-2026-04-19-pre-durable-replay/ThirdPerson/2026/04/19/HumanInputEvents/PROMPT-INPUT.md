# Prompt Input

This file captures the exact batch-specific input text used to generate the April 19 `ThirdPerson` `HumanInputEvents` packet in this directory.

Recovered on April 20, 2026 from Codex session transcripts via `codex-session-search`.

## Current Canonical JSON Rebuild

Transcript source:
[rollout-2026-04-19T16-56-27-019da787-cba0-74d1-9075-eee929c447c5.jsonl#L6](/c:/Users/gregs/.codex/sessions/2026/04/19/rollout-2026-04-19T16-56-27-019da787-cba0-74d1-9075-eee929c447c5.jsonl#L6)

This worker launch produced the current canonical JSON-first packet:

Historical note:
the quoted launch text below predates promotion into the shared `Reports\Interventions` root, so the task-local staging path and earlier `EVENT-*.json` references are preserved verbatim.

```text
You are responsible only for files under C:\Agent\CodexDashboard\Tracking\Task-0007\HumanInputEvents\04\19. You are not alone in the codebase; do not revert unrelated edits and do not write outside that directory.

Task: follow the prompt at C:\Agent\CodexDashboard\Tracking\Task-0007\HumanInputEvents\04\19\INGEST-HUMAN-INPUTS.md exactly. Read only the files the prompt permits. Rebuild the April 19 ThirdPerson ingest as canonical JSON artifacts, including enriching SOURCE-PACKET.jsonl with parsed_envelope, creating EVENT-*.json and INDEX.json, and removing obsolete EVENT-*.md and INDEX.md if successful.

When done, report counts, exact changed files, and any parsing edge cases you encountered.
```

## Superseded Initial Ingest Pass

Transcript source:
[rollout-2026-04-19T16-41-01-019da779-a9b8-72a2-90e5-fb7e7f08c0d0.jsonl#L6](/c:/Users/gregs/.codex/sessions/2026/04/19/rollout-2026-04-19T16-41-01-019da779-a9b8-72a2-90e5-fb7e7f08c0d0.jsonl#L6)

This earlier worker launch created the now-superseded Markdown-first packet before the JSON rebuild above replaced it:

```text
Context-free ingest task. You are not allowed to use this parent thread as context.

Your write scope is only:
C:\Agent\CodexDashboard\Tracking\Task-0007\HumanInputEvents\04\19\

You are not alone in the codebase. Do not revert anyone else's edits. Only create or update files inside that directory.

Read and follow exactly:
C:\Agent\CodexDashboard\Tracking\Task-0007\HumanInputEvents\04\19\INGEST-HUMAN-INPUTS.md

Allowed input files are the ones named in that prompt. Do not read any other files.

Perform the ingest and then report:
- total event files created
- whether INDEX.md was created
- exact file paths changed
```

## Canonical Recipe

[PROMPT.md](../../../../../../../Prompts/Interventions/HumanInputEvents/PROMPT.md)
