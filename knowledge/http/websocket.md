---
id: http-websocket-duplex-connection
title_zh: WebSocket
title_en: WebSocket
type: true-false
body_format: bilingual-section
tags: [http, websocket, realtime]
difficulty: 2
question_zh: "WebSocket 連線建立後，客戶端和伺服器都能主動送訊息。"
question_en: "Once a WebSocket connection is established, both client and server can actively send messages."
answer: true
clickbait_zh: "有些即時功能不是一直輪詢，而是雙方都把話筒打開了。"
clickbait_en: "Some realtime features are not endless polling. Both sides simply keep the microphone on."
review_hint_zh: "WebSocket 是建立後可雙向傳訊的長連線。"
review_hint_en: "WebSocket is a long-lived bidirectional connection."
enabled: true
---

## zh-TW

WebSocket 會先透過 HTTP 升級連線，之後改成長時間存在的雙向通道。
它很適合聊天、即時通知、協作編輯這種需要低延遲推送的場景。

## en

WebSocket starts with an HTTP upgrade and then becomes a long-lived bidirectional channel.
It fits chat, realtime notifications, and collaborative editing where low-latency pushes matter.
