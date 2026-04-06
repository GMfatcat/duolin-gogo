---
id: functional-immutability-stable-data
title_zh: Immutability
title_en: Immutability
type: true-false
body_format: bilingual-section
tags: [functional, immutability, data]
difficulty: 2
question_zh: "immutability 常表示資料建立後不直接修改，而是產生新的值。"
question_en: "Immutability commonly means data is not changed in place after creation, but new values are produced instead."
answer: true
clickbait_zh: "你以為這樣會很麻煩，但很多 bug 正是因為東西被偷偷改掉。"
clickbait_en: "It can sound inconvenient, but many bugs start because something was quietly mutated."
review_hint_zh: "immutability 偏向不原地修改。"
review_hint_en: "Immutability favors not changing values in place."
enabled: true
---

## zh-TW

不可變資料有助於降低共享狀態造成的混亂，特別是在非同步或多執行緒情境裡。
代價通常是需要建立新值，但回報是更清楚的資料流與更少的意外修改。

## en

Immutable data helps reduce confusion caused by shared mutable state, especially in async or concurrent settings.
The tradeoff is creating new values more often, but the payoff is clearer data flow and fewer accidental mutations.
