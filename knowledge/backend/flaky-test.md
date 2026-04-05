---
id: backend-flaky-test-instability
title_zh: Flaky Test
title_en: Flaky Test
type: true-false
body_format: bilingual-section
tags: [backend, testing, flaky]
difficulty: 2
question_zh: "flaky test 指的是在程式沒改的情況下，有時通過、有時失敗的測試。"
question_en: "A flaky test is a test that sometimes passes and sometimes fails even when the code has not changed."
answer: true
clickbait_zh: "最難受的不是測試紅，而是它今天紅、明天又自己綠。"
clickbait_en: "The worst test is not the red one. It is the one that changes its mind by itself."
review_hint_zh: "flaky = 不穩定、結果不可靠。"
review_hint_en: "Flaky means unstable and unreliable."
enabled: true
---

## zh-TW

flaky test 的問題不只是煩人，而是它會削弱團隊對測試結果的信任。
常見來源包括時間依賴、競態條件、外部服務不穩，或測試彼此汙染。

## en

The problem with a flaky test is not only annoyance; it erodes trust in the test suite.
Common sources include time dependence, race conditions, unstable external services, or tests polluting each other.
