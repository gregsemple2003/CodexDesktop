Set-StrictMode -Version Latest

. (Join-Path $PSScriptRoot "LaneHelpers.ps1")

function Get-CodexDashboardRepoRoot {
    $current = Get-OrchestrationRepoRoot
    while ($true) {
        if ((Test-Path -LiteralPath (Join-Path $current "Tracking")) -and
            (Test-Path -LiteralPath (Join-Path $current "AGENTS.md"))) {
            return $current
        }

        $parent = Split-Path -Parent $current
        if ($parent -eq $current -or [string]::IsNullOrWhiteSpace($parent)) {
            return (Get-OrchestrationRepoRoot)
        }
        $current = $parent
    }
}

function Get-HumanLaneBackupInventory {
    $serviceConfig = Get-OrchestrationLaneConfig -Lane "service"
    $repoRoot = Get-CodexDashboardRepoRoot
    $localAppData = Get-OrchestrationLocalAppData
    $dashboardRoot = Join-Path $localAppData "CodexDashboard"
    $jobsSpecsRoot = Join-Path $HOME ".codex\Orchestration\Jobs\specs"
    $volumeName = "$($serviceConfig.ComposeProject)_temporal-postgres-data"

    return [ordered]@{
        generated_at = (Get-Date).ToString("o")
        lane = "human"
        service = [ordered]@{
            bind_address = $serviceConfig.BindAddress
            backend_url = $serviceConfig.JobsBackendUrl
            temporal_address = $serviceConfig.TemporalAddress
            postgres_port = $serviceConfig.PostgresPort
            task_name = $serviceConfig.TaskName
            compose_project = $serviceConfig.ComposeProject
            temporal_postgres_volume = $volumeName
        }
        paths = @(
            [ordered]@{ id = "service-runtime-root"; backup_class = "must"; path = $serviceConfig.RuntimeRoot; kind = "directory" },
            [ordered]@{ id = "service-runs-root"; backup_class = "must"; path = $serviceConfig.RunsRoot; kind = "directory" },
            [ordered]@{ id = "dashboard-current-release"; backup_class = "must"; path = (Join-Path $dashboardRoot "dashboard-current-release.json"); kind = "file" },
            [ordered]@{ id = "dashboard-releases"; backup_class = "must"; path = (Join-Path $dashboardRoot "dashboard-releases"); kind = "directory" },
            [ordered]@{ id = "dashboard-launcher"; backup_class = "must"; path = (Join-Path $dashboardRoot "dashboard-launcher"); kind = "directory" },
            [ordered]@{ id = "dashboard-config"; backup_class = "must"; path = (Join-Path $dashboardRoot "config.json"); kind = "file" },
            [ordered]@{ id = "dashboard-db"; backup_class = "must"; path = (Join-Path $dashboardRoot "dashboard.db"); kind = "file" },
            [ordered]@{ id = "jobs-specs"; backup_class = "must"; path = $jobsSpecsRoot; kind = "directory" }
        )
        repo = [ordered]@{
            root = $repoRoot
            backup_class = "must-backup-delta"
        }
    }
}

function Get-PathSizeBytes {
    param(
        [Parameter(Mandatory = $true)]
        [string]$Path
    )

    if (-not (Test-Path -LiteralPath $Path)) {
        return 0L
    }

    $item = Get-Item -LiteralPath $Path -Force
    if (-not $item.PSIsContainer) {
        return [int64]$item.Length
    }

    $sum = (Get-ChildItem -LiteralPath $Path -Recurse -Force -File -ErrorAction SilentlyContinue |
        Measure-Object -Property Length -Sum).Sum
    if ($null -eq $sum) {
        return 0L
    }
    return [int64]$sum
}

function Get-DockerVolumeSizeBytes {
    param(
        [Parameter(Mandatory = $true)]
        [string]$VolumeName,
        [Parameter(Mandatory = $true)]
        [hashtable]$Config
    )

    $dockerPath = Ensure-DockerEngine -Config $Config
    $output = & $dockerPath run --rm -v "$VolumeName`:/volume:ro" postgres:16-alpine sh -c "du -sk /volume"
    if ($LASTEXITCODE -ne 0) {
        throw "Could not measure Docker volume $VolumeName."
    }
    $kbText = (($output | Select-Object -First 1) -split "\s+")[0]
    return ([int64]$kbText) * 1024L
}

