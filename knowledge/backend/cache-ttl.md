---
id: backend-cache-ttl-expiration-window
title_zh: Cache TTL
title_en: Cache TTL
type: true-false
body_format: bilingual-section
tags: [backend, cache, ttl]
difficulty: 2
question_zh: "TTL 常用來定義 cache 項目可以被視為有效多久。"
question_en: "TTL is commonly used to define how long a cache entry should be considered valid."
answer: true
clickbait_zh: "快取不是放進去就永遠對，很多時候它其實帶著有效期限。"
clickbait_en: "A cache entry is not correct forever. Very often it comes with an expiration window."
review_hint_zh: "TTL 是快取有效時間。"
review_hint_en: "TTL is the validity window of cached data."
enabled: true
---

## zh-TW

TTL 是 time to live，常用來表示快取資料在多久後應該失效或重新抓取。
它是控制新鮮度與效能取捨的重要參數。

## en

TTL stands for time to live and is commonly used to say when cached data should expire or be refreshed.
It is an important parameter in balancing freshness against performance.
