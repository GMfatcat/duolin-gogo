---
id: backend-outbox-pattern-reliable-events
title_zh: Outbox Pattern
title_en: Outbox Pattern
type: true-false
body_format: bilingual-section
tags: [backend, event, consistency]
difficulty: 3
question_zh: "outbox pattern 常用來讓資料庫變更和事件發送更可靠地對齊。"
question_en: "The outbox pattern is commonly used to align database changes and event publishing more reliably."
answer: true
clickbait_zh: "資料存成功但事件沒送出，是很多系統最安靜也最危險的裂縫。"
clickbait_en: "A row saved without its event being published is one of the quietest and most dangerous cracks in backend systems."
review_hint_zh: "outbox pattern 讓資料變更和事件發送透過同一筆持久化銜接。"
review_hint_en: "The outbox pattern bridges data changes and event publication through durable shared persistence."
enabled: true
---

## zh-TW

outbox pattern 會先把事件和業務資料一起寫入資料庫，再由另一個流程把 outbox 裡的事件送出去。
這能降低「資料已寫入，但訊息沒送出」的風險，是常見的可靠事件整合手法。

## en

The outbox pattern writes the business change and the event record into the database first, then a separate process publishes from the outbox.
It reduces the risk of “data committed but event missing,” which is a common reliability problem in event-driven systems.
