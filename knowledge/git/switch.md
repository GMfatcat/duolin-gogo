---
id: git-switch-branch
title: Git Switch
title_zh: git switch 的用途
title_en: Git Switch
type: true-false
body_format: bilingual-section
tags: [git, branches]
difficulty: 1
question_zh: "`git switch` 是比較新的分支切換指令，用來取代部分 `git checkout` 的用途。"
question_en: "`git switch` is a newer branch-switching command that replaces part of `git checkout`."
answer: true
clickbait_zh: "Git 後來把這件事拆開，就是因為太多人切錯"
clickbait_en: "Git split this action out for a reason. Too many people were using the old command wrong."
review_hint_zh: "`git switch` 專注在切分支，比 `checkout` 更明確。"
review_hint_en: "`git switch` focuses on branch switching and is clearer than `checkout`."
confusion_with: [git-checkout-legacy]
metaphor_seed: [換軌, 切車道, 指令分工]
hook_style_tags: [comparison, misunderstood, safer-first]
enabled: true
---

## zh-TW

`git switch` 是較新的 Git 指令，專門拿來切換分支。
它的出現就是為了讓分支切換這件事更清楚，不要再跟還原檔案的操作混在一起。
如果你只是想切分支，通常 `switch` 會比 `checkout` 更直觀。

## en

`git switch` is a newer Git command focused on branch switching.
It was introduced to make branch switching clearer instead of mixing it with file restoration behavior.
If your intention is only to move between branches, `switch` is usually more explicit than `checkout`.
