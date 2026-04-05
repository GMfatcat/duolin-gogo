---
id: linux-pwd-current-directory
title_zh: pwd 目前位置
title_en: PWD
type: true-false
body_format: bilingual-section
tags: [linux, shell, navigation]
difficulty: 1
question_zh: "`pwd` 會顯示你目前所在的工作目錄路徑。"
question_en: "`pwd` shows the path of your current working directory."
answer: true
clickbait_zh: "你以為指令壞了？很多時候只是你根本不在自己以為的位置。"
clickbait_en: "Think the command is broken? Sometimes you are simply not where you think you are."
review_hint_zh: "`pwd` = print working directory。"
review_hint_en: "`pwd` means print working directory."
confusion_with: [linux-ls-list-files]
metaphor_seed: [現在位置, 你站在哪裡]
hook_style_tags: [safer_first, misunderstood]
enabled: true
---

## zh-TW

`pwd` 代表 `print working directory`，會印出你目前所在資料夾的完整路徑。
當你不確定自己在哪個目錄，或腳本行為看起來怪怪的時候，先跑一次 `pwd` 常常很有幫助。
它是 shell 導航裡最基本的定位工具之一。

## en

`pwd` stands for `print working directory` and prints the full path of the directory you are currently in.
When you are not sure where you are, or a script behaves strangely, checking `pwd` is often a good first step.
It is one of the most basic navigation tools in the shell.
