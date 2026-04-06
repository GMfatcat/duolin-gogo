param(
  [string]$BundleName = "duolin-gogo-windows-x64"
)

$ErrorActionPreference = "Stop"

$repoRoot = Split-Path -Parent $PSScriptRoot
$releaseRoot = Join-Path $repoRoot "release"
$bundleRoot = Join-Path $releaseRoot $BundleName
$zipPath = Join-Path $releaseRoot ($BundleName + ".zip")

$appExe = Join-Path $repoRoot "app\\build\\bin\\app.exe"
$knowledgeDir = Join-Path $repoRoot "knowledge"
$dataDir = Join-Path $repoRoot "data"

if (-not (Test-Path $appExe)) {
  throw "app.exe not found. Build the app first."
}

if (Test-Path $bundleRoot) {
  Remove-Item -LiteralPath $bundleRoot -Recurse -Force
}

if (Test-Path $zipPath) {
  Remove-Item -LiteralPath $zipPath -Force
}

New-Item -ItemType Directory -Path $bundleRoot | Out-Null

Copy-Item -LiteralPath $appExe -Destination (Join-Path $bundleRoot "app.exe") -Force
Copy-Item -LiteralPath $knowledgeDir -Destination (Join-Path $bundleRoot "knowledge") -Recurse -Force
Copy-Item -LiteralPath $dataDir -Destination (Join-Path $bundleRoot "data") -Recurse -Force

Compress-Archive -Path (Join-Path $bundleRoot "*") -DestinationPath $zipPath -Force

Write-Host "Release bundle ready:"
Write-Host $bundleRoot
Write-Host $zipPath
