---
id: git-commit-snapshot
title: Git Commit
title_zh: git commit 的用途
title_en: Git Commit
type: true-false
body_format: bilingual-section
tags: [git, commits]
difficulty: 1
question_zh: "`git commit` 會把 staging area 裡的內容記錄成一個新的歷史節點。"
question_en: "`git commit` records the staged changes as a new history point."
answer: true
clickbait_zh: "你以為這一步只是存檔？其實它是在留下歷史證據"
clickbait_en: "You thought this was just save. It is really writing history."
review_hint_zh: "`git commit` 會把 staged changes 寫成新的 commit。"
review_hint_en: "`git commit` writes the staged changes into a new commit."
confusion_with: [git-add-staging]
metaphor_seed: [快照, 歷史證據, 存檔]
hook_style_tags: [comparison, serious, misunderstood]
enabled: true
---

## zh-TW

`git commit` 會把 staging area 裡的內容記錄成一個新的 commit。
這個 commit 會成為版本歷史的一部分，通常也會附上一段說明訊息。
如果你沒有先 `git add`，很多變更其實不會被 commit 進去。

## en

`git commit` records the contents of the staging area as a new commit.
That commit becomes part of the project history and usually includes a commit message.
If you did not stage the changes first, many edits will not be included in the commit.
