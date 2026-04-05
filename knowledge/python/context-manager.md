---
id: python-context-manager-with
title_zh: with 與 context manager
title_en: Python Context Manager
type: true-false
body_format: bilingual-section
tags: [python, context-manager, resource]
difficulty: 3
question_zh: "`with` 常和 context manager 一起用來安全管理資源的開啟與釋放。"
question_en: "`with` is commonly used with a context manager to safely manage opening and releasing resources."
answer: true
clickbait_zh: "你還沒手動關檔案，它其實已經幫你把收尾規矩先安排好了。"
clickbait_en: "Before you manually close anything, the cleanup contract is already in place."
review_hint_zh: "`with` 常用在檔案、鎖、連線等需要收尾的資源。"
review_hint_en: "`with` is often used for files, locks, and connections that need cleanup."
confusion_with: [python-exception-try-except]
metaphor_seed: [進場退場, 收尾, 規矩]
hook_style_tags: [misunderstood, safer_first]
enabled: true
---

## zh-TW

Python 的 context manager 會在進入和離開某段程式時自動處理設定與清理。
最常見的寫法就是 `with`。
它很適合拿來管理檔案、鎖或連線，避免忘記做收尾。

## en

A Python context manager handles setup and cleanup automatically around a block of code.
The most common syntax for it is `with`.
It is especially useful for files, locks, or connections where cleanup should never be forgotten.
