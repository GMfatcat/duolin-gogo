---
id: backend-idempotency-key-safe-retry
title_zh: Idempotency Key
title_en: Idempotency Key
type: true-false
body_format: bilingual-section
tags: [backend, idempotency, api]
difficulty: 3
question_zh: "idempotency key 常用來避免重試請求時重複建立資源或重複扣款。"
question_en: "An idempotency key is commonly used to avoid duplicate creation or duplicate charging when a request is retried."
answer: true
clickbait_zh: "重試不是免費的。如果你沒留辨識碼，同一張單可能真的會下兩次。"
clickbait_en: "Retries are not free. Without an identity key, the same operation may actually happen twice."
review_hint_zh: "idempotency key 幫伺服器辨認這是不是同一筆重試。"
review_hint_en: "An idempotency key helps the server recognize that retries belong to the same logical operation."
enabled: true
---

## zh-TW

當客戶端因 timeout 或網路問題重送請求時，idempotency key 能幫伺服器判斷這是不是同一筆操作。
這在付款、下單、建立資源時特別重要，因為你不想讓重試變成真的重複執行。

## en

When a client retries because of a timeout or network issue, an idempotency key helps the server tell whether it is the same logical operation.
That matters a lot for payments, orders, or resource creation where a retry must not become a real duplicate action.
