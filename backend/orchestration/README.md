# Orchestration Backend

This directory is the Go control-plane home for `Task-0005`.

PASS-0000 establishes:

- the v1 job-spec contract under `C:\Users\gregs\.codex\Orchestration\Jobs`
- a minimal Go service skeleton under `backend/orchestration/`
- a repo-local Temporal plus Postgres dev stack definition
- honest operator notes about the current missing local toolchain

## Current Scope

The scaffold here is intentionally narrow. It does not reconcile or execute jobs yet.

The first backend slice only proves:

- config loading
- spec-file discovery and validation shape
- HTTP server skeleton and health endpoint
- local dev-stack layout for Temporal plus Postgres

Future passes add:

- startup reconcile and explicit sync
- read APIs for the Tk dashboard
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

## Toolchain Status On This Host

Verified during `PASS-0000` on April 6, 2026:

- `codex`: available
- `go`: missing from `PATH`
- `docker`: missing from `PATH`
- `temporal`: missing from `PATH`

That means this pass can land the scaffold and operator docs, but it cannot honestly claim a locally executed Go build or a live Temporal runtime on this host until those tools are installed.

## Suggested Bootstrap Once Tooling Exists

```powershell
cd C:\Agent\CodexDashboard\backend\orchestration
go test ./...
go run .\cmd\controlplane
```
