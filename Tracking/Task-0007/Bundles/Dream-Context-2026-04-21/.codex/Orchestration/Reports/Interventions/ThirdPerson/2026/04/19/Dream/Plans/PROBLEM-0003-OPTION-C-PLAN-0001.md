# Problem 0003 Option C Plan 0001

## Intent

Reduce boundary drift by splitting work into smaller batches so fewer active rules must be held at once.

## What Changes

- Break larger passes into narrower sub-steps.
- Limit each step to one lane, one method, and one closeout target.
- Refresh the active constraints at the start of each batch.

## Files Or Artifact Types That Move

- Task plans and pass definitions that control batch size.
- Milestone notes that define the current narrow objective.
- Optional session scaffolds for active constraints.

## Rollout

1. Identify where `ThirdPerson` work accumulated too many active boundaries at once.
2. Redraw those flows into smaller bounded batches.
3. Add explicit batch-start notes that restate the active constraints.
4. Compare drift rates before and after the narrower batching.

## Success Check

- Active work units carry fewer simultaneous constraints.
- Boundary misses drop on long-running tasks.
- Batch overhead does not outweigh the drift reduction.

## Burden Reduction Under Directional Context

`Truth`: smaller work units make it easier to stay on the right lane.

`Compassion`: some repeat corrections should drop because less state is active at once.

`Tolerance`: it gives the system a simpler frame for handling partial human corrections.
