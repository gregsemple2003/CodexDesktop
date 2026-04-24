# PASS-0002 Backend Smoke 0004

## Type

Unit test / server-only smoke / live validation-lane proof.

## Date

`2026-04-24`

## Scope

Proof for the next `PASS-0002` repair slice:

- `POST /api/v1/task-runs/{run_id}/retry-cleanup`
- cleanup-blocked runs can retry owned-lane restore through the backend
- successful cleanup retry converts the run into `interrupt_review`
- the same durable `follow_up` envelope now carries:
  - backend-worker follow-up after `poke`
  - human or supervisor cleanup repair
  - human or supervisor interrupt review

## Commands

```powershell
cd C:\Agent\CodexDashboard\backend\orchestration
go test ./...
powershell -ExecutionPolicy Bypass -File .\scripts\Start-ValidationLane.ps1

git -C C:\Agent\CodexDashboard worktree list

$runId = 'taskrun--Task-0008--active'
$validRoot = 'C:\Users\gregs\AppData\Local\Temp\cdxow\Task-0008-77be8767\18a967aea8e5c6dc\w'

Invoke-RestMethod -Method Post "http://127.0.0.1:14318/api/v1/task-runs/$runId/interrupt" | ConvertTo-Json -Depth 10

$repair = @'
{
  "repo_lane": {
    "owned_repo_root": "C:\\Users\\gregs\\AppData\\Local\\Temp\\cdxow\\Task-0008-77be8767\\18a967aea8e5c6dc\\w",
    "checkout_mode": "git_worktree_detached",
    "baseline_commit": "8cd1f82b70d1b393be52b39e2c66ee4bbab281f8",
    "approved_restore_commit": "8cd1f82b70d1b393be52b39e2c66ee4bbab281f8",
    "reset_status": "cleanup_blocked",
    "last_reset_target_commit": "8cd1f82b70d1b393be52b39e2c66ee4bbab281f8",
    "reset_failure_summary": "Owned repo root was repaired for cleanup retry."
  },
  "last_progress_summary": "Repaired the recorded owned checkout path before retrying cleanup."
}
'@

Invoke-RestMethod -Method Post `
  -Uri "http://127.0.0.1:14318/api/v1/task-runs/$runId/state" `
  -ContentType 'application/json' `
  -Body $repair | ConvertTo-Json -Depth 10

Invoke-RestMethod -Method Post "http://127.0.0.1:14318/api/v1/task-runs/$runId/retry-cleanup" | ConvertTo-Json -Depth 10
Invoke-RestMethod "http://127.0.0.1:14318/api/v1/tasks/Task-0008" | ConvertTo-Json -Depth 10
```

## Result

Pass.

`go test ./...` passed with the added coverage for:

- successful cleanup retry from `interrupt_cleanup_blocked` into terminal `interrupted`
- rejection of cleanup retry when the run is not in the cleanup-blocked state
- HTTP route coverage for `POST /api/v1/task-runs/{run_id}/retry-cleanup`

The live validation lane then proved:

1. An active run with an intentionally bad `repo_lane.owned_repo_root` hit `interrupt_cleanup_blocked` again on `POST /interrupt`.
2. The run carried:
   - `follow_up.kind = cleanup_repair`
   - `follow_up.status = pending`
   - `repo_lane.reset_status = cleanup_blocked`
3. After restoring the recorded owned-checkout path through `POST /state`, `POST /retry-cleanup` returned:
   - `status = interrupted`
   - `reason_code = interrupt_cleanup_repaired`
   - `repo_lane.reset_status = restored`
   - `follow_up.kind = interrupt_review`
   - `follow_up.status = pending`
4. The next task read returned:
   - `current_story.status = no_active_run`
   - `dispatch_readiness.ready = true`
   - task-level `dispatch.allowed = true`

## Why This Counts

This closes the biggest remaining cleanup gap inside `PASS-0002`.

Cleanup-blocked is no longer just an honest error shape. The backend can now retry cleanup through a dedicated control path, repair the owned checkout, release live-story ownership, and carry the next human obligation through the same durable `follow_up` model used elsewhere in the runtime contract.

That makes cleanup repair part of the backend-owned supervision flow instead of leaving it entirely outside the contract.

## Limitations

- The proof used `git worktree list` plus a direct state update to restore the recorded owned-checkout path before retrying cleanup. That was enough to prove the backend retry path, but it is still a controlled validation setup rather than a full autonomous repair loop.
- `interrupt_review` is now durable and explicit, but there is not yet a separate decision endpoint that resolves that review without relying on later dispatch or broader run lifecycle changes.
- The validation-lane runner can still show intermittent stdout-log lock churn during restart, although the fresh control plane served the proof requests correctly.
- This is not dashboard regression proof.
