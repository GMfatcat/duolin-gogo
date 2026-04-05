---
id: docker-build-image
title_zh: docker build 建立映像
title_en: Docker Build
type: single-choice
body_format: bilingual-section
tags: [docker, image, build]
difficulty: 2
question_zh: "`docker build` 主要在做什麼？"
question_en: "What does `docker build` mainly do?"
choices_zh:
  - "把現有 container 重新啟動"
  - "根據 Dockerfile 建立 image"
  - "查看 image 的執行紀錄"
choices_en:
  - "Restart an existing container"
  - "Build an image from a Dockerfile"
  - "Inspect logs from an image"
answer: 1
clickbait_zh: "你以為 Docker 的起點是 run？很多流程其實早在 build 就決定了。"
clickbait_en: "Think Docker starts at `run`? A lot is decided much earlier at `build`."
review_hint_zh: "`build` 會根據 Dockerfile 產生 image。"
review_hint_en: "`build` turns a Dockerfile into an image."
confusion_with: [docker-run-start-container, docker-images-local-images]
metaphor_seed: [烤箱, 食譜, 成品]
hook_style_tags: [misunderstood, comparison]
enabled: true
---

## zh-TW

`docker build` 會根據目前目錄中的 Dockerfile 與建置內容，產生一個新的 image。
它是在準備可重複使用的映像，不是在直接啟動 container。
常見流程是先 build，再用 `docker run` 啟動。

## en

`docker build` creates a new image from a Dockerfile and the files in the build context.
It prepares a reusable image rather than launching a container directly.
A common workflow is to build first, then run the resulting image.
