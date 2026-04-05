---
id: go-slice-dynamic-view
title_zh: slice 動態視圖
title_en: Slice
type: single-choice
body_format: bilingual-section
tags: [go, collections, memory]
difficulty: 2
question_zh: "在 Go 裡，slice 比較接近哪種概念？"
question_en: "Which idea best matches a slice in Go?"
choices_zh:
  - "固定長度、完全獨立的陣列本體"
  - "指向底層陣列的一個可變長度視圖"
  - "只能存字串的特別容器"
choices_en:
  - "A fixed-length, fully independent array value"
  - "A variable-length view over an underlying array"
  - "A special container that can only store strings"
answer: 1
clickbait_zh: "很多 Go 初學者以為 slice 就是 array，後面才發現資料為什麼一起變。"
clickbait_en: "Many Go beginners assume a slice is just an array, then wonder why related data changes together."
review_hint_zh: "slice 是對底層 array 的視圖，不是單純另一份複製。"
review_hint_en: "A slice is a view over an underlying array, not just a separate copy."
confusion_with: [python-list-vs-tuple]
metaphor_seed: [視窗, 切片, 借用同一塊底層]
hook_style_tags: [misunderstood, fear_of_mistake]
enabled: true
---

## zh-TW

slice 是 Go 很常用的集合型別，但它本質上比較像對底層 array 的一段視圖。
它有長度與容量的概念，而且可以在一定範圍內動態成長。
理解 slice 不是單純獨立陣列，有助於理解 append、共享底層資料等行為。

## en

A slice is one of the most common collection types in Go, but conceptually it is closer to a view over part of an underlying array.
It has both length and capacity, and it can grow within certain limits.
Understanding that a slice is not just a fully separate array helps explain behaviors such as append and shared backing data.
