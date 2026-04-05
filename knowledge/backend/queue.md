---
id: backend-queue-async-work-buffer
title_zh: Queue
title_en: Backend Queue
type: true-false
body_format: bilingual-section
tags: [backend, queue, async]
difficulty: 2
question_zh: "queue 常被用來緩衝非同步工作，讓生產者與消費者不必完全同步。"
question_en: "A queue is often used to buffer async work so producers and consumers do not have to stay perfectly synchronized."
answer: true
clickbait_zh: "不是所有工作都要當下做完，有些系統的祕密其實只是先排隊。"
clickbait_en: "Not every job has to finish immediately. Some systems survive simply by making work wait in line."
review_hint_zh: "queue 能緩衝 async 工作。"
review_hint_en: "A queue buffers asynchronous work."
enabled: true
---

## zh-TW

queue 常見於非同步設計，用來先收下工作，再讓後面的 worker 慢慢處理。
它的價值之一，是把生產速度和處理速度解耦，但也可能帶來堆積與 backpressure 問題。

## en

A queue is common in async design: work is accepted first and processed later by workers.
One of its main values is decoupling production speed from processing speed, though it can also create backlog and backpressure issues.
