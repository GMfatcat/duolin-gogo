---
id: http-cache-control-response-caching
title_zh: Cache-Control
title_en: Cache-Control
type: true-false
body_format: bilingual-section
tags: [http, cache, header]
difficulty: 2
question_zh: "Cache-Control 是用來告訴瀏覽器或代理如何快取回應的 header。"
question_en: "Cache-Control is an HTTP header used to tell browsers or proxies how to cache a response."
answer: true
clickbait_zh: "不是所有回應都該每次重抓。真正的差別常常就藏在一行 header。"
clickbait_en: "Not every response needs to be fetched again. Sometimes the whole difference lives inside one header."
review_hint_zh: "Cache-Control 在描述回應能不能快取、能快取多久。"
review_hint_en: "Cache-Control describes whether and how long a response can be cached."
enabled: true
---

## zh-TW

Cache-Control 會告訴瀏覽器、CDN 或 reverse proxy 這份回應能不能被快取，以及能快取多久。
常見指令像 `max-age`、`no-store`、`public`、`private` 都是在這裡設定。

## en

Cache-Control tells browsers, CDNs, or reverse proxies whether a response may be cached and for how long.
Common directives such as `max-age`, `no-store`, `public`, and `private` are configured there.
