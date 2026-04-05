---
id: docker-rm-remove-container
title_zh: docker rm 刪除容器
title_en: Docker Rm
type: single-choice
body_format: bilingual-section
tags: [docker, container, cleanup]
difficulty: 1
question_zh: "`docker rm` 最主要在做什麼？"
question_en: "What does `docker rm` mainly do?"
choices_zh:
  - "刪除容器"
  - "停止容器"
  - "刪除映像檔"
choices_en:
  - "Remove a container"
  - "Stop a container"
  - "Delete an image"
answer: 0
clickbait_zh: "把容器關掉不代表它消失了，很多人清理時就卡在這裡。"
clickbait_en: "Stopping a container does not make it disappear. Cleanup gets stuck here a lot."
review_hint_zh: "`rm` 會移除容器；停止通常是 `stop`。"
review_hint_en: "`rm` removes a container; stopping is usually `stop`."
confusion_with: [docker-stop-stop-container, docker-images-local-images]
metaphor_seed: [清理, 收走, 丟掉]
hook_style_tags: [comparison, safer_first]
enabled: true
---

## zh-TW

`docker rm` 會把已存在的容器刪掉。
它和 `docker stop` 不同，`stop` 只是停止執行中的容器。
如果你要真正清掉容器本體，通常會用 `rm`。

## en

`docker rm` deletes an existing container.
That is different from `docker stop`, which only stops a running container.
Use `rm` when you want to remove the container itself.
