---
id: git-diff-compare
title: Git Diff
title_zh: git diff 的用途
title_en: Git Diff
type: single-choice
body_format: bilingual-section
tags: [git, inspection]
difficulty: 1
question_zh: "`git diff` 最主要是在做什麼？"
question_en: "What does `git diff` mainly do?"
choices_zh:
  - "比較變更差異"
  - "建立新的 commit"
  - "把分支推上遠端"
choices_en:
  - "Compare changes"
  - "Create a new commit"
  - "Push a branch to the remote"
answer: 0
clickbait_zh: "改了半天卻說不出哪裡變了？這個指令就是你的放大鏡。"
clickbait_en: "Changed a lot but cannot explain what moved? This command is your magnifier."
review_hint_zh: "`git diff` 用來看差異，不是提交，也不是推送。"
review_hint_en: "`git diff` shows differences. It does not commit or push anything."
confusion_with: [git-status-purpose, git-log-history]
metaphor_seed: [放大鏡, 對照表, 找不同]
hook_style_tags: [misunderstood, safer-first]
enabled: true
---

## zh-TW

`git diff` 會把兩個狀態之間的內容差異顯示出來，最常見的是看你目前尚未提交的修改。
它適合在 `git add` 或 `git commit` 前先確認到底改了什麼。
很多人只看 `git status`，但 `status` 只告訴你哪些檔案有變，`diff` 才會把內容細節展開。

## en

`git diff` displays the content differences between two states, often your current edits versus the last committed version.
It is useful before `git add` or `git commit` when you want to inspect the exact changes.
Many people only check `git status`, but `status` tells you which files changed, while `diff` shows what changed inside them.
