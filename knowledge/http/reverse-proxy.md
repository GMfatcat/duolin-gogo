---
id: http-reverse-proxy-front-door
title_zh: Reverse Proxy
title_en: Reverse Proxy
type: true-false
body_format: bilingual-section
tags: [http, proxy, backend]
difficulty: 2
question_zh: "reverse proxy 常被放在應用服務前面，當作對外入口。"
question_en: "A reverse proxy is often placed in front of application servers as the external entry point."
answer: true
clickbait_zh: "很多服務不是直接面對外網，真正站在門口的其實另有其人。"
clickbait_en: "Many services do not face the internet directly. Someone else is standing at the front door."
review_hint_zh: "reverse proxy 在前面收請求，再轉發給後面的服務。"
review_hint_en: "A reverse proxy receives requests first and forwards them to backend services."
enabled: true
---

## zh-TW

reverse proxy 會先接住外部流量，再把請求轉發到內部服務。
它常被用來做 TLS 終結、負載分流、快取、路由或統一入口。

## en

A reverse proxy accepts external traffic first and then forwards requests to internal services.
It is commonly used for TLS termination, traffic routing, caching, and acting as a single public entry point.
