---
id: linux-pipe-command-chain
title_zh: pipe 管線
title_en: Linux Pipe
type: true-false
body_format: bilingual-section
tags: [linux, shell, pipe]
difficulty: 2
question_zh: "Linux pipe `|` 常用來把前一個指令的輸出送給下一個指令當輸入。"
question_en: "A Linux pipe `|` is commonly used to send one command's output into the next command's input."
answer: true
clickbait_zh: "你不一定要先存成檔案，很多時候資料其實可以直接接下去。"
clickbait_en: "You do not always need an intermediate file. Sometimes the output can flow straight into the next tool."
review_hint_zh: "pipe 會把上一個指令的標準輸出接到下一個指令。"
review_hint_en: "A pipe connects one command's standard output to another command's input."
confusion_with: [linux-redirect-output-file]
metaphor_seed: [接水管, 直接接續, 串接]
hook_style_tags: [comparison, misunderstood]
enabled: true
---

## zh-TW

pipe `|` 的核心概念是把一個指令的輸出直接餵給另一個指令，不必先落成檔案。
這讓 shell 很適合把小工具串起來，例如先搜尋，再過濾，再計數。
它是在串接資料流，不是在把輸出寫到檔案裡。

## en

A pipe `|` sends the output of one command directly into another command.
That makes shell workflows good at chaining small tools together, such as searching, filtering, and counting.
It connects data streams instead of writing the output into a file.
