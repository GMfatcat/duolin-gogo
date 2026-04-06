# duolin-gogo

> Local-first desktop microlearning for your own technical notes.

[![Platform](https://img.shields.io/badge/platform-Windows-0f8f9c?style=flat-square)](./RELEASE.md)
[![Stack](https://img.shields.io/badge/stack-Wails%20%7C%20Go%20%7C%20Vue%203-e9c46a?style=flat-square)](./app)
[![Knowledge](https://img.shields.io/badge/knowledge-Markdown%20decks-1c3345?style=flat-square)](./knowledge)
[![Language](https://img.shields.io/badge/language-English%20%7C%20%E7%B9%81%E9%AB%94%E4%B8%AD%E6%96%87-0a5161?style=flat-square)](./README.zh-TW.md)

English | [繁體中文](./README.zh-TW.md)

## Overview

`duolin-gogo` turns your Markdown knowledge base into a local desktop study app with:

- timed study nudges
- concept-first microlearning cards
- spaced review loops
- topic-aware study sessions
- AI-assisted authoring tools
- a DG mascot companion with hidden growth and easter eggs

Everything stays local:

- source knowledge lives in [`knowledge/`](./knowledge)
- runtime state lives in [`data/`](./data)
- the desktop app lives in [`app/`](./app)

## Highlights

- 🌏 Bilingual cards (`zh-TW` / `en`)
- 🧠 Adaptive card selection and review flow
- 🔔 Windows toast notifications and tray background running
- 🗂️ Multi-topic technical study decks
- ✍️ Library + AI authoring workspace
- 🤖 DG mascot with reactions, growth, and easter eggs
- 🚀 First-launch onboarding tour
- ⚡ `cards-cache.gob` knowledge cache for faster startup and rescans

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
- `cards-cache.gob` for parsed knowledge caching

## Repository Layout

- [`knowledge/`](./knowledge): source knowledge cards
- [`data/`](./data): local runtime state and cache
- [`app/`](./app): Wails desktop app
- [`AI_CARD_PROMPT.md`](./AI_CARD_PROMPT.md): AI card authoring prompt
- [`MASCOT_SPEC.md`](./MASCOT_SPEC.md): DG mascot design spec
- [`RELEASE.md`](./RELEASE.md): release preparation and packaging notes

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

Create a distributable release bundle:

```powershell
powershell -ExecutionPolicy Bypass -File .\scripts\create-release-bundle.ps1
```

Recommended release structure:

```text
duolin-gogo/
  app.exe
  knowledge/
  data/
```

See [`RELEASE.md`](./RELEASE.md) for details.
