---
id: backend-backpressure-slow-consumer
title_zh: Backpressure
title_en: Backpressure
type: true-false
body_format: bilingual-section
tags: [backend, async, backpressure]
difficulty: 3
question_zh: "backpressure 常出現在生產速度比消費速度更快時。"
question_en: "Backpressure commonly appears when production is faster than consumption."
answer: true
clickbait_zh: "系統不是只有撐不住才會爆。更多時候，是工作堆太快卻處理不完。"
clickbait_en: "Systems do not fail only when they crash. Often they drown because work arrives faster than it can be drained."
review_hint_zh: "backpressure = 上游產生太快，下游處理跟不上。"
review_hint_en: "Backpressure means upstream producers are outpacing downstream consumers."
enabled: true
---

## zh-TW

backpressure 指的是資料或工作進來的速度比系統能處理的速度還快。
如果沒有節流、排隊上限或丟棄策略，佇列會越積越多，延遲和記憶體壓力也會跟著上升。

## en

Backpressure means work arrives faster than the system can process it.
Without throttling, queue limits, or drop strategies, backlog grows and both latency and memory pressure rise with it.
