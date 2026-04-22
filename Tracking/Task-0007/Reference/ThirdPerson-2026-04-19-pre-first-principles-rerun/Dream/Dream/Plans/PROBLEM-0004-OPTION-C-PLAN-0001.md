# Problem 0004 Option C Plan 0001

## Summary

Attach full raw diffs to approval asks and let the reviewer inspect them directly.

## What Changes

- include the raw diff output for changed files
- add file links, but no extra summary layer
- rely on the reviewer to infer pass framing and decision scope

## Files Or Artifact Types

- approval request template
- diff export helper

## Rollout

1. Export raw diffs before approval asks.
2. Attach them with file links.
3. Trial on one task plan approval.

## Success Checks

- raw diffs are available every time
- links resolve to the changed files
- reviewers can still make sense of the approval request without a summary layer
