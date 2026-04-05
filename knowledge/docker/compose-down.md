---
id: docker-compose-down-services
title_zh: docker compose down 停掉整組服務
title_en: Docker Compose Down
type: true-false
body_format: bilingual-section
tags: [docker, compose, cleanup]
difficulty: 2
question_zh: "`docker compose down` 常用來停止並移除 compose 啟動的服務資源。"
question_en: "`docker compose down` is commonly used to stop and remove compose-managed resources."
answer: true
clickbait_zh: "服務停了不代表現場真的收乾淨，少按這個常常留下一堆東西。"
clickbait_en: "Stopping services is not the same as cleaning the whole scene. This command is where many people stop too early."
review_hint_zh: "`compose down` 會停掉並清掉 compose 啟動的資源。"
review_hint_en: "`compose down` stops and removes compose-managed resources."
confusion_with: [docker-compose-up-services]
metaphor_seed: [收攤, 清場, 整組帶走]
hook_style_tags: [safer_first, comparison]
enabled: true
---

## zh-TW

`docker compose down` 會停止 compose 啟動的服務，並移除相關 container、network 等資源。
如果只是暫時停掉服務，有時候會用其他指令；但當你想把這一組環境收掉時，通常會想到 `down`。
它比單純停止某一個 container 更像整組收攤。

## en

`docker compose down` stops services started by Docker Compose and removes related resources such as containers and networks.
If you want to tear down the whole stack instead of only pausing a process, `down` is the usual command.
It is more like shutting down the entire setup than stopping a single container.
