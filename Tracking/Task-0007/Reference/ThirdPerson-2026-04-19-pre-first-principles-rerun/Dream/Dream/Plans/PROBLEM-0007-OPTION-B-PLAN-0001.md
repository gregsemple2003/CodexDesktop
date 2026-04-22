# Problem 0007 Option B Plan 0001

## Summary

Add a checkpoint protocol that defines when active work continues, when durable state must refresh, and when a message is allowed.

## What Changes

- define durable milestone types for active passes
- require task state, pass state, or bug state to update at those milestones
- define a narrow list of allowed message reasons during active work: approval gate, true external block, or requested progress report
- add a resume rule for non-terminal failed attempts

## Files Or Artifact Types

- shared workflow docs
- active-work prompt bundle
- task-state or pass-checklist helpers

## Rollout

1. Define the milestone and message rules.
2. Tie them to current task-state artifacts and pass closeout artifacts.
3. Add one check that flags a message sent without an allowed reason.
4. Trial on a task that needs several iterations in one pass.

## Success Checks

- active work continues through the next durable milestone by default
- state refreshes stop lagging far behind real work
- unnecessary resume prompts become less common
