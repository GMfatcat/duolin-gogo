---
id: http-http2-multiplexing
title_zh: HTTP/2
title_en: HTTP/2
type: true-false
body_format: bilingual-section
tags: [http, http2, multiplexing]
difficulty: 2
question_zh: "HTTP/2 的重要特性之一，是能在同一連線上多工處理多個請求與回應。"
question_en: "One important feature of HTTP/2 is the ability to multiplex multiple requests and responses over one connection."
answer: true
clickbait_zh: "不是多開幾條線才叫更快，有時候一條線就能同時跑很多件事。"
clickbait_en: "Faster does not always mean more connections. Sometimes one connection can carry many streams at once."
review_hint_zh: "HTTP/2 的關鍵字之一是 multiplexing。"
review_hint_en: "Multiplexing is one of the key ideas of HTTP/2."
enabled: true
---

## zh-TW

HTTP/2 讓多個請求與回應能在同一條連線上交錯傳輸，這就是常說的 multiplexing。
它的目標之一，是改善 HTTP/1.1 在連線使用效率上的限制。

## en

HTTP/2 allows multiple requests and responses to be interleaved over the same connection, which is commonly described as multiplexing.
One of its goals is to improve the connection efficiency limits of HTTP/1.1.
