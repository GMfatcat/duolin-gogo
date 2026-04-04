---
id: git-branch-list
title: Git Branch
title_zh: git branch 的用途
title_en: Git Branch
type: single-choice
body_format: bilingual-section
tags: [git, branching]
difficulty: 1
question_zh: "`git branch` 最常見的用途是什麼？"
question_en: "What is the most common use of `git branch`?"
choices_zh:
  - "列出或建立分支"
  - "直接推送到遠端"
  - "刪除整個儲存庫"
choices_en:
  - "List or create branches"
  - "Push directly to the remote"
  - "Delete the whole repository"
answer: 0
clickbait_zh: "你以為你每天都在切分支，但這個指令其實先負責把名單攤開。"
clickbait_en: "You think this is only for switching, but it often starts with a branch list."
review_hint_zh: "`git branch` 常用來列出分支，也可以建立新分支。"
review_hint_en: "`git branch` often lists branches and can also create a new one."
confusion_with: [git-switch-branch, git-checkout-legacy]
metaphor_seed: [名單, 分流, 路線圖]
hook_style_tags: [comparison, misunderstood]
enabled: true
---

## zh-TW

`git branch` 最常見的用途是列出目前有哪些分支，也可以搭配新名稱去建立分支。
它本身不會直接把你切換到另一條分支，那通常要搭配 `git switch` 或舊式的 `git checkout`。
很多人把「看分支」、「建分支」、「切分支」混在一起記，這張卡就是在拆開它們。

## en

`git branch` is commonly used to list the branches that already exist, and it can also create a new branch when you pass a name.
By itself, it does not switch your working branch. That usually happens with `git switch` or the older `git checkout`.
Many learners mix up listing, creating, and switching branches, so this card separates those ideas.
