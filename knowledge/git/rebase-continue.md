---
id: git-rebase-continue-flow
title: Git Rebase Continue
title_zh: git rebase --continue 的時機
title_en: Git Rebase Continue
type: true-false
body_format: bilingual-section
tags: [git, rebase, conflict]
difficulty: 3
question_zh: "解完 rebase 過程中的衝突並重新 `git add` 後，常見下一步是 `git rebase --continue`。"
question_en: "After resolving conflicts during a rebase and staging the result again, a common next step is `git rebase --continue`."
answer: true
clickbait_zh: "卡在 rebase 半路時，不是重來，而是要告訴 Git 可以繼續走。"
clickbait_en: "When rebase stalls halfway, the next step is often not restart, but continue."
review_hint_zh: "rebase 衝突解完後，常用 `git rebase --continue` 繼續流程。"
review_hint_en: "After fixing rebase conflicts, `git rebase --continue` is the usual next step."
confusion_with: [git-merge-conflict-resolution, git-rebase-vs-merge]
metaphor_seed: [續播, 接關, 半路恢復]
hook_style_tags: [consequence, misunderstood]
enabled: true
---

## zh-TW

當 `git rebase` 途中碰到衝突時，Git 會先停下來，等你處理完目前這一步。
你通常要先修好衝突內容、把結果重新 `git add`，然後用 `git rebase --continue` 讓 rebase 繼續往下跑。
如果不理解這個節奏，很多人會誤以為整個 rebase 失敗了，其實常常只是卡在中間等待處理。

## en

When `git rebase` hits a conflict, Git pauses the process and waits for you to resolve the current step.
You usually fix the conflict, stage the resolved files again, and then run `git rebase --continue` to move forward.
Without that mental model, many people think the rebase has failed completely, when it is often just paused in the middle.
