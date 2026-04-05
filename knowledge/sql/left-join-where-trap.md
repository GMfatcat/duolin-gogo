---
id: sql-left-join-where-trap
title_zh: LEFT JOIN + WHERE 陷阱
title_en: Left Join Where Trap
type: true-false
body_format: bilingual-section
tags: [sql, left-join, where, bug]
difficulty: 3
question_zh: "在 `LEFT JOIN` 後若直接在 `WHERE` 過濾右表欄位，可能把原本想保留的左表未匹配列也排掉。"
question_en: "If you filter right-side columns in `WHERE` after a `LEFT JOIN`, you may accidentally remove the unmatched left-side rows you meant to keep."
answer: true
clickbait_zh: "你明明用了 LEFT JOIN，資料還是消失？很多時候凶手就在後面的 WHERE。"
clickbait_en: "You used LEFT JOIN and still lost rows? The real culprit is often the WHERE clause that came after it."
review_hint_zh: "LEFT JOIN 之後過濾右表欄位時，要小心把未匹配列一起排掉。"
review_hint_en: "After a LEFT JOIN, filtering right-side columns can accidentally drop unmatched rows."
enabled: true
---

## zh-TW

`LEFT JOIN` 的目的之一，是保留左表所有列，即使右表沒有匹配資料也一樣。
但如果你在 `WHERE` 直接要求右表欄位滿足某條件，未匹配列的 `NULL` 往往會被一起排除，效果看起來就更像 `INNER JOIN`。

## en

One purpose of `LEFT JOIN` is to keep all rows from the left table even when the right side has no match.
But if the `WHERE` clause directly filters right-side columns, unmatched rows with `NULL` values often get removed too, making the result behave more like an `INNER JOIN`.
