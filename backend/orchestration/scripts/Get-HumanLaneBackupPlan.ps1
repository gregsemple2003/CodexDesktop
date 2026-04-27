[CmdletBinding()]
param(
    [switch]$SkipDockerVolumeSize
)

$ErrorActionPreference = "Stop"
Set-StrictMode -Version Latest

. (Join-Path $PSScriptRoot "HumanLaneBackupHelpers.ps1")

Get-HumanLaneBackupPlan -SkipDockerVolumeSize:$SkipDockerVolumeSize | ConvertTo-Json -Depth 12
