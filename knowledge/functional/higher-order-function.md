---
id: functional-higher-order-function
title_zh: Higher-Order Function
title_en: Higher-Order Function
type: true-false
body_format: bilingual-section
tags: [functional, higher-order-function, function]
difficulty: 2
question_zh: "higher-order function 常指能接收函式作為參數，或回傳函式的函式。"
question_en: "A higher-order function commonly means a function that takes another function as input or returns one as output."
answer: true
clickbait_zh: "有些函式不處理資料本身，而是專門操控別的函式。"
clickbait_en: "Some functions do not primarily manipulate data. They manipulate other functions."
review_hint_zh: "higher-order function 會操作函式本身。"
review_hint_en: "A higher-order function works with functions themselves."
enabled: true
---

## zh-TW

像 `map`、`filter` 這種工具通常都屬於 higher-order function，因為你會把一個函式傳給它。
這種設計能讓抽象行為被重用，而不是每次都重寫一整段流程。

## en

Tools such as `map` and `filter` are usually higher-order functions because you pass another function into them.
That pattern helps reuse abstract behavior instead of rewriting the whole flow each time.
