# CodexDashboard Regression Checklist

This file is the canonical regression matrix for CodexDashboard.

## Canonical Rule

Task-level regression in this repo starts from the real desktop app surface.

Supporting parser or CLI proof is useful, but it does not replace the required app-surface lane for closure.

Regression must not use the human's personal dashboard lane, config, database, or live Codex data unless the human explicitly authorizes that specific run.
The default or persistent backend lane is not automatically the human lane, but task-closure regression still must run on a separate isolated validation or task-specific regression lane.
Use an isolated lane with task-owned or temp data, task-owned config, and task-owned SQLite persistence.
The current default agent validation lane is documented in [TESTING.md](./TESTING.md):

- backend URL: `http://127.0.0.1:14318`
- Temporal: `127.0.0.1:17233`
- Postgres: `15432`
- runtime root: `%LOCALAPPDATA%\CodexDashboard\orchestration-validation-lane`

If a task creates another regression lane, its ports, data roots, start/stop commands, and cleanup expectations must be documented in [TESTING.md](./TESTING.md) before that lane is used as closure evidence.

When current human-facing functionality changes, this repo-local `REGRESSION.md` must be updated so the changed interaction is covered by a named case. Do not treat a new clickable surface as covered by an older case unless the written steps and expected result explicitly include that interaction.

## Default Regression Lane

### REG-001 Desktop Overlay Launch And Data Smoke

Goal:

Confirm the real desktop app can ingest Codex telemetry, open the real overlay, and render interval data plus budget state.

Steps:

1. Launch the real app from repo root.
2. Point it at a task-owned fixture tree with real-shaped `token_count` events. Do not use `C:\Users\gregs\.codex` unless the human explicitly authorizes that run.
3. Allow at least one ingest cycle to complete.
4. Trigger the real overlay path.
5. Capture an artifact from the running app surface that shows:
   - the selected interval
   - bar data
   - weekly budget state
6. Exit cleanly.

Expected result:

- the app starts without crashing
- the ingest loop persists real token events
- the overlay becomes visible
- interval data appears
- budget and redline state appear

## Required Current Cases

### REG-002 Jobs Tab Interaction And Status Surface

Goal:

Confirm the real desktop app can switch from the default `Usage` surface to the backend-backed `Jobs` surface, show desired-vs-runtime job state from the orchestration backend, and invoke the bounded `Sync now` control without disruptive side effects.

Steps:

1. Launch the real app from repo root.
2. Point it at a task-owned fixture tree with real-shaped `token_count` events. Do not use `C:\Users\gregs\.codex` unless the human explicitly authorizes that run.
3. Start the separate validation-lane orchestration backend and keep it reachable for the app without disturbing the always-on service lane.
4. Set `CODEX_DASHBOARD_JOBS_BACKEND_URL=http://127.0.0.1:14318` for the app process that will be used for the regression run.
5. Allow at least one ingest cycle to complete.
6. Trigger the real overlay path.
7. Click the `Jobs` tab in the running overlay.
8. Verify the tab switch completes immediately and does not trigger `Sync now` work on its own.
9. Verify the `Jobs` surface shows backend-derived job rows with trigger labels, desired/runtime state, and drift status.
10. Open `Details` for one job and verify it shows the backend job payload, including trigger and executor data.
11. Verify `Run now` is visible for manual-capable jobs and disabled for jobs without a manual trigger.
12. Click `Sync now` in the `Jobs` surface and observe the backend-backed state reload to completion.
13. Capture an artifact from the running app surface that shows:
   - the `Jobs` tab as the active surface
   - job summary counts
   - per-job status rows
14. Exit cleanly.

Expected result:

- the tab switch completes without hitching the overlay
- no extra console or app windows are spawned by the interaction
- clicking `Jobs` alone does not trigger `Sync now`
- the `Jobs` surface renders backend-derived jobs, in-sync count, needs-attention count, and per-job rows
- `Details` shows the backend job payload rather than a local Windows task snapshot
- `Sync now` completes through the backend without crashing or spawning stray windows
- `Run now` availability matches whether a job actually supports manual triggering
- visible copy remains human-facing
- the validation lane can be exercised without taking down the default service lane

### REG-003 Tasks Tab Committed-Work Surface

Goal:

Confirm the real desktop app can switch from the default `Usage` surface to the committed-work `Tasks` surface, render backend-shaped task readback from an isolated lane or task-owned fixture, and communicate task state without exposing unpromoted candidates or false progress bars.

Steps:

