# duolin-gogo

把你自己的技術筆記變成桌面微學習工具。

[English README](./README.md)

## 這是什麼

`duolin-gogo` 會把你的 Markdown 知識庫轉成：

- 定時學習提醒
- 先看概念、再答題的短卡流程
- 間隔複習節奏
- 可依主題切換的學習模式
- 帶有 DG 吉祥物陪伴的桌面學習體驗

整套系統以本地為主。知識內容放在 `knowledge/`，執行狀態放在 `data/`。

## 主要功能

- 雙語卡片（`zh-TW` / `en`）
- 以本地 Markdown 為核心的知識 deck
- 自適應選題與複習流程
- Windows 通知與系統匣背景執行
- 多主題技術學習內容
- AI 輔助產卡工作區
- DG 吉祥物陪伴、隱藏成長與彩蛋
- 第一次開啟時的引導導覽

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

## 技術組合

- Wails
- Go
- Vue 3
- 本地 JSON / JSONL runtime state
- `cards-cache.gob` 知識快取

## 專案結構

- `knowledge/`：知識卡來源
- `data/`：本地執行狀態與快取
- `app/`：Wails 桌面程式
- `AI_CARD_PROMPT.md`：AI 產卡提示詞
- `MASCOT_SPEC.md`：DG 吉祥物規格
- `RELEASE.md`：發布前整理與打包說明

## 開發

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

桌面 build：

```powershell
cd app
wails build
```

## 發布

先整理乾淨的發布版資料：

```powershell
powershell -ExecutionPolicy Bypass -File .\scripts\prepare-release.ps1
cd app
wails build
```

建議發布結構：

```text
duolin-gogo/
  app.exe
  knowledge/
  data/
```

詳細請看 [RELEASE.md](./RELEASE.md)。
