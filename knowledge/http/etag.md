---
id: http-etag-version-validator
title_zh: ETag
title_en: ETag
type: true-false
body_format: bilingual-section
tags: [http, cache, etag]
difficulty: 3
question_zh: "ETag 常被用來判斷資源內容有沒有改變。"
question_en: "ETag is commonly used to check whether a resource representation has changed."
answer: true
clickbait_zh: "有些 304 不是魔法，是伺服器和客戶端早就偷偷對好暗號。"
clickbait_en: "Some 304 responses are not magic. The client and server already agreed on a quiet little fingerprint."
review_hint_zh: "ETag 常當作資源版本或內容指紋。"
review_hint_en: "ETag often acts like a version or fingerprint for a resource."
enabled: true
---

## zh-TW

ETag 可以把某份資源標成一個版本指紋，讓客戶端下次請求時帶著它回來比對。
如果資源沒變，伺服器就能回 `304 Not Modified`，減少重複傳輸。

## en

An ETag can label a resource with a version-like fingerprint so the client can send it back on the next request.
If nothing changed, the server can reply with `304 Not Modified` and avoid sending the full body again.
