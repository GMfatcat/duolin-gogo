---
id: go-struct-field-grouping
title_zh: struct 欄位群組
title_en: Go Struct
type: single-choice
body_format: bilingual-section
tags: [go, type-system, struct]
difficulty: 1
question_zh: "Go 的 `struct` 最主要用來做什麼？"
question_en: "What is a Go `struct` mainly used for?"
choices_zh:
  - "把多個欄位資料組成一個型別"
  - "直接啟動背景 goroutine"
  - "建立 key-value 查表"
choices_en:
  - "Group multiple fields into one type"
  - "Launch background goroutines directly"
  - "Create a key-value lookup table"
answer: 0
clickbait_zh: "不是每一段資料都該散在變數裡，有些東西就是該被整理成一個實體。"
clickbait_en: "Not every piece of data should float around as separate variables. Some of it wants to become a single entity."
review_hint_zh: "`struct` 會把多個欄位組成一個型別。"
review_hint_en: "`struct` groups multiple fields into one type."
confusion_with: [go-interface-behavior-contract, go-map-key-value]
metaphor_seed: [表單, 身分卡, 打包]
hook_style_tags: [comparison, safer_first]
enabled: true
---

## zh-TW

`struct` 是 Go 裡用來把多個欄位資料打包成一個型別的方式。
當某些資料天然屬於同一個實體，例如使用者名稱、年齡與 email，就很適合放進同一個 struct。
它是在描述資料形狀，不是在描述行為。

## en

A `struct` in Go groups multiple fields into a single type.
When several pieces of data naturally belong to one entity, such as a user's name, age, and email, a struct is a good fit.
It describes data shape rather than behavior.
