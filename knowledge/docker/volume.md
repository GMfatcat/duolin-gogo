---
id: docker-volume-persist-data
title_zh: docker volume 保存資料
title_en: Docker Volume
type: true-false
body_format: bilingual-section
tags: [docker, storage, volume]
difficulty: 2
question_zh: "Docker volume 常用來讓資料在容器重建後仍然保留。"
question_en: "A Docker volume is commonly used to keep data even if a container is recreated."
answer: true
clickbait_zh: "容器可以隨時換掉，但有些資料你根本不能跟著一起蒸發。"
clickbait_en: "Containers can come and go, but some data cannot evaporate with them."
review_hint_zh: "volume 主要解決資料持久化。"
review_hint_en: "A volume mainly solves data persistence."
confusion_with: [docker-network-container-network]
metaphor_seed: [倉庫, 抽屜, 保存]
hook_style_tags: [misunderstood, safer_first]
enabled: true
---

## zh-TW

`docker volume` 常拿來保存資料，不讓它跟容器生命週期一起消失。
即使容器被刪掉再重建，掛在 volume 上的資料仍可保留。
它處理的是資料持久化，不是容器之間的通訊。

## en

`docker volume` is often used to keep data outside a container's lifecycle.
Even if a container is deleted and recreated, data in the volume can remain.
It is about persistence, not container-to-container communication.
