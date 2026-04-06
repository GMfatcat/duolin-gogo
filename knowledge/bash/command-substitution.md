---
id: bash-command-substitution
title_zh: Command Substitution
title_en: Command Substitution
type: true-false
body_format: bilingual-section
tags: [bash, shell, substitution]
difficulty: 2
question_zh: "`$(...)` 常用來把一個指令的輸出結果嵌進另一個指令裡。"
question_en: "`$(...)` is commonly used to embed the output of one command into another command."
answer: true
clickbait_zh: "有些 shell 指令看起來像在寫巢狀咒語，其實只是把前一個輸出塞進來。"
clickbait_en: "Some shell commands look like nested magic, but they are often just reusing previous output."
review_hint_zh: "`$(...)` 會先執行裡面的指令，再把輸出帶回外層。"
review_hint_en: "`$(...)` runs the inner command first, then substitutes its output."
enabled: true
---

## zh-TW

`$(...)` 是現代 shell 裡很常見的 command substitution 寫法。
它會先執行括號內的指令，然後把輸出文字插回外層命令。

## en

`$(...)` is a common modern shell syntax for command substitution.
It runs the inner command first and inserts the output back into the outer command.
