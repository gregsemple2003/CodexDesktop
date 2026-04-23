# Title

Gate replies with direct-answer and approval-packet validators

## Summary

Introduce a reply-shaping gate that enforces answer-first behavior on direct-question turns and requires a contextual approval packet before plan or pass approval is requested.

## Source Event IDs

```json
[
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-3392",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4379",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4399",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4421",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4512",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4522",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4571",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5001",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5010",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5020",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5030",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5073",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5090",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5281",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5290"
]
```

## Burden Being Reduced

The human should receive a usable first answer or a usable approval packet immediately, rather than reconstructing the reply shape from a mixed narrative.

## Causal Claim

The burden persists because the system has no enforced reply mode and no durable approval artifact that a reviewer can inspect without rebuilding the diff and pass framing manually.

## Evidence

- [../SessionExcerpts/INDEX.json](../SessionExcerpts/INDEX.json) preserves the repair loops for both direct answers and approval surfaces.
- [../BURDEN-ANALYSIS.md](../BURDEN-ANALYSIS.md) treats these as one burden driver because the failure boundary is the human-facing reply surface.

## Goals

- force direct-question turns to answer in the first sentence and requested shape
- require an approval packet for plan/pass approval turns
- keep approval packets contextual, not header-only path dumps

## Non-Goals

- deciding product policy
- replacing task-state storage
- fixing proof-truth by itself

## Implementation Home

- the shared reply-shaping or message-validation layer that runs before a worker sends a human-facing reply
- shared orchestration prompts that classify reply mode
- a new task-owned approval artifact such as `Tracking/Task-<id>/APPROVAL-PACKET-<NNNN>.md`

## Proposed Changes

- classify outgoing replies into at least:
  - direct-answer mode
  - approval-request mode
- in direct-answer mode, require the first sentence to answer the question plainly before added context
- in approval-request mode, require a task-owned approval packet with:
  - changed artifacts
  - contextual summary of what changed
  - links to the changed artifacts
  - explicit pass framing
- block send when the required mode contract is missing

## Expected Resolution

The human no longer has to ask where the pass is, where the links are, or whether the question was answered before they can continue the work.

## What Does Not Count

- a reply that starts with framing and only later answers
- links with no statement of what changed
- asking for approval without a task-owned approval packet

## Why This Mechanism

This addresses both failure modes at the same boundary: the message that reaches the human. It is stronger than reminder language and smaller than a full orchestration rewrite.

## Human Relief If Successful

The human can answer yes or no to approval requests quickly and can evaluate direct answers without repeated clarification loops.

## Remaining Uncertainty

Mode detection must be reliable enough not to misclassify exploratory conversations as approval gates.

## Falsifier

If later packets still contain obvious answer-evasion repairs or approval-surface reconstruction after the validator is active, the validator is too weak or too easy to bypass.

## Acceptance Criteria

- direct-question turns are blocked from sending unless the first sentence contains the direct answer
- approval-seeking turns are blocked from sending unless a contextual approval packet exists
- approval packets include changed-artifact context and links, not just filenames or section headers
- reply-mode classification is explicit enough that reviewers can tell why a given gate fired

