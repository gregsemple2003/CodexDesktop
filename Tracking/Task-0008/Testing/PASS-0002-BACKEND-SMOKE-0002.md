# PASS-0002 Backend Smoke 0002

## Type

Unit test / server-only smoke / live validation-lane proof.

## Date

`2026-04-24`

## Scope

Proof for the next `PASS-0002` supervision refinements:

- stale human-wait supervision
- explicit `waiting_for_human_stale` poke rejection
- dedicated cleanup-blocked readback on interrupt failure
- repo-lane reset failure details and run-level failure summary

## Commands

```powershell
cd C:\Agent\CodexDashboard\backend\orchestration
go test ./...
powershell -ExecutionPolicy Bypass -File .\scripts\Start-ValidationLane.ps1

$body = @'
{
  "state": "waiting_for_human",
  "reason_code": "review_required",
  "state_summary": "Run is waiting for human review.",
  "next_owner": "human",
  "next_expected_event": "Approve the next backend action.",
  "last_progress_summary": "Run moved into a human review wait.",
  "wait_contract": {
    "waiting_on": "human_review",
    "why_blocked": "The next backend action needs human approval.",
    "resume_when": "The human approves the next backend action.",
    "human_action_required": true,
    "human_action_target": {
      "kind": "approval_action",
      "label": "Approve backend review step",
      "uri": "approval://taskrun/Task-0008"
    },
    "stale_after": "2026-04-24T21:00:00Z"
  }
}
'@

Invoke-RestMethod -Method Post `
  -Uri http://127.0.0.1:14318/api/v1/task-runs/taskrun--Task-0008--active/state `
  -ContentType 'application/json' `
  -Body $body | ConvertTo-Json -Depth 10

Invoke-RestMethod http://127.0.0.1:14318/api/v1/task-runs/taskrun--Task-0008--active | ConvertTo-Json -Depth 10
Invoke-RestMethod http://127.0.0.1:14318/api/v1/tasks/Task-0008 | ConvertTo-Json -Depth 10

$badLane = @'
{
  "state": "running",
  "reason_code": "worker_started",
  "state_summary": "Run is actively executing in the owned checkout.",
  "next_owner": "backend",
  "next_expected_event": "Execution worker records the next progress checkpoint.",
  "repo_lane": {
    "owned_repo_root": "C:\\Agent\\CodexDashboard",
    "checkout_mode": "git_worktree_detached",
    "baseline_commit": "8cd1f82b70d1b393be52b39e2c66ee4bbab281f8",
    "approved_restore_commit": "8cd1f82b70d1b393be52b39e2c66ee4bbab281f8",
    "reset_status": "not_run"
  },
  "last_progress_summary": "Run resumed with a deliberately invalid owned lane for cleanup proof."
}
'@

Invoke-RestMethod -Method Post `
  -Uri http://127.0.0.1:14318/api/v1/task-runs/taskrun--Task-0008--active/state `
  -ContentType 'application/json' `
  -Body $badLane | ConvertTo-Json -Depth 10

Invoke-RestMethod -Method Post http://127.0.0.1:14318/api/v1/task-runs/taskrun--Task-0008--active/interrupt | ConvertTo-Json -Depth 10
Invoke-RestMethod http://127.0.0.1:14318/api/v1/task-runs/taskrun--Task-0008--active | ConvertTo-Json -Depth 10
```

## Result

Pass.

`go test ./...` passed with the added coverage for:

- stale human-wait escalation into `human_wait_stale`
- cleanup-blocked interrupt readback
- wait-contract clearing when a run leaves `waiting_for_human`
- first-class repo-lane reset failure fields

The live validation lane then proved both new paths:

1. A run was moved into `waiting_for_human` with an already-expired `wait_contract.stale_after`.
2. The next read-through query returned:
   - `reason_code = human_wait_stale`
   - `attention_level = urgent`
   - `next_owner = human_or_supervisor`
   - `actions.poke.block_reasons[0].code = waiting_for_human_stale`
   - `actions.interrupt.allowed = true`
3. The same run was then given an intentionally invalid `repo_lane.owned_repo_root` outside the backend-owned lane root.
4. `POST /api/v1/task-runs/{run_id}/interrupt` returned a blocked cleanup result with:
   - `reason_code = interrupt_cleanup_blocked`
   - `repo_lane.reset_status = cleanup_blocked`
   - `repo_lane.last_reset_target_commit` populated
   - `repo_lane.reset_failure_summary` populated
   - top-level `failure_summary` populated
5. After the state-hygiene fix, the cleanup-blocked run no longer carried the old human-wait contract into the blocked state.

## Why This Counts

This turns two previously hand-wavy supervision edges into explicit backend truth:

- stale human waits now escalate into a distinct urgent state with durable action gating
- reset failures now surface as first-class repo-lane and run-level readback instead of hiding inside generic error text

That makes the backend much closer to the humane monitoring contract the later `Tasks` tab expects.

## Limitations

- The validation-lane runner showed intermittent stdout-log lock churn during restart, but the fresh control plane still served the proof requests correctly.
- `poke` still records intervention rather than driving a worker-side recovery action automatically.
- Cleanup-blocked is now explicit in readback, but there is not yet a richer repair workflow beyond human or supervisor review.
- This is not dashboard regression proof.
