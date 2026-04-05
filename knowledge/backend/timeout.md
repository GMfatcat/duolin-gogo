---
id: backend-timeout-fail-fast-guard
title_zh: Timeout
title_en: Timeout
type: true-false
body_format: bilingual-section
tags: [backend, timeout, reliability]
difficulty: 2
question_zh: "timeout 常被用來避免請求或工作無限等待。"
question_en: "A timeout is commonly used to prevent a request or job from waiting forever."
answer: true
clickbait_zh: "不是每件事都值得一直等下去，系統有時候更需要的是知道該停。"
clickbait_en: "Not everything is worth waiting for forever. Sometimes the system needs to know when to stop."
review_hint_zh: "timeout 是 fail-fast 保護。"
review_hint_en: "A timeout is a fail-fast safety guard."
enabled: true
---

## zh-TW

timeout 的目標之一，是避免某個外部依賴、請求或工作把整個流程拖住太久。
它不是萬靈丹，但對可靠性與資源控制很重要。

## en

One purpose of a timeout is to stop an external dependency, request, or job from stalling the whole flow for too long.
It is not a cure-all, but it matters a lot for reliability and resource control.
