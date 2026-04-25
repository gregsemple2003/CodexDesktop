# Task 0008 Constraints

This file records explicit human-provided or human-approved constraints that should remain visible throughout implementation.

Add new constraints here as they are given.

## Active Constraints

### 2026-04-24 Implementation Approval

- `PLAN.md` is approved.
- Start implementation immediately.
- Do not stop until the task is done or a real blocker is reached.
- Keep durable task state current at each lifecycle transition.

### 2026-04-24 Runtime Truth And Ingest

- Latest git task docs remain declared task or process truth.
- The backend must preserve rollback-safe runtime truth instead of forgetting waits, interrupts, cleanup, or divergence after git rewinds.
- Day-to-day sync should use task-API read-through reconcile, not a general file watcher.
- Direct agent HTTP calls are not the normal primary truth mechanism.

### 2026-04-24 Execution Lane Safety

- Simple task execution must use an exclusive backend-owned checkout or equivalent isolated repo lane.
- Do not use the human's shared primary worktree as the normal simple-execution lane.
- Restore or cleanup semantics must target a recorded useful commit baseline or later approved restore target.

### 2026-04-24 Human-Facing Contract

- Use `meaning_summary` for the task-level summary of why the task exists or why the human should care.
- Use `state_summary` for the current run or state summary.
- `completed` must mean no further human judgment, review, or approval is required.
- If review or approval is still required, the state must surface as `waiting_for_human` with an explicit approval target.

### 2026-04-24 Constraint Memory

- Write a markdown file `CONSTRAINTS.md` and put that in there, along with any other constraint i give you.

### 2026-04-24 Continuous Ownership

- Plan approved. Start implementation immediately. Do not stop until blocked or done.

### 2026-04-24 Bounded PASS-0002 Slice

- Keep the current slice bounded.
- Prefer explicit interrupt-review decision and resolution behavior over widening scope into real worker-side execution unless execution becomes the smaller honest next step.
- Do not widen scope unnecessarily.

### 2026-04-24 Validation Replay Constraint

- When the fixed active task-run workflow id is reused after workflow-shape changes, reset the disposable validation Temporal volume before live proof so replay failures do not masquerade as task logic regressions.

### 2026-04-24 Validation Warm-Up Constraint

- After resetting the disposable validation Temporal volume, allow a short Temporal warm-up window before launching a clean manual listener or Temporal client init can fail with `error reading server preface: EOF`.

### 2026-04-24 Validation Port Override Constraint

- When starting the disposable validation compose stack directly from `backend/orchestration`, set the validation-lane port overrides explicitly so it does not collide with the service-lane bindings:
  - `CODEX_ORCH_POSTGRES_PORT=15432`
  - `CODEX_ORCH_TEMPORAL_PORT=17233`
  - `CODEX_ORCH_TEMPORAL_UI_PORT=18080`

### 2026-04-24 Owned-Lane Mutation Baseline Constraint

- Keep Task-0008-owned mutation recipes aligned with the current repo baseline; if an owned-lane patch recipe drifts behind the real implementation, the owned-lane `go test` step will correctly fail before later recovery proof can run.

### 2026-04-24 Reconcile Visibility Constraint

- When proving declared-doc drift on an active Temporal task run, do not assume the reconcile signal will be visible in the immediate task-read response while the current workflow activity is still running; poll the active run until `doc_runtime_divergence_status` reflects the reconciled snapshot.
