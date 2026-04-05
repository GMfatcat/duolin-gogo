---
id: linux-find-search-files
title_zh: find 搜尋檔案
title_en: Find
type: single-choice
body_format: bilingual-section
tags: [linux, shell, search]
difficulty: 2
question_zh: "`find` 最常用來做什麼？"
question_en: "What is `find` most commonly used for?"
choices_zh:
  - "在檔案系統中搜尋符合條件的檔案或目錄"
  - "查看目前 shell 的歷史命令"
  - "只搜尋檔案內容中的關鍵字"
choices_en:
  - "Search the filesystem for files or directories that match conditions"
  - "Display shell command history"
  - "Only search inside file contents for keywords"
answer: 0
clickbait_zh: "你以為檔案不見了，可能只是你還不會去整個檔案系統把它找出來。"
clickbait_en: "You may think the file is gone. The real issue might be that you do not know how to search the filesystem properly yet."
review_hint_zh: "`find` 會在檔案系統裡找檔名或條件。"
review_hint_en: "`find` searches the filesystem by name or other conditions."
confusion_with: [linux-grep-search-text]
metaphor_seed: [尋寶, 地毯式搜索, 尋人啟事]
hook_style_tags: [comparison, misunderstood]
enabled: true
---

## zh-TW

`find` 用來在檔案系統中遞迴搜尋檔案或目錄，條件可以是名稱、型別、時間等。
它是在找檔案「在哪裡」，不是找檔案內容裡有沒有某段文字。
如果你常常搞混 `find` 和 `grep`，可以記成前者找位置，後者找內容。

## en

`find` recursively searches the filesystem for files or directories that match certain conditions such as name, type, or time.
It answers where something is located rather than whether a piece of text appears inside a file.
If you mix it up with `grep`, remember that `find` is about location and `grep` is about content.
