---
id: functional-side-effect-external-change
title_zh: Side Effect
title_en: Side Effect
type: true-false
body_format: bilingual-section
tags: [functional, side-effect]
difficulty: 2
question_zh: "副作用常指函式除了回傳值之外，還改變了外部可觀察狀態。"
question_en: "A side effect commonly means a function changes observable external state in addition to returning a value."
answer: true
clickbait_zh: "最難 debug 的不是回傳值，而是它順手在外面留下了什麼。"
clickbait_en: "The hardest thing to debug is often not the return value, but what the function left behind outside."
review_hint_zh: "副作用是對外部世界造成可觀察改變。"
review_hint_en: "A side effect is an observable change to the outside world."
enabled: true
---

## zh-TW

寫檔、改全域變數、送網路請求、改資料庫內容，這些都常被視為副作用。
副作用不是一定不好，但如果沒有被明確管理，程式會更難測試與推理。

## en

Writing files, changing global variables, sending network requests, or mutating a database are common examples of side effects.
Side effects are not always bad, but unmanaged side effects make code harder to test and reason about.
