---
id: docker-compose-up-services
title_zh: docker compose up 啟動服務
title_en: Docker Compose Up
type: true-false
body_format: bilingual-section
tags: [docker, compose, services]
difficulty: 2
question_zh: "`docker compose up` 常用來一起建立並啟動多個服務。"
question_en: "`docker compose up` is commonly used to create and start multiple services together."
answer: true
clickbait_zh: "一個指令把整組服務叫起來，很多人卻還在一台一台手動開。"
clickbait_en: "One command can bring up the whole stack, but many people still start services one by one."
review_hint_zh: "`compose up` 常用來依設定檔一起啟動多服務。"
review_hint_en: "`compose up` commonly starts multiple services from a compose file."
confusion_with: [docker-run-start-container]
metaphor_seed: [整隊出發, 一起開場, 成套啟動]
hook_style_tags: [misunderstood, comparison]
enabled: true
---

## zh-TW

`docker compose up` 會根據 `compose.yaml` 或 `docker-compose.yml` 的設定，一次建立並啟動多個相關服務。
它很適合本機開發環境，像是同時拉起 web、db、cache 這樣的組合。
和單純的 `docker run` 相比，它更強調多服務協作。

## en

`docker compose up` uses a compose file such as `compose.yaml` or `docker-compose.yml` to create and start multiple related services together.
It is especially useful for local development stacks such as web, database, and cache services.
Compared with plain `docker run`, it focuses on coordinated multi-service setups.
