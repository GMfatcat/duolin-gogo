---
id: docker-network-container-network
title_zh: docker network 連線範圍
title_en: Docker Network
type: true-false
body_format: bilingual-section
tags: [docker, network, container]
difficulty: 2
question_zh: "Docker network 主要用來決定容器彼此如何連線。"
question_en: "A Docker network mainly defines how containers connect to each other."
answer: true
clickbait_zh: "容器不是天生就互相看得到，背後其實有一層看不見的交通規則。"
clickbait_en: "Containers do not magically see each other. A hidden traffic rule sits underneath."
review_hint_zh: "network 會影響容器之間的連線與名稱解析。"
review_hint_en: "A network controls connectivity and name resolution between containers."
confusion_with: [docker-volume-persist-data]
metaphor_seed: [道路, 線路, 區域]
hook_style_tags: [misunderstood, comparison]
enabled: true
---

## zh-TW

`docker network` 用來管理容器之間的連線關係。
同一個 network 內的容器通常可以彼此用名稱找到對方。
它處理的是通訊路徑，不是資料保存。

## en

`docker network` manages how containers connect to one another.
Containers on the same network can usually reach each other by name.
It is about communication paths, not data persistence.
