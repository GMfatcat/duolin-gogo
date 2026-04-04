---
id: git-add-staging
title: Git Add
title_zh: git add 的用途
title_en: Git Add
type: single-choice
body_format: bilingual-section
tags: [git, staging]
difficulty: 1
question_zh: "`git add` 最主要是在做什麼？"
question_en: "What does `git add` mainly do?"
choices_zh:
  - "直接把變更送到遠端"
  - "把變更放進 staging area，準備 commit"
  - "永久刪掉工作區的變更"
choices_en:
  - "Pushes changes directly to the remote"
  - "Stages changes so they are ready to commit"
  - "Permanently deletes working tree changes"
answer: 1
clickbait_zh: "這個按鈕像購物車，不是結帳，你有按對嗎？"
clickbait_en: "This command is more like a cart than a checkout. Are you using it right?"
review_hint_zh: "`git add` 是把變更放進 staging area，不是 commit。"
review_hint_en: "`git add` stages changes. It does not create a commit."
confusion_with: [git-commit-snapshot]
metaphor_seed: [購物車, 放進籃子, 暫存清單]
hook_style_tags: [comparison, misunderstood, safer-first]
enabled: true
---

## zh-TW

`git add` 會把你目前的檔案變更放進 staging area。
你可以把它想成「先選好這次要提交哪些內容」，但還沒有真的建立 commit。
常見誤解是把 `git add` 當成儲存或送出，其實它只是準備階段。

## en

`git add` moves your current file changes into the staging area.
You can think of it as selecting what should be included in the next commit, but the commit does not exist yet.
Many beginners confuse it with saving or publishing, but it is only a preparation step.
