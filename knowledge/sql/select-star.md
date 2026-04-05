---
id: sql-select-star-overfetch
title_zh: SELECT *
title_en: SQL SELECT Star
type: true-false
body_format: bilingual-section
tags: [sql, select, performance]
difficulty: 2
question_zh: "`SELECT *` 在很多實際查詢裡不一定是好主意，因為它可能多拿不需要的欄位。"
question_en: "`SELECT *` is not always a good idea in real queries because it can fetch columns you do not need."
answer: true
clickbait_zh: "看起來省事的星號，有時其實是偷偷把多餘成本一起帶回來。"
clickbait_en: "That convenient star can quietly bring back extra cost with it."
review_hint_zh: "`SELECT *` 可能造成 overfetch。"
review_hint_en: "`SELECT *` can cause overfetch."
enabled: true
---

## zh-TW

`SELECT *` 很方便，但它可能讀出比你真正需要還多的欄位，增加傳輸與處理成本。
在正式查詢裡，明確列出需要的欄位通常更穩，也更容易維護。

## en

`SELECT *` is convenient, but it may fetch more columns than you really need, increasing transfer and processing cost.
In production queries, explicitly naming the needed columns is usually clearer and safer.
