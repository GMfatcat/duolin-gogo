---
id: sql-join-match-rows
title_zh: SQL JOIN
title_en: SQL JOIN
type: true-false
body_format: bilingual-section
tags: [sql, join, relation]
difficulty: 2
question_zh: "`JOIN` 常用來把兩張表依相關欄位配對成同一份查詢結果。"
question_en: "`JOIN` is commonly used to match rows from two tables into one query result."
answer: true
clickbait_zh: "真正麻煩的不是表太多，是你忘了它們其實本來就彼此有關。"
clickbait_en: "The hard part is not that there are many tables. It is remembering that they are related."
review_hint_zh: "`JOIN` 會依關聯欄位把資料配對起來。"
review_hint_en: "`JOIN` matches rows through related columns."
enabled: true
---

## zh-TW

`JOIN` 的目的，是把不同資料表中彼此相關的列配對後放進同一份結果。
這在正規化資料庫裡很常見，因為資料通常不會全部塞在同一張表。

## en

The purpose of `JOIN` is to combine related rows from different tables into one result.
That is common in normalized databases because the data is usually split across multiple tables.
