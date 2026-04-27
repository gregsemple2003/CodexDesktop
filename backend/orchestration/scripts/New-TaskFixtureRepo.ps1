[CmdletBinding()]
param(
    [string]$Destination
)

Set-StrictMode -Version Latest
$ErrorActionPreference = "Stop"

if (-not $Destination) {
    $Destination = Join-Path ([System.IO.Path]::GetTempPath()) ("CodexDashboardTaskFixtureRepo-" + [guid]::NewGuid().ToString("n"))
}

if (Test-Path -LiteralPath $Destination) {
    throw "Destination already exists: $Destination. Use a new path or delete the old disposable fixture repo first."
}

New-Item -ItemType Directory -Path $Destination -Force | Out-Null

function Write-FixtureFile {
    param(
        [Parameter(Mandatory=$true)][string]$RelativePath,
        [Parameter(Mandatory=$true)][string]$Content
    )

    $path = Join-Path $Destination $RelativePath
    $parent = Split-Path -Parent $path
    New-Item -ItemType Directory -Path $parent -Force | Out-Null
    Set-Content -Path $path -Value $Content -Encoding UTF8
}

Write-FixtureFile -RelativePath "Tracking/Task-0008/TASK.md" -Content @'
# Task 0008

## Title

Build the backend task dispatch layer.

## Summary

Create the durable backend task-run contract so later clients do not guess state.
'@

Write-FixtureFile -RelativePath "Tracking/Task-0008/TASK-STATE.json" -Content @'
{
  "task_id": "Task-0008",
  "status": "in_progress",
  "phase": "implementation",
  "plan_approved": true,
  "current_pass": "PASS-0001",
  "current_gate": "implementation",
  "blockers": [],
  "updated_at": "2026-04-24T22:00:00-04:00"
}
'@

Write-FixtureFile -RelativePath "Tracking/Task-0008/PLAN.md" -Content @'
# Task-0008 Plan

Build backend readback, dispatch, pause, owned-lane cleanup, and durable run state.
'@

Write-FixtureFile -RelativePath "Tracking/Task-0008/HANDOFF.md" -Content @'
# Task-0008 Handoff

## Current Status

Backend runtime semantics are available for task dispatch and supervision tests.
'@

Write-FixtureFile -RelativePath "Tracking/Task-0009/TASK.md" -Content @'
# Task 0009

## Title

Build the committed-work Tasks tab.

## Summary

Show authored and promoted committed tasks, backend run state, Pause controls, and launch targets without surfacing unpromoted candidates.
'@

Write-FixtureFile -RelativePath "Tracking/Task-0009/TASK-STATE.json" -Content @'
{
  "task_id": "Task-0009",
  "status": "completed",
  "phase": "closure",
  "plan_approved": true,
  "current_pass": "PASS-0004",
  "current_gate": "complete",
  "blockers": [],
  "updated_at": "2026-04-26T21:35:00-04:00"
}
'@

Write-FixtureFile -RelativePath "Tracking/Task-0009/PLAN.md" -Content @'
# Task-0009 Plan

Keep the Tasks tab true to committed work: no unpromoted candidates, no Candidate provenance labels, visible Pause copy, and no false AI-run progress bar.
'@

Write-FixtureFile -RelativePath "Tracking/Task-0009/HANDOFF.md" -Content @'
# Task-0009 Handoff

## Current Status

The Tasks tab consumes backend-shaped task readback and renders the committed-work surface.
'@

git -C $Destination init | Out-Null
git -C $Destination config user.email "taskrepo-fixture@example.com" | Out-Null
git -C $Destination config user.name "Task Repo Fixture" | Out-Null
git -C $Destination add Tracking | Out-Null
git -C $Destination commit -m "fixture: initial task 8 and 9 state" | Out-Null
$initialCommit = (git -C $Destination rev-parse HEAD).Trim()

Write-FixtureFile -RelativePath "Tracking/Task-0008/TASK-STATE.json" -Content @'
{
  "task_id": "Task-0008",
  "status": "in_progress",
  "phase": "implementation",
  "plan_approved": true,
  "current_pass": "PASS-0002",
  "current_gate": "implementation",
  "blockers": [],
  "updated_at": "2026-04-24T23:00:00-04:00"
}
'@

git -C $Destination add Tracking/Task-0008/TASK-STATE.json | Out-Null
git -C $Destination commit -m "fixture: advance task 8 state" | Out-Null
$advancedCommit = (git -C $Destination rev-parse HEAD).Trim()

[ordered]@{
    root = (Resolve-Path -LiteralPath $Destination).Path
    initial_commit = $initialCommit
    advanced_commit = $advancedCommit
    cleanup = "Remove-Item -Recurse -Force -LiteralPath '$((Resolve-Path -LiteralPath $Destination).Path)'"
    note = "Disposable fixture repo. Delete it after smoke testing; do not back up generated fixture repos as durable state."
} | ConvertTo-Json -Depth 4
