# duolin-gogo

Local-first desktop microlearning for your own technical notes.

[繁體中文版](./README.zh-TW.md)

## What It Is

`duolin-gogo` turns your Markdown knowledge base into:

- timed study nudges
- short concept-first learning cards
- spaced review loops
- topic-aware study sessions
- a mascot-led desktop learning experience

Everything runs locally. Your knowledge lives in `knowledge/`, and your runtime state stays in `data/`.

## Main Features

- Bilingual cards (`zh-TW` / `en`)
- Local Markdown-based knowledge decks
- Adaptive card selection and review flow
- Windows toast notifications and tray background running
- Multi-topic technical study decks
- AI-assisted authoring workspace
- DG mascot companion with hidden growth and easter eggs
- First-launch onboarding tour

## Current Topics

- Git
- Docker
- Linux
- Go
- Python
- SQL
- HTTP / Web Basics
- Backend Engineering Basics
- Bash
- Middleware
- OOP
- Functional Programming
- Design Patterns

## Tech Stack

- Wails
- Go
- Vue 3
- Local JSON / JSONL runtime state
- `cards-cache.gob` knowledge cache

## Project Layout

- `knowledge/` - source knowledge cards
- `data/` - local runtime state and cache
- `app/` - Wails desktop app
- `AI_CARD_PROMPT.md` - AI card authoring prompt
- `MASCOT_SPEC.md` - DG mascot design spec
- `RELEASE.md` - release preparation and packaging notes

## Development

Backend tests:

```powershell
cd app
go test ./...
```

Frontend tests:

```powershell
cd app/frontend
npm test -- --run
```

Frontend production build:

```powershell
cd app/frontend
npm run build
```

Desktop build:

```powershell
cd app
wails build
```

## Release

Prepare a clean release build:

```powershell
powershell -ExecutionPolicy Bypass -File .\scripts\prepare-release.ps1
cd app
wails build
```

Recommended release bundle:

```text
duolin-gogo/
  app.exe
  knowledge/
  data/
```

See [RELEASE.md](./RELEASE.md) for details.
