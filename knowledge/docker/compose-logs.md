---
id: docker-compose-logs-service-output
title_zh: docker compose logs
title_en: Docker Compose Logs
type: true-false
body_format: bilingual-section
tags: [docker, compose, logs]
difficulty: 2
question_zh: "`docker compose logs` 可以查看 compose 專案中一個或多個 service 的輸出。"
question_en: "`docker compose logs` can show the output of one or more services in a Compose project."
answer: true
clickbait_zh: "你以為 log 只能一個 container 一個 container 看？compose 其實可以整組拉出來。"
clickbait_en: "You do not have to inspect logs one container at a time. Compose can pull the whole service view together."
review_hint_zh: "compose logs 會聚合 service 容器輸出。"
review_hint_en: "`compose logs` aggregates output by service."
confusion_with: [docker-logs-container-output]
metaphor_seed: [整組輸出, 聚合視角, 同場回放]
hook_style_tags: [comparison, misunderstood]
enabled: true
---

## zh-TW

`docker compose logs` 會把 compose 專案裡各個 service 的輸出整理出來，你也可以只指定某個 service。
這在多容器開發環境中特別好用，因為你能更快看到 web、db、worker 之間發生了什麼。
如果想持續追蹤，也可以搭配 follow 類選項一起看。

## en

`docker compose logs` shows the output from services in a Compose project, and you can narrow it to specific services.
It is especially useful when several containers work together and you want one combined view of what is happening.
You can also keep following the output as new log lines appear.
