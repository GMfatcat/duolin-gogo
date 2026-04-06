---
id: http-cors-preflight-permission-check
title_zh: CORS Preflight
title_en: CORS Preflight
type: true-false
body_format: bilingual-section
tags: [http, cors, browser]
difficulty: 3
question_zh: "CORS preflight 是瀏覽器在某些跨來源請求前先送出的權限檢查。"
question_en: "A CORS preflight is a permission check that the browser sends before certain cross-origin requests."
answer: true
clickbait_zh: "你以為 API 壞了？很多時候只是瀏覽器先偷偷問了一句『我能送嗎？』"
clickbait_en: "You may think the API is broken, but often the browser first whispers: am I allowed to send this?"
review_hint_zh: "preflight 是瀏覽器先送的 OPTIONS 權限檢查。"
review_hint_en: "A preflight is usually a browser-sent OPTIONS permission check."
enabled: true
---

## zh-TW

當請求方法或 header 不屬於簡單請求時，瀏覽器通常會先送一個 OPTIONS preflight。
伺服器要先表明允許哪些來源、方法和 header，真正的跨域請求才會繼續。

## en

When the method or headers are not considered simple, the browser usually sends an OPTIONS preflight first.
The server must say which origins, methods, and headers are allowed before the real cross-origin request proceeds.
