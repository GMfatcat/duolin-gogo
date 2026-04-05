---
id: sql-where-vs-having-trap
title_zh: WHERE vs HAVING
title_en: WHERE vs HAVING
type: true-false
body_format: bilingual-section
tags: [sql, where, having, aggregate]
difficulty: 2
question_zh: "`WHERE` 和 `HAVING` 常被搞混，但前者偏列過濾，後者偏聚合後的群組過濾。"
question_en: "`WHERE` and `HAVING` are often confused, but the former filters rows while the latter filters groups after aggregation."
answer: true
clickbait_zh: "很多 SQL 看起來只是條件寫錯，其實是你把判斷時機放到錯的階段。"
clickbait_en: "Many SQL mistakes are not wrong conditions. They are the right conditions at the wrong stage."
review_hint_zh: "`WHERE` 先篩列，`HAVING` 後篩群組。"
review_hint_en: "`WHERE` filters rows first; `HAVING` filters groups later."
confusion_with: [sql-where-filter-rows, sql-having-post-group-filter]
enabled: true
---

## zh-TW

`WHERE` 主要對原始列資料下條件，`HAVING` 則常用在分組聚合後對群組結果下條件。
如果你把聚合條件放到 `WHERE`，通常就表示判斷時機放錯了。

## en

`WHERE` mainly applies conditions to raw rows, while `HAVING` is usually used after grouping and aggregation.
If you place an aggregate condition in `WHERE`, the problem is often that the condition is being applied at the wrong stage.
