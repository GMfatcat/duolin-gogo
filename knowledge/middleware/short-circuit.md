---
id: middleware-short-circuit-response
title_zh: Short-Circuit Middleware
title_en: Short-Circuit Middleware
type: true-false
body_format: bilingual-section
tags: [middleware, backend, response]
difficulty: 2
question_zh: "某些 middleware 可以直接回傳 response，而不再把 request 傳給後面的 handler。"
question_en: "Some middleware can return a response directly instead of passing the request to later handlers."
answer: true
clickbait_zh: "有些 request 根本進不到你的 endpoint，半路就被 middleware 勸返了。"
clickbait_en: "Some requests never reach your endpoint because middleware turns them away halfway through."
review_hint_zh: "middleware 可以在鏈中提早終止流程。"
review_hint_en: "Middleware can terminate the chain early."
enabled: true
---

## zh-TW

像驗證失敗、缺少權限、健康檢查或某些快取命中情境，都可能直接由 middleware 回傳 response。
這種行為常被稱為 short-circuit，代表 request 沒有再進到真正的業務 handler。

## en

Cases like failed authentication, missing permission, health checks, or some cache hits may return directly from middleware.
That behavior is often called short-circuiting, meaning the request never reaches the business handler.
