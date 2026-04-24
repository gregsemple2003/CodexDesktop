# Task 0008

## Title

Build the backend task dispatch layer and durable execution-state contract.

## Summary

If the dashboard is going to dispatch and supervise work honestly later, it needs a real execution layer behind any UI.

This task is backend-only.

The first honest operator path can be Codex or direct backend interaction against that contract while frontend work waits for later tasks.

That layer must answer questions the current repo cannot answer durably enough yet:

- what exactly was dispatched
- what contract is the agent currently expected to fulfill
- what state is the run in now
- is the agent actively working, legitimately waiting, blocked, or asleep
- when is a human or supervisor allowed to poke the run
- how does interruption work
- what task, thread, or session provenance should be returned when an operator wants deeper context
- which exclusive repository checkout the run owns for execution
- which useful commit that owned checkout can be restored to for proof or cleanup

This task owns that backend layer.

No dashboard tab, launch button, or frontend control ships in this task.

The core product promise is not just `start a task`.
It is:

- dispatch work durably
- keep the state honest
- keep the human informed
- recover when a run drifts, stalls, or goes silent

The task must build on the Temporal-backed orchestration direction delivered by [Task-0005](../Task-0005/TASK.md), not bypass it with legacy scheduler shortcuts or purely local volatile UI state.

The runtime shape is frozen for this task:

- task dispatch will be modeled as a separate backend-owned task-run workflow and API contract under `backend/orchestration/`
- task runs are not Git-tracked recurring job specs
- recurring jobs may trigger task-run creation later, but the task-run model remains a separate runtime concept
- simple task execution in this task's scope will use a backend-owned exclusive repository checkout or equivalent isolated repo lane rather than the human's shared primary worktree

## Writeup Type

Concrete implementation task.

## Burden Being Reduced

The human cannot supervise or recover dispatched work cheaply enough when the system lacks a durable task-run contract.

Without that contract, the human has to guess:

- whether silence is healthy
- whether a run is blocked or sleeping
- whether a poke is justified
- what deeper context should open for recovery

The deeper burden is exported supervision work.

The human becomes the missing execution memory because the runtime cannot explain itself durably enough.

## Current Truth

The repo has a real backend and Temporal foundation from [Task-0005](../Task-0005/TASK.md), but it does not yet have a task-specific dispatch and supervision layer with a durable execution-state vocabulary.

That means current truth is split:

- tasks exist durably in markdown
- jobs exist durably in the backend
- task dispatch and supervision do not yet exist as a first-class durable runtime concept

## Target Truth

The target truth is a backend-owned task-run system that can durably answer:

- what was dispatched
- what state it is in
- what it is waiting on
- when silence is suspicious
- whether a poke or interrupt is justified
- what thread or context should open for deeper recovery
- which owned repo checkout the run is allowed to mutate
- which commit the system can restore that owned checkout to when proof or cleanup needs a known-good baseline

## Causal Claim

If the backend gains a durable task-run contract with:

- explicit state vocabulary
- wait reasons
- progress freshness
- poke rules
- interrupt rules
- launch provenance
- exclusive repo-checkout ownership
- explicit restore-commit semantics

then the dashboard can supervise tasks honestly instead of inferring reality from partial or transient clues.

## Evidence

- [Task-0005](../Task-0005/TASK.md) already proved that this repo wants durable runtime truth on the backend side rather than local scheduler hacks
- [Task-0009](../Task-0009/TASK.md) defines a future `Tasks` surface that explicitly needs sleeping detection, waiting semantics, and durable backend readback
- the task-local contract note at [Design/DURABLE-EXECUTION-STATE-CONTRACT.md](./Design/DURABLE-EXECUTION-STATE-CONTRACT.md) makes the current missing state semantics explicit
- the repo currently lacks a backend-owned execution checkout model that would let simple task runs mutate and reset code safely without leaning on the human's shared worktree

## Why This Mechanism

The right first fix is a durable execution-state and task-run model, not a UI-only workaround and not a weaker reminder-driven process.

This mechanism is chosen because the exported burden is structural:

- the system cannot supervise what it does not represent durably

## Scope Rationale

This task intentionally focuses on the backend dispatch and supervision substrate, not the human-facing `Tasks` tab and not Dream promotion flow.

That split is necessary because:

