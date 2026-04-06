---
id: bash-and-vs-or-command-chain
title_zh: "&& vs ||"
title_en: "&& vs ||"
type: single-choice
body_format: bilingual-section
tags: [bash, shell, control-flow]
difficulty: 2
question_zh: "在 shell 裡，`cmd1 && cmd2` 最常代表什麼？"
question_en: "In a shell, what does `cmd1 && cmd2` most commonly mean?"
choices_zh:
  - "無論 cmd1 是否成功，cmd2 都會執行"
  - "只有 cmd1 成功時，cmd2 才執行"
  - "只有 cmd1 失敗時，cmd2 才執行"
choices_en:
  - "cmd2 runs no matter whether cmd1 succeeds"
  - "cmd2 runs only if cmd1 succeeds"
  - "cmd2 runs only if cmd1 fails"
answer: 1
clickbait_zh: "你以為只是把兩個指令接起來？其實 `&&` 和 `||` 會偷偷改變流程。"
clickbait_en: "You thought you were only chaining commands, but `&&` and `||` quietly change the control flow."
review_hint_zh: "`&&` 偏成功路徑，`||` 偏失敗路徑。"
review_hint_en: "`&&` is for the success path and `||` is for the failure path."
enabled: true
---

## zh-TW

`cmd1 && cmd2` 通常表示只有在 `cmd1` 成功時才執行 `cmd2`。
相對地，`cmd1 || cmd2` 則常用來表達 `cmd1` 失敗後再執行 `cmd2`。

## en

`cmd1 && cmd2` usually means `cmd2` runs only if `cmd1` succeeds.
By contrast, `cmd1 || cmd2` is commonly used to run `cmd2` when `cmd1` fails.
