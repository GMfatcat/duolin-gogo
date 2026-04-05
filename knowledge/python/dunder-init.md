---
id: python-dunder-init-constructor
title_zh: __init__
title_en: Python __init__
type: true-false
body_format: bilingual-section
tags: [python, class, init]
difficulty: 2
question_zh: "`__init__` 常用來在建立物件後初始化它的狀態。"
question_en: "`__init__` is commonly used to initialize an object's state after it is created."
answer: true
clickbait_zh: "物件不是一生出來就自動完整，很多時候第一步是先把自己整理好。"
clickbait_en: "Objects are not born complete. Their first step is often to get themselves into a usable state."
review_hint_zh: "`__init__` 主要負責初始化物件狀態。"
review_hint_en: "`__init__` mainly initializes object state."
confusion_with: [python-class-blueprint-object]
metaphor_seed: [起手整理, 初始化, 開場設定]
hook_style_tags: [safe, misunderstood]
enabled: true
---

## zh-TW

在 Python class 裡，`__init__` 通常會在物件建立後被呼叫，用來設定屬性或初始狀態。
它不是拿來回傳新物件本身，而是讓新建立的物件變得可用。
所以常見心智模型是「建立完成後的初始化步驟」。

## en

In a Python class, `__init__` is usually called after an object is created so its attributes and initial state can be set up.
It is not mainly for returning the object itself.
A useful mental model is "post-creation initialization."
