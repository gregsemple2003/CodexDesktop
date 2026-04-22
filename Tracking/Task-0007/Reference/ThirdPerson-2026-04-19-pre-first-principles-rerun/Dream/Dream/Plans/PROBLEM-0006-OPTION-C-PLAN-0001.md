# Problem 0006 Option C Plan 0001

## Summary

Build always-on deep runtime instrumentation before future root-cause work begins.

## What Changes

- capture richer pose, transform, and writer-state traces on every run
- save them in a stable artifact path
- build views for fast seam inspection

## Files Or Artifact Types

- repo-local instrumentation code
- capture scripts
- artifact viewer or summarizer

## Rollout

1. Define the minimum trace set.
2. Implement capture and storage.
3. Add a review surface for the traces.

## Success Checks

- deep traces exist for every target run
- debugging time drops on later pose defects
- capture overhead stays acceptable on real lanes
