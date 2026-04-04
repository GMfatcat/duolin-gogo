---
id: git-init-repository
title: Git Init
title_zh: git init 的用途
title_en: Git Init
type: true-false
body_format: bilingual-section
tags: [git, setup]
difficulty: 1
question_zh: "`git init` 會在目前資料夾建立一個新的 Git 儲存庫。"
question_en: "`git init` creates a new Git repository in the current directory."
answer: true
clickbait_zh: "還沒連到遠端也能開始？這個指令就是一切的起點。"
clickbait_en: "No remote yet? This command is still enough to start everything."
review_hint_zh: "`git init` 會建立新的本地 Git 儲存庫。"
review_hint_en: "`git init` creates a new local Git repository."
confusion_with: [git-clone-local-copy]
metaphor_seed: [起點, 打地基, 空白專案]
hook_style_tags: [comparison, safer-first]
enabled: true
---

## zh-TW

`git init` 會在目前資料夾建立一個新的 Git 儲存庫，通常會生成 `.git` 目錄來保存版本控制資訊。
它適合用在你本地新開一個專案，還沒有既有遠端可以 clone 的情境。
如果你是接手現有遠端專案，通常會用 `git clone`；如果是自己從零開始，就常從 `git init` 起步。

## en

`git init` creates a new Git repository in the current directory, usually by creating a `.git` folder for version control data.
It fits the case where you are starting a local project from scratch and do not have an existing remote repository to clone.
If you are joining an existing project, you usually use `git clone`; if you are starting fresh, `git init` is a common first step.
