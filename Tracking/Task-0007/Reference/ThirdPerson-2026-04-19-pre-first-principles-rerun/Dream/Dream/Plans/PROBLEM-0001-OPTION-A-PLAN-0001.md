# Problem 0001 Option A Plan 0001

## Summary

Add stronger reminder text that regression closure must use the repo's human default lane.

## What Changes

- update shared closeout prompts
- add a short proof-surface warning to regression-related templates
- add one example that distinguishes supporting proof from closure proof

## Files Or Artifact Types

- shared orchestration prompt files for regression closeout
- shared testing prompt or checklist text

## Rollout

1. Find the closeout prompt and checklist files that talk about regression proof.
2. Add short language that the claimed lane must match the repo's default lane.
3. Add one example from ThirdPerson that shows why headless or operator proof is not enough.

## Success Checks

- future closeout prompts mention the default-lane rule
- reviewers can see the rule in the prompt text without looking elsewhere
- no new runtime validator or gate is required
