# Task 0005 Handoff

## Current Status

`Task-0005` research is complete and the task is now at the planning gate pending explicit human approval of [PLAN.md](/c:/Agent/CodexDashboard/Tracking/Task-0005/PLAN.md).

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

Research output captured so far:

- [RESEARCH-PLAN.md](/c:/Agent/CodexDashboard/Tracking/Task-0005/RESEARCH-PLAN.md)
- [RESEARCH-ANALYSIS.md](/c:/Agent/CodexDashboard/Tracking/Task-0005/RESEARCH-ANALYSIS.md)
- [RESEARCH.md](/c:/Agent/CodexDashboard/Tracking/Task-0005/RESEARCH.md)
- [Cron System with JSON.md](/c:/Agent/CodexDashboard/Tracking/Task-0005/Research/Cron%20System%20with%20JSON.md)

## Next Step

Wait for explicit human approval of [PLAN.md](/c:/Agent/CodexDashboard/Tracking/Task-0005/PLAN.md).

After approval, start `PASS-0000`:

- establish the v1 job-spec layout under `.codex/Orchestration/Jobs`
- scaffold `backend/orchestration/`
- document the repo-local Temporal plus Postgres dev stack
- resolve or honestly record the missing `go`, `docker`, and `temporal` toolchain gap before claiming a runnable backend

## Watchouts

- do not let the dashboard own scheduler logic
- do not let startup reconcile become the only sync path
- do not widen this task into a full agent-graph or self-improvement system
- keep Git as desired state and Temporal as runtime truth
- local `go`, `docker`, and `temporal` binaries are not currently on `PATH`

## References

- [TASK.md](/c:/Agent/CodexDashboard/Tracking/Task-0005/TASK.md)
- [PLAN.md](/c:/Agent/CodexDashboard/Tracking/Task-0005/PLAN.md)
- [RESEARCH.md](/c:/Agent/CodexDashboard/Tracking/Task-0005/RESEARCH.md)
- [Cron System with JSON.md](/c:/Agent/CodexDashboard/Tracking/Task-0005/Research/Cron%20System%20with%20JSON.md)