- [Task-0009](../Task-0009/TASK.md) owns humane surface design
- [Task-0010](../Task-0010/TASK.md) owns daily Dream intake and candidate promotion

Keeping the runtime separate prevents the `Tasks` UI or any Codex-side operator flow from becoming the accidental source of truth for dispatch state.

This task can be validated through backend APIs and Codex-driven backend interactions before any frontend wiring exists.

## Human Relief If Successful

If this task succeeds, the human should gain:

- safer dispatch
- lower-cost recovery from quiet tasks
- honest sleeping versus waiting distinctions
- durable backend readback strong enough that Codex or a later UI can recover the right context without guesswork
- a simple-task runtime that does not require risky cleanup inside the human's live working tree

## Remaining Uncertainty

- the precise freshness thresholds for sleeping detection still need proof
- the final context-readback format may evolve with implementation
- the exact isolated-repo provisioning mechanism can still be implementation-specific as long as exclusive ownership and restore-commit semantics stay explicit

## Falsifier

This task should be considered wrong or incomplete if, after implementation:

- the system still cannot distinguish waiting from sleeping durably
- a `poke` action exists without a durable justification rule
- Codex or a later client still cannot recover the right task context from backend provenance
- task-run truth still lives primarily in client memory instead of backend durability
- simple task execution still depends on mutating a shared human worktree or cannot restore its owned checkout to a known-good commit for proof and cleanup

## Internal Mechanism Map

### Mechanism 1: Durable Task-Run Persistence

Failure reduced:

- dispatched work has no durable identity or recoverable state

Mechanism:

- create and store task-run records with task identity and lifecycle state

### Mechanism 2: Wait Contract

Failure reduced:

- silence cannot be distinguished from legitimate waiting

Mechanism:

- require named wait reasons, resume conditions, and ownership

### Mechanism 3: Supervision And Poke

Failure reduced:

- silent runs drift indefinitely because the system cannot justify intervention

Mechanism:

- classify sleeping or stalled runs from durable contract and progress evidence, then allow poke when justified

### Mechanism 4: Interrupt And Context Provenance

Failure reduced:

- recovery and control actions are unsafe or context-poor

Mechanism:

- persist interrupt state and deep-context provenance alongside the run

### Mechanism 5: Exclusive Repo Ownership And Restore Baseline

Failure reduced:

- simple task execution collides with the human's shared worktree or leaves cleanup ambiguous after unit proof

Mechanism:

- give each simple execution run an exclusive backend-owned repo checkout and record the useful restore commits it may reset to during proof or cleanup

## Goals

- Define a durable execution-state contract that can represent task runs honestly enough for supervision and recovery.
- Distinguish these states durably:
  - queued
  - dispatching
  - running
  - waiting_for_human
  - blocked
  - sleeping_or_stalled
  - interrupted
  - completed
  - failed
- Preserve enough state that the system can decide when a run deserves a poke rather than assuming silence is acceptable.
- Let the human interrupt task runs intentionally and see the result reflected durably.
- Preserve task, thread, and session provenance strongly enough that Codex or later clients can recover the deeper working context from backend readback.
- Give the backend an exclusive repo checkout model for simple task execution so the runtime does not rely on the human's shared primary worktree.
- Preserve enough git baseline information that a run can reset its owned checkout to a known-good useful commit during unit proof or execution cleanup.
- Build on the existing Temporal and backend foundation from [Task-0005](../Task-0005/TASK.md).
- Keep the dispatch layer separate from every client surface so the backend remains the source of truth.
- Leave behind a promotable contract for durable execution state if the task-local design proves stable.

## Non-Goals

- Shipping the whole `Tasks` tab UI in this task.
- Shipping any frontend dispatch controls or dashboard-side launch actions in this task.
- Implementing daily Dream generation or digest email in this task.
- Replacing task-owned markdown as the durable source of task scope and acceptance.
- Treating `no output yet` as a sufficient state model.
- Falling back to legacy Windows Scheduled Tasks for normal dispatch behavior.
- Building a generic agent framework beyond what CodexDashboard needs for honest task dispatch and supervision.
- Building a multi-tenant repo-farm product or generalized remote CI system beyond the exclusive local checkout model needed for simple task execution here.

## Rival Explanations Considered

- `The real problem is only that the UI is missing a better status label set.`
  - rejected because the system cannot honestly expose sleeping, waiting, or poke semantics without durable runtime truth
