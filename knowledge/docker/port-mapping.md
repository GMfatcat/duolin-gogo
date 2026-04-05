---
id: docker-port-mapping-host-container
title_zh: Docker port mapping
title_en: Docker Port Mapping
type: true-false
body_format: bilingual-section
tags: [docker, network, port]
difficulty: 2
question_zh: "port mapping 的目的之一，是讓主機流量能轉到容器內的服務埠。"
question_en: "One purpose of port mapping is to send host traffic to a service port inside the container."
answer: true
clickbait_zh: "服務明明有啟動，外面卻打不到？很多時候不是程式掛了，是門根本沒開。"
clickbait_en: "If the service is running but unreachable, the issue may be the door, not the app."
review_hint_zh: "port mapping 把 host port 對到 container port。"
review_hint_en: "Port mapping connects a host port to a container port."
confusion_with: [docker-network-container-network]
metaphor_seed: [開門, 對外入口, 轉接]
hook_style_tags: [fear_of_mistake, misunderstood]
enabled: true
---

## zh-TW

Docker 容器裡的服務就算正在監聽埠號，也不代表主機外面就能直接連到它。
port mapping 會把主機上的某個 port 轉到容器裡對應的 port，讓流量能進入容器服務。
這比較像對外開入口，不是讓容器彼此互通的 network 設定。

## en

Even if a service is listening inside a container, that does not automatically make it reachable from the host.
Port mapping forwards a host port to a port inside the container so traffic can reach the service.
This is about exposing an entry point, not about container-to-container networking.
