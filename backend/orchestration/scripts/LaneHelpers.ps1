Set-StrictMode -Version Latest

function Get-OrchestrationRepoRoot {
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

    switch ($Lane) {
        "service" {
            $bindPort = 4318
            $temporalPort = 7233
            $temporalUiPort = 8080
            $postgresPort = 5432
            $runtimeRoot = Join-Path $dashboardRoot "orchestration-service-lane"
            $taskName = "CodexDashboard-Orchestration-ServiceLane"
            $description = "Keeps the CodexDashboard orchestration service lane running at user logon."
        }
        "validation" {
            $bindPort = 14318
            $temporalPort = 17233
            $temporalUiPort = 18080
            $postgresPort = 15432
            $runtimeRoot = Join-Path $dashboardRoot "orchestration-validation-lane"
            $taskName = $null
            $description = $null
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
        ComposeProject = "codex-orchestration-$Lane"
        BindAddress = $bindAddress
        JobsBackendUrl = "http://$bindAddress"
        TemporalAddress = "127.0.0.1:$temporalPort"
        TemporalUiUrl = "http://127.0.0.1:$temporalUiPort"
        PostgresPort = $postgresPort
        TemporalPort = $temporalPort
        TemporalUiPort = $temporalUiPort
        RuntimeRoot = $runtimeRoot
        BinaryPath = Join-Path $runtimeRoot "bin\\$binaryName"
        BinaryName = $binaryName
        LogPath = Join-Path $runtimeRoot "logs\\controlplane.log"
        StdoutLogPath = Join-Path $runtimeRoot "logs\\controlplane.stdout.log"
        StderrLogPath = Join-Path $runtimeRoot "logs\\controlplane.stderr.log"
        RunsRoot = Join-Path $dashboardRoot "orchestration-runs\\$Lane-lane"
        TaskName = $taskName
        TaskDescription = $description
        RunnerScriptPath = Join-Path $PSScriptRoot "Run-OrchestrationLane.ps1"
    }
}

function Ensure-OrchestrationLaneDirectories {
    param(
        [Parameter(Mandatory = $true)]
        [hashtable]$Config
    )

    foreach ($path in @(
        $Config.RuntimeRoot,
        (Split-Path -Parent $Config.BinaryPath),
        (Split-Path -Parent $Config.LogPath),
        (Split-Path -Parent $Config.StdoutLogPath),
        (Split-Path -Parent $Config.StderrLogPath),
        $Config.RunsRoot
    )) {
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
    $candidates = @(
        (Join-Path $PSHOME "pwsh.exe"),
        (Join-Path $PSHOME "powershell.exe"),
        "C:\\Program Files\\PowerShell\\7\\pwsh.exe",
        "C:\\Windows\\System32\\WindowsPowerShell\\v1.0\\powershell.exe"
    )

    foreach ($candidate in $candidates) {
        if ([string]::IsNullOrWhiteSpace($candidate)) {
            continue
        }
        if (Test-Path $candidate) {
            return $candidate
        }
    }

    foreach ($commandName in @("pwsh.exe", "powershell.exe")) {
        $command = Get-Command $commandName -ErrorAction SilentlyContinue
        if ($null -ne $command -and -not [string]::IsNullOrWhiteSpace($command.Source)) {
            return $command.Source
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
        [switch]$AllowExistingOnFailure
    )

    Ensure-OrchestrationLaneDirectories -Config $Config
    $goPath = Get-GoExecutablePath

    Push-Location $Config.RepoRoot
    try {
        & $goPath build -o $Config.BinaryPath .\\cmd\\controlplane
        $exitCode = $LASTEXITCODE
    }
    finally {
        Pop-Location
    }

    if ($exitCode -ne 0) {
        if ($AllowExistingOnFailure -and (Test-Path $Config.BinaryPath)) {
            Write-OrchestrationLaneLog -Config $Config -Message "go build failed; reusing existing binary at $($Config.BinaryPath)."
            return $Config.BinaryPath
        }
        throw "go build failed with exit code $exitCode."
    }

    return $Config.BinaryPath
}

function Set-OrchestrationLaneEnvironment {
    param(
        [Parameter(Mandatory = $true)]
        [hashtable]$Config
    )

    Set-Item -Path Env:CODEX_ORCHESTRATION_BIND_ADDRESS -Value $Config.BindAddress
    Set-Item -Path Env:CODEX_ORCHESTRATION_TEMPORAL_ADDRESS -Value $Config.TemporalAddress
    Set-Item -Path Env:CODEX_ORCHESTRATION_RUNS_ROOT -Value $Config.RunsRoot
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

    $runnerScriptPath = $Config.RunnerScriptPath.Replace("\", "\\")
    $runnerProcesses = Get-CimInstance Win32_Process -Filter "Name = 'powershell.exe'" |
        Where-Object {
            $null -ne $_.CommandLine -and
            $_.CommandLine -like "*$runnerScriptPath*" -and
            $_.CommandLine -like "*-Lane $($Config.Lane)*"
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

    $status = [ordered]@{
        lane = $Config.Lane
        jobs_backend_url = $Config.JobsBackendUrl
        temporal_address = $Config.TemporalAddress
        temporal_ui_url = $Config.TemporalUiUrl
        runs_root = $Config.RunsRoot
        log_path = $Config.LogPath
        task_name = $Config.TaskName
        task_state = $null
        process_running = [bool](Get-OrchestrationLaneBinaryProcess -Config $Config)
        health = $null
        jobs = @()
        last_error = $null
    }

    if (-not [string]::IsNullOrWhiteSpace($Config.TaskName)) {
        $task = Get-ScheduledTask -TaskName $Config.TaskName -ErrorAction SilentlyContinue
        if ($null -ne $task) {
            $status.task_state = [string]$task.State
        }
        else {
            $status.task_state = "missing"
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

function Register-ServiceLaneTask {
    param(
        [Parameter(Mandatory = $true)]
        [hashtable]$Config
    )

    if ([string]::IsNullOrWhiteSpace($Config.TaskName)) {
        throw "The requested lane does not define a scheduled task."
    }

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
