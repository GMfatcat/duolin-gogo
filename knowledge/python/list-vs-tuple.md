---
id: python-list-vs-tuple
title_zh: list 與 tuple 差異
title_en: List vs Tuple
type: single-choice
body_format: bilingual-section
tags: [python, collections, basics]
difficulty: 1
question_zh: "哪一個敘述最符合 Python 中 list 和 tuple 的差異？"
question_en: "Which statement best describes the difference between a list and a tuple in Python?"
choices_zh:
  - "list 通常可變，tuple 通常不可變"
  - "tuple 只能存數字，list 只能存字串"
  - "它們完全等價，只是寫法不同"
choices_en:
  - "Lists are usually mutable, while tuples are usually immutable"
  - "Tuples can only store numbers, while lists can only store strings"
  - "They are completely equivalent and only differ in syntax"
answer: 0
clickbait_zh: "看起來只差括號，很多初學者卻在這裡踩了第一個資料結構坑。"
clickbait_en: "They may look like a tiny syntax difference, but this is where many beginners hit an early data-structure trap."
review_hint_zh: "list 常可改，tuple 通常不可改。"
review_hint_en: "Lists are commonly mutable; tuples are usually immutable."
confusion_with: [go-slice-dynamic-view]
metaphor_seed: [可改清單, 固定清單, 容器性格]
hook_style_tags: [comparison, misunderstood]
enabled: true
---

## zh-TW

在 Python 裡，list 和 tuple 都能存放多個值，但它們最常見的差異是可變性。
list 通常可以新增、刪除、修改元素；tuple 則通常建立後就不改動。
這讓 tuple 更適合表達固定結構，list 更適合表達會變動的集合。

## en

In Python, both lists and tuples can store multiple values, but their most common difference is mutability.
A list can usually be changed by adding, removing, or updating elements, while a tuple usually stays fixed after creation.
That makes tuples useful for fixed structures and lists useful for collections that change.
