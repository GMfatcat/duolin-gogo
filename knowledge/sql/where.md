---
id: sql-where-filter-rows
title_zh: SQL WHERE
title_en: SQL WHERE
type: true-false
body_format: bilingual-section
tags: [sql, where, filter]
difficulty: 1
question_zh: "`WHERE` 常用來在查詢結果中先過濾列資料。"
question_en: "`WHERE` is commonly used to filter rows before the final result is returned."
answer: true
clickbait_zh: "不是查出來再慢慢挑，SQL 其實可以一開始就把不要的資料擋在外面。"
clickbait_en: "You do not have to fetch everything and filter later. SQL can reject rows up front."
review_hint_zh: "`WHERE` 先篩列，再決定後續查詢流程。"
review_hint_en: "`WHERE` filters rows before later query steps."
enabled: true
---

## zh-TW

`WHERE` 會在查詢流程前段先篩掉不符合條件的列。
所以它通常用來縮小資料範圍，而不是在結果最後才補過濾。

## en

`WHERE` filters out rows early in the query flow before the final result is produced.
That is why it is commonly used to narrow the dataset before later steps happen.
