---
id: oop-encapsulation-hide-details
title_zh: Encapsulation
title_en: Encapsulation
type: true-false
body_format: bilingual-section
tags: [oop, encapsulation, design]
difficulty: 2
question_zh: "encapsulation 常強調把內部狀態與操作收在同一個物件邊界內。"
question_en: "Encapsulation commonly emphasizes keeping state and related behavior inside the same object boundary."
answer: true
clickbait_zh: "不是所有欄位都該裸奔給外面直接改。"
clickbait_en: "Not every field should be left running around for the outside world to mutate directly."
review_hint_zh: "encapsulation 重點是把資料與操作包在邊界裡。"
review_hint_en: "Encapsulation is about keeping data and behavior behind a boundary."
enabled: true
---

## zh-TW

encapsulation 的核心想法是：物件應該自己管理自己的狀態，而不是讓外部到處直接碰內部細節。
這能降低耦合，也比較容易維持不變條件與一致性。

## en

The core idea of encapsulation is that an object should manage its own state rather than letting outside code touch every internal detail directly.
This helps reduce coupling and makes it easier to preserve invariants and consistency.
