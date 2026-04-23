# PROBLEM-0006 OPTION-A: Task-Local Active-Constraints Ledger (Durable Before Proceeding)

## Title

Record new hard constraints and repo-truth corrections durably in task-owned current-state artifacts before any further work proceeds.

## Summary

The packet shows repeated "constraint drift":

- the human issues hard constraints (for example: "no engine mods") or corrects repo-local truth (for example: do not generalize ThirdPerson lane semantics into shared UE-generic prose)
- the system proceeds or plans as if the constraint did not exist, forcing repeated human restatement and course correction

This winner proposes a task-local active-constraints ledger that is updated immediately when a new constraint arrives, and is consulted before any further edits/builds/plan transitions.

The goal is not to promote constraints into shared memory. The goal is to prevent repeated re-learning and disallowed moves inside the current task.

## Writeup Type

Concrete implementation task (burden-reduction proposal).

## Source Event IDs

```json
[
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-3094",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-3197",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-3463",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-3482",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-3549",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-3585",
  "hie-20260419-019da198-ea2f-7421-8adc-c566da0b6121-4294",
  "hie-20260419-019da19b-19c6-7e20-bdd8-f03aea761fb9-4684"
]
```

## Burden Being Reduced

The human is currently forced to repeatedly do "constraint enforcement" work:

- restating constraints that should have been retained ("no engine mods", "stop at approval gate", "do not generalize repo policy into shared docs")
- correcting disallowed moves after the system already invested time pursuing them
- forcing the system to propose a durable fix rather than applying changes immediately

This is exported labor because the constraints are *governance*, not content: they are conditions under which work is allowed to proceed. If they are not recorded durably, the human must keep reasserting them.

## Current Truth

Today, tasks have two durable current-state surfaces:

- `Tracking/Task-<id>/HANDOFF.md` (human-readable current baseline and next-step guidance)
- `Tracking/Task-<id>/TASK-STATE.json` (machine-readable current orchestration state)

But there is no shared, stable place inside those task-owned artifacts where hard constraints are recorded as a "must read before acting" ledger.

As a result:

- constraints live in chat history
- future sessions can miss them or reinterpret them
- the human has to restate them repeatedly

## Target Truth

When the human issues a new hard constraint or repo-truth correction mid-task, the system must:

1. **Record it durably before proceeding** in the task’s current-state artifacts.
2. **Treat it as binding** for subsequent work until explicitly removed by the human.
3. **Prefer canonical references** for stable repo truths:
   - link to repo-root docs (e.g. ThirdPerson `REGRESSION.md`) rather than copying policy text into task state

The key property is: a later session should not need chat archaeology to know the current constraints.

## Causal Claim

If a task-local ledger is required and is updated immediately when constraints are issued (before further work), then:

- repeated constraint restatement decreases
- disallowed moves decrease because the constraint is visible and checked at the point of action
- later sessions can resume safely because the constraints are carried forward in the current-state artifacts

## Evidence

Evidence is concentrated in [../BURDEN-ANALYSIS.md](../BURDEN-ANALYSIS.md) (`BD-006`) and the stable constraint/truth events in [../../HumanInputEvents/INDEX.json](../../HumanInputEvents/INDEX.json), including:

- explicit prohibition on inventing a watcher/scheduler (scope boundary) (`...-3197`)
- explicit instruction to propose a durable fix rather than apply (`...-3463`, `...-3482`)
- explicit repo-local truth correction: remove generic orchestration UE language and avoid overly specific lane naming (`...-3549`, `...-3585`)
- explicit new hard constraint: "no engine mods" and a reopen of the planning gate (`...-4294`, `...-4684`)

## Why This Mechanism

This mechanism is chosen because it is the narrowest durable fix for mid-task constraint drift:

- Startup rereads help only for stable truths known at task start; the packet shows constraints arriving mid-task and requiring immediate capture.
- A shared global memory ledger is too broad and risks promoting task-specific constraints into permanent policy.

A task-local ledger keeps both truth and compassion:

- truth: what is allowed and disallowed is recorded explicitly
- compassion: it removes repeated human enforcement labor without blaming the human for setting constraints

## Scope Rationale

