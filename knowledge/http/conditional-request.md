---
id: http-conditional-request-state-check
title_zh: Conditional Request
title_en: Conditional Request
type: true-false
body_format: bilingual-section
tags: [http, cache, conditional]
difficulty: 3
question_zh: "conditional request 會在送出請求時附帶條件，讓伺服器決定是否真的傳完整內容。"
question_en: "A conditional request sends extra conditions so the server can decide whether it really needs to return the full content."
answer: true
clickbait_zh: "真正有效率的 HTTP，不是什麼都重傳，而是先問一句：它有變嗎？"
clickbait_en: "Efficient HTTP is not about resending everything. It starts by asking one question: has it changed?"
review_hint_zh: "conditional request 常搭配 ETag 或時間戳做比對。"
review_hint_en: "Conditional requests often use ETag or timestamps for comparison."
enabled: true
---

## zh-TW

conditional request 會用像 `If-None-Match` 或 `If-Modified-Since` 這種條件 header 去問伺服器資料是否有變。
如果條件不成立，伺服器就能回比較輕的結果，而不是整份內容重送。

## en

Conditional requests use headers such as `If-None-Match` or `If-Modified-Since` to ask the server whether data has changed.
When the condition fails, the server can return a lighter response instead of resending the whole representation.
