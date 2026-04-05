---
id: backend-sync-vs-async-flow
title_zh: Sync vs Async
title_en: Sync vs Async
type: true-false
body_format: bilingual-section
tags: [backend, sync, async]
difficulty: 2
question_zh: "sync 與 async 的差異之一，在於呼叫端是否需要當下等到結果才繼續。"
question_en: "One difference between sync and async is whether the caller has to wait for the result before continuing."
answer: true
clickbait_zh: "不是把東西丟到背景就叫比較厲害，先問自己要不要立刻拿到結果。"
clickbait_en: "Throwing work into the background is not automatically smarter. First ask whether you need the result right now."
review_hint_zh: "sync 常立即等待；async 常延後完成。"
review_hint_en: "Sync usually waits now; async usually finishes later."
enabled: true
---

## zh-TW

同步流程通常表示呼叫端要先等結果回來，才能往下走。
非同步流程則允許工作延後完成，呼叫端先做別的事或之後再收結果。

## en

Synchronous flow usually means the caller waits for the result before continuing.
Asynchronous flow allows the work to finish later so the caller can move on or collect the result afterward.
