---
id: sql-limit-top-n
title_zh: SQL LIMIT
title_en: SQL LIMIT
type: true-false
body_format: bilingual-section
tags: [sql, limit, pagination]
difficulty: 1
question_zh: "`LIMIT` 常用來只取前幾筆結果，而不是整張表都讀回來。"
question_en: "`LIMIT` is commonly used to return only the first few rows instead of the whole table."
answer: true
clickbait_zh: "你未必要把整桶資料搬回來，有時候只看前幾筆就夠了。"
clickbait_en: "You do not always need to bring back the whole bucket of data. Sometimes the first few rows are enough."
review_hint_zh: "`LIMIT` 會限制回傳筆數。"
review_hint_en: "`LIMIT` constrains how many rows are returned."
enabled: true
---

## zh-TW

`LIMIT` 用來限制查詢回傳的列數，常見於預覽、列表頁或分頁情境。
如果你想要穩定的前幾筆，通常也應該搭配 `ORDER BY` 一起用。

## en

`LIMIT` restricts the number of rows a query returns and is common in previews, lists, and pagination.
If you want a stable "top N" result, it should usually be combined with `ORDER BY`.
