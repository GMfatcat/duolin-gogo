# Release Notes

## Goal

Prepare a clean distributable build of `duolin-gogo` without shipping local personal study history.

## What Should Be Cleaned

These runtime files should be reset before packaging:

- `data/settings.json`
- `data/progress.json`
- `data/attempts.jsonl`
- `data/import-errors.json`
- `data/pet.json`

These should stay:

- `knowledge/`
- `data/cards-cache.gob`

`cards-cache.gob` is kept so the app can start faster. If it becomes invalid on another machine, the app will rebuild it automatically from `knowledge/`.

## Release Prep Script

Run:

```powershell
cd D:\duolin-gogo
powershell -ExecutionPolicy Bypass -File .\scripts\prepare-release.ps1
```

What it does:

- backs up your current runtime files into `data-release-backup/<timestamp>/`
- resets study progress and attempts
- resets onboarding to unseen
- resets DG hidden pet state
- keeps the knowledge cache gob file

## Build

```powershell
cd D:\duolin-gogo\app
wails build
```

Output:

- [app.exe](D:\duolin-gogo\app\build\bin\app.exe)

## Recommended Release Folder

Recommended distribution layout:

```text
duolin-gogo/
  app.exe
  knowledge/
  data/
```

`knowledge/` and `data/` should live next to the exe in the release bundle so the app can discover them automatically.

## GitHub Release Suggestion

If you publish a GitHub release, ship:

- `app.exe`
- `knowledge/`
- `data/`

Do not ship:

- `data-release-backup/`
- `app/frontend/src/assets/dg-archive/`

