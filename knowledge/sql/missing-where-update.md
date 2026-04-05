---
id: sql-missing-where-update-risk
title_zh: 少了 WHERE 的更新風險
title_en: Missing WHERE Update Risk
type: true-false
body_format: bilingual-section
tags: [sql, update, bug]
difficulty: 2
question_zh: "在 `UPDATE` 或 `DELETE` 中少了 `WHERE`，常見風險之一是整張表都被改掉。"
question_en: "One common risk of missing `WHERE` in `UPDATE` or `DELETE` is that the whole table gets changed."
answer: true
clickbait_zh: "SQL 最可怕的錯有時不是難，而是少打一小段就把整張表一起帶走。"
clickbait_en: "Some of the most dangerous SQL mistakes are not complex at all. They happen when one small clause is missing."
review_hint_zh: "寫入型 SQL 少了 `WHERE` 可能影響整張表。"
review_hint_en: "Write queries without `WHERE` can affect the entire table."
enabled: true
---

## zh-TW

在 `UPDATE` 或 `DELETE` 這類寫入型 SQL 裡，`WHERE` 常用來指定要影響哪些列。
如果忘了加，資料庫通常會把條件視為「全部都符合」，所以整張表都可能被改動。

## en

In write operations like `UPDATE` or `DELETE`, `WHERE` is commonly used to specify which rows should be affected.
If it is omitted, the database will often treat the operation as applying to every row in the table.
