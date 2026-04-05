---
id: linux-ps-process-list
title_zh: ps 查看程序
title_en: PS
type: single-choice
body_format: bilingual-section
tags: [linux, shell, process]
difficulty: 2
question_zh: "`ps` 最常用來做什麼？"
question_en: "What is `ps` most commonly used for?"
choices_zh:
  - "列出目前的 process 狀態"
  - "修改檔案權限"
  - "搜尋檔案名稱"
choices_en:
  - "List current process information"
  - "Change file permissions"
  - "Search for filenames"
answer: 0
clickbait_zh: "程式看起來像消失了，但它可能只是在背景安靜地活著。"
clickbait_en: "A program may look gone while it is still quietly alive in the background."
review_hint_zh: "`ps` 是看 process 狀態與清單。"
review_hint_en: "`ps` shows process information."
confusion_with: [linux-kill-stop-process]
metaphor_seed: [點名, 值班表, 名單]
hook_style_tags: [comparison, misunderstood]
enabled: true
---

## zh-TW

`ps` 用來列出目前系統中的 process 資訊，例如 PID、狀態、命令名稱等。
它本身不是用來停止程序，而是幫你先看誰正在跑。
排查背景程序時，`ps` 幾乎是第一步。

## en

`ps` lists information about current processes, such as PID, status, and command name.
It does not stop a process by itself; it helps you inspect what is running first.
When checking background programs, `ps` is often the first command to use.
