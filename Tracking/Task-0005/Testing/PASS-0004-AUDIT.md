# PASS-0004 Audit

`PASS-0004` corrected the task's operating-model gap by turning the backend from a proof-only stack into a usable local service lane with a separate disposable validation lane.

Delivered in this pass:

- parameterized the repo-local Temporal/Postgres compose stack so different runtime lanes can use different ports
- added service-lane scripts under `backend/orchestration/scripts/` to install, start, stop, and inspect the always-on backend
- added validation-lane scripts so tests and regression can run on separate ports without disturbing the service lane
- taught the desktop app to honor `CODEX_DASHBOARD_JOBS_BACKEND_URL` so regression can target the validation lane explicitly
- updated task and repo docs to make the two-lane local operating model explicit

## Commands Run

```powershell
cd C:\Agent\CodexDashboard\backend\orchestration
powershell -NoProfile -ExecutionPolicy Bypass -File .\scripts\Install-ServiceLane.ps1
powershell -NoProfile -ExecutionPolicy Bypass -File .\scripts\Start-ServiceLane.ps1
powershell -NoProfile -ExecutionPolicy Bypass -File .\scripts\Start-ValidationLane.ps1
& 'C:\Program Files\Go\bin\go.exe' test ./...
python -m unittest discover -s tests -p "test_*.py" -v
python -m app.codex_dashboard --scan-once --print-summary
```

Desktop smoke against the validation lane:

```powershell
$env:CODEX_DASHBOARD_JOBS_BACKEND_URL = "http://127.0.0.1:14318"
python -m app.codex_dashboard --smoke-artifact-dir 'C:\Agent\CodexDashboard\Tracking\Task-0005\Testing\REGRESSION-RUN-0002-jobs-validation-smoke' --smoke-tab jobs
Remove-Item Env:CODEX_DASHBOARD_JOBS_BACKEND_URL -ErrorAction SilentlyContinue
```

## Results

- `Install-ServiceLane.ps1` registered `CodexDashboard-Orchestration-ServiceLane` and the service lane came up healthy on `127.0.0.1:4318` with three in-sync jobs
- `Get-ServiceLaneStatus.ps1` showed the service lane still running while the validation lane was brought up on `127.0.0.1:14318`
- the validation lane's `/healthz` and `/api/v1/jobs` both returned `status: ok`, `job_count: 3`, and the expected next scheduled action time of `2026-04-08T08:00:00Z` for all three tracked jobs
- the validation-lane jobs smoke passed and the captured overlay summary recorded `jobs_backend=http://127.0.0.1:14318`
- `go test ./...` passed across the backend packages
- `python -m unittest discover -s tests -p "test_*.py" -v` passed with 63 tests
- `python -m app.codex_dashboard --scan-once --print-summary` passed

## Exit Bar Review

| Claim | Evidence | Result |
| --- | --- | --- |
| The local machine has a real always-on service lane for scheduled jobs | `Install-ServiceLane.ps1`; `Start-ServiceLane.ps1`; live service-lane status with `task_state: Running`, `process_running: true`, and healthy `/healthz` | Passed |
| Validation work can use a separate runtime lane without taking down the service lane | live service-lane status stayed healthy while validation lane ran on `14318/17233`; [REGRESSION-RUN-0002.md](/c:/Agent/CodexDashboard/Tracking/Task-0005/Testing/REGRESSION-RUN-0002.md) | Passed |
| The desktop app can target the validation lane explicitly | [jobs_backend.py](/c:/Agent/CodexDashboard/app/codex_dashboard/jobs_backend.py); [overlay-summary.txt](/c:/Agent/CodexDashboard/Tracking/Task-0005/Testing/REGRESSION-RUN-0002-jobs-validation-smoke/overlay-summary.txt) | Passed |
| Repo docs and task artifacts describe the two-lane model honestly | [README.md](/c:/Agent/CodexDashboard/backend/orchestration/README.md); [TESTING.md](/c:/Agent/CodexDashboard/TESTING.md); [REGRESSION.md](/c:/Agent/CodexDashboard/REGRESSION.md); [TASK.md](/c:/Agent/CodexDashboard/Tracking/Task-0005/TASK.md); [PLAN.md](/c:/Agent/CodexDashboard/Tracking/Task-0005/PLAN.md) | Passed |

## Caveats

- the service lane intentionally runs under the current interactive Windows user rather than as a machine-wide system service because the tracked `.codex` desired state and the logged-in `codex` CLI session live in that user profile
- the validation lane is disposable by design and should be torn down after proof work
- the earlier executor caveat still applies: the successful live `codex exec` path currently depends on `--dangerously-bypass-approvals-and-sandbox`, which remains tracked in [BUG-0001.md](/c:/Agent/CodexDashboard/Tracking/Task-0005/BUG-0001.md)
