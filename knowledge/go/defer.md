---
id: go-defer-late-execution
title_zh: defer 延後執行
title_en: Defer
type: true-false
body_format: bilingual-section
tags: [go, functions, cleanup]
difficulty: 2
question_zh: "`defer` 會把函式呼叫延後到外層函式結束前再執行。"
question_en: "`defer` delays a function call until just before the surrounding function returns."
answer: true
clickbait_zh: "不是忘記關，而是你應該一開始就安排好最後誰來收尾。"
clickbait_en: "It is not just about remembering cleanup. It is about deciding early how the function should close out."
review_hint_zh: "`defer` 常用來做關閉或清理。"
review_hint_en: "`defer` is commonly used for cleanup or closing resources."
confusion_with: [go-goroutine-concurrency]
metaphor_seed: [收尾, 關門, 最後一步]
hook_style_tags: [safer_first, misunderstood]
enabled: true
---

## zh-TW

`defer` 會把某個函式呼叫延後到外層函式快要結束時再執行。
這很適合用在 `close`、`unlock`、或其他收尾清理工作。
它的好處是你可以在資源打開的地方，就先把收尾邏輯安排好。

## en

`defer` delays a function call until the surrounding function is about to return.
That makes it useful for `close`, `unlock`, or other cleanup tasks.
The benefit is that you can set up the cleanup right where the resource is opened.
