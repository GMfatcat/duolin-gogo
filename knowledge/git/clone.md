---
id: git-clone-local-copy
title: Git Clone
title_zh: git clone 的用途
title_en: Git Clone
type: true-false
body_format: bilingual-section
tags: [git, remote]
difficulty: 1
question_zh: "`git clone` 會把遠端儲存庫複製到本地，並建立初始追蹤設定。"
question_en: "`git clone` copies a remote repository locally and sets up initial tracking."
answer: true
clickbait_zh: "這不是單純下載資料夾，它會把整個 Git 關係也一起帶下來。"
clickbait_en: "This is not just downloading a folder. It brings the Git relationship with it."
review_hint_zh: "`git clone` 會複製儲存庫內容與基本遠端追蹤資訊。"
review_hint_en: "`git clone` copies the repository and its initial remote tracking setup."
confusion_with: [git-init-repository]
metaphor_seed: [搬家, 複製基地, 新據點]
hook_style_tags: [comparison, safer-first]
enabled: true
---

## zh-TW

`git clone` 會把一個現有的 Git 儲存庫完整複製到你的本地電腦。
除了檔案本身，它也會帶下來 `.git` 歷史資料，並幫你設定好預設遠端，通常叫做 `origin`。
所以它不是單純把檔案下載下來，而是把整個版本控制關係一起建立好。

## en

`git clone` creates a local copy of an existing Git repository.
It brings down not only the files, but also the `.git` history and a default remote, usually named `origin`.
That is why cloning is more than downloading files: it also sets up the repository relationship.
