[CmdletBinding()]
param(
    [string]$DestinationRoot = (Join-Path ([Environment]::GetFolderPath("MyDocuments")) "CodexDashboardBackups"),
    [ValidateSet("Dump", "SkipDatabase")]
    [string]$DatabaseMode = "Dump",
    [switch]$IncludeRepoBundle,
    [switch]$SkipDockerVolumeSize,
    [switch]$PlanOnly
)

$ErrorActionPreference = "Stop"
Set-StrictMode -Version Latest

. (Join-Path $PSScriptRoot "HumanLaneBackupHelpers.ps1")

$plan = Get-HumanLaneBackupPlan -SkipDockerVolumeSize:$SkipDockerVolumeSize

if ($PlanOnly) {
    $plan | ConvertTo-Json -Depth 12
    return
}

$timestamp = Get-Date -Format "yyyyMMdd-HHmmss"
$backupRoot = Join-Path $DestinationRoot "human-lane-$timestamp"
$archivesRoot = Join-Path $backupRoot "archives"
$databaseRoot = Join-Path $backupRoot "database"
$repoStateRoot = Join-Path $backupRoot "repo-state"

New-Item -ItemType Directory -Path $archivesRoot -Force | Out-Null
New-Item -ItemType Directory -Path $databaseRoot -Force | Out-Null
New-Item -ItemType Directory -Path $repoStateRoot -Force | Out-Null

$serviceConfig = Get-OrchestrationLaneConfig -Lane "service"
$artifacts = @()

foreach ($entry in @($plan.inventory.paths)) {
    $zipPath = Join-Path $archivesRoot "$($entry.id).zip"
    $archiveResult = Compress-ExistingPath -SourcePath $entry.path -DestinationPath $zipPath
    $artifacts += [ordered]@{
        id = $entry.id
        source = $entry.path
        artifact = if ($archiveResult.created) { $zipPath } else { $null }
        created = [bool]$archiveResult.created
        skipped_count = @($archiveResult.skipped).Count
        skipped = @($archiveResult.skipped)
        error = $archiveResult.error
    }
}

$scheduledTaskPath = Join-Path $backupRoot "service-lane-scheduled-task.xml"
try {
    Export-ScheduledTask -TaskName $serviceConfig.TaskName | Set-Content -LiteralPath $scheduledTaskPath -Encoding UTF8
    $artifacts += [ordered]@{
        id = "service-lane-scheduled-task"
        source = $serviceConfig.TaskName
        artifact = $scheduledTaskPath
        created = $true
    }
}
catch {
    $artifacts += [ordered]@{
        id = "service-lane-scheduled-task"
        source = $serviceConfig.TaskName
        artifact = $null
        created = $false
        error = $_.Exception.Message
    }
}

$databaseArtifact = $null
if ($DatabaseMode -eq "Dump") {
    $databaseArtifact = Join-Path $databaseRoot "temporal-pg-dumpall.sql"
    Invoke-ServiceLanePgDumpAll -Config $serviceConfig -DestinationPath $databaseArtifact
}

Save-RepoDelta -RepoRoot $plan.inventory.repo.root -DestinationRoot $repoStateRoot -IncludeRepoBundle:$IncludeRepoBundle

$manifest = [ordered]@{
    created_at = (Get-Date).ToString("o")
    backup_root = $backupRoot
    database_mode = $DatabaseMode
    database_artifact = $databaseArtifact
    include_repo_bundle = [bool]$IncludeRepoBundle
    source_plan = $plan
    artifacts = $artifacts
}

$manifestPath = Join-Path $backupRoot "manifest.json"
$manifest | ConvertTo-Json -Depth 14 | Set-Content -LiteralPath $manifestPath -Encoding UTF8

[ordered]@{
    backup_root = $backupRoot
    manifest = $manifestPath
    database_mode = $DatabaseMode
    database_artifact = $databaseArtifact
    artifact_count = @($artifacts | Where-Object { $_.created }).Count
    repo_bundle_included = [bool]$IncludeRepoBundle
} | ConvertTo-Json -Depth 8
