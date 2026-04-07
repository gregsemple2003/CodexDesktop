[CmdletBinding()]
param()

$ErrorActionPreference = "Stop"
Set-StrictMode -Version Latest

. (Join-Path $PSScriptRoot "LaneHelpers.ps1")

$config = Get-OrchestrationLaneConfig -Lane "service"

Stop-ScheduledTask -TaskName $config.TaskName -ErrorAction SilentlyContinue | Out-Null
Stop-OrchestrationLaneProcesses -Config $config
Invoke-OrchestrationCompose -Config $config -ComposeArgs @("down") -AllowFailure | Out-Null
Get-OrchestrationLaneStatus -Config $config | ConvertTo-Json -Depth 8
