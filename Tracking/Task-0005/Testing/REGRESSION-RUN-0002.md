# Regression Run 0002

Test type: `regression test`

Claimed lane:

- `REG-002 Jobs Tab Interaction And Status Surface`

Actual flow exercised:

1. Kept the default service lane running on `http://127.0.0.1:4318`.
2. Started the separate validation lane on `http://127.0.0.1:14318`.
3. Launched the real desktop app with the validation-lane backend override:
   - `CODEX_DASHBOARD_JOBS_BACKEND_URL=http://127.0.0.1:14318`
   - `python -m app.codex_dashboard --smoke-artifact-dir 'C:\Agent\CodexDashboard\Tracking\Task-0005\Testing\REGRESSION-RUN-0002-jobs-validation-smoke' --smoke-tab jobs`
4. Let the smoke harness select the real `Jobs` tab, invoke the real `Sync now` widget command, and capture the live overlay.

Why this counts:

- the real Tk app was launched again, not a mocked widget tree
- the `Jobs` surface rendered live backend-derived rows from the validation lane
- the captured overlay summary proves the app process was talking to `http://127.0.0.1:14318`, not the default service lane URL
- the service lane remained running at `http://127.0.0.1:4318`, so this run directly proves the two-lane model can keep regression from disturbing scheduled jobs

Disqualifiers / limitations:

- the smoke harness still invokes live widget commands programmatically instead of using a literal human click
- this rerun only covered `REG-002` because the service-lane follow-up changed backend-lane targeting rather than the general `Usage` surface already covered in [REGRESSION-RUN-0001.md](/c:/Agent/CodexDashboard/Tracking/Task-0005/Testing/REGRESSION-RUN-0001.md)

## Environment

- desktop app repo: `C:\Agent\CodexDashboard`
- service lane backend: `http://127.0.0.1:4318`
- service lane Temporal gRPC: `127.0.0.1:7233`
- validation lane backend: `http://127.0.0.1:14318`
- validation lane Temporal gRPC: `127.0.0.1:17233`

## Results

### REG-002

Result: `passed`

Artifacts:

- [overlay.png](/c:/Agent/CodexDashboard/Tracking/Task-0005/Testing/REGRESSION-RUN-0002-jobs-validation-smoke/overlay.png)
- [overlay-summary.txt](/c:/Agent/CodexDashboard/Tracking/Task-0005/Testing/REGRESSION-RUN-0002-jobs-validation-smoke/overlay-summary.txt)

Observed:

- `active_tab=jobs`
- status line: `Jobs sync completed. 0 schedule changes.`
- `jobs_backend=http://127.0.0.1:14318`
- summary cards showed `03` declared, `03` in sync, `00` needs attention
- the service lane stayed healthy on the default ports throughout the run
