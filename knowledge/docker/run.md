---
id: docker-run-start-container
title_zh: docker run 啟動容器
title_en: Docker Run
type: single-choice
body_format: bilingual-section
tags: [docker, container, runtime]
difficulty: 1
question_zh: "`docker run` 最主要在做什麼？"
question_en: "What does `docker run` mainly do?"
choices_zh:
  - "列出目前所有正在執行的容器"
  - "從 image 建立並啟動一個容器"
  - "進入一個已經存在的容器 shell"
choices_en:
  - "List all currently running containers"
  - "Create and start a container from an image"
  - "Open a shell inside an existing container"
answer: 1
clickbait_zh: "你以為 Docker 第一個指令是 build？很多人其實先卡在 run。"
clickbait_en: "Think `build` is the first Docker command that matters? Many people get stuck on `run` first."
review_hint_zh: "`run` = 用 image 建立並啟動容器。"
review_hint_en: "`run` creates and starts a container from an image."
confusion_with: [docker-exec-running-container, docker-ps-running-containers]
metaphor_seed: [開機, 發車, 啟動]
hook_style_tags: [misunderstood, safer_first]
enabled: true
---

## zh-TW

`docker run` 會根據指定的 image 建立一個新的 container，然後立刻啟動它。
如果 image 還不存在於本機，Docker 也可能先幫你下載。
它是把 image 真正變成可執行環境的常用入口。

## en

`docker run` creates a new container from the given image and starts it immediately.
If the image is not available locally, Docker may pull it first.
It is the common entry point that turns an image into a running environment.
