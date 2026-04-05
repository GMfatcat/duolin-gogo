---
id: sql-null-comparison-trap
title_zh: NULL 比較陷阱
title_en: SQL NULL Comparison Trap
type: true-false
body_format: bilingual-section
tags: [sql, null, bug]
difficulty: 3
question_zh: "在 SQL 中，`NULL = NULL` 不會像一般值比較那樣直接得到 true。"
question_en: "In SQL, `NULL = NULL` does not behave like a normal equality check that directly returns true."
answer: true
clickbait_zh: "你以為空值和空值應該相等，但 SQL 對『不知道』這件事沒那麼直覺。"
clickbait_en: "You may think empty should equal empty, but SQL treats unknown in a less intuitive way."
review_hint_zh: "`NULL` 代表未知，通常要用 `IS NULL` 判斷。"
review_hint_en: "`NULL` means unknown, so it is usually checked with `IS NULL`."
enabled: true
---

## zh-TW

在 SQL 裡，`NULL` 比較像「未知值」，不是一般可直接比較的普通內容。
所以判斷空值時，常見作法不是用 `=`，而是用 `IS NULL` 或 `IS NOT NULL`。

## en

In SQL, `NULL` behaves more like an unknown value than an ordinary comparable value.
That is why null checks are usually written with `IS NULL` or `IS NOT NULL` instead of plain equality.
