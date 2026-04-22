# Problem 0003 Option C Plan 0001

## Summary

Route ThirdPerson work through a fixed specialty prompt lane with preloaded boundaries.

## What Changes

- create a ThirdPerson-specific prompt bundle
- preload repo-local rules and common constraints
- require users to invoke that lane for ThirdPerson work

## Files Or Artifact Types

- specialty prompt files
- repo-specific onboarding or launch docs

## Rollout

1. Gather the stable ThirdPerson rules.
2. Build a dedicated prompt lane around them.
3. Test it on one task reopen.

## Success Checks

- ThirdPerson runs start with the right repo rules in view
- repeated lane mistakes fall on specialized prompts instead of generic ones
- the approach remains maintainable as repo rules change
