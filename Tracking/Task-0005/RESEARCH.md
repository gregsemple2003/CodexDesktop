# Task 0005 Research Summary

## Research Verdict

Task research is planning-ready.

## Key Decisions

- keep `C:\Users\gregs\.codex\Orchestration\Jobs\` as the durable desired-state root, but replace the Windows-shaped `declared-jobs.json` model with one JSON job spec per file plus a shared schema
- add a new Go service area at `backend/orchestration/` and let that service own validation, startup reconcile, explicit sync, manual trigger, webhook intake, and the Temporal worker
- keep Git as desired state and Temporal as runtime truth; do not add scheduler ownership back into the Tk dashboard
- use Temporal schedules for `schedule` triggers and route `manual` plus `webhook` into the same durable run workflow through backend APIs
- keep `codex exec` as the first real executor path and run it only from Activities
- reuse the existing `Jobs` tab as a bounded correctness surface that reads backend state and invokes `Sync now` or `Run now`

## Why This Is The Honest Repo Fit

- the repo already has a durable jobs home under `.codex/Orchestration/Jobs`
- the Tk app already has a usable `Jobs` surface that can become a backend client
- the current Python jobs module is useful reference material but should no longer own runtime reconciliation against Windows objects
- official Temporal schedule features cover the task's required trigger semantics without inventing a second scheduler DSL

## Carry-Forward Constraints

- keep the dashboard hotkey-first and bounded
- keep Git desired state separate from runtime truth
- do not widen v1 into a general agent graph or authoring console
- keep `schedule`, `manual`, and `webhook` as the only required trigger types for closure
- keep `codex exec` as the first required real executor path
- update repo-root `REGRESSION.md` once the backend-backed Jobs surface is concrete

## Environment Watchout

Local verification is not tooling-ready yet:

- `codex` and `codex exec` are available
- `go`, `docker`, and `temporal` are not currently on `PATH`

`PASS-0000` must resolve or document that toolchain gap before the task can honestly claim a runnable Go plus Temporal dev stack.

## Planning Recommendation

Plan the work in four implementation passes:

- `PASS-0000`: job-spec layout, repo-local Temporal dev stack docs, backend scaffold, and environment bootstrap
- `PASS-0001`: desired-state compile plus Temporal reconcile, sync, read APIs, and health
- `PASS-0002`: durable run workflow plus `schedule`, `manual`, `webhook`, and `codex exec` execution proof
- `PASS-0003`: Tk dashboard backend integration, regression-matrix update, and closure preparation

This does not remove the separate task-level regression phase after the planned passes.
