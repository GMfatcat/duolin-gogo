---
id: bash-export-env-var
title_zh: export
title_en: Bash export
type: true-false
body_format: bilingual-section
tags: [bash, shell, env]
difficulty: 1
question_zh: "`export` 常用來把 shell 變數變成子程序也看得到的環境變數。"
question_en: "`export` is commonly used to make a shell variable visible as an environment variable to child processes."
answer: true
clickbait_zh: "你明明有設變數，程式卻還是看不到？問題常常就卡在 export。"
clickbait_en: "You set the variable, but the app still cannot see it? The missing piece is often export."
review_hint_zh: "`export` 會讓子程序也能讀到該變數。"
review_hint_en: "`export` exposes the variable to child processes."
enabled: true
---

## zh-TW

在 shell 裡先設定變數，只代表目前這個 shell session 知道它。
加上 `export` 之後，從這個 shell 啟動的子程序也能讀到這個值。

## en

Setting a variable in a shell only makes it available inside that shell by default.
Using `export` makes the value available to child processes started from that shell.
