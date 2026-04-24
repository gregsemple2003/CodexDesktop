# Orchestration Backend

This directory is the Go control-plane home for `Task-0005`.

The current backend slice establishes:

- the v1 job-spec contract under `C:\Users\gregs\.codex\Orchestration\Jobs`
- a Go control-plane service under `backend/orchestration/`
- declared task-doc readback for `Tracking/Task-<id>/` under the Task-0008 contract
- a repo-local Temporal plus Postgres stack definition that can be run in separate lanes
- startup reconcile plus explicit sync against Temporal schedules
- read APIs for health, job list, job detail, and recent runs
- a Temporal worker-hosted `codex.exec.job` workflow plus `manual` and `webhook` trigger entrypoints

## Current Scope

The current backend slice is intentionally narrow. It reconciles desired `schedule` triggers into Temporal, exposes readback for the dashboard, and routes `schedule`, `manual`, and `webhook` into one durable workflow type.

Task-0008 extends that backend with the first task-readback contract:

- parse declared task docs on demand from the repo worktree
- expose `GET /api/v1/tasks`
- expose `GET /api/v1/tasks/{task_id}`
- expose `POST /api/v1/tasks/{task_id}/dispatch`
- expose `GET /api/v1/task-runs/{run_id}`
- expose `POST /api/v1/task-runs/{run_id}/state`
- expose `POST /api/v1/task-runs/{run_id}/poke`
- expose `POST /api/v1/task-runs/{run_id}/interrupt`
- expose `POST /api/v1/task-runs/{run_id}/retry-cleanup`
- expose `POST /api/v1/task-runs/{run_id}/resolve-interrupt-review`
- keep task meaning, state envelope, dispatch-readiness, and attention inputs in backend readback rather than client heuristics

Task-0008 also starts the first durable dispatch slice:

- dispatch creates a Temporal-backed task run
- dispatch provisions an exclusive backend-owned checkout lane
- dispatch captures the baseline commit the owned lane may restore to later
- dispatch bootstraps the owned lane, captures its current commit, and writes a bootstrap artifact under the backend run-artifact root
- task runs can accept backend-owned post-dispatch state updates
- active-run reads can supervise stale progress into `sleeping_or_stalled`
- active-run reads can supervise stale human waits into `human_wait_stale`
- interrupt can restore the owned checkout to its recorded restore commit
- interrupt cleanup failures surface through dedicated repo-lane reset failure fields
- `poke` creates a durable backend-worker follow-up that later runtime progress can complete
- cleanup retry can restore a cleanup-blocked owned checkout and convert the run into `interrupt_review`
- pending `interrupt_review` blocks redispatch until the review is explicitly resolved
- interrupt review resolution records a durable decision and returns the task to dispatch-ready state
- dispatch can now move a run into backend-produced `running` state without requiring a manual `/state` mutation first
- terminal runs stop owning the task's current live story so the task can become dispatchable again

Real task execution inside the owned checkout remains a future slice.

## Scheduling Boundary

The service-lane Windows Scheduled Task exists only to keep the backend alive at user logon. It is bootstrap infrastructure, not the product scheduling surface for user jobs.

When a human asks to set up or change a job in this repo:

- implement it through tracked v1 job specs plus backend reconcile into Temporal
- verify it through `GET /api/v1/jobs`, Temporal readback, and the dashboard Jobs tab
- do not treat an ad hoc Windows Scheduled Task as an acceptable substitute unless the human explicitly asks for the legacy Windows path

The intended local operating model now has two lanes:

- `service lane`
  - the always-on local backend for real scheduled jobs
  - default dashboard backend URL: `http://127.0.0.1:4318`
  - default Temporal gRPC: `127.0.0.1:7233`
- `validation lane`
  - disposable backend for unit smoke, regression, and debugging work
  - default dashboard override URL: `http://127.0.0.1:14318`
  - default Temporal gRPC: `127.0.0.1:17233`

