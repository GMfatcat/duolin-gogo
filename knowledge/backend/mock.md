---
id: backend-mock-test-double
title_zh: Mock
title_en: Mock
type: true-false
body_format: bilingual-section
tags: [backend, testing, mock]
difficulty: 2
question_zh: "mock 常用來取代真實依賴，並驗證互動是否符合預期。"
question_en: "A mock is commonly used to replace a real dependency and verify that interactions happen as expected."
answer: true
clickbait_zh: "有時候你不是在測結果，而是在測『它到底有沒有照規矩互動』。"
clickbait_en: "Sometimes you are not testing the output first. You are testing whether the interaction happened the right way."
review_hint_zh: "mock 常同時關心替身與互動驗證。"
review_hint_en: "A mock is often about both substitution and interaction verification."
enabled: true
---

## zh-TW

mock 是一種 test double，除了代替真實依賴外，通常還會驗證某些方法有沒有被呼叫、被叫幾次，或帶了什麼參數。
它很適合測互動邏輯，但過度使用也可能讓測試太綁實作細節。

## en

A mock is a kind of test double that not only replaces a real dependency but often verifies whether certain methods were called, how often, or with which arguments.
It is useful for interaction logic, though overuse can make tests too coupled to implementation details.
