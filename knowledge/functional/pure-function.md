---
id: functional-pure-function-no-side-effect
title_zh: Pure Function
title_en: Pure Function
type: true-false
body_format: bilingual-section
tags: [functional, pure-function, side-effect]
difficulty: 2
question_zh: "pure function 在相同輸入下應該回傳相同輸出，且不依賴隱藏外部狀態。"
question_en: "A pure function should return the same output for the same input and should not depend on hidden external state."
answer: true
clickbait_zh: "真正穩定的函式，不該偷偷看全域狀態或順手改點什麼。"
clickbait_en: "A truly stable function should not secretly read global state or quietly modify something on the side."
review_hint_zh: "pure function 重點是同輸入同輸出，且沒有副作用。"
review_hint_en: "A pure function means same input, same output, and no side effects."
enabled: true
---

## zh-TW

純函式的價值在於可預測、容易測試，也比較容易推理。
如果一個函式結果還會受時間、全域變數或外部 I/O 影響，它通常就不算純。

## en

Pure functions are valuable because they are predictable, easy to test, and easier to reason about.
If a function's result also depends on time, globals, or outside I/O, it is usually not pure.
