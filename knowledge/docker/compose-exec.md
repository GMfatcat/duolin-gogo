---
id: docker-compose-exec-service-shell
title_zh: docker compose exec
title_en: Docker Compose Exec
type: true-false
body_format: bilingual-section
tags: [docker, compose, exec]
difficulty: 2
question_zh: "`docker compose exec` 常用來進入 compose 啟動中的某個 service 容器。"
question_en: "`docker compose exec` is commonly used to enter a running service container from a Compose project."
answer: true
clickbait_zh: "不是每次都要自己找 container id，compose 其實早就幫你記好了。"
clickbait_en: "You do not always need the container id. Compose already knows which service you mean."
review_hint_zh: "compose exec 會依 service 名稱進到已啟動的容器。"
review_hint_en: "`compose exec` targets a running container by service name."
confusion_with: [docker-exec-running-container, docker-compose-up-services]
metaphor_seed: [服務名, 指名進入, 已在場]
hook_style_tags: [comparison, safer_first]
enabled: true
---

## zh-TW

`docker compose exec` 會對某個 compose service 對應的執行中容器執行命令，常見用法是開 shell 進去看狀態。
和直接用 `docker exec` 比起來，你不用先手動找出 container id，因為 compose 會依 service 名稱幫你定位。
前提是該 service 容器已經啟動。

## en

`docker compose exec` runs a command inside the running container for a Compose service.
It is often used to open a shell for inspection or debugging.
Unlike plain `docker exec`, you usually target the container by service name instead of looking up the container id yourself.
