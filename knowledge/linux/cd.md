---
id: linux-cd-change-directory
title_zh: cd 切換目錄
title_en: Change Directory
type: true-false
body_format: bilingual-section
tags: [linux, shell, navigation]
difficulty: 1
question_zh: "`cd` 用來切換目前所在的目錄。"
question_en: "`cd` is used to change the current directory."
answer: true
clickbait_zh: "有時候不是指令壞掉，是你人根本走到錯的資料夾了。"
clickbait_en: "Sometimes the command is fine. You are just standing in the wrong folder."
review_hint_zh: "`cd` = change directory。"
review_hint_en: "`cd` means change directory."
confusion_with: [linux-pwd-current-directory, linux-ls-list-files]
metaphor_seed: [走路, 換房間, 移動位置]
hook_style_tags: [safer_first, comparison]
enabled: true
---

## zh-TW

`cd` 是 shell 中最基本的移動指令之一，用來切換目前工作目錄。
很多後續操作都依賴你現在站在哪個資料夾，所以 `cd` 和 `pwd` 常常要一起理解。
當你跑錯地方時，其他指令看起來也會一起出錯。

## en

`cd` is one of the most basic navigation commands in a shell.
It changes your current working directory, which affects where many later commands operate.
If you are in the wrong folder, lots of other commands can appear broken too.
