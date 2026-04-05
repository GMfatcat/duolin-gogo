---
id: backend-eventual-consistency-delay
title_zh: Eventual Consistency
title_en: Eventual Consistency
type: true-false
body_format: bilingual-section
tags: [backend, consistency, distributed]
difficulty: 3
question_zh: "eventual consistency 的核心之一，是系統短時間內可能不一致，但最終會收斂。"
question_en: "One core idea of eventual consistency is that a system may be temporarily inconsistent but will converge later."
answer: true
clickbait_zh: "不是所有『剛剛怎麼看不到』都代表壞掉，有些只是還沒同步完。"
clickbait_en: "Not every \"why can't I see it yet\" means the system is broken. Sometimes it just has not caught up."
review_hint_zh: "eventual consistency 接受短暫不一致。"
review_hint_en: "Eventual consistency accepts temporary inconsistency."
enabled: true
---

## zh-TW

有些分散式系統不追求每個節點立刻同步一致，而是接受短時間差異，之後再慢慢收斂。
這就是 eventual consistency 的典型心智模型。

## en

Some distributed systems do not require every node to become consistent immediately.
Instead, they accept short-lived differences and converge later.
That is the typical mental model behind eventual consistency.
