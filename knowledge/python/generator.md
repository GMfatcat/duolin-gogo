---
id: python-generator-lazy-iteration
title_zh: generator 延後產生資料
title_en: Generator
type: true-false
body_format: bilingual-section
tags: [python, iteration, memory]
difficulty: 2
question_zh: "generator 常用來逐步產生資料，而不是一次把所有結果都建立出來。"
question_en: "A generator is commonly used to produce values step by step instead of creating all results at once."
answer: true
clickbait_zh: "不是每次都要先把整份資料搬進記憶體，高手常常只在需要時才拿下一筆。"
clickbait_en: "You do not always need to load everything into memory first. Often the better move is to produce the next value only when needed."
review_hint_zh: "generator = lazy 產生資料。"
review_hint_en: "A generator produces values lazily."
confusion_with: [python-list-vs-tuple]
metaphor_seed: [一筆一筆吐出, 延後供應, 邊走邊給]
hook_style_tags: [safer_first, misunderstood]
enabled: true
---

## zh-TW

generator 會在需要時才逐步產生下一個值，而不是一開始就把所有結果都建立好。
這種延後產生資料的方式，常有助於節省記憶體，也很適合處理大型序列或串流資料。
你可以把它理解成「邊走邊產生」，不是「先全部準備好」。

## en

A generator produces the next value only when it is needed instead of building all results up front.
That lazy style often helps save memory and is useful for large sequences or streamed data.
You can think of it as "produce while iterating" rather than "prepare everything first."
