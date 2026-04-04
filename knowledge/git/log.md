---
id: git-log-history
title: Git Log
title_zh: git log 的用途
title_en: Git Log
type: single-choice
body_format: bilingual-section
tags: [git, history]
difficulty: 1
question_zh: "`git log` 最主要是在看什麼？"
question_en: "What do you mainly inspect with `git log`?"
choices_zh:
  - "提交歷史"
  - "遠端推送速度"
  - "目前 CPU 使用率"
choices_en:
  - "Commit history"
  - "Remote push speed"
  - "Current CPU usage"
answer: 0
clickbait_zh: "想知道專案到底發生過什麼，先別猜，直接翻歷史。"
clickbait_en: "Want to know what really happened in the repo? Stop guessing and read the history."
review_hint_zh: "`git log` 是看 commit 歷史，不是看工作區差異。"
review_hint_en: "`git log` shows commit history, not working tree diffs."
confusion_with: [git-status-purpose, git-diff-compare]
metaphor_seed: [時間軸, 歷史書, 案件紀錄]
hook_style_tags: [safer-first, serious]
enabled: true
---

## zh-TW

`git log` 用來查看提交歷史，也就是專案一路累積下來的 commit 記錄。
你可以從中看到提交訊息、作者、時間，以及 commit 之間的順序。
當你想理解「這段程式是什麼時候改的」或「最近做了哪些變更」時，`git log` 是很常用的入口。

## en

`git log` is used to inspect commit history, which is the record of how the project evolved over time.
It lets you see commit messages, authors, timestamps, and the order of commits.
When you want to understand when a change happened or what has happened recently, `git log` is a common starting point.
