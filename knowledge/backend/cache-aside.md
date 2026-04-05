---
id: backend-cache-aside-pattern
title_zh: Cache Aside
title_en: Cache Aside
type: true-false
body_format: bilingual-section
tags: [backend, cache, cache-aside]
difficulty: 2
question_zh: "cache aside 的常見流程是先讀 cache，miss 時再查主資料來源並回填 cache。"
question_en: "A common cache-aside flow is to read cache first, then on a miss load from the source of truth and fill the cache."
answer: true
clickbait_zh: "快取不是憑空知道答案，它常常只是先偷看一眼有沒有現成的。"
clickbait_en: "A cache does not magically know the answer. It often just checks whether a ready copy already exists."
review_hint_zh: "cache aside = 先查 cache，miss 再回源並回填。"
review_hint_en: "Cache aside means cache first, then source on miss, then fill cache."
enabled: true
---

## zh-TW

cache aside 是常見快取模式：先查 cache，有命中就直接用；沒命中再查資料庫或主來源，並把結果放回 cache。
它實作直覺，但也要小心失效策略。

## en

Cache aside is a common caching pattern: read from cache first, and if there is a miss, load from the database or source of truth and then populate the cache.
It is intuitive, but invalidation still needs care.
