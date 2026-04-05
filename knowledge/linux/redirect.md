---
id: linux-redirect-output-file
title_zh: 重新導向輸出
title_en: Linux Redirection
type: true-false
body_format: bilingual-section
tags: [linux, shell, redirect]
difficulty: 2
question_zh: "輸出重新導向常用來把指令結果寫進檔案，而不是只顯示在終端上。"
question_en: "Output redirection is commonly used to write command results into a file instead of only showing them in the terminal."
answer: true
clickbait_zh: "不是每個結果都該留在螢幕上，有些東西其實應該被送進檔案。"
clickbait_en: "Not every result belongs on the screen. Some output should be captured as a file instead."
review_hint_zh: "重新導向是把輸出改送到檔案或其他目標。"
review_hint_en: "Redirection changes where command output is sent."
confusion_with: [linux-pipe-command-chain]
metaphor_seed: [改道, 存檔, 轉送]
hook_style_tags: [comparison, safer_first]
enabled: true
---

## zh-TW

shell 的重新導向會改變輸出或輸入的去向，最常見的是把標準輸出寫進檔案。
這和 pipe 不一樣，因為 pipe 是把輸出接到下一個指令，重新導向則是改送到檔案或其他目標。
常見的心智模型是「不是接給下一個工具，而是改到另一個出口」。

## en

Shell redirection changes where input or output goes, most commonly by writing standard output into a file.
That is different from a pipe: a pipe passes output to another command, while redirection sends it to a file or another destination.
It is usually about changing the destination rather than chaining tools together.
