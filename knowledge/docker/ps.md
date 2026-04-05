---
id: docker-ps-running-containers
title_zh: docker ps 查看容器
title_en: Docker PS
type: true-false
body_format: bilingual-section
tags: [docker, container, inspection]
difficulty: 1
question_zh: "`docker ps` 預設會列出目前正在執行的容器。"
question_en: "`docker ps` lists currently running containers by default."
answer: true
clickbait_zh: "找不到容器不一定是它消失了，你可能只是看錯了清單。"
clickbait_en: "If your container seems gone, you may just be looking at the wrong list."
review_hint_zh: "`docker ps` 預設只看正在跑的容器。"
review_hint_en: "`docker ps` shows running containers by default."
confusion_with: [docker-run-start-container]
metaphor_seed: [值班名單, 現場清單]
hook_style_tags: [comparison, misunderstood]
enabled: true
---

## zh-TW

`docker ps` 預設只顯示目前正在執行中的容器。
如果你想看到包含停止狀態在內的所有容器，通常會改用 `docker ps -a`。
這個差異很常讓人誤以為容器不見了。

## en

`docker ps` shows only containers that are currently running.
If you want to include stopped containers, you usually use `docker ps -a`.
That difference often makes people think a container has disappeared.
