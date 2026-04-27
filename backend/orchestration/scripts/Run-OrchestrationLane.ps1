[CmdletBinding()]
param(
    [Parameter(Mandatory = $true)]
    [ValidateSet("service", "validation")]
    [string]$Lane,
    [switch]$Supervise
)

$ErrorActionPreference = "Stop"
Set-StrictMode -Version Latest

. (Join-Path $PSScriptRoot "LaneHelpers.ps1")

$config = Get-OrchestrationLaneConfig -Lane $Lane
Write-OrchestrationLaneLog -Config $config -Message "Starting $Lane lane runner."

while ($true) {
    try {
        $workingDirectory = $config.RepoRoot
        if ($Lane -eq "service") {
            $release = Get-ServiceLaneCurrentRelease -Config $config
            $config.ComposeFile = [string]$release.compose_file_path
            $binaryPath = [string]$release.binary_path
            if (-not [string]::IsNullOrWhiteSpace([string]$release.release_root)) {
                $workingDirectory = [string]$release.release_root
            }
        }
        else {
            $binaryPath = Build-OrchestrationBinary -Config $config -AllowExistingOnFailure
        }

        Invoke-OrchestrationCompose -Config $config -ComposeArgs @("up", "-d") | Out-Null
        Set-OrchestrationLaneEnvironment -Config $config
        Write-OrchestrationLaneLog -Config $config -Message "Launching $binaryPath."

        if (Test-Path $config.StdoutLogPath) {
            Remove-Item $config.StdoutLogPath -Force
        }
        if (Test-Path $config.StderrLogPath) {
            Remove-Item $config.StderrLogPath -Force
        }

        $process = Start-Process -FilePath $binaryPath -WorkingDirectory $workingDirectory -PassThru -Wait -RedirectStandardOutput $config.StdoutLogPath -RedirectStandardError $config.StderrLogPath
        $exitCode = $process.ExitCode

        Write-OrchestrationLaneLog -Config $config -Message "Control plane exited with code $exitCode."
    }
    catch {
        $exitCode = 1
        Write-OrchestrationLaneLog -Config $config -Message "Runner loop failed: $($_.Exception.Message)"
    }

    if (-not $Supervise) {
        exit $exitCode
    }

    Start-Sleep -Seconds 10
}
