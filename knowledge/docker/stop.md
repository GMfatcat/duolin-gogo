---
id: docker-stop-stop-container
title_zh: docker stop 停止容器
title_en: Docker Stop
type: true-false
body_format: bilingual-section
tags: [docker, container, runtime]
difficulty: 1
question_zh: "`docker stop` 主要是讓執行中的容器停下來。"
question_en: "`docker stop` is mainly used to stop a running container."
answer: true
clickbait_zh: "先別急著刪，很多時候你只是想讓它安靜下來。"
clickbait_en: "Do not delete it yet. Sometimes you only want it to go quiet."
review_hint_zh: "`stop` 停的是執行狀態，不是直接刪除容器。"
review_hint_en: "`stop` changes the running state; it does not remove the container."
confusion_with: [docker-rm-remove-container]
metaphor_seed: [暫停, 關機, 安靜]
hook_style_tags: [comparison, misunderstood]
enabled: true
---

## zh-TW

`docker stop` 會要求執行中的容器停止。
它保留容器本身，之後還可以再啟動。
如果你只是想先停掉服務，不需要立刻刪除容器。

## en

`docker stop` tells a running container to stop.
The container still exists and can be started again later.
Use it when you want the service to stop without deleting the container.
