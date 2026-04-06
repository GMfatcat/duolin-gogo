---
id: backend-soft-delete-hide-not-drop
title_zh: Soft Delete
title_en: Soft Delete
type: true-false
body_format: bilingual-section
tags: [backend, crud, delete]
difficulty: 2
question_zh: "soft delete 常指資料沒有真的從資料庫移除，而是被標記成已刪除。"
question_en: "Soft delete commonly means the data is not physically removed from the database but marked as deleted."
answer: true
clickbait_zh: "你以為刪掉了，其實它只是被藏起來，還在資料庫角落看著你。"
clickbait_en: "You thought it was deleted, but it may only be hidden and still sitting in the database."
review_hint_zh: "soft delete 偏標記刪除，不是物理移除。"
review_hint_en: "Soft delete is usually marking, not physical removal."
enabled: true
---

## zh-TW

soft delete 常透過欄位例如 `deleted_at` 或 `is_deleted` 來表示資料已不應被一般流程看到。
它常用在需要恢復資料、保留審計資訊，或避免直接破壞關聯資料的情境。

## en

Soft delete often uses a field such as `deleted_at` or `is_deleted` to mark data as no longer visible to normal flows.
It is common when you need recovery, audit history, or safer handling of related records.
