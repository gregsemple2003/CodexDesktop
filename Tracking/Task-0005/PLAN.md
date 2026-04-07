# Task 0005 Plan

## Planning Verdict

This plan is ready for explicit human approval.

Human gate:

- do not start `PASS-0000` until the human approves this `PLAN.md`

## PASS-0000 Job Spec, Dev Stack, And Backend Scaffold

### Objective

Establish the durable backend foundation for the task:

- the v1 Git job-spec layout under `.codex/Orchestration/Jobs`
- the new `backend/orchestration/` Go module and service skeleton
- repo-local docs for a self-hosted Temporal plus Postgres dev stack
- honest handling of missing local `go`, `docker`, and `temporal` tooling

### Implementation Notes

- add the new backend service area and module bootstrap
- define the first v1 job schema for `schedule`, `manual`, and `webhook`
- migrate or bridge the current declared jobs into the new desired-state layout
- document the local Temporal plus Postgres stack in repo-local docs rather than relying on an archived upstream example at runtime

### Verification

- validate example job specs against the tracked schema
- add focused backend tests for spec loading and validation
- record the exact toolchain state used for Go and Temporal commands

### Exit Bar

- `backend/orchestration/` exists with a buildable service skeleton or the task is honestly blocked on missing toolchain
- the tracked job schema and example specs exist under `.codex/Orchestration/Jobs`
- the repo documents a concrete local Temporal plus Postgres dev workflow
- the task can move into backend reconcile work without another schema reset

## PASS-0001 Reconcile, Sync, And Read APIs

### Objective

Teach the backend to turn Git-managed desired state into Temporal runtime state and expose backend readback for the dashboard.

### Implementation Notes

- load and validate job specs from Git
- compile schedule triggers into Temporal schedules
- compare desired specs against current Temporal state and classify drift
- expose backend endpoints for health, job list, job detail, recent runs, and explicit sync

### Verification

- unit tests for compile or diff behavior and API handlers
- focused proof that startup reconcile and explicit sync both exercise the same reconcile path
- durable task artifacts updated with any environment or runtime caveats

### Exit Bar

- startup reconcile works against the supported spec shape
- explicit sync works without restart
- the backend can answer desired-vs-runtime status for the dashboard
- the backend does not depend on the Tk process being alive

## PASS-0002 Durable Runs, Trigger Paths, And `codex exec`

### Objective

Deliver the first real execution path through Temporal for all required trigger types and capture enough run identity for auditability.

### Implementation Notes

- add a generic durable run workflow and executor Activities
- support `manual` and `webhook` trigger entrypoints through backend APIs
- ensure `schedule` triggers reach the same durable run path
- launch `codex exec` from an Activity with captured job id, desired revision or spec hash, and Temporal identifiers

### Verification

- unit tests around trigger routing and executor command assembly
- focused proof for at least one real trigger path reaching `codex exec`
- durable notes for webhook prerequisites such as local reachability

### Exit Bar

- `schedule`, `manual`, and `webhook` can each start durable runs honestly
- at least one real `codex exec` executor path works end to end
- run records preserve job id, desired revision or spec hash, and Temporal runtime ids
- the system still runs independently of the dashboard process

## PASS-0003 Dashboard Integration And Regression Prep

### Objective

Turn the existing Tk `Jobs` surface into a backend client that shows real desired-vs-runtime state and exposes bounded control actions.

### Implementation Notes

- replace local Windows jobs reconciliation on the `Jobs` tab with backend API consumption
- keep the Jobs surface bounded and operator-clear
- surface desired revision, sync state, recent runs, next run, and last error where applicable
- expose bounded `Sync now` and `Run now` actions where the job supports them
- update repo-root `REGRESSION.md` for the backend-backed Jobs lane

### Verification

- Python unit tests for backend response parsing and UI-state mapping
- supporting smoke that the dashboard can query the backend and render the new state
- prepare the task for the real desktop regression lane after planned passes complete

### Exit Bar

- the dashboard reads backend jobs state instead of local Windows durable state
- the default Jobs surface stays bounded and human-facing
- bounded control actions work through the backend
- repo-root regression requirements are updated honestly for the new surface

## PASS-0004 Service Lane And Validation Separation

### Objective

Align the task with the intended local operating model:

- a persistent service lane that stays up for real scheduled work
- a separate validation lane that can be started and torn down without disturbing scheduled jobs

### Implementation Notes

- parameterize the repo-local Temporal compose stack so different lanes can run on different ports
- add service-lane scripts that install, start, stop, and report the always-on backend
- keep the default dashboard backend URL pointed at the service lane while allowing validation and regression to override it explicitly
- update repo and task docs so future testing uses the validation lane instead of the always-on lane

### Verification

- prove the service lane can be installed and queried live on the default ports
- prove the validation lane can be started on separate ports while the service lane remains up
- capture durable task evidence for the new service-lane baseline and the updated testing story

### Exit Bar

- the local machine has a real always-on service lane for scheduled jobs
- validation work can use a separate runtime lane without taking down the service lane
- repo docs and task artifacts describe the two-lane model honestly
- the task closure wording matches the real delivered operating model

## Task-Level Regression After Planned Passes

After the planned passes close, run the repo-root desktop regression lane that honestly covers the backend-backed Jobs surface. Supporting backend smoke does not replace that real app-surface proof. When the service lane exists, regression should use the separate validation lane unless the task explicitly targets the service-lane install or uptime flow.

## Watchouts

- local `go`, `docker`, and `temporal` tooling are not on `PATH` yet
- do not let the dashboard regain scheduler ownership
- do not widen the schema into a general agent framework in v1
- do not treat backend-only smoke as closure proof for the Tk surface
