---
id: bash-path-command-lookup
title_zh: PATH
title_en: PATH
type: true-false
body_format: bilingual-section
tags: [bash, path, shell]
difficulty: 1
question_zh: "`PATH` 決定 shell 在哪些資料夾裡找可執行檔。"
question_en: "`PATH` determines which directories the shell searches for executable commands."
answer: true
clickbait_zh: "同一個指令在你電腦能跑、別人電腦卻不行？常常就是 PATH 在作怪。"
clickbait_en: "A command works on your machine but not someone else's? PATH is often the difference."
review_hint_zh: "`PATH` 是 shell 搜尋指令的位置清單。"
review_hint_en: "`PATH` is the shell's search list for commands."
enabled: true
---

## zh-TW

當你輸入 `python` 或 `git` 這類指令時，shell 會依序到 `PATH` 裡列出的資料夾查找。
如果對應可執行檔不在這些位置裡，就可能出現 command not found。

## en

When you type commands like `python` or `git`, the shell searches directories listed in `PATH`.
If the executable is not found in those locations, the command may fail with command not found.
