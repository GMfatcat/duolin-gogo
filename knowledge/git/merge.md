---
id: git-merge-purpose
title: Git Merge
title_zh: git merge 的用途
title_en: Git Merge
type: single-choice
body_format: bilingual-section
tags: [git, branching]
difficulty: 2
question_zh: "`git merge` 最主要是在做什麼？"
question_en: "What does `git merge` mainly do?"
choices_zh:
  - "把另一條分支的變更整合進目前分支"
  - "把所有 commit 改寫成線性歷史"
  - "只更新 remote-tracking references"
choices_en:
  - "Integrates another branch into the current branch"
  - "Rewrites all commits into a linear history"
  - "Only updates remote-tracking references"
answer: 0
clickbait_zh: "兩條分支要不要正式在一起，後果差很多"
clickbait_en: "When two branches make it official, the consequences matter."
review_hint_zh: "`git merge` 是把另一條分支整合進目前分支。"
review_hint_en: "`git merge` integrates another branch into the current branch."
confusion_with: [git-rebase-vs-merge, git-pull-composition]
metaphor_seed: [合流, 正式在一起, 匯合]
hook_style_tags: [comparison, consequence]
enabled: true
---

## zh-TW

`git merge` 會把另一條分支的變更整合進你目前所在的分支。
如果兩條分支都各自有新提交，Git 可能會建立一個 merge commit。
它保留了分支匯合的歷史，跟 `rebase` 的思路不同。

## en

`git merge` integrates changes from another branch into your current branch.
If both branches have new commits, Git may create a merge commit.
This preserves the branching history, which is different from the goal of `rebase`.
