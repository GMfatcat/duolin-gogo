---
id: git-pull-composition
title: Git Pull
title_zh: git pull 的組成
title_en: Git Pull
type: single-choice
body_format: bilingual-section
tags: [git, remote]
difficulty: 2
question_zh: "`git pull` 在預設情況下通常等於哪兩個動作的組合？"
question_en: "Under default behavior, what is `git pull` usually a combination of?"
choices_zh:
  - "`git add` + `git commit`"
  - "`git fetch` + `git merge`"
  - "`git status` + `git checkout`"
choices_en:
  - "`git add` + `git commit`"
  - "`git fetch` + `git merge`"
  - "`git status` + `git checkout`"
answer: 1
clickbait_zh: "這個指令看起來一步到位，其實偷偷做了兩件事"
clickbait_en: "This command looks simple, but it secretly does two jobs."
review_hint_zh: "`git pull` 通常是 `fetch` 加上 `merge`。"
review_hint_en: "`git pull` is usually `fetch` plus `merge`."
confusion_with: [git-fetch-basic, git-merge-purpose]
metaphor_seed: [自動打包, 一步到位, 偷偷兩件事]
hook_style_tags: [misunderstood, comparison, consequence]
enabled: true
---

## zh-TW

`git pull` 通常可以理解成先 `git fetch`，再把抓到的遠端變更整合進目前分支。
因為它一步做了兩件事，所以很方便，但也比較容易在你還沒看清楚狀況時就把變更拉進來。
如果你想更保守，通常會先 `fetch` 再自己決定如何整合。

## en

`git pull` is usually understood as doing `git fetch` first and then integrating the fetched changes into your current branch.
It is convenient because it combines two steps, but that convenience can hide important context.
If you want more control, many developers prefer to fetch first and decide how to integrate afterward.
