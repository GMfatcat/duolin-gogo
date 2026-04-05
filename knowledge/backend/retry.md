---
id: backend-retry-transient-failure-policy
title_zh: Retry
title_en: Retry Policy
type: true-false
body_format: bilingual-section
tags: [backend, retry, reliability]
difficulty: 2
question_zh: "retry 通常比較適合處理暫時性失敗，而不是所有錯誤都無條件重試。"
question_en: "Retry is usually more appropriate for transient failures rather than blindly retrying every kind of error."
answer: true
clickbait_zh: "重試不是萬靈丹，有些錯誤你多按幾次只會把災情放大。"
clickbait_en: "Retry is not magic. Some failures only get worse when you repeat them."
review_hint_zh: "retry 要看錯誤是否暫時且可恢復。"
review_hint_en: "Retry works best when the failure is temporary and recoverable."
enabled: true
---

## zh-TW

retry 常見於網路抖動、暫時性服務不可用等情境，但不是所有錯誤都該重試。
如果錯誤本質是邏輯錯、權限錯或參數錯，一直重試通常沒有幫助。

## en

Retry is common in situations like network hiccups or temporary service unavailability, but not every error should be retried.
If the problem is a logic bug, permission issue, or invalid input, repeated retries usually do not help.
