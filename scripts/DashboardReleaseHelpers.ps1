Set-StrictMode -Version Latest

function Get-DashboardRepoRoot {
    return (Resolve-Path (Join-Path $PSScriptRoot "..")).Path
}

function Get-DashboardLocalAppData {
    if ($env:LOCALAPPDATA) {
        return $env:LOCALAPPDATA
    }
    return (Join-Path $HOME "AppData\\Local")
}

function Get-DashboardStartupPath {
    $appData = if ($env:APPDATA) { $env:APPDATA } else { Join-Path $HOME "AppData\\Roaming" }
    return Join-Path $appData "Microsoft\\Windows\\Start Menu\\Programs\\Startup\\CodexDashboard.cmd"
}

function Get-DashboardReleaseConfig {
    $repoRoot = Get-DashboardRepoRoot
    $dashboardRoot = Join-Path (Get-DashboardLocalAppData) "CodexDashboard"
    $launcherRoot = Join-Path $dashboardRoot "dashboard-launcher"
    return @{
        RepoRoot = $repoRoot
        SourceAppRoot = Join-Path $repoRoot "app"
        DashboardRoot = $dashboardRoot
        ReleasesRoot = Join-Path $dashboardRoot "dashboard-releases"
        CurrentReleaseManifestPath = Join-Path $dashboardRoot "dashboard-current-release.json"
        LauncherRoot = $launcherRoot
        LauncherScriptPath = Join-Path $launcherRoot "Start-CodexDashboard.ps1"
        LauncherConfigPath = Join-Path $launcherRoot "launcher-config.json"
        SourceLauncherScriptPath = Join-Path $PSScriptRoot "Start-CodexDashboard.ps1"
        StartupPath = Get-DashboardStartupPath
    }
}

