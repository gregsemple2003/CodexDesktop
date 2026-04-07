[CmdletBinding()]
param(
    [switch]$Detached = $true
)

$ErrorActionPreference = "Stop"
Set-StrictMode -Version Latest

. (Join-Path $PSScriptRoot "LaneHelpers.ps1")

$config = Get-OrchestrationLaneConfig -Lane "validation"
if ($Detached) {
    Stop-OrchestrationLaneProcesses -Config $config
    $powershellPath = Join-Path $PSHOME "powershell.exe"
    $runnerScriptPath = Join-Path $PSScriptRoot "Run-OrchestrationLane.ps1"
    Start-Process -FilePath $powershellPath -ArgumentList @(
        "-NoProfile",
        "-ExecutionPolicy", "Bypass",
        "-WindowStyle", "Hidden",
        "-File", $runnerScriptPath,
        "-Lane", "validation",
        "-Supervise"
    ) | Out-Null
    Wait-OrchestrationLaneHealth -Config $config | Out-Null
    Get-OrchestrationLaneStatus -Config $config | ConvertTo-Json -Depth 8
    return
}

& (Join-Path $PSScriptRoot "Run-OrchestrationLane.ps1") -Lane validation
