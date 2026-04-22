# Problem 0007 Option C Plan 0001

## Summary

Build an automated heartbeat and resume queue service for long-running work.

## What Changes

- add a background service that tracks active passes
- auto-raise reminders when state has not advanced
- queue resume actions after non-terminal failures

## Files Or Artifact Types

- service code
- local runtime storage
- workflow integration points

## Rollout

1. Build the service.
2. Define the heartbeat cadence.
3. Integrate it with active tasks.

## Success Checks

- idle active passes become visible automatically
- resume prompts are generated without manual review
- the service does not create more noise than value
