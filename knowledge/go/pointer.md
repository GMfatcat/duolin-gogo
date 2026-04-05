---
id: go-pointer-memory-address
title_zh: pointer 記憶體位址參考
title_en: Go Pointer
type: true-false
body_format: bilingual-section
tags: [go, memory, pointer]
difficulty: 2
question_zh: "Go 的 pointer 可以讓你參考某個值所在的位置。"
question_en: "A pointer in Go lets you refer to the location of a value."
answer: true
clickbait_zh: "你改的是副本還是原本那份？很多 bug 都卡在這個分不清。"
clickbait_en: "Are you changing a copy or the original? A lot of bugs hide inside that confusion."
review_hint_zh: "pointer 代表某個值的位址參考。"
review_hint_en: "A pointer refers to the address of a value."
confusion_with: [go-struct-field-grouping, go-slice-dynamic-view]
metaphor_seed: [地址, 門牌, 代拿東西]
hook_style_tags: [misunderstood, comparison]
enabled: true
---

## zh-TW

pointer 讓你持有某個值的位址，透過它可以間接讀取或修改原本的值。
這和只操作副本不同，因為副本改了不一定會影響原始資料。
理解 pointer，對理解參數傳遞與資料修改很有幫助。

## en

A pointer lets you hold the address of a value so you can indirectly read or modify the original data.
That is different from working on a separate copy, which may not affect the original value at all.
Understanding pointers helps a lot when reasoning about parameter passing and mutation.
