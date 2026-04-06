# duolin-gogo

`duolin-gogo` is a local-first desktop study app that turns your own Markdown knowledge base into timed nudges, micro-lessons, and review loops.

`duolin-gogo` 是一個以本地為主的桌面學習工具，可以把你自己的 Markdown 知識庫變成定時提醒、微課與複習節奏。

## Features / 功能特色

- Bilingual knowledge cards (`zh-TW` / `en`)
- Local Markdown knowledge base under `knowledge/`
- Adaptive card selection with review mode
- Windows toast notifications and tray background running
- Multi-topic study decks
- AI-assisted authoring flow inside the app
- DG mascot companion with hidden growth and easter eggs

- 雙語知識卡（`zh-TW` / `en`）
- 以 `knowledge/` 為核心的本地 Markdown 題庫
- 自適應選題與複習模式
- Windows 通知與系統匣背景執行
- 多主題學習 deck
- 內建 AI 輔助產卡流程
- 內建 DG 吉祥物陪伴與隱藏成長系統

## Current Topic Coverage / 目前主題

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

## Tech Stack / 技術組合

- Wails
- Go
- Vue 3
- Local JSON / JSONL runtime state
- `cards-cache.gob` knowledge cache

## Project Structure / 專案結構

- [D:\duolin-gogo\knowledge](D:\duolin-gogo\knowledge): source knowledge cards
- [D:\duolin-gogo\data](D:\duolin-gogo\data): local runtime state and cache
- [D:\duolin-gogo\app](D:\duolin-gogo\app): Wails desktop app
- [D:\duolin-gogo\AI_CARD_PROMPT.md](D:\duolin-gogo\AI_CARD_PROMPT.md): AI card authoring prompt
- [D:\duolin-gogo\MASCOT_SPEC.md](D:\duolin-gogo\MASCOT_SPEC.md): DG mascot design spec

## Development / 開發

Backend tests:

```powershell
cd D:\duolin-gogo\app
go test ./...
```

Frontend tests:

```powershell
cd D:\duolin-gogo\app\frontend
npm test -- --run
```

Frontend build:

```powershell
cd D:\duolin-gogo\app\frontend
npm run build
```

Desktop build:

```powershell
cd D:\duolin-gogo\app
wails build
```

## Release / 發布

Before packaging a release build:

1. run the release-prep script to reset local study state
2. build the app
3. ship `app.exe` together with the `knowledge/` and `data/` folders

發布前建議：

1. 先執行 release-prep 腳本，清空本地學習紀錄
2. 再 build 桌面程式
3. 發布時要把 `app.exe`、`knowledge/`、`data/` 一起帶上

Detailed release notes:

- [D:\duolin-gogo\RELEASE.md](D:\duolin-gogo\RELEASE.md)

