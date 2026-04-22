# Problem 0006 Option B Plan 0001

## Summary

Add a first-disagreement debug worksheet that forces one measured bad state, its upstream writer chain, and explicit stop rules before any fix claim.

## What Changes

- define a debug worksheet with required fields for measured bad values, frame or timestamp, lane, and artifact refs
- require boundary-by-boundary tracing from runtime state to the writer that made it wrong
- require contradictory evidence to be preserved instead of hidden
- block `root cause` claims that stop at a category without a concrete seam

## Files Or Artifact Types

- shared debugging workflow templates
- task-owned bug and testing artifact templates
- optional schema or validator for worksheet completeness

## Rollout

1. Define the worksheet sections and required evidence fields.
2. Update debug prompts so reopened runtime bugs start there.
3. Add a review check that rejects category-only cause claims.
4. Trial on a ThirdPerson-style pose or grounding defect.

## Success Checks

- each root-cause claim names one measurable disagreement with values
- writer-chain tracing is explicit and boundary by boundary
- reviewers can tell where the claim is proven and where uncertainty remains
