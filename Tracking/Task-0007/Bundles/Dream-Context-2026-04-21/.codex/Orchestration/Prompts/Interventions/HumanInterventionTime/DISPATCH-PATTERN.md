# Dispatch Pattern

Use this workflow as a sibling of daily harvest, not as part of needs analysis.

Recommended order:

1. Daily source discovery writes `SOURCE-SESSIONS.json`.
2. Daily harvest writes `SOURCE-PACKET.jsonl`.
3. Dispatch the intervention-time worker with [WORKFLOW.md](./WORKFLOW.md).
4. Dispatch needs extraction separately after harvest seals the packet.

Why this split exists:

- harvest owns event extraction
- intervention-time owns cost measurement
- needs extraction owns interpretation

Do not merge those concerns into one pass unless the task later proves that the combined workflow stays truthful and easier to audit.
