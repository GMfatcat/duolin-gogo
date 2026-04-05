---
id: http-put-vs-patch-update
title_zh: PUT vs PATCH
title_en: PUT vs PATCH
type: true-false
body_format: bilingual-section
tags: [http, method, put, patch]
difficulty: 2
question_zh: "`PUT` 和 `PATCH` 都能更新資源，但語意上通常不是完全一樣。"
question_en: "`PUT` and `PATCH` can both update a resource, but their semantics are usually not identical."
answer: true
clickbait_zh: "看起來都叫更新，但很多 API 的坑就埋在『更新方式』不一樣。"
clickbait_en: "They both look like update methods, but many API bugs hide in the difference between how they update."
review_hint_zh: "`PUT` 偏完整替換，`PATCH` 偏部分修改。"
review_hint_en: "`PUT` leans toward full replacement, while `PATCH` leans toward partial modification."
enabled: true
---

## zh-TW

雖然 `PUT` 和 `PATCH` 都用來更新資源，但常見語意是：`PUT` 偏完整替換，`PATCH` 偏局部修改。
實際 API 設計不一定完全一致，但理解這個差異很重要。

## en

Although both `PUT` and `PATCH` update resources, the common semantic distinction is that `PUT` leans toward full replacement while `PATCH` leans toward partial modification.
APIs do not always implement them perfectly, but the distinction is still important.
