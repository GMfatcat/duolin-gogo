---
id: backend-race-condition-shared-state
title_zh: Race Condition
title_en: Race Condition
type: true-false
body_format: bilingual-section
tags: [backend, concurrency, race-condition]
difficulty: 3
question_zh: "race condition 常見於多個執行流程同時碰共享狀態，而結果又依賴執行順序。"
question_en: "A race condition commonly appears when multiple execution flows touch shared state and the result depends on timing or ordering."
answer: true
clickbait_zh: "有些 bug 不是邏輯錯，而是誰先碰到資料這件事本身就不穩。"
clickbait_en: "Some bugs are not about wrong logic. They are about who touched shared state first."
review_hint_zh: "race condition 重點在共享狀態與不穩順序。"
review_hint_en: "Race conditions center on shared state and unstable ordering."
enabled: true
---

## zh-TW

當多個 goroutine、thread 或 worker 同時讀寫共享資料，而結果會隨順序改變時，就可能出現 race condition。
它常讓 bug 難以重現，因為時間一變，結果就變了。

## en

When multiple goroutines, threads, or workers read and write shared data and the outcome changes with ordering, a race condition can appear.
These bugs are often hard to reproduce because timing changes the behavior.
