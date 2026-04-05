---
id: go-goroutine-concurrency
title_zh: goroutine 並行執行
title_en: Goroutine
type: single-choice
body_format: bilingual-section
tags: [go, concurrency, runtime]
difficulty: 2
question_zh: "在 Go 裡，goroutine 最主要是用來做什麼？"
question_en: "What is the main purpose of a goroutine in Go?"
choices_zh:
  - "定義新的 struct 型別"
  - "讓函式以並行方式執行"
  - "把套件編譯成執行檔"
choices_en:
  - "Define a new struct type"
  - "Run a function concurrently"
  - "Compile a package into an executable"
answer: 1
clickbait_zh: "你以為 Go 變快只是因為它夠快？很多時候是因為你敢不敢開 goroutine。"
clickbait_en: "Think Go feels fast only because the language is fast? Often it is about whether you use goroutines well."
review_hint_zh: "goroutine = 輕量並行執行單位。"
review_hint_en: "A goroutine is a lightweight concurrent execution unit."
confusion_with: [go-channel-communication]
metaphor_seed: [分身, 多線同時跑, 分工]
hook_style_tags: [misunderstood, comparison]
enabled: true
---

## zh-TW

goroutine 是 Go 裡非常輕量的並行執行單位。
只要在函式呼叫前加上 `go`，那個函式就可以和其他工作同時進行。
它不是作業系統層級的 thread，但常被拿來處理需要同時做多件事的情境。

## en

A goroutine is a very lightweight unit of concurrent execution in Go.
When you put `go` before a function call, that function can run alongside other work.
It is not the same thing as an operating system thread, but it is often used for doing multiple tasks at the same time.
