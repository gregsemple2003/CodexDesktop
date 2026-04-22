# Problem 0001 Option C Plan 0001

## Summary

Require full runtime video capture for every regression closeout.

## What Changes

- make video capture mandatory for claimed regression passes
- store one canonical runtime recording beside each regression run
- update review guidance to prefer the recording over still images

## Files Or Artifact Types

- regression-run template
- automation capture scripts
- testing docs that define required media

## Rollout

1. Pick one recording format and stable output path.
2. Update regression instructions so closure requires a saved video.
3. Add review guidance for how to inspect the recording.

## Success Checks

- every claimed pass preserves a runtime video
- reviewers can replay the claimed lane afterward
- capture cost remains acceptable for frequent tasks