function Get-RepoStateSummary {
    param(
        [Parameter(Mandatory = $true)]
        [string]$RepoRoot
    )

    $branch = Invoke-GitCapture -RepoRoot $RepoRoot -Arguments @("branch", "--show-current") -AllowFailure
    $head = Invoke-GitCapture -RepoRoot $RepoRoot -Arguments @("rev-parse", "HEAD") -AllowFailure
    $remote = Invoke-GitCapture -RepoRoot $RepoRoot -Arguments @("remote", "get-url", "origin") -AllowFailure
    $status = @((Invoke-GitCapture -RepoRoot $RepoRoot -Arguments @("status", "--short", "--untracked-files=all") -AllowFailure) -split "`n" | Where-Object { $_ })
    $untracked = @((Invoke-GitCapture -RepoRoot $RepoRoot -Arguments @("ls-files", "--others", "--exclude-standard") -AllowFailure) -split "`n" | Where-Object { $_ })
    $modified = @((Invoke-GitCapture -RepoRoot $RepoRoot -Arguments @("diff", "--name-only") -AllowFailure) -split "`n" | Where-Object { $_ })
    $staged = @((Invoke-GitCapture -RepoRoot $RepoRoot -Arguments @("diff", "--cached", "--name-only") -AllowFailure) -split "`n" | Where-Object { $_ })
    $upstream = Invoke-GitCapture -RepoRoot $RepoRoot -Arguments @("rev-parse", "--abbrev-ref", "--symbolic-full-name", "@{u}") -AllowFailure
    $aheadBehind = ""
    if (-not [string]::IsNullOrWhiteSpace($upstream)) {
        $aheadBehind = Invoke-GitCapture -RepoRoot $RepoRoot -Arguments @("rev-list", "--left-right", "--count", "$upstream...HEAD") -AllowFailure
    }

    return [ordered]@{
        root = $RepoRoot
        branch = [string]$branch
        head = [string]$head
        origin = [string]$remote
        upstream = [string]$upstream
        ahead_behind = [string]$aheadBehind
        dirty_entry_count = @($status).Count
        untracked_count = @($untracked).Count
        modified_count = @($modified).Count
        staged_count = @($staged).Count
    }
}

function Invoke-GitCapture {
    param(
        [Parameter(Mandatory = $true)]
        [string]$RepoRoot,
        [Parameter(Mandatory = $true)]
        [string[]]$Arguments,
        [switch]$AllowFailure
    )

    $oldErrorActionPreference = $ErrorActionPreference
    $ErrorActionPreference = "Continue"
    try {
        $output = @(& git -C $RepoRoot @Arguments 2>$null)
        $exitCode = $LASTEXITCODE
    }
    finally {
        $ErrorActionPreference = $oldErrorActionPreference
    }

    if ($exitCode -ne 0 -and -not $AllowFailure) {
        throw "git $($Arguments -join ' ') failed with exit code $exitCode."
    }
    if ($exitCode -ne 0) {
        return ""
    }
    return ($output -join "`n")
}

function Get-HumanLaneBackupPlan {
    param(
        [switch]$SkipDockerVolumeSize
    )

    $inventory = Get-HumanLaneBackupInventory
    $serviceConfig = Get-OrchestrationLaneConfig -Lane "service"
    $items = @()
    $totalBytes = 0L

    foreach ($entry in @($inventory.paths)) {
        $bytes = Get-PathSizeBytes -Path $entry.path
        $totalBytes += $bytes
        $items += [ordered]@{
            id = $entry.id
            backup_class = $entry.backup_class
            path = $entry.path
            kind = $entry.kind
            exists = (Test-Path -LiteralPath $entry.path)
            bytes = $bytes
            mb = [math]::Round($bytes / 1MB, 2)
        }
    }

    $volumeBytes = $null
    $volumeError = $null
    if (-not $SkipDockerVolumeSize) {
        try {
            $volumeBytes = Get-DockerVolumeSizeBytes -VolumeName $inventory.service.temporal_postgres_volume -Config $serviceConfig
            $totalBytes += $volumeBytes
        }
        catch {
            $volumeError = $_.Exception.Message
        }
    }

    $repoSummary = Get-RepoStateSummary -RepoRoot $inventory.repo.root

    return [ordered]@{
        generated_at = (Get-Date).ToString("o")
        inventory = $inventory
        items = $items
        docker_volume = [ordered]@{
            id = "service-temporal-postgres-volume"
            backup_class = "must"
            name = $inventory.service.temporal_postgres_volume
            bytes = $volumeBytes
            mb = if ($null -ne $volumeBytes) { [math]::Round($volumeBytes / 1MB, 2) } else { $null }
            size_error = $volumeError
        }
        repo = $repoSummary
        estimated_known_bytes = $totalBytes
        estimated_known_mb = [math]::Round($totalBytes / 1MB, 2)
    }
}

