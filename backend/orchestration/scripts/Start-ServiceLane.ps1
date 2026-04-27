[CmdletBinding()]
param()

$ErrorActionPreference = "Stop"
Set-StrictMode -Version Latest

. (Join-Path $PSScriptRoot "LaneHelpers.ps1")

$config = Get-OrchestrationLaneConfig -Lane "service"
$task = Get-ScheduledTask -TaskName $config.TaskName -ErrorAction SilentlyContinue
if ($null -eq $task) {
    throw "The service lane is not installed yet. Run Install-ServiceLane.ps1 first."
}

Get-ServiceLaneCurrentRelease -Config $config | Out-Null
Assert-ServiceLaneTaskUsesPinnedLauncher -Config $config
Start-ScheduledTask -TaskName $config.TaskName
Wait-OrchestrationLaneHealth -Config $config | Out-Null
Get-OrchestrationLaneStatus -Config $config | ConvertTo-Json -Depth 8
