# Task 0005

## Title

Build a Git-backed Go + Temporal jobs control plane for CodexDashboard.

## Summary

CodexDashboard currently has a local Python/Tk overlay and a partially explored dashboard-first jobs surface from `Task-0004`.

That is not the right foundation for the durable scheduler and orchestration direction the product now wants. The next honest move is a separate long-running backend:

- Git-tracked JSON job specs as desired state
- a Go control-plane service that reconciles desired state into Temporal
- a local Temporal + Postgres runtime for durable schedules and executions
- dashboard integration that shows desired-vs-runtime correctness without owning scheduling logic

The v1 closure bar is a jobs system with `schedule`, `manual`, and `webhook` triggers. Future agent-oriented primitives should only be included now when they materially improve that v1 shape instead of widening the task.

## Goals

- Introduce a separate Go service for jobs and orchestration control.
- Self-host Temporal plus Postgres as the runtime durability layer for local-first development.
- Keep Git-tracked JSON under `C:\Users\gregs\.codex\Orchestration\Jobs\` as the source of truth for desired job specs.
- Reconcile desired job state into Temporal schedules and workflows on service startup.
- Provide a non-restart sync path so Git changes can be reconciled explicitly while the service is already running.
- Support v1 trigger types:
  - `schedule`
  - `manual`
  - `webhook`
- Support at least one real executor path for Codex-oriented work, starting with `codex exec`.
- Expose backend APIs for job list/detail, recent runs, manual trigger, webhook intake, sync, and health.
- Integrate CodexDashboard so it can display desired state, runtime state, drift, recent runs, and last error for correctness checking.
- Let the dashboard invoke bounded backend actions such as `Sync now` and `Run now`.
- Shape the runtime so future agent orchestration can grow from it without making full agent primitives part of this task's closure bar.

## Non-Goals

- Building a full agent graph engine, autonomous planner, self-improvement loop, or self-healing system in v1.
- Embedding scheduler or clock ownership inside the Tk dashboard.
- Rewriting the dashboard in Go or replacing Tk in this task.
- Multi-node HA, managed cloud deployment, or a production-grade public webhook edge in the first slice.
- Building a full in-dashboard authoring, PR, or Git-conflict workflow for job specs in v1.
- Reworking token ingest, chart math, or other existing dashboard analytics except where the new correctness surface needs integration.
- Finishing the superseded `Task-0004` dashboard-first Jobs implementation path.

## Constraints And Baseline

- `Design/GENERAL-DESIGN.md` still defines CodexDashboard as a hotkey-first private cockpit. Any dashboard job surface added here must stay bounded and operator-clear rather than turning the app into a general admin console.
- `Task-0004` created useful reference material around Git-tracked desired state and dashboard Jobs UI, but that implementation path is now superseded by this backend-first direction.
- The dashboard should remain a client of the job system, not the scheduler host.
- Git is the desired-state record. Temporal plus Postgres is runtime truth. SQLite may remain only as dashboard-local cache or view state when still useful.
- Local-first deployment is acceptable for v1, but webhook support must be honest about host reachability and operator prerequisites.
- Temporal durability and deterministic workflow replay are intentional architectural bets because this system is expected to grow toward broader orchestration later.
- The first slice should prefer a small number of well-defined job JSON shapes over a prematurely general DSL.

## Expected Resolution

- A long-running Go control-plane service and worker can run independently of the dashboard process.
- The service validates job specs from `C:\Users\gregs\.codex\Orchestration\Jobs\`, reconciles them into Temporal at startup, and exposes explicit sync without requiring restart.
- Temporal owns schedules, workflow histories, run state, retries, and runtime coordination. Git remains the human-edited desired state.
- v1 jobs can be triggered by schedule, manual request, or webhook.
- The dashboard can read backend state and show at least:
  - desired spec identity or version
  - runtime deployment or sync state
  - recent run outcomes
  - drift or validation errors
  - last error and next run where applicable
- The dashboard provides bounded correctness controls such as `Sync now` and `Run now` for appropriate jobs.
- The task leaves a clean path for future agent-oriented features such as signals, queries, or child workflows without making them closure requirements now.

## Implementation Home

Keep task-owned artifacts under `Tracking/Task-0005/`.

Implement the Go control-plane and worker code under a new repo-root service area at `backend/orchestration/`.

Keep CodexDashboard integration changes under `app/codex_dashboard/`.

Keep Git-tracked job specs and schemas under `C:\Users\gregs\.codex\Orchestration\Jobs\`.

## Acceptance Criteria

- The repo defines a Git-tracked JSON schema and job layout under `C:\Users\gregs\.codex\Orchestration\Jobs\` for v1 jobs and triggers.
- The first supported trigger types are:
  - `schedule`
  - `manual`
  - `webhook`
- A local self-hosted Temporal plus Postgres dev stack is documented and runnable for this repo's workflow.
- A Go service can load and validate desired job specs and reconcile them into Temporal schedules or workflows on startup.
- The same service exposes an explicit sync path so Git changes do not require a service restart to reconcile.
- Schedule, manual, and webhook trigger paths can each start durable runs through Temporal.
- At least one real executor path is integrated end to end for Codex-oriented work, starting with `codex exec`.
- The backend records enough run identity to audit correctness, including job id plus desired-state version or Git revision and Temporal runtime identifiers.
- The dashboard can query the backend and show desired-vs-runtime status, recent runs, and last error for real job data.
- The dashboard can invoke at least one bounded control path against the backend, such as explicit sync or manual run.
- The system does not require the dashboard process to be running for scheduled, manual, or webhook jobs to execute.
- Focused automated proof exists for spec validation, reconciliation behavior, and at least one real trigger path.
- Repo-root regression coverage is updated or extended honestly for any new or materially changed dashboard job surface required to exercise this backend-backed correctness flow.

## References

- `Design/GENERAL-DESIGN.md`
- `Tracking/Task-0004/TASK.md`
- `Tracking/Task-0004/HANDOFF.md`
- `Tracking/Task-0005/Research/Cron System with JSON.md`
