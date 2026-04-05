---
id: go-struct-tag-metadata
title_zh: struct tag 欄位註記
title_en: Go Struct Tag
type: true-false
body_format: bilingual-section
tags: [go, struct, metadata]
difficulty: 2
question_zh: "Go 的 struct tag 常被用來提供序列化或驗證相關的額外資訊。"
question_en: "Go struct tags are often used to provide extra metadata for serialization or validation."
answer: true
clickbait_zh: "欄位名字看起來沒變，但框架讀到的規則其實藏在旁邊那串字裡。"
clickbait_en: "The field name looks unchanged, but the rule the framework reads is hidden in the tag."
review_hint_zh: "struct tag 常被 JSON、ORM、validation 工具讀取。"
review_hint_en: "Struct tags are commonly read by JSON, ORM, and validation tooling."
confusion_with: [go-struct-field-grouping]
metaphor_seed: [標籤, 備註, 隱藏規則]
hook_style_tags: [misunderstood, comparison]
enabled: true
---

## zh-TW

Go 的 struct tag 是寫在欄位旁邊的額外字串資訊。
很多工具會讀它來決定 JSON key、資料庫欄位映射或驗證規則。
它不是欄位值本身，而是描述欄位如何被外部工具理解。

## en

A Go struct tag is extra string metadata written beside a field.
Many tools read it to decide JSON keys, database mappings, or validation rules.
It is not the field value itself; it describes how outside tooling should interpret the field.
