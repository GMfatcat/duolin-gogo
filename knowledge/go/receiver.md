---
id: go-receiver-method-binding
title_zh: receiver 綁定方法
title_en: Go Receiver
type: true-false
body_format: bilingual-section
tags: [go, method, receiver]
difficulty: 2
question_zh: "Go 的 receiver 會把方法綁在某個型別上。"
question_en: "In Go, a receiver binds a method to a specific type."
answer: true
clickbait_zh: "那不是一般函式，只是它看起來剛好長得很像。"
clickbait_en: "That is not just a regular function. It only happens to look close to one."
review_hint_zh: "receiver 決定方法是掛在哪個型別上。"
review_hint_en: "A receiver decides which type owns the method."
confusion_with: [go-function, go-struct-field-grouping]
metaphor_seed: [掛上去, 綁定, 所屬]
hook_style_tags: [misunderstood, comparison]
enabled: true
---

## zh-TW

在 Go 裡，method 和一般 function 的差別之一是 receiver。
receiver 讓方法和某個型別產生關聯，例如 struct 或自訂型別。
看到 receiver，就可以知道這個方法是屬於哪種值。

## en

In Go, one key difference between a method and a regular function is the receiver.
The receiver ties the method to a type such as a struct or custom type.
When you see the receiver, you know what kind of value the method belongs to.
