---
id: git-reset-head
title: Git Reset
title_zh: git reset 的常見用途
title_en: Git Reset
type: single-choice
body_format: bilingual-section
tags: [git, history, staging]
difficulty: 3
question_zh: "`git reset` 在最常見的教學情境裡，通常是在調整什麼？"
question_en: "In common Git usage, what does `git reset` often adjust?"
choices_zh:
  - "HEAD 或暫存區狀態"
  - "螢幕亮度"
  - "遠端帳號密碼"
choices_en:
  - "HEAD or staging state"
  - "Screen brightness"
  - "Remote account passwords"
answer: 0
clickbait_zh: "後悔藥不是只有一種，這個指令就是 Git 裡最危險也最有用的那種。"
clickbait_en: "Git does have a regret button, but it is also one of the riskiest tools."
review_hint_zh: "`git reset` 常用來移動 HEAD，或把內容從 staging 拿回來。"
review_hint_en: "`git reset` often moves HEAD or unstages changes."
confusion_with: [git-restore-discard, git-revert]
metaphor_seed: [後悔藥, 倒帶, 拔插頭]
hook_style_tags: [fear_of_mistake, misunderstood, consequence]
enabled: true
---

## zh-TW

`git reset` 是一個功能很多、也需要特別小心的指令。
在最常見的使用情境裡，它常被拿來移動 `HEAD`，或把已經 `git add` 的內容從 staging area 拿回來。
因為不同模式像 `--soft`、`--mixed`、`--hard` 的影響差很多，所以學它時最好先抓住核心：它會改變你目前歷史指標或暫存狀態。

## en

`git reset` is a powerful command and one that requires extra care.
In common day-to-day use, it often moves `HEAD` or removes already staged changes from the staging area.
Different modes such as `--soft`, `--mixed`, and `--hard` have very different consequences, so the key first idea is that reset changes your current history pointer or staging state.
