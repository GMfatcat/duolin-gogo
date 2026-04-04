---
id: git-push-upstream
title: Git Push -u
title_zh: git push -u 的用途
title_en: Git Push Upstream
type: single-choice
body_format: bilingual-section
tags: [git, remote, push]
difficulty: 2
question_zh: "`git push -u origin feature-x` 最關鍵的額外效果是什麼？"
question_en: "What is the most important extra effect of `git push -u origin feature-x`?"
choices_zh:
  - "設定上游追蹤分支"
  - "清掉所有本地 commit"
  - "自動解掉 merge conflict"
choices_en:
  - "Set the upstream tracking branch"
  - "Delete all local commits"
  - "Automatically resolve merge conflicts"
answer: 0
clickbait_zh: "第一次 push 時少打一個字母，之後每天都要多做一步。"
clickbait_en: "Miss one small flag on the first push, and you may pay that tax every day after."
review_hint_zh: "`-u` 會設定 upstream，之後 push/pull 會更省事。"
review_hint_en: "`-u` sets the upstream branch so later push/pull commands are simpler."
confusion_with: [git-pull-composition, git-remote-origin]
metaphor_seed: [綁定, 追蹤線, 配對]
hook_style_tags: [safer-first, consequence, misunderstood]
enabled: true
---

## zh-TW

`git push -u origin feature-x` 除了把分支推到遠端，還會順手設定目前本地分支要追蹤哪一條遠端分支。
這樣之後你在同一條分支上通常只要打 `git push` 或 `git pull`，Git 就知道要對應哪個 upstream。
很多人第一次 push 成功就以為結束了，但少了 `-u` 之後常常還要手動補指定分支。

## en

`git push -u origin feature-x` not only pushes your branch to the remote, but also sets the upstream tracking relationship.
That means later `git push` or `git pull` commands usually know which remote branch should be paired with your current local branch.
Many people think the first push is only about uploading, but the `-u` flag also removes later friction.
