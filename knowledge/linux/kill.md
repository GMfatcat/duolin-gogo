---
id: linux-kill-stop-process
title_zh: kill 終止程序
title_en: Kill
type: true-false
body_format: bilingual-section
tags: [linux, shell, process]
difficulty: 2
question_zh: "`kill` 常用來把指定的 process 停止掉。"
question_en: "`kill` is commonly used to stop a specified process."
answer: true
clickbait_zh: "關掉視窗不一定真的停掉程式，有時候你還得找到 process 親手收掉。"
clickbait_en: "Closing a window does not always stop the program. Sometimes you still have to find the process and end it yourself."
review_hint_zh: "`kill` 常用來送訊號給 process，讓它停止。"
review_hint_en: "`kill` commonly sends a signal to a process so it can stop."
confusion_with: [linux-ps-process-list]
metaphor_seed: [斷電, 收尾, 停車]
hook_style_tags: [safer_first, comparison]
enabled: true
---

## zh-TW

`kill` 常用來送訊號給指定的 process，最常見的用途就是把卡住或不需要的程序停掉。
它通常會搭配 `ps` 先找到 process ID，再決定要送出什麼訊號。
雖然名字叫 kill，但實際上它做的是送 signal。

## en

`kill` is commonly used to send a signal to a specific process, often to stop a stuck or unnecessary program.
It is usually paired with `ps` so you can find the process ID first.
Even though the name says kill, the command is really about sending signals.
