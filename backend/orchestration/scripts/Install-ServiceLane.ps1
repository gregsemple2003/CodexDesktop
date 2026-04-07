[CmdletBinding()]
param()

$ErrorActionPreference = "Stop"
Set-StrictMode -Version Latest

. (Join-Path $PSScriptRoot "LaneHelpers.ps1")

$config = Get-OrchestrationLaneConfig -Lane "service"
Ensure-OrchestrationLaneDirectories -Config $config
Register-ServiceLaneTask -Config $config
Start-ScheduledTask -TaskName $config.TaskName
Wait-OrchestrationLaneHealth -Config $config | Out-Null
Get-OrchestrationLaneStatus -Config $config | ConvertTo-Json -Depth 8
