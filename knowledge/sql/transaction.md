---
id: sql-transaction-all-or-nothing
title_zh: Transaction
title_en: SQL Transaction
type: true-false
body_format: bilingual-section
tags: [sql, transaction]
difficulty: 2
question_zh: "transaction 常用來把一組 SQL 操作包成同一個工作單位。"
question_en: "A transaction is commonly used to treat a group of SQL operations as one unit of work."
answer: true
clickbait_zh: "資料庫不是每次都單打一。真正關鍵時，它會把多個操作綁成一條命。"
clickbait_en: "Databases do not always act one statement at a time. Important changes are often tied together as one unit."
review_hint_zh: "transaction 會把多個變更包成同一個工作單位。"
review_hint_en: "A transaction groups multiple changes into one unit of work."
enabled: true
---

## zh-TW

transaction 代表多個 SQL 操作要一起成功，或一起失敗。
它常用在轉帳、下單、扣庫存這種不能只做一半的流程。

## en

A transaction means several SQL operations should succeed together or fail together.
It is useful for flows like transfers, orders, or inventory changes where half-finished work is not acceptable.
