---
id: backend-write-through-vs-write-back
title_zh: Write-through vs Write-back
title_en: Write-through vs Write-back
type: true-false
body_format: bilingual-section
tags: [backend, cache, write-through]
difficulty: 3
question_zh: "write-through 和 write-back 都和寫入 cache 的策略有關，但風險與延遲特性不同。"
question_en: "Write-through and write-back are both cache write strategies, but they differ in risk and latency tradeoffs."
answer: true
clickbait_zh: "寫入快取不是只有快不快，真正麻煩的是資料什麼時候才算真的落地。"
clickbait_en: "Cache writes are not only about speed. The real question is when the data is truly durable."
review_hint_zh: "write-through 比較直接；write-back 通常更快但風險更高。"
review_hint_en: "Write-through is more direct; write-back is often faster but riskier."
enabled: true
---

## zh-TW

write-through 通常表示寫 cache 時也同步寫到主資料來源，所以一致性較直觀。
write-back 則可能先寫 cache，之後再延後刷回主來源，效能可能更好，但失敗風險與一致性複雜度也更高。

## en

Write-through usually means that when data is written to cache, it is also written to the source of truth right away, making consistency easier to reason about.
Write-back may write to cache first and flush later, which can be faster but adds more risk and consistency complexity.
