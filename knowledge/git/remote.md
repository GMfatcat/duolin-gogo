---
id: git-remote-origin
title: Git Remote
title_zh: git remote 的用途
title_en: Git Remote
type: single-choice
body_format: bilingual-section
tags: [git, remote]
difficulty: 2
question_zh: "`git remote` 最常用來管理什麼？"
question_en: "What is `git remote` commonly used to manage?"
choices_zh:
  - "遠端儲存庫名稱與網址"
  - "螢幕解析度"
  - "程式語言版本"
choices_en:
  - "Remote repository names and URLs"
  - "Screen resolution"
  - "Programming language versions"
answer: 0
clickbait_zh: "你以為 `origin` 是魔法字，其實它只是你幫遠端取的綽號。"
clickbait_en: "You thought `origin` was a magical keyword. It is usually just a nickname."
review_hint_zh: "`git remote` 管理遠端名稱與位置，例如 `origin`。"
review_hint_en: "`git remote` manages remote names and locations, such as `origin`."
confusion_with: [git-clone-local-copy, git-push-upstream]
metaphor_seed: [通訊錄, 綽號, 連線目標]
hook_style_tags: [misunderstood, comparison]
enabled: true
---

## zh-TW

`git remote` 用來管理你的遠端儲存庫設定，例如遠端的名稱和網址。
像 `origin` 並不是 Git 內建的神秘保留字，而是 clone 時常見的預設遠端名稱。
當你需要查看、加入、修改或移除遠端時，這組指令會是很常見的入口。

## en

`git remote` manages your configured remote repositories, including their names and URLs.
`origin` is not a magical hardcoded concept. It is just the common default remote name created during cloning.
Whenever you need to inspect, add, rename, or remove a remote, this command family is a common starting point.
