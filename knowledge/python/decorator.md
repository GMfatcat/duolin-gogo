---
id: python-decorator-wrap-function
title_zh: decorator 包裝函式
title_en: Python Decorator
type: true-false
body_format: bilingual-section
tags: [python, function, decorator]
difficulty: 3
question_zh: "Python 的 decorator 常用來在不直接改函式內容時，額外包裝函式行為。"
question_en: "A Python decorator is commonly used to wrap function behavior without rewriting the function body directly."
answer: true
clickbait_zh: "你沒有改函式本體，行為卻變了。這種魔法很多時候就是 decorator。"
clickbait_en: "The function body looks untouched, but the behavior changes anyway. That kind of magic is often a decorator."
review_hint_zh: "decorator 會包一層新行為在函式外面。"
review_hint_en: "A decorator wraps extra behavior around a function."
confusion_with: [python-lambda-anonymous-function]
metaphor_seed: [外套, 包裝紙, 濾鏡]
hook_style_tags: [misunderstood, comparison]
enabled: true
---

## zh-TW

decorator 是 Python 中一種包裝函式的方式，可以在呼叫前後加上額外邏輯，而不用直接改函式內文。
常見用途像是 logging、權限檢查、快取等。
它的重點是「包住原本的函式」，不是只是寫一個更短的函式。

## en

A decorator in Python wraps a function so extra logic can run before or after the original function.
Common uses include logging, permission checks, and caching.
The key idea is wrapping existing behavior, not simply writing a shorter function.
