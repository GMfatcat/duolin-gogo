---
id: git-status-purpose
title: Git Status
title_zh: git status 的用途
title_en: Git Status
type: single-choice
body_format: bilingual-section
tags: [git, basics]
difficulty: 1
question_zh: "`git status` 最主要是在做什麼？"
question_en: "What does `git status` mainly do?"
choices_zh:
  - "顯示工作區、staging area、分支狀態"
  - "把目前變更直接推到遠端"
  - "刪掉所有未追蹤檔案"
choices_en:
  - "Shows the working tree, staging area, and branch status"
  - "Pushes current changes directly to the remote"
  - "Deletes all untracked files"
answer: 0
clickbait_zh: "很多 Git 災難，其實只差這一步有沒有先看"
clickbait_en: "A lot of Git disasters start with skipping this one simple check."
review_hint_zh: "`git status` 是拿來看目前狀態，不會修改檔案。"
review_hint_en: "`git status` inspects the current state. It does not change files."
confusion_with: [git-add-staging, git-commit-snapshot]
metaphor_seed: [儀表板, 體檢, 先看狀況]
hook_style_tags: [safer-first, defensive]
enabled: true
---

## zh-TW

`git status` 會告訴你目前有哪些檔案被修改、哪些已經 staged、哪些還沒被追蹤，以及你現在所在的分支狀態。
它不會改動任何內容，所以很適合在做操作前先確認現況。
很多 Git 錯誤，其實是因為使用者沒有先看清楚狀態。

## en

`git status` tells you which files are modified, which changes are staged, which files are untracked, and what branch state you are in.
It does not modify anything, so it is a safe way to inspect the current situation before taking action.
Many Git mistakes happen because the user did not check the state first.
