---
id: docker-exec-running-container
title_zh: docker exec 進入容器
title_en: Docker Exec
type: single-choice
body_format: bilingual-section
tags: [docker, container, shell]
difficulty: 2
question_zh: "`docker exec` 常見的用途是什麼？"
question_en: "What is a common use of `docker exec`?"
choices_zh:
  - "建立新的 image 並推送到 registry"
  - "在正在執行的容器裡執行命令"
  - "把停止中的容器重新啟動"
choices_en:
  - "Build a new image and push it to a registry"
  - "Run a command inside a running container"
  - "Restart a stopped container"
answer: 1
clickbait_zh: "很多人以為要重跑容器才能進去看，其實不用那麼重。"
clickbait_en: "Many people think they must restart a container just to inspect it. Not necessarily."
review_hint_zh: "`exec` = 在已經 running 的容器裡執行命令。"
review_hint_en: "`exec` runs a command inside an already running container."
confusion_with: [docker-run-start-container]
metaphor_seed: [進控制室, 臨時進場, 進去看]
hook_style_tags: [comparison, safer_first]
enabled: true
---

## zh-TW

`docker exec` 會在一個已經 running 的容器裡執行命令。
很常見的用法是 `docker exec -it <container> sh` 或 `bash`，進去查看檔案、環境變數或程序狀態。
它不會重新建立容器，而是進入既有的執行環境。

## en

`docker exec` runs a command inside a container that is already running.
A common pattern is `docker exec -it <container> sh` or `bash` so you can inspect files, environment variables, or processes.
It does not recreate the container; it enters the existing runtime environment.
