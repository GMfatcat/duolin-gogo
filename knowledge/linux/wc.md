---
id: linux-wc-count-lines-words
title_zh: wc 計數
title_en: Linux wc
type: true-false
body_format: bilingual-section
tags: [linux, shell, text]
difficulty: 1
question_zh: "`wc` 常用來統計檔案或輸入中的行數、字數或位元組數。"
question_en: "`wc` is commonly used to count lines, words, or bytes from a file or input stream."
answer: true
clickbait_zh: "你想知道到底有多少，不一定要自己數，shell 早就有計數員。"
clickbait_en: "You do not have to count by hand. The shell already has a built-in counter."
review_hint_zh: "`wc` 是 shell 裡的基本計數工具。"
review_hint_en: "`wc` is a basic shell counting tool."
confusion_with: [linux-cat-print-file, linux-grep-search-text]
metaphor_seed: [計數員, 數量感, 快速統計]
hook_style_tags: [safe, misunderstood]
enabled: true
---

## zh-TW

`wc` 的名稱來自 word count，但它不只會數字數，也能統計行數和位元組數。
它常和 pipe 一起使用，例如先篩出一批文字，再快速看總共有幾行。
如果你要的是「這份輸出到底有多少」，`wc` 很常就是答案。

## en

`wc` stands for word count, but it can also count lines and bytes.
It is often combined with pipes so you can first filter some text and then quickly measure how much is left.
If your question is "how many?", `wc` is often the right tool.
