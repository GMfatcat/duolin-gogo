---
id: sql-indexed-where-fast-filter
title_zh: 索引欄位上的 WHERE
title_en: Indexed WHERE Filter
type: true-false
body_format: bilingual-section
tags: [sql, index, where, performance]
difficulty: 2
question_zh: "在常見情況下，`WHERE` 條件若能有效利用索引，通常能減少查詢成本。"
question_en: "In common cases, a `WHERE` condition that can use an index usually reduces query cost."
answer: true
clickbait_zh: "有些查詢快不是因為資料變少了，而是因為資料庫知道該先去哪裡找。"
clickbait_en: "Some queries are fast not because there is less data, but because the database knows where to look first."
review_hint_zh: "索引能幫 `WHERE` 更快定位列。"
review_hint_en: "Indexes can help `WHERE` locate rows faster."
enabled: true
---

## zh-TW

如果 `WHERE` 條件正好用在適合的索引欄位上，資料庫通常不必盲目掃整張表。
這也是 SQL 效能優化時，為什麼會特別關心查詢條件和索引是否對得上的原因。

## en

If a `WHERE` condition aligns with a useful indexed column, the database often does not have to scan the entire table blindly.
That is one reason SQL tuning pays close attention to whether query filters and indexes line up well.