Task-local scope is earned:

- the constraints in the packet are task-scoped (for example: "no engine mods" for the current workstream)
- stable repo truths already have canonical homes in repo-root docs; the task should reference them, not rewrite them

## Goals

- Make new constraints durable before any further work proceeds.
- Reduce repeated human constraint restatement and rework caused by disallowed moves.
- Keep stable repo truths referenced canonically rather than copied into task state.

## Non-Goals

- Creating a shared cross-repo "constraints ledger" schema or memory system.
- Duplicating repo policy prose inside every task artifact.
- Solving STOP/ownership semantics (`P-002`) or approval packet shape (`P-004`).

## Implementation Home

- Task-owned current-state artifacts:
  - `Tracking/Task-<id>/HANDOFF.md`
  - `Tracking/Task-<id>/TASK-STATE.json`

## Implementation Home Rationale

- The burden is mid-task and task-scoped; the correct home is the task’s own current-state surfaces, not shared docs.
- `HANDOFF.md` is the human-readable "must read to resume" file; it is the right place for a visible ledger.
- `TASK-STATE.json` is the machine-readable state used to record durable transitions; it is the right place to record that a constraint has triggered a planning gate or a temporary blocked state until the ledger and plan are updated.

## Constraints And Baseline

- Preserve the purpose of `TASK-STATE.json` as compact orchestration state (do not turn it into a narrative dump).
- Do not extend the shared `TASK-STATE` schema in this task-local winner; use `HANDOFF.md` for the durable human-visible ledger and use existing JSON fields for gating transitions.

## Proposed Changes

These are the concrete, reviewable surfaces this winner changes.

1. **Add an explicit active-constraints ledger section to `HANDOFF.md`**
   - Artifact: `Tracking/Task-<id>/HANDOFF.md`
   - Add a stable section, for example `## Active Constraints (Must Read Before Acting)`, containing a small table with (at minimum):
     - `Constraint` (plain-language statement)
     - `Source event id` (when available)
     - `Status` (`active` or `cleared`)
     - `Applies to` (what actions it constrains: edits, engine changes, doc edits, approval gates)
     - `Enforcement / check` (how the task will ensure compliance)
     - `Canonical refs` (links to repo-root docs for stable truths instead of copied prose)
2. **Add a task-state gating rule for new constraints**
   - Artifact: `Tracking/Task-<id>/TASK-STATE.json`
   - When a new hard constraint arrives, require the lead to persist a durable state transition before proceeding:
     - set `status = "blocked"` (or equivalent) until the constraint is recorded in `HANDOFF.md` and any affected plan/gate is reevaluated
     - add a short `blockers[]` entry that points at the ledger section (for example: `"New hard constraint issued; update HANDOFF.md Active Constraints before proceeding."`)
     - once the ledger is updated and the plan/gate is reevaluated, clear the blocker and return the task to `in_progress`
3. **Add a "consult before acting" rule**
   - Artifact: `Tracking/Task-<id>/HANDOFF.md`
   - Add a short rule under the ledger:
     - before any edit/build/plan transition, reread the `Active Constraints` section and explicitly confirm (in the pass audit or in the handoff) that planned actions do not violate active constraints

## Acceptance Criteria

- When a new hard constraint is issued, it appears durably in `HANDOFF.md` under an `Active Constraints` ledger **before** further work proceeds.
- `TASK-STATE.json` reflects a durable "constraint arrived, reevaluation required" gating transition rather than silently continuing.
- Stable repo truths are referenced via canonical doc links (e.g. ThirdPerson `REGRESSION.md`) instead of copied policy prose inside the ledger.
- A later session can resume safely by reading `HANDOFF.md` and does not require chat archaeology to rediscover constraints.

## Expected Resolution

Human-facing outcome:

- The human no longer needs to repeat constraints to keep the task inside allowed boundaries.
- The system makes fewer disallowed moves because constraints are visible and consulted at the point of action.

## Human Relief If Successful

- Less repeated "no, don’t do that" or "re-evaluate" supervision.
- Less rework caused by pursuing disallowed directions.
- Faster, calmer collaboration because constraints are carried forward durably.

## Internal Mechanism Map

