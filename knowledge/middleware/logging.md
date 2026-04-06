---
id: middleware-logging-cross-cutting
title_zh: Logging Middleware
title_en: Logging Middleware
type: true-false
body_format: bilingual-section
tags: [middleware, logging, backend]
difficulty: 1
question_zh: "logging middleware 適合拿來集中記錄 request 路徑、狀態碼與耗時。"
question_en: "Logging middleware is a good place to centralize request path, status code, and latency logging."
answer: true
clickbait_zh: "你可以在每個 endpoint 都自己 print，但通常大家最後都會後悔。"
clickbait_en: "You can print inside every endpoint yourself, but people usually regret that approach."
review_hint_zh: "logging 是很典型的跨切關注點。"
review_hint_en: "Logging is a classic cross-cutting concern."
enabled: true
---

## zh-TW

logging 通常不是某一個業務 endpoint 自己獨有的需求，而是所有 request 都會需要。
把它放進 middleware 可以避免在每個 handler 裡重複寫一樣的紀錄邏輯。

## en

Logging is usually not unique to one business endpoint. It is needed across many requests.
Putting it in middleware avoids repeating the same logging logic in every handler.
