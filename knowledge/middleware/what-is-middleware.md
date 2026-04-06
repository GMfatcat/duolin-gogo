---
id: middleware-request-pipeline-layer
title_zh: Middleware
title_en: Middleware
type: true-false
body_format: bilingual-section
tags: [middleware, backend, web]
difficulty: 1
question_zh: "middleware 常指夾在 request 與 response 流程中的處理層。"
question_en: "Middleware commonly refers to a processing layer that sits inside the request and response flow."
answer: true
clickbait_zh: "你以為只是進 request、出 response？中間其實常藏著一整條攔截帶。"
clickbait_en: "You thought it was just request in, response out, but there is often a whole interception layer in between."
review_hint_zh: "middleware 是流程中的中介處理層。"
review_hint_en: "Middleware is an intermediate processing layer in the flow."
enabled: true
---

## zh-TW

middleware 常見於 web framework，用來在 request 進入實際 handler 前後做共通處理。
像是驗證、記錄、量測、跨域設定，都很適合放在這個位置。

## en

Middleware is common in web frameworks and is used to run shared processing before or after the actual handler.
Authentication, logging, metrics, and CORS handling are all common middleware responsibilities.
