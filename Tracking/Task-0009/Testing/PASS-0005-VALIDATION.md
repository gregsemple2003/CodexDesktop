# PASS-0005 Validation

Date: 2026-04-27

Scope:

- [BUG-0002](../BUG-0002.md)
- frontend dashboard release pinning
- backend service-lane worktree-root correction needed for honest `Tasks` readback
- backup inventory update for dashboard frontend release artifacts

## Focused Validation

Command:

```powershell
python -m unittest tests.test_dashboard_release_scripts tests.test_desktop_support tests.test_service_lane_scripts -v
```

Result: passing, 46 tests.

Coverage:

- dashboard release planning uses `%LOCALAPPDATA%`
- dashboard release publish writes a pinned manifest
- dashboard release publish uses `source_mode = git_commit` by default, so
  unrelated untracked working-tree files are not promoted into the frontend
  release
- startup registration points at the runtime-root dashboard launcher
- dashboard release hash validation catches tampered files
- startup command no longer launches the mutable repo checkout directly
- existing service-lane release isolation coverage still passes

## Script Parse Validation

Command:

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -Command '$files=@("scripts\DashboardReleaseHelpers.ps1","scripts\Publish-DashboardRelease.ps1","scripts\Start-CodexDashboard.ps1","scripts\Start-DashboardRelease.ps1","scripts\Test-DashboardRelease.ps1","backend\orchestration\scripts\LaneHelpers.ps1"); foreach ($f in $files){ [scriptblock]::Create((Get-Content -Raw -LiteralPath $f)) | Out-Null }; "parsed"'
```

Result: passing, output `parsed`.

## Full Unit Validation

Command:

```powershell
python -m unittest discover -s tests -p "test_*.py" -v
```

Result: passing, 96 tests.

## Backend Validation

Command:

```powershell
cd backend\orchestration
go test ./...
```

Result: passing.

## Backup Inventory Validation

Command:

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File .\backend\orchestration\scripts\Get-HumanLaneBackupPlan.ps1 -SkipDockerVolumeSize
```

Result: passing. The plan now includes:

- `dashboard-current-release`
- `dashboard-releases`
- `dashboard-launcher`

## Issues Found During Validation

- Windows PowerShell 5.1 does not expose `[System.IO.Path]::GetRelativePath`.
  The dashboard release helper now uses a local child-path resolver.
- `Test-Path` interpreted the Inter font filename brackets as wildcards. Release
  validation now uses literal-path checks for manifest files.

## Remaining Before Closure

- commit and push the checkpoint
- publish pinned backend and frontend releases from the committed checkpoint
- restart the human-facing backend/frontend only through their pinned launchers
- run release proof scripts and visible Tasks-tab smoke evidence