1. Launch the real app from repo root.
2. Point it at a task-owned fixture tree with real-shaped `token_count` events. Do not use `C:\Users\gregs\.codex` unless the human explicitly authorizes that run.
3. Use an isolated config and isolated SQLite database as documented in [TESTING.md](./TESTING.md).
4. Point task readback at the validation lane URL `http://127.0.0.1:14318` or a task-owned backend-shaped snapshot fixture through `CODEX_DASHBOARD_TASKS_SNAPSHOT_PATH`.
5. Trigger the real overlay path.
6. Click the `Tasks` tab in the running overlay.
7. Verify the tab switch completes immediately and does not dispatch, pause, poke, resume, or otherwise change backend run state by switching tabs.
8. Verify the `Tasks` surface shows:
   - `Needs you`
   - `Sleeping`
   - `Running`
   - `Blocked`
   - `Ready`
9. Verify committed-task rows show title, state, reason, freshness, source provenance, and a selected-task detail pane.
10. Verify unpromoted candidates are not displayed as normal tasks.
11. Verify committed promoted work uses provenance such as `Promoted from Dream` or `Promoted from Review`, not `Candidate` or `Prov: Candidate`.
12. Verify visible stop/hold copy says `Pause` rather than backend-internal `Interrupt`.
13. Verify the surface does not display AI-run progress bars.
14. Capture an artifact from the running app surface that shows:
   - the `Tasks` tab as the active surface
   - committed-work summary counts
   - grouped task rows
   - the selected-task detail pane
15. Exit cleanly.

Expected result:

- the tab switch completes without hitching the overlay
- no extra console or app windows are spawned by the interaction
- switching to `Tasks` is read-only unless the human explicitly clicks a bounded action
- summary counts and rows render from isolated backend-shaped task readback
- candidate/intake work is absent unless it has been promoted into committed work
- promoted committed-work provenance is explicit and does not use candidate labels
- visible control language uses `Pause`, not `Interrupt`
- no progress bar implies false precision for AI task-run progress
- the validation lane or fixture can be exercised without using the persistent service lane, the human's dashboard config, the human's active database, or live Codex data

## Supporting Smoke

### SMOKE-001 Ingest Core

Run:

```powershell
python -m app.codex_dashboard --scan-once --print-summary
```

Expected result:

- the app reads real-shaped telemetry
- SQLite persistence is updated
- a human-readable summary prints

Interpretation:

- this is supporting proof only
- it does not replace `REG-001`

### SMOKE-002 Service Lane Release Isolation

Goal:

Confirm the human service lane is pinned to a promoted release and cannot be
advanced by merely updating or editing the mutable repo checkout.

Precondition:

Only run against the human service lane after the human explicitly authorizes
that lane to be inspected or restarted for this specific run.

Steps:

1. Run the unit tests for service-lane scripts:
   `python -m unittest tests.test_service_lane_scripts -v`
2. Publish the intended service-lane release with
   `backend/orchestration/scripts/Publish-ServiceLaneRelease.ps1`.
3. Restart or install the service lane through the repo scripts only after the
   release has been pinned.
4. Run `backend/orchestration/scripts/Test-ServiceLaneIsolation.ps1`.
5. Run `backend/orchestration/scripts/Get-ServiceLaneStatus.ps1`.
6. Verify no live service-lane runner command line points at the repo-local
   `Run-OrchestrationLane.ps1`.

Expected result:

- the scheduled task uses the runtime-root launcher
- `current-release.json` exists and validates binary and compose-file hashes
- the running process path matches the pinned release binary path
- backend health is reachable
- any dirty-source promotion is explicitly visible in the release manifest

Interpretation:

- this is required operator proof for human-lane release claims
- this is not a substitute for task-level desktop-app regression cases

### SMOKE-003 Dashboard Frontend Release Isolation

Goal:

Confirm the human-facing desktop dashboard frontend is pinned to a promoted
release and cannot be advanced by merely updating or editing the mutable repo
checkout.

Precondition:

Only run against the human dashboard frontend after the human explicitly
authorizes that lane to be inspected or restarted for this specific run.

Steps:

1. Run the unit tests for dashboard release scripts:
   `python -m unittest tests.test_dashboard_release_scripts tests.test_desktop_support -v`
2. Publish the intended dashboard release with `scripts/Publish-DashboardRelease.ps1`.
3. Restart the dashboard through `scripts/Start-DashboardRelease.ps1` or the
   installed runtime launcher.
4. Run `scripts/Test-DashboardRelease.ps1`.
5. Verify no live dashboard command line is `pythonw -m app.codex_dashboard`
   from `C:\Agent\CodexDashboard` without a release id and release root.
6. For visible-surface claims, run an app smoke artifact capture against the
   pinned release and verify the expected tab is active in the generated
   `overlay-summary.txt`.

Expected result:

- `dashboard-current-release.json` exists and validates copied source hashes
- the startup file points at the runtime-root dashboard launcher
- the running dashboard process includes the pinned release id and release root
- any dirty-source promotion is explicitly visible in the release manifest
- visible-surface claims cite a smoke artifact, not backend-only proof

Interpretation:

- this is required operator proof for human-facing dashboard frontend release
  claims
- this is not a substitute for task-level desktop-app regression cases
