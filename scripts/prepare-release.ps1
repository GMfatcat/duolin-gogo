param(
  [switch]$NoBackup
)

$ErrorActionPreference = "Stop"

$repoRoot = Split-Path -Parent $PSScriptRoot
$dataDir = Join-Path $repoRoot "data"
$backupRoot = Join-Path $repoRoot "data-release-backup"

function Ensure-Dir($path) {
  if (-not (Test-Path $path)) {
    New-Item -ItemType Directory -Path $path | Out-Null
  }
}

function Write-JsonFile($path, $object) {
  $json = $object | ConvertTo-Json -Depth 10
  Set-Content -LiteralPath $path -Value $json -Encoding UTF8
}

Ensure-Dir $dataDir

$filesToBackup = @(
  "settings.json",
  "progress.json",
  "attempts.jsonl",
  "import-errors.json",
  "pet.json"
)

if (-not $NoBackup) {
  Ensure-Dir $backupRoot
  $stamp = Get-Date -Format "yyyyMMdd-HHmmss"
  $targetBackup = Join-Path $backupRoot $stamp
  Ensure-Dir $targetBackup

  foreach ($name in $filesToBackup) {
    $source = Join-Path $dataDir $name
    if (Test-Path $source) {
      Copy-Item -LiteralPath $source -Destination (Join-Path $targetBackup $name) -Force
    }
  }
}

$cleanSettings = @{
  version = 1
  knowledge_directories = @()
  selected_topic = "all"
  notification_interval_minutes = 20
  active_hours = @{
    enabled = $true
    start = "09:00"
    end = "23:59"
  }
  review_schedule = @{
    mode = "daily"
    weekday = $null
    time = "21:00"
    batch_size = 8
  }
  language = @{
    default = "zh-TW"
    allow_toggle = $true
  }
  notifications = @{
    style = "chaotic"
    title_mode = "prefer_generated"
  }
  study_rules = @{
    max_new_cards_per_day = 12
    snooze_minutes = 15
    cooldown_after_answer_minutes = 20
    reveal_speed = "normal"
  }
  onboarding = @{
    seen = $false
  }
}

$cleanProgress = @{
  version = 1
  updated_at = $null
  cards = @{}
  daily_summary = @{}
}

$cleanImportErrors = @{
  version = 1
  generated_at = (Get-Date).ToString("yyyy-MM-ddTHH:mm:sszzz")
  errors = @()
}

$cleanPet = @{
  bond_xp = 0
  stage = 0
  reaction_step = 0
  topic_streak_count = 0
  rapid_click_count = 0
}

Write-JsonFile (Join-Path $dataDir "settings.json") $cleanSettings
Write-JsonFile (Join-Path $dataDir "progress.json") $cleanProgress
Write-JsonFile (Join-Path $dataDir "import-errors.json") $cleanImportErrors
Write-JsonFile (Join-Path $dataDir "pet.json") $cleanPet
Set-Content -LiteralPath (Join-Path $dataDir "attempts.jsonl") -Value "" -Encoding UTF8

Write-Host "Release data prepared."
