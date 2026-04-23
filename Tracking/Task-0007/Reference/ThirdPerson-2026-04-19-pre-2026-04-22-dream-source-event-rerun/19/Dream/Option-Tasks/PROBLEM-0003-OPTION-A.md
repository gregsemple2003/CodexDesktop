# Title

Add prompt reminders for answer-first and approval-ready replies

## Summary

Tighten prompt wording so direct questions are answered first and approval requests include links and context more often.

## Source Event IDs

```json
[
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-3392",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4379",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4399",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4512",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5001",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5090",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5281",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5290"
]
```

## Burden Being Reduced

The human should not have to restate the question, ask for smaller words, or ask for diffs and links after already making the needed reply shape obvious.

## Causal Claim

The current prompts do not enforce a reply mode that distinguishes direct-question turns from approval turns.

## Evidence

- [../BURDEN-ANALYSIS.md](../BURDEN-ANALYSIS.md) identifies human-unusable reply shaping as its own burden driver.
- [../SessionExcerpts/INDEX.json](../SessionExcerpts/INDEX.json) keeps both the approval-surface and direct-answer repair sequences.

## Goals

- make prompts say "answer first" more explicitly
- remind approval requests to include contextual links and changed artifacts

## Non-Goals

- adding validators
- adding a durable approval artifact
- changing stop-state or proof-state behavior

## Implementation Home

- shared orchestration prompts that govern worker replies and approval asks

## Proposed Changes

- add direct-question examples that force the first sentence to answer
- add approval examples that include links and changed-artifact context

## Expected Resolution

Some low-effort reply-shape misses disappear.

## What Does Not Count

- adding examples without clear mode-switch language
- reminders that still let the model choose whether to answer first

## Why This Mechanism

It is the fastest possible intervention for reply-shape failures.

## Human Relief If Successful

The human sees more replies in the requested shape even before stronger tooling exists.

## Remaining Uncertainty

Reminder language is exactly what failed repeatedly in this packet.

## Falsifier

If later packets still contain obvious "I asked you a question" or "give me links" repairs, this option is insufficient.

## Acceptance Criteria

- relevant prompts include answer-first wording for direct questions
- relevant prompts include approval-ready wording for plan and pass review asks
- examples in the prompt bundle distinguish the two reply modes

