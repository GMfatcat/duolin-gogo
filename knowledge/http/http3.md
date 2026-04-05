---
id: http-http3-quic-transport
title_zh: HTTP/3
title_en: HTTP/3
type: true-false
body_format: bilingual-section
tags: [http, http3, quic]
difficulty: 3
question_zh: "HTTP/3 的重要特徵之一，是它建立在 QUIC/UDP 之上，而不是傳統 TCP。"
question_en: "One important characteristic of HTTP/3 is that it is built on QUIC over UDP rather than traditional TCP."
answer: true
clickbait_zh: "你以為新版 HTTP 只是更快，其實底下連交通工具都換了。"
clickbait_en: "You may think the new HTTP is just faster, but the transport underneath changed too."
review_hint_zh: "HTTP/3 與 QUIC/UDP 強相關。"
review_hint_en: "HTTP/3 is tightly tied to QUIC over UDP."
enabled: true
---

## zh-TW

HTTP/3 和前面版本最大的不同之一，是它建立在 QUIC 之上，而 QUIC 又使用 UDP。
這讓它在連線建立與某些延遲問題上有新的改善空間。

## en

One of the biggest differences in HTTP/3 is that it runs on top of QUIC, and QUIC uses UDP.
That creates new opportunities to improve connection setup and some latency-related behavior.