The current backend slice proves:

- config loading
- spec-file discovery and validation shape
- startup reconcile against the tracked `.codex` job specs
- explicit `POST /sync` against the same reconcile path
- worker-hosted `codex.exec.job` registration for the shared durable run path
- trigger APIs:
  - `POST /api/v1/jobs/{job_id}/run`
  - `POST /api/v1/webhooks/{path}`
- read APIs:
  - `GET /healthz`
  - `GET /api/v1/jobs`
  - `GET /api/v1/jobs/{job_id}`
  - `GET /api/v1/jobs/{job_id}/runs`
  - `GET /api/v1/tasks`
  - `GET /api/v1/tasks/{task_id}`
  - `GET /api/v1/task-runs/{run_id}`
- task-run mutation API:
- `POST /api/v1/task-runs/{run_id}/state`
- `POST /api/v1/task-runs/{run_id}/poke`
- `POST /api/v1/task-runs/{run_id}/interrupt`
- `POST /api/v1/task-runs/{run_id}/retry-cleanup`
- `POST /api/v1/task-runs/{run_id}/resolve-interrupt-review`
- task dispatch API:
  - `POST /api/v1/tasks/{task_id}/dispatch`
- `codex exec` command assembly with per-run artifact paths for JSONL events and final-message capture
- local dev-stack layout for Temporal plus Postgres

Future passes add:

- dashboard integration
- repo regression coverage for the backend-backed Jobs surface

## Service Lane

Install and start the always-on service lane:

```powershell
cd C:\Agent\CodexDashboard\backend\orchestration
powershell -ExecutionPolicy Bypass -File .\scripts\Install-ServiceLane.ps1
```

That script:

- registers a Scheduled Task named `CodexDashboard-Orchestration-ServiceLane`
- starts the service-lane runner at user logon
- ensures the Temporal/Postgres compose stack is up on the default ports
- builds and launches the control-plane binary under `%LOCALAPPDATA%\CodexDashboard\orchestration-service-lane\`

Useful service-lane commands:

```powershell
cd C:\Agent\CodexDashboard\backend\orchestration
powershell -ExecutionPolicy Bypass -File .\scripts\Get-ServiceLaneStatus.ps1
powershell -ExecutionPolicy Bypass -File .\scripts\Start-ServiceLane.ps1
powershell -ExecutionPolicy Bypass -File .\scripts\Stop-ServiceLane.ps1
```

Service-lane endpoints:

- jobs backend: `http://127.0.0.1:4318`
- Temporal gRPC: `127.0.0.1:7233`
- Temporal Web UI: `http://127.0.0.1:8080`
- Postgres: `127.0.0.1:5432`

Service-lane notes:

- it runs under the current interactive Windows user because `.codex` state and the logged-in `codex` CLI session live there
- it is meant to stay up for real scheduled jobs rather than being torn down after normal proof work

## Validation Lane

The repo-local Temporal stack file is:

- `backend/orchestration/dev/docker-compose.temporal-postgres.yml`

This is a local-first development stack, not a production deployment.

It follows the official Temporal guidance that local Temporal can be run with Docker Compose and that Temporal's CLI/dev-server can also be run from Docker when needed. The archived `temporalio/docker-compose` repo now points people to newer samples, so this repo keeps a local compose definition instead of depending on the archived example at runtime.

Start the disposable validation lane:

```powershell
cd C:\Agent\CodexDashboard\backend\orchestration
powershell -ExecutionPolicy Bypass -File .\scripts\Start-ValidationLane.ps1
```

Stop the disposable validation lane:

```powershell
cd C:\Agent\CodexDashboard\backend\orchestration
powershell -ExecutionPolicy Bypass -File .\scripts\Stop-ValidationLane.ps1
```

Validation-lane endpoints:

- jobs backend: `http://127.0.0.1:14318`
- Temporal gRPC: `127.0.0.1:17233`
- Temporal Web UI: `http://127.0.0.1:18080`
- Postgres: `127.0.0.1:15432`

