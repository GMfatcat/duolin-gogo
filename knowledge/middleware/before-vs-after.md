---
id: middleware-before-vs-after
title_zh: Before vs After Middleware
title_en: Before vs After Middleware
type: true-false
body_format: bilingual-section
tags: [middleware, backend, request]
difficulty: 2
question_zh: "同一個 middleware 可以在把 request 傳下去前做事，也能在 response 回來後再做事。"
question_en: "The same middleware can often do work before passing the request onward and again after the response comes back."
answer: true
clickbait_zh: "它不是只會攔門口，還能在你回程時補一刀。"
clickbait_en: "Middleware does not only block the front door. It can also act on the way back."
review_hint_zh: "middleware 常同時有 pre 與 post 的處理時機。"
review_hint_en: "Middleware often has both pre- and post-processing phases."
enabled: true
---

## zh-TW

很多 middleware 實作會先在 request 進來時做前置處理，再把控制權往下傳。
等 response 回來之後，它還可以補上後處理，例如加 header、記錄耗時或統一包裝輸出。

## en

Many middleware implementations do some work when the request first arrives, then pass control onward.
After the response comes back, the same middleware may still add headers, measure timing, or wrap output consistently.
