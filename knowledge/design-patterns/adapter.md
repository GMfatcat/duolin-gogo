---
id: patterns-adapter-bridge-interface
title_zh: Adapter Pattern
title_en: Adapter Pattern
type: true-false
body_format: bilingual-section
tags: [design-patterns, adapter, integration]
difficulty: 2
question_zh: "adapter pattern 常用來把不相容的介面轉成呼叫端能接受的形式。"
question_en: "The adapter pattern is commonly used to convert an incompatible interface into one that the caller can work with."
answer: true
clickbait_zh: "不是每個外部套件都會照你的世界觀說話，所以常要有人當翻譯。"
clickbait_en: "Not every external library speaks your language, so something often has to translate."
review_hint_zh: "adapter 常像介面翻譯層。"
review_hint_en: "An adapter often acts like an interface translator."
enabled: true
---

## zh-TW

當你接外部 SDK、舊系統或第三方套件時，常會遇到它的呼叫方式跟你自己的抽象不一致。
adapter pattern 讓你在中間做一層轉換，避免整個系統都被外部介面綁住。

## en

When integrating external SDKs, legacy systems, or third-party libraries, their interface often does not match your own abstractions.
The adapter pattern adds a translation layer so the rest of your system does not become tightly coupled to that external API.
