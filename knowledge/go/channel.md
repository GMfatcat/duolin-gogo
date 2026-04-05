---
id: go-channel-communication
title_zh: channel 傳遞資料
title_en: Channel
type: single-choice
body_format: bilingual-section
tags: [go, concurrency, communication]
difficulty: 2
question_zh: "在 Go 裡，channel 常見的用途是什麼？"
question_en: "What is a common use of a channel in Go?"
choices_zh:
  - "在 goroutine 之間傳遞資料或同步"
  - "用來宣告 interface 方法"
  - "直接取代所有 slice"
choices_en:
  - "Pass data or synchronize between goroutines"
  - "Declare interface methods"
  - "Replace all slices directly"
answer: 0
clickbait_zh: "只會開 goroutine 不夠，很多人真正卡住的是資料怎麼安全傳回來。"
clickbait_en: "Starting goroutines is not the hard part. Many people really get stuck on how data comes back safely."
review_hint_zh: "channel 常拿來傳資料和做同步。"
review_hint_en: "Channels are commonly used for data passing and synchronization."
confusion_with: [go-goroutine-concurrency]
metaphor_seed: [傳送帶, 管道, 交接]
hook_style_tags: [comparison, misunderstood, safer_first]
enabled: true
---

## zh-TW

channel 是 Go 用來在 goroutine 之間傳遞資料與同步節奏的重要機制。
你可以把它想成一條管道，一邊送資料，另一邊接資料。
很多 Go 的並行寫法都會搭配 goroutine 和 channel 一起使用。

## en

A channel is one of Go's core tools for passing data and synchronizing work between goroutines.
You can think of it as a pipe: one side sends data and the other side receives it.
Many concurrent Go patterns use goroutines together with channels.
