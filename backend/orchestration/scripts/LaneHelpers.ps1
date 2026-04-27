Set-StrictMode -Version Latest

function Get-OrchestrationLauncherConfigPath {
    return (Join-Path $PSScriptRoot "launcher-config.json")
}

function Get-OrchestrationRepoRoot {
    $launcherConfigPath = Get-OrchestrationLauncherConfigPath
    if (Test-Path $launcherConfigPath) {
        $launcherConfig = Get-Content -Raw -LiteralPath $launcherConfigPath | ConvertFrom-Json
        $sourceRepoRoot = [string]$launcherConfig.source_repo_root
        if ([string]::IsNullOrWhiteSpace($sourceRepoRoot)) {
            throw "Launcher config $launcherConfigPath does not define source_repo_root."
        }
        if (-not (Test-Path $sourceRepoRoot)) {
            throw "Launcher config source_repo_root does not exist: $sourceRepoRoot."
        }
        return (Resolve-Path -LiteralPath $sourceRepoRoot).Path
    }

    return (Resolve-Path (Join-Path $PSScriptRoot "..")).Path
}

function Get-OrchestrationLocalAppData {
    if ($env:LOCALAPPDATA) {
        return $env:LOCALAPPDATA
    }
    return (Join-Path $HOME "AppData\\Local")
}

function Get-OrchestrationLaneConfig {
    param(
        [Parameter(Mandatory = $true)]
        [ValidateSet("service", "validation")]
        [string]$Lane
    )

    $repoRoot = Get-OrchestrationRepoRoot
    $localAppData = Get-OrchestrationLocalAppData
    $dashboardRoot = Join-Path $localAppData "CodexDashboard"
    $sourceRunnerScriptPath = Join-Path $PSScriptRoot "Run-OrchestrationLane.ps1"
    $sourceLaneHelpersPath = Join-Path $PSScriptRoot "LaneHelpers.ps1"

    switch ($Lane) {
        "service" {
            $bindPort = 4318
            $temporalPort = 7233
            $temporalUiPort = 8080
            $postgresPort = 5432
            $runtimeRoot = Join-Path $dashboardRoot "orchestration-service-lane"
            $taskName = "CodexDashboard-Orchestration-ServiceLane"
            $description = "Keeps the CodexDashboard orchestration service lane running at user logon."
            $launcherRoot = Join-Path $runtimeRoot "launcher"
            $runnerScriptPath = Join-Path $launcherRoot "Run-OrchestrationLane.ps1"
        }
        "validation" {
            $bindPort = 14318
            $temporalPort = 17233
            $temporalUiPort = 18080
            $postgresPort = 15432
            $runtimeRoot = Join-Path $dashboardRoot "orchestration-validation-lane"
            $taskName = $null
            $description = $null
            $launcherRoot = $null
            $runnerScriptPath = $sourceRunnerScriptPath
        }
        default {
            throw "Unsupported lane '$Lane'."
        }
    }

    $binaryName = "controlplane-$Lane-lane.exe"
    $bindAddress = "127.0.0.1:$bindPort"

    return @{
        Lane = $Lane
        RepoRoot = $repoRoot
        ComposeFile = Join-Path $repoRoot "dev\\docker-compose.temporal-postgres.yml"
        SourceComposeFile = Join-Path $repoRoot "dev\\docker-compose.temporal-postgres.yml"
        ComposeProject = "codex-orchestration-$Lane"
        BindAddress = $bindAddress
        JobsBackendUrl = "http://$bindAddress"
        TemporalAddress = "127.0.0.1:$temporalPort"
        TemporalUiUrl = "http://127.0.0.1:$temporalUiPort"
        PostgresPort = $postgresPort
        TemporalPort = $temporalPort
        TemporalUiPort = $temporalUiPort
        RuntimeRoot = $runtimeRoot
        ReleasesRoot = Join-Path $runtimeRoot "releases"
        CurrentReleaseManifestPath = Join-Path $runtimeRoot "current-release.json"
        BinaryPath = Join-Path $runtimeRoot "bin\\$binaryName"
        BinaryName = $binaryName
        LogPath = Join-Path $runtimeRoot "logs\\controlplane.log"
        StdoutLogPath = Join-Path $runtimeRoot "logs\\controlplane.stdout.log"
        StderrLogPath = Join-Path $runtimeRoot "logs\\controlplane.stderr.log"
        RunsRoot = Join-Path $dashboardRoot "orchestration-runs\\$Lane-lane"
        TaskName = $taskName
        TaskDescription = $description
        LauncherRoot = $launcherRoot
        LauncherConfigPath = if ($null -eq $launcherRoot) { $null } else { Join-Path $launcherRoot "launcher-config.json" }
        RunnerScriptPath = $runnerScriptPath
        SourceRunnerScriptPath = $sourceRunnerScriptPath
        SourceLaneHelpersPath = $sourceLaneHelpersPath
    }
}

