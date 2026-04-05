---
id: python-dict-key-value
title_zh: dict 鍵值對
title_en: Dict
type: single-choice
body_format: bilingual-section
tags: [python, collections, mapping]
difficulty: 1
question_zh: "Python 的 `dict` 最主要適合用來表示什麼？"
question_en: "What is a Python `dict` mainly used to represent?"
choices_zh:
  - "依照 key 查找 value 的鍵值對資料"
  - "只用來做數學矩陣計算"
  - "專門代表函式回傳值"
choices_en:
  - "Key-value data where values are looked up by key"
  - "Something used only for matrix calculations"
  - "A structure dedicated to function return values"
answer: 0
clickbait_zh: "資料不是越多越難找，真正關鍵是你有沒有給它一把好鑰匙。"
clickbait_en: "The issue is not always too much data. Sometimes it is whether you gave yourself a good key."
review_hint_zh: "`dict` = key-value mapping。"
review_hint_en: "`dict` is a key-value mapping."
confusion_with: [python-list-vs-tuple]
metaphor_seed: [字典, 鑰匙, 對照表]
hook_style_tags: [safer_first, comparison]
enabled: true
---

## zh-TW

`dict` 是 Python 用來表示 key-value 對應關係的常用資料結構。
你可以透過 key 快速找到對應的 value，很適合拿來表示設定、欄位資料或索引。
它不是靠位置取值，而是靠 key。

## en

`dict` is one of Python's standard data structures for representing key-value mappings.
You use a key to look up the corresponding value, which makes it useful for configuration, structured fields, or indexing.
It is based on keys rather than positional access.
