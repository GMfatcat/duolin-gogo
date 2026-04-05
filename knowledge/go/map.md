---
id: go-map-key-value
title_zh: map 鍵值對
title_en: Go Map
type: single-choice
body_format: bilingual-section
tags: [go, data-structure, map]
difficulty: 1
question_zh: "Go 的 `map` 最適合拿來存什麼？"
question_en: "What is a Go `map` best suited for storing?"
choices_zh:
  - "一組依索引順序存放的固定長度資料"
  - "根據 key 取值的鍵值對資料"
  - "只用來跑 goroutine 的同步資訊"
choices_en:
  - "A fixed-length indexed sequence"
  - "Key-value data accessed by key"
  - "Only synchronization data for goroutines"
answer: 1
clickbait_zh: "不是所有資料都該硬排成清單，有些東西天生就該用 key 去找。"
clickbait_en: "Not every collection wants to be a list. Some data is naturally looked up by key."
review_hint_zh: "`map` 是 key-value 結構。"
review_hint_en: "`map` is a key-value structure."
confusion_with: [go-slice-dynamic-view]
metaphor_seed: [字典, 索引卡, 通訊錄]
hook_style_tags: [comparison, safer_first]
enabled: true
---

## zh-TW

Go 的 `map` 用來存放鍵值對資料，適合透過 key 快速查到對應值。
如果你的需求是「用名稱找設定」或「用 ID 找物件」，`map` 就很常見。
它和 slice 不一樣，重點不是順序，而是查找關係。

## en

A Go `map` stores key-value data and is useful when you want to retrieve values by key.
It fits cases like looking up a setting by name or finding an object by ID.
Unlike a slice, the main idea is lookup rather than order.
