---
id: go-interface-nil-trap
title_zh: nil interface 陷阱
title_en: Go Nil Interface Trap
type: true-false
body_format: bilingual-section
tags: [go, interface, nil]
difficulty: 3
question_zh: "在 Go 中，一個 interface 變數裡若裝著 typed nil，它不一定會等於真正的 nil。"
question_en: "In Go, an interface value holding a typed nil does not necessarily compare equal to true nil."
answer: true
clickbait_zh: "你看到 nil，不代表 Go 真的也這樣認為。"
clickbait_en: "You may see nil, but Go might still disagree."
review_hint_zh: "interface 值會同時帶型別與值，typed nil 可能讓它不等於 nil。"
review_hint_en: "An interface carries both type and value, so typed nil can still make the interface non-nil."
confusion_with: [go-interface-behavior-contract, go-pointer-memory-address]
metaphor_seed: [假空值, 外殼還在, nil 陷阱]
hook_style_tags: [fear_of_mistake, misunderstood]
enabled: true
---

## zh-TW

Go 的 interface 值通常可以想成同時包含「動態型別」和「動態值」。
如果裡面的值是某個型別的 nil 指標，型別資訊仍然存在，所以整個 interface 可能不等於真正的 nil。
這是 Go 常見的陷阱之一，尤其在錯誤處理或回傳 interface 時很容易踩到。

## en

A Go interface value can be thought of as carrying both a dynamic type and a dynamic value.
If the stored value is a nil pointer of a concrete type, the type information still exists, so the whole interface may not be equal to true nil.
This is a classic Go trap, especially around error handling and interface returns.
