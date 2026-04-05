---
id: backend-async-backpressure-pressure
title_zh: Async Backpressure
title_en: Async Backpressure
type: true-false
body_format: bilingual-section
tags: [backend, async, backpressure]
difficulty: 3
question_zh: "非同步系統如果接收速度長期高於處理速度，就可能出現 backpressure 問題。"
question_en: "If an async system keeps receiving work faster than it can process it, backpressure can become a problem."
answer: true
clickbait_zh: "不是放進 queue 就結束了，真正的坑常常是後面根本消化不完。"
clickbait_en: "Putting work into a queue is not the end of the story. The real pit appears when the system cannot keep up."
review_hint_zh: "backpressure 本質是上游太快、下游太慢。"
review_hint_en: "Backpressure is fundamentally about upstream outrunning downstream."
enabled: true
---

## zh-TW

非同步架構很容易讓人只看到「先收下再處理」，但如果下游長期來不及消化，queue 就可能越堆越多。
這就是 backpressure 要面對的問題：系統吞吐不平衡。

## en

In async systems, it is easy to focus only on "accept now, process later."
But if downstream stays slower than incoming work, the queue can keep growing.
That is the kind of imbalance backpressure is about.
