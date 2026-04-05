---
id: go-error-value-handling
title_zh: error 值處理
title_en: Go Error Value
type: true-false
body_format: bilingual-section
tags: [go, error-handling, function]
difficulty: 2
question_zh: "Go 常用回傳 `error` 值的方式來處理失敗。"
question_en: "Go commonly handles failures by returning an `error` value."
answer: true
clickbait_zh: "Go 很少把例外藏起來，它通常會直接把錯誤交到你手上。"
clickbait_en: "Go rarely hides failure. It usually hands the error directly back to you."
review_hint_zh: "Go 習慣明確回傳 `error`。"
review_hint_en: "Go favors explicit `error` returns."
confusion_with: [python-exception-try-except]
metaphor_seed: [交回來, 明講, 顯式]
hook_style_tags: [comparison, misunderstood]
enabled: true
---

## zh-TW

Go 常把錯誤當成一般回傳值的一部分。
呼叫端需要檢查 `err != nil`，再決定怎麼處理。
這種做法讓失敗路徑比較明確，也避免把控制流程藏在例外裡。

## en

Go often treats errors as ordinary return values.
The caller checks `err != nil` and decides what to do next.
This keeps failure paths explicit instead of hiding control flow inside exceptions.
