---
id: go-error-wrapping-context
title_zh: Go error wrapping
title_en: Go Error Wrapping
type: true-false
body_format: bilingual-section
tags: [go, error, wrapping]
difficulty: 3
question_zh: "Go 的 error wrapping 常用來保留原始錯誤，同時補上更多上下文。"
question_en: "Go error wrapping is commonly used to keep the original error while adding more context."
answer: true
clickbait_zh: "真正有用的錯誤訊息，不只是說壞了，還會說壞在哪一層。"
clickbait_en: "A useful error does more than say it failed. It tells you where the failure happened."
review_hint_zh: "wrapping = 保留原錯誤，再補上下文。"
review_hint_en: "Wrapping keeps the original error and adds context."
confusion_with: [go-error-value-handling]
metaphor_seed: [多一層線索, 包住原錯, 錯誤脈絡]
hook_style_tags: [safer_first, misunderstood]
enabled: true
---

## zh-TW

在 Go 裡，error wrapping 的重點是不要只回傳一句新的錯誤字串，而是把原始錯誤一起包起來。
這樣你既能保留底層真正發生的錯誤，也能在上層補上「是哪一步失敗」的上下文。
它的價值在於保留錯誤鏈，而不是把舊錯誤藏起來。

## en

In Go, error wrapping is about returning a new error message while still preserving the original error underneath.
That lets you keep the real low-level failure and add higher-level context such as which step failed.
Its value is in preserving the error chain instead of hiding the original cause.
