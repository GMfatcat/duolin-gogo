---
id: http-http1-connection-model
title_zh: HTTP/1.1
title_en: HTTP/1.1
type: true-false
body_format: bilingual-section
tags: [http, http1, protocol]
difficulty: 2
question_zh: "HTTP/1.1 的常見限制之一，是同一條連線上的請求處理效率不如後來的 multiplexing 模型。"
question_en: "One common limitation of HTTP/1.1 is that request handling on a connection is less efficient than later multiplexed models."
answer: true
clickbait_zh: "不是網站慢就一定是伺服器爛，有時候是老協定本來就沒那麼俐落。"
clickbait_en: "A slow web experience does not always mean a weak server. Sometimes the older protocol model is just less elegant."
review_hint_zh: "HTTP/1.1 沒有像 HTTP/2 那樣的 multiplexing。"
review_hint_en: "HTTP/1.1 does not offer multiplexing like HTTP/2."
enabled: true
---

## zh-TW

HTTP/1.1 已經比更早版本成熟很多，但它在同一連線上處理多個請求時的效率仍有限。
這也是後來 HTTP/2 想改善的重要方向之一。

## en

HTTP/1.1 is much more mature than earlier versions, but its efficiency for handling multiple requests on one connection is still limited.
That is one of the important problems HTTP/2 tried to improve.
