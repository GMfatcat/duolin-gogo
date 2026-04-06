---
id: sql-isolation-level-read-boundary
title_zh: Isolation Level
title_en: Isolation Level
type: single-choice
body_format: bilingual-section
tags: [sql, transaction, isolation]
difficulty: 3
question_zh: "isolation level 最主要在控制什麼？"
question_en: "What does an isolation level mainly control?"
choices_zh:
  - "資料表名稱的長度限制"
  - "交易彼此讀寫時能看到多少未完成變更"
  - "索引檔案的大小"
choices_en:
  - "The maximum length of a table name"
  - "How much unfinished work one transaction can observe from another"
  - "The size of an index file"
answer: 1
clickbait_zh: "同一張表，為什麼兩個交易看到的世界會不一樣？"
clickbait_en: "Why can two transactions look at the same table and still see different worlds?"
review_hint_zh: "isolation level 在定義交易之間彼此可見的讀寫邊界。"
review_hint_en: "Isolation level defines the visibility boundary between transactions."
enabled: true
---

## zh-TW

isolation level 在決定一個交易能不能看到另一個交易尚未完成或剛完成的變更。
它影響 dirty read、non-repeatable read、phantom read 這些一致性問題。

## en

An isolation level decides whether one transaction can see unfinished or freshly committed changes from another.
It directly affects consistency problems such as dirty reads, non-repeatable reads, and phantom reads.
