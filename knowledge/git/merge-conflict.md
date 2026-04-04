---
id: git-merge-conflict-resolution
title: Merge Conflict
title_zh: merge conflict 的處理
title_en: Merge Conflict Resolution
type: true-false
body_format: bilingual-section
tags: [git, merge, conflict]
difficulty: 3
question_zh: "發生 merge conflict 時，Git 代表無法自動決定哪些內容該保留。"
question_en: "When a merge conflict happens, Git cannot decide automatically which content to keep."
answer: true
clickbait_zh: "不是 Git 壞掉，是它把最難的決定丟回給你。"
clickbait_en: "Git is not broken. It is handing the hardest decision back to you."
review_hint_zh: "merge conflict 代表 Git 無法自動整合，需要人手處理。"
review_hint_en: "A merge conflict means Git needs you to resolve the integration manually."
confusion_with: [git-merge-purpose, git-rebase-continue-flow]
metaphor_seed: [裁判, 雙方版本, 最終決定]
hook_style_tags: [fear_of_mistake, consequence, misunderstood]
enabled: true
---

## zh-TW

當 Git 在 merge 過程中遇到兩邊都改了同一段內容，而且它無法安全判斷該保留哪個版本時，就會發生 merge conflict。
這不代表儲存庫壞掉，而是代表自動整合停下來，等你手動選擇最終內容。
通常你要先編輯衝突檔案、確認結果，再 `git add`，最後完成 merge。

## en

Git raises a merge conflict when both sides changed overlapping content and it cannot safely decide which version should win.
That does not mean the repository is broken. It means automatic integration has paused and is waiting for your manual decision.
In practice, you edit the conflicted files, confirm the final result, stage them, and then finish the merge.
