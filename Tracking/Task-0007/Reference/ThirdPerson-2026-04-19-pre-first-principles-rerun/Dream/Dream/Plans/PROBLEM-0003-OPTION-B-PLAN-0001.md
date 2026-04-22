# Problem 0003 Option B Plan 0001

## Summary

Add a boundary manifest that captures the live repo, lane, pass, ownership, and hard constraints for the current work.

## What Changes

- define a small manifest artifact for live boundaries
- populate it when a task starts or when the human adds a new hard rule
- require later plan, proof, and handoff steps to compare proposed actions against that manifest

## Files Or Artifact Types

- task-owned state or companion manifest schema
- orchestration prompt steps that read and refresh the manifest
- validator checks for out-of-bound actions

## Rollout

1. Define the minimum fields: repo, lane, pass, ownership, hard constraints, and prohibited surfaces.
2. Add update rules for mid-task corrections like `no engine mods`.
3. Add a check before editing, testing, or claiming closure.
4. Trial the manifest on a reopened ThirdPerson pass.

## Success Checks

- new hard constraints become durable state right away
- later actions can be rejected as out of bounds with a concrete field match
- pass rewrites and repo-crossing drift are easier to catch
