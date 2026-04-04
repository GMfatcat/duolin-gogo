---
id: git-fetch-basic
title: Git Fetch
title_zh: git fetch 的用途
title_en: Git Fetch
type: true-false
body_format: bilingual-section
tags: [git, remote]
difficulty: 1
question_zh: "`git fetch` 只會更新遠端追蹤資訊，不會直接把變更合併到目前分支。"
question_en: "`git fetch` updates remote-tracking information without merging changes into the current branch."
answer: true
clickbait_zh: "先看貨，不下單。這個 Git 動作很多人沒養成"
clickbait_en: "Look first, do not buy yet. Many developers skip this safer Git move."
review_hint_zh: "`git fetch` 會更新遠端資訊，但不會直接 merge。"
review_hint_en: "`git fetch` updates remote info, but does not merge into your branch."
confusion_with: [git-pull-composition]
metaphor_seed: [先看貨, 先觀察, 試探]
hook_style_tags: [safer-first, misunderstood, comparison]
enabled: true
---

## zh-TW

`git fetch` 會從遠端抓取最新資訊，更新本地的 remote-tracking references。
它不會直接改動你目前分支上的內容。
所以當你想先看遠端發生了什麼，再決定要不要合併時，`fetch` 是比較安全的第一步。

## en

`git fetch` downloads new information from the remote and updates local remote-tracking references.
It does not directly change the contents of your current branch.
That makes it a safer first step when you want to inspect remote changes before merging them.
