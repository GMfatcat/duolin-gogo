---
id: http-tls-handshake-trust-setup
title_zh: TLS Handshake
title_en: TLS Handshake
type: true-false
body_format: bilingual-section
tags: [http, tls, handshake]
difficulty: 3
question_zh: "TLS handshake 的作用之一，是在正式傳輸前先建立加密與信任基礎。"
question_en: "One role of the TLS handshake is to establish encryption and trust basics before normal data transfer begins."
answer: true
clickbait_zh: "小鎖不是突然出現的，前面其實先跑了一段彼此確認身份的流程。"
clickbait_en: "That padlock does not appear by magic. A trust-establishing conversation happened first."
review_hint_zh: "TLS handshake 是加密與驗證前置流程。"
review_hint_en: "The TLS handshake is the setup phase for encryption and trust."
enabled: true
---

## zh-TW

在真正傳送應用資料之前，TLS 會先做握手流程，協商參數並建立信任與加密基礎。
這一步通常也和憑證驗證密切相關。

## en

Before application data is sent, TLS performs a handshake to negotiate parameters and establish the basics for trust and encryption.
This step is also closely tied to certificate verification.
