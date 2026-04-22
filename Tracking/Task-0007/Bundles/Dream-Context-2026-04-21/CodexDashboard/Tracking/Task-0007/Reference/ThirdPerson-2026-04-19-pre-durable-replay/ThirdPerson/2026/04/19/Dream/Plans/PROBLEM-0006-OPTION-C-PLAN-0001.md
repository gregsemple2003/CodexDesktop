# Problem 0006 Option C Plan 0001

## Planning Intent

This file turns Problem `0006`, Option `C. Hidden learned profile` from [ORTHOGONAL-SOLUTIONS-MATRIX.md](../ORTHOGONAL-SOLUTIONS-MATRIX.md) into a bounded implementation sequence.

It is an alternative route, not the selected winner task.

## Summary

Store repeated corrections in a local hidden profile that is injected into prompt context automatically, without requiring every lesson to become a tracked repo artifact.

## Fixed Defaults

- scope: local hidden state plus prompt bootstrap rules
- canonical homes:
  - `C:\Users\gregs\.codex\Orchestration\LEARNED-PROFILE.schema.json`
  - `C:\Users\gregs\.codex\state\learned-profile.json`
  - `C:\Users\gregs\.codex\.gitignore`
  - `C:\Users\gregs\.codex\SETUP.md`
  - `C:\Users\gregs\.codex\Orchestration\Prompts\TASK-HARVESTER.md`
  - `C:\Users\gregs\.codex\Orchestration\Prompts\TASK-LEADER.md`
- profile fields:
  - correction summary
  - confidence
  - scope
  - last reinforced at
  - decay or reset flag
- the profile file is local machine state and not user-authored canon

## Pass Plan

### Pass 0000 - Schema And Local-State Contract

Goal:

- define the hidden profile format and its backup rules

Build:

- add `LEARNED-PROFILE.schema.json`
- document `state/learned-profile.json` as ignored local state
- update `.gitignore` and `SETUP.md` with rebuild and reset rules

Unit Proof:

- the ignored path is documented in `SETUP.md`
- the schema names every field used by the local profile

Exit Bar:

- hidden learning has a defined format and honest backup story instead of silent machine state

### Pass 0001 - Profile Update Rules

Goal:

- define how repeated corrections enter or leave the hidden profile

Build:

- update the leader prompts with rules for adding a correction after repeated reinforcement
- define reset and decay conditions so stale lessons do not live forever
- add one example profile entry and one reset case

Unit Proof:

- the prompts reference the same profile fields as the schema
- reset and decay conditions are concrete and reviewable

Exit Bar:

- local hidden learning becomes bounded and reversible instead of opaque drift

### Pass 0002 - Prompt Injection

Goal:

- make the profile influence later work

Build:

- define how the relevant profile slices are injected into prompt context
- require that low-confidence or expired entries are not injected
- add one operator-visible reset path

Unit Proof:

- only relevant scoped entries are eligible for injection
- the reset path clears the profile deterministically

Exit Bar:

- repeated corrections can shape future behavior without manual doc promotion every time

## Testing Strategy

- keep the schema, ignore rules, and prompt references aligned
- reject any rollout that hides local state without documenting rebuild or reset behavior

## Deferred Work

- append-only ledger promotion
- cross-machine sync
- invisible model-side state with no local reset path
