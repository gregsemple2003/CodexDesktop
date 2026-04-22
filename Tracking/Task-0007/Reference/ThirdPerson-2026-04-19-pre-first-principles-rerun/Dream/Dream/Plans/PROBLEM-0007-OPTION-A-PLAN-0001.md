# Problem 0007 Option A Plan 0001

## Summary

Add a reminder to keep going unless there is a real block or approval gate.

## What Changes

- update workflow prompts
- add examples of bad early-stop behavior
- remind workers to update durable state before messaging

## Files Or Artifact Types

- workflow prompts
- handoff or status guidance

## Rollout

1. Add the reminder to the active-work prompt.
2. Add one example from a reopened pass.
3. Trial it on a long-running task.

## Success Checks

- prompts say to continue through the next milestone
- workers see when to avoid unnecessary messages
- no concrete checkpoint protocol exists
