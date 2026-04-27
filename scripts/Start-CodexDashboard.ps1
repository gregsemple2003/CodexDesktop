[CmdletBinding()]
param()

$ErrorActionPreference = "Stop"
Set-StrictMode -Version Latest

function Get-FileSha256 {
    param([Parameter(Mandatory = $true)][string]$Path)
    $stream = [System.IO.File]::OpenRead($Path)
    try {
        $sha256 = [System.Security.Cryptography.SHA256]::Create()
        try {
            return -join ($sha256.ComputeHash($stream) | ForEach-Object { $_.ToString("x2") })
        }
        finally {
            $sha256.Dispose()
        }
    }
    finally {
        $stream.Dispose()
    }
}

$dashboardRoot = Split-Path -Parent $PSScriptRoot
$manifestPath = Join-Path $dashboardRoot "dashboard-current-release.json"
if (-not (Test-Path -LiteralPath $manifestPath)) {
    throw "No pinned dashboard release manifest exists at $manifestPath."
}

$manifest = Get-Content -Raw -LiteralPath $manifestPath | ConvertFrom-Json
if ([string]$manifest.component -ne "dashboard_frontend") {
    throw "Pinned dashboard manifest is for an unexpected component."
}

$releaseRoot = [string]$manifest.release_root
if ([string]::IsNullOrWhiteSpace($releaseRoot) -or -not (Test-Path -LiteralPath $releaseRoot)) {
    throw "Pinned dashboard release root does not exist: $releaseRoot."
}

foreach ($file in @($manifest.files)) {
    $path = Join-Path $releaseRoot ([string]$file.path)
    if (-not (Test-Path -LiteralPath $path)) {
        throw "Pinned dashboard release file is missing: $path."
    }
    if ((Get-FileSha256 -Path $path) -ne [string]$file.sha256) {
        throw "Pinned dashboard release file hash mismatch: $path."
    }
}

$pythonwPath = [string]$manifest.pythonw_path
if ([string]::IsNullOrWhiteSpace($pythonwPath) -or -not (Test-Path -LiteralPath $pythonwPath)) {
    throw "Pinned dashboard pythonw path does not exist: $pythonwPath."
}

$env:PYTHONPATH = $releaseRoot
$env:CODEX_DASHBOARD_RELEASE_ID = [string]$manifest.release_id
$env:CODEX_DASHBOARD_RELEASE_ROOT = $releaseRoot
Start-Process `
    -FilePath $pythonwPath `
    -WorkingDirectory $releaseRoot `
    -ArgumentList @(
        "-m",
        "app.codex_dashboard",
        "--release-id",
        [string]$manifest.release_id,
        "--release-root",
        $releaseRoot
    ) `
    -WindowStyle Hidden | Out-Null
