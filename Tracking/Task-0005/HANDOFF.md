# Task 0005 Handoff

## Current Status

`PASS-0002` is complete with caveats. The bounded final retry proved one real durable `codex exec` run end to end for `codex-daily-agentic-swe-digest`, including a completed Temporal workflow, per-run artifacts, a regenerated report under `.codex\reports`, and a successful Gmail digest send. `PASS-0003` is now the active pass.

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

`PASS-0002` delivered:

- a worker-hosted `codex.exec.job` workflow and activity registration inside the control-plane process
- `schedule` actions updated to start that same workflow type with frozen desired-state identity
- `POST /api/v1/jobs/{job_id}/run` for `manual`
- `POST /api/v1/webhooks/{path}` for `webhook`
- `codex exec` command assembly with a configurable executable path and per-run artifact files under the local runs root
- unit coverage for trigger routing and command assembly
- a short schedule catchup window so startup reconcile stays bounded during proof work
- a successful bounded live manual proof:
  - workflow id: `job/codex-daily-agentic-swe-digest/manual/18a81bff-3355-418e-8559-224407f7f586`
  - run id: `019d686c-3c82-7a12-81f2-50d60cd5ffd5`
  - desired spec hash: `38a97b5d67aa3cf10a8231b972e1f3404fe29c28be37bac1e54e2ce378524b62`
  - artifact root: `C:\Users\gregs\AppData\Local\CodexDashboard\orchestration-runs\pass-0002-final-retry-20260407-104944\codex-daily-agentic-swe-digest\job_codex-daily-agentic-swe-digest_manual_18a81bff-3355-418e-8559-224407f7f586\`
  - final outcome: report regenerated and Gmail digest sent successfully at `2026-04-07T10:58:44-04:00`

Research output captured so far:

- [RESEARCH-PLAN.md](/c:/Agent/CodexDashboard/Tracking/Task-0005/RESEARCH-PLAN.md)
- [RESEARCH-ANALYSIS.md](/c:/Agent/CodexDashboard/Tracking/Task-0005/RESEARCH-ANALYSIS.md)
- [RESEARCH.md](/c:/Agent/CodexDashboard/Tracking/Task-0005/RESEARCH.md)
- [Cron System with JSON.md](/c:/Agent/CodexDashboard/Tracking/Task-0005/Research/Cron%20System%20with%20JSON.md)

## Next Step

Start `PASS-0003`:

1. Replace the Tk `Jobs` tab's local Windows reconciliation path with backend API consumption.
2. Show backend-backed desired-vs-runtime state, recent runs, next run, and last error without making the dashboard the scheduler host.
3. Wire bounded `Sync now` and `Run now` actions through the backend.
4. Update repo-root `REGRESSION.md` for the backend-backed Jobs lane and then run the real desktop regression lane honestly.

## Watchouts

- do not let the dashboard own scheduler logic
- do not let startup reconcile become the only sync path
- do not widen this task into a full agent-graph or self-improvement system
- keep Git as desired state and Temporal as runtime truth
- the original startup over-release and `CreateProcessAsUserW failed: 1920` failure are now historical evidence in [BUG-0001.md](/c:/Agent/CodexDashboard/Tracking/Task-0005/BUG-0001.md), not the active blocker
- the successful Windows executor path currently depends on `--dangerously-bypass-approvals-and-sandbox`
- the real digest proof touched user `.codex` runtime state under `reports\` and `gmail-digest-email\`; treat those as operator/runtime artifacts rather than dashboard product files
- `go`, `docker`, and `temporal` are available on this host, but this shell still needs explicit executable resolution when `PATH` is stale

## References

- [TASK.md](/c:/Agent/CodexDashboard/Tracking/Task-0005/TASK.md)
- [PLAN.md](/c:/Agent/CodexDashboard/Tracking/Task-0005/PLAN.md)
- [RESEARCH.md](/c:/Agent/CodexDashboard/Tracking/Task-0005/RESEARCH.md)
- [Cron System with JSON.md](/c:/Agent/CodexDashboard/Tracking/Task-0005/Research/Cron%20System%20with%20JSON.md)
- [PASS-0000-AUDIT.md](/c:/Agent/CodexDashboard/Tracking/Task-0005/Testing/PASS-0000-AUDIT.md)
- [PASS-0001-AUDIT.md](/c:/Agent/CodexDashboard/Tracking/Task-0005/Testing/PASS-0001-AUDIT.md)
- [BUG-0001.md](/c:/Agent/CodexDashboard/Tracking/Task-0005/BUG-0001.md)
