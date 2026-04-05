---
id: go-package-module-boundary
title_zh: package 程式邊界
title_en: Go Package
type: single-choice
body_format: bilingual-section
tags: [go, package, organization]
difficulty: 2
question_zh: "Go 的 `package` 最主要在幫助什麼？"
question_en: "What does a Go `package` mainly help with?"
choices_zh:
  - "組織程式碼與定義可見邊界"
  - "建立執行緒"
  - "把切片轉成 map"
choices_en:
  - "Organize code and define visibility boundaries"
  - "Create threads"
  - "Turn slices into maps"
answer: 0
clickbait_zh: "不是檔案放一起就算同一塊，真正的邊界其實是 package。"
clickbait_en: "Files sitting together are not the whole story. The real boundary is the package."
review_hint_zh: "`package` 幫助組織程式碼與控制識別範圍。"
review_hint_en: "`package` helps organize code and control visibility."
confusion_with: [go-struct-field-grouping]
metaphor_seed: [邊界, 房間, 區塊]
hook_style_tags: [misunderstood, comparison]
enabled: true
---

## zh-TW

Go 的 `package` 用來把相關程式碼組在一起。
它同時也影響哪些識別名稱可以被外部使用。
理解 package 有助於讀懂程式結構與依賴邊界。

## en

A Go `package` groups related code together.
It also affects which names are visible outside that package.
Understanding packages helps you read structure and dependency boundaries more clearly.
