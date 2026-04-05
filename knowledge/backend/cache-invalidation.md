---
id: backend-cache-invalidation-hard-problem
title_zh: Cache Invalidation
title_en: Cache Invalidation
type: true-false
body_format: bilingual-section
tags: [backend, cache, invalidation]
difficulty: 3
question_zh: "cache invalidation 困難的一個原因，是資料更新後必須確保舊快取不要繼續誤導讀取。"
question_en: "One reason cache invalidation is hard is that stale cached data must stop misleading future reads after the source changes."
answer: true
clickbait_zh: "快取最可怕的不是 miss，而是它很自信地把舊答案拿給你。"
clickbait_en: "The scariest cache failure is not a miss. It is a stale answer delivered with confidence."
review_hint_zh: "cache invalidation 的難點在 stale data。"
review_hint_en: "The difficulty of invalidation is stale data."
enabled: true
---

## zh-TW

快取加速讀取很有用，但資料一旦更新，舊 cache 如果還留著，就會變成 stale data。
cache invalidation 難就難在你要在速度、正確性與複雜度之間做平衡。

## en

Caches are great for speeding up reads, but once the source data changes, old cache entries can turn into stale data.
Cache invalidation is hard because it forces a tradeoff between speed, correctness, and complexity.
