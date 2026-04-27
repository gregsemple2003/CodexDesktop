[CmdletBinding()]
param()

$ErrorActionPreference = "Stop"
Set-StrictMode -Version Latest

. (Join-Path $PSScriptRoot "DashboardReleaseHelpers.ps1")

$config = Get-DashboardReleaseConfig
$manifest = $null
$manifestError = $null
try {
    $manifest = Test-DashboardReleaseManifest -ManifestPath $config.CurrentReleaseManifestPath
}
catch {
    $manifestError = $_.Exception.Message
}

$startupText = $null
if (Test-Path -LiteralPath $config.StartupPath) {
    $startupText = Get-Content -Raw -LiteralPath $config.StartupPath
}

$processes = @()
if ($null -ne $manifest) {
    $releaseId = [string]$manifest.release_id
    $releaseRoot = [string]$manifest.release_root
    $processes = @(
        Get-CimInstance Win32_Process |
            Where-Object {
                $_.Name -eq "pythonw.exe" -and
                $_.CommandLine -like "*app.codex_dashboard*" -and
                $_.CommandLine -like "*--release-id*" -and
                $_.CommandLine -like "*$releaseId*" -and
                $_.CommandLine -like "*$releaseRoot*"
            } |
            Select-Object ProcessId, Name, ExecutablePath, CommandLine
    )
}

$releaseSummary = $null
if ($null -ne $manifest) {
    $releaseSummary = [ordered]@{
        release_id = [string]$manifest.release_id
        git_commit = [string]$manifest.git_commit
        source_mode = [string]$manifest.source_mode
        source_dirty = [bool]$manifest.source_dirty
        repository_dirty = [bool]$manifest.repository_dirty
        release_root = [string]$manifest.release_root
        pythonw_path = [string]$manifest.pythonw_path
        file_count = @($manifest.files).Count
    }
}

[pscustomobject]@{
    current_release_manifest_path = $config.CurrentReleaseManifestPath
    current_release_error = $manifestError
    current_release = $releaseSummary
    launcher_script_path = $config.LauncherScriptPath
    launcher_exists = (Test-Path -LiteralPath $config.LauncherScriptPath)
    startup_path = $config.StartupPath
    startup_uses_pinned_launcher = (
        $null -ne $startupText -and
        $startupText -like "*$($config.LauncherScriptPath)*" -and
        $startupText -notlike "*C:\\Agent\\CodexDashboard*"
    )
    running_process_count = $processes.Count
    running_processes = $processes
} | ConvertTo-Json -Depth 8
