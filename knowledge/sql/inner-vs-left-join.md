---
id: sql-inner-vs-left-join
title_zh: INNER JOIN vs LEFT JOIN
title_en: Inner Join vs Left Join
type: true-false
body_format: bilingual-section
tags: [sql, join, left-join]
difficulty: 2
question_zh: "`LEFT JOIN` 與 `INNER JOIN` 的差別之一，是前者可保留左表沒有匹配到的列。"
question_en: "One difference between `LEFT JOIN` and `INNER JOIN` is that the former can keep left-side rows even when no match exists."
answer: true
clickbait_zh: "有些資料不是消失了，只是你用錯了 join。"
clickbait_en: "Sometimes the rows did not disappear. You just used the wrong join."
review_hint_zh: "`INNER JOIN` 只留匹配列，`LEFT JOIN` 還能保留左側未匹配列。"
review_hint_en: "`INNER JOIN` keeps only matches, while `LEFT JOIN` can keep unmatched left rows."
enabled: true
---

## zh-TW

`INNER JOIN` 只會留下兩邊都能配對成功的列。
`LEFT JOIN` 則會保留左表所有列，就算右表沒找到對應資料也一樣，只是右側欄位會是 `NULL`。

## en

`INNER JOIN` keeps only rows that match on both sides.
`LEFT JOIN` keeps every row from the left table even if the right side has no match, in which case the right-side columns become `NULL`.
