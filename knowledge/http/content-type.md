---
id: http-content-type-media-type
title_zh: Content-Type
title_en: HTTP Content-Type
type: true-false
body_format: bilingual-section
tags: [http, content-type, headers]
difficulty: 1
question_zh: "`Content-Type` 常用來說明 body 是什麼格式，例如 JSON 或 HTML。"
question_en: "`Content-Type` is commonly used to describe the body format, such as JSON or HTML."
answer: true
clickbait_zh: "資料不只要送出去，還要先讓對方知道自己到底是什麼。"
clickbait_en: "Sending data is not enough. The other side also needs to know what it is."
review_hint_zh: "`Content-Type` 說明 payload 的媒體型別。"
review_hint_en: "`Content-Type` tells the media type of the payload."
enabled: true
---

## zh-TW

`Content-Type` 是很常見的 header，用來描述 body 的媒體型別，例如 `application/json` 或 `text/html`。
如果這個資訊錯了，接收方可能會用錯方式解讀資料。

## en

`Content-Type` is a very common header used to describe the media type of the body, such as `application/json` or `text/html`.
If it is wrong, the receiver may parse the data in the wrong way.
