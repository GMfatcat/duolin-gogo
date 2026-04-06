---
id: backend-rate-limit-protect-capacity
title_zh: Rate Limiting
title_en: Rate Limiting
type: true-false
body_format: bilingual-section
tags: [backend, rate-limit, protection]
difficulty: 2
question_zh: "rate limiting 的目的之一，是避免單一使用者或流量來源把系統資源吃光。"
question_en: "One goal of rate limiting is to stop a single user or traffic source from exhausting system capacity."
answer: true
clickbait_zh: "不是所有請求都該被無限歡迎。有時候，保護系統就是先說慢一點。"
clickbait_en: "Not every request deserves unlimited access. Sometimes protecting the system means saying slow down."
review_hint_zh: "rate limiting 在限制單位時間內可接受的請求量。"
review_hint_en: "Rate limiting restricts how many requests are accepted over time."
enabled: true
---

## zh-TW

rate limiting 會限制某個 user、IP 或 token 在一段時間內能送多少請求。
它常用來保護系統容量、降低濫用風險，也能讓資源分配更公平。

## en

Rate limiting restricts how many requests a user, IP, or token can send within a time window.
It helps protect capacity, reduce abuse, and make resource usage more fair.
