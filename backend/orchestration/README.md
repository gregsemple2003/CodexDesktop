# Orchestration Backend

This directory is the Go control-plane home for `Task-0005`.

`PASS-0001` now establishes:

- the v1 job-spec contract under `C:\Users\gregs\.codex\Orchestration\Jobs`
- a Go control-plane service under `backend/orchestration/`
- a repo-local Temporal plus Postgres dev stack definition
- startup reconcile plus explicit sync against Temporal schedules
- read APIs for health, job list, job detail, and recent runs

## Current Scope

The current backend slice is intentionally narrow. It reconciles desired `schedule` triggers into Temporal and exposes readback for the dashboard, but it does not execute real `manual` or `webhook` runs yet.

The current backend slice proves:

- config loading
- spec-file discovery and validation shape
- startup reconcile against the tracked `.codex` job specs
- explicit `POST /sync` against the same reconcile path
- read APIs:
  - `GET /health`
  - `GET /jobs`
  - `GET /jobs/{job_id}`
  - `GET /runs?job_id=<job_id>`
- local dev-stack layout for Temporal plus Postgres

Future passes add:

- durable run workflows
- `schedule`, `manual`, and `webhook` trigger execution
- the real `codex exec` activity path

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
Invoke-WebRequest http://127.0.0.1:4318/health | Select-Object -ExpandProperty Content
Invoke-WebRequest http://127.0.0.1:4318/jobs | Select-Object -ExpandProperty Content
Invoke-WebRequest -Method Post http://127.0.0.1:4318/sync | Select-Object -ExpandProperty Content
temporal schedule list --address 127.0.0.1:7233
```
