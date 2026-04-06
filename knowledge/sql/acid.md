---
id: sql-acid-transaction-guarantees
title_zh: ACID
title_en: ACID
type: true-false
body_format: bilingual-section
tags: [sql, transaction, acid]
difficulty: 2
question_zh: "ACID 常用來描述資料庫 transaction 的四種核心保證。"
question_en: "ACID is commonly used to describe the four core guarantees of a database transaction."
answer: true
clickbait_zh: "交易真的可靠，不是因為資料庫看起來很穩，而是因為背後有 ACID 在撐。"
clickbait_en: "Transactions feel reliable not because the database looks serious, but because ACID is doing real work underneath."
review_hint_zh: "ACID = Atomicity, Consistency, Isolation, Durability。"
review_hint_en: "ACID stands for Atomicity, Consistency, Isolation, and Durability."
enabled: true
---

## zh-TW

ACID 是資料庫交易語意中很經典的一組原則，代表 Atomicity、Consistency、Isolation、Durability。
它幫助你理解為什麼一組資料修改可以被視為一個整體、彼此不亂干擾，而且在提交後應該被可靠保存。

## en

ACID is a classic set of database transaction guarantees: Atomicity, Consistency, Isolation, and Durability.
It helps explain why a group of changes can be treated as one unit, kept from interfering with each other, and preserved after commit.
