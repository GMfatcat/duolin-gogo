---
id: sql-having-post-group-filter
title_zh: SQL HAVING
title_en: SQL HAVING
type: true-false
body_format: bilingual-section
tags: [sql, having, aggregate]
difficulty: 2
question_zh: "`HAVING` 常用來在分組聚合後，再過濾群組結果。"
question_en: "`HAVING` is commonly used to filter grouped results after aggregation."
answer: true
clickbait_zh: "有些條件不是在列上判斷，而是要等整組算完才知道能不能留下。"
clickbait_en: "Some conditions cannot be judged row by row. You only know after the whole group is computed."
review_hint_zh: "`HAVING` 是分組後的過濾。"
review_hint_en: "`HAVING` filters after grouping."
confusion_with: [sql-where-filter-rows, sql-where-vs-having-trap]
enabled: true
---

## zh-TW

`HAVING` 跟 `WHERE` 很像，但它主要用在 `GROUP BY` 之後，針對聚合後的群組做篩選。
常見心智模型是：`WHERE` 篩列，`HAVING` 篩群組。

## en

`HAVING` looks similar to `WHERE`, but it is mainly used after `GROUP BY` to filter grouped results.
A useful mental model is: `WHERE` filters rows, while `HAVING` filters groups.
