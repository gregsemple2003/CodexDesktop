# Problem 0005 Option C Plan 0001

## Planning Intent

This file turns Problem `0005`, Option `C. Background watcher or scheduler` from [ORTHOGONAL-SOLUTIONS-MATRIX.md](../ORTHOGONAL-SOLUTIONS-MATRIX.md) into a bounded implementation sequence.

It is an alternative route, not the selected winner task.

## Summary

Build a backend-controlled watchdog job that detects stalled owned tasks, writes a resume packet, and surfaces the stall in CodexDashboard before the human has to rediscover it manually.

## Fixed Defaults

- scope: orchestration job spec plus CodexDashboard backend and UI
- canonical homes:
  - `C:\Users\gregs\.codex\Orchestration\Jobs\specs\codex-task-watchdog-v1.json`
  - `C:\Agent\CodexDashboard\app\codex_dashboard\job_specs.py`
  - `C:\Agent\CodexDashboard\app\codex_dashboard\jobs_backend.py`
  - `C:\Agent\CodexDashboard\app\codex_dashboard\jobs.py`
  - `C:\Agent\CodexDashboard\app\codex_dashboard\ui.py`
- watched task-state conditions:
  - `task_owner_mode=system_owns_until_block`
  - `continuation_state=continue_autonomously`
  - `last_progress_update_at` older than threshold
  - no named blocker and no approval gate
- first rollout writes a resume packet and raises a visible alert; it does not auto-launch a fresh agent

## Pass Plan

### Pass 0000 - Watchdog Contract And Job Spec

Goal:

- define what counts as a stalled owned task and how the watchdog reports it

Build:

- add `codex-task-watchdog-v1.json`
- define the stale-task predicate and resume-packet payload
- document the operator-visible outcomes for `stalled`, `resumable`, and `hard_blocked`

Unit Proof:

- the job spec names concrete stale-state inputs
- the resume-packet payload is specific enough to render in the dashboard

Exit Bar:

- the watchdog has a bounded contract instead of vague polling behavior

### Pass 0001 - Backend Detection And Packet Writing

Goal:

- make the watchdog runnable against live task state

Build:

- update `job_specs.py`, `jobs_backend.py`, and `jobs.py` to register and run the watchdog
- detect stalled tasks from `TASK-STATE.json`
- write one task-owned resume packet when a task becomes stalled

Unit Proof:

- backend tests or focused smoke slices cover stale, resumable, and blocked cases
- resume packets are not written for tasks that are honestly waiting on approval or a real blocker

Exit Bar:

- the system can surface restart-risk before the human has to wake the task up manually

### Pass 0002 - Dashboard Visibility

Goal:

- make the stalled-task signal visible in the product surface

Build:

- add a watchdog status row or panel in `ui.py`
- show task id, stale reason, last progress time, and resume-packet link
- verify the job appears through the existing jobs backend path

Unit Proof:

- the jobs surface shows the watchdog job and one stale-task example honestly
- the resume-packet link opens the right task artifact

Exit Bar:

- restart supervision becomes an explicit monitored condition instead of a silent human rediscovery

## Testing Strategy

- keep stale-task detection rules aligned with the task-state fields they read
- treat false positives on honest approval gates as rollout failures

## Deferred Work

- autonomous agent relaunch
- hidden retry loops outside the visible jobs path
- non-dashboard notification channels beyond the existing product surfaces
