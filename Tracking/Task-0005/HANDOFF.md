# Task 0005 Handoff

## Current Status

`PASS-0001` is complete. `PASS-0002` is now active.

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
  - `GET /health`
  - `GET /jobs`
  - `GET /jobs/{job_id}`
  - `GET /runs?job_id=<job_id>`
- focused Go tests for compile or diff behavior and HTTP routes
- live proof against the repo-local Temporal plus Postgres stack, including `temporal schedule list`

Research output captured so far:

- [RESEARCH-PLAN.md](/c:/Agent/CodexDashboard/Tracking/Task-0005/RESEARCH-PLAN.md)
- [RESEARCH-ANALYSIS.md](/c:/Agent/CodexDashboard/Tracking/Task-0005/RESEARCH-ANALYSIS.md)
- [RESEARCH.md](/c:/Agent/CodexDashboard/Tracking/Task-0005/RESEARCH.md)
- [Cron System with JSON.md](/c:/Agent/CodexDashboard/Tracking/Task-0005/Research/Cron%20System%20with%20JSON.md)

## Next Step

Continue `PASS-0002`:

- add the first real workflow and worker path for `codex.exec.job`
- route `manual` and `webhook` backend entrypoints into that same durable run path
- prove at least one real `codex exec` executor path end to end
- preserve job id plus desired-state identity and Temporal runtime ids for auditability

## Watchouts

- do not let the dashboard own scheduler logic
- do not let startup reconcile become the only sync path
- do not widen this task into a full agent-graph or self-improvement system
- keep Git as desired state and Temporal as runtime truth
- the local compose stack and control-plane process are currently usable on this host for continued backend work
- `manual` and `webhook` are visible in read models now, but they do not execute yet

## References

- [TASK.md](/c:/Agent/CodexDashboard/Tracking/Task-0005/TASK.md)
- [PLAN.md](/c:/Agent/CodexDashboard/Tracking/Task-0005/PLAN.md)
- [RESEARCH.md](/c:/Agent/CodexDashboard/Tracking/Task-0005/RESEARCH.md)
- [Cron System with JSON.md](/c:/Agent/CodexDashboard/Tracking/Task-0005/Research/Cron%20System%20with%20JSON.md)
- [PASS-0000-AUDIT.md](/c:/Agent/CodexDashboard/Tracking/Task-0005/Testing/PASS-0000-AUDIT.md)
- [PASS-0001-AUDIT.md](/c:/Agent/CodexDashboard/Tracking/Task-0005/Testing/PASS-0001-AUDIT.md)
