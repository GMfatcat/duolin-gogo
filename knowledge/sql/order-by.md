---
id: sql-order-by-sort-rows
title_zh: SQL ORDER BY
title_en: SQL ORDER BY
type: true-false
body_format: bilingual-section
tags: [sql, order-by, sort]
difficulty: 1
question_zh: "`ORDER BY` 用來決定查詢結果的排序方式。"
question_en: "`ORDER BY` is used to decide how query results should be sorted."
answer: true
clickbait_zh: "你以為資料天生就有順序？很多時候沒寫排序，就只是剛好看起來像有。"
clickbait_en: "You may think the rows already have an order, but without sorting that order may just be accidental."
review_hint_zh: "`ORDER BY` 明確指定結果排序。"
review_hint_en: "`ORDER BY` explicitly controls result ordering."
enabled: true
---

## zh-TW

如果沒有 `ORDER BY`，資料庫不一定保證結果順序穩定。
`ORDER BY` 會明確指定你要依哪個欄位排序，以及升冪或降冪。

## en

Without `ORDER BY`, a database does not necessarily guarantee a stable row order.
`ORDER BY` explicitly tells the query which columns should control sorting and in which direction.
