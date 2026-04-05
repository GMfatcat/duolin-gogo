---
id: http-rest-resource-style
title_zh: REST
title_en: REST
type: true-false
body_format: bilingual-section
tags: [http, rest, api]
difficulty: 2
question_zh: "REST 風格常把 API 設計成圍繞資源與標準 HTTP method 的形式。"
question_en: "REST-style APIs commonly organize design around resources and standard HTTP methods."
answer: true
clickbait_zh: "不是 API 都只是幾條網址，很多設計哲學其實都藏在資源和 method 的搭配裡。"
clickbait_en: "An API is not just a set of URLs. A lot of design philosophy hides in how resources and methods are paired."
review_hint_zh: "REST 常圍繞資源與 method 設計。"
review_hint_en: "REST commonly organizes around resources and methods."
enabled: true
---

## zh-TW

REST 常把系統對外能力表達成一組資源，再透過 `GET`、`POST`、`PUT`、`DELETE` 等 method 操作這些資源。
它不只是「回 JSON」，而是一種偏資源導向的 API 風格。

## en

REST often exposes system capabilities as resources and then uses methods like `GET`, `POST`, `PUT`, and `DELETE` to operate on them.
It is more than "returning JSON"; it is a resource-oriented API style.
