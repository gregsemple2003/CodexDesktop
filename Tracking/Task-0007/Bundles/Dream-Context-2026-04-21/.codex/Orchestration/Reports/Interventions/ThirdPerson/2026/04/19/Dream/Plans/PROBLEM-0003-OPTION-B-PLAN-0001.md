# Problem 0003 Option B Plan 0001

## Intent

Echo each new boundary back to the human and ask for confirmation before continuing.

## What Changes

- Summarize the newly heard boundary.
- Pause for human confirmation.
- Continue only after the confirmation lands.

## Files Or Artifact Types That Move

- Shared prompt templates for confirmation turns.
- Conversation examples that show boundary echo behavior.
- Optional state notes that record the confirmed boundary text.

## Rollout

1. Define the boundary types that require confirmation.
2. Write a one-line confirmation template.
3. Pilot it on pass, lane, and ownership changes.
4. Measure whether the extra confirmation cost is lower than the cost of drifting.

## Success Check

- New boundaries are rarely misheard.
- The confirmation text is short and specific.
- The added turn cost stays acceptable.

## Burden Reduction Under Directional Context

`Truth`: it reduces silent misreads.

`Compassion`: it avoids some later correction loops, though it still costs one turn up front.

`Tolerance`: it gives the human a chance to correct a summary before more work continues.
