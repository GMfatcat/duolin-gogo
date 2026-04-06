---
id: middleware-cors-browser-policy
title_zh: CORS Middleware
title_en: CORS Middleware
type: true-false
body_format: bilingual-section
tags: [middleware, cors, http]
difficulty: 2
question_zh: "CORS middleware 常用來控制瀏覽器是否允許跨來源請求。"
question_en: "CORS middleware is commonly used to control whether browsers allow cross-origin requests."
answer: true
clickbait_zh: "API 明明活著，前端卻說被擋？很多時候是 CORS 在門口關燈。"
clickbait_en: "The API is alive but the frontend says it is blocked? CORS is often the bouncer at the door."
review_hint_zh: "CORS 主要是瀏覽器端的跨來源政策。"
review_hint_en: "CORS is mainly a browser-side cross-origin policy."
enabled: true
---

## zh-TW

CORS 不等於一般意義上的後端權限系統，它更像瀏覽器遵守的跨來源規則。
middleware 常用來集中設定哪些來源、方法與 header 被允許。

## en

CORS is not the same thing as a general backend authorization system. It is more like a browser-enforced cross-origin policy.
Middleware is a common place to centrally define which origins, methods, and headers are allowed.
