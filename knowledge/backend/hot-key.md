---
id: backend-hot-key-cache-skew
title_zh: Hot Key
title_en: Cache Hot Key
type: true-false
body_format: bilingual-section
tags: [backend, cache, hot-key]
difficulty: 3
question_zh: "hot key 問題常指某個極熱門 key 造成快取或後端負載高度集中。"
question_en: "A hot key problem commonly means one extremely popular key concentrates too much load on cache or backend systems."
answer: true
clickbait_zh: "快取有時不是平均被打，而是某一個 key 被全世界一起狂敲。"
clickbait_en: "Sometimes cache traffic is not spread out at all. One key gets hammered by everyone."
review_hint_zh: "hot key = 流量集中在少數熱門 key。"
review_hint_en: "A hot key means traffic is heavily concentrated on a few popular keys."
enabled: true
---

## zh-TW

當某個 key 特別熱門時，請求量可能高度集中在同一個快取項目或同一段後端邏輯上。
這會造成局部壅塞，即使整體流量看起來沒有那麼誇張。

## en

When one key becomes extremely popular, request volume can concentrate on a single cache entry or one slice of backend logic.
That creates a local hotspot even if the overall traffic does not look extreme.
