---
id: backend-read-fetch-resource
title_zh: Read
title_en: Read
type: true-false
body_format: bilingual-section
tags: [backend, crud, read]
difficulty: 1
question_zh: "在 CRUD 裡，Read 指的是查詢或取得既有資料。"
question_en: "In CRUD, Read refers to retrieving or querying existing data."
answer: true
clickbait_zh: "你以為只是把資料撈出來？Read 往往才是系統效能的真正壓力源。"
clickbait_en: "You thought it was only fetching data, but reads are often the real pressure point for performance."
review_hint_zh: "Read 的重點是取資料。"
review_hint_en: "Read is about fetching existing data."
enabled: true
---

## zh-TW

Read 包含單筆查詢、列表查詢、搜尋與過濾等常見行為。
在真實系統裡，讀取常比寫入更頻繁，因此效能和快取策略常跟 Read 關係很深。

## en

Read includes fetching one record, listing many records, searching, and filtering.
In real systems, reads are often more frequent than writes, so performance and caching are strongly tied to this operation.
