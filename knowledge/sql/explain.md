---
id: sql-explain-plan-inspect
title_zh: SQL EXPLAIN
title_en: SQL EXPLAIN
type: true-false
body_format: bilingual-section
tags: [sql, explain, performance]
difficulty: 3
question_zh: "`EXPLAIN` 常用來查看資料庫打算如何執行一個查詢。"
question_en: "`EXPLAIN` is commonly used to inspect how the database plans to execute a query."
answer: true
clickbait_zh: "不要只問『這條 SQL 為什麼慢』，先看資料庫到底打算怎麼跑。"
clickbait_en: "Before asking why the query is slow, look at how the database plans to run it."
review_hint_zh: "`EXPLAIN` 用來看查詢計畫。"
review_hint_en: "`EXPLAIN` shows the query plan."
enabled: true
---

## zh-TW

`EXPLAIN` 通常用來查看資料庫的查詢計畫，例如會不會用到 index、是否可能全表掃描，或 join 順序如何。
它是 SQL 效能分析裡很重要的入口。

## en

`EXPLAIN` is commonly used to inspect the query plan, such as whether an index will be used, whether a full table scan may happen, or how joins are ordered.
It is one of the most important starting points for SQL performance analysis.
