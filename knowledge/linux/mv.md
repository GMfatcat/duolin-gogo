---
id: linux-mv-move-rename
title_zh: mv 搬移與重新命名
title_en: Mv
type: single-choice
body_format: bilingual-section
tags: [linux, shell, file]
difficulty: 1
question_zh: "`mv` 最主要在做什麼？"
question_en: "What does `mv` mainly do?"
choices_zh:
  - "搬移或重新命名檔案"
  - "複製檔案"
  - "顯示檔案內容"
choices_en:
  - "Move or rename files"
  - "Copy files"
  - "Print file contents"
answer: 0
clickbait_zh: "有時候它看起來像改名，其實本質上也是一種搬動。"
clickbait_en: "Sometimes it looks like renaming, but underneath it is still a kind of move."
review_hint_zh: "`mv` 可以搬位置，也可以改名稱。"
review_hint_en: "`mv` can change location or rename the target."
confusion_with: [linux-cp-copy-files]
metaphor_seed: [搬家, 改名, 換位置]
hook_style_tags: [comparison, misunderstood]
enabled: true
---

## zh-TW

`mv` 可以把檔案移到別的位置，也能順便改名。
如果目標是在同一個資料夾，它常常就被拿來做重新命名。
它和 `cp` 的差別在於，不是保留一份副本，而是把原本目標搬走。

## en

`mv` can move a file to a different location or rename it.
When the destination stays in the same folder, it is often used just for renaming.
Unlike `cp`, it does not normally leave a duplicate behind.
