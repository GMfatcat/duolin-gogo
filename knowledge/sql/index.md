---
id: sql-index-speed-lookup
title_zh: SQL Index
title_en: SQL Index
type: true-false
body_format: bilingual-section
tags: [sql, index, performance]
difficulty: 2
question_zh: "Index 的主要目的之一，是讓查詢能更快定位資料。"
question_en: "One main purpose of an index is to help queries locate data faster."
answer: true
clickbait_zh: "查詢慢不一定是資料庫笨，可能只是你沒給它地圖。"
clickbait_en: "A slow query does not always mean the database is dumb. It may just be missing a map."
review_hint_zh: "index 像查找地圖，幫助快速定位。"
review_hint_en: "An index acts like a lookup map for faster access."
enabled: true
---

## zh-TW

index 的核心價值在於縮短定位資料的成本，尤其在常被查詢或排序的欄位上很有幫助。
但 index 不是免費的，因為它也會增加寫入與維護成本。

## en

The core value of an index is reducing the cost of finding data, especially on columns that are searched or sorted often.
But an index is not free because it also adds write and maintenance overhead.
