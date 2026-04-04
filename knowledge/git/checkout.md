---
id: git-checkout-legacy
title: Git Checkout
title_zh: git checkout 的舊式用途
title_en: Git Checkout
type: true-false
body_format: bilingual-section
tags: [git, branches]
difficulty: 2
question_zh: "`git checkout` 這個舊指令既能切分支，也能還原檔案內容。"
question_en: "The older `git checkout` command can both switch branches and restore file content."
answer: true
clickbait_zh: "一個指令做太多事，往往就是新手出事的開始"
clickbait_en: "When one command does too much, beginners usually pay for it."
review_hint_zh: "`git checkout` 是舊式多功能指令，能切分支也能還原內容。"
review_hint_en: "`git checkout` is an older multi-purpose command for switching and restoring."
confusion_with: [git-switch-branch, git-restore-discard]
metaphor_seed: [瑞士刀, 一招多用, 容易切錯]
hook_style_tags: [misunderstood, chaotic, comparison]
enabled: true
---

## zh-TW

在較新的 Git 版本裡，`git switch` 和 `git restore` 被拆出來，就是因為 `git checkout` 以前同時肩負太多用途。
它可以切換分支，也可以把某個檔案還原成特定版本。
所以看到 `checkout` 時，要特別看清楚後面的參數，因為它不一定只是「切分支」。

## en

In newer Git versions, `git switch` and `git restore` were introduced because `git checkout` used to handle too many jobs.
It can switch branches, but it can also restore a file to a previous state.
So whenever you see `checkout`, you should check the arguments carefully because it does not always mean branch switching.
