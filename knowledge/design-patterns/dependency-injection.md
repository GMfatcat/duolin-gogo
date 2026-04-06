---
id: patterns-dependency-injection-decouple
title_zh: Dependency Injection
title_en: Dependency Injection
type: true-false
body_format: bilingual-section
tags: [design-patterns, dependency-injection, architecture]
difficulty: 2
question_zh: "dependency injection 常用來把依賴從外部注入，而不是在類別或函式內部直接 new 出來。"
question_en: "Dependency injection commonly means dependencies are supplied from the outside instead of being constructed directly inside a class or function."
answer: true
clickbait_zh: "問題常不是你依賴太多，而是你把依賴寫死在裡面。"
clickbait_en: "The problem is often not having dependencies. It is hard-coding them where they should not be."
review_hint_zh: "DI 的重點是外部提供依賴。"
review_hint_en: "The point of DI is that dependencies come from outside."
enabled: true
---

## zh-TW

依賴注入讓物件專注在使用某個能力，而不是自己決定要怎麼建立那個能力。
這通常有助於測試、替換實作，以及降低模組間耦合。

## en

Dependency injection lets an object focus on using a capability instead of deciding how to construct it.
That usually helps testing, implementation swapping, and reducing coupling between modules.
