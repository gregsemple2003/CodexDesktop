[CmdletBinding()]
param()

$ErrorActionPreference = "Stop"
Set-StrictMode -Version Latest

. (Join-Path $PSScriptRoot "LaneHelpers.ps1")

$config = Get-OrchestrationLaneConfig -Lane "service"

$release = $null
$releaseError = $null
try {
    $release = Get-ServiceLaneCurrentRelease -Config $config
}
catch {
    $releaseError = $_.Exception.Message
}

$releaseSummary = $null
if ($null -ne $release) {
    $releaseSummary = [ordered]@{
        release_id = [string]$release.release_id
        git_commit = [string]$release.git_commit
        source_dirty = [bool]$release.source_dirty
        binary_path = [string]$release.binary_path
        binary_sha256 = [string]$release.binary_sha256
        compose_file_path = [string]$release.compose_file_path
        compose_file_sha256 = [string]$release.compose_file_sha256
    }
}

[pscustomobject]@{
    lane = $config.Lane
    task_name = $config.TaskName
    scheduled_task_uses_pinned_launcher = Test-ServiceLaneTaskUsesPinnedLauncher -Config $config
    runner_script_path = $config.RunnerScriptPath
    current_release_manifest_path = $config.CurrentReleaseManifestPath
    current_release_error = $releaseError
    current_release = $releaseSummary
} | ConvertTo-Json -Depth 8
