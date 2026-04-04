---
id: git-revert-safe-undo
title: Git Revert
title_zh: git revert 的用途
title_en: Git Revert
type: true-false
body_format: bilingual-section
tags: [git, history, safety]
difficulty: 2
question_zh: "`git revert` 會建立一個新的 commit，來反向抵消先前某個 commit 的效果。"
question_en: "`git revert` creates a new commit that reverses the effect of an earlier commit."
answer: true
clickbait_zh: "想反悔又不想改壞歷史？這個指令就是比較安全的退路。"
clickbait_en: "Need to undo something without rewriting history? This is the safer escape route."
review_hint_zh: "`git revert` 不是刪歷史，而是新增一個反向 commit。"
review_hint_en: "`git revert` does not erase history. It adds a reversing commit."
confusion_with: [git-reset-head]
metaphor_seed: [反向動作, 安全後悔藥, 補正紀錄]
hook_style_tags: [safer-first, comparison, fear_of_mistake]
enabled: true
---

## zh-TW

`git revert` 會新增一個新的 commit，用來把某個舊 commit 帶來的變更反向抵消掉。
它和 `git reset` 很容易被混淆，但 `revert` 的重點是保留歷史，只是在歷史上再補一筆「撤回」紀錄。
這讓它在共享分支或已經推上遠端的情境裡，通常比直接重寫歷史更安全。

## en

`git revert` creates a new commit that reverses the effect of an earlier commit.
It is easy to confuse with `git reset`, but revert is focused on preserving history and adding a visible undo record.
That makes it a safer option on shared branches or after commits have already been pushed.
