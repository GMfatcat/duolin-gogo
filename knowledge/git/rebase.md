---
id: git-rebase-vs-merge
title: Rebase vs Merge
type: single-choice
body_format: bilingual-section
tags: [git, branching]
difficulty: 2
question: "What does git rebase mainly do?"
choices:
  - "Creates a merge commit between branches"
  - "Replays commits onto a new base"
  - "Deletes conflicting commits automatically"
answer: 1
clickbait: "Most Git beginners misunderstand rebase. Do you?"
review_hint: "Rebase = replay commits on top of another base."
enabled: true
---

## zh-TW

`git rebase` 會把目前分支上的提交重新套用到另一個 base commit 上。
它常用來讓提交歷史更線性、更乾淨。
和 `merge` 不同，它通常不會產生額外的 merge commit。

## en

`git rebase` takes commits from your current branch and reapplies them onto another base commit.
It is often used to keep history linear and easier to read.
Unlike `merge`, it usually avoids creating an extra merge commit.
