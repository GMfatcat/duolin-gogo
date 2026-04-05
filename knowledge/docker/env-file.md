---
id: docker-env-file-variables
title_zh: Docker env file
title_en: Docker Env File
type: true-false
body_format: bilingual-section
tags: [docker, env, configuration]
difficulty: 2
question_zh: "Docker 的 env file 常用來把環境變數從命令列參數中抽離出來。"
question_en: "A Docker env file is commonly used to move environment variables out of long command arguments."
answer: true
clickbait_zh: "有些命令看起來太長，不是你笨，是設定應該搬家了。"
clickbait_en: "If the command feels too long, the problem may be the configuration, not you."
review_hint_zh: "env file 讓環境變數集中管理。"
review_hint_en: "An env file centralizes environment variable values."
confusion_with: [docker-volume-persist-data, docker-port-mapping-host-container]
metaphor_seed: [設定搬家, 清單化, 集中管理]
hook_style_tags: [safer_first, misunderstood]
enabled: true
---

## zh-TW

env file 通常拿來集中放容器執行時需要的環境變數，讓命令列或 compose 檔不要塞滿一長串設定。
這樣做不代表秘密就自動安全了，但至少設定會更集中、較容易維護。
它主要解決的是設定管理，不是資料持久化或網路連線。

## en

An env file usually stores environment variables needed by a container so you do not have to place every value directly in a long command or Compose file.
This makes configuration easier to read and maintain.
It helps with configuration management, not persistence or networking.
