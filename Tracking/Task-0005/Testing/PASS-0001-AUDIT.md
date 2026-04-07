# Pass 0001 Audit

## Scope

`PASS-0001` moved the backend from scaffold to a bounded live reconcile and readback slice:

- tracked job specs still load from `C:\Users\gregs\.codex\Orchestration\Jobs`
- `schedule` triggers now compile into deterministic Temporal schedule ids
- service startup reconcile and explicit `POST /sync` both exercise the same reconcile path
- backend APIs now expose health, job list, job detail, and recent-runs readback for the dashboard
- the repo-local Temporal compose stack was corrected so the live backend can actually connect and reconcile on this host

This pass deliberately does not claim real `manual` or `webhook` execution, worker-hosted workflows, or the `codex exec` activity path yet. Those remain `PASS-0002` work.

## Verification

Executed from `C:\Agent\CodexDashboard\backend\orchestration` unless noted otherwise:

```powershell
go test ./...
docker compose -f dev\docker-compose.temporal-postgres.yml up -d
docker compose -f dev\docker-compose.temporal-postgres.yml ps
temporal operator cluster health --address 127.0.0.1:7233
go run .\cmd\controlplane
Invoke-WebRequest http://127.0.0.1:4318/health | Select-Object -ExpandProperty Content
Invoke-WebRequest http://127.0.0.1:4318/jobs | Select-Object -ExpandProperty Content
Invoke-WebRequest http://127.0.0.1:4318/jobs/codex-daily-agentic-swe-digest | Select-Object -ExpandProperty Content
Invoke-WebRequest 'http://127.0.0.1:4318/runs?job_id=codex-daily-agentic-swe-digest' | Select-Object -ExpandProperty Content
Invoke-WebRequest -Method Post http://127.0.0.1:4318/sync | Select-Object -ExpandProperty Content
temporal schedule list --address 127.0.0.1:7233
```

Observed result:

- `go test ./...` passed across the new control-plane, HTTP, and Temporal adapter code
- the repo-local compose stack was healthy with Postgres, Temporal, and Temporal UI all running
- `temporal operator cluster health` returned `SERVING`
- control-plane startup reconcile succeeded and `/health` reported `job_count: 3` with a successful last-sync timestamp
- `/jobs` reported all three tracked digest jobs as `in_sync`
- `/jobs/{job_id}` returned the expected desired-state metadata plus runtime schedule status and next-run times
- `/runs?job_id=...` returned a stable empty array when no executions had happened yet
- `POST /sync` completed successfully with no further changes after startup reconcile had already created the schedules
- `temporal schedule list` showed the three managed schedule ids:
  - `codex-job--codex-daily-agentic-swe-digest--00`
  - `codex-job--codex-daily-physical-agents-digest--00`
  - `codex-job--codex-daily-ue-determinism-digest--00`

## Requirement Mapping

| Requirement | Evidence | Result |
| --- | --- | --- |
| Startup reconcile works against the supported spec shape | `go run .\cmd\controlplane`; `/health`; `/jobs`; `temporal schedule list` | Passed |
| Explicit sync works without restart | `POST /sync` returned success after startup reconcile had already converged the runtime | Passed |
| Backend can answer desired-vs-runtime status for the dashboard | `/jobs`; `/jobs/{job_id}`; `/runs?job_id=...`; [controlplane.go](/c:/Agent/CodexDashboard/backend/orchestration/internal/controlplane/controlplane.go); [mux.go](/c:/Agent/CodexDashboard/backend/orchestration/internal/httpapi/mux.go) | Passed |
| Focused automated proof exists for compile/diff and API behavior | [controlplane_test.go](/c:/Agent/CodexDashboard/backend/orchestration/internal/controlplane/controlplane_test.go); [mux_test.go](/c:/Agent/CodexDashboard/backend/orchestration/internal/httpapi/mux_test.go); `go test ./...` | Passed |
| Backend does not depend on the Tk process being alive | `go run .\cmd\controlplane` plus direct HTTP and Temporal CLI proof, with no desktop app involved | Passed |

## Caveat

The current drift comparison is intentionally bounded to the fields the service manages directly in schedule action metadata plus the visible Temporal schedule state returned by the SDK. That is sufficient for this pass and the current desired-state corpus, but it is not yet a full arbitrary-edit detector for every possible out-of-band Temporal schedule mutation.

## Verdict

`ready`
