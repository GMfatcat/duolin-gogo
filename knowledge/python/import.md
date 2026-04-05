---
id: python-import-module-loading
title_zh: import 載入模組
title_en: Python Import
type: true-false
body_format: bilingual-section
tags: [python, module, import]
difficulty: 1
question_zh: "Python 的 `import` 主要用來把其他模組的功能帶進目前程式。"
question_en: "In Python, `import` is mainly used to bring functionality from another module into the current program."
answer: true
clickbait_zh: "你不是把程式碼複製進來，只是向另一個模組借了入口。"
clickbait_en: "You are not copying the whole file in. You are borrowing an entry point into another module."
review_hint_zh: "`import` 會讓你使用其他模組中的名稱。"
review_hint_en: "`import` lets your code use names from another module."
confusion_with: [python-class-blueprint-object]
metaphor_seed: [借入口, 接線, 引進]
hook_style_tags: [safer_first, misunderstood]
enabled: true
---

## zh-TW

`import` 讓 Python 程式可以使用其他模組提供的函式、類別或常數。
它幫你把模組接進目前檔案的作用範圍。
理解 import 很重要，因為大型 Python 專案幾乎都依賴模組拆分。

## en

`import` allows a Python program to use functions, classes, or constants from another module.
It connects that module into the current file's namespace.
Understanding imports matters because larger Python projects rely heavily on modular structure.
