# Orthogonal Solutions Matrix

## Burden-To-Problem Mapping

- `Burden Driver 0001` -> `PROBLEM-0001` (`separate_problem_row`)
- `Burden Driver 0002` -> `PROBLEM-0003` (`separate_problem_row`)
- `Burden Driver 0003` -> `PROBLEM-0002` (`separate_problem_row`)
- `Burden Driver 0004` -> `PROBLEM-0004` (`separate_problem_row`)

No downstream merge was taken. The burden drivers differ in mechanism boundary, enforcement boundary, and falsifier, so merging them would hide distinct exported human work.

## PROBLEM-0001: Keep Regression Closure Pinned To The Human Default Runtime Lane

Source Burden Drivers: `Burden Driver 0001`

Source Event IDs:

```json
[
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-3325",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-3431",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-3473",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-3618",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-3702",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-3814",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4114",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-4202",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4909",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5100",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-6105",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-6451",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-6558",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-6755",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-7446"
]
```

Mechanism Boundary: the repo-local regression proof package that emits and validates closing evidence for the human default runtime lane

Acceptance Test: a would-be closing proof run emits lane metadata, runtime-only artifact declarations, and full-body runtime captures from the default lane; closeout artifacts refuse to treat off-lane or non-runtime material as final proof

Falsifier: a future closure can still pass while relying on an automation-only surrogate, cropped evidence, or non-runtime images for a runtime claim

Options:

- Option A: [Require explicit default-lane proof manifests before regression closure](./Option-Tasks/PROBLEM-0001-OPTION-A.md)
- Option B: [Generate default-lane runtime proof packets from GameAutomation](./Option-Tasks/PROBLEM-0001-OPTION-B.md)

Winner: `Option B`

Why Winner:

`Option A` improves labeling, but it still depends on a human noticing missing or invalid evidence after the run. `Option B` moves the truth boundary earlier by making the runtime lane emit the required proof packet itself, then lets downstream artifacts consume that packet.

## PROBLEM-0002: Make Ownership, Stop State, And New Constraints Durable

Source Burden Drivers: `Burden Driver 0003`

Source Event IDs:

```json
[
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-3447",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-3692",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4015",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-4432",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4294",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-4684",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-4755",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-4919",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5108",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5272",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5300",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5310",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-6558",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-7038"
]
```

Mechanism Boundary: the durable task-state contract that a worker must update before pausing, resuming, handing back, or declaring a checkpoint stop

Acceptance Test: task state carries explicit ownership, stop state, resume trigger, and active constraints; workers cannot stop or hand intermediate failure back without updating those fields, and later workers can continue without the human restating the same constraints

Falsifier: the human still has to restate pause gates, no-engine-mods, pass structure, or "continue under your existing ownership" after those facts were already established

Options:

- Option A: [Add prompt-level continuity and constraint reminders](./Option-Tasks/PROBLEM-0002-OPTION-A.md)
- Option B: [Make stop state, ownership, and active constraints first-class task state](./Option-Tasks/PROBLEM-0002-OPTION-B.md)

Winner: `Option B`

Why Winner:

The packet already shows that reminders did not survive across turns, passes, or idle time. Durable state is the smaller mechanism that can honestly explain both the repeated wake-ups and the repeated constraint restatements.

## PROBLEM-0003: Make Replies Directly Usable For Questions And Approval Gates

Source Burden Drivers: `Burden Driver 0002`

Source Event IDs:

```json
[
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-3392",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4379",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4399",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4421",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4512",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4522",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4571",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5001",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5010",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5020",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5030",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5073",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5090",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5281",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5290"
]
```

Mechanism Boundary: the pre-send reply-shaping layer for human-facing messages

Acceptance Test: direct-question turns answer in the first sentence and in the requested shape; approval-seeking turns attach a contextual approval packet with changed artifacts, links, and pass framing instead of forcing the human to reconstruct them

Falsifier: the human still has to ask "I asked you a question", "give me links", or "where are the new passes" after the reply-shaping layer is applied

Options:

- Option A: [Add prompt reminders for answer-first and approval-ready replies](./Option-Tasks/PROBLEM-0003-OPTION-A.md)
- Option B: [Gate replies with direct-answer and approval-packet validators](./Option-Tasks/PROBLEM-0003-OPTION-B.md)

Winner: `Option B`

Why Winner:

The packet already contains enough reminder language. What is missing is enforcement at the reply boundary and a durable approval artifact the human can actually use.

## PROBLEM-0004: Force Runtime Defects Through First-Disagreement Debugging

Source Burden Drivers: `Burden Driver 0004`

Source Event IDs:

```json
[
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4114",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4768",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5119",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5129",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5154",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-5340",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-6105",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-6755",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-7038",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-7446"
]
```

Mechanism Boundary: the bug/debug packet and runtime probe layer that captures the first disagreement and its writer chain

Acceptance Test: a runtime defect cannot be closed or reframed as "improved" unless the bug artifact names the first disagreement with values, preserves contradictory evidence, and traces the relevant writer chain or explicitly records the remaining unresolved branch

Falsifier: runtime defects can still advance through partial retunes or aesthetic language without a first-disagreement section and without concrete state values

Options:

- Option A: [Strengthen debugging prompts around first-disagreement tracing](./Option-Tasks/PROBLEM-0004-OPTION-A.md)
- Option B: [Require first-disagreement bug packets and runtime disagreement probes](./Option-Tasks/PROBLEM-0004-OPTION-B.md)

Winner: `Option B`

Why Winner:

The packet only becomes clearer when the human explicitly forces the debugging method. `Option B` makes that method durable in artifacts and probes instead of hoping a prompt reminder survives long enough to matter.

## Winner Rollout Order

1. `PROBLEM-0001 / Option B`
   - trust and closure truth fail first if the runtime lane cannot emit a valid proof packet
2. `PROBLEM-0002 / Option B`
   - without durable ownership and stop state, the same work can still fall idle or shed constraints between passes
3. `PROBLEM-0003 / Option B`
   - once proof and state are durable, human-facing replies need to surface them in usable form
4. `PROBLEM-0004 / Option B`
   - this should land after the proof package exists, because the debugging packet should reuse the same runtime-lane artifact roots

## Shared Substrate

These winners intentionally share substrate, but they do not claim the same outcome:

- shared `.codex` process and prompt docs
- repo-local task state and task artifacts in `Tracking/Task-<id>/`
- repo-local GameAutomation runtime capture surfaces

Shared substrate is why the rollout can be staged efficiently. It is not a merge justification.

