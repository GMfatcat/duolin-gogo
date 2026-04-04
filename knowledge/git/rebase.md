---
id: git-rebase-vs-merge
title: Rebase vs Merge
title_zh: Rebase 跟 Merge 的差別
title_en: Rebase vs Merge
type: single-choice
body_format: bilingual-section
tags: [git, branching]
difficulty: 2
question_zh: "git rebase 主要是在做什麼？"
question_en: "What does git rebase mainly do?"
choices_zh:
  - "建立一個 merge commit"
  - "把 commits 重新接到新的 base 上"
  - "自動刪掉所有衝突 commit"
choices_en:
  - "Creates a merge commit between branches"
  - "Replays commits onto a new base"
  - "Deletes conflicting commits automatically"
answer: 1
clickbait_zh: "你真的懂 rebase 跟 merge 的差別嗎？"
clickbait_en: "Most Git beginners misunderstand rebase. Do you?"
review_hint_zh: "Rebase = 把 commits 重放到新的 base 上。"
review_hint_en: "Rebase = replay commits on top of another base."
confusion_with: [git-merge-purpose]
metaphor_seed: [搬家, 換底座, 重排]
hook_style_tags: [misunderstood, comparison]
enabled: true
---

## zh-TW

`git rebase` 會把目前分支上的 commits，重新接到另一個 base commit 上。
它常用來讓提交歷史更線性、更容易閱讀。
跟 `merge` 不同，`rebase` 通常不會額外產生一個 merge commit。

## en

`git rebase` takes commits from your current branch and reapplies them onto another base commit.
It is often used to keep history linear and easier to read.
Unlike `merge`, it usually avoids creating an extra merge commit.