- `The real problem is only weak task docs.`
  - rejected because even perfect task docs do not create live task-run state or provenance

## Rival Mechanisms Considered

- keep dispatch state only inside the dashboard process
  - rejected because it would be fragile and not durable enough for supervision
- reuse the jobs model directly with no task-run-specific contract
  - rejected because jobs and long-running task supervision have different human semantics and state needs
- run simple task execution directly inside the human's current working tree
  - rejected because shared-worktree mutation and cleanup would be too ambiguous for honest dispatch, proof, and rollback

## Constraints And Baseline

- [Task-0005](../Task-0005/TASK.md) already established the intended backend model:
  - Git-tracked desired state
  - Temporal runtime durability
  - dashboard as client, not scheduler host
- The new dispatch layer should reuse that model rather than inventing a separate runtime lane.
- The first operator path may be Codex or direct backend interaction:
  - this task should not require dashboard UI work to prove the contract
- For simple execution in this task's scope, the backend must own the repo lane it mutates:
  - do not assume the human's current checkout is safe to reuse
  - do not rely on shared dirty-worktree cleanup as the normal execution model
- The contract must record at least one known-good restore commit for an owned checkout:
  - the dispatch baseline commit
  - and, when useful, a later explicitly approved restore target created during the run
- The execution-state contract must be durable enough that a sleeping task can be recognized by comparing:
  - last durable progress
  - expected contract
  - current wait reason if any
- Tasks can be interrupted.
- A sleeping run without a matching durable wait contract should be eligible for a poke or recovery action.
- The backend should preserve context provenance strongly enough that later clients can open the right working context without manual search.
- Unit proof and execution cleanup must be able to reset the owned checkout to a useful recorded commit without hand-wavy `clean up afterward` semantics.

## Tradeoffs

- A richer durable state model improves supervision but adds implementation complexity that must stay bounded
- Aggressive sleeping detection reduces silent failure but risks false positives if wait semantics are weak
- Strong context provenance improves recovery but requires careful treatment of local paths and runtime context
- Exclusive owned checkouts and reset semantics improve safety but add repo-provisioning and git-state bookkeeping that must stay bounded to simple task execution

## Shared Substrate

- [Task-0005](../Task-0005/TASK.md) backend and Temporal foundation
- [Task-0009](../Task-0009/TASK.md) task-surface consumer
- [Task-0010](../Task-0010/TASK.md) future promoted-task source
- any future shared `.codex` contract if this task-local model proves durable across repos

## Not Solved Here

- the `Tasks` tab human-facing layout and copy
- frontend dispatch controls or dashboard launch actions
- daily Dream digest rendering
- candidate-task promotion into real task skeletons

## Expected Resolution

- A durable execution-state contract exists and is specific enough to power:
  - task supervision
  - poke decisions
  - interrupt handling
  - backend readback for later task-surface summaries
  - backend context recovery
- The backend can create and track task dispatch runs independently of the UI.
- The backend can provision or bind an exclusive repo checkout for simple task execution and restore it to a recorded useful commit when proof or cleanup requires it.
- Codex and later client surfaces can consume a stable API rather than inferring task-run state from transient local guesses.
- The task leaves behind enough design truth that sleeping-task recovery becomes a real product behavior, not a vague aspiration.

## What Does Not Count

- A `running` label with no durable contract behind it.
- A system that cannot distinguish `waiting intentionally` from `fell asleep`.
- A `poke` action that exists in the UI but has no grounded decision rule.
- A context-recovery path that still makes the human search for the relevant session manually.
- A dispatch system that depends on local in-memory state inside a client process.
- A runtime that mutates the human's shared primary worktree for ordinary simple-task execution.
- Cleanup wording that says the repo can be restored later without naming the recorded commit the owned checkout will reset to.

## Implementation Home

Primary product home:

- `backend/orchestration/`

Task-owned contract material:

- `Tracking/Task-0008/`

Relevant current implementation anchors:

- [backend/orchestration/README.md](../../backend/orchestration/README.md)
- [Task-0005](../Task-0005/TASK.md)

## Implementation Home Rationale

This belongs in `backend/orchestration/` because the key intervention boundary is runtime truth, not presentation.

Codex or later dashboard surfaces must consume this truth, but they should not own it.

