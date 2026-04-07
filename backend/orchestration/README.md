# Orchestration Backend

This directory is the Go control-plane home for `Task-0005`.

The current backend slice establishes:

- the v1 job-spec contract under `C:\Users\gregs\.codex\Orchestration\Jobs`
- a Go control-plane service under `backend/orchestration/`
- a repo-local Temporal plus Postgres dev stack definition
- startup reconcile plus explicit sync against Temporal schedules
- read APIs for health, job list, job detail, and recent runs
- a Temporal worker-hosted `codex.exec.job` workflow plus `manual` and `webhook` trigger entrypoints

## Current Scope

The current backend slice is intentionally narrow. It reconciles desired `schedule` triggers into Temporal, exposes readback for the dashboard, and routes `schedule`, `manual`, and `webhook` into one durable workflow type.

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

## Local Dev Stack

The repo-local Temporal stack file is:

- `backend/orchestration/dev/docker-compose.temporal-postgres.yml`

This is a local-first development stack, not a production deployment.

It follows the official Temporal guidance that local Temporal can be run with Docker Compose and that Temporal's CLI/dev-server can also be run from Docker when needed. The archived `temporalio/docker-compose` repo now points people to newer samples, so this repo keeps a local compose definition instead of depending on the archived example at runtime.

Operator workflow:

```powershell
cd C:\Agent\CodexDashboard\backend\orchestration
docker compose -f dev\docker-compose.temporal-postgres.yml up -d
docker compose -f dev\docker-compose.temporal-postgres.yml ps
```

Expected local endpoints:

- Temporal gRPC: `127.0.0.1:7233`
- Temporal Web UI: `http://127.0.0.1:8080`
- Postgres: `127.0.0.1:5432`

Stop the stack:

```powershell
docker compose -f dev\docker-compose.temporal-postgres.yml down
```

## Suggested Bootstrap

```powershell
cd C:\Agent\CodexDashboard\backend\orchestration
go test ./...
go run .\cmd\controlplane
```

Supporting live checks:

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

Optional environment overrides:

- `CODEX_ORCHESTRATION_CODEX_EXECUTABLE`
  - override the `codex.exe` path when it is not already on `PATH`
- `CODEX_ORCHESTRATION_RUNS_ROOT`
  - override the local artifact root for per-run JSONL, stderr, and final-message files
