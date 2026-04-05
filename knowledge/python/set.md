---
id: python-set-unique-values
title_zh: set 唯一值集合
title_en: Python Set
type: single-choice
body_format: bilingual-section
tags: [python, data-structure, set]
difficulty: 1
question_zh: "Python 的 `set` 最適合處理哪種資料需求？"
question_en: "What kind of data need is a Python `set` best suited for?"
choices_zh:
  - "保留重複值與順序的列表"
  - "只保留唯一值的集合"
  - "用 key 查 value 的映射"
choices_en:
  - "A list that keeps duplicates and order"
  - "A collection of unique values"
  - "A mapping from key to value"
answer: 1
clickbait_zh: "你以為資料很多，其實只是重複很多。這時候 set 會一下子把真相翻出來。"
clickbait_en: "Sometimes your data is not large. It is just repetitive. A `set` reveals that fast."
review_hint_zh: "`set` 的核心是唯一值。"
review_hint_en: "The core idea of a `set` is uniqueness."
confusion_with: [python-list-vs-tuple, python-dict-key-value]
metaphor_seed: [去重, 篩重複, 集合箱]
hook_style_tags: [comparison, safer_first]
enabled: true
---

## zh-TW

`set` 是一種只保留唯一值的集合型別，很適合拿來去重或做成員檢查。
如果你不在意順序，但很在意「有沒有重複」，它通常就很有用。
它和 list、dict 的用途不同，核心是唯一性。

## en

A `set` is a collection type that keeps unique values only.
It is useful for removing duplicates or checking membership quickly.
If order is not your main concern but uniqueness is, a set is often a good choice.
