# Problem 0004 Option A Plan 0001

## Intent

Auto-build a review packet for approval turns with a concise diff summary, exact file links, pass framing, and state delta.

## What Changes

- Every approval ask emits a short review packet instead of a loose link dump.
- The packet explains what changed, where it changed, which pass it belongs to, and what needs approval.
- Durable state changes are shown beside content changes.

## Files Or Artifact Types That Move

- Approval-facing summaries in task `HANDOFF.md`, `PLAN.md`, or task-owned review notes.
- Exact file-link bundles for changed artifacts.
- Shared review-packet templates that can be reused across repos.

## Rollout

1. Define the minimum packet fields: pass id, changed files, short change summary, state change, and approval ask.
2. Generate the packet automatically whenever the workflow reaches an approval gate.
3. Pilot the packet on a `ThirdPerson` planning pass that changes multiple task artifacts.
4. Refine the packet until one screen is enough to understand the approval ask.

## Success Check

- Approval requests always include a readable change summary and exact links.
- Pass framing is visible without the human searching the tree.
- The human can approve or reject from the packet alone in common cases.

## Burden Reduction Under Directional Context

`Truth`: the review surface shows what really changed.

`Compassion`: it removes manual diff hunting and pass reconstruction.

`Tolerance`: it gives a stable review shape even when the human asks abruptly for "links" or "diffs."
