---
id: oop-inheritance-vs-composition
title_zh: Inheritance vs Composition
title_en: Inheritance vs Composition
type: true-false
body_format: bilingual-section
tags: [oop, inheritance, composition]
difficulty: 2
question_zh: "在工程實務裡，composition 常被認為比 inheritance 更靈活。"
question_en: "In practical engineering, composition is often considered more flexible than inheritance."
answer: true
clickbait_zh: "看到 `extends` 很爽，但很多維護災難也是從那裡開始。"
clickbait_en: "Using `extends` feels powerful, but many maintenance disasters start there too."
review_hint_zh: "composition 常比 inheritance 更鬆耦合。"
review_hint_en: "Composition is often less tightly coupled than inheritance."
enabled: true
---

## zh-TW

inheritance 讓你沿用父類別的結構與行為，但也可能把子類別綁得很死。
composition 則偏向把功能拆成可組合的部件，通常更容易替換、測試與重構。

## en

Inheritance lets you reuse a parent class structure and behavior, but it can also tightly bind subclasses to that hierarchy.
Composition instead favors assembling behavior from smaller pieces, which is often easier to swap, test, and refactor.
