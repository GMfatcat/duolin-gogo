---
id: sql-full-table-scan-cost
title_zh: Full Table Scan
title_en: SQL Full Table Scan
type: true-false
body_format: bilingual-section
tags: [sql, performance, scan]
difficulty: 3
question_zh: "Full table scan 的意思，通常是資料庫需要掃過整張表來找符合條件的列。"
question_en: "A full table scan usually means the database has to scan the entire table to find matching rows."
answer: true
clickbait_zh: "有些慢查詢不是因為算很難，而是它根本把整張表都翻過一次。"
clickbait_en: "Some slow queries are not slow because they are smart. They are slow because they read everything."
review_hint_zh: "full table scan = 幾乎整張表都要看過。"
review_hint_en: "A full table scan means most or all rows are inspected."
enabled: true
---

## zh-TW

當資料庫找不到更有效率的方式時，就可能對整張表做掃描，逐列檢查哪些資料符合條件。
在大表上，這通常代表較高成本，所以常是 SQL 優化時要特別注意的訊號。

## en

When the database cannot find a more efficient path, it may scan the whole table and inspect rows one by one.
On large tables, that usually means higher cost, so it is a common signal to watch during SQL tuning.
