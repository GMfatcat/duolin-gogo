---
id: patterns-factory-function-creation
title_zh: Factory Function
title_en: Factory Function
type: true-false
body_format: bilingual-section
tags: [design-patterns, factory, creation]
difficulty: 2
question_zh: "factory function 常用來把物件建立邏輯包在一個函式裡，而不是讓呼叫端自己拼裝。"
question_en: "A factory function is commonly used to wrap object creation logic in one function instead of making the caller assemble everything manually."
answer: true
clickbait_zh: "真正麻煩的不是物件本身，而是每個地方都在偷偷複製同一段建立邏輯。"
clickbait_en: "The annoying part is often not the object itself, but the same creation logic being copied everywhere."
review_hint_zh: "factory function 重點是集中建立邏輯。"
review_hint_en: "A factory function centralizes creation logic."
enabled: true
---

## zh-TW

factory function 能把預設值、條件分支與建立流程收在同一個地方。
這樣呼叫端只要說明需求，不必每次都自己組出完整物件。

## en

A factory function can centralize defaults, conditional branches, and creation flow in one place.
That lets callers describe what they need without rebuilding the full object manually every time.
