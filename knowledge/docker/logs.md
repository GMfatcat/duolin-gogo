---
id: docker-logs-container-output
title_zh: docker logs 查看容器輸出
title_en: Docker Logs
type: single-choice
body_format: bilingual-section
tags: [docker, debugging, logs]
difficulty: 2
question_zh: "`docker logs` 最常用來做什麼？"
question_en: "What is `docker logs` most commonly used for?"
choices_zh:
  - "查看 container 輸出的日誌"
  - "建立新的 image"
  - "把檔案複製進 container"
choices_en:
  - "Read log output from a container"
  - "Create a new image"
  - "Copy files into a container"
answer: 0
clickbait_zh: "容器一片安靜，不代表沒事，有時候答案早就自己寫在 logs 裡。"
clickbait_en: "Silence from a container does not mean nothing happened. Sometimes the answer is already sitting in the logs."
review_hint_zh: "`logs` 主要是看 container 的輸出。"
review_hint_en: "`logs` is mainly for reading container output."
confusion_with: [docker-exec-running-container]
metaphor_seed: [現場錄音, 監視器, 線索]
hook_style_tags: [misunderstood, safer_first]
enabled: true
---

## zh-TW

`docker logs` 會顯示 container 的標準輸出與錯誤輸出，常拿來追查啟動失敗或執行時問題。
它通常是排錯時最先看的地方之一。
如果你想知道 container 剛剛做了什麼、噴了什麼錯，先看 logs 很合理。

## en

`docker logs` shows a container's standard output and standard error streams.
It is commonly used to investigate startup failures or runtime issues.
If you want to know what a container just did or what error it printed, logs are a natural first stop.
