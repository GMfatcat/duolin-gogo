---
id: git-stash-purpose
title: Git Stash
title_zh: git stash 的用途
title_en: Git Stash
type: true-false
body_format: bilingual-section
tags: [git, working-tree]
difficulty: 2
question_zh: "`git stash` 可以先把目前還沒 commit 的變更暫時收起來，之後再拿回來。"
question_en: "`git stash` can temporarily put away uncommitted changes so you can bring them back later."
answer: true
clickbait_zh: "先藏起來，不代表你真的解決了"
clickbait_en: "Hiding the mess is not the same as fixing it."
review_hint_zh: "`git stash` 是暫時收起未提交變更，之後還能再套回來。"
review_hint_en: "`git stash` temporarily stores uncommitted changes so you can reapply them later."
confusion_with: [git-restore-discard]
metaphor_seed: [藏起來, 收納箱, 暫時擱著]
hook_style_tags: [playful, comparison, curiosity-gap]
enabled: true
---

## zh-TW

`git stash` 會把目前尚未 commit 的變更暫時收起來，讓工作區回到比較乾淨的狀態。
這在你臨時要切分支或處理別的事情時很有用。
不過 stash 不是消失，而是暫存起來，之後還需要你手動套回來。

## en

`git stash` temporarily stores your uncommitted changes and gives you a cleaner working tree.
This is useful when you suddenly need to switch branches or handle something else first.
But a stash is not gone forever. It is only set aside until you explicitly apply it again.
