---
id: go-interface-behavior-contract
title_zh: Go interface
title_en: Go Interface
type: true-false
body_format: bilingual-section
tags: [go, interface, type]
difficulty: 2
question_zh: "Go interface 主要描述的是一組行為需求，而不是一份資料欄位清單。"
question_en: "A Go interface mainly describes a required set of behaviors, not a list of data fields."
answer: true
clickbait_zh: "你以為 interface 在定義資料長相？其實它更像在定義『會做什麼』。"
clickbait_en: "If you think an interface describes shape, you are missing the more important part: behavior."
review_hint_zh: "interface 重點在 method set。"
review_hint_en: "An interface is defined by its method set."
confusion_with: [go-struct-field-grouping]
metaphor_seed: [行為契約, 會做什麼, 規格]
hook_style_tags: [misunderstood, comparison]
enabled: true
---

## zh-TW

在 Go 裡，interface 主要不是拿來描述資料欄位，而是描述某個值需要提供哪些方法。
只要一個型別的方法集合符合這份需求，它就能滿足該 interface。
所以 interface 更像行為契約，而不是資料結構模板。

## en

In Go, an interface is mainly about which methods a value provides, not about what fields it stores.
If a type has the required method set, it satisfies the interface.
That is why an interface acts more like a behavior contract than a data template.
