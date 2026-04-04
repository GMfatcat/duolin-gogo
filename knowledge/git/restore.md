---
id: git-restore-discard
title: Git Restore
title_zh: git restore 的用途
title_en: Git Restore
type: single-choice
body_format: bilingual-section
tags: [git, working-tree]
difficulty: 2
question_zh: "`git restore` 常見的用途是什麼？"
question_en: "What is a common use of `git restore`?"
choices_zh:
  - "還原工作區或 staged 的檔案內容"
  - "把目前分支推到遠端"
  - "列出所有 commit 歷史"
choices_en:
  - "Restores file contents in the working tree or staging area"
  - "Pushes the current branch to the remote"
  - "Lists all commit history"
answer: 0
clickbait_zh: "後悔藥真的存在，但亂吃會出事"
clickbait_en: "A regret button exists in Git, but using it carelessly still hurts."
review_hint_zh: "`git restore` 主要用來還原檔案內容。"
review_hint_en: "`git restore` is mainly used to restore file content."
confusion_with: [git-checkout-legacy, git-stash-purpose]
metaphor_seed: [後悔藥, 撤回, 恢復原狀]
hook_style_tags: [consequence, safer-first]
enabled: true
---

## zh-TW

`git restore` 常用來把工作區的檔案還原成某個版本，或把 staged 內容退回去。
它是從 `git checkout` 拆出來的比較清楚的指令之一。
因為它會直接影響你目前的檔案內容，所以使用前最好先確認你是否真的想放棄這些變更。

## en

`git restore` is often used to restore files in the working tree to a previous state or to unstage staged content.
It is one of the clearer commands that was split out from the old `git checkout`.
Because it directly affects your current file contents, you should make sure you truly want to discard those changes before using it.
