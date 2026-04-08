# Task 0004 Handoff

## Current Status

`Task-0004` is cancelled as out of scope.

On `2026-04-06`, the user redirected jobs work away from the dashboard-first Windows jobs surface in this task and toward a separate backend control-plane approach under [Task-0005](/c:/Agent/CodexDashboard/Tracking/Task-0005/TASK.md).

This task did not reach an honest closure bar. The live `Jobs` regression readback remained unfinished, and the product direction changed before that closure work was completed.

## Cancellation Rationale

- The task explored a useful first-pass dashboard Jobs surface.
- It also proved some worthwhile seams around declared jobs, explicit refresh, and bounded reconciliation UI.
- But the durable next move is now a Go plus Temporal backend with dashboard integration, not more iteration on this dashboard-first implementation path.

## Preserved Baseline

The current repo still contains useful reference material from this task, including:

- managed declared-jobs state under `C:\Users\gregs\.codex\Orchestration\Jobs\`
- `Jobs` tab clicks that only switch surfaces
- explicit `Refresh` and `Force Reconcile` actions
- hidden child PowerShell windows for explicit jobs actions
- a scrollable `Jobs` content area and supporting UI smoke evidence

Treat those artifacts as reference material for future work, not as the required completion path now.

## Superseded Next Step

The previous live readback instruction for `REG-002` is superseded for this cancelled task.

Do not reopen `Task-0004` just to finish the old dashboard-first route unless the product direction explicitly swings back to it.

Any further jobs or orchestration work should continue under [Task-0005](/c:/Agent/CodexDashboard/Tracking/Task-0005/TASK.md).

## Watchouts

- treat `Task-0004` artifacts as reference material, not as the new architecture source of truth
- keep scheduler logic out of the dashboard process
- keep Git desired state separate from runtime truth
- update repo-root [REGRESSION.md](/c:/Agent/CodexDashboard/REGRESSION.md) when the new backend-backed job surface materially changes the real UI

## References

- [TASK.md](/c:/Agent/CodexDashboard/Tracking/Task-0004/TASK.md)
- [PLAN.md](/c:/Agent/CodexDashboard/Tracking/Task-0004/PLAN.md)
- [BUG-0001.md](/c:/Agent/CodexDashboard/Tracking/Task-0004/BUG-0001.md)
- [PASS-0002-AUDIT.md](/c:/Agent/CodexDashboard/Tracking/Task-0004/Testing/PASS-0002-AUDIT.md)
- [Task-0005 TASK.md](/c:/Agent/CodexDashboard/Tracking/Task-0005/TASK.md)
