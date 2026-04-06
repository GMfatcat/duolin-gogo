---
id: sql-cursor-pagination-stable-page
title_zh: Cursor Pagination
title_en: Cursor Pagination
type: single-choice
body_format: bilingual-section
tags: [sql, pagination]
difficulty: 3
question_zh: "cursor pagination 相較 offset pagination，通常強在哪裡？"
question_en: "What is cursor pagination usually better at than offset pagination?"
choices_zh:
  - "讓查詢語法看起來更短"
  - "在大量資料和持續新增資料時更穩定"
  - "自動幫你建立索引"
choices_en:
  - "Making the SQL look shorter"
  - "Staying more stable with large datasets and frequent inserts"
  - "Automatically creating indexes"
answer: 1
clickbait_zh: "翻頁不是只有 page=2、page=3。資料一多，舊方法很容易開始飄。"
clickbait_en: "Pagination is not just page 2 and page 3. Once data grows, the old approach starts drifting."
review_hint_zh: "cursor pagination 在大量資料和動態資料集時通常更穩定。"
review_hint_en: "Cursor pagination is usually more stable for large and changing datasets."
enabled: true
---

## zh-TW

cursor pagination 會用上一筆資料的排序鍵當游標，而不是單純跳過前面 N 筆。
這讓它在資料量大或資料持續新增時，通常比 offset pagination 更穩、更有效率。

## en

Cursor pagination uses the sort key of the last seen row as the next cursor instead of simply skipping N rows.
That usually makes it more stable and efficient than offset pagination when datasets are large or constantly changing.
