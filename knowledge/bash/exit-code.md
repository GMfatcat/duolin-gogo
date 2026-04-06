---
id: bash-exit-code-status
title_zh: Exit Code
title_en: Exit Code
type: true-false
body_format: bilingual-section
tags: [bash, shell, process]
difficulty: 2
question_zh: "在 shell 裡，exit code `0` 通常代表成功，非 `0` 代表某種失敗或異常。"
question_en: "In a shell, exit code `0` usually means success, and non-zero usually means some kind of failure or abnormal result."
answer: true
clickbait_zh: "畫面沒報錯不代表真的成功，真正的答案常藏在 exit code。"
clickbait_en: "No visible error does not always mean success. The real answer is often in the exit code."
review_hint_zh: "`0` 通常表示成功。"
review_hint_en: "`0` usually means success."
enabled: true
---

## zh-TW

很多 shell 流程控制都依賴 exit code，而不是單純看畫面輸出了什麼。
慣例上 `0` 表示成功，其他值則表示不同類型的失敗或特殊狀態。

## en

Many shell control-flow decisions depend on the exit code rather than just printed output.
By convention, `0` means success and other values indicate failure or special conditions.
