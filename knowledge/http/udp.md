---
id: http-udp-datagram-transport
title_zh: UDP
title_en: UDP
type: true-false
body_format: bilingual-section
tags: [http, udp, transport]
difficulty: 2
question_zh: "UDP 通常不像 TCP 那樣保證可靠送達與順序。"
question_en: "UDP usually does not guarantee reliable delivery and ordering the way TCP does."
answer: true
clickbait_zh: "比較快的路不一定幫你顧完整，UDP 就是這種取捨代表。"
clickbait_en: "A faster path does not always protect completeness. UDP is a classic example of that tradeoff."
review_hint_zh: "UDP 偏輕量，但不保證可靠與有序。"
review_hint_en: "UDP is lightweight but does not guarantee reliability or order."
enabled: true
---

## zh-TW

UDP 是比較輕量的傳輸協定，不像 TCP 那樣提供完整的可靠送達與順序保證。
它常見於某些更重視低延遲的場景。

## en

UDP is a lighter transport protocol and does not offer the same reliable, ordered guarantees as TCP.
It often appears in scenarios where low latency matters more than strict delivery guarantees.