function Ensure-OrchestrationLaneDirectories {
    param(
        [Parameter(Mandatory = $true)]
        [hashtable]$Config
    )

    foreach ($path in @(
        $Config.RuntimeRoot,
        $Config.ReleasesRoot,
        (Split-Path -Parent $Config.BinaryPath),
        (Split-Path -Parent $Config.LogPath),
        (Split-Path -Parent $Config.StdoutLogPath),
        (Split-Path -Parent $Config.StderrLogPath),
        $Config.RunsRoot
    )) {
        if ([string]::IsNullOrWhiteSpace($path)) {
            continue
        }
        if (-not (Test-Path $path)) {
            New-Item -ItemType Directory -Path $path -Force | Out-Null
        }
    }
}

function Write-OrchestrationLaneLog {
    param(
        [Parameter(Mandatory = $true)]
        [hashtable]$Config,
        [Parameter(Mandatory = $true)]
        [string]$Message
    )

    Ensure-OrchestrationLaneDirectories -Config $Config
    $timestamp = Get-Date -Format o
    Add-Content -Path $Config.LogPath -Value "[$timestamp] $Message"
}

function Get-GoExecutablePath {
    $command = Get-Command "go.exe" -ErrorAction SilentlyContinue
    if ($null -ne $command) {
        return $command.Source
    }

    $fallback = "C:\\Program Files\\Go\\bin\\go.exe"
    if (Test-Path $fallback) {
        return $fallback
    }

    throw "Could not resolve go.exe."
}

function Get-PowerShellExecutablePath {
    $stableCandidates = @()
    if (-not [string]::IsNullOrWhiteSpace($env:ProgramFiles)) {
        $stableCandidates += (Join-Path $env:ProgramFiles "PowerShell\7\pwsh.exe")
    }
    if (-not [string]::IsNullOrWhiteSpace($env:SystemRoot)) {
        $stableCandidates += (Join-Path $env:SystemRoot "System32\WindowsPowerShell\v1.0\powershell.exe")
    }

    foreach ($candidate in $stableCandidates) {
        if ([string]::IsNullOrWhiteSpace($candidate)) {
            continue
        }
        if (Test-Path $candidate) {
            return $candidate
        }
    }

    # Store-installed PowerShell resolves to a versioned WindowsApps package path.
    # Capturing that path in a Scheduled Task breaks after PowerShell updates.
    $fallbackCandidates = @(
        (Join-Path $PSHOME "pwsh.exe"),
        (Join-Path $PSHOME "powershell.exe")
    )
    foreach ($commandName in @("pwsh.exe", "powershell.exe")) {
        $command = Get-Command $commandName -ErrorAction SilentlyContinue
        if ($null -ne $command -and -not [string]::IsNullOrWhiteSpace($command.Source)) {
            $fallbackCandidates += $command.Source
        }
    }

    foreach ($candidate in $fallbackCandidates) {
        if ([string]::IsNullOrWhiteSpace($candidate)) {
            continue
        }
        if ($candidate -like "*\\WindowsApps\\Microsoft.PowerShell_*") {
            continue
        }
        if (Test-Path $candidate) {
            return $candidate
        }
    }

    throw "Could not resolve a PowerShell executable for the scheduled task runner."
}

