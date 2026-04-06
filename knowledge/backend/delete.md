---
id: backend-delete-remove-resource
title_zh: Delete
title_en: Delete
type: true-false
body_format: bilingual-section
tags: [backend, crud, delete]
difficulty: 1
question_zh: "在 CRUD 裡，Delete 指的是刪除或移除既有資源。"
question_en: "In CRUD, Delete refers to removing an existing resource."
answer: true
clickbait_zh: "真正可怕的不是刪不掉，而是你以為刪了一筆結果刪了一片。"
clickbait_en: "The scary part is not failing to delete. It is thinking you deleted one thing and removing far more."
review_hint_zh: "Delete 的重點是移除既有資料。"
review_hint_en: "Delete is about removing existing data."
enabled: true
---

## zh-TW

Delete 可以是硬刪除，也可以是邏輯上的隱藏或封存。
實務上常要考慮資料關聯、審計需求、誤刪風險與是否需要 soft delete。

## en

Delete can mean a hard removal or a logical hide/archive operation.
In practice, you often need to think about relationships, audit requirements, accidental deletion risk, and soft delete strategies.
