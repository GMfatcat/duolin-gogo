---
id: linux-chmod-permissions
title_zh: chmod 修改權限
title_en: Chmod
type: true-false
body_format: bilingual-section
tags: [linux, shell, permissions]
difficulty: 2
question_zh: "`chmod` 常用來修改檔案或資料夾的權限。"
question_en: "`chmod` is commonly used to change file or directory permissions."
answer: true
clickbait_zh: "不能執行不一定是程式壞了，可能只是你沒給它權限。"
clickbait_en: "If something will not run, the problem may be permission rather than the program itself."
review_hint_zh: "`chmod` = change mode，常用來改權限。"
review_hint_en: "`chmod` means change mode and is commonly used for permissions."
confusion_with: [linux-ls-list-files]
metaphor_seed: [門禁, 鑰匙, 開放權限]
hook_style_tags: [fear_of_mistake, misunderstood]
enabled: true
---

## zh-TW

`chmod` 代表 `change mode`，常用來修改檔案或資料夾的權限。
例如你可能會用它讓腳本變得可執行，或限制誰可以讀寫某個檔案。
在 Linux 裡，權限常常直接影響指令能不能成功執行。

## en

`chmod` stands for `change mode` and is commonly used to modify file or directory permissions.
For example, you might use it to make a script executable or to restrict who can read or write a file.
In Linux, permissions often determine whether a command succeeds at all.
