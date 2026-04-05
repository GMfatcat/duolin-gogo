---
id: http-keep-alive-connection-reuse
title_zh: Keep-Alive
title_en: HTTP Keep-Alive
type: true-false
body_format: bilingual-section
tags: [http, keep-alive, connection]
difficulty: 2
question_zh: "Keep-Alive 的核心價值之一，是重用連線，避免每個請求都重新建立連線成本。"
question_en: "One core value of keep-alive is reusing connections so every request does not pay the full setup cost again."
answer: true
clickbait_zh: "有些延遲不是資料太大，而是你每次都在重新敲門。"
clickbait_en: "Some latency is not about payload size. It is about knocking on the door again and again."
review_hint_zh: "keep-alive 讓連線可重用。"
review_hint_en: "Keep-alive makes connection reuse possible."
enabled: true
---

## zh-TW

如果每個 HTTP 請求都重新建立一條新連線，會產生額外延遲與成本。
keep-alive 的價值就在於盡量重用既有連線，減少重複建立連線的負擔。

## en

If every HTTP request had to create a brand new connection, it would add extra latency and cost.
Keep-alive helps by reusing existing connections instead of paying that setup cost every time.
