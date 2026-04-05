---
id: linux-cat-print-file
title_zh: cat 顯示檔案內容
title_en: Cat
type: single-choice
body_format: bilingual-section
tags: [linux, shell, file]
difficulty: 1
question_zh: "`cat` 最常見的用途是什麼？"
question_en: "What is the most common use of `cat`?"
choices_zh:
  - "顯示檔案內容"
  - "改變目錄"
  - "搜尋檔名"
choices_en:
  - "Print file contents"
  - "Change directories"
  - "Search for filenames"
answer: 0
clickbait_zh: "有時你不是要編輯，只是想先偷看檔案裡到底寫了什麼。"
clickbait_en: "Sometimes you do not want to edit. You just want a quick look inside the file."
review_hint_zh: "`cat` 最直接的用途就是把內容印出來。"
review_hint_en: "`cat` is commonly used to print contents directly."
confusion_with: [linux-tail-file-end]
metaphor_seed: [翻開, 看內容, 直接讀]
hook_style_tags: [safer_first, comparison]
enabled: true
---

## zh-TW

`cat` 會把檔案內容直接輸出到終端。
它常用來快速查看設定檔、文字檔或串接到其他指令。
如果你只是想看內容，`cat` 往往是最直接的選擇。

## en

`cat` prints file contents directly to the terminal.
It is often used to quickly inspect config files, text files, or pipe content onward.
If you just want to read a file, `cat` is one of the most direct commands.
