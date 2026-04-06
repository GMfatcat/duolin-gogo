---
id: backend-dead-letter-queue-failed-messages
title_zh: Dead Letter Queue
title_en: Dead Letter Queue
type: true-false
body_format: bilingual-section
tags: [backend, queue, retry]
difficulty: 3
question_zh: "dead letter queue 常用來收容重試多次後仍失敗的訊息。"
question_en: "A dead letter queue is commonly used to hold messages that still fail after repeated retries."
answer: true
clickbait_zh: "有些訊息不是再試一次就會好，硬塞回去只會讓整條流水線更亂。"
clickbait_en: "Some messages do not get better with one more retry. Pushing them back forever only poisons the whole pipeline."
review_hint_zh: "dead letter queue 會收失敗到不適合再立即重試的訊息。"
review_hint_en: "A dead letter queue stores messages that should stop retrying in the main flow."
enabled: true
---

## zh-TW

dead letter queue 會把多次失敗的訊息移到另一個安全區，避免主流程一直被同一批壞訊息拖住。
它讓團隊可以之後再調查、修補或重放，而不是讓正常工作持續被堵塞。

## en

A dead letter queue moves repeatedly failing messages into a separate safe area so the main flow is not blocked forever by the same bad payloads.
That gives teams a place to inspect, repair, or replay them later.
