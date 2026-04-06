---
id: patterns-strategy-swap-behavior
title_zh: Strategy Pattern
title_en: Strategy Pattern
type: true-false
body_format: bilingual-section
tags: [design-patterns, strategy, behavior]
difficulty: 2
question_zh: "strategy pattern 常用來在同一組介面下切換不同演算法或行為。"
question_en: "The strategy pattern is commonly used to switch between different algorithms or behaviors under the same interface."
answer: true
clickbait_zh: "有時候你不是在換資料，而是在換『做這件事的方法』。"
clickbait_en: "Sometimes you are not swapping the data. You are swapping the way the job gets done."
review_hint_zh: "strategy 的重點是可替換的行為。"
review_hint_en: "Strategy is about swappable behavior."
enabled: true
---

## zh-TW

當某個流程有多種可替換實作，例如不同排序、不同付款方式、不同推薦規則時，strategy pattern 很有用。
它讓呼叫端面對的是同一個抽象介面，而不是大量的 if/else 分支。

## en

When one flow can have multiple interchangeable implementations, such as different sorting, payment, or recommendation logic, the strategy pattern is useful.
It lets callers depend on one abstraction instead of a large pile of if/else branches.
