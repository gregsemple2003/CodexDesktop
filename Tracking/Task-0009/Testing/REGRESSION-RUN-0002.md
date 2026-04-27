# REGRESSION-RUN-0002

Date: 2026-04-27

Scope:

- [BUG-0002](../BUG-0002.md)
- [REGRESSION.md](../../../REGRESSION.md) `SMOKE-002`
- [REGRESSION.md](../../../REGRESSION.md) `SMOKE-003`
- visible `Tasks` tab proof from pinned frontend release

## Validation Baseline

Commands:

```powershell
python -m unittest discover -s tests -p "test_*.py" -v
cd backend\orchestration
go test ./...
```

Results:

- Python unit discovery passed: 97 tests.
- Backend Go tests passed: `go test ./...`.

## Backend Release Proof

Published service-lane release:

- release id: `20260427T042647Z-055bdcd6a597`
- git commit: `055bdcd6a597cee5963a3384de012fbe63e5635c`
- source dirty: `false`
- binary path:
  `%LOCALAPPDATA%\CodexDashboard\orchestration-service-lane\releases\20260427T042647Z-055bdcd6a597\bin\controlplane-service-lane.exe`

Commands:

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File .\backend\orchestration\scripts\Publish-ServiceLaneRelease.ps1
powershell -NoProfile -ExecutionPolicy Bypass -File .\backend\orchestration\scripts\Stop-ServiceLane.ps1
powershell -NoProfile -ExecutionPolicy Bypass -File .\backend\orchestration\scripts\Install-ServiceLane.ps1
powershell -NoProfile -ExecutionPolicy Bypass -File .\backend\orchestration\scripts\Test-ServiceLaneIsolation.ps1
powershell -NoProfile -ExecutionPolicy Bypass -File .\backend\orchestration\scripts\Get-ServiceLaneStatus.ps1
```

Observed proof:

- scheduled task uses pinned runtime launcher: `true`
- current release error: `null`
- process running: `true`
- process path matches the pinned release binary
- backend health status: `ok`
- `/api/v1/tasks` returned task readback from `C:\Agent\CodexDashboard\Tracking`

## Frontend Release Proof

Published dashboard frontend release:

- release id: `20260427T042724Z-055bdcd6a597`
- git commit: `055bdcd6a597cee5963a3384de012fbe63e5635c`
- source mode: `git_commit`
- source dirty: `false`
- repository dirty: `true`
- release root:
  `%LOCALAPPDATA%\CodexDashboard\dashboard-releases\20260427T042724Z-055bdcd6a597`

The repository was dirty because unrelated working-tree files were present,
including deleted Task-0009 mockup files and untracked token-time artifacts. The
frontend release was still copied from the committed Git tree, not from those
working-tree changes.

Commands:

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File .\scripts\Publish-DashboardRelease.ps1
powershell -NoProfile -ExecutionPolicy Bypass -File .\scripts\Start-DashboardRelease.ps1
powershell -NoProfile -ExecutionPolicy Bypass -File .\scripts\Test-DashboardRelease.ps1
```

Observed proof:

- current release error: `null`
- startup uses pinned launcher: `true`
- launcher exists: `true`
- running dashboard process count: `1`
- running process command includes:
  - `--release-id 20260427T042724Z-055bdcd6a597`
  - `--release-root %LOCALAPPDATA%\CodexDashboard\dashboard-releases\20260427T042724Z-055bdcd6a597`

## Visible Tasks-Tab Smoke

Artifact directory:

- [PASS-0005-HUMAN-LANE-SMOKE-0001](./PASS-0005-HUMAN-LANE-SMOKE-0001/)

The smoke used the pinned frontend release and task-owned config/SQLite data,
with task readback pointed at the human service-lane backend:

- tasks backend: `http://127.0.0.1:4318`
- config path:
  `Tracking/Task-0009/Testing/Runtime/PASS-0005-HUMAN-LANE-SMOKE-0001/config.json`
- artifact files:
  - [overlay-summary.txt](./PASS-0005-HUMAN-LANE-SMOKE-0001/overlay-summary.txt)
  - [overlay.png](./PASS-0005-HUMAN-LANE-SMOKE-0001/overlay.png)

Summary excerpt:

```text
active_tab=tasks
hotkey_triggered=True
overlay_fallback=False
tasks_backend=http://127.0.0.1:4318
tasks_needs_you=11
tasks_selected=Task-0001
```

The task-owned runtime config, SQLite DB, debug log, and generated fixture job
registry were deleted after the proof run. The committed proof artifacts are the
summary and screenshot.

## Result

Passing.

The human-facing dashboard frontend is now pinned to a release manifest and
runtime launcher. Backend-only proof is no longer used as proof that the visible
dashboard surface is current.