function Get-DockerExecutablePath {
    $command = Get-Command "docker.exe" -ErrorAction SilentlyContinue
    if ($null -ne $command) {
        return $command.Source
    }

    $fallback = "C:\\Program Files\\Docker\\Docker\\resources\\bin\\docker.exe"
    if (Test-Path $fallback) {
        return $fallback
    }

    throw "Could not resolve docker.exe."
}

function Get-DockerDesktopExecutablePath {
    $fallback = "C:\\Program Files\\Docker\\Docker\\Docker Desktop.exe"
    if (Test-Path $fallback) {
        return $fallback
    }
    return $null
}

function Resolve-OrchestrationFullPath {
    param(
        [Parameter(Mandatory = $true)]
        [string]$Path
    )

    return [System.IO.Path]::GetFullPath($Path)
}

function Test-OrchestrationPathWithinRoot {
    param(
        [Parameter(Mandatory = $true)]
        [string]$Path,
        [Parameter(Mandatory = $true)]
        [string]$Root
    )

    $fullPath = Resolve-OrchestrationFullPath -Path $Path
    $fullRoot = (Resolve-OrchestrationFullPath -Path $Root).TrimEnd("\")
    return $fullPath.Equals($fullRoot, [System.StringComparison]::OrdinalIgnoreCase) -or
        $fullPath.StartsWith(($fullRoot + "\"), [System.StringComparison]::OrdinalIgnoreCase)
}

function Get-OrchestrationFileSha256 {
    param(
        [Parameter(Mandatory = $true)]
        [string]$Path
    )

    $stream = [System.IO.File]::OpenRead($Path)
    try {
        $sha256 = [System.Security.Cryptography.SHA256]::Create()
        try {
            $hashBytes = $sha256.ComputeHash($stream)
            return -join ($hashBytes | ForEach-Object { $_.ToString("x2") })
        }
        finally {
            $sha256.Dispose()
        }
    }
    finally {
        $stream.Dispose()
    }
}

function Get-GitExecutablePath {
    $command = Get-Command "git.exe" -ErrorAction SilentlyContinue
    if ($null -ne $command) {
        return $command.Source
    }

    $command = Get-Command "git" -ErrorAction SilentlyContinue
    if ($null -ne $command) {
        return $command.Source
    }

    throw "Could not resolve git."
}

function Get-OrchestrationGitCommit {
    param(
        [Parameter(Mandatory = $true)]
        [string]$RepoRoot
    )

    $gitPath = Get-GitExecutablePath
    $commit = (& $gitPath -C $RepoRoot rev-parse HEAD)
    if ($LASTEXITCODE -ne 0 -or [string]::IsNullOrWhiteSpace($commit)) {
        throw "Could not resolve git commit for $RepoRoot."
    }
    return [string]$commit
}

function Get-OrchestrationGitStatusShort {
    param(
        [Parameter(Mandatory = $true)]
        [string]$RepoRoot
    )

    $gitPath = Get-GitExecutablePath
    $status = @(& $gitPath -C $RepoRoot status --short)
    if ($LASTEXITCODE -ne 0) {
        throw "Could not resolve git status for $RepoRoot."
    }
    return $status
}

function Test-DockerEngineReady {
    param(
        [Parameter(Mandatory = $true)]
        [string]$DockerPath
    )

    & $DockerPath version | Out-Null
    return ($LASTEXITCODE -eq 0)
}

function Ensure-DockerEngine {
    param(
        [Parameter(Mandatory = $true)]
        [hashtable]$Config,
        [int]$TimeoutSeconds = 180
    )

    $dockerPath = Get-DockerExecutablePath
    if (Test-DockerEngineReady -DockerPath $dockerPath) {
        return $dockerPath
    }

    $dockerDesktopPath = Get-DockerDesktopExecutablePath
    if ($null -ne $dockerDesktopPath) {
        Write-OrchestrationLaneLog -Config $Config -Message "Docker engine is not ready. Starting Docker Desktop."
        Start-Process -FilePath $dockerDesktopPath | Out-Null
    }

    $deadline = (Get-Date).AddSeconds($TimeoutSeconds)
    while ((Get-Date) -lt $deadline) {
        Start-Sleep -Seconds 5
        if (Test-DockerEngineReady -DockerPath $dockerPath) {
            return $dockerPath
        }
    }

    throw "Docker engine did not become ready within $TimeoutSeconds seconds."
}

function Invoke-OrchestrationCompose {
    param(
        [Parameter(Mandatory = $true)]
        [hashtable]$Config,
        [Parameter(Mandatory = $true)]
        [string[]]$ComposeArgs,
        [switch]$AllowFailure
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

        & $dockerPath compose --project-name $Config.ComposeProject -f $Config.ComposeFile @ComposeArgs
        $exitCode = $LASTEXITCODE
        if (-not $AllowFailure -and $exitCode -ne 0) {
            throw "docker compose $($ComposeArgs -join ' ') failed with exit code $exitCode."
        }
        return $exitCode
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

function Build-OrchestrationBinary {
    param(
        [Parameter(Mandatory = $true)]
        [hashtable]$Config,
        [switch]$AllowExistingOnFailure,
        [string]$OutputPath
    )

    Ensure-OrchestrationLaneDirectories -Config $Config
    $goPath = Get-GoExecutablePath
    $binaryPath = $Config.BinaryPath
    if (-not [string]::IsNullOrWhiteSpace($OutputPath)) {
        $binaryPath = $OutputPath
    }

    $binaryParent = Split-Path -Parent $binaryPath
    if (-not (Test-Path $binaryParent)) {
        New-Item -ItemType Directory -Path $binaryParent -Force | Out-Null
    }

    Push-Location $Config.RepoRoot
    try {
        & $goPath build -o $binaryPath .\\cmd\\controlplane
        $exitCode = $LASTEXITCODE
    }
    finally {
        Pop-Location
    }

    if ($exitCode -ne 0) {
        if ($AllowExistingOnFailure -and (Test-Path $binaryPath)) {
            Write-OrchestrationLaneLog -Config $Config -Message "go build failed; reusing existing binary at $binaryPath."
            return $binaryPath
        }
        throw "go build failed with exit code $exitCode."
    }

    return $binaryPath
}

function Assert-ServiceLaneConfig {
    param(
        [Parameter(Mandatory = $true)]
        [hashtable]$Config
    )

    if ($Config.Lane -ne "service") {
        throw "Pinned release operations are only valid for the service lane."
    }
}

function Install-ServiceLaneLauncher {
    param(
        [Parameter(Mandatory = $true)]
        [hashtable]$Config
    )

    Assert-ServiceLaneConfig -Config $Config
    Ensure-OrchestrationLaneDirectories -Config $Config
    if (-not (Test-Path $Config.LauncherRoot)) {
        New-Item -ItemType Directory -Path $Config.LauncherRoot -Force | Out-Null
    }

    Copy-Item -LiteralPath $Config.SourceRunnerScriptPath -Destination (Join-Path $Config.LauncherRoot "Run-OrchestrationLane.ps1") -Force
    Copy-Item -LiteralPath $Config.SourceLaneHelpersPath -Destination (Join-Path $Config.LauncherRoot "LaneHelpers.ps1") -Force

    $launcherConfig = [ordered]@{
        schema_version = 1
        source_repo_root = $Config.RepoRoot
        installed_at = (Get-Date).ToUniversalTime().ToString("o")
        runner_script_path = $Config.RunnerScriptPath
    }
    $launcherConfig | ConvertTo-Json -Depth 6 | Set-Content -LiteralPath $Config.LauncherConfigPath -Encoding UTF8
}

function Set-ServiceLaneCurrentRelease {
    param(
        [Parameter(Mandatory = $true)]
        [hashtable]$Config,
        [Parameter(Mandatory = $true)]
        [string]$ManifestPath
    )

    Assert-ServiceLaneConfig -Config $Config
    if (-not (Test-Path $ManifestPath)) {
        throw "Release manifest does not exist: $ManifestPath."
    }

    $manifestJson = Get-Content -Raw -LiteralPath $ManifestPath
    $manifest = $manifestJson | ConvertFrom-Json
    if ([string]$manifest.lane -ne "service") {
        throw "Release manifest $ManifestPath is not for the service lane."
    }

    $currentParent = Split-Path -Parent $Config.CurrentReleaseManifestPath
    if (-not (Test-Path $currentParent)) {
        New-Item -ItemType Directory -Path $currentParent -Force | Out-Null
    }

    $tempPath = "$($Config.CurrentReleaseManifestPath).tmp"
    Set-Content -LiteralPath $tempPath -Value $manifestJson -Encoding UTF8
    Move-Item -LiteralPath $tempPath -Destination $Config.CurrentReleaseManifestPath -Force
}

function Get-ServiceLaneCurrentRelease {
    param(
        [Parameter(Mandatory = $true)]
        [hashtable]$Config
    )

    Assert-ServiceLaneConfig -Config $Config
    if (-not (Test-Path $Config.CurrentReleaseManifestPath)) {
        throw "No pinned service-lane release manifest exists at $($Config.CurrentReleaseManifestPath). Publish a service-lane release before starting the human lane."
    }

    $manifest = Get-Content -Raw -LiteralPath $Config.CurrentReleaseManifestPath | ConvertFrom-Json
    if ([string]$manifest.lane -ne "service") {
        throw "Pinned release manifest is not for the service lane: $($Config.CurrentReleaseManifestPath)."
    }

    $binaryPath = [string]$manifest.binary_path
    $composeFilePath = [string]$manifest.compose_file_path
    if ([string]::IsNullOrWhiteSpace($binaryPath)) {
        throw "Pinned service-lane release manifest does not define binary_path."
    }
    if ([string]::IsNullOrWhiteSpace($composeFilePath)) {
        throw "Pinned service-lane release manifest does not define compose_file_path."
    }
    if (-not (Test-OrchestrationPathWithinRoot -Path $binaryPath -Root $Config.ReleasesRoot)) {
        throw "Pinned service-lane binary is outside the release root: $binaryPath."
    }
    if (-not (Test-OrchestrationPathWithinRoot -Path $composeFilePath -Root $Config.ReleasesRoot)) {
        throw "Pinned service-lane compose file is outside the release root: $composeFilePath."
    }
    if (-not (Test-Path $binaryPath)) {
        throw "Pinned service-lane binary does not exist: $binaryPath."
    }
    if (-not (Test-Path $composeFilePath)) {
        throw "Pinned service-lane compose file does not exist: $composeFilePath."
    }

    $binaryHash = Get-OrchestrationFileSha256 -Path $binaryPath
    if ($binaryHash -ne [string]$manifest.binary_sha256) {
        throw "Pinned service-lane binary hash mismatch for $binaryPath."
    }

    $composeHash = Get-OrchestrationFileSha256 -Path $composeFilePath
    if ($composeHash -ne [string]$manifest.compose_file_sha256) {
        throw "Pinned service-lane compose file hash mismatch for $composeFilePath."
    }

    return $manifest
}

function New-ServiceLaneRelease {
    param(
        [Parameter(Mandatory = $true)]
        [hashtable]$Config,
        [switch]$AllowDirty,
        [switch]$PinCurrent
    )

    Assert-ServiceLaneConfig -Config $Config
    Ensure-OrchestrationLaneDirectories -Config $Config
    if (-not (Test-Path $Config.SourceComposeFile)) {
        throw "Source compose file does not exist: $($Config.SourceComposeFile)."
    }

    $commit = Get-OrchestrationGitCommit -RepoRoot $Config.RepoRoot
    $statusLines = @(Get-OrchestrationGitStatusShort -RepoRoot $Config.RepoRoot)
    if ($statusLines.Count -gt 0 -and -not $AllowDirty) {
        throw "Refusing to publish a service-lane release from a dirty repo. Commit/stash changes or pass -AllowDirty to record the dirty status in the manifest."
    }

    $createdAt = (Get-Date).ToUniversalTime()
    $shortCommit = $commit.Substring(0, [Math]::Min(12, $commit.Length))
    $releaseId = "{0}-{1}" -f $createdAt.ToString("yyyyMMddTHHmmssZ"), $shortCommit
    $releaseRoot = Join-Path $Config.ReleasesRoot $releaseId
    $binaryPath = Join-Path $releaseRoot ("bin\\{0}" -f $Config.BinaryName)
    $composeFilePath = Join-Path $releaseRoot "docker-compose.temporal-postgres.yml"
    $manifestPath = Join-Path $releaseRoot "release-manifest.json"

    New-Item -ItemType Directory -Path (Split-Path -Parent $binaryPath) -Force | Out-Null
    Copy-Item -LiteralPath $Config.SourceComposeFile -Destination $composeFilePath -Force
    $builtBinaryPath = Build-OrchestrationBinary -Config $Config -OutputPath $binaryPath

    $manifest = [ordered]@{
        schema_version = 1
        lane = "service"
        release_id = $releaseId
        release_root = $releaseRoot
        created_at = $createdAt.ToString("o")
        source_repo_root = $Config.RepoRoot
        git_commit = $commit
        source_dirty = ($statusLines.Count -gt 0)
        source_status = $statusLines
        binary_path = $builtBinaryPath
        binary_sha256 = Get-OrchestrationFileSha256 -Path $builtBinaryPath
        compose_file_path = $composeFilePath
        compose_file_sha256 = Get-OrchestrationFileSha256 -Path $composeFilePath
        runner_script_path = $Config.RunnerScriptPath
    }

    $manifest | ConvertTo-Json -Depth 8 | Set-Content -LiteralPath $manifestPath -Encoding UTF8
    if ($PinCurrent) {
        Set-ServiceLaneCurrentRelease -Config $Config -ManifestPath $manifestPath
    }

    return (Get-Content -Raw -LiteralPath $manifestPath | ConvertFrom-Json)
}

function Set-OrchestrationLaneEnvironment {
    param(
        [Parameter(Mandatory = $true)]
        [hashtable]$Config
    )

    Set-Item -Path Env:CODEX_ORCHESTRATION_BIND_ADDRESS -Value $Config.BindAddress
    Set-Item -Path Env:CODEX_ORCHESTRATION_TEMPORAL_ADDRESS -Value $Config.TemporalAddress
    Set-Item -Path Env:CODEX_ORCHESTRATION_RUNS_ROOT -Value $Config.RunsRoot
    Set-Item -Path Env:CODEX_ORCHESTRATION_WORKTREE_ROOT -Value $Config.RepoRoot
    Set-Item -Path Env:CODEX_ORCHESTRATION_TRACKING_ROOT -Value (Join-Path $Config.RepoRoot "Tracking")
}

function Wait-OrchestrationLaneHealth {
    param(
        [Parameter(Mandatory = $true)]
        [hashtable]$Config,
        [int]$TimeoutSeconds = 180
    )

    $deadline = (Get-Date).AddSeconds($TimeoutSeconds)
    while ((Get-Date) -lt $deadline) {
        try {
            return Invoke-RestMethod -Uri "$($Config.JobsBackendUrl)/healthz" -TimeoutSec 5
        }
        catch {
            Start-Sleep -Seconds 3
        }
    }

    throw "Timed out waiting for $($Config.Lane) lane health at $($Config.JobsBackendUrl)/healthz."
}

function Get-OrchestrationLaneBinaryProcess {
    param(
        [Parameter(Mandatory = $true)]
        [hashtable]$Config
    )

    $processName = [System.IO.Path]::GetFileNameWithoutExtension($Config.BinaryName)
    return Get-Process -Name $processName -ErrorAction SilentlyContinue
}

function Stop-OrchestrationLaneProcesses {
    param(
        [Parameter(Mandatory = $true)]
        [hashtable]$Config
    )

    $binaryProcesses = Get-OrchestrationLaneBinaryProcess -Config $Config
    if ($null -ne $binaryProcesses) {
        $binaryProcesses | Stop-Process -Force
    }

    $runnerScriptPaths = @($Config.RunnerScriptPath, $Config.SourceRunnerScriptPath) |
        Where-Object { -not [string]::IsNullOrWhiteSpace($_) } |
        Select-Object -Unique
    $runnerProcesses = Get-CimInstance Win32_Process -Filter "Name = 'powershell.exe'" |
        Where-Object {
            if ($null -eq $_.CommandLine -or $_.CommandLine -notlike "*-Lane $($Config.Lane)*") {
                $false
            }
            else {
                $matchesRunner = $false
                foreach ($runnerScriptPath in $runnerScriptPaths) {
                    $escapedRunnerScriptPath = $runnerScriptPath.Replace("\", "\\")
                    if ($_.CommandLine -like "*$runnerScriptPath*" -or $_.CommandLine -like "*$escapedRunnerScriptPath*") {
                        $matchesRunner = $true
                        break
                    }
                }
                $matchesRunner
            }
        }

    foreach ($process in $runnerProcesses) {
        Stop-Process -Id $process.ProcessId -Force
    }
}

function Get-OrchestrationLaneStatus {
    param(
        [Parameter(Mandatory = $true)]
        [hashtable]$Config
    )

    $binaryProcesses = @(Get-OrchestrationLaneBinaryProcess -Config $Config)
    $processPath = $null
    if ($binaryProcesses.Count -gt 0) {
        try {
            $processPath = [string]$binaryProcesses[0].Path
        }
        catch {
            $processPath = $null
        }
    }

    $status = [ordered]@{
        lane = $Config.Lane
        jobs_backend_url = $Config.JobsBackendUrl
        temporal_address = $Config.TemporalAddress
        temporal_ui_url = $Config.TemporalUiUrl
        runs_root = $Config.RunsRoot
        log_path = $Config.LogPath
        task_name = $Config.TaskName
        task_state = $null
        task_action = $null
        task_uses_pinned_launcher = $null
        current_release = $null
        current_release_error = $null
        process_running = ($binaryProcesses.Count -gt 0)
        process_path = $processPath
        health = $null
        jobs = @()
        last_error = $null
    }

    if (-not [string]::IsNullOrWhiteSpace($Config.TaskName)) {
        $task = Get-ScheduledTask -TaskName $Config.TaskName -ErrorAction SilentlyContinue
        if ($null -ne $task) {
            $status.task_state = [string]$task.State
            $action = @($task.Actions) | Select-Object -First 1
            if ($null -ne $action) {
                $status.task_action = [string]$action.Arguments
            }
            $status.task_uses_pinned_launcher = Test-ServiceLaneTaskUsesPinnedLauncher -Config $Config
        }
        else {
            $status.task_state = "missing"
        }
    }

    if ($Config.Lane -eq "service") {
        try {
            $release = Get-ServiceLaneCurrentRelease -Config $Config
            $status.current_release = [ordered]@{
                release_id = [string]$release.release_id
                git_commit = [string]$release.git_commit
                source_dirty = [bool]$release.source_dirty
                binary_path = [string]$release.binary_path
                binary_sha256 = [string]$release.binary_sha256
                compose_file_path = [string]$release.compose_file_path
                compose_file_sha256 = [string]$release.compose_file_sha256
            }
        }
        catch {
            $status.current_release_error = $_.Exception.Message
        }
    }

    try {
        $status.health = Invoke-RestMethod -Uri "$($Config.JobsBackendUrl)/healthz" -TimeoutSec 5
        $jobsResponse = Invoke-RestMethod -Uri "$($Config.JobsBackendUrl)/api/v1/jobs" -TimeoutSec 5
        foreach ($job in @($jobsResponse.jobs)) {
            $nextActionTime = $null
            $schedules = @($job.schedules)
            if ($schedules.Count -gt 0) {
                $nextActionTimes = @($schedules[0].next_action_times)
                if ($nextActionTimes.Count -gt 0) {
                    $nextActionTime = [string]$nextActionTimes[0]
                }
            }

            $status.jobs += [ordered]@{
                job_id = [string]$job.job_id
                status = [string]$job.status
                next_action_time = $nextActionTime
            }
        }
    }
    catch {
        $status.last_error = $_.Exception.Message
    }

    return [pscustomobject]$status
}

function Get-CurrentInteractiveUser {
    return "{0}\\{1}" -f $env:USERDOMAIN, $env:USERNAME
}

function Test-ServiceLaneTaskUsesPinnedLauncher {
    param(
        [Parameter(Mandatory = $true)]
        [hashtable]$Config
    )

    Assert-ServiceLaneConfig -Config $Config
    $task = Get-ScheduledTask -TaskName $Config.TaskName -ErrorAction SilentlyContinue
    if ($null -eq $task) {
        return $false
    }

    $action = @($task.Actions) | Select-Object -First 1
    if ($null -eq $action) {
        return $false
    }

    $arguments = [string]$action.Arguments
    $expectedRunner = [string]$Config.RunnerScriptPath
    $repoRunner = Join-Path (Join-Path $Config.RepoRoot "scripts") "Run-OrchestrationLane.ps1"
    return ($arguments -like "*$expectedRunner*") -and ($arguments -notlike "*$repoRunner*")
}

function Assert-ServiceLaneTaskUsesPinnedLauncher {
    param(
        [Parameter(Mandatory = $true)]
        [hashtable]$Config
    )

    if (-not (Test-ServiceLaneTaskUsesPinnedLauncher -Config $Config)) {
        throw "Service-lane scheduled task is not pinned to the runtime launcher at $($Config.RunnerScriptPath). Run Install-ServiceLane.ps1 after publishing a service-lane release."
    }
}

function Register-ServiceLaneTask {
    param(
        [Parameter(Mandatory = $true)]
        [hashtable]$Config
    )

    if ([string]::IsNullOrWhiteSpace($Config.TaskName)) {
        throw "The requested lane does not define a scheduled task."
    }

    Install-ServiceLaneLauncher -Config $Config
    $powershellPath = Get-PowerShellExecutablePath
    $currentUser = Get-CurrentInteractiveUser
    $actionArguments = "-NoProfile -ExecutionPolicy Bypass -WindowStyle Hidden -File `"$($Config.RunnerScriptPath)`" -Lane service -Supervise"
    $action = New-ScheduledTaskAction -Execute $powershellPath -Argument $actionArguments
    $trigger = New-ScheduledTaskTrigger -AtLogOn -User $currentUser
    $principal = New-ScheduledTaskPrincipal -UserId $currentUser -LogonType Interactive -RunLevel Limited
    $settings = New-ScheduledTaskSettingsSet `
        -AllowStartIfOnBatteries `
        -DontStopIfGoingOnBatteries `
        -StartWhenAvailable `
        -MultipleInstances IgnoreNew `
        -ExecutionTimeLimit ([TimeSpan]::Zero) `
        -RestartCount 999 `
        -RestartInterval (New-TimeSpan -Minutes 1)

    Register-ScheduledTask -TaskName $Config.TaskName -Action $action -Trigger $trigger -Principal $principal -Settings $settings -Description $Config.TaskDescription -Force | Out-Null
}
