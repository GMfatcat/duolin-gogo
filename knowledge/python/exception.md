---
id: python-exception-try-except
title_zh: try / except 例外處理
title_en: Python Exception Handling
type: single-choice
body_format: bilingual-section
tags: [python, error-handling, exception]
difficulty: 2
question_zh: "Python 的 `try / except` 最主要用來做什麼？"
question_en: "What is `try / except` mainly used for in Python?"
choices_zh:
  - "攔住並處理執行時發生的例外"
  - "宣告新的 class 繼承"
  - "建立匿名函式"
choices_en:
  - "Catch and handle runtime exceptions"
  - "Declare class inheritance"
  - "Create an anonymous function"
answer: 0
clickbait_zh: "程式炸掉不一定是終點，有時候重點是你有沒有先安排好接住它。"
clickbait_en: "A crash is not always the end. Sometimes the real question is whether you prepared to catch it."
review_hint_zh: "`try / except` 讓你處理例外，不是直接讓程式中斷。"
review_hint_en: "`try / except` lets you handle exceptions instead of crashing immediately."
confusion_with: [python-decorator-wrap-function]
metaphor_seed: [安全網, 接球, 緩衝墊]
hook_style_tags: [safer_first, misunderstood]
enabled: true
---

## zh-TW

`try / except` 用來包住可能出錯的程式區段，並在例外發生時提供處理方式。
這讓你的程式在遇到某些錯誤時，不一定要整個中斷。
它的重點不是忽略錯誤，而是有意識地接住並處理。

## en

`try / except` wraps code that might fail and gives you a way to handle exceptions when they occur.
That means your program does not always need to stop entirely when an error appears.
The goal is not to ignore errors, but to catch and handle them deliberately.
