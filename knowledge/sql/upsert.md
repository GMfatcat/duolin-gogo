---
id: sql-upsert-insert-or-update
title_zh: Upsert
title_en: Upsert
type: true-false
body_format: bilingual-section
tags: [sql, write, upsert]
difficulty: 2
question_zh: "upsert 的概念通常是『有資料就更新，沒有就新增』。"
question_en: "Upsert usually means update the row if it exists, otherwise insert it."
answer: true
clickbait_zh: "你以為要先查再決定 insert 還是 update？其實很多資料庫早就幫你合成一招。"
clickbait_en: "You might think you need a read-before-write branch, but many databases already combine insert and update into one move."
review_hint_zh: "upsert = insert or update。"
review_hint_en: "Upsert means insert-or-update."
enabled: true
---

## zh-TW

upsert 常見在唯一鍵衝突時改成更新既有資料，而不是直接報錯。
這種寫法能減少先查再寫的分支判斷，但仍要理解唯一鍵與衝突條件。

## en

Upsert usually updates an existing row when a unique-key conflict happens instead of failing immediately.
It reduces read-before-write branching, but you still need to understand the unique key and conflict condition.
