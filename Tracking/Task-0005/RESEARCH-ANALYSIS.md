# Task 0005 Research Analysis

## Problem 0001 Desired-State Schema And Migration Fit

### Current Repo Facts

The current jobs model under `app/codex_dashboard/jobs.py` is tightly coupled to Windows durable objects:

- desired state is stored as `startup_launcher` and `scheduled_task`
- the reconciler fingerprints startup scripts and exported Task Scheduler XML
- the dashboard reads declared state directly from `declared-jobs.json` and only later computes observed drift

That was the right fit for `Task-0004`, but it is the wrong boundary for `Task-0005`. The backend task needs Git-managed job intent, not a Git export of Windows runtime objects.

### Recommended Desired-State Layout

Keep `C:\Users\gregs\.codex\Orchestration\Jobs\` as the durable root, but move the v1 source of truth to one JSON file per job under a new `specs/` subtree with a shared schema at the root.

Recommended v1 shape:

- `C:\Users\gregs\.codex\Orchestration\Jobs\job.schema.json`
- `C:\Users\gregs\.codex\Orchestration\Jobs\specs\<job-id>.json`

Each job spec should stay intentionally small:

- identity: `job_id`, `title`, optional `description`
- desired state: enabled or disabled
- executor: start with inline `codex_exec` fields instead of a second profile DSL
- triggers: an array supporting only `schedule`, `manual`, and `webhook`
- runtime options: task queue, retry policy, timeouts, and overlap behavior where needed

This keeps the Git record diff-friendly and reviewer-friendly without adding a premature config compiler layer in the first slice.

### Migration Direction

The existing `declared-jobs.json` file is still useful as migration input, but it should stop being the authoritative format. `PASS-0000` should include a bounded migration path that preserves the current declared jobs while converting them into the new schema shape.

## Problem 0002 Runtime Topology And Temporal Boundary

### Current Repo Fit

The current app is still a Python/Tk overlay. There is no Go module, no backend service area, and no existing Temporal runtime inside the repo. The local shell also shows that `go`, `docker`, and `temporal` are not currently available on `PATH`, while `codex` and `codex exec` are available.

That means the task can be planned honestly now, but the first implementation pass must either establish or locate the missing Go and Temporal tooling before claiming a runnable backend stack.

### Recommended Backend Shape

Use one new Go service area under `backend/orchestration/` and keep the v1 runtime intentionally thin:

- an HTTP API layer for list, detail, health, sync, manual run, and webhook intake
- a spec loader and validator that reads Git-managed job specs
- a reconcile layer that compares desired specs against Temporal runtime objects
- a Temporal worker that runs a single generic `JobRunWorkflow` plus executor Activities
- a `codex exec` Activity that shells out, captures machine-readable events, and records stable run metadata

For v1, the same binary can host both the API server and the Temporal worker in local development. The task does not need a separate control-plane database yet if the service can answer the required dashboard queries from Git plus Temporal APIs.

### Temporal Boundary

Temporal should own runtime truth:

- schedules for `schedule` triggers
- workflow histories and retries
- run identifiers and execution state
- durable wait or retry behavior across crashes

Git should own desired state:

- job identity
- trigger declarations
- executor inputs
- human-edited policy

The dashboard should own neither. It should only query the backend and invoke bounded actions.

### Why Temporal Still Fits The Narrow v1

The official schedule docs show the control-plane features this task explicitly wants:

- schedule create and update
- overlap policies
- catchup windows
- pause-on-failure
- manual trigger
- backfill

That lets the task avoid inventing its own scheduler semantics while still keeping Git as the human-facing desired-state record.

### Dev Runtime Direction

The official Temporal PostgreSQL Docker Compose example is enough to justify a repo-local local-dev stack, but the upstream compose repo is now archived. The honest repo fit is to vendor the needed compose file and supporting notes into this repo rather than treating the archived upstream example as a runtime dependency.

## Problem 0003 Dashboard Integration And Proof Strategy

### Current Surface Facts

The existing Tk `Jobs` surface already has useful structure:

- a separate `Jobs` tab
- glanceable counts
- per-job rows
- a details panel
- explicit refresh-style actions
- a backend-blocked fallback state

Those affordances can survive, but the current implementation should stop reading Windows durable state directly from Python.

### Recommended Dashboard Boundary

Reuse the `Jobs` tab as a correctness surface, but change the data source and action semantics:

- `Refresh` becomes a backend readback refresh
- `Force Reconcile` should become explicit `Sync now`
- add bounded `Run now` where the selected job supports it
- details should show desired spec plus backend runtime facts, not raw Windows-only state
- the default view should surface desired revision, sync status, recent runs, last error, and next run where applicable

This preserves the hotkey-first cockpit intent from `Design/GENERAL-DESIGN.md` while making the dashboard a client instead of a scheduler host.

### Proof Strategy

Pass-local proof should stay split by responsibility:

- Go unit tests for schema validation, compile or diff logic, and API behavior
- focused executor tests around `codex exec` command assembly and run-metadata capture
- Python unit tests for dashboard API adaptation and UI-state mapping

Task-level closure still requires a real app-surface regression lane. The current `REG-002` is Windows-registry-specific, so repo-root `REGRESSION.md` must be updated to cover the backend-backed Jobs surface before the task can close honestly.

## Research Verdict

Task research is planning-ready.

## Recommended Direction

- migrate desired state under `.codex/Orchestration/Jobs` from Windows-object registry to one-file-per-job Git specs
- add a thin Go control plane under `backend/orchestration/` that owns validation, reconcile, sync, runtime APIs, and the Temporal worker
- use Temporal schedules only for schedule triggers; manual and webhook triggers should start the same durable run workflow through the backend
- keep `codex exec` in Activities, never in Workflow code
- keep the Tk dashboard as a backend client and update `REGRESSION.md` for the new Jobs lane
- treat missing `go`, `docker`, and `temporal` tooling on `PATH` as a first-pass environment prerequisite, not as something to hand-wave past
