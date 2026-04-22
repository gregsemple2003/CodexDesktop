# Ingest Human Inputs

## Objective

Rebuild one canonical repo-day `HumanInputEvents` packet as JSON artifacts.

Assume the local day boundary for the packet is already resolved. Treat the provided source packet as the full authoritative set for that repo-day and do not search for more sessions or more inputs.

This prompt is for seeded rebuild only.
Seed discovery or bootstrap harvest belongs in [WORKFLOW.md](./WORKFLOW.md).

Use these placeholders:

- `<DAY_ROOT>`
  - canonical repo-day root such as `C:\Users\gregs\.codex\Orchestration\Reports\Interventions\<repo>\<YYYY>\<MM>\<DD>`

## Allowed Inputs Only

You may read only these files:

- `<DAY_ROOT>\HumanInputEvents\SOURCE-PACKET.jsonl`
- `<DAY_ROOT>\HumanInputEvents\SOURCE-SESSIONS.json`
- this prompt file

Do not read any other local files.
Do not search the repo.
Do not search `.codex\sessions`.
Do not browse the web.
Do not infer missing context from memory.

## Output Directory

Write all outputs into:

- `<DAY_ROOT>\HumanInputEvents\`

Do not write outside that directory.

## Required Outputs

1. Rewrite `SOURCE-PACKET.jsonl` in place as an enriched canonical packet.

Each output JSON object must preserve every existing top-level field and add:

- `parsed_envelope`

`parsed_envelope` rules:

- if `raw_text` matches the known Codex IDE wrapper format, set `parsed_envelope` to an object
- if the format is not recognized, set `parsed_envelope` to `null`
- do not add interpretation or analysis

Recognized wrapper format:

- starts with `# Context from my IDE setup:`
- may contain `## Active file:`
- may contain `## Open tabs:`
- contains `## My request for Codex:`

For recognized wrapper messages, `parsed_envelope` must be:

```json
{
  "format": "codex_ide_context_block_v1",
  "active_file": "string or null",
  "open_tabs": ["string", "..."],
  "request_text": "verbatim request body with wrapper removed",
  "has_image": true,
  "image_count": 1
}
```

Parsing rules:

- preserve `raw_text` verbatim at top level
- parse `active_file` from the active-file section when present
- parse `open_tabs` as an array of tab path strings from the open-tabs section when present
- parse `request_text` from the body after `## My request for Codex:`
- remove `<image>` marker blocks from `request_text`, but count them into `image_count`
- set `has_image` from whether one or more `<image>` marker blocks were present
- preserve meaningful internal newlines in `request_text`
- trim leading and trailing blank lines from `request_text`

2. Create `INDEX.json` in the same directory.

`INDEX.json` must be a JSON object with:

- `local_date`
- `repo_ref`
- `event_count`
- `events`

`events` must be an array in ascending `captured_at_local` order. Each array item must include:

- `event_id`
- `captured_at_local`
- `thread_name`
- `source_session_file`
- `source_line`
- `preview`

`preview` rules:

- if `parsed_envelope.request_text` exists and is non-empty, use its first non-empty line
- otherwise use the first non-empty line of `raw_text`
- do not summarize beyond taking that first line

3. Remove obsolete Markdown browse artifacts after the JSON outputs are successfully created:

- `INDEX.md` in this directory

Do not create or preserve `EVENT-*.json` in the promoted canonical packet.

Do not delete:

- `PROMPT-INPUT.md`
- `SOURCE-PACKET.jsonl`
- `SOURCE-SESSIONS.json`

## Ingest Rules

- Ingest every line from `SOURCE-PACKET.jsonl` exactly once.
- Treat each line as one `HumanInputEvent`.
- The source packet is already cleaned to remove control-wrapper records; do not do further filtering unless a line is malformed JSON.
- If a line is malformed, stop and report it instead of inventing output.
- Preserve line order when rewriting `SOURCE-PACKET.jsonl`.
- Do not summarize, redact, or normalize `raw_text`.

## Final Response

At the end, report:

- whether `INDEX.json` was created
- whether obsolete Markdown index files were removed
- the exact file paths changed
