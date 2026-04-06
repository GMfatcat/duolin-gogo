---
id: sql-deadlock-cycle-contention
title_zh: Deadlock
title_en: Deadlock
type: true-false
body_format: bilingual-section
tags: [sql, transaction, lock]
difficulty: 3
question_zh: "deadlock 通常表示兩個交易互相等待對方釋放資源。"
question_en: "A deadlock usually means two transactions are waiting on each other to release resources."
answer: true
clickbait_zh: "資料庫有時不是慢，是兩邊都卡住不肯放手。"
clickbait_en: "Sometimes the database is not slow. Two transactions are just stuck refusing to let go."
review_hint_zh: "deadlock = 互相等待資源，最後必須由系統打斷其中一方。"
review_hint_en: "A deadlock is mutual waiting, so the database must abort one side."
enabled: true
---

## zh-TW

deadlock 常發生在多個交易以不同順序鎖住多筆資料時，最後彼此卡住。
資料庫通常會偵測這種循環等待，並回滾其中一個交易。

## en

Deadlocks often happen when multiple transactions lock several rows in different orders and end up blocking each other.
Databases usually detect that circular wait and roll back one transaction to break the cycle.
