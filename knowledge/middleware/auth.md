---
id: middleware-auth-check-gate
title_zh: Authentication Middleware
title_en: Authentication Middleware
type: true-false
body_format: bilingual-section
tags: [middleware, auth, backend]
difficulty: 2
question_zh: "authentication middleware 常在 request 到達業務 handler 前先檢查身分資訊。"
question_en: "Authentication middleware commonly checks identity information before the request reaches the business handler."
answer: true
clickbait_zh: "真正決定你能不能進門的，常常不是 endpoint 本身，而是前面那層守門員。"
clickbait_en: "The thing that decides whether you get in is often not the endpoint itself, but the gatekeeper before it."
review_hint_zh: "auth middleware 常在 handler 之前就先擋下未授權 request。"
review_hint_en: "Auth middleware often blocks unauthorized requests before the handler runs."
enabled: true
---

## zh-TW

auth middleware 的典型工作包括驗 token、解析 session、確認 request 是否具備基本權限。
這類檢查越早做，越能避免未授權請求繼續浪費後續資源。

## en

Typical auth middleware tasks include validating tokens, reading sessions, and checking basic authorization.
Doing that earlier helps prevent unauthorized requests from consuming later resources.
