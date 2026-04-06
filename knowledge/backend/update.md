---
id: backend-update-change-resource
title_zh: Update
title_en: Update
type: true-false
body_format: bilingual-section
tags: [backend, crud, update]
difficulty: 1
question_zh: "在 CRUD 裡，Update 指的是修改既有資料，而不是建立新資料。"
question_en: "In CRUD, Update means modifying existing data rather than creating new data."
answer: true
clickbait_zh: "最危險的更新不是改太少，而是你以為自己只改一筆。"
clickbait_en: "The most dangerous update is not the one that changes too little, but the one you thought would touch only one row."
review_hint_zh: "Update 的重點是改既有資料。"
review_hint_en: "Update is about changing something that already exists."
enabled: true
---

## zh-TW

Update 代表對既有資源做修改，例如更新使用者暱稱、訂單狀態或某筆設定值。
這裡常牽涉並發、欄位部分更新、驗證與是否可重複執行等問題。

## en

Update means changing an existing resource, such as a username, order status, or config value.
This often involves concurrency, partial updates, validation, and idempotency concerns.
