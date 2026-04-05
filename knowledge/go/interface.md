---
id: go-interface-behavior-contract
title_zh: interface 行為契約
title_en: Go Interface
type: true-false
body_format: bilingual-section
tags: [go, type-system, interface]
difficulty: 2
question_zh: "在 Go 裡，interface 主要是描述一組方法行為。"
question_en: "In Go, an interface mainly describes a set of method behaviors."
answer: true
clickbait_zh: "你以為 interface 是在列欄位？在 Go 裡它更像是在定義『你會做什麼』。"
clickbait_en: "Think an interface is about listing fields? In Go it is much more about what a value can do."
review_hint_zh: "interface 在 Go 裡描述的是方法集合。"
review_hint_en: "An interface in Go describes a method set."
confusion_with: [go-struct-field-grouping]
metaphor_seed: [角色, 契約, 能力清單]
hook_style_tags: [comparison, misunderstood]
enabled: true
---

## zh-TW

在 Go 裡，interface 用來描述一組方法，也就是某個值「會做什麼」。
只要某個型別實作了這些方法，它就能滿足該 interface，不需要顯式宣告。
這讓 Go 的抽象比較偏向行為，而不是繼承樹。

## en

In Go, an interface describes a set of methods, meaning what a value can do.
If a type implements those methods, it satisfies the interface automatically without an explicit declaration.
That makes Go's abstraction model more behavior-oriented than inheritance-oriented.
