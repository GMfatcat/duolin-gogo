---
id: backend-integration-test-boundary
title_zh: Integration Test
title_en: Integration Test
type: true-false
body_format: bilingual-section
tags: [backend, testing, integration-test]
difficulty: 2
question_zh: "integration test 常用來確認多個模組或外部依賴真的能一起工作。"
question_en: "An integration test is commonly used to confirm that multiple modules or external dependencies work together."
answer: true
clickbait_zh: "單元都沒壞，不代表它們接在一起時不會出事。"
clickbait_en: "Even if the units look fine, the seams between them can still break."
review_hint_zh: "integration test 看的是模組之間的接縫。"
review_hint_en: "Integration tests focus on the seams between modules."
enabled: true
---

## zh-TW

integration test 不只看單一函式，而是看多個元件、服務或資料庫互動時有沒有正常協作。
它通常比 unit test 慢，但能抓到接縫問題。

## en

An integration test looks beyond one function and checks whether multiple components, services, or a database interact correctly.
It is often slower than a unit test, but it catches seam-related failures.