function Get-DashboardFileSha256 {
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

function Get-DashboardRelativePath {
    param(
        [Parameter(Mandatory = $true)]
        [string]$Root,
        [Parameter(Mandatory = $true)]
        [string]$Path
    )

    $rootFull = [System.IO.Path]::GetFullPath($Root).TrimEnd([char[]]@("\", "/"))
    $pathFull = [System.IO.Path]::GetFullPath($Path)
    if ($pathFull.Equals($rootFull, [System.StringComparison]::OrdinalIgnoreCase)) {
        return ""
    }
    $rootPrefix = $rootFull + [System.IO.Path]::DirectorySeparatorChar
    if (-not $pathFull.StartsWith($rootPrefix, [System.StringComparison]::OrdinalIgnoreCase)) {
        throw "Path '$Path' is not under root '$Root'."
    }
    return $pathFull.Substring($rootPrefix.Length)
}

function Get-DashboardGitExecutablePath {
    foreach ($name in @("git.exe", "git")) {
        $command = Get-Command $name -ErrorAction SilentlyContinue
        if ($null -ne $command -and -not [string]::IsNullOrWhiteSpace($command.Source)) {
            return $command.Source
        }
    }
    throw "Could not resolve git."
}

function Get-DashboardGitCommit {
    param(
        [Parameter(Mandatory = $true)]
        [string]$RepoRoot
    )

    $gitPath = Get-DashboardGitExecutablePath
    $commit = (& $gitPath -C $RepoRoot rev-parse HEAD)
    if ($LASTEXITCODE -ne 0 -or [string]::IsNullOrWhiteSpace($commit)) {
        throw "Could not resolve git commit for $RepoRoot."
    }
    return [string]$commit
}

function Get-DashboardGitStatusShort {
    param(
        [Parameter(Mandatory = $true)]
        [string]$RepoRoot
    )

    $gitPath = Get-DashboardGitExecutablePath
    $status = @(& $gitPath -C $RepoRoot status --short)
    if ($LASTEXITCODE -ne 0) {
        throw "Could not resolve git status for $RepoRoot."
    }
    return $status
}

function Get-DashboardPythonwPath {
    $candidates = @()
    if (-not [string]::IsNullOrWhiteSpace($env:ProgramFiles)) {
        $candidates += (Join-Path $env:ProgramFiles "Python313\\pythonw.exe")
    }
    foreach ($name in @("pythonw.exe", "pythonw")) {
        $command = Get-Command $name -ErrorAction SilentlyContinue
        if ($null -ne $command -and -not [string]::IsNullOrWhiteSpace($command.Source)) {
            $candidates += $command.Source
        }
    }
    foreach ($candidate in $candidates) {
        if ([string]::IsNullOrWhiteSpace($candidate)) {
            continue
        }
        if ($candidate -like "*\\WindowsApps\\*") {
            continue
        }
        if (Test-Path -LiteralPath $candidate) {
            return $candidate
        }
    }
    throw "Could not resolve a stable pythonw executable."
}

function Get-DashboardPowerShellPath {
    $candidate = Join-Path $env:SystemRoot "System32\\WindowsPowerShell\\v1.0\\powershell.exe"
    if (Test-Path -LiteralPath $candidate) {
        return $candidate
    }
    $command = Get-Command "powershell.exe" -ErrorAction SilentlyContinue
    if ($null -ne $command -and -not [string]::IsNullOrWhiteSpace($command.Source)) {
        return $command.Source
    }
    throw "Could not resolve powershell.exe."
}

function Copy-DashboardReleaseTree {
    param(
        [Parameter(Mandatory = $true)]
        [string]$SourceRoot,
        [Parameter(Mandatory = $true)]
        [string]$DestinationRoot
    )

    $sourceRootFull = [System.IO.Path]::GetFullPath($SourceRoot)
    foreach ($item in Get-ChildItem -LiteralPath $sourceRootFull -Recurse -Force) {
        $relative = Get-DashboardRelativePath -Root $sourceRootFull -Path $item.FullName
        if ($relative -split "[\\/]" | Where-Object { $_ -eq "__pycache__" }) {
            continue
        }
        if (-not $item.PSIsContainer -and $item.Name -like "*.pyc") {
            continue
        }
        $destination = Join-Path $DestinationRoot $relative
        if ($item.PSIsContainer) {
            New-Item -ItemType Directory -Path $destination -Force | Out-Null
        }
        else {
            $parent = Split-Path -Parent $destination
            if (-not (Test-Path -LiteralPath $parent)) {
                New-Item -ItemType Directory -Path $parent -Force | Out-Null
            }
            Copy-Item -LiteralPath $item.FullName -Destination $destination -Force
        }
    }
}

function Copy-DashboardGitTree {
    param(
        [Parameter(Mandatory = $true)]
        [string]$RepoRoot,
        [Parameter(Mandatory = $true)]
        [string]$Commit,
        [Parameter(Mandatory = $true)]
        [string]$DestinationRoot
    )

    $gitPath = Get-DashboardGitExecutablePath
    $tempRoot = Join-Path ([System.IO.Path]::GetTempPath()) ("codex-dashboard-release-" + [guid]::NewGuid().ToString("n"))
    $archivePath = Join-Path $tempRoot "app.zip"
    New-Item -ItemType Directory -Path $tempRoot -Force | Out-Null
    try {
        & $gitPath -C $RepoRoot archive --format=zip -o $archivePath $Commit app
        if ($LASTEXITCODE -ne 0) {
            throw "git archive failed with exit code $LASTEXITCODE."
        }
        Expand-Archive -LiteralPath $archivePath -DestinationPath $DestinationRoot -Force
    }
    finally {
        Remove-Item -LiteralPath $tempRoot -Recurse -Force -ErrorAction SilentlyContinue
    }
}

function Get-DashboardReleaseFileManifest {
    param(
        [Parameter(Mandatory = $true)]
        [string]$ReleaseRoot
    )

    $releaseRootFull = [System.IO.Path]::GetFullPath($ReleaseRoot)
    $files = @()
    foreach ($file in Get-ChildItem -LiteralPath (Join-Path $releaseRootFull "app") -Recurse -File | Sort-Object FullName) {
        $relative = (Get-DashboardRelativePath -Root $releaseRootFull -Path $file.FullName).Replace("\", "/")
        $files += [ordered]@{
            path = $relative
            sha256 = Get-DashboardFileSha256 -Path $file.FullName
            bytes = $file.Length
        }
    }
    return $files
}

function Install-DashboardLauncher {
    param(
        [Parameter(Mandatory = $true)]
        [hashtable]$Config
    )

    New-Item -ItemType Directory -Path $Config.LauncherRoot -Force | Out-Null
    Copy-Item -LiteralPath $Config.SourceLauncherScriptPath -Destination $Config.LauncherScriptPath -Force
    [ordered]@{
        schema_version = 1
        source_repo_root = $Config.RepoRoot
        installed_at = (Get-Date).ToUniversalTime().ToString("o")
        launcher_script_path = $Config.LauncherScriptPath
    } | ConvertTo-Json -Depth 6 | Set-Content -LiteralPath $Config.LauncherConfigPath -Encoding UTF8
}

function Set-DashboardCurrentRelease {
    param(
        [Parameter(Mandatory = $true)]
        [hashtable]$Config,
        [Parameter(Mandatory = $true)]
        [string]$ManifestPath
    )

    if (-not (Test-Path -LiteralPath $ManifestPath)) {
        throw "Dashboard release manifest does not exist: $ManifestPath."
    }
    $manifestJson = Get-Content -Raw -LiteralPath $ManifestPath
    $manifest = $manifestJson | ConvertFrom-Json
    if ([string]$manifest.component -ne "dashboard_frontend") {
        throw "Dashboard release manifest is for an unexpected component."
    }
    $tempPath = "$($Config.CurrentReleaseManifestPath).tmp"
    Set-Content -LiteralPath $tempPath -Value $manifestJson -Encoding UTF8
    Move-Item -LiteralPath $tempPath -Destination $Config.CurrentReleaseManifestPath -Force
}

function Test-DashboardReleaseManifest {
    param(
        [Parameter(Mandatory = $true)]
        [string]$ManifestPath
    )

    if (-not (Test-Path -LiteralPath $ManifestPath)) {
        throw "No pinned dashboard release manifest exists at $ManifestPath."
    }
    $manifest = Get-Content -Raw -LiteralPath $ManifestPath | ConvertFrom-Json
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
        $actualHash = Get-DashboardFileSha256 -Path $path
        if ($actualHash -ne [string]$file.sha256) {
            throw "Pinned dashboard release file hash mismatch: $path."
        }
    }
    return $manifest
}

function New-DashboardRelease {
    param(
        [Parameter(Mandatory = $true)]
        [hashtable]$Config,
        [switch]$AllowDirty,
        [switch]$PinCurrent,
        [switch]$FromWorkingTree
    )

    if (-not (Test-Path -LiteralPath $Config.SourceAppRoot)) {
        throw "Dashboard source app root does not exist: $($Config.SourceAppRoot)."
    }
    $commit = Get-DashboardGitCommit -RepoRoot $Config.RepoRoot
    $statusLines = @(Get-DashboardGitStatusShort -RepoRoot $Config.RepoRoot)
    if ($FromWorkingTree -and $statusLines.Count -gt 0 -and -not $AllowDirty) {
        throw "Refusing to publish a dashboard release from a dirty repo. Commit/stash changes or pass -AllowDirty to record the dirty status in the manifest."
    }
    $createdAt = (Get-Date).ToUniversalTime()
    $shortCommit = $commit.Substring(0, [Math]::Min(12, $commit.Length))
    $releaseId = "{0}-{1}" -f $createdAt.ToString("yyyyMMddTHHmmssZ"), $shortCommit
    $releaseRoot = Join-Path $Config.ReleasesRoot $releaseId
    $releaseAppRoot = Join-Path $releaseRoot "app"
    $manifestPath = Join-Path $releaseRoot "dashboard-release-manifest.json"

    New-Item -ItemType Directory -Path $releaseRoot -Force | Out-Null
    if ($FromWorkingTree) {
        Copy-DashboardReleaseTree -SourceRoot $Config.SourceAppRoot -DestinationRoot $releaseAppRoot
    }
    else {
        Copy-DashboardGitTree -RepoRoot $Config.RepoRoot -Commit $commit -DestinationRoot $releaseRoot
    }
    $pythonwPath = Get-DashboardPythonwPath

    $manifest = [ordered]@{
        schema_version = 1
        component = "dashboard_frontend"
        release_id = $releaseId
        release_root = $releaseRoot
        created_at = $createdAt.ToString("o")
        source_repo_root = $Config.RepoRoot
        git_commit = $commit
        source_mode = if ($FromWorkingTree) { "working_tree" } else { "git_commit" }
        source_dirty = ($FromWorkingTree -and $statusLines.Count -gt 0)
        repository_dirty = ($statusLines.Count -gt 0)
        source_status = $statusLines
        pythonw_path = $pythonwPath
        launcher_script_path = $Config.LauncherScriptPath
        startup_path = $Config.StartupPath
        files = @(Get-DashboardReleaseFileManifest -ReleaseRoot $releaseRoot)
    }
    $manifest | ConvertTo-Json -Depth 12 | Set-Content -LiteralPath $manifestPath -Encoding UTF8
    if ($PinCurrent) {
        Set-DashboardCurrentRelease -Config $Config -ManifestPath $manifestPath
    }
    return (Get-Content -Raw -LiteralPath $manifestPath | ConvertFrom-Json)
}

function Install-DashboardStartup {
    param(
        [Parameter(Mandatory = $true)]
        [hashtable]$Config
    )

    $powershellPath = Get-DashboardPowerShellPath
    $startupParent = Split-Path -Parent $Config.StartupPath
    if (-not (Test-Path -LiteralPath $startupParent)) {
        New-Item -ItemType Directory -Path $startupParent -Force | Out-Null
    }
    $content = @(
        "@echo off",
        "`"$powershellPath`" -NoProfile -ExecutionPolicy Bypass -WindowStyle Hidden -File `"$($Config.LauncherScriptPath)`""
    ) -join "`r`n"
    Set-Content -LiteralPath $Config.StartupPath -Value ($content + "`r`n") -Encoding ASCII
}
