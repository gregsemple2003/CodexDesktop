# Problem 0004 Option B Plan 0001

## Summary

Auto-build a human approval packet that shows change shape, links, pass framing, and the exact decision needed.

## What Changes

- generate a short approval summary before approval is requested
- list touched files with context links
- summarize the change in plain language and name the decision the human is making
- include pass id and whether the work is new, reopened, or audit-only

## Files Or Artifact Types

- approval packet template
- helper that gathers changed files and short summaries
- pass or handoff prompts that require the packet

## Rollout

1. Define the approval packet sections.
2. Teach the workflow to generate the packet before approval asks.
3. Link each file with enough context for fast review.
4. Trial it on a plan approval case with a reopened pass.

## Success Checks

- approval asks come with a small review packet every time
- the packet names what changed and what needs approval
- a human can make the decision without hunting through unrelated files