When running the desktop app or regression against the validation lane, override the backend URL:

```powershell
$env:CODEX_DASHBOARD_JOBS_BACKEND_URL = "http://127.0.0.1:14318"
python -m app.codex_dashboard
```

Clear that override when you want the app to talk to the service lane again:

```powershell
Remove-Item Env:CODEX_DASHBOARD_JOBS_BACKEND_URL -ErrorAction SilentlyContinue
```

## Supporting Live Checks

```powershell
Invoke-WebRequest http://127.0.0.1:4318/healthz | Select-Object -ExpandProperty Content
Invoke-WebRequest http://127.0.0.1:4318/api/v1/jobs | Select-Object -ExpandProperty Content
Invoke-WebRequest http://127.0.0.1:4318/api/v1/jobs/codex-daily-agentic-swe-digest | Select-Object -ExpandProperty Content
Invoke-WebRequest http://127.0.0.1:4318/api/v1/jobs/codex-daily-agentic-swe-digest/runs | Select-Object -ExpandProperty Content
Invoke-WebRequest http://127.0.0.1:4318/api/v1/tasks | Select-Object -ExpandProperty Content
Invoke-WebRequest http://127.0.0.1:4318/api/v1/tasks/Task-0008 | Select-Object -ExpandProperty Content
Invoke-WebRequest -Method Post http://127.0.0.1:4318/api/v1/tasks/Task-0008/dispatch | Select-Object -ExpandProperty Content
$dispatch = Invoke-RestMethod -Method Post http://127.0.0.1:4318/api/v1/tasks/Task-0008/dispatch
Get-Content -Raw $dispatch.repo_lane.bootstrap_artifact_path
Invoke-WebRequest http://127.0.0.1:4318/api/v1/task-runs/taskrun--Task-0008--active | Select-Object -ExpandProperty Content
$body = '{"state":"waiting_for_human","reason_code":"approval_required","state_summary":"Run is waiting for approval.","next_owner":"human","next_expected_event":"Approve or redirect the next backend step."}'
Invoke-WebRequest -Method Post -Uri http://127.0.0.1:4318/api/v1/task-runs/taskrun--Task-0008--active/state -ContentType 'application/json' -Body $body | Select-Object -ExpandProperty Content
Invoke-WebRequest -Method Post http://127.0.0.1:4318/api/v1/task-runs/taskrun--Task-0008--active/poke | Select-Object -ExpandProperty Content
Invoke-WebRequest -Method Post http://127.0.0.1:4318/api/v1/task-runs/taskrun--Task-0008--active/interrupt | Select-Object -ExpandProperty Content
Invoke-WebRequest -Method Post http://127.0.0.1:4318/api/v1/task-runs/taskrun--Task-0008--active/retry-cleanup | Select-Object -ExpandProperty Content
Invoke-WebRequest -Method Post http://127.0.0.1:4318/api/v1/jobs/codex-daily-agentic-swe-digest/run | Select-Object -ExpandProperty Content
Invoke-WebRequest -Method Post http://127.0.0.1:4318/api/v1/webhooks/digests/physical-agents | Select-Object -ExpandProperty Content
Invoke-WebRequest -Method Post http://127.0.0.1:4318/sync | Select-Object -ExpandProperty Content
temporal schedule list --address 127.0.0.1:7233
```

Validation-lane variants of those same checks use `14318` and `17233`.

Optional environment overrides:

- `CODEX_ORCHESTRATION_CODEX_EXECUTABLE`
  - override the `codex.exe` path when it is not already on `PATH`
- `CODEX_ORCHESTRATION_RUNS_ROOT`
  - override the local artifact root for per-run JSONL, stderr, and final-message files
- `CODEX_DASHBOARD_JOBS_BACKEND_URL`
  - override the dashboard's backend URL when the app should talk to the validation lane instead of the service lane
