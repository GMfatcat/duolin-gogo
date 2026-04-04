---
id: git-tag-release-marker
title: Git Tag
title_zh: git tag 的用途
title_en: Git Tag
type: single-choice
body_format: bilingual-section
tags: [git, release]
difficulty: 2
question_zh: "`git tag` 最常用來標記什麼？"
question_en: "What is `git tag` commonly used to mark?"
choices_zh:
  - "某個重要的版本點"
  - "今天午餐吃什麼"
  - "本地 CPU 溫度"
choices_en:
  - "An important version point"
  - "What you had for lunch"
  - "The local CPU temperature"
answer: 0
clickbait_zh: "版本真的要發了，光靠 commit 訊息還不夠，這時候要插旗。"
clickbait_en: "When a release really matters, a commit message alone is often not enough. You plant a flag."
review_hint_zh: "`git tag` 常用來標記版本，例如 release 或 milestone。"
review_hint_en: "`git tag` is often used to mark versions such as releases or milestones."
confusion_with: [git-log-history, git-branch-list]
metaphor_seed: [插旗, 里程碑, 版本標記]
hook_style_tags: [serious, consequence]
enabled: true
---

## zh-TW

`git tag` 常用來標記一個重要版本點，例如 `v1.0.0` 這種 release 節點。
和分支不同，tag 通常不是拿來持續開發，而是用來指出「這個 commit 很值得被記住」。
所以當你要標示版本發布、里程碑或可回溯的穩定點時，tag 很有用。

## en

`git tag` is commonly used to mark an important version point, such as a release like `v1.0.0`.
Unlike a branch, a tag is usually not meant for ongoing development. It is meant to say, "this commit matters and should be easy to find again."
That makes tags useful for releases, milestones, and other stable reference points.
