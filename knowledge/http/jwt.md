---
id: http-jwt-stateless-token
title_zh: JWT
title_en: JWT
type: true-false
body_format: bilingual-section
tags: [http, jwt, auth]
difficulty: 2
question_zh: "JWT 常被用來攜帶可驗證的身份資訊，讓系統在某些情境下減少伺服器端 session 依賴。"
question_en: "A JWT is often used to carry verifiable identity data so a system can reduce server-side session dependency in some cases."
answer: true
clickbait_zh: "不是所有登入狀態都要靠 session 留在伺服器裡，有些會直接把憑證帶在身上。"
clickbait_en: "Not every auth flow stores everything on the server. Some carry the credential with the request."
review_hint_zh: "JWT 是可驗證 token，不等於自動安全。"
review_hint_en: "A JWT is a verifiable token, but that does not make it automatically safe."
enabled: true
---

## zh-TW

JWT 常用來攜帶簽名過的身份資訊，讓伺服器能在收到 token 時自行驗證內容。
它常被拿來做 stateless 風格的身份傳遞，但這不代表就沒有安全與失效管理問題。

## en

JWTs are often used to carry signed identity information so a server can verify the token contents when it receives a request.
They are common in stateless-style auth flows, but that does not remove security or revocation concerns.
