# Prompt Input

This file captures the exact batch-specific input text used to generate the `HumanNeeds` packet in this directory.

Recovered on April 20, 2026 from Codex session transcripts via `codex-session-search`.

## Worker Launch Input

Transcript source:
[rollout-2026-04-20T10-30-41-019dab4c-f794-7f82-9e8c-6bcbb208808f.jsonl#L6](/c:/Users/gregs/.codex/sessions/2026/04/20/rollout-2026-04-20T10-30-41-019dab4c-f794-7f82-9e8c-6bcbb208808f.jsonl#L6)

This worker launch generated [PACKET-TRIAGE.json](./PACKET-TRIAGE.json) and [PACKET-RECORD.json](./PACKET-RECORD.json):

Historical note:
the quoted launch text below predates promotion into the shared `Reports\Interventions` root, so the task-local generation path is preserved verbatim.

```text
You are responsible only for files under C:\Agent\CodexDashboard\Tracking\Task-0007\NeedsExtraction\2026-04-19-thirdperson-prompt-pack-c. You are not alone in the codebase; do not revert unrelated edits and do not write outside that directory.

Task: follow C:\Agent\CodexDashboard\Tracking\Task-0007\NeedsExtraction\2026-04-19-thirdperson-prompt-pack-c\RUN-ANALYSIS.md exactly. Read only the files it permits. Produce the required JSON outputs in that same folder and nothing else.

When done, report the exact files written, representative event count, and whether a clarifying question remained necessary.
```

## Canonical Recipe And Local Context

- [PROMPT.md](../../../../../../../Prompts/Interventions/HumanNeeds/PROMPT.md)
- [WORKFLOW.md](../../../../../../../Prompts/Interventions/HumanNeeds/WORKFLOW.md)
- [LOCAL-CONTEXT.json](./LOCAL-CONTEXT.json)
