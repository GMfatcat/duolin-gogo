---
id: backend-stub-fixed-response-double
title_zh: Stub
title_en: Backend Stub
type: true-false
body_format: bilingual-section
tags: [backend, testing, stub]
difficulty: 2
question_zh: "stub 常用來提供固定輸出，好讓測試不必依賴真實外部系統。"
question_en: "A stub is commonly used to return fixed outputs so tests do not need a real external system."
answer: true
clickbait_zh: "有時候測試不是要驗證真世界，而是先把世界固定住。"
clickbait_en: "Sometimes a test is not trying to validate the real world. It is trying to freeze the world first."
review_hint_zh: "stub 偏固定回應，不一定驗證互動。"
review_hint_en: "A stub is more about fixed responses than interaction verification."
enabled: true
---

## zh-TW

stub 也是 test double，但它通常偏向提供固定回應，讓測試可以在可預期條件下運行。
和 mock 相比，它通常比較少驗證互動細節。

## en

A stub is also a test double, but it usually focuses on returning fixed responses so a test can run in predictable conditions.
Compared with a mock, it often cares less about verifying interaction details.
