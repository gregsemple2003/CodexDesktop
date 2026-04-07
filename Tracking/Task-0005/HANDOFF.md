# Task 0005 Handoff

## Current Status

`PASS-0001` is complete. `PASS-0002` is active again after a bounded runtime fence fix: the first live proof attempt released all three tracked schedules at startup and hit `CreateProcessAsUserW failed: 1920`, and the control-plane now carries an explicit short schedule catchup window so existing schedules will be updated away from that unsafe startup behavior on the next reconcile.

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

- a bounded live proof that reaches exactly one intended real Codex execution
- end-to-end evidence that the chosen trigger path records Temporal ids plus desired spec hash
- confirmation of whether the Windows Codex runtime failure `CreateProcessAsUserW failed: 1920` still occurs after the bounded retry

Research output captured so far:

- [RESEARCH-PLAN.md](/c:/Agent/CodexDashboard/Tracking/Task-0005/RESEARCH-PLAN.md)
- [RESEARCH-ANALYSIS.md](/c:/Agent/CodexDashboard/Tracking/Task-0005/RESEARCH-ANALYSIS.md)
- [RESEARCH.md](/c:/Agent/CodexDashboard/Tracking/Task-0005/RESEARCH.md)
- [Cron System with JSON.md](/c:/Agent/CodexDashboard/Tracking/Task-0005/Research/Cron%20System%20with%20JSON.md)

## Next Step

Continue `PASS-0002`:

1. Start the control-plane against the local Temporal stack so reconcile updates the managed schedules onto the new short catchup window.
2. Verify the startup path no longer releases the three overdue digest runs.
3. Drive one bounded real manual run and capture the resulting Temporal ids, spec hash, and Codex artifacts.
4. If the run still fails with `CreateProcessAsUserW failed: 1920`, preserve that executor failure as the honest PASS-0002 blocker; otherwise close `PASS-0002` and continue into `PASS-0003`.

## Watchouts

- do not let the dashboard own scheduler logic
- do not let startup reconcile become the only sync path
- do not widen this task into a full agent-graph or self-improvement system
- keep Git as desired state and Temporal as runtime truth
- the contained proof launch on `2026-04-07` showed that starting the worker-backed control-plane can release all three overdue daily digest schedules at once unless the schedule path is fenced
- the fence fix is now in code and covered by backend tests; existing schedules will pick it up on the next reconcile because catchup window drift is now part of desired-vs-runtime comparison
- the resulting real Codex runs wrote artifacts under `C:\Users\gregs\AppData\Local\CodexDashboard\orchestration-runs\...` and failed with `CreateProcessAsUserW failed: 1920`
- `go`, `docker`, and `temporal` are available on this host, but this shell needed PATH refresh or explicit executable resolution
- `manual` and `webhook` already route into the durable workflow path; the remaining task is bounded live proof plus honest handling of the Windows Codex sandbox behavior

## References

- [TASK.md](/c:/Agent/CodexDashboard/Tracking/Task-0005/TASK.md)
- [PLAN.md](/c:/Agent/CodexDashboard/Tracking/Task-0005/PLAN.md)
- [RESEARCH.md](/c:/Agent/CodexDashboard/Tracking/Task-0005/RESEARCH.md)
- [Cron System with JSON.md](/c:/Agent/CodexDashboard/Tracking/Task-0005/Research/Cron%20System%20with%20JSON.md)
- [PASS-0000-AUDIT.md](/c:/Agent/CodexDashboard/Tracking/Task-0005/Testing/PASS-0000-AUDIT.md)
- [PASS-0001-AUDIT.md](/c:/Agent/CodexDashboard/Tracking/Task-0005/Testing/PASS-0001-AUDIT.md)
- [BUG-0001.md](/c:/Agent/CodexDashboard/Tracking/Task-0005/BUG-0001.md)
