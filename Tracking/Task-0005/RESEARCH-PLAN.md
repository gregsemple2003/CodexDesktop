# Task 0005 Research Plan

## Research Goal

Turn the backend-first Go plus Temporal direction in `Task-0005` into a planning-ready implementation path that fits the current Tk dashboard, the existing `.codex/Orchestration/Jobs` state, and the real toolchain available on this machine.

## Decision-Shaping Problems

### Problem 0001 Desired-State Schema And Migration Fit

Determine how the current `declared-jobs.json` Windows registry should evolve into a Git-friendly v1 job-spec layout for `schedule`, `manual`, and `webhook` triggers without keeping Windows durable state as the source of truth.

### Problem 0002 Runtime Topology And Temporal Boundary

Determine the smallest honest Go service shape for startup reconcile, explicit sync, durable execution, and `codex exec` integration while keeping Temporal as runtime truth and Git as desired state.

### Problem 0003 Dashboard Integration And Proof Strategy

Determine how the Tk `Jobs` surface should change from local Windows reconciliation to backend readback and bounded controls, and what task-level proof is required for the new backend-backed surface.

## Research Inputs

- `Tracking/Task-0005/TASK.md`
- `Tracking/Task-0005/HANDOFF.md`
- `Tracking/Task-0005/Research/Cron System with JSON.md`
- `Tracking/Task-0004/TASK.md`
- `Tracking/Task-0004/HANDOFF.md`
- `Design/GENERAL-DESIGN.md`
- `app/codex_dashboard/jobs.py`
- `app/codex_dashboard/paths.py`
- `app/codex_dashboard/ui.py`
- `tests/test_jobs.py`
- `C:\Users\gregs\.codex\Orchestration\Jobs\declared-jobs.json`
- `C:\Users\gregs\.codex\Orchestration\Jobs\declared-jobs.schema.json`
- repo-root `REGRESSION.md`
- repo-root `TESTING.md`
- local `codex --help`
- local `codex exec --help`
- official Temporal docs:
  - `https://docs.temporal.io/develop/go`
  - `https://docs.temporal.io/develop/go/schedules`
  - `https://docs.temporal.io/cli/schedule`
  - `https://raw.githubusercontent.com/temporalio/docker-compose/refs/heads/main/docker-compose-postgres.yml`

## Intended Outputs

- `RESEARCH-ANALYSIS.md`
- `RESEARCH.md`
- a planning-ready `PLAN.md`

## Exit Bar

- the desired-state layout under `.codex/Orchestration/Jobs` is concrete enough to implement without another architecture reset
- the backend responsibility split between Git, the Go service, Temporal, and the dashboard is explicit
- the pass plan names a realistic path to `schedule`, `manual`, and `webhook` support plus one real `codex exec` executor
- the proof plan distinguishes pass-local unit coverage from the required real app-surface regression lane
- local environment prerequisites are preserved honestly instead of being discovered halfway through implementation
