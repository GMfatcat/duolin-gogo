---
id: git-cherry-pick-purpose
title: Cherry-pick Purpose
type: true-false
body_format: bilingual-section
tags: [git, commits]
difficulty: 2
question: "`git cherry-pick` is used to apply a specific commit from another branch onto the current branch."
answer: true
clickbait: "One Git command can steal just one commit. Know which?"
review_hint: "Cherry-pick copies selected commit changes onto your current branch."
enabled: true
---

## zh-TW

`git cherry-pick` 可以把某一個特定 commit，或少數幾個 commit，套用到目前分支上。
當你只想拿一個局部修改，而不是整條 branch 都合進來時，這個指令很有用。

## en

`git cherry-pick` lets you take one specific commit, or a small set of commits, and apply them to your current branch.
It is useful when you want a targeted change without merging an entire branch.
