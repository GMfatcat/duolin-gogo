---
id: bash-shell-vs-env-variable
title_zh: Shell Variable vs Env Var
title_en: Shell Variable vs Environment Variable
type: true-false
body_format: bilingual-section
tags: [bash, shell, env]
difficulty: 2
question_zh: "一般 shell 變數不一定會自動傳給子程序，環境變數則會。"
question_en: "A normal shell variable is not automatically inherited by child processes, but an environment variable is."
answer: true
clickbait_zh: "你以為有設就算有設，但 shell 跟環境變數其實不是同一件事。"
clickbait_en: "You thought a variable was set, but shell variables and environment variables are not the same thing."
review_hint_zh: "差別在於子程序能不能看到它。"
review_hint_en: "The difference is whether child processes can see it."
enabled: true
---

## zh-TW

shell variable 常只存在於目前 shell 自己的上下文裡。
environment variable 則會跟著程序環境一起傳給從它啟動的子程序。

## en

A shell variable usually lives only inside the current shell context.
An environment variable is carried into child processes started from that shell.
