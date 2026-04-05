---
id: sql-join-too-many-cost
title_zh: 太多 JOIN 的成本
title_en: Too Many Joins Cost
type: true-false
body_format: bilingual-section
tags: [sql, join, performance]
difficulty: 3
question_zh: "JOIN 越多不代表一定錯，但通常也代表查詢計畫與成本更需要小心檢查。"
question_en: "More joins do not automatically mean a query is wrong, but they usually mean the plan and cost deserve closer inspection."
answer: true
clickbait_zh: "表不是不能接很多張，但每多接一次，資料庫就多一層要想的成本。"
clickbait_en: "It is not forbidden to join many tables, but each join adds another layer of cost the database must reason about."
review_hint_zh: "多 JOIN 不一定錯，但更需要看 `EXPLAIN` 與索引。"
review_hint_en: "Many joins are not automatically wrong, but they deserve more `EXPLAIN` and index scrutiny."
enabled: true
---

## zh-TW

多個 `JOIN` 在商業查詢裡很常見，但它們通常也會讓查詢計畫更複雜。
這不代表一定不能寫，而是代表你更該確認索引、條件和 `EXPLAIN` 結果是否合理。

## en

Multiple joins are common in real business queries, but they often make the execution plan more complex.
That does not mean they are forbidden; it means indexes, conditions, and the `EXPLAIN` output deserve closer review.
