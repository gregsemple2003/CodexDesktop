# Task 0005 Handoff

## Current Status

`PASS-0001` is complete. `PASS-0002` implementation is in place, but the pass is waiting on explicit human approval for a real paid `codex exec` proof run.

## Current Baseline

The agreed direction for this task is:

- a separate Go service under `backend/orchestration/` owns scheduling and orchestration control
- Temporal plus Postgres owns runtime durability
- Git-tracked JSON under `C:\Users\gregs\.codex\Orchestration\Jobs\` remains the desired-state source of truth
- startup reconcile is required, but the first slice also needs an explicit non-restart sync path
- v1 trigger scope is:
  - `schedule`
  - `manual`
  - `webhook`
- CodexDashboard integration is required as a correctness surface, not as the scheduler host
- the task should stay narrow and avoid widening into a full agent framework in its first slice
- the existing `declared-jobs.json` Windows registry is useful migration input, but it is not the long-term desired-state format for this task
- the Tk `Jobs` surface should survive as a backend client, not as a local Windows reconciler

`PASS-0000` delivered:

- `backend/orchestration/` Go scaffold with config, health endpoint, and spec-loader skeleton
- repo-local Temporal plus Postgres dev-stack docs and compose file
- new `.codex/Orchestration/Jobs/job-spec.schema.json`
- new `.codex/Orchestration/Jobs/specs/*.json` starter desired-state specs
- Python-side spec loading and validation helpers with unit coverage

`PASS-0001` delivered:

- startup reconcile from tracked `.codex` job specs into Temporal schedule ids
- explicit `POST /sync` through the same reconcile path used at startup
- backend read APIs for:
  - `GET /healthz`
  - `GET /api/v1/jobs`
  - `GET /api/v1/jobs/{job_id}`
  - `GET /api/v1/jobs/{job_id}/runs`
- focused Go tests for compile or diff behavior and HTTP routes
- live proof against the repo-local Temporal plus Postgres stack, including `temporal schedule list`

`PASS-0002` now has implementation-ready code for:

- a worker-hosted `codex.exec.job` workflow and activity registration inside the control-plane process
- `schedule` actions updated to start that same workflow type with frozen desired-state identity
- `POST /api/v1/jobs/{job_id}/run` for `manual`
- `POST /api/v1/webhooks/{path}` for `webhook`
- `codex exec` command assembly with a configurable executable path and per-run artifact files under the local runs root
- unit coverage for trigger routing and command assembly

What is still missing before `PASS-0002` can close honestly:

- one explicit human approval to run a real paid `codex exec` job through the worker
- live proof that at least one trigger path reaches a real Codex execution and records Temporal ids plus desired spec hash

Research output captured so far:

- [RESEARCH-PLAN.md](/c:/Agent/CodexDashboard/Tracking/Task-0005/RESEARCH-PLAN.md)
- [RESEARCH-ANALYSIS.md](/c:/Agent/CodexDashboard/Tracking/Task-0005/RESEARCH-ANALYSIS.md)
- [RESEARCH.md](/c:/Agent/CodexDashboard/Tracking/Task-0005/RESEARCH.md)
- [Cron System with JSON.md](/c:/Agent/CodexDashboard/Tracking/Task-0005/Research/Cron%20System%20with%20JSON.md)

## Next Step

Continue `PASS-0002`:

1. Get explicit human approval for one paid `codex exec` validation run.
2. Start the backend with the local Temporal stack and worker active.
3. Drive one real `manual` or `webhook` run through the backend and capture the resulting Temporal ids, spec hash, and Codex artifacts.
4. Close `PASS-0002`, then continue into `PASS-0003`.

## Watchouts

- do not let the dashboard own scheduler logic
- do not let startup reconcile become the only sync path
- do not widen this task into a full agent-graph or self-improvement system
- keep Git as desired state and Temporal as runtime truth
- the local compose stack is usable on this host; task-owned control-plane validation processes were cleaned up after proof runs
- `manual` and `webhook` now route into the durable workflow path, but a real Codex execution has not been approved or proven yet
- `CODEX_ORCHESTRATION_CODEX_EXECUTABLE` may be needed on hosts where `codex.exe` is not already on `PATH`

## References

- [TASK.md](/c:/Agent/CodexDashboard/Tracking/Task-0005/TASK.md)
- [PLAN.md](/c:/Agent/CodexDashboard/Tracking/Task-0005/PLAN.md)
- [RESEARCH.md](/c:/Agent/CodexDashboard/Tracking/Task-0005/RESEARCH.md)
- [Cron System with JSON.md](/c:/Agent/CodexDashboard/Tracking/Task-0005/Research/Cron%20System%20with%20JSON.md)
- [PASS-0000-AUDIT.md](/c:/Agent/CodexDashboard/Tracking/Task-0005/Testing/PASS-0000-AUDIT.md)
- [PASS-0001-AUDIT.md](/c:/Agent/CodexDashboard/Tracking/Task-0005/Testing/PASS-0001-AUDIT.md)
