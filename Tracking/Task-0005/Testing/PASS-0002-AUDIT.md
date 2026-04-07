# Pass 0002 Audit

## Scope

`PASS-0002` completed the first durable execution slice for `Task-0005`:

- the worker-hosted `codex.exec.job` workflow and `codex.exec.activity`
- `schedule`, `manual`, and `webhook` trigger entrypoints routed into the same durable run path
- bounded startup reconcile behavior for managed schedules
- a real `codex exec` proof with captured job id, desired spec hash, Temporal workflow identity, and per-run artifacts

This pass still does not claim the dashboard-backed Jobs surface or repo-root desktop regression. Those remain `PASS-0003` work.

## Verification

Executed from `C:\Agent\CodexDashboard\backend\orchestration` unless noted otherwise:

```powershell
& 'C:\Program Files\Go\bin\go.exe' test ./...
python -m unittest discover -s tests -p "test_*.py" -v
python -m app.codex_dashboard --scan-once --print-summary
Invoke-WebRequest http://127.0.0.1:4318/healthz | Select-Object -ExpandProperty Content
& 'C:\Users\gregs\AppData\Local\Microsoft\WinGet\Links\temporal.exe' workflow describe --workflow-id 'job/codex-daily-agentic-swe-digest/manual/18a81bff-3355-418e-8559-224407f7f586' --address 127.0.0.1:7233
Get-Content -Raw 'C:\Users\gregs\AppData\Local\CodexDashboard\orchestration-runs\pass-0002-final-retry-20260407-104944\codex-daily-agentic-swe-digest\job_codex-daily-agentic-swe-digest_manual_18a81bff-3355-418e-8559-224407f7f586\final-message.txt'
Get-Content 'C:\Users\gregs\.codex\gmail-digest-email\sent-log.jsonl' -Tail 1
```

Observed result:

- `go test ./...` passed across the control-plane, HTTP, Temporal adapter, and `jobexec` packages
- repo-root Python unit tests still passed: `56` tests
- the supporting ingest smoke still completed successfully against the live telemetry tree
- `/healthz` reported `status: ok`, `job_count: 3`, and a successful last-sync timestamp after the bounded retry startup reconcile
- the final manual proof workflow completed successfully:
  - workflow id: `job/codex-daily-agentic-swe-digest/manual/18a81bff-3355-418e-8559-224407f7f586`
  - run id: `019d686c-3c82-7a12-81f2-50d60cd5ffd5`
  - workflow type: `codex.exec.job`
  - status: `COMPLETED`
- the Temporal result preserved the full executor command, including:
  - job id: `codex-daily-agentic-swe-digest`
  - trigger type: `manual`
  - desired spec hash: `38a97b5d67aa3cf10a8231b972e1f3404fe29c28be37bac1e54e2ce378524b62`
  - workflow id and run id
  - per-run artifact paths for `events.jsonl`, `stderr.txt`, and `final-message.txt`
- the successful run wrote:
  - report: `C:\Users\gregs\.codex\reports\2026-04-07-SWE-Orchestration.md`
  - final message: `C:\Users\gregs\AppData\Local\CodexDashboard\orchestration-runs\pass-0002-final-retry-20260407-104944\codex-daily-agentic-swe-digest\job_codex-daily-agentic-swe-digest_manual_18a81bff-3355-418e-8559-224407f7f586\final-message.txt`
- the configured Gmail digest mailer recorded a successful send at `2026-04-07T10:58:44`

## Requirement Mapping

| Requirement | Evidence | Result |
| --- | --- | --- |
| `schedule`, `manual`, and `webhook` can each start durable runs through Temporal | [controlplane.go](/c:/Agent/CodexDashboard/backend/orchestration/internal/controlplane/controlplane.go); [mux.go](/c:/Agent/CodexDashboard/backend/orchestration/internal/httpapi/mux.go); [mux_test.go](/c:/Agent/CodexDashboard/backend/orchestration/internal/httpapi/mux_test.go); [controlplane_test.go](/c:/Agent/CodexDashboard/backend/orchestration/internal/controlplane/controlplane_test.go) | Passed |
| At least one real `codex exec` executor path works end to end | successful manual workflow `job/codex-daily-agentic-swe-digest/manual/18a81bff-3355-418e-8559-224407f7f586`; `final-message.txt`; Gmail sent-log entry | Passed |
| Run records preserve job id, desired spec hash, and Temporal ids | Temporal workflow result from `workflow describe`; [jobexec.go](/c:/Agent/CodexDashboard/backend/orchestration/internal/jobexec/jobexec.go) | Passed |
| Startup reconcile remains bounded enough for controlled proof work | [controlplane.go](/c:/Agent/CodexDashboard/backend/orchestration/internal/controlplane/controlplane.go); `/healthz`; successful final retry run root under `pass-0002-final-retry-20260407-104944` | Passed |
| Focused automated proof exists for trigger routing and executor command assembly | [mux_test.go](/c:/Agent/CodexDashboard/backend/orchestration/internal/httpapi/mux_test.go); [jobexec_test.go](/c:/Agent/CodexDashboard/backend/orchestration/internal/jobexec/jobexec_test.go); `go test ./...` | Passed |
| Backend execution remains independent of the Tk process | all live proof steps ran through Temporal, HTTP, and `codex exec` without launching the desktop UI | Passed |

## Caveats

The successful Windows executor path currently relies on `codex exec --dangerously-bypass-approvals-and-sandbox`.

That is an intentional `PASS-0002` unblocker for this host, not a claim that the safer sandboxed path is fixed. The earlier `CreateProcessAsUserW failed: 1920` failure and the startup over-release issue are preserved in [BUG-0001.md](/c:/Agent/CodexDashboard/Tracking/Task-0005/BUG-0001.md).

The live digest job also installed the Gmail skill's declared Python dependencies into the user Python environment and touched `.codex` report and mailer state as part of the real `--email` workflow. That is honest job-side behavior for this proof, but it should be treated as operator/runtime state rather than as CodexDashboard product code.

Repo-root desktop regression was not run in this pass because the dashboard integration work is still pending in `PASS-0003`.

## Verdict

`ready_with_caveats`
