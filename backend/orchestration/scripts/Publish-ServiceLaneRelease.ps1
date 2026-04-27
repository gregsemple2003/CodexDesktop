[CmdletBinding()]
param(
    [switch]$AllowDirty,
    [switch]$NoPin,
    [switch]$PlanOnly
)

$ErrorActionPreference = "Stop"
Set-StrictMode -Version Latest

. (Join-Path $PSScriptRoot "LaneHelpers.ps1")

$config = Get-OrchestrationLaneConfig -Lane "service"

if ($PlanOnly) {
    [pscustomobject]@{
        lane = $config.Lane
        source_repo_root = $config.RepoRoot
        releases_root = $config.ReleasesRoot
        current_release_manifest_path = $config.CurrentReleaseManifestPath
        runner_script_path = $config.RunnerScriptPath
        source_compose_file = $config.SourceComposeFile
        would_pin_current = (-not $NoPin)
    } | ConvertTo-Json -Depth 8
    exit 0
}

Install-ServiceLaneLauncher -Config $config
$release = New-ServiceLaneRelease -Config $config -AllowDirty:$AllowDirty -PinCurrent:(-not $NoPin)
$release | ConvertTo-Json -Depth 8
