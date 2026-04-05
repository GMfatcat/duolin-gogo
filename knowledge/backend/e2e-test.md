---
id: backend-e2e-test-full-journey
title_zh: E2E Test
title_en: End-to-End Test
type: true-false
body_format: bilingual-section
tags: [backend, testing, e2e]
difficulty: 2
question_zh: "end-to-end test 常試圖驗證一整段真實使用流程，而不只是單一函式或單一模組。"
question_en: "An end-to-end test commonly tries to validate a complete user flow rather than just one function or module."
answer: true
clickbait_zh: "單元都過了還是不放心？那就把整條路真的走一遍。"
clickbait_en: "If unit tests still leave doubt, run the whole journey from start to finish."
review_hint_zh: "E2E 看的是完整流程。"
review_hint_en: "E2E focuses on the complete journey."
enabled: true
---

## zh-TW

E2E 測試會盡量模擬接近真實使用的完整流程，例如從請求進來、服務處理，到結果產出。
它通常比 unit test 慢，但更接近真實情境。

## en

An end-to-end test tries to simulate a realistic full flow, such as a request entering the system, the service handling it, and the final result being produced.
It is usually slower than a unit test, but closer to real usage.
