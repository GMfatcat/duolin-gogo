---
id: go-context-cancellation
title_zh: Go context
title_en: Go Context
type: true-false
body_format: bilingual-section
tags: [go, context, concurrency]
difficulty: 3
question_zh: "Go 的 context 常用來傳遞取消訊號、deadline，或 request-scoped metadata。"
question_en: "Go context is commonly used to pass cancellation signals, deadlines, or request-scoped metadata."
answer: true
clickbait_zh: "有些 goroutine 不是做太慢，是根本沒人通知它該停了。"
clickbait_en: "Some goroutines do not run too long because they are slow. They run too long because nobody told them to stop."
review_hint_zh: "context 常見於取消、逾時與 request 範圍資訊。"
review_hint_en: "Context is commonly used for cancellation, timeout, and request-scoped data."
confusion_with: [go-goroutine-concurrency, go-channel-communication]
metaphor_seed: [停損, 傳遞截止, 共同時鐘]
hook_style_tags: [fear_of_mistake, safer_first]
enabled: true
---

## zh-TW

`context` 在 Go 裡很常用來控制一串相關工作的生命週期，例如 request 超時、手動取消，或傳遞與這次請求有關的資料。
它的重點不是拿來裝一般業務資料，而是攜帶控制資訊與範圍性的 metadata。
在 goroutine、I/O、HTTP handler 這類情境中尤其常見。

## en

`context` in Go is commonly used to control the lifetime of related work, such as request timeouts, cancellation, or request-scoped metadata.
It is not mainly for general business data; it carries control information and scoped metadata.
That is why it appears so often in goroutines, I/O flows, and HTTP handlers.