The first honest slice should prove the backend contract directly. Frontend integration can follow later under [Task-0009](../Task-0009/TASK.md).

## Proposed Changes

- add a separate backend-owned task-run workflow and persistence model under `backend/orchestration/` rather than extending Git-tracked recurring job specs directly
- add an exclusive backend-owned repository-checkout model for simple task execution under `backend/orchestration/` so runs do not depend on the human's shared primary worktree
- persist run identity, task identity, status, wait reason, progress freshness, interrupt reason, and launch provenance for each task run
- persist repo-execution fields such as:
  - owned checkout identity
  - repository root
  - baseline commit
  - current commit
  - allowed restore commit or commits for proof and cleanup
- add backend endpoints for:
  - `POST /api/v1/tasks/{task_id}/dispatch`
  - `GET /api/v1/tasks`
  - `GET /api/v1/tasks/{task_id}`
  - `GET /api/v1/task-runs/{run_id}`
  - `POST /api/v1/task-runs/{run_id}/poke`
  - `POST /api/v1/task-runs/{run_id}/interrupt`
- expose context readback strong enough for Codex or future clients to recover the right task or thread context without reconstructing provenance manually
- expose enough repo-ownership and restore-baseline readback that unit proof and later operators can tell which checkout is safe to reset and which commit it will restore to

## Acceptance Criteria

- The backend can create a durable task run through `POST /api/v1/tasks/{task_id}/dispatch`.
- The created task run persists:
  - task identity
  - run identity
  - status
  - wait contract
  - last meaningful progress
  - interrupt state
  - deep-context provenance
- For simple execution, the created task run also persists:
  - an exclusive owned checkout or equivalent isolated repo lane
  - the baseline commit captured at dispatch
  - the current commit when it changes materially
  - the useful restore commit or commits the runtime may reset to during proof or cleanup
- `GET /api/v1/task-runs/{run_id}` returns enough state to distinguish:
  - running
  - waiting_for_human
  - blocked
  - sleeping_or_stalled
  - interrupted
  - completed
  - failed
- `POST /api/v1/task-runs/{run_id}/poke` is rejected when the durable wait contract says silence is legitimate and accepted when a run is sleeping or stalled under the task rules.
- `POST /api/v1/task-runs/{run_id}/interrupt` records interrupted state durably and leaves the run recoverable for later review.
- The backend exposes context provenance strong enough for Codex or a future UI to recover the right task or thread context.
- The backend does not require ordinary simple-task execution to mutate the human's shared primary worktree.
- Unit proof and execution cleanup can reset the owned checkout to a recorded useful commit baseline without ambiguous manual cleanup steps.
- Focused automated tests exist for task-run state transitions, wait semantics, sleeping detection, poke gating, and interrupt handling.
- Real proof exists for one dispatch, one legitimate wait, one sleeping detection, one poke, and one interrupt path through direct backend interaction or Codex-driven backend calls.

## Proof Plan

- add focused tests for task-run state transitions
- add focused tests for wait-contract and sleeping-detection rules
- add focused tests for owned-checkout selection and restore-commit reset behavior
- capture real proof of:
  - a task dispatch
  - a task entering a durable wait state
  - a justified poke
  - an interrupt
  - backend context readback from durable provenance
  - owned-checkout reset to a recorded useful commit during proof or cleanup
- leave frontend proof to later tasks once [Task-0009](../Task-0009/TASK.md) consumes the backend contract

## Open Questions

- What heartbeat or progress evidence is strong enough to distinguish healthy silence from sleeping?
- Which parts of context provenance should be generic and which parts should remain CodexDashboard-specific?
- How should interrupt semantics differ between:
  - queued
  - dispatching
  - running
  - sleeping
  - waiting_for_human
- Which repo-isolation mechanism is simplest in this repo while still giving Task-0008 exclusive ownership and reliable restore-commit resets for simple tasks?

## References

- [Task-0005](../Task-0005/TASK.md)
- [Task-0009](../Task-0009/TASK.md)
- [Task-0010](../Task-0010/TASK.md)
- [Design/DURABLE-EXECUTION-STATE-CONTRACT.md](./Design/DURABLE-EXECUTION-STATE-CONTRACT.md)
- [backend/orchestration/README.md](../../backend/orchestration/README.md)
- [TASK-STATE.md](../../../../Users/gregs/.codex/Orchestration/TASK-STATE.md)
