---
id: oop-polymorphism-same-interface
title_zh: Polymorphism
title_en: Polymorphism
type: true-false
body_format: bilingual-section
tags: [oop, polymorphism, interface]
difficulty: 2
question_zh: "polymorphism 常表示不同物件可以透過相同介面被統一使用。"
question_en: "Polymorphism commonly means different objects can be used through the same interface."
answer: true
clickbait_zh: "外表看起來一樣，底下做的事其實可以完全不同。"
clickbait_en: "Things can look identical from the outside while doing very different work underneath."
review_hint_zh: "polymorphism 的重點是同介面、多實作。"
review_hint_en: "Polymorphism is about one interface with multiple implementations."
enabled: true
---

## zh-TW

多型常讓呼叫端只依賴抽象介面，而不是依賴某個具體類型。
這樣不同實作可以被替換，而呼叫端不需要大改。

## en

Polymorphism often lets calling code depend on an abstract interface rather than a single concrete type.
That makes it easier to swap implementations without changing every caller.
