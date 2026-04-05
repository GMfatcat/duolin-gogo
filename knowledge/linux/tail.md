---
id: linux-tail-file-end
title_zh: tail 看檔案尾端
title_en: Tail
type: true-false
body_format: bilingual-section
tags: [linux, shell, log]
difficulty: 2
question_zh: "`tail` 常用來查看檔案最後幾行，特別是 log。"
question_en: "`tail` is often used to inspect the last lines of a file, especially logs."
answer: true
clickbait_zh: "你不用從頭看到尾，很多時候真正有用的線索就在最後面。"
clickbait_en: "You do not need to read from the top. The useful clue is often at the end."
review_hint_zh: "`tail` 特別適合看 log 尾端。"
review_hint_en: "`tail` is especially useful for the end of logs."
confusion_with: [linux-cat-print-file]
metaphor_seed: [最後面, 尾巴, 最新]
hook_style_tags: [safer_first, comparison]
enabled: true
---

## zh-TW

`tail` 會顯示檔案最後幾行內容。
它很常拿來看 log，因為最新的事件通常出現在尾端。
搭配 `-f` 時，還能持續追蹤檔案的新輸出。

## en

`tail` shows the last lines of a file.
It is commonly used for logs because the newest events usually appear at the end.
With `-f`, it can continue following new output as the file grows.
