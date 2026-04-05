---
id: sql-group-by-bucket
title_zh: SQL GROUP BY
title_en: SQL GROUP BY
type: true-false
body_format: bilingual-section
tags: [sql, group-by, aggregate]
difficulty: 2
question_zh: "`GROUP BY` 主要用來把列依某個欄位分組，方便做聚合統計。"
question_en: "`GROUP BY` is mainly used to bucket rows by a column so aggregates can be computed."
answer: true
clickbait_zh: "不是每次都要看單筆資料，有時候真正有用的是把它們先分群。"
clickbait_en: "Sometimes the useful view is not the individual row. It is the grouped picture."
review_hint_zh: "`GROUP BY` 是聚合前的分桶步驟。"
review_hint_en: "`GROUP BY` creates buckets before aggregation."
enabled: true
---

## zh-TW

`GROUP BY` 會依欄位值把多列資料分成多個群組，常拿來搭配 `COUNT`、`SUM`、`AVG` 等聚合函式。
它不是單純排序，而是把資料切成多個統計桶。

## en

`GROUP BY` splits rows into groups based on column values and is often used with aggregate functions like `COUNT`, `SUM`, or `AVG`.
It is not just sorting; it creates buckets for summary calculations.
