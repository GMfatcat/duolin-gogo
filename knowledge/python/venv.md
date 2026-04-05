---
id: python-venv-isolation
title_zh: venv 虛擬環境
title_en: Python Venv
type: true-false
body_format: bilingual-section
tags: [python, environment, tooling]
difficulty: 1
question_zh: "`venv` 常用來建立彼此隔離的 Python 套件環境。"
question_en: "`venv` is commonly used to create isolated Python package environments."
answer: true
clickbait_zh: "你以為套件壞掉，其實可能只是你把不同專案的環境全混在一起了。"
clickbait_en: "Think the package setup is broken? Sometimes the real issue is that multiple projects share the same environment."
review_hint_zh: "`venv` = 隔離不同專案的 Python 環境。"
review_hint_en: "`venv` isolates Python environments between projects."
confusion_with: [python-dict-key-value]
metaphor_seed: [獨立房間, 隔離艙, 分開管理]
hook_style_tags: [safer_first, misunderstood]
enabled: true
---

## zh-TW

`venv` 常用來替不同 Python 專案建立彼此隔離的套件環境。
這樣每個專案可以有自己的依賴版本，不會互相污染。
它是避免「這台機器明明昨天可以跑，今天卻壞掉」的重要工具。

## en

`venv` is commonly used to create isolated package environments for different Python projects.
That lets each project keep its own dependency versions without polluting others.
It is one of the key tools for avoiding "it worked yesterday, but now the environment is broken" problems.
