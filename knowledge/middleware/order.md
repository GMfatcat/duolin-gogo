---
id: middleware-order-chain-effects
title_zh: Middleware Order
title_en: Middleware Order
type: true-false
body_format: bilingual-section
tags: [middleware, backend, order]
difficulty: 2
question_zh: "middleware 的註冊順序常會影響實際執行結果。"
question_en: "The registration order of middleware often affects the actual execution result."
answer: true
clickbait_zh: "不是每層都只是疊上去而已，順序一錯整條 request 鏈就會變味。"
clickbait_en: "Middleware layers do not merely stack. A wrong order can change the whole request chain."
review_hint_zh: "middleware 是有順序語意的。"
review_hint_en: "Middleware order carries real meaning."
enabled: true
---

## zh-TW

例如你可能希望先記錄 request，再做驗證，或先做錯誤處理包裝，再做回應轉換。
如果順序反了，log、auth、exception handling 的結果都可能跟預期不同。

## en

For example, you may want to log before auth, or wrap errors before transforming responses.
If the order is reversed, logging, auth, and exception handling can all behave differently than expected.
