---
id: http-idempotent-request-semantics
title_zh: Idempotent
title_en: HTTP Idempotency
type: true-false
body_format: bilingual-section
tags: [http, idempotent, api]
difficulty: 2
question_zh: "在 HTTP 語意中，idempotent 指的是同一請求重複執行多次，最終效果應該一致。"
question_en: "In HTTP semantics, idempotent means repeating the same request should lead to the same final effect."
answer: true
clickbait_zh: "有些請求可以重送，有些重送一次就出事，差別就在這個字。"
clickbait_en: "Some requests can be retried safely, while others can cause trouble. The difference often hides in this one word."
review_hint_zh: "idempotent 重點在重複請求後的最終效果一致。"
review_hint_en: "Idempotent focuses on the same final effect after repeated requests."
enabled: true
---

## zh-TW

idempotent 的意思不是「只做一次」，而是「做很多次後，最終狀態和做一次相同」。
這個概念對 retry、安全更新與 API 設計都很重要。

## en

Idempotent does not mean "only once." It means that repeating the same request leads to the same final state as doing it once.
That idea matters a lot for retries, safe updates, and API design.
