---
id: linux-xargs-build-arguments
title_zh: xargs 組參數
title_en: Linux xargs
type: true-false
body_format: bilingual-section
tags: [linux, shell, xargs]
difficulty: 3
question_zh: "`xargs` 常用來把標準輸入整理成另一個指令的參數列。"
question_en: "`xargs` is commonly used to turn standard input into arguments for another command."
answer: true
clickbait_zh: "有些輸出不是拿來看的，是拿來變成下一個指令的參數。"
clickbait_en: "Some output is not meant to be read. It is meant to become the next command's arguments."
review_hint_zh: "`xargs` 會把輸入轉成命令列參數。"
review_hint_en: "`xargs` converts input into command-line arguments."
confusion_with: [linux-pipe-command-chain, linux-find-search-files]
metaphor_seed: [代辦清單, 排成參數, 接力]
hook_style_tags: [comparison, misunderstood]
enabled: true
---

## zh-TW

`xargs` 會把標準輸入拆成一組組參數，再交給另一個指令使用。
這在你先用某個工具列出一串目標，再想把它們當成下一個指令參數時很有用。
它不是單純把輸出「接」過去，而是重組成命令列參數。

## en

`xargs` takes standard input and builds command-line arguments from it.
It is useful when one command lists targets and another command expects those targets as explicit arguments.
It does more than pass output along; it reshapes the input into an argument list.
