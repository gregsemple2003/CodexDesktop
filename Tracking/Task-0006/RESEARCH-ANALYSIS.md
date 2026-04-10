# Task 0006 Research Analysis

## Problem 0001 Incident Contract Boundary And Why Chain

### Current Task Facts

`Task-0006` is intentionally not a general "make orchestration smarter" bucket. The trigger is narrower:

- a human-facing outcome diverged from intent
- the human had to step in and disagree explicitly
- that disagreement should become durable evidence

That means the contract must preserve both the human-facing miss and the orchestration context that let it through.

The user-level correction to the earlier examples exposed one more hard requirement: the record must preserve the grounded pre-correction state, including active course when the incident was about drift. Otherwise a good bug note can masquerade as an incident.

The rejected first plan exposed the real missing piece: a flat incident summary is not enough. But the earlier `goal_stack` shape over-corrected into too much machinery. The record still needs an upward explanation, but it should be `why_chains` that start from the concrete event and progressively generalize the target state the human was protecting.

One more constraint is now clear from the example review: the record should not wait for perfect diagnosis before it exists. A first-pass incident may be grounded and useful while still being too shallow on mechanism. That is acceptable as long as the process expects a second-pass root-cause refinement later.

### Recommended Minimum Record Shape

The record needs three kinds of fields from the start:

- human-facing fields:
  - expected state
  - actual outcome
  - why the gap mattered
- qualification fields:
  - grounded pre-correction state
  - human intervention summary and evidence
- orchestration fields:
  - evidence links
  - primary family
- `why_chains` fields:
  - one or more ordered linear rationale paths
  - a schema-constrained principle category for each entry
  - the target state being expressed at each entry

Without the qualification layer, the artifact becomes a bug note. Without the `why_chains`, it becomes an ungrounded complaint.

But the first pass should not also have to solve every downstream cause. The honest split is:

- first pass
  - get the event, intervention, and protected target state on disk
- second pass
  - refine the incident against the actual causal chain once code, tests, screenshots, or later investigation make that clearer

### Recommended Why Chains

Each chain should move from most concrete to most generalized:

- immediate target state in this incident
- the broader review or product target behind that event
- the more general design or workflow target
- the broadest human principle honestly expressed by the correction

Use multiple chains when the human expressed multiple sibling rationales that do not recursively answer each other.

The purpose is not to encode every possible hypothesis. The purpose is to explain, step by step, where the human is coming from.

### Qualification Rule

A qualified incident must be able to answer three falsifiable questions:

- What concrete state or active course existed before the human stepped in?
- Where is the direct evidence that the human corrected that course?
- What concrete event should have happened instead?

If any one of those is missing, the record should be stored as a bug, review note, or task input instead of a `Task-0006` incident.

## Problem 0002 Storage And Workflow Fit

### Current Repo Fit

The task is shared-orchestration work, but the evidence is still task-local and transcript-backed. The honest first step is:

- keep research, sample incidents, and task history under `Tracking/Task-0006/`
- only promote cross-repo contract or workflow rules into `C:\Users\gregs\.codex\Orchestration\` once the shape is proven locally

### Recommended First Storage Direction

Use task-owned markdown records for the seed set before inventing a new index or store. That keeps the first slice reviewable and cheap to refine.

If a normalized shared schema is added later, it should be informed by these task-local `why_chains` records instead of replacing them.

## Problem 0003 Seed Backfill And Coverage Standard

### Current Research Finding

A useful seed set should not be five versions of copy polish. The last five days of history already show at least five distinct miss families:

- truthful status communication
- information architecture and CTA framing
- UI terminology and scan-language mismatch
- producer or workflow misunderstanding
- proxy-proof or debugging misunderstanding

### Recommended Coverage Rule

The first five backfilled incidents should deliberately span:

- at least one usability or state-truth case
- at least one explicit UI semantics case
- at least one broader general misunderstanding
- at least one incident where the durable fix belongs in orchestration or workflow rather than the surface artifact alone

That mix is more valuable for the capture system than a prettier but narrower all-UI sample.

The same rule applies upstream: the seed set should not collapse into one kind of human principle either. It should test whether the `why_chains` can preserve multiple distinct kinds of correction:

- truthfulness
- directness
- reviewability
- boundedness
- operator clarity
- human-world interpretation

## Research Verdict

Task research is planning-ready.

## Recommended Direction

- define a narrow incident contract with explicit `why_chains` first
- backfill five recent incidents with intentional coverage across usability, UI, and general misunderstandings
- revisit accepted incidents for root-cause refinement so the corpus preserves mechanism as well as objection
- use those records to pressure-test the template before promoting any shared schema or workflow changes into `.codex`
