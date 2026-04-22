# Problem 0005 Option B Plan 0001

## Intent

Add timed self-resume ownership so long-running open passes wake themselves back up after quiet checkpoints.

## What Changes

- Set a resume timer when the task pauses without closing.
- Re-check the task after the timer expires.
- Continue automatically if the blocker cleared or if the pause was only a weak checkpoint.

## Files Or Artifact Types That Move

- Workflow timers or scheduler-backed resume records.
- Task-state fields for next resume time and resume reason.
- Logs or notes showing when self-resume fired.

## Rollout

1. Define which open states qualify for timed self-resume.
2. Store the timer in durable task state.
3. Pilot the mechanism on one pass with known restart friction.
4. Measure whether it reduces wake-up messages without creating noisy restarts.

## Success Check

- Long open pauses are revisited automatically.
- Quiet checkpoints do not silently turn into stalled work.
- Self-resume does not override real wait gates from the human.

## Burden Reduction Under Directional Context

`Truth`: it exposes silent stalls sooner.

`Compassion`: it can reduce some restart burden, though it does not fix bad stop decisions by itself.

`Tolerance`: it respects genuine wait instructions while treating idle drift as a recoverable failure mode.
