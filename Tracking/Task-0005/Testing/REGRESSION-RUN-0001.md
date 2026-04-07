# Regression Run 0001

Test type: `regression test`

Claimed lane:

- `REG-001 Desktop Overlay Launch And Data Smoke`
- `REG-002 Jobs Tab Interaction And Status Surface`

Actual flow exercised:

1. Launched the real desktop app with smoke artifact capture for the default `Usage` lane:
   - `python -m app.codex_dashboard --smoke-artifact-dir 'C:\Agent\CodexDashboard\Tracking\Task-0005\Testing\REGRESSION-RUN-0001-usage-smoke' --smoke-tab usage`
2. Launched the real desktop app again with smoke artifact capture for the `Jobs` lane:
   - `python -m app.codex_dashboard --smoke-artifact-dir 'C:\Agent\CodexDashboard\Tracking\Task-0005\Testing\REGRESSION-RUN-0001-jobs-smoke' --smoke-tab jobs`
3. In the `Jobs` smoke run, the live app selected the real `Jobs` tab and invoked the real `Sync now` widget command before capturing the overlay.

Why this counts:

- both runs started the real Tk app and opened the real overlay
- both runs captured the live overlay window to PNG, not just a text-only backend dump
- the `Jobs` run rendered the backend-backed rows, visible action buttons, and the post-sync status line on the real app surface
- the `Jobs` run used the actual widget command path for `Sync now`, so the bounded backend action was exercised from inside the live UI rather than from a backend-only script

Disqualifiers / limitations:

- the smoke harness invokes live widget commands programmatically instead of requiring a literal human mouse click
- the `Run now` button was not executed in regression because that path is paid/email-backed on the current live job corpus; it was already proven end to end in `PASS-0002`

## Environment

- desktop app repo: `C:\Agent\CodexDashboard`
- live backend: `http://127.0.0.1:4318`
- Temporal gRPC: `127.0.0.1:7233`

## Results

### REG-001

Result: `passed`

Artifacts:

- [overlay.png](/c:/Agent/CodexDashboard/Tracking/Task-0005/Testing/REGRESSION-RUN-0001-usage-smoke/overlay.png)
- [overlay-summary.txt](/c:/Agent/CodexDashboard/Tracking/Task-0005/Testing/REGRESSION-RUN-0001-usage-smoke/overlay-summary.txt)
- [overlay-chart.ps](/c:/Agent/CodexDashboard/Tracking/Task-0005/Testing/REGRESSION-RUN-0001-usage-smoke/overlay-chart.ps)

Observed:

- `active_tab=usage`
- overlay visible without fallback-only behavior
- interval, budget, and redline state were rendered on the live surface

### REG-002

Result: `passed`

Artifacts:

- [overlay.png](/c:/Agent/CodexDashboard/Tracking/Task-0005/Testing/REGRESSION-RUN-0001-jobs-smoke/overlay.png)
- [overlay-summary.txt](/c:/Agent/CodexDashboard/Tracking/Task-0005/Testing/REGRESSION-RUN-0001-jobs-smoke/overlay-summary.txt)

Observed:

- `active_tab=jobs`
- status line: `Jobs sync completed. 0 schedule changes.`
- summary cards showed `03` declared, `03` in sync, `00` needs attention
- rows rendered trigger labels, desired/runtime state, status chips, and `Run now`/`Details` actions
- the webhook-only job showed a disabled `Run now` control while manual-capable jobs showed enabled `Run now`
