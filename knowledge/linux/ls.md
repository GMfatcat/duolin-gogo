---
id: linux-ls-list-files
title_zh: ls 列出檔案
title_en: LS
type: single-choice
body_format: bilingual-section
tags: [linux, shell, files]
difficulty: 1
question_zh: "`ls` 最主要在做什麼？"
question_en: "What does `ls` mainly do?"
choices_zh:
  - "列出目前目錄中的檔案和資料夾"
  - "切換到另一個目錄"
  - "顯示目前登入的使用者"
choices_en:
  - "List files and directories in the current location"
  - "Switch to another directory"
  - "Show the currently logged-in user"
answer: 0
clickbait_zh: "你以為資料夾是空的？先別急，也許只是你還沒真的看。"
clickbait_en: "Think the folder is empty? Maybe you just have not actually looked yet."
review_hint_zh: "`ls` = 列出檔案與資料夾。"
review_hint_en: "`ls` lists files and directories."
confusion_with: [linux-pwd-current-directory]
metaphor_seed: [看清單, 看貨架, 打開抽屜]
hook_style_tags: [safer_first, comparison]
enabled: true
---

## zh-TW

`ls` 會列出目前目錄中的檔案和資料夾。
常見變化像 `ls -l` 可以看更詳細資訊，`ls -a` 可以連隱藏檔一起看。
它是你進到一個目錄後最常先用來確認內容的指令之一。

## en

`ls` lists files and directories in the current location.
Common variations include `ls -l` for more detail and `ls -a` to include hidden files.
It is one of the first commands people use after entering a directory to inspect its contents.
