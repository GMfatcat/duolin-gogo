---
id: python-venv-pip-install-flow
title_zh: venv 與 pip
title_en: Python venv and pip
type: true-false
body_format: bilingual-section
tags: [python, venv, pip]
difficulty: 2
question_zh: "在專案裡先進入 venv，再用 pip 安裝套件，通常能減少全域套件污染。"
question_en: "Activating a venv before installing packages with pip usually reduces global package pollution."
answer: true
clickbait_zh: "你不是一定裝錯套件，有時候只是裝到了錯的世界。"
clickbait_en: "Sometimes the package is fine. It just got installed into the wrong world."
review_hint_zh: "venv 隔離環境，pip 會裝到目前啟用的環境。"
review_hint_en: "A venv isolates the environment, and pip installs into the active environment."
confusion_with: [python-venv-isolation, python-import-module-loading]
metaphor_seed: [錯的世界, 隔離房間, 裝到哪裡]
hook_style_tags: [fear_of_mistake, safer_first]
enabled: true
---

## zh-TW

`venv` 的作用是建立隔離的 Python 執行環境，避免不同專案共用一堆彼此衝突的套件版本。
當你啟用某個 venv 後，再用 `pip` 安裝套件，通常就會裝進那個環境，而不是全域 Python。
這是 Python 專案管理中很基本也很重要的習慣。

## en

`venv` creates an isolated Python environment so different projects do not have to fight over the same global package versions.
After you activate a venv, running `pip install` usually installs packages into that environment instead of the global Python setup.
That is one of the most basic and important Python workflow habits.
