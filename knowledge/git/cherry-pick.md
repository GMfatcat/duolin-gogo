---
id: git-cherry-pick-purpose
title: Cherry-pick Purpose
title_zh: Cherry-pick 的用途
title_en: Cherry-pick Purpose
type: true-false
body_format: bilingual-section
tags: [git, commits]
difficulty: 2
question_zh: "`git cherry-pick` 會把指定的一個 commit 套用到目前分支上。"
question_en: "`git cherry-pick` is used to apply a specific commit from another branch onto the current branch."
answer: true
clickbait_zh: "哪個 Git 指令可以只拿走一個 commit？"
clickbait_en: "One Git command can steal just one commit. Know which?"
review_hint_zh: "Cherry-pick 會把選定 commit 的變更套到目前分支。"
review_hint_en: "Cherry-pick copies selected commit changes onto your current branch."
confusion_with: [git-merge-purpose]
metaphor_seed: [摘櫻桃, 單點搬運, 只拿一個]
hook_style_tags: [targeted, curiosity-gap]
enabled: true
---

## zh-TW

`git cherry-pick` 可以把某一個特定 commit，從別的分支套用到你目前所在的分支上。
這在你只想拿某個小修正、但不想整條 branch 一起 merge 進來時特別有用。

## en

`git cherry-pick` lets you take one specific commit, or a small set of commits, and apply them to your current branch.
It is useful when you want a targeted change without merging an entire branch.
