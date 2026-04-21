# Orchestration Backend

This directory is the Go control-plane home for `Task-0005`.

The current backend slice establishes:

- the v1 job-spec contract under `C:\Users\gregs\.codex\Orchestration\Jobs`
- a Go control-plane service under `backend/orchestration/`
- a repo-local Temporal plus Postgres stack definition that can be run in separate lanes
- startup reconcile plus explicit sync against Temporal schedules
- read APIs for health, job list, job detail, and recent runs
- a Temporal worker-hosted `codex.exec.job` workflow plus `manual` and `webhook` trigger entrypoints

## Current Scope

The current backend slice is intentionally narrow. It reconciles desired `schedule` triggers into Temporal, exposes readback for the dashboard, and routes `schedule`, `manual`, and `webhook` into one durable workflow type.

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
