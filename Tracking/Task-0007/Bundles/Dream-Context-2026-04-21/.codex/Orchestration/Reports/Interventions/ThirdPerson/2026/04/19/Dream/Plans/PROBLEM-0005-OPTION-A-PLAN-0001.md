# Problem 0005 Option A Plan 0001

## Intent

Add a stop-eligibility gate so open work does not stop while the next justified action is known and no real external block exists.

## What Changes

- Evaluate every proposed stop against open-task status, known next action, and external blockers.
- Block checkpoint-style stops that would only force the human to wake the work back up.
- Require an explicit blocker record when stopping before the human's done bar.

## Files Or Artifact Types That Move

- Shared workflow rules for milestone, pause, and stop behavior.
- Task-state fields for `next_action`, `blocking_reason`, and stop eligibility.
- Pass closeout or handoff artifacts that record why a stop was valid.

## Rollout

1. Define the stop-eligibility test.
2. Add it before milestone summaries and pause decisions.
3. Make blocker recording mandatory when the gate says the task cannot honestly stop cleanly.
4. Pilot it on task flows with measured wake-up loss like the `ThirdPerson` packet.
5. Promote the rule after it proves it reduces resume-prod events.

## Success Check

- Open work with a known next action continues without a human wake-up.
- Stop events that remain all have explicit real blockers.
- Resume-prod messages and charged stall loss decline on similar tasks.

## Burden Reduction Under Directional Context

`Truth`: it stops fake completion and fake stop points.

`Compassion`: it directly targets the packet's largest measured stall cost.

`Tolerance`: it treats wake-up messages as failure telemetry, not as normal coordination traffic.
