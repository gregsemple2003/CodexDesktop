[CmdletBinding()]
param(
    [switch]$AllowDirty,
    [switch]$NoPin,
    [switch]$NoStartup,
    [switch]$PlanOnly
)

$ErrorActionPreference = "Stop"
Set-StrictMode -Version Latest

. (Join-Path $PSScriptRoot "DashboardReleaseHelpers.ps1")

$config = Get-DashboardReleaseConfig

if ($PlanOnly) {
    [pscustomobject]@{
        source_repo_root = $config.RepoRoot
        source_app_root = $config.SourceAppRoot
        releases_root = $config.ReleasesRoot
        current_release_manifest_path = $config.CurrentReleaseManifestPath
        launcher_script_path = $config.LauncherScriptPath
        startup_path = $config.StartupPath
        would_pin_current = (-not $NoPin)
        would_install_startup = (-not $NoStartup)
    } | ConvertTo-Json -Depth 8
    exit 0
}

Install-DashboardLauncher -Config $config
$release = New-DashboardRelease -Config $config -AllowDirty:$AllowDirty -PinCurrent:(-not $NoPin)
if (-not $NoStartup) {
    Install-DashboardStartup -Config $config
}
$release | ConvertTo-Json -Depth 12