function Write-TextFile {
    param(
        [Parameter(Mandatory = $true)]
        [string]$Path,
        [Parameter(Mandatory = $true)]
        [AllowEmptyString()]
        [string]$Value
    )

    $parent = Split-Path -Parent $Path
    if (-not (Test-Path -LiteralPath $parent)) {
        New-Item -ItemType Directory -Path $parent -Force | Out-Null
    }
    Set-Content -LiteralPath $Path -Encoding UTF8 -Value $Value
}

function Compress-ExistingPath {
    param(
        [Parameter(Mandatory = $true)]
        [string]$SourcePath,
        [Parameter(Mandatory = $true)]
        [string]$DestinationPath
    )

    $result = [ordered]@{
        created = $false
        skipped = @()
        error = $null
    }

    if (-not (Test-Path -LiteralPath $SourcePath)) {
        $result.error = "Source path does not exist."
        return $result
    }

    $parent = Split-Path -Parent $DestinationPath
    if (-not (Test-Path -LiteralPath $parent)) {
        New-Item -ItemType Directory -Path $parent -Force | Out-Null
    }

    $stagingRoot = Join-Path ([System.IO.Path]::GetTempPath()) ("codex-dashboard-backup-stage-" + [guid]::NewGuid().ToString("n"))
    New-Item -ItemType Directory -Path $stagingRoot -Force | Out-Null

    try {
        $sourceItem = Get-Item -LiteralPath $SourcePath -Force
        if ($sourceItem.PSIsContainer) {
            $stageSource = Join-Path $stagingRoot $sourceItem.Name
            New-Item -ItemType Directory -Path $stageSource -Force | Out-Null
            $sourcePrefix = $sourceItem.FullName.TrimEnd('\') + '\'
            foreach ($file in Get-ChildItem -LiteralPath $SourcePath -Recurse -Force -File -ErrorAction SilentlyContinue) {
                $relative = $file.FullName.Substring($sourcePrefix.Length)
                $target = Join-Path $stageSource $relative
                $targetParent = Split-Path -Parent $target
                New-Item -ItemType Directory -Path $targetParent -Force | Out-Null
                try {
                    Copy-Item -LiteralPath $file.FullName -Destination $target -Force -ErrorAction Stop
                }
                catch {
                    $result.skipped += [ordered]@{
                        path = $file.FullName
                        reason = $_.Exception.Message
                    }
                }
            }
            Compress-Archive -LiteralPath $stageSource -DestinationPath $DestinationPath -Force
        }
        else {
            $stageFile = Join-Path $stagingRoot $sourceItem.Name
            try {
                Copy-Item -LiteralPath $SourcePath -Destination $stageFile -Force -ErrorAction Stop
            }
            catch {
                $result.skipped += [ordered]@{
                    path = $SourcePath
                    reason = $_.Exception.Message
                }
            }
            if (Test-Path -LiteralPath $stageFile) {
                Compress-Archive -LiteralPath $stageFile -DestinationPath $DestinationPath -Force
            }
        }

        $result.created = Test-Path -LiteralPath $DestinationPath
        return $result
    }
    catch {
        $result.error = $_.Exception.Message
        return $result
    }
    finally {
        Remove-Item -LiteralPath $stagingRoot -Recurse -Force -ErrorAction SilentlyContinue
    }
}

function Invoke-ServiceLanePgDumpAll {
    param(
        [Parameter(Mandatory = $true)]
        [hashtable]$Config,
        [Parameter(Mandatory = $true)]
        [string]$DestinationPath
    )

    $dockerPath = Ensure-DockerEngine -Config $Config
    $savedEnvironment = @{
        CODEX_ORCH_POSTGRES_PORT = $env:CODEX_ORCH_POSTGRES_PORT
        CODEX_ORCH_TEMPORAL_PORT = $env:CODEX_ORCH_TEMPORAL_PORT
        CODEX_ORCH_TEMPORAL_UI_PORT = $env:CODEX_ORCH_TEMPORAL_UI_PORT
    }

    try {
        Set-Item -Path Env:CODEX_ORCH_POSTGRES_PORT -Value ([string]$Config.PostgresPort)
        Set-Item -Path Env:CODEX_ORCH_TEMPORAL_PORT -Value ([string]$Config.TemporalPort)
        Set-Item -Path Env:CODEX_ORCH_TEMPORAL_UI_PORT -Value ([string]$Config.TemporalUiPort)

        & $dockerPath compose --project-name $Config.ComposeProject -f $Config.ComposeFile exec -T -e PGPASSWORD=temporal postgres pg_dumpall -U temporal |
            Set-Content -LiteralPath $DestinationPath -Encoding UTF8
        if ($LASTEXITCODE -ne 0) {
            throw "pg_dumpall failed with exit code $LASTEXITCODE."
        }
    }
    finally {
        foreach ($entry in $savedEnvironment.GetEnumerator()) {
            if ([string]::IsNullOrEmpty($entry.Value)) {
                Remove-Item -Path ("Env:{0}" -f $entry.Key) -ErrorAction SilentlyContinue
            }
            else {
                Set-Item -Path ("Env:{0}" -f $entry.Key) -Value $entry.Value
            }
        }
    }
}

function Save-RepoDelta {
    param(
        [Parameter(Mandatory = $true)]
        [string]$RepoRoot,
        [Parameter(Mandatory = $true)]
        [string]$DestinationRoot,
        [switch]$IncludeRepoBundle
    )

    if (-not (Test-Path -LiteralPath $DestinationRoot)) {
        New-Item -ItemType Directory -Path $DestinationRoot -Force | Out-Null
    }

    Write-TextFile -Path (Join-Path $DestinationRoot "branch.txt") -Value (Invoke-GitCapture -RepoRoot $RepoRoot -Arguments @("branch", "--show-current") -AllowFailure)
    Write-TextFile -Path (Join-Path $DestinationRoot "head.txt") -Value (Invoke-GitCapture -RepoRoot $RepoRoot -Arguments @("rev-parse", "HEAD") -AllowFailure)
    Write-TextFile -Path (Join-Path $DestinationRoot "origin.txt") -Value (Invoke-GitCapture -RepoRoot $RepoRoot -Arguments @("remote", "get-url", "origin") -AllowFailure)
    Write-TextFile -Path (Join-Path $DestinationRoot "status-short.txt") -Value (Invoke-GitCapture -RepoRoot $RepoRoot -Arguments @("status", "--short", "--untracked-files=all") -AllowFailure)
    Write-TextFile -Path (Join-Path $DestinationRoot "status.txt") -Value (Invoke-GitCapture -RepoRoot $RepoRoot -Arguments @("status", "--untracked-files=all") -AllowFailure)
    Write-TextFile -Path (Join-Path $DestinationRoot "diff.patch") -Value (Invoke-GitCapture -RepoRoot $RepoRoot -Arguments @("diff", "--binary") -AllowFailure)
    Write-TextFile -Path (Join-Path $DestinationRoot "diff-staged.patch") -Value (Invoke-GitCapture -RepoRoot $RepoRoot -Arguments @("diff", "--cached", "--binary") -AllowFailure)

    $deltaFiles = @()
    $deltaFiles += @((Invoke-GitCapture -RepoRoot $RepoRoot -Arguments @("diff", "--name-only") -AllowFailure) -split "`n")
    $deltaFiles += @((Invoke-GitCapture -RepoRoot $RepoRoot -Arguments @("diff", "--cached", "--name-only") -AllowFailure) -split "`n")
    $deltaFiles += @((Invoke-GitCapture -RepoRoot $RepoRoot -Arguments @("ls-files", "--others", "--exclude-standard") -AllowFailure) -split "`n")
    $deltaFiles = @($deltaFiles | Where-Object { -not [string]::IsNullOrWhiteSpace($_) } | Sort-Object -Unique)
    Write-TextFile -Path (Join-Path $DestinationRoot "delta-files.txt") -Value ($deltaFiles -join "`n")

    if ($deltaFiles.Count -gt 0) {
        $staging = Join-Path $DestinationRoot "delta-file-staging"
        New-Item -ItemType Directory -Path $staging -Force | Out-Null
        foreach ($relative in $deltaFiles) {
            $source = Join-Path $RepoRoot $relative
            if (Test-Path -LiteralPath $source -PathType Leaf) {
                $target = Join-Path $staging $relative
                $targetParent = Split-Path -Parent $target
                New-Item -ItemType Directory -Path $targetParent -Force | Out-Null
                Copy-Item -LiteralPath $source -Destination $target -Force
            }
        }
        Compress-Archive -LiteralPath $staging -DestinationPath (Join-Path $DestinationRoot "delta-files.zip") -Force
        Remove-Item -LiteralPath $staging -Recurse -Force
    }

    if ($IncludeRepoBundle) {
        & git -C $RepoRoot bundle create (Join-Path $DestinationRoot "all-refs.bundle") --all | Out-Null
        if ($LASTEXITCODE -ne 0) {
            throw "git bundle create failed with exit code $LASTEXITCODE."
        }
    }
}
