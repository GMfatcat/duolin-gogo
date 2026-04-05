---
id: python-iterator-next-protocol
title_zh: Python iterator
title_en: Python Iterator
type: true-false
body_format: bilingual-section
tags: [python, iterator, iteration]
difficulty: 2
question_zh: "Python iterator 的核心，是能逐步提供下一個值，而不是一次產出全部結果。"
question_en: "The core idea of a Python iterator is to provide the next value step by step instead of producing everything at once."
answer: true
clickbait_zh: "不是每次迭代都要先把整包東西攤開，有時候下一個值就夠了。"
clickbait_en: "You do not always need the whole collection up front. Sometimes the next value is enough."
review_hint_zh: "iterator 會一個一個提供值。"
review_hint_en: "An iterator yields values one at a time."
confusion_with: [python-generator-lazy-iteration]
metaphor_seed: [下一個, 逐步供應, 不一次攤開]
hook_style_tags: [comparison, misunderstood]
enabled: true
---

## zh-TW

iterator 的重點是它能在需要時提供下一個值，而不是先建立整份結果。
這讓它很適合處理大量資料或逐步消費資料流。
在 Python 裡，`for` 迴圈背後就大量依賴 iterator protocol。

## en

An iterator provides the next value when asked instead of building the full result in advance.
That makes it useful for large datasets or streaming-style consumption.
In Python, `for` loops rely heavily on the iterator protocol behind the scenes.
