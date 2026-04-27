[CmdletBinding()]
param()

$ErrorActionPreference = "Stop"
Set-StrictMode -Version Latest

. (Join-Path $PSScriptRoot "DashboardReleaseHelpers.ps1")

$config = Get-DashboardReleaseConfig
Test-DashboardReleaseManifest -ManifestPath $config.CurrentReleaseManifestPath | Out-Null
if (-not (Test-Path -LiteralPath $config.LauncherScriptPath)) {
    throw "Dashboard launcher is missing at $($config.LauncherScriptPath). Run Publish-DashboardRelease.ps1 first."
}

& $config.LauncherScriptPath
Start-Sleep -Seconds 2
& (Join-Path $PSScriptRoot "Test-DashboardRelease.ps1")
