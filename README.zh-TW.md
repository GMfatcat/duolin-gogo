# duolin-gogo

> 把自己的技術筆記變成一個本地優先的桌面微學習工具。

[![平台](https://img.shields.io/badge/platform-Windows-0f8f9c?style=flat-square)](./RELEASE.md)
[![技術棧](https://img.shields.io/badge/stack-Wails%20%7C%20Go%20%7C%20Vue%203-e9c46a?style=flat-square)](./app)
[![知識來源](https://img.shields.io/badge/knowledge-Markdown%20decks-1c3345?style=flat-square)](./knowledge)
[![語言](https://img.shields.io/badge/language-English%20%7C%20%E7%B9%81%E9%AB%94%E4%B8%AD%E6%96%87-0a5161?style=flat-square)](./README.md)

[English](./README.md) | 繁體中文

## 專案簡介

`duolin-gogo` 會把你的 Markdown 知識庫轉成一個本地桌面學習 app，包含：

- 定時學習提醒
- 以概念為核心的微學習卡片
- 間隔複習循環
- 依主題切換的學習模式
- AI 輔助的產卡工作流
- 具備隱藏成長和彩蛋的 DG 吉祥物陪伴

所有資料都留在本機：

- 原始知識內容放在 [`knowledge/`](./knowledge)
- 執行時狀態放在 [`data/`](./data)
- 桌面 app 程式放在 [`app/`](./app)

## 主要特色

- 🌏 雙語卡片（`zh-TW` / `en`）
- 🧠 自適應選題與複習流程
- 🔔 Windows toast 通知與背景常駐
- 🗂️ 多主題技術知識 deck
- ✍️ Library + AI 產卡工作區
- 🤖 DG 吉祥物陪伴、互動、成長與彩蛋
- 🚀 第一次啟動引導導覽
- ⚡ 使用 `cards-cache.gob` 提升啟動與重新掃描速度

## 開發時間線

- `MVP`：本地優先學習流程、複習循環、通知與背景常駐
- `Authoring`：Library / AI 產卡工作區、批次審查、修正建議、匯入報告
- `Expansion`：擴充為多主題技術 deck，涵蓋工具、後端與程式設計概念
- `DG`：加入吉祥物陪伴、隱藏成長、引導導覽與彩蛋

## 目前主題

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

## 技術棧

- Wails
- Go
- Vue 3
- 本地 JSON / JSONL runtime state
- `cards-cache.gob` 知識快取

## 專案結構

- [`knowledge/`](./knowledge)：知識卡來源
- [`data/`](./data)：本地執行狀態與快取
- [`app/`](./app)：Wails 桌面應用程式
- [`AI_CARD_PROMPT.md`](./AI_CARD_PROMPT.md)：AI 產卡 prompt
- [`MASCOT_SPEC.md`](./MASCOT_SPEC.md)：DG 吉祥物設計規格
- [`RELEASE.md`](./RELEASE.md)：發布前整理與打包說明

## 開發方式

Backend 測試：

```powershell
cd app
go test ./...
```

Frontend 測試：

```powershell
cd app/frontend
npm test -- --run
```

Frontend build：

```powershell
cd app/frontend
npm run build
```

桌面版 build：

```powershell
cd app
wails build
```

## 發布

先整理出乾淨的發布資料：

```powershell
powershell -ExecutionPolicy Bypass -File .\scripts\prepare-release.ps1
cd app
wails build
```

再建立可分發的 release bundle：

```powershell
powershell -ExecutionPolicy Bypass -File .\scripts\create-release-bundle.ps1
```

建議的發布結構：

```text
duolin-gogo/
  app.exe
  knowledge/
  data/
```

詳細流程請看 [`RELEASE.md`](./RELEASE.md)。
