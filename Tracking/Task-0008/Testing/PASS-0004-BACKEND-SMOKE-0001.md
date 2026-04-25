# PASS-0004 Backend Smoke 0001

## Type

Unit test / server-only smoke / closeout validation.

## Date

`2026-04-24`

## Scope

Closeout proof that the approved Task-0008 backend/runtime scope is now satisfied:

- durable dispatch and task-run persistence
- supervision, poke, interrupt, cleanup, and recovery behavior
- exclusive owned-lane execution plus restore-baseline semantics
- deep-context readback for operators and later clients
- declared-doc drift reconcile with preserved runtime truth

## Commands

```powershell
cd C:\Agent\CodexDashboard\backend\orchestration
go test ./...
```

## Result

Pass.

The full backend suite passed after the declared-doc drift slice landed.

Task-owned proof now covers the approved contract end to end:

- dispatch and durable task/run readback in [PASS-0000-BACKEND-SMOKE-0001.md](./PASS-0000-BACKEND-SMOKE-0001.md), [PASS-0001-BACKEND-SMOKE-0001.md](./PASS-0001-BACKEND-SMOKE-0001.md), and [PASS-0001-BACKEND-SMOKE-0002.md](./PASS-0001-BACKEND-SMOKE-0002.md)
- supervision, poke, interrupt, cleanup, owned-lane execution, workload recovery, and bounded natural failure proof across [PASS-0002-BACKEND-SMOKE-0001.md](./PASS-0002-BACKEND-SMOKE-0001.md) through [PASS-0002-BACKEND-SMOKE-0020.md](./PASS-0002-BACKEND-SMOKE-0020.md)
- deep-context operator readback in [PASS-0003-BACKEND-SMOKE-0001.md](./PASS-0003-BACKEND-SMOKE-0001.md)
- declared-doc drift and runtime-divergence reconcile proof in [PASS-0003-BACKEND-SMOKE-0002.md](./PASS-0003-BACKEND-SMOKE-0002.md)

## Regression Applicability

Repo-root regression is honestly `not_applicable` for Task-0008.

Per [../../../REGRESSION.md](../../../REGRESSION.md), canonical regression in this repo starts from the real desktop app surface.

Task-0008 shipped backend-only orchestration and readback behavior under `backend/orchestration/` and did not change the Tk desktop overlay or other app-surface interactions. Backend smoke and unit proof are therefore the correct closeout lane for this task.

## Why This Counts

This closeout run satisfies the remaining PASS-0004 burden from [../PLAN.md](../PLAN.md):

- state transitions and supervision rules are covered by focused tests
- task-owned evidence proves dispatch, poke, interrupt, cleanup, retry, deep-context, and divergence behavior through direct backend interaction
- owned-lane reset semantics are concrete and recorded against useful commits
- backend readback, not client memory, owns state envelopes, attention, wait validity, action gating, and divergence reporting

## Residual Notes

- The default validation lane on `14318` can still be less trustworthy than a clean manual listener when proof must guarantee the newest binary and avoid stdout-log lock races.
- The declared-doc drift proof is bounded to live task-doc mutation plus read-through reconcile rather than a literal committed git rewind, but it exercises the same contract requirement: declared docs changed while runtime truth remained preserved and was reported honestly.
