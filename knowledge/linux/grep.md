---
id: linux-grep-search-text
title_zh: grep 搜尋文字
title_en: Grep
type: single-choice
body_format: bilingual-section
tags: [linux, shell, search]
difficulty: 2
question_zh: "`grep` 最主要的用途是什麼？"
question_en: "What is the main purpose of `grep`?"
choices_zh:
  - "搜尋文字內容是否符合某個模式"
  - "建立新的資料夾"
  - "變更檔案權限"
choices_en:
  - "Search text content for lines that match a pattern"
  - "Create a new directory"
  - "Change file permissions"
answer: 0
clickbait_zh: "不是檔案不見，是你還沒學會怎麼把關鍵字挖出來。"
clickbait_en: "The file may not be missing. You may just not know how to pull the keyword out yet."
review_hint_zh: "`grep` 用來找符合模式的文字行。"
review_hint_en: "`grep` finds text lines that match a pattern."
confusion_with: [linux-ls-list-files]
metaphor_seed: [放大鏡, 關鍵字搜索, 從一堆字裡挖答案]
hook_style_tags: [misunderstood, comparison]
enabled: true
---

## zh-TW

`grep` 會在文字內容中搜尋符合指定模式的行。
你可以拿它查 log、設定檔或指令輸出裡是否出現某個關鍵字。
它不是找檔名，而是找內容。

## en

`grep` searches text content for lines that match a given pattern.
You can use it on logs, config files, or command output to find lines containing a keyword.
It searches content, not filenames.