1. New constraint arrives.
2. Task is durably gated (state + handoff ledger update required).
3. Constraint is recorded in `HANDOFF.md` with source id and enforcement/check.
4. Plan/pass continues with explicit constraint consultation.

## Rival Explanations Considered

- "This is just a one-time misunderstanding."
  - Rejected: the packet includes multiple separate constraint/truth corrections; drift is recurrent.
- "Just reread repo docs at startup."
  - Insufficient: the packet shows constraints arriving mid-task; startup rereads do not capture those.

## Rival Mechanisms Considered

- `P-006 / Option B` (repo-local preflight loader for stable truth only):
  - Rejected as winner: it helps stable truths at task start but does not capture task-specific mid-task constraints like "no engine mods" or "stop at approval gate."
- Shared global constraints ledger:
  - Rejected: too broad; risks promoting task-specific constraints into permanent policy and creates a new shared substrate the packet did not justify.

## Tradeoffs

- Extra bookkeeping at constraint boundaries:
  - Intentional: it prevents larger rework and repeated human enforcement.
- Potential duplication:
  - Minimized by requiring canonical refs for stable repo truths instead of copying prose.

## Shared Substrate

- Task artifact roles and structure: [../../../../../../../../../AGENTS.md](../../../../../../../../../AGENTS.md)
- Shared task-state contract and schema (this task does not modify it): [../../../../../../../../TASK-STATE.md](../../../../../../../../TASK-STATE.md)
- ThirdPerson repo truth docs that should be referenced canonically when relevant:
  - [ThirdPerson/AGENTS.md](/c:/Agent/ThirdPerson/AGENTS.md)
  - [ThirdPerson/REGRESSION.md](/c:/Agent/ThirdPerson/REGRESSION.md)
  - [ThirdPerson/TESTING.md](/c:/Agent/ThirdPerson/TESTING.md)

## Not Solved Here

- Shared STOP/ownership semantics (`P-002`).
- Approval packet contract (`P-004`).
- Regression claim gates (`P-001`) and debugging gates (`P-005`).

## What Does Not Count

- Mentioning a constraint in chat without recording it durably.
- Recording constraints only in `TASK-STATE.json` blockers without a human-visible ledger in `HANDOFF.md`.
- Copying repo policy prose into the ledger instead of linking the canonical doc.

## Remaining Uncertainty

- The minimal ledger table fields that will remain useful across many tasks without growing into noisy chatter capture.
- The best convention for uniquely identifying constraints (simple short ids vs full-text only).

## Falsifier

This proposal is falsified if, after implementation:

- a new constraint is issued and later must be repeated because it was not recorded durably before work continued
- or a disallowed move still occurs despite the constraint being issued, because the current-state artifacts did not carry the constraint forward

## Proof Plan

1. Updating the shared exemplars is *not required* for this task-local winner, but implementers should update a real task’s `HANDOFF.md` with the new ledger section.
2. Rehearse two scenarios on a real task:
   - "new hard constraint" arrives (e.g. "no engine mods"): task state transitions to blocked, ledger updated, plan reevaluated, then unblocked
   - "repo-truth correction" arrives (e.g. "do not generalize UE wording"): ledger updated with canonical doc refs, and the next action explicitly confirms compliance
3. Verify in a later session that the constraints are visible and consulted without rereading chat history.

## Open Questions

- Should the ledger explicitly require a "cleared by" record when a constraint is removed, or is `Status: cleared` sufficient?
- Should tasks standardize a short "constraint consult" line in each pass audit to make compliance reviewable?

## References

- Burden driver: [../BURDEN-ANALYSIS.md](../BURDEN-ANALYSIS.md) (`BD-006`)
- Designed options: [../SOLUTION-DESIGN.md](../SOLUTION-DESIGN.md#p-006-durable-constraint-ledger)
- Frozen winner boundary: [../WINNER-SYNTHESIS.md](../WINNER-SYNTHESIS.md#w-006-task-local-active-constraints-ledger)
- Final matrix row: [../ORTHOGONAL-SOLUTIONS-MATRIX.md](../ORTHOGONAL-SOLUTIONS-MATRIX.md#p-006-durable-constraint-ledger)
