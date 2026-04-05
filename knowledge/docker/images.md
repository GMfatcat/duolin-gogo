---
id: docker-images-local-images
title_zh: docker images 查看本機映像
title_en: Docker Images
type: true-false
body_format: bilingual-section
tags: [docker, image, inspection]
difficulty: 1
question_zh: "`docker images` 會列出目前本機可用的 images。"
question_en: "`docker images` lists images available on the local machine."
answer: true
clickbait_zh: "不是每次出錯都要怪 registry，有時候你只是根本不知道本機現在有什麼 image。"
clickbait_en: "Not every problem starts at the registry. Sometimes you simply do not know what images you already have."
review_hint_zh: "`docker images` 是看本機 image 清單。"
review_hint_en: "`docker images` shows local images."
confusion_with: [docker-ps-running-containers]
metaphor_seed: [倉庫, 庫存, 清單]
hook_style_tags: [misunderstood, comparison]
enabled: true
---

## zh-TW

`docker images` 會列出本機目前有的 image，例如 repository、tag、image ID 與大小。
它是在看映像庫存，不是在看目前正在跑的 container。
如果你想知道某個 image 有沒有先被拉下來，這個指令很常用。

## en

`docker images` lists the images currently stored on your machine, including repository, tag, image ID, and size.
It shows your local image inventory rather than running containers.
It is useful when you want to check whether an image already exists locally.
